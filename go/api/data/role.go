package data

import (
	"api/db"
	"database/sql"
	"encoding/json"
	"time"
)

//////////////////////////////////////////////////
// Structures and Variables

// Role identifies a Role via common attributes
type Role struct {
	Id        *string    `json:"id,omitempty"`
	RoleName  *string    `json:"role_name,omitempty"`
	DateAdded *time.Time `json:"date_added,omitempty"`
}

// Roles represent a splice of Role
type Roles []*Role

// UserRole identifies a UserRole via common attributes
type UserRole struct {
	RoleId    *string    `json:"role_id,omitempty"`
	UserId    *string    `json:"user_id,omitempty"`
	DateAdded *time.Time `json:"date_added,omitempty"`
}

// Roles represent a splice of Role
type UserRoles []*UserRole

// RoleDA provides access to the database for authentication purposes
type RoleDA struct {
	access *db.Access
}

// NewRoleDA creates a new data access from a underlying shared data access
func NewRoleDA(access *db.Access) *RoleDA {
	return &RoleDA{
		access: access,
	}
}

// Commit commits the current implicit transaction
func (access *RoleDA) Commit() error {
	return access.access.Commit()
}

//////////////////////////////////////////////////
// Mappers

func mapRole(rows *sql.Rows) (interface{}, error) {
	var identifier Role
	err := rows.Scan(
		&identifier.Id,
		&identifier.RoleName,
		&identifier.DateAdded,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

func mapUserRole(rows *sql.Rows) (interface{}, error) {
	var identifier UserRole
	err := rows.Scan(
		&identifier.RoleId,
		&identifier.UserId,
		&identifier.DateAdded,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

//////////////////////////////////////////////////
// Functions

// StoreIdentifier stores an identifier
func (access *RoleDA) StoreIdentifier(identifier *Role) error {
	_, err := access.access.Query(
		`SELECT 1 FROM role.identifier_store($1, $2)`, nil,
		identifier.Id, identifier.RoleName)
	if err != nil {
		return err
	}
	return nil
}

//FindIdentifier finds an identifier
func (access *RoleDA) FindIdentifier(identifier *Role, permissions *Permissions) (Roles, error) {
	permissionContent, err := json.Marshal(*permissions)
	if err != nil {
		return nil, err
	}
	results, err := access.access.Query(
		`SELECT * FROM role.identifier_find($1, $2, $3, $4)`, mapRole,
		identifier.Id, identifier.RoleName, identifier.DateAdded, permissionContent)
	if err != nil {
		return nil, err
	}
	tmp := make([]*Role, 0)
	for r, _ := range results {
		if value, ok := results[r].(Role); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//DeleteIdentifier finds an identifier
func (access *RoleDA) DeleteIdentifier(identifier *Role) (*Role, error) {
	results, err := access.access.Query(
		`SELECT * FROM role.identifier_remove($1)`, mapRole,
		identifier.Id)
	if err != nil {
		return nil, err
	}
	var tmp Roles
	tmp = make([]*Role, 0)
	for r, _ := range results {
		if value, ok := results[r].(Role); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp.FindHead(), nil
}

//FindHead returns the first Booking
func (roles Roles) FindHead() *Role {
	if len(roles) == 0 {
		return nil
	}
	return roles[0]
}

// StoreUserRole stores an identifier
func (access *RoleDA) StoreUserRole(identifier *UserRole) error {
	_, err := access.access.Query(
		`SELECT 1 FROM role.user_store($1, $2)`, nil,
		identifier.RoleId, identifier.UserId)
	if err != nil {
		return err
	}
	return nil
}

//FindUserRole finds an identifier
func (access *RoleDA) FindUserRole(identifier *UserRole, permissions *Permissions) (UserRoles, error) {
	permissionContent, err := json.Marshal(*permissions)
	if err != nil {
		return nil, err
	}
	results, err := access.access.Query(
		`SELECT * FROM role.user_find($1, $2, $3, $4)`, mapUserRole,
		identifier.RoleId, identifier.UserId, identifier.DateAdded, permissionContent)
	if err != nil {
		return nil, err
	}
	tmp := make([]*UserRole, 0)
	for r, _ := range results {
		if value, ok := results[r].(UserRole); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//DeleteUserRole finds an identifier
func (access *RoleDA) DeleteUserRole(identifier *UserRole) (*UserRole, error) {
	results, err := access.access.Query(
		`SELECT * FROM role.user_remove($1, $2)`, mapUserRole,
		identifier.RoleId, identifier.UserId)
	if err != nil {
		return nil, err
	}
	var tmp UserRoles
	tmp = make([]*UserRole, 0)
	for r, _ := range results {
		if value, ok := results[r].(UserRole); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp.FindHead(), nil
}

//FindHead returns the first Booking
func (roles UserRoles) FindHead() *UserRole {
	if len(roles) == 0 {
		return nil
	}
	return roles[0]
}
