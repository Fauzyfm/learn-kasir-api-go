package service

import (
	repositories "kasir-api/internal/Repositories"
	"kasir-api/internal/models"
	"time"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailyReport() (*models.Report, error) {
	return s.repo.GetDailyReport()
}

func (s *ReportService) GetReportByDateRange(startDate, endDate time.Time) (*models.Report, error) {
	return s.repo.GetReportByDateRange(startDate, endDate)
}
