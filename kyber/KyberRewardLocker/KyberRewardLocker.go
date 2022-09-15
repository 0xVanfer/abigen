// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KyberRewardLocker

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

// IKyberRewardLockerVestingSchedule is an auto generated low-level Go binding around an user-defined struct.
type IKyberRewardLockerVestingSchedule struct {
	StartBlock     uint64
	EndBlock       uint64
	Quantity       *big.Int
	VestedQuantity *big.Int
}

// KyberRewardLockerMetaData contains all meta data concerning the KyberRewardLocker contract.
var KyberRewardLockerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountEscrowedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountVestedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardContract\",\"type\":\"address\"}],\"name\":\"addRewardsContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getRewardContractsPerToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"rewardContracts\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVestingScheduleAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"quantity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"vestedQuantity\",\"type\":\"uint128\"}],\"internalType\":\"structIKyberRewardLocker.VestingSchedule\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getVestingSchedules\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"quantity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"vestedQuantity\",\"type\":\"uint128\"}],\"internalType\":\"structIKyberRewardLocker.VestingSchedule[]\",\"name\":\"schedules\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"name\":\"lockWithStartBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"numVestingSchedules\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardContract\",\"type\":\"address\"}],\"name\":\"removeRewardsContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_vestingDuration\",\"type\":\"uint64\"}],\"name\":\"setVestingDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"vestCompletedSchedules\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"vestCompletedSchedulesForMultipleTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"vestedAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"vestScheduleAtIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"indices\",\"type\":\"uint256[][]\"}],\"name\":\"vestScheduleForMultipleTokensAtIndices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"vestedAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"vestSchedulesInRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vestingDurationPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// KyberRewardLockerABI is the input ABI used to generate the binding from.
// Deprecated: Use KyberRewardLockerMetaData.ABI instead.
var KyberRewardLockerABI = KyberRewardLockerMetaData.ABI

// KyberRewardLocker is an auto generated Go binding around an Ethereum contract.
type KyberRewardLocker struct {
	KyberRewardLockerCaller     // Read-only binding to the contract
	KyberRewardLockerTransactor // Write-only binding to the contract
	KyberRewardLockerFilterer   // Log filterer for contract events
}

// KyberRewardLockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberRewardLockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberRewardLockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberRewardLockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberRewardLockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberRewardLockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberRewardLockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberRewardLockerSession struct {
	Contract     *KyberRewardLocker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KyberRewardLockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberRewardLockerCallerSession struct {
	Contract *KyberRewardLockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// KyberRewardLockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberRewardLockerTransactorSession struct {
	Contract     *KyberRewardLockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// KyberRewardLockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberRewardLockerRaw struct {
	Contract *KyberRewardLocker // Generic contract binding to access the raw methods on
}

// KyberRewardLockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberRewardLockerCallerRaw struct {
	Contract *KyberRewardLockerCaller // Generic read-only contract binding to access the raw methods on
}

// KyberRewardLockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberRewardLockerTransactorRaw struct {
	Contract *KyberRewardLockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberRewardLocker creates a new instance of KyberRewardLocker, bound to a specific deployed contract.
func NewKyberRewardLocker(address common.Address, backend bind.ContractBackend) (*KyberRewardLocker, error) {
	contract, err := bindKyberRewardLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KyberRewardLocker{KyberRewardLockerCaller: KyberRewardLockerCaller{contract: contract}, KyberRewardLockerTransactor: KyberRewardLockerTransactor{contract: contract}, KyberRewardLockerFilterer: KyberRewardLockerFilterer{contract: contract}}, nil
}

// NewKyberRewardLockerCaller creates a new read-only instance of KyberRewardLocker, bound to a specific deployed contract.
func NewKyberRewardLockerCaller(address common.Address, caller bind.ContractCaller) (*KyberRewardLockerCaller, error) {
	contract, err := bindKyberRewardLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberRewardLockerCaller{contract: contract}, nil
}

