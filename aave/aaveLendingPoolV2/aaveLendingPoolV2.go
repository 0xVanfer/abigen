// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveLendingPoolV2

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

// DataTypesReserveConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveConfigurationMap struct {
	Data *big.Int
}

// DataTypesReserveData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveData struct {
	Configuration               DataTypesReserveConfigurationMap
	LiquidityIndex              *big.Int
	VariableBorrowIndex         *big.Int
	CurrentLiquidityRate        *big.Int
	CurrentVariableBorrowRate   *big.Int
	CurrentStableBorrowRate     *big.Int
	LastUpdateTimestamp         *big.Int
	ATokenAddress               common.Address
	StableDebtTokenAddress      common.Address
	VariableDebtTokenAddress    common.Address
	InterestRateStrategyAddress common.Address
	Id                          uint8
}

// DataTypesUserConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesUserConfigurationMap struct {
	Data *big.Int
}

// AaveLendingPoolV2MetaData contains all meta data concerning the AaveLendingPoolV2 contract.
var AaveLendingPoolV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_referralCode\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_onBehalfOf\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"modes\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentStableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"internalType\":\"structDataTypes.ReserveData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rateMode\",\"type\":\"uint256\"}],\"name\":\"swapBorrowRateMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveLendingPoolV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveLendingPoolV2MetaData.ABI instead.
var AaveLendingPoolV2ABI = AaveLendingPoolV2MetaData.ABI

// AaveLendingPoolV2 is an auto generated Go binding around an Ethereum contract.
type AaveLendingPoolV2 struct {
	AaveLendingPoolV2Caller     // Read-only binding to the contract
	AaveLendingPoolV2Transactor // Write-only binding to the contract
	AaveLendingPoolV2Filterer   // Log filterer for contract events
}

// AaveLendingPoolV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveLendingPoolV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveLendingPoolV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveLendingPoolV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveLendingPoolV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveLendingPoolV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveLendingPoolV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveLendingPoolV2Session struct {
	Contract     *AaveLendingPoolV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AaveLendingPoolV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveLendingPoolV2CallerSession struct {
	Contract *AaveLendingPoolV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AaveLendingPoolV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveLendingPoolV2TransactorSession struct {
	Contract     *AaveLendingPoolV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AaveLendingPoolV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveLendingPoolV2Raw struct {
	Contract *AaveLendingPoolV2 // Generic contract binding to access the raw methods on
}

// AaveLendingPoolV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveLendingPoolV2CallerRaw struct {
	Contract *AaveLendingPoolV2Caller // Generic read-only contract binding to access the raw methods on
}

// AaveLendingPoolV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveLendingPoolV2TransactorRaw struct {
	Contract *AaveLendingPoolV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveLendingPoolV2 creates a new instance of AaveLendingPoolV2, bound to a specific deployed contract.
func NewAaveLendingPoolV2(address common.Address, backend bind.ContractBackend) (*AaveLendingPoolV2, error) {
	contract, err := bindAaveLendingPoolV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveLendingPoolV2{AaveLendingPoolV2Caller: AaveLendingPoolV2Caller{contract: contract}, AaveLendingPoolV2Transactor: AaveLendingPoolV2Transactor{contract: contract}, AaveLendingPoolV2Filterer: AaveLendingPoolV2Filterer{contract: contract}}, nil
}

// NewAaveLendingPoolV2Caller creates a new read-only instance of AaveLendingPoolV2, bound to a specific deployed contract.
func NewAaveLendingPoolV2Caller(address common.Address, caller bind.ContractCaller) (*AaveLendingPoolV2Caller, error) {
	contract, err := bindAaveLendingPoolV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveLendingPoolV2Caller{contract: contract}, nil
}

// NewAaveLendingPoolV2Transactor creates a new write-only instance of AaveLendingPoolV2, bound to a specific deployed contract.
func NewAaveLendingPoolV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveLendingPoolV2Transactor, error) {
	contract, err := bindAaveLendingPoolV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveLendingPoolV2Transactor{contract: contract}, nil
}

// NewAaveLendingPoolV2Filterer creates a new log filterer instance of AaveLendingPoolV2, bound to a specific deployed contract.
func NewAaveLendingPoolV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveLendingPoolV2Filterer, error) {
	contract, err := bindAaveLendingPoolV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveLendingPoolV2Filterer{contract: contract}, nil
}

