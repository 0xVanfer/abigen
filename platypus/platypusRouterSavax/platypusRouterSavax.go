// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platypusRouterSavax

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

// PlatypusRouterSavaxMetaData contains all meta data concerning the PlatypusRouterSavax contract.
var PlatypusRouterSavaxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"addAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"assetOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getC1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHaircutRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxPriceDeviation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRetentionRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlippageParamK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlippageParamN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getXThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"weth_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantedToken\",\"type\":\"address\"}],\"name\":\"quoteMaxInitialAssetWithdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxInitialAssetAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"}],\"name\":\"quotePotentialSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"potentialOutcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"haircut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"quotePotentialWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enoughCash\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"quotePotentialWithdrawFromOtherAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"}],\"name\":\"removeAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sAvax\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"haircutRate_\",\"type\":\"uint256\"}],\"name\":\"setHaircutRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxPriceDeviation_\",\"type\":\"uint256\"}],\"name\":\"setMaxPriceDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"retentionRatio_\",\"type\":\"uint256\"}],\"name\":\"setRetentionRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sAvax_\",\"type\":\"address\"}],\"name\":\"setSAvax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"k_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"c1_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"xThreshold_\",\"type\":\"uint256\"}],\"name\":\"setSlippageParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wethForwarder_\",\"type\":\"address\"}],\"name\":\"setWETHForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumToAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualToAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"haircut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumToAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapFromETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualToAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"haircut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumToAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapToETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualToAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"haircut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wantedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"withdrawFromOtherAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PlatypusRouterSavaxABI is the input ABI used to generate the binding from.
// Deprecated: Use PlatypusRouterSavaxMetaData.ABI instead.
var PlatypusRouterSavaxABI = PlatypusRouterSavaxMetaData.ABI

// PlatypusRouterSavax is an auto generated Go binding around an Ethereum contract.
type PlatypusRouterSavax struct {
	PlatypusRouterSavaxCaller     // Read-only binding to the contract
	PlatypusRouterSavaxTransactor // Write-only binding to the contract
	PlatypusRouterSavaxFilterer   // Log filterer for contract events
}

// PlatypusRouterSavaxCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlatypusRouterSavaxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusRouterSavaxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlatypusRouterSavaxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusRouterSavaxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlatypusRouterSavaxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusRouterSavaxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlatypusRouterSavaxSession struct {
	Contract     *PlatypusRouterSavax // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PlatypusRouterSavaxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlatypusRouterSavaxCallerSession struct {
	Contract *PlatypusRouterSavaxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PlatypusRouterSavaxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlatypusRouterSavaxTransactorSession struct {
	Contract     *PlatypusRouterSavaxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PlatypusRouterSavaxRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlatypusRouterSavaxRaw struct {
	Contract *PlatypusRouterSavax // Generic contract binding to access the raw methods on
}

// PlatypusRouterSavaxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlatypusRouterSavaxCallerRaw struct {
	Contract *PlatypusRouterSavaxCaller // Generic read-only contract binding to access the raw methods on
}

// PlatypusRouterSavaxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlatypusRouterSavaxTransactorRaw struct {
	Contract *PlatypusRouterSavaxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlatypusRouterSavax creates a new instance of PlatypusRouterSavax, bound to a specific deployed contract.
