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

	userRouter := router.PathPrefix("/api/user").Subrouter()
	err = endpoints.RegisterUserHandlers(userRouter) // registers endpoints/user.go
	if err != nil {
		logger.Error.Fatal(err)
		os.Exit(-1)
	}

	// Start API on port 8080 in its docker container
	logger.Info.Println("Starting API on 8080")
	logger.Error.Fatal(http.ListenAndServe(":8080", router))
}
