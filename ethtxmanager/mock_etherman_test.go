// Code generated by mockery. DO NOT EDIT.

package ethtxmanager

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/ethereum/go-ethereum/core/types"
)

// ethermanMock is an autogenerated mock type for the ethermanInterface type
type ethermanMock struct {
	mock.Mock
}

type ethermanMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ethermanMock) EXPECT() *ethermanMock_Expecter {
	return &ethermanMock_Expecter{mock: &_m.Mock}
}

// CheckTxWasMined provides a mock function with given fields: ctx, txHash
func (_m *ethermanMock) CheckTxWasMined(ctx context.Context, txHash common.Hash) (bool, *types.Receipt, error) {
	ret := _m.Called(ctx, txHash)

	if len(ret) == 0 {
		panic("no return value specified for CheckTxWasMined")
	}

	var r0 bool
	var r1 *types.Receipt
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (bool, *types.Receipt, error)); ok {
		return rf(ctx, txHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) bool); ok {
		r0 = rf(ctx, txHash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) *types.Receipt); ok {
		r1 = rf(ctx, txHash)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Receipt)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, common.Hash) error); ok {
		r2 = rf(ctx, txHash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ethermanMock_CheckTxWasMined_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckTxWasMined'
type ethermanMock_CheckTxWasMined_Call struct {
	*mock.Call
}

// CheckTxWasMined is a helper method to define mock.On call
//   - ctx context.Context
//   - txHash common.Hash
func (_e *ethermanMock_Expecter) CheckTxWasMined(ctx interface{}, txHash interface{}) *ethermanMock_CheckTxWasMined_Call {
	return &ethermanMock_CheckTxWasMined_Call{Call: _e.mock.On("CheckTxWasMined", ctx, txHash)}
}

func (_c *ethermanMock_CheckTxWasMined_Call) Run(run func(ctx context.Context, txHash common.Hash)) *ethermanMock_CheckTxWasMined_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *ethermanMock_CheckTxWasMined_Call) Return(_a0 bool, _a1 *types.Receipt, _a2 error) *ethermanMock_CheckTxWasMined_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ethermanMock_CheckTxWasMined_Call) RunAndReturn(run func(context.Context, common.Hash) (bool, *types.Receipt, error)) *ethermanMock_CheckTxWasMined_Call {
	_c.Call.Return(run)
	return _c
}

// CurrentNonce provides a mock function with given fields: ctx, account
func (_m *ethermanMock) CurrentNonce(ctx context.Context, account common.Address) (uint64, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for CurrentNonce")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (uint64, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ethermanMock_CurrentNonce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CurrentNonce'
type ethermanMock_CurrentNonce_Call struct {
	*mock.Call
}

// CurrentNonce is a helper method to define mock.On call
//   - ctx context.Context
//   - account common.Address
func (_e *ethermanMock_Expecter) CurrentNonce(ctx interface{}, account interface{}) *ethermanMock_CurrentNonce_Call {
	return &ethermanMock_CurrentNonce_Call{Call: _e.mock.On("CurrentNonce", ctx, account)}
}

func (_c *ethermanMock_CurrentNonce_Call) Run(run func(ctx context.Context, account common.Address)) *ethermanMock_CurrentNonce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address))
	})
	return _c
}

func (_c *ethermanMock_CurrentNonce_Call) Return(_a0 uint64, _a1 error) *ethermanMock_CurrentNonce_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ethermanMock_CurrentNonce_Call) RunAndReturn(run func(context.Context, common.Address) (uint64, error)) *ethermanMock_CurrentNonce_Call {
	_c.Call.Return(run)
	return _c
}