func NewPlatypusRouterSavax(address common.Address, backend bind.ContractBackend) (*PlatypusRouterSavax, error) {
	contract, err := bindPlatypusRouterSavax(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlatypusRouterSavax{PlatypusRouterSavaxCaller: PlatypusRouterSavaxCaller{contract: contract}, PlatypusRouterSavaxTransactor: PlatypusRouterSavaxTransactor{contract: contract}, PlatypusRouterSavaxFilterer: PlatypusRouterSavaxFilterer{contract: contract}}, nil
}

// NewPlatypusRouterSavaxCaller creates a new read-only instance of PlatypusRouterSavax, bound to a specific deployed contract.
func NewPlatypusRouterSavaxCaller(address common.Address, caller bind.ContractCaller) (*PlatypusRouterSavaxCaller, error) {
	contract, err := bindPlatypusRouterSavax(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlatypusRouterSavaxCaller{contract: contract}, nil
}

// NewPlatypusRouterSavaxTransactor creates a new write-only instance of PlatypusRouterSavax, bound to a specific deployed contract.
func NewPlatypusRouterSavaxTransactor(address common.Address, transactor bind.ContractTransactor) (*PlatypusRouterSavaxTransactor, error) {
	contract, err := bindPlatypusRouterSavax(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlatypusRouterSavaxTransactor{contract: contract}, nil
}

// NewPlatypusRouterSavaxFilterer creates a new log filterer instance of PlatypusRouterSavax, bound to a specific deployed contract.
func NewPlatypusRouterSavaxFilterer(address common.Address, filterer bind.ContractFilterer) (*PlatypusRouterSavaxFilterer, error) {
	contract, err := bindPlatypusRouterSavax(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlatypusRouterSavaxFilterer{contract: contract}, nil
}

// bindPlatypusRouterSavax binds a generic wrapper to an already deployed contract.
func bindPlatypusRouterSavax(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlatypusRouterSavaxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatypusRouterSavax *PlatypusRouterSavaxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatypusRouterSavax.Contract.PlatypusRouterSavaxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatypusRouterSavax *PlatypusRouterSavaxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.PlatypusRouterSavaxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatypusRouterSavax *PlatypusRouterSavaxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.PlatypusRouterSavaxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatypusRouterSavax.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.contract.Transact(opts, method, params...)
}

// AssetOf is a free data retrieval call binding the contract method 0x71f96211.
//
// Solidity: function assetOf(address token) view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) AssetOf(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "assetOf", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetOf is a free data retrieval call binding the contract method 0x71f96211.
//
// Solidity: function assetOf(address token) view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) AssetOf(token common.Address) (common.Address, error) {
	return _PlatypusRouterSavax.Contract.AssetOf(&_PlatypusRouterSavax.CallOpts, token)
}

// AssetOf is a free data retrieval call binding the contract method 0x71f96211.
//
// Solidity: function assetOf(address token) view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) AssetOf(token common.Address) (common.Address, error) {
	return _PlatypusRouterSavax.Contract.AssetOf(&_PlatypusRouterSavax.CallOpts, token)
}

// GetC1 is a free data retrieval call binding the contract method 0xa76f54d2.
//
// Solidity: function getC1() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) GetC1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "getC1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetC1 is a free data retrieval call binding the contract method 0xa76f54d2.
//
// Solidity: function getC1() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) GetC1() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetC1(&_PlatypusRouterSavax.CallOpts)
}

// GetC1 is a free data retrieval call binding the contract method 0xa76f54d2.
//
// Solidity: function getC1() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) GetC1() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetC1(&_PlatypusRouterSavax.CallOpts)
}

// GetDev is a free data retrieval call binding the contract method 0x09bb9267.
//
// Solidity: function getDev() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) GetDev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "getDev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDev is a free data retrieval call binding the contract method 0x09bb9267.
//
// Solidity: function getDev() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) GetDev() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.GetDev(&_PlatypusRouterSavax.CallOpts)
}

// GetDev is a free data retrieval call binding the contract method 0x09bb9267.
//
// Solidity: function getDev() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) GetDev() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.GetDev(&_PlatypusRouterSavax.CallOpts)
}

// GetHaircutRate is a free data retrieval call binding the contract method 0x7fdd5a8e.
//
// Solidity: function getHaircutRate() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) GetHaircutRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "getHaircutRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHaircutRate is a free data retrieval call binding the contract method 0x7fdd5a8e.
//
// Solidity: function getHaircutRate() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) GetHaircutRate() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetHaircutRate(&_PlatypusRouterSavax.CallOpts)
}

// GetHaircutRate is a free data retrieval call binding the contract method 0x7fdd5a8e.
//
// Solidity: function getHaircutRate() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) GetHaircutRate() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetHaircutRate(&_PlatypusRouterSavax.CallOpts)
}

// GetMaxPriceDeviation is a free data retrieval call binding the contract method 0xddcbc516.
//
// Solidity: function getMaxPriceDeviation() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) GetMaxPriceDeviation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "getMaxPriceDeviation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPriceDeviation is a free data retrieval call binding the contract method 0xddcbc516.
//
// Solidity: function getMaxPriceDeviation() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) GetMaxPriceDeviation() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetMaxPriceDeviation(&_PlatypusRouterSavax.CallOpts)
}

// GetMaxPriceDeviation is a free data retrieval call binding the contract method 0xddcbc516.
//
// Solidity: function getMaxPriceDeviation() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) GetMaxPriceDeviation() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetMaxPriceDeviation(&_PlatypusRouterSavax.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) GetPriceOracle() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.GetPriceOracle(&_PlatypusRouterSavax.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) GetPriceOracle() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.GetPriceOracle(&_PlatypusRouterSavax.CallOpts)
}

