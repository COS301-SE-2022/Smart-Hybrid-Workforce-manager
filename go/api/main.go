package main

import (
	"api/db"
	"api/endpoints"
	"lib/logger"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	// Create Database connection pool
	err := db.RegisterAccess()
	if err != nil {
		logger.Error.Fatal(err)
		os.Exit(-1)
	}

	// Route endpoints
	router := mux.NewRouter().StrictSlash(true)

	// User endpoints
	userRouter := router.PathPrefix("/api/user").Subrouter()
	err = endpoints.UserHandlers(userRouter)
	if err != nil {
		logger.Error.Fatal(err)
		os.Exit(-1)
	}

	// Team endpoints
	teamRouter := router.PathPrefix("/api/team").Subrouter()
	err = endpoints.TeamHandlers(teamRouter)
	if err != nil {
		logger.Error.Fatal(err)
		os.Exit(-1)
	}

	// Booking endpoints
	bookingRouter := router.PathPrefix("/api/booking").Subrouter()
	err = endpoints.BookingHandlers(bookingRouter)
	if err != nil {
		logger.Error.Fatal(err)
		os.Exit(-1)
	}

	//Notification endpoints
	notificationRouter := router.PathPrefix("/api/notification").Subrouter()
	err = endpoints.NotificationHandlers(notificationRouter)
	if err != nil {
		logger.Error.Fatal(err)
		os.Exit(-1)
	}

	// Resource endpoints
	resourceRouter := router.PathPrefix("/api/resource").Subrouter()
	err = endpoints.ResourceHandlers(resourceRouter)
	if err != nil {
		logger.Error.Fatal(err)
		os.Exit(-1)
	}

	// Role endpoints
	roleRouter := router.PathPrefix("/api/role").Subrouter()
	err = endpoints.RoleHandlers(roleRouter)
	if err != nil {
		logger.Error.Fatal(err)
		os.Exit(-1)
	}

	// Permission endpoints
	permissionRouter := router.PathPrefix("/api/permission").Subrouter()
	err = endpoints.PermissionHandlers(permissionRouter)
	if err != nil {
		logger.Error.Fatal(err)
		os.Exit(-1)
	}

	batchBookingRouter := router.PathPrefix("/api/batch-booking").Subrouter()
	err = endpoints.BatchBookingHandlers(batchBookingRouter)
	if err != nil {
		logger.Error.Fatal(err)
		os.Exit(-1)
	}

	// Start API on port 8080 in its docker container
	logger.Info.Println("Starting API on 8080")
	logger.Error.Fatal(http.ListenAndServe(":8080", router))
}
