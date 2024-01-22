// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package openOceanRouter

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

// IOpenOceanExchangeCallDescription is an auto generated low-level Go binding around an user-defined struct.
type IOpenOceanExchangeCallDescription struct {
	Target   *big.Int
	GasLimit *big.Int
	Value    *big.Int
	Data     []byte
}

// IOpenOceanExchangeSwapDescription is an auto generated low-level Go binding around an user-defined struct.
type IOpenOceanExchangeSwapDescription struct {
	SrcToken         common.Address
	DstToken         common.Address
	SrcReceiver      common.Address
	DstReceiver      common.Address
	Amount           *big.Int
	MinReturnAmount  *big.Int
	GuaranteedAmount *big.Int
	Flags            *big.Int
	Referrer         common.Address
	Permit           []byte
}

// OpenOceanRouterMetaData contains all meta data concerning the OpenOceanRouter contract.
var OpenOceanRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"pools\",\"type\":\"bytes32[]\"}],\"name\":\"callUniswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"callUniswapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"pools\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"callUniswapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"pools\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"callUniswapWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guaranteedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structIOpenOceanExchange.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIOpenOceanExchange.CallDescription[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3Swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3SwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OpenOceanRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use OpenOceanRouterMetaData.ABI instead.
var OpenOceanRouterABI = OpenOceanRouterMetaData.ABI

// OpenOceanRouter is an auto generated Go binding around an Ethereum contract.
type OpenOceanRouter struct {
	OpenOceanRouterCaller     // Read-only binding to the contract
	OpenOceanRouterTransactor // Write-only binding to the contract
	OpenOceanRouterFilterer   // Log filterer for contract events
}

// OpenOceanRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpenOceanRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenOceanRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpenOceanRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenOceanRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpenOceanRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenOceanRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpenOceanRouterSession struct {
	Contract     *OpenOceanRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpenOceanRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpenOceanRouterCallerSession struct {
	Contract *OpenOceanRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// OpenOceanRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpenOceanRouterTransactorSession struct {
	Contract     *OpenOceanRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// OpenOceanRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpenOceanRouterRaw struct {
	Contract *OpenOceanRouter // Generic contract binding to access the raw methods on
}

// OpenOceanRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpenOceanRouterCallerRaw struct {
	Contract *OpenOceanRouterCaller // Generic read-only contract binding to access the raw methods on
}

// OpenOceanRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpenOceanRouterTransactorRaw struct {
	Contract *OpenOceanRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpenOceanRouter creates a new instance of OpenOceanRouter, bound to a specific deployed contract.
func NewOpenOceanRouter(address common.Address, backend bind.ContractBackend) (*OpenOceanRouter, error) {
	contract, err := bindOpenOceanRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpenOceanRouter{OpenOceanRouterCaller: OpenOceanRouterCaller{contract: contract}, OpenOceanRouterTransactor: OpenOceanRouterTransactor{contract: contract}, OpenOceanRouterFilterer: OpenOceanRouterFilterer{contract: contract}}, nil
}

// NewOpenOceanRouterCaller creates a new read-only instance of OpenOceanRouter, bound to a specific deployed contract.
func NewOpenOceanRouterCaller(address common.Address, caller bind.ContractCaller) (*OpenOceanRouterCaller, error) {
	contract, err := bindOpenOceanRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpenOceanRouterCaller{contract: contract}, nil
}

// NewOpenOceanRouterTransactor creates a new write-only instance of OpenOceanRouter, bound to a specific deployed contract.
func NewOpenOceanRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*OpenOceanRouterTransactor, error) {
	contract, err := bindOpenOceanRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpenOceanRouterTransactor{contract: contract}, nil
}

// NewOpenOceanRouterFilterer creates a new log filterer instance of OpenOceanRouter, bound to a specific deployed contract.
func NewOpenOceanRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*OpenOceanRouterFilterer, error) {
	contract, err := bindOpenOceanRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpenOceanRouterFilterer{contract: contract}, nil
}

