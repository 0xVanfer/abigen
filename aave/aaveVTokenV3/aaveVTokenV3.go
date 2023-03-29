// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveVTokenV3

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

// AaveVTokenV3MetaData contains all meta data concerning the AaveVTokenV3 contract.
var AaveVTokenV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEBT_TOKEN_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING_ASSET_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toUser\",\"type\":\"address\"}],\"name\":\"borrowAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIncentivesController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getScaledUserBalanceAndSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentivesController\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"debtTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"debtTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"debtTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"scaledBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaledTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveVTokenV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveVTokenV3MetaData.ABI instead.
var AaveVTokenV3ABI = AaveVTokenV3MetaData.ABI

// AaveVTokenV3 is an auto generated Go binding around an Ethereum contract.
type AaveVTokenV3 struct {
	AaveVTokenV3Caller     // Read-only binding to the contract
	AaveVTokenV3Transactor // Write-only binding to the contract
	AaveVTokenV3Filterer   // Log filterer for contract events
}

// AaveVTokenV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveVTokenV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveVTokenV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveVTokenV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveVTokenV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveVTokenV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveVTokenV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveVTokenV3Session struct {
	Contract     *AaveVTokenV3     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveVTokenV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveVTokenV3CallerSession struct {
	Contract *AaveVTokenV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AaveVTokenV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveVTokenV3TransactorSession struct {
	Contract     *AaveVTokenV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AaveVTokenV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveVTokenV3Raw struct {
	Contract *AaveVTokenV3 // Generic contract binding to access the raw methods on
}

// AaveVTokenV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveVTokenV3CallerRaw struct {
	Contract *AaveVTokenV3Caller // Generic read-only contract binding to access the raw methods on
}

// AaveVTokenV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveVTokenV3TransactorRaw struct {
	Contract *AaveVTokenV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveVTokenV3 creates a new instance of AaveVTokenV3, bound to a specific deployed contract.
func NewAaveVTokenV3(address common.Address, backend bind.ContractBackend) (*AaveVTokenV3, error) {
	contract, err := bindAaveVTokenV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveVTokenV3{AaveVTokenV3Caller: AaveVTokenV3Caller{contract: contract}, AaveVTokenV3Transactor: AaveVTokenV3Transactor{contract: contract}, AaveVTokenV3Filterer: AaveVTokenV3Filterer{contract: contract}}, nil
}

// NewAaveVTokenV3Caller creates a new read-only instance of AaveVTokenV3, bound to a specific deployed contract.
func NewAaveVTokenV3Caller(address common.Address, caller bind.ContractCaller) (*AaveVTokenV3Caller, error) {
	contract, err := bindAaveVTokenV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveVTokenV3Caller{contract: contract}, nil
}

// NewAaveVTokenV3Transactor creates a new write-only instance of AaveVTokenV3, bound to a specific deployed contract.
func NewAaveVTokenV3Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveVTokenV3Transactor, error) {
	contract, err := bindAaveVTokenV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveVTokenV3Transactor{contract: contract}, nil
}

// NewAaveVTokenV3Filterer creates a new log filterer instance of AaveVTokenV3, bound to a specific deployed contract.
func NewAaveVTokenV3Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveVTokenV3Filterer, error) {
	contract, err := bindAaveVTokenV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveVTokenV3Filterer{contract: contract}, nil
}

