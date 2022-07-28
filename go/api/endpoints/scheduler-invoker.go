package endpoints

import (
	"api/scheduler"
	"api/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//BookingHandlers handles booking requests
func SchedulerHandlers(router *mux.Router) error {
	router.HandleFunc("/execute", SchedulerInvoker).Methods("POST")
	// router.HandleFunc("/scheduler/execute", security.Validate(InformationMeetingRoomBookingHandler,
	// 	&data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")})).Methods("POST")

	return nil
}

func SchedulerInvoker(writer http.ResponseWriter, request *http.Request) {
	now := time.Now()
	nextMonday := scheduler.TimeOfNextWeekDay(now, "Monday")            // Start of next week
	nextSaturday := scheduler.TimeOfNextWeekDay(nextMonday, "Saturday") // End of next work-week
	schedulerData, err := scheduler.GetSchedulerData(nextMonday, nextSaturday)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	err = scheduler.Call(schedulerData) // TODO: @JonathanEnslin handle the return data
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	utils.Ok(writer, request)
}
