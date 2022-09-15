// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderjoeveJOEBoostChef

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

// TraderjoeveJOEBoostChefMetaData contains all meta data concerning the TraderjoeveJOEBoostChef contract.
var TraderjoeveJOEBoostChefMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"JOE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_CHEF_V2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_PID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VEJOE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_allocPoint\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"_veJoeShareBp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimableJoe\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestFromMasterChef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dummyToken\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MASTER_CHEF_V2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_joe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_veJoe\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_MASTER_PID\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joePerSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingJoe\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pendingBonusToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"allocPoint\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"accJoePerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accJoePerFactorPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rewarder\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"veJoeShareBp\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"totalFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalLpSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pools\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_allocPoint\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"_veJoeShareBp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_overwrite\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newVeJoeBalance\",\"type\":\"uint256\"}],\"name\":\"updateFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"factor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TraderjoeveJOEBoostChefABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderjoeveJOEBoostChefMetaData.ABI instead.
var TraderjoeveJOEBoostChefABI = TraderjoeveJOEBoostChefMetaData.ABI

// TraderjoeveJOEBoostChef is an auto generated Go binding around an Ethereum contract.
type TraderjoeveJOEBoostChef struct {
	TraderjoeveJOEBoostChefCaller     // Read-only binding to the contract
	TraderjoeveJOEBoostChefTransactor // Write-only binding to the contract
	TraderjoeveJOEBoostChefFilterer   // Log filterer for contract events
}

// TraderjoeveJOEBoostChefCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderjoeveJOEBoostChefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeveJOEBoostChefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderjoeveJOEBoostChefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeveJOEBoostChefFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderjoeveJOEBoostChefFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeveJOEBoostChefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderjoeveJOEBoostChefSession struct {
	Contract     *TraderjoeveJOEBoostChef // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TraderjoeveJOEBoostChefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderjoeveJOEBoostChefCallerSession struct {
	Contract *TraderjoeveJOEBoostChefCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// TraderjoeveJOEBoostChefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderjoeveJOEBoostChefTransactorSession struct {
	Contract     *TraderjoeveJOEBoostChefTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// TraderjoeveJOEBoostChefRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderjoeveJOEBoostChefRaw struct {
	Contract *TraderjoeveJOEBoostChef // Generic contract binding to access the raw methods on
}

// TraderjoeveJOEBoostChefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderjoeveJOEBoostChefCallerRaw struct {
	Contract *TraderjoeveJOEBoostChefCaller // Generic read-only contract binding to access the raw methods on
}

// TraderjoeveJOEBoostChefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderjoeveJOEBoostChefTransactorRaw struct {
	Contract *TraderjoeveJOEBoostChefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderjoeveJOEBoostChef creates a new instance of TraderjoeveJOEBoostChef, bound to a specific deployed contract.
