// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// GamePhotoStorage is an autogenerated mock type for the GamePhotoStorage type
type GamePhotoStorage struct {
	mock.Mock
}

type GamePhotoStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *GamePhotoStorage) EXPECT() *GamePhotoStorage_Expecter {
	return &GamePhotoStorage_Expecter{mock: &_m.Mock}
}

// GetGameIDsWithPhotos provides a mock function with given fields: ctx, limit
func (_m *GamePhotoStorage) GetGameIDsWithPhotos(ctx context.Context, limit uint32) ([]int32, error) {
	ret := _m.Called(ctx, limit)

	var r0 []int32
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []int32); ok {
		r0 = rf(ctx, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int32)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamePhotoStorage_GetGameIDsWithPhotos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGameIDsWithPhotos'
type GamePhotoStorage_GetGameIDsWithPhotos_Call struct {
	*mock.Call
}

// GetGameIDsWithPhotos is a helper method to define mock.On call
//  - ctx context.Context
//  - limit uint32
func (_e *GamePhotoStorage_Expecter) GetGameIDsWithPhotos(ctx interface{}, limit interface{}) *GamePhotoStorage_GetGameIDsWithPhotos_Call {
	return &GamePhotoStorage_GetGameIDsWithPhotos_Call{Call: _e.mock.On("GetGameIDsWithPhotos", ctx, limit)}
}

func (_c *GamePhotoStorage_GetGameIDsWithPhotos_Call) Run(run func(ctx context.Context, limit uint32)) *GamePhotoStorage_GetGameIDsWithPhotos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *GamePhotoStorage_GetGameIDsWithPhotos_Call) Return(_a0 []int32, _a1 error) *GamePhotoStorage_GetGameIDsWithPhotos_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetGamePhotosByGameID provides a mock function with given fields: ctx, gameID
func (_m *GamePhotoStorage) GetGamePhotosByGameID(ctx context.Context, gameID int32) ([]model.GamePhoto, error) {
	ret := _m.Called(ctx, gameID)

	var r0 []model.GamePhoto
	if rf, ok := ret.Get(0).(func(context.Context, int32) []model.GamePhoto); ok {
		r0 = rf(ctx, gameID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.GamePhoto)
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

// GamePhotoStorage_GetGamePhotosByGameID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGamePhotosByGameID'
type GamePhotoStorage_GetGamePhotosByGameID_Call struct {
	*mock.Call
}

// GetGamePhotosByGameID is a helper method to define mock.On call
//  - ctx context.Context
//  - gameID int32
func (_e *GamePhotoStorage_Expecter) GetGamePhotosByGameID(ctx interface{}, gameID interface{}) *GamePhotoStorage_GetGamePhotosByGameID_Call {
	return &GamePhotoStorage_GetGamePhotosByGameID_Call{Call: _e.mock.On("GetGamePhotosByGameID", ctx, gameID)}
}

func (_c *GamePhotoStorage_GetGamePhotosByGameID_Call) Run(run func(ctx context.Context, gameID int32)) *GamePhotoStorage_GetGamePhotosByGameID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *GamePhotoStorage_GetGamePhotosByGameID_Call) Return(_a0 []model.GamePhoto, _a1 error) *GamePhotoStorage_GetGamePhotosByGameID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Insert provides a mock function with given fields: ctx, gamePhoto
func (_m *GamePhotoStorage) Insert(ctx context.Context, gamePhoto model.GamePhoto) (int32, error) {
	ret := _m.Called(ctx, gamePhoto)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, model.GamePhoto) int32); ok {
		r0 = rf(ctx, gamePhoto)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.GamePhoto) error); ok {
		r1 = rf(ctx, gamePhoto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GamePhotoStorage_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type GamePhotoStorage_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//  - ctx context.Context
//  - gamePhoto model.GamePhoto
func (_e *GamePhotoStorage_Expecter) Insert(ctx interface{}, gamePhoto interface{}) *GamePhotoStorage_Insert_Call {
	return &GamePhotoStorage_Insert_Call{Call: _e.mock.On("Insert", ctx, gamePhoto)}
}

func (_c *GamePhotoStorage_Insert_Call) Run(run func(ctx context.Context, gamePhoto model.GamePhoto)) *GamePhotoStorage_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.GamePhoto))
	})
	return _c
}

func (_c *GamePhotoStorage_Insert_Call) Return(_a0 int32, _a1 error) *GamePhotoStorage_Insert_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewGamePhotoStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewGamePhotoStorage creates a new instance of GamePhotoStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGamePhotoStorage(t mockConstructorTestingTNewGamePhotoStorage) *GamePhotoStorage {
	mock := &GamePhotoStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