// GetRetentionRatio is a free data retrieval call binding the contract method 0xcb733d7a.
//
// Solidity: function getRetentionRatio() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) GetRetentionRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "getRetentionRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRetentionRatio is a free data retrieval call binding the contract method 0xcb733d7a.
//
// Solidity: function getRetentionRatio() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) GetRetentionRatio() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetRetentionRatio(&_PlatypusRouterSavax.CallOpts)
}

// GetRetentionRatio is a free data retrieval call binding the contract method 0xcb733d7a.
//
// Solidity: function getRetentionRatio() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) GetRetentionRatio() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetRetentionRatio(&_PlatypusRouterSavax.CallOpts)
}

// GetSlippageParamK is a free data retrieval call binding the contract method 0x55af008a.
//
// Solidity: function getSlippageParamK() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) GetSlippageParamK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "getSlippageParamK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlippageParamK is a free data retrieval call binding the contract method 0x55af008a.
//
// Solidity: function getSlippageParamK() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) GetSlippageParamK() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetSlippageParamK(&_PlatypusRouterSavax.CallOpts)
}

// GetSlippageParamK is a free data retrieval call binding the contract method 0x55af008a.
//
// Solidity: function getSlippageParamK() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) GetSlippageParamK() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetSlippageParamK(&_PlatypusRouterSavax.CallOpts)
}

// GetSlippageParamN is a free data retrieval call binding the contract method 0x7727c655.
//
// Solidity: function getSlippageParamN() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) GetSlippageParamN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "getSlippageParamN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlippageParamN is a free data retrieval call binding the contract method 0x7727c655.
//
// Solidity: function getSlippageParamN() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) GetSlippageParamN() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetSlippageParamN(&_PlatypusRouterSavax.CallOpts)
}

// GetSlippageParamN is a free data retrieval call binding the contract method 0x7727c655.
//
// Solidity: function getSlippageParamN() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) GetSlippageParamN() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetSlippageParamN(&_PlatypusRouterSavax.CallOpts)
}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() view returns(address[])
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) GetTokenAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "getTokenAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() view returns(address[])
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) GetTokenAddresses() ([]common.Address, error) {
	return _PlatypusRouterSavax.Contract.GetTokenAddresses(&_PlatypusRouterSavax.CallOpts)
}

// GetTokenAddresses is a free data retrieval call binding the contract method 0xee8c24b8.
//
// Solidity: function getTokenAddresses() view returns(address[])
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) GetTokenAddresses() ([]common.Address, error) {
	return _PlatypusRouterSavax.Contract.GetTokenAddresses(&_PlatypusRouterSavax.CallOpts)
}

// GetXThreshold is a free data retrieval call binding the contract method 0x7a1c36d1.
//
// Solidity: function getXThreshold() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) GetXThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "getXThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetXThreshold is a free data retrieval call binding the contract method 0x7a1c36d1.
//
// Solidity: function getXThreshold() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) GetXThreshold() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetXThreshold(&_PlatypusRouterSavax.CallOpts)
}

// GetXThreshold is a free data retrieval call binding the contract method 0x7a1c36d1.
//
// Solidity: function getXThreshold() view returns(uint256)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) GetXThreshold() (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.GetXThreshold(&_PlatypusRouterSavax.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) Owner() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.Owner(&_PlatypusRouterSavax.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) Owner() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.Owner(&_PlatypusRouterSavax.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) Paused() (bool, error) {
	return _PlatypusRouterSavax.Contract.Paused(&_PlatypusRouterSavax.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) Paused() (bool, error) {
	return _PlatypusRouterSavax.Contract.Paused(&_PlatypusRouterSavax.CallOpts)
}

// QuoteMaxInitialAssetWithdrawable is a free data retrieval call binding the contract method 0x2a803b5e.
//
// Solidity: function quoteMaxInitialAssetWithdrawable(address initialToken, address wantedToken) view returns(uint256 maxInitialAssetAmount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) QuoteMaxInitialAssetWithdrawable(opts *bind.CallOpts, initialToken common.Address, wantedToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "quoteMaxInitialAssetWithdrawable", initialToken, wantedToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteMaxInitialAssetWithdrawable is a free data retrieval call binding the contract method 0x2a803b5e.
//
// Solidity: function quoteMaxInitialAssetWithdrawable(address initialToken, address wantedToken) view returns(uint256 maxInitialAssetAmount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) QuoteMaxInitialAssetWithdrawable(initialToken common.Address, wantedToken common.Address) (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.QuoteMaxInitialAssetWithdrawable(&_PlatypusRouterSavax.CallOpts, initialToken, wantedToken)
}

