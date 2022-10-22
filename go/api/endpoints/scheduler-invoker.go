package endpoints

import (
	"api/data"
	"api/db"
	"api/scheduler"
	"fmt"
	"lib/logger"
	"lib/testutils"
	"lib/utils"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type SchedulerRequest struct {
	StartDate *time.Time                 `json:"start_date,omitempty"`
	NumDays   *int                       `json:"num_days,omitempty"` // Used for weekly scheduler, not necessarily daily scheduler
	Config    *scheduler.SchedulerConfig `json:"config,omitempty"`
}

//BookingHandlers handles booking requests
func SchedulerHandlers(router *mux.Router) error {
	router.HandleFunc("/execute", SchedulerInvoker).Methods("POST")
	router.HandleFunc("/execute/weekly", WeeklyScheduler).Methods("POST")
	router.HandleFunc("/execute/daily", DailyScheduler).Methods("POST")
	router.HandleFunc("/execute/meeting_room", MeetingroomScheduler).Methods("POST")
	router.HandleFunc("/delete", RemoveAutomatedBookings).Methods("POST")

	return nil
}

// SchedulerInvoker will invoke the weekly scheduler and then the daily schedulers for each day of the week
func SchedulerInvoker(writer http.ResponseWriter, request *http.Request) {
	weeklyEndpointURL := os.Getenv("SCHEDULER_ADDR") + "/weekly"
	dailyEndpointURL := os.Getenv("SCHEDULER_ADDR") + "/daily"
	var requestedData SchedulerRequest
	err := utils.UnmarshalJSON(writer, request, &requestedData)
	if err != nil {
		utils.BadRequest(writer, request, fmt.Sprintf("invalid_request: %v", err))
		return
	}
	now := time.Now()
	if requestedData.StartDate != nil { // Use passed in date if a date was supplied
		now = *requestedData.StartDate
		now = now.AddDate(0, 0, -7)
		// Set start to previous day, so that scheduler is called for the requested day
	}
	nextMonday := scheduler.TimeOfNextWeekDay(now, "Monday")            // Start of next week
	nextSaturday := scheduler.TimeOfNextWeekDay(nextMonday, "Saturday") // End of next work-week
	deskType := "DESK"
	schedulerData, err := scheduler.GetSchedulerData(nextMonday, nextSaturday, &deskType, &deskType)
	schedulerData.Config = requestedData.Config
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	buildingGroups := scheduler.GroupByBuilding(schedulerData)
	weeklyBuildingWg := sync.WaitGroup{}
	for _, data := range buildingGroups {
		weeklyBuildingWg.Add(1)
		go func(schedulerData *scheduler.SchedulerData) {
			defer weeklyBuildingWg.Done()
			logger.Debug.Println(testutils.Scolourf(testutils.GREEN, "Running weekly scheduler from %v -> %v for building: %v", nextMonday, nextSaturday, *schedulerData.Buildings[0].Id))
			if err != nil {
				utils.InternalServerError(writer, request, err)
				return
			}
			err = scheduler.Call(schedulerData, weeklyEndpointURL)
			if err != nil {
				utils.InternalServerError(writer, request, err)
				return
			}
		}(data)
	}
	weeklyBuildingWg.Wait()
	// Call daily scheduler 5 times
	now = nextMonday
	yyyy, mm, dd := now.Date()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(_now time.Time, _i int) {
			defer wg.Done()
			startDate := time.Date(yyyy, mm, dd+_i, 0, 0, 0, 0, now.Location())
			endDate := startDate.AddDate(0, 0, 1) // Add one day
			schedulerData, err := scheduler.GetSchedulerData(startDate, endDate, &deskType, &deskType)
			schedulerData.Config = requestedData.Config
			buildingGroups := scheduler.GroupByBuilding(schedulerData)
			dailyBuildingWg := sync.WaitGroup{}
			for _, data := range buildingGroups {
				dailyBuildingWg.Add(1)
				go func(schedulerData *scheduler.SchedulerData) {
					defer dailyBuildingWg.Done()
					logger.Debug.Println(testutils.Scolourf(testutils.GREEN, "Running daily scheduler from %v -> %v for building: %v", startDate, endDate, *schedulerData.Buildings[0].Id))
					if err != nil {
						utils.InternalServerError(writer, request, err)
						return
					}
					err = scheduler.Call(schedulerData, dailyEndpointURL)
					if err != nil {
						utils.InternalServerError(writer, request, err)
						return
					}
				}(data)
			}
			dailyBuildingWg.Wait()
		}(now, i)
	}
	wg.Wait()
	utils.Ok(writer, request)
}

