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

type CreateResourceStruct struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Capacity    *int    `json:"capacity,omitempty"`
	Picture     *string `json:"picture,omitempty"`
}

/////////////////////////////////////////////
// Endpoints

//ResourceHandlers
func ResourceHandlers(router *mux.Router) error {
	router.HandleFunc("/create", TempResourceHandler).Methods("POST")
	router.HandleFunc("/update", TempResourceHandler).Methods("POST")
	router.HandleFunc("/remove", TempResourceHandler).Methods("POST")
	router.HandleFunc("/list", TempResourceHandler).Methods("POST")
	router.HandleFunc("/room/create", TempResourceHandler).Methods("POST")
	router.HandleFunc("/room/update", TempResourceHandler).Methods("POST")
	router.HandleFunc("/room/remove", TempResourceHandler).Methods("POST")
	router.HandleFunc("/room/list", TempResourceHandler).Methods("POST")
	router.HandleFunc("/building/create", TempResourceHandler).Methods("POST")
	router.HandleFunc("/building/update", TempResourceHandler).Methods("POST")
	router.HandleFunc("/building/remove", TempResourceHandler).Methods("POST")
	router.HandleFunc("/building/list", TempResourceHandler).Methods("POST")
	router.HandleFunc("/workspace/create", TempResourceHandler).Methods("POST")
	router.HandleFunc("/workspace/update", TempResourceHandler).Methods("POST")
	router.HandleFunc("/workspace/remove", TempResourceHandler).Methods("POST")
	router.HandleFunc("/workspace/list", TempResourceHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

// DeleteBookingHandler rewrites fields for booking where applicable
func TempResourceHandler(writer http.ResponseWriter, request *http.Request){
	utils.Ok(writer, request)
}
