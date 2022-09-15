// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KyberZap

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

// KyberZapMetaData contains all meta data concerning the KyberZap contract.
var KyberZapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minLpQty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"zapIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minLpQty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"zapInEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"zapOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"zapOutPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KyberZapABI is the input ABI used to generate the binding from.
// Deprecated: Use KyberZapMetaData.ABI instead.
var KyberZapABI = KyberZapMetaData.ABI

// KyberZap is an auto generated Go binding around an Ethereum contract.
type KyberZap struct {
	KyberZapCaller     // Read-only binding to the contract
	KyberZapTransactor // Write-only binding to the contract
	KyberZapFilterer   // Log filterer for contract events
}

// KyberZapCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberZapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberZapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberZapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberZapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberZapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberZapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberZapSession struct {
	Contract     *KyberZap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KyberZapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberZapCallerSession struct {
	Contract *KyberZapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// KyberZapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberZapTransactorSession struct {
	Contract     *KyberZapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KyberZapRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberZapRaw struct {
	Contract *KyberZap // Generic contract binding to access the raw methods on
}

// KyberZapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberZapCallerRaw struct {
	Contract *KyberZapCaller // Generic read-only contract binding to access the raw methods on
}

// KyberZapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberZapTransactorRaw struct {
	Contract *KyberZapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberZap creates a new instance of KyberZap, bound to a specific deployed contract.
func NewKyberZap(address common.Address, backend bind.ContractBackend) (*KyberZap, error) {
	contract, err := bindKyberZap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KyberZap{KyberZapCaller: KyberZapCaller{contract: contract}, KyberZapTransactor: KyberZapTransactor{contract: contract}, KyberZapFilterer: KyberZapFilterer{contract: contract}}, nil
}

// NewKyberZapCaller creates a new read-only instance of KyberZap, bound to a specific deployed contract.
func NewKyberZapCaller(address common.Address, caller bind.ContractCaller) (*KyberZapCaller, error) {
	contract, err := bindKyberZap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberZapCaller{contract: contract}, nil
}

// NewKyberZapTransactor creates a new write-only instance of KyberZap, bound to a specific deployed contract.
func NewKyberZapTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberZapTransactor, error) {
	contract, err := bindKyberZap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberZapTransactor{contract: contract}, nil
}

// NewKyberZapFilterer creates a new log filterer instance of KyberZap, bound to a specific deployed contract.
func NewKyberZapFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberZapFilterer, error) {
	contract, err := bindKyberZap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberZapFilterer{contract: contract}, nil
}

// bindKyberZap binds a generic wrapper to an already deployed contract.
func bindKyberZap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberZapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberZap *KyberZapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberZap.Contract.KyberZapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberZap *KyberZapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberZap.Contract.KyberZapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberZap *KyberZapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberZap.Contract.KyberZapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberZap *KyberZapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberZap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberZap *KyberZapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberZap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberZap *KyberZapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberZap.Contract.contract.Transact(opts, method, params...)
}

// ZapIn is a paid mutator transaction binding the contract method 0x5ce9303a.
//
// Solidity: function zapIn(address tokenIn, address tokenOut, uint256 userIn, address pool, address to, uint256 minLpQty, uint256 deadline) returns()
func (_KyberZap *KyberZapTransactor) ZapIn(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, userIn *big.Int, pool common.Address, to common.Address, minLpQty *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _KyberZap.contract.Transact(opts, "zapIn", tokenIn, tokenOut, userIn, pool, to, minLpQty, deadline)
}

// ZapIn is a paid mutator transaction binding the contract method 0x5ce9303a.
//
// Solidity: function zapIn(address tokenIn, address tokenOut, uint256 userIn, address pool, address to, uint256 minLpQty, uint256 deadline) returns()
func (_KyberZap *KyberZapSession) ZapIn(tokenIn common.Address, tokenOut common.Address, userIn *big.Int, pool common.Address, to common.Address, minLpQty *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _KyberZap.Contract.ZapIn(&_KyberZap.TransactOpts, tokenIn, tokenOut, userIn, pool, to, minLpQty, deadline)
}

// ZapIn is a paid mutator transaction binding the contract method 0x5ce9303a.
//
// Solidity: function zapIn(address tokenIn, address tokenOut, uint256 userIn, address pool, address to, uint256 minLpQty, uint256 deadline) returns()
func (_KyberZap *KyberZapTransactorSession) ZapIn(tokenIn common.Address, tokenOut common.Address, userIn *big.Int, pool common.Address, to common.Address, minLpQty *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _KyberZap.Contract.ZapIn(&_KyberZap.TransactOpts, tokenIn, tokenOut, userIn, pool, to, minLpQty, deadline)
}

