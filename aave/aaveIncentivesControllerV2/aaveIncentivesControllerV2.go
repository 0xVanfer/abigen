// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveIncentivesControllerV2

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

// AaveIncentivesControllerV2MetaData contains all meta data concerning the AaveIncentivesControllerV2 contract.
var AaveIncentivesControllerV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DISTRIBUTION_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMISSION_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"uint104\",\"name\":\"emissionPerSecond\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"index\",\"type\":\"uint104\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claimRewardsOnBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"emissionsPerSecond\",\"type\":\"uint256[]\"}],\"name\":\"configureAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getClaimer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDistributionEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getRewardsBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardsVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getUserAssetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserUnclaimedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBalance\",\"type\":\"uint256\"}],\"name\":\"handleAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardsVault\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"setClaimer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionEnd\",\"type\":\"uint256\"}],\"name\":\"setDistributionEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardsVault\",\"type\":\"address\"}],\"name\":\"setRewardsVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveIncentivesControllerV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveIncentivesControllerV2MetaData.ABI instead.
var AaveIncentivesControllerV2ABI = AaveIncentivesControllerV2MetaData.ABI

// AaveIncentivesControllerV2 is an auto generated Go binding around an Ethereum contract.
type AaveIncentivesControllerV2 struct {
	AaveIncentivesControllerV2Caller     // Read-only binding to the contract
	AaveIncentivesControllerV2Transactor // Write-only binding to the contract
	AaveIncentivesControllerV2Filterer   // Log filterer for contract events
}

// AaveIncentivesControllerV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveIncentivesControllerV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveIncentivesControllerV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveIncentivesControllerV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveIncentivesControllerV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveIncentivesControllerV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveIncentivesControllerV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveIncentivesControllerV2Session struct {
	Contract     *AaveIncentivesControllerV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AaveIncentivesControllerV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveIncentivesControllerV2CallerSession struct {
	Contract *AaveIncentivesControllerV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// AaveIncentivesControllerV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveIncentivesControllerV2TransactorSession struct {
	Contract     *AaveIncentivesControllerV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// AaveIncentivesControllerV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveIncentivesControllerV2Raw struct {
	Contract *AaveIncentivesControllerV2 // Generic contract binding to access the raw methods on
}

// AaveIncentivesControllerV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveIncentivesControllerV2CallerRaw struct {
	Contract *AaveIncentivesControllerV2Caller // Generic read-only contract binding to access the raw methods on
}

// AaveIncentivesControllerV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveIncentivesControllerV2TransactorRaw struct {
	Contract *AaveIncentivesControllerV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveIncentivesControllerV2 creates a new instance of AaveIncentivesControllerV2, bound to a specific deployed contract.
func NewAaveIncentivesControllerV2(address common.Address, backend bind.ContractBackend) (*AaveIncentivesControllerV2, error) {
	contract, err := bindAaveIncentivesControllerV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerV2{AaveIncentivesControllerV2Caller: AaveIncentivesControllerV2Caller{contract: contract}, AaveIncentivesControllerV2Transactor: AaveIncentivesControllerV2Transactor{contract: contract}, AaveIncentivesControllerV2Filterer: AaveIncentivesControllerV2Filterer{contract: contract}}, nil
}

// NewAaveIncentivesControllerV2Caller creates a new read-only instance of AaveIncentivesControllerV2, bound to a specific deployed contract.
func NewAaveIncentivesControllerV2Caller(address common.Address, caller bind.ContractCaller) (*AaveIncentivesControllerV2Caller, error) {
	contract, err := bindAaveIncentivesControllerV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerV2Caller{contract: contract}, nil
}

// NewAaveIncentivesControllerV2Transactor creates a new write-only instance of AaveIncentivesControllerV2, bound to a specific deployed contract.
func NewAaveIncentivesControllerV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveIncentivesControllerV2Transactor, error) {
	contract, err := bindAaveIncentivesControllerV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerV2Transactor{contract: contract}, nil
}

// NewAaveIncentivesControllerV2Filterer creates a new log filterer instance of AaveIncentivesControllerV2, bound to a specific deployed contract.
func NewAaveIncentivesControllerV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveIncentivesControllerV2Filterer, error) {
	contract, err := bindAaveIncentivesControllerV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerV2Filterer{contract: contract}, nil
}

