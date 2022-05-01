package data

import (
	"api/db"
	"database/sql"
	"time"
)

//////////////////////////////////////////////////
// Structures and Variables

// Identifier identifies a user via common attributes
type Identifier struct {
	Identifier  string    `json:"identifier,omitempty"`
	FirstName   *string   `json:"first_name,omitempty"`
	LastName    *string   `json:"last_name,omitempty"`
	Email       *string   `json:"email,omitempty"`
	Picture     *string   `json:"picture,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
}

// Credential identifies a login (not a user)
type Credential struct {
	ID             string    `json:"id,omitempty"`
	Secret         string    `json:"secret,omitempty"`
	Active         bool      `json:"active,omitempty"`
	Identifier     string    `json:"identifier,omitempty"`
	FailedAttempts int       `json:"failed_attempts,omitempty"`
	LastAccessed   time.Time `json:"last_accessed,omitempty"`
}

// UserDA provides access to the database for authentication purposes
type UserDA struct {
	access *db.Access
}

// NewUserDA creates a new data access from a underlying shared data access
func NewUserDA(access *db.Access) *UserDA {
	return &UserDA{
		access: access,
	}
}

//////////////////////////////////////////////////
// Mappers

func mapIdentifier(rows *sql.Rows) (interface{}, error) {
	var identifier Identifier
	err := rows.Scan(
		&identifier.Identifier,
		&identifier.FirstName,
		&identifier.LastName,
		&identifier.Email,
		&identifier.Picture,
		&identifier.DateCreated,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

func mapCredential(rows *sql.Rows) (interface{}, error) {
	var cred Credential
	err := rows.Scan(
		&cred.ID,
		&cred.Secret,
		&cred.Active,
		&cred.FailedAttempts,
		&cred.LastAccessed,
		&cred.Identifier,
	)
	if err != nil {
		return nil, err
	}
	return cred, nil
}

//////////////////////////////////////////////////
// Functions

//StoreIdentifier stores an identifier
func (da *UserDA) StoreIdentifier(identifier *Identifier) error {
	_, err := da.access.Query(
		`SELECT 1 FROM "user".identifier_store($1, $2, $3, $4, $5)`, nil,
		identifier.Identifier, identifier.FirstName, identifier.LastName, identifier.Email, identifier.Picture)
	if err != nil {
		return err
	}
	return nil
}

//Find finds an identifier
func (access *UserDA) FindIdentifier(id string) (*Identifier, error) {
	results, err := access.access.Query(
		`SELECT * FROM "user".identifier_find($1)`, mapIdentifier, id)
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		if identifier, ok := result.(Identifier); ok {
			return &identifier, nil
		}
	}
	return nil, nil
}

// StoreCredential stores a credential
func (access *UserDA) StoreCredential(Id string, secret *string, identifier string) error {
	_, err := access.access.Query(`SELECT * FROM "user".credential_store($1, $2, $3)`, nil, Id, secret, identifier)
	if err != nil {
		return err
	}
	return nil
}

// Commit commits the current implicit transaction
func (access *UserDA) Commit() error {
	return access.access.Commit()
}
