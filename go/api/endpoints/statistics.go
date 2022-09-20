package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
)

//StatisticsHandlers registers the user
func StatisticsHandlers(router *mux.Router) error {
	router.HandleFunc("/resource_utilisation", ResourceUtilisation).Methods("POST")
	return nil
}

func ResourceUtilisation(writer http.ResponseWriter, request *http.Request) {

}
