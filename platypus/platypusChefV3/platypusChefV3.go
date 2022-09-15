// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platypusChefV3

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

// PlatypusChefV3MetaData contains all meta data concerning the PlatypusChefV3 contract.
var PlatypusChefV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseAllocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimablePtp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dialutingRepartition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPtpWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ptp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vePtp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ptpPerSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dialutingRepartition\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPoolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"name\":\"multiClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newMasterPlatypus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonDialutingRepartition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCandidate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingPtp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pendingBonusToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"poolAdjustFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAllocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accPtpPerShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewarder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sumOfFactors\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accPtpPerFactorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adjustedAllocPoint\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"proposeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ptp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ptpPerSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"rewarderBonusTokenInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseAllocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPoolLength\",\"type\":\"uint256\"}],\"name\":\"setMaxPoolLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMasterPlatypus\",\"type\":\"address\"}],\"name\":\"setNewMasterPlatypus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVePtp\",\"type\":\"address\"}],\"name\":\"setVePtp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAdjustedAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ptpPerSec\",\"type\":\"uint256\"}],\"name\":\"updateEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dialutingRepartition\",\"type\":\"uint256\"}],\"name\":\"updateEmissionRepartition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newVePtpBalance\",\"type\":\"uint256\"}],\"name\":\"updateFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"factor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vePtp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PlatypusChefV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use PlatypusChefV3MetaData.ABI instead.
var PlatypusChefV3ABI = PlatypusChefV3MetaData.ABI

// PlatypusChefV3 is an auto generated Go binding around an Ethereum contract.
type PlatypusChefV3 struct {
	PlatypusChefV3Caller     // Read-only binding to the contract
	PlatypusChefV3Transactor // Write-only binding to the contract
	PlatypusChefV3Filterer   // Log filterer for contract events
}

// PlatypusChefV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type PlatypusChefV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusChefV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PlatypusChefV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusChefV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlatypusChefV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusChefV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlatypusChefV3Session struct {
	Contract     *PlatypusChefV3   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlatypusChefV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlatypusChefV3CallerSession struct {
	Contract *PlatypusChefV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PlatypusChefV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlatypusChefV3TransactorSession struct {
	Contract     *PlatypusChefV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PlatypusChefV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type PlatypusChefV3Raw struct {
	Contract *PlatypusChefV3 // Generic contract binding to access the raw methods on
}

// PlatypusChefV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlatypusChefV3CallerRaw struct {
	Contract *PlatypusChefV3Caller // Generic read-only contract binding to access the raw methods on
}

// PlatypusChefV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlatypusChefV3TransactorRaw struct {
	Contract *PlatypusChefV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPlatypusChefV3 creates a new instance of PlatypusChefV3, bound to a specific deployed contract.
func NewPlatypusChefV3(address common.Address, backend bind.ContractBackend) (*PlatypusChefV3, error) {
	contract, err := bindPlatypusChefV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlatypusChefV3{PlatypusChefV3Caller: PlatypusChefV3Caller{contract: contract}, PlatypusChefV3Transactor: PlatypusChefV3Transactor{contract: contract}, PlatypusChefV3Filterer: PlatypusChefV3Filterer{contract: contract}}, nil
}

// NewPlatypusChefV3Caller creates a new read-only instance of PlatypusChefV3, bound to a specific deployed contract.
func NewPlatypusChefV3Caller(address common.Address, caller bind.ContractCaller) (*PlatypusChefV3Caller, error) {
	contract, err := bindPlatypusChefV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlatypusChefV3Caller{contract: contract}, nil
}

// NewPlatypusChefV3Transactor creates a new write-only instance of PlatypusChefV3, bound to a specific deployed contract.
func NewPlatypusChefV3Transactor(address common.Address, transactor bind.ContractTransactor) (*PlatypusChefV3Transactor, error) {
	contract, err := bindPlatypusChefV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlatypusChefV3Transactor{contract: contract}, nil
}

// NewPlatypusChefV3Filterer creates a new log filterer instance of PlatypusChefV3, bound to a specific deployed contract.
func NewPlatypusChefV3Filterer(address common.Address, filterer bind.ContractFilterer) (*PlatypusChefV3Filterer, error) {
	contract, err := bindPlatypusChefV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlatypusChefV3Filterer{contract: contract}, nil
}

// bindPlatypusChefV3 binds a generic wrapper to an already deployed contract.
func bindPlatypusChefV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlatypusChefV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatypusChefV3 *PlatypusChefV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatypusChefV3.Contract.PlatypusChefV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatypusChefV3 *PlatypusChefV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.PlatypusChefV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatypusChefV3 *PlatypusChefV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.PlatypusChefV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatypusChefV3 *PlatypusChefV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatypusChefV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatypusChefV3 *PlatypusChefV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatypusChefV3 *PlatypusChefV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.contract.Transact(opts, method, params...)
}

