package main

import (
	"net/http"
	"webapp/routes"
)

func main() {
	http.HandleFunc("/", routes.List)
	http.HandleFunc("/new", routes.Create)
	http.HandleFunc("/delete", routes.Delete)
	http.ListenAndServe(":8080", nil)
}
