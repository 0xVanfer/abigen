// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveGaugeFactory

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

// CurveGaugeFactoryMetaData contains all meta data concerning the CurveGaugeFactory contract.
var CurveGaugeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"accept_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"call_proxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_future_owner\",\"type\":\"address\"}],\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"deploy_gauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"deploy_gauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"gauge_data\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"get_gauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_gauge_count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"get_gauge_from_lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"is_mirrored\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"is_valid_gauge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"last_request\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[32]\",\"name\":\"_gauges\",\"type\":\"address[32]\"}],\"name\":\"mint_many\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"arg1\",\"type\":\"address\"}],\"name\":\"minted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_call_proxy\",\"type\":\"address\"}],\"name\":\"set_call_proxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"set_implementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_mirrored\",\"type\":\"bool\"}],\"name\":\"set_mirrored\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voting_escrow\",\"type\":\"address\"}],\"name\":\"set_voting_escrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voting_escrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CurveGaugeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveGaugeFactoryMetaData.ABI instead.
var CurveGaugeFactoryABI = CurveGaugeFactoryMetaData.ABI

// CurveGaugeFactory is an auto generated Go binding around an Ethereum contract.
type CurveGaugeFactory struct {
	CurveGaugeFactoryCaller     // Read-only binding to the contract
	CurveGaugeFactoryTransactor // Write-only binding to the contract
	CurveGaugeFactoryFilterer   // Log filterer for contract events
}

// CurveGaugeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveGaugeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveGaugeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveGaugeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveGaugeFactorySession struct {
	Contract     *CurveGaugeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CurveGaugeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveGaugeFactoryCallerSession struct {
	Contract *CurveGaugeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CurveGaugeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveGaugeFactoryTransactorSession struct {
	Contract     *CurveGaugeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CurveGaugeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveGaugeFactoryRaw struct {
	Contract *CurveGaugeFactory // Generic contract binding to access the raw methods on
}

// CurveGaugeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveGaugeFactoryCallerRaw struct {
	Contract *CurveGaugeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CurveGaugeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveGaugeFactoryTransactorRaw struct {
	Contract *CurveGaugeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveGaugeFactory creates a new instance of CurveGaugeFactory, bound to a specific deployed contract.
func NewCurveGaugeFactory(address common.Address, backend bind.ContractBackend) (*CurveGaugeFactory, error) {
	contract, err := bindCurveGaugeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeFactory{CurveGaugeFactoryCaller: CurveGaugeFactoryCaller{contract: contract}, CurveGaugeFactoryTransactor: CurveGaugeFactoryTransactor{contract: contract}, CurveGaugeFactoryFilterer: CurveGaugeFactoryFilterer{contract: contract}}, nil
}

// NewCurveGaugeFactoryCaller creates a new read-only instance of CurveGaugeFactory, bound to a specific deployed contract.
func NewCurveGaugeFactoryCaller(address common.Address, caller bind.ContractCaller) (*CurveGaugeFactoryCaller, error) {
	contract, err := bindCurveGaugeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeFactoryCaller{contract: contract}, nil
}

// NewCurveGaugeFactoryTransactor creates a new write-only instance of CurveGaugeFactory, bound to a specific deployed contract.
func NewCurveGaugeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveGaugeFactoryTransactor, error) {
	contract, err := bindCurveGaugeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeFactoryTransactor{contract: contract}, nil
}

// NewCurveGaugeFactoryFilterer creates a new log filterer instance of CurveGaugeFactory, bound to a specific deployed contract.
func NewCurveGaugeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveGaugeFactoryFilterer, error) {
	contract, err := bindCurveGaugeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeFactoryFilterer{contract: contract}, nil
}

// bindCurveGaugeFactory binds a generic wrapper to an already deployed contract.
func bindCurveGaugeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveGaugeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGaugeFactory *CurveGaugeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGaugeFactory.Contract.CurveGaugeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGaugeFactory *CurveGaugeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.CurveGaugeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGaugeFactory *CurveGaugeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.CurveGaugeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGaugeFactory *CurveGaugeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGaugeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.contract.Transact(opts, method, params...)
}

// CallProxy is a free data retrieval call binding the contract method 0xf81c6c3e.
//
// Solidity: function call_proxy() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) CallProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "call_proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CallProxy is a free data retrieval call binding the contract method 0xf81c6c3e.
//
// Solidity: function call_proxy() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactorySession) CallProxy() (common.Address, error) {
	return _CurveGaugeFactory.Contract.CallProxy(&_CurveGaugeFactory.CallOpts)
}

