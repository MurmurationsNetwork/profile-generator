package main

import (
	"log"
	"net/http"

	"github.com/MurmurationsNetwork/profile-generator/internal/controller"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/profile/small", controller.GetShortProfile).Methods("GET")
	router.HandleFunc("/profile/medium", controller.GetMediumProfile).Methods("GET")
	router.HandleFunc("/profile/large", controller.GetLargeProfile).Methods("GET")

	router.HandleFunc("/profile/small/{rest:.*}", controller.GetShortProfile).Methods("GET")
	router.HandleFunc("/profile/medium/{rest:.*}", controller.GetMediumProfile).Methods("GET")
	router.HandleFunc("/profile/large/{rest:.*}", controller.GetLargeProfile).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
