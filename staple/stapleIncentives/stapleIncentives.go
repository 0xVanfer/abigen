// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stapleIncentives

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

// IncentivesDataTypesDistribution is an auto generated low-level Go binding around an user-defined struct.
type IncentivesDataTypesDistribution struct {
	Addr            common.Address
	StartTime       *big.Int
	EndTime         *big.Int
	EmissionPerTick *big.Int
}

// IncentivesDataTypesGlobal is an auto generated low-level Go binding around an user-defined struct.
type IncentivesDataTypesGlobal struct {
	LastUpdateTick    *big.Int
	LastClaimEmission *big.Int
}

// IncentivesDataTypesUser is an auto generated low-level Go binding around an user-defined struct.
type IncentivesDataTypesUser struct {
	LastClaimTick     *big.Int
	LastClaimEmission *big.Int
	Unclaimed         *big.Int
}

// StapleIncentivesMetaData contains all meta data concerning the StapleIncentives contract.
var StapleIncentivesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"incentivesID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimIncentives\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"incentivesID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"incentivesToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emissionPerTick\",\"type\":\"uint256\"}],\"name\":\"UpdateIncentives\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"claimFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"estimateIncentivesForUser\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"incentivesIDs_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"incentivesID\",\"type\":\"uint256\"}],\"name\":\"getDistributionInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"emissionPerTick\",\"type\":\"uint256\"}],\"internalType\":\"structIncentivesDataTypes.Distribution\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"incentivesID\",\"type\":\"uint256\"}],\"name\":\"getGlobalIncentivesInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lastUpdateTick\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimEmission\",\"type\":\"uint256\"}],\"internalType\":\"structIncentivesDataTypes.Global\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"getTokenActiveIncentivesIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"getTokenIncentivesIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"incentivesID\",\"type\":\"uint256\"}],\"name\":\"getUserIncentivesInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lastClaimTick\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimEmission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unclaimed\",\"type\":\"uint256\"}],\"internalType\":\"structIncentivesDataTypes.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"refresh\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StapleIncentivesABI is the input ABI used to generate the binding from.
// Deprecated: Use StapleIncentivesMetaData.ABI instead.
var StapleIncentivesABI = StapleIncentivesMetaData.ABI

// StapleIncentives is an auto generated Go binding around an Ethereum contract.
type StapleIncentives struct {
	StapleIncentivesCaller     // Read-only binding to the contract
	StapleIncentivesTransactor // Write-only binding to the contract
	StapleIncentivesFilterer   // Log filterer for contract events
}