// CallProxy is a free data retrieval call binding the contract method 0xf81c6c3e.
//
// Solidity: function call_proxy() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) CallProxy() (common.Address, error) {
	return _CurveGaugeFactory.Contract.CallProxy(&_CurveGaugeFactory.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactorySession) FutureOwner() (common.Address, error) {
	return _CurveGaugeFactory.Contract.FutureOwner(&_CurveGaugeFactory.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) FutureOwner() (common.Address, error) {
	return _CurveGaugeFactory.Contract.FutureOwner(&_CurveGaugeFactory.CallOpts)
}

// GaugeData is a free data retrieval call binding the contract method 0xf0ce32f8.
//
// Solidity: function gauge_data(address arg0) view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) GaugeData(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "gauge_data", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GaugeData is a free data retrieval call binding the contract method 0xf0ce32f8.
//
// Solidity: function gauge_data(address arg0) view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactorySession) GaugeData(arg0 common.Address) (*big.Int, error) {
	return _CurveGaugeFactory.Contract.GaugeData(&_CurveGaugeFactory.CallOpts, arg0)
}

// GaugeData is a free data retrieval call binding the contract method 0xf0ce32f8.
//
// Solidity: function gauge_data(address arg0) view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) GaugeData(arg0 common.Address) (*big.Int, error) {
	return _CurveGaugeFactory.Contract.GaugeData(&_CurveGaugeFactory.CallOpts, arg0)
}

// GetGauge is a free data retrieval call binding the contract method 0x28521848.
//
// Solidity: function get_gauge(uint256 arg0) view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) GetGauge(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "get_gauge", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge is a free data retrieval call binding the contract method 0x28521848.
//
// Solidity: function get_gauge(uint256 arg0) view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactorySession) GetGauge(arg0 *big.Int) (common.Address, error) {
	return _CurveGaugeFactory.Contract.GetGauge(&_CurveGaugeFactory.CallOpts, arg0)
}

// GetGauge is a free data retrieval call binding the contract method 0x28521848.
//
// Solidity: function get_gauge(uint256 arg0) view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) GetGauge(arg0 *big.Int) (common.Address, error) {
	return _CurveGaugeFactory.Contract.GetGauge(&_CurveGaugeFactory.CallOpts, arg0)
}

// GetGaugeCount is a free data retrieval call binding the contract method 0xf111569c.
//
// Solidity: function get_gauge_count() view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) GetGaugeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "get_gauge_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGaugeCount is a free data retrieval call binding the contract method 0xf111569c.
//
// Solidity: function get_gauge_count() view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactorySession) GetGaugeCount() (*big.Int, error) {
	return _CurveGaugeFactory.Contract.GetGaugeCount(&_CurveGaugeFactory.CallOpts)
}

// GetGaugeCount is a free data retrieval call binding the contract method 0xf111569c.
//
// Solidity: function get_gauge_count() view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) GetGaugeCount() (*big.Int, error) {
	return _CurveGaugeFactory.Contract.GetGaugeCount(&_CurveGaugeFactory.CallOpts)
}

// GetGaugeFromLpToken is a free data retrieval call binding the contract method 0x5d95c65e.
//
// Solidity: function get_gauge_from_lp_token(address arg0) view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) GetGaugeFromLpToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "get_gauge_from_lp_token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGaugeFromLpToken is a free data retrieval call binding the contract method 0x5d95c65e.
//
// Solidity: function get_gauge_from_lp_token(address arg0) view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactorySession) GetGaugeFromLpToken(arg0 common.Address) (common.Address, error) {
	return _CurveGaugeFactory.Contract.GetGaugeFromLpToken(&_CurveGaugeFactory.CallOpts, arg0)
}

// GetGaugeFromLpToken is a free data retrieval call binding the contract method 0x5d95c65e.
//
// Solidity: function get_gauge_from_lp_token(address arg0) view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) GetGaugeFromLpToken(arg0 common.Address) (common.Address, error) {
	return _CurveGaugeFactory.Contract.GetGaugeFromLpToken(&_CurveGaugeFactory.CallOpts, arg0)
}

