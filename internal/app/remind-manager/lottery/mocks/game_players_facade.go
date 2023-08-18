// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GamePlayersFacade is an autogenerated mock type for the GamePlayersFacade type
type GamePlayersFacade struct {
	mock.Mock
}

type GamePlayersFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *GamePlayersFacade) EXPECT() *GamePlayersFacade_Expecter {
	return &GamePlayersFacade_Expecter{mock: &_m.Mock}
}

// GetGamePlayersByGameID provides a mock function with given fields: ctx, gameID
func (_m *GamePlayersFacade) GetGamePlayersByGameID(ctx context.Context, gameID int32) ([]model.GamePlayer, error) {
	ret := _m.Called(ctx, gameID)

	var r0 []model.GamePlayer
	if rf, ok := ret.Get(0).(func(context.Context, int32) []model.GamePlayer); ok {
		r0 = rf(ctx, gameID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.GamePlayer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, gameID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamePlayersFacade_GetGamePlayersByGameID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGamePlayersByGameID'
type GamePlayersFacade_GetGamePlayersByGameID_Call struct {
	*mock.Call
}

// GetGamePlayersByGameID is a helper method to define mock.On call
//  - ctx context.Context
//  - gameID int32
func (_e *GamePlayersFacade_Expecter) GetGamePlayersByGameID(ctx interface{}, gameID interface{}) *GamePlayersFacade_GetGamePlayersByGameID_Call {
	return &GamePlayersFacade_GetGamePlayersByGameID_Call{Call: _e.mock.On("GetGamePlayersByGameID", ctx, gameID)}
}

func (_c *GamePlayersFacade_GetGamePlayersByGameID_Call) Run(run func(ctx context.Context, gameID int32)) *GamePlayersFacade_GetGamePlayersByGameID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *GamePlayersFacade_GetGamePlayersByGameID_Call) Return(_a0 []model.GamePlayer, _a1 error) *GamePlayersFacade_GetGamePlayersByGameID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewGamePlayersFacade interface {
	mock.TestingT
	Cleanup(func())
}

// NewGamePlayersFacade creates a new instance of GamePlayersFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGamePlayersFacade(t mockConstructorTestingTNewGamePlayersFacade) *GamePlayersFacade {
	mock := &GamePlayersFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}