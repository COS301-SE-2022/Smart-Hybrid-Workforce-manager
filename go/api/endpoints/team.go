package endpoints

import (
	"api/data"
	"api/db"
	"api/security"
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
	router.HandleFunc("/create", security.Validate(CreateTeamHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "TEAM", "IDENTIFIER")})).Methods("POST")
	router.HandleFunc("/information", security.Validate(InformationTeamHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "TEAM", "IDENTIFIER")})).Methods("POST")
	router.HandleFunc("/remove", security.Validate(DeleteTeamHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "TEAM", "IDENTIFIER")})).Methods("POST")

	router.HandleFunc("/user/create", security.Validate(CreateUserTeamHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "TEAM", "USER")})).Methods("POST")
	router.HandleFunc("/user/information", security.Validate(InformationUserTeamHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "TEAM", "USER")})).Methods("POST")
	router.HandleFunc("/user/remove", security.Validate(DeleteUserTeamHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "TEAM", "USER")})).Methods("POST")

	router.HandleFunc("/association/create", security.Validate(CreateTeamAssociationHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "TEAM", "ASSOCIATION")})).Methods("POST")
	router.HandleFunc("/association/information", security.Validate(InformationTeamAssociationHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "TEAM", "ASSOCIATION")})).Methods("POST")
	router.HandleFunc("/association/remove", security.Validate(DeleteTeamAssociationHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "TEAM", "ASSOCIATION")})).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

///////////////
// Team

// CreateTeamHandler creates a new team
func CreateTeamHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Team
	var team data.Team
	err := utils.UnmarshalJSON(writer, request, &team)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Create a database connection
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

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v team created\n", team.Name)
	utils.Ok(writer, request)
}

// InformationTeamHandler gets teams
func InformationTeamHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Team
	var team data.Team
	err := utils.UnmarshalJSON(writer, request, &team)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewTeamDA(access)
	teams, err := da.FindIdentifier(&team, permissions)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v team information requested\n", team.Id)
	utils.JSONResponse(writer, request, teams)
}

// DeleteTeamHandler removes a team
func DeleteTeamHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Team
	var team data.Team
	err := utils.UnmarshalJSON(writer, request, &team)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// Get booking information if no user is defined
	da := data.NewTeamDA(access)
	if team.Id == nil {
		temp, err := da.FindIdentifier(&team, &data.Permissions{data.CreateGenericPermission("VIEW", "TEAM", "IDENTIFIER")})
		team = *temp.FindHead()
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
	}

	// Check if user has permission to delete a team
	if team.Id != nil {
		authorized := false
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == team.Id || permission.PermissionTenantId == nil {
				authorized = true
			}
		}
		if !authorized {
			utils.AccessDenied(writer, request, fmt.Errorf("doesn't have permission to execute query")) // TODO [KP]: Be more descriptive
			return
		}
	}

	teamRemoved, err := da.DeleteIdentifier(&team)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Commit transaction
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

// CreateUserTeamHandler creates user team association
func CreateUserTeamHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Team User
	var userTeam data.UserTeam
	err := utils.UnmarshalJSON(writer, request, &userTeam)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Check if user has permission to create a team for the incomming user
	authorized := false
	if userTeam.UserId != nil {
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == userTeam.UserId || permission.PermissionTenantId == nil {
				authorized = true
			}
		}
	}
	if !authorized {
		utils.AccessDenied(writer, request, fmt.Errorf("doesn't have permission to execute query")) // TODO [KP]: Be more descriptive
		return
	}

	// Create a database connection
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

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v User Team entry created\n", userTeam.TeamId)
	utils.Ok(writer, request)
}

// InformationUserTeamHandler gets user team association
func InformationUserTeamHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Team User
	var userTeam data.UserTeam
	err := utils.UnmarshalJSON(writer, request, &userTeam)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()
	da := data.NewTeamDA(access)

	userTeams, err := da.FindUserTeam(&userTeam, permissions)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v User Team information requested\n", userTeam.TeamId)
	utils.JSONResponse(writer, request, userTeams)
}

// DeleteUserTeamHandler removes a user team associaiton
func DeleteUserTeamHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Team User
	var userTeam data.UserTeam
	err := utils.UnmarshalJSON(writer, request, &userTeam)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Check if user has permission to delete a team for the incomming booking user
	if userTeam.UserId != nil {
		authorized := false
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == userTeam.UserId || permission.PermissionTenantId == nil {
				authorized = true
			}
		}
		if !authorized {
			utils.AccessDenied(writer, request, fmt.Errorf("doesn't have permission to execute query")) // TODO [KP]: Be more descriptive
			return
		}
	}

	// Create a database connection
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

	// Commit transaction
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

// CreateTeamAssociationHandler creates a new team association
func CreateTeamAssociationHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Team Association
	var teamAssociation data.TeamAssociation
	err := utils.UnmarshalJSON(writer, request, &teamAssociation)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Create a database connection
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

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v User Team entry created\n", teamAssociation.TeamId)
	utils.Ok(writer, request)
}

// InformationTeamAssociationHandler gets team associations
func InformationTeamAssociationHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Team Association
	var teamAssociation data.TeamAssociation
	err := utils.UnmarshalJSON(writer, request, &teamAssociation)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewTeamDA(access)
	teamAssociations, err := da.FindTeamAssociation(&teamAssociation, permissions)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v User Team information requested\n", teamAssociation.TeamId)
	utils.JSONResponse(writer, request, teamAssociations)
}

// DeleteTeamAssociationHandler removes a team association
func DeleteTeamAssociationHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Team Association
	var teamAssociation data.TeamAssociation
	err := utils.UnmarshalJSON(writer, request, &teamAssociation)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Create a database connection
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

	// Commit transaction
	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	logger.Access.Printf("%v User Team removed\n", teamAssociation.TeamId)
	utils.JSONResponse(writer, request, teamAssociationRemoved)
}
