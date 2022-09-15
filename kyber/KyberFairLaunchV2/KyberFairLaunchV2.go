// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KyberFairLaunchV2

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

// KyberFairLaunchV2MetaData contains all meta data concerning the KyberFairLaunchV2 contract.
var KyberFairLaunchV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_endTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_vestingDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint256[]\",\"name\":\"_totalRewards\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"_tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenSymbol\",\"type\":\"string\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardTokenIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_shouldHarvest\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"generatedToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastRewardTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"vestingDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint256[]\",\"name\":\"rewardPerSeconds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewardMultipliers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"accRewardPerShares\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"unclaimedRewards\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lastRewardPerShares\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"name\":\"harvestMultiplePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"multipliers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"memoryrewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_endTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_vestingDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint256[]\",\"name\":\"_totalRewards\",\"type\":\"uint256[]\"}],\"name\":\"renewPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardLocker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_endTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_vestingDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint256[]\",\"name\":\"_totalRewards\",\"type\":\"uint256[]\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePoolRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KyberFairLaunchV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use KyberFairLaunchV2MetaData.ABI instead.
var KyberFairLaunchV2ABI = KyberFairLaunchV2MetaData.ABI

// KyberFairLaunchV2 is an auto generated Go binding around an Ethereum contract.
type KyberFairLaunchV2 struct {
	KyberFairLaunchV2Caller     // Read-only binding to the contract
	KyberFairLaunchV2Transactor // Write-only binding to the contract
	KyberFairLaunchV2Filterer   // Log filterer for contract events
}

// KyberFairLaunchV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type KyberFairLaunchV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberFairLaunchV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberFairLaunchV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberFairLaunchV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberFairLaunchV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberFairLaunchV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberFairLaunchV2Session struct {
	Contract     *KyberFairLaunchV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KyberFairLaunchV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberFairLaunchV2CallerSession struct {
	Contract *KyberFairLaunchV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// KyberFairLaunchV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberFairLaunchV2TransactorSession struct {
	Contract     *KyberFairLaunchV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// KyberFairLaunchV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type KyberFairLaunchV2Raw struct {
	Contract *KyberFairLaunchV2 // Generic contract binding to access the raw methods on
}

// KyberFairLaunchV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberFairLaunchV2CallerRaw struct {
	Contract *KyberFairLaunchV2Caller // Generic read-only contract binding to access the raw methods on
}

// KyberFairLaunchV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberFairLaunchV2TransactorRaw struct {
	Contract *KyberFairLaunchV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberFairLaunchV2 creates a new instance of KyberFairLaunchV2, bound to a specific deployed contract.
func NewKyberFairLaunchV2(address common.Address, backend bind.ContractBackend) (*KyberFairLaunchV2, error) {
	contract, err := bindKyberFairLaunchV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KyberFairLaunchV2{KyberFairLaunchV2Caller: KyberFairLaunchV2Caller{contract: contract}, KyberFairLaunchV2Transactor: KyberFairLaunchV2Transactor{contract: contract}, KyberFairLaunchV2Filterer: KyberFairLaunchV2Filterer{contract: contract}}, nil
}

// NewKyberFairLaunchV2Caller creates a new read-only instance of KyberFairLaunchV2, bound to a specific deployed contract.
func NewKyberFairLaunchV2Caller(address common.Address, caller bind.ContractCaller) (*KyberFairLaunchV2Caller, error) {
	contract, err := bindKyberFairLaunchV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberFairLaunchV2Caller{contract: contract}, nil
}

// NewKyberFairLaunchV2Transactor creates a new write-only instance of KyberFairLaunchV2, bound to a specific deployed contract.
func NewKyberFairLaunchV2Transactor(address common.Address, transactor bind.ContractTransactor) (*KyberFairLaunchV2Transactor, error) {
	contract, err := bindKyberFairLaunchV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberFairLaunchV2Transactor{contract: contract}, nil
}

// NewKyberFairLaunchV2Filterer creates a new log filterer instance of KyberFairLaunchV2, bound to a specific deployed contract.
func NewKyberFairLaunchV2Filterer(address common.Address, filterer bind.ContractFilterer) (*KyberFairLaunchV2Filterer, error) {
	contract, err := bindKyberFairLaunchV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberFairLaunchV2Filterer{contract: contract}, nil
}

