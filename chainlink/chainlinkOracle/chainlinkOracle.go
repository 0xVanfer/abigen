// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainlinkOracle

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

// ChainlinkOracleMetaData contains all meta data concerning the ChainlinkOracle contract.
var ChainlinkOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"confirmAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"phaseAggregators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"phaseId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"proposeAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"proposedGetRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedLatestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accessController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ChainlinkOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use ChainlinkOracleMetaData.ABI instead.
var ChainlinkOracleABI = ChainlinkOracleMetaData.ABI

// ChainlinkOracle is an auto generated Go binding around an Ethereum contract.
type ChainlinkOracle struct {
	ChainlinkOracleCaller     // Read-only binding to the contract
	ChainlinkOracleTransactor // Write-only binding to the contract
	ChainlinkOracleFilterer   // Log filterer for contract events
}

// ChainlinkOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainlinkOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainlinkOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainlinkOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainlinkOracleSession struct {
	Contract     *ChainlinkOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainlinkOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainlinkOracleCallerSession struct {
	Contract *ChainlinkOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ChainlinkOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainlinkOracleTransactorSession struct {
	Contract     *ChainlinkOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ChainlinkOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainlinkOracleRaw struct {
	Contract *ChainlinkOracle // Generic contract binding to access the raw methods on
}

// ChainlinkOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainlinkOracleCallerRaw struct {
	Contract *ChainlinkOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ChainlinkOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainlinkOracleTransactorRaw struct {
	Contract *ChainlinkOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkOracle creates a new instance of ChainlinkOracle, bound to a specific deployed contract.
func NewChainlinkOracle(address common.Address, backend bind.ContractBackend) (*ChainlinkOracle, error) {
	contract, err := bindChainlinkOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracle{ChainlinkOracleCaller: ChainlinkOracleCaller{contract: contract}, ChainlinkOracleTransactor: ChainlinkOracleTransactor{contract: contract}, ChainlinkOracleFilterer: ChainlinkOracleFilterer{contract: contract}}, nil
}

// NewChainlinkOracleCaller creates a new read-only instance of ChainlinkOracle, bound to a specific deployed contract.
func NewChainlinkOracleCaller(address common.Address, caller bind.ContractCaller) (*ChainlinkOracleCaller, error) {
	contract, err := bindChainlinkOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleCaller{contract: contract}, nil
}

// NewChainlinkOracleTransactor creates a new write-only instance of ChainlinkOracle, bound to a specific deployed contract.
func NewChainlinkOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainlinkOracleTransactor, error) {
	contract, err := bindChainlinkOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleTransactor{contract: contract}, nil
}

// NewChainlinkOracleFilterer creates a new log filterer instance of ChainlinkOracle, bound to a specific deployed contract.
func NewChainlinkOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainlinkOracleFilterer, error) {
	contract, err := bindChainlinkOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleFilterer{contract: contract}, nil
}

// bindChainlinkOracle binds a generic wrapper to an already deployed contract.
func bindChainlinkOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainlinkOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkOracle *ChainlinkOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkOracle.Contract.ChainlinkOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkOracle *ChainlinkOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.ChainlinkOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkOracle *ChainlinkOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.ChainlinkOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkOracle *ChainlinkOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkOracle *ChainlinkOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkOracle *ChainlinkOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.contract.Transact(opts, method, params...)
}

// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
//
// Solidity: function accessController() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCaller) AccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "accessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
//
// Solidity: function accessController() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleSession) AccessController() (common.Address, error) {
	return _ChainlinkOracle.Contract.AccessController(&_ChainlinkOracle.CallOpts)
}

