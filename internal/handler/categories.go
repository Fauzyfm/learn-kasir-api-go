package handler

import (
	"encoding/json"
	"kasir-api/internal/models"
	"kasir-api/internal/service"
	"net/http"
	"strconv"
	"strings"
)

// GetAllCategories - handler untuk GET /api/categories
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	Categories := service.GetAllCategories()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Categories)
}

// CreateCategories - handler untuk POST /api/categories
func CreateCategories(w http.ResponseWriter, r *http.Request) {
	var categorie models.Categories
	err := json.NewDecoder(r.Body).Decode(&categorie)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	newCategories := service.CreateCategories(categorie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategories)
}

// GetCategorieByID - handler untuk GET /api/categorie/{id}
func GetCategorieByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categorie/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	categorie, _ := service.GetCategorieByID(id)
	if categorie == nil {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categorie)
}

// UpdateCategorie - handler untuk PUT /api/categorie/{id}
func UpdateCategorie(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categorie/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var categorie models.Categories
	err = json.NewDecoder(r.Body).Decode(&categorie)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	updateCategorie, _ := service.UpdateCategorie(id, categorie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updateCategorie)

}

// DeleteCategorie - handler untuk DELETE /api/categorie/{id}
func DeleteCategorie(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categorie/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	success := service.DeleteCategorie(id)
	if !success {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product deleted successfully",
	})
}