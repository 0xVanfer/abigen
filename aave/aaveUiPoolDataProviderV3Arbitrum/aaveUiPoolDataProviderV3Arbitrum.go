// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveUiPoolDataProviderV3Arbitrum

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

// aaveUiPoolDataProviderV3ArbitrumAggregatedReserveData is an auto generated low-level Go binding around an user-defined struct.
type aaveUiPoolDataProviderV3ArbitrumAggregatedReserveData struct {
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
	FlashLoanEnabled               bool
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

// aaveUiPoolDataProviderV3ArbitrumBaseCurrencyInfo is an auto generated low-level Go binding around an user-defined struct.
type aaveUiPoolDataProviderV3ArbitrumBaseCurrencyInfo struct {
	MarketReferenceCurrencyUnit       *big.Int
	MarketReferenceCurrencyPriceInUsd *big.Int
	NetworkBaseTokenPriceInUsd        *big.Int
	NetworkBaseTokenPriceDecimals     uint8
}

// aaveUiPoolDataProviderV3ArbitrumUserReserveData is an auto generated low-level Go binding around an user-defined struct.
type aaveUiPoolDataProviderV3ArbitrumUserReserveData struct {
	UnderlyingAsset                 common.Address
	ScaledATokenBalance             *big.Int
	UsageAsCollateralEnabledOnUser  bool
	StableBorrowRate                *big.Int
	ScaledVariableDebt              *big.Int
	PrincipalStableDebt             *big.Int
	StableBorrowLastUpdateTimestamp *big.Int
}

// AaveUiPoolDataProviderV3ArbitrumMetaData contains all meta data concerning the AaveUiPoolDataProviderV3Arbitrum contract.
var AaveUiPoolDataProviderV3ArbitrumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ETH_CURRENCY_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MKR_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseLTVasCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"stableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrincipalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableDebtLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalScaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInMarketReferenceCurrency\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseVariableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimalUsageRatio\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSiloedBorrowing\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"accruedToTreasury\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"unbacked\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"isolationModeTotalDebt\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"flashLoanEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"debtCeiling\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtCeilingDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"eModeCategoryId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyCap\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"eModeLtv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eModeLiquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eModeLiquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"eModePriceSource\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"eModeLabel\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"borrowableInIsolation\",\"type\":\"bool\"}],\"internalType\":\"structaaveUiPoolDataProviderV3Arbitrum.AggregatedReserveData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketReferenceCurrencyUnit\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketReferenceCurrencyPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"networkBaseTokenPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"networkBaseTokenPriceDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structaaveUiPoolDataProviderV3Arbitrum.BaseCurrencyInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scaledATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabledOnUser\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowLastUpdateTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structaaveUiPoolDataProviderV3Arbitrum.UserReserveData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketReferenceCurrencyPriceInUsdProxyAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkBaseTokenPriceInUsdProxyAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0496f53a": "ETH_CURRENCY_UNIT()",
		"825ffd92": "MKR_ADDRESS()",
		"9201de55": "bytes32ToString(bytes32)",
		"ec489c21": "getReservesData(address)",
		"586c1442": "getReservesList(address)",
		"51974cc0": "getUserReservesData(address,address)",
		"d22cf68a": "marketReferenceCurrencyPriceInUsdProxyAggregator()",
		"3c1740ed": "networkBaseTokenPriceInUsdProxyAggregator()",
	},
}

// AaveUiPoolDataProviderV3ArbitrumABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveUiPoolDataProviderV3ArbitrumMetaData.ABI instead.
var AaveUiPoolDataProviderV3ArbitrumABI = AaveUiPoolDataProviderV3ArbitrumMetaData.ABI

// Deprecated: Use AaveUiPoolDataProviderV3ArbitrumMetaData.Sigs instead.
// AaveUiPoolDataProviderV3ArbitrumFuncSigs maps the 4-byte function signature to its string representation.
var AaveUiPoolDataProviderV3ArbitrumFuncSigs = AaveUiPoolDataProviderV3ArbitrumMetaData.Sigs

// AaveUiPoolDataProviderV3Arbitrum is an auto generated Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3Arbitrum struct {
	AaveUiPoolDataProviderV3ArbitrumCaller     // Read-only binding to the contract
	AaveUiPoolDataProviderV3ArbitrumTransactor // Write-only binding to the contract
	AaveUiPoolDataProviderV3ArbitrumFilterer   // Log filterer for contract events
}

