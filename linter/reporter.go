package linter

import "time"

//Report specifies a list of validation errors
//and time details of modifying the report
type Report struct {
	Name       string
	Created    *time.Time
	Updated    *time.Time
	Finished   *time.Time
	Error      string // TODO: do we need this one?
	Problems   []*Problem
	Statistics Count
}

// Count specifies a number of reviewed(examined) document
type Count struct {
	Total     int64
	Inspected int64
	Valid     int64
	Invalid   int64
}

//Problem specifies details of error
//where original is a string representation of the original file
//that contains problem(s)
type Problem struct {
	Id       string
	Original string
	Details  []*ProblemDetails
}

//ProblemDetails specifies information about validation errors
//where id is identifier of the problem's location in the original file
type ProblemDetails struct {
	Id          string
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
