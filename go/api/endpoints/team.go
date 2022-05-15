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
	router.HandleFunc("/remove", DeleteTeamHandler).Methods("POST")

	router.HandleFunc("/user/create", CreateUserTeamHandler).Methods("POST")
	router.HandleFunc("/user/information", InformationUserTeamHandler).Methods("POST")
	router.HandleFunc("/user/remove", DeleteUserTeamHandler).Methods("POST")

	router.HandleFunc("/association/create", CreateTeamAssociationHandler).Methods("POST")
	router.HandleFunc("/association/information", InformationTeamAssociationHandler).Methods("POST")
	router.HandleFunc("/association/remove", DeleteTeamAssociationHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

///////////////
// Team

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

// InformationTeamHandler gets teams
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

// DeleteTeamHandler removes a booking
func DeleteTeamHandler(writer http.ResponseWriter, request *http.Request) {
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

	teamRemoved, err := da.DeleteIdentifier(&team)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v team removed\n", team.Id)

	utils.JSONResponse(writer, request, teamRemoved)
}

///////////////
// UserTeam

// CreateUserTeamHandler creates a new Userteam
func CreateUserTeamHandler(writer http.ResponseWriter, request *http.Request) {
	var userTeam data.UserTeam

	err := utils.UnmarshalJSON(writer, request, &userTeam)
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

	err = da.CreateUserTeam(&userTeam)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v User Team entry created\n", userTeam.TeamId)

	utils.Ok(writer, request)
}

// InformationUserTeamHandler gets User Teams
func InformationUserTeamHandler(writer http.ResponseWriter, request *http.Request) {
	var userTeam data.UserTeam

	err := utils.UnmarshalJSON(writer, request, &userTeam)
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

	userTeams, err := da.FindUserTeam(&userTeam)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v User Team information requested\n", userTeam.TeamId)

	utils.JSONResponse(writer, request, userTeams)
}

// DeleteUserTeamHandler removes a booking
func DeleteUserTeamHandler(writer http.ResponseWriter, request *http.Request) {
	var userTeam data.UserTeam

	err := utils.UnmarshalJSON(writer, request, &userTeam)
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

	userTeamRemoved, err := da.DeleteUserTeam(&userTeam)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v User Team removed\n", userTeam.TeamId)

	utils.JSONResponse(writer, request, userTeamRemoved)
}

///////////////
// TeamAssociation

// CreateTeamAssociationHandler creates a new teamAssociation
func CreateTeamAssociationHandler(writer http.ResponseWriter, request *http.Request) {
	var teamAssociation data.TeamAssociation

	err := utils.UnmarshalJSON(writer, request, &teamAssociation)
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

	err = da.CreateTeamAssociation(&teamAssociation)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v User Team entry created\n", teamAssociation.TeamId)

	utils.Ok(writer, request)
}

// InformationTeamAssociationHandler gets User Teams
func InformationTeamAssociationHandler(writer http.ResponseWriter, request *http.Request) {
	var teamAssociation data.TeamAssociation

	err := utils.UnmarshalJSON(writer, request, &teamAssociation)
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

	teamAssociations, err := da.FindTeamAssociation(&teamAssociation)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v User Team information requested\n", teamAssociation.TeamId)

	utils.JSONResponse(writer, request, teamAssociations)
}

// DeleteTeamAssociationHandler removes a booking
func DeleteTeamAssociationHandler(writer http.ResponseWriter, request *http.Request) {
	var teamAssociation data.TeamAssociation

	err := utils.UnmarshalJSON(writer, request, &teamAssociation)
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

	teamAssociationRemoved, err := da.DeleteTeamAssociation(&teamAssociation)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v User Team removed\n", teamAssociation.TeamId)

	utils.JSONResponse(writer, request, teamAssociationRemoved)
}
