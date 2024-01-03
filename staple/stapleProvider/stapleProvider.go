// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stapleProvider

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

// PoolDataTypesPoolParams is an auto generated low-level Go binding around an user-defined struct.
type PoolDataTypesPoolParams struct {
	Id                 *big.Int
	Underlying         common.Address
	LpAddr             common.Address
	UnderlyingDecimals uint8
	Tl                 uint8
	RelatedVtps        []*big.Int
}

// UIDataTypesPool is an auto generated low-level Go binding around an user-defined struct.
type UIDataTypesPool struct {
	Params PoolDataTypesPoolParams
	Status UIDataTypesPoolStatus
}

// UIDataTypesPoolStatus is an auto generated low-level Go binding around an user-defined struct.
type UIDataTypesPoolStatus struct {
	Asset     *big.Int
	Liability *big.Int
	Invested  *big.Int
}

// UIDataTypesUserIncentivesInfo is an auto generated low-level Go binding around an user-defined struct.
type UIDataTypesUserIncentivesInfo struct {
	Token           common.Address
	Symbol          string
	Decimals        *big.Int
	Amount          *big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	EmissionPerTick *big.Int
}

// UIDataTypesUserPoolInfo is an auto generated low-level Go binding around an user-defined struct.
type UIDataTypesUserPoolInfo struct {
	Underlying         common.Address
	Deposition         *big.Int
	Liability          *big.Int
	UnlockedLp         *big.Int
	UnlockedUnderlying *big.Int
	TotalCredits       *big.Int
	CreditUsed         *big.Int
	AllocatedVtps      []*big.Int
	UserVtps           []UIDataTypesUserVtpTokenInfo
	AllRewards         *big.Int
}

// UIDataTypesUserVtpTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type UIDataTypesUserVtpTokenInfo struct {
	Id                *big.Int
	Alloction         *big.Int
	AlloctedRate      *big.Int
	MaxDeallocateRate *big.Int
	Shares            *big.Int
	Rewards           *big.Int
	Incentives        []UIDataTypesUserIncentivesInfo
}

// StapleProviderMetaData contains all meta data concerning the StapleProvider contract.
var StapleProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIControllerEx\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPools\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lpAddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"underlyingDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tl\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"relatedVtps\",\"type\":\"uint256[]\"}],\"internalType\":\"structPoolDataTypes.PoolParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liability\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invested\",\"type\":\"uint256\"}],\"internalType\":\"structUIDataTypes.PoolStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"internalType\":\"structUIDataTypes.Pool[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllUnderlyings\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"getDz\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getDzWithoutPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"}],\"name\":\"getPool\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lpAddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"underlyingDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tl\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"relatedVtps\",\"type\":\"uint256[]\"}],\"internalType\":\"structPoolDataTypes.PoolParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liability\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invested\",\"type\":\"uint256\"}],\"internalType\":\"structUIDataTypes.PoolStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"internalType\":\"structUIDataTypes.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"idList\",\"type\":\"uint256[]\"}],\"name\":\"getPools\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lpAddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"underlyingDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tl\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"relatedVtps\",\"type\":\"uint256[]\"}],\"internalType\":\"structPoolDataTypes.PoolParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liability\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invested\",\"type\":\"uint256\"}],\"internalType\":\"structUIDataTypes.PoolStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"internalType\":\"structUIDataTypes.Pool[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"}],\"name\":\"getTokensByPath\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incentivesController\",\"outputs\":[{\"internalType\":\"contractIIncentivesController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposition\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liability\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedLp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedUnderlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCredits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"allocatedVtps\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alloction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alloctedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDeallocateRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"emissionPerTick\",\"type\":\"uint256\"}],\"internalType\":\"structUIDataTypes.UserIncentivesInfo[]\",\"name\":\"incentives\",\"type\":\"tuple[]\"}],\"internalType\":\"structUIDataTypes.UserVtpTokenInfo[]\",\"name\":\"userVtps\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"allRewards\",\"type\":\"uint256\"}],\"internalType\":\"structUIDataTypes.UserPoolInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"userPoolInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposition\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liability\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedLp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedUnderlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCredits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"allocatedVtps\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alloction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alloctedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDeallocateRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"emissionPerTick\",\"type\":\"uint256\"}],\"internalType\":\"structUIDataTypes.UserIncentivesInfo[]\",\"name\":\"incentives\",\"type\":\"tuple[]\"}],\"internalType\":\"structUIDataTypes.UserVtpTokenInfo[]\",\"name\":\"userVtps\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"allRewards\",\"type\":\"uint256\"}],\"internalType\":\"structUIDataTypes.UserPoolInfo\",\"name\":\"userPoolInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"userVtpTokenIncentives\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"emissionPerTick\",\"type\":\"uint256\"}],\"internalType\":\"structUIDataTypes.UserIncentivesInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StapleProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use StapleProviderMetaData.ABI instead.
var StapleProviderABI = StapleProviderMetaData.ABI