// NewKyberRewardLockerTransactor creates a new write-only instance of KyberRewardLocker, bound to a specific deployed contract.
func NewKyberRewardLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberRewardLockerTransactor, error) {
	contract, err := bindKyberRewardLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberRewardLockerTransactor{contract: contract}, nil
}

// NewKyberRewardLockerFilterer creates a new log filterer instance of KyberRewardLocker, bound to a specific deployed contract.
func NewKyberRewardLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberRewardLockerFilterer, error) {
	contract, err := bindKyberRewardLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberRewardLockerFilterer{contract: contract}, nil
}

// bindKyberRewardLocker binds a generic wrapper to an already deployed contract.
func bindKyberRewardLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberRewardLockerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberRewardLocker *KyberRewardLockerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberRewardLocker.Contract.KyberRewardLockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberRewardLocker *KyberRewardLockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.KyberRewardLockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberRewardLocker *KyberRewardLockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.KyberRewardLockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberRewardLocker *KyberRewardLockerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberRewardLocker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberRewardLocker *KyberRewardLockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberRewardLocker *KyberRewardLockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.contract.Transact(opts, method, params...)
}

// AccountEscrowedBalance is a free data retrieval call binding the contract method 0x6e732b70.
//
// Solidity: function accountEscrowedBalance(address , address ) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerCaller) AccountEscrowedBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KyberRewardLocker.contract.Call(opts, &out, "accountEscrowedBalance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountEscrowedBalance is a free data retrieval call binding the contract method 0x6e732b70.
//
// Solidity: function accountEscrowedBalance(address , address ) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerSession) AccountEscrowedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _KyberRewardLocker.Contract.AccountEscrowedBalance(&_KyberRewardLocker.CallOpts, arg0, arg1)
}

// AccountEscrowedBalance is a free data retrieval call binding the contract method 0x6e732b70.
//
// Solidity: function accountEscrowedBalance(address , address ) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerCallerSession) AccountEscrowedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _KyberRewardLocker.Contract.AccountEscrowedBalance(&_KyberRewardLocker.CallOpts, arg0, arg1)
}

// AccountVestedBalance is a free data retrieval call binding the contract method 0x3b5bfa8b.
//
// Solidity: function accountVestedBalance(address , address ) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerCaller) AccountVestedBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KyberRewardLocker.contract.Call(opts, &out, "accountVestedBalance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountVestedBalance is a free data retrieval call binding the contract method 0x3b5bfa8b.
//
// Solidity: function accountVestedBalance(address , address ) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerSession) AccountVestedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _KyberRewardLocker.Contract.AccountVestedBalance(&_KyberRewardLocker.CallOpts, arg0, arg1)
}

// AccountVestedBalance is a free data retrieval call binding the contract method 0x3b5bfa8b.
//
// Solidity: function accountVestedBalance(address , address ) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerCallerSession) AccountVestedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _KyberRewardLocker.Contract.AccountVestedBalance(&_KyberRewardLocker.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_KyberRewardLocker *KyberRewardLockerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KyberRewardLocker.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_KyberRewardLocker *KyberRewardLockerSession) Admin() (common.Address, error) {
	return _KyberRewardLocker.Contract.Admin(&_KyberRewardLocker.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_KyberRewardLocker *KyberRewardLockerCallerSession) Admin() (common.Address, error) {
	return _KyberRewardLocker.Contract.Admin(&_KyberRewardLocker.CallOpts)
}

