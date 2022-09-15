// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platypusChefV2

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

// PlatypusChefV2MetaData contains all meta data concerning the PlatypusChefV2 contract.
var PlatypusChefV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseAllocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimablePtp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dialutingRepartition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPtpWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ptp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vePtp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ptpPerSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dialutingRepartition\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPoolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"name\":\"multiClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newMasterPlatypus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonDialutingRepartition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCandidate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingPtp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pendingBonusToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"poolAdjustFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAllocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accPtpPerShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewarder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sumOfFactors\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accPtpPerFactorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adjustedAllocPoint\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"proposeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ptp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ptpPerSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"rewarderBonusTokenInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseAllocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPoolLength\",\"type\":\"uint256\"}],\"name\":\"setMaxPoolLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMasterPlatypus\",\"type\":\"address\"}],\"name\":\"setNewMasterPlatypus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVePtp\",\"type\":\"address\"}],\"name\":\"setVePtp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAdjustedAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ptpPerSec\",\"type\":\"uint256\"}],\"name\":\"updateEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dialutingRepartition\",\"type\":\"uint256\"}],\"name\":\"updateEmissionRepartition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newVePtpBalance\",\"type\":\"uint256\"}],\"name\":\"updateFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"factor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vePtp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PlatypusChefV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use PlatypusChefV2MetaData.ABI instead.
var PlatypusChefV2ABI = PlatypusChefV2MetaData.ABI

// PlatypusChefV2 is an auto generated Go binding around an Ethereum contract.
type PlatypusChefV2 struct {
	PlatypusChefV2Caller     // Read-only binding to the contract
	PlatypusChefV2Transactor // Write-only binding to the contract
	PlatypusChefV2Filterer   // Log filterer for contract events
}

// PlatypusChefV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type PlatypusChefV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusChefV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PlatypusChefV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusChefV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlatypusChefV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusChefV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlatypusChefV2Session struct {
	Contract     *PlatypusChefV2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlatypusChefV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlatypusChefV2CallerSession struct {
	Contract *PlatypusChefV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PlatypusChefV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlatypusChefV2TransactorSession struct {
	Contract     *PlatypusChefV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PlatypusChefV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type PlatypusChefV2Raw struct {
	Contract *PlatypusChefV2 // Generic contract binding to access the raw methods on
}

// PlatypusChefV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlatypusChefV2CallerRaw struct {
	Contract *PlatypusChefV2Caller // Generic read-only contract binding to access the raw methods on
}

// PlatypusChefV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlatypusChefV2TransactorRaw struct {
	Contract *PlatypusChefV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPlatypusChefV2 creates a new instance of PlatypusChefV2, bound to a specific deployed contract.
func NewPlatypusChefV2(address common.Address, backend bind.ContractBackend) (*PlatypusChefV2, error) {
	contract, err := bindPlatypusChefV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlatypusChefV2{PlatypusChefV2Caller: PlatypusChefV2Caller{contract: contract}, PlatypusChefV2Transactor: PlatypusChefV2Transactor{contract: contract}, PlatypusChefV2Filterer: PlatypusChefV2Filterer{contract: contract}}, nil
}

// NewPlatypusChefV2Caller creates a new read-only instance of PlatypusChefV2, bound to a specific deployed contract.
func NewPlatypusChefV2Caller(address common.Address, caller bind.ContractCaller) (*PlatypusChefV2Caller, error) {
	contract, err := bindPlatypusChefV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlatypusChefV2Caller{contract: contract}, nil
}

// NewPlatypusChefV2Transactor creates a new write-only instance of PlatypusChefV2, bound to a specific deployed contract.
func NewPlatypusChefV2Transactor(address common.Address, transactor bind.ContractTransactor) (*PlatypusChefV2Transactor, error) {
	contract, err := bindPlatypusChefV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlatypusChefV2Transactor{contract: contract}, nil
}

// NewPlatypusChefV2Filterer creates a new log filterer instance of PlatypusChefV2, bound to a specific deployed contract.
func NewPlatypusChefV2Filterer(address common.Address, filterer bind.ContractFilterer) (*PlatypusChefV2Filterer, error) {
	contract, err := bindPlatypusChefV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlatypusChefV2Filterer{contract: contract}, nil
}

// bindPlatypusChefV2 binds a generic wrapper to an already deployed contract.
func bindPlatypusChefV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlatypusChefV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatypusChefV2 *PlatypusChefV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatypusChefV2.Contract.PlatypusChefV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatypusChefV2 *PlatypusChefV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.PlatypusChefV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatypusChefV2 *PlatypusChefV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.PlatypusChefV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatypusChefV2 *PlatypusChefV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatypusChefV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatypusChefV2 *PlatypusChefV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatypusChefV2 *PlatypusChefV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.contract.Transact(opts, method, params...)
}