// bindKyberFairLaunchV2 binds a generic wrapper to an already deployed contract.
func bindKyberFairLaunchV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberFairLaunchV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberFairLaunchV2 *KyberFairLaunchV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberFairLaunchV2.Contract.KyberFairLaunchV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberFairLaunchV2 *KyberFairLaunchV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.KyberFairLaunchV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberFairLaunchV2 *KyberFairLaunchV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.KyberFairLaunchV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberFairLaunchV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Caller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KyberFairLaunchV2.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) Admin() (common.Address, error) {
	return _KyberFairLaunchV2.Contract.Admin(&_KyberFairLaunchV2.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerSession) Admin() (common.Address, error) {
	return _KyberFairLaunchV2.Contract.Admin(&_KyberFairLaunchV2.CallOpts)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 _pid) view returns(uint256 totalStake, address stakeToken, address generatedToken, uint32 startTime, uint32 endTime, uint32 lastRewardTime, uint32 vestingDuration, uint256[] rewardPerSeconds, uint256[] rewardMultipliers, uint256[] accRewardPerShares)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Caller) GetPoolInfo(opts *bind.CallOpts, _pid *big.Int) (struct {
	TotalStake         *big.Int
	StakeToken         common.Address
	GeneratedToken     common.Address
	StartTime          uint32
	EndTime            uint32
	LastRewardTime     uint32
	VestingDuration    uint32
	RewardPerSeconds   []*big.Int
	RewardMultipliers  []*big.Int
	AccRewardPerShares []*big.Int
}, error) {
	var out []interface{}
	err := _KyberFairLaunchV2.contract.Call(opts, &out, "getPoolInfo", _pid)

	outstruct := new(struct {
		TotalStake         *big.Int
		StakeToken         common.Address
		GeneratedToken     common.Address
		StartTime          uint32
		EndTime            uint32
		LastRewardTime     uint32
		VestingDuration    uint32
		RewardPerSeconds   []*big.Int
		RewardMultipliers  []*big.Int
		AccRewardPerShares []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalStake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StakeToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.GeneratedToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.StartTime = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.EndTime = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.LastRewardTime = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.VestingDuration = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.RewardPerSeconds = *abi.ConvertType(out[7], new([]*big.Int)).(*[]*big.Int)
	outstruct.RewardMultipliers = *abi.ConvertType(out[8], new([]*big.Int)).(*[]*big.Int)
	outstruct.AccRewardPerShares = *abi.ConvertType(out[9], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 _pid) view returns(uint256 totalStake, address stakeToken, address generatedToken, uint32 startTime, uint32 endTime, uint32 lastRewardTime, uint32 vestingDuration, uint256[] rewardPerSeconds, uint256[] rewardMultipliers, uint256[] accRewardPerShares)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) GetPoolInfo(_pid *big.Int) (struct {
	TotalStake         *big.Int
	StakeToken         common.Address
	GeneratedToken     common.Address
	StartTime          uint32
	EndTime            uint32
	LastRewardTime     uint32
	VestingDuration    uint32
	RewardPerSeconds   []*big.Int
	RewardMultipliers  []*big.Int
	AccRewardPerShares []*big.Int
}, error) {
	return _KyberFairLaunchV2.Contract.GetPoolInfo(&_KyberFairLaunchV2.CallOpts, _pid)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 _pid) view returns(uint256 totalStake, address stakeToken, address generatedToken, uint32 startTime, uint32 endTime, uint32 lastRewardTime, uint32 vestingDuration, uint256[] rewardPerSeconds, uint256[] rewardMultipliers, uint256[] accRewardPerShares)
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerSession) GetPoolInfo(_pid *big.Int) (struct {
	TotalStake         *big.Int
	StakeToken         common.Address
	GeneratedToken     common.Address
	StartTime          uint32
	EndTime            uint32
	LastRewardTime     uint32
	VestingDuration    uint32
	RewardPerSeconds   []*big.Int
	RewardMultipliers  []*big.Int
	AccRewardPerShares []*big.Int
}, error) {
	return _KyberFairLaunchV2.Contract.GetPoolInfo(&_KyberFairLaunchV2.CallOpts, _pid)
}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_KyberFairLaunchV2 *KyberFairLaunchV2Caller) GetRewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _KyberFairLaunchV2.contract.Call(opts, &out, "getRewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) GetRewardTokens() ([]common.Address, error) {
	return _KyberFairLaunchV2.Contract.GetRewardTokens(&_KyberFairLaunchV2.CallOpts)
}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerSession) GetRewardTokens() ([]common.Address, error) {
	return _KyberFairLaunchV2.Contract.GetRewardTokens(&_KyberFairLaunchV2.CallOpts)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x1069f3b5.
//
// Solidity: function getUserInfo(uint256 _pid, address _account) view returns(uint256 amount, uint256[] unclaimedRewards, uint256[] lastRewardPerShares)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Caller) GetUserInfo(opts *bind.CallOpts, _pid *big.Int, _account common.Address) (struct {
	Amount              *big.Int
	UnclaimedRewards    []*big.Int
	LastRewardPerShares []*big.Int
}, error) {
	var out []interface{}
	err := _KyberFairLaunchV2.contract.Call(opts, &out, "getUserInfo", _pid, _account)

	outstruct := new(struct {
		Amount              *big.Int
		UnclaimedRewards    []*big.Int
		LastRewardPerShares []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnclaimedRewards = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.LastRewardPerShares = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x1069f3b5.
//
// Solidity: function getUserInfo(uint256 _pid, address _account) view returns(uint256 amount, uint256[] unclaimedRewards, uint256[] lastRewardPerShares)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) GetUserInfo(_pid *big.Int, _account common.Address) (struct {
	Amount              *big.Int
	UnclaimedRewards    []*big.Int
	LastRewardPerShares []*big.Int
}, error) {
	return _KyberFairLaunchV2.Contract.GetUserInfo(&_KyberFairLaunchV2.CallOpts, _pid, _account)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x1069f3b5.
//
// Solidity: function getUserInfo(uint256 _pid, address _account) view returns(uint256 amount, uint256[] unclaimedRewards, uint256[] lastRewardPerShares)
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerSession) GetUserInfo(_pid *big.Int, _account common.Address) (struct {
	Amount              *big.Int
	UnclaimedRewards    []*big.Int
	LastRewardPerShares []*big.Int
}, error) {
	return _KyberFairLaunchV2.Contract.GetUserInfo(&_KyberFairLaunchV2.CallOpts, _pid, _account)
}

// Multipliers is a free data retrieval call binding the contract method 0xac7fc263.
//
// Solidity: function multipliers(uint256 ) view returns(uint256)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Caller) Multipliers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KyberFairLaunchV2.contract.Call(opts, &out, "multipliers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multipliers is a free data retrieval call binding the contract method 0xac7fc263.
//
// Solidity: function multipliers(uint256 ) view returns(uint256)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) Multipliers(arg0 *big.Int) (*big.Int, error) {
	return _KyberFairLaunchV2.Contract.Multipliers(&_KyberFairLaunchV2.CallOpts, arg0)
}

