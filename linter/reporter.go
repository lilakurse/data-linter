package linter

import (
	"time"
)

// Report contains a list of problems (i.e., various validation errors)
// and time details of modifications made to the report.
type Report struct {
	Name       string
	Created    *time.Time
	Updated    *time.Time
	Finished   *time.Time
	Error      string
	Problems   []*Problem
	Statistics Count
}

// Count contains numeric information on inspected documents.
type Count struct {
	Total     int64
	Inspected int64
	Valid     int64
	Invalid   int64
}

// Problem refers to a particular file that contains problem(s).
// Original is the string representation of that file.
type Problem struct {
	Id       string
	Original string
	Details  []*ProblemDetails
}

// ProblemDetails contains information about a specific validation error.
// Id identifies the problem's location in the original file.
type ProblemDetails struct {
	Id          string
	Fragment    string
	Description string
}

// Reporter embeds read and write functionality for reports.
type Reporter struct {
	ReportReader
	ReportWriter
}

// ReportWriter instantiates a report, fills it with problems, and closes it.
type ReportWriter interface {
	Start() (*Report, error)
	Finish(*Report) error
	Commit(*Report, []*Problem) error
}

// ReportReader provides functionality for accessing various report-related information.
type ReportReader interface {
	GetAllReports() []*Report
	GetReportByName(Name string) *Report
	GetReportByCreationTime(time *time.Time) *Report
	GetReportByUpdateTime(time *time.Time) *Report
	GetReportByCommitTime(time *time.Time) *Report
	TotalReportsCount() *ReportsCount
}

// ReportCount contains statistics on reports.
type ReportsCount struct {
	Total      int64
	Failed     int64
	Successful int64
}
