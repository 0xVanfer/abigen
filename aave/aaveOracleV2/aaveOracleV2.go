// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveOracleV2

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

// AaveOracleV2MetaData contains all meta data concerning the AaveOracleV2 contract.
var AaveOracleV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"getAssetsPrices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFallbackOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getSourceOfAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sources\",\"type\":\"address[]\"}],\"name\":\"setAssetSources\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"}],\"name\":\"setFallbackOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveOracleV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveOracleV2MetaData.ABI instead.
var AaveOracleV2ABI = AaveOracleV2MetaData.ABI

// AaveOracleV2 is an auto generated Go binding around an Ethereum contract.
type AaveOracleV2 struct {
	AaveOracleV2Caller     // Read-only binding to the contract
	AaveOracleV2Transactor // Write-only binding to the contract
	AaveOracleV2Filterer   // Log filterer for contract events
}

// AaveOracleV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveOracleV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveOracleV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveOracleV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveOracleV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveOracleV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveOracleV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveOracleV2Session struct {
	Contract     *AaveOracleV2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveOracleV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveOracleV2CallerSession struct {
	Contract *AaveOracleV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AaveOracleV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveOracleV2TransactorSession struct {
	Contract     *AaveOracleV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AaveOracleV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveOracleV2Raw struct {
	Contract *AaveOracleV2 // Generic contract binding to access the raw methods on
}

// AaveOracleV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveOracleV2CallerRaw struct {
	Contract *AaveOracleV2Caller // Generic read-only contract binding to access the raw methods on
}

// AaveOracleV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveOracleV2TransactorRaw struct {
	Contract *AaveOracleV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveOracleV2 creates a new instance of AaveOracleV2, bound to a specific deployed contract.
func NewAaveOracleV2(address common.Address, backend bind.ContractBackend) (*AaveOracleV2, error) {
	contract, err := bindAaveOracleV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveOracleV2{AaveOracleV2Caller: AaveOracleV2Caller{contract: contract}, AaveOracleV2Transactor: AaveOracleV2Transactor{contract: contract}, AaveOracleV2Filterer: AaveOracleV2Filterer{contract: contract}}, nil
}

// NewAaveOracleV2Caller creates a new read-only instance of AaveOracleV2, bound to a specific deployed contract.
func NewAaveOracleV2Caller(address common.Address, caller bind.ContractCaller) (*AaveOracleV2Caller, error) {
	contract, err := bindAaveOracleV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveOracleV2Caller{contract: contract}, nil
}

// NewAaveOracleV2Transactor creates a new write-only instance of AaveOracleV2, bound to a specific deployed contract.
func NewAaveOracleV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveOracleV2Transactor, error) {
	contract, err := bindAaveOracleV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveOracleV2Transactor{contract: contract}, nil
}

// NewAaveOracleV2Filterer creates a new log filterer instance of AaveOracleV2, bound to a specific deployed contract.
func NewAaveOracleV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveOracleV2Filterer, error) {
	contract, err := bindAaveOracleV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveOracleV2Filterer{contract: contract}, nil
}

