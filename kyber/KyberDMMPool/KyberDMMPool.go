// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KyberDMMPool

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

// KyberDMMPoolMetaData contains all meta data concerning the KyberDMMPool contract.
var KyberDMMPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ampBps\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTradeInfo\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_vReserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_vReserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint256\",\"name\":\"feeInPrecision\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVolumeTrendData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_shortEMA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_longEMA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_currentBlockVolume\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_lastTradeBlock\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_ampBps\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KyberDMMPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use KyberDMMPoolMetaData.ABI instead.
var KyberDMMPoolABI = KyberDMMPoolMetaData.ABI

// KyberDMMPool is an auto generated Go binding around an Ethereum contract.
type KyberDMMPool struct {
	KyberDMMPoolCaller     // Read-only binding to the contract
	KyberDMMPoolTransactor // Write-only binding to the contract
	KyberDMMPoolFilterer   // Log filterer for contract events
}

// KyberDMMPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberDMMPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberDMMPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberDMMPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberDMMPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberDMMPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberDMMPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberDMMPoolSession struct {
	Contract     *KyberDMMPool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KyberDMMPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberDMMPoolCallerSession struct {
	Contract *KyberDMMPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KyberDMMPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberDMMPoolTransactorSession struct {
	Contract     *KyberDMMPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KyberDMMPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberDMMPoolRaw struct {
	Contract *KyberDMMPool // Generic contract binding to access the raw methods on
}

// KyberDMMPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberDMMPoolCallerRaw struct {
	Contract *KyberDMMPoolCaller // Generic read-only contract binding to access the raw methods on
}

// KyberDMMPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberDMMPoolTransactorRaw struct {
	Contract *KyberDMMPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberDMMPool creates a new instance of KyberDMMPool, bound to a specific deployed contract.
func NewKyberDMMPool(address common.Address, backend bind.ContractBackend) (*KyberDMMPool, error) {
	contract, err := bindKyberDMMPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KyberDMMPool{KyberDMMPoolCaller: KyberDMMPoolCaller{contract: contract}, KyberDMMPoolTransactor: KyberDMMPoolTransactor{contract: contract}, KyberDMMPoolFilterer: KyberDMMPoolFilterer{contract: contract}}, nil
}

// NewKyberDMMPoolCaller creates a new read-only instance of KyberDMMPool, bound to a specific deployed contract.
func NewKyberDMMPoolCaller(address common.Address, caller bind.ContractCaller) (*KyberDMMPoolCaller, error) {
	contract, err := bindKyberDMMPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberDMMPoolCaller{contract: contract}, nil
}

// NewKyberDMMPoolTransactor creates a new write-only instance of KyberDMMPool, bound to a specific deployed contract.
func NewKyberDMMPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberDMMPoolTransactor, error) {
	contract, err := bindKyberDMMPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberDMMPoolTransactor{contract: contract}, nil
}

// NewKyberDMMPoolFilterer creates a new log filterer instance of KyberDMMPool, bound to a specific deployed contract.
func NewKyberDMMPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberDMMPoolFilterer, error) {
	contract, err := bindKyberDMMPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberDMMPoolFilterer{contract: contract}, nil
}

