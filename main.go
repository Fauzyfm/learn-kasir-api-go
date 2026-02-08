package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	repositories "kasir-api/internal/Repositories"
	"kasir-api/internal/handler"
	"kasir-api/internal/models"
	"kasir-api/internal/router"
	"kasir-api/internal/service"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := models.Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Inisialisasi database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// Report
	reportRepo := repositories.NewReportRepository(db)
	reportService := service.NewReportService(reportRepo)
	reportHandler := handler.NewReportHandler(reportService)

	// Setup semua routes
	router.SetupRoutes(productHandler, categoryHandler, transactionHandler, reportHandler)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "Server is Running...",
		})
	})

	fmt.Println("Server started on port", config.Port)
	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
