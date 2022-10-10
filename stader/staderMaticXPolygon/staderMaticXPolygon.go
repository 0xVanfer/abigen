// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staderMaticXPolygon

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

// StaderMaticXPolygonMetaData contains all meta data concerning the StaderMaticXPolygon contract.
var StaderMaticXPolygonMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CHILD_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHILD_CHAIN_ID_BYTES\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSITOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC712_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_CHAIN_ID_BYTES\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"name\":\"executeMetaTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeperator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"childChainManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StaderMaticXPolygonABI is the input ABI used to generate the binding from.
// Deprecated: Use StaderMaticXPolygonMetaData.ABI instead.
var StaderMaticXPolygonABI = StaderMaticXPolygonMetaData.ABI

// StaderMaticXPolygon is an auto generated Go binding around an Ethereum contract.
type StaderMaticXPolygon struct {
	StaderMaticXPolygonCaller     // Read-only binding to the contract
	StaderMaticXPolygonTransactor // Write-only binding to the contract
	StaderMaticXPolygonFilterer   // Log filterer for contract events
}

// StaderMaticXPolygonCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaderMaticXPolygonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderMaticXPolygonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaderMaticXPolygonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderMaticXPolygonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaderMaticXPolygonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderMaticXPolygonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaderMaticXPolygonSession struct {
	Contract     *StaderMaticXPolygon // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StaderMaticXPolygonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaderMaticXPolygonCallerSession struct {
	Contract *StaderMaticXPolygonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// StaderMaticXPolygonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaderMaticXPolygonTransactorSession struct {
	Contract     *StaderMaticXPolygonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// StaderMaticXPolygonRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaderMaticXPolygonRaw struct {
	Contract *StaderMaticXPolygon // Generic contract binding to access the raw methods on
}

// StaderMaticXPolygonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaderMaticXPolygonCallerRaw struct {
	Contract *StaderMaticXPolygonCaller // Generic read-only contract binding to access the raw methods on
}

// StaderMaticXPolygonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaderMaticXPolygonTransactorRaw struct {
	Contract *StaderMaticXPolygonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaderMaticXPolygon creates a new instance of StaderMaticXPolygon, bound to a specific deployed contract.
func NewStaderMaticXPolygon(address common.Address, backend bind.ContractBackend) (*StaderMaticXPolygon, error) {
	contract, err := bindStaderMaticXPolygon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaderMaticXPolygon{StaderMaticXPolygonCaller: StaderMaticXPolygonCaller{contract: contract}, StaderMaticXPolygonTransactor: StaderMaticXPolygonTransactor{contract: contract}, StaderMaticXPolygonFilterer: StaderMaticXPolygonFilterer{contract: contract}}, nil
}

// NewStaderMaticXPolygonCaller creates a new read-only instance of StaderMaticXPolygon, bound to a specific deployed contract.
func NewStaderMaticXPolygonCaller(address common.Address, caller bind.ContractCaller) (*StaderMaticXPolygonCaller, error) {
	contract, err := bindStaderMaticXPolygon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaderMaticXPolygonCaller{contract: contract}, nil
}

// NewStaderMaticXPolygonTransactor creates a new write-only instance of StaderMaticXPolygon, bound to a specific deployed contract.
func NewStaderMaticXPolygonTransactor(address common.Address, transactor bind.ContractTransactor) (*StaderMaticXPolygonTransactor, error) {
	contract, err := bindStaderMaticXPolygon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaderMaticXPolygonTransactor{contract: contract}, nil
}

// NewStaderMaticXPolygonFilterer creates a new log filterer instance of StaderMaticXPolygon, bound to a specific deployed contract.
func NewStaderMaticXPolygonFilterer(address common.Address, filterer bind.ContractFilterer) (*StaderMaticXPolygonFilterer, error) {
	contract, err := bindStaderMaticXPolygon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaderMaticXPolygonFilterer{contract: contract}, nil
}

