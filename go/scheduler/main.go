package main

import (
	"lib/logger"
	"net/http"
	"os"
	"scheduler/endpoints"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Route endpoints
	router := mux.NewRouter().StrictSlash(true)

	// Scheduler endpoints
	schedulerRouter := router.PathPrefix("/api/scheduler").Subrouter()
	err := endpoints.SchedulerHandlers(schedulerRouter)
	if err != nil {
		logger.Error.Fatal(err)
		os.Exit(-1)
	}

	// Start API on port 8080 in its docker container
	logger.Info.Println("Starting API on 8080")
	server := http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	logger.Error.Fatal(server.ListenAndServe())
}
