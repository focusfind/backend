package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/focusfind/backend/pkg/models"
)

func (h handler) ListSpots(w http.ResponseWriter, r *http.Request) {
	var spots []models.Spot
	if result := h.DB.Find(&spots); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(spots)
}
