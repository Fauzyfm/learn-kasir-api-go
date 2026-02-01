package service

import (
	"kasir-api/internal/Repositories"
	"kasir-api/internal/models"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]models.Categories, error) {
	return s.repo.GetAllCategories()
}

func (s *CategoryService) CreateCategory(category *models.Categories) error {
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) GetCategoryByID(id int) (*models.Categories, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *CategoryService) UpdateCategory(id int, category *models.Categories) (*models.Categories, error) {
	return s.repo.UpdateCategory(id, category)
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.repo.DeleteCategory(id)
}