// GetRewardContractsPerToken is a free data retrieval call binding the contract method 0xbbd8a328.
//
// Solidity: function getRewardContractsPerToken(address token) view returns(address[] rewardContracts)
func (_KyberRewardLocker *KyberRewardLockerCaller) GetRewardContractsPerToken(opts *bind.CallOpts, token common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _KyberRewardLocker.contract.Call(opts, &out, "getRewardContractsPerToken", token)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardContractsPerToken is a free data retrieval call binding the contract method 0xbbd8a328.
//
// Solidity: function getRewardContractsPerToken(address token) view returns(address[] rewardContracts)
func (_KyberRewardLocker *KyberRewardLockerSession) GetRewardContractsPerToken(token common.Address) ([]common.Address, error) {
	return _KyberRewardLocker.Contract.GetRewardContractsPerToken(&_KyberRewardLocker.CallOpts, token)
}

// GetRewardContractsPerToken is a free data retrieval call binding the contract method 0xbbd8a328.
//
// Solidity: function getRewardContractsPerToken(address token) view returns(address[] rewardContracts)
func (_KyberRewardLocker *KyberRewardLockerCallerSession) GetRewardContractsPerToken(token common.Address) ([]common.Address, error) {
	return _KyberRewardLocker.Contract.GetRewardContractsPerToken(&_KyberRewardLocker.CallOpts, token)
}

// GetVestingScheduleAtIndex is a free data retrieval call binding the contract method 0x0a3b7e31.
//
// Solidity: function getVestingScheduleAtIndex(address account, address token, uint256 index) view returns((uint64,uint64,uint128,uint128))
func (_KyberRewardLocker *KyberRewardLockerCaller) GetVestingScheduleAtIndex(opts *bind.CallOpts, account common.Address, token common.Address, index *big.Int) (IKyberRewardLockerVestingSchedule, error) {
	var out []interface{}
	err := _KyberRewardLocker.contract.Call(opts, &out, "getVestingScheduleAtIndex", account, token, index)

	if err != nil {
		return *new(IKyberRewardLockerVestingSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new(IKyberRewardLockerVestingSchedule)).(*IKyberRewardLockerVestingSchedule)

	return out0, err

}

// GetVestingScheduleAtIndex is a free data retrieval call binding the contract method 0x0a3b7e31.
//
// Solidity: function getVestingScheduleAtIndex(address account, address token, uint256 index) view returns((uint64,uint64,uint128,uint128))
func (_KyberRewardLocker *KyberRewardLockerSession) GetVestingScheduleAtIndex(account common.Address, token common.Address, index *big.Int) (IKyberRewardLockerVestingSchedule, error) {
	return _KyberRewardLocker.Contract.GetVestingScheduleAtIndex(&_KyberRewardLocker.CallOpts, account, token, index)
}

// GetVestingScheduleAtIndex is a free data retrieval call binding the contract method 0x0a3b7e31.
//
// Solidity: function getVestingScheduleAtIndex(address account, address token, uint256 index) view returns((uint64,uint64,uint128,uint128))
func (_KyberRewardLocker *KyberRewardLockerCallerSession) GetVestingScheduleAtIndex(account common.Address, token common.Address, index *big.Int) (IKyberRewardLockerVestingSchedule, error) {
	return _KyberRewardLocker.Contract.GetVestingScheduleAtIndex(&_KyberRewardLocker.CallOpts, account, token, index)
}

// GetVestingSchedules is a free data retrieval call binding the contract method 0x2f50bd72.
//
// Solidity: function getVestingSchedules(address account, address token) view returns((uint64,uint64,uint128,uint128)[] schedules)
func (_KyberRewardLocker *KyberRewardLockerCaller) GetVestingSchedules(opts *bind.CallOpts, account common.Address, token common.Address) ([]IKyberRewardLockerVestingSchedule, error) {
	var out []interface{}
	err := _KyberRewardLocker.contract.Call(opts, &out, "getVestingSchedules", account, token)

	if err != nil {
		return *new([]IKyberRewardLockerVestingSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new([]IKyberRewardLockerVestingSchedule)).(*[]IKyberRewardLockerVestingSchedule)

	return out0, err

}

// GetVestingSchedules is a free data retrieval call binding the contract method 0x2f50bd72.
//
// Solidity: function getVestingSchedules(address account, address token) view returns((uint64,uint64,uint128,uint128)[] schedules)
func (_KyberRewardLocker *KyberRewardLockerSession) GetVestingSchedules(account common.Address, token common.Address) ([]IKyberRewardLockerVestingSchedule, error) {
	return _KyberRewardLocker.Contract.GetVestingSchedules(&_KyberRewardLocker.CallOpts, account, token)
}

// GetVestingSchedules is a free data retrieval call binding the contract method 0x2f50bd72.
//
// Solidity: function getVestingSchedules(address account, address token) view returns((uint64,uint64,uint128,uint128)[] schedules)
func (_KyberRewardLocker *KyberRewardLockerCallerSession) GetVestingSchedules(account common.Address, token common.Address) ([]IKyberRewardLockerVestingSchedule, error) {
	return _KyberRewardLocker.Contract.GetVestingSchedules(&_KyberRewardLocker.CallOpts, account, token)
}

// NumVestingSchedules is a free data retrieval call binding the contract method 0x679f7f77.
//
// Solidity: function numVestingSchedules(address account, address token) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerCaller) NumVestingSchedules(opts *bind.CallOpts, account common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KyberRewardLocker.contract.Call(opts, &out, "numVestingSchedules", account, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumVestingSchedules is a free data retrieval call binding the contract method 0x679f7f77.
//
// Solidity: function numVestingSchedules(address account, address token) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerSession) NumVestingSchedules(account common.Address, token common.Address) (*big.Int, error) {
	return _KyberRewardLocker.Contract.NumVestingSchedules(&_KyberRewardLocker.CallOpts, account, token)
}

// NumVestingSchedules is a free data retrieval call binding the contract method 0x679f7f77.
//
// Solidity: function numVestingSchedules(address account, address token) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerCallerSession) NumVestingSchedules(account common.Address, token common.Address) (*big.Int, error) {
	return _KyberRewardLocker.Contract.NumVestingSchedules(&_KyberRewardLocker.CallOpts, account, token)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_KyberRewardLocker *KyberRewardLockerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KyberRewardLocker.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_KyberRewardLocker *KyberRewardLockerSession) PendingAdmin() (common.Address, error) {
	return _KyberRewardLocker.Contract.PendingAdmin(&_KyberRewardLocker.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_KyberRewardLocker *KyberRewardLockerCallerSession) PendingAdmin() (common.Address, error) {
	return _KyberRewardLocker.Contract.PendingAdmin(&_KyberRewardLocker.CallOpts)
}

// VestingDurationPerToken is a free data retrieval call binding the contract method 0xaaf54385.
//
// Solidity: function vestingDurationPerToken(address ) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerCaller) VestingDurationPerToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KyberRewardLocker.contract.Call(opts, &out, "vestingDurationPerToken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingDurationPerToken is a free data retrieval call binding the contract method 0xaaf54385.
//
// Solidity: function vestingDurationPerToken(address ) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerSession) VestingDurationPerToken(arg0 common.Address) (*big.Int, error) {
	return _KyberRewardLocker.Contract.VestingDurationPerToken(&_KyberRewardLocker.CallOpts, arg0)
}

// VestingDurationPerToken is a free data retrieval call binding the contract method 0xaaf54385.
//
// Solidity: function vestingDurationPerToken(address ) view returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerCallerSession) VestingDurationPerToken(arg0 common.Address) (*big.Int, error) {
	return _KyberRewardLocker.Contract.VestingDurationPerToken(&_KyberRewardLocker.CallOpts, arg0)
}