// ClaimablePtp is a free data retrieval call binding the contract method 0x8b4d83a3.
//
// Solidity: function claimablePtp(uint256 , address ) view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Caller) ClaimablePtp(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "claimablePtp", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimablePtp is a free data retrieval call binding the contract method 0x8b4d83a3.
//
// Solidity: function claimablePtp(uint256 , address ) view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) ClaimablePtp(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _PlatypusChefV2.Contract.ClaimablePtp(&_PlatypusChefV2.CallOpts, arg0, arg1)
}

// ClaimablePtp is a free data retrieval call binding the contract method 0x8b4d83a3.
//
// Solidity: function claimablePtp(uint256 , address ) view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) ClaimablePtp(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _PlatypusChefV2.Contract.ClaimablePtp(&_PlatypusChefV2.CallOpts, arg0, arg1)
}

// DialutingRepartition is a free data retrieval call binding the contract method 0x05ed1de4.
//
// Solidity: function dialutingRepartition() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Caller) DialutingRepartition(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "dialutingRepartition")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DialutingRepartition is a free data retrieval call binding the contract method 0x05ed1de4.
//
// Solidity: function dialutingRepartition() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) DialutingRepartition() (*big.Int, error) {
	return _PlatypusChefV2.Contract.DialutingRepartition(&_PlatypusChefV2.CallOpts)
}

// DialutingRepartition is a free data retrieval call binding the contract method 0x05ed1de4.
//
// Solidity: function dialutingRepartition() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) DialutingRepartition() (*big.Int, error) {
	return _PlatypusChefV2.Contract.DialutingRepartition(&_PlatypusChefV2.CallOpts)
}

// MaxPoolLength is a free data retrieval call binding the contract method 0x7b62a738.
//
// Solidity: function maxPoolLength() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Caller) MaxPoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "maxPoolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPoolLength is a free data retrieval call binding the contract method 0x7b62a738.
//
// Solidity: function maxPoolLength() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) MaxPoolLength() (*big.Int, error) {
	return _PlatypusChefV2.Contract.MaxPoolLength(&_PlatypusChefV2.CallOpts)
}

// MaxPoolLength is a free data retrieval call binding the contract method 0x7b62a738.
//
// Solidity: function maxPoolLength() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) MaxPoolLength() (*big.Int, error) {
	return _PlatypusChefV2.Contract.MaxPoolLength(&_PlatypusChefV2.CallOpts)
}

// NewMasterPlatypus is a free data retrieval call binding the contract method 0x01126816.
//
// Solidity: function newMasterPlatypus() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2Caller) NewMasterPlatypus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "newMasterPlatypus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewMasterPlatypus is a free data retrieval call binding the contract method 0x01126816.
//
// Solidity: function newMasterPlatypus() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2Session) NewMasterPlatypus() (common.Address, error) {
	return _PlatypusChefV2.Contract.NewMasterPlatypus(&_PlatypusChefV2.CallOpts)
}

// NewMasterPlatypus is a free data retrieval call binding the contract method 0x01126816.
//
// Solidity: function newMasterPlatypus() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) NewMasterPlatypus() (common.Address, error) {
	return _PlatypusChefV2.Contract.NewMasterPlatypus(&_PlatypusChefV2.CallOpts)
}

// NonDialutingRepartition is a free data retrieval call binding the contract method 0xf87bbc56.
//
// Solidity: function nonDialutingRepartition() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Caller) NonDialutingRepartition(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "nonDialutingRepartition")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonDialutingRepartition is a free data retrieval call binding the contract method 0xf87bbc56.
//
// Solidity: function nonDialutingRepartition() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) NonDialutingRepartition() (*big.Int, error) {
	return _PlatypusChefV2.Contract.NonDialutingRepartition(&_PlatypusChefV2.CallOpts)
}

