// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	builder "github.com/go-xorm/builder"

	mock "github.com/stretchr/testify/mock"

	mysql "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

// GamePlayerStorage is an autogenerated mock type for the GamePlayerStorage type
type GamePlayerStorage struct {
	mock.Mock
}

type GamePlayerStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *GamePlayerStorage) EXPECT() *GamePlayerStorage_Expecter {
	return &GamePlayerStorage_Expecter{mock: &_m.Mock}
}

// CreateGamePlayer provides a mock function with given fields: ctx, gamePlayer
func (_m *GamePlayerStorage) CreateGamePlayer(ctx context.Context, gamePlayer mysql.GamePlayer) (int, error) {
	ret := _m.Called(ctx, gamePlayer)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, mysql.GamePlayer) (int, error)); ok {
		return rf(ctx, gamePlayer)
	}
	if rf, ok := ret.Get(0).(func(context.Context, mysql.GamePlayer) int); ok {
		r0 = rf(ctx, gamePlayer)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, mysql.GamePlayer) error); ok {
		r1 = rf(ctx, gamePlayer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamePlayerStorage_CreateGamePlayer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateGamePlayer'
type GamePlayerStorage_CreateGamePlayer_Call struct {
	*mock.Call
}

// CreateGamePlayer is a helper method to define mock.On call
//   - ctx context.Context
//   - gamePlayer mysql.GamePlayer
func (_e *GamePlayerStorage_Expecter) CreateGamePlayer(ctx interface{}, gamePlayer interface{}) *GamePlayerStorage_CreateGamePlayer_Call {
	return &GamePlayerStorage_CreateGamePlayer_Call{Call: _e.mock.On("CreateGamePlayer", ctx, gamePlayer)}
}

func (_c *GamePlayerStorage_CreateGamePlayer_Call) Run(run func(ctx context.Context, gamePlayer mysql.GamePlayer)) *GamePlayerStorage_CreateGamePlayer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(mysql.GamePlayer))
	})
	return _c
}

func (_c *GamePlayerStorage_CreateGamePlayer_Call) Return(_a0 int, _a1 error) *GamePlayerStorage_CreateGamePlayer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GamePlayerStorage_CreateGamePlayer_Call) RunAndReturn(run func(context.Context, mysql.GamePlayer) (int, error)) *GamePlayerStorage_CreateGamePlayer_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteGamePlayer provides a mock function with given fields: ctx, id
func (_m *GamePlayerStorage) DeleteGamePlayer(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GamePlayerStorage_DeleteGamePlayer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteGamePlayer'
type GamePlayerStorage_DeleteGamePlayer_Call struct {
	*mock.Call
}

// DeleteGamePlayer is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
func (_e *GamePlayerStorage_Expecter) DeleteGamePlayer(ctx interface{}, id interface{}) *GamePlayerStorage_DeleteGamePlayer_Call {
	return &GamePlayerStorage_DeleteGamePlayer_Call{Call: _e.mock.On("DeleteGamePlayer", ctx, id)}
}

func (_c *GamePlayerStorage_DeleteGamePlayer_Call) Run(run func(ctx context.Context, id int)) *GamePlayerStorage_DeleteGamePlayer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *GamePlayerStorage_DeleteGamePlayer_Call) Return(_a0 error) *GamePlayerStorage_DeleteGamePlayer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GamePlayerStorage_DeleteGamePlayer_Call) RunAndReturn(run func(context.Context, int) error) *GamePlayerStorage_DeleteGamePlayer_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: ctx, q
func (_m *GamePlayerStorage) Find(ctx context.Context, q builder.Cond) ([]mysql.GamePlayer, error) {
	ret := _m.Called(ctx, q)

	var r0 []mysql.GamePlayer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, builder.Cond) ([]mysql.GamePlayer, error)); ok {
		return rf(ctx, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, builder.Cond) []mysql.GamePlayer); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mysql.GamePlayer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, builder.Cond) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamePlayerStorage_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type GamePlayerStorage_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - ctx context.Context
