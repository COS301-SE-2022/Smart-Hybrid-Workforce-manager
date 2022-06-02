package endpoints

import (
	"api/data"
	"api/db"
	"api/utils"
	"lib/logger"
	"net/http"

	"github.com/gorilla/mux"
)

func mockHandler(writer http.ResponseWriter, request *http.Request) {
	CreateBatchBookingHandler(writer, request, nil)
}

func BatchBookingHandlers(router *mux.Router) error {
	// TODO @JonathanEnslin Add router handlers
	router.HandleFunc("/create", mockHandler).Methods("POST")
	return nil
}

////////////////////////////////////////
// Functions

// Creates (or updates) a batch of bookings
func CreateBatchBookingHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permission) {
	// Unmarhsal bookings
	var bookings []data.Booking
	err := utils.UnmarshalJSON(writer, request, &bookings)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// check permissions
	// authFailed := false
	// for _, booking := range bookings {
	// 	if authFailed {
	// 		break // If authorisation fails once, abort all
	// 	}
	// }
	// if authFailed {
	// 	utils.AccessDenied(writer, request, fmt.Errorf(""))
	// }

	// Connect to db
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// Query db to create all bookings
	for i := range bookings {
		da := data.NewBookingDA(access)
		err = da.StoreIdentifier(&bookings[i])
		if err != nil { // if error occurs nothing will be comitted
			utils.InternalServerError(writer, request, err)
			return
		}
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Log all created bookings
	for _, booking := range bookings {
		logger.Access.Printf("%v created\n", booking.Id)
	}
	utils.Ok(writer, request)
}
