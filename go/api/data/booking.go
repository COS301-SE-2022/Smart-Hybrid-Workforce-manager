package data

import (
	"api/db"
	"database/sql"
	"encoding/json"
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
	ResourceId           *string    `json:"resource_id,omitempty"`
	Start                *time.Time `json:"start,omitempty"`
	End                  *time.Time `json:"end,omitempty"`
	Booked               *bool      `json:"booked,omitempty"`
	DateCreated          *time.Time `json:"date_created,omitempty"`
}

// Bookings represent a splice of Booking
type Bookings []*Booking

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

// Commit commits the current implicit transaction
func (access *BookingDA) Commit() error {
	return access.access.Commit()
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
		&identifier.ResourceId,
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
		`SELECT 1 FROM booking.identifier_store($1, $2, $3, $4, $5, $6, $7, $8)`, nil,
		identifier.Id, identifier.UserId, identifier.ResourceType, identifier.ResourcePreferenceId, identifier.ResourceId, identifier.Start, identifier.End, identifier.Booked)
	if err != nil {
		return err
	}
	return nil
}

//FindIdentifier finds an identifier
func (access *BookingDA) FindIdentifier(identifier *Booking, permissions *Permissions) (Bookings, error) {
	content, err := json.Marshal(*permissions)
	if err != nil {
		return nil, err
	}
	results, err := access.access.Query(
		`SELECT * FROM booking.identifier_find($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, mapBooking,
		identifier.Id, identifier.UserId, identifier.ResourceType, identifier.ResourcePreferenceId,
		identifier.ResourceId, identifier.Start, identifier.End, identifier.Booked, identifier.DateCreated, content)
	if err != nil {
		return nil, err
	}
	tmp := make([]*Booking, 0)
	for r, _ := range results {
		if value, ok := results[r].(Booking); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//DeleteIdentifier finds an identifier
func (access *BookingDA) DeleteIdentifier(identifier *Booking) (*Booking, error) {
	results, err := access.access.Query(
		`SELECT * FROM booking.identifier_remove($1)`, mapBooking,
		identifier.Id)
	if err != nil {
		return nil, err
	}
	var tmp Bookings
	tmp = make([]*Booking, 0)
	for r, _ := range results {
		if value, ok := results[r].(Booking); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp.FindHead(), nil
}

//FindHead returns the first Booking
func (bookings Bookings) FindHead() *Booking {
	if len(bookings) == 0 {
		return nil
	}
	return bookings[0]
}
