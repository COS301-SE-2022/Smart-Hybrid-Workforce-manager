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
	err = endpoints.RegisterUserHandlers(userRouter)
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

	// Start API on port 8080 in its docker container
	logger.Info.Println("Starting API on 8080")
	logger.Error.Fatal(http.ListenAndServe(":8080", router))
}
