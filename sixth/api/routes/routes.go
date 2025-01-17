package routes

import (
	c "api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", c.Home)
	r.HandleFunc("/accounts", c.AllAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}", c.GetAccount).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
