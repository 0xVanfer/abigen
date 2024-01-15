// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staplePathFinder

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

// StaplePathFinderMetaData contains all meta data concerning the StaplePathFinder contract.
var StaplePathFinderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxLayers\",\"type\":\"uint256\"}],\"name\":\"findAllPaths\",\"outputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"paths\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[][]\",\"name\":\"allPaths\",\"type\":\"uint256[][]\"}],\"name\":\"findBestPath\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"estiOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StaplePathFinderABI is the input ABI used to generate the binding from.
// Deprecated: Use StaplePathFinderMetaData.ABI instead.
var StaplePathFinderABI = StaplePathFinderMetaData.ABI

// StaplePathFinder is an auto generated Go binding around an Ethereum contract.
type StaplePathFinder struct {
	StaplePathFinderCaller     // Read-only binding to the contract
	StaplePathFinderTransactor // Write-only binding to the contract
	StaplePathFinderFilterer   // Log filterer for contract events
}

// StaplePathFinderCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaplePathFinderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaplePathFinderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaplePathFinderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaplePathFinderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaplePathFinderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaplePathFinderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaplePathFinderSession struct {
	Contract     *StaplePathFinder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaplePathFinderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaplePathFinderCallerSession struct {
	Contract *StaplePathFinderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StaplePathFinderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaplePathFinderTransactorSession struct {
	Contract     *StaplePathFinderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StaplePathFinderRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaplePathFinderRaw struct {
	Contract *StaplePathFinder // Generic contract binding to access the raw methods on
}

// StaplePathFinderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaplePathFinderCallerRaw struct {
	Contract *StaplePathFinderCaller // Generic read-only contract binding to access the raw methods on
}

// StaplePathFinderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaplePathFinderTransactorRaw struct {
	Contract *StaplePathFinderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaplePathFinder creates a new instance of StaplePathFinder, bound to a specific deployed contract.
func NewStaplePathFinder(address common.Address, backend bind.ContractBackend) (*StaplePathFinder, error) {
	contract, err := bindStaplePathFinder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaplePathFinder{StaplePathFinderCaller: StaplePathFinderCaller{contract: contract}, StaplePathFinderTransactor: StaplePathFinderTransactor{contract: contract}, StaplePathFinderFilterer: StaplePathFinderFilterer{contract: contract}}, nil
}

// NewStaplePathFinderCaller creates a new read-only instance of StaplePathFinder, bound to a specific deployed contract.
func NewStaplePathFinderCaller(address common.Address, caller bind.ContractCaller) (*StaplePathFinderCaller, error) {
	contract, err := bindStaplePathFinder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaplePathFinderCaller{contract: contract}, nil
}

// NewStaplePathFinderTransactor creates a new write-only instance of StaplePathFinder, bound to a specific deployed contract.
func NewStaplePathFinderTransactor(address common.Address, transactor bind.ContractTransactor) (*StaplePathFinderTransactor, error) {
	contract, err := bindStaplePathFinder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaplePathFinderTransactor{contract: contract}, nil
}

// NewStaplePathFinderFilterer creates a new log filterer instance of StaplePathFinder, bound to a specific deployed contract.
func NewStaplePathFinderFilterer(address common.Address, filterer bind.ContractFilterer) (*StaplePathFinderFilterer, error) {
	contract, err := bindStaplePathFinder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaplePathFinderFilterer{contract: contract}, nil
}

// bindStaplePathFinder binds a generic wrapper to an already deployed contract.
func bindStaplePathFinder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StaplePathFinderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaplePathFinder *StaplePathFinderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaplePathFinder.Contract.StaplePathFinderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaplePathFinder *StaplePathFinderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaplePathFinder.Contract.StaplePathFinderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaplePathFinder *StaplePathFinderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaplePathFinder.Contract.StaplePathFinderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaplePathFinder *StaplePathFinderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaplePathFinder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaplePathFinder *StaplePathFinderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaplePathFinder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaplePathFinder *StaplePathFinderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaplePathFinder.Contract.contract.Transact(opts, method, params...)
}

