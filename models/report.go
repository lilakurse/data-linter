package models

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

// ReportCount contains statistics on reports.
type ReportsCount struct {
	Total      int64
	Failed     int64
	Successful int64
}