// StapleIncentivesCaller is an auto generated read-only Go binding around an Ethereum contract.
type StapleIncentivesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StapleIncentivesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StapleIncentivesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StapleIncentivesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StapleIncentivesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StapleIncentivesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StapleIncentivesSession struct {
	Contract     *StapleIncentives // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StapleIncentivesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StapleIncentivesCallerSession struct {
	Contract *StapleIncentivesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StapleIncentivesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StapleIncentivesTransactorSession struct {
	Contract     *StapleIncentivesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StapleIncentivesRaw is an auto generated low-level Go binding around an Ethereum contract.
type StapleIncentivesRaw struct {
	Contract *StapleIncentives // Generic contract binding to access the raw methods on
}

// StapleIncentivesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StapleIncentivesCallerRaw struct {
	Contract *StapleIncentivesCaller // Generic read-only contract binding to access the raw methods on
}

// StapleIncentivesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StapleIncentivesTransactorRaw struct {
	Contract *StapleIncentivesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStapleIncentives creates a new instance of StapleIncentives, bound to a specific deployed contract.
func NewStapleIncentives(address common.Address, backend bind.ContractBackend) (*StapleIncentives, error) {
	contract, err := bindStapleIncentives(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StapleIncentives{StapleIncentivesCaller: StapleIncentivesCaller{contract: contract}, StapleIncentivesTransactor: StapleIncentivesTransactor{contract: contract}, StapleIncentivesFilterer: StapleIncentivesFilterer{contract: contract}}, nil
}

// NewStapleIncentivesCaller creates a new read-only instance of StapleIncentives, bound to a specific deployed contract.
func NewStapleIncentivesCaller(address common.Address, caller bind.ContractCaller) (*StapleIncentivesCaller, error) {
	contract, err := bindStapleIncentives(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StapleIncentivesCaller{contract: contract}, nil
}

// NewStapleIncentivesTransactor creates a new write-only instance of StapleIncentives, bound to a specific deployed contract.
func NewStapleIncentivesTransactor(address common.Address, transactor bind.ContractTransactor) (*StapleIncentivesTransactor, error) {
	contract, err := bindStapleIncentives(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StapleIncentivesTransactor{contract: contract}, nil
}

// NewStapleIncentivesFilterer creates a new log filterer instance of StapleIncentives, bound to a specific deployed contract.
func NewStapleIncentivesFilterer(address common.Address, filterer bind.ContractFilterer) (*StapleIncentivesFilterer, error) {
	contract, err := bindStapleIncentives(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StapleIncentivesFilterer{contract: contract}, nil
}

// bindStapleIncentives binds a generic wrapper to an already deployed contract.
func bindStapleIncentives(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StapleIncentivesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StapleIncentives *StapleIncentivesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StapleIncentives.Contract.StapleIncentivesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StapleIncentives *StapleIncentivesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StapleIncentives.Contract.StapleIncentivesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StapleIncentives *StapleIncentivesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StapleIncentives.Contract.StapleIncentivesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StapleIncentives *StapleIncentivesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StapleIncentives.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StapleIncentives *StapleIncentivesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StapleIncentives.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StapleIncentives *StapleIncentivesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StapleIncentives.Contract.contract.Transact(opts, method, params...)
}

// EstimateIncentivesForUser is a free data retrieval call binding the contract method 0xb803c28c.
//
// Solidity: function estimateIncentivesForUser(address user, address underlying, uint256 vtpID) view returns(uint256[] incentivesIDs_, uint256[] amounts_)
func (_StapleIncentives *StapleIncentivesCaller) EstimateIncentivesForUser(opts *bind.CallOpts, user common.Address, underlying common.Address, vtpID *big.Int) (struct {
	IncentivesIDs []*big.Int
	Amounts       []*big.Int
}, error) {
	var out []interface{}
	err := _StapleIncentives.contract.Call(opts, &out, "estimateIncentivesForUser", user, underlying, vtpID)

	outstruct := new(struct {
		IncentivesIDs []*big.Int
		Amounts       []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IncentivesIDs = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// EstimateIncentivesForUser is a free data retrieval call binding the contract method 0xb803c28c.
//
// Solidity: function estimateIncentivesForUser(address user, address underlying, uint256 vtpID) view returns(uint256[] incentivesIDs_, uint256[] amounts_)
func (_StapleIncentives *StapleIncentivesSession) EstimateIncentivesForUser(user common.Address, underlying common.Address, vtpID *big.Int) (struct {
	IncentivesIDs []*big.Int
	Amounts       []*big.Int
}, error) {
	return _StapleIncentives.Contract.EstimateIncentivesForUser(&_StapleIncentives.CallOpts, user, underlying, vtpID)
}

// EstimateIncentivesForUser is a free data retrieval call binding the contract method 0xb803c28c.
//
// Solidity: function estimateIncentivesForUser(address user, address underlying, uint256 vtpID) view returns(uint256[] incentivesIDs_, uint256[] amounts_)
func (_StapleIncentives *StapleIncentivesCallerSession) EstimateIncentivesForUser(user common.Address, underlying common.Address, vtpID *big.Int) (struct {
	IncentivesIDs []*big.Int
	Amounts       []*big.Int
}, error) {
	return _StapleIncentives.Contract.EstimateIncentivesForUser(&_StapleIncentives.CallOpts, user, underlying, vtpID)
}

// GetDistributionInfo is a free data retrieval call binding the contract method 0xd6ec965d.
//
// Solidity: function getDistributionInfo(address underlying, uint256 vtpID, uint256 incentivesID) view returns((address,uint256,uint256,uint256))
func (_StapleIncentives *StapleIncentivesCaller) GetDistributionInfo(opts *bind.CallOpts, underlying common.Address, vtpID *big.Int, incentivesID *big.Int) (IncentivesDataTypesDistribution, error) {
	var out []interface{}
	err := _StapleIncentives.contract.Call(opts, &out, "getDistributionInfo", underlying, vtpID, incentivesID)

	if err != nil {
		return *new(IncentivesDataTypesDistribution), err
	}

	out0 := *abi.ConvertType(out[0], new(IncentivesDataTypesDistribution)).(*IncentivesDataTypesDistribution)

	return out0, err

}

// GetDistributionInfo is a free data retrieval call binding the contract method 0xd6ec965d.
//
// Solidity: function getDistributionInfo(address underlying, uint256 vtpID, uint256 incentivesID) view returns((address,uint256,uint256,uint256))
func (_StapleIncentives *StapleIncentivesSession) GetDistributionInfo(underlying common.Address, vtpID *big.Int, incentivesID *big.Int) (IncentivesDataTypesDistribution, error) {
	return _StapleIncentives.Contract.GetDistributionInfo(&_StapleIncentives.CallOpts, underlying, vtpID, incentivesID)
}

// GetDistributionInfo is a free data retrieval call binding the contract method 0xd6ec965d.
//
// Solidity: function getDistributionInfo(address underlying, uint256 vtpID, uint256 incentivesID) view returns((address,uint256,uint256,uint256))
func (_StapleIncentives *StapleIncentivesCallerSession) GetDistributionInfo(underlying common.Address, vtpID *big.Int, incentivesID *big.Int) (IncentivesDataTypesDistribution, error) {
	return _StapleIncentives.Contract.GetDistributionInfo(&_StapleIncentives.CallOpts, underlying, vtpID, incentivesID)
}

// GetGlobalIncentivesInfo is a free data retrieval call binding the contract method 0x4c97aba0.
//
// Solidity: function getGlobalIncentivesInfo(address underlying, uint256 vtpID, uint256 incentivesID) view returns((uint256,uint256))
func (_StapleIncentives *StapleIncentivesCaller) GetGlobalIncentivesInfo(opts *bind.CallOpts, underlying common.Address, vtpID *big.Int, incentivesID *big.Int) (IncentivesDataTypesGlobal, error) {
	var out []interface{}
	err := _StapleIncentives.contract.Call(opts, &out, "getGlobalIncentivesInfo", underlying, vtpID, incentivesID)

	if err != nil {
		return *new(IncentivesDataTypesGlobal), err
	}

	out0 := *abi.ConvertType(out[0], new(IncentivesDataTypesGlobal)).(*IncentivesDataTypesGlobal)

	return out0, err

}

// GetGlobalIncentivesInfo is a free data retrieval call binding the contract method 0x4c97aba0.
//
// Solidity: function getGlobalIncentivesInfo(address underlying, uint256 vtpID, uint256 incentivesID) view returns((uint256,uint256))
func (_StapleIncentives *StapleIncentivesSession) GetGlobalIncentivesInfo(underlying common.Address, vtpID *big.Int, incentivesID *big.Int) (IncentivesDataTypesGlobal, error) {
	return _StapleIncentives.Contract.GetGlobalIncentivesInfo(&_StapleIncentives.CallOpts, underlying, vtpID, incentivesID)
}

// GetGlobalIncentivesInfo is a free data retrieval call binding the contract method 0x4c97aba0.
//
// Solidity: function getGlobalIncentivesInfo(address underlying, uint256 vtpID, uint256 incentivesID) view returns((uint256,uint256))
func (_StapleIncentives *StapleIncentivesCallerSession) GetGlobalIncentivesInfo(underlying common.Address, vtpID *big.Int, incentivesID *big.Int) (IncentivesDataTypesGlobal, error) {
	return _StapleIncentives.Contract.GetGlobalIncentivesInfo(&_StapleIncentives.CallOpts, underlying, vtpID, incentivesID)
}

// GetTokenActiveIncentivesIDs is a free data retrieval call binding the contract method 0xf7d5be56.
//
// Solidity: function getTokenActiveIncentivesIDs(address underlying, uint256 vtpID) view returns(uint256[])
func (_StapleIncentives *StapleIncentivesCaller) GetTokenActiveIncentivesIDs(opts *bind.CallOpts, underlying common.Address, vtpID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _StapleIncentives.contract.Call(opts, &out, "getTokenActiveIncentivesIDs", underlying, vtpID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenActiveIncentivesIDs is a free data retrieval call binding the contract method 0xf7d5be56.
//
// Solidity: function getTokenActiveIncentivesIDs(address underlying, uint256 vtpID) view returns(uint256[])
func (_StapleIncentives *StapleIncentivesSession) GetTokenActiveIncentivesIDs(underlying common.Address, vtpID *big.Int) ([]*big.Int, error) {
	return _StapleIncentives.Contract.GetTokenActiveIncentivesIDs(&_StapleIncentives.CallOpts, underlying, vtpID)
}

// GetTokenActiveIncentivesIDs is a free data retrieval call binding the contract method 0xf7d5be56.
//
// Solidity: function getTokenActiveIncentivesIDs(address underlying, uint256 vtpID) view returns(uint256[])
func (_StapleIncentives *StapleIncentivesCallerSession) GetTokenActiveIncentivesIDs(underlying common.Address, vtpID *big.Int) ([]*big.Int, error) {
	return _StapleIncentives.Contract.GetTokenActiveIncentivesIDs(&_StapleIncentives.CallOpts, underlying, vtpID)
}

// GetTokenIncentivesIDs is a free data retrieval call binding the contract method 0x8b06665b.
//
// Solidity: function getTokenIncentivesIDs(address underlying, uint256 vtpID) view returns(uint256[])
func (_StapleIncentives *StapleIncentivesCaller) GetTokenIncentivesIDs(opts *bind.CallOpts, underlying common.Address, vtpID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _StapleIncentives.contract.Call(opts, &out, "getTokenIncentivesIDs", underlying, vtpID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenIncentivesIDs is a free data retrieval call binding the contract method 0x8b06665b.
//
// Solidity: function getTokenIncentivesIDs(address underlying, uint256 vtpID) view returns(uint256[])
func (_StapleIncentives *StapleIncentivesSession) GetTokenIncentivesIDs(underlying common.Address, vtpID *big.Int) ([]*big.Int, error) {
	return _StapleIncentives.Contract.GetTokenIncentivesIDs(&_StapleIncentives.CallOpts, underlying, vtpID)
}

// GetTokenIncentivesIDs is a free data retrieval call binding the contract method 0x8b06665b.
//
// Solidity: function getTokenIncentivesIDs(address underlying, uint256 vtpID) view returns(uint256[])
func (_StapleIncentives *StapleIncentivesCallerSession) GetTokenIncentivesIDs(underlying common.Address, vtpID *big.Int) ([]*big.Int, error) {
	return _StapleIncentives.Contract.GetTokenIncentivesIDs(&_StapleIncentives.CallOpts, underlying, vtpID)
}

// GetUserIncentivesInfo is a free data retrieval call binding the contract method 0x6683ea12.
//
// Solidity: function getUserIncentivesInfo(address user, address underlying, uint256 vtpID, uint256 incentivesID) view returns((uint256,uint256,uint256))
func (_StapleIncentives *StapleIncentivesCaller) GetUserIncentivesInfo(opts *bind.CallOpts, user common.Address, underlying common.Address, vtpID *big.Int, incentivesID *big.Int) (IncentivesDataTypesUser, error) {
	var out []interface{}
	err := _StapleIncentives.contract.Call(opts, &out, "getUserIncentivesInfo", user, underlying, vtpID, incentivesID)

	if err != nil {
		return *new(IncentivesDataTypesUser), err
	}

	out0 := *abi.ConvertType(out[0], new(IncentivesDataTypesUser)).(*IncentivesDataTypesUser)

	return out0, err

}

// GetUserIncentivesInfo is a free data retrieval call binding the contract method 0x6683ea12.
//
// Solidity: function getUserIncentivesInfo(address user, address underlying, uint256 vtpID, uint256 incentivesID) view returns((uint256,uint256,uint256))
func (_StapleIncentives *StapleIncentivesSession) GetUserIncentivesInfo(user common.Address, underlying common.Address, vtpID *big.Int, incentivesID *big.Int) (IncentivesDataTypesUser, error) {
	return _StapleIncentives.Contract.GetUserIncentivesInfo(&_StapleIncentives.CallOpts, user, underlying, vtpID, incentivesID)
}

// GetUserIncentivesInfo is a free data retrieval call binding the contract method 0x6683ea12.
//
// Solidity: function getUserIncentivesInfo(address user, address underlying, uint256 vtpID, uint256 incentivesID) view returns((uint256,uint256,uint256))
func (_StapleIncentives *StapleIncentivesCallerSession) GetUserIncentivesInfo(user common.Address, underlying common.Address, vtpID *big.Int, incentivesID *big.Int) (IncentivesDataTypesUser, error) {
	return _StapleIncentives.Contract.GetUserIncentivesInfo(&_StapleIncentives.CallOpts, user, underlying, vtpID, incentivesID)
}

// ClaimFor is a paid mutator transaction binding the contract method 0x44dae040.
//
// Solidity: function claimFor(address _for, address underlying, uint256 vtpID) returns()
func (_StapleIncentives *StapleIncentivesTransactor) ClaimFor(opts *bind.TransactOpts, _for common.Address, underlying common.Address, vtpID *big.Int) (*types.Transaction, error) {
	return _StapleIncentives.contract.Transact(opts, "claimFor", _for, underlying, vtpID)
}

// ClaimFor is a paid mutator transaction binding the contract method 0x44dae040.
//
// Solidity: function claimFor(address _for, address underlying, uint256 vtpID) returns()
func (_StapleIncentives *StapleIncentivesSession) ClaimFor(_for common.Address, underlying common.Address, vtpID *big.Int) (*types.Transaction, error) {
	return _StapleIncentives.Contract.ClaimFor(&_StapleIncentives.TransactOpts, _for, underlying, vtpID)
}

// ClaimFor is a paid mutator transaction binding the contract method 0x44dae040.
//
// Solidity: function claimFor(address _for, address underlying, uint256 vtpID) returns()
func (_StapleIncentives *StapleIncentivesTransactorSession) ClaimFor(_for common.Address, underlying common.Address, vtpID *big.Int) (*types.Transaction, error) {
	return _StapleIncentives.Contract.ClaimFor(&_StapleIncentives.TransactOpts, _for, underlying, vtpID)
}

// Refresh is a paid mutator transaction binding the contract method 0xf25eca49.
//
// Solidity: function refresh(address user, address underlying, uint256 vtpID) returns()
func (_StapleIncentives *StapleIncentivesTransactor) Refresh(opts *bind.TransactOpts, user common.Address, underlying common.Address, vtpID *big.Int) (*types.Transaction, error) {
	return _StapleIncentives.contract.Transact(opts, "refresh", user, underlying, vtpID)
}

// Refresh is a paid mutator transaction binding the contract method 0xf25eca49.
//
// Solidity: function refresh(address user, address underlying, uint256 vtpID) returns()
func (_StapleIncentives *StapleIncentivesSession) Refresh(user common.Address, underlying common.Address, vtpID *big.Int) (*types.Transaction, error) {
	return _StapleIncentives.Contract.Refresh(&_StapleIncentives.TransactOpts, user, underlying, vtpID)
}

// Refresh is a paid mutator transaction binding the contract method 0xf25eca49.
//
// Solidity: function refresh(address user, address underlying, uint256 vtpID) returns()
func (_StapleIncentives *StapleIncentivesTransactorSession) Refresh(user common.Address, underlying common.Address, vtpID *big.Int) (*types.Transaction, error) {
	return _StapleIncentives.Contract.Refresh(&_StapleIncentives.TransactOpts, user, underlying, vtpID)
}

// StapleIncentivesClaimIncentivesIterator is returned from FilterClaimIncentives and is used to iterate over the raw logs and unpacked data for ClaimIncentives events raised by the StapleIncentives contract.
type StapleIncentivesClaimIncentivesIterator struct {
	Event *StapleIncentivesClaimIncentives // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StapleIncentivesClaimIncentivesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleIncentivesClaimIncentives)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StapleIncentivesClaimIncentives)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StapleIncentivesClaimIncentivesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleIncentivesClaimIncentivesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleIncentivesClaimIncentives represents a ClaimIncentives event raised by the StapleIncentives contract.
type StapleIncentivesClaimIncentives struct {
	Owner        common.Address
	Receiver     common.Address
	Underlying   common.Address
	VtpID        *big.Int
	IncentivesID *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaimIncentives is a free log retrieval operation binding the contract event 0x7b2cdf9bab072af59077ca3041e3e90dac022f02c42fc1b9573273f757402fb3.
//
// Solidity: event ClaimIncentives(address indexed owner, address indexed receiver, address underlying, uint256 vtpID, uint256 incentivesID, uint256 amount)
func (_StapleIncentives *StapleIncentivesFilterer) FilterClaimIncentives(opts *bind.FilterOpts, owner []common.Address, receiver []common.Address) (*StapleIncentivesClaimIncentivesIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _StapleIncentives.contract.FilterLogs(opts, "ClaimIncentives", ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &StapleIncentivesClaimIncentivesIterator{contract: _StapleIncentives.contract, event: "ClaimIncentives", logs: logs, sub: sub}, nil
}

// WatchClaimIncentives is a free log subscription operation binding the contract event 0x7b2cdf9bab072af59077ca3041e3e90dac022f02c42fc1b9573273f757402fb3.
//
// Solidity: event ClaimIncentives(address indexed owner, address indexed receiver, address underlying, uint256 vtpID, uint256 incentivesID, uint256 amount)
func (_StapleIncentives *StapleIncentivesFilterer) WatchClaimIncentives(opts *bind.WatchOpts, sink chan<- *StapleIncentivesClaimIncentives, owner []common.Address, receiver []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _StapleIncentives.contract.WatchLogs(opts, "ClaimIncentives", ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleIncentivesClaimIncentives)
				if err := _StapleIncentives.contract.UnpackLog(event, "ClaimIncentives", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimIncentives is a log parse operation binding the contract event 0x7b2cdf9bab072af59077ca3041e3e90dac022f02c42fc1b9573273f757402fb3.
//
// Solidity: event ClaimIncentives(address indexed owner, address indexed receiver, address underlying, uint256 vtpID, uint256 incentivesID, uint256 amount)
func (_StapleIncentives *StapleIncentivesFilterer) ParseClaimIncentives(log types.Log) (*StapleIncentivesClaimIncentives, error) {
	event := new(StapleIncentivesClaimIncentives)
	if err := _StapleIncentives.contract.UnpackLog(event, "ClaimIncentives", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleIncentivesUpdateControllerIterator is returned from FilterUpdateController and is used to iterate over the raw logs and unpacked data for UpdateController events raised by the StapleIncentives contract.
type StapleIncentivesUpdateControllerIterator struct {
	Event *StapleIncentivesUpdateController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StapleIncentivesUpdateControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleIncentivesUpdateController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StapleIncentivesUpdateController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StapleIncentivesUpdateControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleIncentivesUpdateControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleIncentivesUpdateController represents a UpdateController event raised by the StapleIncentives contract.
type StapleIncentivesUpdateController struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateController is a free log retrieval operation binding the contract event 0xfce593255b7647c9b730cf207aed89b0761a9c7ff374009a869068dffa08fb2a.
//
// Solidity: event UpdateController(address newAddress)
func (_StapleIncentives *StapleIncentivesFilterer) FilterUpdateController(opts *bind.FilterOpts) (*StapleIncentivesUpdateControllerIterator, error) {

	logs, sub, err := _StapleIncentives.contract.FilterLogs(opts, "UpdateController")
	if err != nil {
		return nil, err
	}
	return &StapleIncentivesUpdateControllerIterator{contract: _StapleIncentives.contract, event: "UpdateController", logs: logs, sub: sub}, nil
}

// WatchUpdateController is a free log subscription operation binding the contract event 0xfce593255b7647c9b730cf207aed89b0761a9c7ff374009a869068dffa08fb2a.
//
// Solidity: event UpdateController(address newAddress)
func (_StapleIncentives *StapleIncentivesFilterer) WatchUpdateController(opts *bind.WatchOpts, sink chan<- *StapleIncentivesUpdateController) (event.Subscription, error) {

	logs, sub, err := _StapleIncentives.contract.WatchLogs(opts, "UpdateController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleIncentivesUpdateController)
				if err := _StapleIncentives.contract.UnpackLog(event, "UpdateController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateController is a log parse operation binding the contract event 0xfce593255b7647c9b730cf207aed89b0761a9c7ff374009a869068dffa08fb2a.
//
// Solidity: event UpdateController(address newAddress)
func (_StapleIncentives *StapleIncentivesFilterer) ParseUpdateController(log types.Log) (*StapleIncentivesUpdateController, error) {
	event := new(StapleIncentivesUpdateController)
	if err := _StapleIncentives.contract.UnpackLog(event, "UpdateController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleIncentivesUpdateIncentivesIterator is returned from FilterUpdateIncentives and is used to iterate over the raw logs and unpacked data for UpdateIncentives events raised by the StapleIncentives contract.
type StapleIncentivesUpdateIncentivesIterator struct {
	Event *StapleIncentivesUpdateIncentives // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StapleIncentivesUpdateIncentivesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleIncentivesUpdateIncentives)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StapleIncentivesUpdateIncentives)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StapleIncentivesUpdateIncentivesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleIncentivesUpdateIncentivesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleIncentivesUpdateIncentives represents a UpdateIncentives event raised by the StapleIncentives contract.
type StapleIncentivesUpdateIncentives struct {
	Underlying      common.Address
	VtpID           *big.Int
	IncentivesID    *big.Int
	IncentivesToken common.Address
	StartTime       *big.Int
	EndTime         *big.Int
	EmissionPerTick *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateIncentives is a free log retrieval operation binding the contract event 0xda743ac955b3eb50ca52b072cc175d8975250179f5e498f5876f665af710c1d0.
//
// Solidity: event UpdateIncentives(address indexed underlying, uint256 indexed vtpID, uint256 indexed incentivesID, address incentivesToken, uint256 startTime, uint256 endTime, uint256 emissionPerTick)
func (_StapleIncentives *StapleIncentivesFilterer) FilterUpdateIncentives(opts *bind.FilterOpts, underlying []common.Address, vtpID []*big.Int, incentivesID []*big.Int) (*StapleIncentivesUpdateIncentivesIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}
	var incentivesIDRule []interface{}
	for _, incentivesIDItem := range incentivesID {
		incentivesIDRule = append(incentivesIDRule, incentivesIDItem)
	}

	logs, sub, err := _StapleIncentives.contract.FilterLogs(opts, "UpdateIncentives", underlyingRule, vtpIDRule, incentivesIDRule)
	if err != nil {
		return nil, err
	}
	return &StapleIncentivesUpdateIncentivesIterator{contract: _StapleIncentives.contract, event: "UpdateIncentives", logs: logs, sub: sub}, nil
}

// WatchUpdateIncentives is a free log subscription operation binding the contract event 0xda743ac955b3eb50ca52b072cc175d8975250179f5e498f5876f665af710c1d0.
//
// Solidity: event UpdateIncentives(address indexed underlying, uint256 indexed vtpID, uint256 indexed incentivesID, address incentivesToken, uint256 startTime, uint256 endTime, uint256 emissionPerTick)
func (_StapleIncentives *StapleIncentivesFilterer) WatchUpdateIncentives(opts *bind.WatchOpts, sink chan<- *StapleIncentivesUpdateIncentives, underlying []common.Address, vtpID []*big.Int, incentivesID []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}
	var incentivesIDRule []interface{}
	for _, incentivesIDItem := range incentivesID {
		incentivesIDRule = append(incentivesIDRule, incentivesIDItem)
	}

	logs, sub, err := _StapleIncentives.contract.WatchLogs(opts, "UpdateIncentives", underlyingRule, vtpIDRule, incentivesIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleIncentivesUpdateIncentives)
				if err := _StapleIncentives.contract.UnpackLog(event, "UpdateIncentives", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateIncentives is a log parse operation binding the contract event 0xda743ac955b3eb50ca52b072cc175d8975250179f5e498f5876f665af710c1d0.
//
// Solidity: event UpdateIncentives(address indexed underlying, uint256 indexed vtpID, uint256 indexed incentivesID, address incentivesToken, uint256 startTime, uint256 endTime, uint256 emissionPerTick)
func (_StapleIncentives *StapleIncentivesFilterer) ParseUpdateIncentives(log types.Log) (*StapleIncentivesUpdateIncentives, error) {
	event := new(StapleIncentivesUpdateIncentives)
	if err := _StapleIncentives.contract.UnpackLog(event, "UpdateIncentives", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