// ZapInEth is a paid mutator transaction binding the contract method 0x71679309.
//
// Solidity: function zapInEth(address tokenOut, address pool, address to, uint256 minLpQty, uint256 deadline) payable returns()
func (_KyberZap *KyberZapTransactor) ZapInEth(opts *bind.TransactOpts, tokenOut common.Address, pool common.Address, to common.Address, minLpQty *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _KyberZap.contract.Transact(opts, "zapInEth", tokenOut, pool, to, minLpQty, deadline)
}

// ZapInEth is a paid mutator transaction binding the contract method 0x71679309.
//
// Solidity: function zapInEth(address tokenOut, address pool, address to, uint256 minLpQty, uint256 deadline) payable returns()
func (_KyberZap *KyberZapSession) ZapInEth(tokenOut common.Address, pool common.Address, to common.Address, minLpQty *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _KyberZap.Contract.ZapInEth(&_KyberZap.TransactOpts, tokenOut, pool, to, minLpQty, deadline)
}

// ZapInEth is a paid mutator transaction binding the contract method 0x71679309.
//
// Solidity: function zapInEth(address tokenOut, address pool, address to, uint256 minLpQty, uint256 deadline) payable returns()
func (_KyberZap *KyberZapTransactorSession) ZapInEth(tokenOut common.Address, pool common.Address, to common.Address, minLpQty *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _KyberZap.Contract.ZapInEth(&_KyberZap.TransactOpts, tokenOut, pool, to, minLpQty, deadline)
}

// ZapOut is a paid mutator transaction binding the contract method 0xd4f9ec73.
//
// Solidity: function zapOut(address tokenIn, address tokenOut, uint256 liquidity, address pool, address to, uint256 minTokenOut, uint256 deadline) returns()
func (_KyberZap *KyberZapTransactor) ZapOut(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, liquidity *big.Int, pool common.Address, to common.Address, minTokenOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _KyberZap.contract.Transact(opts, "zapOut", tokenIn, tokenOut, liquidity, pool, to, minTokenOut, deadline)
}

// ZapOut is a paid mutator transaction binding the contract method 0xd4f9ec73.
//
// Solidity: function zapOut(address tokenIn, address tokenOut, uint256 liquidity, address pool, address to, uint256 minTokenOut, uint256 deadline) returns()
func (_KyberZap *KyberZapSession) ZapOut(tokenIn common.Address, tokenOut common.Address, liquidity *big.Int, pool common.Address, to common.Address, minTokenOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _KyberZap.Contract.ZapOut(&_KyberZap.TransactOpts, tokenIn, tokenOut, liquidity, pool, to, minTokenOut, deadline)
}

// ZapOut is a paid mutator transaction binding the contract method 0xd4f9ec73.
//
// Solidity: function zapOut(address tokenIn, address tokenOut, uint256 liquidity, address pool, address to, uint256 minTokenOut, uint256 deadline) returns()
func (_KyberZap *KyberZapTransactorSession) ZapOut(tokenIn common.Address, tokenOut common.Address, liquidity *big.Int, pool common.Address, to common.Address, minTokenOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _KyberZap.Contract.ZapOut(&_KyberZap.TransactOpts, tokenIn, tokenOut, liquidity, pool, to, minTokenOut, deadline)
}

// ZapOutPermit is a paid mutator transaction binding the contract method 0x2bb1a4ec.
//
// Solidity: function zapOutPermit(address tokenIn, address tokenOut, uint256 liquidity, address pool, address to, uint256 minTokenOut, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns()
func (_KyberZap *KyberZapTransactor) ZapOutPermit(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, liquidity *big.Int, pool common.Address, to common.Address, minTokenOut *big.Int, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberZap.contract.Transact(opts, "zapOutPermit", tokenIn, tokenOut, liquidity, pool, to, minTokenOut, deadline, approveMax, v, r, s)
}

// ZapOutPermit is a paid mutator transaction binding the contract method 0x2bb1a4ec.
//
// Solidity: function zapOutPermit(address tokenIn, address tokenOut, uint256 liquidity, address pool, address to, uint256 minTokenOut, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns()
func (_KyberZap *KyberZapSession) ZapOutPermit(tokenIn common.Address, tokenOut common.Address, liquidity *big.Int, pool common.Address, to common.Address, minTokenOut *big.Int, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberZap.Contract.ZapOutPermit(&_KyberZap.TransactOpts, tokenIn, tokenOut, liquidity, pool, to, minTokenOut, deadline, approveMax, v, r, s)
}

// ZapOutPermit is a paid mutator transaction binding the contract method 0x2bb1a4ec.
//
// Solidity: function zapOutPermit(address tokenIn, address tokenOut, uint256 liquidity, address pool, address to, uint256 minTokenOut, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns()
func (_KyberZap *KyberZapTransactorSession) ZapOutPermit(tokenIn common.Address, tokenOut common.Address, liquidity *big.Int, pool common.Address, to common.Address, minTokenOut *big.Int, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberZap.Contract.ZapOutPermit(&_KyberZap.TransactOpts, tokenIn, tokenOut, liquidity, pool, to, minTokenOut, deadline, approveMax, v, r, s)
}