// bindAaveIncentivesControllerV2 binds a generic wrapper to an already deployed contract.
func bindAaveIncentivesControllerV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveIncentivesControllerV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveIncentivesControllerV2.Contract.AaveIncentivesControllerV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.AaveIncentivesControllerV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.AaveIncentivesControllerV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveIncentivesControllerV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.contract.Transact(opts, method, params...)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) DISTRIBUTIONEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "DISTRIBUTION_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) DISTRIBUTIONEND() (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.DISTRIBUTIONEND(&_AaveIncentivesControllerV2.CallOpts)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.DISTRIBUTIONEND(&_AaveIncentivesControllerV2.CallOpts)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) EMISSIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "EMISSION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) EMISSIONMANAGER() (common.Address, error) {
	return _AaveIncentivesControllerV2.Contract.EMISSIONMANAGER(&_AaveIncentivesControllerV2.CallOpts)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) EMISSIONMANAGER() (common.Address, error) {
	return _AaveIncentivesControllerV2.Contract.EMISSIONMANAGER(&_AaveIncentivesControllerV2.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) PRECISION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) PRECISION() (uint8, error) {
	return _AaveIncentivesControllerV2.Contract.PRECISION(&_AaveIncentivesControllerV2.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) PRECISION() (uint8, error) {
	return _AaveIncentivesControllerV2.Contract.PRECISION(&_AaveIncentivesControllerV2.CallOpts)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) REVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) REVISION() (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.REVISION(&_AaveIncentivesControllerV2.CallOpts)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) REVISION() (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.REVISION(&_AaveIncentivesControllerV2.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) REWARDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "REWARD_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) REWARDTOKEN() (common.Address, error) {
	return _AaveIncentivesControllerV2.Contract.REWARDTOKEN(&_AaveIncentivesControllerV2.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) REWARDTOKEN() (common.Address, error) {
	return _AaveIncentivesControllerV2.Contract.REWARDTOKEN(&_AaveIncentivesControllerV2.CallOpts)
}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint104 emissionPerSecond, uint104 index, uint40 lastUpdateTimestamp)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) Assets(opts *bind.CallOpts, arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	Index               *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "assets", arg0)

	outstruct := new(struct {
		EmissionPerSecond   *big.Int
		Index               *big.Int
		LastUpdateTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EmissionPerSecond = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint104 emissionPerSecond, uint104 index, uint40 lastUpdateTimestamp)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) Assets(arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	Index               *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _AaveIncentivesControllerV2.Contract.Assets(&_AaveIncentivesControllerV2.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint104 emissionPerSecond, uint104 index, uint40 lastUpdateTimestamp)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) Assets(arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	Index               *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _AaveIncentivesControllerV2.Contract.Assets(&_AaveIncentivesControllerV2.CallOpts, arg0)
}

// GetAssetData is a free data retrieval call binding the contract method 0x1652e7b7.
//
// Solidity: function getAssetData(address asset) view returns(uint256, uint256, uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) GetAssetData(opts *bind.CallOpts, asset common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "getAssetData", asset)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAssetData is a free data retrieval call binding the contract method 0x1652e7b7.
//
// Solidity: function getAssetData(address asset) view returns(uint256, uint256, uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) GetAssetData(asset common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.GetAssetData(&_AaveIncentivesControllerV2.CallOpts, asset)
}

// GetAssetData is a free data retrieval call binding the contract method 0x1652e7b7.
//
// Solidity: function getAssetData(address asset) view returns(uint256, uint256, uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) GetAssetData(asset common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.GetAssetData(&_AaveIncentivesControllerV2.CallOpts, asset)
}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) GetClaimer(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "getClaimer", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) GetClaimer(user common.Address) (common.Address, error) {
	return _AaveIncentivesControllerV2.Contract.GetClaimer(&_AaveIncentivesControllerV2.CallOpts, user)
}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) GetClaimer(user common.Address) (common.Address, error) {
	return _AaveIncentivesControllerV2.Contract.GetClaimer(&_AaveIncentivesControllerV2.CallOpts, user)
}

// GetDistributionEnd is a free data retrieval call binding the contract method 0xcc69afec.
//
// Solidity: function getDistributionEnd() view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) GetDistributionEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "getDistributionEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionEnd is a free data retrieval call binding the contract method 0xcc69afec.
//
// Solidity: function getDistributionEnd() view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) GetDistributionEnd() (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.GetDistributionEnd(&_AaveIncentivesControllerV2.CallOpts)
}

