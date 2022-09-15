// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platypusLp

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

// PlatypusLpMetaData contains all meta data concerning the PlatypusLp contract.
var PlatypusLpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addCash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addLiability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregateAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlyingToken_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"aggregateAccount_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liability\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeCash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeLiability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregateAccount_\",\"type\":\"address\"}],\"name\":\"setAggregateAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply_\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferUnderlyingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PlatypusLpABI is the input ABI used to generate the binding from.
// Deprecated: Use PlatypusLpMetaData.ABI instead.
var PlatypusLpABI = PlatypusLpMetaData.ABI

// PlatypusLp is an auto generated Go binding around an Ethereum contract.
type PlatypusLp struct {
	PlatypusLpCaller     // Read-only binding to the contract
	PlatypusLpTransactor // Write-only binding to the contract
	PlatypusLpFilterer   // Log filterer for contract events
}

// PlatypusLpCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlatypusLpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusLpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlatypusLpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusLpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlatypusLpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusLpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlatypusLpSession struct {
	Contract     *PlatypusLp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlatypusLpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlatypusLpCallerSession struct {
	Contract *PlatypusLpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PlatypusLpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlatypusLpTransactorSession struct {
	Contract     *PlatypusLpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PlatypusLpRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlatypusLpRaw struct {
	Contract *PlatypusLp // Generic contract binding to access the raw methods on
}

// PlatypusLpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlatypusLpCallerRaw struct {
	Contract *PlatypusLpCaller // Generic read-only contract binding to access the raw methods on
}

// PlatypusLpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlatypusLpTransactorRaw struct {
	Contract *PlatypusLpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlatypusLp creates a new instance of PlatypusLp, bound to a specific deployed contract.
func NewPlatypusLp(address common.Address, backend bind.ContractBackend) (*PlatypusLp, error) {
	contract, err := bindPlatypusLp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlatypusLp{PlatypusLpCaller: PlatypusLpCaller{contract: contract}, PlatypusLpTransactor: PlatypusLpTransactor{contract: contract}, PlatypusLpFilterer: PlatypusLpFilterer{contract: contract}}, nil
}

// NewPlatypusLpCaller creates a new read-only instance of PlatypusLp, bound to a specific deployed contract.
func NewPlatypusLpCaller(address common.Address, caller bind.ContractCaller) (*PlatypusLpCaller, error) {
	contract, err := bindPlatypusLp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlatypusLpCaller{contract: contract}, nil
}

// NewPlatypusLpTransactor creates a new write-only instance of PlatypusLp, bound to a specific deployed contract.
func NewPlatypusLpTransactor(address common.Address, transactor bind.ContractTransactor) (*PlatypusLpTransactor, error) {
	contract, err := bindPlatypusLp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlatypusLpTransactor{contract: contract}, nil
}

// NewPlatypusLpFilterer creates a new log filterer instance of PlatypusLp, bound to a specific deployed contract.
func NewPlatypusLpFilterer(address common.Address, filterer bind.ContractFilterer) (*PlatypusLpFilterer, error) {
	contract, err := bindPlatypusLp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlatypusLpFilterer{contract: contract}, nil
}

// bindPlatypusLp binds a generic wrapper to an already deployed contract.
func bindPlatypusLp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlatypusLpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatypusLp *PlatypusLpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatypusLp.Contract.PlatypusLpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatypusLp *PlatypusLpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusLp.Contract.PlatypusLpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatypusLp *PlatypusLpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatypusLp.Contract.PlatypusLpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatypusLp *PlatypusLpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatypusLp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatypusLp *PlatypusLpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusLp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatypusLp *PlatypusLpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatypusLp.Contract.contract.Transact(opts, method, params...)
}

