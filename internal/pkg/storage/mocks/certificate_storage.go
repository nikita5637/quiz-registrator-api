// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mysql "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	mock "github.com/stretchr/testify/mock"
)

// CertificateStorage is an autogenerated mock type for the CertificateStorage type
type CertificateStorage struct {
	mock.Mock
}

type CertificateStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *CertificateStorage) EXPECT() *CertificateStorage_Expecter {
	return &CertificateStorage_Expecter{mock: &_m.Mock}
}

// CreateCertificate provides a mock function with given fields: ctx, dbCertificate
func (_m *CertificateStorage) CreateCertificate(ctx context.Context, dbCertificate mysql.Certificate) (int32, error) {
	ret := _m.Called(ctx, dbCertificate)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, mysql.Certificate) int32); ok {
		r0 = rf(ctx, dbCertificate)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, mysql.Certificate) error); ok {
		r1 = rf(ctx, dbCertificate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CertificateStorage_CreateCertificate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCertificate'
type CertificateStorage_CreateCertificate_Call struct {
	*mock.Call
}

// CreateCertificate is a helper method to define mock.On call
//  - ctx context.Context
//  - dbCertificate mysql.Certificate
func (_e *CertificateStorage_Expecter) CreateCertificate(ctx interface{}, dbCertificate interface{}) *CertificateStorage_CreateCertificate_Call {
	return &CertificateStorage_CreateCertificate_Call{Call: _e.mock.On("CreateCertificate", ctx, dbCertificate)}
}

func (_c *CertificateStorage_CreateCertificate_Call) Run(run func(ctx context.Context, dbCertificate mysql.Certificate)) *CertificateStorage_CreateCertificate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(mysql.Certificate))
	})
	return _c
}

func (_c *CertificateStorage_CreateCertificate_Call) Return(_a0 int32, _a1 error) *CertificateStorage_CreateCertificate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteCertificate provides a mock function with given fields: ctx, id
func (_m *CertificateStorage) DeleteCertificate(ctx context.Context, id int32) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CertificateStorage_DeleteCertificate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCertificate'
type CertificateStorage_DeleteCertificate_Call struct {
	*mock.Call
}

// DeleteCertificate is a helper method to define mock.On call
//  - ctx context.Context
//  - id int32
func (_e *CertificateStorage_Expecter) DeleteCertificate(ctx interface{}, id interface{}) *CertificateStorage_DeleteCertificate_Call {
	return &CertificateStorage_DeleteCertificate_Call{Call: _e.mock.On("DeleteCertificate", ctx, id)}
}

func (_c *CertificateStorage_DeleteCertificate_Call) Run(run func(ctx context.Context, id int32)) *CertificateStorage_DeleteCertificate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *CertificateStorage_DeleteCertificate_Call) Return(_a0 error) *CertificateStorage_DeleteCertificate_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetCertificateByID provides a mock function with given fields: ctx, id
func (_m *CertificateStorage) GetCertificateByID(ctx context.Context, id int32) (*mysql.Certificate, error) {
	ret := _m.Called(ctx, id)

	var r0 *mysql.Certificate
	if rf, ok := ret.Get(0).(func(context.Context, int32) *mysql.Certificate); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mysql.Certificate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CertificateStorage_GetCertificateByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCertificateByID'
type CertificateStorage_GetCertificateByID_Call struct {
	*mock.Call
}

// GetCertificateByID is a helper method to define mock.On call
//  - ctx context.Context
//  - id int32
func (_e *CertificateStorage_Expecter) GetCertificateByID(ctx interface{}, id interface{}) *CertificateStorage_GetCertificateByID_Call {
	return &CertificateStorage_GetCertificateByID_Call{Call: _e.mock.On("GetCertificateByID", ctx, id)}
}

func (_c *CertificateStorage_GetCertificateByID_Call) Run(run func(ctx context.Context, id int32)) *CertificateStorage_GetCertificateByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *CertificateStorage_GetCertificateByID_Call) Return(_a0 *mysql.Certificate, _a1 error) *CertificateStorage_GetCertificateByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCertificates provides a mock function with given fields: ctx
func (_m *CertificateStorage) GetCertificates(ctx context.Context) ([]mysql.Certificate, error) {
	ret := _m.Called(ctx)

	var r0 []mysql.Certificate
	if rf, ok := ret.Get(0).(func(context.Context) []mysql.Certificate); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mysql.Certificate)
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

// CertificateStorage_GetCertificates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCertificates'
type CertificateStorage_GetCertificates_Call struct {
	*mock.Call
}

// GetCertificates is a helper method to define mock.On call
//  - ctx context.Context
func (_e *CertificateStorage_Expecter) GetCertificates(ctx interface{}) *CertificateStorage_GetCertificates_Call {
	return &CertificateStorage_GetCertificates_Call{Call: _e.mock.On("GetCertificates", ctx)}
}

func (_c *CertificateStorage_GetCertificates_Call) Run(run func(ctx context.Context)) *CertificateStorage_GetCertificates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *CertificateStorage_GetCertificates_Call) Return(_a0 []mysql.Certificate, _a1 error) *CertificateStorage_GetCertificates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// PatchCertificate provides a mock function with given fields: ctx, dbCertificate
func (_m *CertificateStorage) PatchCertificate(ctx context.Context, dbCertificate mysql.Certificate) error {
	ret := _m.Called(ctx, dbCertificate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mysql.Certificate) error); ok {
		r0 = rf(ctx, dbCertificate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CertificateStorage_PatchCertificate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchCertificate'
type CertificateStorage_PatchCertificate_Call struct {
	*mock.Call
}

// PatchCertificate is a helper method to define mock.On call
//  - ctx context.Context
//  - dbCertificate mysql.Certificate
func (_e *CertificateStorage_Expecter) PatchCertificate(ctx interface{}, dbCertificate interface{}) *CertificateStorage_PatchCertificate_Call {
	return &CertificateStorage_PatchCertificate_Call{Call: _e.mock.On("PatchCertificate", ctx, dbCertificate)}
}

func (_c *CertificateStorage_PatchCertificate_Call) Run(run func(ctx context.Context, dbCertificate mysql.Certificate)) *CertificateStorage_PatchCertificate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(mysql.Certificate))
	})
	return _c
}

func (_c *CertificateStorage_PatchCertificate_Call) Return(_a0 error) *CertificateStorage_PatchCertificate_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewCertificateStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewCertificateStorage creates a new instance of CertificateStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCertificateStorage(t mockConstructorTestingTNewCertificateStorage) *CertificateStorage {
	mock := &CertificateStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
