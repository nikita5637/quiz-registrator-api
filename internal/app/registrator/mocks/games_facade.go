// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// GamesFacade is an autogenerated mock type for the GamesFacade type
type GamesFacade struct {
	mock.Mock
}

type GamesFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *GamesFacade) EXPECT() *GamesFacade_Expecter {
	return &GamesFacade_Expecter{mock: &_m.Mock}
}

// AddGame provides a mock function with given fields: ctx, game
func (_m *GamesFacade) AddGame(ctx context.Context, game model.Game) (int32, error) {
	ret := _m.Called(ctx, game)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, model.Game) int32); ok {
		r0 = rf(ctx, game)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Game) error); ok {
		r1 = rf(ctx, game)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamesFacade_AddGame_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddGame'
type GamesFacade_AddGame_Call struct {
	*mock.Call
}

// AddGame is a helper method to define mock.On call
//  - ctx context.Context
//  - game model.Game
func (_e *GamesFacade_Expecter) AddGame(ctx interface{}, game interface{}) *GamesFacade_AddGame_Call {
	return &GamesFacade_AddGame_Call{Call: _e.mock.On("AddGame", ctx, game)}
}

func (_c *GamesFacade_AddGame_Call) Run(run func(ctx context.Context, game model.Game)) *GamesFacade_AddGame_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Game))
	})
	return _c
}

func (_c *GamesFacade_AddGame_Call) Return(_a0 int32, _a1 error) *GamesFacade_AddGame_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// AddGames provides a mock function with given fields: ctx, games
func (_m *GamesFacade) AddGames(ctx context.Context, games []model.Game) error {
	ret := _m.Called(ctx, games)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []model.Game) error); ok {
		r0 = rf(ctx, games)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GamesFacade_AddGames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddGames'
type GamesFacade_AddGames_Call struct {
	*mock.Call
}

// AddGames is a helper method to define mock.On call
//  - ctx context.Context
//  - games []model.Game
func (_e *GamesFacade_Expecter) AddGames(ctx interface{}, games interface{}) *GamesFacade_AddGames_Call {
	return &GamesFacade_AddGames_Call{Call: _e.mock.On("AddGames", ctx, games)}
}

func (_c *GamesFacade_AddGames_Call) Run(run func(ctx context.Context, games []model.Game)) *GamesFacade_AddGames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]model.Game))
	})
	return _c
}

func (_c *GamesFacade_AddGames_Call) Return(_a0 error) *GamesFacade_AddGames_Call {
	_c.Call.Return(_a0)
	return _c
}

// DeleteGame provides a mock function with given fields: ctx, gameID
func (_m *GamesFacade) DeleteGame(ctx context.Context, gameID int32) error {
	ret := _m.Called(ctx, gameID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, gameID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GamesFacade_DeleteGame_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteGame'
type GamesFacade_DeleteGame_Call struct {
	*mock.Call
}

// DeleteGame is a helper method to define mock.On call
//  - ctx context.Context
//  - gameID int32
func (_e *GamesFacade_Expecter) DeleteGame(ctx interface{}, gameID interface{}) *GamesFacade_DeleteGame_Call {
	return &GamesFacade_DeleteGame_Call{Call: _e.mock.On("DeleteGame", ctx, gameID)}
}

func (_c *GamesFacade_DeleteGame_Call) Run(run func(ctx context.Context, gameID int32)) *GamesFacade_DeleteGame_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *GamesFacade_DeleteGame_Call) Return(_a0 error) *GamesFacade_DeleteGame_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetGameByID provides a mock function with given fields: ctx, id
func (_m *GamesFacade) GetGameByID(ctx context.Context, id int32) (model.Game, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Game
	if rf, ok := ret.Get(0).(func(context.Context, int32) model.Game); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Game)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamesFacade_GetGameByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGameByID'
type GamesFacade_GetGameByID_Call struct {
	*mock.Call
}

// GetGameByID is a helper method to define mock.On call
//  - ctx context.Context
//  - id int32
func (_e *GamesFacade_Expecter) GetGameByID(ctx interface{}, id interface{}) *GamesFacade_GetGameByID_Call {
	return &GamesFacade_GetGameByID_Call{Call: _e.mock.On("GetGameByID", ctx, id)}
}

func (_c *GamesFacade_GetGameByID_Call) Run(run func(ctx context.Context, id int32)) *GamesFacade_GetGameByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *GamesFacade_GetGameByID_Call) Return(_a0 model.Game, _a1 error) *GamesFacade_GetGameByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetGames provides a mock function with given fields: ctx
func (_m *GamesFacade) GetGames(ctx context.Context) ([]model.Game, error) {
	ret := _m.Called(ctx)

	var r0 []model.Game
	if rf, ok := ret.Get(0).(func(context.Context) []model.Game); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Game)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamesFacade_GetGames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGames'
