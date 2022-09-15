// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveUiIncentiveDataProviderV3

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

// IUiIncentiveDataProviderV3AggregatedReserveIncentiveData is an auto generated low-level Go binding around an user-defined struct.
type IUiIncentiveDataProviderV3AggregatedReserveIncentiveData struct {
	UnderlyingAsset common.Address
	AIncentiveData  IUiIncentiveDataProviderV3IncentiveData
	VIncentiveData  IUiIncentiveDataProviderV3IncentiveData
	SIncentiveData  IUiIncentiveDataProviderV3IncentiveData
}

// IUiIncentiveDataProviderV3IncentiveData is an auto generated low-level Go binding around an user-defined struct.
type IUiIncentiveDataProviderV3IncentiveData struct {
	TokenAddress               common.Address
	IncentiveControllerAddress common.Address
	RewardsTokenInformation    []IUiIncentiveDataProviderV3RewardInfo
}

// IUiIncentiveDataProviderV3RewardInfo is an auto generated low-level Go binding around an user-defined struct.
type IUiIncentiveDataProviderV3RewardInfo struct {
	RewardTokenSymbol             string
	RewardTokenAddress            common.Address
	RewardOracleAddress           common.Address
	EmissionPerSecond             *big.Int
	IncentivesLastUpdateTimestamp *big.Int
	TokenIncentivesIndex          *big.Int
	EmissionEndTimestamp          *big.Int
	RewardPriceFeed               *big.Int
	RewardTokenDecimals           uint8
	Precision                     uint8
	PriceFeedDecimals             uint8
}

// IUiIncentiveDataProviderV3UserIncentiveData is an auto generated low-level Go binding around an user-defined struct.
type IUiIncentiveDataProviderV3UserIncentiveData struct {
	TokenAddress               common.Address
	IncentiveControllerAddress common.Address
	UserRewardsInformation     []IUiIncentiveDataProviderV3UserRewardInfo
}

// IUiIncentiveDataProviderV3UserReserveIncentiveData is an auto generated low-level Go binding around an user-defined struct.
type IUiIncentiveDataProviderV3UserReserveIncentiveData struct {
	UnderlyingAsset          common.Address
	ATokenIncentivesUserData IUiIncentiveDataProviderV3UserIncentiveData
	VTokenIncentivesUserData IUiIncentiveDataProviderV3UserIncentiveData
	STokenIncentivesUserData IUiIncentiveDataProviderV3UserIncentiveData
}

// IUiIncentiveDataProviderV3UserRewardInfo is an auto generated low-level Go binding around an user-defined struct.
type IUiIncentiveDataProviderV3UserRewardInfo struct {
	RewardTokenSymbol        string
	RewardOracleAddress      common.Address
	RewardTokenAddress       common.Address
	UserUnclaimedRewards     *big.Int
	TokenIncentivesUserIndex *big.Int
	RewardPriceFeed          *big.Int
	PriceFeedDecimals        uint8
	RewardTokenDecimals      uint8
}

