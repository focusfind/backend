# focusfind backend

## Spots CRUD

```go
router.HandleFunc("/spots", h.ListSpots).Methods(http.MethodGet)
router.HandleFunc("/spots/{id}", h.GetSpotById).Methods(http.MethodGet)
router.HandleFunc("/spots", h.CreateSpot).Methods(http.MethodPost)
router.HandleFunc("/spots/{id}", h.DeleteSpotById).Methods(http.MethodDelete)
router.HandleFunc("/spots/{id}", h.UpdateSpotById).Methods(http.MethodPut)
```

## Account CRUD

```go
router.HandleFunc("/accounts/", h.CreateAccount).Methods(http.MethodPost)
router.HandleFunc("/accounts/{id}", h.GetAccountById).Methods(http.MethodGet)
router.HandleFunc("/accounts/{id}", h.DeleteAccountById).Methods(http.MethodDelete)
router.HandleFunc("/accounts/{id}", h.UpdateAccountById).Methods(http.MethodPut)
```
