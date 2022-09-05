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
	router.HandleFunc("/execute/weekly", WeeklyScheduler).Methods("POST")
	router.HandleFunc("/execute/daily", DailyScheduler).Methods("POST")
	router.HandleFunc("/delete", RemoveAutomatedBookings).Methods("POST")

	return nil
}

// SchedulerInvoker will invoke the weekly scheduler and then the daily schedulers for each day of the week
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

// WeeklyScheduler will call and execute the weekly scheduers
func WeeklyScheduler(writer http.ResponseWriter, request *http.Request) {
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

// DailyScheduler will call and execute the daily scheduers
func DailyScheduler(writer http.ResponseWriter, request *http.Request) {
	utils.Ok(writer, request)
}

// RemoveAutomatedBookings removes all automated bookings from the database
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
