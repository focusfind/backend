package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/focusfind/backend/pkg/models"
)

const minDistanceMeters = 100

func (h handler) CreateSpot(w http.ResponseWriter, r *http.Request) {
	var spot models.Spot
	err := json.NewDecoder(r.Body).Decode(&spot)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Check for nearby spots
	var nearbySpots []models.Spot
	result := h.DB.Select("id, name, type, ST_AsEWKB(coordinates) as coordinates, description, busy_index").
		Where("ST_DWithin(coordinates::geography, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?)",
			spot.Coordinates.Longitude, spot.Coordinates.Latitude, minDistanceMeters).
		Find(&nearbySpots)

	if result.Error != nil {
		log.Println("Error checking nearby spots:", result.Error)
		http.Error(w, "Failed to check nearby spots", http.StatusInternalServerError)
		return
	}

	if len(nearbySpots) > 0 {
		http.Error(w, "A spot already exists within 100 meters of this location", http.StatusBadRequest)
		return
	}

	// Append to Spots table
	result = h.DB.Create(&spot)
	if result.Error != nil {
		// Check for unique constraint violations
		if strings.Contains(result.Error.Error(), "idx_coordinates") {
			http.Error(w, "Duplicate entry for coordinates", http.StatusBadRequest)
			return
		} else if strings.Contains(result.Error.Error(), "uni_spots_name") {
			http.Error(w, "Duplicate entry for name", http.StatusBadRequest)
			return
		}
		// Handle other database errors
		log.Println("Error creating spot:", result.Error)
		http.Error(w, "Failed to create spot", http.StatusInternalServerError)
		return
	}

	// Send response 201 Created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
