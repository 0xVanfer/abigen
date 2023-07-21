// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lidowstETH

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

// LidowstETHMetaData contains all meta data concerning the LidowstETH contract.
var LidowstETHMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wstETHAmount\",\"type\":\"uint256\"}],\"name\":\"getStETHByWstETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stETHAmount\",\"type\":\"uint256\"}],\"name\":\"getWstETHByStETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stEthPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensPerStEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wstETHAmount\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stETHAmount\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LidowstETHABI is the input ABI used to generate the binding from.
// Deprecated: Use LidowstETHMetaData.ABI instead.
var LidowstETHABI = LidowstETHMetaData.ABI

// LidowstETH is an auto generated Go binding around an Ethereum contract.
type LidowstETH struct {
	LidowstETHCaller     // Read-only binding to the contract
	LidowstETHTransactor // Write-only binding to the contract
	LidowstETHFilterer   // Log filterer for contract events
}

// LidowstETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type LidowstETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidowstETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LidowstETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidowstETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LidowstETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidowstETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LidowstETHSession struct {
	Contract     *LidowstETH       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LidowstETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LidowstETHCallerSession struct {
	Contract *LidowstETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LidowstETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LidowstETHTransactorSession struct {
	Contract     *LidowstETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LidowstETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type LidowstETHRaw struct {
	Contract *LidowstETH // Generic contract binding to access the raw methods on
}

// LidowstETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LidowstETHCallerRaw struct {
	Contract *LidowstETHCaller // Generic read-only contract binding to access the raw methods on
}

// LidowstETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LidowstETHTransactorRaw struct {
	Contract *LidowstETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLidowstETH creates a new instance of LidowstETH, bound to a specific deployed contract.
func NewLidowstETH(address common.Address, backend bind.ContractBackend) (*LidowstETH, error) {
	contract, err := bindLidowstETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LidowstETH{LidowstETHCaller: LidowstETHCaller{contract: contract}, LidowstETHTransactor: LidowstETHTransactor{contract: contract}, LidowstETHFilterer: LidowstETHFilterer{contract: contract}}, nil
}

// NewLidowstETHCaller creates a new read-only instance of LidowstETH, bound to a specific deployed contract.
func NewLidowstETHCaller(address common.Address, caller bind.ContractCaller) (*LidowstETHCaller, error) {
	contract, err := bindLidowstETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LidowstETHCaller{contract: contract}, nil
}

// NewLidowstETHTransactor creates a new write-only instance of LidowstETH, bound to a specific deployed contract.
func NewLidowstETHTransactor(address common.Address, transactor bind.ContractTransactor) (*LidowstETHTransactor, error) {
	contract, err := bindLidowstETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LidowstETHTransactor{contract: contract}, nil
}

// NewLidowstETHFilterer creates a new log filterer instance of LidowstETH, bound to a specific deployed contract.
func NewLidowstETHFilterer(address common.Address, filterer bind.ContractFilterer) (*LidowstETHFilterer, error) {
	contract, err := bindLidowstETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LidowstETHFilterer{contract: contract}, nil
}

// bindLidowstETH binds a generic wrapper to an already deployed contract.
func bindLidowstETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LidowstETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LidowstETH *LidowstETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LidowstETH.Contract.LidowstETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LidowstETH *LidowstETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidowstETH.Contract.LidowstETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LidowstETH *LidowstETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LidowstETH.Contract.LidowstETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LidowstETH *LidowstETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LidowstETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LidowstETH *LidowstETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidowstETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LidowstETH *LidowstETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LidowstETH.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LidowstETH *LidowstETHCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LidowstETH *LidowstETHSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LidowstETH.Contract.DOMAINSEPARATOR(&_LidowstETH.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LidowstETH *LidowstETHCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LidowstETH.Contract.DOMAINSEPARATOR(&_LidowstETH.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LidowstETH *LidowstETHCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LidowstETH *LidowstETHSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LidowstETH.Contract.Allowance(&_LidowstETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LidowstETH *LidowstETHCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LidowstETH.Contract.Allowance(&_LidowstETH.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LidowstETH *LidowstETHCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LidowstETH *LidowstETHSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LidowstETH.Contract.BalanceOf(&_LidowstETH.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LidowstETH *LidowstETHCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LidowstETH.Contract.BalanceOf(&_LidowstETH.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LidowstETH *LidowstETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LidowstETH *LidowstETHSession) Decimals() (uint8, error) {
	return _LidowstETH.Contract.Decimals(&_LidowstETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LidowstETH *LidowstETHCallerSession) Decimals() (uint8, error) {
	return _LidowstETH.Contract.Decimals(&_LidowstETH.CallOpts)
}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_LidowstETH *LidowstETHCaller) GetStETHByWstETH(opts *bind.CallOpts, _wstETHAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "getStETHByWstETH", _wstETHAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_LidowstETH *LidowstETHSession) GetStETHByWstETH(_wstETHAmount *big.Int) (*big.Int, error) {
	return _LidowstETH.Contract.GetStETHByWstETH(&_LidowstETH.CallOpts, _wstETHAmount)
}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_LidowstETH *LidowstETHCallerSession) GetStETHByWstETH(_wstETHAmount *big.Int) (*big.Int, error) {
	return _LidowstETH.Contract.GetStETHByWstETH(&_LidowstETH.CallOpts, _wstETHAmount)
}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_LidowstETH *LidowstETHCaller) GetWstETHByStETH(opts *bind.CallOpts, _stETHAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "getWstETHByStETH", _stETHAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_LidowstETH *LidowstETHSession) GetWstETHByStETH(_stETHAmount *big.Int) (*big.Int, error) {
	return _LidowstETH.Contract.GetWstETHByStETH(&_LidowstETH.CallOpts, _stETHAmount)
}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_LidowstETH *LidowstETHCallerSession) GetWstETHByStETH(_stETHAmount *big.Int) (*big.Int, error) {
	return _LidowstETH.Contract.GetWstETHByStETH(&_LidowstETH.CallOpts, _stETHAmount)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LidowstETH *LidowstETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LidowstETH *LidowstETHSession) Name() (string, error) {
	return _LidowstETH.Contract.Name(&_LidowstETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LidowstETH *LidowstETHCallerSession) Name() (string, error) {
	return _LidowstETH.Contract.Name(&_LidowstETH.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_LidowstETH *LidowstETHCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_LidowstETH *LidowstETHSession) Nonces(owner common.Address) (*big.Int, error) {
	return _LidowstETH.Contract.Nonces(&_LidowstETH.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_LidowstETH *LidowstETHCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _LidowstETH.Contract.Nonces(&_LidowstETH.CallOpts, owner)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_LidowstETH *LidowstETHCaller) StETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "stETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_LidowstETH *LidowstETHSession) StETH() (common.Address, error) {
	return _LidowstETH.Contract.StETH(&_LidowstETH.CallOpts)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_LidowstETH *LidowstETHCallerSession) StETH() (common.Address, error) {
	return _LidowstETH.Contract.StETH(&_LidowstETH.CallOpts)
}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_LidowstETH *LidowstETHCaller) StEthPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "stEthPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_LidowstETH *LidowstETHSession) StEthPerToken() (*big.Int, error) {
	return _LidowstETH.Contract.StEthPerToken(&_LidowstETH.CallOpts)
}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_LidowstETH *LidowstETHCallerSession) StEthPerToken() (*big.Int, error) {
	return _LidowstETH.Contract.StEthPerToken(&_LidowstETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LidowstETH *LidowstETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LidowstETH *LidowstETHSession) Symbol() (string, error) {
	return _LidowstETH.Contract.Symbol(&_LidowstETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LidowstETH *LidowstETHCallerSession) Symbol() (string, error) {
	return _LidowstETH.Contract.Symbol(&_LidowstETH.CallOpts)
}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_LidowstETH *LidowstETHCaller) TokensPerStEth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "tokensPerStEth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_LidowstETH *LidowstETHSession) TokensPerStEth() (*big.Int, error) {
	return _LidowstETH.Contract.TokensPerStEth(&_LidowstETH.CallOpts)
}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_LidowstETH *LidowstETHCallerSession) TokensPerStEth() (*big.Int, error) {
	return _LidowstETH.Contract.TokensPerStEth(&_LidowstETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidowstETH *LidowstETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidowstETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidowstETH *LidowstETHSession) TotalSupply() (*big.Int, error) {
	return _LidowstETH.Contract.TotalSupply(&_LidowstETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidowstETH *LidowstETHCallerSession) TotalSupply() (*big.Int, error) {
	return _LidowstETH.Contract.TotalSupply(&_LidowstETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LidowstETH *LidowstETHTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LidowstETH *LidowstETHSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.Approve(&_LidowstETH.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LidowstETH *LidowstETHTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.Approve(&_LidowstETH.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LidowstETH *LidowstETHTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidowstETH.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LidowstETH *LidowstETHSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.DecreaseAllowance(&_LidowstETH.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LidowstETH *LidowstETHTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.DecreaseAllowance(&_LidowstETH.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LidowstETH *LidowstETHTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LidowstETH.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LidowstETH *LidowstETHSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.IncreaseAllowance(&_LidowstETH.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LidowstETH *LidowstETHTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.IncreaseAllowance(&_LidowstETH.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LidowstETH *LidowstETHTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LidowstETH.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LidowstETH *LidowstETHSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LidowstETH.Contract.Permit(&_LidowstETH.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LidowstETH *LidowstETHTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LidowstETH.Contract.Permit(&_LidowstETH.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LidowstETH *LidowstETHTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LidowstETH *LidowstETHSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.Transfer(&_LidowstETH.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LidowstETH *LidowstETHTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.Transfer(&_LidowstETH.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LidowstETH *LidowstETHTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LidowstETH *LidowstETHSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.TransferFrom(&_LidowstETH.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LidowstETH *LidowstETHTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.TransferFrom(&_LidowstETH.TransactOpts, sender, recipient, amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_LidowstETH *LidowstETHTransactor) Unwrap(opts *bind.TransactOpts, _wstETHAmount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.contract.Transact(opts, "unwrap", _wstETHAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_LidowstETH *LidowstETHSession) Unwrap(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.Unwrap(&_LidowstETH.TransactOpts, _wstETHAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_LidowstETH *LidowstETHTransactorSession) Unwrap(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.Unwrap(&_LidowstETH.TransactOpts, _wstETHAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_LidowstETH *LidowstETHTransactor) Wrap(opts *bind.TransactOpts, _stETHAmount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.contract.Transact(opts, "wrap", _stETHAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_LidowstETH *LidowstETHSession) Wrap(_stETHAmount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.Wrap(&_LidowstETH.TransactOpts, _stETHAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_LidowstETH *LidowstETHTransactorSession) Wrap(_stETHAmount *big.Int) (*types.Transaction, error) {
	return _LidowstETH.Contract.Wrap(&_LidowstETH.TransactOpts, _stETHAmount)
}
