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
	Id                 *string    `json:"id,omitempty"`
	Identifier         *string    `json:"identifier,omitempty"`
	FirstName          *string    `json:"first_name,omitempty"`
	LastName           *string    `json:"last_name,omitempty"`
	Email              *string    `json:"email,omitempty"`
	Picture            *string    `json:"picture,omitempty"`
	DateCreated        time.Time  `json:"date_created,omitempty"`
	WorkFromHome       *bool      `json:"work_from_home,omitempty"`
	Parking            *string    `json:"parking,omitempty"`
	OfficeDays         *int       `json:"office_days,omitempty"`
	PreferredStartTime *time.Time `json:"preferred_start_time,omitempty"`
	PreferredEndTime   *time.Time `json:"preferred_end_time,omitempty"`
}

// Users represent a splice of User
type Users []*User

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
		&identifier.WorkFromHome,
		&identifier.Parking,
		&identifier.OfficeDays,
		&identifier.PreferredStartTime,
		&identifier.PreferredEndTime,
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
func (access *UserDA) StoreIdentifier(identifier *User) (string, error) {
	results, err := access.access.Query(
		`SELECT * FROM "user".identifier_store($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, mapString,
		identifier.Id, identifier.Identifier, identifier.FirstName, identifier.LastName, identifier.Email, identifier.Picture, identifier.WorkFromHome,
		identifier.Parking, identifier.OfficeDays, identifier.PreferredStartTime, identifier.PreferredEndTime)
	if err != nil {
		return "", err
	}
	for r, _ := range results {
		if value, ok := results[r].(string); ok {
			return value, nil
		}
	}
	return "", nil
}

//FindIdentifier finds an identifier
func (access *UserDA) FindIdentifier(identifier *User) (Users, error) {
	results, err := access.access.Query(
		`SELECT * FROM "user".identifier_find($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`, mapUser,
		identifier.Id, identifier.Identifier, identifier.FirstName, identifier.LastName, identifier.Email, identifier.Picture, identifier.DateCreated, identifier.WorkFromHome,
		identifier.Parking, identifier.OfficeDays, identifier.PreferredStartTime, identifier.PreferredEndTime)
	if err != nil {
		return nil, err
	}
	tmp := make([]*User, 0)
	for r, _ := range results {
		if value, ok := results[r].(User); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

// StoreCredential stores a credential
func (access *UserDA) StoreCredential(Id string, secret *string, identifier string) error {
	_, err := access.access.Query(`SELECT * FROM "user".credential_store($1, $2, $3)`, nil, Id, secret, identifier)
	if err != nil {
		return err
	}
	return nil
}

//FindCredential finds a user according to Credentials
func (access *UserDA) FindCredential(credential *Credential) (Users, error) {
	results, err := access.access.Query(
		`SELECT * FROM "user".credential_find($1, $2, $3, $4, $5, $6, $7)`, mapCredential,
		credential.Id, credential.Secret, credential.Identifier, nil, credential.Active, nil, credential.LastAccessed)
	if err != nil {
		return nil, err
	}
	tmp := make([]*User, 0)
	for r, _ := range results {
		if value, ok := results[r].(User); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//FindHead returns the first User
func (users Users) FindHead() *User {
	if len(users) == 0 {
		return nil
	}
	return users[0]
}
