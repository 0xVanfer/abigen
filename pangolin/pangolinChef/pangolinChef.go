// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pangolinChef

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

// PangolinChefMetaData contains all meta data concerning the PangolinChef contract.
var PangolinChefMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"REWARD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_funder\",\"type\":\"address\"}],\"name\":\"addFunder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_allocPoints\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_lpTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_rewarders\",\"type\":\"address[]\"}],\"name\":\"addPools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"depositWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extension\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFunding\",\"type\":\"uint256\"}],\"name\":\"extendRewardsViaDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"funding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minExtension\",\"type\":\"uint256\"}],\"name\":\"extendRewardsViaFunding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"funding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"fundRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_funder\",\"type\":\"address\"}],\"name\":\"isFunder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdateAllPools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pids\",\"type\":\"uint256[]\"}],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrationDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accRewardPerShare\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"lastRewardTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"allocPoint\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pools\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_funder\",\"type\":\"address\"}],\"name\":\"removeFunder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"resetRewardsDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsExpiration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"allocPoints\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"rewarders\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"overwrites\",\"type\":\"bool[]\"}],\"name\":\"setPools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardDebt\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawAndHarvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PangolinChefABI is the input ABI used to generate the binding from.
// Deprecated: Use PangolinChefMetaData.ABI instead.
var PangolinChefABI = PangolinChefMetaData.ABI

// PangolinChef is an auto generated Go binding around an Ethereum contract.
type PangolinChef struct {
	PangolinChefCaller     // Read-only binding to the contract
	PangolinChefTransactor // Write-only binding to the contract
	PangolinChefFilterer   // Log filterer for contract events
}