// bindOpenOceanRouter binds a generic wrapper to an already deployed contract.
func bindOpenOceanRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenOceanRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenOceanRouter *OpenOceanRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenOceanRouter.Contract.OpenOceanRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenOceanRouter *OpenOceanRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.OpenOceanRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenOceanRouter *OpenOceanRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.OpenOceanRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenOceanRouter *OpenOceanRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenOceanRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenOceanRouter *OpenOceanRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenOceanRouter *OpenOceanRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpenOceanRouter *OpenOceanRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpenOceanRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpenOceanRouter *OpenOceanRouterSession) Owner() (common.Address, error) {
	return _OpenOceanRouter.Contract.Owner(&_OpenOceanRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OpenOceanRouter *OpenOceanRouterCallerSession) Owner() (common.Address, error) {
	return _OpenOceanRouter.Contract.Owner(&_OpenOceanRouter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OpenOceanRouter *OpenOceanRouterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OpenOceanRouter.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OpenOceanRouter *OpenOceanRouterSession) Paused() (bool, error) {
	return _OpenOceanRouter.Contract.Paused(&_OpenOceanRouter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OpenOceanRouter *OpenOceanRouterCallerSession) Paused() (bool, error) {
	return _OpenOceanRouter.Contract.Paused(&_OpenOceanRouter.CallOpts)
}

// CallUniswap is a paid mutator transaction binding the contract method 0x8980041a.
//
// Solidity: function callUniswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactor) CallUniswap(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "callUniswap", srcToken, amount, minReturn, pools)
}

// CallUniswap is a paid mutator transaction binding the contract method 0x8980041a.
//
// Solidity: function callUniswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterSession) CallUniswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.CallUniswap(&_OpenOceanRouter.TransactOpts, srcToken, amount, minReturn, pools)
}

// CallUniswap is a paid mutator transaction binding the contract method 0x8980041a.
//
// Solidity: function callUniswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) CallUniswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.CallUniswap(&_OpenOceanRouter.TransactOpts, srcToken, amount, minReturn, pools)
}

// CallUniswapTo is a paid mutator transaction binding the contract method 0x6b58f2f0.
//
// Solidity: function callUniswapTo(address srcToken, uint256 amount, uint256 minReturn, bytes32[] , address recipient) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactor) CallUniswapTo(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte, recipient common.Address) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "callUniswapTo", srcToken, amount, minReturn, arg3, recipient)
}

// CallUniswapTo is a paid mutator transaction binding the contract method 0x6b58f2f0.
//
// Solidity: function callUniswapTo(address srcToken, uint256 amount, uint256 minReturn, bytes32[] , address recipient) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterSession) CallUniswapTo(srcToken common.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte, recipient common.Address) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.CallUniswapTo(&_OpenOceanRouter.TransactOpts, srcToken, amount, minReturn, arg3, recipient)
}

// CallUniswapTo is a paid mutator transaction binding the contract method 0x6b58f2f0.
//
// Solidity: function callUniswapTo(address srcToken, uint256 amount, uint256 minReturn, bytes32[] , address recipient) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) CallUniswapTo(srcToken common.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte, recipient common.Address) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.CallUniswapTo(&_OpenOceanRouter.TransactOpts, srcToken, amount, minReturn, arg3, recipient)
}

// CallUniswapToWithPermit is a paid mutator transaction binding the contract method 0x22dca3d7.
//
// Solidity: function callUniswapToWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit, address recipient) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactor) CallUniswapToWithPermit(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte, recipient common.Address) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "callUniswapToWithPermit", srcToken, amount, minReturn, pools, permit, recipient)
}

// CallUniswapToWithPermit is a paid mutator transaction binding the contract method 0x22dca3d7.
//
// Solidity: function callUniswapToWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit, address recipient) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterSession) CallUniswapToWithPermit(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte, recipient common.Address) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.CallUniswapToWithPermit(&_OpenOceanRouter.TransactOpts, srcToken, amount, minReturn, pools, permit, recipient)
}

// CallUniswapToWithPermit is a paid mutator transaction binding the contract method 0x22dca3d7.
//
// Solidity: function callUniswapToWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit, address recipient) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) CallUniswapToWithPermit(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte, recipient common.Address) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.CallUniswapToWithPermit(&_OpenOceanRouter.TransactOpts, srcToken, amount, minReturn, pools, permit, recipient)
}

// CallUniswapWithPermit is a paid mutator transaction binding the contract method 0xd57360fc.
//
// Solidity: function callUniswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactor) CallUniswapWithPermit(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "callUniswapWithPermit", srcToken, amount, minReturn, pools, permit)
}