// QuoteMaxInitialAssetWithdrawable is a free data retrieval call binding the contract method 0x2a803b5e.
//
// Solidity: function quoteMaxInitialAssetWithdrawable(address initialToken, address wantedToken) view returns(uint256 maxInitialAssetAmount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) QuoteMaxInitialAssetWithdrawable(initialToken common.Address, wantedToken common.Address) (*big.Int, error) {
	return _PlatypusRouterSavax.Contract.QuoteMaxInitialAssetWithdrawable(&_PlatypusRouterSavax.CallOpts, initialToken, wantedToken)
}

// QuotePotentialSwap is a free data retrieval call binding the contract method 0x43c2e2f5.
//
// Solidity: function quotePotentialSwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 potentialOutcome, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) QuotePotentialSwap(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int) (struct {
	PotentialOutcome *big.Int
	Haircut          *big.Int
}, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "quotePotentialSwap", fromToken, toToken, fromAmount)

	outstruct := new(struct {
		PotentialOutcome *big.Int
		Haircut          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PotentialOutcome = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Haircut = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuotePotentialSwap is a free data retrieval call binding the contract method 0x43c2e2f5.
//
// Solidity: function quotePotentialSwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 potentialOutcome, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) QuotePotentialSwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (struct {
	PotentialOutcome *big.Int
	Haircut          *big.Int
}, error) {
	return _PlatypusRouterSavax.Contract.QuotePotentialSwap(&_PlatypusRouterSavax.CallOpts, fromToken, toToken, fromAmount)
}

// QuotePotentialSwap is a free data retrieval call binding the contract method 0x43c2e2f5.
//
// Solidity: function quotePotentialSwap(address fromToken, address toToken, uint256 fromAmount) view returns(uint256 potentialOutcome, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) QuotePotentialSwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int) (struct {
	PotentialOutcome *big.Int
	Haircut          *big.Int
}, error) {
	return _PlatypusRouterSavax.Contract.QuotePotentialSwap(&_PlatypusRouterSavax.CallOpts, fromToken, toToken, fromAmount)
}

// QuotePotentialWithdraw is a free data retrieval call binding the contract method 0x907448ed.
//
// Solidity: function quotePotentialWithdraw(address token, uint256 liquidity) view returns(uint256 amount, uint256 fee, bool enoughCash)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) QuotePotentialWithdraw(opts *bind.CallOpts, token common.Address, liquidity *big.Int) (struct {
	Amount     *big.Int
	Fee        *big.Int
	EnoughCash bool
}, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "quotePotentialWithdraw", token, liquidity)

	outstruct := new(struct {
		Amount     *big.Int
		Fee        *big.Int
		EnoughCash bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EnoughCash = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// QuotePotentialWithdraw is a free data retrieval call binding the contract method 0x907448ed.
//
// Solidity: function quotePotentialWithdraw(address token, uint256 liquidity) view returns(uint256 amount, uint256 fee, bool enoughCash)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) QuotePotentialWithdraw(token common.Address, liquidity *big.Int) (struct {
	Amount     *big.Int
	Fee        *big.Int
	EnoughCash bool
}, error) {
	return _PlatypusRouterSavax.Contract.QuotePotentialWithdraw(&_PlatypusRouterSavax.CallOpts, token, liquidity)
}

// QuotePotentialWithdraw is a free data retrieval call binding the contract method 0x907448ed.
//
// Solidity: function quotePotentialWithdraw(address token, uint256 liquidity) view returns(uint256 amount, uint256 fee, bool enoughCash)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) QuotePotentialWithdraw(token common.Address, liquidity *big.Int) (struct {
	Amount     *big.Int
	Fee        *big.Int
	EnoughCash bool
}, error) {
	return _PlatypusRouterSavax.Contract.QuotePotentialWithdraw(&_PlatypusRouterSavax.CallOpts, token, liquidity)
}

// QuotePotentialWithdrawFromOtherAsset is a free data retrieval call binding the contract method 0xa4275ceb.
//
// Solidity: function quotePotentialWithdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity) view returns(uint256 amount, uint256 fee)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) QuotePotentialWithdrawFromOtherAsset(opts *bind.CallOpts, initialToken common.Address, wantedToken common.Address, liquidity *big.Int) (struct {
	Amount *big.Int
	Fee    *big.Int
}, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "quotePotentialWithdrawFromOtherAsset", initialToken, wantedToken, liquidity)

	outstruct := new(struct {
		Amount *big.Int
		Fee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuotePotentialWithdrawFromOtherAsset is a free data retrieval call binding the contract method 0xa4275ceb.
//
// Solidity: function quotePotentialWithdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity) view returns(uint256 amount, uint256 fee)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) QuotePotentialWithdrawFromOtherAsset(initialToken common.Address, wantedToken common.Address, liquidity *big.Int) (struct {
	Amount *big.Int
	Fee    *big.Int
}, error) {
	return _PlatypusRouterSavax.Contract.QuotePotentialWithdrawFromOtherAsset(&_PlatypusRouterSavax.CallOpts, initialToken, wantedToken, liquidity)
}

