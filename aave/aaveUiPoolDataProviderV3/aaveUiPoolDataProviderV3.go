// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveUiPoolDataProviderV3

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

// IUiPoolDataProviderV3AggregatedReserveData is an auto generated low-level Go binding around an user-defined struct.
type IUiPoolDataProviderV3AggregatedReserveData struct {
	UnderlyingAsset                common.Address
	Name                           string
	Symbol                         string
	Decimals                       *big.Int
	BaseLTVasCollateral            *big.Int
	ReserveLiquidationThreshold    *big.Int
	ReserveLiquidationBonus        *big.Int
	ReserveFactor                  *big.Int
	UsageAsCollateralEnabled       bool
	BorrowingEnabled               bool
	StableBorrowRateEnabled        bool
	IsActive                       bool
	IsFrozen                       bool
	LiquidityIndex                 *big.Int
	VariableBorrowIndex            *big.Int
	LiquidityRate                  *big.Int
	VariableBorrowRate             *big.Int
	StableBorrowRate               *big.Int
	LastUpdateTimestamp            *big.Int
	ATokenAddress                  common.Address
	StableDebtTokenAddress         common.Address
	VariableDebtTokenAddress       common.Address
	InterestRateStrategyAddress    common.Address
	AvailableLiquidity             *big.Int
	TotalPrincipalStableDebt       *big.Int
	AverageStableRate              *big.Int
	StableDebtLastUpdateTimestamp  *big.Int
	TotalScaledVariableDebt        *big.Int
	PriceInMarketReferenceCurrency *big.Int
	PriceOracle                    common.Address
	VariableRateSlope1             *big.Int
	VariableRateSlope2             *big.Int
	StableRateSlope1               *big.Int
	StableRateSlope2               *big.Int
	BaseStableBorrowRate           *big.Int
	BaseVariableBorrowRate         *big.Int
	OptimalUsageRatio              *big.Int
	IsPaused                       bool
	AccruedToTreasury              *big.Int
	Unbacked                       *big.Int
	IsolationModeTotalDebt         *big.Int
	DebtCeiling                    *big.Int
	DebtCeilingDecimals            *big.Int
	EModeCategoryId                uint8
	BorrowCap                      *big.Int
	SupplyCap                      *big.Int
	EModeLtv                       uint16
	EModeLiquidationThreshold      uint16
	EModeLiquidationBonus          uint16
	EModePriceSource               common.Address
	EModeLabel                     string
	BorrowableInIsolation          bool
}

// IUiPoolDataProviderV3BaseCurrencyInfo is an auto generated low-level Go binding around an user-defined struct.
type IUiPoolDataProviderV3BaseCurrencyInfo struct {
	MarketReferenceCurrencyUnit       *big.Int
	MarketReferenceCurrencyPriceInUsd *big.Int
	NetworkBaseTokenPriceInUsd        *big.Int
	NetworkBaseTokenPriceDecimals     uint8
}

// IUiPoolDataProviderV3UserReserveData is an auto generated low-level Go binding around an user-defined struct.
type IUiPoolDataProviderV3UserReserveData struct {
	UnderlyingAsset                 common.Address
	ScaledATokenBalance             *big.Int
	UsageAsCollateralEnabledOnUser  bool
	StableBorrowRate                *big.Int
	ScaledVariableDebt              *big.Int
	PrincipalStableDebt             *big.Int
	StableBorrowLastUpdateTimestamp *big.Int
}

// AaveUiPoolDataProviderV3MetaData contains all meta data concerning the AaveUiPoolDataProviderV3 contract.
var AaveUiPoolDataProviderV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProviderV3\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseLTVasCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"stableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrincipalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableDebtLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalScaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInMarketReferenceCurrency\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseVariableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimalUsageRatio\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"accruedToTreasury\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"unbacked\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"isolationModeTotalDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"debtCeiling\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtCeilingDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"eModeCategoryId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyCap\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"eModeLtv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eModeLiquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eModeLiquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"eModePriceSource\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"eModeLabel\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"borrowableInIsolation\",\"type\":\"bool\"}],\"internalType\":\"structIUiPoolDataProviderV3.AggregatedReserveData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketReferenceCurrencyUnit\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketReferenceCurrencyPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"networkBaseTokenPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"networkBaseTokenPriceDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiPoolDataProviderV3.BaseCurrencyInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProviderV3\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProviderV3\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scaledATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabledOnUser\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowLastUpdateTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIUiPoolDataProviderV3.UserReserveData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AaveUiPoolDataProviderV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveUiPoolDataProviderV3MetaData.ABI instead.
var AaveUiPoolDataProviderV3ABI = AaveUiPoolDataProviderV3MetaData.ABI

// AaveUiPoolDataProviderV3 is an auto generated Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3 struct {
	AaveUiPoolDataProviderV3Caller     // Read-only binding to the contract
	AaveUiPoolDataProviderV3Transactor // Write-only binding to the contract
	AaveUiPoolDataProviderV3Filterer   // Log filterer for contract events
}

