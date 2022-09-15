// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderjoeFactory

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

// TraderjoeFactoryMetaData contains all meta data concerning the TraderjoeFactory contract.
var TraderjoeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TraderjoeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderjoeFactoryMetaData.ABI instead.
var TraderjoeFactoryABI = TraderjoeFactoryMetaData.ABI

// TraderjoeFactory is an auto generated Go binding around an Ethereum contract.
type TraderjoeFactory struct {
	TraderjoeFactoryCaller     // Read-only binding to the contract
	TraderjoeFactoryTransactor // Write-only binding to the contract
	TraderjoeFactoryFilterer   // Log filterer for contract events
}

// TraderjoeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderjoeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderjoeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderjoeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderjoeFactorySession struct {
	Contract     *TraderjoeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraderjoeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderjoeFactoryCallerSession struct {
	Contract *TraderjoeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TraderjoeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderjoeFactoryTransactorSession struct {
	Contract     *TraderjoeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TraderjoeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderjoeFactoryRaw struct {
	Contract *TraderjoeFactory // Generic contract binding to access the raw methods on
}

// TraderjoeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderjoeFactoryCallerRaw struct {
	Contract *TraderjoeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TraderjoeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderjoeFactoryTransactorRaw struct {
	Contract *TraderjoeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderjoeFactory creates a new instance of TraderjoeFactory, bound to a specific deployed contract.
func NewTraderjoeFactory(address common.Address, backend bind.ContractBackend) (*TraderjoeFactory, error) {
	contract, err := bindTraderjoeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraderjoeFactory{TraderjoeFactoryCaller: TraderjoeFactoryCaller{contract: contract}, TraderjoeFactoryTransactor: TraderjoeFactoryTransactor{contract: contract}, TraderjoeFactoryFilterer: TraderjoeFactoryFilterer{contract: contract}}, nil
}

// NewTraderjoeFactoryCaller creates a new read-only instance of TraderjoeFactory, bound to a specific deployed contract.
func NewTraderjoeFactoryCaller(address common.Address, caller bind.ContractCaller) (*TraderjoeFactoryCaller, error) {
	contract, err := bindTraderjoeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeFactoryCaller{contract: contract}, nil
}

// NewTraderjoeFactoryTransactor creates a new write-only instance of TraderjoeFactory, bound to a specific deployed contract.
func NewTraderjoeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderjoeFactoryTransactor, error) {
	contract, err := bindTraderjoeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeFactoryTransactor{contract: contract}, nil
}

// NewTraderjoeFactoryFilterer creates a new log filterer instance of TraderjoeFactory, bound to a specific deployed contract.
func NewTraderjoeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderjoeFactoryFilterer, error) {
	contract, err := bindTraderjoeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderjoeFactoryFilterer{contract: contract}, nil
}

// bindTraderjoeFactory binds a generic wrapper to an already deployed contract.
func bindTraderjoeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraderjoeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeFactory *TraderjoeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeFactory.Contract.TraderjoeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeFactory *TraderjoeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.TraderjoeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeFactory *TraderjoeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.TraderjoeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeFactory *TraderjoeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeFactory *TraderjoeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeFactory *TraderjoeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_TraderjoeFactory *TraderjoeFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeFactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_TraderjoeFactory *TraderjoeFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _TraderjoeFactory.Contract.AllPairs(&_TraderjoeFactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_TraderjoeFactory *TraderjoeFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _TraderjoeFactory.Contract.AllPairs(&_TraderjoeFactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_TraderjoeFactory *TraderjoeFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeFactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_TraderjoeFactory *TraderjoeFactorySession) AllPairsLength() (*big.Int, error) {
	return _TraderjoeFactory.Contract.AllPairsLength(&_TraderjoeFactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_TraderjoeFactory *TraderjoeFactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _TraderjoeFactory.Contract.AllPairsLength(&_TraderjoeFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_TraderjoeFactory *TraderjoeFactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeFactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_TraderjoeFactory *TraderjoeFactorySession) FeeTo() (common.Address, error) {
	return _TraderjoeFactory.Contract.FeeTo(&_TraderjoeFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_TraderjoeFactory *TraderjoeFactoryCallerSession) FeeTo() (common.Address, error) {
	return _TraderjoeFactory.Contract.FeeTo(&_TraderjoeFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_TraderjoeFactory *TraderjoeFactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeFactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_TraderjoeFactory *TraderjoeFactorySession) FeeToSetter() (common.Address, error) {
	return _TraderjoeFactory.Contract.FeeToSetter(&_TraderjoeFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_TraderjoeFactory *TraderjoeFactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _TraderjoeFactory.Contract.FeeToSetter(&_TraderjoeFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_TraderjoeFactory *TraderjoeFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_TraderjoeFactory *TraderjoeFactorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _TraderjoeFactory.Contract.GetPair(&_TraderjoeFactory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_TraderjoeFactory *TraderjoeFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _TraderjoeFactory.Contract.GetPair(&_TraderjoeFactory.CallOpts, tokenA, tokenB)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_TraderjoeFactory *TraderjoeFactoryCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeFactory.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_TraderjoeFactory *TraderjoeFactorySession) Migrator() (common.Address, error) {
	return _TraderjoeFactory.Contract.Migrator(&_TraderjoeFactory.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_TraderjoeFactory *TraderjoeFactoryCallerSession) Migrator() (common.Address, error) {
	return _TraderjoeFactory.Contract.Migrator(&_TraderjoeFactory.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_TraderjoeFactory *TraderjoeFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_TraderjoeFactory *TraderjoeFactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.CreatePair(&_TraderjoeFactory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_TraderjoeFactory *TraderjoeFactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.CreatePair(&_TraderjoeFactory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_TraderjoeFactory *TraderjoeFactoryTransactor) SetFeeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.contract.Transact(opts, "setFeeTo", arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_TraderjoeFactory *TraderjoeFactorySession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.SetFeeTo(&_TraderjoeFactory.TransactOpts, arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_TraderjoeFactory *TraderjoeFactoryTransactorSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.SetFeeTo(&_TraderjoeFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_TraderjoeFactory *TraderjoeFactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.contract.Transact(opts, "setFeeToSetter", arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_TraderjoeFactory *TraderjoeFactorySession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.SetFeeToSetter(&_TraderjoeFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_TraderjoeFactory *TraderjoeFactoryTransactorSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.SetFeeToSetter(&_TraderjoeFactory.TransactOpts, arg0)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address ) returns()
func (_TraderjoeFactory *TraderjoeFactoryTransactor) SetMigrator(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.contract.Transact(opts, "setMigrator", arg0)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address ) returns()
func (_TraderjoeFactory *TraderjoeFactorySession) SetMigrator(arg0 common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.SetMigrator(&_TraderjoeFactory.TransactOpts, arg0)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address ) returns()
func (_TraderjoeFactory *TraderjoeFactoryTransactorSession) SetMigrator(arg0 common.Address) (*types.Transaction, error) {
	return _TraderjoeFactory.Contract.SetMigrator(&_TraderjoeFactory.TransactOpts, arg0)
}

// TraderjoeFactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the TraderjoeFactory contract.
type TraderjoeFactoryPairCreatedIterator struct {
	Event *TraderjoeFactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *TraderjoeFactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoeFactoryPairCreated)
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
		it.Event = new(TraderjoeFactoryPairCreated)
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
func (it *TraderjoeFactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoeFactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoeFactoryPairCreated represents a PairCreated event raised by the TraderjoeFactory contract.
type TraderjoeFactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_TraderjoeFactory *TraderjoeFactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*TraderjoeFactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _TraderjoeFactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &TraderjoeFactoryPairCreatedIterator{contract: _TraderjoeFactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_TraderjoeFactory *TraderjoeFactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *TraderjoeFactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _TraderjoeFactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoeFactoryPairCreated)
				if err := _TraderjoeFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_TraderjoeFactory *TraderjoeFactoryFilterer) ParsePairCreated(log types.Log) (*TraderjoeFactoryPairCreated, error) {
	event := new(TraderjoeFactoryPairCreated)
	if err := _TraderjoeFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
