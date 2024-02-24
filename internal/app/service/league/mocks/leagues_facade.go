// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// LeaguesFacade is an autogenerated mock type for the LeaguesFacade type
type LeaguesFacade struct {
	mock.Mock
}

type LeaguesFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *LeaguesFacade) EXPECT() *LeaguesFacade_Expecter {
	return &LeaguesFacade_Expecter{mock: &_m.Mock}
}

// GetLeagueByID provides a mock function with given fields: ctx, id
func (_m *LeaguesFacade) GetLeagueByID(ctx context.Context, id int32) (model.League, error) {
	ret := _m.Called(ctx, id)

	var r0 model.League
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (model.League, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) model.League); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.League)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LeaguesFacade_GetLeagueByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLeagueByID'
type LeaguesFacade_GetLeagueByID_Call struct {
	*mock.Call
}

// GetLeagueByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id int32
func (_e *LeaguesFacade_Expecter) GetLeagueByID(ctx interface{}, id interface{}) *LeaguesFacade_GetLeagueByID_Call {
	return &LeaguesFacade_GetLeagueByID_Call{Call: _e.mock.On("GetLeagueByID", ctx, id)}
}

func (_c *LeaguesFacade_GetLeagueByID_Call) Run(run func(ctx context.Context, id int32)) *LeaguesFacade_GetLeagueByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *LeaguesFacade_GetLeagueByID_Call) Return(_a0 model.League, _a1 error) *LeaguesFacade_GetLeagueByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LeaguesFacade_GetLeagueByID_Call) RunAndReturn(run func(context.Context, int32) (model.League, error)) *LeaguesFacade_GetLeagueByID_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewLeaguesFacade interface {
	mock.TestingT
	Cleanup(func())
}

// NewLeaguesFacade creates a new instance of LeaguesFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLeaguesFacade(t mockConstructorTestingTNewLeaguesFacade) *LeaguesFacade {
	mock := &LeaguesFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
