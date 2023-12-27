// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// MathProblemsFacade is an autogenerated mock type for the MathProblemsFacade type
type MathProblemsFacade struct {
	mock.Mock
}

type MathProblemsFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *MathProblemsFacade) EXPECT() *MathProblemsFacade_Expecter {
	return &MathProblemsFacade_Expecter{mock: &_m.Mock}
}

// CreateMathProblem provides a mock function with given fields: ctx, mathProblem
func (_m *MathProblemsFacade) CreateMathProblem(ctx context.Context, mathProblem model.MathProblem) (model.MathProblem, error) {
	ret := _m.Called(ctx, mathProblem)

	var r0 model.MathProblem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.MathProblem) (model.MathProblem, error)); ok {
		return rf(ctx, mathProblem)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.MathProblem) model.MathProblem); ok {
		r0 = rf(ctx, mathProblem)
	} else {
		r0 = ret.Get(0).(model.MathProblem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.MathProblem) error); ok {
		r1 = rf(ctx, mathProblem)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MathProblemsFacade_CreateMathProblem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMathProblem'
type MathProblemsFacade_CreateMathProblem_Call struct {
	*mock.Call
}

// CreateMathProblem is a helper method to define mock.On call
//   - ctx context.Context
//   - mathProblem model.MathProblem
func (_e *MathProblemsFacade_Expecter) CreateMathProblem(ctx interface{}, mathProblem interface{}) *MathProblemsFacade_CreateMathProblem_Call {
	return &MathProblemsFacade_CreateMathProblem_Call{Call: _e.mock.On("CreateMathProblem", ctx, mathProblem)}
}

func (_c *MathProblemsFacade_CreateMathProblem_Call) Run(run func(ctx context.Context, mathProblem model.MathProblem)) *MathProblemsFacade_CreateMathProblem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.MathProblem))
	})
	return _c
}

func (_c *MathProblemsFacade_CreateMathProblem_Call) Return(_a0 model.MathProblem, _a1 error) *MathProblemsFacade_CreateMathProblem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MathProblemsFacade_CreateMathProblem_Call) RunAndReturn(run func(context.Context, model.MathProblem) (model.MathProblem, error)) *MathProblemsFacade_CreateMathProblem_Call {
	_c.Call.Return(run)
	return _c
}

// GetMathProblemByGameID provides a mock function with given fields: ctx, gameID
func (_m *MathProblemsFacade) GetMathProblemByGameID(ctx context.Context, gameID int32) (model.MathProblem, error) {
	ret := _m.Called(ctx, gameID)

	var r0 model.MathProblem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (model.MathProblem, error)); ok {
		return rf(ctx, gameID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) model.MathProblem); ok {
		r0 = rf(ctx, gameID)
	} else {
		r0 = ret.Get(0).(model.MathProblem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, gameID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MathProblemsFacade_GetMathProblemByGameID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMathProblemByGameID'
type MathProblemsFacade_GetMathProblemByGameID_Call struct {
	*mock.Call
}

// GetMathProblemByGameID is a helper method to define mock.On call
//   - ctx context.Context
//   - gameID int32
func (_e *MathProblemsFacade_Expecter) GetMathProblemByGameID(ctx interface{}, gameID interface{}) *MathProblemsFacade_GetMathProblemByGameID_Call {
	return &MathProblemsFacade_GetMathProblemByGameID_Call{Call: _e.mock.On("GetMathProblemByGameID", ctx, gameID)}
}

func (_c *MathProblemsFacade_GetMathProblemByGameID_Call) Run(run func(ctx context.Context, gameID int32)) *MathProblemsFacade_GetMathProblemByGameID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MathProblemsFacade_GetMathProblemByGameID_Call) Return(_a0 model.MathProblem, _a1 error) *MathProblemsFacade_GetMathProblemByGameID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MathProblemsFacade_GetMathProblemByGameID_Call) RunAndReturn(run func(context.Context, int32) (model.MathProblem, error)) *MathProblemsFacade_GetMathProblemByGameID_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMathProblemsFacade interface {
	mock.TestingT
	Cleanup(func())
}

// NewMathProblemsFacade creates a new instance of MathProblemsFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMathProblemsFacade(t mockConstructorTestingTNewMathProblemsFacade) *MathProblemsFacade {
	mock := &MathProblemsFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
