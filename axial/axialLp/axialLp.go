// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package axialLp

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

// AxialLpMetaData contains all meta data concerning the AxialLp contract.
var AxialLpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AxialLpABI is the input ABI used to generate the binding from.
// Deprecated: Use AxialLpMetaData.ABI instead.
var AxialLpABI = AxialLpMetaData.ABI

// AxialLp is an auto generated Go binding around an Ethereum contract.
type AxialLp struct {
	AxialLpCaller     // Read-only binding to the contract
	AxialLpTransactor // Write-only binding to the contract
	AxialLpFilterer   // Log filterer for contract events
}

// AxialLpCaller is an auto generated read-only Go binding around an Ethereum contract.
type AxialLpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxialLpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AxialLpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxialLpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AxialLpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxialLpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AxialLpSession struct {
	Contract     *AxialLp          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AxialLpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AxialLpCallerSession struct {
	Contract *AxialLpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AxialLpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AxialLpTransactorSession struct {
	Contract     *AxialLpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AxialLpRaw is an auto generated low-level Go binding around an Ethereum contract.
type AxialLpRaw struct {
	Contract *AxialLp // Generic contract binding to access the raw methods on
}

// AxialLpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AxialLpCallerRaw struct {
	Contract *AxialLpCaller // Generic read-only contract binding to access the raw methods on
}

// AxialLpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AxialLpTransactorRaw struct {
	Contract *AxialLpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAxialLp creates a new instance of AxialLp, bound to a specific deployed contract.
func NewAxialLp(address common.Address, backend bind.ContractBackend) (*AxialLp, error) {
	contract, err := bindAxialLp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AxialLp{AxialLpCaller: AxialLpCaller{contract: contract}, AxialLpTransactor: AxialLpTransactor{contract: contract}, AxialLpFilterer: AxialLpFilterer{contract: contract}}, nil
}

// NewAxialLpCaller creates a new read-only instance of AxialLp, bound to a specific deployed contract.
func NewAxialLpCaller(address common.Address, caller bind.ContractCaller) (*AxialLpCaller, error) {
	contract, err := bindAxialLp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AxialLpCaller{contract: contract}, nil
}

// NewAxialLpTransactor creates a new write-only instance of AxialLp, bound to a specific deployed contract.
func NewAxialLpTransactor(address common.Address, transactor bind.ContractTransactor) (*AxialLpTransactor, error) {
	contract, err := bindAxialLp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AxialLpTransactor{contract: contract}, nil
}

// NewAxialLpFilterer creates a new log filterer instance of AxialLp, bound to a specific deployed contract.
func NewAxialLpFilterer(address common.Address, filterer bind.ContractFilterer) (*AxialLpFilterer, error) {
	contract, err := bindAxialLp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AxialLpFilterer{contract: contract}, nil
}

// bindAxialLp binds a generic wrapper to an already deployed contract.
func bindAxialLp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AxialLpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AxialLp *AxialLpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AxialLp.Contract.AxialLpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AxialLp *AxialLpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AxialLp.Contract.AxialLpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AxialLp *AxialLpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AxialLp.Contract.AxialLpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AxialLp *AxialLpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AxialLp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AxialLp *AxialLpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AxialLp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AxialLp *AxialLpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AxialLp.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AxialLp *AxialLpCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AxialLp.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AxialLp *AxialLpSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AxialLp.Contract.Allowance(&_AxialLp.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AxialLp *AxialLpCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AxialLp.Contract.Allowance(&_AxialLp.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AxialLp *AxialLpCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AxialLp.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AxialLp *AxialLpSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AxialLp.Contract.BalanceOf(&_AxialLp.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AxialLp *AxialLpCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AxialLp.Contract.BalanceOf(&_AxialLp.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AxialLp *AxialLpCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AxialLp.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AxialLp *AxialLpSession) Decimals() (uint8, error) {
	return _AxialLp.Contract.Decimals(&_AxialLp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AxialLp *AxialLpCallerSession) Decimals() (uint8, error) {
	return _AxialLp.Contract.Decimals(&_AxialLp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AxialLp *AxialLpCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AxialLp.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AxialLp *AxialLpSession) Name() (string, error) {
	return _AxialLp.Contract.Name(&_AxialLp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AxialLp *AxialLpCallerSession) Name() (string, error) {
	return _AxialLp.Contract.Name(&_AxialLp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AxialLp *AxialLpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AxialLp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AxialLp *AxialLpSession) Owner() (common.Address, error) {
	return _AxialLp.Contract.Owner(&_AxialLp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AxialLp *AxialLpCallerSession) Owner() (common.Address, error) {
	return _AxialLp.Contract.Owner(&_AxialLp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AxialLp *AxialLpCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AxialLp.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AxialLp *AxialLpSession) Symbol() (string, error) {
	return _AxialLp.Contract.Symbol(&_AxialLp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AxialLp *AxialLpCallerSession) Symbol() (string, error) {
	return _AxialLp.Contract.Symbol(&_AxialLp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AxialLp *AxialLpCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AxialLp.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AxialLp *AxialLpSession) TotalSupply() (*big.Int, error) {
	return _AxialLp.Contract.TotalSupply(&_AxialLp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AxialLp *AxialLpCallerSession) TotalSupply() (*big.Int, error) {
	return _AxialLp.Contract.TotalSupply(&_AxialLp.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AxialLp *AxialLpTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AxialLp *AxialLpSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.Approve(&_AxialLp.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AxialLp *AxialLpTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.Approve(&_AxialLp.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_AxialLp *AxialLpTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_AxialLp *AxialLpSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.Burn(&_AxialLp.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_AxialLp *AxialLpTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.Burn(&_AxialLp.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_AxialLp *AxialLpTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_AxialLp *AxialLpSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.BurnFrom(&_AxialLp.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_AxialLp *AxialLpTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.BurnFrom(&_AxialLp.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AxialLp *AxialLpTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AxialLp.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AxialLp *AxialLpSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.DecreaseAllowance(&_AxialLp.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AxialLp *AxialLpTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.DecreaseAllowance(&_AxialLp.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AxialLp *AxialLpTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AxialLp.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AxialLp *AxialLpSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.IncreaseAllowance(&_AxialLp.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AxialLp *AxialLpTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.IncreaseAllowance(&_AxialLp.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns(bool)
func (_AxialLp *AxialLpTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _AxialLp.contract.Transact(opts, "initialize", name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns(bool)
func (_AxialLp *AxialLpSession) Initialize(name string, symbol string) (*types.Transaction, error) {
	return _AxialLp.Contract.Initialize(&_AxialLp.TransactOpts, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns(bool)
func (_AxialLp *AxialLpTransactorSession) Initialize(name string, symbol string) (*types.Transaction, error) {
	return _AxialLp.Contract.Initialize(&_AxialLp.TransactOpts, name, symbol)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_AxialLp *AxialLpTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.contract.Transact(opts, "mint", recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_AxialLp *AxialLpSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.Mint(&_AxialLp.TransactOpts, recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_AxialLp *AxialLpTransactorSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.Mint(&_AxialLp.TransactOpts, recipient, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AxialLp *AxialLpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AxialLp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AxialLp *AxialLpSession) RenounceOwnership() (*types.Transaction, error) {
	return _AxialLp.Contract.RenounceOwnership(&_AxialLp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AxialLp *AxialLpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AxialLp.Contract.RenounceOwnership(&_AxialLp.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AxialLp *AxialLpTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AxialLp *AxialLpSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.Transfer(&_AxialLp.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AxialLp *AxialLpTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.Transfer(&_AxialLp.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AxialLp *AxialLpTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AxialLp *AxialLpSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.TransferFrom(&_AxialLp.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AxialLp *AxialLpTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AxialLp.Contract.TransferFrom(&_AxialLp.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AxialLp *AxialLpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AxialLp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AxialLp *AxialLpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AxialLp.Contract.TransferOwnership(&_AxialLp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AxialLp *AxialLpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AxialLp.Contract.TransferOwnership(&_AxialLp.TransactOpts, newOwner)
}