// GetImplementation is a free data retrieval call binding the contract method 0xc781c668.
//
// Solidity: function get_implementation() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "get_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xc781c668.
//
// Solidity: function get_implementation() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactorySession) GetImplementation() (common.Address, error) {
	return _CurveGaugeFactory.Contract.GetImplementation(&_CurveGaugeFactory.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xc781c668.
//
// Solidity: function get_implementation() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) GetImplementation() (common.Address, error) {
	return _CurveGaugeFactory.Contract.GetImplementation(&_CurveGaugeFactory.CallOpts)
}

// IsMirrored is a free data retrieval call binding the contract method 0x8a42bd82.
//
// Solidity: function is_mirrored(address _gauge) view returns(bool)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) IsMirrored(opts *bind.CallOpts, _gauge common.Address) (bool, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "is_mirrored", _gauge)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMirrored is a free data retrieval call binding the contract method 0x8a42bd82.
//
// Solidity: function is_mirrored(address _gauge) view returns(bool)
func (_CurveGaugeFactory *CurveGaugeFactorySession) IsMirrored(_gauge common.Address) (bool, error) {
	return _CurveGaugeFactory.Contract.IsMirrored(&_CurveGaugeFactory.CallOpts, _gauge)
}

// IsMirrored is a free data retrieval call binding the contract method 0x8a42bd82.
//
// Solidity: function is_mirrored(address _gauge) view returns(bool)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) IsMirrored(_gauge common.Address) (bool, error) {
	return _CurveGaugeFactory.Contract.IsMirrored(&_CurveGaugeFactory.CallOpts, _gauge)
}

// IsValidGauge is a free data retrieval call binding the contract method 0x4b920379.
//
// Solidity: function is_valid_gauge(address _gauge) view returns(bool)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) IsValidGauge(opts *bind.CallOpts, _gauge common.Address) (bool, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "is_valid_gauge", _gauge)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidGauge is a free data retrieval call binding the contract method 0x4b920379.
//
// Solidity: function is_valid_gauge(address _gauge) view returns(bool)
func (_CurveGaugeFactory *CurveGaugeFactorySession) IsValidGauge(_gauge common.Address) (bool, error) {
	return _CurveGaugeFactory.Contract.IsValidGauge(&_CurveGaugeFactory.CallOpts, _gauge)
}

// IsValidGauge is a free data retrieval call binding the contract method 0x4b920379.
//
// Solidity: function is_valid_gauge(address _gauge) view returns(bool)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) IsValidGauge(_gauge common.Address) (bool, error) {
	return _CurveGaugeFactory.Contract.IsValidGauge(&_CurveGaugeFactory.CallOpts, _gauge)
}

// LastRequest is a free data retrieval call binding the contract method 0x51bd4db5.
//
// Solidity: function last_request(address _gauge) view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) LastRequest(opts *bind.CallOpts, _gauge common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "last_request", _gauge)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequest is a free data retrieval call binding the contract method 0x51bd4db5.
//
// Solidity: function last_request(address _gauge) view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactorySession) LastRequest(_gauge common.Address) (*big.Int, error) {
	return _CurveGaugeFactory.Contract.LastRequest(&_CurveGaugeFactory.CallOpts, _gauge)
}

// LastRequest is a free data retrieval call binding the contract method 0x51bd4db5.
//
// Solidity: function last_request(address _gauge) view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) LastRequest(_gauge common.Address) (*big.Int, error) {
	return _CurveGaugeFactory.Contract.LastRequest(&_CurveGaugeFactory.CallOpts, _gauge)
}

// Minted is a free data retrieval call binding the contract method 0x8b752bb0.
//
// Solidity: function minted(address arg0, address arg1) view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) Minted(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "minted", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x8b752bb0.
//
// Solidity: function minted(address arg0, address arg1) view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactorySession) Minted(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGaugeFactory.Contract.Minted(&_CurveGaugeFactory.CallOpts, arg0, arg1)
}

