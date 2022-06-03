package data

import (
	"api/db"
	"encoding/json"
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


// Finds and returns bookings as specified by the identifiers in the Bookings array
func (access *BatchBookingDA) FindIdentifiers(indentifiers *BatchBooking, permissions *Permissions) (Bookings, error) {
	permissionContent, err := json.Marshal(*permissions)
	if err != nil {
		return nil, err
	}
	allResults := make([]interface{}, 0)
	// Get the results for each booking
	// Could update to just use a BookingDA
	for _, identifier := range indentifiers.Bookings {
		results, err := access.access.Query(
			`SELECT * FROM booking.identifier_find($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, mapBooking,
			identifier.Id, identifier.UserId, identifier.ResourceType, identifier.ResourcePreferenceId,
			identifier.ResourceId, identifier.Start, identifier.End, identifier.Booked, identifier.DateCreated, permissionContent,
		)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, results...)
	}
	tmp := make([]*Booking, 0)
	for r := range allResults {
		if val, ok := allResults[r].(Booking); ok {
			tmp = append(tmp, &val)
		}
	}
	// filter out duplicates (incase)
	idMap := make(map[string]bool)
	bookingSet := make([]*Booking, 0)
	for _, booking := range tmp {
		// If booking not yet added
		if _, added := idMap[*booking.Id]; !added {
			bookingSet = append(bookingSet, booking) // add booking to set
			idMap[*booking.Id] = true
		}
	}
	return bookingSet, nil
}

// Delete identifiers deletes (and returns) multiple bookings
func (access *BatchBookingDA) DeleteIdentifiers(identifiers *BatchBooking) (Bookings, error) {
	bookingDA := NewBookingDA(access.access)
	allDeleted := make([]*Booking, 0)
	for _, booking := range identifiers.Bookings {
		result, err := bookingDA.DeleteIdentifier(booking)
		if err != nil {
			return nil, err
		}
		allDeleted = append(allDeleted, result)
	}
	return allDeleted, nil
}