// bindStaderMaticXPolygon binds a generic wrapper to an already deployed contract.
func bindStaderMaticXPolygon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StaderMaticXPolygonABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderMaticXPolygon *StaderMaticXPolygonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderMaticXPolygon.Contract.StaderMaticXPolygonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderMaticXPolygon *StaderMaticXPolygonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.StaderMaticXPolygonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderMaticXPolygon *StaderMaticXPolygonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.StaderMaticXPolygonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderMaticXPolygon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.contract.Transact(opts, method, params...)
}

// CHILDCHAINID is a free data retrieval call binding the contract method 0x626381a0.
//
// Solidity: function CHILD_CHAIN_ID() view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) CHILDCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "CHILD_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHILDCHAINID is a free data retrieval call binding the contract method 0x626381a0.
//
// Solidity: function CHILD_CHAIN_ID() view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) CHILDCHAINID() (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.CHILDCHAINID(&_StaderMaticXPolygon.CallOpts)
}

// CHILDCHAINID is a free data retrieval call binding the contract method 0x626381a0.
//
// Solidity: function CHILD_CHAIN_ID() view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) CHILDCHAINID() (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.CHILDCHAINID(&_StaderMaticXPolygon.CallOpts)
}

// CHILDCHAINIDBYTES is a free data retrieval call binding the contract method 0x0b54817c.
//
// Solidity: function CHILD_CHAIN_ID_BYTES() view returns(bytes)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) CHILDCHAINIDBYTES(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "CHILD_CHAIN_ID_BYTES")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CHILDCHAINIDBYTES is a free data retrieval call binding the contract method 0x0b54817c.
//
// Solidity: function CHILD_CHAIN_ID_BYTES() view returns(bytes)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) CHILDCHAINIDBYTES() ([]byte, error) {
	return _StaderMaticXPolygon.Contract.CHILDCHAINIDBYTES(&_StaderMaticXPolygon.CallOpts)
}

// CHILDCHAINIDBYTES is a free data retrieval call binding the contract method 0x0b54817c.
//
// Solidity: function CHILD_CHAIN_ID_BYTES() view returns(bytes)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) CHILDCHAINIDBYTES() ([]byte, error) {
	return _StaderMaticXPolygon.Contract.CHILDCHAINIDBYTES(&_StaderMaticXPolygon.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StaderMaticXPolygon.Contract.DEFAULTADMINROLE(&_StaderMaticXPolygon.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StaderMaticXPolygon.Contract.DEFAULTADMINROLE(&_StaderMaticXPolygon.CallOpts)
}

// DEPOSITORROLE is a free data retrieval call binding the contract method 0xa3b0b5a3.
//
// Solidity: function DEPOSITOR_ROLE() view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) DEPOSITORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "DEPOSITOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITORROLE is a free data retrieval call binding the contract method 0xa3b0b5a3.
//
// Solidity: function DEPOSITOR_ROLE() view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) DEPOSITORROLE() ([32]byte, error) {
	return _StaderMaticXPolygon.Contract.DEPOSITORROLE(&_StaderMaticXPolygon.CallOpts)
}

// DEPOSITORROLE is a free data retrieval call binding the contract method 0xa3b0b5a3.
//
// Solidity: function DEPOSITOR_ROLE() view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) DEPOSITORROLE() ([32]byte, error) {
	return _StaderMaticXPolygon.Contract.DEPOSITORROLE(&_StaderMaticXPolygon.CallOpts)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) ERC712VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "ERC712_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) ERC712VERSION() (string, error) {
	return _StaderMaticXPolygon.Contract.ERC712VERSION(&_StaderMaticXPolygon.CallOpts)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) ERC712VERSION() (string, error) {
	return _StaderMaticXPolygon.Contract.ERC712VERSION(&_StaderMaticXPolygon.CallOpts)
}

