package endpoints

import (
	"api/data"
	"api/db"
	"api/security"
	"api/utils"
	"fmt"
	"lib/logger"
	"lib/testutils"
	"net/http"

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

	// TODO [KP]: Do more checks like if they already have a booking etc

	da := data.NewBookingDA(access)
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
	logger.Access.Printf("%v created\n", booking.Id) // TODO [KP]: Be more descriptive
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

// TODO: @JonathanEnslin create bookings for role and team members
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
	// TODO: @JonathanEnslin fix these permissions, add perms for creating bookings for certain teams, roles etc... Or leave it up to the scheduler
	authorized := false
	if meetingRoomBooking.Booking.UserId != nil {
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == meetingRoomBooking.Booking.UserId || permission.PermissionTenantId == nil {
				authorized = authorized && true
			}
		}
	}
	if meetingRoomBooking.RoleId != nil {
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == meetingRoomBooking.RoleId || permission.PermissionTenantId == nil {
				authorized = authorized && true
			}
		}
	}
	if meetingRoomBooking.TeamId != nil {
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == meetingRoomBooking.TeamId || permission.PermissionTenantId == nil {
				authorized = authorized && true
			}
		}
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
	// TODO: @JonathanEnslin check if booking exists first
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
	fmt.Println(testutils.Scolour(testutils.RED, fmt.Sprint(meetingRoomBooking.DesksAttendees)))
	if meetingRoomBooking.DesksAttendees != nil && *meetingRoomBooking.DesksAttendees == true {
		fmt.Println(testutils.Scolour(testutils.RED, fmt.Sprint(*meetingRoomBooking.DesksAttendees)))
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
		deskResourceType := "DESK"
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
	}
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v meeting room booking created\n", bookingId)
}

// TODO: @JonathanEnslin determine if unnecessary because db has on delete cascade
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