func NewTraderjoeveJOEBoostChef(address common.Address, backend bind.ContractBackend) (*TraderjoeveJOEBoostChef, error) {
	contract, err := bindTraderjoeveJOEBoostChef(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraderjoeveJOEBoostChef{TraderjoeveJOEBoostChefCaller: TraderjoeveJOEBoostChefCaller{contract: contract}, TraderjoeveJOEBoostChefTransactor: TraderjoeveJOEBoostChefTransactor{contract: contract}, TraderjoeveJOEBoostChefFilterer: TraderjoeveJOEBoostChefFilterer{contract: contract}}, nil
}

// NewTraderjoeveJOEBoostChefCaller creates a new read-only instance of TraderjoeveJOEBoostChef, bound to a specific deployed contract.
func NewTraderjoeveJOEBoostChefCaller(address common.Address, caller bind.ContractCaller) (*TraderjoeveJOEBoostChefCaller, error) {
	contract, err := bindTraderjoeveJOEBoostChef(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeveJOEBoostChefCaller{contract: contract}, nil
}

// NewTraderjoeveJOEBoostChefTransactor creates a new write-only instance of TraderjoeveJOEBoostChef, bound to a specific deployed contract.
func NewTraderjoeveJOEBoostChefTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderjoeveJOEBoostChefTransactor, error) {
	contract, err := bindTraderjoeveJOEBoostChef(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeveJOEBoostChefTransactor{contract: contract}, nil
}

// NewTraderjoeveJOEBoostChefFilterer creates a new log filterer instance of TraderjoeveJOEBoostChef, bound to a specific deployed contract.
func NewTraderjoeveJOEBoostChefFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderjoeveJOEBoostChefFilterer, error) {
	contract, err := bindTraderjoeveJOEBoostChef(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderjoeveJOEBoostChefFilterer{contract: contract}, nil
}

// bindTraderjoeveJOEBoostChef binds a generic wrapper to an already deployed contract.
func bindTraderjoeveJOEBoostChef(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraderjoeveJOEBoostChefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeveJOEBoostChef.Contract.TraderjoeveJOEBoostChefCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.TraderjoeveJOEBoostChefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.TraderjoeveJOEBoostChefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeveJOEBoostChef.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.contract.Transact(opts, method, params...)
}

// JOE is a free data retrieval call binding the contract method 0xffebad30.
//
// Solidity: function JOE() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) JOE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "JOE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JOE is a free data retrieval call binding the contract method 0xffebad30.
//
// Solidity: function JOE() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) JOE() (common.Address, error) {
	return _TraderjoeveJOEBoostChef.Contract.JOE(&_TraderjoeveJOEBoostChef.CallOpts)
}

// JOE is a free data retrieval call binding the contract method 0xffebad30.
//
// Solidity: function JOE() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) JOE() (common.Address, error) {
	return _TraderjoeveJOEBoostChef.Contract.JOE(&_TraderjoeveJOEBoostChef.CallOpts)
}

// MASTERCHEFV2 is a free data retrieval call binding the contract method 0x27bf88ad.
//
// Solidity: function MASTER_CHEF_V2() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) MASTERCHEFV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "MASTER_CHEF_V2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MASTERCHEFV2 is a free data retrieval call binding the contract method 0x27bf88ad.
//
// Solidity: function MASTER_CHEF_V2() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) MASTERCHEFV2() (common.Address, error) {
	return _TraderjoeveJOEBoostChef.Contract.MASTERCHEFV2(&_TraderjoeveJOEBoostChef.CallOpts)
}

// MASTERCHEFV2 is a free data retrieval call binding the contract method 0x27bf88ad.
//
// Solidity: function MASTER_CHEF_V2() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) MASTERCHEFV2() (common.Address, error) {
	return _TraderjoeveJOEBoostChef.Contract.MASTERCHEFV2(&_TraderjoeveJOEBoostChef.CallOpts)
}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) MASTERPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "MASTER_PID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) MASTERPID() (*big.Int, error) {
	return _TraderjoeveJOEBoostChef.Contract.MASTERPID(&_TraderjoeveJOEBoostChef.CallOpts)
}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) MASTERPID() (*big.Int, error) {
	return _TraderjoeveJOEBoostChef.Contract.MASTERPID(&_TraderjoeveJOEBoostChef.CallOpts)
}

// VEJOE is a free data retrieval call binding the contract method 0xe76fdb7e.
//
// Solidity: function VEJOE() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) VEJOE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "VEJOE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VEJOE is a free data retrieval call binding the contract method 0xe76fdb7e.
//
// Solidity: function VEJOE() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) VEJOE() (common.Address, error) {
	return _TraderjoeveJOEBoostChef.Contract.VEJOE(&_TraderjoeveJOEBoostChef.CallOpts)
}

// VEJOE is a free data retrieval call binding the contract method 0xe76fdb7e.
//
// Solidity: function VEJOE() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) VEJOE() (common.Address, error) {
	return _TraderjoeveJOEBoostChef.Contract.VEJOE(&_TraderjoeveJOEBoostChef.CallOpts)
}

