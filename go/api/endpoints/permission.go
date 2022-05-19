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

//////////////////////////////////////////////////
// Structures and Variables

/////////////////////////////////////////////
// Endpoints

//PermissionHandlers
func PermissionHandlers(router *mux.Router) error {
	// Permissions Roles
	router.HandleFunc("/role/create", security.Validate(CreatePermissionRoleHandler,
		data.CreateGenericPermission("CREATE", "PERMISSION", "ROLE"))).Methods("POST")
	router.HandleFunc("/role/information", security.Validate(InformationPermissionRoleHandler,
		data.CreateGenericPermission("VIEW", "PERMISSION", "ROLE"))).Methods("POST")
	router.HandleFunc("/role/remove", DeletePermissionRoleHandler).Methods("POST")

	// Permissions Users
	router.HandleFunc("/user/create", CreatePermissionUserHandler).Methods("POST")
	router.HandleFunc("/user/information", InformationPermissionUserHandler).Methods("POST")
	router.HandleFunc("/user/remove", DeletePermissionUserHandler).Methods("POST")

	return nil
}

/////////////////////////////////////////////
// Functions

/////////////////////
// Role Permissions

// CreatePermissionRoleHandler creates or updates a role permission
func CreatePermissionRoleHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Permission
	var rolePermission data.Permission
	err := utils.UnmarshalJSON(writer, request, &rolePermission)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Check if user has permission to create a role permission for the incomming role
	authorized := false
	if rolePermission.Id != nil {
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == rolePermission.Id || permission.PermissionTenantId == nil {
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

	// TODO [KP]: Do more checks like if they already have a permission etc

	da := data.NewPermissionDA(access)
	err = da.StoreRolePermission(&rolePermission)
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
	logger.Access.Printf("%v created\n", rolePermission.Id) // TODO [KP]: Be more descriptive
	utils.Ok(writer, request)
}

// InformationPermissionRoleHandler gets role permissions
func InformationPermissionRoleHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Booking
	var rolePermission data.Permission
	err := utils.UnmarshalJSON(writer, request, &rolePermission)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// No check for permissions the database handles information permissions

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	// TODO [KP]: null checks etc.

	da := data.NewPermissionDA(access)
	rolePermissions, err := da.FindRolePermission(&rolePermission)
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

	logger.Access.Printf("%v permission information requested\n", rolePermission.Id)

	utils.JSONResponse(writer, request, rolePermissions)
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

/////////////////////
// User Permissions

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