// ClaimablePtp is a free data retrieval call binding the contract method 0x8b4d83a3.
//
// Solidity: function claimablePtp(uint256 , address ) view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Caller) ClaimablePtp(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "claimablePtp", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimablePtp is a free data retrieval call binding the contract method 0x8b4d83a3.
//
// Solidity: function claimablePtp(uint256 , address ) view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) ClaimablePtp(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _PlatypusChefV3.Contract.ClaimablePtp(&_PlatypusChefV3.CallOpts, arg0, arg1)
}

// ClaimablePtp is a free data retrieval call binding the contract method 0x8b4d83a3.
//
// Solidity: function claimablePtp(uint256 , address ) view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) ClaimablePtp(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _PlatypusChefV3.Contract.ClaimablePtp(&_PlatypusChefV3.CallOpts, arg0, arg1)
}

// DialutingRepartition is a free data retrieval call binding the contract method 0x05ed1de4.
//
// Solidity: function dialutingRepartition() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Caller) DialutingRepartition(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "dialutingRepartition")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DialutingRepartition is a free data retrieval call binding the contract method 0x05ed1de4.
//
// Solidity: function dialutingRepartition() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) DialutingRepartition() (*big.Int, error) {
	return _PlatypusChefV3.Contract.DialutingRepartition(&_PlatypusChefV3.CallOpts)
}

// DialutingRepartition is a free data retrieval call binding the contract method 0x05ed1de4.
//
// Solidity: function dialutingRepartition() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) DialutingRepartition() (*big.Int, error) {
	return _PlatypusChefV3.Contract.DialutingRepartition(&_PlatypusChefV3.CallOpts)
}

// MaxPoolLength is a free data retrieval call binding the contract method 0x7b62a738.
//
// Solidity: function maxPoolLength() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Caller) MaxPoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "maxPoolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPoolLength is a free data retrieval call binding the contract method 0x7b62a738.
//
// Solidity: function maxPoolLength() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) MaxPoolLength() (*big.Int, error) {
	return _PlatypusChefV3.Contract.MaxPoolLength(&_PlatypusChefV3.CallOpts)
}

// MaxPoolLength is a free data retrieval call binding the contract method 0x7b62a738.
//
// Solidity: function maxPoolLength() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) MaxPoolLength() (*big.Int, error) {
	return _PlatypusChefV3.Contract.MaxPoolLength(&_PlatypusChefV3.CallOpts)
}

// NewMasterPlatypus is a free data retrieval call binding the contract method 0x01126816.
//
// Solidity: function newMasterPlatypus() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3Caller) NewMasterPlatypus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "newMasterPlatypus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewMasterPlatypus is a free data retrieval call binding the contract method 0x01126816.
//
// Solidity: function newMasterPlatypus() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3Session) NewMasterPlatypus() (common.Address, error) {
	return _PlatypusChefV3.Contract.NewMasterPlatypus(&_PlatypusChefV3.CallOpts)
}

// NewMasterPlatypus is a free data retrieval call binding the contract method 0x01126816.
//
// Solidity: function newMasterPlatypus() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) NewMasterPlatypus() (common.Address, error) {
	return _PlatypusChefV3.Contract.NewMasterPlatypus(&_PlatypusChefV3.CallOpts)
}

