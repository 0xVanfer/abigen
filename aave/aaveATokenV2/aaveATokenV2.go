// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveATokenV2

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

// AaveATokenV2MetaData contains all meta data concerning the AaveATokenV2 contract.
var AaveATokenV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ATOKEN_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_REVISION\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_TREASURY_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING_ASSET_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiverOfUnderlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIncentivesController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getScaledUserBalanceAndSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"handleRepayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentivesController\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"aTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"aTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"aTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"mintToTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"scaledBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaledTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferOnLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferUnderlyingTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveATokenV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveATokenV2MetaData.ABI instead.
var AaveATokenV2ABI = AaveATokenV2MetaData.ABI

// AaveATokenV2 is an auto generated Go binding around an Ethereum contract.
type AaveATokenV2 struct {
	AaveATokenV2Caller     // Read-only binding to the contract
	AaveATokenV2Transactor // Write-only binding to the contract
	AaveATokenV2Filterer   // Log filterer for contract events
}

// AaveATokenV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveATokenV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveATokenV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveATokenV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveATokenV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveATokenV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveATokenV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveATokenV2Session struct {
	Contract     *AaveATokenV2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveATokenV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveATokenV2CallerSession struct {
	Contract *AaveATokenV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AaveATokenV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveATokenV2TransactorSession struct {
	Contract     *AaveATokenV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AaveATokenV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveATokenV2Raw struct {
	Contract *AaveATokenV2 // Generic contract binding to access the raw methods on
}

// AaveATokenV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveATokenV2CallerRaw struct {
	Contract *AaveATokenV2Caller // Generic read-only contract binding to access the raw methods on
}

// AaveATokenV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveATokenV2TransactorRaw struct {
	Contract *AaveATokenV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveATokenV2 creates a new instance of AaveATokenV2, bound to a specific deployed contract.
func NewAaveATokenV2(address common.Address, backend bind.ContractBackend) (*AaveATokenV2, error) {
	contract, err := bindAaveATokenV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveATokenV2{AaveATokenV2Caller: AaveATokenV2Caller{contract: contract}, AaveATokenV2Transactor: AaveATokenV2Transactor{contract: contract}, AaveATokenV2Filterer: AaveATokenV2Filterer{contract: contract}}, nil
}

// NewAaveATokenV2Caller creates a new read-only instance of AaveATokenV2, bound to a specific deployed contract.
func NewAaveATokenV2Caller(address common.Address, caller bind.ContractCaller) (*AaveATokenV2Caller, error) {
	contract, err := bindAaveATokenV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveATokenV2Caller{contract: contract}, nil
}

// NewAaveATokenV2Transactor creates a new write-only instance of AaveATokenV2, bound to a specific deployed contract.
func NewAaveATokenV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveATokenV2Transactor, error) {
	contract, err := bindAaveATokenV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveATokenV2Transactor{contract: contract}, nil
}

// NewAaveATokenV2Filterer creates a new log filterer instance of AaveATokenV2, bound to a specific deployed contract.
func NewAaveATokenV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveATokenV2Filterer, error) {
	contract, err := bindAaveATokenV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveATokenV2Filterer{contract: contract}, nil
}

// bindAaveATokenV2 binds a generic wrapper to an already deployed contract.
func bindAaveATokenV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveATokenV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveATokenV2 *AaveATokenV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveATokenV2.Contract.AaveATokenV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveATokenV2 *AaveATokenV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.AaveATokenV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveATokenV2 *AaveATokenV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.AaveATokenV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveATokenV2 *AaveATokenV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveATokenV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveATokenV2 *AaveATokenV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveATokenV2 *AaveATokenV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.contract.Transact(opts, method, params...)
}

// ATOKENREVISION is a free data retrieval call binding the contract method 0x0bd7ad3b.
//
// Solidity: function ATOKEN_REVISION() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Caller) ATOKENREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "ATOKEN_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ATOKENREVISION is a free data retrieval call binding the contract method 0x0bd7ad3b.
//
// Solidity: function ATOKEN_REVISION() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Session) ATOKENREVISION() (*big.Int, error) {
	return _AaveATokenV2.Contract.ATOKENREVISION(&_AaveATokenV2.CallOpts)
}