// AaveUiIncentiveDataProviderV3MetaData contains all meta data concerning the AaveUiIncentiveDataProviderV3 contract.
var AaveUiIncentiveDataProviderV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProviderV3\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getFullReservesIncentiveData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"emissionPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"incentivesLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"emissionEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.RewardInfo[]\",\"name\":\"rewardsTokenInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.IncentiveData\",\"name\":\"aIncentiveData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"emissionPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"incentivesLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"emissionEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.RewardInfo[]\",\"name\":\"rewardsTokenInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.IncentiveData\",\"name\":\"vIncentiveData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"emissionPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"incentivesLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"emissionEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.RewardInfo[]\",\"name\":\"rewardsTokenInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.IncentiveData\",\"name\":\"sIncentiveData\",\"type\":\"tuple\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.AggregatedReserveIncentiveData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userUnclaimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesUserIndex\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserRewardInfo[]\",\"name\":\"userRewardsInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserIncentiveData\",\"name\":\"aTokenIncentivesUserData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userUnclaimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesUserIndex\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserRewardInfo[]\",\"name\":\"userRewardsInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserIncentiveData\",\"name\":\"vTokenIncentivesUserData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userUnclaimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesUserIndex\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserRewardInfo[]\",\"name\":\"userRewardsInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserIncentiveData\",\"name\":\"sTokenIncentivesUserData\",\"type\":\"tuple\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserReserveIncentiveData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProviderV3\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesIncentivesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"emissionPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"incentivesLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"emissionEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.RewardInfo[]\",\"name\":\"rewardsTokenInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.IncentiveData\",\"name\":\"aIncentiveData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"emissionPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"incentivesLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"emissionEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.RewardInfo[]\",\"name\":\"rewardsTokenInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.IncentiveData\",\"name\":\"vIncentiveData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"emissionPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"incentivesLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"emissionEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.RewardInfo[]\",\"name\":\"rewardsTokenInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.IncentiveData\",\"name\":\"sIncentiveData\",\"type\":\"tuple\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.AggregatedReserveIncentiveData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProviderV3\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReservesIncentivesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userUnclaimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesUserIndex\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserRewardInfo[]\",\"name\":\"userRewardsInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserIncentiveData\",\"name\":\"aTokenIncentivesUserData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userUnclaimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesUserIndex\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserRewardInfo[]\",\"name\":\"userRewardsInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserIncentiveData\",\"name\":\"vTokenIncentivesUserData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incentiveControllerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"rewardTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"rewardOracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userUnclaimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIncentivesUserIndex\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardPriceFeed\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"priceFeedDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rewardTokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserRewardInfo[]\",\"name\":\"userRewardsInformation\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserIncentiveData\",\"name\":\"sTokenIncentivesUserData\",\"type\":\"tuple\"}],\"internalType\":\"structIUiIncentiveDataProviderV3.UserReserveIncentiveData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AaveUiIncentiveDataProviderV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveUiIncentiveDataProviderV3MetaData.ABI instead.
var AaveUiIncentiveDataProviderV3ABI = AaveUiIncentiveDataProviderV3MetaData.ABI

// AaveUiIncentiveDataProviderV3 is an auto generated Go binding around an Ethereum contract.
type AaveUiIncentiveDataProviderV3 struct {
	AaveUiIncentiveDataProviderV3Caller     // Read-only binding to the contract
	AaveUiIncentiveDataProviderV3Transactor // Write-only binding to the contract
	AaveUiIncentiveDataProviderV3Filterer   // Log filterer for contract events
}

// AaveUiIncentiveDataProviderV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type AaveUiIncentiveDataProviderV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiIncentiveDataProviderV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveUiIncentiveDataProviderV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiIncentiveDataProviderV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveUiIncentiveDataProviderV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveUiIncentiveDataProviderV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveUiIncentiveDataProviderV3Session struct {
	Contract     *AaveUiIncentiveDataProviderV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AaveUiIncentiveDataProviderV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveUiIncentiveDataProviderV3CallerSession struct {
	Contract *AaveUiIncentiveDataProviderV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// AaveUiIncentiveDataProviderV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveUiIncentiveDataProviderV3TransactorSession struct {
	Contract     *AaveUiIncentiveDataProviderV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// AaveUiIncentiveDataProviderV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type AaveUiIncentiveDataProviderV3Raw struct {
	Contract *AaveUiIncentiveDataProviderV3 // Generic contract binding to access the raw methods on
}

// AaveUiIncentiveDataProviderV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveUiIncentiveDataProviderV3CallerRaw struct {
	Contract *AaveUiIncentiveDataProviderV3Caller // Generic read-only contract binding to access the raw methods on
}

// AaveUiIncentiveDataProviderV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveUiIncentiveDataProviderV3TransactorRaw struct {
	Contract *AaveUiIncentiveDataProviderV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveUiIncentiveDataProviderV3 creates a new instance of AaveUiIncentiveDataProviderV3, bound to a specific deployed contract.
