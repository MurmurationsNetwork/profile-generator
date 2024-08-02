package main

import (
	"log"
	"net/http"

	"github.com/MurmurationsNetwork/profile-generator/internal/controller"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profile", controller.GetProfile).Methods("GET")
	router.HandleFunc("/profile/{rest:.*}", controller.GetProfile).Methods("GET") // Wildcard for any path starting with /profile

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
