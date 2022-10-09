// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveUiPoolDataProviderV3Polygon

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
	IsSiloedBorrowing              bool
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

// AaveUiPoolDataProviderV3PolygonMetaData contains all meta data concerning the AaveUiPoolDataProviderV3Polygon contract.
var AaveUiPoolDataProviderV3PolygonMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ETH_CURRENCY_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MKR_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseLTVasCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"stableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrincipalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableDebtLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalScaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInMarketReferenceCurrency\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseVariableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimalUsageRatio\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSiloedBorrowing\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"accruedToTreasury\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"unbacked\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"isolationModeTotalDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"debtCeiling\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtCeilingDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"eModeCategoryId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyCap\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"eModeLtv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eModeLiquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eModeLiquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"eModePriceSource\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"eModeLabel\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"borrowableInIsolation\",\"type\":\"bool\"}],\"internalType\":\"structIUiPoolDataProviderV3.AggregatedReserveData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketReferenceCurrencyUnit\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketReferenceCurrencyPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"networkBaseTokenPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"networkBaseTokenPriceDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiPoolDataProviderV3.BaseCurrencyInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketReferenceCurrencyPriceInUsdProxyAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkBaseTokenPriceInUsdProxyAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AaveUiPoolDataProviderV3PolygonABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveUiPoolDataProviderV3PolygonMetaData.ABI instead.
var AaveUiPoolDataProviderV3PolygonABI = AaveUiPoolDataProviderV3PolygonMetaData.ABI

// AaveUiPoolDataProviderV3Polygon is an auto generated Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3Polygon struct {
	AaveUiPoolDataProviderV3PolygonCaller     // Read-only binding to the contract
	AaveUiPoolDataProviderV3PolygonTransactor // Write-only binding to the contract
	AaveUiPoolDataProviderV3PolygonFilterer   // Log filterer for contract events
}

// AaveUiPoolDataProviderV3PolygonCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3PolygonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiPoolDataProviderV3PolygonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3PolygonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiPoolDataProviderV3PolygonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveUiPoolDataProviderV3PolygonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiPoolDataProviderV3PolygonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveUiPoolDataProviderV3PolygonSession struct {
	Contract     *AaveUiPoolDataProviderV3Polygon // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AaveUiPoolDataProviderV3PolygonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveUiPoolDataProviderV3PolygonCallerSession struct {
	Contract *AaveUiPoolDataProviderV3PolygonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// AaveUiPoolDataProviderV3PolygonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveUiPoolDataProviderV3PolygonTransactorSession struct {
	Contract     *AaveUiPoolDataProviderV3PolygonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// AaveUiPoolDataProviderV3PolygonRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3PolygonRaw struct {
	Contract *AaveUiPoolDataProviderV3Polygon // Generic contract binding to access the raw methods on
}

// AaveUiPoolDataProviderV3PolygonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3PolygonCallerRaw struct {
	Contract *AaveUiPoolDataProviderV3PolygonCaller // Generic read-only contract binding to access the raw methods on
}

// AaveUiPoolDataProviderV3PolygonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3PolygonTransactorRaw struct {
	Contract *AaveUiPoolDataProviderV3PolygonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveUiPoolDataProviderV3Polygon creates a new instance of AaveUiPoolDataProviderV3Polygon, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3Polygon(address common.Address, backend bind.ContractBackend) (*AaveUiPoolDataProviderV3Polygon, error) {
	contract, err := bindAaveUiPoolDataProviderV3Polygon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3Polygon{AaveUiPoolDataProviderV3PolygonCaller: AaveUiPoolDataProviderV3PolygonCaller{contract: contract}, AaveUiPoolDataProviderV3PolygonTransactor: AaveUiPoolDataProviderV3PolygonTransactor{contract: contract}, AaveUiPoolDataProviderV3PolygonFilterer: AaveUiPoolDataProviderV3PolygonFilterer{contract: contract}}, nil
}

// NewAaveUiPoolDataProviderV3PolygonCaller creates a new read-only instance of AaveUiPoolDataProviderV3Polygon, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3PolygonCaller(address common.Address, caller bind.ContractCaller) (*AaveUiPoolDataProviderV3PolygonCaller, error) {
	contract, err := bindAaveUiPoolDataProviderV3Polygon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3PolygonCaller{contract: contract}, nil
}

// NewAaveUiPoolDataProviderV3PolygonTransactor creates a new write-only instance of AaveUiPoolDataProviderV3Polygon, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3PolygonTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveUiPoolDataProviderV3PolygonTransactor, error) {
	contract, err := bindAaveUiPoolDataProviderV3Polygon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3PolygonTransactor{contract: contract}, nil
}