// NonDialutingRepartition is a free data retrieval call binding the contract method 0xf87bbc56.
//
// Solidity: function nonDialutingRepartition() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) NonDialutingRepartition() (*big.Int, error) {
	return _PlatypusChefV2.Contract.NonDialutingRepartition(&_PlatypusChefV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2Session) Owner() (common.Address, error) {
	return _PlatypusChefV2.Contract.Owner(&_PlatypusChefV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) Owner() (common.Address, error) {
	return _PlatypusChefV2.Contract.Owner(&_PlatypusChefV2.CallOpts)
}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2Caller) OwnerCandidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "ownerCandidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2Session) OwnerCandidate() (common.Address, error) {
	return _PlatypusChefV2.Contract.OwnerCandidate(&_PlatypusChefV2.CallOpts)
}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) OwnerCandidate() (common.Address, error) {
	return _PlatypusChefV2.Contract.OwnerCandidate(&_PlatypusChefV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusChefV2 *PlatypusChefV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusChefV2 *PlatypusChefV2Session) Paused() (bool, error) {
	return _PlatypusChefV2.Contract.Paused(&_PlatypusChefV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) Paused() (bool, error) {
	return _PlatypusChefV2.Contract.Paused(&_PlatypusChefV2.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingPtp, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_PlatypusChefV2 *PlatypusChefV2Caller) PendingTokens(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (struct {
	PendingPtp        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "pendingTokens", _pid, _user)

	outstruct := new(struct {
		PendingPtp        *big.Int
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
		PendingBonusToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingPtp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BonusTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PendingBonusToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingPtp, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_PlatypusChefV2 *PlatypusChefV2Session) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingPtp        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _PlatypusChefV2.Contract.PendingTokens(&_PlatypusChefV2.CallOpts, _pid, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingPtp, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingPtp        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _PlatypusChefV2.Contract.PendingTokens(&_PlatypusChefV2.CallOpts, _pid, _user)
}

// PoolAdjustFactor is a free data retrieval call binding the contract method 0x2b685580.
//
// Solidity: function poolAdjustFactor(uint256 pid) view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Caller) PoolAdjustFactor(opts *bind.CallOpts, pid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "poolAdjustFactor", pid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAdjustFactor is a free data retrieval call binding the contract method 0x2b685580.
//
// Solidity: function poolAdjustFactor(uint256 pid) view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) PoolAdjustFactor(pid *big.Int) (*big.Int, error) {
	return _PlatypusChefV2.Contract.PoolAdjustFactor(&_PlatypusChefV2.CallOpts, pid)
}

// PoolAdjustFactor is a free data retrieval call binding the contract method 0x2b685580.
//
// Solidity: function poolAdjustFactor(uint256 pid) view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) PoolAdjustFactor(pid *big.Int) (*big.Int, error) {
	return _PlatypusChefV2.Contract.PoolAdjustFactor(&_PlatypusChefV2.CallOpts, pid)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 baseAllocPoint, uint256 lastRewardTimestamp, uint256 accPtpPerShare, address rewarder, uint256 sumOfFactors, uint256 accPtpPerFactorShare, uint256 adjustedAllocPoint)
func (_PlatypusChefV2 *PlatypusChefV2Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken              common.Address
	BaseAllocPoint       *big.Int
	LastRewardTimestamp  *big.Int
	AccPtpPerShare       *big.Int
	Rewarder             common.Address
	SumOfFactors         *big.Int
	AccPtpPerFactorShare *big.Int
	AdjustedAllocPoint   *big.Int
}, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken              common.Address
		BaseAllocPoint       *big.Int
		LastRewardTimestamp  *big.Int
		AccPtpPerShare       *big.Int
		Rewarder             common.Address
		SumOfFactors         *big.Int
		AccPtpPerFactorShare *big.Int
		AdjustedAllocPoint   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BaseAllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccPtpPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Rewarder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.SumOfFactors = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.AccPtpPerFactorShare = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.AdjustedAllocPoint = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 baseAllocPoint, uint256 lastRewardTimestamp, uint256 accPtpPerShare, address rewarder, uint256 sumOfFactors, uint256 accPtpPerFactorShare, uint256 adjustedAllocPoint)
func (_PlatypusChefV2 *PlatypusChefV2Session) PoolInfo(arg0 *big.Int) (struct {
	LpToken              common.Address
	BaseAllocPoint       *big.Int
	LastRewardTimestamp  *big.Int
	AccPtpPerShare       *big.Int
	Rewarder             common.Address
	SumOfFactors         *big.Int
	AccPtpPerFactorShare *big.Int
	AdjustedAllocPoint   *big.Int
}, error) {
	return _PlatypusChefV2.Contract.PoolInfo(&_PlatypusChefV2.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 baseAllocPoint, uint256 lastRewardTimestamp, uint256 accPtpPerShare, address rewarder, uint256 sumOfFactors, uint256 accPtpPerFactorShare, uint256 adjustedAllocPoint)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken              common.Address
	BaseAllocPoint       *big.Int
	LastRewardTimestamp  *big.Int
	AccPtpPerShare       *big.Int
	Rewarder             common.Address
	SumOfFactors         *big.Int
	AccPtpPerFactorShare *big.Int
	AdjustedAllocPoint   *big.Int
}, error) {
	return _PlatypusChefV2.Contract.PoolInfo(&_PlatypusChefV2.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) PoolLength() (*big.Int, error) {
	return _PlatypusChefV2.Contract.PoolLength(&_PlatypusChefV2.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) PoolLength() (*big.Int, error) {
	return _PlatypusChefV2.Contract.PoolLength(&_PlatypusChefV2.CallOpts)
}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2Caller) Ptp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "ptp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2Session) Ptp() (common.Address, error) {
	return _PlatypusChefV2.Contract.Ptp(&_PlatypusChefV2.CallOpts)
}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) Ptp() (common.Address, error) {
	return _PlatypusChefV2.Contract.Ptp(&_PlatypusChefV2.CallOpts)
}

// PtpPerSec is a free data retrieval call binding the contract method 0x9702d3e2.
//
// Solidity: function ptpPerSec() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Caller) PtpPerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "ptpPerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PtpPerSec is a free data retrieval call binding the contract method 0x9702d3e2.
//
// Solidity: function ptpPerSec() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) PtpPerSec() (*big.Int, error) {
	return _PlatypusChefV2.Contract.PtpPerSec(&_PlatypusChefV2.CallOpts)
}