// ROOTCHAINID is a free data retrieval call binding the contract method 0x8acfcaf7.
//
// Solidity: function ROOT_CHAIN_ID() view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) ROOTCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "ROOT_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROOTCHAINID is a free data retrieval call binding the contract method 0x8acfcaf7.
//
// Solidity: function ROOT_CHAIN_ID() view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) ROOTCHAINID() (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.ROOTCHAINID(&_StaderMaticXPolygon.CallOpts)
}

// ROOTCHAINID is a free data retrieval call binding the contract method 0x8acfcaf7.
//
// Solidity: function ROOT_CHAIN_ID() view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) ROOTCHAINID() (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.ROOTCHAINID(&_StaderMaticXPolygon.CallOpts)
}

// ROOTCHAINIDBYTES is a free data retrieval call binding the contract method 0x0dd7531a.
//
// Solidity: function ROOT_CHAIN_ID_BYTES() view returns(bytes)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) ROOTCHAINIDBYTES(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "ROOT_CHAIN_ID_BYTES")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ROOTCHAINIDBYTES is a free data retrieval call binding the contract method 0x0dd7531a.
//
// Solidity: function ROOT_CHAIN_ID_BYTES() view returns(bytes)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) ROOTCHAINIDBYTES() ([]byte, error) {
	return _StaderMaticXPolygon.Contract.ROOTCHAINIDBYTES(&_StaderMaticXPolygon.CallOpts)
}