// Multipliers is a free data retrieval call binding the contract method 0xac7fc263.
//
// Solidity: function multipliers(uint256 ) view returns(uint256)
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerSession) Multipliers(arg0 *big.Int) (*big.Int, error) {
	return _KyberFairLaunchV2.Contract.Multipliers(&_KyberFairLaunchV2.CallOpts, arg0)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Caller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KyberFairLaunchV2.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) PendingAdmin() (common.Address, error) {
	return _KyberFairLaunchV2.Contract.PendingAdmin(&_KyberFairLaunchV2.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerSession) PendingAdmin() (common.Address, error) {
	return _KyberFairLaunchV2.Contract.PendingAdmin(&_KyberFairLaunchV2.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0xd18df53c.
//
// Solidity: function pendingRewards(uint256 _pid, address _user) view returns(uint256[] memoryrewards)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Caller) PendingRewards(opts *bind.CallOpts, _pid *big.Int, _user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _KyberFairLaunchV2.contract.Call(opts, &out, "pendingRewards", _pid, _user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0xd18df53c.
//
// Solidity: function pendingRewards(uint256 _pid, address _user) view returns(uint256[] memoryrewards)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) PendingRewards(_pid *big.Int, _user common.Address) ([]*big.Int, error) {
	return _KyberFairLaunchV2.Contract.PendingRewards(&_KyberFairLaunchV2.CallOpts, _pid, _user)
}

// PendingRewards is a free data retrieval call binding the contract method 0xd18df53c.
//
// Solidity: function pendingRewards(uint256 _pid, address _user) view returns(uint256[] memoryrewards)
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerSession) PendingRewards(_pid *big.Int, _user common.Address) ([]*big.Int, error) {
	return _KyberFairLaunchV2.Contract.PendingRewards(&_KyberFairLaunchV2.CallOpts, _pid, _user)
}

// PoolExists is a free data retrieval call binding the contract method 0x1e1c6a07.
//
// Solidity: function poolExists(address ) view returns(bool)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Caller) PoolExists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _KyberFairLaunchV2.contract.Call(opts, &out, "poolExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolExists is a free data retrieval call binding the contract method 0x1e1c6a07.
//
// Solidity: function poolExists(address ) view returns(bool)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) PoolExists(arg0 common.Address) (bool, error) {
	return _KyberFairLaunchV2.Contract.PoolExists(&_KyberFairLaunchV2.CallOpts, arg0)
}