// ATOKENREVISION is a free data retrieval call binding the contract method 0x0bd7ad3b.
//
// Solidity: function ATOKEN_REVISION() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2CallerSession) ATOKENREVISION() (*big.Int, error) {
	return _AaveATokenV2.Contract.ATOKENREVISION(&_AaveATokenV2.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AaveATokenV2 *AaveATokenV2Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AaveATokenV2 *AaveATokenV2Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _AaveATokenV2.Contract.DOMAINSEPARATOR(&_AaveATokenV2.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AaveATokenV2 *AaveATokenV2CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _AaveATokenV2.Contract.DOMAINSEPARATOR(&_AaveATokenV2.CallOpts)
}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_AaveATokenV2 *AaveATokenV2Caller) EIP712REVISION(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "EIP712_REVISION")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_AaveATokenV2 *AaveATokenV2Session) EIP712REVISION() ([]byte, error) {
	return _AaveATokenV2.Contract.EIP712REVISION(&_AaveATokenV2.CallOpts)
}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_AaveATokenV2 *AaveATokenV2CallerSession) EIP712REVISION() ([]byte, error) {
	return _AaveATokenV2.Contract.EIP712REVISION(&_AaveATokenV2.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_AaveATokenV2 *AaveATokenV2Caller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_AaveATokenV2 *AaveATokenV2Session) PERMITTYPEHASH() ([32]byte, error) {
	return _AaveATokenV2.Contract.PERMITTYPEHASH(&_AaveATokenV2.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_AaveATokenV2 *AaveATokenV2CallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _AaveATokenV2.Contract.PERMITTYPEHASH(&_AaveATokenV2.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Caller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Session) POOL() (common.Address, error) {
	return _AaveATokenV2.Contract.POOL(&_AaveATokenV2.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveATokenV2 *AaveATokenV2CallerSession) POOL() (common.Address, error) {
	return _AaveATokenV2.Contract.POOL(&_AaveATokenV2.CallOpts)
}

// RESERVETREASURYADDRESS is a free data retrieval call binding the contract method 0xae167335.
//
// Solidity: function RESERVE_TREASURY_ADDRESS() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Caller) RESERVETREASURYADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "RESERVE_TREASURY_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RESERVETREASURYADDRESS is a free data retrieval call binding the contract method 0xae167335.
//
// Solidity: function RESERVE_TREASURY_ADDRESS() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Session) RESERVETREASURYADDRESS() (common.Address, error) {
	return _AaveATokenV2.Contract.RESERVETREASURYADDRESS(&_AaveATokenV2.CallOpts)
}

// RESERVETREASURYADDRESS is a free data retrieval call binding the contract method 0xae167335.
//
// Solidity: function RESERVE_TREASURY_ADDRESS() view returns(address)
func (_AaveATokenV2 *AaveATokenV2CallerSession) RESERVETREASURYADDRESS() (common.Address, error) {
	return _AaveATokenV2.Contract.RESERVETREASURYADDRESS(&_AaveATokenV2.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Caller) UNDERLYINGASSETADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "UNDERLYING_ASSET_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Session) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveATokenV2.Contract.UNDERLYINGASSETADDRESS(&_AaveATokenV2.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveATokenV2 *AaveATokenV2CallerSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveATokenV2.Contract.UNDERLYINGASSETADDRESS(&_AaveATokenV2.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "_nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.Nonces(&_AaveATokenV2.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.Nonces(&_AaveATokenV2.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.Allowance(&_AaveATokenV2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.Allowance(&_AaveATokenV2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Caller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Session) BalanceOf(user common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.BalanceOf(&_AaveATokenV2.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2CallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.BalanceOf(&_AaveATokenV2.CallOpts, user)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveATokenV2 *AaveATokenV2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveATokenV2 *AaveATokenV2Session) Decimals() (uint8, error) {
	return _AaveATokenV2.Contract.Decimals(&_AaveATokenV2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveATokenV2 *AaveATokenV2CallerSession) Decimals() (uint8, error) {
	return _AaveATokenV2.Contract.Decimals(&_AaveATokenV2.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Caller) GetIncentivesController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "getIncentivesController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveATokenV2 *AaveATokenV2Session) GetIncentivesController() (common.Address, error) {
	return _AaveATokenV2.Contract.GetIncentivesController(&_AaveATokenV2.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveATokenV2 *AaveATokenV2CallerSession) GetIncentivesController() (common.Address, error) {
	return _AaveATokenV2.Contract.GetIncentivesController(&_AaveATokenV2.CallOpts)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveATokenV2 *AaveATokenV2Caller) GetScaledUserBalanceAndSupply(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "getScaledUserBalanceAndSupply", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveATokenV2 *AaveATokenV2Session) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _AaveATokenV2.Contract.GetScaledUserBalanceAndSupply(&_AaveATokenV2.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveATokenV2 *AaveATokenV2CallerSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _AaveATokenV2.Contract.GetScaledUserBalanceAndSupply(&_AaveATokenV2.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveATokenV2 *AaveATokenV2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveATokenV2 *AaveATokenV2Session) Name() (string, error) {
	return _AaveATokenV2.Contract.Name(&_AaveATokenV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveATokenV2 *AaveATokenV2CallerSession) Name() (string, error) {
	return _AaveATokenV2.Contract.Name(&_AaveATokenV2.CallOpts)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Caller) ScaledBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "scaledBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Session) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.ScaledBalanceOf(&_AaveATokenV2.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2CallerSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveATokenV2.Contract.ScaledBalanceOf(&_AaveATokenV2.CallOpts, user)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Caller) ScaledTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "scaledTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Session) ScaledTotalSupply() (*big.Int, error) {
	return _AaveATokenV2.Contract.ScaledTotalSupply(&_AaveATokenV2.CallOpts)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2CallerSession) ScaledTotalSupply() (*big.Int, error) {
	return _AaveATokenV2.Contract.ScaledTotalSupply(&_AaveATokenV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveATokenV2 *AaveATokenV2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveATokenV2 *AaveATokenV2Session) Symbol() (string, error) {
	return _AaveATokenV2.Contract.Symbol(&_AaveATokenV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveATokenV2 *AaveATokenV2CallerSession) Symbol() (string, error) {
	return _AaveATokenV2.Contract.Symbol(&_AaveATokenV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveATokenV2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Session) TotalSupply() (*big.Int, error) {
	return _AaveATokenV2.Contract.TotalSupply(&_AaveATokenV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveATokenV2 *AaveATokenV2CallerSession) TotalSupply() (*big.Int, error) {
	return _AaveATokenV2.Contract.TotalSupply(&_AaveATokenV2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Approve(&_AaveATokenV2.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Approve(&_AaveATokenV2.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address user, address receiverOfUnderlying, uint256 amount, uint256 index) returns()
func (_AaveATokenV2 *AaveATokenV2Transactor) Burn(opts *bind.TransactOpts, user common.Address, receiverOfUnderlying common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "burn", user, receiverOfUnderlying, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address user, address receiverOfUnderlying, uint256 amount, uint256 index) returns()
func (_AaveATokenV2 *AaveATokenV2Session) Burn(user common.Address, receiverOfUnderlying common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Burn(&_AaveATokenV2.TransactOpts, user, receiverOfUnderlying, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address user, address receiverOfUnderlying, uint256 amount, uint256 index) returns()
func (_AaveATokenV2 *AaveATokenV2TransactorSession) Burn(user common.Address, receiverOfUnderlying common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Burn(&_AaveATokenV2.TransactOpts, user, receiverOfUnderlying, amount, index)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.DecreaseAllowance(&_AaveATokenV2.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveATokenV2 *AaveATokenV2TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.DecreaseAllowance(&_AaveATokenV2.TransactOpts, spender, subtractedValue)
}

// HandleRepayment is a paid mutator transaction binding the contract method 0x88dd91a1.
//
// Solidity: function handleRepayment(address user, uint256 amount) returns()
func (_AaveATokenV2 *AaveATokenV2Transactor) HandleRepayment(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "handleRepayment", user, amount)
}

// HandleRepayment is a paid mutator transaction binding the contract method 0x88dd91a1.
//
// Solidity: function handleRepayment(address user, uint256 amount) returns()
func (_AaveATokenV2 *AaveATokenV2Session) HandleRepayment(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.HandleRepayment(&_AaveATokenV2.TransactOpts, user, amount)
}

// HandleRepayment is a paid mutator transaction binding the contract method 0x88dd91a1.
//
// Solidity: function handleRepayment(address user, uint256 amount) returns()
func (_AaveATokenV2 *AaveATokenV2TransactorSession) HandleRepayment(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.HandleRepayment(&_AaveATokenV2.TransactOpts, user, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.IncreaseAllowance(&_AaveATokenV2.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveATokenV2 *AaveATokenV2TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.IncreaseAllowance(&_AaveATokenV2.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x183fb413.
//
// Solidity: function initialize(address pool, address treasury, address underlyingAsset, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params) returns()
func (_AaveATokenV2 *AaveATokenV2Transactor) Initialize(opts *bind.TransactOpts, pool common.Address, treasury common.Address, underlyingAsset common.Address, incentivesController common.Address, aTokenDecimals uint8, aTokenName string, aTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "initialize", pool, treasury, underlyingAsset, incentivesController, aTokenDecimals, aTokenName, aTokenSymbol, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x183fb413.
//
// Solidity: function initialize(address pool, address treasury, address underlyingAsset, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params) returns()
func (_AaveATokenV2 *AaveATokenV2Session) Initialize(pool common.Address, treasury common.Address, underlyingAsset common.Address, incentivesController common.Address, aTokenDecimals uint8, aTokenName string, aTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Initialize(&_AaveATokenV2.TransactOpts, pool, treasury, underlyingAsset, incentivesController, aTokenDecimals, aTokenName, aTokenSymbol, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x183fb413.
//
// Solidity: function initialize(address pool, address treasury, address underlyingAsset, address incentivesController, uint8 aTokenDecimals, string aTokenName, string aTokenSymbol, bytes params) returns()
func (_AaveATokenV2 *AaveATokenV2TransactorSession) Initialize(pool common.Address, treasury common.Address, underlyingAsset common.Address, incentivesController common.Address, aTokenDecimals uint8, aTokenName string, aTokenSymbol string, params []byte) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Initialize(&_AaveATokenV2.TransactOpts, pool, treasury, underlyingAsset, incentivesController, aTokenDecimals, aTokenName, aTokenSymbol, params)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address user, uint256 amount, uint256 index) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Transactor) Mint(opts *bind.TransactOpts, user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "mint", user, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address user, uint256 amount, uint256 index) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Session) Mint(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Mint(&_AaveATokenV2.TransactOpts, user, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address user, uint256 amount, uint256 index) returns(bool)
func (_AaveATokenV2 *AaveATokenV2TransactorSession) Mint(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Mint(&_AaveATokenV2.TransactOpts, user, amount, index)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x7df5bd3b.
//
// Solidity: function mintToTreasury(uint256 amount, uint256 index) returns()
func (_AaveATokenV2 *AaveATokenV2Transactor) MintToTreasury(opts *bind.TransactOpts, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "mintToTreasury", amount, index)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x7df5bd3b.
//
// Solidity: function mintToTreasury(uint256 amount, uint256 index) returns()
func (_AaveATokenV2 *AaveATokenV2Session) MintToTreasury(amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.MintToTreasury(&_AaveATokenV2.TransactOpts, amount, index)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x7df5bd3b.
//
// Solidity: function mintToTreasury(uint256 amount, uint256 index) returns()
func (_AaveATokenV2 *AaveATokenV2TransactorSession) MintToTreasury(amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.MintToTreasury(&_AaveATokenV2.TransactOpts, amount, index)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveATokenV2 *AaveATokenV2Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveATokenV2 *AaveATokenV2Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Permit(&_AaveATokenV2.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveATokenV2 *AaveATokenV2TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Permit(&_AaveATokenV2.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Transfer(&_AaveATokenV2.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.Transfer(&_AaveATokenV2.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.TransferFrom(&_AaveATokenV2.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveATokenV2 *AaveATokenV2TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.TransferFrom(&_AaveATokenV2.TransactOpts, sender, recipient, amount)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address from, address to, uint256 value) returns()
func (_AaveATokenV2 *AaveATokenV2Transactor) TransferOnLiquidation(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "transferOnLiquidation", from, to, value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address from, address to, uint256 value) returns()
func (_AaveATokenV2 *AaveATokenV2Session) TransferOnLiquidation(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.TransferOnLiquidation(&_AaveATokenV2.TransactOpts, from, to, value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address from, address to, uint256 value) returns()
func (_AaveATokenV2 *AaveATokenV2TransactorSession) TransferOnLiquidation(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.TransferOnLiquidation(&_AaveATokenV2.TransactOpts, from, to, value)
}

// TransferUnderlyingTo is a paid mutator transaction binding the contract method 0x4efecaa5.
//
// Solidity: function transferUnderlyingTo(address target, uint256 amount) returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Transactor) TransferUnderlyingTo(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.contract.Transact(opts, "transferUnderlyingTo", target, amount)
}

// TransferUnderlyingTo is a paid mutator transaction binding the contract method 0x4efecaa5.
//
// Solidity: function transferUnderlyingTo(address target, uint256 amount) returns(uint256)
func (_AaveATokenV2 *AaveATokenV2Session) TransferUnderlyingTo(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.TransferUnderlyingTo(&_AaveATokenV2.TransactOpts, target, amount)
}

// TransferUnderlyingTo is a paid mutator transaction binding the contract method 0x4efecaa5.
//
// Solidity: function transferUnderlyingTo(address target, uint256 amount) returns(uint256)
func (_AaveATokenV2 *AaveATokenV2TransactorSession) TransferUnderlyingTo(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveATokenV2.Contract.TransferUnderlyingTo(&_AaveATokenV2.TransactOpts, target, amount)
}
