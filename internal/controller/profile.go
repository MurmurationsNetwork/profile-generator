package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/bxcodec/faker/v3"
)

type Profile struct {
	LinkedSchemas []string `json:"linked_schemas"`
	Name          string   `json:"name"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
}

const (
	// Coordinates for a central point in France (e.g., Paris)
	centerLatitude  = 48.8566
	centerLongitude = 2.3522
	// Radius in kilometers
	radiusKm = 5.0
)

// randomCoordinate generates a random coordinate within a radius in km
func randomCoordinate(centerLat, centerLon, radiusKm float64) (float64, float64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	radiusDeg := radiusKm / 111.0 // Approx conversion from km to degrees

	latOffset := (r.Float64() - 0.5) * 2 * radiusDeg
	lonOffset := (r.Float64() - 0.5) * 2 * radiusDeg

	lat := centerLat + latOffset
	lon := centerLon + lonOffset

	return lat, lon
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	latitude, longitude := randomCoordinate(centerLatitude, centerLongitude, radiusKm)
	name := faker.Name()

	profile := Profile{
		LinkedSchemas: []string{"test_schema-v2.0.0"},
		Name:          name,
		Latitude:      latitude,
		Longitude:     longitude,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