// bindAaveLendingPoolV2 binds a generic wrapper to an already deployed contract.
func bindAaveLendingPoolV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveLendingPoolV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveLendingPoolV2 *AaveLendingPoolV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveLendingPoolV2.Contract.AaveLendingPoolV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveLendingPoolV2 *AaveLendingPoolV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.AaveLendingPoolV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveLendingPoolV2 *AaveLendingPoolV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.AaveLendingPoolV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveLendingPoolV2 *AaveLendingPoolV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveLendingPoolV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveLendingPoolV2 *AaveLendingPoolV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveLendingPoolV2 *AaveLendingPoolV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.contract.Transact(opts, method, params...)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint8))
func (_AaveLendingPoolV2 *AaveLendingPoolV2Caller) GetReserveData(opts *bind.CallOpts, asset common.Address) (DataTypesReserveData, error) {
	var out []interface{}
	err := _AaveLendingPoolV2.contract.Call(opts, &out, "getReserveData", asset)

	if err != nil {
		return *new(DataTypesReserveData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveData)).(*DataTypesReserveData)

	return out0, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint8))
func (_AaveLendingPoolV2 *AaveLendingPoolV2Session) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _AaveLendingPoolV2.Contract.GetReserveData(&_AaveLendingPoolV2.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint8))
func (_AaveLendingPoolV2 *AaveLendingPoolV2CallerSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _AaveLendingPoolV2.Contract.GetReserveData(&_AaveLendingPoolV2.CallOpts, asset)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AaveLendingPoolV2 *AaveLendingPoolV2Caller) GetReservesList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AaveLendingPoolV2.contract.Call(opts, &out, "getReservesList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AaveLendingPoolV2 *AaveLendingPoolV2Session) GetReservesList() ([]common.Address, error) {
	return _AaveLendingPoolV2.Contract.GetReservesList(&_AaveLendingPoolV2.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AaveLendingPoolV2 *AaveLendingPoolV2CallerSession) GetReservesList() ([]common.Address, error) {
	return _AaveLendingPoolV2.Contract.GetReservesList(&_AaveLendingPoolV2.CallOpts)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralETH, uint256 totalDebtETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AaveLendingPoolV2 *AaveLendingPoolV2Caller) GetUserAccountData(opts *bind.CallOpts, user common.Address) (struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	var out []interface{}
	err := _AaveLendingPoolV2.contract.Call(opts, &out, "getUserAccountData", user)

	outstruct := new(struct {
		TotalCollateralETH          *big.Int
		TotalDebtETH                *big.Int
		AvailableBorrowsETH         *big.Int
		CurrentLiquidationThreshold *big.Int
		Ltv                         *big.Int
		HealthFactor                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCollateralETH = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalDebtETH = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsETH = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentLiquidationThreshold = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HealthFactor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralETH, uint256 totalDebtETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AaveLendingPoolV2 *AaveLendingPoolV2Session) GetUserAccountData(user common.Address) (struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _AaveLendingPoolV2.Contract.GetUserAccountData(&_AaveLendingPoolV2.CallOpts, user)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralETH, uint256 totalDebtETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AaveLendingPoolV2 *AaveLendingPoolV2CallerSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _AaveLendingPoolV2.Contract.GetUserAccountData(&_AaveLendingPoolV2.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AaveLendingPoolV2 *AaveLendingPoolV2Caller) GetUserConfiguration(opts *bind.CallOpts, user common.Address) (DataTypesUserConfigurationMap, error) {
	var out []interface{}
	err := _AaveLendingPoolV2.contract.Call(opts, &out, "getUserConfiguration", user)

	if err != nil {
		return *new(DataTypesUserConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesUserConfigurationMap)).(*DataTypesUserConfigurationMap)

	return out0, err

}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AaveLendingPoolV2 *AaveLendingPoolV2Session) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _AaveLendingPoolV2.Contract.GetUserConfiguration(&_AaveLendingPoolV2.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AaveLendingPoolV2 *AaveLendingPoolV2CallerSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _AaveLendingPoolV2.Contract.GetUserConfiguration(&_AaveLendingPoolV2.CallOpts, user)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address _asset, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode, address _onBehalfOf) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Transactor) Borrow(opts *bind.TransactOpts, _asset common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingPoolV2.contract.Transact(opts, "borrow", _asset, _amount, _interestRateMode, _referralCode, _onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address _asset, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode, address _onBehalfOf) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Session) Borrow(_asset common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.Borrow(&_AaveLendingPoolV2.TransactOpts, _asset, _amount, _interestRateMode, _referralCode, _onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address _asset, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode, address _onBehalfOf) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2TransactorSession) Borrow(_asset common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.Borrow(&_AaveLendingPoolV2.TransactOpts, _asset, _amount, _interestRateMode, _referralCode, _onBehalfOf)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address _asset, uint256 _amount, address _onBehalfOf, uint16 _referralCode) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Transactor) Deposit(opts *bind.TransactOpts, _asset common.Address, _amount *big.Int, _onBehalfOf common.Address, _referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingPoolV2.contract.Transact(opts, "deposit", _asset, _amount, _onBehalfOf, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address _asset, uint256 _amount, address _onBehalfOf, uint16 _referralCode) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Session) Deposit(_asset common.Address, _amount *big.Int, _onBehalfOf common.Address, _referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.Deposit(&_AaveLendingPoolV2.TransactOpts, _asset, _amount, _onBehalfOf, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address _asset, uint256 _amount, address _onBehalfOf, uint16 _referralCode) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2TransactorSession) Deposit(_asset common.Address, _amount *big.Int, _onBehalfOf common.Address, _referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.Deposit(&_AaveLendingPoolV2.TransactOpts, _asset, _amount, _onBehalfOf, _referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] modes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Transactor) FlashLoan(opts *bind.TransactOpts, receiverAddress common.Address, assets []common.Address, amounts []*big.Int, modes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingPoolV2.contract.Transact(opts, "flashLoan", receiverAddress, assets, amounts, modes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] modes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Session) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, modes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.FlashLoan(&_AaveLendingPoolV2.TransactOpts, receiverAddress, assets, amounts, modes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] modes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2TransactorSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, modes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.FlashLoan(&_AaveLendingPoolV2.TransactOpts, receiverAddress, assets, amounts, modes, onBehalfOf, params, referralCode)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address _asset, uint256 _amount, uint256 _rateMode, address _onBehalfOf) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Transactor) Repay(opts *bind.TransactOpts, _asset common.Address, _amount *big.Int, _rateMode *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingPoolV2.contract.Transact(opts, "repay", _asset, _amount, _rateMode, _onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address _asset, uint256 _amount, uint256 _rateMode, address _onBehalfOf) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Session) Repay(_asset common.Address, _amount *big.Int, _rateMode *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.Repay(&_AaveLendingPoolV2.TransactOpts, _asset, _amount, _rateMode, _onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address _asset, uint256 _amount, uint256 _rateMode, address _onBehalfOf) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2TransactorSession) Repay(_asset common.Address, _amount *big.Int, _rateMode *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.Repay(&_AaveLendingPoolV2.TransactOpts, _asset, _amount, _rateMode, _onBehalfOf)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _asset, bool _useAsCollateral) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Transactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, _asset common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _AaveLendingPoolV2.contract.Transact(opts, "setUserUseReserveAsCollateral", _asset, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _asset, bool _useAsCollateral) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Session) SetUserUseReserveAsCollateral(_asset common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.SetUserUseReserveAsCollateral(&_AaveLendingPoolV2.TransactOpts, _asset, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _asset, bool _useAsCollateral) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2TransactorSession) SetUserUseReserveAsCollateral(_asset common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.SetUserUseReserveAsCollateral(&_AaveLendingPoolV2.TransactOpts, _asset, _useAsCollateral)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address _asset, uint256 _rateMode) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Transactor) SwapBorrowRateMode(opts *bind.TransactOpts, _asset common.Address, _rateMode *big.Int) (*types.Transaction, error) {
	return _AaveLendingPoolV2.contract.Transact(opts, "swapBorrowRateMode", _asset, _rateMode)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address _asset, uint256 _rateMode) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Session) SwapBorrowRateMode(_asset common.Address, _rateMode *big.Int) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.SwapBorrowRateMode(&_AaveLendingPoolV2.TransactOpts, _asset, _rateMode)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address _asset, uint256 _rateMode) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2TransactorSession) SwapBorrowRateMode(_asset common.Address, _rateMode *big.Int) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.SwapBorrowRateMode(&_AaveLendingPoolV2.TransactOpts, _asset, _rateMode)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address _asset, uint256 _amount, address _to) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Transactor) Withdraw(opts *bind.TransactOpts, _asset common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _AaveLendingPoolV2.contract.Transact(opts, "withdraw", _asset, _amount, _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address _asset, uint256 _amount, address _to) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2Session) Withdraw(_asset common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.Withdraw(&_AaveLendingPoolV2.TransactOpts, _asset, _amount, _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address _asset, uint256 _amount, address _to) returns()
func (_AaveLendingPoolV2 *AaveLendingPoolV2TransactorSession) Withdraw(_asset common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _AaveLendingPoolV2.Contract.Withdraw(&_AaveLendingPoolV2.TransactOpts, _asset, _amount, _to)
}