type GamesFacade_GetGames_Call struct {
	*mock.Call
}

// GetGames is a helper method to define mock.On call
//  - ctx context.Context
func (_e *GamesFacade_Expecter) GetGames(ctx interface{}) *GamesFacade_GetGames_Call {
	return &GamesFacade_GetGames_Call{Call: _e.mock.On("GetGames", ctx)}
}

func (_c *GamesFacade_GetGames_Call) Run(run func(ctx context.Context)) *GamesFacade_GetGames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *GamesFacade_GetGames_Call) Return(_a0 []model.Game, _a1 error) *GamesFacade_GetGames_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetGamesByUserID provides a mock function with given fields: ctx, userID
func (_m *GamesFacade) GetGamesByUserID(ctx context.Context, userID int32) ([]model.Game, error) {
	ret := _m.Called(ctx, userID)

	var r0 []model.Game
	if rf, ok := ret.Get(0).(func(context.Context, int32) []model.Game); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Game)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamesFacade_GetGamesByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGamesByUserID'
type GamesFacade_GetGamesByUserID_Call struct {
	*mock.Call
}

// GetGamesByUserID is a helper method to define mock.On call
//  - ctx context.Context
//  - userID int32
func (_e *GamesFacade_Expecter) GetGamesByUserID(ctx interface{}, userID interface{}) *GamesFacade_GetGamesByUserID_Call {
	return &GamesFacade_GetGamesByUserID_Call{Call: _e.mock.On("GetGamesByUserID", ctx, userID)}
}

func (_c *GamesFacade_GetGamesByUserID_Call) Run(run func(ctx context.Context, userID int32)) *GamesFacade_GetGamesByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *GamesFacade_GetGamesByUserID_Call) Return(_a0 []model.Game, _a1 error) *GamesFacade_GetGamesByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetRegisteredGames provides a mock function with given fields: ctx
func (_m *GamesFacade) GetRegisteredGames(ctx context.Context) ([]model.Game, error) {
	ret := _m.Called(ctx)

	var r0 []model.Game
	if rf, ok := ret.Get(0).(func(context.Context) []model.Game); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Game)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamesFacade_GetRegisteredGames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRegisteredGames'
type GamesFacade_GetRegisteredGames_Call struct {
	*mock.Call
}

// GetRegisteredGames is a helper method to define mock.On call
//  - ctx context.Context
func (_e *GamesFacade_Expecter) GetRegisteredGames(ctx interface{}) *GamesFacade_GetRegisteredGames_Call {
	return &GamesFacade_GetRegisteredGames_Call{Call: _e.mock.On("GetRegisteredGames", ctx)}
}