// PoolExists is a free data retrieval call binding the contract method 0x1e1c6a07.
//
// Solidity: function poolExists(address ) view returns(bool)
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerSession) PoolExists(arg0 common.Address) (bool, error) {
	return _KyberFairLaunchV2.Contract.PoolExists(&_KyberFairLaunchV2.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KyberFairLaunchV2.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) PoolLength() (*big.Int, error) {
	return _KyberFairLaunchV2.Contract.PoolLength(&_KyberFairLaunchV2.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerSession) PoolLength() (*big.Int, error) {
	return _KyberFairLaunchV2.Contract.PoolLength(&_KyberFairLaunchV2.CallOpts)
}

// RewardLocker is a free data retrieval call binding the contract method 0x3892601c.
//
// Solidity: function rewardLocker() view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Caller) RewardLocker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KyberFairLaunchV2.contract.Call(opts, &out, "rewardLocker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardLocker is a free data retrieval call binding the contract method 0x3892601c.
//
// Solidity: function rewardLocker() view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) RewardLocker() (common.Address, error) {
	return _KyberFairLaunchV2.Contract.RewardLocker(&_KyberFairLaunchV2.CallOpts)
}

// RewardLocker is a free data retrieval call binding the contract method 0x3892601c.
//
// Solidity: function rewardLocker() view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerSession) RewardLocker() (common.Address, error) {
	return _KyberFairLaunchV2.Contract.RewardLocker(&_KyberFairLaunchV2.CallOpts)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Caller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KyberFairLaunchV2.contract.Call(opts, &out, "rewardTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _KyberFairLaunchV2.Contract.RewardTokens(&_KyberFairLaunchV2.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_KyberFairLaunchV2 *KyberFairLaunchV2CallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _KyberFairLaunchV2.Contract.RewardTokens(&_KyberFairLaunchV2.CallOpts, arg0)
}

// AddPool is a paid mutator transaction binding the contract method 0xfef12687.
//
// Solidity: function addPool(address _stakeToken, uint32 _startTime, uint32 _endTime, uint32 _vestingDuration, uint256[] _totalRewards, string _tokenName, string _tokenSymbol) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) AddPool(opts *bind.TransactOpts, _stakeToken common.Address, _startTime uint32, _endTime uint32, _vestingDuration uint32, _totalRewards []*big.Int, _tokenName string, _tokenSymbol string) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "addPool", _stakeToken, _startTime, _endTime, _vestingDuration, _totalRewards, _tokenName, _tokenSymbol)
}

// AddPool is a paid mutator transaction binding the contract method 0xfef12687.
//
// Solidity: function addPool(address _stakeToken, uint32 _startTime, uint32 _endTime, uint32 _vestingDuration, uint256[] _totalRewards, string _tokenName, string _tokenSymbol) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) AddPool(_stakeToken common.Address, _startTime uint32, _endTime uint32, _vestingDuration uint32, _totalRewards []*big.Int, _tokenName string, _tokenSymbol string) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.AddPool(&_KyberFairLaunchV2.TransactOpts, _stakeToken, _startTime, _endTime, _vestingDuration, _totalRewards, _tokenName, _tokenSymbol)
}

