package endpoints

import (
	"api/data"
	"api/db"
	"api/utils"
	"lib/logger"
	"net/http"

	"github.com/gorilla/mux"
)

//////////////////////////////////////////////////
// Structures and Variables

type CreateTeamStruct struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Capacity    *int    `json:"capacity,omitempty"`
	Picture     *string `json:"picture,omitempty"`
}

/////////////////////////////////////////////
// Endpoints

//TeamHandlers manages teams
func TeamHandlers(router *mux.Router) error {
	router.HandleFunc("/create", CreateTeamHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

// CreateTeamHandler creates a new team
func CreateTeamHandler(writer http.ResponseWriter, request *http.Request) {
	var createTeamStruct CreateTeamStruct

	err := utils.UnmarshalJSON(writer, request, &createTeamStruct)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewTeamDA(access)

	team := &data.Team{
		Name:        createTeamStruct.Name,
		Description: createTeamStruct.Description,
		Capacity:    createTeamStruct.Capacity,
		Picture:     createTeamStruct.Picture,
	}

	err = da.CreateTeam(team)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v team created\n", team.Name)

	utils.Ok(writer, request)
}
