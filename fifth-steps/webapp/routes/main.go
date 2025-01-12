package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"webapp/db"
	"webapp/models"
)

var Templ = template.Must(template.ParseGlob("templates/*.html"))

func List(w http.ResponseWriter, r *http.Request) {
	dbConection, _ := db.ConectDB()
	defer dbConection.Close()
	products := models.ListProduct(dbConection)
	Templ.ExecuteTemplate(w, "Index", products)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		dbConection, _ := db.ConectDB()
		defer dbConection.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var tempProduct models.Product
		err = json.Unmarshal(body, &tempProduct)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		res, err := tempProduct.CreateProduct(dbConection)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, res)
	} else {
		Templ.ExecuteTemplate(w, "New", nil)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")

		dbConection, _ := db.ConectDB()
		defer dbConection.Close()
		conv, _ := strconv.Atoi(id)
		var product models.Product
		product.ID = conv

		res, err := product.DeleteProduct(dbConection)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, res)
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, "Wrong method")
}
