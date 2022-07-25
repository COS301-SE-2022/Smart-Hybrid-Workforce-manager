package scheduler

import (
	"api/data"
	"api/db"
)

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

// func getBookings(from time.Time, to time.Time) (data.Bookings, error) {
// 	// Unmarhsal bookings
// 	var bookings data.BatchBooking
// 	// Connect to db
// 	access, err := db.Open()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer access.Close()

// 	// TODO @JonathanEnslin do other necessary checks

// 	da := data.NewBatchBookingDA(access)
// 	bookingsInfo, err := da.FindIdentifiers(&bookings, security.RemoveRolePermissions(permissions))
// 	if err != nil {
// 		return nil, err
// 	}

// 	// commit transaction
// 	err = access.Commit()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bookingsInfo, nil
// }
