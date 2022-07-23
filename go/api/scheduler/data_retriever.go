package scheduler

import (
	"api/data"
	"api/db"
	"api/security"
	"time"
)

// getUsers Retrieves all the users from the database
func getUsers() (data.Users, error) {
	access, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer access.Close()

	da := data.NewUserDA(access)

	user := data.User{}
	users, err := da.FindIdentifier(&user)
	if err != nil {
		return nil, err
	}

	err = access.Commit()
	if err != nil {
		return nil, err
	}
	// TODO: @JonathanEnslin remove unnecessary fields from users
	return users, nil
}

// getBookings Retrieves all the bookings from the database between from from to to
func getBookings(from time.Time, to time.Time) (data.Bookings, error) {
	permissions := &data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")}
	// Connect to db
	access, err := db.Open() // TODO: @JonathanEnslin
	if err != nil {
		return nil, err
	}
	defer access.Close()

	bookingFilter := data.Booking{
		Start: &from,
		End:   &to,
	}

	da := data.NewBookingDA(access)
	bookings, err := da.FindIdentifier(&bookingFilter, security.RemoveRolePermissions(permissions))
	if err != nil {
		return nil, err
	}

	// commit transaction
	err = access.Commit()
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