// NewAaveUiPoolDataProviderV3PolygonFilterer creates a new log filterer instance of AaveUiPoolDataProviderV3Polygon, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3PolygonFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveUiPoolDataProviderV3PolygonFilterer, error) {
	contract, err := bindAaveUiPoolDataProviderV3Polygon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3PolygonFilterer{contract: contract}, nil
}

// bindAaveUiPoolDataProviderV3Polygon binds a generic wrapper to an already deployed contract.
func bindAaveUiPoolDataProviderV3Polygon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveUiPoolDataProviderV3PolygonABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveUiPoolDataProviderV3Polygon.Contract.AaveUiPoolDataProviderV3PolygonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.AaveUiPoolDataProviderV3PolygonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.AaveUiPoolDataProviderV3PolygonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveUiPoolDataProviderV3Polygon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.contract.Transact(opts, method, params...)
}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCaller) ETHCURRENCYUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Polygon.contract.Call(opts, &out, "ETH_CURRENCY_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonSession) ETHCURRENCYUNIT() (*big.Int, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.ETHCURRENCYUNIT(&_AaveUiPoolDataProviderV3Polygon.CallOpts)
}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCallerSession) ETHCURRENCYUNIT() (*big.Int, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.ETHCURRENCYUNIT(&_AaveUiPoolDataProviderV3Polygon.CallOpts)
}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCaller) MKRADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Polygon.contract.Call(opts, &out, "MKR_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonSession) MKRADDRESS() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.MKRADDRESS(&_AaveUiPoolDataProviderV3Polygon.CallOpts)
}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCallerSession) MKRADDRESS() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.MKRADDRESS(&_AaveUiPoolDataProviderV3Polygon.CallOpts)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCaller) Bytes32ToString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Polygon.contract.Call(opts, &out, "bytes32ToString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.Bytes32ToString(&_AaveUiPoolDataProviderV3Polygon.CallOpts, _bytes32)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCallerSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.Bytes32ToString(&_AaveUiPoolDataProviderV3Polygon.CallOpts, _bytes32)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCaller) GetReservesData(opts *bind.CallOpts, provider common.Address) ([]IUiPoolDataProviderV3AggregatedReserveData, IUiPoolDataProviderV3BaseCurrencyInfo, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Polygon.contract.Call(opts, &out, "getReservesData", provider)

	if err != nil {
		return *new([]IUiPoolDataProviderV3AggregatedReserveData), *new(IUiPoolDataProviderV3BaseCurrencyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUiPoolDataProviderV3AggregatedReserveData)).(*[]IUiPoolDataProviderV3AggregatedReserveData)
	out1 := *abi.ConvertType(out[1], new(IUiPoolDataProviderV3BaseCurrencyInfo)).(*IUiPoolDataProviderV3BaseCurrencyInfo)

	return out0, out1, err

}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonSession) GetReservesData(provider common.Address) ([]IUiPoolDataProviderV3AggregatedReserveData, IUiPoolDataProviderV3BaseCurrencyInfo, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.GetReservesData(&_AaveUiPoolDataProviderV3Polygon.CallOpts, provider)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCallerSession) GetReservesData(provider common.Address) ([]IUiPoolDataProviderV3AggregatedReserveData, IUiPoolDataProviderV3BaseCurrencyInfo, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.GetReservesData(&_AaveUiPoolDataProviderV3Polygon.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCaller) GetReservesList(opts *bind.CallOpts, provider common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Polygon.contract.Call(opts, &out, "getReservesList", provider)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonSession) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.GetReservesList(&_AaveUiPoolDataProviderV3Polygon.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCallerSession) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.GetReservesList(&_AaveUiPoolDataProviderV3Polygon.CallOpts, provider)
}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCaller) MarketReferenceCurrencyPriceInUsdProxyAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Polygon.contract.Call(opts, &out, "marketReferenceCurrencyPriceInUsdProxyAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonSession) MarketReferenceCurrencyPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.MarketReferenceCurrencyPriceInUsdProxyAggregator(&_AaveUiPoolDataProviderV3Polygon.CallOpts)
}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCallerSession) MarketReferenceCurrencyPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.MarketReferenceCurrencyPriceInUsdProxyAggregator(&_AaveUiPoolDataProviderV3Polygon.CallOpts)
}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCaller) NetworkBaseTokenPriceInUsdProxyAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Polygon.contract.Call(opts, &out, "networkBaseTokenPriceInUsdProxyAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonSession) NetworkBaseTokenPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.NetworkBaseTokenPriceInUsdProxyAggregator(&_AaveUiPoolDataProviderV3Polygon.CallOpts)
}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Polygon *AaveUiPoolDataProviderV3PolygonCallerSession) NetworkBaseTokenPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Polygon.Contract.NetworkBaseTokenPriceInUsdProxyAggregator(&_AaveUiPoolDataProviderV3Polygon.CallOpts)
}
