// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// PlacesFacade is an autogenerated mock type for the PlacesFacade type
type PlacesFacade struct {
	mock.Mock
}

type PlacesFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *PlacesFacade) EXPECT() *PlacesFacade_Expecter {
	return &PlacesFacade_Expecter{mock: &_m.Mock}
}

// GetPlaceByID provides a mock function with given fields: ctx, placeID
func (_m *PlacesFacade) GetPlaceByID(ctx context.Context, placeID int32) (model.Place, error) {
	ret := _m.Called(ctx, placeID)

	var r0 model.Place
	if rf, ok := ret.Get(0).(func(context.Context, int32) model.Place); ok {
		r0 = rf(ctx, placeID)
	} else {
		r0 = ret.Get(0).(model.Place)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, placeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlacesFacade_GetPlaceByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlaceByID'
type PlacesFacade_GetPlaceByID_Call struct {
	*mock.Call
}

// GetPlaceByID is a helper method to define mock.On call
//  - ctx context.Context
//  - placeID int32
func (_e *PlacesFacade_Expecter) GetPlaceByID(ctx interface{}, placeID interface{}) *PlacesFacade_GetPlaceByID_Call {
	return &PlacesFacade_GetPlaceByID_Call{Call: _e.mock.On("GetPlaceByID", ctx, placeID)}
}

func (_c *PlacesFacade_GetPlaceByID_Call) Run(run func(ctx context.Context, placeID int32)) *PlacesFacade_GetPlaceByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *PlacesFacade_GetPlaceByID_Call) Return(_a0 model.Place, _a1 error) *PlacesFacade_GetPlaceByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewPlacesFacade interface {
	mock.TestingT
	Cleanup(func())
}

// NewPlacesFacade creates a new instance of PlacesFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPlacesFacade(t mockConstructorTestingTNewPlacesFacade) *PlacesFacade {
	mock := &PlacesFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