// Minted is a free data retrieval call binding the contract method 0x8b752bb0.
//
// Solidity: function minted(address arg0, address arg1) view returns(uint256)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) Minted(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGaugeFactory.Contract.Minted(&_CurveGaugeFactory.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactorySession) Owner() (common.Address, error) {
	return _CurveGaugeFactory.Contract.Owner(&_CurveGaugeFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) Owner() (common.Address, error) {
	return _CurveGaugeFactory.Contract.Owner(&_CurveGaugeFactory.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCaller) VotingEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGaugeFactory.contract.Call(opts, &out, "voting_escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactorySession) VotingEscrow() (common.Address, error) {
	return _CurveGaugeFactory.Contract.VotingEscrow(&_CurveGaugeFactory.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryCallerSession) VotingEscrow() (common.Address, error) {
	return _CurveGaugeFactory.Contract.VotingEscrow(&_CurveGaugeFactory.CallOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactor) AcceptTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGaugeFactory.contract.Transact(opts, "accept_transfer_ownership")
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveGaugeFactory *CurveGaugeFactorySession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.AcceptTransferOwnership(&_CurveGaugeFactory.TransactOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.AcceptTransferOwnership(&_CurveGaugeFactory.TransactOpts)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _future_owner) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _future_owner common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.contract.Transact(opts, "commit_transfer_ownership", _future_owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _future_owner) returns()
func (_CurveGaugeFactory *CurveGaugeFactorySession) CommitTransferOwnership(_future_owner common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.CommitTransferOwnership(&_CurveGaugeFactory.TransactOpts, _future_owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _future_owner) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorSession) CommitTransferOwnership(_future_owner common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.CommitTransferOwnership(&_CurveGaugeFactory.TransactOpts, _future_owner)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x6be320d2.
//
// Solidity: function deploy_gauge(address _lp_token, bytes32 _salt, address _manager) returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryTransactor) DeployGauge(opts *bind.TransactOpts, _lp_token common.Address, _salt [32]byte, _manager common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.contract.Transact(opts, "deploy_gauge", _lp_token, _salt, _manager)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x6be320d2.
//
// Solidity: function deploy_gauge(address _lp_token, bytes32 _salt, address _manager) returns(address)
func (_CurveGaugeFactory *CurveGaugeFactorySession) DeployGauge(_lp_token common.Address, _salt [32]byte, _manager common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.DeployGauge(&_CurveGaugeFactory.TransactOpts, _lp_token, _salt, _manager)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x6be320d2.
//
// Solidity: function deploy_gauge(address _lp_token, bytes32 _salt, address _manager) returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorSession) DeployGauge(_lp_token common.Address, _salt [32]byte, _manager common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.DeployGauge(&_CurveGaugeFactory.TransactOpts, _lp_token, _salt, _manager)
}

// DeployGauge0 is a paid mutator transaction binding the contract method 0x8db98b5c.
//
// Solidity: function deploy_gauge(address _lp_token, bytes32 _salt) returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryTransactor) DeployGauge0(opts *bind.TransactOpts, _lp_token common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _CurveGaugeFactory.contract.Transact(opts, "deploy_gauge0", _lp_token, _salt)
}

// DeployGauge0 is a paid mutator transaction binding the contract method 0x8db98b5c.
//
// Solidity: function deploy_gauge(address _lp_token, bytes32 _salt) returns(address)
func (_CurveGaugeFactory *CurveGaugeFactorySession) DeployGauge0(_lp_token common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.DeployGauge0(&_CurveGaugeFactory.TransactOpts, _lp_token, _salt)
}

// DeployGauge0 is a paid mutator transaction binding the contract method 0x8db98b5c.
//
// Solidity: function deploy_gauge(address _lp_token, bytes32 _salt) returns(address)
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorSession) DeployGauge0(_lp_token common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.DeployGauge0(&_CurveGaugeFactory.TransactOpts, _lp_token, _salt)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _gauge) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactor) Mint(opts *bind.TransactOpts, _gauge common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.contract.Transact(opts, "mint", _gauge)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _gauge) returns()
func (_CurveGaugeFactory *CurveGaugeFactorySession) Mint(_gauge common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.Mint(&_CurveGaugeFactory.TransactOpts, _gauge)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _gauge) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorSession) Mint(_gauge common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.Mint(&_CurveGaugeFactory.TransactOpts, _gauge)
}

// MintMany is a paid mutator transaction binding the contract method 0x55ec6708.
//
// Solidity: function mint_many(address[32] _gauges) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactor) MintMany(opts *bind.TransactOpts, _gauges [32]common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.contract.Transact(opts, "mint_many", _gauges)
}

// MintMany is a paid mutator transaction binding the contract method 0x55ec6708.
//
// Solidity: function mint_many(address[32] _gauges) returns()
func (_CurveGaugeFactory *CurveGaugeFactorySession) MintMany(_gauges [32]common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.MintMany(&_CurveGaugeFactory.TransactOpts, _gauges)
}