// AddRewardsContract is a paid mutator transaction binding the contract method 0xf793d77e.
//
// Solidity: function addRewardsContract(address token, address _rewardContract) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactor) AddRewardsContract(opts *bind.TransactOpts, token common.Address, _rewardContract common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "addRewardsContract", token, _rewardContract)
}

// AddRewardsContract is a paid mutator transaction binding the contract method 0xf793d77e.
//
// Solidity: function addRewardsContract(address token, address _rewardContract) returns()
func (_KyberRewardLocker *KyberRewardLockerSession) AddRewardsContract(token common.Address, _rewardContract common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.AddRewardsContract(&_KyberRewardLocker.TransactOpts, token, _rewardContract)
}

// AddRewardsContract is a paid mutator transaction binding the contract method 0xf793d77e.
//
// Solidity: function addRewardsContract(address token, address _rewardContract) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) AddRewardsContract(token common.Address, _rewardContract common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.AddRewardsContract(&_KyberRewardLocker.TransactOpts, token, _rewardContract)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberRewardLocker *KyberRewardLockerTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberRewardLocker *KyberRewardLockerSession) ClaimAdmin() (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.ClaimAdmin(&_KyberRewardLocker.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.ClaimAdmin(&_KyberRewardLocker.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0x7750c9f0.
//
// Solidity: function lock(address token, address account, uint256 quantity) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactor) Lock(opts *bind.TransactOpts, token common.Address, account common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "lock", token, account, quantity)
}

