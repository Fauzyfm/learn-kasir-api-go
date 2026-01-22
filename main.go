package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID int `json:"id"`
	NamaBarang string `json:"nama_barang"`
	HargaBarang int `json:"harga_barang"`
	Stok int `json:"stok"`
}

var product = []Product{
	{ID: 1, NamaBarang: "Pensil", HargaBarang: 2000, Stok: 100},
	{ID: 2, NamaBarang: "Buku Tulis", HargaBarang: 5000, Stok: 200},
	{ID: 3, NamaBarang: "Penghapus", HargaBarang: 1500, Stok: 150},
}


func getProductByID(w http.ResponseWriter, r *http.Request){
		idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Product ID", http.StatusBadRequest)
			return
		}

		for _, p := range product {
			if p.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(p)
				return
			}
		}

		http.Error(w, "Product Not Found", http.StatusNotFound)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Product ID", http.StatusBadRequest)
			return
		}

		var aptProduct Product
		err = json.NewDecoder(r.Body).Decode(&aptProduct)
		if err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		for i := range product {
			if product[i].ID == id {
				aptProduct.ID = id
				product[i] = aptProduct

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(aptProduct)
				return
			}
		}

		http.Error(w, "Product Not Found", http.StatusNotFound)

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Product ID", http.StatusBadRequest)
			return
		}

		for i := range product {
			if product[i].ID == id {
				product = append(product[:i],  product[i+1:]...)

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Product deleted successfully",
				})
				return
			}

		}

		http.Error(w, "Product Not Found", http.StatusNotFound)


}

func getNextID() int {
    maxID := 0
    for _, p := range product {
        if p.ID > maxID {
            maxID = p.ID
        }
    }
    return maxID + 1
}
func main() {

	// GET & PUT api/product/{id}
	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProductByID(w, r)
		} else if r.Method == "PUT" {
			updateProduct(w,r)
		} else if r.Method == "DELETE" {
			deleteProduct(w, r)
		}
	})


	// POST & GET api/products
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request){

		// POST
		if r.Method == "POST" {
			var NewProduct Product
			err := json.NewDecoder(r.Body).Decode(&NewProduct)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			NewProduct.ID = getNextID()
			product = append(product, NewProduct)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(NewProduct)
			return
		}

		// GET
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)


	})



	// GET /health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "oke",
			"meassage": "API Running",
		})
	})


	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}

}