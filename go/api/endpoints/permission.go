package endpoints

import (
	// "api/data"
	// "api/db"
	"api/utils"
	// "fmt"
	// "lib/logger"
	"net/http"
	// "regexp"

	"github.com/gorilla/mux"
)

//////////////////////////////////////////////////
// Structures and Variables


/////////////////////////////////////////////
// Endpoints

//PermissionHandlers
func PermissionHandlers(router *mux.Router) error {
	router.HandleFunc("/create", TempHandler).Methods("POST")
	router.HandleFunc("/allocate", TempHandler).Methods("POST")
	router.HandleFunc("/remove", TempHandler).Methods("POST")
	router.HandleFunc("/list", TempHandler).Methods("POST")
	router.HandleFunc("/update", TempHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

// DeleteBookingHandler rewrites fields for booking where applicable
func TempHandler(writer http.ResponseWriter, request *http.Request){
	utils.Ok(writer, request)
}