// StapleProvider is an auto generated Go binding around an Ethereum contract.
type StapleProvider struct {
	StapleProviderCaller     // Read-only binding to the contract
	StapleProviderTransactor // Write-only binding to the contract
	StapleProviderFilterer   // Log filterer for contract events
}

// StapleProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type StapleProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StapleProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StapleProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StapleProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StapleProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StapleProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StapleProviderSession struct {
	Contract     *StapleProvider   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StapleProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StapleProviderCallerSession struct {
	Contract *StapleProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StapleProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StapleProviderTransactorSession struct {
	Contract     *StapleProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StapleProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type StapleProviderRaw struct {
	Contract *StapleProvider // Generic contract binding to access the raw methods on
}

// StapleProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StapleProviderCallerRaw struct {
	Contract *StapleProviderCaller // Generic read-only contract binding to access the raw methods on
}

// StapleProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StapleProviderTransactorRaw struct {
	Contract *StapleProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStapleProvider creates a new instance of StapleProvider, bound to a specific deployed contract.
func NewStapleProvider(address common.Address, backend bind.ContractBackend) (*StapleProvider, error) {
	contract, err := bindStapleProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StapleProvider{StapleProviderCaller: StapleProviderCaller{contract: contract}, StapleProviderTransactor: StapleProviderTransactor{contract: contract}, StapleProviderFilterer: StapleProviderFilterer{contract: contract}}, nil
}

// NewStapleProviderCaller creates a new read-only instance of StapleProvider, bound to a specific deployed contract.
func NewStapleProviderCaller(address common.Address, caller bind.ContractCaller) (*StapleProviderCaller, error) {
	contract, err := bindStapleProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StapleProviderCaller{contract: contract}, nil
}

// NewStapleProviderTransactor creates a new write-only instance of StapleProvider, bound to a specific deployed contract.
func NewStapleProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*StapleProviderTransactor, error) {
	contract, err := bindStapleProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StapleProviderTransactor{contract: contract}, nil
}

// NewStapleProviderFilterer creates a new log filterer instance of StapleProvider, bound to a specific deployed contract.
func NewStapleProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*StapleProviderFilterer, error) {
	contract, err := bindStapleProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StapleProviderFilterer{contract: contract}, nil
}