// NonDialutingRepartition is a free data retrieval call binding the contract method 0xf87bbc56.
//
// Solidity: function nonDialutingRepartition() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Caller) NonDialutingRepartition(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "nonDialutingRepartition")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonDialutingRepartition is a free data retrieval call binding the contract method 0xf87bbc56.
//
// Solidity: function nonDialutingRepartition() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) NonDialutingRepartition() (*big.Int, error) {
	return _PlatypusChefV3.Contract.NonDialutingRepartition(&_PlatypusChefV3.CallOpts)
}

// NonDialutingRepartition is a free data retrieval call binding the contract method 0xf87bbc56.
//
// Solidity: function nonDialutingRepartition() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) NonDialutingRepartition() (*big.Int, error) {
	return _PlatypusChefV3.Contract.NonDialutingRepartition(&_PlatypusChefV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3Session) Owner() (common.Address, error) {
	return _PlatypusChefV3.Contract.Owner(&_PlatypusChefV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) Owner() (common.Address, error) {
	return _PlatypusChefV3.Contract.Owner(&_PlatypusChefV3.CallOpts)
}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3Caller) OwnerCandidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "ownerCandidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3Session) OwnerCandidate() (common.Address, error) {
	return _PlatypusChefV3.Contract.OwnerCandidate(&_PlatypusChefV3.CallOpts)
}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) OwnerCandidate() (common.Address, error) {
	return _PlatypusChefV3.Contract.OwnerCandidate(&_PlatypusChefV3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusChefV3 *PlatypusChefV3Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusChefV3 *PlatypusChefV3Session) Paused() (bool, error) {
	return _PlatypusChefV3.Contract.Paused(&_PlatypusChefV3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) Paused() (bool, error) {
	return _PlatypusChefV3.Contract.Paused(&_PlatypusChefV3.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingPtp, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_PlatypusChefV3 *PlatypusChefV3Caller) PendingTokens(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (struct {
	PendingPtp        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "pendingTokens", _pid, _user)

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
func (_PlatypusChefV3 *PlatypusChefV3Session) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingPtp        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _PlatypusChefV3.Contract.PendingTokens(&_PlatypusChefV3.CallOpts, _pid, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingPtp, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingPtp        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _PlatypusChefV3.Contract.PendingTokens(&_PlatypusChefV3.CallOpts, _pid, _user)
}

// PoolAdjustFactor is a free data retrieval call binding the contract method 0x2b685580.
//
// Solidity: function poolAdjustFactor(uint256 pid) view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Caller) PoolAdjustFactor(opts *bind.CallOpts, pid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "poolAdjustFactor", pid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAdjustFactor is a free data retrieval call binding the contract method 0x2b685580.
//
// Solidity: function poolAdjustFactor(uint256 pid) view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) PoolAdjustFactor(pid *big.Int) (*big.Int, error) {
	return _PlatypusChefV3.Contract.PoolAdjustFactor(&_PlatypusChefV3.CallOpts, pid)
}

// PoolAdjustFactor is a free data retrieval call binding the contract method 0x2b685580.
//
// Solidity: function poolAdjustFactor(uint256 pid) view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) PoolAdjustFactor(pid *big.Int) (*big.Int, error) {
	return _PlatypusChefV3.Contract.PoolAdjustFactor(&_PlatypusChefV3.CallOpts, pid)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 baseAllocPoint, uint256 lastRewardTimestamp, uint256 accPtpPerShare, address rewarder, uint256 sumOfFactors, uint256 accPtpPerFactorShare, uint256 adjustedAllocPoint)
func (_PlatypusChefV3 *PlatypusChefV3Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _PlatypusChefV3.contract.Call(opts, &out, "poolInfo", arg0)

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
func (_PlatypusChefV3 *PlatypusChefV3Session) PoolInfo(arg0 *big.Int) (struct {
	LpToken              common.Address
	BaseAllocPoint       *big.Int
	LastRewardTimestamp  *big.Int
	AccPtpPerShare       *big.Int
	Rewarder             common.Address
	SumOfFactors         *big.Int
	AccPtpPerFactorShare *big.Int
	AdjustedAllocPoint   *big.Int
}, error) {
	return _PlatypusChefV3.Contract.PoolInfo(&_PlatypusChefV3.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 baseAllocPoint, uint256 lastRewardTimestamp, uint256 accPtpPerShare, address rewarder, uint256 sumOfFactors, uint256 accPtpPerFactorShare, uint256 adjustedAllocPoint)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken              common.Address
	BaseAllocPoint       *big.Int
	LastRewardTimestamp  *big.Int
	AccPtpPerShare       *big.Int
	Rewarder             common.Address
	SumOfFactors         *big.Int
	AccPtpPerFactorShare *big.Int
	AdjustedAllocPoint   *big.Int
}, error) {
	return _PlatypusChefV3.Contract.PoolInfo(&_PlatypusChefV3.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) PoolLength() (*big.Int, error) {
	return _PlatypusChefV3.Contract.PoolLength(&_PlatypusChefV3.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) PoolLength() (*big.Int, error) {
	return _PlatypusChefV3.Contract.PoolLength(&_PlatypusChefV3.CallOpts)
}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3Caller) Ptp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "ptp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3Session) Ptp() (common.Address, error) {
	return _PlatypusChefV3.Contract.Ptp(&_PlatypusChefV3.CallOpts)
}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) Ptp() (common.Address, error) {
	return _PlatypusChefV3.Contract.Ptp(&_PlatypusChefV3.CallOpts)
}

// PtpPerSec is a free data retrieval call binding the contract method 0x9702d3e2.
//
// Solidity: function ptpPerSec() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Caller) PtpPerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "ptpPerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PtpPerSec is a free data retrieval call binding the contract method 0x9702d3e2.
//
// Solidity: function ptpPerSec() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) PtpPerSec() (*big.Int, error) {
	return _PlatypusChefV3.Contract.PtpPerSec(&_PlatypusChefV3.CallOpts)
}

// PtpPerSec is a free data retrieval call binding the contract method 0x9702d3e2.
//
// Solidity: function ptpPerSec() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) PtpPerSec() (*big.Int, error) {
	return _PlatypusChefV3.Contract.PtpPerSec(&_PlatypusChefV3.CallOpts)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_PlatypusChefV3 *PlatypusChefV3Caller) RewarderBonusTokenInfo(opts *bind.CallOpts, _pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "rewarderBonusTokenInfo", _pid)

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
func (_PlatypusChefV3 *PlatypusChefV3Session) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	return _PlatypusChefV3.Contract.RewarderBonusTokenInfo(&_PlatypusChefV3.CallOpts, _pid)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	return _PlatypusChefV3.Contract.RewarderBonusTokenInfo(&_PlatypusChefV3.CallOpts, _pid)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Caller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "startTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) StartTimestamp() (*big.Int, error) {
	return _PlatypusChefV3.Contract.StartTimestamp(&_PlatypusChefV3.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) StartTimestamp() (*big.Int, error) {
	return _PlatypusChefV3.Contract.StartTimestamp(&_PlatypusChefV3.CallOpts)
}

