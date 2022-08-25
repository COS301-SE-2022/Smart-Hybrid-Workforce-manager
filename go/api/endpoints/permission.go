package endpoints

import (
	"api/data"
	"api/db"
	"api/security"
	"lib/utils"
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
	router.HandleFunc("/create", security.Validate(CreatePermissionIdentifierHandler,
		&data.Permissions{data.CreateGenericPermission("CREATE", "PERMISSION", "IDENTIFIER")})).Methods("POST")
	router.HandleFunc("/information", security.Validate(InformationPermissionIdentifierHandler,
		&data.Permissions{data.CreateGenericPermission("VIEW", "PERMISSION", "IDENTIFIER")})).Methods("POST")
	router.HandleFunc("/remove", security.Validate(DeletePermissionIdentifierHandler,
		&data.Permissions{data.CreateGenericPermission("DELETE", "PERMISSION", "IDENTIFIER")})).Methods("POST")

	return nil
}

/////////////////////////////////////////////
// Functions

// CreatePermissionIdentifierHandler creates or updates a permission identifier
func CreatePermissionIdentifierHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Permission
	var permission data.Permission
	err := utils.UnmarshalJSON(writer, request, &permission)
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

	// TODO [KP]: Do more checks like if they already have a permission etc

	da := data.NewPermissionDA(access)
	err = da.StorePermission(&permission)
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
	logger.Access.Printf("%v created\n", permission.Id) // TODO [KP]: Be more descriptive
	utils.Ok(writer, request)
}

// InformationPermissionIdentifierHandler gets role permissions
func InformationPermissionIdentifierHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Booking
	var permission data.Permission
	err := utils.UnmarshalJSON(writer, request, &permission)
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
	tempPermissions, err := da.FindPermission(&permission, security.RemoveUserPermissions(permissions))
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
	logger.Access.Printf("%v permission information requested\n", permission.Id) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, tempPermissions)
}

// DeletePermissionIdentifierHandler removes role permission
func DeletePermissionIdentifierHandler(writer http.ResponseWriter, request *http.Request, permissions *data.Permissions) {
	// Unmarshal Permission
	var permission data.Permission
	err := utils.UnmarshalJSON(writer, request, &permission)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	logger.Access.Printf("%v PERMISSION ID\n", *permission.Id)

	// Create a database connection
	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()
	da := data.NewPermissionDA(access)
	permissionRemoved, err := da.DeletePermission(&permission)
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
	logger.Access.Printf("%v permission removed\n", permissionRemoved) // TODO [KP]: Be more descriptive
	utils.JSONResponse(writer, request, permissionRemoved)
}