// Lock is a paid mutator transaction binding the contract method 0x7750c9f0.
//
// Solidity: function lock(address token, address account, uint256 quantity) returns()
func (_KyberRewardLocker *KyberRewardLockerSession) Lock(token common.Address, account common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.Lock(&_KyberRewardLocker.TransactOpts, token, account, quantity)
}

// Lock is a paid mutator transaction binding the contract method 0x7750c9f0.
//
// Solidity: function lock(address token, address account, uint256 quantity) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) Lock(token common.Address, account common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.Lock(&_KyberRewardLocker.TransactOpts, token, account, quantity)
}

// LockWithStartBlock is a paid mutator transaction binding the contract method 0xa5d21e31.
//
// Solidity: function lockWithStartBlock(address token, address account, uint256 quantity, uint256 startBlock) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactor) LockWithStartBlock(opts *bind.TransactOpts, token common.Address, account common.Address, quantity *big.Int, startBlock *big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "lockWithStartBlock", token, account, quantity, startBlock)
}

// LockWithStartBlock is a paid mutator transaction binding the contract method 0xa5d21e31.
//
// Solidity: function lockWithStartBlock(address token, address account, uint256 quantity, uint256 startBlock) returns()
func (_KyberRewardLocker *KyberRewardLockerSession) LockWithStartBlock(token common.Address, account common.Address, quantity *big.Int, startBlock *big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.LockWithStartBlock(&_KyberRewardLocker.TransactOpts, token, account, quantity, startBlock)
}

// LockWithStartBlock is a paid mutator transaction binding the contract method 0xa5d21e31.
//
// Solidity: function lockWithStartBlock(address token, address account, uint256 quantity, uint256 startBlock) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) LockWithStartBlock(token common.Address, account common.Address, quantity *big.Int, startBlock *big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.LockWithStartBlock(&_KyberRewardLocker.TransactOpts, token, account, quantity, startBlock)
}

// RemoveRewardsContract is a paid mutator transaction binding the contract method 0x4c9d00cc.
//
// Solidity: function removeRewardsContract(address token, address _rewardContract) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactor) RemoveRewardsContract(opts *bind.TransactOpts, token common.Address, _rewardContract common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "removeRewardsContract", token, _rewardContract)
}

// RemoveRewardsContract is a paid mutator transaction binding the contract method 0x4c9d00cc.
//
// Solidity: function removeRewardsContract(address token, address _rewardContract) returns()
func (_KyberRewardLocker *KyberRewardLockerSession) RemoveRewardsContract(token common.Address, _rewardContract common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.RemoveRewardsContract(&_KyberRewardLocker.TransactOpts, token, _rewardContract)
}

// RemoveRewardsContract is a paid mutator transaction binding the contract method 0x4c9d00cc.
//
// Solidity: function removeRewardsContract(address token, address _rewardContract) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) RemoveRewardsContract(token common.Address, _rewardContract common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.RemoveRewardsContract(&_KyberRewardLocker.TransactOpts, token, _rewardContract)
}