// TotalAdjustedAllocPoint is a free data retrieval call binding the contract method 0x91ea1d68.
//
// Solidity: function totalAdjustedAllocPoint() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Caller) TotalAdjustedAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "totalAdjustedAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAdjustedAllocPoint is a free data retrieval call binding the contract method 0x91ea1d68.
//
// Solidity: function totalAdjustedAllocPoint() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) TotalAdjustedAllocPoint() (*big.Int, error) {
	return _PlatypusChefV3.Contract.TotalAdjustedAllocPoint(&_PlatypusChefV3.CallOpts)
}

// TotalAdjustedAllocPoint is a free data retrieval call binding the contract method 0x91ea1d68.
//
// Solidity: function totalAdjustedAllocPoint() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) TotalAdjustedAllocPoint() (*big.Int, error) {
	return _PlatypusChefV3.Contract.TotalAdjustedAllocPoint(&_PlatypusChefV3.CallOpts)
}

// TotalBaseAllocPoint is a free data retrieval call binding the contract method 0x33e045fc.
//
// Solidity: function totalBaseAllocPoint() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Caller) TotalBaseAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "totalBaseAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaseAllocPoint is a free data retrieval call binding the contract method 0x33e045fc.
//
// Solidity: function totalBaseAllocPoint() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) TotalBaseAllocPoint() (*big.Int, error) {
	return _PlatypusChefV3.Contract.TotalBaseAllocPoint(&_PlatypusChefV3.CallOpts)
}