// AddPool is a paid mutator transaction binding the contract method 0xfef12687.
//
// Solidity: function addPool(address _stakeToken, uint32 _startTime, uint32 _endTime, uint32 _vestingDuration, uint256[] _totalRewards, string _tokenName, string _tokenSymbol) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) AddPool(_stakeToken common.Address, _startTime uint32, _endTime uint32, _vestingDuration uint32, _totalRewards []*big.Int, _tokenName string, _tokenSymbol string) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.AddPool(&_KyberFairLaunchV2.TransactOpts, _stakeToken, _startTime, _endTime, _vestingDuration, _totalRewards, _tokenName, _tokenSymbol)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0xd12e4ceb.
//
// Solidity: function adminWithdraw(uint256 rewardTokenIndex, uint256 amount) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) AdminWithdraw(opts *bind.TransactOpts, rewardTokenIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "adminWithdraw", rewardTokenIndex, amount)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0xd12e4ceb.
//
// Solidity: function adminWithdraw(uint256 rewardTokenIndex, uint256 amount) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) AdminWithdraw(rewardTokenIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.AdminWithdraw(&_KyberFairLaunchV2.TransactOpts, rewardTokenIndex, amount)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0xd12e4ceb.
//
// Solidity: function adminWithdraw(uint256 rewardTokenIndex, uint256 amount) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) AdminWithdraw(rewardTokenIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.AdminWithdraw(&_KyberFairLaunchV2.TransactOpts, rewardTokenIndex, amount)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) ClaimAdmin() (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.ClaimAdmin(&_KyberFairLaunchV2.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.ClaimAdmin(&_KyberFairLaunchV2.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _shouldHarvest) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _shouldHarvest bool) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "deposit", _pid, _amount, _shouldHarvest)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _shouldHarvest) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) Deposit(_pid *big.Int, _amount *big.Int, _shouldHarvest bool) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.Deposit(&_KyberFairLaunchV2.TransactOpts, _pid, _amount, _shouldHarvest)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _shouldHarvest) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) Deposit(_pid *big.Int, _amount *big.Int, _shouldHarvest bool) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.Deposit(&_KyberFairLaunchV2.TransactOpts, _pid, _amount, _shouldHarvest)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.EmergencyWithdraw(&_KyberFairLaunchV2.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.EmergencyWithdraw(&_KyberFairLaunchV2.TransactOpts, _pid)
}

// Harvest is a paid mutator transaction binding the contract method 0xddc63262.
//
// Solidity: function harvest(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) Harvest(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "harvest", _pid)
}

// Harvest is a paid mutator transaction binding the contract method 0xddc63262.
//
// Solidity: function harvest(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) Harvest(_pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.Harvest(&_KyberFairLaunchV2.TransactOpts, _pid)
}

// Harvest is a paid mutator transaction binding the contract method 0xddc63262.
//
// Solidity: function harvest(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) Harvest(_pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.Harvest(&_KyberFairLaunchV2.TransactOpts, _pid)
}

// HarvestMultiplePools is a paid mutator transaction binding the contract method 0x52fd9f2f.
//
// Solidity: function harvestMultiplePools(uint256[] _pids) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) HarvestMultiplePools(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "harvestMultiplePools", _pids)
}

// HarvestMultiplePools is a paid mutator transaction binding the contract method 0x52fd9f2f.
//
// Solidity: function harvestMultiplePools(uint256[] _pids) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) HarvestMultiplePools(_pids []*big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.HarvestMultiplePools(&_KyberFairLaunchV2.TransactOpts, _pids)
}

// HarvestMultiplePools is a paid mutator transaction binding the contract method 0x52fd9f2f.
//
// Solidity: function harvestMultiplePools(uint256[] _pids) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) HarvestMultiplePools(_pids []*big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.HarvestMultiplePools(&_KyberFairLaunchV2.TransactOpts, _pids)
}

// RenewPool is a paid mutator transaction binding the contract method 0xc401fb4f.
//
// Solidity: function renewPool(uint256 _pid, uint32 _startTime, uint32 _endTime, uint32 _vestingDuration, uint256[] _totalRewards) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) RenewPool(opts *bind.TransactOpts, _pid *big.Int, _startTime uint32, _endTime uint32, _vestingDuration uint32, _totalRewards []*big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "renewPool", _pid, _startTime, _endTime, _vestingDuration, _totalRewards)
}

