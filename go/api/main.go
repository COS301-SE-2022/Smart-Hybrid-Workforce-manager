package main

import (
	"fmt"
	"lib/logger"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You got the Arche api to work congratulations!")
	logger.Info.Println("Endpoint Hit: arche homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