// EstimateGas provides a mock function with given fields: ctx, from, to, value, data
func (_m *ethermanMock) EstimateGas(ctx context.Context, from common.Address, to *common.Address, value *big.Int, data []byte) (uint64, error) {
	ret := _m.Called(ctx, from, to, value, data)

	if len(ret) == 0 {
		panic("no return value specified for EstimateGas")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *common.Address, *big.Int, []byte) (uint64, error)); ok {
		return rf(ctx, from, to, value, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *common.Address, *big.Int, []byte) uint64); ok {
		r0 = rf(ctx, from, to, value, data)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *common.Address, *big.Int, []byte) error); ok {
		r1 = rf(ctx, from, to, value, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ethermanMock_EstimateGas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EstimateGas'
type ethermanMock_EstimateGas_Call struct {
	*mock.Call
}

// EstimateGas is a helper method to define mock.On call
//   - ctx context.Context
//   - from common.Address
//   - to *common.Address
//   - value *big.Int
//   - data []byte
func (_e *ethermanMock_Expecter) EstimateGas(ctx interface{}, from interface{}, to interface{}, value interface{}, data interface{}) *ethermanMock_EstimateGas_Call {
	return &ethermanMock_EstimateGas_Call{Call: _e.mock.On("EstimateGas", ctx, from, to, value, data)}
}

func (_c *ethermanMock_EstimateGas_Call) Run(run func(ctx context.Context, from common.Address, to *common.Address, value *big.Int, data []byte)) *ethermanMock_EstimateGas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(*common.Address), args[3].(*big.Int), args[4].([]byte))
	})
	return _c
}

func (_c *ethermanMock_EstimateGas_Call) Return(_a0 uint64, _a1 error) *ethermanMock_EstimateGas_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ethermanMock_EstimateGas_Call) RunAndReturn(run func(context.Context, common.Address, *common.Address, *big.Int, []byte) (uint64, error)) *ethermanMock_EstimateGas_Call {
	_c.Call.Return(run)
	return _c
}

// GetRevertMessage provides a mock function with given fields: ctx, tx
func (_m *ethermanMock) GetRevertMessage(ctx context.Context, tx *types.Transaction) (string, error) {
	ret := _m.Called(ctx, tx)

	if len(ret) == 0 {
		panic("no return value specified for GetRevertMessage")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) (string, error)); ok {
		return rf(ctx, tx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) string); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Transaction) error); ok {
		r1 = rf(ctx, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ethermanMock_GetRevertMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRevertMessage'
type ethermanMock_GetRevertMessage_Call struct {
	*mock.Call
}

// GetRevertMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - tx *types.Transaction
func (_e *ethermanMock_Expecter) GetRevertMessage(ctx interface{}, tx interface{}) *ethermanMock_GetRevertMessage_Call {
	return &ethermanMock_GetRevertMessage_Call{Call: _e.mock.On("GetRevertMessage", ctx, tx)}
}

func (_c *ethermanMock_GetRevertMessage_Call) Run(run func(ctx context.Context, tx *types.Transaction)) *ethermanMock_GetRevertMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.Transaction))
	})
	return _c
}

func (_c *ethermanMock_GetRevertMessage_Call) Return(_a0 string, _a1 error) *ethermanMock_GetRevertMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ethermanMock_GetRevertMessage_Call) RunAndReturn(run func(context.Context, *types.Transaction) (string, error)) *ethermanMock_GetRevertMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetTx provides a mock function with given fields: ctx, txHash
func (_m *ethermanMock) GetTx(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	ret := _m.Called(ctx, txHash)

	if len(ret) == 0 {
		panic("no return value specified for GetTx")
	}

	var r0 *types.Transaction
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Transaction, bool, error)); ok {
		return rf(ctx, txHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Transaction); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) bool); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, common.Hash) error); ok {
		r2 = rf(ctx, txHash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ethermanMock_GetTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTx'
type ethermanMock_GetTx_Call struct {
	*mock.Call
}

// GetTx is a helper method to define mock.On call
//   - ctx context.Context
//   - txHash common.Hash
func (_e *ethermanMock_Expecter) GetTx(ctx interface{}, txHash interface{}) *ethermanMock_GetTx_Call {
	return &ethermanMock_GetTx_Call{Call: _e.mock.On("GetTx", ctx, txHash)}
}

func (_c *ethermanMock_GetTx_Call) Run(run func(ctx context.Context, txHash common.Hash)) *ethermanMock_GetTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *ethermanMock_GetTx_Call) Return(_a0 *types.Transaction, _a1 bool, _a2 error) *ethermanMock_GetTx_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ethermanMock_GetTx_Call) RunAndReturn(run func(context.Context, common.Hash) (*types.Transaction, bool, error)) *ethermanMock_GetTx_Call {
	_c.Call.Return(run)
	return _c
}