// bindStapleProvider binds a generic wrapper to an already deployed contract.
func bindStapleProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StapleProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StapleProvider *StapleProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StapleProvider.Contract.StapleProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StapleProvider *StapleProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StapleProvider.Contract.StapleProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StapleProvider *StapleProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StapleProvider.Contract.StapleProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StapleProvider *StapleProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StapleProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StapleProvider *StapleProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StapleProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StapleProvider *StapleProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StapleProvider.Contract.contract.Transact(opts, method, params...)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_StapleProvider *StapleProviderCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_StapleProvider *StapleProviderSession) Controller() (common.Address, error) {
	return _StapleProvider.Contract.Controller(&_StapleProvider.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_StapleProvider *StapleProviderCallerSession) Controller() (common.Address, error) {
	return _StapleProvider.Contract.Controller(&_StapleProvider.CallOpts)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() view returns(((uint256,address,address,uint8,uint8,uint256[]),(uint256,uint256,uint256))[])
func (_StapleProvider *StapleProviderCaller) GetAllPools(opts *bind.CallOpts) ([]UIDataTypesPool, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "getAllPools")

	if err != nil {
		return *new([]UIDataTypesPool), err
	}

	out0 := *abi.ConvertType(out[0], new([]UIDataTypesPool)).(*[]UIDataTypesPool)

	return out0, err

}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() view returns(((uint256,address,address,uint8,uint8,uint256[]),(uint256,uint256,uint256))[])
func (_StapleProvider *StapleProviderSession) GetAllPools() ([]UIDataTypesPool, error) {
	return _StapleProvider.Contract.GetAllPools(&_StapleProvider.CallOpts)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() view returns(((uint256,address,address,uint8,uint8,uint256[]),(uint256,uint256,uint256))[])
func (_StapleProvider *StapleProviderCallerSession) GetAllPools() ([]UIDataTypesPool, error) {
	return _StapleProvider.Contract.GetAllPools(&_StapleProvider.CallOpts)
}

// GetAllUnderlyings is a free data retrieval call binding the contract method 0xc3f80bc4.
//
// Solidity: function getAllUnderlyings() view returns(address[])
func (_StapleProvider *StapleProviderCaller) GetAllUnderlyings(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "getAllUnderlyings")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllUnderlyings is a free data retrieval call binding the contract method 0xc3f80bc4.
//
// Solidity: function getAllUnderlyings() view returns(address[])
func (_StapleProvider *StapleProviderSession) GetAllUnderlyings() ([]common.Address, error) {
	return _StapleProvider.Contract.GetAllUnderlyings(&_StapleProvider.CallOpts)
}

// GetAllUnderlyings is a free data retrieval call binding the contract method 0xc3f80bc4.
//
// Solidity: function getAllUnderlyings() view returns(address[])
func (_StapleProvider *StapleProviderCallerSession) GetAllUnderlyings() ([]common.Address, error) {
	return _StapleProvider.Contract.GetAllUnderlyings(&_StapleProvider.CallOpts)
}

// GetDz is a free data retrieval call binding the contract method 0x579eda78.
//
// Solidity: function getDz(address user, address tokenIn, uint256[] path, uint256 amount, uint256[] prices) view returns(uint256)
func (_StapleProvider *StapleProviderCaller) GetDz(opts *bind.CallOpts, user common.Address, tokenIn common.Address, path []*big.Int, amount *big.Int, prices []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "getDz", user, tokenIn, path, amount, prices)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDz is a free data retrieval call binding the contract method 0x579eda78.
//
// Solidity: function getDz(address user, address tokenIn, uint256[] path, uint256 amount, uint256[] prices) view returns(uint256)
func (_StapleProvider *StapleProviderSession) GetDz(user common.Address, tokenIn common.Address, path []*big.Int, amount *big.Int, prices []*big.Int) (*big.Int, error) {
	return _StapleProvider.Contract.GetDz(&_StapleProvider.CallOpts, user, tokenIn, path, amount, prices)
}

// GetDz is a free data retrieval call binding the contract method 0x579eda78.
//
// Solidity: function getDz(address user, address tokenIn, uint256[] path, uint256 amount, uint256[] prices) view returns(uint256)
func (_StapleProvider *StapleProviderCallerSession) GetDz(user common.Address, tokenIn common.Address, path []*big.Int, amount *big.Int, prices []*big.Int) (*big.Int, error) {
	return _StapleProvider.Contract.GetDz(&_StapleProvider.CallOpts, user, tokenIn, path, amount, prices)
}

// GetDzWithoutPrices is a free data retrieval call binding the contract method 0x5e7d1a9f.
//
// Solidity: function getDzWithoutPrices(address user, address tokenIn, uint256[] path, uint256 amount) view returns(uint256)
func (_StapleProvider *StapleProviderCaller) GetDzWithoutPrices(opts *bind.CallOpts, user common.Address, tokenIn common.Address, path []*big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "getDzWithoutPrices", user, tokenIn, path, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDzWithoutPrices is a free data retrieval call binding the contract method 0x5e7d1a9f.
//
// Solidity: function getDzWithoutPrices(address user, address tokenIn, uint256[] path, uint256 amount) view returns(uint256)
func (_StapleProvider *StapleProviderSession) GetDzWithoutPrices(user common.Address, tokenIn common.Address, path []*big.Int, amount *big.Int) (*big.Int, error) {
	return _StapleProvider.Contract.GetDzWithoutPrices(&_StapleProvider.CallOpts, user, tokenIn, path, amount)
}

// GetDzWithoutPrices is a free data retrieval call binding the contract method 0x5e7d1a9f.
//
// Solidity: function getDzWithoutPrices(address user, address tokenIn, uint256[] path, uint256 amount) view returns(uint256)
func (_StapleProvider *StapleProviderCallerSession) GetDzWithoutPrices(user common.Address, tokenIn common.Address, path []*big.Int, amount *big.Int) (*big.Int, error) {
	return _StapleProvider.Contract.GetDzWithoutPrices(&_StapleProvider.CallOpts, user, tokenIn, path, amount)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 poolID) view returns(((uint256,address,address,uint8,uint8,uint256[]),(uint256,uint256,uint256)))
func (_StapleProvider *StapleProviderCaller) GetPool(opts *bind.CallOpts, poolID *big.Int) (UIDataTypesPool, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "getPool", poolID)

	if err != nil {
		return *new(UIDataTypesPool), err
	}

	out0 := *abi.ConvertType(out[0], new(UIDataTypesPool)).(*UIDataTypesPool)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 poolID) view returns(((uint256,address,address,uint8,uint8,uint256[]),(uint256,uint256,uint256)))
func (_StapleProvider *StapleProviderSession) GetPool(poolID *big.Int) (UIDataTypesPool, error) {
	return _StapleProvider.Contract.GetPool(&_StapleProvider.CallOpts, poolID)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 poolID) view returns(((uint256,address,address,uint8,uint8,uint256[]),(uint256,uint256,uint256)))
func (_StapleProvider *StapleProviderCallerSession) GetPool(poolID *big.Int) (UIDataTypesPool, error) {
	return _StapleProvider.Contract.GetPool(&_StapleProvider.CallOpts, poolID)
}

// GetPools is a free data retrieval call binding the contract method 0x2952dde8.
//
// Solidity: function getPools(uint256[] idList) view returns(((uint256,address,address,uint8,uint8,uint256[]),(uint256,uint256,uint256))[])
func (_StapleProvider *StapleProviderCaller) GetPools(opts *bind.CallOpts, idList []*big.Int) ([]UIDataTypesPool, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "getPools", idList)

	if err != nil {
		return *new([]UIDataTypesPool), err
	}

	out0 := *abi.ConvertType(out[0], new([]UIDataTypesPool)).(*[]UIDataTypesPool)

	return out0, err

}

// GetPools is a free data retrieval call binding the contract method 0x2952dde8.
//
// Solidity: function getPools(uint256[] idList) view returns(((uint256,address,address,uint8,uint8,uint256[]),(uint256,uint256,uint256))[])
func (_StapleProvider *StapleProviderSession) GetPools(idList []*big.Int) ([]UIDataTypesPool, error) {
	return _StapleProvider.Contract.GetPools(&_StapleProvider.CallOpts, idList)
}

// GetPools is a free data retrieval call binding the contract method 0x2952dde8.
//
// Solidity: function getPools(uint256[] idList) view returns(((uint256,address,address,uint8,uint8,uint256[]),(uint256,uint256,uint256))[])
func (_StapleProvider *StapleProviderCallerSession) GetPools(idList []*big.Int) ([]UIDataTypesPool, error) {
	return _StapleProvider.Contract.GetPools(&_StapleProvider.CallOpts, idList)
}

// GetTokensByPath is a free data retrieval call binding the contract method 0x01ae405a.
//
// Solidity: function getTokensByPath(address tokenIn, uint256[] path) view returns(address[] tokens)
func (_StapleProvider *StapleProviderCaller) GetTokensByPath(opts *bind.CallOpts, tokenIn common.Address, path []*big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "getTokensByPath", tokenIn, path)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokensByPath is a free data retrieval call binding the contract method 0x01ae405a.
//
// Solidity: function getTokensByPath(address tokenIn, uint256[] path) view returns(address[] tokens)
func (_StapleProvider *StapleProviderSession) GetTokensByPath(tokenIn common.Address, path []*big.Int) ([]common.Address, error) {
	return _StapleProvider.Contract.GetTokensByPath(&_StapleProvider.CallOpts, tokenIn, path)
}

// GetTokensByPath is a free data retrieval call binding the contract method 0x01ae405a.
//
// Solidity: function getTokensByPath(address tokenIn, uint256[] path) view returns(address[] tokens)
func (_StapleProvider *StapleProviderCallerSession) GetTokensByPath(tokenIn common.Address, path []*big.Int) ([]common.Address, error) {
	return _StapleProvider.Contract.GetTokensByPath(&_StapleProvider.CallOpts, tokenIn, path)
}

// IncentivesController is a free data retrieval call binding the contract method 0xaf1df255.
//
// Solidity: function incentivesController() view returns(address)
func (_StapleProvider *StapleProviderCaller) IncentivesController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "incentivesController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IncentivesController is a free data retrieval call binding the contract method 0xaf1df255.
//
// Solidity: function incentivesController() view returns(address)
func (_StapleProvider *StapleProviderSession) IncentivesController() (common.Address, error) {
	return _StapleProvider.Contract.IncentivesController(&_StapleProvider.CallOpts)
}

// IncentivesController is a free data retrieval call binding the contract method 0xaf1df255.
//
// Solidity: function incentivesController() view returns(address)
func (_StapleProvider *StapleProviderCallerSession) IncentivesController() (common.Address, error) {
	return _StapleProvider.Contract.IncentivesController(&_StapleProvider.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address user) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256[],(uint256,uint256,uint256,uint256,uint256,uint256,(address,string,uint256,uint256,uint256,uint256,uint256)[])[],uint256)[])
