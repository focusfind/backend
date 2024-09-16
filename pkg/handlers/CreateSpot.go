package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/focusfind/backend/pkg/models"
)

func (h handler) CreateSpot(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var spot models.Spot
	json.Unmarshal(body, &spot)

	// Append to Spots table
	if result := h.DB.Create(&spot); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send response 201 Created
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
