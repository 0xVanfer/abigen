// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapRouter

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

// ISwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// ISwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// ISwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// UniswapRouterMetaData contains all meta data concerning the UniswapRouter contract.
var UniswapRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniswapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapRouterMetaData.ABI instead.
var UniswapRouterABI = UniswapRouterMetaData.ABI

// UniswapRouter is an auto generated Go binding around an Ethereum contract.
type UniswapRouter struct {
	UniswapRouterCaller     // Read-only binding to the contract
	UniswapRouterTransactor // Write-only binding to the contract
	UniswapRouterFilterer   // Log filterer for contract events
}

// UniswapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapRouterSession struct {
	Contract     *UniswapRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapRouterCallerSession struct {
	Contract *UniswapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UniswapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapRouterTransactorSession struct {
	Contract     *UniswapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UniswapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapRouterRaw struct {
	Contract *UniswapRouter // Generic contract binding to access the raw methods on
}

// UniswapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapRouterCallerRaw struct {
	Contract *UniswapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapRouterTransactorRaw struct {
	Contract *UniswapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapRouter creates a new instance of UniswapRouter, bound to a specific deployed contract.
func NewUniswapRouter(address common.Address, backend bind.ContractBackend) (*UniswapRouter, error) {
	contract, err := bindUniswapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapRouter{UniswapRouterCaller: UniswapRouterCaller{contract: contract}, UniswapRouterTransactor: UniswapRouterTransactor{contract: contract}, UniswapRouterFilterer: UniswapRouterFilterer{contract: contract}}, nil
}

// NewUniswapRouterCaller creates a new read-only instance of UniswapRouter, bound to a specific deployed contract.
func NewUniswapRouterCaller(address common.Address, caller bind.ContractCaller) (*UniswapRouterCaller, error) {
	contract, err := bindUniswapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapRouterCaller{contract: contract}, nil
}

// NewUniswapRouterTransactor creates a new write-only instance of UniswapRouter, bound to a specific deployed contract.
func NewUniswapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapRouterTransactor, error) {
	contract, err := bindUniswapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapRouterTransactor{contract: contract}, nil
}

// NewUniswapRouterFilterer creates a new log filterer instance of UniswapRouter, bound to a specific deployed contract.
func NewUniswapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapRouterFilterer, error) {
	contract, err := bindUniswapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapRouterFilterer{contract: contract}, nil
}

// bindUniswapRouter binds a generic wrapper to an already deployed contract.
func bindUniswapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapRouter *UniswapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapRouter.Contract.UniswapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapRouter *UniswapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapRouter.Contract.UniswapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapRouter *UniswapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapRouter.Contract.UniswapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapRouter *UniswapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapRouter *UniswapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapRouter *UniswapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapRouter.Contract.contract.Transact(opts, method, params...)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_UniswapRouter *UniswapRouterTransactor) ExactInput(opts *bind.TransactOpts, params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapRouter.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_UniswapRouter *UniswapRouterSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapRouter.Contract.ExactInput(&_UniswapRouter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_UniswapRouter *UniswapRouterTransactorSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapRouter.Contract.ExactInput(&_UniswapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_UniswapRouter *UniswapRouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapRouter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_UniswapRouter *UniswapRouterSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapRouter.Contract.ExactInputSingle(&_UniswapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_UniswapRouter *UniswapRouterTransactorSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapRouter.Contract.ExactInputSingle(&_UniswapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_UniswapRouter *UniswapRouterTransactor) ExactOutput(opts *bind.TransactOpts, params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapRouter.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_UniswapRouter *UniswapRouterSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapRouter.Contract.ExactOutput(&_UniswapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_UniswapRouter *UniswapRouterTransactorSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapRouter.Contract.ExactOutput(&_UniswapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_UniswapRouter *UniswapRouterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _UniswapRouter.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_UniswapRouter *UniswapRouterSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _UniswapRouter.Contract.ExactOutputSingle(&_UniswapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_UniswapRouter *UniswapRouterTransactorSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _UniswapRouter.Contract.ExactOutputSingle(&_UniswapRouter.TransactOpts, params)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapRouter *UniswapRouterTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapRouter.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapRouter *UniswapRouterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapRouter.Contract.UniswapV3SwapCallback(&_UniswapRouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_UniswapRouter *UniswapRouterTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapRouter.Contract.UniswapV3SwapCallback(&_UniswapRouter.TransactOpts, amount0Delta, amount1Delta, data)
}
