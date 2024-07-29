// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveDefaultReserveInterestRateStrategyV2

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

// IInterestRateStrategyV2CalcInterestRatesLocalVars is an auto generated low-level Go binding around an user-defined struct.
type IInterestRateStrategyV2CalcInterestRatesLocalVars struct {
	AvailableLiquidity         *big.Int
	TotalDebt                  *big.Int
	CurrentVariableBorrowRate  *big.Int
	CurrentLiquidityRate       *big.Int
	BorrowUsageRatio           *big.Int
	SupplyUsageRatio           *big.Int
	AvailableLiquidityPlusDebt *big.Int
}

// IInterestRateStrategyV2InterestRateData is an auto generated low-level Go binding around an user-defined struct.
type IInterestRateStrategyV2InterestRateData struct {
	OptimalUsageRatio      uint16
	BaseVariableBorrowRate uint32
	VariableRateSlope1     uint32
	VariableRateSlope2     uint32
}

// AaveDefaultReserveInterestRateStrategyV2MetaData contains all meta data concerning the AaveDefaultReserveInterestRateStrategyV2 contract.
var AaveDefaultReserveInterestRateStrategyV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optimalUsageRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseVariableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableRateSlope1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableRateSlope2\",\"type\":\"uint256\"}],\"name\":\"RateDataUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BORROW_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_OPTIMAL_POINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_OPTIMAL_POINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowUsageRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyUsageRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidityPlusDebt\",\"type\":\"uint256\"}],\"internalType\":\"structIInterestRateStrategyV2.CalcInterestRatesLocalVars\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"calculateInterestRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"getBaseVariableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"getInterestRateData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"optimalUsageRatio\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"baseVariableBorrowRate\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"variableRateSlope1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"variableRateSlope2\",\"type\":\"uint32\"}],\"internalType\":\"structIInterestRateStrategyV2.InterestRateData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"getInterestRateDataBps\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"optimalUsageRatio\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"baseVariableBorrowRate\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"variableRateSlope1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"variableRateSlope2\",\"type\":\"uint32\"}],\"internalType\":\"structIInterestRateStrategyV2.InterestRateData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"getMaxVariableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"getOptimalUsageRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"getVariableRateSlope1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"getVariableRateSlope2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AaveDefaultReserveInterestRateStrategyV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveDefaultReserveInterestRateStrategyV2MetaData.ABI instead.
var AaveDefaultReserveInterestRateStrategyV2ABI = AaveDefaultReserveInterestRateStrategyV2MetaData.ABI

// AaveDefaultReserveInterestRateStrategyV2 is an auto generated Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategyV2 struct {
	AaveDefaultReserveInterestRateStrategyV2Caller     // Read-only binding to the contract
	AaveDefaultReserveInterestRateStrategyV2Transactor // Write-only binding to the contract
	AaveDefaultReserveInterestRateStrategyV2Filterer   // Log filterer for contract events
}

// AaveDefaultReserveInterestRateStrategyV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategyV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveDefaultReserveInterestRateStrategyV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategyV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveDefaultReserveInterestRateStrategyV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveDefaultReserveInterestRateStrategyV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveDefaultReserveInterestRateStrategyV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveDefaultReserveInterestRateStrategyV2Session struct {
	Contract     *AaveDefaultReserveInterestRateStrategyV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                             // Call options to use throughout this session
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// AaveDefaultReserveInterestRateStrategyV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveDefaultReserveInterestRateStrategyV2CallerSession struct {
	Contract *AaveDefaultReserveInterestRateStrategyV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                   // Call options to use throughout this session
}

// AaveDefaultReserveInterestRateStrategyV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveDefaultReserveInterestRateStrategyV2TransactorSession struct {
	Contract     *AaveDefaultReserveInterestRateStrategyV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                   // Transaction auth options to use throughout this session
}

// AaveDefaultReserveInterestRateStrategyV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategyV2Raw struct {
	Contract *AaveDefaultReserveInterestRateStrategyV2 // Generic contract binding to access the raw methods on
}

// AaveDefaultReserveInterestRateStrategyV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategyV2CallerRaw struct {
	Contract *AaveDefaultReserveInterestRateStrategyV2Caller // Generic read-only contract binding to access the raw methods on
}

