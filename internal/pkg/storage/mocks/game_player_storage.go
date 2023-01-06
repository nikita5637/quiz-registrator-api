// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	builder "github.com/go-xorm/builder"

	mock "github.com/stretchr/testify/mock"

	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
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

// Delete provides a mock function with given fields: ctx, id
func (_m *GamePlayerStorage) Delete(ctx context.Context, id int32) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GamePlayerStorage_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type GamePlayerStorage_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - ctx context.Context
//  - id int32
func (_e *GamePlayerStorage_Expecter) Delete(ctx interface{}, id interface{}) *GamePlayerStorage_Delete_Call {
	return &GamePlayerStorage_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *GamePlayerStorage_Delete_Call) Run(run func(ctx context.Context, id int32)) *GamePlayerStorage_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *GamePlayerStorage_Delete_Call) Return(_a0 error) *GamePlayerStorage_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Find provides a mock function with given fields: ctx, q
func (_m *GamePlayerStorage) Find(ctx context.Context, q builder.Cond) ([]model.GamePlayer, error) {
	ret := _m.Called(ctx, q)

	var r0 []model.GamePlayer
	if rf, ok := ret.Get(0).(func(context.Context, builder.Cond) []model.GamePlayer); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.GamePlayer)
		}
	}

	var r1 error
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
//  - ctx context.Context
//  - q builder.Cond
func (_e *GamePlayerStorage_Expecter) Find(ctx interface{}, q interface{}) *GamePlayerStorage_Find_Call {
	return &GamePlayerStorage_Find_Call{Call: _e.mock.On("Find", ctx, q)}
}

func (_c *GamePlayerStorage_Find_Call) Run(run func(ctx context.Context, q builder.Cond)) *GamePlayerStorage_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(builder.Cond))
	})
	return _c
}

func (_c *GamePlayerStorage_Find_Call) Return(_a0 []model.GamePlayer, _a1 error) *GamePlayerStorage_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Insert provides a mock function with given fields: ctx, gamePlayer
func (_m *GamePlayerStorage) Insert(ctx context.Context, gamePlayer model.GamePlayer) (int32, error) {
	ret := _m.Called(ctx, gamePlayer)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, model.GamePlayer) int32); ok {
		r0 = rf(ctx, gamePlayer)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.GamePlayer) error); ok {
		r1 = rf(ctx, gamePlayer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamePlayerStorage_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type GamePlayerStorage_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//  - ctx context.Context
//  - gamePlayer model.GamePlayer
func (_e *GamePlayerStorage_Expecter) Insert(ctx interface{}, gamePlayer interface{}) *GamePlayerStorage_Insert_Call {
	return &GamePlayerStorage_Insert_Call{Call: _e.mock.On("Insert", ctx, gamePlayer)}
}

func (_c *GamePlayerStorage_Insert_Call) Run(run func(ctx context.Context, gamePlayer model.GamePlayer)) *GamePlayerStorage_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.GamePlayer))
	})
	return _c
}

func (_c *GamePlayerStorage_Insert_Call) Return(_a0 int32, _a1 error) *GamePlayerStorage_Insert_Call {
	_c.Call.Return(_a0, _a1)
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
