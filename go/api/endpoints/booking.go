package endpoints

import (
	"api/data"
	"api/db"
	"api/google_api"
	"api/security"
	"fmt"
	"lib/collectionutils"
	"lib/logger"
	"lib/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/////////////////////////////////////////////
// Endpoints

//BookingHandlers handles booking requests
func BookingHandlers(router *mux.Router) error {
	router.HandleFunc("/create", security.Validate(CreateBookingHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "BOOKING", "USER")})).Methods("POST")

	router.HandleFunc("/information", security.Validate(InformationBookingHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")})).Methods("POST")

	router.HandleFunc("/remove", security.Validate(DeleteBookingHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "BOOKING", "USER")})).Methods("POST")

	router.HandleFunc("/meetingroom/create", security.Validate(CreateMeetingRoomBookingHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "BOOKING", "USER"),
			data.CreateGenericPermission("CREATE", "BOOKING", "TEAM"),
			data.CreateGenericPermission("CREATE", "BOOKING", "ROLE")})).Methods("POST")

	router.HandleFunc("/meetingroom/information", security.Validate(InformationMeetingRoomBookingHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")})).Methods("POST")

	return nil
}

/////////////////////////////////////////////
// Functions

// CreateBookingHandler creates or updates a booking
func CreateBookingHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Booking
	var booking data.Booking
	err := utils.UnmarshalJSON(writer, request, &booking)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}
	// Set booked to nil
	booking.Booked = nil

	// Check if user has permission to create a booking for the incomming user
	authorized := false
	if booking.UserId != nil {
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == booking.UserId || permission.PermissionTenantId == nil {
				authorized = true
			}
		}
	}
	if !authorized {
		utils.AccessDenied(writer, request, fmt.Errorf("doesn't have permission to execute query")) // TODO [KP]: Be more descriptive
		return
	}

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// Make access connection to db
	da := data.NewBookingDA(access)
	dr := data.NewResourceDA(access)

	// Check if they already have a desk booking
	floor := time.Date(booking.Start.Year(), booking.Start.Month(), booking.Start.Day(), 0, 0, 0, 0, booking.Start.Location())
	ceil := time.Date(booking.Start.Year(), booking.Start.Month(), booking.Start.Day(), 23, 59, 59, 0, booking.Start.Location())
	desk := "DESK"

	presentBooking := &data.Booking{
		UserId:       booking.UserId,
		Start:        &floor,
		End:          &ceil,
		ResourceType: &desk,
	}
	bookings, err := da.FindIdentifier(presentBooking, &data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")})
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	if len(bookings) > 0 {
		for _, existingBooking := range bookings {
			if (booking.Start.Before(*existingBooking.End) && booking.Start.After(*existingBooking.Start)) || (booking.End.After(*existingBooking.Start) && booking.End.Before(*existingBooking.End)) || (booking.Start.Equal(*existingBooking.Start) && booking.End.Equal(*existingBooking.End)) {
				utils.BadRequest(writer, request, "booking_exists")
				return
			}
		}
	}

	// On Demand scheduler
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	daysInAdvance := 2
	onDemandCutoffDate := now.AddDate(0, 0, daysInAdvance)

	// Assign Resource on Demand
	if (*booking.Start).Before(onDemandCutoffDate) {
		// Get all resources booked on the day
		day := &data.Booking{
			Start: &floor,
			End:   &ceil,
		}
		bookings, err := da.FindIdentifier(day, &data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")})
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
		// Create a list of Resources booked
		var resourcesBooked []string
		for _, booking := range bookings {
			if booking.ResourceId != nil {
				resourcesBooked = append(resourcesBooked, *booking.ResourceId)
			}
		}
		// Get all resources available
		resourceType := &data.Resource{
			ResourceType: &desk,
		}
		resources, err := dr.FindIdentifier(resourceType, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "IDENTIFIER")})
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
		// Find first available resource and bookit
		yes := true
		for _, resource := range resources {
			if resource.Id != nil && !collectionutils.Contains(resourcesBooked, *resource.Id) {
				booking.ResourceId = resource.Id
				booking.Booked = &yes
				break
			}
		}
		// If no available resources error
		if !(*booking.Booked) {
			utils.BadRequest(writer, request, "no_available_resources")
			return
		}
	}

	// Create Booking
	_, err = da.StoreIdentifier(&booking)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Create Google Calendar Event
	du := data.NewUserDA(access)
	users, err := du.FindIdentifier(&data.User{Id: booking.UserId})
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	user := users.FindHead()
	logger.Access.Println("\nHERE1")
	logger.Access.Printf("\nData5: %v", booking)
	logger.Access.Printf("\nData5: %v", booking.Id)
	err = google_api.CreateBooking(user, &booking)
	if err != nil {
		logger.Error.Printf("Error occured while creating google calendar event: %v", err)
	}
	logger.Access.Println("\nWE GOT TILL HERE BEFORE IT ... UP")

	logger.Access.Println("\nHERE2")
	logger.Access.Printf("%v created\n", booking.Id) // TODO [KP]: Be more descriptive
	logger.Access.Println("\nHERE3")
	utils.Ok(writer, request)
}

