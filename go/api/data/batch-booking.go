package data

import (
	"api/db"
)

type BatchBooking struct {
	UserId   *string  `json:"user_id,omitempty"`
	Bookings Bookings `json:"bookings,omitempty"`
}

type BatchBookingDA struct {
	access *db.Access
}

// Creates DA from passed in access
func NewBatchBookingDA(access *db.Access) *BatchBookingDA {
	return &BatchBookingDA{
		access: access,
	}
}

// Commits transaction
func (access *BatchBookingDA) Commit() error {
	return access.access.Commit()
}

//////////////////////////////////////////////////
// Functions

// StoreIdentifiers stores multiple Bookings
// Look into more efficient implementation, less traffic to db
func (access *BatchBookingDA) StoreIdentifiers(identifiers *BatchBooking) error {
	bookingDA := NewBookingDA(access.access)
	for _, booking := range identifiers.Bookings {
		err := bookingDA.StoreIdentifier(booking)
		if err != nil {
			return err
		}
	}
	return nil
}