// AccessController is a free data retrieval call binding the contract method 0xbc43cbaf.
//
// Solidity: function accessController() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) AccessController() (common.Address, error) {
	return _ChainlinkOracle.Contract.AccessController(&_ChainlinkOracle.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleSession) Aggregator() (common.Address, error) {
	return _ChainlinkOracle.Contract.Aggregator(&_ChainlinkOracle.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) Aggregator() (common.Address, error) {
	return _ChainlinkOracle.Contract.Aggregator(&_ChainlinkOracle.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkOracle *ChainlinkOracleCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkOracle *ChainlinkOracleSession) Decimals() (uint8, error) {
	return _ChainlinkOracle.Contract.Decimals(&_ChainlinkOracle.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) Decimals() (uint8, error) {
	return _ChainlinkOracle.Contract.Decimals(&_ChainlinkOracle.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ChainlinkOracle *ChainlinkOracleCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ChainlinkOracle *ChainlinkOracleSession) Description() (string, error) {
	return _ChainlinkOracle.Contract.Description(&_ChainlinkOracle.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) Description() (string, error) {
	return _ChainlinkOracle.Contract.Description(&_ChainlinkOracle.CallOpts)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "getAnswer", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetAnswer(&_ChainlinkOracle.CallOpts, _roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetAnswer(&_ChainlinkOracle.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkOracle.Contract.GetRoundData(&_ChainlinkOracle.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkOracle.Contract.GetRoundData(&_ChainlinkOracle.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "getTimestamp", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetTimestamp(&_ChainlinkOracle.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetTimestamp(&_ChainlinkOracle.CallOpts, _roundId)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleSession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestAnswer(&_ChainlinkOracle.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestAnswer(&_ChainlinkOracle.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleSession) LatestRound() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestRound(&_ChainlinkOracle.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) LatestRound() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestRound(&_ChainlinkOracle.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkOracle.Contract.LatestRoundData(&_ChainlinkOracle.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkOracle.Contract.LatestRoundData(&_ChainlinkOracle.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleSession) LatestTimestamp() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestTimestamp(&_ChainlinkOracle.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) LatestTimestamp() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestTimestamp(&_ChainlinkOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleSession) Owner() (common.Address, error) {
	return _ChainlinkOracle.Contract.Owner(&_ChainlinkOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) Owner() (common.Address, error) {
	return _ChainlinkOracle.Contract.Owner(&_ChainlinkOracle.CallOpts)
}

// PhaseAggregators is a free data retrieval call binding the contract method 0xc1597304.
//
// Solidity: function phaseAggregators(uint16 ) view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCaller) PhaseAggregators(opts *bind.CallOpts, arg0 uint16) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "phaseAggregators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PhaseAggregators is a free data retrieval call binding the contract method 0xc1597304.
//
// Solidity: function phaseAggregators(uint16 ) view returns(address)
func (_ChainlinkOracle *ChainlinkOracleSession) PhaseAggregators(arg0 uint16) (common.Address, error) {
	return _ChainlinkOracle.Contract.PhaseAggregators(&_ChainlinkOracle.CallOpts, arg0)
}

// PhaseAggregators is a free data retrieval call binding the contract method 0xc1597304.
//
// Solidity: function phaseAggregators(uint16 ) view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) PhaseAggregators(arg0 uint16) (common.Address, error) {
	return _ChainlinkOracle.Contract.PhaseAggregators(&_ChainlinkOracle.CallOpts, arg0)
}

// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
//
// Solidity: function phaseId() view returns(uint16)
func (_ChainlinkOracle *ChainlinkOracleCaller) PhaseId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "phaseId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
//
// Solidity: function phaseId() view returns(uint16)
func (_ChainlinkOracle *ChainlinkOracleSession) PhaseId() (uint16, error) {
	return _ChainlinkOracle.Contract.PhaseId(&_ChainlinkOracle.CallOpts)
}

// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
//
// Solidity: function phaseId() view returns(uint16)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) PhaseId() (uint16, error) {
	return _ChainlinkOracle.Contract.PhaseId(&_ChainlinkOracle.CallOpts)
}

// ProposedAggregator is a free data retrieval call binding the contract method 0xe8c4be30.
//
// Solidity: function proposedAggregator() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCaller) ProposedAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "proposedAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedAggregator is a free data retrieval call binding the contract method 0xe8c4be30.
//
// Solidity: function proposedAggregator() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleSession) ProposedAggregator() (common.Address, error) {
	return _ChainlinkOracle.Contract.ProposedAggregator(&_ChainlinkOracle.CallOpts)
}

// ProposedAggregator is a free data retrieval call binding the contract method 0xe8c4be30.
//
// Solidity: function proposedAggregator() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) ProposedAggregator() (common.Address, error) {
	return _ChainlinkOracle.Contract.ProposedAggregator(&_ChainlinkOracle.CallOpts)
}

// ProposedGetRoundData is a free data retrieval call binding the contract method 0x6001ac53.
//
// Solidity: function proposedGetRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleCaller) ProposedGetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "proposedGetRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposedGetRoundData is a free data retrieval call binding the contract method 0x6001ac53.
//
// Solidity: function proposedGetRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleSession) ProposedGetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkOracle.Contract.ProposedGetRoundData(&_ChainlinkOracle.CallOpts, _roundId)
}

// ProposedGetRoundData is a free data retrieval call binding the contract method 0x6001ac53.
//
// Solidity: function proposedGetRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) ProposedGetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkOracle.Contract.ProposedGetRoundData(&_ChainlinkOracle.CallOpts, _roundId)
}

// ProposedLatestRoundData is a free data retrieval call binding the contract method 0x8f6b4d91.
//
// Solidity: function proposedLatestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleCaller) ProposedLatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "proposedLatestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposedLatestRoundData is a free data retrieval call binding the contract method 0x8f6b4d91.
//
// Solidity: function proposedLatestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleSession) ProposedLatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkOracle.Contract.ProposedLatestRoundData(&_ChainlinkOracle.CallOpts)
}

// ProposedLatestRoundData is a free data retrieval call binding the contract method 0x8f6b4d91.
//
// Solidity: function proposedLatestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) ProposedLatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkOracle.Contract.ProposedLatestRoundData(&_ChainlinkOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkOracle.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleSession) Version() (*big.Int, error) {
	return _ChainlinkOracle.Contract.Version(&_ChainlinkOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) Version() (*big.Int, error) {
	return _ChainlinkOracle.Contract.Version(&_ChainlinkOracle.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkOracle *ChainlinkOracleSession) AcceptOwnership() (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.AcceptOwnership(&_ChainlinkOracle.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.AcceptOwnership(&_ChainlinkOracle.TransactOpts)
}

// ConfirmAggregator is a paid mutator transaction binding the contract method 0xa928c096.
//
// Solidity: function confirmAggregator(address _aggregator) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) ConfirmAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "confirmAggregator", _aggregator)
}

// ConfirmAggregator is a paid mutator transaction binding the contract method 0xa928c096.
//
// Solidity: function confirmAggregator(address _aggregator) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) ConfirmAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.ConfirmAggregator(&_ChainlinkOracle.TransactOpts, _aggregator)
}

// ConfirmAggregator is a paid mutator transaction binding the contract method 0xa928c096.
//
// Solidity: function confirmAggregator(address _aggregator) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) ConfirmAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.ConfirmAggregator(&_ChainlinkOracle.TransactOpts, _aggregator)
}

// ProposeAggregator is a paid mutator transaction binding the contract method 0xf8a2abd3.
//
// Solidity: function proposeAggregator(address _aggregator) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) ProposeAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "proposeAggregator", _aggregator)
}

// ProposeAggregator is a paid mutator transaction binding the contract method 0xf8a2abd3.
//
// Solidity: function proposeAggregator(address _aggregator) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) ProposeAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.ProposeAggregator(&_ChainlinkOracle.TransactOpts, _aggregator)
}

// ProposeAggregator is a paid mutator transaction binding the contract method 0xf8a2abd3.
//
// Solidity: function proposeAggregator(address _aggregator) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) ProposeAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.ProposeAggregator(&_ChainlinkOracle.TransactOpts, _aggregator)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _accessController) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) SetController(opts *bind.TransactOpts, _accessController common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "setController", _accessController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _accessController) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) SetController(_accessController common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.SetController(&_ChainlinkOracle.TransactOpts, _accessController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _accessController) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) SetController(_accessController common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.SetController(&_ChainlinkOracle.TransactOpts, _accessController)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) TransferOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "transferOwnership", _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.TransferOwnership(&_ChainlinkOracle.TransactOpts, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.TransferOwnership(&_ChainlinkOracle.TransactOpts, _to)
}
