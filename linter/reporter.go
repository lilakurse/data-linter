package linter

import (
	"github.com/GabbyyLS/data-linter/models"
	"time"
)

// Reporter embeds read and write functionality for reports.
type Reporter struct {
	ReportReader
	ReportWriter
}

// ReportWriter instantiates a report, fills it with problems, and closes it.
type ReportWriter interface {
	Start() (*models.Report, error)
	Finish(*models.Report) error
	Commit(*models.Report, []*models.Problem) error
}

func Start(reportWriter ReportWriter) (*models.Report, error) {
	return reportWriter.Start()
}
func Finish(reportWriter ReportWriter, report *models.Report) error {
	return reportWriter.Finish(report)
}

func Commit(reportWriter ReportWriter, report *models.Report, problem []*models.Problem) error {
	return reportWriter.Commit(report, problem)
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
