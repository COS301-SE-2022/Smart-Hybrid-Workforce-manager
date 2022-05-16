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

//PermissionHandlers
func PermissionHandlers(router *mux.Router) error {
	router.HandleFunc("/role/create", CreatePermissionRoleHandler).Methods("POST")
	router.HandleFunc("/role/information", InformationPermissionRoleHandler).Methods("POST")
	router.HandleFunc("/role/remove", DeletePermissionRoleHandler).Methods("POST")

	router.HandleFunc("/user/create", CreatePermissionUserHandler).Methods("POST")
	router.HandleFunc("/user/information", InformationPermissionUserHandler).Methods("POST")
	router.HandleFunc("/user/remove", DeletePermissionUserHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

// CreatePermissionRoleHandler creates a role permission
func CreatePermissionRoleHandler(writer http.ResponseWriter, request *http.Request) {
	var permission data.Permission

	err := utils.UnmarshalJSON(writer, request, &permission)
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

	da := data.NewPermissionDA(access)

	// TODO [KP]: Do more checks like if they already have a permission etc

	err = da.StoreRolePermission(&permission)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v created\n", permission.Id)

	utils.Ok(writer, request)
}

// CreatePermissionUserHandler creates a user permission
func CreatePermissionUserHandler(writer http.ResponseWriter, request *http.Request) {
	var permission data.Permission

	err := utils.UnmarshalJSON(writer, request, &permission)
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

	da := data.NewPermissionDA(access)

	// TODO [KP]: Do more checks like if they already have a permission etc

	err = da.StoreUserPermission(&permission)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v created\n", permission.Id)

	utils.Ok(writer, request)
}

// InformationPermissionUserHandler gets user permissions
func InformationPermissionUserHandler(writer http.ResponseWriter, request *http.Request) {
	var permission data.Permission

	err := utils.UnmarshalJSON(writer, request, &permission)
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

	da := data.NewPermissionDA(access)

	permissions, err := da.FindUserPermission(&permission)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v permission information requested\n", permission.Id)

	utils.JSONResponse(writer, request, permissions)
}

// InformationPermissionRoleHandler gets role permissions
func InformationPermissionRoleHandler(writer http.ResponseWriter, request *http.Request) {
	var permission data.Permission

	err := utils.UnmarshalJSON(writer, request, &permission)
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

	da := data.NewPermissionDA(access)

	permissions, err := da.FindRolePermission(&permission)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v permission information requested\n", permission.Id)

	utils.JSONResponse(writer, request, permissions)
}

// DeletePermissionUserHandler removes user permission
func DeletePermissionUserHandler(writer http.ResponseWriter, request *http.Request) {
	var permission data.Permission

	err := utils.UnmarshalJSON(writer, request, &permission)
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

	da := data.NewPermissionDA(access)

	permissionRemoved, err := da.DeleteUserPermission(&permission)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v permission removed\n", permission.Id)

	utils.JSONResponse(writer, request, permissionRemoved)
}

// DeletePermissionRoleHandler removes role permission
func DeletePermissionRoleHandler(writer http.ResponseWriter, request *http.Request) {
	var permission data.Permission

	err := utils.UnmarshalJSON(writer, request, &permission)
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

	da := data.NewPermissionDA(access)

	permissionRemoved, err := da.DeleteRolePermission(&permission)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v permission removed\n", permission.Id)

	utils.JSONResponse(writer, request, permissionRemoved)
}
