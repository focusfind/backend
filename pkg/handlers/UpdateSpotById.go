package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/focusfind/backend/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) UpdateSpotById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var updatedSpot models.Spot
	json.Unmarshal(body, &updatedSpot)

	var spot models.Spot
	if result := h.DB.First(&spot, id); result.Error != nil {
		fmt.Println(result.Error)
		http.Error(w, "Spot not found", http.StatusNotFound)
		return
	}

	spot.Name = updatedSpot.Name
	spot.Type = updatedSpot.Type
	spot.Coordinates = updatedSpot.Coordinates
	spot.Description = updatedSpot.Description
	spot.BusyIndex = updatedSpot.BusyIndex

	result := h.DB.Save(&spot)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
			http.Error(w, "Duplicate entry for coordinates", http.StatusBadRequest)
			return
		} else if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
			http.Error(w, "Duplicate entry for name", http.StatusBadRequest)
			return
		}
		log.Println("Error updating spot:", result.Error)
		http.Error(w, "Failed to update spot", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("Updated spot %d", id))
}
