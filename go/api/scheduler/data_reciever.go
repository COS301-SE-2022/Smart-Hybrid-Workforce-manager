package scheduler

import (
	"api/data"

	"api/db"
)

type CandidateBookings []data.Bookings

// makeBookings stores the created bookings in the database
func makeBookings(candidates CandidateBookings) error {
	choose := 0 // Right now just choose any set of bookings, no heuristic yet
	if len(candidates) == 0 {
		return nil // No bookings to be made
	}
	access, err := db.Open()
	if err != nil {
		return err
	}
	defer access.Close()

	bookings := data.BatchBooking{
		UserId:   nil,
		Bookings: candidates[choose],
	}

	da := data.NewBatchBookingDA(access)
	err = da.StoreIdentifiers(&bookings)
	if err != nil {
		return err
	}

	err = access.Commit()
	if err != nil {
		return err
	}
	return nil
}