// RenewPool is a paid mutator transaction binding the contract method 0xc401fb4f.
//
// Solidity: function renewPool(uint256 _pid, uint32 _startTime, uint32 _endTime, uint32 _vestingDuration, uint256[] _totalRewards) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) RenewPool(_pid *big.Int, _startTime uint32, _endTime uint32, _vestingDuration uint32, _totalRewards []*big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.RenewPool(&_KyberFairLaunchV2.TransactOpts, _pid, _startTime, _endTime, _vestingDuration, _totalRewards)
}

// RenewPool is a paid mutator transaction binding the contract method 0xc401fb4f.
//
// Solidity: function renewPool(uint256 _pid, uint32 _startTime, uint32 _endTime, uint32 _vestingDuration, uint256[] _totalRewards) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) RenewPool(_pid *big.Int, _startTime uint32, _endTime uint32, _vestingDuration uint32, _totalRewards []*big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.RenewPool(&_KyberFairLaunchV2.TransactOpts, _pid, _startTime, _endTime, _vestingDuration, _totalRewards)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.TransferAdmin(&_KyberFairLaunchV2.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.TransferAdmin(&_KyberFairLaunchV2.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.TransferAdminQuickly(&_KyberFairLaunchV2.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.TransferAdminQuickly(&_KyberFairLaunchV2.TransactOpts, newAdmin)
}

// UpdatePool is a paid mutator transaction binding the contract method 0xa3a19dfb.
//
// Solidity: function updatePool(uint256 _pid, uint32 _endTime, uint32 _vestingDuration, uint256[] _totalRewards) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int, _endTime uint32, _vestingDuration uint32, _totalRewards []*big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "updatePool", _pid, _endTime, _vestingDuration, _totalRewards)
}

// UpdatePool is a paid mutator transaction binding the contract method 0xa3a19dfb.
//
// Solidity: function updatePool(uint256 _pid, uint32 _endTime, uint32 _vestingDuration, uint256[] _totalRewards) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) UpdatePool(_pid *big.Int, _endTime uint32, _vestingDuration uint32, _totalRewards []*big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.UpdatePool(&_KyberFairLaunchV2.TransactOpts, _pid, _endTime, _vestingDuration, _totalRewards)
}

// UpdatePool is a paid mutator transaction binding the contract method 0xa3a19dfb.
//
// Solidity: function updatePool(uint256 _pid, uint32 _endTime, uint32 _vestingDuration, uint256[] _totalRewards) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) UpdatePool(_pid *big.Int, _endTime uint32, _vestingDuration uint32, _totalRewards []*big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.UpdatePool(&_KyberFairLaunchV2.TransactOpts, _pid, _endTime, _vestingDuration, _totalRewards)
}

// UpdatePoolRewards is a paid mutator transaction binding the contract method 0xadb82b31.
//
// Solidity: function updatePoolRewards(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) UpdatePoolRewards(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "updatePoolRewards", _pid)
}

// UpdatePoolRewards is a paid mutator transaction binding the contract method 0xadb82b31.
//
// Solidity: function updatePoolRewards(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) UpdatePoolRewards(_pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.UpdatePoolRewards(&_KyberFairLaunchV2.TransactOpts, _pid)
}

// UpdatePoolRewards is a paid mutator transaction binding the contract method 0xadb82b31.
//
// Solidity: function updatePoolRewards(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) UpdatePoolRewards(_pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.UpdatePoolRewards(&_KyberFairLaunchV2.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.Withdraw(&_KyberFairLaunchV2.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.Withdraw(&_KyberFairLaunchV2.TransactOpts, _pid, _amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Transactor) WithdrawAll(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.contract.Transact(opts, "withdrawAll", _pid)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2Session) WithdrawAll(_pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.WithdrawAll(&_KyberFairLaunchV2.TransactOpts, _pid)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns()
func (_KyberFairLaunchV2 *KyberFairLaunchV2TransactorSession) WithdrawAll(_pid *big.Int) (*types.Transaction, error) {
	return _KyberFairLaunchV2.Contract.WithdrawAll(&_KyberFairLaunchV2.TransactOpts, _pid)
}