// TotalBaseAllocPoint is a free data retrieval call binding the contract method 0x33e045fc.
//
// Solidity: function totalBaseAllocPoint() view returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) TotalBaseAllocPoint() (*big.Int, error) {
	return _PlatypusChefV3.Contract.TotalBaseAllocPoint(&_PlatypusChefV3.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 factor)
func (_PlatypusChefV3 *PlatypusChefV3Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	Factor     *big.Int
}, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "userInfo", arg0, arg1)

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
func (_PlatypusChefV3 *PlatypusChefV3Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	Factor     *big.Int
}, error) {
	return _PlatypusChefV3.Contract.UserInfo(&_PlatypusChefV3.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 factor)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	Factor     *big.Int
}, error) {
	return _PlatypusChefV3.Contract.UserInfo(&_PlatypusChefV3.CallOpts, arg0, arg1)
}

// VePtp is a free data retrieval call binding the contract method 0x82c780a1.
//
// Solidity: function vePtp() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3Caller) VePtp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "vePtp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VePtp is a free data retrieval call binding the contract method 0x82c780a1.
//
// Solidity: function vePtp() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3Session) VePtp() (common.Address, error) {
	return _PlatypusChefV3.Contract.VePtp(&_PlatypusChefV3.CallOpts)
}

// VePtp is a free data retrieval call binding the contract method 0x82c780a1.
//
// Solidity: function vePtp() view returns(address)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) VePtp() (common.Address, error) {
	return _PlatypusChefV3.Contract.VePtp(&_PlatypusChefV3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusChefV3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) Version() (*big.Int, error) {
	return _PlatypusChefV3.Contract.Version(&_PlatypusChefV3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_PlatypusChefV3 *PlatypusChefV3CallerSession) Version() (*big.Int, error) {
	return _PlatypusChefV3.Contract.Version(&_PlatypusChefV3.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) AcceptOwnership() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.AcceptOwnership(&_PlatypusChefV3.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.AcceptOwnership(&_PlatypusChefV3.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _baseAllocPoint, address _lpToken, address _rewarder) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) Add(opts *bind.TransactOpts, _baseAllocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "add", _baseAllocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _baseAllocPoint, address _lpToken, address _rewarder) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) Add(_baseAllocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Add(&_PlatypusChefV3.TransactOpts, _baseAllocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _baseAllocPoint, address _lpToken, address _rewarder) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) Add(_baseAllocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Add(&_PlatypusChefV3.TransactOpts, _baseAllocPoint, _lpToken, _rewarder)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV3 *PlatypusChefV3Transactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Deposit(&_PlatypusChefV3.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Deposit(&_PlatypusChefV3.TransactOpts, _pid, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) DepositFor(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "depositFor", _pid, _amount, _user)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) DepositFor(_pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.DepositFor(&_PlatypusChefV3.TransactOpts, _pid, _amount, _user)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) DepositFor(_pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.DepositFor(&_PlatypusChefV3.TransactOpts, _pid, _amount, _user)
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) EmergencyPtpWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "emergencyPtpWithdraw")
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) EmergencyPtpWithdraw() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.EmergencyPtpWithdraw(&_PlatypusChefV3.TransactOpts)
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) EmergencyPtpWithdraw() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.EmergencyPtpWithdraw(&_PlatypusChefV3.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.EmergencyWithdraw(&_PlatypusChefV3.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.EmergencyWithdraw(&_PlatypusChefV3.TransactOpts, _pid)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _ptp, address _vePtp, uint256 _ptpPerSec, uint256 _dialutingRepartition, uint256 _startTimestamp) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) Initialize(opts *bind.TransactOpts, _ptp common.Address, _vePtp common.Address, _ptpPerSec *big.Int, _dialutingRepartition *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "initialize", _ptp, _vePtp, _ptpPerSec, _dialutingRepartition, _startTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _ptp, address _vePtp, uint256 _ptpPerSec, uint256 _dialutingRepartition, uint256 _startTimestamp) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) Initialize(_ptp common.Address, _vePtp common.Address, _ptpPerSec *big.Int, _dialutingRepartition *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Initialize(&_PlatypusChefV3.TransactOpts, _ptp, _vePtp, _ptpPerSec, _dialutingRepartition, _startTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _ptp, address _vePtp, uint256 _ptpPerSec, uint256 _dialutingRepartition, uint256 _startTimestamp) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) Initialize(_ptp common.Address, _vePtp common.Address, _ptpPerSec *big.Int, _dialutingRepartition *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Initialize(&_PlatypusChefV3.TransactOpts, _ptp, _vePtp, _ptpPerSec, _dialutingRepartition, _startTimestamp)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) MassUpdatePools() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.MassUpdatePools(&_PlatypusChefV3.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.MassUpdatePools(&_PlatypusChefV3.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) Migrate(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "migrate", _pids)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) Migrate(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Migrate(&_PlatypusChefV3.TransactOpts, _pids)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) Migrate(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Migrate(&_PlatypusChefV3.TransactOpts, _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[])
func (_PlatypusChefV3 *PlatypusChefV3Transactor) MultiClaim(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "multiClaim", _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[])
func (_PlatypusChefV3 *PlatypusChefV3Session) MultiClaim(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.MultiClaim(&_PlatypusChefV3.TransactOpts, _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[])
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) MultiClaim(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.MultiClaim(&_PlatypusChefV3.TransactOpts, _pids)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) Pause() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Pause(&_PlatypusChefV3.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) Pause() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Pause(&_PlatypusChefV3.TransactOpts)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) ProposeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "proposeOwner", newOwner)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) ProposeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.ProposeOwner(&_PlatypusChefV3.TransactOpts, newOwner)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) ProposeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.ProposeOwner(&_PlatypusChefV3.TransactOpts, newOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.RenounceOwnership(&_PlatypusChefV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.RenounceOwnership(&_PlatypusChefV3.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _baseAllocPoint, address _rewarder, bool overwrite) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) Set(opts *bind.TransactOpts, _pid *big.Int, _baseAllocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "set", _pid, _baseAllocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _baseAllocPoint, address _rewarder, bool overwrite) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) Set(_pid *big.Int, _baseAllocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Set(&_PlatypusChefV3.TransactOpts, _pid, _baseAllocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _baseAllocPoint, address _rewarder, bool overwrite) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) Set(_pid *big.Int, _baseAllocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Set(&_PlatypusChefV3.TransactOpts, _pid, _baseAllocPoint, _rewarder, overwrite)
}

