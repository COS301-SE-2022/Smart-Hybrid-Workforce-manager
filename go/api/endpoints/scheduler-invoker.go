package endpoints

import (
	"api/data"
	"api/db"
	"api/scheduler"
	"lib/logger"
	"lib/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//BookingHandlers handles booking requests
func SchedulerHandlers(router *mux.Router) error {
	router.HandleFunc("/execute", SchedulerInvoker).Methods("POST")
	router.HandleFunc("/delete", RemoveAutomatedBookings).Methods("POST")
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

func RemoveAutomatedBookings(writer http.ResponseWriter, request *http.Request) {
	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	var booking data.Booking
	yes := true
	booking.Automated = &yes
	da := data.NewBookingDA(access)
	bookings, err := da.FindIdentifier(&booking, &data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")})
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	for _, abooking := range bookings {
		_, err = da.DeleteIdentifier(abooking)
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Println("All automatic booking deleted") // TODO [KP]: Be more descriptive
	utils.Ok(writer, request)
}