// InformationBookingHandler gets bookings
func InformationBookingHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Booking
	var booking data.Booking
	err := utils.UnmarshalJSON(writer, request, &booking)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// No check for permissions the database handles information permissions

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// TODO [KP]: null checks etc.

	da := data.NewBookingDA(access)
	bookings, err := da.FindIdentifier(&booking, security.RemoveRolePermissions(permissions))
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v booking information requested\n", booking.Id) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, bookings)
}

// DeleteBookingHandler removes a booking
func DeleteBookingHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Booking
	var booking data.Booking
	err := utils.UnmarshalJSON(writer, request, &booking)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// Get booking information if no user is defined
	da := data.NewBookingDA(access)
	if booking.UserId == nil {
		// TODO: fix call when non existent booking is deleted
		temp, err := da.FindIdentifier(&booking, &data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")})
		booking = *temp.FindHead()
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
	}

	// Check if user has permission to delete a booking for the incomming booking user
	if booking.UserId != nil {
		authorized := false
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == booking.UserId || permission.PermissionTenantId == nil {
				authorized = true
			}
		}
		if !authorized {
			utils.AccessDenied(writer, request, fmt.Errorf("doesn't have permission to execute query")) // TODO [KP]: Be more descriptive
			return
		}
	}

	bookingRemoved, err := da.DeleteIdentifier(&booking)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v booking removed\n", booking.Id) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, bookingRemoved)
}

// CreateMeetingRoomBookingHandler creates or updates a booking
func CreateMeetingRoomBookingHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshall MeetingRoomBooking
	var meetingRoomBooking data.MeetingRoomBooking
	err := utils.UnmarshalJSON(writer, request, &meetingRoomBooking)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Check if the user has permission to create or update a booking for the incoming meeting room
	authorized := false
	if meetingRoomBooking.Booking.UserId != nil {
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == meetingRoomBooking.Booking.UserId || permission.PermissionTenantId == nil {
				authorized = true
			}
		}
	}
	if !authorized {
		utils.AccessDenied(writer, request, fmt.Errorf("doesn't have permission to execute query"))
		return
	}

	authorized = false
	if meetingRoomBooking.RoleId != nil {
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == meetingRoomBooking.RoleId || permission.PermissionTenantId == nil {
				authorized = true
			}
		}
	} else {
		authorized = true
	}
	if !authorized {
		utils.AccessDenied(writer, request, fmt.Errorf("doesn't have permission to execute query"))
		return
	}

	authorized = false
	if meetingRoomBooking.TeamId != nil {
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == meetingRoomBooking.TeamId || permission.PermissionTenantId == nil {
				authorized = true
			}
		}
	} else {
		authorized = true
	}
	if !authorized {
		utils.AccessDenied(writer, request, fmt.Errorf("doesn't have permission to execute query"))
		return
	}

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewBookingDA(access)

	var bookingId *string
	if meetingRoomBooking.Booking != nil {
		bookingId, err = da.StoreIdentifier(meetingRoomBooking.Booking)
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
	} else {
		if meetingRoomBooking.BookingId == nil {
			utils.BadRequest(writer, request, "no booking_id passed")
			return
		}
		bookingId = meetingRoomBooking.BookingId
	}

	meetingRoomBooking.BookingId = bookingId
	err = da.StoreBookingMeetingRoom(&meetingRoomBooking)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	// // Add desks for role/team members in meeting
	teamUsers := data.UserTeams{}
	roleUsers := data.UserRoles{}

	if meetingRoomBooking.TeamId != nil {
		teamUsers, err = GetUserTeams(meetingRoomBooking.TeamId)
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
	}
	if meetingRoomBooking.RoleId != nil {
		roleUsers, err = GetUserRoles(meetingRoomBooking.RoleId)
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
	}

	roleTeamUsersMap := make(map[string]bool)
	for _, user := range teamUsers {
		if user != nil {
			roleTeamUsersMap[*user.UserId] = true
		}
	}
	for _, user := range roleUsers {
		if user != nil {
			roleTeamUsersMap[*user.UserId] = true
		}
	}

	bookings := data.Bookings{}
	// Make the actual bookings
	deskResourceType := "MEETINGROOM"
	for key, _ := range roleTeamUsersMap {
		bookings = append(bookings, &data.Booking{
			UserId:       &key,
			ResourceType: &deskResourceType,
			Start:        meetingRoomBooking.Booking.Start,
			End:          meetingRoomBooking.Booking.End,
			Dependent:    bookingId,
		})
	}

	// Store the bookings
	batchBooking := data.BatchBooking{
		UserId:   nil,
		Bookings: bookings,
	}

	batchBookingDa := data.NewBatchBookingDA(access)
	err = batchBookingDa.StoreIdentifiers(&batchBooking)
	if err != nil {
		utils.InternalServerError(writer, request, err)
	}

	bookings = data.Bookings{}

	// Add desks bookings
	if meetingRoomBooking.DesksAttendees != nil && *meetingRoomBooking.DesksAttendees == true {
		// // Add desks for role/team members in meeting
		// Make the actual bookings
		deskResourceType = "DESK"
		for key, _ := range roleTeamUsersMap {
			bookings = append(bookings, &data.Booking{
				UserId:       &key,
				ResourceType: &deskResourceType,
				Start:        meetingRoomBooking.Booking.Start,
				End:          meetingRoomBooking.Booking.End,
				Dependent:    bookingId,
			})
		}

		// Store the bookings
		batchBooking = data.BatchBooking{
			UserId:   nil,
			Bookings: bookings,
		}

		err = batchBookingDa.StoreIdentifiers(&batchBooking)
		if err != nil {
			utils.InternalServerError(writer, request, err)
		}
	}
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v meeting room booking created\n", bookingId)
}