// AaveDefaultReserveInterestRateStrategyV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveDefaultReserveInterestRateStrategyV2TransactorRaw struct {
	Contract *AaveDefaultReserveInterestRateStrategyV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveDefaultReserveInterestRateStrategyV2 creates a new instance of AaveDefaultReserveInterestRateStrategyV2, bound to a specific deployed contract.
func NewAaveDefaultReserveInterestRateStrategyV2(address common.Address, backend bind.ContractBackend) (*AaveDefaultReserveInterestRateStrategyV2, error) {
	contract, err := bindAaveDefaultReserveInterestRateStrategyV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveDefaultReserveInterestRateStrategyV2{AaveDefaultReserveInterestRateStrategyV2Caller: AaveDefaultReserveInterestRateStrategyV2Caller{contract: contract}, AaveDefaultReserveInterestRateStrategyV2Transactor: AaveDefaultReserveInterestRateStrategyV2Transactor{contract: contract}, AaveDefaultReserveInterestRateStrategyV2Filterer: AaveDefaultReserveInterestRateStrategyV2Filterer{contract: contract}}, nil
}

// NewAaveDefaultReserveInterestRateStrategyV2Caller creates a new read-only instance of AaveDefaultReserveInterestRateStrategyV2, bound to a specific deployed contract.
func NewAaveDefaultReserveInterestRateStrategyV2Caller(address common.Address, caller bind.ContractCaller) (*AaveDefaultReserveInterestRateStrategyV2Caller, error) {
	contract, err := bindAaveDefaultReserveInterestRateStrategyV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveDefaultReserveInterestRateStrategyV2Caller{contract: contract}, nil
}

// NewAaveDefaultReserveInterestRateStrategyV2Transactor creates a new write-only instance of AaveDefaultReserveInterestRateStrategyV2, bound to a specific deployed contract.
func NewAaveDefaultReserveInterestRateStrategyV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveDefaultReserveInterestRateStrategyV2Transactor, error) {
	contract, err := bindAaveDefaultReserveInterestRateStrategyV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveDefaultReserveInterestRateStrategyV2Transactor{contract: contract}, nil
}

// NewAaveDefaultReserveInterestRateStrategyV2Filterer creates a new log filterer instance of AaveDefaultReserveInterestRateStrategyV2, bound to a specific deployed contract.
func NewAaveDefaultReserveInterestRateStrategyV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveDefaultReserveInterestRateStrategyV2Filterer, error) {
	contract, err := bindAaveDefaultReserveInterestRateStrategyV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveDefaultReserveInterestRateStrategyV2Filterer{contract: contract}, nil
}

// bindAaveDefaultReserveInterestRateStrategyV2 binds a generic wrapper to an already deployed contract.
func bindAaveDefaultReserveInterestRateStrategyV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveDefaultReserveInterestRateStrategyV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.AaveDefaultReserveInterestRateStrategyV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.AaveDefaultReserveInterestRateStrategyV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.AaveDefaultReserveInterestRateStrategyV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.ADDRESSESPROVIDER(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.ADDRESSESPROVIDER(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts)
}

// MAXBORROWRATE is a free data retrieval call binding the contract method 0x7a0c5ebf.
//
// Solidity: function MAX_BORROW_RATE() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) MAXBORROWRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "MAX_BORROW_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBORROWRATE is a free data retrieval call binding the contract method 0x7a0c5ebf.
//
// Solidity: function MAX_BORROW_RATE() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) MAXBORROWRATE() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.MAXBORROWRATE(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts)
}

// MAXBORROWRATE is a free data retrieval call binding the contract method 0x7a0c5ebf.
//
// Solidity: function MAX_BORROW_RATE() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) MAXBORROWRATE() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.MAXBORROWRATE(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts)
}

// MAXOPTIMALPOINT is a free data retrieval call binding the contract method 0x7a24bd7e.
//
// Solidity: function MAX_OPTIMAL_POINT() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) MAXOPTIMALPOINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "MAX_OPTIMAL_POINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOPTIMALPOINT is a free data retrieval call binding the contract method 0x7a24bd7e.
//
// Solidity: function MAX_OPTIMAL_POINT() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) MAXOPTIMALPOINT() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.MAXOPTIMALPOINT(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts)
}