// bindAaveOracleV2 binds a generic wrapper to an already deployed contract.
func bindAaveOracleV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveOracleV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveOracleV2 *AaveOracleV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveOracleV2.Contract.AaveOracleV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveOracleV2 *AaveOracleV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveOracleV2.Contract.AaveOracleV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveOracleV2 *AaveOracleV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveOracleV2.Contract.AaveOracleV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveOracleV2 *AaveOracleV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveOracleV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveOracleV2 *AaveOracleV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveOracleV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveOracleV2 *AaveOracleV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveOracleV2.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_AaveOracleV2 *AaveOracleV2Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveOracleV2.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_AaveOracleV2 *AaveOracleV2Session) WETH() (common.Address, error) {
	return _AaveOracleV2.Contract.WETH(&_AaveOracleV2.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_AaveOracleV2 *AaveOracleV2CallerSession) WETH() (common.Address, error) {
	return _AaveOracleV2.Contract.WETH(&_AaveOracleV2.CallOpts)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_AaveOracleV2 *AaveOracleV2Caller) GetAssetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveOracleV2.contract.Call(opts, &out, "getAssetPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_AaveOracleV2 *AaveOracleV2Session) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _AaveOracleV2.Contract.GetAssetPrice(&_AaveOracleV2.CallOpts, asset)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_AaveOracleV2 *AaveOracleV2CallerSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _AaveOracleV2.Contract.GetAssetPrice(&_AaveOracleV2.CallOpts, asset)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_AaveOracleV2 *AaveOracleV2Caller) GetAssetsPrices(opts *bind.CallOpts, assets []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AaveOracleV2.contract.Call(opts, &out, "getAssetsPrices", assets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_AaveOracleV2 *AaveOracleV2Session) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _AaveOracleV2.Contract.GetAssetsPrices(&_AaveOracleV2.CallOpts, assets)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_AaveOracleV2 *AaveOracleV2CallerSession) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _AaveOracleV2.Contract.GetAssetsPrices(&_AaveOracleV2.CallOpts, assets)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_AaveOracleV2 *AaveOracleV2Caller) GetFallbackOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveOracleV2.contract.Call(opts, &out, "getFallbackOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_AaveOracleV2 *AaveOracleV2Session) GetFallbackOracle() (common.Address, error) {
	return _AaveOracleV2.Contract.GetFallbackOracle(&_AaveOracleV2.CallOpts)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_AaveOracleV2 *AaveOracleV2CallerSession) GetFallbackOracle() (common.Address, error) {
	return _AaveOracleV2.Contract.GetFallbackOracle(&_AaveOracleV2.CallOpts)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_AaveOracleV2 *AaveOracleV2Caller) GetSourceOfAsset(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _AaveOracleV2.contract.Call(opts, &out, "getSourceOfAsset", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_AaveOracleV2 *AaveOracleV2Session) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _AaveOracleV2.Contract.GetSourceOfAsset(&_AaveOracleV2.CallOpts, asset)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_AaveOracleV2 *AaveOracleV2CallerSession) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _AaveOracleV2.Contract.GetSourceOfAsset(&_AaveOracleV2.CallOpts, asset)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveOracleV2 *AaveOracleV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveOracleV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveOracleV2 *AaveOracleV2Session) Owner() (common.Address, error) {
	return _AaveOracleV2.Contract.Owner(&_AaveOracleV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveOracleV2 *AaveOracleV2CallerSession) Owner() (common.Address, error) {
	return _AaveOracleV2.Contract.Owner(&_AaveOracleV2.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveOracleV2 *AaveOracleV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveOracleV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveOracleV2 *AaveOracleV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _AaveOracleV2.Contract.RenounceOwnership(&_AaveOracleV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveOracleV2 *AaveOracleV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveOracleV2.Contract.RenounceOwnership(&_AaveOracleV2.TransactOpts)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_AaveOracleV2 *AaveOracleV2Transactor) SetAssetSources(opts *bind.TransactOpts, assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _AaveOracleV2.contract.Transact(opts, "setAssetSources", assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_AaveOracleV2 *AaveOracleV2Session) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _AaveOracleV2.Contract.SetAssetSources(&_AaveOracleV2.TransactOpts, assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_AaveOracleV2 *AaveOracleV2TransactorSession) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _AaveOracleV2.Contract.SetAssetSources(&_AaveOracleV2.TransactOpts, assets, sources)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_AaveOracleV2 *AaveOracleV2Transactor) SetFallbackOracle(opts *bind.TransactOpts, fallbackOracle common.Address) (*types.Transaction, error) {
	return _AaveOracleV2.contract.Transact(opts, "setFallbackOracle", fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_AaveOracleV2 *AaveOracleV2Session) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _AaveOracleV2.Contract.SetFallbackOracle(&_AaveOracleV2.TransactOpts, fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_AaveOracleV2 *AaveOracleV2TransactorSession) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _AaveOracleV2.Contract.SetFallbackOracle(&_AaveOracleV2.TransactOpts, fallbackOracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveOracleV2 *AaveOracleV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AaveOracleV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveOracleV2 *AaveOracleV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveOracleV2.Contract.TransferOwnership(&_AaveOracleV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveOracleV2 *AaveOracleV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveOracleV2.Contract.TransferOwnership(&_AaveOracleV2.TransactOpts, newOwner)
}
