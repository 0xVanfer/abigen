// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveLp

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

// CurveLpMetaData contains all meta data concerning the CurveLp contract.
var CurveLpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"arg1\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtracted_value\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_added_value\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"frac\",\"type\":\"uint256\"}],\"name\":\"mint_relative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"set_minter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"set_name\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CurveLpABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveLpMetaData.ABI instead.
var CurveLpABI = CurveLpMetaData.ABI

// CurveLp is an auto generated Go binding around an Ethereum contract.
type CurveLp struct {
	CurveLpCaller     // Read-only binding to the contract
	CurveLpTransactor // Write-only binding to the contract
	CurveLpFilterer   // Log filterer for contract events
}

// CurveLpCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveLpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveLpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveLpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveLpSession struct {
	Contract     *CurveLp          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveLpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveLpCallerSession struct {
	Contract *CurveLpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CurveLpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveLpTransactorSession struct {
	Contract     *CurveLpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CurveLpRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveLpRaw struct {
	Contract *CurveLp // Generic contract binding to access the raw methods on
}

// CurveLpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveLpCallerRaw struct {
	Contract *CurveLpCaller // Generic read-only contract binding to access the raw methods on
}

// CurveLpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveLpTransactorRaw struct {
	Contract *CurveLpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveLp creates a new instance of CurveLp, bound to a specific deployed contract.
func NewCurveLp(address common.Address, backend bind.ContractBackend) (*CurveLp, error) {
	contract, err := bindCurveLp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveLp{CurveLpCaller: CurveLpCaller{contract: contract}, CurveLpTransactor: CurveLpTransactor{contract: contract}, CurveLpFilterer: CurveLpFilterer{contract: contract}}, nil
}

// NewCurveLpCaller creates a new read-only instance of CurveLp, bound to a specific deployed contract.
func NewCurveLpCaller(address common.Address, caller bind.ContractCaller) (*CurveLpCaller, error) {
	contract, err := bindCurveLp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveLpCaller{contract: contract}, nil
}

// NewCurveLpTransactor creates a new write-only instance of CurveLp, bound to a specific deployed contract.
func NewCurveLpTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveLpTransactor, error) {
	contract, err := bindCurveLp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveLpTransactor{contract: contract}, nil
}

// NewCurveLpFilterer creates a new log filterer instance of CurveLp, bound to a specific deployed contract.
func NewCurveLpFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveLpFilterer, error) {
	contract, err := bindCurveLp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveLpFilterer{contract: contract}, nil
}

