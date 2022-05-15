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

//////////////////////////////////////////////////
// Structures and Variables

/////////////////////////////////////////////
// Endpoints

//ResourceHandlers
func RoleHandlers(router *mux.Router) error {
	router.HandleFunc("/create", CreateRoleHandler).Methods("POST")
	router.HandleFunc("/information", InformationRolesHandler).Methods("POST")
	router.HandleFunc("/remove", DeleteRoleHandler).Methods("POST")

	router.HandleFunc("/user/create", CreateUserRoleHandler).Methods("POST")
	router.HandleFunc("/user/information", InformationUserRolesHandler).Methods("POST")
	router.HandleFunc("/user/remove", DeleteUserRoleHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

// CreateRoleHandler creates or updates a Role
func CreateRoleHandler(writer http.ResponseWriter, request *http.Request) {
	var role data.Role

	err := utils.UnmarshalJSON(writer, request, &role)
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

	da := data.NewRoleDA(access)

	// TODO [KP]: Do more checks like if there exists a Role already etc

	err = da.StoreIdentifier(&role)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v created\n", role.Id)

	utils.Ok(writer, request)
}

// InformationRolesHandler gets Roles
func InformationRolesHandler(writer http.ResponseWriter, request *http.Request) {
	var role data.Role

	err := utils.UnmarshalJSON(writer, request, &role)
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

	da := data.NewRoleDA(access)

	roles, err := da.FindIdentifier(&role)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v Role information requested\n", role.Id)

	utils.JSONResponse(writer, request, roles)
}

// DeleteRoleHandler removes a Role
func DeleteRoleHandler(writer http.ResponseWriter, request *http.Request) {
	var role data.Role

	err := utils.UnmarshalJSON(writer, request, &role)
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

	da := data.NewRoleDA(access)

	roleRemoved, err := da.DeleteIdentifier(&role)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v Role removed\n", role.Id)

	utils.JSONResponse(writer, request, roleRemoved)
}

// CreateUserRoleHandler creates or updates a Role
func CreateUserRoleHandler(writer http.ResponseWriter, request *http.Request) {
	var role data.UserRole

	err := utils.UnmarshalJSON(writer, request, &role)
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

	da := data.NewRoleDA(access)

	// TODO [KP]: Do more checks like if there exists a Role already etc

	err = da.StoreUserRole(&role)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v created\n", role.RoleId)

	utils.Ok(writer, request)
}

// InformationUserRolesHandler gets Roles
func InformationUserRolesHandler(writer http.ResponseWriter, request *http.Request) {
	var role data.UserRole

	err := utils.UnmarshalJSON(writer, request, &role)
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

	da := data.NewRoleDA(access)

	roles, err := da.FindUserRole(&role)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v Role information requested\n", role.RoleId)

	utils.JSONResponse(writer, request, roles)
}

// DeleteUserRoleHandler removes a Role
func DeleteUserRoleHandler(writer http.ResponseWriter, request *http.Request) {
	var role data.UserRole

	err := utils.UnmarshalJSON(writer, request, &role)
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

	da := data.NewRoleDA(access)

	roleRemoved, err := da.DeleteUserRole(&role)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v Role removed\n", role.RoleId)

	utils.JSONResponse(writer, request, roleRemoved)
}