// QuotePotentialWithdrawFromOtherAsset is a free data retrieval call binding the contract method 0xa4275ceb.
//
// Solidity: function quotePotentialWithdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity) view returns(uint256 amount, uint256 fee)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) QuotePotentialWithdrawFromOtherAsset(initialToken common.Address, wantedToken common.Address, liquidity *big.Int) (struct {
	Amount *big.Int
	Fee    *big.Int
}, error) {
	return _PlatypusRouterSavax.Contract.QuotePotentialWithdrawFromOtherAsset(&_PlatypusRouterSavax.CallOpts, initialToken, wantedToken, liquidity)
}

// SAvax is a free data retrieval call binding the contract method 0x3d05f2c3.
//
// Solidity: function sAvax() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) SAvax(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "sAvax")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SAvax is a free data retrieval call binding the contract method 0x3d05f2c3.
//
// Solidity: function sAvax() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) SAvax() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.SAvax(&_PlatypusRouterSavax.CallOpts)
}

// SAvax is a free data retrieval call binding the contract method 0x3d05f2c3.
//
// Solidity: function sAvax() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) SAvax() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.SAvax(&_PlatypusRouterSavax.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) Weth() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.Weth(&_PlatypusRouterSavax.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) Weth() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.Weth(&_PlatypusRouterSavax.CallOpts)
}

// WethForwarder is a free data retrieval call binding the contract method 0x03b2f7aa.
//
// Solidity: function wethForwarder() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCaller) WethForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusRouterSavax.contract.Call(opts, &out, "wethForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethForwarder is a free data retrieval call binding the contract method 0x03b2f7aa.
//
// Solidity: function wethForwarder() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) WethForwarder() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.WethForwarder(&_PlatypusRouterSavax.CallOpts)
}

// WethForwarder is a free data retrieval call binding the contract method 0x03b2f7aa.
//
// Solidity: function wethForwarder() view returns(address)
func (_PlatypusRouterSavax *PlatypusRouterSavaxCallerSession) WethForwarder() (common.Address, error) {
	return _PlatypusRouterSavax.Contract.WethForwarder(&_PlatypusRouterSavax.CallOpts)
}

// AddAsset is a paid mutator transaction binding the contract method 0xda489997.
//
// Solidity: function addAsset(address token, address asset) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) AddAsset(opts *bind.TransactOpts, token common.Address, asset common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "addAsset", token, asset)
}

// AddAsset is a paid mutator transaction binding the contract method 0xda489997.
//
// Solidity: function addAsset(address token, address asset) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) AddAsset(token common.Address, asset common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.AddAsset(&_PlatypusRouterSavax.TransactOpts, token, asset)
}

// AddAsset is a paid mutator transaction binding the contract method 0xda489997.
//
// Solidity: function addAsset(address token, address asset) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) AddAsset(token common.Address, asset common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.AddAsset(&_PlatypusRouterSavax.TransactOpts, token, asset)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address token, uint256 amount, address to, uint256 deadline) returns(uint256 liquidity)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "deposit", token, amount, to, deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address token, uint256 amount, address to, uint256 deadline) returns(uint256 liquidity)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) Deposit(token common.Address, amount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Deposit(&_PlatypusRouterSavax.TransactOpts, token, amount, to, deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address token, uint256 amount, address to, uint256 deadline) returns(uint256 liquidity)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) Deposit(token common.Address, amount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Deposit(&_PlatypusRouterSavax.TransactOpts, token, amount, to, deadline)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address to, uint256 deadline) payable returns(uint256 liquidity)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) DepositETH(opts *bind.TransactOpts, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "depositETH", to, deadline)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address to, uint256 deadline) payable returns(uint256 liquidity)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) DepositETH(to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.DepositETH(&_PlatypusRouterSavax.TransactOpts, to, deadline)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address to, uint256 deadline) payable returns(uint256 liquidity)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) DepositETH(to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.DepositETH(&_PlatypusRouterSavax.TransactOpts, to, deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address weth_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) Initialize(opts *bind.TransactOpts, weth_ common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "initialize", weth_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address weth_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) Initialize(weth_ common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Initialize(&_PlatypusRouterSavax.TransactOpts, weth_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address weth_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) Initialize(weth_ common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Initialize(&_PlatypusRouterSavax.TransactOpts, weth_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) Pause() (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Pause(&_PlatypusRouterSavax.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) Pause() (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Pause(&_PlatypusRouterSavax.TransactOpts)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x4a5e42b1.
//
// Solidity: function removeAsset(address key) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) RemoveAsset(opts *bind.TransactOpts, key common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "removeAsset", key)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x4a5e42b1.
//
// Solidity: function removeAsset(address key) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) RemoveAsset(key common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.RemoveAsset(&_PlatypusRouterSavax.TransactOpts, key)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x4a5e42b1.
//
// Solidity: function removeAsset(address key) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) RemoveAsset(key common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.RemoveAsset(&_PlatypusRouterSavax.TransactOpts, key)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) RenounceOwnership() (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.RenounceOwnership(&_PlatypusRouterSavax.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.RenounceOwnership(&_PlatypusRouterSavax.TransactOpts)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) SetDev(opts *bind.TransactOpts, dev common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "setDev", dev)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) SetDev(dev common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetDev(&_PlatypusRouterSavax.TransactOpts, dev)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) SetDev(dev common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetDev(&_PlatypusRouterSavax.TransactOpts, dev)
}

