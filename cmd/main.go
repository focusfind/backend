package main

import (
	"log"
	"net/http"

	"github.com/focusfind/backend/pkg/db"
	"github.com/focusfind/backend/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	db := db.Init()
	h := handlers.New(db)
	router := mux.NewRouter()

	// Spots CRUD
	router.HandleFunc("/spots", h.ListSpots).Methods(http.MethodGet)
	router.HandleFunc("/spots/{id}", h.GetSpotById).Methods(http.MethodGet)
	router.HandleFunc("/spots", h.CreateSpot).Methods(http.MethodPost)
	router.HandleFunc("/spots/{id}", h.DeleteSpotById).Methods(http.MethodDelete)
	router.HandleFunc("/spots/{id}", h.UpdateSpotById).Methods(http.MethodPut)

	// Account CRUD
	router.HandleFunc("/accounts/", h.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{id}", h.GetAccountById).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{id}", h.DeleteAccountById).Methods(http.MethodDelete)
	router.HandleFunc("/accounts/{id}", h.UpdateAccountById).Methods(http.MethodPut)

	log.Println("API is running!")
	http.ListenAndServe(":42069", router)
}