// SetMaxPoolLength is a paid mutator transaction binding the contract method 0x53c5eb44.
//
// Solidity: function setMaxPoolLength(uint256 _maxPoolLength) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) SetMaxPoolLength(opts *bind.TransactOpts, _maxPoolLength *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "setMaxPoolLength", _maxPoolLength)
}

// SetMaxPoolLength is a paid mutator transaction binding the contract method 0x53c5eb44.
//
// Solidity: function setMaxPoolLength(uint256 _maxPoolLength) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) SetMaxPoolLength(_maxPoolLength *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.SetMaxPoolLength(&_PlatypusChefV3.TransactOpts, _maxPoolLength)
}

// SetMaxPoolLength is a paid mutator transaction binding the contract method 0x53c5eb44.
//
// Solidity: function setMaxPoolLength(uint256 _maxPoolLength) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) SetMaxPoolLength(_maxPoolLength *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.SetMaxPoolLength(&_PlatypusChefV3.TransactOpts, _maxPoolLength)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) SetNewMasterPlatypus(opts *bind.TransactOpts, _newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "setNewMasterPlatypus", _newMasterPlatypus)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) SetNewMasterPlatypus(_newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.SetNewMasterPlatypus(&_PlatypusChefV3.TransactOpts, _newMasterPlatypus)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) SetNewMasterPlatypus(_newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.SetNewMasterPlatypus(&_PlatypusChefV3.TransactOpts, _newMasterPlatypus)
}

