package scheduler

import (
	"api/data"
	"api/db"
	"api/google_api"
	"lib/logger"
)

type CandidateBookings []data.Bookings

// makeBookings stores the created bookings in the database
func makeBookings(candidates CandidateBookings, schedulerData *SchedulerData) error {
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
	go func() { // Make bookings
		accessGoogle, err := db.Open()
		if err != nil {
			logger.Error.Print("Could not create access: ", err)
			return
		}
		defer accessGoogle.Close()
		for _, booking := range candidates[choose] {
			if booking.ResourceId != nil {
				du := data.NewUserDA(accessGoogle)
				users, err := du.FindIdentifier(&data.User{Id: booking.UserId})
				if err != nil {
					logger.Error.Printf("User not found when creating calendar booking. User ID: %v, err: %v\n", *booking.UserId, err)
					continue
				}
				user := users.FindHead()
				err = google_api.CreateUpdateBooking(user, booking)
				if err != nil {
					logger.Error.Println("User not found when creating calendar booking")
					continue
				}
				err = du.Commit()
				if err != nil {
					logger.Error.Printf("Could not commit %v", err)
				}
			}
		}
	}()
	return nil
}
