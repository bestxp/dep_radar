package mocks

import interfaces "github.com/stamm/dep_radar/interfaces"
import mock "github.com/stretchr/testify/mock"

// IDepTool is an autogenerated mock type for the IDepTool type
type IDepTool struct {
	mock.Mock
}

// Deps provides a mock function with given fields: _a0
func (_m *IDepTool) Deps(_a0 interfaces.IApp) (interfaces.AppDeps, error) {
	ret := _m.Called(_a0)

	var r0 interfaces.AppDeps
	if rf, ok := ret.Get(0).(func(interfaces.IApp) interfaces.AppDeps); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(interfaces.AppDeps)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interfaces.IApp) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *IDepTool) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

var _ interfaces.IDepTool = (*IDepTool)(nil)