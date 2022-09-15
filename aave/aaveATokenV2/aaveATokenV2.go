// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveATokenV2

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

// AaveATokenV2MetaData contains all meta data concerning the AaveATokenV2 contract.
var AaveATokenV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING_ASSET_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveATokenV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveATokenV2MetaData.ABI instead.
var AaveATokenV2ABI = AaveATokenV2MetaData.ABI

// AaveATokenV2 is an auto generated Go binding around an Ethereum contract.
type AaveATokenV2 struct {
	AaveATokenV2Caller     // Read-only binding to the contract
	AaveATokenV2Transactor // Write-only binding to the contract
	AaveATokenV2Filterer   // Log filterer for contract events
}

// AaveATokenV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveATokenV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveATokenV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveATokenV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveATokenV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveATokenV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveATokenV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveATokenV2Session struct {
	Contract     *AaveATokenV2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveATokenV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveATokenV2CallerSession struct {
	Contract *AaveATokenV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AaveATokenV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveATokenV2TransactorSession struct {
	Contract     *AaveATokenV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AaveATokenV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveATokenV2Raw struct {
	Contract *AaveATokenV2 // Generic contract binding to access the raw methods on
}

// AaveATokenV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveATokenV2CallerRaw struct {
	Contract *AaveATokenV2Caller // Generic read-only contract binding to access the raw methods on
}

// AaveATokenV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveATokenV2TransactorRaw struct {
	Contract *AaveATokenV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveATokenV2 creates a new instance of AaveATokenV2, bound to a specific deployed contract.
func NewAaveATokenV2(address common.Address, backend bind.ContractBackend) (*AaveATokenV2, error) {
	contract, err := bindAaveATokenV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveATokenV2{AaveATokenV2Caller: AaveATokenV2Caller{contract: contract}, AaveATokenV2Transactor: AaveATokenV2Transactor{contract: contract}, AaveATokenV2Filterer: AaveATokenV2Filterer{contract: contract}}, nil
}

// NewAaveATokenV2Caller creates a new read-only instance of AaveATokenV2, bound to a specific deployed contract.
func NewAaveATokenV2Caller(address common.Address, caller bind.ContractCaller) (*AaveATokenV2Caller, error) {
	contract, err := bindAaveATokenV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveATokenV2Caller{contract: contract}, nil
}

// NewAaveATokenV2Transactor creates a new write-only instance of AaveATokenV2, bound to a specific deployed contract.
func NewAaveATokenV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveATokenV2Transactor, error) {
	contract, err := bindAaveATokenV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveATokenV2Transactor{contract: contract}, nil
}

// NewAaveATokenV2Filterer creates a new log filterer instance of AaveATokenV2, bound to a specific deployed contract.
func NewAaveATokenV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveATokenV2Filterer, error) {
	contract, err := bindAaveATokenV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveATokenV2Filterer{contract: contract}, nil
}

// bindAaveATokenV2 binds a generic wrapper to an already deployed contract.
func bindAaveATokenV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveATokenV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveATokenV2 *AaveATokenV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveATokenV2.Contract.AaveATokenV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveATokenV2 *AaveATokenV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.AaveATokenV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveATokenV2 *AaveATokenV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.AaveATokenV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveATokenV2 *AaveATokenV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveATokenV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveATokenV2 *AaveATokenV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveATokenV2 *AaveATokenV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.contract.Transact(opts, method, params...)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Caller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Session) POOL() (common.Address, error) {
	return _AaveATokenV2.Contract.POOL(&_AaveATokenV2.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveATokenV2 *AaveATokenV2CallerSession) POOL() (common.Address, error) {
	return _AaveATokenV2.Contract.POOL(&_AaveATokenV2.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Caller) UNDERLYINGASSETADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "UNDERLYING_ASSET_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Session) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveATokenV2.Contract.UNDERLYINGASSETADDRESS(&_AaveATokenV2.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveATokenV2 *AaveATokenV2CallerSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveATokenV2.Contract.UNDERLYINGASSETADDRESS(&_AaveATokenV2.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.Allowance(&_AaveATokenV2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.Allowance(&_AaveATokenV2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Caller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Session) BalanceOf(_user common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.BalanceOf(&_AaveATokenV2.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2CallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.BalanceOf(&_AaveATokenV2.CallOpts, _user)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Session) Decimals() (*big.Int, error) {
	return _AaveATokenV2.Contract.Decimals(&_AaveATokenV2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2CallerSession) Decimals() (*big.Int, error) {
	return _AaveATokenV2.Contract.Decimals(&_AaveATokenV2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Approve(&_AaveATokenV2.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Approve(&_AaveATokenV2.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_AaveATokenV2 *AaveATokenV2Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_AaveATokenV2 *AaveATokenV2Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Transfer(&_AaveATokenV2.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_AaveATokenV2 *AaveATokenV2TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Transfer(&_AaveATokenV2.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.TransferFrom(&_AaveATokenV2.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.TransferFrom(&_AaveATokenV2.TransactOpts, sender, recipient, amount)
}
