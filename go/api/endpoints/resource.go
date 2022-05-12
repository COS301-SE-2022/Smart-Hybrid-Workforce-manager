package endpoints

import "github.com/gorilla/mux"

//////////////////////////////////////////////////
// Structures and Variables

type CreateResourceStruct struct {
	//Name        *string `json:"name,omitempty"`
	//Description *string `json:"description,omitempty"`
	//Capacity    *int    `json:"capacity,omitempty"`
	//Picture     *string `json:"picture,omitempty"`
}

/////////////////////////////////////////////
// Endpoints

//ResourceHandlers manages Resources
func ResourceHandlers(router *mux.Router) error {
	//router.HandleFunc("/create", CreateTeamHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions
