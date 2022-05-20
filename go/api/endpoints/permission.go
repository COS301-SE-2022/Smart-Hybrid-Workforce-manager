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
	router.HandleFunc("/role/remove", security.Validate(DeletePermissionRoleHandler,
		data.CreateGenericPermission("DELETE", "PERMISSION", "ROLE"))).Methods("POST")

	// Permissions Users
	router.HandleFunc("/user/create", security.Validate(CreatePermissionUserHandler,
		data.CreateGenericPermission("CREATE", "PERMISSION", "USER"))).Methods("POST")
	router.HandleFunc("/user/information", security.Validate(InformationPermissionUserHandler,
		data.CreateGenericPermission("VIEW", "PERMISSION", "USER"))).Methods("POST")
	router.HandleFunc("/user/remove", security.Validate(DeletePermissionUserHandler,
		data.CreateGenericPermission("DELETE", "PERMISSION", "USER"))).Methods("POST")

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
	rolePermissions, err := da.FindRolePermission(&rolePermission, security.RemoveUserPermissions(permissions))
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
	logger.Access.Printf("%v permission information requested\n", rolePermission.Id) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, rolePermissions)
}

// DeletePermissionRoleHandler removes role permission
func DeletePermissionRoleHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Permission
	var rolePermission data.Permission
	err := utils.UnmarshalJSON(writer, request, &rolePermission)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Check if user has permission to delete a permission for the incomming role
	authorized := false
	for _, permission := range *permissions {
		// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
		if permission.PermissionTenantId == rolePermission.Id || permission.PermissionTenantId == nil {
			authorized = true
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
	da := data.NewPermissionDA(access)
	permissionRemoved, err := da.DeleteRolePermission(&rolePermission)
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
	logger.Access.Printf("%v permission removed\n", permissionRemoved.Id) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, permissionRemoved)
}

/////////////////////
// User Permissions

// CreatePermissionUserHandler creates a user permission
func CreatePermissionUserHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Permission
	var userPermission data.Permission
	err := utils.UnmarshalJSON(writer, request, &userPermission)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Check if user has permission to create a permission for the incomming user
	authorized := false
	for _, permission := range *permissions {
		// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
		if permission.PermissionTenantId == userPermission.Id || permission.PermissionTenantId == nil {
			authorized = true
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
	err = da.StoreUserPermission(&userPermission)
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
	logger.Access.Printf("%v created\n", userPermission.Id) // TODO [KP]: Be more descriptive
	utils.Ok(writer, request)
}

// InformationPermissionUserHandler gets user permissions
func InformationPermissionUserHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Permission
	var userPermission data.Permission
	err := utils.UnmarshalJSON(writer, request, &userPermission)
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
	userPermissions, err := da.FindUserPermission(&userPermission, security.RemoveRolePermissions(permissions))
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
	logger.Access.Printf("%v permission information requested\n", userPermission.Id) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, userPermissions)
}

// DeletePermissionUserHandler removes user permission
func DeletePermissionUserHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Permission
	var userPermission data.Permission
	err := utils.UnmarshalJSON(writer, request, &userPermission)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Check if user has permission to delete a permission for the incomming user
	authorized := false
	for _, permission := range *permissions {
		// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
		if permission.PermissionTenantId == userPermission.Id || permission.PermissionTenantId == nil {
			authorized = true
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

	da := data.NewPermissionDA(access)
	permissionRemoved, err := da.DeleteUserPermission(&userPermission)
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
	logger.Access.Printf("%v permission removed\n", userPermission.Id) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, permissionRemoved)
}
