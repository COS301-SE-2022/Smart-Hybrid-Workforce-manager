package security

import (
	"api/data"
	"api/db"
	"api/utils"
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func Validate(function HandlerFunc, permissionRequired *data.Permission) HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		access, err := db.Open()
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
		defer access.Close()

		user_id := "11111111-3333-4a06-9983-8b374586e459"
		permissions, err := GetPermissionsUserId(&user_id, access) // TODO [KP]: Fix this once redis is up and running
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}

		authorized := false
		for _, permission := range permissions {
			if permissionRequired.CompareTo(permission) {
				authorized = true
			}
		}
		if !authorized {
			utils.AccessDenied(writer, request, fmt.Errorf("doesn't have permissions to execute query"))
			return
		}
		function(writer, request)
	}
}