// ROOTCHAINIDBYTES is a free data retrieval call binding the contract method 0x0dd7531a.
//
// Solidity: function ROOT_CHAIN_ID_BYTES() view returns(bytes)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) ROOTCHAINIDBYTES() ([]byte, error) {
	return _StaderMaticXPolygon.Contract.ROOTCHAINIDBYTES(&_StaderMaticXPolygon.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.Allowance(&_StaderMaticXPolygon.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.Allowance(&_StaderMaticXPolygon.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.BalanceOf(&_StaderMaticXPolygon.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.BalanceOf(&_StaderMaticXPolygon.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) Decimals() (uint8, error) {
	return _StaderMaticXPolygon.Contract.Decimals(&_StaderMaticXPolygon.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) Decimals() (uint8, error) {
	return _StaderMaticXPolygon.Contract.Decimals(&_StaderMaticXPolygon.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) GetChainId() (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.GetChainId(&_StaderMaticXPolygon.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() pure returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) GetChainId() (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.GetChainId(&_StaderMaticXPolygon.CallOpts)
}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) GetDomainSeperator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "getDomainSeperator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) GetDomainSeperator() ([32]byte, error) {
	return _StaderMaticXPolygon.Contract.GetDomainSeperator(&_StaderMaticXPolygon.CallOpts)
}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) GetDomainSeperator() ([32]byte, error) {
	return _StaderMaticXPolygon.Contract.GetDomainSeperator(&_StaderMaticXPolygon.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) GetNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "getNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) GetNonce(user common.Address) (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.GetNonce(&_StaderMaticXPolygon.CallOpts, user)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) GetNonce(user common.Address) (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.GetNonce(&_StaderMaticXPolygon.CallOpts, user)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StaderMaticXPolygon.Contract.GetRoleAdmin(&_StaderMaticXPolygon.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StaderMaticXPolygon.Contract.GetRoleAdmin(&_StaderMaticXPolygon.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StaderMaticXPolygon.Contract.GetRoleMember(&_StaderMaticXPolygon.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StaderMaticXPolygon.Contract.GetRoleMember(&_StaderMaticXPolygon.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.GetRoleMemberCount(&_StaderMaticXPolygon.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.GetRoleMemberCount(&_StaderMaticXPolygon.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StaderMaticXPolygon.Contract.HasRole(&_StaderMaticXPolygon.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StaderMaticXPolygon.Contract.HasRole(&_StaderMaticXPolygon.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) Name() (string, error) {
	return _StaderMaticXPolygon.Contract.Name(&_StaderMaticXPolygon.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) Name() (string, error) {
	return _StaderMaticXPolygon.Contract.Name(&_StaderMaticXPolygon.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) Symbol() (string, error) {
	return _StaderMaticXPolygon.Contract.Symbol(&_StaderMaticXPolygon.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) Symbol() (string, error) {
	return _StaderMaticXPolygon.Contract.Symbol(&_StaderMaticXPolygon.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXPolygon.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) TotalSupply() (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.TotalSupply(&_StaderMaticXPolygon.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StaderMaticXPolygon *StaderMaticXPolygonCallerSession) TotalSupply() (*big.Int, error) {
	return _StaderMaticXPolygon.Contract.TotalSupply(&_StaderMaticXPolygon.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.Approve(&_StaderMaticXPolygon.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.Approve(&_StaderMaticXPolygon.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.DecreaseAllowance(&_StaderMaticXPolygon.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.DecreaseAllowance(&_StaderMaticXPolygon.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xcf2c52cb.
//
// Solidity: function deposit(address user, bytes depositData) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) Deposit(opts *bind.TransactOpts, user common.Address, depositData []byte) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "deposit", user, depositData)
}

// Deposit is a paid mutator transaction binding the contract method 0xcf2c52cb.
//
// Solidity: function deposit(address user, bytes depositData) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) Deposit(user common.Address, depositData []byte) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.Deposit(&_StaderMaticXPolygon.TransactOpts, user, depositData)
}

// Deposit is a paid mutator transaction binding the contract method 0xcf2c52cb.
//
// Solidity: function deposit(address user, bytes depositData) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) Deposit(user common.Address, depositData []byte) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.Deposit(&_StaderMaticXPolygon.TransactOpts, user, depositData)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns(bytes)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) ExecuteMetaTransaction(opts *bind.TransactOpts, userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "executeMetaTransaction", userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns(bytes)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.ExecuteMetaTransaction(&_StaderMaticXPolygon.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns(bytes)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.ExecuteMetaTransaction(&_StaderMaticXPolygon.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.GrantRole(&_StaderMaticXPolygon.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.GrantRole(&_StaderMaticXPolygon.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.IncreaseAllowance(&_StaderMaticXPolygon.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.IncreaseAllowance(&_StaderMaticXPolygon.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xde7ea79d.
//
// Solidity: function initialize(string name_, string symbol_, uint8 decimals_, address childChainManager) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, decimals_ uint8, childChainManager common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "initialize", name_, symbol_, decimals_, childChainManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xde7ea79d.
//
// Solidity: function initialize(string name_, string symbol_, uint8 decimals_, address childChainManager) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) Initialize(name_ string, symbol_ string, decimals_ uint8, childChainManager common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.Initialize(&_StaderMaticXPolygon.TransactOpts, name_, symbol_, decimals_, childChainManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xde7ea79d.
//
// Solidity: function initialize(string name_, string symbol_, uint8 decimals_, address childChainManager) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) Initialize(name_ string, symbol_ string, decimals_ uint8, childChainManager common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.Initialize(&_StaderMaticXPolygon.TransactOpts, name_, symbol_, decimals_, childChainManager)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.RenounceRole(&_StaderMaticXPolygon.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.RenounceRole(&_StaderMaticXPolygon.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.RevokeRole(&_StaderMaticXPolygon.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.RevokeRole(&_StaderMaticXPolygon.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.Transfer(&_StaderMaticXPolygon.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.Transfer(&_StaderMaticXPolygon.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.TransferFrom(&_StaderMaticXPolygon.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.TransferFrom(&_StaderMaticXPolygon.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.Withdraw(&_StaderMaticXPolygon.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_StaderMaticXPolygon *StaderMaticXPolygonTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXPolygon.Contract.Withdraw(&_StaderMaticXPolygon.TransactOpts, amount)
}