// AaveUiPoolDataProviderV3ArbitrumCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3ArbitrumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiPoolDataProviderV3ArbitrumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3ArbitrumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiPoolDataProviderV3ArbitrumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveUiPoolDataProviderV3ArbitrumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiPoolDataProviderV3ArbitrumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveUiPoolDataProviderV3ArbitrumSession struct {
	Contract     *AaveUiPoolDataProviderV3Arbitrum // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AaveUiPoolDataProviderV3ArbitrumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveUiPoolDataProviderV3ArbitrumCallerSession struct {
	Contract *AaveUiPoolDataProviderV3ArbitrumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// AaveUiPoolDataProviderV3ArbitrumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveUiPoolDataProviderV3ArbitrumTransactorSession struct {
	Contract     *AaveUiPoolDataProviderV3ArbitrumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// AaveUiPoolDataProviderV3ArbitrumRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3ArbitrumRaw struct {
	Contract *AaveUiPoolDataProviderV3Arbitrum // Generic contract binding to access the raw methods on
}

// AaveUiPoolDataProviderV3ArbitrumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3ArbitrumCallerRaw struct {
	Contract *AaveUiPoolDataProviderV3ArbitrumCaller // Generic read-only contract binding to access the raw methods on
}

// AaveUiPoolDataProviderV3ArbitrumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveUiPoolDataProviderV3ArbitrumTransactorRaw struct {
	Contract *AaveUiPoolDataProviderV3ArbitrumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveUiPoolDataProviderV3Arbitrum creates a new instance of AaveUiPoolDataProviderV3Arbitrum, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3Arbitrum(address common.Address, backend bind.ContractBackend) (*AaveUiPoolDataProviderV3Arbitrum, error) {
	contract, err := bindAaveUiPoolDataProviderV3Arbitrum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3Arbitrum{AaveUiPoolDataProviderV3ArbitrumCaller: AaveUiPoolDataProviderV3ArbitrumCaller{contract: contract}, AaveUiPoolDataProviderV3ArbitrumTransactor: AaveUiPoolDataProviderV3ArbitrumTransactor{contract: contract}, AaveUiPoolDataProviderV3ArbitrumFilterer: AaveUiPoolDataProviderV3ArbitrumFilterer{contract: contract}}, nil
}

// NewAaveUiPoolDataProviderV3ArbitrumCaller creates a new read-only instance of AaveUiPoolDataProviderV3Arbitrum, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3ArbitrumCaller(address common.Address, caller bind.ContractCaller) (*AaveUiPoolDataProviderV3ArbitrumCaller, error) {
	contract, err := bindAaveUiPoolDataProviderV3Arbitrum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3ArbitrumCaller{contract: contract}, nil
}

// NewAaveUiPoolDataProviderV3ArbitrumTransactor creates a new write-only instance of AaveUiPoolDataProviderV3Arbitrum, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3ArbitrumTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveUiPoolDataProviderV3ArbitrumTransactor, error) {
	contract, err := bindAaveUiPoolDataProviderV3Arbitrum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3ArbitrumTransactor{contract: contract}, nil
}

// NewAaveUiPoolDataProviderV3ArbitrumFilterer creates a new log filterer instance of AaveUiPoolDataProviderV3Arbitrum, bound to a specific deployed contract.
func NewAaveUiPoolDataProviderV3ArbitrumFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveUiPoolDataProviderV3ArbitrumFilterer, error) {
	contract, err := bindAaveUiPoolDataProviderV3Arbitrum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveUiPoolDataProviderV3ArbitrumFilterer{contract: contract}, nil
}

