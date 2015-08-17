package linter

import (
	"time"
	"github.com/GabbyyLS/data-linter/models"
)

// Reporter embeds read and write functionality for reports.
type Reporter struct {
	ReportReader
	ReportWriter
}

// ReportWriter instantiates a report, fills it with problems, and closes it.
type ReportWriter interface {
	Start() (*models.Report, error)
	Finish(currentReport *models.Report) (*models.Report, error)
	Commit(currentReport *models.Report, []*models.Problem) (*models.Report, error)
}

// ReportReader provides functionality for accessing various report-related information.
type ReportReader interface {
	GetAllReports() []*models.Report
	GetReportByName(Name string) *models.Report
	GetReportByCreationTime(time *time.Time) *models.Report
	GetReportByUpdateTime(time *time.Time) *models.Report
	GetReportByCommitTime(time *time.Time) *models.Report
	TotalReportsCount() *models.ReportsCount
}



