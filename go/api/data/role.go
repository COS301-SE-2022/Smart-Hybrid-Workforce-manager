package data

import (
	"api/db"
	"database/sql"
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
