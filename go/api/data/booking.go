package data

import (
	"api/db"
	"database/sql"
	"time"
)

//////////////////////////////////////////////////
// Structures and Variables

// Booking identifies a booking via common attributes
type Booking struct {
	Id                   *string    `json:"id,omitempty"`
	UserId               *string    `json:"user_id,omitempty"`
	ResourceType         *string    `json:"resource_type,omitempty"`
	ResourcePreferenceId *string    `json:"resource_preference_id,omitempty"`
	Start                *time.Time `json:"start,omitempty"`
	End                  *time.Time `json:"end,omitempty"`
	Booked               *bool      `json:"booked,omitempty"`
	DateCreated          *time.Time `json:"date_created,omitempty"`
}

// BookingDA provides access to the database for authentication purposes
type BookingDA struct {
	access *db.Access
}

// NewBookingDA creates a new data access from a underlying shared data access
func NewBookingDA(access *db.Access) *BookingDA {
	return &BookingDA{
		access: access,
	}
}

//////////////////////////////////////////////////
// Mappers

func mapBooking(rows *sql.Rows) (interface{}, error) {
	var identifier Booking
	err := rows.Scan(
		&identifier.Id,
		&identifier.UserId,
		&identifier.ResourceType,
		&identifier.ResourcePreferenceId,
		&identifier.Start,
		&identifier.End,
		&identifier.Booked,
		&identifier.DateCreated,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

//////////////////////////////////////////////////
// Functions

// StoreIdentifier stores an identifier
func (access *BookingDA) StoreIdentifier(identifier *Booking) error {
	_, err := access.access.Query(
		`SELECT 1 FROM booking.identifier_store($1, $2, $3, $4, $5, $6, $7)`, nil,
		identifier.Id, identifier.UserId, identifier.ResourceType, identifier.ResourcePreferenceId, identifier.Start, identifier.End, identifier.Booked)
	if err != nil {
		return err
	}
	return nil
}

//FindIdentifier finds an identifier
func (access *BookingDA) FindIdentifier(identifier *Booking) (*Booking, error) {
	results, err := access.access.Query(
		`SELECT * FROM booking.identifier_find($1, $2, $3, $4, $5, $6, $7, $8)`, mapUser,
		identifier.Id, identifier.UserId, identifier.ResourceType, identifier.ResourcePreferenceId, identifier.Start, identifier.End, identifier.Booked, identifier.DateCreated)
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		if identifier, ok := result.(Booking); ok {
			return &identifier, nil
		}
	}
	return nil, nil
}