// CallUniswapWithPermit is a paid mutator transaction binding the contract method 0xd57360fc.
//
// Solidity: function callUniswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterSession) CallUniswapWithPermit(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.CallUniswapWithPermit(&_OpenOceanRouter.TransactOpts, srcToken, amount, minReturn, pools, permit)
}

// CallUniswapWithPermit is a paid mutator transaction binding the contract method 0xd57360fc.
//
// Solidity: function callUniswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) CallUniswapWithPermit(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.CallUniswapWithPermit(&_OpenOceanRouter.TransactOpts, srcToken, amount, minReturn, pools, permit)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OpenOceanRouter *OpenOceanRouterTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OpenOceanRouter *OpenOceanRouterSession) Initialize() (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.Initialize(&_OpenOceanRouter.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) Initialize() (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.Initialize(&_OpenOceanRouter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OpenOceanRouter *OpenOceanRouterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OpenOceanRouter *OpenOceanRouterSession) Pause() (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.Pause(&_OpenOceanRouter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) Pause() (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.Pause(&_OpenOceanRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpenOceanRouter *OpenOceanRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpenOceanRouter *OpenOceanRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.RenounceOwnership(&_OpenOceanRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.RenounceOwnership(&_OpenOceanRouter.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_OpenOceanRouter *OpenOceanRouterTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_OpenOceanRouter *OpenOceanRouterSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.RescueFunds(&_OpenOceanRouter.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.RescueFunds(&_OpenOceanRouter.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x90411a32.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,uint256,address,bytes) desc, (uint256,uint256,uint256,bytes)[] calls) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactor) Swap(opts *bind.TransactOpts, caller common.Address, desc IOpenOceanExchangeSwapDescription, calls []IOpenOceanExchangeCallDescription) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "swap", caller, desc, calls)
}

// Swap is a paid mutator transaction binding the contract method 0x90411a32.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,uint256,address,bytes) desc, (uint256,uint256,uint256,bytes)[] calls) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterSession) Swap(caller common.Address, desc IOpenOceanExchangeSwapDescription, calls []IOpenOceanExchangeCallDescription) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.Swap(&_OpenOceanRouter.TransactOpts, caller, desc, calls)
}

// Swap is a paid mutator transaction binding the contract method 0x90411a32.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,uint256,address,bytes) desc, (uint256,uint256,uint256,bytes)[] calls) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) Swap(caller common.Address, desc IOpenOceanExchangeSwapDescription, calls []IOpenOceanExchangeCallDescription) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.Swap(&_OpenOceanRouter.TransactOpts, caller, desc, calls)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpenOceanRouter *OpenOceanRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpenOceanRouter *OpenOceanRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.TransferOwnership(&_OpenOceanRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.TransferOwnership(&_OpenOceanRouter.TransactOpts, newOwner)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactor) UniswapV3Swap(opts *bind.TransactOpts, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "uniswapV3Swap", amount, minReturn, pools)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterSession) UniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.UniswapV3Swap(&_OpenOceanRouter.TransactOpts, amount, minReturn, pools)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) UniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.UniswapV3Swap(&_OpenOceanRouter.TransactOpts, amount, minReturn, pools)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_OpenOceanRouter *OpenOceanRouterTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_OpenOceanRouter *OpenOceanRouterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.UniswapV3SwapCallback(&_OpenOceanRouter.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.UniswapV3SwapCallback(&_OpenOceanRouter.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactor) UniswapV3SwapTo(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "uniswapV3SwapTo", recipient, amount, minReturn, pools)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterSession) UniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.UniswapV3SwapTo(&_OpenOceanRouter.TransactOpts, recipient, amount, minReturn, pools)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) UniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.UniswapV3SwapTo(&_OpenOceanRouter.TransactOpts, recipient, amount, minReturn, pools)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactor) UniswapV3SwapToWithPermit(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _OpenOceanRouter.contract.Transact(opts, "uniswapV3SwapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterSession) UniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.UniswapV3SwapToWithPermit(&_OpenOceanRouter.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_OpenOceanRouter *OpenOceanRouterTransactorSession) UniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _OpenOceanRouter.Contract.UniswapV3SwapToWithPermit(&_OpenOceanRouter.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}