// SetVestingDuration is a paid mutator transaction binding the contract method 0xcb1d086b.
//
// Solidity: function setVestingDuration(address token, uint64 _vestingDuration) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactor) SetVestingDuration(opts *bind.TransactOpts, token common.Address, _vestingDuration uint64) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "setVestingDuration", token, _vestingDuration)
}

// SetVestingDuration is a paid mutator transaction binding the contract method 0xcb1d086b.
//
// Solidity: function setVestingDuration(address token, uint64 _vestingDuration) returns()
func (_KyberRewardLocker *KyberRewardLockerSession) SetVestingDuration(token common.Address, _vestingDuration uint64) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.SetVestingDuration(&_KyberRewardLocker.TransactOpts, token, _vestingDuration)
}

// SetVestingDuration is a paid mutator transaction binding the contract method 0xcb1d086b.
//
// Solidity: function setVestingDuration(address token, uint64 _vestingDuration) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) SetVestingDuration(token common.Address, _vestingDuration uint64) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.SetVestingDuration(&_KyberRewardLocker.TransactOpts, token, _vestingDuration)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberRewardLocker *KyberRewardLockerSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.TransferAdmin(&_KyberRewardLocker.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.TransferAdmin(&_KyberRewardLocker.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberRewardLocker *KyberRewardLockerSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.TransferAdminQuickly(&_KyberRewardLocker.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.TransferAdminQuickly(&_KyberRewardLocker.TransactOpts, newAdmin)
}

// VestCompletedSchedules is a paid mutator transaction binding the contract method 0xfdfaaa05.
//
// Solidity: function vestCompletedSchedules(address token) returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerTransactor) VestCompletedSchedules(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "vestCompletedSchedules", token)
}

// VestCompletedSchedules is a paid mutator transaction binding the contract method 0xfdfaaa05.
//
// Solidity: function vestCompletedSchedules(address token) returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerSession) VestCompletedSchedules(token common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.VestCompletedSchedules(&_KyberRewardLocker.TransactOpts, token)
}

// VestCompletedSchedules is a paid mutator transaction binding the contract method 0xfdfaaa05.
//
// Solidity: function vestCompletedSchedules(address token) returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) VestCompletedSchedules(token common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.VestCompletedSchedules(&_KyberRewardLocker.TransactOpts, token)
}

// VestCompletedSchedulesForMultipleTokens is a paid mutator transaction binding the contract method 0x9059aa6a.
//
// Solidity: function vestCompletedSchedulesForMultipleTokens(address[] tokens) returns(uint256[] vestedAmounts)
func (_KyberRewardLocker *KyberRewardLockerTransactor) VestCompletedSchedulesForMultipleTokens(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "vestCompletedSchedulesForMultipleTokens", tokens)
}

// VestCompletedSchedulesForMultipleTokens is a paid mutator transaction binding the contract method 0x9059aa6a.
//
// Solidity: function vestCompletedSchedulesForMultipleTokens(address[] tokens) returns(uint256[] vestedAmounts)
func (_KyberRewardLocker *KyberRewardLockerSession) VestCompletedSchedulesForMultipleTokens(tokens []common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.VestCompletedSchedulesForMultipleTokens(&_KyberRewardLocker.TransactOpts, tokens)
}

// VestCompletedSchedulesForMultipleTokens is a paid mutator transaction binding the contract method 0x9059aa6a.
//
// Solidity: function vestCompletedSchedulesForMultipleTokens(address[] tokens) returns(uint256[] vestedAmounts)
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) VestCompletedSchedulesForMultipleTokens(tokens []common.Address) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.VestCompletedSchedulesForMultipleTokens(&_KyberRewardLocker.TransactOpts, tokens)
}