func NewAaveUiIncentiveDataProviderV3(address common.Address, backend bind.ContractBackend) (*AaveUiIncentiveDataProviderV3, error) {
	contract, err := bindAaveUiIncentiveDataProviderV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveUiIncentiveDataProviderV3{AaveUiIncentiveDataProviderV3Caller: AaveUiIncentiveDataProviderV3Caller{contract: contract}, AaveUiIncentiveDataProviderV3Transactor: AaveUiIncentiveDataProviderV3Transactor{contract: contract}, AaveUiIncentiveDataProviderV3Filterer: AaveUiIncentiveDataProviderV3Filterer{contract: contract}}, nil
}

// NewAaveUiIncentiveDataProviderV3Caller creates a new read-only instance of AaveUiIncentiveDataProviderV3, bound to a specific deployed contract.
func NewAaveUiIncentiveDataProviderV3Caller(address common.Address, caller bind.ContractCaller) (*AaveUiIncentiveDataProviderV3Caller, error) {
	contract, err := bindAaveUiIncentiveDataProviderV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveUiIncentiveDataProviderV3Caller{contract: contract}, nil
}

// NewAaveUiIncentiveDataProviderV3Transactor creates a new write-only instance of AaveUiIncentiveDataProviderV3, bound to a specific deployed contract.
func NewAaveUiIncentiveDataProviderV3Transactor(address common.Address, transactor bind.ContractTransactor) (*AaveUiIncentiveDataProviderV3Transactor, error) {
	contract, err := bindAaveUiIncentiveDataProviderV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveUiIncentiveDataProviderV3Transactor{contract: contract}, nil
}

// NewAaveUiIncentiveDataProviderV3Filterer creates a new log filterer instance of AaveUiIncentiveDataProviderV3, bound to a specific deployed contract.
func NewAaveUiIncentiveDataProviderV3Filterer(address common.Address, filterer bind.ContractFilterer) (*AaveUiIncentiveDataProviderV3Filterer, error) {
	contract, err := bindAaveUiIncentiveDataProviderV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveUiIncentiveDataProviderV3Filterer{contract: contract}, nil
}

// bindAaveUiIncentiveDataProviderV3 binds a generic wrapper to an already deployed contract.
func bindAaveUiIncentiveDataProviderV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveUiIncentiveDataProviderV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveUiIncentiveDataProviderV3.Contract.AaveUiIncentiveDataProviderV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveUiIncentiveDataProviderV3.Contract.AaveUiIncentiveDataProviderV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveUiIncentiveDataProviderV3.Contract.AaveUiIncentiveDataProviderV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveUiIncentiveDataProviderV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveUiIncentiveDataProviderV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveUiIncentiveDataProviderV3.Contract.contract.Transact(opts, method, params...)
}

// GetFullReservesIncentiveData is a free data retrieval call binding the contract method 0x47637536.
//
// Solidity: function getFullReservesIncentiveData(address provider, address user) view returns((address,(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]))[], (address,(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]))[])
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3Caller) GetFullReservesIncentiveData(opts *bind.CallOpts, provider common.Address, user common.Address) ([]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData, []IUiIncentiveDataProviderV3UserReserveIncentiveData, error) {
	var out []interface{}
	err := _AaveUiIncentiveDataProviderV3.contract.Call(opts, &out, "getFullReservesIncentiveData", provider, user)

	if err != nil {
		return *new([]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData), *new([]IUiIncentiveDataProviderV3UserReserveIncentiveData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData)).(*[]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData)
	out1 := *abi.ConvertType(out[1], new([]IUiIncentiveDataProviderV3UserReserveIncentiveData)).(*[]IUiIncentiveDataProviderV3UserReserveIncentiveData)

	return out0, out1, err

}

// GetFullReservesIncentiveData is a free data retrieval call binding the contract method 0x47637536.
//
// Solidity: function getFullReservesIncentiveData(address provider, address user) view returns((address,(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]))[], (address,(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]))[])
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3Session) GetFullReservesIncentiveData(provider common.Address, user common.Address) ([]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData, []IUiIncentiveDataProviderV3UserReserveIncentiveData, error) {
	return _AaveUiIncentiveDataProviderV3.Contract.GetFullReservesIncentiveData(&_AaveUiIncentiveDataProviderV3.CallOpts, provider, user)
}

