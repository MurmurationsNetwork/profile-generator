package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// Constants for coordinates and radius
const (
	centerLatitude  = 48.8566
	centerLongitude = 2.3522
	radiusKm        = 5.0
)

// randomCoordinate generates a random coordinate within a radius in km
func randomCoordinate(centerLat, centerLon, radiusKm float64) (float64, float64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	radiusDeg := radiusKm / 111.0

	latOffset := (r.Float64() - 0.5) * 2 * radiusDeg
	lonOffset := (r.Float64() - 0.5) * 2 * radiusDeg

	lat := centerLat + latOffset
	lon := centerLon + lonOffset

	return lat, lon
}

func GetShortProfile(w http.ResponseWriter, r *http.Request) {
	latitude, longitude := randomCoordinate(centerLatitude, centerLongitude, radiusKm)

	profile := map[string]interface{}{
		"linked_schemas": []string{"media_schema-v0.1.0"},
		"title":          "A media article",
		"media_url":      "https://foo.bar",
		"author":         "The Dude",
		"tags":           []string{"test"},
		"description":    "A small profile for testing purposes",
		"name":           "A Small Profile",
		"timestamp":      time.Now().Unix(),
		"geolocation": map[string]float64{
			"lat": latitude,
			"lon": longitude,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func GetMediumProfile(w http.ResponseWriter, r *http.Request) {
	latitude, longitude := randomCoordinate(centerLatitude, centerLongitude, radiusKm)

	profile := map[string]interface{}{
		"linked_schemas": []string{"organizations_schema-v1.0.0"},
		"nickname":       "some project",
		"primary_url":    "https://dev.null",
		"tags":           []string{"test"},
		"urls": []map[string]string{
			{
				"name": "X",
				"url":  "https://x.com/test",
			},
		},
		"description":      "A medium sized profile for testing purposes.",
		"mission":          "A medium sized profile for testing purposes.",
		"status":           "on_hold",
		"full_address":     "A medium sized profile for testing purposes.",
		"country_iso_3166": "AF",
		"image":            "https://dev.null/logo.png",
		"header_image":     "https://dev.null/header.png",
		"images": []map[string]string{
			{
				"name": "Some image",
				"url":  "https://dev.null/image.png",
			},
		},
		"rss":       "https://dev.null/rss",
		"starts_at": 1818181818,
		"ends_at":   1919191919,
		"contact_details": map[string]string{
			"email": "dude@dev.null",
		},
		"telephone":        "+12125551212",
		"geographic_scope": "international",
		"unique_id":        "sdjsdiwej39fjskc",
		"name":             "A Medium Profile",
		"timestamp":        time.Now().Unix(),
		"geolocation": map[string]float64{
			"lat": latitude,
			"lon": longitude,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func GetLargeProfile(w http.ResponseWriter, r *http.Request) {
	latitude, longitude := randomCoordinate(centerLatitude, centerLongitude, radiusKm)

	profile := map[string]interface{}{
		"linked_schemas": []string{
			"permaculture_addon_schema-v1.0.0",
			"organizations_schema-v1.0.0",
		},
		"purpose": "A RESTful API service typically has numerous endpoints. No assumptions " +
			"should be made about the performance characteristics of one endpoint, by " +
			"testing another. This simple fact leads to the realization that every " +
			"endpoint should be tested with different assumptions, metrics and thresholds. " +
			"Starting with individual endpoints is a smart way to begin your API " +
			"performance testing." +
			"A RESTful API service typically has numerous endpoints. No assumptions " +
			"should be made about the performance characteristics of one endpoint, by " +
			"testing another. This simple fact leads to the realization that every " +
			"endpoint should be tested with different assumptions, metrics and thresholds. " +
			"Starting with individual endpoints is a smart way to begin your API " +
			"performance testing." +
			"A RESTful API service typically has numerous endpoints. No assumptions " +
			"should be made about the performance characteristics of one endpoint, by " +
			"testing another. This simple fact leads to the realization that every " +
			"endpoint should be tested with different assumptions, metrics and thresholds. " +
			"Starting with individual endpoints is a smart way to begin your API " +
			"performance testing." +
			"A RESTful API service typically has numerous endpoints. No assumptions " +
			"should be made about the performance characteristics of one endpoint, by " +
			"testing another. This simple fact leads to the realization that every " +
			"endpoint should be tested with different assumptions, metrics and thresholds. " +
			"Starting with individual endpoints is a smart way to begin your API " +
			"performance testing." +
			"A RESTful API service typically has numerous endpoints. No assumptions " +
			"should be made about the performance characteristics of one endpoint, by " +
			"testing another. This simple fact leads to the realization that every " +
			"endpoint should be tested with different assumptions, metrics and thresholds. " +
			"Starting with individual endpoints is a smart way to begin your API " +
			"performance testing." +
			"A RESTful API service typically has numerous endpoints. No assumptions " +
			"should be made about the performance characteristics of one endpoint, by " +
			"testing another. This simple fact leads to the realization that every " +
			"endpoint should be tested with different assumptions, metrics and thresholds. " +
			"Starting with individual endpoints is a smart way to begin your API " +
			"performance testing.",
		"date_added":                   1919191919,
		"org_type":                     []string{"type of org"},
		"number_of_members":            23920,
		"number_of_beneficiaries":      2320392309,
		"geographic_scope_description": "dkfjsdkf",
		"leadership_tags":              []string{"sdsdas"},
		"permaculture_domains":         []string{"sdfsdf"},
		"nickname":                     "xfsdfdsf",
		"primary_url":                  "https://ic3.dev/person1",
		"description": "A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description",
		"mission": "A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description",
		"status": "completed",
		"full_address": "A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description A long description A long " +
			"description A long description A long description",
		"country_iso_3166": "BF",
		"image":            "https://dev.null/logo.png",
		"header_image":     "https://dev.null/header.png",
		"images": []map[string]string{
			{
				"name": "Some image",
				"url":  "https://dev.null/image.png",
			},
		},
		"rss":       "https://dev.null/rss",
		"starts_at": 1020102,
		"ends_at":   34023233232,
		"contact_details": map[string]string{
			"email": "dude@dev.null",
		},
		"telephone":        "+12125551212",
		"geographic_scope": "regional",
		"unique_id": "The unique identifier of the entity, optionally defined by the group " +
			"recording/tracking the entity The unique identifier of the entity, optionally " +
			"defined by the group recording/tracking the entity The unique identifier of the " +
			"entity, optionally defined by the group recording/tracking the entity The unique " +
			"identifier of the entity, optionally defined by the group recording/tracking the " +
			"entity The unique identifier of the entity, optionally defined by the group " +
			"recording/tracking the entity The unique identifier of the entity, optionally " +
			"defined by the group recording/tracking the entity The unique identifier of the " +
			"entity, optionally defined by the group recording/tracking the entity The unique " +
			"identifier of the entity, optionally defined by the group recording/tracking the " +
			"entity The unique identifier of the entity, optionally defined by the group " +
			"recording/tracking the entity The unique identifier of the entity, optionally " +
			"defined by the group recording/tracking the entity The unique identifier of the " +
			"entity, optionally defined by the group recording/tracking the entity The unique " +
			"identifier of the entity, optionally defined by the group recording/tracking the " +
			"entity The unique identifier of the entity, optionally defined by the group " +
			"recording/tracking the entity",
		"name":      "A Large Profile",
		"timestamp": time.Now().Unix(),
		"geolocation": map[string]float64{
			"lat": latitude,
			"lon": longitude,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