// FindAllPaths is a free data retrieval call binding the contract method 0xc274873f.
//
// Solidity: function findAllPaths(address tokenIn, address tokenOut, uint256 maxLayers) view returns(uint256[][] paths)
func (_StaplePathFinder *StaplePathFinderCaller) FindAllPaths(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address, maxLayers *big.Int) ([][]*big.Int, error) {
	var out []interface{}
	err := _StaplePathFinder.contract.Call(opts, &out, "findAllPaths", tokenIn, tokenOut, maxLayers)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// FindAllPaths is a free data retrieval call binding the contract method 0xc274873f.
//
// Solidity: function findAllPaths(address tokenIn, address tokenOut, uint256 maxLayers) view returns(uint256[][] paths)
func (_StaplePathFinder *StaplePathFinderSession) FindAllPaths(tokenIn common.Address, tokenOut common.Address, maxLayers *big.Int) ([][]*big.Int, error) {
	return _StaplePathFinder.Contract.FindAllPaths(&_StaplePathFinder.CallOpts, tokenIn, tokenOut, maxLayers)
}

// FindAllPaths is a free data retrieval call binding the contract method 0xc274873f.
//
// Solidity: function findAllPaths(address tokenIn, address tokenOut, uint256 maxLayers) view returns(uint256[][] paths)
func (_StaplePathFinder *StaplePathFinderCallerSession) FindAllPaths(tokenIn common.Address, tokenOut common.Address, maxLayers *big.Int) ([][]*big.Int, error) {
	return _StaplePathFinder.Contract.FindAllPaths(&_StaplePathFinder.CallOpts, tokenIn, tokenOut, maxLayers)
}

// FindBestPath is a free data retrieval call binding the contract method 0x18e40959.
//
// Solidity: function findBestPath(address user, address tokenIn, uint256 amount, uint256[][] allPaths) view returns(uint256 estiOut, uint256[] path)
func (_StaplePathFinder *StaplePathFinderCaller) FindBestPath(opts *bind.CallOpts, user common.Address, tokenIn common.Address, amount *big.Int, allPaths [][]*big.Int) (struct {
	EstiOut *big.Int
	Path    []*big.Int
}, error) {
	var out []interface{}
	err := _StaplePathFinder.contract.Call(opts, &out, "findBestPath", user, tokenIn, amount, allPaths)

	outstruct := new(struct {
		EstiOut *big.Int
		Path    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EstiOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Path = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// FindBestPath is a free data retrieval call binding the contract method 0x18e40959.
//
// Solidity: function findBestPath(address user, address tokenIn, uint256 amount, uint256[][] allPaths) view returns(uint256 estiOut, uint256[] path)
func (_StaplePathFinder *StaplePathFinderSession) FindBestPath(user common.Address, tokenIn common.Address, amount *big.Int, allPaths [][]*big.Int) (struct {
	EstiOut *big.Int
	Path    []*big.Int
}, error) {
	return _StaplePathFinder.Contract.FindBestPath(&_StaplePathFinder.CallOpts, user, tokenIn, amount, allPaths)
}

// FindBestPath is a free data retrieval call binding the contract method 0x18e40959.
//
// Solidity: function findBestPath(address user, address tokenIn, uint256 amount, uint256[][] allPaths) view returns(uint256 estiOut, uint256[] path)
func (_StaplePathFinder *StaplePathFinderCallerSession) FindBestPath(user common.Address, tokenIn common.Address, amount *big.Int, allPaths [][]*big.Int) (struct {
	EstiOut *big.Int
	Path    []*big.Int
}, error) {
	return _StaplePathFinder.Contract.FindBestPath(&_StaplePathFinder.CallOpts, user, tokenIn, amount, allPaths)
}