// PtpPerSec is a free data retrieval call binding the contract method 0x9702d3e2.
//
// Solidity: function ptpPerSec() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) PtpPerSec() (*big.Int, error) {
	return _PlatypusChefV2.Contract.PtpPerSec(&_PlatypusChefV2.CallOpts)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_PlatypusChefV2 *PlatypusChefV2Caller) RewarderBonusTokenInfo(opts *bind.CallOpts, _pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "rewarderBonusTokenInfo", _pid)

	outstruct := new(struct {
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BonusTokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_PlatypusChefV2 *PlatypusChefV2Session) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	return _PlatypusChefV2.Contract.RewarderBonusTokenInfo(&_PlatypusChefV2.CallOpts, _pid)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	return _PlatypusChefV2.Contract.RewarderBonusTokenInfo(&_PlatypusChefV2.CallOpts, _pid)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Caller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "startTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) StartTimestamp() (*big.Int, error) {
	return _PlatypusChefV2.Contract.StartTimestamp(&_PlatypusChefV2.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) StartTimestamp() (*big.Int, error) {
	return _PlatypusChefV2.Contract.StartTimestamp(&_PlatypusChefV2.CallOpts)
}

// TotalAdjustedAllocPoint is a free data retrieval call binding the contract method 0x91ea1d68.
//
// Solidity: function totalAdjustedAllocPoint() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Caller) TotalAdjustedAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "totalAdjustedAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAdjustedAllocPoint is a free data retrieval call binding the contract method 0x91ea1d68.
//
// Solidity: function totalAdjustedAllocPoint() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) TotalAdjustedAllocPoint() (*big.Int, error) {
	return _PlatypusChefV2.Contract.TotalAdjustedAllocPoint(&_PlatypusChefV2.CallOpts)
}

