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

func (h handler) UpdateAccountById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var updatedAccount models.Account
	json.Unmarshal(body, &updatedAccount)

	var account models.Account
	if result := h.DB.First(&account, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	account.Name = updatedAccount.Name
	account.Email = updatedAccount.Email
	account.Password = updatedAccount.Password

	h.DB.Save(&account)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("Updated account %d", id))
}