//   - q builder.Cond
func (_e *GamePlayerStorage_Expecter) Find(ctx interface{}, q interface{}) *GamePlayerStorage_Find_Call {
	return &GamePlayerStorage_Find_Call{Call: _e.mock.On("Find", ctx, q)}
}

func (_c *GamePlayerStorage_Find_Call) Run(run func(ctx context.Context, q builder.Cond)) *GamePlayerStorage_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(builder.Cond))
	})
	return _c
}

func (_c *GamePlayerStorage_Find_Call) Return(_a0 []mysql.GamePlayer, _a1 error) *GamePlayerStorage_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GamePlayerStorage_Find_Call) RunAndReturn(run func(context.Context, builder.Cond) ([]mysql.GamePlayer, error)) *GamePlayerStorage_Find_Call {
	_c.Call.Return(run)
	return _c
}

// GetGamePlayer provides a mock function with given fields: ctx, id
func (_m *GamePlayerStorage) GetGamePlayer(ctx context.Context, id int) (*mysql.GamePlayer, error) {
	ret := _m.Called(ctx, id)

	var r0 *mysql.GamePlayer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*mysql.GamePlayer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *mysql.GamePlayer); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mysql.GamePlayer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamePlayerStorage_GetGamePlayer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGamePlayer'
type GamePlayerStorage_GetGamePlayer_Call struct {
	*mock.Call
}

// GetGamePlayer is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
func (_e *GamePlayerStorage_Expecter) GetGamePlayer(ctx interface{}, id interface{}) *GamePlayerStorage_GetGamePlayer_Call {
	return &GamePlayerStorage_GetGamePlayer_Call{Call: _e.mock.On("GetGamePlayer", ctx, id)}
}

func (_c *GamePlayerStorage_GetGamePlayer_Call) Run(run func(ctx context.Context, id int)) *GamePlayerStorage_GetGamePlayer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *GamePlayerStorage_GetGamePlayer_Call) Return(_a0 *mysql.GamePlayer, _a1 error) *GamePlayerStorage_GetGamePlayer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GamePlayerStorage_GetGamePlayer_Call) RunAndReturn(run func(context.Context, int) (*mysql.GamePlayer, error)) *GamePlayerStorage_GetGamePlayer_Call {
	_c.Call.Return(run)
	return _c
}

// PatchGamePlayer provides a mock function with given fields: ctx, gamePlayer
func (_m *GamePlayerStorage) PatchGamePlayer(ctx context.Context, gamePlayer mysql.GamePlayer) error {
	ret := _m.Called(ctx, gamePlayer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mysql.GamePlayer) error); ok {
		r0 = rf(ctx, gamePlayer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GamePlayerStorage_PatchGamePlayer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchGamePlayer'
type GamePlayerStorage_PatchGamePlayer_Call struct {
	*mock.Call
}

// PatchGamePlayer is a helper method to define mock.On call
//   - ctx context.Context
//   - gamePlayer mysql.GamePlayer
func (_e *GamePlayerStorage_Expecter) PatchGamePlayer(ctx interface{}, gamePlayer interface{}) *GamePlayerStorage_PatchGamePlayer_Call {
	return &GamePlayerStorage_PatchGamePlayer_Call{Call: _e.mock.On("PatchGamePlayer", ctx, gamePlayer)}
}

func (_c *GamePlayerStorage_PatchGamePlayer_Call) Run(run func(ctx context.Context, gamePlayer mysql.GamePlayer)) *GamePlayerStorage_PatchGamePlayer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(mysql.GamePlayer))
	})
	return _c
}

func (_c *GamePlayerStorage_PatchGamePlayer_Call) Return(_a0 error) *GamePlayerStorage_PatchGamePlayer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GamePlayerStorage_PatchGamePlayer_Call) RunAndReturn(run func(context.Context, mysql.GamePlayer) error) *GamePlayerStorage_PatchGamePlayer_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewGamePlayerStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewGamePlayerStorage creates a new instance of GamePlayerStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGamePlayerStorage(t mockConstructorTestingTNewGamePlayerStorage) *GamePlayerStorage {
	mock := &GamePlayerStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
