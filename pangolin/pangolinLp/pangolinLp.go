// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pangolinLp

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

// PangolinLpMetaData contains all meta data concerning the PangolinLp contract.
var PangolinLpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PangolinLpABI is the input ABI used to generate the binding from.
// Deprecated: Use PangolinLpMetaData.ABI instead.
var PangolinLpABI = PangolinLpMetaData.ABI

// PangolinLp is an auto generated Go binding around an Ethereum contract.
type PangolinLp struct {
	PangolinLpCaller     // Read-only binding to the contract
	PangolinLpTransactor // Write-only binding to the contract
	PangolinLpFilterer   // Log filterer for contract events
}

// PangolinLpCaller is an auto generated read-only Go binding around an Ethereum contract.
type PangolinLpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PangolinLpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PangolinLpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PangolinLpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PangolinLpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PangolinLpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PangolinLpSession struct {
	Contract     *PangolinLp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PangolinLpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PangolinLpCallerSession struct {
	Contract *PangolinLpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PangolinLpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PangolinLpTransactorSession struct {
	Contract     *PangolinLpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PangolinLpRaw is an auto generated low-level Go binding around an Ethereum contract.
type PangolinLpRaw struct {
	Contract *PangolinLp // Generic contract binding to access the raw methods on
}

// PangolinLpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PangolinLpCallerRaw struct {
	Contract *PangolinLpCaller // Generic read-only contract binding to access the raw methods on
}

// PangolinLpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PangolinLpTransactorRaw struct {
	Contract *PangolinLpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPangolinLp creates a new instance of PangolinLp, bound to a specific deployed contract.
func NewPangolinLp(address common.Address, backend bind.ContractBackend) (*PangolinLp, error) {
	contract, err := bindPangolinLp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PangolinLp{PangolinLpCaller: PangolinLpCaller{contract: contract}, PangolinLpTransactor: PangolinLpTransactor{contract: contract}, PangolinLpFilterer: PangolinLpFilterer{contract: contract}}, nil
}

// NewPangolinLpCaller creates a new read-only instance of PangolinLp, bound to a specific deployed contract.
func NewPangolinLpCaller(address common.Address, caller bind.ContractCaller) (*PangolinLpCaller, error) {
	contract, err := bindPangolinLp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PangolinLpCaller{contract: contract}, nil
}

// NewPangolinLpTransactor creates a new write-only instance of PangolinLp, bound to a specific deployed contract.
func NewPangolinLpTransactor(address common.Address, transactor bind.ContractTransactor) (*PangolinLpTransactor, error) {
	contract, err := bindPangolinLp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PangolinLpTransactor{contract: contract}, nil
}

// NewPangolinLpFilterer creates a new log filterer instance of PangolinLp, bound to a specific deployed contract.
func NewPangolinLpFilterer(address common.Address, filterer bind.ContractFilterer) (*PangolinLpFilterer, error) {
	contract, err := bindPangolinLp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PangolinLpFilterer{contract: contract}, nil
}

// bindPangolinLp binds a generic wrapper to an already deployed contract.
func bindPangolinLp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PangolinLpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PangolinLp *PangolinLpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PangolinLp.Contract.PangolinLpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PangolinLp *PangolinLpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PangolinLp.Contract.PangolinLpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PangolinLp *PangolinLpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PangolinLp.Contract.PangolinLpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PangolinLp *PangolinLpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PangolinLp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PangolinLp *PangolinLpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PangolinLp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PangolinLp *PangolinLpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PangolinLp.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PangolinLp *PangolinLpCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PangolinLp *PangolinLpSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _PangolinLp.Contract.DOMAINSEPARATOR(&_PangolinLp.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PangolinLp *PangolinLpCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _PangolinLp.Contract.DOMAINSEPARATOR(&_PangolinLp.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_PangolinLp *PangolinLpCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_PangolinLp *PangolinLpSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _PangolinLp.Contract.MINIMUMLIQUIDITY(&_PangolinLp.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_PangolinLp *PangolinLpCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _PangolinLp.Contract.MINIMUMLIQUIDITY(&_PangolinLp.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_PangolinLp *PangolinLpCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_PangolinLp *PangolinLpSession) PERMITTYPEHASH() ([32]byte, error) {
	return _PangolinLp.Contract.PERMITTYPEHASH(&_PangolinLp.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_PangolinLp *PangolinLpCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _PangolinLp.Contract.PERMITTYPEHASH(&_PangolinLp.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_PangolinLp *PangolinLpCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_PangolinLp *PangolinLpSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PangolinLp.Contract.Allowance(&_PangolinLp.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_PangolinLp *PangolinLpCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PangolinLp.Contract.Allowance(&_PangolinLp.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_PangolinLp *PangolinLpCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_PangolinLp *PangolinLpSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _PangolinLp.Contract.BalanceOf(&_PangolinLp.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_PangolinLp *PangolinLpCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _PangolinLp.Contract.BalanceOf(&_PangolinLp.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PangolinLp *PangolinLpCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PangolinLp *PangolinLpSession) Decimals() (uint8, error) {
	return _PangolinLp.Contract.Decimals(&_PangolinLp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PangolinLp *PangolinLpCallerSession) Decimals() (uint8, error) {
	return _PangolinLp.Contract.Decimals(&_PangolinLp.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PangolinLp *PangolinLpCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PangolinLp *PangolinLpSession) Factory() (common.Address, error) {
	return _PangolinLp.Contract.Factory(&_PangolinLp.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PangolinLp *PangolinLpCallerSession) Factory() (common.Address, error) {
	return _PangolinLp.Contract.Factory(&_PangolinLp.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_PangolinLp *PangolinLpCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_PangolinLp *PangolinLpSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _PangolinLp.Contract.GetReserves(&_PangolinLp.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_PangolinLp *PangolinLpCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _PangolinLp.Contract.GetReserves(&_PangolinLp.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_PangolinLp *PangolinLpCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_PangolinLp *PangolinLpSession) KLast() (*big.Int, error) {
	return _PangolinLp.Contract.KLast(&_PangolinLp.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_PangolinLp *PangolinLpCallerSession) KLast() (*big.Int, error) {
	return _PangolinLp.Contract.KLast(&_PangolinLp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PangolinLp *PangolinLpCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PangolinLp *PangolinLpSession) Name() (string, error) {
	return _PangolinLp.Contract.Name(&_PangolinLp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PangolinLp *PangolinLpCallerSession) Name() (string, error) {
	return _PangolinLp.Contract.Name(&_PangolinLp.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PangolinLp *PangolinLpCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PangolinLp *PangolinLpSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PangolinLp.Contract.Nonces(&_PangolinLp.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PangolinLp *PangolinLpCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PangolinLp.Contract.Nonces(&_PangolinLp.CallOpts, arg0)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_PangolinLp *PangolinLpCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_PangolinLp *PangolinLpSession) Price0CumulativeLast() (*big.Int, error) {
	return _PangolinLp.Contract.Price0CumulativeLast(&_PangolinLp.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_PangolinLp *PangolinLpCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _PangolinLp.Contract.Price0CumulativeLast(&_PangolinLp.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_PangolinLp *PangolinLpCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_PangolinLp *PangolinLpSession) Price1CumulativeLast() (*big.Int, error) {
	return _PangolinLp.Contract.Price1CumulativeLast(&_PangolinLp.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_PangolinLp *PangolinLpCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _PangolinLp.Contract.Price1CumulativeLast(&_PangolinLp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PangolinLp *PangolinLpCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PangolinLp *PangolinLpSession) Symbol() (string, error) {
	return _PangolinLp.Contract.Symbol(&_PangolinLp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PangolinLp *PangolinLpCallerSession) Symbol() (string, error) {
	return _PangolinLp.Contract.Symbol(&_PangolinLp.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PangolinLp *PangolinLpCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PangolinLp *PangolinLpSession) Token0() (common.Address, error) {
	return _PangolinLp.Contract.Token0(&_PangolinLp.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PangolinLp *PangolinLpCallerSession) Token0() (common.Address, error) {
	return _PangolinLp.Contract.Token0(&_PangolinLp.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PangolinLp *PangolinLpCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PangolinLp *PangolinLpSession) Token1() (common.Address, error) {
	return _PangolinLp.Contract.Token1(&_PangolinLp.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PangolinLp *PangolinLpCallerSession) Token1() (common.Address, error) {
	return _PangolinLp.Contract.Token1(&_PangolinLp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PangolinLp *PangolinLpCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PangolinLp.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PangolinLp *PangolinLpSession) TotalSupply() (*big.Int, error) {
	return _PangolinLp.Contract.TotalSupply(&_PangolinLp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PangolinLp *PangolinLpCallerSession) TotalSupply() (*big.Int, error) {
	return _PangolinLp.Contract.TotalSupply(&_PangolinLp.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PangolinLp *PangolinLpTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PangolinLp.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PangolinLp *PangolinLpSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PangolinLp.Contract.Approve(&_PangolinLp.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PangolinLp *PangolinLpTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PangolinLp.Contract.Approve(&_PangolinLp.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_PangolinLp *PangolinLpTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _PangolinLp.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_PangolinLp *PangolinLpSession) Burn(to common.Address) (*types.Transaction, error) {
	return _PangolinLp.Contract.Burn(&_PangolinLp.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_PangolinLp *PangolinLpTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _PangolinLp.Contract.Burn(&_PangolinLp.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_PangolinLp *PangolinLpTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _PangolinLp.contract.Transact(opts, "initialize", _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_PangolinLp *PangolinLpSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _PangolinLp.Contract.Initialize(&_PangolinLp.TransactOpts, _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_PangolinLp *PangolinLpTransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _PangolinLp.Contract.Initialize(&_PangolinLp.TransactOpts, _token0, _token1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_PangolinLp *PangolinLpTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _PangolinLp.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_PangolinLp *PangolinLpSession) Mint(to common.Address) (*types.Transaction, error) {
	return _PangolinLp.Contract.Mint(&_PangolinLp.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_PangolinLp *PangolinLpTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _PangolinLp.Contract.Mint(&_PangolinLp.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PangolinLp *PangolinLpTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PangolinLp.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PangolinLp *PangolinLpSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PangolinLp.Contract.Permit(&_PangolinLp.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PangolinLp *PangolinLpTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PangolinLp.Contract.Permit(&_PangolinLp.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_PangolinLp *PangolinLpTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _PangolinLp.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_PangolinLp *PangolinLpSession) Skim(to common.Address) (*types.Transaction, error) {
	return _PangolinLp.Contract.Skim(&_PangolinLp.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_PangolinLp *PangolinLpTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _PangolinLp.Contract.Skim(&_PangolinLp.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_PangolinLp *PangolinLpTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _PangolinLp.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_PangolinLp *PangolinLpSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _PangolinLp.Contract.Swap(&_PangolinLp.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_PangolinLp *PangolinLpTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _PangolinLp.Contract.Swap(&_PangolinLp.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_PangolinLp *PangolinLpTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PangolinLp.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_PangolinLp *PangolinLpSession) Sync() (*types.Transaction, error) {
	return _PangolinLp.Contract.Sync(&_PangolinLp.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_PangolinLp *PangolinLpTransactorSession) Sync() (*types.Transaction, error) {
	return _PangolinLp.Contract.Sync(&_PangolinLp.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PangolinLp *PangolinLpTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PangolinLp.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PangolinLp *PangolinLpSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PangolinLp.Contract.Transfer(&_PangolinLp.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PangolinLp *PangolinLpTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PangolinLp.Contract.Transfer(&_PangolinLp.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PangolinLp *PangolinLpTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PangolinLp.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PangolinLp *PangolinLpSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PangolinLp.Contract.TransferFrom(&_PangolinLp.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PangolinLp *PangolinLpTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PangolinLp.Contract.TransferFrom(&_PangolinLp.TransactOpts, from, to, value)
}
