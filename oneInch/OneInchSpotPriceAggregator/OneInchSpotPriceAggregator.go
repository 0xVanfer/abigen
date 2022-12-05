// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OneInchSpotPriceAggregator

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

// OneInchSpotPriceAggregatorMetaData contains all meta data concerning the OneInchSpotPriceAggregator contract.
var OneInchSpotPriceAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"addConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"oracleKind\",\"type\":\"uint8\"}],\"name\":\"addOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"allConnectors\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useWrappers\",\"type\":\"bool\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useSrcWrappers\",\"type\":\"bool\"}],\"name\":\"getRateToEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiWrapper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"allOracles\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"oracleTypes\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"removeConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"oracleKind\",\"type\":\"uint8\"}],\"name\":\"removeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_multiWrapper\",\"type\":\"address\"}],\"name\":\"setMultiWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OneInchSpotPriceAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use OneInchSpotPriceAggregatorMetaData.ABI instead.
var OneInchSpotPriceAggregatorABI = OneInchSpotPriceAggregatorMetaData.ABI

// OneInchSpotPriceAggregator is an auto generated Go binding around an Ethereum contract.
type OneInchSpotPriceAggregator struct {
	OneInchSpotPriceAggregatorCaller     // Read-only binding to the contract
	OneInchSpotPriceAggregatorTransactor // Write-only binding to the contract
	OneInchSpotPriceAggregatorFilterer   // Log filterer for contract events
}

// OneInchSpotPriceAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneInchSpotPriceAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchSpotPriceAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneInchSpotPriceAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchSpotPriceAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneInchSpotPriceAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchSpotPriceAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneInchSpotPriceAggregatorSession struct {
	Contract     *OneInchSpotPriceAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OneInchSpotPriceAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneInchSpotPriceAggregatorCallerSession struct {
	Contract *OneInchSpotPriceAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// OneInchSpotPriceAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneInchSpotPriceAggregatorTransactorSession struct {
	Contract     *OneInchSpotPriceAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// OneInchSpotPriceAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneInchSpotPriceAggregatorRaw struct {
	Contract *OneInchSpotPriceAggregator // Generic contract binding to access the raw methods on
}

// OneInchSpotPriceAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneInchSpotPriceAggregatorCallerRaw struct {
	Contract *OneInchSpotPriceAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// OneInchSpotPriceAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneInchSpotPriceAggregatorTransactorRaw struct {
	Contract *OneInchSpotPriceAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneInchSpotPriceAggregator creates a new instance of OneInchSpotPriceAggregator, bound to a specific deployed contract.
func NewOneInchSpotPriceAggregator(address common.Address, backend bind.ContractBackend) (*OneInchSpotPriceAggregator, error) {
	contract, err := bindOneInchSpotPriceAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceAggregator{OneInchSpotPriceAggregatorCaller: OneInchSpotPriceAggregatorCaller{contract: contract}, OneInchSpotPriceAggregatorTransactor: OneInchSpotPriceAggregatorTransactor{contract: contract}, OneInchSpotPriceAggregatorFilterer: OneInchSpotPriceAggregatorFilterer{contract: contract}}, nil
}

// NewOneInchSpotPriceAggregatorCaller creates a new read-only instance of OneInchSpotPriceAggregator, bound to a specific deployed contract.
func NewOneInchSpotPriceAggregatorCaller(address common.Address, caller bind.ContractCaller) (*OneInchSpotPriceAggregatorCaller, error) {
	contract, err := bindOneInchSpotPriceAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceAggregatorCaller{contract: contract}, nil
}

// NewOneInchSpotPriceAggregatorTransactor creates a new write-only instance of OneInchSpotPriceAggregator, bound to a specific deployed contract.
func NewOneInchSpotPriceAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OneInchSpotPriceAggregatorTransactor, error) {
	contract, err := bindOneInchSpotPriceAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceAggregatorTransactor{contract: contract}, nil
}

// NewOneInchSpotPriceAggregatorFilterer creates a new log filterer instance of OneInchSpotPriceAggregator, bound to a specific deployed contract.
func NewOneInchSpotPriceAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OneInchSpotPriceAggregatorFilterer, error) {
	contract, err := bindOneInchSpotPriceAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceAggregatorFilterer{contract: contract}, nil
}

