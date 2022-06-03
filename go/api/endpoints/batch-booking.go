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

// func mockHandler(writer http.ResponseWriter, request *http.Request) {
// 	CreateBatchBookingHandler(writer, request, nil)
// }

func BatchBookingHandlers(router *mux.Router) error {
	router.HandleFunc("/create", security.Validate(CreateBatchBookingHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "BOOKING", "USER")})).Methods("POST")
	router.HandleFunc("/information", security.Validate(InformationBatchBookingHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")})).Methods("POST")
	router.HandleFunc("/remove", security.Validate(DeleteBatchBookingsHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "BOOKING", "USER")})).Methods("POST")
	return nil
}

////////////////////////////////////////
// Functions

// Creates (or updates) a batch of bookings
func CreateBatchBookingHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarhsal bookings
	var bookings data.BatchBooking
	err := utils.UnmarshalJSON(writer, request, &bookings)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	hasNilTenant := false                  // checks if user has permissions over all
	permTenantIDs := make(map[string]bool) // foor lookup of permissionTenantIDs
	for _, permission := range *permissions {
		if permission.PermissionTenantId == nil {
			hasNilTenant = true // Permission applies to all, no need to check further
			break
		}
		permTenantIDs[*permission.PermissionTenantId] = true // true is just a stub
	}

	// check permissions
	authFailed := false
	if bookings.UserId == nil { // request must have an associated ID
		authFailed = true
	} else if !hasNilTenant { // if the tenantId was nil, permission applies to all, no need to do furhter checks
		for _, booking := range bookings.Bookings {
			if booking.UserId == nil { // all bookings must have an associated UserId
				authFailed = true
				break
			}
			if _, contained := permTenantIDs[*booking.UserId]; !contained { // if the user making request does not have perms to create booking for this user
				authFailed = true // auth has failed
				break
			}
		}
	}

	if authFailed {
		utils.AccessDenied(writer, request, fmt.Errorf("does not have permissions to create one or more of the specified booking"))
		return
	}

	// Connect to db
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// TODO @JonathanEnslin do other necessary checks

	da := data.NewBatchBookingDA(access)
	err = da.StoreIdentifiers(&bookings)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Log all created bookings
	for _, booking := range bookings.Bookings {
		logger.Access.Printf("%v created\n", booking.Id)
	}
	utils.Ok(writer, request)
}

func InformationBatchBookingHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarhsal bookings
	var bookings data.BatchBooking
	err := utils.UnmarshalJSON(writer, request, &bookings)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Connect to db
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// TODO @JonathanEnslin do other necessary checks

	da := data.NewBatchBookingDA(access)
	bookingsInfo, err := da.FindIdentifiers(&bookings, security.RemoveRolePermissions(permissions))
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	for _, booking := range bookings.Bookings {
		logger.Access.Printf("%v, batch booking information requested\n", booking.Id)
	}
	utils.JSONResponse(writer, request, bookingsInfo)
}

func DeleteBatchBookingsHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	var bookings data.BatchBooking
	err := utils.UnmarshalJSON(writer, request, &bookings)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Connect to db
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	hasNilTenant := false                  // checks if user has permissions over all
	permTenantIDs := make(map[string]bool) // foor lookup of permissionTenantIDs
	for _, permission := range *permissions {
		if permission.PermissionTenantId == nil {
			hasNilTenant = true // Permission applies to all, no need to check further
			break
		}
		permTenantIDs[*permission.PermissionTenantId] = true // true is just a stub
	}

	da := data.NewBatchBookingDA(access)
	// Get all full bookings (meaning all info included), use a generic booking since nil tenant_id means all data can be requested
	fullBookings, err := da.FindIdentifiers(&bookings, &data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")})
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	authFailed := false
	if bookings.UserId == nil {
		authFailed = true
	} else if !hasNilTenant { // no need to check if nil tenenant_id is presenting since user has perms over all
		// Check if user that made batch delete request has permissions to delete the bookings
		for _, booking := range fullBookings {
			if booking.UserId == nil {
				authFailed = true
				break
			}
			if _, ok := permTenantIDs[*booking.UserId]; !ok {
				// if there is no permission tenant_id for this user, then auth has failed
				authFailed = true
				break
			}
		}
	}

	if authFailed {
		utils.AccessDenied(writer, request, fmt.Errorf("do not have permissions to delete one or more of the specified bookings"))
		return
	}

	// Delete bookings
	deleted, err := da.DeleteIdentifiers(&bookings)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	for _, deletedBooking := range deleted {
		logger.Access.Printf("%v booking removed\n", deletedBooking.Id)
	}
	utils.JSONResponse(writer, request, deleted)
}
