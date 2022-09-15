// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KyberFactory

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

// KyberFactoryMetaData contains all meta data concerning the KyberFactory contract.
var KyberFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"ampBps\",\"type\":\"uint32\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeConfiguration\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_governmentFeeBps\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPoolAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"name\":\"getPools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokenPools\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"name\":\"getPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getUnamplifiedPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_governmentFeeBps\",\"type\":\"uint16\"}],\"name\":\"setFeeConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KyberFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use KyberFactoryMetaData.ABI instead.
var KyberFactoryABI = KyberFactoryMetaData.ABI

// KyberFactory is an auto generated Go binding around an Ethereum contract.
type KyberFactory struct {
	KyberFactoryCaller     // Read-only binding to the contract
	KyberFactoryTransactor // Write-only binding to the contract
	KyberFactoryFilterer   // Log filterer for contract events
}

// KyberFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberFactorySession struct {
	Contract     *KyberFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KyberFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberFactoryCallerSession struct {
	Contract *KyberFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KyberFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberFactoryTransactorSession struct {
	Contract     *KyberFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KyberFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberFactoryRaw struct {
	Contract *KyberFactory // Generic contract binding to access the raw methods on
}

// KyberFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberFactoryCallerRaw struct {
	Contract *KyberFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// KyberFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberFactoryTransactorRaw struct {
	Contract *KyberFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberFactory creates a new instance of KyberFactory, bound to a specific deployed contract.
func NewKyberFactory(address common.Address, backend bind.ContractBackend) (*KyberFactory, error) {
	contract, err := bindKyberFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KyberFactory{KyberFactoryCaller: KyberFactoryCaller{contract: contract}, KyberFactoryTransactor: KyberFactoryTransactor{contract: contract}, KyberFactoryFilterer: KyberFactoryFilterer{contract: contract}}, nil
}

// NewKyberFactoryCaller creates a new read-only instance of KyberFactory, bound to a specific deployed contract.
func NewKyberFactoryCaller(address common.Address, caller bind.ContractCaller) (*KyberFactoryCaller, error) {
	contract, err := bindKyberFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberFactoryCaller{contract: contract}, nil
}

// NewKyberFactoryTransactor creates a new write-only instance of KyberFactory, bound to a specific deployed contract.
func NewKyberFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberFactoryTransactor, error) {
	contract, err := bindKyberFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberFactoryTransactor{contract: contract}, nil
}

// NewKyberFactoryFilterer creates a new log filterer instance of KyberFactory, bound to a specific deployed contract.
func NewKyberFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberFactoryFilterer, error) {
	contract, err := bindKyberFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberFactoryFilterer{contract: contract}, nil
}

