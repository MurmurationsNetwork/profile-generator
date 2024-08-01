package controller

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	LinkedSchemas []string `json:"linked_schemas"`
	Name          string   `json:"name"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	profile := Profile{
		LinkedSchemas: []string{"test_schema-v2.0.0"},
		Name:          "IC3 Dev",
		Latitude:      11.111111,
		Longitude:     12.121212,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
