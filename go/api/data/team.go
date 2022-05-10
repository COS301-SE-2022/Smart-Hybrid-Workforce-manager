package data

import (
	"api/db"
	"database/sql"
	"time"
)

//////////////////////////////////////////////////
// Structures and Variables

// Identifier identifies a user via common attributes
type Team struct {
	Id          *string   `json:"id,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Capacity    *int      `json:"capacity,omitempty"`
	Picture     *string   `json:"picture,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
}

// TeamDA provides access to the database for team management
type TeamDA struct {
	access *db.Access
}

// NewTeamDA creates a new data access from a underlying shared data access
func NewTeamDA(access *db.Access) *TeamDA {
	return &TeamDA{
		access: access,
	}
}

//////////////////////////////////////////////////
// Mappers

func mapTeam(rows *sql.Rows) (interface{}, error) {
	var identifier Team
	err := rows.Scan(
		&identifier.Id,
		&identifier.Name,
		&identifier.Description,
		&identifier.Capacity,
		&identifier.Picture,
		&identifier.DateCreated,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

//////////////////////////////////////////////////
// Functions

//CreateTeam creates a team
func (da *TeamDA) CreateTeam(identifier *Team) error {
	_, err := da.access.Query(
		`SELECT 1 FROM team.identifier_store($1, $2, $3, $4, $5)`, nil,
		identifier.Id, identifier.Name, identifier.Description, identifier.Capacity, identifier.Picture)
	if err != nil {
		return err
	}
	return nil
}
