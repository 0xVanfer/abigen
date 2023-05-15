// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveDefaultReserveInterestRateStrategy

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

// IDefaultReserveInterestRateStrategyCalculateInterestRatesParams is an auto generated low-level Go binding around an user-defined struct.
type IDefaultReserveInterestRateStrategyCalculateInterestRatesParams struct {
	Unbacked                *big.Int
	LiquidityAdded          *big.Int
	LiquidityTaken          *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	AverageStableBorrowRate *big.Int
	ReserveFactor           *big.Int
	Reserve                 common.Address
	AToken                  common.Address
}

// AaveDefaultReserveInterestRateStrategyMetaData contains all meta data concerning the AaveDefaultReserveInterestRateStrategy contract.
var AaveDefaultReserveInterestRateStrategyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_EXCESS_USAGE_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPTIMAL_USAGE_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"unbacked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityAdded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityTaken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aToken\",\"type\":\"address\"}],\"internalType\":\"structIDefaultReserveInterestRateStrategy.CalculateInterestRatesParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"calculateInterestRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseStableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseVariableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxVariableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStableRateExcessOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStableRateSlope1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStableRateSlope2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVariableRateSlope1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVariableRateSlope2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AaveDefaultReserveInterestRateStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveDefaultReserveInterestRateStrategyMetaData.ABI instead.
var AaveDefaultReserveInterestRateStrategyABI = AaveDefaultReserveInterestRateStrategyMetaData.ABI

// AaveDefaultReserveInterestRateStrategy is an auto generated Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategy struct {
	AaveDefaultReserveInterestRateStrategyCaller     // Read-only binding to the contract
	AaveDefaultReserveInterestRateStrategyTransactor // Write-only binding to the contract
	AaveDefaultReserveInterestRateStrategyFilterer   // Log filterer for contract events
}

// AaveDefaultReserveInterestRateStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveDefaultReserveInterestRateStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveDefaultReserveInterestRateStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveDefaultReserveInterestRateStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveDefaultReserveInterestRateStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveDefaultReserveInterestRateStrategySession struct {
	Contract     *AaveDefaultReserveInterestRateStrategy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                           // Call options to use throughout this session
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// AaveDefaultReserveInterestRateStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveDefaultReserveInterestRateStrategyCallerSession struct {
	Contract *AaveDefaultReserveInterestRateStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                 // Call options to use throughout this session
}

// AaveDefaultReserveInterestRateStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveDefaultReserveInterestRateStrategyTransactorSession struct {
	Contract     *AaveDefaultReserveInterestRateStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                 // Transaction auth options to use throughout this session
}

// AaveDefaultReserveInterestRateStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategyRaw struct {
	Contract *AaveDefaultReserveInterestRateStrategy // Generic contract binding to access the raw methods on
}

// AaveDefaultReserveInterestRateStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategyCallerRaw struct {
	Contract *AaveDefaultReserveInterestRateStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// AaveDefaultReserveInterestRateStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategyTransactorRaw struct {
	Contract *AaveDefaultReserveInterestRateStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveDefaultReserveInterestRateStrategy creates a new instance of AaveDefaultReserveInterestRateStrategy, bound to a specific deployed contract.
func NewAaveDefaultReserveInterestRateStrategy(address common.Address, backend bind.ContractBackend) (*AaveDefaultReserveInterestRateStrategy, error) {
	contract, err := bindAaveDefaultReserveInterestRateStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveDefaultReserveInterestRateStrategy{AaveDefaultReserveInterestRateStrategyCaller: AaveDefaultReserveInterestRateStrategyCaller{contract: contract}, AaveDefaultReserveInterestRateStrategyTransactor: AaveDefaultReserveInterestRateStrategyTransactor{contract: contract}, AaveDefaultReserveInterestRateStrategyFilterer: AaveDefaultReserveInterestRateStrategyFilterer{contract: contract}}, nil
}

// NewAaveDefaultReserveInterestRateStrategyCaller creates a new read-only instance of AaveDefaultReserveInterestRateStrategy, bound to a specific deployed contract.
func NewAaveDefaultReserveInterestRateStrategyCaller(address common.Address, caller bind.ContractCaller) (*AaveDefaultReserveInterestRateStrategyCaller, error) {
	contract, err := bindAaveDefaultReserveInterestRateStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveDefaultReserveInterestRateStrategyCaller{contract: contract}, nil
}

// NewAaveDefaultReserveInterestRateStrategyTransactor creates a new write-only instance of AaveDefaultReserveInterestRateStrategy, bound to a specific deployed contract.
func NewAaveDefaultReserveInterestRateStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveDefaultReserveInterestRateStrategyTransactor, error) {
	contract, err := bindAaveDefaultReserveInterestRateStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveDefaultReserveInterestRateStrategyTransactor{contract: contract}, nil
}

