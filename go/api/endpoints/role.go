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

//ResourceHandlers
func RoleHandlers(router *mux.Router) error {
	router.HandleFunc("/create", security.Validate(CreateRoleHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "ROLE", "IDENTIFIER")})).Methods("POST")
	router.HandleFunc("/information", security.Validate(InformationRolesHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "ROLE", "IDENTIFIER")})).Methods("POST")
	router.HandleFunc("/remove", security.Validate(DeleteRoleHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "ROLE", "IDENTIFIER")})).Methods("POST")

	router.HandleFunc("/user/create", security.Validate(CreateUserRoleHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "ROLE", "USER")})).Methods("POST")
	router.HandleFunc("/user/information", security.Validate(InformationUserRolesHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "ROLE", "USER")})).Methods("POST")
	router.HandleFunc("/user/remove", security.Validate(DeleteUserRoleHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "ROLE", "USER")})).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

// CreateRoleHandler creates or updates a role
func CreateRoleHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Role
	var role data.Role
	err := utils.UnmarshalJSON(writer, request, &role)
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

	// TODO [KP]: Do more checks like if there exists a Role already etc

	da := data.NewRoleDA(access)
	err = da.StoreIdentifier(&role)
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
	logger.Access.Printf("%v created\n", role.Id)
	utils.Ok(writer, request)
}

// InformationRolesHandler gets roles
func InformationRolesHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Role
	var role data.Role
	err := utils.UnmarshalJSON(writer, request, &role)
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

	da := data.NewRoleDA(access)
	roles, err := da.FindIdentifier(&role, permissions)
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
	logger.Access.Printf("%v Role information requested\n", role.Id)
	utils.JSONResponse(writer, request, roles)
}

// DeleteRoleHandler removes a role
func DeleteRoleHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Role
	var role data.Role
	err := utils.UnmarshalJSON(writer, request, &role)
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

	// Get role information if no role id is defined
	da := data.NewRoleDA(access)
	if role.Id == nil {
		temp, err := da.FindIdentifier(&role, &data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")})
		role = *temp.FindHead()
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
	}

	// Check if user has permission to delete the role
	if role.Id != nil {
		authorized := false
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == role.Id || permission.PermissionTenantId == nil {
				authorized = true
			}
		}
		if !authorized {
			utils.AccessDenied(writer, request, fmt.Errorf("doesn't have permission to execute query")) // TODO [KP]: Be more descriptive
			return
		}
	}

	roleRemoved, err := da.DeleteIdentifier(&role)
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
	logger.Access.Printf("%v Role removed\n", role.Id)
	utils.JSONResponse(writer, request, roleRemoved)
}

// CreateUserRoleHandler creates or updates a role
func CreateUserRoleHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Role
	var role data.UserRole
	err := utils.UnmarshalJSON(writer, request, &role)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Check if user has permission to create a role for the incomming user
	authorized := false
	if role.UserId != nil {
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == role.UserId || permission.PermissionTenantId == nil {
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

	// TODO [KP]: Do more checks like if there exists a Role already etc

	da := data.NewRoleDA(access)
	err = da.StoreUserRole(&role)
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
	logger.Access.Printf("%v created\n", role.RoleId)
	utils.Ok(writer, request)
}

// InformationUserRolesHandler gets roles
func InformationUserRolesHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Role
	var role data.UserRole
	err := utils.UnmarshalJSON(writer, request, &role)
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

	da := data.NewRoleDA(access)
	roles, err := da.FindUserRole(&role, permissions)
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
	logger.Access.Printf("%v Role information requested\n", role.RoleId)
	utils.JSONResponse(writer, request, roles)
}

// DeleteUserRoleHandler removes a role
func DeleteUserRoleHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Role
	var role data.UserRole
	err := utils.UnmarshalJSON(writer, request, &role)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	// Check if user has permission to delete a role for the incomming user
	if role.UserId != nil {
		authorized := false
		for _, permission := range *permissions {
			// A permission tenant id of nil means the user is allowed to perform the action on all tenants of the type
			if permission.PermissionTenantId == role.UserId || permission.PermissionTenantId == nil {
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

	da := data.NewRoleDA(access)
	roleRemoved, err := da.DeleteUserRole(&role)
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
	logger.Access.Printf("%v Role removed\n", role.RoleId)
	utils.JSONResponse(writer, request, roleRemoved)
}
