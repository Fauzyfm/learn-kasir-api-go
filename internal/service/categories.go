package service

import (
	"kasir-api/internal/models"
)

var Categories = []models.Categories{
	{ID: 1, Name: "Sepatu Pria", Description: "Penjualan Khusus Sepatu Untuk Pria"},
	{ID: 2, Name: "Sepatu Wanita", Description: "Penjualan Khusus Sepatu Untuk Wanita"},
	{ID: 3, Name: "Jacket Unisex", Description: "Penjualan Khusus Jacket Pria dan Wanita"},
}

func GetAllCategories() []models.Categories {
	return Categories
}

func CreateCategories(categorie models.Categories) models.Categories {
	categorie.ID = GetNextID()
	Categories = append(Categories, categorie)
	return categorie
}

func GetCategorieByID(id int) (*models.Categories, error) {
	for i, c := range Categories {
		if c.ID == id {
			return &Categories[i], nil
		}
	}
	return nil, nil
}

func UpdateCategorie(id int, categorie models.Categories) (*models.Categories, error) {
	for i := range Categories {
		if Categories[i].ID == id {
			categorie.ID = id
			Categories[i] = categorie
			return &Categories[i], nil
		}
	}
	return nil, nil
}

func DeleteCategorie(id int) bool {
	for i := range Categories {
		if Categories[i].ID == id {
			Categories = append(Categories[:i], Categories[i+1:]...)
			return true
		}
	}
	return false
}
