package service

import (
	"kasir-api/internal/Repositories"
	"kasir-api/internal/models"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}


func (s *ProductService) GetAllProducts(name string) ([]models.Product, error) {
	return s.repo.GetAllProducts(name)
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)

}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	return s.repo.GetProductByID(id)
	
}


func (s *ProductService) UpdateProduct(id int, product models.Product) (*models.Product, error) {
	return s.repo.UpdateProduct(id, product)
}


func (s *ProductService) DeleteProduct(id int) error {
	return  s.repo.DeleteProduct(id)
}


func (s *ProductService) GetProductByIDWithCategory(id int) (*models.ProductWithCategory, error) {
	return s.repo.GetProductByIDWithCategory(id)
}
