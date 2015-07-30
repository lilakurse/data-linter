package mocks
import (
	"errors"
	"strconv"
)

func NewMockIterator() *Iterator {
	mockObj := new(Iterator)
	mockObj.On("Name").Return(errors.New("Bad Name"))
	mockObj.On("Name").Return("Good Name")
	mockObj.On("Count").Return(errors.New(strconv.Itoa(666)))
	mockObj.On("Count").Return(Count)
	mockObj.On("Next",mock.Step).Return()
	mockObj.On("Next",mock.Step).Return()

	// TODO: expectations for tests

	return mockObj
}