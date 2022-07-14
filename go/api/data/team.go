package data

import (
	"api/db"
	"database/sql"
	"encoding/json"
	"time"
)

//////////////////////////////////////////////////
// Structures and Variables

// Identifier identifies a user via common attributes
type Team struct {
	Id          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Capacity    *int       `json:"capacity,omitempty"`
	Picture     *string    `json:"picture,omitempty"`
	Priority    *int       `json:"priority,omitempty"`
	TeamLeadId  *string    `json:"team_lead_id,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
}

// Teams represent a splice of Team
type Teams []*Team

// Identifier identifies a user via common attributes
type UserTeam struct {
	TeamId    *string `json:"team_id,omitempty"`
	UserId    *string `json:"user_id,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
}

// UserTeams represent a splice of UserTeam
type UserTeams []*UserTeam

// Identifier identifies a user via common attributes
type TeamAssociation struct {
	TeamId            *string `json:"team_id,omitempty"`
	TeamIdAssociation *string `json:"team_id_association,omitempty"`
}

// TeamAssociations represent a splice of TeamAssociation
type TeamAssociations []*TeamAssociation

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

// Commit commits the current implicit transaction
func (access *TeamDA) Commit() error {
	return access.access.Commit()
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
		&identifier.Priority,
		&identifier.TeamLeadId,
		&identifier.DateCreated,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

func mapUserTeam(rows *sql.Rows) (interface{}, error) {
	var identifier UserTeam
	err := rows.Scan(
		&identifier.TeamId,
		&identifier.UserId,
		&identifier.DateAdded,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

func mapTeamAssociation(rows *sql.Rows) (interface{}, error) {
	var identifier TeamAssociation
	err := rows.Scan(
		&identifier.TeamId,
		&identifier.TeamIdAssociation,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

//////////////////////////////////////////////////
// Functions

///////////////
// Team

//CreateTeam creates a team
func (access *TeamDA) CreateTeam(identifier *Team) error {
	_, err := access.access.Query(
		`SELECT 1 FROM team.identifier_store($1, $2, $3, $4, $5, $6, $7)`, nil,
		identifier.Id, identifier.Name, identifier.Description, identifier.Capacity, identifier.Picture, identifier.Priority, identifier.TeamLeadId)
	if err != nil {
		return err
	}
	return nil
}

//FindIdentifier finds a team
func (access *TeamDA) FindIdentifier(identifier *Team, permissions *Permissions) (Teams, error) {
	permissionContent, err := json.Marshal(*permissions)
	if err != nil {
		return nil, err
	}
	results, err := access.access.Query(
		`SELECT * FROM team.identifier_find($1, $2, $3, $4, $5, $6, $7, $8, $9)`, mapTeam,
		identifier.Id, identifier.Name, identifier.Description, identifier.Capacity, identifier.Picture, identifier.Priority, identifier.TeamLeadId, identifier.DateCreated, permissionContent)
	if err != nil {
		return nil, err
	}
	tmp := make([]*Team, 0)
	for r, _ := range results {
		if value, ok := results[r].(Team); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//DeleteIdentifier finds an identifier
func (access *TeamDA) DeleteIdentifier(identifier *Team) (*Team, error) {
	results, err := access.access.Query(
		`SELECT * FROM team.identifier_remove($1)`, mapTeam,
		identifier.Id)
	if err != nil {
		return nil, err
	}
	var tmp Teams
	tmp = make([]*Team, 0)
	for r, _ := range results {
		if value, ok := results[r].(Team); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp.FindHead(), nil
}

//FindHead returns the first Team
func (teams Teams) FindHead() *Team {
	if len(teams) == 0 {
		return nil
	}
	return teams[0]
}

///////////////
// UserTeam

//CreateUserTeam connects a user to a team
func (access *TeamDA) CreateUserTeam(identifier *UserTeam) error {
	_, err := access.access.Query(
		`SELECT 1 FROM team.user_store($1, $2)`, nil,
		identifier.TeamId, identifier.UserId)
	if err != nil {
		return err
	}
	return nil
}

//FindUserTeam finds a users team
func (access *TeamDA) FindUserTeam(identifier *UserTeam, permissions *Permissions) (UserTeams, error) {
	permissionContent, err := json.Marshal(*permissions)
	if err != nil {
		return nil, err
	}
	results, err := access.access.Query(
		`SELECT * FROM team.user_find($1, $2, $3, $4)`, mapUserTeam,
		identifier.TeamId, identifier.UserId, identifier.DateAdded, permissionContent)
	if err != nil {
		return nil, err
	}
	tmp := make([]*UserTeam, 0)
	for r, _ := range results {
		if value, ok := results[r].(UserTeam); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//DeleteUserTeam removes a teams user
func (access *TeamDA) DeleteUserTeam(identifier *UserTeam) (*UserTeam, error) {
	results, err := access.access.Query(
		`SELECT * FROM team.user_remove($1, $2)`, mapUserTeam,
		identifier.TeamId, identifier.UserId)
	if err != nil {
		return nil, err
	}
	var tmp UserTeams
	tmp = make([]*UserTeam, 0)
	for r, _ := range results {
		if value, ok := results[r].(UserTeam); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp.FindHead(), nil
}

//FindHead returns the first Team
func (teams UserTeams) FindHead() *UserTeam {
	if len(teams) == 0 {
		return nil
	}
	return teams[0]
}

///////////////
// TeamAssociation

//CreateTeamAssociation connects a user to a team
func (access *TeamDA) CreateTeamAssociation(identifier *TeamAssociation) error {
	_, err := access.access.Query(
		`SELECT 1 FROM team.association_store($1, $2)`, nil,
		identifier.TeamId, identifier.TeamIdAssociation)
	if err != nil {
		return err
	}
	return nil
}

//FindTeamAssociation finds a users team
func (access *TeamDA) FindTeamAssociation(identifier *TeamAssociation, permissions *Permissions) (TeamAssociations, error) {
	permissionContent, err := json.Marshal(*permissions)
	if err != nil {
		return nil, err
	}
	results, err := access.access.Query(
		`SELECT * FROM team.association_find($1, $2, $3)`, mapTeamAssociation,
		identifier.TeamId, identifier.TeamIdAssociation, permissionContent)
	if err != nil {
		return nil, err
	}
	tmp := make([]*TeamAssociation, 0)
	for r, _ := range results {
		if value, ok := results[r].(TeamAssociation); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//DeleteTeamAssociation removes a teams user
func (access *TeamDA) DeleteTeamAssociation(identifier *TeamAssociation) (*TeamAssociation, error) {
	results, err := access.access.Query(
		`SELECT * FROM team.association_remove($1, $2)`, mapTeamAssociation,
		identifier.TeamId, identifier.TeamIdAssociation)
	if err != nil {
		return nil, err
	}
	var tmp TeamAssociations
	tmp = make([]*TeamAssociation, 0)
	for r, _ := range results {
		if value, ok := results[r].(TeamAssociation); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp.FindHead(), nil
}

//FindHead returns the first Team
func (teams TeamAssociations) FindHead() *TeamAssociation {
	if len(teams) == 0 {
		return nil
	}
	return teams[0]
}
