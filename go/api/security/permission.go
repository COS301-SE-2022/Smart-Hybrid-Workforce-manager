package security

import (
	"api/data"
	"api/db"
)

// GetUserPermissions
func GetUserPermissions(userId *string, access *db.Access) (data.Permissions, error) {
	// Create data-access
	da := data.NewPermissionDA(access)

	// Get user permissions
	temp := "USER"
	permissions, err := da.FindPermission(&data.Permission{PermissionId: userId, PermissionIdType: &temp}, &data.Permissions{data.CreateGenericPermission("VIEW", "PERMISSION", "USER")})
	if err != nil {
		return nil, err
	}

	// for _, permission := range permissions {
	// 	logger.Access.Printf("%v %v %v %v\n", *permission.PermissionIdType, *permission.PermissionType, *permission.PermissionCategory, *permission.PermissionTenant)
	// }

	// Get user roles
	dr := data.NewRoleDA(access)
	roles, err := dr.FindUserRole(&data.UserRole{UserId: userId}, &data.Permissions{data.CreateGenericPermission("VIEW", "ROLE", "USER"), data.CreateGenericPermission("VIEW", "USER", "ROLE")})
	if err != nil {
		return nil, err
	}

	// for _, role := range roles {
	// 	logger.Access.Printf("ROLE: %v\n", *role.RoleId)
	// }

	// Get role permissions
	temp = "ROLE"
	for _, role := range roles {
		permission, err := da.FindPermission(&data.Permission{PermissionId: role.RoleId, PermissionIdType: &temp}, &data.Permissions{data.CreateGenericPermission("VIEW", "PERMISSION", "ROLE")})
		if err != nil {
			return nil, err
		}
		for _, entry := range permission {
			permissions = append(permissions, entry)
		}
	}

	// Get user teams
	dt := data.NewTeamDA(access)
	teams, err := dt.FindUserTeam(&data.UserTeam{UserId: userId}, &data.Permissions{data.CreateGenericPermission("VIEW", "TEAM", "USER"), data.CreateGenericPermission("VIEW", "USER", "TEAM")})
	if err != nil {
		return nil, err
	}

	// Get team permissions
	temp = "TEAM"
	for _, team := range teams {
		permission, err := da.FindPermission(&data.Permission{PermissionId: team.TeamId, PermissionIdType: &temp}, &data.Permissions{data.CreateGenericPermission("VIEW", "PERMISSION", "TEAM")})
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
			roleUsers, err := dr.FindUserRole(&data.UserRole{RoleId: permission.PermissionTenantId}, &data.Permissions{data.CreateGenericPermission("VIEW", "ROLE", "USER")})
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