func (_StapleProvider *StapleProviderCaller) UserInfo(opts *bind.CallOpts, user common.Address) ([]UIDataTypesUserPoolInfo, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "userInfo", user)

	if err != nil {
		return *new([]UIDataTypesUserPoolInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]UIDataTypesUserPoolInfo)).(*[]UIDataTypesUserPoolInfo)

	return out0, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address user) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256[],(uint256,uint256,uint256,uint256,uint256,uint256,(address,string,uint256,uint256,uint256,uint256,uint256)[])[],uint256)[])
func (_StapleProvider *StapleProviderSession) UserInfo(user common.Address) ([]UIDataTypesUserPoolInfo, error) {
	return _StapleProvider.Contract.UserInfo(&_StapleProvider.CallOpts, user)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address user) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256[],(uint256,uint256,uint256,uint256,uint256,uint256,(address,string,uint256,uint256,uint256,uint256,uint256)[])[],uint256)[])
func (_StapleProvider *StapleProviderCallerSession) UserInfo(user common.Address) ([]UIDataTypesUserPoolInfo, error) {
	return _StapleProvider.Contract.UserInfo(&_StapleProvider.CallOpts, user)
}

// UserPoolInfo is a free data retrieval call binding the contract method 0x81a48cbd.
//
// Solidity: function userPoolInfo(address user, address underlying) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256[],(uint256,uint256,uint256,uint256,uint256,uint256,(address,string,uint256,uint256,uint256,uint256,uint256)[])[],uint256) userPoolInfo)
func (_StapleProvider *StapleProviderCaller) UserPoolInfo(opts *bind.CallOpts, user common.Address, underlying common.Address) (UIDataTypesUserPoolInfo, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "userPoolInfo", user, underlying)

	if err != nil {
		return *new(UIDataTypesUserPoolInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(UIDataTypesUserPoolInfo)).(*UIDataTypesUserPoolInfo)

	return out0, err

}