// bindAaveUiPoolDataProviderV3Arbitrum binds a generic wrapper to an already deployed contract.
func bindAaveUiPoolDataProviderV3Arbitrum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveUiPoolDataProviderV3ArbitrumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.AaveUiPoolDataProviderV3ArbitrumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.AaveUiPoolDataProviderV3ArbitrumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.AaveUiPoolDataProviderV3ArbitrumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.contract.Transact(opts, method, params...)
}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCaller) ETHCURRENCYUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Arbitrum.contract.Call(opts, &out, "ETH_CURRENCY_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumSession) ETHCURRENCYUNIT() (*big.Int, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.ETHCURRENCYUNIT(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts)
}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCallerSession) ETHCURRENCYUNIT() (*big.Int, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.ETHCURRENCYUNIT(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts)
}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCaller) MKRADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Arbitrum.contract.Call(opts, &out, "MKR_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumSession) MKRADDRESS() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.MKRADDRESS(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts)
}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCallerSession) MKRADDRESS() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.MKRADDRESS(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCaller) Bytes32ToString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Arbitrum.contract.Call(opts, &out, "bytes32ToString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.Bytes32ToString(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts, _bytes32)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCallerSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.Bytes32ToString(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts, _bytes32)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,bool,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCaller) GetReservesData(opts *bind.CallOpts, provider common.Address) ([]aaveUiPoolDataProviderV3ArbitrumAggregatedReserveData, aaveUiPoolDataProviderV3ArbitrumBaseCurrencyInfo, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Arbitrum.contract.Call(opts, &out, "getReservesData", provider)

	if err != nil {
		return *new([]aaveUiPoolDataProviderV3ArbitrumAggregatedReserveData), *new(aaveUiPoolDataProviderV3ArbitrumBaseCurrencyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]aaveUiPoolDataProviderV3ArbitrumAggregatedReserveData)).(*[]aaveUiPoolDataProviderV3ArbitrumAggregatedReserveData)
	out1 := *abi.ConvertType(out[1], new(aaveUiPoolDataProviderV3ArbitrumBaseCurrencyInfo)).(*aaveUiPoolDataProviderV3ArbitrumBaseCurrencyInfo)

	return out0, out1, err

}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,bool,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumSession) GetReservesData(provider common.Address) ([]aaveUiPoolDataProviderV3ArbitrumAggregatedReserveData, aaveUiPoolDataProviderV3ArbitrumBaseCurrencyInfo, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.GetReservesData(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts, provider)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,bool,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCallerSession) GetReservesData(provider common.Address) ([]aaveUiPoolDataProviderV3ArbitrumAggregatedReserveData, aaveUiPoolDataProviderV3ArbitrumBaseCurrencyInfo, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.GetReservesData(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCaller) GetReservesList(opts *bind.CallOpts, provider common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Arbitrum.contract.Call(opts, &out, "getReservesList", provider)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumSession) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.GetReservesList(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCallerSession) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.GetReservesList(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts, provider)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCaller) GetUserReservesData(opts *bind.CallOpts, provider common.Address, user common.Address) ([]aaveUiPoolDataProviderV3ArbitrumUserReserveData, uint8, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Arbitrum.contract.Call(opts, &out, "getUserReservesData", provider, user)

	if err != nil {
		return *new([]aaveUiPoolDataProviderV3ArbitrumUserReserveData), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]aaveUiPoolDataProviderV3ArbitrumUserReserveData)).(*[]aaveUiPoolDataProviderV3ArbitrumUserReserveData)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumSession) GetUserReservesData(provider common.Address, user common.Address) ([]aaveUiPoolDataProviderV3ArbitrumUserReserveData, uint8, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.GetUserReservesData(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts, provider, user)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCallerSession) GetUserReservesData(provider common.Address, user common.Address) ([]aaveUiPoolDataProviderV3ArbitrumUserReserveData, uint8, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.GetUserReservesData(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts, provider, user)
}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCaller) MarketReferenceCurrencyPriceInUsdProxyAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Arbitrum.contract.Call(opts, &out, "marketReferenceCurrencyPriceInUsdProxyAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumSession) MarketReferenceCurrencyPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.MarketReferenceCurrencyPriceInUsdProxyAggregator(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts)
}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCallerSession) MarketReferenceCurrencyPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.MarketReferenceCurrencyPriceInUsdProxyAggregator(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts)
}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCaller) NetworkBaseTokenPriceInUsdProxyAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveUiPoolDataProviderV3Arbitrum.contract.Call(opts, &out, "networkBaseTokenPriceInUsdProxyAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumSession) NetworkBaseTokenPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.NetworkBaseTokenPriceInUsdProxyAggregator(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts)
}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_AaveUiPoolDataProviderV3Arbitrum *AaveUiPoolDataProviderV3ArbitrumCallerSession) NetworkBaseTokenPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveUiPoolDataProviderV3Arbitrum.Contract.NetworkBaseTokenPriceInUsdProxyAggregator(&_AaveUiPoolDataProviderV3Arbitrum.CallOpts)
}
