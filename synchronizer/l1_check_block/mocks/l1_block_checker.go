// Code generated by mockery. DO NOT EDIT.

package mock_l1_check_block

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// L1BlockChecker is an autogenerated mock type for the L1BlockChecker type
type L1BlockChecker struct {
	mock.Mock
}

type L1BlockChecker_Expecter struct {
	mock *mock.Mock
}

func (_m *L1BlockChecker) EXPECT() *L1BlockChecker_Expecter {
	return &L1BlockChecker_Expecter{mock: &_m.Mock}
}

// Step provides a mock function with given fields: ctx
func (_m *L1BlockChecker) Step(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Step")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// L1BlockChecker_Step_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Step'
type L1BlockChecker_Step_Call struct {
	*mock.Call
}

// Step is a helper method to define mock.On call
//   - ctx context.Context
func (_e *L1BlockChecker_Expecter) Step(ctx interface{}) *L1BlockChecker_Step_Call {
	return &L1BlockChecker_Step_Call{Call: _e.mock.On("Step", ctx)}
}

func (_c *L1BlockChecker_Step_Call) Run(run func(ctx context.Context)) *L1BlockChecker_Step_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *L1BlockChecker_Step_Call) Return(_a0 error) *L1BlockChecker_Step_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *L1BlockChecker_Step_Call) RunAndReturn(run func(context.Context) error) *L1BlockChecker_Step_Call {
	_c.Call.Return(run)
	return _c
}

// NewL1BlockChecker creates a new instance of L1BlockChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewL1BlockChecker(t interface {
	mock.TestingT
	Cleanup(func())
}) *L1BlockChecker {
	mock := &L1BlockChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}