// GetTxReceipt provides a mock function with given fields: ctx, txHash
func (_m *ethermanMock) GetTxReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ret := _m.Called(ctx, txHash)

	if len(ret) == 0 {
		panic("no return value specified for GetTxReceipt")
	}

	var r0 *types.Receipt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Receipt, error)); ok {
		return rf(ctx, txHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Receipt); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ethermanMock_GetTxReceipt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTxReceipt'
type ethermanMock_GetTxReceipt_Call struct {
	*mock.Call
}

// GetTxReceipt is a helper method to define mock.On call
//   - ctx context.Context
//   - txHash common.Hash
func (_e *ethermanMock_Expecter) GetTxReceipt(ctx interface{}, txHash interface{}) *ethermanMock_GetTxReceipt_Call {
	return &ethermanMock_GetTxReceipt_Call{Call: _e.mock.On("GetTxReceipt", ctx, txHash)}
}

func (_c *ethermanMock_GetTxReceipt_Call) Run(run func(ctx context.Context, txHash common.Hash)) *ethermanMock_GetTxReceipt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *ethermanMock_GetTxReceipt_Call) Return(_a0 *types.Receipt, _a1 error) *ethermanMock_GetTxReceipt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ethermanMock_GetTxReceipt_Call) RunAndReturn(run func(context.Context, common.Hash) (*types.Receipt, error)) *ethermanMock_GetTxReceipt_Call {
	_c.Call.Return(run)
	return _c
}

// PendingNonce provides a mock function with given fields: ctx, account
func (_m *ethermanMock) PendingNonce(ctx context.Context, account common.Address) (uint64, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for PendingNonce")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (uint64, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ethermanMock_PendingNonce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PendingNonce'
type ethermanMock_PendingNonce_Call struct {
	*mock.Call
}

// PendingNonce is a helper method to define mock.On call
//   - ctx context.Context
//   - account common.Address
func (_e *ethermanMock_Expecter) PendingNonce(ctx interface{}, account interface{}) *ethermanMock_PendingNonce_Call {
	return &ethermanMock_PendingNonce_Call{Call: _e.mock.On("PendingNonce", ctx, account)}
}

func (_c *ethermanMock_PendingNonce_Call) Run(run func(ctx context.Context, account common.Address)) *ethermanMock_PendingNonce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address))
	})
	return _c
}

func (_c *ethermanMock_PendingNonce_Call) Return(_a0 uint64, _a1 error) *ethermanMock_PendingNonce_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ethermanMock_PendingNonce_Call) RunAndReturn(run func(context.Context, common.Address) (uint64, error)) *ethermanMock_PendingNonce_Call {
	_c.Call.Return(run)
	return _c
}

// SendTx provides a mock function with given fields: ctx, tx
func (_m *ethermanMock) SendTx(ctx context.Context, tx *types.Transaction) error {
	ret := _m.Called(ctx, tx)

	if len(ret) == 0 {
		panic("no return value specified for SendTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) error); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ethermanMock_SendTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendTx'
type ethermanMock_SendTx_Call struct {
	*mock.Call
}

// SendTx is a helper method to define mock.On call
//   - ctx context.Context
//   - tx *types.Transaction
func (_e *ethermanMock_Expecter) SendTx(ctx interface{}, tx interface{}) *ethermanMock_SendTx_Call {
	return &ethermanMock_SendTx_Call{Call: _e.mock.On("SendTx", ctx, tx)}
}

func (_c *ethermanMock_SendTx_Call) Run(run func(ctx context.Context, tx *types.Transaction)) *ethermanMock_SendTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.Transaction))
	})
	return _c
}

func (_c *ethermanMock_SendTx_Call) Return(_a0 error) *ethermanMock_SendTx_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ethermanMock_SendTx_Call) RunAndReturn(run func(context.Context, *types.Transaction) error) *ethermanMock_SendTx_Call {
	_c.Call.Return(run)
	return _c
}

