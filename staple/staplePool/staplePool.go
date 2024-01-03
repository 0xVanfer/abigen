// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staplePool

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

// StaplePoolMetaData contains all meta data concerning the StaplePool contract.
var StaplePoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"CollectFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Recover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"assetTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"collectFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAssets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAssets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recordSwapIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recordSwapOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tl\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalManagedAssets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"transferByController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newTl\",\"type\":\"uint8\"}],\"name\":\"updateTL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userLiability\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StaplePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use StaplePoolMetaData.ABI instead.
var StaplePoolABI = StaplePoolMetaData.ABI

// StaplePool is an auto generated Go binding around an Ethereum contract.
type StaplePool struct {
	StaplePoolCaller     // Read-only binding to the contract
	StaplePoolTransactor // Write-only binding to the contract
	StaplePoolFilterer   // Log filterer for contract events
}

// StaplePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaplePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaplePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaplePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaplePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaplePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaplePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaplePoolSession struct {
	Contract     *StaplePool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaplePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaplePoolCallerSession struct {
	Contract *StaplePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StaplePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaplePoolTransactorSession struct {
	Contract     *StaplePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StaplePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaplePoolRaw struct {
	Contract *StaplePool // Generic contract binding to access the raw methods on
}

// StaplePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaplePoolCallerRaw struct {
	Contract *StaplePoolCaller // Generic read-only contract binding to access the raw methods on
}

// StaplePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaplePoolTransactorRaw struct {
	Contract *StaplePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaplePool creates a new instance of StaplePool, bound to a specific deployed contract.
func NewStaplePool(address common.Address, backend bind.ContractBackend) (*StaplePool, error) {
	contract, err := bindStaplePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaplePool{StaplePoolCaller: StaplePoolCaller{contract: contract}, StaplePoolTransactor: StaplePoolTransactor{contract: contract}, StaplePoolFilterer: StaplePoolFilterer{contract: contract}}, nil
}

// NewStaplePoolCaller creates a new read-only instance of StaplePool, bound to a specific deployed contract.
func NewStaplePoolCaller(address common.Address, caller bind.ContractCaller) (*StaplePoolCaller, error) {
	contract, err := bindStaplePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaplePoolCaller{contract: contract}, nil
}

// NewStaplePoolTransactor creates a new write-only instance of StaplePool, bound to a specific deployed contract.
func NewStaplePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*StaplePoolTransactor, error) {
	contract, err := bindStaplePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaplePoolTransactor{contract: contract}, nil
}

// NewStaplePoolFilterer creates a new log filterer instance of StaplePool, bound to a specific deployed contract.
func NewStaplePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*StaplePoolFilterer, error) {
	contract, err := bindStaplePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaplePoolFilterer{contract: contract}, nil
}

