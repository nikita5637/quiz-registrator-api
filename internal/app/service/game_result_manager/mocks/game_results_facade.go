// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GameResultsFacade is an autogenerated mock type for the GameResultsFacade type
type GameResultsFacade struct {
	mock.Mock
}

type GameResultsFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *GameResultsFacade) EXPECT() *GameResultsFacade_Expecter {
	return &GameResultsFacade_Expecter{mock: &_m.Mock}
}

// CreateGameResult provides a mock function with given fields: ctx, gameResult
func (_m *GameResultsFacade) CreateGameResult(ctx context.Context, gameResult model.GameResult) (model.GameResult, error) {
	ret := _m.Called(ctx, gameResult)

	var r0 model.GameResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GameResult) (model.GameResult, error)); ok {
		return rf(ctx, gameResult)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GameResult) model.GameResult); ok {
		r0 = rf(ctx, gameResult)
	} else {
		r0 = ret.Get(0).(model.GameResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GameResult) error); ok {
		r1 = rf(ctx, gameResult)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GameResultsFacade_CreateGameResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateGameResult'
type GameResultsFacade_CreateGameResult_Call struct {
	*mock.Call
}

// CreateGameResult is a helper method to define mock.On call
//   - ctx context.Context
//   - gameResult model.GameResult
func (_e *GameResultsFacade_Expecter) CreateGameResult(ctx interface{}, gameResult interface{}) *GameResultsFacade_CreateGameResult_Call {
	return &GameResultsFacade_CreateGameResult_Call{Call: _e.mock.On("CreateGameResult", ctx, gameResult)}
}

func (_c *GameResultsFacade_CreateGameResult_Call) Run(run func(ctx context.Context, gameResult model.GameResult)) *GameResultsFacade_CreateGameResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.GameResult))
	})
	return _c
}

func (_c *GameResultsFacade_CreateGameResult_Call) Return(_a0 model.GameResult, _a1 error) *GameResultsFacade_CreateGameResult_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GameResultsFacade_CreateGameResult_Call) RunAndReturn(run func(context.Context, model.GameResult) (model.GameResult, error)) *GameResultsFacade_CreateGameResult_Call {
	_c.Call.Return(run)
	return _c
}

// ListGameResults provides a mock function with given fields: ctx
func (_m *GameResultsFacade) ListGameResults(ctx context.Context) ([]model.GameResult, error) {
	ret := _m.Called(ctx)

	var r0 []model.GameResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.GameResult, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.GameResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.GameResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GameResultsFacade_ListGameResults_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListGameResults'
type GameResultsFacade_ListGameResults_Call struct {
	*mock.Call
}

// ListGameResults is a helper method to define mock.On call
//   - ctx context.Context
func (_e *GameResultsFacade_Expecter) ListGameResults(ctx interface{}) *GameResultsFacade_ListGameResults_Call {
	return &GameResultsFacade_ListGameResults_Call{Call: _e.mock.On("ListGameResults", ctx)}
}

func (_c *GameResultsFacade_ListGameResults_Call) Run(run func(ctx context.Context)) *GameResultsFacade_ListGameResults_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *GameResultsFacade_ListGameResults_Call) Return(_a0 []model.GameResult, _a1 error) *GameResultsFacade_ListGameResults_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GameResultsFacade_ListGameResults_Call) RunAndReturn(run func(context.Context) ([]model.GameResult, error)) *GameResultsFacade_ListGameResults_Call {
	_c.Call.Return(run)
	return _c
}

// PatchGameResult provides a mock function with given fields: ctx, gameResult, paths
func (_m *GameResultsFacade) PatchGameResult(ctx context.Context, gameResult model.GameResult, paths []string) (model.GameResult, error) {
	ret := _m.Called(ctx, gameResult, paths)

	var r0 model.GameResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GameResult, []string) (model.GameResult, error)); ok {
		return rf(ctx, gameResult, paths)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GameResult, []string) model.GameResult); ok {
		r0 = rf(ctx, gameResult, paths)
	} else {
		r0 = ret.Get(0).(model.GameResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GameResult, []string) error); ok {
		r1 = rf(ctx, gameResult, paths)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GameResultsFacade_PatchGameResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchGameResult'
type GameResultsFacade_PatchGameResult_Call struct {
	*mock.Call
}

// PatchGameResult is a helper method to define mock.On call
//   - ctx context.Context
//   - gameResult model.GameResult
//   - paths []string
func (_e *GameResultsFacade_Expecter) PatchGameResult(ctx interface{}, gameResult interface{}, paths interface{}) *GameResultsFacade_PatchGameResult_Call {
	return &GameResultsFacade_PatchGameResult_Call{Call: _e.mock.On("PatchGameResult", ctx, gameResult, paths)}
}

func (_c *GameResultsFacade_PatchGameResult_Call) Run(run func(ctx context.Context, gameResult model.GameResult, paths []string)) *GameResultsFacade_PatchGameResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.GameResult), args[2].([]string))
	})
	return _c
}

func (_c *GameResultsFacade_PatchGameResult_Call) Return(_a0 model.GameResult, _a1 error) *GameResultsFacade_PatchGameResult_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GameResultsFacade_PatchGameResult_Call) RunAndReturn(run func(context.Context, model.GameResult, []string) (model.GameResult, error)) *GameResultsFacade_PatchGameResult_Call {
	_c.Call.Return(run)
	return _c
}

// SearchGameResultByGameID provides a mock function with given fields: ctx, gameID
func (_m *GameResultsFacade) SearchGameResultByGameID(ctx context.Context, gameID int32) (model.GameResult, error) {
	ret := _m.Called(ctx, gameID)

	var r0 model.GameResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (model.GameResult, error)); ok {
		return rf(ctx, gameID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) model.GameResult); ok {
		r0 = rf(ctx, gameID)
	} else {
		r0 = ret.Get(0).(model.GameResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, gameID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GameResultsFacade_SearchGameResultByGameID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchGameResultByGameID'
type GameResultsFacade_SearchGameResultByGameID_Call struct {
	*mock.Call
}

// SearchGameResultByGameID is a helper method to define mock.On call
//   - ctx context.Context
//   - gameID int32
func (_e *GameResultsFacade_Expecter) SearchGameResultByGameID(ctx interface{}, gameID interface{}) *GameResultsFacade_SearchGameResultByGameID_Call {
	return &GameResultsFacade_SearchGameResultByGameID_Call{Call: _e.mock.On("SearchGameResultByGameID", ctx, gameID)}
}

func (_c *GameResultsFacade_SearchGameResultByGameID_Call) Run(run func(ctx context.Context, gameID int32)) *GameResultsFacade_SearchGameResultByGameID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *GameResultsFacade_SearchGameResultByGameID_Call) Return(_a0 model.GameResult, _a1 error) *GameResultsFacade_SearchGameResultByGameID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GameResultsFacade_SearchGameResultByGameID_Call) RunAndReturn(run func(context.Context, int32) (model.GameResult, error)) *GameResultsFacade_SearchGameResultByGameID_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewGameResultsFacade interface {
	mock.TestingT
	Cleanup(func())
}

// NewGameResultsFacade creates a new instance of GameResultsFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGameResultsFacade(t mockConstructorTestingTNewGameResultsFacade) *GameResultsFacade {
	mock := &GameResultsFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
