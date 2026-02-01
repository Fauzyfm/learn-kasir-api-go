package handler

import (
	"encoding/json"
	"kasir-api/internal/models"
	"kasir-api/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}


func (h *ProductHandler) HandleProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			h.GetAllProducts(w, r)
		case http.MethodPost:
			h.CreateProduct(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
}



func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	product, err := h.service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}


func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}	

    if product.NamaBarang == "" {
        http.Error(w, "nama_barang is required", http.StatusBadRequest)
        return
    }
    
    if product.HargaBarang <= 0 {
        http.Error(w, "harga_barang must be greater than 0", http.StatusBadRequest)
        return
    }
    
    if product.CategoryID <= 0 {
        http.Error(w, "category_id is required", http.StatusBadRequest)
        return
    }


	err = h.service.CreateProduct(&product)
	if err != nil {	
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}


func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			h.GetProductByID(w, r)
		case http.MethodPut:
			h.UpdateProduct(w, r)
		case http.MethodDelete:
			h.DeleteProduct(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			
	}
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w  http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updateProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	newProduct, err := h.service.UpdateProduct(id, updateProduct)
	if err != nil {
		http.Error(w, "Product Not Found", http.StatusBadRequest)
		return	
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)

}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	

	succes := h.service.DeleteProduct(id)
	if succes != nil {
		http.Error(w, "Product Not Found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product deleted successfully",
	})
}


func (h *ProductHandler) GetProductDetail(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/api/products/detail/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }
    
    product, err := h.service.GetProductByIDWithCategory(id)
    if err != nil {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}

// Health - handler untuk GET /health
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "oke",
		"message": "API Running",
	})
}