// bindStaplePool binds a generic wrapper to an already deployed contract.
func bindStaplePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StaplePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaplePool *StaplePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaplePool.Contract.StaplePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaplePool *StaplePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaplePool.Contract.StaplePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaplePool *StaplePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaplePool.Contract.StaplePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaplePool *StaplePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaplePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaplePool *StaplePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaplePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaplePool *StaplePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaplePool.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StaplePool *StaplePoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StaplePool *StaplePoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StaplePool.Contract.Allowance(&_StaplePool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StaplePool *StaplePoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StaplePool.Contract.Allowance(&_StaplePool.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address assetTokenAddress)
func (_StaplePool *StaplePoolCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address assetTokenAddress)
func (_StaplePool *StaplePoolSession) Asset() (common.Address, error) {
	return _StaplePool.Contract.Asset(&_StaplePool.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address assetTokenAddress)
func (_StaplePool *StaplePoolCallerSession) Asset() (common.Address, error) {
	return _StaplePool.Contract.Asset(&_StaplePool.CallOpts)
}

// Assets is a free data retrieval call binding the contract method 0x71a97305.
//
// Solidity: function assets() view returns(uint256)
func (_StaplePool *StaplePoolCaller) Assets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "assets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Assets is a free data retrieval call binding the contract method 0x71a97305.
//
// Solidity: function assets() view returns(uint256)
func (_StaplePool *StaplePoolSession) Assets() (*big.Int, error) {
	return _StaplePool.Contract.Assets(&_StaplePool.CallOpts)
}

// Assets is a free data retrieval call binding the contract method 0x71a97305.
//
// Solidity: function assets() view returns(uint256)
func (_StaplePool *StaplePoolCallerSession) Assets() (*big.Int, error) {
	return _StaplePool.Contract.Assets(&_StaplePool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StaplePool *StaplePoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StaplePool *StaplePoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StaplePool.Contract.BalanceOf(&_StaplePool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StaplePool *StaplePoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StaplePool.Contract.BalanceOf(&_StaplePool.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_StaplePool *StaplePoolCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_StaplePool *StaplePoolSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.ConvertToAssets(&_StaplePool.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_StaplePool *StaplePoolCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.ConvertToAssets(&_StaplePool.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_StaplePool *StaplePoolCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_StaplePool *StaplePoolSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.ConvertToShares(&_StaplePool.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_StaplePool *StaplePoolCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.ConvertToShares(&_StaplePool.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StaplePool *StaplePoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StaplePool *StaplePoolSession) Decimals() (uint8, error) {
	return _StaplePool.Contract.Decimals(&_StaplePool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StaplePool *StaplePoolCallerSession) Decimals() (uint8, error) {
	return _StaplePool.Contract.Decimals(&_StaplePool.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256 maxAssets)
func (_StaplePool *StaplePoolCaller) MaxDeposit(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "maxDeposit", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256 maxAssets)
func (_StaplePool *StaplePoolSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _StaplePool.Contract.MaxDeposit(&_StaplePool.CallOpts, receiver)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256 maxAssets)
func (_StaplePool *StaplePoolCallerSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _StaplePool.Contract.MaxDeposit(&_StaplePool.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256 maxShares)
func (_StaplePool *StaplePoolCaller) MaxMint(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "maxMint", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256 maxShares)
func (_StaplePool *StaplePoolSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _StaplePool.Contract.MaxMint(&_StaplePool.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256 maxShares)
func (_StaplePool *StaplePoolCallerSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _StaplePool.Contract.MaxMint(&_StaplePool.CallOpts, receiver)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 maxShares)
func (_StaplePool *StaplePoolCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 maxShares)
func (_StaplePool *StaplePoolSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _StaplePool.Contract.MaxRedeem(&_StaplePool.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 maxShares)
func (_StaplePool *StaplePoolCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _StaplePool.Contract.MaxRedeem(&_StaplePool.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 maxAssets)
func (_StaplePool *StaplePoolCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 maxAssets)
func (_StaplePool *StaplePoolSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _StaplePool.Contract.MaxWithdraw(&_StaplePool.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 maxAssets)
func (_StaplePool *StaplePoolCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _StaplePool.Contract.MaxWithdraw(&_StaplePool.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StaplePool *StaplePoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StaplePool *StaplePoolSession) Name() (string, error) {
	return _StaplePool.Contract.Name(&_StaplePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StaplePool *StaplePoolCallerSession) Name() (string, error) {
	return _StaplePool.Contract.Name(&_StaplePool.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256 shares)
func (_StaplePool *StaplePoolCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256 shares)
func (_StaplePool *StaplePoolSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.PreviewDeposit(&_StaplePool.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256 shares)
func (_StaplePool *StaplePoolCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.PreviewDeposit(&_StaplePool.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256 assets)
func (_StaplePool *StaplePoolCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256 assets)
func (_StaplePool *StaplePoolSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.PreviewMint(&_StaplePool.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256 assets)
func (_StaplePool *StaplePoolCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.PreviewMint(&_StaplePool.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_StaplePool *StaplePoolCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_StaplePool *StaplePoolSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.PreviewRedeem(&_StaplePool.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_StaplePool *StaplePoolCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.PreviewRedeem(&_StaplePool.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_StaplePool *StaplePoolCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_StaplePool *StaplePoolSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.PreviewWithdraw(&_StaplePool.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_StaplePool *StaplePoolCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _StaplePool.Contract.PreviewWithdraw(&_StaplePool.CallOpts, assets)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StaplePool *StaplePoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StaplePool *StaplePoolSession) Symbol() (string, error) {
	return _StaplePool.Contract.Symbol(&_StaplePool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StaplePool *StaplePoolCallerSession) Symbol() (string, error) {
	return _StaplePool.Contract.Symbol(&_StaplePool.CallOpts)
}

// Tl is a free data retrieval call binding the contract method 0xaaab671c.
//
// Solidity: function tl() view returns(uint8)
func (_StaplePool *StaplePoolCaller) Tl(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "tl")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Tl is a free data retrieval call binding the contract method 0xaaab671c.
//
// Solidity: function tl() view returns(uint8)
func (_StaplePool *StaplePoolSession) Tl() (uint8, error) {
	return _StaplePool.Contract.Tl(&_StaplePool.CallOpts)
}

// Tl is a free data retrieval call binding the contract method 0xaaab671c.
//
// Solidity: function tl() view returns(uint8)
func (_StaplePool *StaplePoolCallerSession) Tl() (uint8, error) {
	return _StaplePool.Contract.Tl(&_StaplePool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 totalManagedAssets)
func (_StaplePool *StaplePoolCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 totalManagedAssets)
func (_StaplePool *StaplePoolSession) TotalAssets() (*big.Int, error) {
	return _StaplePool.Contract.TotalAssets(&_StaplePool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 totalManagedAssets)
func (_StaplePool *StaplePoolCallerSession) TotalAssets() (*big.Int, error) {
	return _StaplePool.Contract.TotalAssets(&_StaplePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StaplePool *StaplePoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StaplePool *StaplePoolSession) TotalSupply() (*big.Int, error) {
	return _StaplePool.Contract.TotalSupply(&_StaplePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StaplePool *StaplePoolCallerSession) TotalSupply() (*big.Int, error) {
	return _StaplePool.Contract.TotalSupply(&_StaplePool.CallOpts)
}

// UserLiability is a free data retrieval call binding the contract method 0xab5fd3d4.
//
// Solidity: function userLiability(address user) view returns(uint256)
func (_StaplePool *StaplePoolCaller) UserLiability(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaplePool.contract.Call(opts, &out, "userLiability", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLiability is a free data retrieval call binding the contract method 0xab5fd3d4.
//
// Solidity: function userLiability(address user) view returns(uint256)
func (_StaplePool *StaplePoolSession) UserLiability(user common.Address) (*big.Int, error) {
	return _StaplePool.Contract.UserLiability(&_StaplePool.CallOpts, user)
}

// UserLiability is a free data retrieval call binding the contract method 0xab5fd3d4.
//
// Solidity: function userLiability(address user) view returns(uint256)
func (_StaplePool *StaplePoolCallerSession) UserLiability(user common.Address) (*big.Int, error) {
	return _StaplePool.Contract.UserLiability(&_StaplePool.CallOpts, user)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StaplePool *StaplePoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StaplePool *StaplePoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.Approve(&_StaplePool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StaplePool *StaplePoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.Approve(&_StaplePool.TransactOpts, spender, amount)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 assets, address user) returns(uint256 shares)
func (_StaplePool *StaplePoolTransactor) CollectFee(opts *bind.TransactOpts, assets *big.Int, user common.Address) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "collectFee", assets, user)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 assets, address user) returns(uint256 shares)
func (_StaplePool *StaplePoolSession) CollectFee(assets *big.Int, user common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.CollectFee(&_StaplePool.TransactOpts, assets, user)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 assets, address user) returns(uint256 shares)
func (_StaplePool *StaplePoolTransactorSession) CollectFee(assets *big.Int, user common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.CollectFee(&_StaplePool.TransactOpts, assets, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_StaplePool *StaplePoolTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_StaplePool *StaplePoolSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.Deposit(&_StaplePool.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_StaplePool *StaplePoolTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.Deposit(&_StaplePool.TransactOpts, assets, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_StaplePool *StaplePoolTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_StaplePool *StaplePoolSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.Mint(&_StaplePool.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_StaplePool *StaplePoolTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.Mint(&_StaplePool.TransactOpts, shares, receiver)
}

// RecordSwapIn is a paid mutator transaction binding the contract method 0x4783b564.
//
// Solidity: function recordSwapIn(uint256 amount) returns()
func (_StaplePool *StaplePoolTransactor) RecordSwapIn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "recordSwapIn", amount)
}

// RecordSwapIn is a paid mutator transaction binding the contract method 0x4783b564.
//
// Solidity: function recordSwapIn(uint256 amount) returns()
func (_StaplePool *StaplePoolSession) RecordSwapIn(amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.RecordSwapIn(&_StaplePool.TransactOpts, amount)
}

// RecordSwapIn is a paid mutator transaction binding the contract method 0x4783b564.
//
// Solidity: function recordSwapIn(uint256 amount) returns()
func (_StaplePool *StaplePoolTransactorSession) RecordSwapIn(amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.RecordSwapIn(&_StaplePool.TransactOpts, amount)
}

// RecordSwapOut is a paid mutator transaction binding the contract method 0x94d5fc84.
//
// Solidity: function recordSwapOut(uint256 amount) returns()
func (_StaplePool *StaplePoolTransactor) RecordSwapOut(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "recordSwapOut", amount)
}

// RecordSwapOut is a paid mutator transaction binding the contract method 0x94d5fc84.
//
// Solidity: function recordSwapOut(uint256 amount) returns()
func (_StaplePool *StaplePoolSession) RecordSwapOut(amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.RecordSwapOut(&_StaplePool.TransactOpts, amount)
}

// RecordSwapOut is a paid mutator transaction binding the contract method 0x94d5fc84.
//
// Solidity: function recordSwapOut(uint256 amount) returns()
func (_StaplePool *StaplePoolTransactorSession) RecordSwapOut(amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.RecordSwapOut(&_StaplePool.TransactOpts, amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xd04323c5.
//
// Solidity: function recoverToken(address tokenType, uint256 amount, address receiver) returns()
func (_StaplePool *StaplePoolTransactor) RecoverToken(opts *bind.TransactOpts, tokenType common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "recoverToken", tokenType, amount, receiver)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xd04323c5.
//
// Solidity: function recoverToken(address tokenType, uint256 amount, address receiver) returns()
func (_StaplePool *StaplePoolSession) RecoverToken(tokenType common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.RecoverToken(&_StaplePool.TransactOpts, tokenType, amount, receiver)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xd04323c5.
//
// Solidity: function recoverToken(address tokenType, uint256 amount, address receiver) returns()
func (_StaplePool *StaplePoolTransactorSession) RecoverToken(tokenType common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.RecoverToken(&_StaplePool.TransactOpts, tokenType, amount, receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_StaplePool *StaplePoolTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_StaplePool *StaplePoolSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.Redeem(&_StaplePool.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_StaplePool *StaplePoolTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.Redeem(&_StaplePool.TransactOpts, shares, receiver, owner)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StaplePool *StaplePoolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StaplePool *StaplePoolSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.Transfer(&_StaplePool.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StaplePool *StaplePoolTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.Transfer(&_StaplePool.TransactOpts, to, amount)
}

// TransferByController is a paid mutator transaction binding the contract method 0xb777638a.
//
// Solidity: function transferByController(address receiver, uint256 assets) returns()
func (_StaplePool *StaplePoolTransactor) TransferByController(opts *bind.TransactOpts, receiver common.Address, assets *big.Int) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "transferByController", receiver, assets)
}

// TransferByController is a paid mutator transaction binding the contract method 0xb777638a.
//
// Solidity: function transferByController(address receiver, uint256 assets) returns()
func (_StaplePool *StaplePoolSession) TransferByController(receiver common.Address, assets *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.TransferByController(&_StaplePool.TransactOpts, receiver, assets)
}

// TransferByController is a paid mutator transaction binding the contract method 0xb777638a.
//
// Solidity: function transferByController(address receiver, uint256 assets) returns()
func (_StaplePool *StaplePoolTransactorSession) TransferByController(receiver common.Address, assets *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.TransferByController(&_StaplePool.TransactOpts, receiver, assets)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StaplePool *StaplePoolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StaplePool *StaplePoolSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.TransferFrom(&_StaplePool.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StaplePool *StaplePoolTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaplePool.Contract.TransferFrom(&_StaplePool.TransactOpts, from, to, amount)
}

// UpdateTL is a paid mutator transaction binding the contract method 0xa3ff7270.
//
// Solidity: function updateTL(uint8 newTl) returns()
func (_StaplePool *StaplePoolTransactor) UpdateTL(opts *bind.TransactOpts, newTl uint8) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "updateTL", newTl)
}

// UpdateTL is a paid mutator transaction binding the contract method 0xa3ff7270.
//
// Solidity: function updateTL(uint8 newTl) returns()
func (_StaplePool *StaplePoolSession) UpdateTL(newTl uint8) (*types.Transaction, error) {
	return _StaplePool.Contract.UpdateTL(&_StaplePool.TransactOpts, newTl)
}

// UpdateTL is a paid mutator transaction binding the contract method 0xa3ff7270.
//
// Solidity: function updateTL(uint8 newTl) returns()
func (_StaplePool *StaplePoolTransactorSession) UpdateTL(newTl uint8) (*types.Transaction, error) {
	return _StaplePool.Contract.UpdateTL(&_StaplePool.TransactOpts, newTl)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_StaplePool *StaplePoolTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _StaplePool.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_StaplePool *StaplePoolSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.Withdraw(&_StaplePool.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_StaplePool *StaplePoolTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _StaplePool.Contract.Withdraw(&_StaplePool.TransactOpts, assets, receiver, owner)
}

// StaplePoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StaplePool contract.
type StaplePoolApprovalIterator struct {
	Event *StaplePoolApproval // Event containing the contract specifics and raw log

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
func (it *StaplePoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaplePoolApproval)
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
		it.Event = new(StaplePoolApproval)
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
func (it *StaplePoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaplePoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaplePoolApproval represents a Approval event raised by the StaplePool contract.
type StaplePoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StaplePool *StaplePoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StaplePoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StaplePool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StaplePoolApprovalIterator{contract: _StaplePool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StaplePool *StaplePoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StaplePoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StaplePool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaplePoolApproval)
				if err := _StaplePool.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StaplePool *StaplePoolFilterer) ParseApproval(log types.Log) (*StaplePoolApproval, error) {
	event := new(StaplePoolApproval)
	if err := _StaplePool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaplePoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the StaplePool contract.
type StaplePoolBurnIterator struct {
	Event *StaplePoolBurn // Event containing the contract specifics and raw log

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
func (it *StaplePoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaplePoolBurn)
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
		it.Event = new(StaplePoolBurn)
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
func (it *StaplePoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaplePoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaplePoolBurn represents a Burn event raised by the StaplePool contract.
type StaplePoolBurn struct {
	User   common.Address
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 shares)
func (_StaplePool *StaplePoolFilterer) FilterBurn(opts *bind.FilterOpts, user []common.Address) (*StaplePoolBurnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StaplePool.contract.FilterLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return &StaplePoolBurnIterator{contract: _StaplePool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 shares)
func (_StaplePool *StaplePoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *StaplePoolBurn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StaplePool.contract.WatchLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaplePoolBurn)
				if err := _StaplePool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 shares)
func (_StaplePool *StaplePoolFilterer) ParseBurn(log types.Log) (*StaplePoolBurn, error) {
	event := new(StaplePoolBurn)
	if err := _StaplePool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaplePoolCollectFeeIterator is returned from FilterCollectFee and is used to iterate over the raw logs and unpacked data for CollectFee events raised by the StaplePool contract.
type StaplePoolCollectFeeIterator struct {
	Event *StaplePoolCollectFee // Event containing the contract specifics and raw log

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
func (it *StaplePoolCollectFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaplePoolCollectFee)
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
		it.Event = new(StaplePoolCollectFee)
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
func (it *StaplePoolCollectFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaplePoolCollectFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaplePoolCollectFee represents a CollectFee event raised by the StaplePool contract.
type StaplePoolCollectFee struct {
	User   common.Address
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollectFee is a free log retrieval operation binding the contract event 0xdfeef5879c25440d1418a79a16489f0e739d22ac040c8c76f7b998fc704edeb8.
//
// Solidity: event CollectFee(address indexed user, uint256 shares)
func (_StaplePool *StaplePoolFilterer) FilterCollectFee(opts *bind.FilterOpts, user []common.Address) (*StaplePoolCollectFeeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StaplePool.contract.FilterLogs(opts, "CollectFee", userRule)
	if err != nil {
		return nil, err
	}
	return &StaplePoolCollectFeeIterator{contract: _StaplePool.contract, event: "CollectFee", logs: logs, sub: sub}, nil
}

// WatchCollectFee is a free log subscription operation binding the contract event 0xdfeef5879c25440d1418a79a16489f0e739d22ac040c8c76f7b998fc704edeb8.
//
// Solidity: event CollectFee(address indexed user, uint256 shares)
func (_StaplePool *StaplePoolFilterer) WatchCollectFee(opts *bind.WatchOpts, sink chan<- *StaplePoolCollectFee, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StaplePool.contract.WatchLogs(opts, "CollectFee", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaplePoolCollectFee)
				if err := _StaplePool.contract.UnpackLog(event, "CollectFee", log); err != nil {
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

// ParseCollectFee is a log parse operation binding the contract event 0xdfeef5879c25440d1418a79a16489f0e739d22ac040c8c76f7b998fc704edeb8.
//
// Solidity: event CollectFee(address indexed user, uint256 shares)
func (_StaplePool *StaplePoolFilterer) ParseCollectFee(log types.Log) (*StaplePoolCollectFee, error) {
	event := new(StaplePoolCollectFee)
	if err := _StaplePool.contract.UnpackLog(event, "CollectFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaplePoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the StaplePool contract.
type StaplePoolDepositIterator struct {
	Event *StaplePoolDeposit // Event containing the contract specifics and raw log

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
func (it *StaplePoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaplePoolDeposit)
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
		it.Event = new(StaplePoolDeposit)
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
func (it *StaplePoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaplePoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaplePoolDeposit represents a Deposit event raised by the StaplePool contract.
type StaplePoolDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_StaplePool *StaplePoolFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*StaplePoolDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StaplePool.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &StaplePoolDepositIterator{contract: _StaplePool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_StaplePool *StaplePoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StaplePoolDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StaplePool.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaplePoolDeposit)
				if err := _StaplePool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_StaplePool *StaplePoolFilterer) ParseDeposit(log types.Log) (*StaplePoolDeposit, error) {
	event := new(StaplePoolDeposit)
	if err := _StaplePool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaplePoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the StaplePool contract.
type StaplePoolMintIterator struct {
	Event *StaplePoolMint // Event containing the contract specifics and raw log

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
func (it *StaplePoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaplePoolMint)
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
		it.Event = new(StaplePoolMint)
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
func (it *StaplePoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaplePoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaplePoolMint represents a Mint event raised by the StaplePool contract.
type StaplePoolMint struct {
	User   common.Address
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed user, uint256 shares)
func (_StaplePool *StaplePoolFilterer) FilterMint(opts *bind.FilterOpts, user []common.Address) (*StaplePoolMintIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StaplePool.contract.FilterLogs(opts, "Mint", userRule)
	if err != nil {
		return nil, err
	}
	return &StaplePoolMintIterator{contract: _StaplePool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed user, uint256 shares)
func (_StaplePool *StaplePoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *StaplePoolMint, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StaplePool.contract.WatchLogs(opts, "Mint", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaplePoolMint)
				if err := _StaplePool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed user, uint256 shares)
func (_StaplePool *StaplePoolFilterer) ParseMint(log types.Log) (*StaplePoolMint, error) {
	event := new(StaplePoolMint)
	if err := _StaplePool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaplePoolRecoverIterator is returned from FilterRecover and is used to iterate over the raw logs and unpacked data for Recover events raised by the StaplePool contract.
type StaplePoolRecoverIterator struct {
	Event *StaplePoolRecover // Event containing the contract specifics and raw log

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
func (it *StaplePoolRecoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaplePoolRecover)
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
		it.Event = new(StaplePoolRecover)
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
func (it *StaplePoolRecoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaplePoolRecoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaplePoolRecover represents a Recover event raised by the StaplePool contract.
type StaplePoolRecover struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRecover is a free log retrieval operation binding the contract event 0xc609d4d70061ca85bd518176467e44ee689191f8654e28c186e4889816071f9f.
//
// Solidity: event Recover(address token)
func (_StaplePool *StaplePoolFilterer) FilterRecover(opts *bind.FilterOpts) (*StaplePoolRecoverIterator, error) {

	logs, sub, err := _StaplePool.contract.FilterLogs(opts, "Recover")
	if err != nil {
		return nil, err
	}
	return &StaplePoolRecoverIterator{contract: _StaplePool.contract, event: "Recover", logs: logs, sub: sub}, nil
}

// WatchRecover is a free log subscription operation binding the contract event 0xc609d4d70061ca85bd518176467e44ee689191f8654e28c186e4889816071f9f.
//
// Solidity: event Recover(address token)
func (_StaplePool *StaplePoolFilterer) WatchRecover(opts *bind.WatchOpts, sink chan<- *StaplePoolRecover) (event.Subscription, error) {

	logs, sub, err := _StaplePool.contract.WatchLogs(opts, "Recover")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaplePoolRecover)
				if err := _StaplePool.contract.UnpackLog(event, "Recover", log); err != nil {
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

// ParseRecover is a log parse operation binding the contract event 0xc609d4d70061ca85bd518176467e44ee689191f8654e28c186e4889816071f9f.
//
// Solidity: event Recover(address token)
func (_StaplePool *StaplePoolFilterer) ParseRecover(log types.Log) (*StaplePoolRecover, error) {
	event := new(StaplePoolRecover)
	if err := _StaplePool.contract.UnpackLog(event, "Recover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaplePoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StaplePool contract.
type StaplePoolTransferIterator struct {
	Event *StaplePoolTransfer // Event containing the contract specifics and raw log

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
func (it *StaplePoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaplePoolTransfer)
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
		it.Event = new(StaplePoolTransfer)
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
func (it *StaplePoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaplePoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaplePoolTransfer represents a Transfer event raised by the StaplePool contract.
type StaplePoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StaplePool *StaplePoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StaplePoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StaplePool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StaplePoolTransferIterator{contract: _StaplePool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StaplePool *StaplePoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StaplePoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StaplePool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaplePoolTransfer)
				if err := _StaplePool.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StaplePool *StaplePoolFilterer) ParseTransfer(log types.Log) (*StaplePoolTransfer, error) {
	event := new(StaplePoolTransfer)
	if err := _StaplePool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaplePoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the StaplePool contract.
type StaplePoolWithdrawIterator struct {
	Event *StaplePoolWithdraw // Event containing the contract specifics and raw log

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
func (it *StaplePoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaplePoolWithdraw)
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
		it.Event = new(StaplePoolWithdraw)
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
func (it *StaplePoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaplePoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaplePoolWithdraw represents a Withdraw event raised by the StaplePool contract.
type StaplePoolWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_StaplePool *StaplePoolFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*StaplePoolWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StaplePool.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &StaplePoolWithdrawIterator{contract: _StaplePool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_StaplePool *StaplePoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StaplePoolWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StaplePool.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaplePoolWithdraw)
				if err := _StaplePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_StaplePool *StaplePoolFilterer) ParseWithdraw(log types.Log) (*StaplePoolWithdraw, error) {
	event := new(StaplePoolWithdraw)
	if err := _StaplePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