// GetDistributionEnd is a free data retrieval call binding the contract method 0xcc69afec.
//
// Solidity: function getDistributionEnd() view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) GetDistributionEnd() (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.GetDistributionEnd(&_AaveIncentivesControllerV2.CallOpts)
}

// GetRewardsBalance is a free data retrieval call binding the contract method 0x8b599f26.
//
// Solidity: function getRewardsBalance(address[] assets, address user) view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) GetRewardsBalance(opts *bind.CallOpts, assets []common.Address, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "getRewardsBalance", assets, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardsBalance is a free data retrieval call binding the contract method 0x8b599f26.
//
// Solidity: function getRewardsBalance(address[] assets, address user) view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) GetRewardsBalance(assets []common.Address, user common.Address) (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.GetRewardsBalance(&_AaveIncentivesControllerV2.CallOpts, assets, user)
}

// GetRewardsBalance is a free data retrieval call binding the contract method 0x8b599f26.
//
// Solidity: function getRewardsBalance(address[] assets, address user) view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) GetRewardsBalance(assets []common.Address, user common.Address) (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.GetRewardsBalance(&_AaveIncentivesControllerV2.CallOpts, assets, user)
}

// GetRewardsVault is a free data retrieval call binding the contract method 0xe23ddec5.
//
// Solidity: function getRewardsVault() view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) GetRewardsVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "getRewardsVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardsVault is a free data retrieval call binding the contract method 0xe23ddec5.
//
// Solidity: function getRewardsVault() view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) GetRewardsVault() (common.Address, error) {
	return _AaveIncentivesControllerV2.Contract.GetRewardsVault(&_AaveIncentivesControllerV2.CallOpts)
}

// GetRewardsVault is a free data retrieval call binding the contract method 0xe23ddec5.
//
// Solidity: function getRewardsVault() view returns(address)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) GetRewardsVault() (common.Address, error) {
	return _AaveIncentivesControllerV2.Contract.GetRewardsVault(&_AaveIncentivesControllerV2.CallOpts)
}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) GetUserAssetData(opts *bind.CallOpts, user common.Address, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "getUserAssetData", user, asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) GetUserAssetData(user common.Address, asset common.Address) (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.GetUserAssetData(&_AaveIncentivesControllerV2.CallOpts, user, asset)
}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) GetUserAssetData(user common.Address, asset common.Address) (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.GetUserAssetData(&_AaveIncentivesControllerV2.CallOpts, user, asset)
}

// GetUserUnclaimedRewards is a free data retrieval call binding the contract method 0x198fa81e.
//
// Solidity: function getUserUnclaimedRewards(address _user) view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Caller) GetUserUnclaimedRewards(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesControllerV2.contract.Call(opts, &out, "getUserUnclaimedRewards", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserUnclaimedRewards is a free data retrieval call binding the contract method 0x198fa81e.
//
// Solidity: function getUserUnclaimedRewards(address _user) view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) GetUserUnclaimedRewards(_user common.Address) (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.GetUserUnclaimedRewards(&_AaveIncentivesControllerV2.CallOpts, _user)
}