// AggregateAccount is a free data retrieval call binding the contract method 0x7e1317fa.
//
// Solidity: function aggregateAccount() view returns(address)
func (_PlatypusLp *PlatypusLpCaller) AggregateAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "aggregateAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggregateAccount is a free data retrieval call binding the contract method 0x7e1317fa.
//
// Solidity: function aggregateAccount() view returns(address)
func (_PlatypusLp *PlatypusLpSession) AggregateAccount() (common.Address, error) {
	return _PlatypusLp.Contract.AggregateAccount(&_PlatypusLp.CallOpts)
}

// AggregateAccount is a free data retrieval call binding the contract method 0x7e1317fa.
//
// Solidity: function aggregateAccount() view returns(address)
func (_PlatypusLp *PlatypusLpCallerSession) AggregateAccount() (common.Address, error) {
	return _PlatypusLp.Contract.AggregateAccount(&_PlatypusLp.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PlatypusLp *PlatypusLpCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PlatypusLp *PlatypusLpSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PlatypusLp.Contract.Allowance(&_PlatypusLp.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PlatypusLp *PlatypusLpCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PlatypusLp.Contract.Allowance(&_PlatypusLp.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PlatypusLp *PlatypusLpCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PlatypusLp *PlatypusLpSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PlatypusLp.Contract.BalanceOf(&_PlatypusLp.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PlatypusLp *PlatypusLpCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PlatypusLp.Contract.BalanceOf(&_PlatypusLp.CallOpts, account)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_PlatypusLp *PlatypusLpCaller) Cash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_PlatypusLp *PlatypusLpSession) Cash() (*big.Int, error) {
	return _PlatypusLp.Contract.Cash(&_PlatypusLp.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_PlatypusLp *PlatypusLpCallerSession) Cash() (*big.Int, error) {
	return _PlatypusLp.Contract.Cash(&_PlatypusLp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PlatypusLp *PlatypusLpCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PlatypusLp *PlatypusLpSession) Decimals() (uint8, error) {
	return _PlatypusLp.Contract.Decimals(&_PlatypusLp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PlatypusLp *PlatypusLpCallerSession) Decimals() (uint8, error) {
	return _PlatypusLp.Contract.Decimals(&_PlatypusLp.CallOpts)
}

// Liability is a free data retrieval call binding the contract method 0x705727b5.
//
// Solidity: function liability() view returns(uint256)
func (_PlatypusLp *PlatypusLpCaller) Liability(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "liability")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liability is a free data retrieval call binding the contract method 0x705727b5.
//
// Solidity: function liability() view returns(uint256)
func (_PlatypusLp *PlatypusLpSession) Liability() (*big.Int, error) {
	return _PlatypusLp.Contract.Liability(&_PlatypusLp.CallOpts)
}

// Liability is a free data retrieval call binding the contract method 0x705727b5.
//
// Solidity: function liability() view returns(uint256)
func (_PlatypusLp *PlatypusLpCallerSession) Liability() (*big.Int, error) {
	return _PlatypusLp.Contract.Liability(&_PlatypusLp.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_PlatypusLp *PlatypusLpCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_PlatypusLp *PlatypusLpSession) MaxSupply() (*big.Int, error) {
	return _PlatypusLp.Contract.MaxSupply(&_PlatypusLp.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_PlatypusLp *PlatypusLpCallerSession) MaxSupply() (*big.Int, error) {
	return _PlatypusLp.Contract.MaxSupply(&_PlatypusLp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PlatypusLp *PlatypusLpCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PlatypusLp *PlatypusLpSession) Name() (string, error) {
	return _PlatypusLp.Contract.Name(&_PlatypusLp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PlatypusLp *PlatypusLpCallerSession) Name() (string, error) {
	return _PlatypusLp.Contract.Name(&_PlatypusLp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusLp *PlatypusLpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusLp *PlatypusLpSession) Owner() (common.Address, error) {
	return _PlatypusLp.Contract.Owner(&_PlatypusLp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusLp *PlatypusLpCallerSession) Owner() (common.Address, error) {
	return _PlatypusLp.Contract.Owner(&_PlatypusLp.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PlatypusLp *PlatypusLpCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PlatypusLp *PlatypusLpSession) Pool() (common.Address, error) {
	return _PlatypusLp.Contract.Pool(&_PlatypusLp.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PlatypusLp *PlatypusLpCallerSession) Pool() (common.Address, error) {
	return _PlatypusLp.Contract.Pool(&_PlatypusLp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PlatypusLp *PlatypusLpCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PlatypusLp *PlatypusLpSession) Symbol() (string, error) {
	return _PlatypusLp.Contract.Symbol(&_PlatypusLp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PlatypusLp *PlatypusLpCallerSession) Symbol() (string, error) {
	return _PlatypusLp.Contract.Symbol(&_PlatypusLp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PlatypusLp *PlatypusLpCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PlatypusLp *PlatypusLpSession) TotalSupply() (*big.Int, error) {
	return _PlatypusLp.Contract.TotalSupply(&_PlatypusLp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PlatypusLp *PlatypusLpCallerSession) TotalSupply() (*big.Int, error) {
	return _PlatypusLp.Contract.TotalSupply(&_PlatypusLp.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PlatypusLp *PlatypusLpCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PlatypusLp *PlatypusLpSession) UnderlyingToken() (common.Address, error) {
	return _PlatypusLp.Contract.UnderlyingToken(&_PlatypusLp.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PlatypusLp *PlatypusLpCallerSession) UnderlyingToken() (common.Address, error) {
	return _PlatypusLp.Contract.UnderlyingToken(&_PlatypusLp.CallOpts)
}

// UnderlyingTokenBalance is a free data retrieval call binding the contract method 0x99c91a64.
//
// Solidity: function underlyingTokenBalance() view returns(uint256)
func (_PlatypusLp *PlatypusLpCaller) UnderlyingTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusLp.contract.Call(opts, &out, "underlyingTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingTokenBalance is a free data retrieval call binding the contract method 0x99c91a64.
//
// Solidity: function underlyingTokenBalance() view returns(uint256)
func (_PlatypusLp *PlatypusLpSession) UnderlyingTokenBalance() (*big.Int, error) {
	return _PlatypusLp.Contract.UnderlyingTokenBalance(&_PlatypusLp.CallOpts)
}

// UnderlyingTokenBalance is a free data retrieval call binding the contract method 0x99c91a64.
//
// Solidity: function underlyingTokenBalance() view returns(uint256)
func (_PlatypusLp *PlatypusLpCallerSession) UnderlyingTokenBalance() (*big.Int, error) {
	return _PlatypusLp.Contract.UnderlyingTokenBalance(&_PlatypusLp.CallOpts)
}

// AddCash is a paid mutator transaction binding the contract method 0x16c9e7a0.
//
// Solidity: function addCash(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactor) AddCash(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "addCash", amount)
}

// AddCash is a paid mutator transaction binding the contract method 0x16c9e7a0.
//
// Solidity: function addCash(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpSession) AddCash(amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.AddCash(&_PlatypusLp.TransactOpts, amount)
}

// AddCash is a paid mutator transaction binding the contract method 0x16c9e7a0.
//
// Solidity: function addCash(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) AddCash(amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.AddCash(&_PlatypusLp.TransactOpts, amount)
}

// AddLiability is a paid mutator transaction binding the contract method 0xa0f0f604.
//
// Solidity: function addLiability(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactor) AddLiability(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "addLiability", amount)
}

// AddLiability is a paid mutator transaction binding the contract method 0xa0f0f604.
//
// Solidity: function addLiability(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpSession) AddLiability(amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.AddLiability(&_PlatypusLp.TransactOpts, amount)
}

// AddLiability is a paid mutator transaction binding the contract method 0xa0f0f604.
//
// Solidity: function addLiability(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) AddLiability(amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.AddLiability(&_PlatypusLp.TransactOpts, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PlatypusLp *PlatypusLpTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PlatypusLp *PlatypusLpSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.Approve(&_PlatypusLp.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PlatypusLp *PlatypusLpTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.Approve(&_PlatypusLp.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactor) Burn(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "burn", to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_PlatypusLp *PlatypusLpSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.Burn(&_PlatypusLp.TransactOpts, to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.Burn(&_PlatypusLp.TransactOpts, to, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PlatypusLp *PlatypusLpTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PlatypusLp *PlatypusLpSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.DecreaseAllowance(&_PlatypusLp.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PlatypusLp *PlatypusLpTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.DecreaseAllowance(&_PlatypusLp.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PlatypusLp *PlatypusLpTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PlatypusLp *PlatypusLpSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.IncreaseAllowance(&_PlatypusLp.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PlatypusLp *PlatypusLpTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.IncreaseAllowance(&_PlatypusLp.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x613d25bb.
//
// Solidity: function initialize(address underlyingToken_, string name_, string symbol_, address aggregateAccount_) returns()
func (_PlatypusLp *PlatypusLpTransactor) Initialize(opts *bind.TransactOpts, underlyingToken_ common.Address, name_ string, symbol_ string, aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "initialize", underlyingToken_, name_, symbol_, aggregateAccount_)
}

// Initialize is a paid mutator transaction binding the contract method 0x613d25bb.
//
// Solidity: function initialize(address underlyingToken_, string name_, string symbol_, address aggregateAccount_) returns()
func (_PlatypusLp *PlatypusLpSession) Initialize(underlyingToken_ common.Address, name_ string, symbol_ string, aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _PlatypusLp.Contract.Initialize(&_PlatypusLp.TransactOpts, underlyingToken_, name_, symbol_, aggregateAccount_)
}

// Initialize is a paid mutator transaction binding the contract method 0x613d25bb.
//
// Solidity: function initialize(address underlyingToken_, string name_, string symbol_, address aggregateAccount_) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) Initialize(underlyingToken_ common.Address, name_ string, symbol_ string, aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _PlatypusLp.Contract.Initialize(&_PlatypusLp.TransactOpts, underlyingToken_, name_, symbol_, aggregateAccount_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PlatypusLp *PlatypusLpSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.Mint(&_PlatypusLp.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.Mint(&_PlatypusLp.TransactOpts, to, amount)
}

// RemoveCash is a paid mutator transaction binding the contract method 0x9f9ef988.
//
// Solidity: function removeCash(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactor) RemoveCash(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "removeCash", amount)
}

// RemoveCash is a paid mutator transaction binding the contract method 0x9f9ef988.
//
// Solidity: function removeCash(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpSession) RemoveCash(amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.RemoveCash(&_PlatypusLp.TransactOpts, amount)
}

// RemoveCash is a paid mutator transaction binding the contract method 0x9f9ef988.
//
// Solidity: function removeCash(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) RemoveCash(amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.RemoveCash(&_PlatypusLp.TransactOpts, amount)
}

// RemoveLiability is a paid mutator transaction binding the contract method 0xd8b87853.
//
// Solidity: function removeLiability(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactor) RemoveLiability(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "removeLiability", amount)
}

// RemoveLiability is a paid mutator transaction binding the contract method 0xd8b87853.
//
// Solidity: function removeLiability(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpSession) RemoveLiability(amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.RemoveLiability(&_PlatypusLp.TransactOpts, amount)
}

// RemoveLiability is a paid mutator transaction binding the contract method 0xd8b87853.
//
// Solidity: function removeLiability(uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) RemoveLiability(amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.RemoveLiability(&_PlatypusLp.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusLp *PlatypusLpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusLp *PlatypusLpSession) RenounceOwnership() (*types.Transaction, error) {
	return _PlatypusLp.Contract.RenounceOwnership(&_PlatypusLp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusLp *PlatypusLpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PlatypusLp.Contract.RenounceOwnership(&_PlatypusLp.TransactOpts)
}

// SetAggregateAccount is a paid mutator transaction binding the contract method 0x95d6f7b9.
//
// Solidity: function setAggregateAccount(address aggregateAccount_) returns()
func (_PlatypusLp *PlatypusLpTransactor) SetAggregateAccount(opts *bind.TransactOpts, aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "setAggregateAccount", aggregateAccount_)
}

// SetAggregateAccount is a paid mutator transaction binding the contract method 0x95d6f7b9.
//
// Solidity: function setAggregateAccount(address aggregateAccount_) returns()
func (_PlatypusLp *PlatypusLpSession) SetAggregateAccount(aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _PlatypusLp.Contract.SetAggregateAccount(&_PlatypusLp.TransactOpts, aggregateAccount_)
}

// SetAggregateAccount is a paid mutator transaction binding the contract method 0x95d6f7b9.
//
// Solidity: function setAggregateAccount(address aggregateAccount_) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) SetAggregateAccount(aggregateAccount_ common.Address) (*types.Transaction, error) {
	return _PlatypusLp.Contract.SetAggregateAccount(&_PlatypusLp.TransactOpts, aggregateAccount_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_PlatypusLp *PlatypusLpTransactor) SetMaxSupply(opts *bind.TransactOpts, maxSupply_ *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "setMaxSupply", maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_PlatypusLp *PlatypusLpSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.SetMaxSupply(&_PlatypusLp.TransactOpts, maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.SetMaxSupply(&_PlatypusLp.TransactOpts, maxSupply_)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address pool_) returns()
func (_PlatypusLp *PlatypusLpTransactor) SetPool(opts *bind.TransactOpts, pool_ common.Address) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "setPool", pool_)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address pool_) returns()
func (_PlatypusLp *PlatypusLpSession) SetPool(pool_ common.Address) (*types.Transaction, error) {
	return _PlatypusLp.Contract.SetPool(&_PlatypusLp.TransactOpts, pool_)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address pool_) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) SetPool(pool_ common.Address) (*types.Transaction, error) {
	return _PlatypusLp.Contract.SetPool(&_PlatypusLp.TransactOpts, pool_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PlatypusLp *PlatypusLpTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PlatypusLp *PlatypusLpSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.Transfer(&_PlatypusLp.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PlatypusLp *PlatypusLpTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.Transfer(&_PlatypusLp.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PlatypusLp *PlatypusLpTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PlatypusLp *PlatypusLpSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.TransferFrom(&_PlatypusLp.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PlatypusLp *PlatypusLpTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.TransferFrom(&_PlatypusLp.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PlatypusLp *PlatypusLpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PlatypusLp *PlatypusLpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusLp.Contract.TransferOwnership(&_PlatypusLp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusLp.Contract.TransferOwnership(&_PlatypusLp.TransactOpts, newOwner)
}

// TransferUnderlyingToken is a paid mutator transaction binding the contract method 0x9e79eaa5.
//
// Solidity: function transferUnderlyingToken(address to, uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactor) TransferUnderlyingToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.contract.Transact(opts, "transferUnderlyingToken", to, amount)
}

// TransferUnderlyingToken is a paid mutator transaction binding the contract method 0x9e79eaa5.
//
// Solidity: function transferUnderlyingToken(address to, uint256 amount) returns()
func (_PlatypusLp *PlatypusLpSession) TransferUnderlyingToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.TransferUnderlyingToken(&_PlatypusLp.TransactOpts, to, amount)
}

// TransferUnderlyingToken is a paid mutator transaction binding the contract method 0x9e79eaa5.
//
// Solidity: function transferUnderlyingToken(address to, uint256 amount) returns()
func (_PlatypusLp *PlatypusLpTransactorSession) TransferUnderlyingToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlatypusLp.Contract.TransferUnderlyingToken(&_PlatypusLp.TransactOpts, to, amount)
}
