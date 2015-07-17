package linter

import "time"

type Report struct {
	Name       string
	Created    *time.Time
	Updated    *time.Time
	Finished   *time.Time
	Error      string // TODO: do we need this one?
	Problems   []*Problem
	Statistics Count
}

type Count struct {
	Total     int64
	Inspected int64
	Valid     int64
	Invalid   int64
}

type Problem struct {
	Id       string
	Original string // string representation of the original file that contains problems(s)
	Details  []*ProblemDetails
}

type ProblemDetails struct {
	Id          string // identifier of the problem's location in the original file
	Fragment    string
	Description string
}

type Reporter struct {
	ReportReader
	ReportWriter
}

type ReportWriter interface {
	Start(*Report) error
	Finish(*Report) error
	Commit(*Report, []*Problem) error
}

type ReportReader interface {
	// TODO: complete
	Count()
	Find()
}