// GetFullReservesIncentiveData is a free data retrieval call binding the contract method 0x47637536.
//
// Solidity: function getFullReservesIncentiveData(address provider, address user) view returns((address,(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]))[], (address,(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]))[])
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3CallerSession) GetFullReservesIncentiveData(provider common.Address, user common.Address) ([]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData, []IUiIncentiveDataProviderV3UserReserveIncentiveData, error) {
	return _AaveUiIncentiveDataProviderV3.Contract.GetFullReservesIncentiveData(&_AaveUiIncentiveDataProviderV3.CallOpts, provider, user)
}

// GetReservesIncentivesData is a free data retrieval call binding the contract method 0x976fafc5.
//
// Solidity: function getReservesIncentivesData(address provider) view returns((address,(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]))[])
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3Caller) GetReservesIncentivesData(opts *bind.CallOpts, provider common.Address) ([]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData, error) {
	var out []interface{}
	err := _AaveUiIncentiveDataProviderV3.contract.Call(opts, &out, "getReservesIncentivesData", provider)

	if err != nil {
		return *new([]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData)).(*[]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData)

	return out0, err

}

// GetReservesIncentivesData is a free data retrieval call binding the contract method 0x976fafc5.
//
// Solidity: function getReservesIncentivesData(address provider) view returns((address,(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]))[])
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3Session) GetReservesIncentivesData(provider common.Address) ([]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData, error) {
	return _AaveUiIncentiveDataProviderV3.Contract.GetReservesIncentivesData(&_AaveUiIncentiveDataProviderV3.CallOpts, provider)
}

// GetReservesIncentivesData is a free data retrieval call binding the contract method 0x976fafc5.
//
// Solidity: function getReservesIncentivesData(address provider) view returns((address,(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,uint256,uint256,int256,uint8,uint8,uint8)[]))[])
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3CallerSession) GetReservesIncentivesData(provider common.Address) ([]IUiIncentiveDataProviderV3AggregatedReserveIncentiveData, error) {
	return _AaveUiIncentiveDataProviderV3.Contract.GetReservesIncentivesData(&_AaveUiIncentiveDataProviderV3.CallOpts, provider)
}

// GetUserReservesIncentivesData is a free data retrieval call binding the contract method 0x799bdcf5.
//
// Solidity: function getUserReservesIncentivesData(address provider, address user) view returns((address,(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]))[])
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3Caller) GetUserReservesIncentivesData(opts *bind.CallOpts, provider common.Address, user common.Address) ([]IUiIncentiveDataProviderV3UserReserveIncentiveData, error) {
	var out []interface{}
	err := _AaveUiIncentiveDataProviderV3.contract.Call(opts, &out, "getUserReservesIncentivesData", provider, user)

	if err != nil {
		return *new([]IUiIncentiveDataProviderV3UserReserveIncentiveData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUiIncentiveDataProviderV3UserReserveIncentiveData)).(*[]IUiIncentiveDataProviderV3UserReserveIncentiveData)

	return out0, err

}

// GetUserReservesIncentivesData is a free data retrieval call binding the contract method 0x799bdcf5.
//
// Solidity: function getUserReservesIncentivesData(address provider, address user) view returns((address,(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]))[])
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3Session) GetUserReservesIncentivesData(provider common.Address, user common.Address) ([]IUiIncentiveDataProviderV3UserReserveIncentiveData, error) {
	return _AaveUiIncentiveDataProviderV3.Contract.GetUserReservesIncentivesData(&_AaveUiIncentiveDataProviderV3.CallOpts, provider, user)
}

// GetUserReservesIncentivesData is a free data retrieval call binding the contract method 0x799bdcf5.
//
// Solidity: function getUserReservesIncentivesData(address provider, address user) view returns((address,(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]),(address,address,(string,address,address,uint256,uint256,int256,uint8,uint8)[]))[])
func (_AaveUiIncentiveDataProviderV3 *AaveUiIncentiveDataProviderV3CallerSession) GetUserReservesIncentivesData(provider common.Address, user common.Address) ([]IUiIncentiveDataProviderV3UserReserveIncentiveData, error) {
	return _AaveUiIncentiveDataProviderV3.Contract.GetUserReservesIncentivesData(&_AaveUiIncentiveDataProviderV3.CallOpts, provider, user)
}
