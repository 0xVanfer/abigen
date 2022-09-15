// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderjoeSimpleReward

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

// TraderjoeSimpleRewardMetaData contains all meta data concerning the TraderjoeSimpleReward contract.
var TraderjoeSimpleRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MCJ\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isNative\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lpAmount\",\"type\":\"uint256\"}],\"name\":\"onJoeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accTokenPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenPerSec\",\"type\":\"uint256\"}],\"name\":\"setRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPerSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"direct\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"renounce\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TraderjoeSimpleRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderjoeSimpleRewardMetaData.ABI instead.
var TraderjoeSimpleRewardABI = TraderjoeSimpleRewardMetaData.ABI

// TraderjoeSimpleReward is an auto generated Go binding around an Ethereum contract.
type TraderjoeSimpleReward struct {
	TraderjoeSimpleRewardCaller     // Read-only binding to the contract
	TraderjoeSimpleRewardTransactor // Write-only binding to the contract
	TraderjoeSimpleRewardFilterer   // Log filterer for contract events
}

// TraderjoeSimpleRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderjoeSimpleRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeSimpleRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderjoeSimpleRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeSimpleRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderjoeSimpleRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeSimpleRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderjoeSimpleRewardSession struct {
	Contract     *TraderjoeSimpleReward // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TraderjoeSimpleRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderjoeSimpleRewardCallerSession struct {
	Contract *TraderjoeSimpleRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TraderjoeSimpleRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderjoeSimpleRewardTransactorSession struct {
	Contract     *TraderjoeSimpleRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TraderjoeSimpleRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderjoeSimpleRewardRaw struct {
	Contract *TraderjoeSimpleReward // Generic contract binding to access the raw methods on
}

// TraderjoeSimpleRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderjoeSimpleRewardCallerRaw struct {
	Contract *TraderjoeSimpleRewardCaller // Generic read-only contract binding to access the raw methods on
}

// TraderjoeSimpleRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderjoeSimpleRewardTransactorRaw struct {
	Contract *TraderjoeSimpleRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderjoeSimpleReward creates a new instance of TraderjoeSimpleReward, bound to a specific deployed contract.
func NewTraderjoeSimpleReward(address common.Address, backend bind.ContractBackend) (*TraderjoeSimpleReward, error) {
	contract, err := bindTraderjoeSimpleReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraderjoeSimpleReward{TraderjoeSimpleRewardCaller: TraderjoeSimpleRewardCaller{contract: contract}, TraderjoeSimpleRewardTransactor: TraderjoeSimpleRewardTransactor{contract: contract}, TraderjoeSimpleRewardFilterer: TraderjoeSimpleRewardFilterer{contract: contract}}, nil
}

// NewTraderjoeSimpleRewardCaller creates a new read-only instance of TraderjoeSimpleReward, bound to a specific deployed contract.
func NewTraderjoeSimpleRewardCaller(address common.Address, caller bind.ContractCaller) (*TraderjoeSimpleRewardCaller, error) {
	contract, err := bindTraderjoeSimpleReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeSimpleRewardCaller{contract: contract}, nil
}

// NewTraderjoeSimpleRewardTransactor creates a new write-only instance of TraderjoeSimpleReward, bound to a specific deployed contract.
func NewTraderjoeSimpleRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderjoeSimpleRewardTransactor, error) {
	contract, err := bindTraderjoeSimpleReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeSimpleRewardTransactor{contract: contract}, nil
}

// NewTraderjoeSimpleRewardFilterer creates a new log filterer instance of TraderjoeSimpleReward, bound to a specific deployed contract.
func NewTraderjoeSimpleRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderjoeSimpleRewardFilterer, error) {
	contract, err := bindTraderjoeSimpleReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderjoeSimpleRewardFilterer{contract: contract}, nil
}

// bindTraderjoeSimpleReward binds a generic wrapper to an already deployed contract.
func bindTraderjoeSimpleReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraderjoeSimpleRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeSimpleReward.Contract.TraderjoeSimpleRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.TraderjoeSimpleRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.TraderjoeSimpleRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeSimpleReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.contract.Transact(opts, method, params...)
}