// NewAaveDefaultReserveInterestRateStrategyFilterer creates a new log filterer instance of AaveDefaultReserveInterestRateStrategy, bound to a specific deployed contract.
func NewAaveDefaultReserveInterestRateStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveDefaultReserveInterestRateStrategyFilterer, error) {
	contract, err := bindAaveDefaultReserveInterestRateStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveDefaultReserveInterestRateStrategyFilterer{contract: contract}, nil
}

// bindAaveDefaultReserveInterestRateStrategy binds a generic wrapper to an already deployed contract.
func bindAaveDefaultReserveInterestRateStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveDefaultReserveInterestRateStrategyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveDefaultReserveInterestRateStrategy.Contract.AaveDefaultReserveInterestRateStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.AaveDefaultReserveInterestRateStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.AaveDefaultReserveInterestRateStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveDefaultReserveInterestRateStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.ADDRESSESPROVIDER(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.ADDRESSESPROVIDER(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// MAXEXCESSSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0xfe5fd698.
//
// Solidity: function MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) MAXEXCESSSTABLETOTOTALDEBTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXEXCESSSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0xfe5fd698.
//
// Solidity: function MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) MAXEXCESSSTABLETOTOTALDEBTRATIO() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.MAXEXCESSSTABLETOTOTALDEBTRATIO(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// MAXEXCESSSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0xfe5fd698.
//
// Solidity: function MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) MAXEXCESSSTABLETOTOTALDEBTRATIO() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.MAXEXCESSSTABLETOTOTALDEBTRATIO(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// MAXEXCESSUSAGERATIO is a free data retrieval call binding the contract method 0xa9c622f8.
//
// Solidity: function MAX_EXCESS_USAGE_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) MAXEXCESSUSAGERATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "MAX_EXCESS_USAGE_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXEXCESSUSAGERATIO is a free data retrieval call binding the contract method 0xa9c622f8.
//
// Solidity: function MAX_EXCESS_USAGE_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) MAXEXCESSUSAGERATIO() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.MAXEXCESSUSAGERATIO(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// MAXEXCESSUSAGERATIO is a free data retrieval call binding the contract method 0xa9c622f8.
//
// Solidity: function MAX_EXCESS_USAGE_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) MAXEXCESSUSAGERATIO() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.MAXEXCESSUSAGERATIO(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// OPTIMALSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0x6fb92589.
//
// Solidity: function OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) OPTIMALSTABLETOTOTALDEBTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPTIMALSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0x6fb92589.
//
// Solidity: function OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) OPTIMALSTABLETOTOTALDEBTRATIO() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.OPTIMALSTABLETOTOTALDEBTRATIO(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// OPTIMALSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0x6fb92589.
//
// Solidity: function OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) OPTIMALSTABLETOTOTALDEBTRATIO() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.OPTIMALSTABLETOTOTALDEBTRATIO(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// OPTIMALUSAGERATIO is a free data retrieval call binding the contract method 0x54c365c6.
//
// Solidity: function OPTIMAL_USAGE_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) OPTIMALUSAGERATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "OPTIMAL_USAGE_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPTIMALUSAGERATIO is a free data retrieval call binding the contract method 0x54c365c6.
//
// Solidity: function OPTIMAL_USAGE_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) OPTIMALUSAGERATIO() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.OPTIMALUSAGERATIO(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// OPTIMALUSAGERATIO is a free data retrieval call binding the contract method 0x54c365c6.
//
// Solidity: function OPTIMAL_USAGE_RATIO() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) OPTIMALUSAGERATIO() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.OPTIMALUSAGERATIO(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// CalculateInterestRates is a free data retrieval call binding the contract method 0xa5898709.
//
// Solidity: function calculateInterestRates((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address) params) view returns(uint256, uint256, uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) CalculateInterestRates(opts *bind.CallOpts, params IDefaultReserveInterestRateStrategyCalculateInterestRatesParams) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "calculateInterestRates", params)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CalculateInterestRates is a free data retrieval call binding the contract method 0xa5898709.
//
// Solidity: function calculateInterestRates((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address) params) view returns(uint256, uint256, uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) CalculateInterestRates(params IDefaultReserveInterestRateStrategyCalculateInterestRatesParams) (*big.Int, *big.Int, *big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.CalculateInterestRates(&_AaveDefaultReserveInterestRateStrategy.CallOpts, params)
}

// CalculateInterestRates is a free data retrieval call binding the contract method 0xa5898709.
//
// Solidity: function calculateInterestRates((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address) params) view returns(uint256, uint256, uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) CalculateInterestRates(params IDefaultReserveInterestRateStrategyCalculateInterestRatesParams) (*big.Int, *big.Int, *big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.CalculateInterestRates(&_AaveDefaultReserveInterestRateStrategy.CallOpts, params)
}