// TotalAdjustedAllocPoint is a free data retrieval call binding the contract method 0x91ea1d68.
//
// Solidity: function totalAdjustedAllocPoint() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) TotalAdjustedAllocPoint() (*big.Int, error) {
	return _PlatypusChefV2.Contract.TotalAdjustedAllocPoint(&_PlatypusChefV2.CallOpts)
}

// TotalBaseAllocPoint is a free data retrieval call binding the contract method 0x33e045fc.
//
// Solidity: function totalBaseAllocPoint() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Caller) TotalBaseAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "totalBaseAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaseAllocPoint is a free data retrieval call binding the contract method 0x33e045fc.
//
// Solidity: function totalBaseAllocPoint() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) TotalBaseAllocPoint() (*big.Int, error) {
	return _PlatypusChefV2.Contract.TotalBaseAllocPoint(&_PlatypusChefV2.CallOpts)
}

// TotalBaseAllocPoint is a free data retrieval call binding the contract method 0x33e045fc.
//
// Solidity: function totalBaseAllocPoint() view returns(uint256)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) TotalBaseAllocPoint() (*big.Int, error) {
	return _PlatypusChefV2.Contract.TotalBaseAllocPoint(&_PlatypusChefV2.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 factor)
func (_PlatypusChefV2 *PlatypusChefV2Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	Factor     *big.Int
}, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "userInfo", arg0, arg1)

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
func (_PlatypusChefV2 *PlatypusChefV2Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	Factor     *big.Int
}, error) {
	return _PlatypusChefV2.Contract.UserInfo(&_PlatypusChefV2.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 factor)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	Factor     *big.Int
}, error) {
	return _PlatypusChefV2.Contract.UserInfo(&_PlatypusChefV2.CallOpts, arg0, arg1)
}

// VePtp is a free data retrieval call binding the contract method 0x82c780a1.
//
// Solidity: function vePtp() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2Caller) VePtp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusChefV2.contract.Call(opts, &out, "vePtp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VePtp is a free data retrieval call binding the contract method 0x82c780a1.
//
// Solidity: function vePtp() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2Session) VePtp() (common.Address, error) {
	return _PlatypusChefV2.Contract.VePtp(&_PlatypusChefV2.CallOpts)
}