// SetHaircutRate is a paid mutator transaction binding the contract method 0xf57e84d5.
//
// Solidity: function setHaircutRate(uint256 haircutRate_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) SetHaircutRate(opts *bind.TransactOpts, haircutRate_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "setHaircutRate", haircutRate_)
}

// SetHaircutRate is a paid mutator transaction binding the contract method 0xf57e84d5.
//
// Solidity: function setHaircutRate(uint256 haircutRate_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) SetHaircutRate(haircutRate_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetHaircutRate(&_PlatypusRouterSavax.TransactOpts, haircutRate_)
}

// SetHaircutRate is a paid mutator transaction binding the contract method 0xf57e84d5.
//
// Solidity: function setHaircutRate(uint256 haircutRate_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) SetHaircutRate(haircutRate_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetHaircutRate(&_PlatypusRouterSavax.TransactOpts, haircutRate_)
}

// SetMaxPriceDeviation is a paid mutator transaction binding the contract method 0x9ee4c057.
//
// Solidity: function setMaxPriceDeviation(uint256 maxPriceDeviation_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) SetMaxPriceDeviation(opts *bind.TransactOpts, maxPriceDeviation_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "setMaxPriceDeviation", maxPriceDeviation_)
}

// SetMaxPriceDeviation is a paid mutator transaction binding the contract method 0x9ee4c057.
//
// Solidity: function setMaxPriceDeviation(uint256 maxPriceDeviation_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) SetMaxPriceDeviation(maxPriceDeviation_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetMaxPriceDeviation(&_PlatypusRouterSavax.TransactOpts, maxPriceDeviation_)
}

// SetMaxPriceDeviation is a paid mutator transaction binding the contract method 0x9ee4c057.
//
// Solidity: function setMaxPriceDeviation(uint256 maxPriceDeviation_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) SetMaxPriceDeviation(maxPriceDeviation_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetMaxPriceDeviation(&_PlatypusRouterSavax.TransactOpts, maxPriceDeviation_)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) SetPriceOracle(opts *bind.TransactOpts, priceOracle common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "setPriceOracle", priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetPriceOracle(&_PlatypusRouterSavax.TransactOpts, priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetPriceOracle(&_PlatypusRouterSavax.TransactOpts, priceOracle)
}

// SetRetentionRatio is a paid mutator transaction binding the contract method 0x44db964a.
//
// Solidity: function setRetentionRatio(uint256 retentionRatio_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) SetRetentionRatio(opts *bind.TransactOpts, retentionRatio_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "setRetentionRatio", retentionRatio_)
}

// SetRetentionRatio is a paid mutator transaction binding the contract method 0x44db964a.
//
// Solidity: function setRetentionRatio(uint256 retentionRatio_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) SetRetentionRatio(retentionRatio_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetRetentionRatio(&_PlatypusRouterSavax.TransactOpts, retentionRatio_)
}

// SetRetentionRatio is a paid mutator transaction binding the contract method 0x44db964a.
//
// Solidity: function setRetentionRatio(uint256 retentionRatio_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) SetRetentionRatio(retentionRatio_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetRetentionRatio(&_PlatypusRouterSavax.TransactOpts, retentionRatio_)
}

