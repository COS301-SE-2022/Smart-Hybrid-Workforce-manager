package security

import (
	"api/data"
	"api/db"
)

// CheckUserPermission checks whether the passed user has the passed parameter
func CheckUserPermission(user *data.User, permission *data.Permission) (bool, error) {
	return false, nil
}

// GetPermissionsUserId
func GetPermissionsUserId(userId *string, access *db.Access) (data.Permissions, error) {
	da := data.NewPermissionDA(access)

	permissions, err := da.FindUserPermission(&data.Permission{Id: userId})
	if err != nil {
		return nil, err
	}

	// get user roles
	dr := data.NewRoleDA(access)

	roles, err := dr.FindUserRole(&data.UserRole{UserId: userId})
	if err != nil {
		return nil, err
	}

	// get role permissions
	for _, role := range roles {
		permission, err := da.FindRolePermission(&data.Permission{Id: role.RoleId})
		if err != nil {
			return nil, err
		}
		for _, entry := range permission {
			permissions = append(permissions, entry)
		}
	}

	var result data.Permissions
	// convert Role Id permissions to User Id permissions
	for _, permission := range permissions {
		if *permission.PermissionTenant == "USER" {
			result = append(result, permission)
		} else {
			roleUsers, err := dr.FindUserRole(&data.UserRole{RoleId: permission.PermissionTenantId})
			if err != nil {
				return nil, err
			}
			for _, roleUser := range roleUsers {
				permissionTenant := "USER"
				result = append(result, &data.Permission{Id: permission.Id, PermissionType: permission.PermissionType,
					PermissionCategory: permission.PermissionCategory, PermissionTenant: &permissionTenant, PermissionTenantId: roleUser.UserId, DateAdded: permission.DateAdded})
			}
		}
	}

	return result, nil
}