// VePtp is a free data retrieval call binding the contract method 0x82c780a1.
//
// Solidity: function vePtp() view returns(address)
func (_PlatypusChefV2 *PlatypusChefV2CallerSession) VePtp() (common.Address, error) {
	return _PlatypusChefV2.Contract.VePtp(&_PlatypusChefV2.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.AcceptOwnership(&_PlatypusChefV2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.AcceptOwnership(&_PlatypusChefV2.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _baseAllocPoint, address _lpToken, address _rewarder) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) Add(opts *bind.TransactOpts, _baseAllocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "add", _baseAllocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _baseAllocPoint, address _lpToken, address _rewarder) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) Add(_baseAllocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Add(&_PlatypusChefV2.TransactOpts, _baseAllocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _baseAllocPoint, address _lpToken, address _rewarder) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) Add(_baseAllocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Add(&_PlatypusChefV2.TransactOpts, _baseAllocPoint, _lpToken, _rewarder)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV2 *PlatypusChefV2Transactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Deposit(&_PlatypusChefV2.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Deposit(&_PlatypusChefV2.TransactOpts, _pid, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) DepositFor(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "depositFor", _pid, _amount, _user)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) DepositFor(_pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.DepositFor(&_PlatypusChefV2.TransactOpts, _pid, _amount, _user)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) DepositFor(_pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.DepositFor(&_PlatypusChefV2.TransactOpts, _pid, _amount, _user)
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) EmergencyPtpWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "emergencyPtpWithdraw")
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) EmergencyPtpWithdraw() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.EmergencyPtpWithdraw(&_PlatypusChefV2.TransactOpts)
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) EmergencyPtpWithdraw() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.EmergencyPtpWithdraw(&_PlatypusChefV2.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.EmergencyWithdraw(&_PlatypusChefV2.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.EmergencyWithdraw(&_PlatypusChefV2.TransactOpts, _pid)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _ptp, address _vePtp, uint256 _ptpPerSec, uint256 _dialutingRepartition, uint256 _startTimestamp) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) Initialize(opts *bind.TransactOpts, _ptp common.Address, _vePtp common.Address, _ptpPerSec *big.Int, _dialutingRepartition *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "initialize", _ptp, _vePtp, _ptpPerSec, _dialutingRepartition, _startTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _ptp, address _vePtp, uint256 _ptpPerSec, uint256 _dialutingRepartition, uint256 _startTimestamp) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) Initialize(_ptp common.Address, _vePtp common.Address, _ptpPerSec *big.Int, _dialutingRepartition *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Initialize(&_PlatypusChefV2.TransactOpts, _ptp, _vePtp, _ptpPerSec, _dialutingRepartition, _startTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _ptp, address _vePtp, uint256 _ptpPerSec, uint256 _dialutingRepartition, uint256 _startTimestamp) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) Initialize(_ptp common.Address, _vePtp common.Address, _ptpPerSec *big.Int, _dialutingRepartition *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Initialize(&_PlatypusChefV2.TransactOpts, _ptp, _vePtp, _ptpPerSec, _dialutingRepartition, _startTimestamp)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) MassUpdatePools() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.MassUpdatePools(&_PlatypusChefV2.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.MassUpdatePools(&_PlatypusChefV2.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) Migrate(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "migrate", _pids)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) Migrate(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Migrate(&_PlatypusChefV2.TransactOpts, _pids)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) Migrate(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Migrate(&_PlatypusChefV2.TransactOpts, _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[])
func (_PlatypusChefV2 *PlatypusChefV2Transactor) MultiClaim(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "multiClaim", _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[])
func (_PlatypusChefV2 *PlatypusChefV2Session) MultiClaim(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.MultiClaim(&_PlatypusChefV2.TransactOpts, _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[])
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) MultiClaim(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.MultiClaim(&_PlatypusChefV2.TransactOpts, _pids)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) Pause() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Pause(&_PlatypusChefV2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) Pause() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Pause(&_PlatypusChefV2.TransactOpts)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) ProposeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "proposeOwner", newOwner)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) ProposeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.ProposeOwner(&_PlatypusChefV2.TransactOpts, newOwner)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) ProposeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.ProposeOwner(&_PlatypusChefV2.TransactOpts, newOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.RenounceOwnership(&_PlatypusChefV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.RenounceOwnership(&_PlatypusChefV2.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _baseAllocPoint, address _rewarder, bool overwrite) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) Set(opts *bind.TransactOpts, _pid *big.Int, _baseAllocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "set", _pid, _baseAllocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _baseAllocPoint, address _rewarder, bool overwrite) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) Set(_pid *big.Int, _baseAllocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Set(&_PlatypusChefV2.TransactOpts, _pid, _baseAllocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _baseAllocPoint, address _rewarder, bool overwrite) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) Set(_pid *big.Int, _baseAllocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Set(&_PlatypusChefV2.TransactOpts, _pid, _baseAllocPoint, _rewarder, overwrite)
}

// SetMaxPoolLength is a paid mutator transaction binding the contract method 0x53c5eb44.
//
// Solidity: function setMaxPoolLength(uint256 _maxPoolLength) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) SetMaxPoolLength(opts *bind.TransactOpts, _maxPoolLength *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "setMaxPoolLength", _maxPoolLength)
}

// SetMaxPoolLength is a paid mutator transaction binding the contract method 0x53c5eb44.
//
// Solidity: function setMaxPoolLength(uint256 _maxPoolLength) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) SetMaxPoolLength(_maxPoolLength *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.SetMaxPoolLength(&_PlatypusChefV2.TransactOpts, _maxPoolLength)
}

