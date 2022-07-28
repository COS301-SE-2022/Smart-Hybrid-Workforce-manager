package security

import (
	"api/data"
	"api/db"
	"api/utils"
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, *data.Permissions)
type HandlerFuncOut func(http.ResponseWriter, *http.Request)

// Validate validates if the sender has permissions to execute the endpoint
func Validate(function HandlerFunc, permissionRequired *data.Permissions) HandlerFuncOut {
	return func(writer http.ResponseWriter, request *http.Request) {
		access, err := db.Open()
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}
		defer access.Close()

		// if the user does not require permissions
		if permissionRequired == nil {
			function(writer, request, nil)
			return
		}

		// userInfo,err := redis.GetUserInfo(request)
		// if err != nil{
		// 	logger.Error.Println(err)
		// 	utils.BadRequest(writer, request, "Invalid Authorization Token")
		// 	return
		// }
		// // Check if user data is null
		// user_id := userInfo.User_id
		user_id := "00000000-0000-0000-0000-000000000000"
		permissions, err := GetUserPermissions(&user_id, access)
		if err != nil {
			utils.InternalServerError(writer, request, err)
			return
		}

		// TODO [KP]: unpack permissions

		// filter permissions based on the permission required
		var filteredPermissions data.Permissions
		for _, permission := range permissions {
			if permissionRequired.CompareTo(permission) {
				filteredPermissions = append(filteredPermissions, permission)
			}
		}
		if len(filteredPermissions) == 0 {
			utils.AccessDenied(writer, request, fmt.Errorf("the user doesn't have permission to execute query")) // TODO [KP]: Be more descriptive
			return
		}
		function(writer, request, &filteredPermissions)
	}
}
