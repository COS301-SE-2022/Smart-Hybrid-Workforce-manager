package endpoints

import (
	"api/data"
	"api/db"
	"api/security"
	"api/utils"
	"fmt"
	"lib/logger"
	"net/http"

	"github.com/gorilla/mux"
)

/////////////////////////////////////////////
// Endpoints

//BookingHandlers handles booking requests
func BookingHandlers(router *mux.Router) error {
	router.HandleFunc("/create", security.Validate(CreateBookingHandler,
		data.CreateGenericPermission("CREATE", "BOOKING", "USER"))).Methods("POST")

	router.HandleFunc("/information", security.Validate(InformationBookingHandler,
		data.CreateGenericPermission("VIEW", "BOOKING", "USER"))).Methods("POST")

	router.HandleFunc("/remove", security.Validate(DeleteBookingHandler,
		data.CreateGenericPermission("DELETE", "BOOKING", "USER"))).Methods("POST")
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
	err = da.StoreIdentifier(&booking)
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
