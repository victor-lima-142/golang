package controllers

import (
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllAccounts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Accounts)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, acc := range models.Accounts {
		if acc.ID == id {
			json.NewEncoder(w).Encode(acc)
			return
		}
	}
	json.NewEncoder(w).Encode([]models.Account{})
}
