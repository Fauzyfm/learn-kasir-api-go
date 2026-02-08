package repositories

import (
	"database/sql"
	"kasir-api/internal/models"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetDailyReport() (*models.Report, error) {
	// Get today's start and end date
	today := time.Now()
	startDate := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	endDate := startDate.Add(24 * time.Hour)

	return repo.GetReportByDateRange(startDate, endDate)
}

func (repo *ReportRepository) GetReportByDateRange(startDate, endDate time.Time) (*models.Report, error) {
	// Get total revenue
	var totalRevenue int
	err := repo.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0) FROM transactions 
		WHERE created_at >= $1 AND created_at < $2
	`, startDate, endDate).Scan(&totalRevenue)
	if err != nil {
		return nil, err
	}

	// Get total transactions
	var totalTransaksi int
	err = repo.db.QueryRow(`
		SELECT COALESCE(COUNT(*), 0) FROM transactions 
		WHERE created_at >= $1 AND created_at < $2
	`, startDate, endDate).Scan(&totalTransaksi)
	if err != nil {
		return nil, err
	}

	// Get best selling product
	var produkName string
	var qtyTerjual int
	err = repo.db.QueryRow(`
		SELECT p.nama_barang, SUM(td.quantity) as total_qty
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.nama_barang
		ORDER BY total_qty DESC
		LIMIT 1
	`, startDate, endDate).Scan(&produkName, &qtyTerjual)

	if err == sql.ErrNoRows {
		// If no transactions, return with empty product
		produkName = ""
		qtyTerjual = 0
	} else if err != nil {
		return nil, err
	}

	// Get all transactions with details in date range
	transactions := make([]models.ReportTransaction, 0)
	rows, err := repo.db.Query(`
		SELECT id, total_amount FROM transactions 
		WHERE created_at >= $1 AND created_at < $2
		ORDER BY id DESC
	`, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transID int
		var totalAmount int
		err := rows.Scan(&transID, &totalAmount)
		if err != nil {
			return nil, err
		}

		// Get details for this transaction
		details := make([]models.TransactionDetails, 0)
		detailRows, err := repo.db.Query(`
			SELECT id, transaction_id, product_id, quantity, subtotal FROM transaction_details 
			WHERE transaction_id = $1
		`, transID)
		if err != nil {
			return nil, err
		}

		for detailRows.Next() {
			var detail models.TransactionDetails
			err := detailRows.Scan(&detail.ID, &detail.TransactionID, &detail.ProductID, &detail.Quantity, &detail.Subtotal)
			if err != nil {
				detailRows.Close()
				return nil, err
			}

			// Get product name
			var productName string
			err = repo.db.QueryRow(`
				SELECT nama_barang FROM products WHERE id = $1
			`, detail.ProductID).Scan(&productName)
			if err == nil {
				detail.ProductName = productName
			}

			details = append(details, detail)
		}
		detailRows.Close()

		transactions = append(transactions, models.ReportTransaction{
			ID:          transID,
			TotalAmount: totalAmount,
			Details:     details,
		})
	}

	return &models.Report{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: models.ProdukTerlaris{
			Nama:       produkName,
			QtyTerjual: qtyTerjual,
		},
		Transactions: transactions,
	}, nil
}