func (_c *GamesFacade_GetRegisteredGames_Call) Run(run func(ctx context.Context)) *GamesFacade_GetRegisteredGames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *GamesFacade_GetRegisteredGames_Call) Return(_a0 []model.Game, _a1 error) *GamesFacade_GetRegisteredGames_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// RegisterGame provides a mock function with given fields: ctx, gameID
func (_m *GamesFacade) RegisterGame(ctx context.Context, gameID int32) (model.RegisterGameStatus, error) {
	ret := _m.Called(ctx, gameID)

	var r0 model.RegisterGameStatus
	if rf, ok := ret.Get(0).(func(context.Context, int32) model.RegisterGameStatus); ok {
		r0 = rf(ctx, gameID)
	} else {
		r0 = ret.Get(0).(model.RegisterGameStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, gameID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamesFacade_RegisterGame_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterGame'
type GamesFacade_RegisterGame_Call struct {
	*mock.Call
}

// RegisterGame is a helper method to define mock.On call
//  - ctx context.Context
//  - gameID int32
func (_e *GamesFacade_Expecter) RegisterGame(ctx interface{}, gameID interface{}) *GamesFacade_RegisterGame_Call {
	return &GamesFacade_RegisterGame_Call{Call: _e.mock.On("RegisterGame", ctx, gameID)}
}

func (_c *GamesFacade_RegisterGame_Call) Run(run func(ctx context.Context, gameID int32)) *GamesFacade_RegisterGame_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *GamesFacade_RegisterGame_Call) Return(_a0 model.RegisterGameStatus, _a1 error) *GamesFacade_RegisterGame_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UnregisterGame provides a mock function with given fields: ctx, gameID
func (_m *GamesFacade) UnregisterGame(ctx context.Context, gameID int32) (model.UnregisterGameStatus, error) {
	ret := _m.Called(ctx, gameID)

	var r0 model.UnregisterGameStatus
	if rf, ok := ret.Get(0).(func(context.Context, int32) model.UnregisterGameStatus); ok {
		r0 = rf(ctx, gameID)
	} else {
		r0 = ret.Get(0).(model.UnregisterGameStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, gameID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamesFacade_UnregisterGame_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnregisterGame'
type GamesFacade_UnregisterGame_Call struct {
	*mock.Call
}

// UnregisterGame is a helper method to define mock.On call
//  - ctx context.Context
//  - gameID int32
func (_e *GamesFacade_Expecter) UnregisterGame(ctx interface{}, gameID interface{}) *GamesFacade_UnregisterGame_Call {
	return &GamesFacade_UnregisterGame_Call{Call: _e.mock.On("UnregisterGame", ctx, gameID)}
}

func (_c *GamesFacade_UnregisterGame_Call) Run(run func(ctx context.Context, gameID int32)) *GamesFacade_UnregisterGame_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *GamesFacade_UnregisterGame_Call) Return(_a0 model.UnregisterGameStatus, _a1 error) *GamesFacade_UnregisterGame_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdatePayment provides a mock function with given fields: ctx, gameID, payment
func (_m *GamesFacade) UpdatePayment(ctx context.Context, gameID int32, payment int32) error {
	ret := _m.Called(ctx, gameID, payment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) error); ok {
		r0 = rf(ctx, gameID, payment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GamesFacade_UpdatePayment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePayment'
type GamesFacade_UpdatePayment_Call struct {
	*mock.Call
}

// UpdatePayment is a helper method to define mock.On call
//  - ctx context.Context
//  - gameID int32
//  - payment int32
func (_e *GamesFacade_Expecter) UpdatePayment(ctx interface{}, gameID interface{}, payment interface{}) *GamesFacade_UpdatePayment_Call {
	return &GamesFacade_UpdatePayment_Call{Call: _e.mock.On("UpdatePayment", ctx, gameID, payment)}
}

func (_c *GamesFacade_UpdatePayment_Call) Run(run func(ctx context.Context, gameID int32, payment int32)) *GamesFacade_UpdatePayment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].(int32))
	})
	return _c
}

func (_c *GamesFacade_UpdatePayment_Call) Return(_a0 error) *GamesFacade_UpdatePayment_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewGamesFacade interface {
	mock.TestingT
	Cleanup(func())
}

// NewGamesFacade creates a new instance of GamesFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGamesFacade(t mockConstructorTestingTNewGamesFacade) *GamesFacade {
	mock := &GamesFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
