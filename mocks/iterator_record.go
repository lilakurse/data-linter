package mocks

func NewMockEmptyIterator() *Iterator {
	mockObj := new(Iterator)
	mockObj.On("Name").Return("Mock iterator with NO data")
	mockObj.On("Count").Return(0)
	mockObj.On("Next", mock.Step).Return()  // should return []*Checker
											// We already have these:
											//	NewMockValidDoc() *Checker
											//	NewMockInvalidDoc() *Checker
											//	NewMockBadDoc() *Checker

	return mockObj
}

func NewMockFilledIterator() *Iterator {
	mockObj := new(Iterator)
	mockObj.On("Name").Return("Mock iterator with data")
	mockObj.On("Count").Return(100)
	mockObj.On("Next", mock.Step).Return()  // should return []*Checker
											// We already have these:
											//	NewMockValidDoc() *Checker
											//	NewMockInvalidDoc() *Checker
											//	NewMockBadDoc() *Checker

	return mockObj
}
