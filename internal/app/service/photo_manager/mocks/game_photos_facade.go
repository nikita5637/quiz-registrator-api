// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// GamePhotosFacade is an autogenerated mock type for the GamePhotosFacade type
type GamePhotosFacade struct {
	mock.Mock
}

type GamePhotosFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *GamePhotosFacade) EXPECT() *GamePhotosFacade_Expecter {
	return &GamePhotosFacade_Expecter{mock: &_m.Mock}
}

// AddGamePhotos provides a mock function with given fields: ctx, gameID, urls
func (_m *GamePhotosFacade) AddGamePhotos(ctx context.Context, gameID int32, urls []string) error {
	ret := _m.Called(ctx, gameID, urls)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, []string) error); ok {
		r0 = rf(ctx, gameID, urls)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GamePhotosFacade_AddGamePhotos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddGamePhotos'
type GamePhotosFacade_AddGamePhotos_Call struct {
	*mock.Call
}

// AddGamePhotos is a helper method to define mock.On call
//  - ctx context.Context
//  - gameID int32
//  - urls []string
func (_e *GamePhotosFacade_Expecter) AddGamePhotos(ctx interface{}, gameID interface{}, urls interface{}) *GamePhotosFacade_AddGamePhotos_Call {
	return &GamePhotosFacade_AddGamePhotos_Call{Call: _e.mock.On("AddGamePhotos", ctx, gameID, urls)}
}

func (_c *GamePhotosFacade_AddGamePhotos_Call) Run(run func(ctx context.Context, gameID int32, urls []string)) *GamePhotosFacade_AddGamePhotos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].([]string))
	})
	return _c
}

func (_c *GamePhotosFacade_AddGamePhotos_Call) Return(_a0 error) *GamePhotosFacade_AddGamePhotos_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetPhotosByGameID provides a mock function with given fields: ctx, gameID
func (_m *GamePhotosFacade) GetPhotosByGameID(ctx context.Context, gameID int32) ([]string, error) {
	ret := _m.Called(ctx, gameID)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, int32) []string); ok {
		r0 = rf(ctx, gameID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// GamePhotosFacade_GetPhotosByGameID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPhotosByGameID'
type GamePhotosFacade_GetPhotosByGameID_Call struct {
	*mock.Call
}

// GetPhotosByGameID is a helper method to define mock.On call
//  - ctx context.Context
//  - gameID int32
func (_e *GamePhotosFacade_Expecter) GetPhotosByGameID(ctx interface{}, gameID interface{}) *GamePhotosFacade_GetPhotosByGameID_Call {
	return &GamePhotosFacade_GetPhotosByGameID_Call{Call: _e.mock.On("GetPhotosByGameID", ctx, gameID)}
}

func (_c *GamePhotosFacade_GetPhotosByGameID_Call) Run(run func(ctx context.Context, gameID int32)) *GamePhotosFacade_GetPhotosByGameID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *GamePhotosFacade_GetPhotosByGameID_Call) Return(_a0 []string, _a1 error) *GamePhotosFacade_GetPhotosByGameID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewGamePhotosFacade interface {
	mock.TestingT
	Cleanup(func())
}

// NewGamePhotosFacade creates a new instance of GamePhotosFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGamePhotosFacade(t mockConstructorTestingTNewGamePhotosFacade) *GamePhotosFacade {
	mock := &GamePhotosFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