// bindCurveLp binds a generic wrapper to an already deployed contract.
func bindCurveLp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveLpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveLp *CurveLpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveLp.Contract.CurveLpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveLp *CurveLpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveLp.Contract.CurveLpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveLp *CurveLpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveLp.Contract.CurveLpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveLp *CurveLpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveLp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveLp *CurveLpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveLp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveLp *CurveLpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveLp.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveLp *CurveLpCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLp.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveLp *CurveLpSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveLp.Contract.Allowance(&_CurveLp.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveLp *CurveLpCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveLp.Contract.Allowance(&_CurveLp.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveLp *CurveLpCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLp.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveLp *CurveLpSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveLp.Contract.BalanceOf(&_CurveLp.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveLp *CurveLpCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveLp.Contract.BalanceOf(&_CurveLp.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveLp *CurveLpCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLp.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveLp *CurveLpSession) Decimals() (*big.Int, error) {
	return _CurveLp.Contract.Decimals(&_CurveLp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveLp *CurveLpCallerSession) Decimals() (*big.Int, error) {
	return _CurveLp.Contract.Decimals(&_CurveLp.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CurveLp *CurveLpCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveLp.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CurveLp *CurveLpSession) Minter() (common.Address, error) {
	return _CurveLp.Contract.Minter(&_CurveLp.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CurveLp *CurveLpCallerSession) Minter() (common.Address, error) {
	return _CurveLp.Contract.Minter(&_CurveLp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveLp *CurveLpCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveLp.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveLp *CurveLpSession) Name() (string, error) {
	return _CurveLp.Contract.Name(&_CurveLp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveLp *CurveLpCallerSession) Name() (string, error) {
	return _CurveLp.Contract.Name(&_CurveLp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveLp *CurveLpCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveLp.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveLp *CurveLpSession) Symbol() (string, error) {
	return _CurveLp.Contract.Symbol(&_CurveLp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveLp *CurveLpCallerSession) Symbol() (string, error) {
	return _CurveLp.Contract.Symbol(&_CurveLp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveLp *CurveLpCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLp.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveLp *CurveLpSession) TotalSupply() (*big.Int, error) {
	return _CurveLp.Contract.TotalSupply(&_CurveLp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveLp *CurveLpCallerSession) TotalSupply() (*big.Int, error) {
	return _CurveLp.Contract.TotalSupply(&_CurveLp.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveLp *CurveLpTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveLp *CurveLpSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.Approve(&_CurveLp.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveLp *CurveLpTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.Approve(&_CurveLp.TransactOpts, _spender, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpTransactor) BurnFrom(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.contract.Transact(opts, "burnFrom", _to, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpSession) BurnFrom(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.BurnFrom(&_CurveLp.TransactOpts, _to, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpTransactorSession) BurnFrom(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.BurnFrom(&_CurveLp.TransactOpts, _to, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveLp *CurveLpTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveLp.contract.Transact(opts, "decreaseAllowance", _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveLp *CurveLpSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.DecreaseAllowance(&_CurveLp.TransactOpts, _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveLp *CurveLpTransactorSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.DecreaseAllowance(&_CurveLp.TransactOpts, _spender, _subtracted_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveLp *CurveLpTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveLp.contract.Transact(opts, "increaseAllowance", _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveLp *CurveLpSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.IncreaseAllowance(&_CurveLp.TransactOpts, _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveLp *CurveLpTransactorSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.IncreaseAllowance(&_CurveLp.TransactOpts, _spender, _added_value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.Mint(&_CurveLp.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpTransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.Mint(&_CurveLp.TransactOpts, _to, _value)
}

// MintRelative is a paid mutator transaction binding the contract method 0x6962f845.
//
// Solidity: function mint_relative(address _to, uint256 frac) returns(uint256)
func (_CurveLp *CurveLpTransactor) MintRelative(opts *bind.TransactOpts, _to common.Address, frac *big.Int) (*types.Transaction, error) {
	return _CurveLp.contract.Transact(opts, "mint_relative", _to, frac)
}

// MintRelative is a paid mutator transaction binding the contract method 0x6962f845.
//
// Solidity: function mint_relative(address _to, uint256 frac) returns(uint256)
func (_CurveLp *CurveLpSession) MintRelative(_to common.Address, frac *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.MintRelative(&_CurveLp.TransactOpts, _to, frac)
}

// MintRelative is a paid mutator transaction binding the contract method 0x6962f845.
//
// Solidity: function mint_relative(address _to, uint256 frac) returns(uint256)
func (_CurveLp *CurveLpTransactorSession) MintRelative(_to common.Address, frac *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.MintRelative(&_CurveLp.TransactOpts, _to, frac)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_CurveLp *CurveLpTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _CurveLp.contract.Transact(opts, "set_minter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_CurveLp *CurveLpSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CurveLp.Contract.SetMinter(&_CurveLp.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_CurveLp *CurveLpTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CurveLp.Contract.SetMinter(&_CurveLp.TransactOpts, _minter)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_CurveLp *CurveLpTransactor) SetName(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _CurveLp.contract.Transact(opts, "set_name", _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_CurveLp *CurveLpSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _CurveLp.Contract.SetName(&_CurveLp.TransactOpts, _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_CurveLp *CurveLpTransactorSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _CurveLp.Contract.SetName(&_CurveLp.TransactOpts, _name, _symbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.Transfer(&_CurveLp.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.Transfer(&_CurveLp.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.TransferFrom(&_CurveLp.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveLp *CurveLpTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLp.Contract.TransferFrom(&_CurveLp.TransactOpts, _from, _to, _value)
}