// bindKyberFactory binds a generic wrapper to an already deployed contract.
func bindKyberFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberFactory *KyberFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberFactory.Contract.KyberFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberFactory *KyberFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberFactory.Contract.KyberFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberFactory *KyberFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberFactory.Contract.KyberFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberFactory *KyberFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberFactory *KyberFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberFactory *KyberFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberFactory.Contract.contract.Transact(opts, method, params...)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_KyberFactory *KyberFactoryCaller) AllPools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KyberFactory.contract.Call(opts, &out, "allPools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_KyberFactory *KyberFactorySession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _KyberFactory.Contract.AllPools(&_KyberFactory.CallOpts, arg0)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_KyberFactory *KyberFactoryCallerSession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _KyberFactory.Contract.AllPools(&_KyberFactory.CallOpts, arg0)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_KyberFactory *KyberFactoryCaller) AllPoolsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KyberFactory.contract.Call(opts, &out, "allPoolsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_KyberFactory *KyberFactorySession) AllPoolsLength() (*big.Int, error) {
	return _KyberFactory.Contract.AllPoolsLength(&_KyberFactory.CallOpts)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_KyberFactory *KyberFactoryCallerSession) AllPoolsLength() (*big.Int, error) {
	return _KyberFactory.Contract.AllPoolsLength(&_KyberFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_KyberFactory *KyberFactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KyberFactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_KyberFactory *KyberFactorySession) FeeToSetter() (common.Address, error) {
	return _KyberFactory.Contract.FeeToSetter(&_KyberFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_KyberFactory *KyberFactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _KyberFactory.Contract.FeeToSetter(&_KyberFactory.CallOpts)
}

// GetFeeConfiguration is a free data retrieval call binding the contract method 0xad5cb2e1.
//
// Solidity: function getFeeConfiguration() view returns(address _feeTo, uint16 _governmentFeeBps)
func (_KyberFactory *KyberFactoryCaller) GetFeeConfiguration(opts *bind.CallOpts) (struct {
	FeeTo            common.Address
	GovernmentFeeBps uint16
}, error) {
	var out []interface{}
	err := _KyberFactory.contract.Call(opts, &out, "getFeeConfiguration")

	outstruct := new(struct {
		FeeTo            common.Address
		GovernmentFeeBps uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeTo = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.GovernmentFeeBps = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetFeeConfiguration is a free data retrieval call binding the contract method 0xad5cb2e1.
//
// Solidity: function getFeeConfiguration() view returns(address _feeTo, uint16 _governmentFeeBps)
func (_KyberFactory *KyberFactorySession) GetFeeConfiguration() (struct {
	FeeTo            common.Address
	GovernmentFeeBps uint16
}, error) {
	return _KyberFactory.Contract.GetFeeConfiguration(&_KyberFactory.CallOpts)
}

// GetFeeConfiguration is a free data retrieval call binding the contract method 0xad5cb2e1.
//
// Solidity: function getFeeConfiguration() view returns(address _feeTo, uint16 _governmentFeeBps)
func (_KyberFactory *KyberFactoryCallerSession) GetFeeConfiguration() (struct {
	FeeTo            common.Address
	GovernmentFeeBps uint16
}, error) {
	return _KyberFactory.Contract.GetFeeConfiguration(&_KyberFactory.CallOpts)
}

// GetPoolAtIndex is a free data retrieval call binding the contract method 0x65da9289.
//
// Solidity: function getPoolAtIndex(address token0, address token1, uint256 index) view returns(address pool)
func (_KyberFactory *KyberFactoryCaller) GetPoolAtIndex(opts *bind.CallOpts, token0 common.Address, token1 common.Address, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KyberFactory.contract.Call(opts, &out, "getPoolAtIndex", token0, token1, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAtIndex is a free data retrieval call binding the contract method 0x65da9289.
//
// Solidity: function getPoolAtIndex(address token0, address token1, uint256 index) view returns(address pool)
func (_KyberFactory *KyberFactorySession) GetPoolAtIndex(token0 common.Address, token1 common.Address, index *big.Int) (common.Address, error) {
	return _KyberFactory.Contract.GetPoolAtIndex(&_KyberFactory.CallOpts, token0, token1, index)
}

// GetPoolAtIndex is a free data retrieval call binding the contract method 0x65da9289.
//
// Solidity: function getPoolAtIndex(address token0, address token1, uint256 index) view returns(address pool)
func (_KyberFactory *KyberFactoryCallerSession) GetPoolAtIndex(token0 common.Address, token1 common.Address, index *big.Int) (common.Address, error) {
	return _KyberFactory.Contract.GetPoolAtIndex(&_KyberFactory.CallOpts, token0, token1, index)
}

// GetPools is a free data retrieval call binding the contract method 0x5b1dc86f.
//
// Solidity: function getPools(address token0, address token1) view returns(address[] _tokenPools)
func (_KyberFactory *KyberFactoryCaller) GetPools(opts *bind.CallOpts, token0 common.Address, token1 common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _KyberFactory.contract.Call(opts, &out, "getPools", token0, token1)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPools is a free data retrieval call binding the contract method 0x5b1dc86f.
//
// Solidity: function getPools(address token0, address token1) view returns(address[] _tokenPools)
func (_KyberFactory *KyberFactorySession) GetPools(token0 common.Address, token1 common.Address) ([]common.Address, error) {
	return _KyberFactory.Contract.GetPools(&_KyberFactory.CallOpts, token0, token1)
}

// GetPools is a free data retrieval call binding the contract method 0x5b1dc86f.
//
// Solidity: function getPools(address token0, address token1) view returns(address[] _tokenPools)
func (_KyberFactory *KyberFactoryCallerSession) GetPools(token0 common.Address, token1 common.Address) ([]common.Address, error) {
	return _KyberFactory.Contract.GetPools(&_KyberFactory.CallOpts, token0, token1)
}

// GetPoolsLength is a free data retrieval call binding the contract method 0x538633df.
//
// Solidity: function getPoolsLength(address token0, address token1) view returns(uint256)
func (_KyberFactory *KyberFactoryCaller) GetPoolsLength(opts *bind.CallOpts, token0 common.Address, token1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KyberFactory.contract.Call(opts, &out, "getPoolsLength", token0, token1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolsLength is a free data retrieval call binding the contract method 0x538633df.
//
// Solidity: function getPoolsLength(address token0, address token1) view returns(uint256)
func (_KyberFactory *KyberFactorySession) GetPoolsLength(token0 common.Address, token1 common.Address) (*big.Int, error) {
	return _KyberFactory.Contract.GetPoolsLength(&_KyberFactory.CallOpts, token0, token1)
}

// GetPoolsLength is a free data retrieval call binding the contract method 0x538633df.
//
// Solidity: function getPoolsLength(address token0, address token1) view returns(uint256)
func (_KyberFactory *KyberFactoryCallerSession) GetPoolsLength(token0 common.Address, token1 common.Address) (*big.Int, error) {
	return _KyberFactory.Contract.GetPoolsLength(&_KyberFactory.CallOpts, token0, token1)
}

// GetUnamplifiedPool is a free data retrieval call binding the contract method 0x3d82497e.
//
// Solidity: function getUnamplifiedPool(address , address ) view returns(address)
func (_KyberFactory *KyberFactoryCaller) GetUnamplifiedPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _KyberFactory.contract.Call(opts, &out, "getUnamplifiedPool", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUnamplifiedPool is a free data retrieval call binding the contract method 0x3d82497e.
//
// Solidity: function getUnamplifiedPool(address , address ) view returns(address)
func (_KyberFactory *KyberFactorySession) GetUnamplifiedPool(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _KyberFactory.Contract.GetUnamplifiedPool(&_KyberFactory.CallOpts, arg0, arg1)
}

// GetUnamplifiedPool is a free data retrieval call binding the contract method 0x3d82497e.
//
// Solidity: function getUnamplifiedPool(address , address ) view returns(address)
func (_KyberFactory *KyberFactoryCallerSession) GetUnamplifiedPool(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _KyberFactory.Contract.GetUnamplifiedPool(&_KyberFactory.CallOpts, arg0, arg1)
}

// IsPool is a free data retrieval call binding the contract method 0xeb787f61.
//
// Solidity: function isPool(address token0, address token1, address pool) view returns(bool)
func (_KyberFactory *KyberFactoryCaller) IsPool(opts *bind.CallOpts, token0 common.Address, token1 common.Address, pool common.Address) (bool, error) {
	var out []interface{}
	err := _KyberFactory.contract.Call(opts, &out, "isPool", token0, token1, pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPool is a free data retrieval call binding the contract method 0xeb787f61.
//
// Solidity: function isPool(address token0, address token1, address pool) view returns(bool)
func (_KyberFactory *KyberFactorySession) IsPool(token0 common.Address, token1 common.Address, pool common.Address) (bool, error) {
	return _KyberFactory.Contract.IsPool(&_KyberFactory.CallOpts, token0, token1, pool)
}

// IsPool is a free data retrieval call binding the contract method 0xeb787f61.
//
// Solidity: function isPool(address token0, address token1, address pool) view returns(bool)
func (_KyberFactory *KyberFactoryCallerSession) IsPool(token0 common.Address, token1 common.Address, pool common.Address) (bool, error) {
	return _KyberFactory.Contract.IsPool(&_KyberFactory.CallOpts, token0, token1, pool)
}

// CreatePool is a paid mutator transaction binding the contract method 0x8fd64840.
//
// Solidity: function createPool(address tokenA, address tokenB, uint32 ampBps) returns(address pool)
func (_KyberFactory *KyberFactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, ampBps uint32) (*types.Transaction, error) {
	return _KyberFactory.contract.Transact(opts, "createPool", tokenA, tokenB, ampBps)
}

// CreatePool is a paid mutator transaction binding the contract method 0x8fd64840.
//
// Solidity: function createPool(address tokenA, address tokenB, uint32 ampBps) returns(address pool)
func (_KyberFactory *KyberFactorySession) CreatePool(tokenA common.Address, tokenB common.Address, ampBps uint32) (*types.Transaction, error) {
	return _KyberFactory.Contract.CreatePool(&_KyberFactory.TransactOpts, tokenA, tokenB, ampBps)
}

// CreatePool is a paid mutator transaction binding the contract method 0x8fd64840.
//
// Solidity: function createPool(address tokenA, address tokenB, uint32 ampBps) returns(address pool)
func (_KyberFactory *KyberFactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, ampBps uint32) (*types.Transaction, error) {
	return _KyberFactory.Contract.CreatePool(&_KyberFactory.TransactOpts, tokenA, tokenB, ampBps)
}

// SetFeeConfiguration is a paid mutator transaction binding the contract method 0x2900909d.
//
// Solidity: function setFeeConfiguration(address _feeTo, uint16 _governmentFeeBps) returns()
func (_KyberFactory *KyberFactoryTransactor) SetFeeConfiguration(opts *bind.TransactOpts, _feeTo common.Address, _governmentFeeBps uint16) (*types.Transaction, error) {
	return _KyberFactory.contract.Transact(opts, "setFeeConfiguration", _feeTo, _governmentFeeBps)
}

// SetFeeConfiguration is a paid mutator transaction binding the contract method 0x2900909d.
//
// Solidity: function setFeeConfiguration(address _feeTo, uint16 _governmentFeeBps) returns()
func (_KyberFactory *KyberFactorySession) SetFeeConfiguration(_feeTo common.Address, _governmentFeeBps uint16) (*types.Transaction, error) {
	return _KyberFactory.Contract.SetFeeConfiguration(&_KyberFactory.TransactOpts, _feeTo, _governmentFeeBps)
}

// SetFeeConfiguration is a paid mutator transaction binding the contract method 0x2900909d.
//
// Solidity: function setFeeConfiguration(address _feeTo, uint16 _governmentFeeBps) returns()
func (_KyberFactory *KyberFactoryTransactorSession) SetFeeConfiguration(_feeTo common.Address, _governmentFeeBps uint16) (*types.Transaction, error) {
	return _KyberFactory.Contract.SetFeeConfiguration(&_KyberFactory.TransactOpts, _feeTo, _governmentFeeBps)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_KyberFactory *KyberFactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _KyberFactory.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_KyberFactory *KyberFactorySession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _KyberFactory.Contract.SetFeeToSetter(&_KyberFactory.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_KyberFactory *KyberFactoryTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _KyberFactory.Contract.SetFeeToSetter(&_KyberFactory.TransactOpts, _feeToSetter)
}