// MCJ is a free data retrieval call binding the contract method 0x30628e15.
//
// Solidity: function MCJ() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCaller) MCJ(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeSimpleReward.contract.Call(opts, &out, "MCJ")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MCJ is a free data retrieval call binding the contract method 0x30628e15.
//
// Solidity: function MCJ() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) MCJ() (common.Address, error) {
	return _TraderjoeSimpleReward.Contract.MCJ(&_TraderjoeSimpleReward.CallOpts)
}

// MCJ is a free data retrieval call binding the contract method 0x30628e15.
//
// Solidity: function MCJ() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerSession) MCJ() (common.Address, error) {
	return _TraderjoeSimpleReward.Contract.MCJ(&_TraderjoeSimpleReward.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeSimpleReward.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) Balance() (*big.Int, error) {
	return _TraderjoeSimpleReward.Contract.Balance(&_TraderjoeSimpleReward.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerSession) Balance() (*big.Int, error) {
	return _TraderjoeSimpleReward.Contract.Balance(&_TraderjoeSimpleReward.CallOpts)
}

// IsNative is a free data retrieval call binding the contract method 0x73cfc6b2.
//
// Solidity: function isNative() view returns(bool)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCaller) IsNative(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TraderjoeSimpleReward.contract.Call(opts, &out, "isNative")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNative is a free data retrieval call binding the contract method 0x73cfc6b2.
//
// Solidity: function isNative() view returns(bool)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) IsNative() (bool, error) {
	return _TraderjoeSimpleReward.Contract.IsNative(&_TraderjoeSimpleReward.CallOpts)
}