// DeleteMeetingRoomBookingHandler deletes a meeting room booking
// func DeleteMeetingRoomBookingHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
// 	// Unmarshall MeetingRoomBooking
// 	var meetingRoomBooking data.MeetingRoomBooking
// 	err := utils.UnmarshalJSON(writer, request, &meetingRoomBooking)
// 	if err != nil {
// 		utils.BadRequest(writer, request, "invalid_request")
// 		return
// 	}

// 	// Create a database connection
// 	access, err := db.Open()
// 	if err != nil {
// 		utils.InternalServerError(writer, request, err)
// 		return
// 	}
// 	defer access.Close()

// 	// Check if user has permission to delete a booking for the incoming meeting room booking

// }

func InformationMeetingRoomBookingHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal MeetingRoomBooking
	var meetingRoomBooking data.MeetingRoomBooking
	err := utils.UnmarshalJSON(writer, request, &meetingRoomBooking)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// No check for permissions the database handles information permissions

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// TODO: null checks etc.

	da := data.NewBookingDA(access)
	bookings, err := da.FindMeeetingRoomBooking(&meetingRoomBooking, security.RemoveRolePermissions(permissions))
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v meetingroom booking information requested\n", meetingRoomBooking.Id) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, bookings)
}

// GetUserTeams will return all the userteam pairs
func GetUserTeams(teamId *string) (data.UserTeams, error) {
	// Create a database connection
	access, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer access.Close()
	da := data.NewTeamDA(access)

	userTeam := data.UserTeam{
		TeamId: teamId,
	}
	permissions := &data.Permissions{data.CreateGenericPermission("VIEW", "TEAM", "USER"),
		data.CreateGenericPermission("VIEW", "USER", "TEAM")}
	userTeams, err := da.FindUserTeam(&userTeam, permissions)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		return nil, err
	}
	return userTeams, nil
}

// GetUserRoles will return all the userteam pairs
func GetUserRoles(roleId *string) (data.UserRoles, error) {
	// Create a database connection
	access, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer access.Close()
	da := data.NewRoleDA(access)

	userRole := data.UserRole{
		RoleId: roleId,
	}
	permissions := &data.Permissions{data.CreateGenericPermission("VIEW", "ROLE", "USER"),
		data.CreateGenericPermission("VIEW", "USER", "ROLE")}
	userRoles, err := da.FindUserRole(&userRole, permissions)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}
