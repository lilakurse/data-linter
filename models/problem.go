package models

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