// GetBaseStableBorrowRate is a free data retrieval call binding the contract method 0xacd78686.
//
// Solidity: function getBaseStableBorrowRate() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) GetBaseStableBorrowRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "getBaseStableBorrowRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseStableBorrowRate is a free data retrieval call binding the contract method 0xacd78686.
//
// Solidity: function getBaseStableBorrowRate() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) GetBaseStableBorrowRate() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetBaseStableBorrowRate(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetBaseStableBorrowRate is a free data retrieval call binding the contract method 0xacd78686.
//
// Solidity: function getBaseStableBorrowRate() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) GetBaseStableBorrowRate() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetBaseStableBorrowRate(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetBaseVariableBorrowRate is a free data retrieval call binding the contract method 0x34762ca5.
//
// Solidity: function getBaseVariableBorrowRate() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) GetBaseVariableBorrowRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "getBaseVariableBorrowRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseVariableBorrowRate is a free data retrieval call binding the contract method 0x34762ca5.
//
// Solidity: function getBaseVariableBorrowRate() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) GetBaseVariableBorrowRate() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetBaseVariableBorrowRate(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetBaseVariableBorrowRate is a free data retrieval call binding the contract method 0x34762ca5.
//
// Solidity: function getBaseVariableBorrowRate() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) GetBaseVariableBorrowRate() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetBaseVariableBorrowRate(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetMaxVariableBorrowRate is a free data retrieval call binding the contract method 0x80031e37.
//
// Solidity: function getMaxVariableBorrowRate() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) GetMaxVariableBorrowRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "getMaxVariableBorrowRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxVariableBorrowRate is a free data retrieval call binding the contract method 0x80031e37.
//
// Solidity: function getMaxVariableBorrowRate() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) GetMaxVariableBorrowRate() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetMaxVariableBorrowRate(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetMaxVariableBorrowRate is a free data retrieval call binding the contract method 0x80031e37.
//
// Solidity: function getMaxVariableBorrowRate() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) GetMaxVariableBorrowRate() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetMaxVariableBorrowRate(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetStableRateExcessOffset is a free data retrieval call binding the contract method 0xbc626908.
//
// Solidity: function getStableRateExcessOffset() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) GetStableRateExcessOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "getStableRateExcessOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStableRateExcessOffset is a free data retrieval call binding the contract method 0xbc626908.
//
// Solidity: function getStableRateExcessOffset() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) GetStableRateExcessOffset() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetStableRateExcessOffset(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetStableRateExcessOffset is a free data retrieval call binding the contract method 0xbc626908.
//
// Solidity: function getStableRateExcessOffset() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) GetStableRateExcessOffset() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetStableRateExcessOffset(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetStableRateSlope1 is a free data retrieval call binding the contract method 0xd5cd7391.
//
// Solidity: function getStableRateSlope1() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) GetStableRateSlope1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "getStableRateSlope1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStableRateSlope1 is a free data retrieval call binding the contract method 0xd5cd7391.
//
// Solidity: function getStableRateSlope1() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) GetStableRateSlope1() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetStableRateSlope1(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetStableRateSlope1 is a free data retrieval call binding the contract method 0xd5cd7391.
//
// Solidity: function getStableRateSlope1() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) GetStableRateSlope1() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetStableRateSlope1(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetStableRateSlope2 is a free data retrieval call binding the contract method 0x14e32da4.
//
// Solidity: function getStableRateSlope2() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) GetStableRateSlope2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "getStableRateSlope2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStableRateSlope2 is a free data retrieval call binding the contract method 0x14e32da4.
//
// Solidity: function getStableRateSlope2() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) GetStableRateSlope2() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetStableRateSlope2(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetStableRateSlope2 is a free data retrieval call binding the contract method 0x14e32da4.
//
// Solidity: function getStableRateSlope2() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) GetStableRateSlope2() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetStableRateSlope2(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetVariableRateSlope1 is a free data retrieval call binding the contract method 0x0b3429a2.
//
// Solidity: function getVariableRateSlope1() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) GetVariableRateSlope1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "getVariableRateSlope1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVariableRateSlope1 is a free data retrieval call binding the contract method 0x0b3429a2.
//
// Solidity: function getVariableRateSlope1() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) GetVariableRateSlope1() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetVariableRateSlope1(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetVariableRateSlope1 is a free data retrieval call binding the contract method 0x0b3429a2.
//
// Solidity: function getVariableRateSlope1() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) GetVariableRateSlope1() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetVariableRateSlope1(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetVariableRateSlope2 is a free data retrieval call binding the contract method 0xf4202409.
//
// Solidity: function getVariableRateSlope2() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCaller) GetVariableRateSlope2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategy.contract.Call(opts, &out, "getVariableRateSlope2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVariableRateSlope2 is a free data retrieval call binding the contract method 0xf4202409.
//
// Solidity: function getVariableRateSlope2() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategySession) GetVariableRateSlope2() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetVariableRateSlope2(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}

// GetVariableRateSlope2 is a free data retrieval call binding the contract method 0xf4202409.
//
// Solidity: function getVariableRateSlope2() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategy *AaveDefaultReserveInterestRateStrategyCallerSession) GetVariableRateSlope2() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategy.Contract.GetVariableRateSlope2(&_AaveDefaultReserveInterestRateStrategy.CallOpts)
}