// UserPoolInfo is a free data retrieval call binding the contract method 0x81a48cbd.
//
// Solidity: function userPoolInfo(address user, address underlying) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256[],(uint256,uint256,uint256,uint256,uint256,uint256,(address,string,uint256,uint256,uint256,uint256,uint256)[])[],uint256) userPoolInfo)
func (_StapleProvider *StapleProviderSession) UserPoolInfo(user common.Address, underlying common.Address) (UIDataTypesUserPoolInfo, error) {
	return _StapleProvider.Contract.UserPoolInfo(&_StapleProvider.CallOpts, user, underlying)
}

// UserPoolInfo is a free data retrieval call binding the contract method 0x81a48cbd.
//
// Solidity: function userPoolInfo(address user, address underlying) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256[],(uint256,uint256,uint256,uint256,uint256,uint256,(address,string,uint256,uint256,uint256,uint256,uint256)[])[],uint256) userPoolInfo)
func (_StapleProvider *StapleProviderCallerSession) UserPoolInfo(user common.Address, underlying common.Address) (UIDataTypesUserPoolInfo, error) {
	return _StapleProvider.Contract.UserPoolInfo(&_StapleProvider.CallOpts, user, underlying)
}

// UserVtpTokenIncentives is a free data retrieval call binding the contract method 0x31499fc8.
//
// Solidity: function userVtpTokenIncentives(address user, address underlying, uint256 vtpID) view returns((address,string,uint256,uint256,uint256,uint256,uint256)[])
func (_StapleProvider *StapleProviderCaller) UserVtpTokenIncentives(opts *bind.CallOpts, user common.Address, underlying common.Address, vtpID *big.Int) ([]UIDataTypesUserIncentivesInfo, error) {
	var out []interface{}
	err := _StapleProvider.contract.Call(opts, &out, "userVtpTokenIncentives", user, underlying, vtpID)

	if err != nil {
		return *new([]UIDataTypesUserIncentivesInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]UIDataTypesUserIncentivesInfo)).(*[]UIDataTypesUserIncentivesInfo)

	return out0, err

}

// UserVtpTokenIncentives is a free data retrieval call binding the contract method 0x31499fc8.
//
// Solidity: function userVtpTokenIncentives(address user, address underlying, uint256 vtpID) view returns((address,string,uint256,uint256,uint256,uint256,uint256)[])
func (_StapleProvider *StapleProviderSession) UserVtpTokenIncentives(user common.Address, underlying common.Address, vtpID *big.Int) ([]UIDataTypesUserIncentivesInfo, error) {
	return _StapleProvider.Contract.UserVtpTokenIncentives(&_StapleProvider.CallOpts, user, underlying, vtpID)
}

// UserVtpTokenIncentives is a free data retrieval call binding the contract method 0x31499fc8.
//
// Solidity: function userVtpTokenIncentives(address user, address underlying, uint256 vtpID) view returns((address,string,uint256,uint256,uint256,uint256,uint256)[])
func (_StapleProvider *StapleProviderCallerSession) UserVtpTokenIncentives(user common.Address, underlying common.Address, vtpID *big.Int) ([]UIDataTypesUserIncentivesInfo, error) {
	return _StapleProvider.Contract.UserVtpTokenIncentives(&_StapleProvider.CallOpts, user, underlying, vtpID)
}