// AaveUiPoolDataProviderV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiPoolDataProviderV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiPoolDataProviderV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveUiPoolDataProviderV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiPoolDataProviderV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveUiPoolDataProviderV3Session struct {
	Contract     *AaveUiPoolDataProviderV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AaveUiPoolDataProviderV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveUiPoolDataProviderV3CallerSession struct {
	Contract *AaveUiPoolDataProviderV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AaveUiPoolDataProviderV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveUiPoolDataProviderV3TransactorSession struct {
	Contract     *AaveUiPoolDataProviderV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AaveUiPoolDataProviderV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3Raw struct {
	Contract *AaveUiPoolDataProviderV3 // Generic contract binding to access the raw methods on
}

// AaveUiPoolDataProviderV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3CallerRaw struct {
	Contract *AaveUiPoolDataProviderV3Caller // Generic read-only contract binding to access the raw methods on
}

// AaveUiPoolDataProviderV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3TransactorRaw struct {
	Contract *AaveUiPoolDataProviderV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveUiPoolDataProviderV3 creates a new instance of AaveUiPoolDataProviderV3, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3(address common.Address, backend bind.ContractBackend) (*AaveUiPoolDataProviderV3, error) {
	contract, err := bindAaveUiPoolDataProviderV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3{AaveUiPoolDataProviderV3Caller: AaveUiPoolDataProviderV3Caller{contract: contract}, AaveUiPoolDataProviderV3Transactor: AaveUiPoolDataProviderV3Transactor{contract: contract}, AaveUiPoolDataProviderV3Filterer: AaveUiPoolDataProviderV3Filterer{contract: contract}}, nil
}

// NewAaveUiPoolDataProviderV3Caller creates a new read-only instance of AaveUiPoolDataProviderV3, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3Caller(address common.Address, caller bind.ContractCaller) (*AaveUiPoolDataProviderV3Caller, error) {
	contract, err := bindAaveUiPoolDataProviderV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3Caller{contract: contract}, nil
}

// NewAaveUiPoolDataProviderV3Transactor creates a new write-only instance of AaveUiPoolDataProviderV3, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveUiPoolDataProviderV3Transactor, error) {
	contract, err := bindAaveUiPoolDataProviderV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3Transactor{contract: contract}, nil
}

// NewAaveUiPoolDataProviderV3Filterer creates a new log filterer instance of AaveUiPoolDataProviderV3, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveUiPoolDataProviderV3Filterer, error) {
	contract, err := bindAaveUiPoolDataProviderV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3Filterer{contract: contract}, nil
}

// bindAaveUiPoolDataProviderV3 binds a generic wrapper to an already deployed contract.
func bindAaveUiPoolDataProviderV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveUiPoolDataProviderV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveUiPoolDataProviderV3.Contract.AaveUiPoolDataProviderV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3.Contract.AaveUiPoolDataProviderV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3.Contract.AaveUiPoolDataProviderV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveUiPoolDataProviderV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3.Contract.contract.Transact(opts, method, params...)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint128,uint128,uint128,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3Caller) GetReservesData(opts *bind.CallOpts, provider common.Address) ([]IUiPoolDataProviderV3AggregatedReserveData, IUiPoolDataProviderV3BaseCurrencyInfo, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3.contract.Call(opts, &out, "getReservesData", provider)

	if err != nil {
		return *new([]IUiPoolDataProviderV3AggregatedReserveData), *new(IUiPoolDataProviderV3BaseCurrencyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUiPoolDataProviderV3AggregatedReserveData)).(*[]IUiPoolDataProviderV3AggregatedReserveData)
	out1 := *abi.ConvertType(out[1], new(IUiPoolDataProviderV3BaseCurrencyInfo)).(*IUiPoolDataProviderV3BaseCurrencyInfo)

	return out0, out1, err

}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint128,uint128,uint128,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3Session) GetReservesData(provider common.Address) ([]IUiPoolDataProviderV3AggregatedReserveData, IUiPoolDataProviderV3BaseCurrencyInfo, error) {
	return _AaveUiPoolDataProviderV3.Contract.GetReservesData(&_AaveUiPoolDataProviderV3.CallOpts, provider)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint128,uint128,uint128,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3CallerSession) GetReservesData(provider common.Address) ([]IUiPoolDataProviderV3AggregatedReserveData, IUiPoolDataProviderV3BaseCurrencyInfo, error) {
	return _AaveUiPoolDataProviderV3.Contract.GetReservesData(&_AaveUiPoolDataProviderV3.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3Caller) GetReservesList(opts *bind.CallOpts, provider common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3.contract.Call(opts, &out, "getReservesList", provider)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3Session) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _AaveUiPoolDataProviderV3.Contract.GetReservesList(&_AaveUiPoolDataProviderV3.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3CallerSession) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _AaveUiPoolDataProviderV3.Contract.GetReservesList(&_AaveUiPoolDataProviderV3.CallOpts, provider)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3Caller) GetUserReservesData(opts *bind.CallOpts, provider common.Address, user common.Address) ([]IUiPoolDataProviderV3UserReserveData, uint8, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3.contract.Call(opts, &out, "getUserReservesData", provider, user)

	if err != nil {
		return *new([]IUiPoolDataProviderV3UserReserveData), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUiPoolDataProviderV3UserReserveData)).(*[]IUiPoolDataProviderV3UserReserveData)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3Session) GetUserReservesData(provider common.Address, user common.Address) ([]IUiPoolDataProviderV3UserReserveData, uint8, error) {
	return _AaveUiPoolDataProviderV3.Contract.GetUserReservesData(&_AaveUiPoolDataProviderV3.CallOpts, provider, user)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_AaveUiPoolDataProviderV3 *AaveUiPoolDataProviderV3CallerSession) GetUserReservesData(provider common.Address, user common.Address) ([]IUiPoolDataProviderV3UserReserveData, uint8, error) {
	return _AaveUiPoolDataProviderV3.Contract.GetUserReservesData(&_AaveUiPoolDataProviderV3.CallOpts, provider, user)
}
