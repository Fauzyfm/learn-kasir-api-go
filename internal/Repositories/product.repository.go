package repositories

import (
	"database/sql"
	"kasir-api/internal/models"

)


type ProductRepository struct {
	db *sql.DB
}
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAllProducts() ([]models.Product, error) {
	rows, err := repo.db.Query("SELECT id, nama_barang, harga_barang, stok FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.NamaBarang, &p.HargaBarang, &p.Stok)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (repo *ProductRepository) CreateProduct(product *models.Product) error {
	query := "INSERT INTO products (nama_barang, harga_barang, stok, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.NamaBarang, product.HargaBarang, product.Stok, product.CategoryID).Scan(&product.ID)
	if err != nil {
		return err
	}
	return nil
}


func (repo *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	query := "SELECT id, nama_barang, harga_barang, stok FROM products WHERE id = $1"
	var product models.Product

	err := repo.db.QueryRow(query, id).Scan(&product.ID, &product.NamaBarang, &product.HargaBarang, &product.Stok)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepository) UpdateProduct(id int, product models.Product) (*models.Product, error) {
	query := "UPDATE products SET nama_barang = $1, harga_barang = $2, stok = $3 WHERE id = $4"
	
	_, err := repo.db.Exec(query, product.NamaBarang, product.HargaBarang, product.Stok, id)
	if err != nil {
		return nil, err
	}
	
	var updateProduct models.Product

	err = repo.db.QueryRow("SELECT id, nama_barang, harga_barang, stok FROM products WHERE id = $1", id).Scan(&updateProduct.ID, &updateProduct.NamaBarang, &updateProduct.HargaBarang, &updateProduct.Stok)
	if err != nil {
		return  nil, err
	}

	return &updateProduct, nil

}

func (repo *ProductRepository) DeleteProduct(id int) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()


	result, err := tx.Exec("DELETE FROM products WHERE id = $1", id)
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

func (repo *ProductRepository) GetProductByIDWithCategory(id int) (*models.ProductWithCategory, error) {
    query := `
        SELECT 
            p.id, 
            p.nama_barang, 
            p.harga_barang, 
            p.stok, 
            p.category_id,
            c.id as cat_id,
            c.name as cat_name,
            c.description as cat_description
        FROM products p
        LEFT JOIN categories c ON p.category_id = c.id
        WHERE p.id = $1
    `
    
    var product models.ProductWithCategory
    var category models.Categories
    
    err := repo.db.QueryRow(query, id).Scan(
        &product.ID,
        &product.NamaBarang,
        &product.HargaBarang,
        &product.Stok,
        &product.CategoryID,
        &category.ID,
        &category.Name,
        &category.Description,
    )
    
    if err != nil {
        return nil, err
    }
    
    product.Category = category
    return &product, nil
}

