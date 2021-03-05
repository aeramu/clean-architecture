// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/aeramu/clean-architecture/service/api"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: ctx, req
func (_m *Service) CreateBook(ctx context.Context, req api.CreateBookReq) (*api.CreateBookRes, error) {
	ret := _m.Called(ctx, req)

	var r0 *api.CreateBookRes
	if rf, ok := ret.Get(0).(func(context.Context, api.CreateBookReq) *api.CreateBookRes); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.CreateBookRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.CreateBookReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBook provides a mock function with given fields: ctx, req
func (_m *Service) GetBook(ctx context.Context, req api.GetBookReq) (*api.GetBookRes, error) {
	ret := _m.Called(ctx, req)

	var r0 *api.GetBookRes
	if rf, ok := ret.Get(0).(func(context.Context, api.GetBookReq) *api.GetBookRes); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.GetBookRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.GetBookReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
