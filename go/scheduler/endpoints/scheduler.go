package endpoints

import (
	"lib/utils"
	"net/http"
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
	domain.Terminals = data.ExtractUserIdsDuplicates(&schedulerData)
	domain.Config = &config
	domain.SchedulerData = &schedulerData

	results := ga.GA(domain, ga.DayVResourceCrossover, ga.StubFitness, ga.DayVResouceMutate, ga.TournamentSelection, ga.DayVResourcePopulationGenerator)

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

// TODO: FIX THIS FUNCTION SOMEBODY
// func parseConfig(path string) (*data.Config, error) {
// 	filePath := filepath.Clean(path)
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func() {
// 		err := file.Close()
// 		if err != nil {
// Log stuff
// 			panic(err)
// 		}
// 	}()
// 	decoder := json.NewDecoder(file)
// 	Config := data.Config{}
// 	err = decoder.Decode(&Config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Config, nil
// }