// PangolinChefCaller is an auto generated read-only Go binding around an Ethereum contract.
type PangolinChefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PangolinChefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PangolinChefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PangolinChefFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PangolinChefFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PangolinChefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PangolinChefSession struct {
	Contract     *PangolinChef     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PangolinChefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PangolinChefCallerSession struct {
	Contract *PangolinChefCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PangolinChefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PangolinChefTransactorSession struct {
	Contract     *PangolinChefTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PangolinChefRaw is an auto generated low-level Go binding around an Ethereum contract.
type PangolinChefRaw struct {
	Contract *PangolinChef // Generic contract binding to access the raw methods on
}

// PangolinChefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PangolinChefCallerRaw struct {
	Contract *PangolinChefCaller // Generic read-only contract binding to access the raw methods on
}

// PangolinChefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PangolinChefTransactorRaw struct {
	Contract *PangolinChefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPangolinChef creates a new instance of PangolinChef, bound to a specific deployed contract.
func NewPangolinChef(address common.Address, backend bind.ContractBackend) (*PangolinChef, error) {
	contract, err := bindPangolinChef(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PangolinChef{PangolinChefCaller: PangolinChefCaller{contract: contract}, PangolinChefTransactor: PangolinChefTransactor{contract: contract}, PangolinChefFilterer: PangolinChefFilterer{contract: contract}}, nil
}

// NewPangolinChefCaller creates a new read-only instance of PangolinChef, bound to a specific deployed contract.
func NewPangolinChefCaller(address common.Address, caller bind.ContractCaller) (*PangolinChefCaller, error) {
	contract, err := bindPangolinChef(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PangolinChefCaller{contract: contract}, nil
}

// NewPangolinChefTransactor creates a new write-only instance of PangolinChef, bound to a specific deployed contract.
func NewPangolinChefTransactor(address common.Address, transactor bind.ContractTransactor) (*PangolinChefTransactor, error) {
	contract, err := bindPangolinChef(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PangolinChefTransactor{contract: contract}, nil
}

// NewPangolinChefFilterer creates a new log filterer instance of PangolinChef, bound to a specific deployed contract.
func NewPangolinChefFilterer(address common.Address, filterer bind.ContractFilterer) (*PangolinChefFilterer, error) {
	contract, err := bindPangolinChef(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PangolinChefFilterer{contract: contract}, nil
}

// bindPangolinChef binds a generic wrapper to an already deployed contract.
func bindPangolinChef(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PangolinChefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PangolinChef *PangolinChefRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PangolinChef.Contract.PangolinChefCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PangolinChef *PangolinChefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PangolinChef.Contract.PangolinChefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PangolinChef *PangolinChefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PangolinChef.Contract.PangolinChefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PangolinChef *PangolinChefCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PangolinChef.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PangolinChef *PangolinChefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PangolinChef.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PangolinChef *PangolinChefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PangolinChef.Contract.contract.Transact(opts, method, params...)
}

// REWARD is a free data retrieval call binding the contract method 0xcab34c08.
//
// Solidity: function REWARD() view returns(address)
func (_PangolinChef *PangolinChefCaller) REWARD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "REWARD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARD is a free data retrieval call binding the contract method 0xcab34c08.
//
// Solidity: function REWARD() view returns(address)
func (_PangolinChef *PangolinChefSession) REWARD() (common.Address, error) {
	return _PangolinChef.Contract.REWARD(&_PangolinChef.CallOpts)
}

// REWARD is a free data retrieval call binding the contract method 0xcab34c08.
//
// Solidity: function REWARD() view returns(address)
func (_PangolinChef *PangolinChefCallerSession) REWARD() (common.Address, error) {
	return _PangolinChef.Contract.REWARD(&_PangolinChef.CallOpts)
}

// AddedTokens is a free data retrieval call binding the contract method 0x79d12ffb.
//
// Solidity: function addedTokens(address ) view returns(bool)
func (_PangolinChef *PangolinChefCaller) AddedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "addedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddedTokens is a free data retrieval call binding the contract method 0x79d12ffb.
//
// Solidity: function addedTokens(address ) view returns(bool)
func (_PangolinChef *PangolinChefSession) AddedTokens(arg0 common.Address) (bool, error) {
	return _PangolinChef.Contract.AddedTokens(&_PangolinChef.CallOpts, arg0)
}

// AddedTokens is a free data retrieval call binding the contract method 0x79d12ffb.
//
// Solidity: function addedTokens(address ) view returns(bool)
func (_PangolinChef *PangolinChefCallerSession) AddedTokens(arg0 common.Address) (bool, error) {
	return _PangolinChef.Contract.AddedTokens(&_PangolinChef.CallOpts, arg0)
}

// IsFunder is a free data retrieval call binding the contract method 0x1ea48870.
//
// Solidity: function isFunder(address _funder) view returns(bool allowed)
func (_PangolinChef *PangolinChefCaller) IsFunder(opts *bind.CallOpts, _funder common.Address) (bool, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "isFunder", _funder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFunder is a free data retrieval call binding the contract method 0x1ea48870.
//
// Solidity: function isFunder(address _funder) view returns(bool allowed)
func (_PangolinChef *PangolinChefSession) IsFunder(_funder common.Address) (bool, error) {
	return _PangolinChef.Contract.IsFunder(&_PangolinChef.CallOpts, _funder)
}

// IsFunder is a free data retrieval call binding the contract method 0x1ea48870.
//
// Solidity: function isFunder(address _funder) view returns(bool allowed)
func (_PangolinChef *PangolinChefCallerSession) IsFunder(_funder common.Address) (bool, error) {
	return _PangolinChef.Contract.IsFunder(&_PangolinChef.CallOpts, _funder)
}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_PangolinChef *PangolinChefCaller) LpToken(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "lpToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_PangolinChef *PangolinChefSession) LpToken(arg0 *big.Int) (common.Address, error) {
	return _PangolinChef.Contract.LpToken(&_PangolinChef.CallOpts, arg0)
}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_PangolinChef *PangolinChefCallerSession) LpToken(arg0 *big.Int) (common.Address, error) {
	return _PangolinChef.Contract.LpToken(&_PangolinChef.CallOpts, arg0)
}

// LpTokens is a free data retrieval call binding the contract method 0x34c0b46b.
//
// Solidity: function lpTokens() view returns(address[])
func (_PangolinChef *PangolinChefCaller) LpTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "lpTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// LpTokens is a free data retrieval call binding the contract method 0x34c0b46b.
//
// Solidity: function lpTokens() view returns(address[])
func (_PangolinChef *PangolinChefSession) LpTokens() ([]common.Address, error) {
	return _PangolinChef.Contract.LpTokens(&_PangolinChef.CallOpts)
}

// LpTokens is a free data retrieval call binding the contract method 0x34c0b46b.
//
// Solidity: function lpTokens() view returns(address[])
func (_PangolinChef *PangolinChefCallerSession) LpTokens() ([]common.Address, error) {
	return _PangolinChef.Contract.LpTokens(&_PangolinChef.CallOpts)
}

// MigrationDisabled is a free data retrieval call binding the contract method 0x0c812af3.
//
// Solidity: function migrationDisabled() view returns(bool)
func (_PangolinChef *PangolinChefCaller) MigrationDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "migrationDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MigrationDisabled is a free data retrieval call binding the contract method 0x0c812af3.
//
// Solidity: function migrationDisabled() view returns(bool)
func (_PangolinChef *PangolinChefSession) MigrationDisabled() (bool, error) {
	return _PangolinChef.Contract.MigrationDisabled(&_PangolinChef.CallOpts)
}

// MigrationDisabled is a free data retrieval call binding the contract method 0x0c812af3.
//
// Solidity: function migrationDisabled() view returns(bool)
func (_PangolinChef *PangolinChefCallerSession) MigrationDisabled() (bool, error) {
	return _PangolinChef.Contract.MigrationDisabled(&_PangolinChef.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_PangolinChef *PangolinChefCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_PangolinChef *PangolinChefSession) Migrator() (common.Address, error) {
	return _PangolinChef.Contract.Migrator(&_PangolinChef.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_PangolinChef *PangolinChefCallerSession) Migrator() (common.Address, error) {
	return _PangolinChef.Contract.Migrator(&_PangolinChef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PangolinChef *PangolinChefCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PangolinChef *PangolinChefSession) Owner() (common.Address, error) {
	return _PangolinChef.Contract.Owner(&_PangolinChef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PangolinChef *PangolinChefCallerSession) Owner() (common.Address, error) {
	return _PangolinChef.Contract.Owner(&_PangolinChef.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256 pending)
func (_PangolinChef *PangolinChefCaller) PendingReward(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "pendingReward", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256 pending)
func (_PangolinChef *PangolinChefSession) PendingReward(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _PangolinChef.Contract.PendingReward(&_PangolinChef.CallOpts, _pid, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256 pending)
func (_PangolinChef *PangolinChefCallerSession) PendingReward(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _PangolinChef.Contract.PendingReward(&_PangolinChef.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint128 accRewardPerShare, uint64 lastRewardTime, uint64 allocPoint)
func (_PangolinChef *PangolinChefCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AccRewardPerShare *big.Int
	LastRewardTime    uint64
	AllocPoint        uint64
}, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		AccRewardPerShare *big.Int
		LastRewardTime    uint64
		AllocPoint        uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccRewardPerShare = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTime = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.AllocPoint = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint128 accRewardPerShare, uint64 lastRewardTime, uint64 allocPoint)
func (_PangolinChef *PangolinChefSession) PoolInfo(arg0 *big.Int) (struct {
	AccRewardPerShare *big.Int
	LastRewardTime    uint64
	AllocPoint        uint64
}, error) {
	return _PangolinChef.Contract.PoolInfo(&_PangolinChef.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint128 accRewardPerShare, uint64 lastRewardTime, uint64 allocPoint)
func (_PangolinChef *PangolinChefCallerSession) PoolInfo(arg0 *big.Int) (struct {
	AccRewardPerShare *big.Int
	LastRewardTime    uint64
	AllocPoint        uint64
}, error) {
	return _PangolinChef.Contract.PoolInfo(&_PangolinChef.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_PangolinChef *PangolinChefCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_PangolinChef *PangolinChefSession) PoolLength() (*big.Int, error) {
	return _PangolinChef.Contract.PoolLength(&_PangolinChef.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_PangolinChef *PangolinChefCallerSession) PoolLength() (*big.Int, error) {
	return _PangolinChef.Contract.PoolLength(&_PangolinChef.CallOpts)
}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_PangolinChef *PangolinChefCaller) RewardPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "rewardPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_PangolinChef *PangolinChefSession) RewardPerSecond() (*big.Int, error) {
	return _PangolinChef.Contract.RewardPerSecond(&_PangolinChef.CallOpts)
}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_PangolinChef *PangolinChefCallerSession) RewardPerSecond() (*big.Int, error) {
	return _PangolinChef.Contract.RewardPerSecond(&_PangolinChef.CallOpts)
}

// Rewarder is a free data retrieval call binding the contract method 0xc346253d.
//
// Solidity: function rewarder(uint256 ) view returns(address)
func (_PangolinChef *PangolinChefCaller) Rewarder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "rewarder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewarder is a free data retrieval call binding the contract method 0xc346253d.
//
// Solidity: function rewarder(uint256 ) view returns(address)
func (_PangolinChef *PangolinChefSession) Rewarder(arg0 *big.Int) (common.Address, error) {
	return _PangolinChef.Contract.Rewarder(&_PangolinChef.CallOpts, arg0)
}

// Rewarder is a free data retrieval call binding the contract method 0xc346253d.
//
// Solidity: function rewarder(uint256 ) view returns(address)
func (_PangolinChef *PangolinChefCallerSession) Rewarder(arg0 *big.Int) (common.Address, error) {
	return _PangolinChef.Contract.Rewarder(&_PangolinChef.CallOpts, arg0)
}

// RewardsExpiration is a free data retrieval call binding the contract method 0xb76f4aeb.
//
// Solidity: function rewardsExpiration() view returns(uint256)
func (_PangolinChef *PangolinChefCaller) RewardsExpiration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "rewardsExpiration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsExpiration is a free data retrieval call binding the contract method 0xb76f4aeb.
//
// Solidity: function rewardsExpiration() view returns(uint256)
func (_PangolinChef *PangolinChefSession) RewardsExpiration() (*big.Int, error) {
	return _PangolinChef.Contract.RewardsExpiration(&_PangolinChef.CallOpts)
}

// RewardsExpiration is a free data retrieval call binding the contract method 0xb76f4aeb.
//
// Solidity: function rewardsExpiration() view returns(uint256)
func (_PangolinChef *PangolinChefCallerSession) RewardsExpiration() (*big.Int, error) {
	return _PangolinChef.Contract.RewardsExpiration(&_PangolinChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_PangolinChef *PangolinChefCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_PangolinChef *PangolinChefSession) TotalAllocPoint() (*big.Int, error) {
	return _PangolinChef.Contract.TotalAllocPoint(&_PangolinChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_PangolinChef *PangolinChefCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _PangolinChef.Contract.TotalAllocPoint(&_PangolinChef.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, int256 rewardDebt)
func (_PangolinChef *PangolinChefCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _PangolinChef.contract.Call(opts, &out, "userInfo", arg0, arg1)

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

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, int256 rewardDebt)
func (_PangolinChef *PangolinChefSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _PangolinChef.Contract.UserInfo(&_PangolinChef.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, int256 rewardDebt)
func (_PangolinChef *PangolinChefCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _PangolinChef.Contract.UserInfo(&_PangolinChef.CallOpts, arg0, arg1)
}

// AddFunder is a paid mutator transaction binding the contract method 0x1bbe9d8c.
//
// Solidity: function addFunder(address _funder) returns()
func (_PangolinChef *PangolinChefTransactor) AddFunder(opts *bind.TransactOpts, _funder common.Address) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "addFunder", _funder)
}

// AddFunder is a paid mutator transaction binding the contract method 0x1bbe9d8c.
//
// Solidity: function addFunder(address _funder) returns()
func (_PangolinChef *PangolinChefSession) AddFunder(_funder common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.AddFunder(&_PangolinChef.TransactOpts, _funder)
}

// AddFunder is a paid mutator transaction binding the contract method 0x1bbe9d8c.
//
// Solidity: function addFunder(address _funder) returns()
func (_PangolinChef *PangolinChefTransactorSession) AddFunder(_funder common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.AddFunder(&_PangolinChef.TransactOpts, _funder)
}

// AddPool is a paid mutator transaction binding the contract method 0xf624d2c5.
//
// Solidity: function addPool(uint256 _allocPoint, address _lpToken, address _rewarder) returns()
func (_PangolinChef *PangolinChefTransactor) AddPool(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "addPool", _allocPoint, _lpToken, _rewarder)
}

// AddPool is a paid mutator transaction binding the contract method 0xf624d2c5.
//
// Solidity: function addPool(uint256 _allocPoint, address _lpToken, address _rewarder) returns()
func (_PangolinChef *PangolinChefSession) AddPool(_allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.AddPool(&_PangolinChef.TransactOpts, _allocPoint, _lpToken, _rewarder)
}

// AddPool is a paid mutator transaction binding the contract method 0xf624d2c5.
//
// Solidity: function addPool(uint256 _allocPoint, address _lpToken, address _rewarder) returns()
func (_PangolinChef *PangolinChefTransactorSession) AddPool(_allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.AddPool(&_PangolinChef.TransactOpts, _allocPoint, _lpToken, _rewarder)
}

// AddPools is a paid mutator transaction binding the contract method 0x7973aa7b.
//
// Solidity: function addPools(uint256[] _allocPoints, address[] _lpTokens, address[] _rewarders) returns()
func (_PangolinChef *PangolinChefTransactor) AddPools(opts *bind.TransactOpts, _allocPoints []*big.Int, _lpTokens []common.Address, _rewarders []common.Address) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "addPools", _allocPoints, _lpTokens, _rewarders)
}

// AddPools is a paid mutator transaction binding the contract method 0x7973aa7b.
//
// Solidity: function addPools(uint256[] _allocPoints, address[] _lpTokens, address[] _rewarders) returns()
func (_PangolinChef *PangolinChefSession) AddPools(_allocPoints []*big.Int, _lpTokens []common.Address, _rewarders []common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.AddPools(&_PangolinChef.TransactOpts, _allocPoints, _lpTokens, _rewarders)
}

// AddPools is a paid mutator transaction binding the contract method 0x7973aa7b.
//
// Solidity: function addPools(uint256[] _allocPoints, address[] _lpTokens, address[] _rewarders) returns()
func (_PangolinChef *PangolinChefTransactorSession) AddPools(_allocPoints []*big.Int, _lpTokens []common.Address, _rewarders []common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.AddPools(&_PangolinChef.TransactOpts, _allocPoints, _lpTokens, _rewarders)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 pid, uint256 amount, address to) returns()
func (_PangolinChef *PangolinChefTransactor) Deposit(opts *bind.TransactOpts, pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "deposit", pid, amount, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 pid, uint256 amount, address to) returns()
func (_PangolinChef *PangolinChefSession) Deposit(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.Deposit(&_PangolinChef.TransactOpts, pid, amount, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 pid, uint256 amount, address to) returns()
func (_PangolinChef *PangolinChefTransactorSession) Deposit(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.Deposit(&_PangolinChef.TransactOpts, pid, amount, to)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x2273ee6e.
//
// Solidity: function depositWithPermit(uint256 pid, uint256 amount, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PangolinChef *PangolinChefTransactor) DepositWithPermit(opts *bind.TransactOpts, pid *big.Int, amount *big.Int, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "depositWithPermit", pid, amount, to, deadline, v, r, s)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x2273ee6e.
//
// Solidity: function depositWithPermit(uint256 pid, uint256 amount, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PangolinChef *PangolinChefSession) DepositWithPermit(pid *big.Int, amount *big.Int, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PangolinChef.Contract.DepositWithPermit(&_PangolinChef.TransactOpts, pid, amount, to, deadline, v, r, s)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x2273ee6e.
//
// Solidity: function depositWithPermit(uint256 pid, uint256 amount, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PangolinChef *PangolinChefTransactorSession) DepositWithPermit(pid *big.Int, amount *big.Int, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PangolinChef.Contract.DepositWithPermit(&_PangolinChef.TransactOpts, pid, amount, to, deadline, v, r, s)
}

// DisableMigrator is a paid mutator transaction binding the contract method 0xf6d676be.
//
// Solidity: function disableMigrator() returns()
func (_PangolinChef *PangolinChefTransactor) DisableMigrator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "disableMigrator")
}

// DisableMigrator is a paid mutator transaction binding the contract method 0xf6d676be.
//
// Solidity: function disableMigrator() returns()
func (_PangolinChef *PangolinChefSession) DisableMigrator() (*types.Transaction, error) {
	return _PangolinChef.Contract.DisableMigrator(&_PangolinChef.TransactOpts)
}

// DisableMigrator is a paid mutator transaction binding the contract method 0xf6d676be.
//
// Solidity: function disableMigrator() returns()
func (_PangolinChef *PangolinChefTransactorSession) DisableMigrator() (*types.Transaction, error) {
	return _PangolinChef.Contract.DisableMigrator(&_PangolinChef.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x2f940c70.
//
// Solidity: function emergencyWithdraw(uint256 pid, address to) returns()
func (_PangolinChef *PangolinChefTransactor) EmergencyWithdraw(opts *bind.TransactOpts, pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "emergencyWithdraw", pid, to)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x2f940c70.
//
// Solidity: function emergencyWithdraw(uint256 pid, address to) returns()
func (_PangolinChef *PangolinChefSession) EmergencyWithdraw(pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.EmergencyWithdraw(&_PangolinChef.TransactOpts, pid, to)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x2f940c70.
//
// Solidity: function emergencyWithdraw(uint256 pid, address to) returns()
func (_PangolinChef *PangolinChefTransactorSession) EmergencyWithdraw(pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.EmergencyWithdraw(&_PangolinChef.TransactOpts, pid, to)
}

// ExtendRewardsViaDuration is a paid mutator transaction binding the contract method 0x933f084f.
//
// Solidity: function extendRewardsViaDuration(uint256 extension, uint256 maxFunding) returns()
func (_PangolinChef *PangolinChefTransactor) ExtendRewardsViaDuration(opts *bind.TransactOpts, extension *big.Int, maxFunding *big.Int) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "extendRewardsViaDuration", extension, maxFunding)
}

// ExtendRewardsViaDuration is a paid mutator transaction binding the contract method 0x933f084f.
//
// Solidity: function extendRewardsViaDuration(uint256 extension, uint256 maxFunding) returns()
func (_PangolinChef *PangolinChefSession) ExtendRewardsViaDuration(extension *big.Int, maxFunding *big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.ExtendRewardsViaDuration(&_PangolinChef.TransactOpts, extension, maxFunding)
}

// ExtendRewardsViaDuration is a paid mutator transaction binding the contract method 0x933f084f.
//
// Solidity: function extendRewardsViaDuration(uint256 extension, uint256 maxFunding) returns()
func (_PangolinChef *PangolinChefTransactorSession) ExtendRewardsViaDuration(extension *big.Int, maxFunding *big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.ExtendRewardsViaDuration(&_PangolinChef.TransactOpts, extension, maxFunding)
}

// ExtendRewardsViaFunding is a paid mutator transaction binding the contract method 0x19610b2b.
//
// Solidity: function extendRewardsViaFunding(uint256 funding, uint256 minExtension) returns()
func (_PangolinChef *PangolinChefTransactor) ExtendRewardsViaFunding(opts *bind.TransactOpts, funding *big.Int, minExtension *big.Int) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "extendRewardsViaFunding", funding, minExtension)
}

// ExtendRewardsViaFunding is a paid mutator transaction binding the contract method 0x19610b2b.
//
// Solidity: function extendRewardsViaFunding(uint256 funding, uint256 minExtension) returns()
func (_PangolinChef *PangolinChefSession) ExtendRewardsViaFunding(funding *big.Int, minExtension *big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.ExtendRewardsViaFunding(&_PangolinChef.TransactOpts, funding, minExtension)
}

// ExtendRewardsViaFunding is a paid mutator transaction binding the contract method 0x19610b2b.
//
// Solidity: function extendRewardsViaFunding(uint256 funding, uint256 minExtension) returns()
func (_PangolinChef *PangolinChefTransactorSession) ExtendRewardsViaFunding(funding *big.Int, minExtension *big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.ExtendRewardsViaFunding(&_PangolinChef.TransactOpts, funding, minExtension)
}

// FundRewards is a paid mutator transaction binding the contract method 0x28662551.
//
// Solidity: function fundRewards(uint256 funding, uint256 duration) returns()
func (_PangolinChef *PangolinChefTransactor) FundRewards(opts *bind.TransactOpts, funding *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "fundRewards", funding, duration)
}

// FundRewards is a paid mutator transaction binding the contract method 0x28662551.
//
// Solidity: function fundRewards(uint256 funding, uint256 duration) returns()
func (_PangolinChef *PangolinChefSession) FundRewards(funding *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.FundRewards(&_PangolinChef.TransactOpts, funding, duration)
}

// FundRewards is a paid mutator transaction binding the contract method 0x28662551.
//
// Solidity: function fundRewards(uint256 funding, uint256 duration) returns()
func (_PangolinChef *PangolinChefTransactorSession) FundRewards(funding *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.FundRewards(&_PangolinChef.TransactOpts, funding, duration)
}

// Harvest is a paid mutator transaction binding the contract method 0x18fccc76.
//
// Solidity: function harvest(uint256 pid, address to) returns()
func (_PangolinChef *PangolinChefTransactor) Harvest(opts *bind.TransactOpts, pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "harvest", pid, to)
}

// Harvest is a paid mutator transaction binding the contract method 0x18fccc76.
//
// Solidity: function harvest(uint256 pid, address to) returns()
func (_PangolinChef *PangolinChefSession) Harvest(pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.Harvest(&_PangolinChef.TransactOpts, pid, to)
}

// Harvest is a paid mutator transaction binding the contract method 0x18fccc76.
//
// Solidity: function harvest(uint256 pid, address to) returns()
func (_PangolinChef *PangolinChefTransactorSession) Harvest(pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.Harvest(&_PangolinChef.TransactOpts, pid, to)
}

// MassUpdateAllPools is a paid mutator transaction binding the contract method 0xc1bcb493.
//
// Solidity: function massUpdateAllPools() returns()
func (_PangolinChef *PangolinChefTransactor) MassUpdateAllPools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "massUpdateAllPools")
}

// MassUpdateAllPools is a paid mutator transaction binding the contract method 0xc1bcb493.
//
// Solidity: function massUpdateAllPools() returns()
func (_PangolinChef *PangolinChefSession) MassUpdateAllPools() (*types.Transaction, error) {
	return _PangolinChef.Contract.MassUpdateAllPools(&_PangolinChef.TransactOpts)
}

// MassUpdateAllPools is a paid mutator transaction binding the contract method 0xc1bcb493.
//
// Solidity: function massUpdateAllPools() returns()
func (_PangolinChef *PangolinChefTransactorSession) MassUpdateAllPools() (*types.Transaction, error) {
	return _PangolinChef.Contract.MassUpdateAllPools(&_PangolinChef.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_PangolinChef *PangolinChefTransactor) MassUpdatePools(opts *bind.TransactOpts, pids []*big.Int) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "massUpdatePools", pids)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_PangolinChef *PangolinChefSession) MassUpdatePools(pids []*big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.MassUpdatePools(&_PangolinChef.TransactOpts, pids)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_PangolinChef *PangolinChefTransactorSession) MassUpdatePools(pids []*big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.MassUpdatePools(&_PangolinChef.TransactOpts, pids)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_PangolinChef *PangolinChefTransactor) Migrate(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "migrate", _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_PangolinChef *PangolinChefSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.Migrate(&_PangolinChef.TransactOpts, _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_PangolinChef *PangolinChefTransactorSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.Migrate(&_PangolinChef.TransactOpts, _pid)
}

// RemoveFunder is a paid mutator transaction binding the contract method 0xdcd18dd4.
//
// Solidity: function removeFunder(address _funder) returns()
func (_PangolinChef *PangolinChefTransactor) RemoveFunder(opts *bind.TransactOpts, _funder common.Address) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "removeFunder", _funder)
}

// RemoveFunder is a paid mutator transaction binding the contract method 0xdcd18dd4.
//
// Solidity: function removeFunder(address _funder) returns()
func (_PangolinChef *PangolinChefSession) RemoveFunder(_funder common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.RemoveFunder(&_PangolinChef.TransactOpts, _funder)
}

// RemoveFunder is a paid mutator transaction binding the contract method 0xdcd18dd4.
//
// Solidity: function removeFunder(address _funder) returns()
func (_PangolinChef *PangolinChefTransactorSession) RemoveFunder(_funder common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.RemoveFunder(&_PangolinChef.TransactOpts, _funder)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PangolinChef *PangolinChefTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PangolinChef *PangolinChefSession) RenounceOwnership() (*types.Transaction, error) {
	return _PangolinChef.Contract.RenounceOwnership(&_PangolinChef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PangolinChef *PangolinChefTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PangolinChef.Contract.RenounceOwnership(&_PangolinChef.TransactOpts)
}

// ResetRewardsDuration is a paid mutator transaction binding the contract method 0x4b87e90f.
//
// Solidity: function resetRewardsDuration(uint256 duration) returns()
func (_PangolinChef *PangolinChefTransactor) ResetRewardsDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "resetRewardsDuration", duration)
}

// ResetRewardsDuration is a paid mutator transaction binding the contract method 0x4b87e90f.
//
// Solidity: function resetRewardsDuration(uint256 duration) returns()
func (_PangolinChef *PangolinChefSession) ResetRewardsDuration(duration *big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.ResetRewardsDuration(&_PangolinChef.TransactOpts, duration)
}

// ResetRewardsDuration is a paid mutator transaction binding the contract method 0x4b87e90f.
//
// Solidity: function resetRewardsDuration(uint256 duration) returns()
func (_PangolinChef *PangolinChefTransactorSession) ResetRewardsDuration(duration *big.Int) (*types.Transaction, error) {
	return _PangolinChef.Contract.ResetRewardsDuration(&_PangolinChef.TransactOpts, duration)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_PangolinChef *PangolinChefTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_PangolinChef *PangolinChefSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.SetMigrator(&_PangolinChef.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_PangolinChef *PangolinChefTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.SetMigrator(&_PangolinChef.TransactOpts, _migrator)
}

// SetPool is a paid mutator transaction binding the contract method 0xb763e7c4.
//
// Solidity: function setPool(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_PangolinChef *PangolinChefTransactor) SetPool(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "setPool", _pid, _allocPoint, _rewarder, overwrite)
}

// SetPool is a paid mutator transaction binding the contract method 0xb763e7c4.
//
// Solidity: function setPool(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_PangolinChef *PangolinChefSession) SetPool(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _PangolinChef.Contract.SetPool(&_PangolinChef.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// SetPool is a paid mutator transaction binding the contract method 0xb763e7c4.
//
// Solidity: function setPool(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_PangolinChef *PangolinChefTransactorSession) SetPool(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _PangolinChef.Contract.SetPool(&_PangolinChef.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// SetPools is a paid mutator transaction binding the contract method 0x3ef6db91.
//
// Solidity: function setPools(uint256[] pids, uint256[] allocPoints, address[] rewarders, bool[] overwrites) returns()
func (_PangolinChef *PangolinChefTransactor) SetPools(opts *bind.TransactOpts, pids []*big.Int, allocPoints []*big.Int, rewarders []common.Address, overwrites []bool) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "setPools", pids, allocPoints, rewarders, overwrites)
}

