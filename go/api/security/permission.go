package security

import (
	"api/data"
	"api/db"
)

// GetUserPermissions
func GetUserPermissions(userId *string, access *db.Access) (data.Permissions, error) {
	// Create data-access
	da := data.NewPermissionDA(access)
	permissions, err := da.FindUserPermission(&data.Permission{Id: userId})
	if err != nil {
		return nil, err
	}

	// Get user roles
	dr := data.NewRoleDA(access)
	roles, err := dr.FindUserRole(&data.UserRole{UserId: userId})
	if err != nil {
		return nil, err
	}

	// Get role permissions
	for _, role := range roles {
		permission, err := da.FindRolePermission(&data.Permission{Id: role.RoleId})
		if err != nil {
			return nil, err
		}
		for _, entry := range permission {
			permissions = append(permissions, entry)
		}
	}

	// Convert Role Id permissions to User Id permissions
	for _, permission := range permissions {
		if *permission.PermissionTenant == "ROLE" {
			roleUsers, err := dr.FindUserRole(&data.UserRole{RoleId: permission.PermissionTenantId})
			if err != nil {
				return nil, err
			}
			for _, roleUser := range roleUsers {
				permissionTenant := "USER"
				permissions = append(permissions, &data.Permission{Id: permission.Id, PermissionType: permission.PermissionType,
					PermissionCategory: permission.PermissionCategory, PermissionTenant: &permissionTenant, PermissionTenantId: roleUser.UserId, DateAdded: permission.DateAdded})
			}
		}
	}

	return permissions, nil
}

//RemoveRolePermissions removes role permissions from array
func RemoveRolePermissions(permissions *data.Permissions) *data.Permissions {
	var result data.Permissions
	for _, permission := range *permissions {
		if *permission.PermissionTenant != "ROLE" {
			result = append(result, permission)
		}
	}
	return &result
}

//RemoveRolePermissions removes user permissions from array
func RemoveUserPermissions(permissions *data.Permissions) *data.Permissions {
	var result data.Permissions
	for _, permission := range *permissions {
		if *permission.PermissionTenant != "USER" {
			result = append(result, permission)
		}
	}
	return &result
}
