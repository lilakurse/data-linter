package mocks

//import "errors"

//func NewMockEmptyIterator() *Iterator {
//	mockObj := new(Iterator)
//	mockObj.On("Name").Return("Mock iterator with NO data")
//	mockObj.On("Count").Return(0)
//	mockObj.On("Next", mock.Step).Return(  ) // TODO: should return []*Checker
//	// We already have these:
//	//	NewMockValidDoc() *Checker
//	//	NewMockInvalidDoc() *Checker
//	//	NewMockBadDoc() *Checker
//
//	return mockObj
//}
//
//func NewMockFilledIterator() *Iterator {
//	mockObj := new(Iterator)
//	mockObj.On("Name").Return("Mock iterator with data")
//	mockObj.On("Count").Return(mock.Count)
//	mockObj.On("Next", mock.Step).Return(NewMockValidDoc()) // should return []*Checker
//	// We already have these:
//	//	NewMockValidDoc() *Checker
//	//	NewMockInvalidDoc() *Checker
//	//	NewMockBadDoc() *Checker
//
//	return mockObj
//}
//
//func NewMockWithErrorIterator() *Iterator {
//	mockObj := new(Iterator)
//	mockObj.On("Name").Return("Mock iterator with error")
//	mockObj.On("Count").Return(mock.Count)
//	mockObj.On("Next", mock.Step).Return(errors.New(NewMockBadDoc()))
//	return mockObj
//}
