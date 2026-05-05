package mocks

import (
	"github.com/stretchr/testify/mock"
)

// URLGetter is a mock implementation of redirect.URLGetter interface
type URLGetter struct {
	mock.Mock
}

// GetURL mocks the GetURL method
func (m *URLGetter) GetURL(alias string) (string, error) {
	args := m.Called(alias)
	return args.String(0), args.Error(1)
}

// NewURLGetter creates a new instance of URLGetter
func NewURLGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *URLGetter {
	m := &URLGetter{}
	m.Mock.Test(t)
	t.Cleanup(func() { m.AssertExpectations(t) })
	return m
}
