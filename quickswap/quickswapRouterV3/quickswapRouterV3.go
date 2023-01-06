// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quickswapRouterV3

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IQuickSwapRouterV3ExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type IQuickSwapRouterV3ExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// IQuickSwapRouterV3ExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuickSwapRouterV3ExactInputSingleParams struct {
	TokenIn          common.Address
	TokenOut         common.Address
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
	LimitSqrtPrice   *big.Int
}

// IQuickSwapRouterV3ExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type IQuickSwapRouterV3ExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// IQuickSwapRouterV3ExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuickSwapRouterV3ExactOutputSingleParams struct {
	TokenIn         common.Address
	TokenOut        common.Address
	Fee             *big.Int
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
	LimitSqrtPrice  *big.Int
}

// QuickswapRouterV3MetaData contains all meta data concerning the QuickswapRouterV3 contract.
var QuickswapRouterV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"algebraSwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structIQuickSwapRouterV3.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structIQuickSwapRouterV3.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structIQuickSwapRouterV3.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingleSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structIQuickSwapRouterV3.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structIQuickSwapRouterV3.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// QuickswapRouterV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use QuickswapRouterV3MetaData.ABI instead.
var QuickswapRouterV3ABI = QuickswapRouterV3MetaData.ABI

// QuickswapRouterV3 is an auto generated Go binding around an Ethereum contract.
type QuickswapRouterV3 struct {
	QuickswapRouterV3Caller     // Read-only binding to the contract
	QuickswapRouterV3Transactor // Write-only binding to the contract
	QuickswapRouterV3Filterer   // Log filterer for contract events
}

// QuickswapRouterV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type QuickswapRouterV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuickswapRouterV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type QuickswapRouterV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuickswapRouterV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuickswapRouterV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuickswapRouterV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuickswapRouterV3Session struct {
	Contract     *QuickswapRouterV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// QuickswapRouterV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuickswapRouterV3CallerSession struct {
	Contract *QuickswapRouterV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// QuickswapRouterV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuickswapRouterV3TransactorSession struct {
	Contract     *QuickswapRouterV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// QuickswapRouterV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type QuickswapRouterV3Raw struct {
	Contract *QuickswapRouterV3 // Generic contract binding to access the raw methods on
}

// QuickswapRouterV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuickswapRouterV3CallerRaw struct {
	Contract *QuickswapRouterV3Caller // Generic read-only contract binding to access the raw methods on
}

// QuickswapRouterV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuickswapRouterV3TransactorRaw struct {
	Contract *QuickswapRouterV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewQuickswapRouterV3 creates a new instance of QuickswapRouterV3, bound to a specific deployed contract.
func NewQuickswapRouterV3(address common.Address, backend bind.ContractBackend) (*QuickswapRouterV3, error) {
	contract, err := bindQuickswapRouterV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuickswapRouterV3{QuickswapRouterV3Caller: QuickswapRouterV3Caller{contract: contract}, QuickswapRouterV3Transactor: QuickswapRouterV3Transactor{contract: contract}, QuickswapRouterV3Filterer: QuickswapRouterV3Filterer{contract: contract}}, nil
}

// NewQuickswapRouterV3Caller creates a new read-only instance of QuickswapRouterV3, bound to a specific deployed contract.
func NewQuickswapRouterV3Caller(address common.Address, caller bind.ContractCaller) (*QuickswapRouterV3Caller, error) {
	contract, err := bindQuickswapRouterV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuickswapRouterV3Caller{contract: contract}, nil
}

// NewQuickswapRouterV3Transactor creates a new write-only instance of QuickswapRouterV3, bound to a specific deployed contract.
func NewQuickswapRouterV3Transactor(address common.Address, transactor bind.ContractTransactor) (*QuickswapRouterV3Transactor, error) {
	contract, err := bindQuickswapRouterV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuickswapRouterV3Transactor{contract: contract}, nil
}

// NewQuickswapRouterV3Filterer creates a new log filterer instance of QuickswapRouterV3, bound to a specific deployed contract.
func NewQuickswapRouterV3Filterer(address common.Address, filterer bind.ContractFilterer) (*QuickswapRouterV3Filterer, error) {
	contract, err := bindQuickswapRouterV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuickswapRouterV3Filterer{contract: contract}, nil
}

