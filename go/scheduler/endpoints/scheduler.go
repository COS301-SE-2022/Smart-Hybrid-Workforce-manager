package endpoints

import (
	"lib/utils"
	"net/http"
	"scheduler/data"

	"github.com/gorilla/mux"
)

/////////////////////////////////////////////
// Endpoints

//SchedulerHandlers maintains scheduler endpoints
func SchedulerHandlers(router *mux.Router) error {
	router.HandleFunc("/test", TEST).Methods("POST") // TODO [KP]: REMOVE THIS
	router.HandleFunc("/weekly", weeklyScheduler).Methods("POST")
	router.HandleFunc("/daily", dailyScheduler).Methods("POST")

	return nil
}

func TEST(writer http.ResponseWriter, request *http.Request) {
	utils.Ok(writer, request)
}

func weeklyScheduler(writer http.ResponseWriter, request *http.Request) {
	var schedulerData data.SchedulerData

	err := utils.UnmarshalJSON(writer, request, &schedulerData)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Perform Magic

	var bookings []data.Bookings
	utils.JSONResponse(writer, request, bookings)
}

func dailyScheduler(writer http.ResponseWriter, request *http.Request) {
	var schedulerData data.SchedulerData

	err := utils.UnmarshalJSON(writer, request, &schedulerData)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Perform Magic

	var bookings []data.Bookings
	utils.JSONResponse(writer, request, bookings)
}
