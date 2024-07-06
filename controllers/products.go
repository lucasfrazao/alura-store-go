package controllers

import (
	"alura-store/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	tmp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price:", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity:", err)
		}

		models.CreateNewProduct(name, description, priceConverted, quantityConverted)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.ProductById(productId)

	tmp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConverted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error converting id:", err)
		}

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price:", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity:", err)
		}

		models.UpdateProduct(idConverted, quantityConverted, name, description, priceConverted)
	}
	http.Redirect(w, r, "/", 301)
}
