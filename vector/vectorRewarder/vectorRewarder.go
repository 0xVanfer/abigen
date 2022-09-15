// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vectorRewarder

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

// VectorRewarderMetaData contains all meta data concerning the VectorRewarder contract.
var VectorRewarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountReward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"name\":\"donateRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRewardToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainRewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountReward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"name\":\"queueNewRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"name\":\"rewardDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queuedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"historicalRewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"updateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdrawFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VectorRewarderABI is the input ABI used to generate the binding from.
// Deprecated: Use VectorRewarderMetaData.ABI instead.
var VectorRewarderABI = VectorRewarderMetaData.ABI

// VectorRewarder is an auto generated Go binding around an Ethereum contract.
type VectorRewarder struct {
	VectorRewarderCaller     // Read-only binding to the contract
	VectorRewarderTransactor // Write-only binding to the contract
	VectorRewarderFilterer   // Log filterer for contract events
}

// VectorRewarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type VectorRewarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorRewarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VectorRewarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorRewarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VectorRewarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorRewarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VectorRewarderSession struct {
	Contract     *VectorRewarder   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VectorRewarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VectorRewarderCallerSession struct {
	Contract *VectorRewarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VectorRewarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VectorRewarderTransactorSession struct {
	Contract     *VectorRewarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VectorRewarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type VectorRewarderRaw struct {
	Contract *VectorRewarder // Generic contract binding to access the raw methods on
}

// VectorRewarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VectorRewarderCallerRaw struct {
	Contract *VectorRewarderCaller // Generic read-only contract binding to access the raw methods on
}

// VectorRewarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VectorRewarderTransactorRaw struct {
	Contract *VectorRewarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVectorRewarder creates a new instance of VectorRewarder, bound to a specific deployed contract.
func NewVectorRewarder(address common.Address, backend bind.ContractBackend) (*VectorRewarder, error) {
	contract, err := bindVectorRewarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VectorRewarder{VectorRewarderCaller: VectorRewarderCaller{contract: contract}, VectorRewarderTransactor: VectorRewarderTransactor{contract: contract}, VectorRewarderFilterer: VectorRewarderFilterer{contract: contract}}, nil
}

// NewVectorRewarderCaller creates a new read-only instance of VectorRewarder, bound to a specific deployed contract.
func NewVectorRewarderCaller(address common.Address, caller bind.ContractCaller) (*VectorRewarderCaller, error) {
	contract, err := bindVectorRewarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VectorRewarderCaller{contract: contract}, nil
}

// NewVectorRewarderTransactor creates a new write-only instance of VectorRewarder, bound to a specific deployed contract.
func NewVectorRewarderTransactor(address common.Address, transactor bind.ContractTransactor) (*VectorRewarderTransactor, error) {
	contract, err := bindVectorRewarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VectorRewarderTransactor{contract: contract}, nil
}

// NewVectorRewarderFilterer creates a new log filterer instance of VectorRewarder, bound to a specific deployed contract.
func NewVectorRewarderFilterer(address common.Address, filterer bind.ContractFilterer) (*VectorRewarderFilterer, error) {
	contract, err := bindVectorRewarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VectorRewarderFilterer{contract: contract}, nil
}

// bindVectorRewarder binds a generic wrapper to an already deployed contract.
func bindVectorRewarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VectorRewarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VectorRewarder *VectorRewarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VectorRewarder.Contract.VectorRewarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VectorRewarder *VectorRewarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorRewarder.Contract.VectorRewarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VectorRewarder *VectorRewarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VectorRewarder.Contract.VectorRewarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VectorRewarder *VectorRewarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VectorRewarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VectorRewarder *VectorRewarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorRewarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VectorRewarder *VectorRewarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VectorRewarder.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_VectorRewarder *VectorRewarderCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_VectorRewarder *VectorRewarderSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.BalanceOf(&_VectorRewarder.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_VectorRewarder *VectorRewarderCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.BalanceOf(&_VectorRewarder.CallOpts, _account)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address _account, address _rewardToken) view returns(uint256)
func (_VectorRewarder *VectorRewarderCaller) Earned(opts *bind.CallOpts, _account common.Address, _rewardToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "earned", _account, _rewardToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address _account, address _rewardToken) view returns(uint256)
func (_VectorRewarder *VectorRewarderSession) Earned(_account common.Address, _rewardToken common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.Earned(&_VectorRewarder.CallOpts, _account, _rewardToken)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address _account, address _rewardToken) view returns(uint256)
func (_VectorRewarder *VectorRewarderCallerSession) Earned(_account common.Address, _rewardToken common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.Earned(&_VectorRewarder.CallOpts, _account, _rewardToken)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_VectorRewarder *VectorRewarderCaller) GetStakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "getStakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_VectorRewarder *VectorRewarderSession) GetStakingToken() (common.Address, error) {
	return _VectorRewarder.Contract.GetStakingToken(&_VectorRewarder.CallOpts)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_VectorRewarder *VectorRewarderCallerSession) GetStakingToken() (common.Address, error) {
	return _VectorRewarder.Contract.GetStakingToken(&_VectorRewarder.CallOpts)
}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address ) view returns(bool)
func (_VectorRewarder *VectorRewarderCaller) IsRewardToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "isRewardToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address ) view returns(bool)
func (_VectorRewarder *VectorRewarderSession) IsRewardToken(arg0 common.Address) (bool, error) {
	return _VectorRewarder.Contract.IsRewardToken(&_VectorRewarder.CallOpts, arg0)
}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address ) view returns(bool)
func (_VectorRewarder *VectorRewarderCallerSession) IsRewardToken(arg0 common.Address) (bool, error) {
	return _VectorRewarder.Contract.IsRewardToken(&_VectorRewarder.CallOpts, arg0)
}

