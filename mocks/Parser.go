// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// Parser is an autogenerated mock type for the Parser type
type Parser struct {
	mock.Mock
}

// Parse provides a mock function with given fields: reader
func (_m *Parser) Parse(reader io.Reader) ([]string, error) {
	ret := _m.Called(reader)

	var r0 []string
	if rf, ok := ret.Get(0).(func(io.Reader) []string); ok {
		r0 = rf(reader)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader) error); ok {
		r1 = rf(reader)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}