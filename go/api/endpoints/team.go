package endpoints

import (
	"api/data"
	"api/db"
	"api/utils"
	"fmt"
	"lib/logger"
	"net/http"

	"github.com/gorilla/mux"
)

/////////////////////////////////////////////
// Endpoints

//TeamHandlers manages teams
func TeamHandlers(router *mux.Router) error {
	router.HandleFunc("/create", CreateTeamHandler).Methods("POST")
	router.HandleFunc("/information", InformationTeamHandler).Methods("POST")
	router.HandleFunc("/update", UpdateTeamHandler).Methods("POST")
	router.HandleFunc("/members", UpdateTeamMemberHandler).Methods("POST")
	router.HandleFunc("/list", RemoveTeamMember).Methods("POST")
	router.HandleFunc("/remove", DeleteTeam).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

// CreateTeamHandler creates a new team
func CreateTeamHandler(writer http.ResponseWriter, request *http.Request) {
	var team data.Team

	err := utils.UnmarshalJSON(writer, request, &team)
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

	err = da.CreateTeam(&team)
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

func InformationTeamHandler(writer http.ResponseWriter, request *http.Request) {
	var team data.Team

	err := utils.UnmarshalJSON(writer, request, &team)
	if err != nil {
		fmt.Println(err)
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

	teams, err := da.FindIdentifier(&team)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v team information requested\n", team.Id)

	utils.JSONResponse(writer, request, teams)
}

func UpdateTeamHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Info.Println("team update requested")
	name := "Team#1"
	t := data.Team{
		Name: &name,
	}
	utils.JSONResponse(writer, request, t)
}

func UpdateTeamMemberHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Info.Println("team member addition requested")
	name := "Team#1"
	t := data.Team{
		Name: &name,
	}
	utils.JSONResponse(writer, request, t)
}

func RemoveTeamMember(writer http.ResponseWriter, request *http.Request) {
	logger.Info.Println("team member remove requested")
	name := "Team#1"
	t := data.Team{
		Name: &name,
	}
	utils.JSONResponse(writer, request, t)
}

func DeleteTeam(writer http.ResponseWriter, request *http.Request) {
	logger.Info.Println("team delete requested")
	name := "Team#1"
	t := data.Team{
		Name: &name,
	}
	utils.JSONResponse(writer, request, t)
}