// GetUserUnclaimedRewards is a free data retrieval call binding the contract method 0x198fa81e.
//
// Solidity: function getUserUnclaimedRewards(address _user) view returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2CallerSession) GetUserUnclaimedRewards(_user common.Address) (*big.Int, error) {
	return _AaveIncentivesControllerV2.Contract.GetUserUnclaimedRewards(&_AaveIncentivesControllerV2.CallOpts, _user)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x3111e7b3.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to) returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Transactor) ClaimRewards(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.contract.Transact(opts, "claimRewards", assets, amount, to)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x3111e7b3.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to) returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) ClaimRewards(assets []common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.ClaimRewards(&_AaveIncentivesControllerV2.TransactOpts, assets, amount, to)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x3111e7b3.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to) returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2TransactorSession) ClaimRewards(assets []common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.ClaimRewards(&_AaveIncentivesControllerV2.TransactOpts, assets, amount, to)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x6d34b96e.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to) returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Transactor) ClaimRewardsOnBehalf(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, user common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.contract.Transact(opts, "claimRewardsOnBehalf", assets, amount, user, to)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x6d34b96e.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to) returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) ClaimRewardsOnBehalf(assets []common.Address, amount *big.Int, user common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.ClaimRewardsOnBehalf(&_AaveIncentivesControllerV2.TransactOpts, assets, amount, user, to)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x6d34b96e.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to) returns(uint256)
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2TransactorSession) ClaimRewardsOnBehalf(assets []common.Address, amount *big.Int, user common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.ClaimRewardsOnBehalf(&_AaveIncentivesControllerV2.TransactOpts, assets, amount, user, to)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x79f171b2.
//
// Solidity: function configureAssets(address[] assets, uint256[] emissionsPerSecond) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Transactor) ConfigureAssets(opts *bind.TransactOpts, assets []common.Address, emissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.contract.Transact(opts, "configureAssets", assets, emissionsPerSecond)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x79f171b2.
//
// Solidity: function configureAssets(address[] assets, uint256[] emissionsPerSecond) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) ConfigureAssets(assets []common.Address, emissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.ConfigureAssets(&_AaveIncentivesControllerV2.TransactOpts, assets, emissionsPerSecond)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x79f171b2.
//
// Solidity: function configureAssets(address[] assets, uint256[] emissionsPerSecond) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2TransactorSession) ConfigureAssets(assets []common.Address, emissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.ConfigureAssets(&_AaveIncentivesControllerV2.TransactOpts, assets, emissionsPerSecond)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Transactor) HandleAction(opts *bind.TransactOpts, user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.contract.Transact(opts, "handleAction", user, totalSupply, userBalance)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) HandleAction(user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.HandleAction(&_AaveIncentivesControllerV2.TransactOpts, user, totalSupply, userBalance)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2TransactorSession) HandleAction(user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.HandleAction(&_AaveIncentivesControllerV2.TransactOpts, user, totalSupply, userBalance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rewardsVault) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Transactor) Initialize(opts *bind.TransactOpts, rewardsVault common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.contract.Transact(opts, "initialize", rewardsVault)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rewardsVault) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) Initialize(rewardsVault common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.Initialize(&_AaveIncentivesControllerV2.TransactOpts, rewardsVault)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rewardsVault) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2TransactorSession) Initialize(rewardsVault common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.Initialize(&_AaveIncentivesControllerV2.TransactOpts, rewardsVault)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Transactor) SetClaimer(opts *bind.TransactOpts, user common.Address, caller common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.contract.Transact(opts, "setClaimer", user, caller)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) SetClaimer(user common.Address, caller common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.SetClaimer(&_AaveIncentivesControllerV2.TransactOpts, user, caller)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2TransactorSession) SetClaimer(user common.Address, caller common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.SetClaimer(&_AaveIncentivesControllerV2.TransactOpts, user, caller)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0x39ccbdd3.
//
// Solidity: function setDistributionEnd(uint256 distributionEnd) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Transactor) SetDistributionEnd(opts *bind.TransactOpts, distributionEnd *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.contract.Transact(opts, "setDistributionEnd", distributionEnd)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0x39ccbdd3.
//
// Solidity: function setDistributionEnd(uint256 distributionEnd) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) SetDistributionEnd(distributionEnd *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.SetDistributionEnd(&_AaveIncentivesControllerV2.TransactOpts, distributionEnd)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0x39ccbdd3.
//
// Solidity: function setDistributionEnd(uint256 distributionEnd) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2TransactorSession) SetDistributionEnd(distributionEnd *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.SetDistributionEnd(&_AaveIncentivesControllerV2.TransactOpts, distributionEnd)
}

// SetRewardsVault is a paid mutator transaction binding the contract method 0xf5bb3e02.
//
// Solidity: function setRewardsVault(address rewardsVault) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Transactor) SetRewardsVault(opts *bind.TransactOpts, rewardsVault common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.contract.Transact(opts, "setRewardsVault", rewardsVault)
}

// SetRewardsVault is a paid mutator transaction binding the contract method 0xf5bb3e02.
//
// Solidity: function setRewardsVault(address rewardsVault) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2Session) SetRewardsVault(rewardsVault common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.SetRewardsVault(&_AaveIncentivesControllerV2.TransactOpts, rewardsVault)
}

// SetRewardsVault is a paid mutator transaction binding the contract method 0xf5bb3e02.
//
// Solidity: function setRewardsVault(address rewardsVault) returns()
func (_AaveIncentivesControllerV2 *AaveIncentivesControllerV2TransactorSession) SetRewardsVault(rewardsVault common.Address) (*types.Transaction, error) {
	return _AaveIncentivesControllerV2.Contract.SetRewardsVault(&_AaveIncentivesControllerV2.TransactOpts, rewardsVault)
}