// SetVePtp is a paid mutator transaction binding the contract method 0x90d9c1c3.
//
// Solidity: function setVePtp(address _newVePtp) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) SetVePtp(opts *bind.TransactOpts, _newVePtp common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "setVePtp", _newVePtp)
}

// SetVePtp is a paid mutator transaction binding the contract method 0x90d9c1c3.
//
// Solidity: function setVePtp(address _newVePtp) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) SetVePtp(_newVePtp common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.SetVePtp(&_PlatypusChefV3.TransactOpts, _newVePtp)
}

// SetVePtp is a paid mutator transaction binding the contract method 0x90d9c1c3.
//
// Solidity: function setVePtp(address _newVePtp) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) SetVePtp(_newVePtp common.Address) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.SetVePtp(&_PlatypusChefV3.TransactOpts, _newVePtp)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) Unpause() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Unpause(&_PlatypusChefV3.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) Unpause() (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Unpause(&_PlatypusChefV3.TransactOpts)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _ptpPerSec) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) UpdateEmissionRate(opts *bind.TransactOpts, _ptpPerSec *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "updateEmissionRate", _ptpPerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _ptpPerSec) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) UpdateEmissionRate(_ptpPerSec *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.UpdateEmissionRate(&_PlatypusChefV3.TransactOpts, _ptpPerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _ptpPerSec) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) UpdateEmissionRate(_ptpPerSec *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.UpdateEmissionRate(&_PlatypusChefV3.TransactOpts, _ptpPerSec)
}

// UpdateEmissionRepartition is a paid mutator transaction binding the contract method 0xe0a4ed43.
//
// Solidity: function updateEmissionRepartition(uint256 _dialutingRepartition) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) UpdateEmissionRepartition(opts *bind.TransactOpts, _dialutingRepartition *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "updateEmissionRepartition", _dialutingRepartition)
}

// UpdateEmissionRepartition is a paid mutator transaction binding the contract method 0xe0a4ed43.
//
// Solidity: function updateEmissionRepartition(uint256 _dialutingRepartition) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) UpdateEmissionRepartition(_dialutingRepartition *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.UpdateEmissionRepartition(&_PlatypusChefV3.TransactOpts, _dialutingRepartition)
}

// UpdateEmissionRepartition is a paid mutator transaction binding the contract method 0xe0a4ed43.
//
// Solidity: function updateEmissionRepartition(uint256 _dialutingRepartition) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) UpdateEmissionRepartition(_dialutingRepartition *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.UpdateEmissionRepartition(&_PlatypusChefV3.TransactOpts, _dialutingRepartition)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVePtpBalance) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) UpdateFactor(opts *bind.TransactOpts, _user common.Address, _newVePtpBalance *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "updateFactor", _user, _newVePtpBalance)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVePtpBalance) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) UpdateFactor(_user common.Address, _newVePtpBalance *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.UpdateFactor(&_PlatypusChefV3.TransactOpts, _user, _newVePtpBalance)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVePtpBalance) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) UpdateFactor(_user common.Address, _newVePtpBalance *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.UpdateFactor(&_PlatypusChefV3.TransactOpts, _user, _newVePtpBalance)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PlatypusChefV3 *PlatypusChefV3Transactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PlatypusChefV3 *PlatypusChefV3Session) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.UpdatePool(&_PlatypusChefV3.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.UpdatePool(&_PlatypusChefV3.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV3 *PlatypusChefV3Transactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV3 *PlatypusChefV3Session) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Withdraw(&_PlatypusChefV3.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256)
func (_PlatypusChefV3 *PlatypusChefV3TransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusChefV3.Contract.Withdraw(&_PlatypusChefV3.TransactOpts, _pid, _amount)
}
