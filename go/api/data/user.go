package data

import (
	"api/db"
	"database/sql"
	"time"
)

//////////////////////////////////////////////////
// Structures and Variables

// User identifies a user via common attributes
type User struct {
	Id          *string   `json:"id,omitempty"`
	Identifier  *string   `json:"identifier,omitempty"`
	FirstName   *string   `json:"first_name,omitempty"`
	LastName    *string   `json:"last_name,omitempty"`
	Email       *string   `json:"email,omitempty"`
	Picture     *string   `json:"picture,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
}

// Credential identifies a login (not a user)
type Credential struct {
	Id             string    `json:"id,omitempty"`
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

// Commit commits the current implicit transaction
func (access *UserDA) Commit() error {
	return access.access.Commit()
}

//////////////////////////////////////////////////
// Mappers

func mapUser(rows *sql.Rows) (interface{}, error) {
	var identifier User
	err := rows.Scan(
		&identifier.Id,
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
		&cred.Id,
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
func (access *UserDA) StoreIdentifier(identifier *User) error {
	_, err := access.access.Query(
		`SELECT 1 FROM "user".identifier_store($1, $2, $3, $4, $5, $6)`, nil,
		identifier.Id, identifier.Identifier, identifier.FirstName, identifier.LastName, identifier.Email, identifier.Picture)
	if err != nil {
		return err
	}
	return nil
}

//FindIdentifier finds an identifier
func (access *UserDA) FindIdentifier(identifier *User) (*User, error) {
	results, err := access.access.Query(
		`SELECT * FROM "user".identifier_find($1, $2, $3, $4, $5, $6, $7)`, mapUser,
		identifier.Id, identifier.Identifier, identifier.FirstName, identifier.LastName, identifier.Email, identifier.Picture, identifier.DateCreated)
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		if identifier, ok := result.(User); ok {
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
