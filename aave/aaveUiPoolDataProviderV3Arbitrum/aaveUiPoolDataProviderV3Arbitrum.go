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

// GeneratedInterfaceAggregatedReserveData is an auto generated low-level Go binding around an user-defined struct.
type GeneratedInterfaceAggregatedReserveData struct {
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

// GeneratedInterfaceBaseCurrencyInfo is an auto generated low-level Go binding around an user-defined struct.
type GeneratedInterfaceBaseCurrencyInfo struct {
	MarketReferenceCurrencyUnit       *big.Int
	MarketReferenceCurrencyPriceInUsd *big.Int
	NetworkBaseTokenPriceInUsd        *big.Int
	NetworkBaseTokenPriceDecimals     uint8
}

// GeneratedInterfaceUserReserveData is an auto generated low-level Go binding around an user-defined struct.
type GeneratedInterfaceUserReserveData struct {
	UnderlyingAsset                 common.Address
	ScaledATokenBalance             *big.Int
	UsageAsCollateralEnabledOnUser  bool
	StableBorrowRate                *big.Int
	ScaledVariableDebt              *big.Int
	PrincipalStableDebt             *big.Int
	StableBorrowLastUpdateTimestamp *big.Int
}

// GeneratedInterfaceMetaData contains all meta data concerning the GeneratedInterface contract.
var GeneratedInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ETH_CURRENCY_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MKR_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseLTVasCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"stableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrincipalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableDebtLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalScaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInMarketReferenceCurrency\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseVariableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimalUsageRatio\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSiloedBorrowing\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"accruedToTreasury\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"unbacked\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"isolationModeTotalDebt\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"flashLoanEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"debtCeiling\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtCeilingDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"eModeCategoryId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyCap\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"eModeLtv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eModeLiquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eModeLiquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"eModePriceSource\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"eModeLabel\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"borrowableInIsolation\",\"type\":\"bool\"}],\"internalType\":\"structGeneratedInterface.AggregatedReserveData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketReferenceCurrencyUnit\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketReferenceCurrencyPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"networkBaseTokenPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"networkBaseTokenPriceDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structGeneratedInterface.BaseCurrencyInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scaledATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabledOnUser\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowLastUpdateTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structGeneratedInterface.UserReserveData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketReferenceCurrencyPriceInUsdProxyAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkBaseTokenPriceInUsdProxyAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GeneratedInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use GeneratedInterfaceMetaData.ABI instead.
var GeneratedInterfaceABI = GeneratedInterfaceMetaData.ABI

// Deprecated: Use GeneratedInterfaceMetaData.Sigs instead.
// GeneratedInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var GeneratedInterfaceFuncSigs = GeneratedInterfaceMetaData.Sigs

// GeneratedInterface is an auto generated Go binding around an Ethereum contract.
type GeneratedInterface struct {
	GeneratedInterfaceCaller     // Read-only binding to the contract
	GeneratedInterfaceTransactor // Write-only binding to the contract
	GeneratedInterfaceFilterer   // Log filterer for contract events
}

// GeneratedInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GeneratedInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneratedInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GeneratedInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneratedInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GeneratedInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneratedInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GeneratedInterfaceSession struct {
	Contract     *GeneratedInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GeneratedInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GeneratedInterfaceCallerSession struct {
	Contract *GeneratedInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GeneratedInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GeneratedInterfaceTransactorSession struct {
	Contract     *GeneratedInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GeneratedInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GeneratedInterfaceRaw struct {
	Contract *GeneratedInterface // Generic contract binding to access the raw methods on
}

// GeneratedInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GeneratedInterfaceCallerRaw struct {
	Contract *GeneratedInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// GeneratedInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GeneratedInterfaceTransactorRaw struct {
	Contract *GeneratedInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGeneratedInterface creates a new instance of GeneratedInterface, bound to a specific deployed contract.
func NewGeneratedInterface(address common.Address, backend bind.ContractBackend) (*GeneratedInterface, error) {
	contract, err := bindGeneratedInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GeneratedInterface{GeneratedInterfaceCaller: GeneratedInterfaceCaller{contract: contract}, GeneratedInterfaceTransactor: GeneratedInterfaceTransactor{contract: contract}, GeneratedInterfaceFilterer: GeneratedInterfaceFilterer{contract: contract}}, nil
}

// NewGeneratedInterfaceCaller creates a new read-only instance of GeneratedInterface, bound to a specific deployed contract.
func NewGeneratedInterfaceCaller(address common.Address, caller bind.ContractCaller) (*GeneratedInterfaceCaller, error) {
	contract, err := bindGeneratedInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GeneratedInterfaceCaller{contract: contract}, nil
}

// NewGeneratedInterfaceTransactor creates a new write-only instance of GeneratedInterface, bound to a specific deployed contract.
func NewGeneratedInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*GeneratedInterfaceTransactor, error) {
	contract, err := bindGeneratedInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GeneratedInterfaceTransactor{contract: contract}, nil
}

// NewGeneratedInterfaceFilterer creates a new log filterer instance of GeneratedInterface, bound to a specific deployed contract.
func NewGeneratedInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*GeneratedInterfaceFilterer, error) {
	contract, err := bindGeneratedInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GeneratedInterfaceFilterer{contract: contract}, nil
}

// bindGeneratedInterface binds a generic wrapper to an already deployed contract.
func bindGeneratedInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GeneratedInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneratedInterface *GeneratedInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GeneratedInterface.Contract.GeneratedInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneratedInterface *GeneratedInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneratedInterface.Contract.GeneratedInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneratedInterface *GeneratedInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneratedInterface.Contract.GeneratedInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneratedInterface *GeneratedInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GeneratedInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneratedInterface *GeneratedInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneratedInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneratedInterface *GeneratedInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneratedInterface.Contract.contract.Transact(opts, method, params...)
}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_GeneratedInterface *GeneratedInterfaceCaller) ETHCURRENCYUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GeneratedInterface.contract.Call(opts, &out, "ETH_CURRENCY_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_GeneratedInterface *GeneratedInterfaceSession) ETHCURRENCYUNIT() (*big.Int, error) {
	return _GeneratedInterface.Contract.ETHCURRENCYUNIT(&_GeneratedInterface.CallOpts)
}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_GeneratedInterface *GeneratedInterfaceCallerSession) ETHCURRENCYUNIT() (*big.Int, error) {
	return _GeneratedInterface.Contract.ETHCURRENCYUNIT(&_GeneratedInterface.CallOpts)
}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_GeneratedInterface *GeneratedInterfaceCaller) MKRADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GeneratedInterface.contract.Call(opts, &out, "MKR_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_GeneratedInterface *GeneratedInterfaceSession) MKRADDRESS() (common.Address, error) {
	return _GeneratedInterface.Contract.MKRADDRESS(&_GeneratedInterface.CallOpts)
}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_GeneratedInterface *GeneratedInterfaceCallerSession) MKRADDRESS() (common.Address, error) {
	return _GeneratedInterface.Contract.MKRADDRESS(&_GeneratedInterface.CallOpts)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_GeneratedInterface *GeneratedInterfaceCaller) Bytes32ToString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _GeneratedInterface.contract.Call(opts, &out, "bytes32ToString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_GeneratedInterface *GeneratedInterfaceSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _GeneratedInterface.Contract.Bytes32ToString(&_GeneratedInterface.CallOpts, _bytes32)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_GeneratedInterface *GeneratedInterfaceCallerSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _GeneratedInterface.Contract.Bytes32ToString(&_GeneratedInterface.CallOpts, _bytes32)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,bool,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_GeneratedInterface *GeneratedInterfaceCaller) GetReservesData(opts *bind.CallOpts, provider common.Address) ([]GeneratedInterfaceAggregatedReserveData, GeneratedInterfaceBaseCurrencyInfo, error) {
	var out []interface{}
	err := _GeneratedInterface.contract.Call(opts, &out, "getReservesData", provider)

	if err != nil {
		return *new([]GeneratedInterfaceAggregatedReserveData), *new(GeneratedInterfaceBaseCurrencyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]GeneratedInterfaceAggregatedReserveData)).(*[]GeneratedInterfaceAggregatedReserveData)
	out1 := *abi.ConvertType(out[1], new(GeneratedInterfaceBaseCurrencyInfo)).(*GeneratedInterfaceBaseCurrencyInfo)

	return out0, out1, err

}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,bool,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_GeneratedInterface *GeneratedInterfaceSession) GetReservesData(provider common.Address) ([]GeneratedInterfaceAggregatedReserveData, GeneratedInterfaceBaseCurrencyInfo, error) {
	return _GeneratedInterface.Contract.GetReservesData(&_GeneratedInterface.CallOpts, provider)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,bool,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool)[], (uint256,int256,int256,uint8))
