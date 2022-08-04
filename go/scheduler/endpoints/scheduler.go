package endpoints

import (
	"lib/utils"
	"net/http"

	"github.com/gorilla/mux"
)

/////////////////////////////////////////////
// Endpoints

//SchedulerHandlers maintains scheduler endpoints
func SchedulerHandlers(router *mux.Router) error {
	router.HandleFunc("/test", TEST).Methods("POST") // TODO [KP]: REMOVE THIS
	return nil
}

func TEST(writer http.ResponseWriter, request *http.Request) {
	utils.Ok(writer, request)
}
