package router

import (
	"net/http"

	"kasir-api/internal/handler"
)

// SetupRoutes - mengatur semua routes untuk API
func SetupRoutes(productHandler *handler.ProductHandler, categoryHandler *handler.CategoryHandler, transactionHandler *handler.TransactionHandler, reportHandler *handler.ReportHandler) {
	// GET & PUT & DELETE /api/product/{id}
	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			productHandler.HandleProductByID(w, r)
		} else if r.Method == "PUT" {
			productHandler.HandleProductByID(w, r)
		} else if r.Method == "DELETE" {
			productHandler.HandleProductByID(w, r)
		}
	})

	// POST & GET /api/products
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			productHandler.HandleProduct(w, r)
		} else if r.Method == "POST" {
			productHandler.HandleProduct(w, r)
		}
	})

	// // GET & POST /api/categories
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			categoryHandler.HandleCategory(w, r)
		} else if r.Method == "POST" {
			categoryHandler.HandleCategory(w, r)
		}
	})

	// // GET & PUT & DELETE /api/categorie/{id}
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			categoryHandler.HandleCategoryByID(w, r)
		} else if r.Method == "PUT" {
			categoryHandler.HandleCategoryByID(w, r)
		} else if r.Method == "DELETE" {
			categoryHandler.HandleCategoryByID(w, r)
		}

	})

	http.HandleFunc("/api/products/detail/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			productHandler.GetProductDetail(w, r)
		}
	})

	// POST /api/checkout
	http.HandleFunc("/api/checkout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			transactionHandler.HandleCheckout(w, r)
		}
	})

	// GET /api/report
	http.HandleFunc("/api/report", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			reportHandler.HandleReport(w, r)
		}
	})

	// GET /api/report/hari-ini
	http.HandleFunc("/api/report/hari-ini", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			reportHandler.HandleReport(w, r)
		}
	})
}