// bindAaveVTokenV3 binds a generic wrapper to an already deployed contract.
func bindAaveVTokenV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveVTokenV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveVTokenV3 *AaveVTokenV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveVTokenV3.Contract.AaveVTokenV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveVTokenV3 *AaveVTokenV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.AaveVTokenV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveVTokenV3 *AaveVTokenV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.AaveVTokenV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveVTokenV3 *AaveVTokenV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveVTokenV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveVTokenV3 *AaveVTokenV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveVTokenV3 *AaveVTokenV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.contract.Transact(opts, method, params...)
}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Caller) DEBTTOKENREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "DEBT_TOKEN_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Session) DEBTTOKENREVISION() (*big.Int, error) {
	return _AaveVTokenV3.Contract.DEBTTOKENREVISION(&_AaveVTokenV3.CallOpts)
}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) DEBTTOKENREVISION() (*big.Int, error) {
	return _AaveVTokenV3.Contract.DEBTTOKENREVISION(&_AaveVTokenV3.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveVTokenV3 *AaveVTokenV3Caller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveVTokenV3 *AaveVTokenV3Session) POOL() (common.Address, error) {
	return _AaveVTokenV3.Contract.POOL(&_AaveVTokenV3.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) POOL() (common.Address, error) {
	return _AaveVTokenV3.Contract.POOL(&_AaveVTokenV3.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveVTokenV3 *AaveVTokenV3Caller) UNDERLYINGASSETADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "UNDERLYING_ASSET_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveVTokenV3 *AaveVTokenV3Session) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveVTokenV3.Contract.UNDERLYINGASSETADDRESS(&_AaveVTokenV3.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveVTokenV3.Contract.UNDERLYINGASSETADDRESS(&_AaveVTokenV3.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveVTokenV3.Contract.Allowance(&_AaveVTokenV3.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveVTokenV3.Contract.Allowance(&_AaveVTokenV3.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Caller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Session) BalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVTokenV3.Contract.BalanceOf(&_AaveVTokenV3.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVTokenV3.Contract.BalanceOf(&_AaveVTokenV3.CallOpts, user)
}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Caller) BorrowAllowance(opts *bind.CallOpts, fromUser common.Address, toUser common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "borrowAllowance", fromUser, toUser)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Session) BorrowAllowance(fromUser common.Address, toUser common.Address) (*big.Int, error) {
	return _AaveVTokenV3.Contract.BorrowAllowance(&_AaveVTokenV3.CallOpts, fromUser, toUser)
}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) BorrowAllowance(fromUser common.Address, toUser common.Address) (*big.Int, error) {
	return _AaveVTokenV3.Contract.BorrowAllowance(&_AaveVTokenV3.CallOpts, fromUser, toUser)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveVTokenV3 *AaveVTokenV3Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveVTokenV3 *AaveVTokenV3Session) Decimals() (uint8, error) {
	return _AaveVTokenV3.Contract.Decimals(&_AaveVTokenV3.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) Decimals() (uint8, error) {
	return _AaveVTokenV3.Contract.Decimals(&_AaveVTokenV3.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveVTokenV3 *AaveVTokenV3Caller) GetIncentivesController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "getIncentivesController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveVTokenV3 *AaveVTokenV3Session) GetIncentivesController() (common.Address, error) {
	return _AaveVTokenV3.Contract.GetIncentivesController(&_AaveVTokenV3.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) GetIncentivesController() (common.Address, error) {
	return _AaveVTokenV3.Contract.GetIncentivesController(&_AaveVTokenV3.CallOpts)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveVTokenV3 *AaveVTokenV3Caller) GetScaledUserBalanceAndSupply(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "getScaledUserBalanceAndSupply", user)

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
func (_AaveVTokenV3 *AaveVTokenV3Session) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _AaveVTokenV3.Contract.GetScaledUserBalanceAndSupply(&_AaveVTokenV3.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _AaveVTokenV3.Contract.GetScaledUserBalanceAndSupply(&_AaveVTokenV3.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveVTokenV3 *AaveVTokenV3Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveVTokenV3 *AaveVTokenV3Session) Name() (string, error) {
	return _AaveVTokenV3.Contract.Name(&_AaveVTokenV3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) Name() (string, error) {
	return _AaveVTokenV3.Contract.Name(&_AaveVTokenV3.CallOpts)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Caller) ScaledBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "scaledBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Session) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVTokenV3.Contract.ScaledBalanceOf(&_AaveVTokenV3.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVTokenV3.Contract.ScaledBalanceOf(&_AaveVTokenV3.CallOpts, user)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Caller) ScaledTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "scaledTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Session) ScaledTotalSupply() (*big.Int, error) {
	return _AaveVTokenV3.Contract.ScaledTotalSupply(&_AaveVTokenV3.CallOpts)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) ScaledTotalSupply() (*big.Int, error) {
	return _AaveVTokenV3.Contract.ScaledTotalSupply(&_AaveVTokenV3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveVTokenV3 *AaveVTokenV3Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveVTokenV3 *AaveVTokenV3Session) Symbol() (string, error) {
	return _AaveVTokenV3.Contract.Symbol(&_AaveVTokenV3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) Symbol() (string, error) {
	return _AaveVTokenV3.Contract.Symbol(&_AaveVTokenV3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveVTokenV3.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3Session) TotalSupply() (*big.Int, error) {
	return _AaveVTokenV3.Contract.TotalSupply(&_AaveVTokenV3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveVTokenV3 *AaveVTokenV3CallerSession) TotalSupply() (*big.Int, error) {
	return _AaveVTokenV3.Contract.TotalSupply(&_AaveVTokenV3.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.Approve(&_AaveVTokenV3.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.Approve(&_AaveVTokenV3.TransactOpts, spender, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveVTokenV3 *AaveVTokenV3Transactor) ApproveDelegation(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.contract.Transact(opts, "approveDelegation", delegatee, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveVTokenV3 *AaveVTokenV3Session) ApproveDelegation(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.ApproveDelegation(&_AaveVTokenV3.TransactOpts, delegatee, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveVTokenV3 *AaveVTokenV3TransactorSession) ApproveDelegation(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.ApproveDelegation(&_AaveVTokenV3.TransactOpts, delegatee, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveVTokenV3 *AaveVTokenV3Transactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.contract.Transact(opts, "burn", user, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveVTokenV3 *AaveVTokenV3Session) Burn(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.Burn(&_AaveVTokenV3.TransactOpts, user, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveVTokenV3 *AaveVTokenV3TransactorSession) Burn(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.Burn(&_AaveVTokenV3.TransactOpts, user, amount, index)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.DecreaseAllowance(&_AaveVTokenV3.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.DecreaseAllowance(&_AaveVTokenV3.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.IncreaseAllowance(&_AaveVTokenV3.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.IncreaseAllowance(&_AaveVTokenV3.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc222ec8a.
//
// Solidity: function initialize(address pool, address underlyingAsset, address incentivesController, uint8 debtTokenDecimals, string debtTokenName, string debtTokenSymbol, bytes params) returns()
func (_AaveVTokenV3 *AaveVTokenV3Transactor) Initialize(opts *bind.TransactOpts, pool common.Address, underlyingAsset common.Address, incentivesController common.Address, debtTokenDecimals uint8, debtTokenName string, debtTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _AaveVTokenV3.contract.Transact(opts, "initialize", pool, underlyingAsset, incentivesController, debtTokenDecimals, debtTokenName, debtTokenSymbol, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xc222ec8a.
//
// Solidity: function initialize(address pool, address underlyingAsset, address incentivesController, uint8 debtTokenDecimals, string debtTokenName, string debtTokenSymbol, bytes params) returns()
func (_AaveVTokenV3 *AaveVTokenV3Session) Initialize(pool common.Address, underlyingAsset common.Address, incentivesController common.Address, debtTokenDecimals uint8, debtTokenName string, debtTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.Initialize(&_AaveVTokenV3.TransactOpts, pool, underlyingAsset, incentivesController, debtTokenDecimals, debtTokenName, debtTokenSymbol, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xc222ec8a.
//
// Solidity: function initialize(address pool, address underlyingAsset, address incentivesController, uint8 debtTokenDecimals, string debtTokenName, string debtTokenSymbol, bytes params) returns()
func (_AaveVTokenV3 *AaveVTokenV3TransactorSession) Initialize(pool common.Address, underlyingAsset common.Address, incentivesController common.Address, debtTokenDecimals uint8, debtTokenName string, debtTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.Initialize(&_AaveVTokenV3.TransactOpts, pool, underlyingAsset, incentivesController, debtTokenDecimals, debtTokenName, debtTokenSymbol, params)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Transactor) Mint(opts *bind.TransactOpts, user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.contract.Transact(opts, "mint", user, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Session) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.Mint(&_AaveVTokenV3.TransactOpts, user, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3TransactorSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.Mint(&_AaveVTokenV3.TransactOpts, user, onBehalfOf, amount, index)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.Transfer(&_AaveVTokenV3.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.Transfer(&_AaveVTokenV3.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.TransferFrom(&_AaveVTokenV3.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveVTokenV3 *AaveVTokenV3TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVTokenV3.Contract.TransferFrom(&_AaveVTokenV3.TransactOpts, sender, recipient, amount)
}