func (_GeneratedInterface *GeneratedInterfaceCallerSession) GetReservesData(provider common.Address) ([]GeneratedInterfaceAggregatedReserveData, GeneratedInterfaceBaseCurrencyInfo, error) {
	return _GeneratedInterface.Contract.GetReservesData(&_GeneratedInterface.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_GeneratedInterface *GeneratedInterfaceCaller) GetReservesList(opts *bind.CallOpts, provider common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _GeneratedInterface.contract.Call(opts, &out, "getReservesList", provider)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_GeneratedInterface *GeneratedInterfaceSession) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _GeneratedInterface.Contract.GetReservesList(&_GeneratedInterface.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_GeneratedInterface *GeneratedInterfaceCallerSession) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _GeneratedInterface.Contract.GetReservesList(&_GeneratedInterface.CallOpts, provider)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_GeneratedInterface *GeneratedInterfaceCaller) GetUserReservesData(opts *bind.CallOpts, provider common.Address, user common.Address) ([]GeneratedInterfaceUserReserveData, uint8, error) {
	var out []interface{}
	err := _GeneratedInterface.contract.Call(opts, &out, "getUserReservesData", provider, user)

	if err != nil {
		return *new([]GeneratedInterfaceUserReserveData), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]GeneratedInterfaceUserReserveData)).(*[]GeneratedInterfaceUserReserveData)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_GeneratedInterface *GeneratedInterfaceSession) GetUserReservesData(provider common.Address, user common.Address) ([]GeneratedInterfaceUserReserveData, uint8, error) {
	return _GeneratedInterface.Contract.GetUserReservesData(&_GeneratedInterface.CallOpts, provider, user)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_GeneratedInterface *GeneratedInterfaceCallerSession) GetUserReservesData(provider common.Address, user common.Address) ([]GeneratedInterfaceUserReserveData, uint8, error) {
	return _GeneratedInterface.Contract.GetUserReservesData(&_GeneratedInterface.CallOpts, provider, user)
}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_GeneratedInterface *GeneratedInterfaceCaller) MarketReferenceCurrencyPriceInUsdProxyAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GeneratedInterface.contract.Call(opts, &out, "marketReferenceCurrencyPriceInUsdProxyAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_GeneratedInterface *GeneratedInterfaceSession) MarketReferenceCurrencyPriceInUsdProxyAggregator() (common.Address, error) {
	return _GeneratedInterface.Contract.MarketReferenceCurrencyPriceInUsdProxyAggregator(&_GeneratedInterface.CallOpts)
}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_GeneratedInterface *GeneratedInterfaceCallerSession) MarketReferenceCurrencyPriceInUsdProxyAggregator() (common.Address, error) {
	return _GeneratedInterface.Contract.MarketReferenceCurrencyPriceInUsdProxyAggregator(&_GeneratedInterface.CallOpts)
}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_GeneratedInterface *GeneratedInterfaceCaller) NetworkBaseTokenPriceInUsdProxyAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GeneratedInterface.contract.Call(opts, &out, "networkBaseTokenPriceInUsdProxyAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_GeneratedInterface *GeneratedInterfaceSession) NetworkBaseTokenPriceInUsdProxyAggregator() (common.Address, error) {
	return _GeneratedInterface.Contract.NetworkBaseTokenPriceInUsdProxyAggregator(&_GeneratedInterface.CallOpts)
}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_GeneratedInterface *GeneratedInterfaceCallerSession) NetworkBaseTokenPriceInUsdProxyAggregator() (common.Address, error) {
	return _GeneratedInterface.Contract.NetworkBaseTokenPriceInUsdProxyAggregator(&_GeneratedInterface.CallOpts)
}
