package routes

import (
	c "api/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", c.Home)
	http.HandleFunc("/all", c.AllAccounts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
