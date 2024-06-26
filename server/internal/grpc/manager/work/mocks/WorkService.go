// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"
	paymentsmodel "server/server/internal/models/work"

	mock "github.com/stretchr/testify/mock"
)

// WorkService is an autogenerated mock type for the WorkService type
type WorkService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, rec
func (_m *WorkService) Create(ctx context.Context, rec paymentsmodel.CreateWork) error {
	ret := _m.Called(ctx, rec)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, paymentsmodel.CreateWork) error); ok {
		r0 = rf(ctx, rec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, rec
func (_m *WorkService) Delete(ctx context.Context, rec paymentsmodel.DeleteWork) error {
	ret := _m.Called(ctx, rec)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, paymentsmodel.DeleteWork) error); ok {
		r0 = rf(ctx, rec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, rec
func (_m *WorkService) Get(ctx context.Context, rec paymentsmodel.GetAllWork) ([]paymentsmodel.Work, error) {
	ret := _m.Called(ctx, rec)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []paymentsmodel.Work
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, paymentsmodel.GetAllWork) ([]paymentsmodel.Work, error)); ok {
		return rf(ctx, rec)
	}
	if rf, ok := ret.Get(0).(func(context.Context, paymentsmodel.GetAllWork) []paymentsmodel.Work); ok {
		r0 = rf(ctx, rec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]paymentsmodel.Work)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, paymentsmodel.GetAllWork) error); ok {
		r1 = rf(ctx, rec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByDate provides a mock function with given fields: ctx, rec
func (_m *WorkService) GetByDate(ctx context.Context, rec paymentsmodel.GetAllWorkByDate) ([]paymentsmodel.Work, error) {
	ret := _m.Called(ctx, rec)

	if len(ret) == 0 {
		panic("no return value specified for GetByDate")
	}

	var r0 []paymentsmodel.Work
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, paymentsmodel.GetAllWorkByDate) ([]paymentsmodel.Work, error)); ok {
		return rf(ctx, rec)
	}
	if rf, ok := ret.Get(0).(func(context.Context, paymentsmodel.GetAllWorkByDate) []paymentsmodel.Work); ok {
		r0 = rf(ctx, rec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]paymentsmodel.Work)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, paymentsmodel.GetAllWorkByDate) error); ok {
		r1 = rf(ctx, rec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, rec
func (_m *WorkService) Update(ctx context.Context, rec paymentsmodel.UpdateWork) error {
	ret := _m.Called(ctx, rec)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, paymentsmodel.UpdateWork) error); ok {
		r0 = rf(ctx, rec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWorkService creates a new instance of WorkService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWorkService(t interface {
	mock.TestingT
	Cleanup(func())
}) *WorkService {
	mock := &WorkService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