// bindKyberDMMPool binds a generic wrapper to an already deployed contract.
func bindKyberDMMPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberDMMPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberDMMPool *KyberDMMPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberDMMPool.Contract.KyberDMMPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberDMMPool *KyberDMMPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.KyberDMMPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberDMMPool *KyberDMMPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.KyberDMMPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberDMMPool *KyberDMMPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberDMMPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberDMMPool *KyberDMMPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberDMMPool *KyberDMMPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _KyberDMMPool.Contract.MINIMUMLIQUIDITY(&_KyberDMMPool.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _KyberDMMPool.Contract.MINIMUMLIQUIDITY(&_KyberDMMPool.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_KyberDMMPool *KyberDMMPoolCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_KyberDMMPool *KyberDMMPoolSession) PERMITTYPEHASH() ([32]byte, error) {
	return _KyberDMMPool.Contract.PERMITTYPEHASH(&_KyberDMMPool.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_KyberDMMPool *KyberDMMPoolCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _KyberDMMPool.Contract.PERMITTYPEHASH(&_KyberDMMPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KyberDMMPool.Contract.Allowance(&_KyberDMMPool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KyberDMMPool.Contract.Allowance(&_KyberDMMPool.CallOpts, owner, spender)
}

// AmpBps is a free data retrieval call binding the contract method 0x49386b16.
//
// Solidity: function ampBps() view returns(uint32)
func (_KyberDMMPool *KyberDMMPoolCaller) AmpBps(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "ampBps")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AmpBps is a free data retrieval call binding the contract method 0x49386b16.
//
// Solidity: function ampBps() view returns(uint32)
func (_KyberDMMPool *KyberDMMPoolSession) AmpBps() (uint32, error) {
	return _KyberDMMPool.Contract.AmpBps(&_KyberDMMPool.CallOpts)
}

// AmpBps is a free data retrieval call binding the contract method 0x49386b16.
//
// Solidity: function ampBps() view returns(uint32)
func (_KyberDMMPool *KyberDMMPoolCallerSession) AmpBps() (uint32, error) {
	return _KyberDMMPool.Contract.AmpBps(&_KyberDMMPool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KyberDMMPool.Contract.BalanceOf(&_KyberDMMPool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KyberDMMPool.Contract.BalanceOf(&_KyberDMMPool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KyberDMMPool *KyberDMMPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KyberDMMPool *KyberDMMPoolSession) Decimals() (uint8, error) {
	return _KyberDMMPool.Contract.Decimals(&_KyberDMMPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KyberDMMPool *KyberDMMPoolCallerSession) Decimals() (uint8, error) {
	return _KyberDMMPool.Contract.Decimals(&_KyberDMMPool.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_KyberDMMPool *KyberDMMPoolCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_KyberDMMPool *KyberDMMPoolSession) DomainSeparator() ([32]byte, error) {
	return _KyberDMMPool.Contract.DomainSeparator(&_KyberDMMPool.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_KyberDMMPool *KyberDMMPoolCallerSession) DomainSeparator() ([32]byte, error) {
	return _KyberDMMPool.Contract.DomainSeparator(&_KyberDMMPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KyberDMMPool *KyberDMMPoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KyberDMMPool *KyberDMMPoolSession) Factory() (common.Address, error) {
	return _KyberDMMPool.Contract.Factory(&_KyberDMMPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KyberDMMPool *KyberDMMPoolCallerSession) Factory() (common.Address, error) {
	return _KyberDMMPool.Contract.Factory(&_KyberDMMPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1)
func (_KyberDMMPool *KyberDMMPoolCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0 *big.Int
		Reserve1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1)
func (_KyberDMMPool *KyberDMMPoolSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _KyberDMMPool.Contract.GetReserves(&_KyberDMMPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1)
func (_KyberDMMPool *KyberDMMPoolCallerSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _KyberDMMPool.Contract.GetReserves(&_KyberDMMPool.CallOpts)
}

// GetTradeInfo is a free data retrieval call binding the contract method 0xd6694027.
//
// Solidity: function getTradeInfo() view returns(uint112 _reserve0, uint112 _reserve1, uint112 _vReserve0, uint112 _vReserve1, uint256 feeInPrecision)
func (_KyberDMMPool *KyberDMMPoolCaller) GetTradeInfo(opts *bind.CallOpts) (struct {
	Reserve0       *big.Int
	Reserve1       *big.Int
	VReserve0      *big.Int
	VReserve1      *big.Int
	FeeInPrecision *big.Int
}, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "getTradeInfo")

	outstruct := new(struct {
		Reserve0       *big.Int
		Reserve1       *big.Int
		VReserve0      *big.Int
		VReserve1      *big.Int
		FeeInPrecision *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VReserve0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.VReserve1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FeeInPrecision = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTradeInfo is a free data retrieval call binding the contract method 0xd6694027.
//
// Solidity: function getTradeInfo() view returns(uint112 _reserve0, uint112 _reserve1, uint112 _vReserve0, uint112 _vReserve1, uint256 feeInPrecision)
func (_KyberDMMPool *KyberDMMPoolSession) GetTradeInfo() (struct {
	Reserve0       *big.Int
	Reserve1       *big.Int
	VReserve0      *big.Int
	VReserve1      *big.Int
	FeeInPrecision *big.Int
}, error) {
	return _KyberDMMPool.Contract.GetTradeInfo(&_KyberDMMPool.CallOpts)
}

// GetTradeInfo is a free data retrieval call binding the contract method 0xd6694027.
//
// Solidity: function getTradeInfo() view returns(uint112 _reserve0, uint112 _reserve1, uint112 _vReserve0, uint112 _vReserve1, uint256 feeInPrecision)
func (_KyberDMMPool *KyberDMMPoolCallerSession) GetTradeInfo() (struct {
	Reserve0       *big.Int
	Reserve1       *big.Int
	VReserve0      *big.Int
	VReserve1      *big.Int
	FeeInPrecision *big.Int
}, error) {
	return _KyberDMMPool.Contract.GetTradeInfo(&_KyberDMMPool.CallOpts)
}

// GetVolumeTrendData is a free data retrieval call binding the contract method 0x0d94d50b.
//
// Solidity: function getVolumeTrendData() view returns(uint128 _shortEMA, uint128 _longEMA, uint128 _currentBlockVolume, uint128 _lastTradeBlock)
func (_KyberDMMPool *KyberDMMPoolCaller) GetVolumeTrendData(opts *bind.CallOpts) (struct {
	ShortEMA           *big.Int
	LongEMA            *big.Int
	CurrentBlockVolume *big.Int
	LastTradeBlock     *big.Int
}, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "getVolumeTrendData")

	outstruct := new(struct {
		ShortEMA           *big.Int
		LongEMA            *big.Int
		CurrentBlockVolume *big.Int
		LastTradeBlock     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ShortEMA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LongEMA = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentBlockVolume = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastTradeBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVolumeTrendData is a free data retrieval call binding the contract method 0x0d94d50b.
//
// Solidity: function getVolumeTrendData() view returns(uint128 _shortEMA, uint128 _longEMA, uint128 _currentBlockVolume, uint128 _lastTradeBlock)
func (_KyberDMMPool *KyberDMMPoolSession) GetVolumeTrendData() (struct {
	ShortEMA           *big.Int
	LongEMA            *big.Int
	CurrentBlockVolume *big.Int
	LastTradeBlock     *big.Int
}, error) {
	return _KyberDMMPool.Contract.GetVolumeTrendData(&_KyberDMMPool.CallOpts)
}

// GetVolumeTrendData is a free data retrieval call binding the contract method 0x0d94d50b.
//
// Solidity: function getVolumeTrendData() view returns(uint128 _shortEMA, uint128 _longEMA, uint128 _currentBlockVolume, uint128 _lastTradeBlock)
func (_KyberDMMPool *KyberDMMPoolCallerSession) GetVolumeTrendData() (struct {
	ShortEMA           *big.Int
	LongEMA            *big.Int
	CurrentBlockVolume *big.Int
	LastTradeBlock     *big.Int
}, error) {
	return _KyberDMMPool.Contract.GetVolumeTrendData(&_KyberDMMPool.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolSession) KLast() (*big.Int, error) {
	return _KyberDMMPool.Contract.KLast(&_KyberDMMPool.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCallerSession) KLast() (*big.Int, error) {
	return _KyberDMMPool.Contract.KLast(&_KyberDMMPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KyberDMMPool *KyberDMMPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KyberDMMPool *KyberDMMPoolSession) Name() (string, error) {
	return _KyberDMMPool.Contract.Name(&_KyberDMMPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KyberDMMPool *KyberDMMPoolCallerSession) Name() (string, error) {
	return _KyberDMMPool.Contract.Name(&_KyberDMMPool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _KyberDMMPool.Contract.Nonces(&_KyberDMMPool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _KyberDMMPool.Contract.Nonces(&_KyberDMMPool.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KyberDMMPool *KyberDMMPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KyberDMMPool *KyberDMMPoolSession) Symbol() (string, error) {
	return _KyberDMMPool.Contract.Symbol(&_KyberDMMPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KyberDMMPool *KyberDMMPoolCallerSession) Symbol() (string, error) {
	return _KyberDMMPool.Contract.Symbol(&_KyberDMMPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KyberDMMPool *KyberDMMPoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KyberDMMPool *KyberDMMPoolSession) Token0() (common.Address, error) {
	return _KyberDMMPool.Contract.Token0(&_KyberDMMPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KyberDMMPool *KyberDMMPoolCallerSession) Token0() (common.Address, error) {
	return _KyberDMMPool.Contract.Token0(&_KyberDMMPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KyberDMMPool *KyberDMMPoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KyberDMMPool *KyberDMMPoolSession) Token1() (common.Address, error) {
	return _KyberDMMPool.Contract.Token1(&_KyberDMMPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KyberDMMPool *KyberDMMPoolCallerSession) Token1() (common.Address, error) {
	return _KyberDMMPool.Contract.Token1(&_KyberDMMPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KyberDMMPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolSession) TotalSupply() (*big.Int, error) {
	return _KyberDMMPool.Contract.TotalSupply(&_KyberDMMPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KyberDMMPool *KyberDMMPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _KyberDMMPool.Contract.TotalSupply(&_KyberDMMPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KyberDMMPool *KyberDMMPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KyberDMMPool *KyberDMMPoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Approve(&_KyberDMMPool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KyberDMMPool *KyberDMMPoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Approve(&_KyberDMMPool.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_KyberDMMPool *KyberDMMPoolTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_KyberDMMPool *KyberDMMPoolSession) Burn(to common.Address) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Burn(&_KyberDMMPool.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_KyberDMMPool *KyberDMMPoolTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Burn(&_KyberDMMPool.TransactOpts, to)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KyberDMMPool *KyberDMMPoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KyberDMMPool *KyberDMMPoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.DecreaseAllowance(&_KyberDMMPool.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KyberDMMPool *KyberDMMPoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.DecreaseAllowance(&_KyberDMMPool.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KyberDMMPool *KyberDMMPoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KyberDMMPool *KyberDMMPoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.IncreaseAllowance(&_KyberDMMPool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KyberDMMPool *KyberDMMPoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.IncreaseAllowance(&_KyberDMMPool.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecf2b22.
//
// Solidity: function initialize(address _token0, address _token1, uint32 _ampBps) returns()
func (_KyberDMMPool *KyberDMMPoolTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address, _ampBps uint32) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "initialize", _token0, _token1, _ampBps)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecf2b22.
//
// Solidity: function initialize(address _token0, address _token1, uint32 _ampBps) returns()
func (_KyberDMMPool *KyberDMMPoolSession) Initialize(_token0 common.Address, _token1 common.Address, _ampBps uint32) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Initialize(&_KyberDMMPool.TransactOpts, _token0, _token1, _ampBps)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecf2b22.
//
// Solidity: function initialize(address _token0, address _token1, uint32 _ampBps) returns()
func (_KyberDMMPool *KyberDMMPoolTransactorSession) Initialize(_token0 common.Address, _token1 common.Address, _ampBps uint32) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Initialize(&_KyberDMMPool.TransactOpts, _token0, _token1, _ampBps)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_KyberDMMPool *KyberDMMPoolTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_KyberDMMPool *KyberDMMPoolSession) Mint(to common.Address) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Mint(&_KyberDMMPool.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_KyberDMMPool *KyberDMMPoolTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Mint(&_KyberDMMPool.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_KyberDMMPool *KyberDMMPoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_KyberDMMPool *KyberDMMPoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Permit(&_KyberDMMPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_KyberDMMPool *KyberDMMPoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Permit(&_KyberDMMPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_KyberDMMPool *KyberDMMPoolTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_KyberDMMPool *KyberDMMPoolSession) Skim(to common.Address) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Skim(&_KyberDMMPool.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_KyberDMMPool *KyberDMMPoolTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Skim(&_KyberDMMPool.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes callbackData) returns()
func (_KyberDMMPool *KyberDMMPoolTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, callbackData []byte) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "swap", amount0Out, amount1Out, to, callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes callbackData) returns()
func (_KyberDMMPool *KyberDMMPoolSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, callbackData []byte) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Swap(&_KyberDMMPool.TransactOpts, amount0Out, amount1Out, to, callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes callbackData) returns()
func (_KyberDMMPool *KyberDMMPoolTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, callbackData []byte) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Swap(&_KyberDMMPool.TransactOpts, amount0Out, amount1Out, to, callbackData)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_KyberDMMPool *KyberDMMPoolTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_KyberDMMPool *KyberDMMPoolSession) Sync() (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Sync(&_KyberDMMPool.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_KyberDMMPool *KyberDMMPoolTransactorSession) Sync() (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Sync(&_KyberDMMPool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KyberDMMPool *KyberDMMPoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KyberDMMPool *KyberDMMPoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Transfer(&_KyberDMMPool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KyberDMMPool *KyberDMMPoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.Transfer(&_KyberDMMPool.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KyberDMMPool *KyberDMMPoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KyberDMMPool *KyberDMMPoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.TransferFrom(&_KyberDMMPool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KyberDMMPool *KyberDMMPoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KyberDMMPool.Contract.TransferFrom(&_KyberDMMPool.TransactOpts, sender, recipient, amount)
}