// bindOneInchSpotPriceAggregator binds a generic wrapper to an already deployed contract.
func bindOneInchSpotPriceAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneInchSpotPriceAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchSpotPriceAggregator.Contract.OneInchSpotPriceAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.OneInchSpotPriceAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.OneInchSpotPriceAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchSpotPriceAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.contract.Transact(opts, method, params...)
}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCaller) Connectors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OneInchSpotPriceAggregator.contract.Call(opts, &out, "connectors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) Connectors() ([]common.Address, error) {
	return _OneInchSpotPriceAggregator.Contract.Connectors(&_OneInchSpotPriceAggregator.CallOpts)
}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCallerSession) Connectors() ([]common.Address, error) {
	return _OneInchSpotPriceAggregator.Contract.Connectors(&_OneInchSpotPriceAggregator.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCaller) GetRate(opts *bind.CallOpts, srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	var out []interface{}
	err := _OneInchSpotPriceAggregator.contract.Call(opts, &out, "getRate", srcToken, dstToken, useWrappers)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) GetRate(srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	return _OneInchSpotPriceAggregator.Contract.GetRate(&_OneInchSpotPriceAggregator.CallOpts, srcToken, dstToken, useWrappers)
}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCallerSession) GetRate(srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	return _OneInchSpotPriceAggregator.Contract.GetRate(&_OneInchSpotPriceAggregator.CallOpts, srcToken, dstToken, useWrappers)
}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCaller) GetRateToEth(opts *bind.CallOpts, srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	var out []interface{}
	err := _OneInchSpotPriceAggregator.contract.Call(opts, &out, "getRateToEth", srcToken, useSrcWrappers)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) GetRateToEth(srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	return _OneInchSpotPriceAggregator.Contract.GetRateToEth(&_OneInchSpotPriceAggregator.CallOpts, srcToken, useSrcWrappers)
}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCallerSession) GetRateToEth(srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	return _OneInchSpotPriceAggregator.Contract.GetRateToEth(&_OneInchSpotPriceAggregator.CallOpts, srcToken, useSrcWrappers)
}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCaller) MultiWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchSpotPriceAggregator.contract.Call(opts, &out, "multiWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) MultiWrapper() (common.Address, error) {
	return _OneInchSpotPriceAggregator.Contract.MultiWrapper(&_OneInchSpotPriceAggregator.CallOpts)
}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCallerSession) MultiWrapper() (common.Address, error) {
	return _OneInchSpotPriceAggregator.Contract.MultiWrapper(&_OneInchSpotPriceAggregator.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCaller) Oracles(opts *bind.CallOpts) (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	var out []interface{}
	err := _OneInchSpotPriceAggregator.contract.Call(opts, &out, "oracles")

	outstruct := new(struct {
		AllOracles  []common.Address
		OracleTypes []uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllOracles = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OracleTypes = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) Oracles() (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	return _OneInchSpotPriceAggregator.Contract.Oracles(&_OneInchSpotPriceAggregator.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCallerSession) Oracles() (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	return _OneInchSpotPriceAggregator.Contract.Oracles(&_OneInchSpotPriceAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchSpotPriceAggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) Owner() (common.Address, error) {
	return _OneInchSpotPriceAggregator.Contract.Owner(&_OneInchSpotPriceAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorCallerSession) Owner() (common.Address, error) {
	return _OneInchSpotPriceAggregator.Contract.Owner(&_OneInchSpotPriceAggregator.CallOpts)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactor) AddConnector(opts *bind.TransactOpts, connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.contract.Transact(opts, "addConnector", connector)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) AddConnector(connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.AddConnector(&_OneInchSpotPriceAggregator.TransactOpts, connector)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactorSession) AddConnector(connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.AddConnector(&_OneInchSpotPriceAggregator.TransactOpts, connector)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactor) AddOracle(opts *bind.TransactOpts, oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.contract.Transact(opts, "addOracle", oracle, oracleKind)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) AddOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.AddOracle(&_OneInchSpotPriceAggregator.TransactOpts, oracle, oracleKind)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactorSession) AddOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.AddOracle(&_OneInchSpotPriceAggregator.TransactOpts, oracle, oracleKind)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactor) RemoveConnector(opts *bind.TransactOpts, connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.contract.Transact(opts, "removeConnector", connector)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) RemoveConnector(connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.RemoveConnector(&_OneInchSpotPriceAggregator.TransactOpts, connector)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactorSession) RemoveConnector(connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.RemoveConnector(&_OneInchSpotPriceAggregator.TransactOpts, connector)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactor) RemoveOracle(opts *bind.TransactOpts, oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.contract.Transact(opts, "removeOracle", oracle, oracleKind)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) RemoveOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.RemoveOracle(&_OneInchSpotPriceAggregator.TransactOpts, oracle, oracleKind)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactorSession) RemoveOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.RemoveOracle(&_OneInchSpotPriceAggregator.TransactOpts, oracle, oracleKind)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.RenounceOwnership(&_OneInchSpotPriceAggregator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.RenounceOwnership(&_OneInchSpotPriceAggregator.TransactOpts)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactor) SetMultiWrapper(opts *bind.TransactOpts, _multiWrapper common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.contract.Transact(opts, "setMultiWrapper", _multiWrapper)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) SetMultiWrapper(_multiWrapper common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.SetMultiWrapper(&_OneInchSpotPriceAggregator.TransactOpts, _multiWrapper)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactorSession) SetMultiWrapper(_multiWrapper common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.SetMultiWrapper(&_OneInchSpotPriceAggregator.TransactOpts, _multiWrapper)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.TransferOwnership(&_OneInchSpotPriceAggregator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchSpotPriceAggregator *OneInchSpotPriceAggregatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchSpotPriceAggregator.Contract.TransferOwnership(&_OneInchSpotPriceAggregator.TransactOpts, newOwner)
}
