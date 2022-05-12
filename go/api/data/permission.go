package data

import (
	"api/db"
	"database/sql"
	"time"
)

//////////////////////////////////////////////////
// Structures and Variables

// Permission identifies a permission via common attributes
type Permission struct {
	Id                 *string    `json:"id,omitempty"`
	PermissionType     *string    `json:"permission_type,omitempty"`
	PermissionCategory *string    `json:"permission_category,omitempty"`
	PermissionTenant   *string    `json:"permission_tenant,omitempty"`
	PermissionTenantId *string    `json:"permission_tenant_id,omitempty"`
	DateAdded          *time.Time `json:"date_added,omitempty"`
}

// Permissions represent a splice of Permission
type Permissions []*Permission

// PermissionDA provides access to the database for authentication purposes
type PermissionDA struct {
	access *db.Access
}

// NewPermissionDA creates a new data access from a underlying shared data access
func NewPermissionDA(access *db.Access) *PermissionDA {
	return &PermissionDA{
		access: access,
	}
}

//////////////////////////////////////////////////
// Mappers

func mapPermission(rows *sql.Rows) (interface{}, error) {
	var identifier Permission
	err := rows.Scan(
		&identifier.Id,
		&identifier.PermissionType,
		&identifier.PermissionCategory,
		&identifier.PermissionTenant,
		&identifier.PermissionTenantId,
		&identifier.DateAdded,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

//////////////////////////////////////////////////
// Functions

// StoreUserIdentifier stores an identifier
func (access *PermissionDA) StoreUserIdentifier(identifier *Permission) error {
	_, err := access.access.Query(
		`SELECT 1 FROM permission.user_store($1, $2, $3, $4, $5)`, nil,
		identifier.Id, identifier.PermissionType, identifier.PermissionCategory, identifier.PermissionTenant, identifier.PermissionTenantId)
	if err != nil {
		return err
	}
	return nil
}

// StoreUserIdentifier stores an identifier
func (access *PermissionDA) StoreRoleIdentifier(identifier *Permission) error {
	_, err := access.access.Query(
		`SELECT 1 FROM permission.role_store($1, $2, $3, $4, $5)`, nil,
		identifier.Id, identifier.PermissionType, identifier.PermissionCategory, identifier.PermissionTenant, identifier.PermissionTenantId)
	if err != nil {
		return err
	}
	return nil
}
