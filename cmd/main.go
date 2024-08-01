package main

import (
	"log"
	"net/http"
	"profile-generator/internal/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profile", controller.GetProfile).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
