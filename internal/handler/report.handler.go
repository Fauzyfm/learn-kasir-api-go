package handler

import (
	"encoding/json"
	"fmt"
	"kasir-api/internal/service"
	"net/http"
	"time"
)

type ReportHandler struct {
	service *service.ReportService
}

func NewReportHandler(service *service.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	var report interface{}
	var err error

	// If date range parameters provided, use them
	if startDateStr != "" && endDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			http.Error(w, "Invalid start_date format. Use YYYY-MM-DD", http.StatusBadRequest)
			return
		}

		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			http.Error(w, "Invalid end_date format. Use YYYY-MM-DD", http.StatusBadRequest)
			return
		}

		// Add one day to end_date to include the entire end date
		endDate = endDate.Add(24 * time.Hour)

		report, err = h.service.GetReportByDateRange(startDate, endDate)
	} else {
		// Get daily report
		report, err = h.service.GetDailyReport()
	}

	if err != nil {
		fmt.Printf("DEBUG: Report error: %v\n", err)
		http.Error(w, "Failed to get report: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