// ClaimableJoe is a free data retrieval call binding the contract method 0xd00ffa8e.
//
// Solidity: function claimableJoe(uint256 , address ) view returns(uint256)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) ClaimableJoe(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "claimableJoe", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableJoe is a free data retrieval call binding the contract method 0xd00ffa8e.
//
// Solidity: function claimableJoe(uint256 , address ) view returns(uint256)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) ClaimableJoe(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _TraderjoeveJOEBoostChef.Contract.ClaimableJoe(&_TraderjoeveJOEBoostChef.CallOpts, arg0, arg1)
}

// ClaimableJoe is a free data retrieval call binding the contract method 0xd00ffa8e.
//
// Solidity: function claimableJoe(uint256 , address ) view returns(uint256)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) ClaimableJoe(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _TraderjoeveJOEBoostChef.Contract.ClaimableJoe(&_TraderjoeveJOEBoostChef.CallOpts, arg0, arg1)
}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256 amount)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) JoePerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "joePerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256 amount)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) JoePerSec() (*big.Int, error) {
	return _TraderjoeveJOEBoostChef.Contract.JoePerSec(&_TraderjoeveJOEBoostChef.CallOpts)
}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256 amount)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) JoePerSec() (*big.Int, error) {
	return _TraderjoeveJOEBoostChef.Contract.JoePerSec(&_TraderjoeveJOEBoostChef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) Owner() (common.Address, error) {
	return _TraderjoeveJOEBoostChef.Contract.Owner(&_TraderjoeveJOEBoostChef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) Owner() (common.Address, error) {
	return _TraderjoeveJOEBoostChef.Contract.Owner(&_TraderjoeveJOEBoostChef.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) PendingTokens(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "pendingTokens", _pid, _user)

	outstruct := new(struct {
		PendingJoe        *big.Int
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
		PendingBonusToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingJoe = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BonusTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PendingBonusToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _TraderjoeveJOEBoostChef.Contract.PendingTokens(&_TraderjoeveJOEBoostChef.CallOpts, _pid, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _TraderjoeveJOEBoostChef.Contract.PendingTokens(&_TraderjoeveJOEBoostChef.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint96 allocPoint, uint256 accJoePerShare, uint256 accJoePerFactorPerShare, uint64 lastRewardTimestamp, address rewarder, uint32 veJoeShareBp, uint256 totalFactor, uint256 totalLpSupply)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken                 common.Address
	AllocPoint              *big.Int
	AccJoePerShare          *big.Int
	AccJoePerFactorPerShare *big.Int
	LastRewardTimestamp     uint64
	Rewarder                common.Address
	VeJoeShareBp            uint32
	TotalFactor             *big.Int
	TotalLpSupply           *big.Int
}, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken                 common.Address
		AllocPoint              *big.Int
		AccJoePerShare          *big.Int
		AccJoePerFactorPerShare *big.Int
		LastRewardTimestamp     uint64
		Rewarder                common.Address
		VeJoeShareBp            uint32
		TotalFactor             *big.Int
		TotalLpSupply           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccJoePerShare = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccJoePerFactorPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.Rewarder = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.VeJoeShareBp = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.TotalFactor = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TotalLpSupply = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint96 allocPoint, uint256 accJoePerShare, uint256 accJoePerFactorPerShare, uint64 lastRewardTimestamp, address rewarder, uint32 veJoeShareBp, uint256 totalFactor, uint256 totalLpSupply)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken                 common.Address
	AllocPoint              *big.Int
	AccJoePerShare          *big.Int
	AccJoePerFactorPerShare *big.Int
	LastRewardTimestamp     uint64
	Rewarder                common.Address
	VeJoeShareBp            uint32
	TotalFactor             *big.Int
	TotalLpSupply           *big.Int
}, error) {
	return _TraderjoeveJOEBoostChef.Contract.PoolInfo(&_TraderjoeveJOEBoostChef.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint96 allocPoint, uint256 accJoePerShare, uint256 accJoePerFactorPerShare, uint64 lastRewardTimestamp, address rewarder, uint32 veJoeShareBp, uint256 totalFactor, uint256 totalLpSupply)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken                 common.Address
	AllocPoint              *big.Int
	AccJoePerShare          *big.Int
	AccJoePerFactorPerShare *big.Int
	LastRewardTimestamp     uint64
	Rewarder                common.Address
	VeJoeShareBp            uint32
	TotalFactor             *big.Int
	TotalLpSupply           *big.Int
}, error) {
	return _TraderjoeveJOEBoostChef.Contract.PoolInfo(&_TraderjoeveJOEBoostChef.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) PoolLength() (*big.Int, error) {
	return _TraderjoeveJOEBoostChef.Contract.PoolLength(&_TraderjoeveJOEBoostChef.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) PoolLength() (*big.Int, error) {
	return _TraderjoeveJOEBoostChef.Contract.PoolLength(&_TraderjoeveJOEBoostChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) TotalAllocPoint() (*big.Int, error) {
	return _TraderjoeveJOEBoostChef.Contract.TotalAllocPoint(&_TraderjoeveJOEBoostChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _TraderjoeveJOEBoostChef.Contract.TotalAllocPoint(&_TraderjoeveJOEBoostChef.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 factor)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	Factor     *big.Int
}, error) {
	var out []interface{}
	err := _TraderjoeveJOEBoostChef.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
		Factor     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Factor = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 factor)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	Factor     *big.Int
}, error) {
	return _TraderjoeveJOEBoostChef.Contract.UserInfo(&_TraderjoeveJOEBoostChef.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 factor)
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	Factor     *big.Int
}, error) {
	return _TraderjoeveJOEBoostChef.Contract.UserInfo(&_TraderjoeveJOEBoostChef.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x561be05a.
//
// Solidity: function add(uint96 _allocPoint, uint32 _veJoeShareBp, address _lpToken, address _rewarder) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _veJoeShareBp uint32, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "add", _allocPoint, _veJoeShareBp, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0x561be05a.
//
// Solidity: function add(uint96 _allocPoint, uint32 _veJoeShareBp, address _lpToken, address _rewarder) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) Add(_allocPoint *big.Int, _veJoeShareBp uint32, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Add(&_TraderjoeveJOEBoostChef.TransactOpts, _allocPoint, _veJoeShareBp, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0x561be05a.
//
// Solidity: function add(uint96 _allocPoint, uint32 _veJoeShareBp, address _lpToken, address _rewarder) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) Add(_allocPoint *big.Int, _veJoeShareBp uint32, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Add(&_TraderjoeveJOEBoostChef.TransactOpts, _allocPoint, _veJoeShareBp, _lpToken, _rewarder)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Deposit(&_TraderjoeveJOEBoostChef.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Deposit(&_TraderjoeveJOEBoostChef.TransactOpts, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.EmergencyWithdraw(&_TraderjoeveJOEBoostChef.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.EmergencyWithdraw(&_TraderjoeveJOEBoostChef.TransactOpts, _pid)
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) HarvestFromMasterChef(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "harvestFromMasterChef")
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) HarvestFromMasterChef() (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.HarvestFromMasterChef(&_TraderjoeveJOEBoostChef.TransactOpts)
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) HarvestFromMasterChef() (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.HarvestFromMasterChef(&_TraderjoeveJOEBoostChef.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _dummyToken) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) Init(opts *bind.TransactOpts, _dummyToken common.Address) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "init", _dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _dummyToken) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) Init(_dummyToken common.Address) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Init(&_TraderjoeveJOEBoostChef.TransactOpts, _dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _dummyToken) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) Init(_dummyToken common.Address) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Init(&_TraderjoeveJOEBoostChef.TransactOpts, _dummyToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _MASTER_CHEF_V2, address _joe, address _veJoe, uint256 _MASTER_PID) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) Initialize(opts *bind.TransactOpts, _MASTER_CHEF_V2 common.Address, _joe common.Address, _veJoe common.Address, _MASTER_PID *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "initialize", _MASTER_CHEF_V2, _joe, _veJoe, _MASTER_PID)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _MASTER_CHEF_V2, address _joe, address _veJoe, uint256 _MASTER_PID) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) Initialize(_MASTER_CHEF_V2 common.Address, _joe common.Address, _veJoe common.Address, _MASTER_PID *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Initialize(&_TraderjoeveJOEBoostChef.TransactOpts, _MASTER_CHEF_V2, _joe, _veJoe, _MASTER_PID)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _MASTER_CHEF_V2, address _joe, address _veJoe, uint256 _MASTER_PID) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) Initialize(_MASTER_CHEF_V2 common.Address, _joe common.Address, _veJoe common.Address, _MASTER_PID *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Initialize(&_TraderjoeveJOEBoostChef.TransactOpts, _MASTER_CHEF_V2, _joe, _veJoe, _MASTER_PID)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) MassUpdatePools() (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.MassUpdatePools(&_TraderjoeveJOEBoostChef.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.MassUpdatePools(&_TraderjoeveJOEBoostChef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) RenounceOwnership() (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.RenounceOwnership(&_TraderjoeveJOEBoostChef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.RenounceOwnership(&_TraderjoeveJOEBoostChef.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x980e5f90.
//
// Solidity: function set(uint256 _pid, uint96 _allocPoint, uint32 _veJoeShareBp, address _rewarder, bool _overwrite) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _veJoeShareBp uint32, _rewarder common.Address, _overwrite bool) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "set", _pid, _allocPoint, _veJoeShareBp, _rewarder, _overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x980e5f90.
//
// Solidity: function set(uint256 _pid, uint96 _allocPoint, uint32 _veJoeShareBp, address _rewarder, bool _overwrite) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) Set(_pid *big.Int, _allocPoint *big.Int, _veJoeShareBp uint32, _rewarder common.Address, _overwrite bool) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Set(&_TraderjoeveJOEBoostChef.TransactOpts, _pid, _allocPoint, _veJoeShareBp, _rewarder, _overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x980e5f90.
//
// Solidity: function set(uint256 _pid, uint96 _allocPoint, uint32 _veJoeShareBp, address _rewarder, bool _overwrite) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _veJoeShareBp uint32, _rewarder common.Address, _overwrite bool) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Set(&_TraderjoeveJOEBoostChef.TransactOpts, _pid, _allocPoint, _veJoeShareBp, _rewarder, _overwrite)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.TransferOwnership(&_TraderjoeveJOEBoostChef.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.TransferOwnership(&_TraderjoeveJOEBoostChef.TransactOpts, newOwner)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVeJoeBalance) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) UpdateFactor(opts *bind.TransactOpts, _user common.Address, _newVeJoeBalance *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "updateFactor", _user, _newVeJoeBalance)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVeJoeBalance) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) UpdateFactor(_user common.Address, _newVeJoeBalance *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.UpdateFactor(&_TraderjoeveJOEBoostChef.TransactOpts, _user, _newVeJoeBalance)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVeJoeBalance) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) UpdateFactor(_user common.Address, _newVeJoeBalance *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.UpdateFactor(&_TraderjoeveJOEBoostChef.TransactOpts, _user, _newVeJoeBalance)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.UpdatePool(&_TraderjoeveJOEBoostChef.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.UpdatePool(&_TraderjoeveJOEBoostChef.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Withdraw(&_TraderjoeveJOEBoostChef.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_TraderjoeveJOEBoostChef *TraderjoeveJOEBoostChefTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeveJOEBoostChef.Contract.Withdraw(&_TraderjoeveJOEBoostChef.TransactOpts, _pid, _amount)
}