// SetPools is a paid mutator transaction binding the contract method 0x3ef6db91.
//
// Solidity: function setPools(uint256[] pids, uint256[] allocPoints, address[] rewarders, bool[] overwrites) returns()
func (_PangolinChef *PangolinChefSession) SetPools(pids []*big.Int, allocPoints []*big.Int, rewarders []common.Address, overwrites []bool) (*types.Transaction, error) {
	return _PangolinChef.Contract.SetPools(&_PangolinChef.TransactOpts, pids, allocPoints, rewarders, overwrites)
}

// SetPools is a paid mutator transaction binding the contract method 0x3ef6db91.
//
// Solidity: function setPools(uint256[] pids, uint256[] allocPoints, address[] rewarders, bool[] overwrites) returns()
func (_PangolinChef *PangolinChefTransactorSession) SetPools(pids []*big.Int, allocPoints []*big.Int, rewarders []common.Address, overwrites []bool) (*types.Transaction, error) {
	return _PangolinChef.Contract.SetPools(&_PangolinChef.TransactOpts, pids, allocPoints, rewarders, overwrites)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PangolinChef *PangolinChefTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PangolinChef *PangolinChefSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.TransferOwnership(&_PangolinChef.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PangolinChef *PangolinChefTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.TransferOwnership(&_PangolinChef.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0ad58d2f.
//
// Solidity: function withdraw(uint256 pid, uint256 amount, address to) returns()
func (_PangolinChef *PangolinChefTransactor) Withdraw(opts *bind.TransactOpts, pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "withdraw", pid, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0ad58d2f.
//
// Solidity: function withdraw(uint256 pid, uint256 amount, address to) returns()
func (_PangolinChef *PangolinChefSession) Withdraw(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.Withdraw(&_PangolinChef.TransactOpts, pid, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0ad58d2f.
//
// Solidity: function withdraw(uint256 pid, uint256 amount, address to) returns()
func (_PangolinChef *PangolinChefTransactorSession) Withdraw(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.Withdraw(&_PangolinChef.TransactOpts, pid, amount, to)
}

// WithdrawAndHarvest is a paid mutator transaction binding the contract method 0xd1abb907.
//
// Solidity: function withdrawAndHarvest(uint256 pid, uint256 amount, address to) returns()
func (_PangolinChef *PangolinChefTransactor) WithdrawAndHarvest(opts *bind.TransactOpts, pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.contract.Transact(opts, "withdrawAndHarvest", pid, amount, to)
}

// WithdrawAndHarvest is a paid mutator transaction binding the contract method 0xd1abb907.
//
// Solidity: function withdrawAndHarvest(uint256 pid, uint256 amount, address to) returns()
func (_PangolinChef *PangolinChefSession) WithdrawAndHarvest(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.WithdrawAndHarvest(&_PangolinChef.TransactOpts, pid, amount, to)
}

// WithdrawAndHarvest is a paid mutator transaction binding the contract method 0xd1abb907.
//
// Solidity: function withdrawAndHarvest(uint256 pid, uint256 amount, address to) returns()
func (_PangolinChef *PangolinChefTransactorSession) WithdrawAndHarvest(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _PangolinChef.Contract.WithdrawAndHarvest(&_PangolinChef.TransactOpts, pid, amount, to)
}
