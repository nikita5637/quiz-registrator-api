// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// GameResultStorage is an autogenerated mock type for the GameResultStorage type
type GameResultStorage struct {
	mock.Mock
}

type GameResultStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *GameResultStorage) EXPECT() *GameResultStorage_Expecter {
	return &GameResultStorage_Expecter{mock: &_m.Mock}
}

// GetGameResultByFkGameID provides a mock function with given fields: ctx, fkGameID
func (_m *GameResultStorage) GetGameResultByFkGameID(ctx context.Context, fkGameID int32) (model.GameResult, error) {
	ret := _m.Called(ctx, fkGameID)

	var r0 model.GameResult
	if rf, ok := ret.Get(0).(func(context.Context, int32) model.GameResult); ok {
		r0 = rf(ctx, fkGameID)
	} else {
		r0 = ret.Get(0).(model.GameResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, fkGameID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GameResultStorage_GetGameResultByFkGameID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGameResultByFkGameID'
type GameResultStorage_GetGameResultByFkGameID_Call struct {
	*mock.Call
}

// GetGameResultByFkGameID is a helper method to define mock.On call
//  - ctx context.Context
//  - fkGameID int32
func (_e *GameResultStorage_Expecter) GetGameResultByFkGameID(ctx interface{}, fkGameID interface{}) *GameResultStorage_GetGameResultByFkGameID_Call {
	return &GameResultStorage_GetGameResultByFkGameID_Call{Call: _e.mock.On("GetGameResultByFkGameID", ctx, fkGameID)}
}

func (_c *GameResultStorage_GetGameResultByFkGameID_Call) Run(run func(ctx context.Context, fkGameID int32)) *GameResultStorage_GetGameResultByFkGameID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *GameResultStorage_GetGameResultByFkGameID_Call) Return(_a0 model.GameResult, _a1 error) *GameResultStorage_GetGameResultByFkGameID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewGameResultStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewGameResultStorage creates a new instance of GameResultStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGameResultStorage(t mockConstructorTestingTNewGameResultStorage) *GameResultStorage {
	mock := &GameResultStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}