package endpoints

import (
	"api/data"
	"api/db"
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
	router.HandleFunc("/create", CreateBookingHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

// CreateBookingHandler registers a new user
func CreateBookingHandler(writer http.ResponseWriter, request *http.Request) {
	var booking data.Booking

	err := utils.UnmarshalJSON(writer, request, &booking)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewBookingDA(access)

	// TODO [KP]: Do more checks like if they already have a booking etc

	err = da.StoreIdentifier(&booking)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v created\n", booking.Id)

	utils.Ok(writer, request)
}