// SetSAvax is a paid mutator transaction binding the contract method 0x77dd6d7e.
//
// Solidity: function setSAvax(address sAvax_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) SetSAvax(opts *bind.TransactOpts, sAvax_ common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "setSAvax", sAvax_)
}

// SetSAvax is a paid mutator transaction binding the contract method 0x77dd6d7e.
//
// Solidity: function setSAvax(address sAvax_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) SetSAvax(sAvax_ common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetSAvax(&_PlatypusRouterSavax.TransactOpts, sAvax_)
}

// SetSAvax is a paid mutator transaction binding the contract method 0x77dd6d7e.
//
// Solidity: function setSAvax(address sAvax_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) SetSAvax(sAvax_ common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetSAvax(&_PlatypusRouterSavax.TransactOpts, sAvax_)
}

// SetSlippageParams is a paid mutator transaction binding the contract method 0xdf220181.
//
// Solidity: function setSlippageParams(uint256 k_, uint256 n_, uint256 c1_, uint256 xThreshold_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) SetSlippageParams(opts *bind.TransactOpts, k_ *big.Int, n_ *big.Int, c1_ *big.Int, xThreshold_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "setSlippageParams", k_, n_, c1_, xThreshold_)
}

// SetSlippageParams is a paid mutator transaction binding the contract method 0xdf220181.
//
// Solidity: function setSlippageParams(uint256 k_, uint256 n_, uint256 c1_, uint256 xThreshold_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) SetSlippageParams(k_ *big.Int, n_ *big.Int, c1_ *big.Int, xThreshold_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetSlippageParams(&_PlatypusRouterSavax.TransactOpts, k_, n_, c1_, xThreshold_)
}

// SetSlippageParams is a paid mutator transaction binding the contract method 0xdf220181.
//
// Solidity: function setSlippageParams(uint256 k_, uint256 n_, uint256 c1_, uint256 xThreshold_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) SetSlippageParams(k_ *big.Int, n_ *big.Int, c1_ *big.Int, xThreshold_ *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetSlippageParams(&_PlatypusRouterSavax.TransactOpts, k_, n_, c1_, xThreshold_)
}

// SetWETHForwarder is a paid mutator transaction binding the contract method 0x57eab9b2.
//
// Solidity: function setWETHForwarder(address wethForwarder_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) SetWETHForwarder(opts *bind.TransactOpts, wethForwarder_ common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "setWETHForwarder", wethForwarder_)
}

// SetWETHForwarder is a paid mutator transaction binding the contract method 0x57eab9b2.
//
// Solidity: function setWETHForwarder(address wethForwarder_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) SetWETHForwarder(wethForwarder_ common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetWETHForwarder(&_PlatypusRouterSavax.TransactOpts, wethForwarder_)
}

// SetWETHForwarder is a paid mutator transaction binding the contract method 0x57eab9b2.
//
// Solidity: function setWETHForwarder(address wethForwarder_) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) SetWETHForwarder(wethForwarder_ common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SetWETHForwarder(&_PlatypusRouterSavax.TransactOpts, wethForwarder_)
}

// Swap is a paid mutator transaction binding the contract method 0x9908fc8b.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 actualToAmount, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) Swap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "swap", fromToken, toToken, fromAmount, minimumToAmount, to, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x9908fc8b.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 actualToAmount, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Swap(&_PlatypusRouterSavax.TransactOpts, fromToken, toToken, fromAmount, minimumToAmount, to, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x9908fc8b.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 actualToAmount, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) Swap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Swap(&_PlatypusRouterSavax.TransactOpts, fromToken, toToken, fromAmount, minimumToAmount, to, deadline)
}

// SwapFromETH is a paid mutator transaction binding the contract method 0x2d59d767.
//
// Solidity: function swapFromETH(address toToken, uint256 minimumToAmount, address to, uint256 deadline) payable returns(uint256 actualToAmount, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) SwapFromETH(opts *bind.TransactOpts, toToken common.Address, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "swapFromETH", toToken, minimumToAmount, to, deadline)
}

// SwapFromETH is a paid mutator transaction binding the contract method 0x2d59d767.
//
// Solidity: function swapFromETH(address toToken, uint256 minimumToAmount, address to, uint256 deadline) payable returns(uint256 actualToAmount, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) SwapFromETH(toToken common.Address, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SwapFromETH(&_PlatypusRouterSavax.TransactOpts, toToken, minimumToAmount, to, deadline)
}

