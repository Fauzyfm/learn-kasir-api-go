package service

import "kasir-api/internal/models"

var Products = []models.Product{
	{ID: 1, NamaBarang: "Pensil", HargaBarang: 2000, Stok: 100},
	{ID: 2, NamaBarang: "Buku Tulis", HargaBarang: 5000, Stok: 200},
	{ID: 3, NamaBarang: "Penghapus", HargaBarang: 1500, Stok: 150},
}

// GetProductByID - mendapatkan product berdasarkan ID
func GetProductByID(id int) (*models.Product, error) {
	for i, p := range Products {
		if p.ID == id {
			return &Products[i], nil
		}
	}
	return nil, nil
}

// GetAllProducts - mendapatkan semua products
func GetAllProducts() []models.Product {
	return Products
}

// CreateProduct - membuat product baru
func CreateProduct(product models.Product) models.Product {
	product.ID = GetNextID()
	Products = append(Products, product)
	return product
}

// UpdateProduct - update product berdasarkan ID
func UpdateProduct(id int, product models.Product) (*models.Product, error) {
	for i := range Products {
		if Products[i].ID == id {
			product.ID = id
			Products[i] = product
			return &Products[i], nil
		}
	}
	return nil, nil
}

// DeleteProduct - delete product berdasarkan ID
func DeleteProduct(id int) bool {
	for i := range Products {
		if Products[i].ID == id {
			Products = append(Products[:i], Products[i+1:]...)
			return true
		}
	}
	return false
}

// GetNextID - mendapatkan ID terbaru
func GetNextID() int {
	maxID := 0
	for _, p := range Products {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	return maxID + 1
}
