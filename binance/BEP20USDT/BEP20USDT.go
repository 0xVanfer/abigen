// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BEP20USDT

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

// BEP20USDTMetaData contains all meta data concerning the BEP20USDT contract.
var BEP20USDTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BEP20USDTABI is the input ABI used to generate the binding from.
// Deprecated: Use BEP20USDTMetaData.ABI instead.
var BEP20USDTABI = BEP20USDTMetaData.ABI

// BEP20USDT is an auto generated Go binding around an Ethereum contract.
type BEP20USDT struct {
	BEP20USDTCaller     // Read-only binding to the contract
	BEP20USDTTransactor // Write-only binding to the contract
	BEP20USDTFilterer   // Log filterer for contract events
}

// BEP20USDTCaller is an auto generated read-only Go binding around an Ethereum contract.
type BEP20USDTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20USDTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BEP20USDTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20USDTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BEP20USDTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20USDTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BEP20USDTSession struct {
	Contract     *BEP20USDT        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BEP20USDTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BEP20USDTCallerSession struct {
	Contract *BEP20USDTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BEP20USDTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BEP20USDTTransactorSession struct {
	Contract     *BEP20USDTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BEP20USDTRaw is an auto generated low-level Go binding around an Ethereum contract.
type BEP20USDTRaw struct {
	Contract *BEP20USDT // Generic contract binding to access the raw methods on
}

// BEP20USDTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BEP20USDTCallerRaw struct {
	Contract *BEP20USDTCaller // Generic read-only contract binding to access the raw methods on
}

// BEP20USDTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BEP20USDTTransactorRaw struct {
	Contract *BEP20USDTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBEP20USDT creates a new instance of BEP20USDT, bound to a specific deployed contract.
func NewBEP20USDT(address common.Address, backend bind.ContractBackend) (*BEP20USDT, error) {
	contract, err := bindBEP20USDT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BEP20USDT{BEP20USDTCaller: BEP20USDTCaller{contract: contract}, BEP20USDTTransactor: BEP20USDTTransactor{contract: contract}, BEP20USDTFilterer: BEP20USDTFilterer{contract: contract}}, nil
}

// NewBEP20USDTCaller creates a new read-only instance of BEP20USDT, bound to a specific deployed contract.
func NewBEP20USDTCaller(address common.Address, caller bind.ContractCaller) (*BEP20USDTCaller, error) {
	contract, err := bindBEP20USDT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20USDTCaller{contract: contract}, nil
}

// NewBEP20USDTTransactor creates a new write-only instance of BEP20USDT, bound to a specific deployed contract.
func NewBEP20USDTTransactor(address common.Address, transactor bind.ContractTransactor) (*BEP20USDTTransactor, error) {
	contract, err := bindBEP20USDT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20USDTTransactor{contract: contract}, nil
}

// NewBEP20USDTFilterer creates a new log filterer instance of BEP20USDT, bound to a specific deployed contract.
func NewBEP20USDTFilterer(address common.Address, filterer bind.ContractFilterer) (*BEP20USDTFilterer, error) {
	contract, err := bindBEP20USDT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BEP20USDTFilterer{contract: contract}, nil
}

// bindBEP20USDT binds a generic wrapper to an already deployed contract.
func bindBEP20USDT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BEP20USDTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20USDT *BEP20USDTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20USDT.Contract.BEP20USDTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20USDT *BEP20USDTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20USDT.Contract.BEP20USDTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20USDT *BEP20USDTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20USDT.Contract.BEP20USDTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20USDT *BEP20USDTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20USDT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20USDT *BEP20USDTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20USDT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20USDT *BEP20USDTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20USDT.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BEP20USDT *BEP20USDTCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BEP20USDT.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BEP20USDT *BEP20USDTSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BEP20USDT.Contract.Allowance(&_BEP20USDT.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BEP20USDT *BEP20USDTCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BEP20USDT.Contract.Allowance(&_BEP20USDT.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BEP20USDT *BEP20USDTCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BEP20USDT.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BEP20USDT *BEP20USDTSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BEP20USDT.Contract.BalanceOf(&_BEP20USDT.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BEP20USDT *BEP20USDTCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BEP20USDT.Contract.BalanceOf(&_BEP20USDT.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BEP20USDT *BEP20USDTCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BEP20USDT.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BEP20USDT *BEP20USDTSession) Decimals() (uint8, error) {
	return _BEP20USDT.Contract.Decimals(&_BEP20USDT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BEP20USDT *BEP20USDTCallerSession) Decimals() (uint8, error) {
	return _BEP20USDT.Contract.Decimals(&_BEP20USDT.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BEP20USDT *BEP20USDTCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20USDT.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BEP20USDT *BEP20USDTSession) GetOwner() (common.Address, error) {
	return _BEP20USDT.Contract.GetOwner(&_BEP20USDT.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BEP20USDT *BEP20USDTCallerSession) GetOwner() (common.Address, error) {
	return _BEP20USDT.Contract.GetOwner(&_BEP20USDT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BEP20USDT *BEP20USDTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BEP20USDT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BEP20USDT *BEP20USDTSession) Name() (string, error) {
	return _BEP20USDT.Contract.Name(&_BEP20USDT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BEP20USDT *BEP20USDTCallerSession) Name() (string, error) {
	return _BEP20USDT.Contract.Name(&_BEP20USDT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20USDT *BEP20USDTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20USDT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20USDT *BEP20USDTSession) Owner() (common.Address, error) {
	return _BEP20USDT.Contract.Owner(&_BEP20USDT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20USDT *BEP20USDTCallerSession) Owner() (common.Address, error) {
	return _BEP20USDT.Contract.Owner(&_BEP20USDT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BEP20USDT *BEP20USDTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BEP20USDT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BEP20USDT *BEP20USDTSession) Symbol() (string, error) {
	return _BEP20USDT.Contract.Symbol(&_BEP20USDT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BEP20USDT *BEP20USDTCallerSession) Symbol() (string, error) {
	return _BEP20USDT.Contract.Symbol(&_BEP20USDT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BEP20USDT *BEP20USDTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20USDT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BEP20USDT *BEP20USDTSession) TotalSupply() (*big.Int, error) {
	return _BEP20USDT.Contract.TotalSupply(&_BEP20USDT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BEP20USDT *BEP20USDTCallerSession) TotalSupply() (*big.Int, error) {
	return _BEP20USDT.Contract.TotalSupply(&_BEP20USDT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.Approve(&_BEP20USDT.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.Approve(&_BEP20USDT.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.Burn(&_BEP20USDT.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.Burn(&_BEP20USDT.TransactOpts, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BEP20USDT *BEP20USDTTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BEP20USDT *BEP20USDTSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.DecreaseAllowance(&_BEP20USDT.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BEP20USDT *BEP20USDTTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.DecreaseAllowance(&_BEP20USDT.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BEP20USDT *BEP20USDTTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BEP20USDT *BEP20USDTSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.IncreaseAllowance(&_BEP20USDT.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BEP20USDT *BEP20USDTTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.IncreaseAllowance(&_BEP20USDT.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.Mint(&_BEP20USDT.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.Mint(&_BEP20USDT.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20USDT *BEP20USDTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20USDT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20USDT *BEP20USDTSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20USDT.Contract.RenounceOwnership(&_BEP20USDT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20USDT *BEP20USDTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20USDT.Contract.RenounceOwnership(&_BEP20USDT.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.Transfer(&_BEP20USDT.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.Transfer(&_BEP20USDT.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.TransferFrom(&_BEP20USDT.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BEP20USDT *BEP20USDTTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20USDT.Contract.TransferFrom(&_BEP20USDT.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20USDT *BEP20USDTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BEP20USDT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20USDT *BEP20USDTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20USDT.Contract.TransferOwnership(&_BEP20USDT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20USDT *BEP20USDTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20USDT.Contract.TransferOwnership(&_BEP20USDT.TransactOpts, newOwner)
}
