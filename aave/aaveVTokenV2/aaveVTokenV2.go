// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveVTokenV2

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

// AaveVTokenV2MetaData contains all meta data concerning the AaveVTokenV2 contract.
var AaveVTokenV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toUser\",\"type\":\"address\"}],\"name\":\"borrowAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getScaledUserBalanceAndSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"scaledBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaledTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AaveVTokenV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveVTokenV2MetaData.ABI instead.
var AaveVTokenV2ABI = AaveVTokenV2MetaData.ABI

// AaveVTokenV2 is an auto generated Go binding around an Ethereum contract.
type AaveVTokenV2 struct {
	AaveVTokenV2Caller     // Read-only binding to the contract
	AaveVTokenV2Transactor // Write-only binding to the contract
	AaveVTokenV2Filterer   // Log filterer for contract events
}

// AaveVTokenV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveVTokenV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveVTokenV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveVTokenV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveVTokenV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveVTokenV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveVTokenV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveVTokenV2Session struct {
	Contract     *AaveVTokenV2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveVTokenV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveVTokenV2CallerSession struct {
	Contract *AaveVTokenV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AaveVTokenV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveVTokenV2TransactorSession struct {
	Contract     *AaveVTokenV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AaveVTokenV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveVTokenV2Raw struct {
	Contract *AaveVTokenV2 // Generic contract binding to access the raw methods on
}

// AaveVTokenV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveVTokenV2CallerRaw struct {
	Contract *AaveVTokenV2Caller // Generic read-only contract binding to access the raw methods on
}

// AaveVTokenV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveVTokenV2TransactorRaw struct {
	Contract *AaveVTokenV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveVTokenV2 creates a new instance of AaveVTokenV2, bound to a specific deployed contract.
func NewAaveVTokenV2(address common.Address, backend bind.ContractBackend) (*AaveVTokenV2, error) {
	contract, err := bindAaveVTokenV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveVTokenV2{AaveVTokenV2Caller: AaveVTokenV2Caller{contract: contract}, AaveVTokenV2Transactor: AaveVTokenV2Transactor{contract: contract}, AaveVTokenV2Filterer: AaveVTokenV2Filterer{contract: contract}}, nil
}

// NewAaveVTokenV2Caller creates a new read-only instance of AaveVTokenV2, bound to a specific deployed contract.
func NewAaveVTokenV2Caller(address common.Address, caller bind.ContractCaller) (*AaveVTokenV2Caller, error) {
	contract, err := bindAaveVTokenV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveVTokenV2Caller{contract: contract}, nil
}

// NewAaveVTokenV2Transactor creates a new write-only instance of AaveVTokenV2, bound to a specific deployed contract.
func NewAaveVTokenV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveVTokenV2Transactor, error) {
	contract, err := bindAaveVTokenV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveVTokenV2Transactor{contract: contract}, nil
}

// NewAaveVTokenV2Filterer creates a new log filterer instance of AaveVTokenV2, bound to a specific deployed contract.
func NewAaveVTokenV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveVTokenV2Filterer, error) {
	contract, err := bindAaveVTokenV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveVTokenV2Filterer{contract: contract}, nil
}

// bindAaveVTokenV2 binds a generic wrapper to an already deployed contract.
func bindAaveVTokenV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveVTokenV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveVTokenV2 *AaveVTokenV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveVTokenV2.Contract.AaveVTokenV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveVTokenV2 *AaveVTokenV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.AaveVTokenV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveVTokenV2 *AaveVTokenV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.AaveVTokenV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveVTokenV2 *AaveVTokenV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveVTokenV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveVTokenV2 *AaveVTokenV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveVTokenV2 *AaveVTokenV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.contract.Transact(opts, method, params...)
}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Caller) BorrowAllowance(opts *bind.CallOpts, fromUser common.Address, toUser common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "borrowAllowance", fromUser, toUser)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Session) BorrowAllowance(fromUser common.Address, toUser common.Address) (*big.Int, error) {
	return _AaveVTokenV2.Contract.BorrowAllowance(&_AaveVTokenV2.CallOpts, fromUser, toUser)
}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) BorrowAllowance(fromUser common.Address, toUser common.Address) (*big.Int, error) {
	return _AaveVTokenV2.Contract.BorrowAllowance(&_AaveVTokenV2.CallOpts, fromUser, toUser)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveVTokenV2 *AaveVTokenV2Caller) GetScaledUserBalanceAndSupply(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "getScaledUserBalanceAndSupply", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveVTokenV2 *AaveVTokenV2Session) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _AaveVTokenV2.Contract.GetScaledUserBalanceAndSupply(&_AaveVTokenV2.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _AaveVTokenV2.Contract.GetScaledUserBalanceAndSupply(&_AaveVTokenV2.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Caller) ScaledBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "scaledBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Session) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVTokenV2.Contract.ScaledBalanceOf(&_AaveVTokenV2.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVTokenV2.Contract.ScaledBalanceOf(&_AaveVTokenV2.CallOpts, user)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Caller) ScaledTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "scaledTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Session) ScaledTotalSupply() (*big.Int, error) {
	return _AaveVTokenV2.Contract.ScaledTotalSupply(&_AaveVTokenV2.CallOpts)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) ScaledTotalSupply() (*big.Int, error) {
	return _AaveVTokenV2.Contract.ScaledTotalSupply(&_AaveVTokenV2.CallOpts)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveVTokenV2 *AaveVTokenV2Transactor) ApproveDelegation(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.contract.Transact(opts, "approveDelegation", delegatee, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveVTokenV2 *AaveVTokenV2Session) ApproveDelegation(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.ApproveDelegation(&_AaveVTokenV2.TransactOpts, delegatee, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveVTokenV2 *AaveVTokenV2TransactorSession) ApproveDelegation(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.ApproveDelegation(&_AaveVTokenV2.TransactOpts, delegatee, amount)
}
