// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// Croupier is an autogenerated mock type for the Croupier type
type Croupier struct {
	mock.Mock
}

type Croupier_Expecter struct {
	mock *mock.Mock
}

func (_m *Croupier) EXPECT() *Croupier_Expecter {
	return &Croupier_Expecter{mock: &_m.Mock}
}

// GetIsLotteryActive provides a mock function with given fields: ctx, game
func (_m *Croupier) GetIsLotteryActive(ctx context.Context, game model.Game) bool {
	ret := _m.Called(ctx, game)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, model.Game) bool); ok {
		r0 = rf(ctx, game)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Croupier_GetIsLotteryActive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIsLotteryActive'
type Croupier_GetIsLotteryActive_Call struct {
	*mock.Call
}

// GetIsLotteryActive is a helper method to define mock.On call
//   - ctx context.Context
//   - game model.Game
func (_e *Croupier_Expecter) GetIsLotteryActive(ctx interface{}, game interface{}) *Croupier_GetIsLotteryActive_Call {
	return &Croupier_GetIsLotteryActive_Call{Call: _e.mock.On("GetIsLotteryActive", ctx, game)}
}

func (_c *Croupier_GetIsLotteryActive_Call) Run(run func(ctx context.Context, game model.Game)) *Croupier_GetIsLotteryActive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Game))
	})
	return _c
}

func (_c *Croupier_GetIsLotteryActive_Call) Return(_a0 bool) *Croupier_GetIsLotteryActive_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Croupier_GetIsLotteryActive_Call) RunAndReturn(run func(context.Context, model.Game) bool) *Croupier_GetIsLotteryActive_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterForLottery provides a mock function with given fields: ctx, game, user
func (_m *Croupier) RegisterForLottery(ctx context.Context, game model.Game, user model.User) (int32, error) {
	ret := _m.Called(ctx, game, user)

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Game, model.User) (int32, error)); ok {
		return rf(ctx, game, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Game, model.User) int32); ok {
		r0 = rf(ctx, game, user)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Game, model.User) error); ok {
		r1 = rf(ctx, game, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Croupier_RegisterForLottery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterForLottery'
type Croupier_RegisterForLottery_Call struct {
	*mock.Call
}

// RegisterForLottery is a helper method to define mock.On call
//   - ctx context.Context
//   - game model.Game
//   - user model.User
func (_e *Croupier_Expecter) RegisterForLottery(ctx interface{}, game interface{}, user interface{}) *Croupier_RegisterForLottery_Call {
	return &Croupier_RegisterForLottery_Call{Call: _e.mock.On("RegisterForLottery", ctx, game, user)}
}

func (_c *Croupier_RegisterForLottery_Call) Run(run func(ctx context.Context, game model.Game, user model.User)) *Croupier_RegisterForLottery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Game), args[2].(model.User))
	})
	return _c
}

func (_c *Croupier_RegisterForLottery_Call) Return(_a0 int32, _a1 error) *Croupier_RegisterForLottery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Croupier_RegisterForLottery_Call) RunAndReturn(run func(context.Context, model.Game, model.User) (int32, error)) *Croupier_RegisterForLottery_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCroupier interface {
	mock.TestingT
	Cleanup(func())
}

// NewCroupier creates a new instance of Croupier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCroupier(t mockConstructorTestingTNewCroupier) *Croupier {
	mock := &Croupier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