func MeetingroomScheduler(writer http.ResponseWriter, request *http.Request) {
	var requestedDate SchedulerRequest
	err := utils.UnmarshalJSON(writer, request, &requestedDate)
	if err != nil {
		utils.BadRequest(writer, request, fmt.Sprintf("invalid_request: %v", err))
		return
	}
	now := time.Now()
	if requestedDate.StartDate != nil { // Use passed in date if a date was supplied
		now = *requestedDate.StartDate
	}
	err = scheduler.CallMeetingRoomScheduler(0, now)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
}

// WeeklyScheduler will call and execute the weekly scheduers
func WeeklyScheduler(writer http.ResponseWriter, request *http.Request) {
	weeklyEndpointURL := os.Getenv("SCHEDULER_ADDR") + "/weekly"

	var requestedData SchedulerRequest
	err := utils.UnmarshalJSON(writer, request, &requestedData)
	if err != nil {
		utils.BadRequest(writer, request, fmt.Sprintf("invalid_request: %v", err))
		return
	}
	now := time.Now()
	if requestedData.StartDate != nil { // Use passed in date if a date was supplied
		now = *requestedData.StartDate
		now = now.AddDate(0, 0, -7)
		// Set start to previous day, so that scheduler is called for the requested day
	}
	deskType := "DESK"
	nextMonday := scheduler.TimeOfNextWeekDay(now, "Monday")            // Start of next week
	nextSaturday := scheduler.TimeOfNextWeekDay(nextMonday, "Saturday") // End of next work-week
	schedulerData, err := scheduler.GetSchedulerData(nextMonday, nextSaturday, &deskType, &deskType)
	schedulerData.Config = requestedData.Config
	buildingGroups := scheduler.GroupByBuilding(schedulerData)
	weeklyBuildingWg := sync.WaitGroup{}
	for _, data := range buildingGroups {
		weeklyBuildingWg.Add(1)
		go func(schedulerData *scheduler.SchedulerData) {
			defer weeklyBuildingWg.Done()
			logger.Debug.Println(testutils.Scolourf(testutils.GREEN, "Running weekly scheduler from %v -> %v for building: %v", nextMonday, nextSaturday, *schedulerData.Buildings[0].Id))
			if err != nil {
				utils.InternalServerError(writer, request, err)
				return
			}
			err = scheduler.Call(schedulerData, weeklyEndpointURL)
			if err != nil {
				utils.InternalServerError(writer, request, err)
				return
			}
		}(data)
	}
	weeklyBuildingWg.Wait()
	utils.Ok(writer, request)
}

// var meetingRoomBooking data.MeetingRoomBooking
// 	err := utils.UnmarshalJSON(writer, request, &meetingRoomBooking)

// DailyScheduler will call and execute the daily scheduers
func DailyScheduler(writer http.ResponseWriter, request *http.Request) {
	dailyEndpointURL := os.Getenv("SCHEDULER_ADDR") + "/daily"

	var requestedData SchedulerRequest
	err := utils.UnmarshalJSON(writer, request, &requestedData)
	if err != nil {
		utils.BadRequest(writer, request, fmt.Sprintf("invalid_request: %v", err))
		return
	}
	now := time.Now()
	if requestedData.StartDate != nil { // Use passed in date if a date was supplied
		now = *requestedData.StartDate
		// Set start to previous day, so that scheduler is called for the requested day
		now = now.AddDate(0, 0, -1)
	}
	deskType := "DESK"
	yyyy, mm, dd := now.Date()
	startDate := time.Date(yyyy, mm, dd+1, 0, 0, 0, 0, now.Location())
	endDate := startDate.AddDate(0, 0, 1) // Add one day
	// Get data between start and end of date
	schedulerData, err := scheduler.GetSchedulerData(startDate, endDate, &deskType, &deskType)
	schedulerData.Config = requestedData.Config
	buildingGroups := scheduler.GroupByBuilding(schedulerData)
	dailyBuildingWg := sync.WaitGroup{}
	for _, data := range buildingGroups {
		dailyBuildingWg.Add(1)
		go func(schedulerData *scheduler.SchedulerData) {
			defer dailyBuildingWg.Done()
			logger.Debug.Println(testutils.Scolourf(testutils.GREEN, "Running daily scheduler from %v -> %v for building: %v", startDate, endDate, *schedulerData.Buildings[0].Id))
			if err != nil {
				utils.InternalServerError(writer, request, err)
				return
			}
			err = scheduler.Call(schedulerData, dailyEndpointURL)
			if err != nil {
				utils.InternalServerError(writer, request, err)
				return
			}
		}(data)
	}
	dailyBuildingWg.Wait()
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
