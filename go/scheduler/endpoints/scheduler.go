package endpoints

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"lib/logger"
	"lib/utils"
	"net/http"
	"os"
	"path/filepath"
	"scheduler/data"
	meetingroom "scheduler/meeting_room"
	"scheduler/overseer"

	"github.com/gorilla/mux"
)

var weekdays = []string{"Mon", "Tue", "Wed"}

/////////////////////////////////////////////
// Endpoints

//SchedulerHandlers maintains scheduler endpoints
func SchedulerHandlers(router *mux.Router) error {
	router.HandleFunc("/test", TEST).Methods("POST") // TODO [KP]: REMOVE THIS
	router.HandleFunc("/weekly", weeklyScheduler).Methods("POST")
	router.HandleFunc("/daily", dailyScheduler).Methods("POST")
	router.HandleFunc("/meeting_room", meetingRoomScheduler).Methods("POST")

	return nil
}

func TEST(writer http.ResponseWriter, request *http.Request) {
	utils.Ok(writer, request)
}

func weeklyScheduler(writer http.ResponseWriter, request *http.Request) {
	config, _ := parseConfig("/run/secrets/config.json")
	var schedulerData data.SchedulerData

	err := utils.UnmarshalJSON(writer, request, &schedulerData)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}
	schedulerData.ApplyMapping()
	bookings := overseer.WeeklyOverseer(schedulerData, config)

	utils.JSONResponse(writer, request, bookings)
}

func meetingRoomScheduler(writer http.ResponseWriter, request *http.Request) {
	// config, _ := parseConfig("/run/secrets/config.json")
	var schedulerData data.SchedulerData

	err := utils.UnmarshalJSON(writer, request, &schedulerData)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}
	schedulerData.ApplyMapping()
	bookings := meetingroom.AssignMeetingRoomsToBookings(&schedulerData)
	utils.JSONResponse(writer, request, bookings)
}

func dailyScheduler(writer http.ResponseWriter, request *http.Request) {
	config, _ := parseConfig("/run/secrets/config.json")
	var schedulerData data.SchedulerData

	err := utils.UnmarshalJSON(writer, request, &schedulerData)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}
	schedulerData.ApplyMapping()
	bookings := overseer.DailyOverseer(schedulerData, config)

	utils.JSONResponse(writer, request, bookings)
	logger.Debug.Println("AAAAAMMMM IIIII HEEEEERRRREEEEEEE")
}

// Loads the schedulers config file uwu
func parseConfig(filePath string) (*data.SchedulerConfig, error) {
	if _, err := os.Stat(filepath.Join("", filepath.Clean(filePath))); errors.Is(err, os.ErrNotExist) {
		logger.Info.Println("Could not find scheduler config file")
		return nil, nil
	} else {
		logger.Info.Println("Loading scheduler config")
		configJson, err := os.Open(filepath.Join("", filepath.Clean(filePath)))
		if err != nil {
			logger.Error.Println("Could not open scheduler config file")
			return nil, err
		}
		fileBytes, err := ioutil.ReadAll(configJson)
		closeErr := configJson.Close()
		if closeErr != nil {
			logger.Error.Println("Could not close config file")
		}
		if err != nil {
			logger.Error.Println("Could not read scheduler config file, err: ", err)
			return nil, err
		}
		var schedulerConfig data.SchedulerConfig
		err = json.Unmarshal(fileBytes, &schedulerConfig)
		if err != nil {
			logger.Error.Println("Could not parse JSON")
			return nil, err
		}
		return &schedulerConfig, nil
	}
}
