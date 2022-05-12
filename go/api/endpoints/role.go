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

//ResourceHandlers
func RoleHandlers(router *mux.Router) error {
	router.HandleFunc("/create", TempRoleHandler).Methods("POST")
	router.HandleFunc("/information", TempRoleHandler).Methods("POST")
	router.HandleFunc("/users", TempRoleHandler).Methods("POST")
	router.HandleFunc("/remove", TempRoleHandler).Methods("POST")
	router.HandleFunc("/update", TempRoleHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

// DeleteBookingHandler rewrites fields for booking where applicable
func TempRoleHandler(writer http.ResponseWriter, request *http.Request){
	utils.Ok(writer, request)
}