// MainRewardToken is a free data retrieval call binding the contract method 0xf3fc7c2b.
//
// Solidity: function mainRewardToken() view returns(address)
func (_VectorRewarder *VectorRewarderCaller) MainRewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "mainRewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainRewardToken is a free data retrieval call binding the contract method 0xf3fc7c2b.
//
// Solidity: function mainRewardToken() view returns(address)
func (_VectorRewarder *VectorRewarderSession) MainRewardToken() (common.Address, error) {
	return _VectorRewarder.Contract.MainRewardToken(&_VectorRewarder.CallOpts)
}

// MainRewardToken is a free data retrieval call binding the contract method 0xf3fc7c2b.
//
// Solidity: function mainRewardToken() view returns(address)
func (_VectorRewarder *VectorRewarderCallerSession) MainRewardToken() (common.Address, error) {
	return _VectorRewarder.Contract.MainRewardToken(&_VectorRewarder.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_VectorRewarder *VectorRewarderCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_VectorRewarder *VectorRewarderSession) Operator() (common.Address, error) {
	return _VectorRewarder.Contract.Operator(&_VectorRewarder.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_VectorRewarder *VectorRewarderCallerSession) Operator() (common.Address, error) {
	return _VectorRewarder.Contract.Operator(&_VectorRewarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VectorRewarder *VectorRewarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VectorRewarder *VectorRewarderSession) Owner() (common.Address, error) {
	return _VectorRewarder.Contract.Owner(&_VectorRewarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VectorRewarder *VectorRewarderCallerSession) Owner() (common.Address, error) {
	return _VectorRewarder.Contract.Owner(&_VectorRewarder.CallOpts)
}

// RewardDecimals is a free data retrieval call binding the contract method 0x99da1729.
//
// Solidity: function rewardDecimals(address _rewardToken) view returns(uint256)
func (_VectorRewarder *VectorRewarderCaller) RewardDecimals(opts *bind.CallOpts, _rewardToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "rewardDecimals", _rewardToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardDecimals is a free data retrieval call binding the contract method 0x99da1729.
//
// Solidity: function rewardDecimals(address _rewardToken) view returns(uint256)
func (_VectorRewarder *VectorRewarderSession) RewardDecimals(_rewardToken common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.RewardDecimals(&_VectorRewarder.CallOpts, _rewardToken)
}

// RewardDecimals is a free data retrieval call binding the contract method 0x99da1729.
//
// Solidity: function rewardDecimals(address _rewardToken) view returns(uint256)
func (_VectorRewarder *VectorRewarderCallerSession) RewardDecimals(_rewardToken common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.RewardDecimals(&_VectorRewarder.CallOpts, _rewardToken)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_VectorRewarder *VectorRewarderCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_VectorRewarder *VectorRewarderSession) RewardManager() (common.Address, error) {
	return _VectorRewarder.Contract.RewardManager(&_VectorRewarder.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_VectorRewarder *VectorRewarderCallerSession) RewardManager() (common.Address, error) {
	return _VectorRewarder.Contract.RewardManager(&_VectorRewarder.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardToken) view returns(uint256)
func (_VectorRewarder *VectorRewarderCaller) RewardPerToken(opts *bind.CallOpts, _rewardToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "rewardPerToken", _rewardToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardToken) view returns(uint256)
func (_VectorRewarder *VectorRewarderSession) RewardPerToken(_rewardToken common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.RewardPerToken(&_VectorRewarder.CallOpts, _rewardToken)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardToken) view returns(uint256)
func (_VectorRewarder *VectorRewarderCallerSession) RewardPerToken(_rewardToken common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.RewardPerToken(&_VectorRewarder.CallOpts, _rewardToken)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_VectorRewarder *VectorRewarderCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "rewardTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_VectorRewarder *VectorRewarderSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _VectorRewarder.Contract.RewardTokens(&_VectorRewarder.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_VectorRewarder *VectorRewarderCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _VectorRewarder.Contract.RewardTokens(&_VectorRewarder.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(address rewardToken, uint256 rewardPerTokenStored, uint256 queuedRewards, uint256 historicalRewards)
func (_VectorRewarder *VectorRewarderCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (struct {
	RewardToken          common.Address
	RewardPerTokenStored *big.Int
	QueuedRewards        *big.Int
	HistoricalRewards    *big.Int
}, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "rewards", arg0)

	outstruct := new(struct {
		RewardToken          common.Address
		RewardPerTokenStored *big.Int
		QueuedRewards        *big.Int
		HistoricalRewards    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RewardPerTokenStored = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.QueuedRewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.HistoricalRewards = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(address rewardToken, uint256 rewardPerTokenStored, uint256 queuedRewards, uint256 historicalRewards)
func (_VectorRewarder *VectorRewarderSession) Rewards(arg0 common.Address) (struct {
	RewardToken          common.Address
	RewardPerTokenStored *big.Int
	QueuedRewards        *big.Int
	HistoricalRewards    *big.Int
}, error) {
	return _VectorRewarder.Contract.Rewards(&_VectorRewarder.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(address rewardToken, uint256 rewardPerTokenStored, uint256 queuedRewards, uint256 historicalRewards)
func (_VectorRewarder *VectorRewarderCallerSession) Rewards(arg0 common.Address) (struct {
	RewardToken          common.Address
	RewardPerTokenStored *big.Int
	QueuedRewards        *big.Int
	HistoricalRewards    *big.Int
}, error) {
	return _VectorRewarder.Contract.Rewards(&_VectorRewarder.CallOpts, arg0)
}

// StakingDecimals is a free data retrieval call binding the contract method 0xc11f5344.
//
// Solidity: function stakingDecimals() view returns(uint256)
func (_VectorRewarder *VectorRewarderCaller) StakingDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "stakingDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingDecimals is a free data retrieval call binding the contract method 0xc11f5344.
//
// Solidity: function stakingDecimals() view returns(uint256)
func (_VectorRewarder *VectorRewarderSession) StakingDecimals() (*big.Int, error) {
	return _VectorRewarder.Contract.StakingDecimals(&_VectorRewarder.CallOpts)
}

// StakingDecimals is a free data retrieval call binding the contract method 0xc11f5344.
//
// Solidity: function stakingDecimals() view returns(uint256)
func (_VectorRewarder *VectorRewarderCallerSession) StakingDecimals() (*big.Int, error) {
	return _VectorRewarder.Contract.StakingDecimals(&_VectorRewarder.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_VectorRewarder *VectorRewarderCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_VectorRewarder *VectorRewarderSession) StakingToken() (common.Address, error) {
	return _VectorRewarder.Contract.StakingToken(&_VectorRewarder.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_VectorRewarder *VectorRewarderCallerSession) StakingToken() (common.Address, error) {
	return _VectorRewarder.Contract.StakingToken(&_VectorRewarder.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VectorRewarder *VectorRewarderCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VectorRewarder *VectorRewarderSession) TotalSupply() (*big.Int, error) {
	return _VectorRewarder.Contract.TotalSupply(&_VectorRewarder.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VectorRewarder *VectorRewarderCallerSession) TotalSupply() (*big.Int, error) {
	return _VectorRewarder.Contract.TotalSupply(&_VectorRewarder.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_VectorRewarder *VectorRewarderCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_VectorRewarder *VectorRewarderSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.UserRewardPerTokenPaid(&_VectorRewarder.CallOpts, arg0, arg1)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_VectorRewarder *VectorRewarderCallerSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.UserRewardPerTokenPaid(&_VectorRewarder.CallOpts, arg0, arg1)
}

// UserRewards is a free data retrieval call binding the contract method 0xa980356a.
//
// Solidity: function userRewards(address , address ) view returns(uint256)
func (_VectorRewarder *VectorRewarderCaller) UserRewards(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VectorRewarder.contract.Call(opts, &out, "userRewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewards is a free data retrieval call binding the contract method 0xa980356a.
//
// Solidity: function userRewards(address , address ) view returns(uint256)
func (_VectorRewarder *VectorRewarderSession) UserRewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.UserRewards(&_VectorRewarder.CallOpts, arg0, arg1)
}

// UserRewards is a free data retrieval call binding the contract method 0xa980356a.
//
// Solidity: function userRewards(address , address ) view returns(uint256)
func (_VectorRewarder *VectorRewarderCallerSession) UserRewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VectorRewarder.Contract.UserRewards(&_VectorRewarder.CallOpts, arg0, arg1)
}

// DonateRewards is a paid mutator transaction binding the contract method 0x5bc59ce7.
//
// Solidity: function donateRewards(uint256 _amountReward, address _rewardToken) returns(bool)
func (_VectorRewarder *VectorRewarderTransactor) DonateRewards(opts *bind.TransactOpts, _amountReward *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _VectorRewarder.contract.Transact(opts, "donateRewards", _amountReward, _rewardToken)
}

// DonateRewards is a paid mutator transaction binding the contract method 0x5bc59ce7.
//
// Solidity: function donateRewards(uint256 _amountReward, address _rewardToken) returns(bool)
func (_VectorRewarder *VectorRewarderSession) DonateRewards(_amountReward *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _VectorRewarder.Contract.DonateRewards(&_VectorRewarder.TransactOpts, _amountReward, _rewardToken)
}

// DonateRewards is a paid mutator transaction binding the contract method 0x5bc59ce7.
//
// Solidity: function donateRewards(uint256 _amountReward, address _rewardToken) returns(bool)
func (_VectorRewarder *VectorRewarderTransactorSession) DonateRewards(_amountReward *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _VectorRewarder.Contract.DonateRewards(&_VectorRewarder.TransactOpts, _amountReward, _rewardToken)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns(bool)
func (_VectorRewarder *VectorRewarderTransactor) GetReward(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _VectorRewarder.contract.Transact(opts, "getReward", _account)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns(bool)
func (_VectorRewarder *VectorRewarderSession) GetReward(_account common.Address) (*types.Transaction, error) {
	return _VectorRewarder.Contract.GetReward(&_VectorRewarder.TransactOpts, _account)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns(bool)
func (_VectorRewarder *VectorRewarderTransactorSession) GetReward(_account common.Address) (*types.Transaction, error) {
	return _VectorRewarder.Contract.GetReward(&_VectorRewarder.TransactOpts, _account)
}

// GetRewardUser is a paid mutator transaction binding the contract method 0x18838d33.
//
// Solidity: function getRewardUser() returns(bool)
func (_VectorRewarder *VectorRewarderTransactor) GetRewardUser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorRewarder.contract.Transact(opts, "getRewardUser")
}

// GetRewardUser is a paid mutator transaction binding the contract method 0x18838d33.
//
// Solidity: function getRewardUser() returns(bool)
func (_VectorRewarder *VectorRewarderSession) GetRewardUser() (*types.Transaction, error) {
	return _VectorRewarder.Contract.GetRewardUser(&_VectorRewarder.TransactOpts)
}

// GetRewardUser is a paid mutator transaction binding the contract method 0x18838d33.
//
// Solidity: function getRewardUser() returns(bool)
func (_VectorRewarder *VectorRewarderTransactorSession) GetRewardUser() (*types.Transaction, error) {
	return _VectorRewarder.Contract.GetRewardUser(&_VectorRewarder.TransactOpts)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x8fcf4822.
//
// Solidity: function queueNewRewards(uint256 _amountReward, address _rewardToken) returns(bool)
func (_VectorRewarder *VectorRewarderTransactor) QueueNewRewards(opts *bind.TransactOpts, _amountReward *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _VectorRewarder.contract.Transact(opts, "queueNewRewards", _amountReward, _rewardToken)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x8fcf4822.
//
// Solidity: function queueNewRewards(uint256 _amountReward, address _rewardToken) returns(bool)
func (_VectorRewarder *VectorRewarderSession) QueueNewRewards(_amountReward *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _VectorRewarder.Contract.QueueNewRewards(&_VectorRewarder.TransactOpts, _amountReward, _rewardToken)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x8fcf4822.
//
// Solidity: function queueNewRewards(uint256 _amountReward, address _rewardToken) returns(bool)
func (_VectorRewarder *VectorRewarderTransactorSession) QueueNewRewards(_amountReward *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _VectorRewarder.Contract.QueueNewRewards(&_VectorRewarder.TransactOpts, _amountReward, _rewardToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VectorRewarder *VectorRewarderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorRewarder.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VectorRewarder *VectorRewarderSession) RenounceOwnership() (*types.Transaction, error) {
	return _VectorRewarder.Contract.RenounceOwnership(&_VectorRewarder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VectorRewarder *VectorRewarderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VectorRewarder.Contract.RenounceOwnership(&_VectorRewarder.TransactOpts)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _for, uint256 _amount) returns(bool)
func (_VectorRewarder *VectorRewarderTransactor) StakeFor(opts *bind.TransactOpts, _for common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VectorRewarder.contract.Transact(opts, "stakeFor", _for, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _for, uint256 _amount) returns(bool)
func (_VectorRewarder *VectorRewarderSession) StakeFor(_for common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VectorRewarder.Contract.StakeFor(&_VectorRewarder.TransactOpts, _for, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _for, uint256 _amount) returns(bool)
func (_VectorRewarder *VectorRewarderTransactorSession) StakeFor(_for common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VectorRewarder.Contract.StakeFor(&_VectorRewarder.TransactOpts, _for, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VectorRewarder *VectorRewarderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VectorRewarder.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VectorRewarder *VectorRewarderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VectorRewarder.Contract.TransferOwnership(&_VectorRewarder.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VectorRewarder *VectorRewarderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VectorRewarder.Contract.TransferOwnership(&_VectorRewarder.TransactOpts, newOwner)
}

// UpdateFor is a paid mutator transaction binding the contract method 0x0e0a5968.
//
// Solidity: function updateFor(address _account) returns()
func (_VectorRewarder *VectorRewarderTransactor) UpdateFor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _VectorRewarder.contract.Transact(opts, "updateFor", _account)
}

// UpdateFor is a paid mutator transaction binding the contract method 0x0e0a5968.
//
// Solidity: function updateFor(address _account) returns()
func (_VectorRewarder *VectorRewarderSession) UpdateFor(_account common.Address) (*types.Transaction, error) {
	return _VectorRewarder.Contract.UpdateFor(&_VectorRewarder.TransactOpts, _account)
}

// UpdateFor is a paid mutator transaction binding the contract method 0x0e0a5968.
//
// Solidity: function updateFor(address _account) returns()
func (_VectorRewarder *VectorRewarderTransactorSession) UpdateFor(_account common.Address) (*types.Transaction, error) {
	return _VectorRewarder.Contract.UpdateFor(&_VectorRewarder.TransactOpts, _account)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x0e19c699.
//
// Solidity: function withdrawFor(address _for, uint256 _amount, bool claim) returns(bool)
func (_VectorRewarder *VectorRewarderTransactor) WithdrawFor(opts *bind.TransactOpts, _for common.Address, _amount *big.Int, claim bool) (*types.Transaction, error) {
	return _VectorRewarder.contract.Transact(opts, "withdrawFor", _for, _amount, claim)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x0e19c699.
//
// Solidity: function withdrawFor(address _for, uint256 _amount, bool claim) returns(bool)
func (_VectorRewarder *VectorRewarderSession) WithdrawFor(_for common.Address, _amount *big.Int, claim bool) (*types.Transaction, error) {
	return _VectorRewarder.Contract.WithdrawFor(&_VectorRewarder.TransactOpts, _for, _amount, claim)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x0e19c699.
//
// Solidity: function withdrawFor(address _for, uint256 _amount, bool claim) returns(bool)
func (_VectorRewarder *VectorRewarderTransactorSession) WithdrawFor(_for common.Address, _amount *big.Int, claim bool) (*types.Transaction, error) {
	return _VectorRewarder.Contract.WithdrawFor(&_VectorRewarder.TransactOpts, _for, _amount, claim)
}
