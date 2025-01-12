package main

import (
	"database/sql"
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func conectDB() (*sql.DB, error) {
	conectionUrl := "user=postgres dbname=goproducts password=123456 hostname=localhost sslmode=disable"

	db, err := sql.Open("postgres", conectionUrl)
	return db, err
}

var templ = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db, _ := conectDB()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{
			Name:        "Laptop",
			Description: "Asus",
			Price:       1000,
			Quantity:    10,
		},
		{
			Name:        "Mobile",
			Description: "Samsung",
			Price:       500,
			Quantity:    20,
		},
		{
			Name:        "Tablet",
			Description: "Apple",
			Price:       800,
			Quantity:    15,
		},
	}
	templ.ExecuteTemplate(w, "Index", products)
}