// MAXOPTIMALPOINT is a free data retrieval call binding the contract method 0x7a24bd7e.
//
// Solidity: function MAX_OPTIMAL_POINT() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) MAXOPTIMALPOINT() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.MAXOPTIMALPOINT(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts)
}

// MINOPTIMALPOINT is a free data retrieval call binding the contract method 0xf7e0fe67.
//
// Solidity: function MIN_OPTIMAL_POINT() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) MINOPTIMALPOINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "MIN_OPTIMAL_POINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINOPTIMALPOINT is a free data retrieval call binding the contract method 0xf7e0fe67.
//
// Solidity: function MIN_OPTIMAL_POINT() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) MINOPTIMALPOINT() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.MINOPTIMALPOINT(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts)
}

// MINOPTIMALPOINT is a free data retrieval call binding the contract method 0xf7e0fe67.
//
// Solidity: function MIN_OPTIMAL_POINT() view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) MINOPTIMALPOINT() (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.MINOPTIMALPOINT(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts)
}

// CalculateInterestRates is a free data retrieval call binding the contract method 0xbd1b0ced.
//
// Solidity: function calculateInterestRates((uint256,uint256,uint256,uint256,uint256,uint256,uint256) params) view returns(uint256, uint256, uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) CalculateInterestRates(opts *bind.CallOpts, params IInterestRateStrategyV2CalcInterestRatesLocalVars) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "calculateInterestRates", params)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CalculateInterestRates is a free data retrieval call binding the contract method 0xbd1b0ced.
//
// Solidity: function calculateInterestRates((uint256,uint256,uint256,uint256,uint256,uint256,uint256) params) view returns(uint256, uint256, uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) CalculateInterestRates(params IInterestRateStrategyV2CalcInterestRatesLocalVars) (*big.Int, *big.Int, *big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.CalculateInterestRates(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, params)
}

// CalculateInterestRates is a free data retrieval call binding the contract method 0xbd1b0ced.
//
// Solidity: function calculateInterestRates((uint256,uint256,uint256,uint256,uint256,uint256,uint256) params) view returns(uint256, uint256, uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) CalculateInterestRates(params IInterestRateStrategyV2CalcInterestRatesLocalVars) (*big.Int, *big.Int, *big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.CalculateInterestRates(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, params)
}

// GetBaseVariableBorrowRate is a free data retrieval call binding the contract method 0xcca22ea1.
//
// Solidity: function getBaseVariableBorrowRate(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) GetBaseVariableBorrowRate(opts *bind.CallOpts, reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "getBaseVariableBorrowRate", reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseVariableBorrowRate is a free data retrieval call binding the contract method 0xcca22ea1.
//
// Solidity: function getBaseVariableBorrowRate(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) GetBaseVariableBorrowRate(reserve common.Address) (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetBaseVariableBorrowRate(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetBaseVariableBorrowRate is a free data retrieval call binding the contract method 0xcca22ea1.
//
// Solidity: function getBaseVariableBorrowRate(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) GetBaseVariableBorrowRate(reserve common.Address) (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetBaseVariableBorrowRate(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetInterestRateData is a free data retrieval call binding the contract method 0x131e889c.
//
// Solidity: function getInterestRateData(address reserve) view returns((uint16,uint32,uint32,uint32))
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) GetInterestRateData(opts *bind.CallOpts, reserve common.Address) (IInterestRateStrategyV2InterestRateData, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "getInterestRateData", reserve)

	if err != nil {
		return *new(IInterestRateStrategyV2InterestRateData), err
	}

	out0 := *abi.ConvertType(out[0], new(IInterestRateStrategyV2InterestRateData)).(*IInterestRateStrategyV2InterestRateData)

	return out0, err

}

