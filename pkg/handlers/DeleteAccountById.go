package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/focusfind/backend/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteAccountById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var account models.Account

	if result := h.DB.First(&account, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&account)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
