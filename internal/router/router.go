package router

import (
	"net/http"

	"kasir-api/internal/handler"
)

// SetupRoutes - mengatur semua routes untuk API
func SetupRoutes() {
	// GET & PUT & DELETE /api/product/{id}
	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handler.GetProductByID(w, r)
		} else if r.Method == "PUT" {
			handler.UpdateProduct(w, r)
		} else if r.Method == "DELETE" {
			handler.DeleteProduct(w, r)
		}
	})

	// POST & GET /api/products
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handler.CreateProduct(w, r)
		} else if r.Method == "GET" {
			handler.GetAllProducts(w, r)
		}
	})

	// GET & POST /api/categories
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handler.GetAllCategories(w, r)
		} else if r.Method == "POST" {
			handler.CreateCategories(w, r)
		}
	})

	// GET & PUT & DELETE /api/categorie/{id}
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handler.GetCategorieByID(w, r)
		} else if r.Method == "PUT" {
			handler.UpdateCategorie(w, r)
		} else if r.Method == "DELETE" {
			handler.DeleteCategorie(w, r)
		}

	})

	// GET /health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		handler.Health(w, r)
	})
}
