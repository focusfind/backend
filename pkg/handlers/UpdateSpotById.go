package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

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
	}

	spot.Name = updatedSpot.Name
	spot.Type = updatedSpot.Type
	spot.Location = updatedSpot.Location
	spot.Description = updatedSpot.Description
	spot.BusyIndex = updatedSpot.BusyIndex

	h.DB.Save(&spot)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("Updated spot %d", id))
}
