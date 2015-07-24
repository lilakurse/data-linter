package linter

import (
	"time"
)

// The Report struct contains a list of problems (i.e., various validation errors)
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

// The auxiliary Count struct contains numeric information on inspected documents.
type Count struct {
	Total     int64
	Inspected int64
	Valid     int64
	Invalid   int64
}

// The Problem struct refers to a particular file that contains problem(s).
// Original is the string representation of that file.
type Problem struct {
	Id       string
	Original string
	Details  []*ProblemDetails
}

// The auxiliary ProblemDetails struct contains information about specific validation error.
// Id identifies the problem's location in the original file.
type ProblemDetails struct {
	Id          string
	Fragment    string
	Description string
}

// The Reporter struct contains information about read and write operations on reporter.
type Reporter struct {
	ReportReader
	ReportWriter
}

// TODO: document
// The interface ReportWriter contains three methods that creates a report,
// filling it with inspected problems and closes it.
type ReportWriter interface {
	Start() (*Report, error)
	Finish(*Report) error
	Commit(*Report, []*Problem) error
}

// TODO: document
// The interface ReportReader contains methods that shows all created reports and there total count
// and provides a search on different fields.
type ReportReader interface {
	GetAllReports() []*Report                        // Get all reports that presently exist.
	GetReportByName(Name string) *Report             // Find a report by its name. //
	GetReportByCreationTime(time *time.Time) *Report // Find a report by its creation date.
	GetReportByUpdateTime(time *time.Time) *Report   // Find a report by its update date.
	GetReportByCommitTime(time *time.Time) *Report   // Find a report by its commit date.
	TotalReportsCount() *ReportsCount                // Statistics on reports (how many failed, how many were successfull etc.).
}

// TODO: document
// The ReportCount struct contains statistics about total number of reports.
type ReportsCount struct {
	Total       int64
	Failed      int64
	Successfull int64
}