// bindQuickswapRouterV3 binds a generic wrapper to an already deployed contract.
func bindQuickswapRouterV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuickswapRouterV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuickswapRouterV3 *QuickswapRouterV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuickswapRouterV3.Contract.QuickswapRouterV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuickswapRouterV3 *QuickswapRouterV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.QuickswapRouterV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuickswapRouterV3 *QuickswapRouterV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.QuickswapRouterV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuickswapRouterV3 *QuickswapRouterV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuickswapRouterV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuickswapRouterV3 *QuickswapRouterV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuickswapRouterV3 *QuickswapRouterV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.contract.Transact(opts, method, params...)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_QuickswapRouterV3 *QuickswapRouterV3Transactor) AlgebraSwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _QuickswapRouterV3.contract.Transact(opts, "algebraSwapCallback", amount0Delta, amount1Delta, data)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_QuickswapRouterV3 *QuickswapRouterV3Session) AlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.AlgebraSwapCallback(&_QuickswapRouterV3.TransactOpts, amount0Delta, amount1Delta, data)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_QuickswapRouterV3 *QuickswapRouterV3TransactorSession) AlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.AlgebraSwapCallback(&_QuickswapRouterV3.TransactOpts, amount0Delta, amount1Delta, data)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_QuickswapRouterV3 *QuickswapRouterV3Transactor) ExactInput(opts *bind.TransactOpts, params IQuickSwapRouterV3ExactInputParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_QuickswapRouterV3 *QuickswapRouterV3Session) ExactInput(params IQuickSwapRouterV3ExactInputParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.ExactInput(&_QuickswapRouterV3.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_QuickswapRouterV3 *QuickswapRouterV3TransactorSession) ExactInput(params IQuickSwapRouterV3ExactInputParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.ExactInput(&_QuickswapRouterV3.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xbc651188.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_QuickswapRouterV3 *QuickswapRouterV3Transactor) ExactInputSingle(opts *bind.TransactOpts, params IQuickSwapRouterV3ExactInputSingleParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xbc651188.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_QuickswapRouterV3 *QuickswapRouterV3Session) ExactInputSingle(params IQuickSwapRouterV3ExactInputSingleParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.ExactInputSingle(&_QuickswapRouterV3.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xbc651188.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_QuickswapRouterV3 *QuickswapRouterV3TransactorSession) ExactInputSingle(params IQuickSwapRouterV3ExactInputSingleParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.ExactInputSingle(&_QuickswapRouterV3.TransactOpts, params)
}

// ExactInputSingleSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb87d2524.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_QuickswapRouterV3 *QuickswapRouterV3Transactor) ExactInputSingleSupportingFeeOnTransferTokens(opts *bind.TransactOpts, params IQuickSwapRouterV3ExactInputSingleParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.contract.Transact(opts, "exactInputSingleSupportingFeeOnTransferTokens", params)
}

// ExactInputSingleSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb87d2524.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_QuickswapRouterV3 *QuickswapRouterV3Session) ExactInputSingleSupportingFeeOnTransferTokens(params IQuickSwapRouterV3ExactInputSingleParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.ExactInputSingleSupportingFeeOnTransferTokens(&_QuickswapRouterV3.TransactOpts, params)
}

// ExactInputSingleSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb87d2524.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_QuickswapRouterV3 *QuickswapRouterV3TransactorSession) ExactInputSingleSupportingFeeOnTransferTokens(params IQuickSwapRouterV3ExactInputSingleParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.ExactInputSingleSupportingFeeOnTransferTokens(&_QuickswapRouterV3.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_QuickswapRouterV3 *QuickswapRouterV3Transactor) ExactOutput(opts *bind.TransactOpts, params IQuickSwapRouterV3ExactOutputParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_QuickswapRouterV3 *QuickswapRouterV3Session) ExactOutput(params IQuickSwapRouterV3ExactOutputParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.ExactOutput(&_QuickswapRouterV3.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_QuickswapRouterV3 *QuickswapRouterV3TransactorSession) ExactOutput(params IQuickSwapRouterV3ExactOutputParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.ExactOutput(&_QuickswapRouterV3.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_QuickswapRouterV3 *QuickswapRouterV3Transactor) ExactOutputSingle(opts *bind.TransactOpts, params IQuickSwapRouterV3ExactOutputSingleParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_QuickswapRouterV3 *QuickswapRouterV3Session) ExactOutputSingle(params IQuickSwapRouterV3ExactOutputSingleParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.ExactOutputSingle(&_QuickswapRouterV3.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_QuickswapRouterV3 *QuickswapRouterV3TransactorSession) ExactOutputSingle(params IQuickSwapRouterV3ExactOutputSingleParams) (*types.Transaction, error) {
	return _QuickswapRouterV3.Contract.ExactOutputSingle(&_QuickswapRouterV3.TransactOpts, params)
}