// MintMany is a paid mutator transaction binding the contract method 0x55ec6708.
//
// Solidity: function mint_many(address[32] _gauges) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorSession) MintMany(_gauges [32]common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.MintMany(&_CurveGaugeFactory.TransactOpts, _gauges)
}

// SetCallProxy is a paid mutator transaction binding the contract method 0x5ecb9e14.
//
// Solidity: function set_call_proxy(address _new_call_proxy) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactor) SetCallProxy(opts *bind.TransactOpts, _new_call_proxy common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.contract.Transact(opts, "set_call_proxy", _new_call_proxy)
}

// SetCallProxy is a paid mutator transaction binding the contract method 0x5ecb9e14.
//
// Solidity: function set_call_proxy(address _new_call_proxy) returns()
func (_CurveGaugeFactory *CurveGaugeFactorySession) SetCallProxy(_new_call_proxy common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.SetCallProxy(&_CurveGaugeFactory.TransactOpts, _new_call_proxy)
}

// SetCallProxy is a paid mutator transaction binding the contract method 0x5ecb9e14.
//
// Solidity: function set_call_proxy(address _new_call_proxy) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorSession) SetCallProxy(_new_call_proxy common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.SetCallProxy(&_CurveGaugeFactory.TransactOpts, _new_call_proxy)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x4cd69da0.
//
// Solidity: function set_implementation(address _implementation) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactor) SetImplementation(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.contract.Transact(opts, "set_implementation", _implementation)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x4cd69da0.
//
// Solidity: function set_implementation(address _implementation) returns()
func (_CurveGaugeFactory *CurveGaugeFactorySession) SetImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.SetImplementation(&_CurveGaugeFactory.TransactOpts, _implementation)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x4cd69da0.
//
// Solidity: function set_implementation(address _implementation) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorSession) SetImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.SetImplementation(&_CurveGaugeFactory.TransactOpts, _implementation)
}

// SetMirrored is a paid mutator transaction binding the contract method 0x4b29cac8.
//
// Solidity: function set_mirrored(address _gauge, bool _mirrored) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactor) SetMirrored(opts *bind.TransactOpts, _gauge common.Address, _mirrored bool) (*types.Transaction, error) {
	return _CurveGaugeFactory.contract.Transact(opts, "set_mirrored", _gauge, _mirrored)
}

// SetMirrored is a paid mutator transaction binding the contract method 0x4b29cac8.
//
// Solidity: function set_mirrored(address _gauge, bool _mirrored) returns()
func (_CurveGaugeFactory *CurveGaugeFactorySession) SetMirrored(_gauge common.Address, _mirrored bool) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.SetMirrored(&_CurveGaugeFactory.TransactOpts, _gauge, _mirrored)
}

// SetMirrored is a paid mutator transaction binding the contract method 0x4b29cac8.
//
// Solidity: function set_mirrored(address _gauge, bool _mirrored) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorSession) SetMirrored(_gauge common.Address, _mirrored bool) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.SetMirrored(&_CurveGaugeFactory.TransactOpts, _gauge, _mirrored)
}

// SetVotingEscrow is a paid mutator transaction binding the contract method 0x23fc5a47.
//
// Solidity: function set_voting_escrow(address _voting_escrow) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactor) SetVotingEscrow(opts *bind.TransactOpts, _voting_escrow common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.contract.Transact(opts, "set_voting_escrow", _voting_escrow)
}

// SetVotingEscrow is a paid mutator transaction binding the contract method 0x23fc5a47.
//
// Solidity: function set_voting_escrow(address _voting_escrow) returns()
func (_CurveGaugeFactory *CurveGaugeFactorySession) SetVotingEscrow(_voting_escrow common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.SetVotingEscrow(&_CurveGaugeFactory.TransactOpts, _voting_escrow)
}

// SetVotingEscrow is a paid mutator transaction binding the contract method 0x23fc5a47.
//
// Solidity: function set_voting_escrow(address _voting_escrow) returns()
func (_CurveGaugeFactory *CurveGaugeFactoryTransactorSession) SetVotingEscrow(_voting_escrow common.Address) (*types.Transaction, error) {
	return _CurveGaugeFactory.Contract.SetVotingEscrow(&_CurveGaugeFactory.TransactOpts, _voting_escrow)
}