// VestScheduleAtIndices is a paid mutator transaction binding the contract method 0x0f5636c3.
//
// Solidity: function vestScheduleAtIndices(address token, uint256[] indexes) returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerTransactor) VestScheduleAtIndices(opts *bind.TransactOpts, token common.Address, indexes []*big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "vestScheduleAtIndices", token, indexes)
}

// VestScheduleAtIndices is a paid mutator transaction binding the contract method 0x0f5636c3.
//
// Solidity: function vestScheduleAtIndices(address token, uint256[] indexes) returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerSession) VestScheduleAtIndices(token common.Address, indexes []*big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.VestScheduleAtIndices(&_KyberRewardLocker.TransactOpts, token, indexes)
}

// VestScheduleAtIndices is a paid mutator transaction binding the contract method 0x0f5636c3.
//
// Solidity: function vestScheduleAtIndices(address token, uint256[] indexes) returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) VestScheduleAtIndices(token common.Address, indexes []*big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.VestScheduleAtIndices(&_KyberRewardLocker.TransactOpts, token, indexes)
}

// VestScheduleForMultipleTokensAtIndices is a paid mutator transaction binding the contract method 0x4d4f3d93.
//
// Solidity: function vestScheduleForMultipleTokensAtIndices(address[] tokens, uint256[][] indices) returns(uint256[] vestedAmounts)
func (_KyberRewardLocker *KyberRewardLockerTransactor) VestScheduleForMultipleTokensAtIndices(opts *bind.TransactOpts, tokens []common.Address, indices [][]*big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "vestScheduleForMultipleTokensAtIndices", tokens, indices)
}

// VestScheduleForMultipleTokensAtIndices is a paid mutator transaction binding the contract method 0x4d4f3d93.
//
// Solidity: function vestScheduleForMultipleTokensAtIndices(address[] tokens, uint256[][] indices) returns(uint256[] vestedAmounts)
func (_KyberRewardLocker *KyberRewardLockerSession) VestScheduleForMultipleTokensAtIndices(tokens []common.Address, indices [][]*big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.VestScheduleForMultipleTokensAtIndices(&_KyberRewardLocker.TransactOpts, tokens, indices)
}

// VestScheduleForMultipleTokensAtIndices is a paid mutator transaction binding the contract method 0x4d4f3d93.
//
// Solidity: function vestScheduleForMultipleTokensAtIndices(address[] tokens, uint256[][] indices) returns(uint256[] vestedAmounts)
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) VestScheduleForMultipleTokensAtIndices(tokens []common.Address, indices [][]*big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.VestScheduleForMultipleTokensAtIndices(&_KyberRewardLocker.TransactOpts, tokens, indices)
}

// VestSchedulesInRange is a paid mutator transaction binding the contract method 0xc33fddf8.
//
// Solidity: function vestSchedulesInRange(address token, uint256 startIndex, uint256 endIndex) returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerTransactor) VestSchedulesInRange(opts *bind.TransactOpts, token common.Address, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.contract.Transact(opts, "vestSchedulesInRange", token, startIndex, endIndex)
}

// VestSchedulesInRange is a paid mutator transaction binding the contract method 0xc33fddf8.
//
// Solidity: function vestSchedulesInRange(address token, uint256 startIndex, uint256 endIndex) returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerSession) VestSchedulesInRange(token common.Address, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.VestSchedulesInRange(&_KyberRewardLocker.TransactOpts, token, startIndex, endIndex)
}

// VestSchedulesInRange is a paid mutator transaction binding the contract method 0xc33fddf8.
//
// Solidity: function vestSchedulesInRange(address token, uint256 startIndex, uint256 endIndex) returns(uint256)
func (_KyberRewardLocker *KyberRewardLockerTransactorSession) VestSchedulesInRange(token common.Address, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _KyberRewardLocker.Contract.VestSchedulesInRange(&_KyberRewardLocker.TransactOpts, token, startIndex, endIndex)
}