// SwapFromETH is a paid mutator transaction binding the contract method 0x2d59d767.
//
// Solidity: function swapFromETH(address toToken, uint256 minimumToAmount, address to, uint256 deadline) payable returns(uint256 actualToAmount, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) SwapFromETH(toToken common.Address, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SwapFromETH(&_PlatypusRouterSavax.TransactOpts, toToken, minimumToAmount, to, deadline)
}

// SwapToETH is a paid mutator transaction binding the contract method 0x1192903f.
//
// Solidity: function swapToETH(address fromToken, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 actualToAmount, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) SwapToETH(opts *bind.TransactOpts, fromToken common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "swapToETH", fromToken, fromAmount, minimumToAmount, to, deadline)
}

// SwapToETH is a paid mutator transaction binding the contract method 0x1192903f.
//
// Solidity: function swapToETH(address fromToken, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 actualToAmount, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) SwapToETH(fromToken common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SwapToETH(&_PlatypusRouterSavax.TransactOpts, fromToken, fromAmount, minimumToAmount, to, deadline)
}

// SwapToETH is a paid mutator transaction binding the contract method 0x1192903f.
//
// Solidity: function swapToETH(address fromToken, uint256 fromAmount, uint256 minimumToAmount, address to, uint256 deadline) returns(uint256 actualToAmount, uint256 haircut)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) SwapToETH(fromToken common.Address, fromAmount *big.Int, minimumToAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.SwapToETH(&_PlatypusRouterSavax.TransactOpts, fromToken, fromAmount, minimumToAmount, to, deadline)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.TransferOwnership(&_PlatypusRouterSavax.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.TransferOwnership(&_PlatypusRouterSavax.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) Unpause() (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Unpause(&_PlatypusRouterSavax.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) Unpause() (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Unpause(&_PlatypusRouterSavax.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x09a5fca3.
//
// Solidity: function withdraw(address token, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "withdraw", token, liquidity, minimumAmount, to, deadline)
}

// Withdraw is a paid mutator transaction binding the contract method 0x09a5fca3.
//
// Solidity: function withdraw(address token, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) Withdraw(token common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Withdraw(&_PlatypusRouterSavax.TransactOpts, token, liquidity, minimumAmount, to, deadline)
}

// Withdraw is a paid mutator transaction binding the contract method 0x09a5fca3.
//
// Solidity: function withdraw(address token, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) Withdraw(token common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.Withdraw(&_PlatypusRouterSavax.TransactOpts, token, liquidity, minimumAmount, to, deadline)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x7fc4b3ec.
//
// Solidity: function withdrawETH(uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) WithdrawETH(opts *bind.TransactOpts, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "withdrawETH", liquidity, minimumAmount, to, deadline)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x7fc4b3ec.
//
// Solidity: function withdrawETH(uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) WithdrawETH(liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.WithdrawETH(&_PlatypusRouterSavax.TransactOpts, liquidity, minimumAmount, to, deadline)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x7fc4b3ec.
//
// Solidity: function withdrawETH(uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) WithdrawETH(liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.WithdrawETH(&_PlatypusRouterSavax.TransactOpts, liquidity, minimumAmount, to, deadline)
}

// WithdrawFromOtherAsset is a paid mutator transaction binding the contract method 0x0f91f06f.
//
// Solidity: function withdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactor) WithdrawFromOtherAsset(opts *bind.TransactOpts, initialToken common.Address, wantedToken common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.contract.Transact(opts, "withdrawFromOtherAsset", initialToken, wantedToken, liquidity, minimumAmount, to, deadline)
}

// WithdrawFromOtherAsset is a paid mutator transaction binding the contract method 0x0f91f06f.
//
// Solidity: function withdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxSession) WithdrawFromOtherAsset(initialToken common.Address, wantedToken common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.WithdrawFromOtherAsset(&_PlatypusRouterSavax.TransactOpts, initialToken, wantedToken, liquidity, minimumAmount, to, deadline)
}

// WithdrawFromOtherAsset is a paid mutator transaction binding the contract method 0x0f91f06f.
//
// Solidity: function withdrawFromOtherAsset(address initialToken, address wantedToken, uint256 liquidity, uint256 minimumAmount, address to, uint256 deadline) returns(uint256 amount)
func (_PlatypusRouterSavax *PlatypusRouterSavaxTransactorSession) WithdrawFromOtherAsset(initialToken common.Address, wantedToken common.Address, liquidity *big.Int, minimumAmount *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PlatypusRouterSavax.Contract.WithdrawFromOtherAsset(&_PlatypusRouterSavax.TransactOpts, initialToken, wantedToken, liquidity, minimumAmount, to, deadline)
}
