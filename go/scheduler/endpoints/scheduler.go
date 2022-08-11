package endpoints

import (
	"encoding/json"
	"lib/utils"
	"net/http"
	"os"
	"scheduler/data"
	"scheduler/ga"

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

	var config data.Config

	// Perform Magic
	var bookings []data.Bookings

	// Create domain
	var domain ga.Domain
	domain.Terminals = weekdays

	results := ga.GA(schedulerData, config, domain, ga.StubCrossOver, ga.StubFitness, ga.StubMutate, ga.StubSelection, ga.StubPopulationGenerator)

	if len(results) == 0 {

	}
	// parse results as bookings

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

func parseConfig(path string) (*data.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := data.Config{}
	err = decoder.Decode(&Config)
	if err != nil {
		return nil, err
	}
	return &Config, nil
}
