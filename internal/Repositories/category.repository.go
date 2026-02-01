package repositories

import (
	"database/sql"
	"kasir-api/internal/models"
)

type CategoryRepository struct {
	db *sql.DB
}


func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAllCategories() ([]models.Categories, error) {
	rows, err := repo.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return  nil, err
	}
	defer rows.Close()

	categories := make([]models.Categories, 0)
	for rows.Next() {
		var c models.Categories
		err := rows.Scan(&c.ID, &c.Name, &c.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return  categories, nil
}

func (repo *CategoryRepository) CreateCategory(category *models.Categories)  error {
	query := "INSERT into categories ( name, description) VALUES ($1, $2) RETURNING id"
	
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CategoryRepository) GetCategoryByID(id int) (*models.Categories, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"

	var category models.Categories

	err := repo.db.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}

	return &category, nil
}


func (repo *CategoryRepository) UpdateCategory(id int, category *models.Categories) (*models.Categories, error) {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"

	_, err := repo.db.Exec(query, category.Name, category.Description, id)
	if err != nil {
		return nil, err
	}

	var updateCategory models.Categories

	err = repo.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id).Scan(&updateCategory.ID, &updateCategory.Name, &updateCategory.Description)
	if err != nil {
		return  nil, err
	}

	return  &updateCategory, nil
}


func (repo *CategoryRepository) DeleteCategory(id int) error {
		tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()


	result, err := tx.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return	err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return tx.Commit()
}