// SignTx provides a mock function with given fields: ctx, sender, tx
func (_m *ethermanMock) SignTx(ctx context.Context, sender common.Address, tx *types.Transaction) (*types.Transaction, error) {
	ret := _m.Called(ctx, sender, tx)

	if len(ret) == 0 {
		panic("no return value specified for SignTx")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *types.Transaction) (*types.Transaction, error)); ok {
		return rf(ctx, sender, tx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *types.Transaction) *types.Transaction); ok {
		r0 = rf(ctx, sender, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *types.Transaction) error); ok {
		r1 = rf(ctx, sender, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ethermanMock_SignTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignTx'
type ethermanMock_SignTx_Call struct {
	*mock.Call
}

// SignTx is a helper method to define mock.On call
//   - ctx context.Context
//   - sender common.Address
//   - tx *types.Transaction
func (_e *ethermanMock_Expecter) SignTx(ctx interface{}, sender interface{}, tx interface{}) *ethermanMock_SignTx_Call {
	return &ethermanMock_SignTx_Call{Call: _e.mock.On("SignTx", ctx, sender, tx)}
}

func (_c *ethermanMock_SignTx_Call) Run(run func(ctx context.Context, sender common.Address, tx *types.Transaction)) *ethermanMock_SignTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(*types.Transaction))
	})
	return _c
}

func (_c *ethermanMock_SignTx_Call) Return(_a0 *types.Transaction, _a1 error) *ethermanMock_SignTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ethermanMock_SignTx_Call) RunAndReturn(run func(context.Context, common.Address, *types.Transaction) (*types.Transaction, error)) *ethermanMock_SignTx_Call {
	_c.Call.Return(run)
	return _c
}

// SuggestedGasPrice provides a mock function with given fields: ctx
func (_m *ethermanMock) SuggestedGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SuggestedGasPrice")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ethermanMock_SuggestedGasPrice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SuggestedGasPrice'
type ethermanMock_SuggestedGasPrice_Call struct {
	*mock.Call
}

// SuggestedGasPrice is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ethermanMock_Expecter) SuggestedGasPrice(ctx interface{}) *ethermanMock_SuggestedGasPrice_Call {
	return &ethermanMock_SuggestedGasPrice_Call{Call: _e.mock.On("SuggestedGasPrice", ctx)}
}

func (_c *ethermanMock_SuggestedGasPrice_Call) Run(run func(ctx context.Context)) *ethermanMock_SuggestedGasPrice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ethermanMock_SuggestedGasPrice_Call) Return(_a0 *big.Int, _a1 error) *ethermanMock_SuggestedGasPrice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ethermanMock_SuggestedGasPrice_Call) RunAndReturn(run func(context.Context) (*big.Int, error)) *ethermanMock_SuggestedGasPrice_Call {
	_c.Call.Return(run)
	return _c
}

// WaitTxToBeMined provides a mock function with given fields: ctx, tx, timeout
func (_m *ethermanMock) WaitTxToBeMined(ctx context.Context, tx *types.Transaction, timeout time.Duration) (bool, error) {
	ret := _m.Called(ctx, tx, timeout)

	if len(ret) == 0 {
		panic("no return value specified for WaitTxToBeMined")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction, time.Duration) (bool, error)); ok {
		return rf(ctx, tx, timeout)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction, time.Duration) bool); ok {
		r0 = rf(ctx, tx, timeout)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Transaction, time.Duration) error); ok {
		r1 = rf(ctx, tx, timeout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ethermanMock_WaitTxToBeMined_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WaitTxToBeMined'
type ethermanMock_WaitTxToBeMined_Call struct {
	*mock.Call
}

// WaitTxToBeMined is a helper method to define mock.On call
//   - ctx context.Context
//   - tx *types.Transaction
//   - timeout time.Duration
func (_e *ethermanMock_Expecter) WaitTxToBeMined(ctx interface{}, tx interface{}, timeout interface{}) *ethermanMock_WaitTxToBeMined_Call {
	return &ethermanMock_WaitTxToBeMined_Call{Call: _e.mock.On("WaitTxToBeMined", ctx, tx, timeout)}
}

func (_c *ethermanMock_WaitTxToBeMined_Call) Run(run func(ctx context.Context, tx *types.Transaction, timeout time.Duration)) *ethermanMock_WaitTxToBeMined_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.Transaction), args[2].(time.Duration))
	})
	return _c
}

func (_c *ethermanMock_WaitTxToBeMined_Call) Return(_a0 bool, _a1 error) *ethermanMock_WaitTxToBeMined_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ethermanMock_WaitTxToBeMined_Call) RunAndReturn(run func(context.Context, *types.Transaction, time.Duration) (bool, error)) *ethermanMock_WaitTxToBeMined_Call {
	_c.Call.Return(run)
	return _c
}

// newEthermanMock creates a new instance of ethermanMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newEthermanMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ethermanMock {
	mock := &ethermanMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