// GetInterestRateData is a free data retrieval call binding the contract method 0x131e889c.
//
// Solidity: function getInterestRateData(address reserve) view returns((uint16,uint32,uint32,uint32))
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) GetInterestRateData(reserve common.Address) (IInterestRateStrategyV2InterestRateData, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetInterestRateData(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetInterestRateData is a free data retrieval call binding the contract method 0x131e889c.
//
// Solidity: function getInterestRateData(address reserve) view returns((uint16,uint32,uint32,uint32))
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) GetInterestRateData(reserve common.Address) (IInterestRateStrategyV2InterestRateData, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetInterestRateData(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetInterestRateDataBps is a free data retrieval call binding the contract method 0xc79ce42e.
//
// Solidity: function getInterestRateDataBps(address reserve) view returns((uint16,uint32,uint32,uint32))
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) GetInterestRateDataBps(opts *bind.CallOpts, reserve common.Address) (IInterestRateStrategyV2InterestRateData, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "getInterestRateDataBps", reserve)

	if err != nil {
		return *new(IInterestRateStrategyV2InterestRateData), err
	}

	out0 := *abi.ConvertType(out[0], new(IInterestRateStrategyV2InterestRateData)).(*IInterestRateStrategyV2InterestRateData)

	return out0, err

}

// GetInterestRateDataBps is a free data retrieval call binding the contract method 0xc79ce42e.
//
// Solidity: function getInterestRateDataBps(address reserve) view returns((uint16,uint32,uint32,uint32))
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) GetInterestRateDataBps(reserve common.Address) (IInterestRateStrategyV2InterestRateData, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetInterestRateDataBps(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetInterestRateDataBps is a free data retrieval call binding the contract method 0xc79ce42e.
//
// Solidity: function getInterestRateDataBps(address reserve) view returns((uint16,uint32,uint32,uint32))
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) GetInterestRateDataBps(reserve common.Address) (IInterestRateStrategyV2InterestRateData, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetInterestRateDataBps(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetMaxVariableBorrowRate is a free data retrieval call binding the contract method 0x6a00178e.
//
// Solidity: function getMaxVariableBorrowRate(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) GetMaxVariableBorrowRate(opts *bind.CallOpts, reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "getMaxVariableBorrowRate", reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxVariableBorrowRate is a free data retrieval call binding the contract method 0x6a00178e.
//
// Solidity: function getMaxVariableBorrowRate(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) GetMaxVariableBorrowRate(reserve common.Address) (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetMaxVariableBorrowRate(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetMaxVariableBorrowRate is a free data retrieval call binding the contract method 0x6a00178e.
//
// Solidity: function getMaxVariableBorrowRate(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) GetMaxVariableBorrowRate(reserve common.Address) (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetMaxVariableBorrowRate(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetOptimalUsageRatio is a free data retrieval call binding the contract method 0xaa33f063.
//
// Solidity: function getOptimalUsageRatio(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) GetOptimalUsageRatio(opts *bind.CallOpts, reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "getOptimalUsageRatio", reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOptimalUsageRatio is a free data retrieval call binding the contract method 0xaa33f063.
//
// Solidity: function getOptimalUsageRatio(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) GetOptimalUsageRatio(reserve common.Address) (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetOptimalUsageRatio(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetOptimalUsageRatio is a free data retrieval call binding the contract method 0xaa33f063.
//
// Solidity: function getOptimalUsageRatio(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) GetOptimalUsageRatio(reserve common.Address) (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetOptimalUsageRatio(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetVariableRateSlope1 is a free data retrieval call binding the contract method 0x5b651bae.
//
// Solidity: function getVariableRateSlope1(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) GetVariableRateSlope1(opts *bind.CallOpts, reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "getVariableRateSlope1", reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVariableRateSlope1 is a free data retrieval call binding the contract method 0x5b651bae.
//
// Solidity: function getVariableRateSlope1(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) GetVariableRateSlope1(reserve common.Address) (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetVariableRateSlope1(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetVariableRateSlope1 is a free data retrieval call binding the contract method 0x5b651bae.
//
// Solidity: function getVariableRateSlope1(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) GetVariableRateSlope1(reserve common.Address) (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetVariableRateSlope1(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetVariableRateSlope2 is a free data retrieval call binding the contract method 0x8f4b0d5d.
//
// Solidity: function getVariableRateSlope2(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Caller) GetVariableRateSlope2(opts *bind.CallOpts, reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveDefaultReserveInterestRateStrategyV2.contract.Call(opts, &out, "getVariableRateSlope2", reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVariableRateSlope2 is a free data retrieval call binding the contract method 0x8f4b0d5d.
//
// Solidity: function getVariableRateSlope2(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Session) GetVariableRateSlope2(reserve common.Address) (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetVariableRateSlope2(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// GetVariableRateSlope2 is a free data retrieval call binding the contract method 0x8f4b0d5d.
//
// Solidity: function getVariableRateSlope2(address reserve) view returns(uint256)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2CallerSession) GetVariableRateSlope2(reserve common.Address) (*big.Int, error) {
	return _AaveDefaultReserveInterestRateStrategyV2.Contract.GetVariableRateSlope2(&_AaveDefaultReserveInterestRateStrategyV2.CallOpts, reserve)
}

// AaveDefaultReserveInterestRateStrategyV2RateDataUpdateIterator is returned from FilterRateDataUpdate and is used to iterate over the raw logs and unpacked data for RateDataUpdate events raised by the AaveDefaultReserveInterestRateStrategyV2 contract.
type AaveDefaultReserveInterestRateStrategyV2RateDataUpdateIterator struct {
	Event *AaveDefaultReserveInterestRateStrategyV2RateDataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveDefaultReserveInterestRateStrategyV2RateDataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveDefaultReserveInterestRateStrategyV2RateDataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveDefaultReserveInterestRateStrategyV2RateDataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveDefaultReserveInterestRateStrategyV2RateDataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveDefaultReserveInterestRateStrategyV2RateDataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveDefaultReserveInterestRateStrategyV2RateDataUpdate represents a RateDataUpdate event raised by the AaveDefaultReserveInterestRateStrategyV2 contract.
type AaveDefaultReserveInterestRateStrategyV2RateDataUpdate struct {
	Reserve                common.Address
	OptimalUsageRatio      *big.Int
	BaseVariableBorrowRate *big.Int
	VariableRateSlope1     *big.Int
	VariableRateSlope2     *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRateDataUpdate is a free log retrieval operation binding the contract event 0x5d123bea2036a4052274206f59d99350b9741e17da56ffae335d809b25ee0942.
//
// Solidity: event RateDataUpdate(address indexed reserve, uint256 optimalUsageRatio, uint256 baseVariableBorrowRate, uint256 variableRateSlope1, uint256 variableRateSlope2)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Filterer) FilterRateDataUpdate(opts *bind.FilterOpts, reserve []common.Address) (*AaveDefaultReserveInterestRateStrategyV2RateDataUpdateIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AaveDefaultReserveInterestRateStrategyV2.contract.FilterLogs(opts, "RateDataUpdate", reserveRule)
	if err != nil {
		return nil, err
	}
	return &AaveDefaultReserveInterestRateStrategyV2RateDataUpdateIterator{contract: _AaveDefaultReserveInterestRateStrategyV2.contract, event: "RateDataUpdate", logs: logs, sub: sub}, nil
}

// WatchRateDataUpdate is a free log subscription operation binding the contract event 0x5d123bea2036a4052274206f59d99350b9741e17da56ffae335d809b25ee0942.
//
// Solidity: event RateDataUpdate(address indexed reserve, uint256 optimalUsageRatio, uint256 baseVariableBorrowRate, uint256 variableRateSlope1, uint256 variableRateSlope2)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Filterer) WatchRateDataUpdate(opts *bind.WatchOpts, sink chan<- *AaveDefaultReserveInterestRateStrategyV2RateDataUpdate, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AaveDefaultReserveInterestRateStrategyV2.contract.WatchLogs(opts, "RateDataUpdate", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveDefaultReserveInterestRateStrategyV2RateDataUpdate)
				if err := _AaveDefaultReserveInterestRateStrategyV2.contract.UnpackLog(event, "RateDataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRateDataUpdate is a log parse operation binding the contract event 0x5d123bea2036a4052274206f59d99350b9741e17da56ffae335d809b25ee0942.
//
// Solidity: event RateDataUpdate(address indexed reserve, uint256 optimalUsageRatio, uint256 baseVariableBorrowRate, uint256 variableRateSlope1, uint256 variableRateSlope2)
func (_AaveDefaultReserveInterestRateStrategyV2 *AaveDefaultReserveInterestRateStrategyV2Filterer) ParseRateDataUpdate(log types.Log) (*AaveDefaultReserveInterestRateStrategyV2RateDataUpdate, error) {
	event := new(AaveDefaultReserveInterestRateStrategyV2RateDataUpdate)
	if err := _AaveDefaultReserveInterestRateStrategyV2.contract.UnpackLog(event, "RateDataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