// SetMaxPoolLength is a paid mutator transaction binding the contract method 0x53c5eb44.
//
// Solidity: function setMaxPoolLength(uint256 _maxPoolLength) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) SetMaxPoolLength(_maxPoolLength *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.SetMaxPoolLength(&_PlatypusChefV2.TransactOpts, _maxPoolLength)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) SetNewMasterPlatypus(opts *bind.TransactOpts, _newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "setNewMasterPlatypus", _newMasterPlatypus)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) SetNewMasterPlatypus(_newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.SetNewMasterPlatypus(&_PlatypusChefV2.TransactOpts, _newMasterPlatypus)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) SetNewMasterPlatypus(_newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.SetNewMasterPlatypus(&_PlatypusChefV2.TransactOpts, _newMasterPlatypus)
}

// SetVePtp is a paid mutator transaction binding the contract method 0x90d9c1c3.
//
// Solidity: function setVePtp(address _newVePtp) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) SetVePtp(opts *bind.TransactOpts, _newVePtp common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "setVePtp", _newVePtp)
}

// SetVePtp is a paid mutator transaction binding the contract method 0x90d9c1c3.
//
// Solidity: function setVePtp(address _newVePtp) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) SetVePtp(_newVePtp common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.SetVePtp(&_PlatypusChefV2.TransactOpts, _newVePtp)
}

// SetVePtp is a paid mutator transaction binding the contract method 0x90d9c1c3.
//
// Solidity: function setVePtp(address _newVePtp) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) SetVePtp(_newVePtp common.Address) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.SetVePtp(&_PlatypusChefV2.TransactOpts, _newVePtp)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) Unpause() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Unpause(&_PlatypusChefV2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) Unpause() (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Unpause(&_PlatypusChefV2.TransactOpts)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _ptpPerSec) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) UpdateEmissionRate(opts *bind.TransactOpts, _ptpPerSec *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "updateEmissionRate", _ptpPerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _ptpPerSec) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) UpdateEmissionRate(_ptpPerSec *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.UpdateEmissionRate(&_PlatypusChefV2.TransactOpts, _ptpPerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _ptpPerSec) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) UpdateEmissionRate(_ptpPerSec *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.UpdateEmissionRate(&_PlatypusChefV2.TransactOpts, _ptpPerSec)
}

// UpdateEmissionRepartition is a paid mutator transaction binding the contract method 0xe0a4ed43.
//
// Solidity: function updateEmissionRepartition(uint256 _dialutingRepartition) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) UpdateEmissionRepartition(opts *bind.TransactOpts, _dialutingRepartition *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "updateEmissionRepartition", _dialutingRepartition)
}

// UpdateEmissionRepartition is a paid mutator transaction binding the contract method 0xe0a4ed43.
//
// Solidity: function updateEmissionRepartition(uint256 _dialutingRepartition) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) UpdateEmissionRepartition(_dialutingRepartition *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.UpdateEmissionRepartition(&_PlatypusChefV2.TransactOpts, _dialutingRepartition)
}

// UpdateEmissionRepartition is a paid mutator transaction binding the contract method 0xe0a4ed43.
//
// Solidity: function updateEmissionRepartition(uint256 _dialutingRepartition) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) UpdateEmissionRepartition(_dialutingRepartition *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.UpdateEmissionRepartition(&_PlatypusChefV2.TransactOpts, _dialutingRepartition)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVePtpBalance) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) UpdateFactor(opts *bind.TransactOpts, _user common.Address, _newVePtpBalance *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "updateFactor", _user, _newVePtpBalance)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVePtpBalance) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) UpdateFactor(_user common.Address, _newVePtpBalance *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.UpdateFactor(&_PlatypusChefV2.TransactOpts, _user, _newVePtpBalance)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVePtpBalance) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) UpdateFactor(_user common.Address, _newVePtpBalance *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.UpdateFactor(&_PlatypusChefV2.TransactOpts, _user, _newVePtpBalance)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PlatypusChefV2 *PlatypusChefV2Transactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PlatypusChefV2 *PlatypusChefV2Session) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.UpdatePool(&_PlatypusChefV2.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.UpdatePool(&_PlatypusChefV2.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV2 *PlatypusChefV2Transactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV2 *PlatypusChefV2Session) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Withdraw(&_PlatypusChefV2.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV2 *PlatypusChefV2TransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV2.Contract.Withdraw(&_PlatypusChefV2.TransactOpts, _pid, _amount)
}
