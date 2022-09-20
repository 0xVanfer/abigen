// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package savaxOracle

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

// SavaxOracleMetaData contains all meta data concerning the SavaxOracle contract.
var SavaxOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AVAX_USD_FEED\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SAVAX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SavaxOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use SavaxOracleMetaData.ABI instead.
var SavaxOracleABI = SavaxOracleMetaData.ABI

// SavaxOracle is an auto generated Go binding around an Ethereum contract.
type SavaxOracle struct {
	SavaxOracleCaller     // Read-only binding to the contract
	SavaxOracleTransactor // Write-only binding to the contract
	SavaxOracleFilterer   // Log filterer for contract events
}

// SavaxOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SavaxOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavaxOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SavaxOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavaxOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SavaxOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavaxOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SavaxOracleSession struct {
	Contract     *SavaxOracle      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SavaxOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SavaxOracleCallerSession struct {
	Contract *SavaxOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SavaxOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SavaxOracleTransactorSession struct {
	Contract     *SavaxOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SavaxOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SavaxOracleRaw struct {
	Contract *SavaxOracle // Generic contract binding to access the raw methods on
}

// SavaxOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SavaxOracleCallerRaw struct {
	Contract *SavaxOracleCaller // Generic read-only contract binding to access the raw methods on
}

// SavaxOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SavaxOracleTransactorRaw struct {
	Contract *SavaxOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSavaxOracle creates a new instance of SavaxOracle, bound to a specific deployed contract.
func NewSavaxOracle(address common.Address, backend bind.ContractBackend) (*SavaxOracle, error) {
	contract, err := bindSavaxOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SavaxOracle{SavaxOracleCaller: SavaxOracleCaller{contract: contract}, SavaxOracleTransactor: SavaxOracleTransactor{contract: contract}, SavaxOracleFilterer: SavaxOracleFilterer{contract: contract}}, nil
}

// NewSavaxOracleCaller creates a new read-only instance of SavaxOracle, bound to a specific deployed contract.
func NewSavaxOracleCaller(address common.Address, caller bind.ContractCaller) (*SavaxOracleCaller, error) {
	contract, err := bindSavaxOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SavaxOracleCaller{contract: contract}, nil
}

// NewSavaxOracleTransactor creates a new write-only instance of SavaxOracle, bound to a specific deployed contract.
func NewSavaxOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*SavaxOracleTransactor, error) {
	contract, err := bindSavaxOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SavaxOracleTransactor{contract: contract}, nil
}

// NewSavaxOracleFilterer creates a new log filterer instance of SavaxOracle, bound to a specific deployed contract.
func NewSavaxOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*SavaxOracleFilterer, error) {
	contract, err := bindSavaxOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SavaxOracleFilterer{contract: contract}, nil
}

// bindSavaxOracle binds a generic wrapper to an already deployed contract.
func bindSavaxOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SavaxOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SavaxOracle *SavaxOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SavaxOracle.Contract.SavaxOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SavaxOracle *SavaxOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SavaxOracle.Contract.SavaxOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SavaxOracle *SavaxOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SavaxOracle.Contract.SavaxOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SavaxOracle *SavaxOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SavaxOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SavaxOracle *SavaxOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SavaxOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SavaxOracle *SavaxOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SavaxOracle.Contract.contract.Transact(opts, method, params...)
}

// AVAXUSDFEED is a free data retrieval call binding the contract method 0x4c68bf68.
//
// Solidity: function AVAX_USD_FEED() view returns(address)
func (_SavaxOracle *SavaxOracleCaller) AVAXUSDFEED(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SavaxOracle.contract.Call(opts, &out, "AVAX_USD_FEED")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AVAXUSDFEED is a free data retrieval call binding the contract method 0x4c68bf68.
//
// Solidity: function AVAX_USD_FEED() view returns(address)
func (_SavaxOracle *SavaxOracleSession) AVAXUSDFEED() (common.Address, error) {
	return _SavaxOracle.Contract.AVAXUSDFEED(&_SavaxOracle.CallOpts)
}

// AVAXUSDFEED is a free data retrieval call binding the contract method 0x4c68bf68.
//
// Solidity: function AVAX_USD_FEED() view returns(address)
func (_SavaxOracle *SavaxOracleCallerSession) AVAXUSDFEED() (common.Address, error) {
	return _SavaxOracle.Contract.AVAXUSDFEED(&_SavaxOracle.CallOpts)
}

// SAVAX is a free data retrieval call binding the contract method 0x17a2f6c9.
//
// Solidity: function SAVAX() view returns(address)
func (_SavaxOracle *SavaxOracleCaller) SAVAX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SavaxOracle.contract.Call(opts, &out, "SAVAX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SAVAX is a free data retrieval call binding the contract method 0x17a2f6c9.
//
// Solidity: function SAVAX() view returns(address)
func (_SavaxOracle *SavaxOracleSession) SAVAX() (common.Address, error) {
	return _SavaxOracle.Contract.SAVAX(&_SavaxOracle.CallOpts)
}

// SAVAX is a free data retrieval call binding the contract method 0x17a2f6c9.
//
// Solidity: function SAVAX() view returns(address)
func (_SavaxOracle *SavaxOracleCallerSession) SAVAX() (common.Address, error) {
	return _SavaxOracle.Contract.SAVAX(&_SavaxOracle.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_SavaxOracle *SavaxOracleCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SavaxOracle.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_SavaxOracle *SavaxOracleSession) LatestAnswer() (*big.Int, error) {
	return _SavaxOracle.Contract.LatestAnswer(&_SavaxOracle.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_SavaxOracle *SavaxOracleCallerSession) LatestAnswer() (*big.Int, error) {
	return _SavaxOracle.Contract.LatestAnswer(&_SavaxOracle.CallOpts)
}
