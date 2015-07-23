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
	Error      string // TODO: do we need this one?
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
type ReportWriter interface {
	Start() (*Report, error) // TODO: think about this signature ( was: Start(*Report) error )
	Finish(*Report) error
	Commit(*Report, []*Problem) error
}

// TODO: document
type ReportReader interface {
	// TODO: complete
	GetAllReports() []Report            // Get all reports that presently exist.
	GetReportByName(Name string) Report // Find a report by its name.
	TotalReportsCount() int64           // Total amount of scanned docs.
	Logger() string                     // Report time information.
	// TODO: how about searching by creation/update/commit dates?
	// TODO: how about getting a particular report's statistics?
}
