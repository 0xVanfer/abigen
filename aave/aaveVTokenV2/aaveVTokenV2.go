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
	ABI: "[{\"inputs\":[],\"name\":\"DEBT_TOKEN_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING_ASSET_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toUser\",\"type\":\"address\"}],\"name\":\"borrowAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIncentivesController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getScaledUserBalanceAndSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentivesController\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"debtTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"debtTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"debtTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"scaledBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaledTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Caller) DEBTTOKENREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "DEBT_TOKEN_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Session) DEBTTOKENREVISION() (*big.Int, error) {
	return _AaveVTokenV2.Contract.DEBTTOKENREVISION(&_AaveVTokenV2.CallOpts)
}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) DEBTTOKENREVISION() (*big.Int, error) {
	return _AaveVTokenV2.Contract.DEBTTOKENREVISION(&_AaveVTokenV2.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveVTokenV2 *AaveVTokenV2Caller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveVTokenV2 *AaveVTokenV2Session) POOL() (common.Address, error) {
	return _AaveVTokenV2.Contract.POOL(&_AaveVTokenV2.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) POOL() (common.Address, error) {
	return _AaveVTokenV2.Contract.POOL(&_AaveVTokenV2.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveVTokenV2 *AaveVTokenV2Caller) UNDERLYINGASSETADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "UNDERLYING_ASSET_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveVTokenV2 *AaveVTokenV2Session) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveVTokenV2.Contract.UNDERLYINGASSETADDRESS(&_AaveVTokenV2.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveVTokenV2.Contract.UNDERLYINGASSETADDRESS(&_AaveVTokenV2.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveVTokenV2.Contract.Allowance(&_AaveVTokenV2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveVTokenV2.Contract.Allowance(&_AaveVTokenV2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Caller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Session) BalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVTokenV2.Contract.BalanceOf(&_AaveVTokenV2.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVTokenV2.Contract.BalanceOf(&_AaveVTokenV2.CallOpts, user)
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

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveVTokenV2 *AaveVTokenV2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveVTokenV2 *AaveVTokenV2Session) Decimals() (uint8, error) {
	return _AaveVTokenV2.Contract.Decimals(&_AaveVTokenV2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) Decimals() (uint8, error) {
	return _AaveVTokenV2.Contract.Decimals(&_AaveVTokenV2.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveVTokenV2 *AaveVTokenV2Caller) GetIncentivesController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "getIncentivesController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveVTokenV2 *AaveVTokenV2Session) GetIncentivesController() (common.Address, error) {
	return _AaveVTokenV2.Contract.GetIncentivesController(&_AaveVTokenV2.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) GetIncentivesController() (common.Address, error) {
	return _AaveVTokenV2.Contract.GetIncentivesController(&_AaveVTokenV2.CallOpts)
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

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveVTokenV2 *AaveVTokenV2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveVTokenV2 *AaveVTokenV2Session) Name() (string, error) {
	return _AaveVTokenV2.Contract.Name(&_AaveVTokenV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) Name() (string, error) {
	return _AaveVTokenV2.Contract.Name(&_AaveVTokenV2.CallOpts)
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

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveVTokenV2 *AaveVTokenV2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveVTokenV2 *AaveVTokenV2Session) Symbol() (string, error) {
	return _AaveVTokenV2.Contract.Symbol(&_AaveVTokenV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) Symbol() (string, error) {
	return _AaveVTokenV2.Contract.Symbol(&_AaveVTokenV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2Session) TotalSupply() (*big.Int, error) {
	return _AaveVTokenV2.Contract.TotalSupply(&_AaveVTokenV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveVTokenV2 *AaveVTokenV2CallerSession) TotalSupply() (*big.Int, error) {
	return _AaveVTokenV2.Contract.TotalSupply(&_AaveVTokenV2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.Approve(&_AaveVTokenV2.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.Approve(&_AaveVTokenV2.TransactOpts, spender, amount)
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

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveVTokenV2 *AaveVTokenV2Transactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.contract.Transact(opts, "burn", user, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveVTokenV2 *AaveVTokenV2Session) Burn(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.Burn(&_AaveVTokenV2.TransactOpts, user, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveVTokenV2 *AaveVTokenV2TransactorSession) Burn(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.Burn(&_AaveVTokenV2.TransactOpts, user, amount, index)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.DecreaseAllowance(&_AaveVTokenV2.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.DecreaseAllowance(&_AaveVTokenV2.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.IncreaseAllowance(&_AaveVTokenV2.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.IncreaseAllowance(&_AaveVTokenV2.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc222ec8a.
//
// Solidity: function initialize(address pool, address underlyingAsset, address incentivesController, uint8 debtTokenDecimals, string debtTokenName, string debtTokenSymbol, bytes params) returns()
func (_AaveVTokenV2 *AaveVTokenV2Transactor) Initialize(opts *bind.TransactOpts, pool common.Address, underlyingAsset common.Address, incentivesController common.Address, debtTokenDecimals uint8, debtTokenName string, debtTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _AaveVTokenV2.contract.Transact(opts, "initialize", pool, underlyingAsset, incentivesController, debtTokenDecimals, debtTokenName, debtTokenSymbol, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xc222ec8a.
//
// Solidity: function initialize(address pool, address underlyingAsset, address incentivesController, uint8 debtTokenDecimals, string debtTokenName, string debtTokenSymbol, bytes params) returns()
func (_AaveVTokenV2 *AaveVTokenV2Session) Initialize(pool common.Address, underlyingAsset common.Address, incentivesController common.Address, debtTokenDecimals uint8, debtTokenName string, debtTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.Initialize(&_AaveVTokenV2.TransactOpts, pool, underlyingAsset, incentivesController, debtTokenDecimals, debtTokenName, debtTokenSymbol, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xc222ec8a.
//
// Solidity: function initialize(address pool, address underlyingAsset, address incentivesController, uint8 debtTokenDecimals, string debtTokenName, string debtTokenSymbol, bytes params) returns()
func (_AaveVTokenV2 *AaveVTokenV2TransactorSession) Initialize(pool common.Address, underlyingAsset common.Address, incentivesController common.Address, debtTokenDecimals uint8, debtTokenName string, debtTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.Initialize(&_AaveVTokenV2.TransactOpts, pool, underlyingAsset, incentivesController, debtTokenDecimals, debtTokenName, debtTokenSymbol, params)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Transactor) Mint(opts *bind.TransactOpts, user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.contract.Transact(opts, "mint", user, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Session) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.Mint(&_AaveVTokenV2.TransactOpts, user, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2TransactorSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.Mint(&_AaveVTokenV2.TransactOpts, user, onBehalfOf, amount, index)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.Transfer(&_AaveVTokenV2.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.Transfer(&_AaveVTokenV2.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.TransferFrom(&_AaveVTokenV2.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV2 *AaveVTokenV2TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV2.Contract.TransferFrom(&_AaveVTokenV2.TransactOpts, sender, recipient, amount)
}
