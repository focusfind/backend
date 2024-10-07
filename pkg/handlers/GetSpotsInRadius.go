package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/focusfind/backend/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetSpotsInRadius(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	latitude, _ := strconv.Atoi(vars["lat"])
	longitude, _ := strconv.Atoi(vars["long"])
	radius, _ := strconv.Atoi(vars["radius"])

	var spots []models.Spot
	h.DB.Where("ST_DWithin(coordinates, ST_SetSRID(ST_MakePoint(?, ?), 4326), ?)", longitude, latitude, radius).Find(&spots)

	fmt.Println(spots)
}