// IsNative is a free data retrieval call binding the contract method 0x73cfc6b2.
//
// Solidity: function isNative() view returns(bool)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerSession) IsNative() (bool, error) {
	return _TraderjoeSimpleReward.Contract.IsNative(&_TraderjoeSimpleReward.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeSimpleReward.contract.Call(opts, &out, "lpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) LpToken() (common.Address, error) {
	return _TraderjoeSimpleReward.Contract.LpToken(&_TraderjoeSimpleReward.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerSession) LpToken() (common.Address, error) {
	return _TraderjoeSimpleReward.Contract.LpToken(&_TraderjoeSimpleReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeSimpleReward.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) Owner() (common.Address, error) {
	return _TraderjoeSimpleReward.Contract.Owner(&_TraderjoeSimpleReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerSession) Owner() (common.Address, error) {
	return _TraderjoeSimpleReward.Contract.Owner(&_TraderjoeSimpleReward.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeSimpleReward.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) PendingOwner() (common.Address, error) {
	return _TraderjoeSimpleReward.Contract.PendingOwner(&_TraderjoeSimpleReward.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerSession) PendingOwner() (common.Address, error) {
	return _TraderjoeSimpleReward.Contract.PendingOwner(&_TraderjoeSimpleReward.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xc031a66f.
//
// Solidity: function pendingTokens(address _user) view returns(uint256 pending)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCaller) PendingTokens(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeSimpleReward.contract.Call(opts, &out, "pendingTokens", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xc031a66f.
//
// Solidity: function pendingTokens(address _user) view returns(uint256 pending)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) PendingTokens(_user common.Address) (*big.Int, error) {
	return _TraderjoeSimpleReward.Contract.PendingTokens(&_TraderjoeSimpleReward.CallOpts, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xc031a66f.
//
// Solidity: function pendingTokens(address _user) view returns(uint256 pending)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerSession) PendingTokens(_user common.Address) (*big.Int, error) {
	return _TraderjoeSimpleReward.Contract.PendingTokens(&_TraderjoeSimpleReward.CallOpts, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x5a2f3d09.
//
// Solidity: function poolInfo() view returns(uint256 accTokenPerShare, uint256 lastRewardTimestamp)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCaller) PoolInfo(opts *bind.CallOpts) (struct {
	AccTokenPerShare    *big.Int
	LastRewardTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _TraderjoeSimpleReward.contract.Call(opts, &out, "poolInfo")

	outstruct := new(struct {
		AccTokenPerShare    *big.Int
		LastRewardTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccTokenPerShare = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x5a2f3d09.
//
// Solidity: function poolInfo() view returns(uint256 accTokenPerShare, uint256 lastRewardTimestamp)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) PoolInfo() (struct {
	AccTokenPerShare    *big.Int
	LastRewardTimestamp *big.Int
}, error) {
	return _TraderjoeSimpleReward.Contract.PoolInfo(&_TraderjoeSimpleReward.CallOpts)
}

// PoolInfo is a free data retrieval call binding the contract method 0x5a2f3d09.
//
// Solidity: function poolInfo() view returns(uint256 accTokenPerShare, uint256 lastRewardTimestamp)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerSession) PoolInfo() (struct {
	AccTokenPerShare    *big.Int
	LastRewardTimestamp *big.Int
}, error) {
	return _TraderjoeSimpleReward.Contract.PoolInfo(&_TraderjoeSimpleReward.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeSimpleReward.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) RewardToken() (common.Address, error) {
	return _TraderjoeSimpleReward.Contract.RewardToken(&_TraderjoeSimpleReward.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerSession) RewardToken() (common.Address, error) {
	return _TraderjoeSimpleReward.Contract.RewardToken(&_TraderjoeSimpleReward.CallOpts)
}

// TokenPerSec is a free data retrieval call binding the contract method 0x50fd1f3e.
//
// Solidity: function tokenPerSec() view returns(uint256)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCaller) TokenPerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeSimpleReward.contract.Call(opts, &out, "tokenPerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPerSec is a free data retrieval call binding the contract method 0x50fd1f3e.
//
// Solidity: function tokenPerSec() view returns(uint256)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) TokenPerSec() (*big.Int, error) {
	return _TraderjoeSimpleReward.Contract.TokenPerSec(&_TraderjoeSimpleReward.CallOpts)
}

// TokenPerSec is a free data retrieval call binding the contract method 0x50fd1f3e.
//
// Solidity: function tokenPerSec() view returns(uint256)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerSession) TokenPerSec() (*big.Int, error) {
	return _TraderjoeSimpleReward.Contract.TokenPerSec(&_TraderjoeSimpleReward.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _TraderjoeSimpleReward.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _TraderjoeSimpleReward.Contract.UserInfo(&_TraderjoeSimpleReward.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _TraderjoeSimpleReward.Contract.UserInfo(&_TraderjoeSimpleReward.CallOpts, arg0)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) ClaimOwnership() (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.ClaimOwnership(&_TraderjoeSimpleReward.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.ClaimOwnership(&_TraderjoeSimpleReward.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.EmergencyWithdraw(&_TraderjoeSimpleReward.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.EmergencyWithdraw(&_TraderjoeSimpleReward.TransactOpts)
}

// OnJoeReward is a paid mutator transaction binding the contract method 0x1a7af8b0.
//
// Solidity: function onJoeReward(address _user, uint256 _lpAmount) returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactor) OnJoeReward(opts *bind.TransactOpts, _user common.Address, _lpAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.contract.Transact(opts, "onJoeReward", _user, _lpAmount)
}

// OnJoeReward is a paid mutator transaction binding the contract method 0x1a7af8b0.
//
// Solidity: function onJoeReward(address _user, uint256 _lpAmount) returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) OnJoeReward(_user common.Address, _lpAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.OnJoeReward(&_TraderjoeSimpleReward.TransactOpts, _user, _lpAmount)
}

// OnJoeReward is a paid mutator transaction binding the contract method 0x1a7af8b0.
//
// Solidity: function onJoeReward(address _user, uint256 _lpAmount) returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactorSession) OnJoeReward(_user common.Address, _lpAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.OnJoeReward(&_TraderjoeSimpleReward.TransactOpts, _user, _lpAmount)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _tokenPerSec) returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactor) SetRewardRate(opts *bind.TransactOpts, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.contract.Transact(opts, "setRewardRate", _tokenPerSec)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _tokenPerSec) returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) SetRewardRate(_tokenPerSec *big.Int) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.SetRewardRate(&_TraderjoeSimpleReward.TransactOpts, _tokenPerSec)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _tokenPerSec) returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactorSession) SetRewardRate(_tokenPerSec *big.Int) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.SetRewardRate(&_TraderjoeSimpleReward.TransactOpts, _tokenPerSec)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.contract.Transact(opts, "transferOwnership", newOwner, direct, renounce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardSession) TransferOwnership(newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.TransferOwnership(&_TraderjoeSimpleReward.TransactOpts, newOwner, direct, renounce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_TraderjoeSimpleReward *TraderjoeSimpleRewardTransactorSession) TransferOwnership(newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _TraderjoeSimpleReward.Contract.TransferOwnership(&_TraderjoeSimpleReward.TransactOpts, newOwner, direct, renounce)
}
