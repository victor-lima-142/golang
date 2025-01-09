package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
