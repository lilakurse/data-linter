package mocks

// Valid writer with empty list of problems.
func NewMockValidReportWriter() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return()
	mockObj.On("Finish").Return()
	mockObj.On("Commit").Return()
	return mockObj
}

// Invalid writer with some problems.
func NewMockInvalidRepotWriter() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return()
	mockObj.On("Finish").Return()
	mockObj.On("Commit").Return()
	return mockObj
}

// Bad writer with error and no report.
func NewMockBadReportWriter() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return()
	mockObj.On("Finish").Return()
	mockObj.On("Commit").Return()
	return mockObj
}
