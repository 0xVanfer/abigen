// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stapleController

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

// PoolDataTypesPoolActionInput is an auto generated low-level Go binding around an user-defined struct.
type PoolDataTypesPoolActionInput struct {
	Underlying common.Address
	Amount     *big.Int
	Receiver   common.Address
	Owner      common.Address
}

// PoolDataTypesPoolParams is an auto generated low-level Go binding around an user-defined struct.
type PoolDataTypesPoolParams struct {
	Id                 *big.Int
	Underlying         common.Address
	LpAddr             common.Address
	UnderlyingDecimals uint8
	Tl                 uint8
	RelatedVtps        []*big.Int
}

// SwapDataTypesSwapSingleInput is an auto generated low-level Go binding around an user-defined struct.
type SwapDataTypesSwapSingleInput struct {
	User      common.Address
	TokenIn   common.Address
	TokenOut  common.Address
	VtpID     *big.Int
	AmountIn  *big.Int
	PriceRate *big.Int
}

// SwapDataTypesSwapSingleProcess is an auto generated low-level Go binding around an user-defined struct.
type SwapDataTypesSwapSingleProcess struct {
	AmountIn     *big.Int
	FeeIn        *big.Int
	RealIn       *big.Int
	Pav          *big.Int
	SwapGet      *big.Int
	FeeOut       *big.Int
	EstiOut      *big.Int
	Punishment   *big.Int
	IsPunishment bool
	RealOut      *big.Int
}

// VtpDataTypesAllocateInput is an auto generated low-level Go binding around an user-defined struct.
type VtpDataTypesAllocateInput struct {
	User       common.Address
	Underlying common.Address
	VtpID      *big.Int
	Amount     *big.Int
	PriceRate  *big.Int
	MaxFeeRate *big.Int
}

// VtpDataTypesDeallocateInput is an auto generated low-level Go binding around an user-defined struct.
type VtpDataTypesDeallocateInput struct {
	User       common.Address
	Underlying common.Address
	VtpID      *big.Int
	Amount     *big.Int
	PriceRate  *big.Int
	MaxFeeRate *big.Int
}

// VtpDataTypesVtp is an auto generated low-level Go binding around an user-defined struct.
type VtpDataTypesVtp struct {
	Params VtpDataTypesVtpParams
	Status VtpDataTypesVtpStatus
	Token0 VtpDataTypesVtpToken
	Token1 VtpDataTypesVtpToken
}

// VtpDataTypesVtpParams is an auto generated low-level Go binding around an user-defined struct.
type VtpDataTypesVtpParams struct {
	Vl uint8
	Id *big.Int
	N  *big.Int
	P  *big.Int
}

// VtpDataTypesVtpStatus is an auto generated low-level Go binding around an user-defined struct.
type VtpDataTypesVtpStatus struct {
	Po *big.Int
	Pa *big.Int
}

// VtpDataTypesVtpToken is an auto generated low-level Go binding around an user-defined struct.
type VtpDataTypesVtpToken struct {
	Params VtpDataTypesVtpTokenParams
	Status VtpDataTypesVtpTokenStatus
}

// VtpDataTypesVtpTokenParams is an auto generated low-level Go binding around an user-defined struct.
type VtpDataTypesVtpTokenParams struct {
	Addr            common.Address
	LpAddr          common.Address
	Decimals        uint8
	Tl              uint8
	Id              *big.Int
	MaxAllocateRate *big.Int
	SwapFeeIn       *big.Int
	SwapFeeOut      *big.Int
	ProtocolFeeRate *big.Int
	AlrLowerBound   *big.Int
}

// VtpDataTypesVtpTokenStatus is an auto generated low-level Go binding around an user-defined struct.
type VtpDataTypesVtpTokenStatus struct {
	Allocated            *big.Int
	Liability            *big.Int
	Asset                *big.Int
	Alr                  *big.Int
	TotalFeeCollected    *big.Int
	FeeCollected         *big.Int
	FeeCollectedProtocol *big.Int
	TotalShares          *big.Int
	Paused               bool
}

// StapleControllerMetaData contains all meta data concerning the StapleController contract.
var StapleControllerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBound\",\"type\":\"uint256\"}],\"name\":\"UpdateAlrLowerBound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateDefaultTokenImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateIncentivesController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxAllocateRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newStatus\",\"type\":\"bool\"}],\"name\":\"UpdatePauseStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"UpdateProtocolFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateProxyAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creditLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"vls\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"credits\",\"type\":\"uint256[]\"}],\"name\":\"UpdateRiskInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateRouter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOut\",\"type\":\"uint256\"}],\"name\":\"UpdateSwapFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newTL\",\"type\":\"uint8\"}],\"name\":\"UpdateTL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"discounts\",\"type\":\"uint256[]\"}],\"name\":\"UpdateWhiteList\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeeRate\",\"type\":\"uint256\"}],\"internalType\":\"structVtpDataTypes.AllocateInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeeRate\",\"type\":\"uint256\"}],\"internalType\":\"structVtpDataTypes.DeallocateInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"deallocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structPoolDataTypes.PoolActionInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeeRate\",\"type\":\"uint256\"}],\"internalType\":\"structVtpDataTypes.AllocateInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"estimateAllocationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeeRate\",\"type\":\"uint256\"}],\"internalType\":\"structVtpDataTypes.DeallocateInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"estimateDeallocationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceRate\",\"type\":\"uint256\"}],\"internalType\":\"structSwapDataTypes.SwapSingleInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"estimateSwapSingleOut\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"realIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pav\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapGet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estiOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"punishment\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPunishment\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"realOut\",\"type\":\"uint256\"}],\"internalType\":\"structSwapDataTypes.SwapSingleProcess\",\"name\":\"process\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"getCreditCostOfVtp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreditTotalLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"getPairedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"getPoolParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lpAddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"underlyingDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tl\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"relatedVtps\",\"type\":\"uint256[]\"}],\"internalType\":\"structPoolDataTypes.PoolParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"getVtp\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"vl\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"p\",\"type\":\"uint256\"}],\"internalType\":\"structVtpDataTypes.VtpParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"po\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pa\",\"type\":\"uint256\"}],\"internalType\":\"structVtpDataTypes.VtpStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lpAddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tl\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAllocateRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFeeIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFeeOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alrLowerBound\",\"type\":\"uint256\"}],\"internalType\":\"structVtpDataTypes.VtpTokenParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liability\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFeeCollected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeCollected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeCollectedProtocol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"internalType\":\"structVtpDataTypes.VtpTokenStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"internalType\":\"structVtpDataTypes.VtpToken\",\"name\":\"token0\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lpAddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tl\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAllocateRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFeeIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFeeOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alrLowerBound\",\"type\":\"uint256\"}],\"internalType\":\"structVtpDataTypes.VtpTokenParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liability\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFeeCollected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeCollected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeCollectedProtocol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"internalType\":\"structVtpDataTypes.VtpTokenStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"internalType\":\"structVtpDataTypes.VtpToken\",\"name\":\"token1\",\"type\":\"tuple\"}],\"internalType\":\"structVtpDataTypes.Vtp\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"getVtpToken\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lpAddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tl\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAllocateRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFeeIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFeeOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alrLowerBound\",\"type\":\"uint256\"}],\"internalType\":\"structVtpDataTypes.VtpTokenParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liability\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFeeCollected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeCollected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeCollectedProtocol\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"internalType\":\"structVtpDataTypes.VtpTokenStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"internalType\":\"structVtpDataTypes.VtpToken\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"id2underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incentivesController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceRate\",\"type\":\"uint256\"}],\"internalType\":\"structSwapDataTypes.SwapSingleInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"realOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"userAllocatedVtps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"vtpIDs\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"userAllocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allocation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"userDeposition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userDiscount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vtpID\",\"type\":\"uint256\"}],\"name\":\"userMaxAllocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAllocation_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vtpID\",\"type\":\"uint256\"}],\"name\":\"userMaxDeallocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxDeallocation_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"userRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vtpID\",\"type\":\"uint256\"}],\"name\":\"userUnchargedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"userUnlockedUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unlocked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vtpCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structPoolDataTypes.PoolActionInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StapleControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use StapleControllerMetaData.ABI instead.
var StapleControllerABI = StapleControllerMetaData.ABI

// StapleController is an auto generated Go binding around an Ethereum contract.
type StapleController struct {
	StapleControllerCaller     // Read-only binding to the contract
	StapleControllerTransactor // Write-only binding to the contract
	StapleControllerFilterer   // Log filterer for contract events
}

// StapleControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StapleControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StapleControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StapleControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StapleControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StapleControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StapleControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StapleControllerSession struct {
	Contract     *StapleController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StapleControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StapleControllerCallerSession struct {
	Contract *StapleControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StapleControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StapleControllerTransactorSession struct {
	Contract     *StapleControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StapleControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StapleControllerRaw struct {
	Contract *StapleController // Generic contract binding to access the raw methods on
}

// StapleControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StapleControllerCallerRaw struct {
	Contract *StapleControllerCaller // Generic read-only contract binding to access the raw methods on
}

// StapleControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StapleControllerTransactorRaw struct {
	Contract *StapleControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStapleController creates a new instance of StapleController, bound to a specific deployed contract.
func NewStapleController(address common.Address, backend bind.ContractBackend) (*StapleController, error) {
	contract, err := bindStapleController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StapleController{StapleControllerCaller: StapleControllerCaller{contract: contract}, StapleControllerTransactor: StapleControllerTransactor{contract: contract}, StapleControllerFilterer: StapleControllerFilterer{contract: contract}}, nil
}

// NewStapleControllerCaller creates a new read-only instance of StapleController, bound to a specific deployed contract.
func NewStapleControllerCaller(address common.Address, caller bind.ContractCaller) (*StapleControllerCaller, error) {
	contract, err := bindStapleController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StapleControllerCaller{contract: contract}, nil
}

// NewStapleControllerTransactor creates a new write-only instance of StapleController, bound to a specific deployed contract.
func NewStapleControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*StapleControllerTransactor, error) {
	contract, err := bindStapleController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StapleControllerTransactor{contract: contract}, nil
}

// NewStapleControllerFilterer creates a new log filterer instance of StapleController, bound to a specific deployed contract.
func NewStapleControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*StapleControllerFilterer, error) {
	contract, err := bindStapleController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StapleControllerFilterer{contract: contract}, nil
}

// bindStapleController binds a generic wrapper to an already deployed contract.
func bindStapleController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StapleControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StapleController *StapleControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StapleController.Contract.StapleControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StapleController *StapleControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StapleController.Contract.StapleControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StapleController *StapleControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StapleController.Contract.StapleControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StapleController *StapleControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StapleController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StapleController *StapleControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StapleController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StapleController *StapleControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StapleController.Contract.contract.Transact(opts, method, params...)
}

// DefaultImpl is a free data retrieval call binding the contract method 0x631ffbf3.
//
// Solidity: function defaultImpl() view returns(address)
func (_StapleController *StapleControllerCaller) DefaultImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "defaultImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultImpl is a free data retrieval call binding the contract method 0x631ffbf3.
//
// Solidity: function defaultImpl() view returns(address)
func (_StapleController *StapleControllerSession) DefaultImpl() (common.Address, error) {
	return _StapleController.Contract.DefaultImpl(&_StapleController.CallOpts)
}

// DefaultImpl is a free data retrieval call binding the contract method 0x631ffbf3.
//
// Solidity: function defaultImpl() view returns(address)
func (_StapleController *StapleControllerCallerSession) DefaultImpl() (common.Address, error) {
	return _StapleController.Contract.DefaultImpl(&_StapleController.CallOpts)
}

// EstimateAllocationFee is a free data retrieval call binding the contract method 0x6805d640.
//
// Solidity: function estimateAllocationFee((address,address,uint256,uint256,uint256,uint256) input) view returns(uint256 fee)
func (_StapleController *StapleControllerCaller) EstimateAllocationFee(opts *bind.CallOpts, input VtpDataTypesAllocateInput) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "estimateAllocationFee", input)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateAllocationFee is a free data retrieval call binding the contract method 0x6805d640.
//
// Solidity: function estimateAllocationFee((address,address,uint256,uint256,uint256,uint256) input) view returns(uint256 fee)
func (_StapleController *StapleControllerSession) EstimateAllocationFee(input VtpDataTypesAllocateInput) (*big.Int, error) {
	return _StapleController.Contract.EstimateAllocationFee(&_StapleController.CallOpts, input)
}

// EstimateAllocationFee is a free data retrieval call binding the contract method 0x6805d640.
//
// Solidity: function estimateAllocationFee((address,address,uint256,uint256,uint256,uint256) input) view returns(uint256 fee)
func (_StapleController *StapleControllerCallerSession) EstimateAllocationFee(input VtpDataTypesAllocateInput) (*big.Int, error) {
	return _StapleController.Contract.EstimateAllocationFee(&_StapleController.CallOpts, input)
}

// EstimateDeallocationFee is a free data retrieval call binding the contract method 0x258fc746.
//
// Solidity: function estimateDeallocationFee((address,address,uint256,uint256,uint256,uint256) input) view returns(uint256 fee)
func (_StapleController *StapleControllerCaller) EstimateDeallocationFee(opts *bind.CallOpts, input VtpDataTypesDeallocateInput) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "estimateDeallocationFee", input)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateDeallocationFee is a free data retrieval call binding the contract method 0x258fc746.
//
// Solidity: function estimateDeallocationFee((address,address,uint256,uint256,uint256,uint256) input) view returns(uint256 fee)
func (_StapleController *StapleControllerSession) EstimateDeallocationFee(input VtpDataTypesDeallocateInput) (*big.Int, error) {
	return _StapleController.Contract.EstimateDeallocationFee(&_StapleController.CallOpts, input)
}

// EstimateDeallocationFee is a free data retrieval call binding the contract method 0x258fc746.
//
// Solidity: function estimateDeallocationFee((address,address,uint256,uint256,uint256,uint256) input) view returns(uint256 fee)
func (_StapleController *StapleControllerCallerSession) EstimateDeallocationFee(input VtpDataTypesDeallocateInput) (*big.Int, error) {
	return _StapleController.Contract.EstimateDeallocationFee(&_StapleController.CallOpts, input)
}

// EstimateSwapSingleOut is a free data retrieval call binding the contract method 0x1f5a7298.
//
// Solidity: function estimateSwapSingleOut((address,address,address,uint256,uint256,uint256) input) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256) process)
func (_StapleController *StapleControllerCaller) EstimateSwapSingleOut(opts *bind.CallOpts, input SwapDataTypesSwapSingleInput) (SwapDataTypesSwapSingleProcess, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "estimateSwapSingleOut", input)

	if err != nil {
		return *new(SwapDataTypesSwapSingleProcess), err
	}

	out0 := *abi.ConvertType(out[0], new(SwapDataTypesSwapSingleProcess)).(*SwapDataTypesSwapSingleProcess)

	return out0, err

}

// EstimateSwapSingleOut is a free data retrieval call binding the contract method 0x1f5a7298.
//
// Solidity: function estimateSwapSingleOut((address,address,address,uint256,uint256,uint256) input) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256) process)
func (_StapleController *StapleControllerSession) EstimateSwapSingleOut(input SwapDataTypesSwapSingleInput) (SwapDataTypesSwapSingleProcess, error) {
	return _StapleController.Contract.EstimateSwapSingleOut(&_StapleController.CallOpts, input)
}

// EstimateSwapSingleOut is a free data retrieval call binding the contract method 0x1f5a7298.
//
// Solidity: function estimateSwapSingleOut((address,address,address,uint256,uint256,uint256) input) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256) process)
func (_StapleController *StapleControllerCallerSession) EstimateSwapSingleOut(input SwapDataTypesSwapSingleInput) (SwapDataTypesSwapSingleProcess, error) {
	return _StapleController.Contract.EstimateSwapSingleOut(&_StapleController.CallOpts, input)
}

// GetCreditCostOfVtp is a free data retrieval call binding the contract method 0x586a6d06.
//
// Solidity: function getCreditCostOfVtp(uint256 vtpID) view returns(uint256)
func (_StapleController *StapleControllerCaller) GetCreditCostOfVtp(opts *bind.CallOpts, vtpID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "getCreditCostOfVtp", vtpID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCreditCostOfVtp is a free data retrieval call binding the contract method 0x586a6d06.
//
// Solidity: function getCreditCostOfVtp(uint256 vtpID) view returns(uint256)
func (_StapleController *StapleControllerSession) GetCreditCostOfVtp(vtpID *big.Int) (*big.Int, error) {
	return _StapleController.Contract.GetCreditCostOfVtp(&_StapleController.CallOpts, vtpID)
}

// GetCreditCostOfVtp is a free data retrieval call binding the contract method 0x586a6d06.
//
// Solidity: function getCreditCostOfVtp(uint256 vtpID) view returns(uint256)
func (_StapleController *StapleControllerCallerSession) GetCreditCostOfVtp(vtpID *big.Int) (*big.Int, error) {
	return _StapleController.Contract.GetCreditCostOfVtp(&_StapleController.CallOpts, vtpID)
}

// GetCreditTotalLimit is a free data retrieval call binding the contract method 0x7ed605dc.
//
// Solidity: function getCreditTotalLimit() view returns(uint256)
func (_StapleController *StapleControllerCaller) GetCreditTotalLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "getCreditTotalLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCreditTotalLimit is a free data retrieval call binding the contract method 0x7ed605dc.
//
// Solidity: function getCreditTotalLimit() view returns(uint256)
func (_StapleController *StapleControllerSession) GetCreditTotalLimit() (*big.Int, error) {
	return _StapleController.Contract.GetCreditTotalLimit(&_StapleController.CallOpts)
}

// GetCreditTotalLimit is a free data retrieval call binding the contract method 0x7ed605dc.
//
// Solidity: function getCreditTotalLimit() view returns(uint256)
func (_StapleController *StapleControllerCallerSession) GetCreditTotalLimit() (*big.Int, error) {
	return _StapleController.Contract.GetCreditTotalLimit(&_StapleController.CallOpts)
}

// GetPairedToken is a free data retrieval call binding the contract method 0xcb760918.
//
// Solidity: function getPairedToken(address underlying, uint256 vtpID) view returns(address)
func (_StapleController *StapleControllerCaller) GetPairedToken(opts *bind.CallOpts, underlying common.Address, vtpID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "getPairedToken", underlying, vtpID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPairedToken is a free data retrieval call binding the contract method 0xcb760918.
//
// Solidity: function getPairedToken(address underlying, uint256 vtpID) view returns(address)
func (_StapleController *StapleControllerSession) GetPairedToken(underlying common.Address, vtpID *big.Int) (common.Address, error) {
	return _StapleController.Contract.GetPairedToken(&_StapleController.CallOpts, underlying, vtpID)
}

// GetPairedToken is a free data retrieval call binding the contract method 0xcb760918.
//
// Solidity: function getPairedToken(address underlying, uint256 vtpID) view returns(address)
func (_StapleController *StapleControllerCallerSession) GetPairedToken(underlying common.Address, vtpID *big.Int) (common.Address, error) {
	return _StapleController.Contract.GetPairedToken(&_StapleController.CallOpts, underlying, vtpID)
}

// GetPoolParams is a free data retrieval call binding the contract method 0xa2b78ff4.
//
// Solidity: function getPoolParams(address underlying) view returns((uint256,address,address,uint8,uint8,uint256[]))
func (_StapleController *StapleControllerCaller) GetPoolParams(opts *bind.CallOpts, underlying common.Address) (PoolDataTypesPoolParams, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "getPoolParams", underlying)

	if err != nil {
		return *new(PoolDataTypesPoolParams), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolDataTypesPoolParams)).(*PoolDataTypesPoolParams)

	return out0, err

}

// GetPoolParams is a free data retrieval call binding the contract method 0xa2b78ff4.
//
// Solidity: function getPoolParams(address underlying) view returns((uint256,address,address,uint8,uint8,uint256[]))
func (_StapleController *StapleControllerSession) GetPoolParams(underlying common.Address) (PoolDataTypesPoolParams, error) {
	return _StapleController.Contract.GetPoolParams(&_StapleController.CallOpts, underlying)
}

// GetPoolParams is a free data retrieval call binding the contract method 0xa2b78ff4.
//
// Solidity: function getPoolParams(address underlying) view returns((uint256,address,address,uint8,uint8,uint256[]))
func (_StapleController *StapleControllerCallerSession) GetPoolParams(underlying common.Address) (PoolDataTypesPoolParams, error) {
	return _StapleController.Contract.GetPoolParams(&_StapleController.CallOpts, underlying)
}

// GetVtp is a free data retrieval call binding the contract method 0xa01bdbef.
//
// Solidity: function getVtp(uint256 vtpID) view returns(((uint8,uint256,uint256,uint256),(uint256,uint256),((address,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)),((address,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))))
func (_StapleController *StapleControllerCaller) GetVtp(opts *bind.CallOpts, vtpID *big.Int) (VtpDataTypesVtp, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "getVtp", vtpID)

	if err != nil {
		return *new(VtpDataTypesVtp), err
	}

	out0 := *abi.ConvertType(out[0], new(VtpDataTypesVtp)).(*VtpDataTypesVtp)

	return out0, err

}

// GetVtp is a free data retrieval call binding the contract method 0xa01bdbef.
//
// Solidity: function getVtp(uint256 vtpID) view returns(((uint8,uint256,uint256,uint256),(uint256,uint256),((address,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)),((address,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))))
func (_StapleController *StapleControllerSession) GetVtp(vtpID *big.Int) (VtpDataTypesVtp, error) {
	return _StapleController.Contract.GetVtp(&_StapleController.CallOpts, vtpID)
}

// GetVtp is a free data retrieval call binding the contract method 0xa01bdbef.
//
// Solidity: function getVtp(uint256 vtpID) view returns(((uint8,uint256,uint256,uint256),(uint256,uint256),((address,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)),((address,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))))
func (_StapleController *StapleControllerCallerSession) GetVtp(vtpID *big.Int) (VtpDataTypesVtp, error) {
	return _StapleController.Contract.GetVtp(&_StapleController.CallOpts, vtpID)
}

// GetVtpToken is a free data retrieval call binding the contract method 0x533f927c.
//
// Solidity: function getVtpToken(address underlying, uint256 vtpID) view returns(((address,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)))
func (_StapleController *StapleControllerCaller) GetVtpToken(opts *bind.CallOpts, underlying common.Address, vtpID *big.Int) (VtpDataTypesVtpToken, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "getVtpToken", underlying, vtpID)

	if err != nil {
		return *new(VtpDataTypesVtpToken), err
	}

	out0 := *abi.ConvertType(out[0], new(VtpDataTypesVtpToken)).(*VtpDataTypesVtpToken)

	return out0, err

}

// GetVtpToken is a free data retrieval call binding the contract method 0x533f927c.
//
// Solidity: function getVtpToken(address underlying, uint256 vtpID) view returns(((address,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)))
func (_StapleController *StapleControllerSession) GetVtpToken(underlying common.Address, vtpID *big.Int) (VtpDataTypesVtpToken, error) {
	return _StapleController.Contract.GetVtpToken(&_StapleController.CallOpts, underlying, vtpID)
}

// GetVtpToken is a free data retrieval call binding the contract method 0x533f927c.
//
// Solidity: function getVtpToken(address underlying, uint256 vtpID) view returns(((address,address,uint8,uint8,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)))
func (_StapleController *StapleControllerCallerSession) GetVtpToken(underlying common.Address, vtpID *big.Int) (VtpDataTypesVtpToken, error) {
	return _StapleController.Contract.GetVtpToken(&_StapleController.CallOpts, underlying, vtpID)
}

// Id2underlying is a free data retrieval call binding the contract method 0x9dfd9826.
//
// Solidity: function id2underlying(uint256 id) view returns(address underlying)
func (_StapleController *StapleControllerCaller) Id2underlying(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "id2underlying", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Id2underlying is a free data retrieval call binding the contract method 0x9dfd9826.
//
// Solidity: function id2underlying(uint256 id) view returns(address underlying)
func (_StapleController *StapleControllerSession) Id2underlying(id *big.Int) (common.Address, error) {
	return _StapleController.Contract.Id2underlying(&_StapleController.CallOpts, id)
}

// Id2underlying is a free data retrieval call binding the contract method 0x9dfd9826.
//
// Solidity: function id2underlying(uint256 id) view returns(address underlying)
func (_StapleController *StapleControllerCallerSession) Id2underlying(id *big.Int) (common.Address, error) {
	return _StapleController.Contract.Id2underlying(&_StapleController.CallOpts, id)
}

// IncentivesController is a free data retrieval call binding the contract method 0xaf1df255.
//
// Solidity: function incentivesController() view returns(address)
func (_StapleController *StapleControllerCaller) IncentivesController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "incentivesController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IncentivesController is a free data retrieval call binding the contract method 0xaf1df255.
//
// Solidity: function incentivesController() view returns(address)
func (_StapleController *StapleControllerSession) IncentivesController() (common.Address, error) {
	return _StapleController.Contract.IncentivesController(&_StapleController.CallOpts)
}

// IncentivesController is a free data retrieval call binding the contract method 0xaf1df255.
//
// Solidity: function incentivesController() view returns(address)
func (_StapleController *StapleControllerCallerSession) IncentivesController() (common.Address, error) {
	return _StapleController.Contract.IncentivesController(&_StapleController.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_StapleController *StapleControllerCaller) ProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "proxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_StapleController *StapleControllerSession) ProxyAdmin() (common.Address, error) {
	return _StapleController.Contract.ProxyAdmin(&_StapleController.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_StapleController *StapleControllerCallerSession) ProxyAdmin() (common.Address, error) {
	return _StapleController.Contract.ProxyAdmin(&_StapleController.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_StapleController *StapleControllerCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_StapleController *StapleControllerSession) Router() (common.Address, error) {
	return _StapleController.Contract.Router(&_StapleController.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_StapleController *StapleControllerCallerSession) Router() (common.Address, error) {
	return _StapleController.Contract.Router(&_StapleController.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() view returns(uint256)
func (_StapleController *StapleControllerCaller) TokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "tokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() view returns(uint256)
func (_StapleController *StapleControllerSession) TokenCount() (*big.Int, error) {
	return _StapleController.Contract.TokenCount(&_StapleController.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() view returns(uint256)
func (_StapleController *StapleControllerCallerSession) TokenCount() (*big.Int, error) {
	return _StapleController.Contract.TokenCount(&_StapleController.CallOpts)
}

// UserAllocatedVtps is a free data retrieval call binding the contract method 0xd0366519.
//
// Solidity: function userAllocatedVtps(address user, address underlying) view returns(uint256[] vtpIDs)
func (_StapleController *StapleControllerCaller) UserAllocatedVtps(opts *bind.CallOpts, user common.Address, underlying common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "userAllocatedVtps", user, underlying)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// UserAllocatedVtps is a free data retrieval call binding the contract method 0xd0366519.
//
// Solidity: function userAllocatedVtps(address user, address underlying) view returns(uint256[] vtpIDs)
func (_StapleController *StapleControllerSession) UserAllocatedVtps(user common.Address, underlying common.Address) ([]*big.Int, error) {
	return _StapleController.Contract.UserAllocatedVtps(&_StapleController.CallOpts, user, underlying)
}

// UserAllocatedVtps is a free data retrieval call binding the contract method 0xd0366519.
//
// Solidity: function userAllocatedVtps(address user, address underlying) view returns(uint256[] vtpIDs)
func (_StapleController *StapleControllerCallerSession) UserAllocatedVtps(user common.Address, underlying common.Address) ([]*big.Int, error) {
	return _StapleController.Contract.UserAllocatedVtps(&_StapleController.CallOpts, user, underlying)
}

// UserAllocation is a free data retrieval call binding the contract method 0xe7ea0458.
//
// Solidity: function userAllocation(address user, address underlying, uint256 vtpID) view returns(uint256 allocation, uint256 shares)
func (_StapleController *StapleControllerCaller) UserAllocation(opts *bind.CallOpts, user common.Address, underlying common.Address, vtpID *big.Int) (struct {
	Allocation *big.Int
	Shares     *big.Int
}, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "userAllocation", user, underlying, vtpID)

	outstruct := new(struct {
		Allocation *big.Int
		Shares     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Allocation = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Shares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserAllocation is a free data retrieval call binding the contract method 0xe7ea0458.
//
// Solidity: function userAllocation(address user, address underlying, uint256 vtpID) view returns(uint256 allocation, uint256 shares)
func (_StapleController *StapleControllerSession) UserAllocation(user common.Address, underlying common.Address, vtpID *big.Int) (struct {
	Allocation *big.Int
	Shares     *big.Int
}, error) {
	return _StapleController.Contract.UserAllocation(&_StapleController.CallOpts, user, underlying, vtpID)
}

// UserAllocation is a free data retrieval call binding the contract method 0xe7ea0458.
//
// Solidity: function userAllocation(address user, address underlying, uint256 vtpID) view returns(uint256 allocation, uint256 shares)
func (_StapleController *StapleControllerCallerSession) UserAllocation(user common.Address, underlying common.Address, vtpID *big.Int) (struct {
	Allocation *big.Int
	Shares     *big.Int
}, error) {
	return _StapleController.Contract.UserAllocation(&_StapleController.CallOpts, user, underlying, vtpID)
}

// UserDeposition is a free data retrieval call binding the contract method 0x37a70556.
//
// Solidity: function userDeposition(address user, address underlying) view returns(uint256 amount)
func (_StapleController *StapleControllerCaller) UserDeposition(opts *bind.CallOpts, user common.Address, underlying common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "userDeposition", user, underlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserDeposition is a free data retrieval call binding the contract method 0x37a70556.
//
// Solidity: function userDeposition(address user, address underlying) view returns(uint256 amount)
func (_StapleController *StapleControllerSession) UserDeposition(user common.Address, underlying common.Address) (*big.Int, error) {
	return _StapleController.Contract.UserDeposition(&_StapleController.CallOpts, user, underlying)
}

// UserDeposition is a free data retrieval call binding the contract method 0x37a70556.
//
// Solidity: function userDeposition(address user, address underlying) view returns(uint256 amount)
func (_StapleController *StapleControllerCallerSession) UserDeposition(user common.Address, underlying common.Address) (*big.Int, error) {
	return _StapleController.Contract.UserDeposition(&_StapleController.CallOpts, user, underlying)
}

// UserDiscount is a free data retrieval call binding the contract method 0x2ceb46e2.
//
// Solidity: function userDiscount(address user) view returns(uint256 discount)
func (_StapleController *StapleControllerCaller) UserDiscount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "userDiscount", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserDiscount is a free data retrieval call binding the contract method 0x2ceb46e2.
//
// Solidity: function userDiscount(address user) view returns(uint256 discount)
func (_StapleController *StapleControllerSession) UserDiscount(user common.Address) (*big.Int, error) {
	return _StapleController.Contract.UserDiscount(&_StapleController.CallOpts, user)
}

// UserDiscount is a free data retrieval call binding the contract method 0x2ceb46e2.
//
// Solidity: function userDiscount(address user) view returns(uint256 discount)
func (_StapleController *StapleControllerCallerSession) UserDiscount(user common.Address) (*big.Int, error) {
	return _StapleController.Contract.UserDiscount(&_StapleController.CallOpts, user)
}

// UserMaxAllocation is a free data retrieval call binding the contract method 0x90602940.
//
// Solidity: function userMaxAllocation(address _user, address _underlying, uint256 _vtpID) view returns(uint256 maxAllocation_)
func (_StapleController *StapleControllerCaller) UserMaxAllocation(opts *bind.CallOpts, _user common.Address, _underlying common.Address, _vtpID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "userMaxAllocation", _user, _underlying, _vtpID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMaxAllocation is a free data retrieval call binding the contract method 0x90602940.
//
// Solidity: function userMaxAllocation(address _user, address _underlying, uint256 _vtpID) view returns(uint256 maxAllocation_)
func (_StapleController *StapleControllerSession) UserMaxAllocation(_user common.Address, _underlying common.Address, _vtpID *big.Int) (*big.Int, error) {
	return _StapleController.Contract.UserMaxAllocation(&_StapleController.CallOpts, _user, _underlying, _vtpID)
}

// UserMaxAllocation is a free data retrieval call binding the contract method 0x90602940.
//
// Solidity: function userMaxAllocation(address _user, address _underlying, uint256 _vtpID) view returns(uint256 maxAllocation_)
func (_StapleController *StapleControllerCallerSession) UserMaxAllocation(_user common.Address, _underlying common.Address, _vtpID *big.Int) (*big.Int, error) {
	return _StapleController.Contract.UserMaxAllocation(&_StapleController.CallOpts, _user, _underlying, _vtpID)
}

// UserMaxDeallocation is a free data retrieval call binding the contract method 0x51ca8936.
//
// Solidity: function userMaxDeallocation(address _user, address _underlying, uint256 _vtpID) view returns(uint256 maxDeallocation_)
func (_StapleController *StapleControllerCaller) UserMaxDeallocation(opts *bind.CallOpts, _user common.Address, _underlying common.Address, _vtpID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "userMaxDeallocation", _user, _underlying, _vtpID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMaxDeallocation is a free data retrieval call binding the contract method 0x51ca8936.
//
// Solidity: function userMaxDeallocation(address _user, address _underlying, uint256 _vtpID) view returns(uint256 maxDeallocation_)
func (_StapleController *StapleControllerSession) UserMaxDeallocation(_user common.Address, _underlying common.Address, _vtpID *big.Int) (*big.Int, error) {
	return _StapleController.Contract.UserMaxDeallocation(&_StapleController.CallOpts, _user, _underlying, _vtpID)
}

// UserMaxDeallocation is a free data retrieval call binding the contract method 0x51ca8936.
//
// Solidity: function userMaxDeallocation(address _user, address _underlying, uint256 _vtpID) view returns(uint256 maxDeallocation_)
func (_StapleController *StapleControllerCallerSession) UserMaxDeallocation(_user common.Address, _underlying common.Address, _vtpID *big.Int) (*big.Int, error) {
	return _StapleController.Contract.UserMaxDeallocation(&_StapleController.CallOpts, _user, _underlying, _vtpID)
}

// UserRewards is a free data retrieval call binding the contract method 0xc6e18b28.
//
// Solidity: function userRewards(address user, address underlying, uint256 vtpID) view returns(uint256 rewards)
func (_StapleController *StapleControllerCaller) UserRewards(opts *bind.CallOpts, user common.Address, underlying common.Address, vtpID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "userRewards", user, underlying, vtpID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewards is a free data retrieval call binding the contract method 0xc6e18b28.
//
// Solidity: function userRewards(address user, address underlying, uint256 vtpID) view returns(uint256 rewards)
func (_StapleController *StapleControllerSession) UserRewards(user common.Address, underlying common.Address, vtpID *big.Int) (*big.Int, error) {
	return _StapleController.Contract.UserRewards(&_StapleController.CallOpts, user, underlying, vtpID)
}

// UserRewards is a free data retrieval call binding the contract method 0xc6e18b28.
//
// Solidity: function userRewards(address user, address underlying, uint256 vtpID) view returns(uint256 rewards)
func (_StapleController *StapleControllerCallerSession) UserRewards(user common.Address, underlying common.Address, vtpID *big.Int) (*big.Int, error) {
	return _StapleController.Contract.UserRewards(&_StapleController.CallOpts, user, underlying, vtpID)
}

// UserUnchargedFee is a free data retrieval call binding the contract method 0xc5495459.
//
// Solidity: function userUnchargedFee(address user, address underlying, uint256 vtpID) view returns(uint256)
func (_StapleController *StapleControllerCaller) UserUnchargedFee(opts *bind.CallOpts, user common.Address, underlying common.Address, vtpID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "userUnchargedFee", user, underlying, vtpID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnchargedFee is a free data retrieval call binding the contract method 0xc5495459.
//
// Solidity: function userUnchargedFee(address user, address underlying, uint256 vtpID) view returns(uint256)
func (_StapleController *StapleControllerSession) UserUnchargedFee(user common.Address, underlying common.Address, vtpID *big.Int) (*big.Int, error) {
	return _StapleController.Contract.UserUnchargedFee(&_StapleController.CallOpts, user, underlying, vtpID)
}

// UserUnchargedFee is a free data retrieval call binding the contract method 0xc5495459.
//
// Solidity: function userUnchargedFee(address user, address underlying, uint256 vtpID) view returns(uint256)
func (_StapleController *StapleControllerCallerSession) UserUnchargedFee(user common.Address, underlying common.Address, vtpID *big.Int) (*big.Int, error) {
	return _StapleController.Contract.UserUnchargedFee(&_StapleController.CallOpts, user, underlying, vtpID)
}

// UserUnlockedUnderlying is a free data retrieval call binding the contract method 0x4969f69f.
//
// Solidity: function userUnlockedUnderlying(address user, address underlying) view returns(uint256 unlocked)
func (_StapleController *StapleControllerCaller) UserUnlockedUnderlying(opts *bind.CallOpts, user common.Address, underlying common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "userUnlockedUnderlying", user, underlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnlockedUnderlying is a free data retrieval call binding the contract method 0x4969f69f.
//
// Solidity: function userUnlockedUnderlying(address user, address underlying) view returns(uint256 unlocked)
func (_StapleController *StapleControllerSession) UserUnlockedUnderlying(user common.Address, underlying common.Address) (*big.Int, error) {
	return _StapleController.Contract.UserUnlockedUnderlying(&_StapleController.CallOpts, user, underlying)
}

// UserUnlockedUnderlying is a free data retrieval call binding the contract method 0x4969f69f.
//
// Solidity: function userUnlockedUnderlying(address user, address underlying) view returns(uint256 unlocked)
func (_StapleController *StapleControllerCallerSession) UserUnlockedUnderlying(user common.Address, underlying common.Address) (*big.Int, error) {
	return _StapleController.Contract.UserUnlockedUnderlying(&_StapleController.CallOpts, user, underlying)
}

// VtpCount is a free data retrieval call binding the contract method 0xdd4402ca.
//
// Solidity: function vtpCount() view returns(uint256)
func (_StapleController *StapleControllerCaller) VtpCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StapleController.contract.Call(opts, &out, "vtpCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VtpCount is a free data retrieval call binding the contract method 0xdd4402ca.
//
// Solidity: function vtpCount() view returns(uint256)
func (_StapleController *StapleControllerSession) VtpCount() (*big.Int, error) {
	return _StapleController.Contract.VtpCount(&_StapleController.CallOpts)
}

// VtpCount is a free data retrieval call binding the contract method 0xdd4402ca.
//
// Solidity: function vtpCount() view returns(uint256)
func (_StapleController *StapleControllerCallerSession) VtpCount() (*big.Int, error) {
	return _StapleController.Contract.VtpCount(&_StapleController.CallOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0x45948e47.
//
// Solidity: function allocate((address,address,uint256,uint256,uint256,uint256) input) returns()
func (_StapleController *StapleControllerTransactor) Allocate(opts *bind.TransactOpts, input VtpDataTypesAllocateInput) (*types.Transaction, error) {
	return _StapleController.contract.Transact(opts, "allocate", input)
}

// Allocate is a paid mutator transaction binding the contract method 0x45948e47.
//
// Solidity: function allocate((address,address,uint256,uint256,uint256,uint256) input) returns()
func (_StapleController *StapleControllerSession) Allocate(input VtpDataTypesAllocateInput) (*types.Transaction, error) {
	return _StapleController.Contract.Allocate(&_StapleController.TransactOpts, input)
}

// Allocate is a paid mutator transaction binding the contract method 0x45948e47.
//
// Solidity: function allocate((address,address,uint256,uint256,uint256,uint256) input) returns()
func (_StapleController *StapleControllerTransactorSession) Allocate(input VtpDataTypesAllocateInput) (*types.Transaction, error) {
	return _StapleController.Contract.Allocate(&_StapleController.TransactOpts, input)
}

// Deallocate is a paid mutator transaction binding the contract method 0xffeca21a.
//
// Solidity: function deallocate((address,address,uint256,uint256,uint256,uint256) input) returns()
func (_StapleController *StapleControllerTransactor) Deallocate(opts *bind.TransactOpts, input VtpDataTypesDeallocateInput) (*types.Transaction, error) {
	return _StapleController.contract.Transact(opts, "deallocate", input)
}

// Deallocate is a paid mutator transaction binding the contract method 0xffeca21a.
//
// Solidity: function deallocate((address,address,uint256,uint256,uint256,uint256) input) returns()
func (_StapleController *StapleControllerSession) Deallocate(input VtpDataTypesDeallocateInput) (*types.Transaction, error) {
	return _StapleController.Contract.Deallocate(&_StapleController.TransactOpts, input)
}

// Deallocate is a paid mutator transaction binding the contract method 0xffeca21a.
//
// Solidity: function deallocate((address,address,uint256,uint256,uint256,uint256) input) returns()
func (_StapleController *StapleControllerTransactorSession) Deallocate(input VtpDataTypesDeallocateInput) (*types.Transaction, error) {
	return _StapleController.Contract.Deallocate(&_StapleController.TransactOpts, input)
}

// Deposit is a paid mutator transaction binding the contract method 0xe3b08e79.
//
// Solidity: function deposit((address,uint256,address,address) input) returns(uint256 shares)
func (_StapleController *StapleControllerTransactor) Deposit(opts *bind.TransactOpts, input PoolDataTypesPoolActionInput) (*types.Transaction, error) {
	return _StapleController.contract.Transact(opts, "deposit", input)
}

// Deposit is a paid mutator transaction binding the contract method 0xe3b08e79.
//
// Solidity: function deposit((address,uint256,address,address) input) returns(uint256 shares)
func (_StapleController *StapleControllerSession) Deposit(input PoolDataTypesPoolActionInput) (*types.Transaction, error) {
	return _StapleController.Contract.Deposit(&_StapleController.TransactOpts, input)
}

// Deposit is a paid mutator transaction binding the contract method 0xe3b08e79.
//
// Solidity: function deposit((address,uint256,address,address) input) returns(uint256 shares)
func (_StapleController *StapleControllerTransactorSession) Deposit(input PoolDataTypesPoolActionInput) (*types.Transaction, error) {
	return _StapleController.Contract.Deposit(&_StapleController.TransactOpts, input)
}

// Swap is a paid mutator transaction binding the contract method 0x441a12de.
//
// Solidity: function swap((address,address,address,uint256,uint256,uint256) input) returns(uint256 realOut)
func (_StapleController *StapleControllerTransactor) Swap(opts *bind.TransactOpts, input SwapDataTypesSwapSingleInput) (*types.Transaction, error) {
	return _StapleController.contract.Transact(opts, "swap", input)
}

// Swap is a paid mutator transaction binding the contract method 0x441a12de.
//
// Solidity: function swap((address,address,address,uint256,uint256,uint256) input) returns(uint256 realOut)
func (_StapleController *StapleControllerSession) Swap(input SwapDataTypesSwapSingleInput) (*types.Transaction, error) {
	return _StapleController.Contract.Swap(&_StapleController.TransactOpts, input)
}

// Swap is a paid mutator transaction binding the contract method 0x441a12de.
//
// Solidity: function swap((address,address,address,uint256,uint256,uint256) input) returns(uint256 realOut)
func (_StapleController *StapleControllerTransactorSession) Swap(input SwapDataTypesSwapSingleInput) (*types.Transaction, error) {
	return _StapleController.Contract.Swap(&_StapleController.TransactOpts, input)
}

// Withdraw is a paid mutator transaction binding the contract method 0x175b87d1.
//
// Solidity: function withdraw((address,uint256,address,address) input) returns(uint256 shares)
func (_StapleController *StapleControllerTransactor) Withdraw(opts *bind.TransactOpts, input PoolDataTypesPoolActionInput) (*types.Transaction, error) {
	return _StapleController.contract.Transact(opts, "withdraw", input)
}

// Withdraw is a paid mutator transaction binding the contract method 0x175b87d1.
//
// Solidity: function withdraw((address,uint256,address,address) input) returns(uint256 shares)
func (_StapleController *StapleControllerSession) Withdraw(input PoolDataTypesPoolActionInput) (*types.Transaction, error) {
	return _StapleController.Contract.Withdraw(&_StapleController.TransactOpts, input)
}

// Withdraw is a paid mutator transaction binding the contract method 0x175b87d1.
//
// Solidity: function withdraw((address,uint256,address,address) input) returns(uint256 shares)
func (_StapleController *StapleControllerTransactorSession) Withdraw(input PoolDataTypesPoolActionInput) (*types.Transaction, error) {
	return _StapleController.Contract.Withdraw(&_StapleController.TransactOpts, input)
}

// StapleControllerUpdateAlrLowerBoundIterator is returned from FilterUpdateAlrLowerBound and is used to iterate over the raw logs and unpacked data for UpdateAlrLowerBound events raised by the StapleController contract.
type StapleControllerUpdateAlrLowerBoundIterator struct {
	Event *StapleControllerUpdateAlrLowerBound // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdateAlrLowerBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdateAlrLowerBound)
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
		it.Event = new(StapleControllerUpdateAlrLowerBound)
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
func (it *StapleControllerUpdateAlrLowerBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdateAlrLowerBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdateAlrLowerBound represents a UpdateAlrLowerBound event raised by the StapleController contract.
type StapleControllerUpdateAlrLowerBound struct {
	Underlying common.Address
	VtpID      *big.Int
	NewBound   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAlrLowerBound is a free log retrieval operation binding the contract event 0x22231b437df11c3f9fa00a0987b876c1008026f0f0a2a8799ef8fc33aa4da170.
//
// Solidity: event UpdateAlrLowerBound(address indexed underlying, uint256 indexed vtpID, uint256 newBound)
func (_StapleController *StapleControllerFilterer) FilterUpdateAlrLowerBound(opts *bind.FilterOpts, underlying []common.Address, vtpID []*big.Int) (*StapleControllerUpdateAlrLowerBoundIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdateAlrLowerBound", underlyingRule, vtpIDRule)
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdateAlrLowerBoundIterator{contract: _StapleController.contract, event: "UpdateAlrLowerBound", logs: logs, sub: sub}, nil
}

// WatchUpdateAlrLowerBound is a free log subscription operation binding the contract event 0x22231b437df11c3f9fa00a0987b876c1008026f0f0a2a8799ef8fc33aa4da170.
//
// Solidity: event UpdateAlrLowerBound(address indexed underlying, uint256 indexed vtpID, uint256 newBound)
func (_StapleController *StapleControllerFilterer) WatchUpdateAlrLowerBound(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdateAlrLowerBound, underlying []common.Address, vtpID []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdateAlrLowerBound", underlyingRule, vtpIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdateAlrLowerBound)
				if err := _StapleController.contract.UnpackLog(event, "UpdateAlrLowerBound", log); err != nil {
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

// ParseUpdateAlrLowerBound is a log parse operation binding the contract event 0x22231b437df11c3f9fa00a0987b876c1008026f0f0a2a8799ef8fc33aa4da170.
//
// Solidity: event UpdateAlrLowerBound(address indexed underlying, uint256 indexed vtpID, uint256 newBound)
func (_StapleController *StapleControllerFilterer) ParseUpdateAlrLowerBound(log types.Log) (*StapleControllerUpdateAlrLowerBound, error) {
	event := new(StapleControllerUpdateAlrLowerBound)
	if err := _StapleController.contract.UnpackLog(event, "UpdateAlrLowerBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleControllerUpdateDefaultTokenImplementationIterator is returned from FilterUpdateDefaultTokenImplementation and is used to iterate over the raw logs and unpacked data for UpdateDefaultTokenImplementation events raised by the StapleController contract.
type StapleControllerUpdateDefaultTokenImplementationIterator struct {
	Event *StapleControllerUpdateDefaultTokenImplementation // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdateDefaultTokenImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdateDefaultTokenImplementation)
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
		it.Event = new(StapleControllerUpdateDefaultTokenImplementation)
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
func (it *StapleControllerUpdateDefaultTokenImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdateDefaultTokenImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdateDefaultTokenImplementation represents a UpdateDefaultTokenImplementation event raised by the StapleController contract.
type StapleControllerUpdateDefaultTokenImplementation struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultTokenImplementation is a free log retrieval operation binding the contract event 0x6d9b42355f02ed2247346532854e06c51ce3d30394c073147c736d9b6f1a3d37.
//
// Solidity: event UpdateDefaultTokenImplementation(address newAddress)
func (_StapleController *StapleControllerFilterer) FilterUpdateDefaultTokenImplementation(opts *bind.FilterOpts) (*StapleControllerUpdateDefaultTokenImplementationIterator, error) {

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdateDefaultTokenImplementation")
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdateDefaultTokenImplementationIterator{contract: _StapleController.contract, event: "UpdateDefaultTokenImplementation", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultTokenImplementation is a free log subscription operation binding the contract event 0x6d9b42355f02ed2247346532854e06c51ce3d30394c073147c736d9b6f1a3d37.
//
// Solidity: event UpdateDefaultTokenImplementation(address newAddress)
func (_StapleController *StapleControllerFilterer) WatchUpdateDefaultTokenImplementation(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdateDefaultTokenImplementation) (event.Subscription, error) {

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdateDefaultTokenImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdateDefaultTokenImplementation)
				if err := _StapleController.contract.UnpackLog(event, "UpdateDefaultTokenImplementation", log); err != nil {
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

// ParseUpdateDefaultTokenImplementation is a log parse operation binding the contract event 0x6d9b42355f02ed2247346532854e06c51ce3d30394c073147c736d9b6f1a3d37.
//
// Solidity: event UpdateDefaultTokenImplementation(address newAddress)
func (_StapleController *StapleControllerFilterer) ParseUpdateDefaultTokenImplementation(log types.Log) (*StapleControllerUpdateDefaultTokenImplementation, error) {
	event := new(StapleControllerUpdateDefaultTokenImplementation)
	if err := _StapleController.contract.UnpackLog(event, "UpdateDefaultTokenImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleControllerUpdateIncentivesControllerIterator is returned from FilterUpdateIncentivesController and is used to iterate over the raw logs and unpacked data for UpdateIncentivesController events raised by the StapleController contract.
type StapleControllerUpdateIncentivesControllerIterator struct {
	Event *StapleControllerUpdateIncentivesController // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdateIncentivesControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdateIncentivesController)
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
		it.Event = new(StapleControllerUpdateIncentivesController)
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
func (it *StapleControllerUpdateIncentivesControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdateIncentivesControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdateIncentivesController represents a UpdateIncentivesController event raised by the StapleController contract.
type StapleControllerUpdateIncentivesController struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateIncentivesController is a free log retrieval operation binding the contract event 0x5a1dc234bd6cbfeb202293a0c793a2418ed708afd888dae2dfe67b518fa31c2a.
//
// Solidity: event UpdateIncentivesController(address newAddress)
func (_StapleController *StapleControllerFilterer) FilterUpdateIncentivesController(opts *bind.FilterOpts) (*StapleControllerUpdateIncentivesControllerIterator, error) {

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdateIncentivesController")
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdateIncentivesControllerIterator{contract: _StapleController.contract, event: "UpdateIncentivesController", logs: logs, sub: sub}, nil
}

// WatchUpdateIncentivesController is a free log subscription operation binding the contract event 0x5a1dc234bd6cbfeb202293a0c793a2418ed708afd888dae2dfe67b518fa31c2a.
//
// Solidity: event UpdateIncentivesController(address newAddress)
func (_StapleController *StapleControllerFilterer) WatchUpdateIncentivesController(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdateIncentivesController) (event.Subscription, error) {

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdateIncentivesController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdateIncentivesController)
				if err := _StapleController.contract.UnpackLog(event, "UpdateIncentivesController", log); err != nil {
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

// ParseUpdateIncentivesController is a log parse operation binding the contract event 0x5a1dc234bd6cbfeb202293a0c793a2418ed708afd888dae2dfe67b518fa31c2a.
//
// Solidity: event UpdateIncentivesController(address newAddress)
func (_StapleController *StapleControllerFilterer) ParseUpdateIncentivesController(log types.Log) (*StapleControllerUpdateIncentivesController, error) {
	event := new(StapleControllerUpdateIncentivesController)
	if err := _StapleController.contract.UnpackLog(event, "UpdateIncentivesController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleControllerUpdateMaxAllocateRateIterator is returned from FilterUpdateMaxAllocateRate and is used to iterate over the raw logs and unpacked data for UpdateMaxAllocateRate events raised by the StapleController contract.
type StapleControllerUpdateMaxAllocateRateIterator struct {
	Event *StapleControllerUpdateMaxAllocateRate // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdateMaxAllocateRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdateMaxAllocateRate)
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
		it.Event = new(StapleControllerUpdateMaxAllocateRate)
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
func (it *StapleControllerUpdateMaxAllocateRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdateMaxAllocateRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdateMaxAllocateRate represents a UpdateMaxAllocateRate event raised by the StapleController contract.
type StapleControllerUpdateMaxAllocateRate struct {
	Underlying common.Address
	VtpID      *big.Int
	NewRate    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxAllocateRate is a free log retrieval operation binding the contract event 0xb4071e4c644e32b4e5ea39e374b1e446646e1ce8d713d64bb14a0e0d6b49e875.
//
// Solidity: event UpdateMaxAllocateRate(address indexed underlying, uint256 indexed vtpID, uint256 newRate)
func (_StapleController *StapleControllerFilterer) FilterUpdateMaxAllocateRate(opts *bind.FilterOpts, underlying []common.Address, vtpID []*big.Int) (*StapleControllerUpdateMaxAllocateRateIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdateMaxAllocateRate", underlyingRule, vtpIDRule)
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdateMaxAllocateRateIterator{contract: _StapleController.contract, event: "UpdateMaxAllocateRate", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxAllocateRate is a free log subscription operation binding the contract event 0xb4071e4c644e32b4e5ea39e374b1e446646e1ce8d713d64bb14a0e0d6b49e875.
//
// Solidity: event UpdateMaxAllocateRate(address indexed underlying, uint256 indexed vtpID, uint256 newRate)
func (_StapleController *StapleControllerFilterer) WatchUpdateMaxAllocateRate(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdateMaxAllocateRate, underlying []common.Address, vtpID []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdateMaxAllocateRate", underlyingRule, vtpIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdateMaxAllocateRate)
				if err := _StapleController.contract.UnpackLog(event, "UpdateMaxAllocateRate", log); err != nil {
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

// ParseUpdateMaxAllocateRate is a log parse operation binding the contract event 0xb4071e4c644e32b4e5ea39e374b1e446646e1ce8d713d64bb14a0e0d6b49e875.
//
// Solidity: event UpdateMaxAllocateRate(address indexed underlying, uint256 indexed vtpID, uint256 newRate)
func (_StapleController *StapleControllerFilterer) ParseUpdateMaxAllocateRate(log types.Log) (*StapleControllerUpdateMaxAllocateRate, error) {
	event := new(StapleControllerUpdateMaxAllocateRate)
	if err := _StapleController.contract.UnpackLog(event, "UpdateMaxAllocateRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleControllerUpdatePauseStatusIterator is returned from FilterUpdatePauseStatus and is used to iterate over the raw logs and unpacked data for UpdatePauseStatus events raised by the StapleController contract.
type StapleControllerUpdatePauseStatusIterator struct {
	Event *StapleControllerUpdatePauseStatus // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdatePauseStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdatePauseStatus)
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
		it.Event = new(StapleControllerUpdatePauseStatus)
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
func (it *StapleControllerUpdatePauseStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdatePauseStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdatePauseStatus represents a UpdatePauseStatus event raised by the StapleController contract.
type StapleControllerUpdatePauseStatus struct {
	Underlying common.Address
	VtpID      *big.Int
	NewStatus  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdatePauseStatus is a free log retrieval operation binding the contract event 0x7d19969ff682d20681fd82ea5b58ce7338ad651b6373bd81a3e91e1bfdca3c38.
//
// Solidity: event UpdatePauseStatus(address indexed underlying, uint256 indexed vtpID, bool newStatus)
func (_StapleController *StapleControllerFilterer) FilterUpdatePauseStatus(opts *bind.FilterOpts, underlying []common.Address, vtpID []*big.Int) (*StapleControllerUpdatePauseStatusIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdatePauseStatus", underlyingRule, vtpIDRule)
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdatePauseStatusIterator{contract: _StapleController.contract, event: "UpdatePauseStatus", logs: logs, sub: sub}, nil
}

// WatchUpdatePauseStatus is a free log subscription operation binding the contract event 0x7d19969ff682d20681fd82ea5b58ce7338ad651b6373bd81a3e91e1bfdca3c38.
//
// Solidity: event UpdatePauseStatus(address indexed underlying, uint256 indexed vtpID, bool newStatus)
func (_StapleController *StapleControllerFilterer) WatchUpdatePauseStatus(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdatePauseStatus, underlying []common.Address, vtpID []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdatePauseStatus", underlyingRule, vtpIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdatePauseStatus)
				if err := _StapleController.contract.UnpackLog(event, "UpdatePauseStatus", log); err != nil {
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

// ParseUpdatePauseStatus is a log parse operation binding the contract event 0x7d19969ff682d20681fd82ea5b58ce7338ad651b6373bd81a3e91e1bfdca3c38.
//
// Solidity: event UpdatePauseStatus(address indexed underlying, uint256 indexed vtpID, bool newStatus)
func (_StapleController *StapleControllerFilterer) ParseUpdatePauseStatus(log types.Log) (*StapleControllerUpdatePauseStatus, error) {
	event := new(StapleControllerUpdatePauseStatus)
	if err := _StapleController.contract.UnpackLog(event, "UpdatePauseStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleControllerUpdateProtocolFeeRateIterator is returned from FilterUpdateProtocolFeeRate and is used to iterate over the raw logs and unpacked data for UpdateProtocolFeeRate events raised by the StapleController contract.
type StapleControllerUpdateProtocolFeeRateIterator struct {
	Event *StapleControllerUpdateProtocolFeeRate // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdateProtocolFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdateProtocolFeeRate)
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
		it.Event = new(StapleControllerUpdateProtocolFeeRate)
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
func (it *StapleControllerUpdateProtocolFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdateProtocolFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdateProtocolFeeRate represents a UpdateProtocolFeeRate event raised by the StapleController contract.
type StapleControllerUpdateProtocolFeeRate struct {
	Underlying common.Address
	VtpID      *big.Int
	NewRate    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateProtocolFeeRate is a free log retrieval operation binding the contract event 0x624d54423f9f0e4b4e50ed18d6a5e38c8d75cd8dc58fe041d4e71ea323830d31.
//
// Solidity: event UpdateProtocolFeeRate(address indexed underlying, uint256 indexed vtpID, uint256 newRate)
func (_StapleController *StapleControllerFilterer) FilterUpdateProtocolFeeRate(opts *bind.FilterOpts, underlying []common.Address, vtpID []*big.Int) (*StapleControllerUpdateProtocolFeeRateIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdateProtocolFeeRate", underlyingRule, vtpIDRule)
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdateProtocolFeeRateIterator{contract: _StapleController.contract, event: "UpdateProtocolFeeRate", logs: logs, sub: sub}, nil
}

// WatchUpdateProtocolFeeRate is a free log subscription operation binding the contract event 0x624d54423f9f0e4b4e50ed18d6a5e38c8d75cd8dc58fe041d4e71ea323830d31.
//
// Solidity: event UpdateProtocolFeeRate(address indexed underlying, uint256 indexed vtpID, uint256 newRate)
func (_StapleController *StapleControllerFilterer) WatchUpdateProtocolFeeRate(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdateProtocolFeeRate, underlying []common.Address, vtpID []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdateProtocolFeeRate", underlyingRule, vtpIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdateProtocolFeeRate)
				if err := _StapleController.contract.UnpackLog(event, "UpdateProtocolFeeRate", log); err != nil {
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

// ParseUpdateProtocolFeeRate is a log parse operation binding the contract event 0x624d54423f9f0e4b4e50ed18d6a5e38c8d75cd8dc58fe041d4e71ea323830d31.
//
// Solidity: event UpdateProtocolFeeRate(address indexed underlying, uint256 indexed vtpID, uint256 newRate)
func (_StapleController *StapleControllerFilterer) ParseUpdateProtocolFeeRate(log types.Log) (*StapleControllerUpdateProtocolFeeRate, error) {
	event := new(StapleControllerUpdateProtocolFeeRate)
	if err := _StapleController.contract.UnpackLog(event, "UpdateProtocolFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleControllerUpdateProxyAdminIterator is returned from FilterUpdateProxyAdmin and is used to iterate over the raw logs and unpacked data for UpdateProxyAdmin events raised by the StapleController contract.
type StapleControllerUpdateProxyAdminIterator struct {
	Event *StapleControllerUpdateProxyAdmin // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdateProxyAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdateProxyAdmin)
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
		it.Event = new(StapleControllerUpdateProxyAdmin)
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
func (it *StapleControllerUpdateProxyAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdateProxyAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdateProxyAdmin represents a UpdateProxyAdmin event raised by the StapleController contract.
type StapleControllerUpdateProxyAdmin struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateProxyAdmin is a free log retrieval operation binding the contract event 0x8cb0f1efb2fb9a441d1759cb883b073bfdcb20b9abae9ca276d2e73b4fed27b0.
//
// Solidity: event UpdateProxyAdmin(address newAddress)
func (_StapleController *StapleControllerFilterer) FilterUpdateProxyAdmin(opts *bind.FilterOpts) (*StapleControllerUpdateProxyAdminIterator, error) {

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdateProxyAdmin")
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdateProxyAdminIterator{contract: _StapleController.contract, event: "UpdateProxyAdmin", logs: logs, sub: sub}, nil
}

// WatchUpdateProxyAdmin is a free log subscription operation binding the contract event 0x8cb0f1efb2fb9a441d1759cb883b073bfdcb20b9abae9ca276d2e73b4fed27b0.
//
// Solidity: event UpdateProxyAdmin(address newAddress)
func (_StapleController *StapleControllerFilterer) WatchUpdateProxyAdmin(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdateProxyAdmin) (event.Subscription, error) {

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdateProxyAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdateProxyAdmin)
				if err := _StapleController.contract.UnpackLog(event, "UpdateProxyAdmin", log); err != nil {
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

// ParseUpdateProxyAdmin is a log parse operation binding the contract event 0x8cb0f1efb2fb9a441d1759cb883b073bfdcb20b9abae9ca276d2e73b4fed27b0.
//
// Solidity: event UpdateProxyAdmin(address newAddress)
func (_StapleController *StapleControllerFilterer) ParseUpdateProxyAdmin(log types.Log) (*StapleControllerUpdateProxyAdmin, error) {
	event := new(StapleControllerUpdateProxyAdmin)
	if err := _StapleController.contract.UnpackLog(event, "UpdateProxyAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleControllerUpdateRiskInfoIterator is returned from FilterUpdateRiskInfo and is used to iterate over the raw logs and unpacked data for UpdateRiskInfo events raised by the StapleController contract.
type StapleControllerUpdateRiskInfoIterator struct {
	Event *StapleControllerUpdateRiskInfo // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdateRiskInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdateRiskInfo)
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
		it.Event = new(StapleControllerUpdateRiskInfo)
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
func (it *StapleControllerUpdateRiskInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdateRiskInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdateRiskInfo represents a UpdateRiskInfo event raised by the StapleController contract.
type StapleControllerUpdateRiskInfo struct {
	CreditLimit *big.Int
	Vls         []*big.Int
	Credits     []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateRiskInfo is a free log retrieval operation binding the contract event 0xcf4dfb38125a89162de306f2291d59baf39b3fd7c9b92de80dff22839bbdeb7b.
//
// Solidity: event UpdateRiskInfo(uint256 creditLimit, uint256[] vls, uint256[] credits)
func (_StapleController *StapleControllerFilterer) FilterUpdateRiskInfo(opts *bind.FilterOpts) (*StapleControllerUpdateRiskInfoIterator, error) {

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdateRiskInfo")
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdateRiskInfoIterator{contract: _StapleController.contract, event: "UpdateRiskInfo", logs: logs, sub: sub}, nil
}

// WatchUpdateRiskInfo is a free log subscription operation binding the contract event 0xcf4dfb38125a89162de306f2291d59baf39b3fd7c9b92de80dff22839bbdeb7b.
//
// Solidity: event UpdateRiskInfo(uint256 creditLimit, uint256[] vls, uint256[] credits)
func (_StapleController *StapleControllerFilterer) WatchUpdateRiskInfo(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdateRiskInfo) (event.Subscription, error) {

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdateRiskInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdateRiskInfo)
				if err := _StapleController.contract.UnpackLog(event, "UpdateRiskInfo", log); err != nil {
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

// ParseUpdateRiskInfo is a log parse operation binding the contract event 0xcf4dfb38125a89162de306f2291d59baf39b3fd7c9b92de80dff22839bbdeb7b.
//
// Solidity: event UpdateRiskInfo(uint256 creditLimit, uint256[] vls, uint256[] credits)
func (_StapleController *StapleControllerFilterer) ParseUpdateRiskInfo(log types.Log) (*StapleControllerUpdateRiskInfo, error) {
	event := new(StapleControllerUpdateRiskInfo)
	if err := _StapleController.contract.UnpackLog(event, "UpdateRiskInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleControllerUpdateRouterIterator is returned from FilterUpdateRouter and is used to iterate over the raw logs and unpacked data for UpdateRouter events raised by the StapleController contract.
type StapleControllerUpdateRouterIterator struct {
	Event *StapleControllerUpdateRouter // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdateRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdateRouter)
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
		it.Event = new(StapleControllerUpdateRouter)
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
func (it *StapleControllerUpdateRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdateRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdateRouter represents a UpdateRouter event raised by the StapleController contract.
type StapleControllerUpdateRouter struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateRouter is a free log retrieval operation binding the contract event 0xd5f5b4d616f94f5e10b2955392470790b3fdde7c0c0b68bd1f3ea635e2caa8d7.
//
// Solidity: event UpdateRouter(address newAddress)
func (_StapleController *StapleControllerFilterer) FilterUpdateRouter(opts *bind.FilterOpts) (*StapleControllerUpdateRouterIterator, error) {

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdateRouter")
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdateRouterIterator{contract: _StapleController.contract, event: "UpdateRouter", logs: logs, sub: sub}, nil
}

// WatchUpdateRouter is a free log subscription operation binding the contract event 0xd5f5b4d616f94f5e10b2955392470790b3fdde7c0c0b68bd1f3ea635e2caa8d7.
//
// Solidity: event UpdateRouter(address newAddress)
func (_StapleController *StapleControllerFilterer) WatchUpdateRouter(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdateRouter) (event.Subscription, error) {

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdateRouter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdateRouter)
				if err := _StapleController.contract.UnpackLog(event, "UpdateRouter", log); err != nil {
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

// ParseUpdateRouter is a log parse operation binding the contract event 0xd5f5b4d616f94f5e10b2955392470790b3fdde7c0c0b68bd1f3ea635e2caa8d7.
//
// Solidity: event UpdateRouter(address newAddress)
func (_StapleController *StapleControllerFilterer) ParseUpdateRouter(log types.Log) (*StapleControllerUpdateRouter, error) {
	event := new(StapleControllerUpdateRouter)
	if err := _StapleController.contract.UnpackLog(event, "UpdateRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleControllerUpdateSwapFeeIterator is returned from FilterUpdateSwapFee and is used to iterate over the raw logs and unpacked data for UpdateSwapFee events raised by the StapleController contract.
type StapleControllerUpdateSwapFeeIterator struct {
	Event *StapleControllerUpdateSwapFee // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdateSwapFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdateSwapFee)
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
		it.Event = new(StapleControllerUpdateSwapFee)
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
func (it *StapleControllerUpdateSwapFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdateSwapFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdateSwapFee represents a UpdateSwapFee event raised by the StapleController contract.
type StapleControllerUpdateSwapFee struct {
	Underlying common.Address
	VtpID      *big.Int
	NewIn      *big.Int
	NewOut     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateSwapFee is a free log retrieval operation binding the contract event 0x6a7f6d6b57893e857c5d30e0d3629ea5501650ace6705647d6cd740005c19a34.
//
// Solidity: event UpdateSwapFee(address indexed underlying, uint256 indexed vtpID, uint256 newIn, uint256 newOut)
func (_StapleController *StapleControllerFilterer) FilterUpdateSwapFee(opts *bind.FilterOpts, underlying []common.Address, vtpID []*big.Int) (*StapleControllerUpdateSwapFeeIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdateSwapFee", underlyingRule, vtpIDRule)
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdateSwapFeeIterator{contract: _StapleController.contract, event: "UpdateSwapFee", logs: logs, sub: sub}, nil
}

// WatchUpdateSwapFee is a free log subscription operation binding the contract event 0x6a7f6d6b57893e857c5d30e0d3629ea5501650ace6705647d6cd740005c19a34.
//
// Solidity: event UpdateSwapFee(address indexed underlying, uint256 indexed vtpID, uint256 newIn, uint256 newOut)
func (_StapleController *StapleControllerFilterer) WatchUpdateSwapFee(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdateSwapFee, underlying []common.Address, vtpID []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vtpIDRule []interface{}
	for _, vtpIDItem := range vtpID {
		vtpIDRule = append(vtpIDRule, vtpIDItem)
	}

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdateSwapFee", underlyingRule, vtpIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdateSwapFee)
				if err := _StapleController.contract.UnpackLog(event, "UpdateSwapFee", log); err != nil {
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

// ParseUpdateSwapFee is a log parse operation binding the contract event 0x6a7f6d6b57893e857c5d30e0d3629ea5501650ace6705647d6cd740005c19a34.
//
// Solidity: event UpdateSwapFee(address indexed underlying, uint256 indexed vtpID, uint256 newIn, uint256 newOut)
func (_StapleController *StapleControllerFilterer) ParseUpdateSwapFee(log types.Log) (*StapleControllerUpdateSwapFee, error) {
	event := new(StapleControllerUpdateSwapFee)
	if err := _StapleController.contract.UnpackLog(event, "UpdateSwapFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleControllerUpdateTLIterator is returned from FilterUpdateTL and is used to iterate over the raw logs and unpacked data for UpdateTL events raised by the StapleController contract.
type StapleControllerUpdateTLIterator struct {
	Event *StapleControllerUpdateTL // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdateTLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdateTL)
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
		it.Event = new(StapleControllerUpdateTL)
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
func (it *StapleControllerUpdateTLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdateTLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdateTL represents a UpdateTL event raised by the StapleController contract.
type StapleControllerUpdateTL struct {
	Underlying common.Address
	NewTL      uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateTL is a free log retrieval operation binding the contract event 0xac065ae6ed4bd1cc5125544fea35e4011f475af7fe8dabf8a300e1b23af793da.
//
// Solidity: event UpdateTL(address indexed underlying, uint8 newTL)
func (_StapleController *StapleControllerFilterer) FilterUpdateTL(opts *bind.FilterOpts, underlying []common.Address) (*StapleControllerUpdateTLIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdateTL", underlyingRule)
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdateTLIterator{contract: _StapleController.contract, event: "UpdateTL", logs: logs, sub: sub}, nil
}

// WatchUpdateTL is a free log subscription operation binding the contract event 0xac065ae6ed4bd1cc5125544fea35e4011f475af7fe8dabf8a300e1b23af793da.
//
// Solidity: event UpdateTL(address indexed underlying, uint8 newTL)
func (_StapleController *StapleControllerFilterer) WatchUpdateTL(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdateTL, underlying []common.Address) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdateTL", underlyingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdateTL)
				if err := _StapleController.contract.UnpackLog(event, "UpdateTL", log); err != nil {
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

// ParseUpdateTL is a log parse operation binding the contract event 0xac065ae6ed4bd1cc5125544fea35e4011f475af7fe8dabf8a300e1b23af793da.
//
// Solidity: event UpdateTL(address indexed underlying, uint8 newTL)
func (_StapleController *StapleControllerFilterer) ParseUpdateTL(log types.Log) (*StapleControllerUpdateTL, error) {
	event := new(StapleControllerUpdateTL)
	if err := _StapleController.contract.UnpackLog(event, "UpdateTL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StapleControllerUpdateWhiteListIterator is returned from FilterUpdateWhiteList and is used to iterate over the raw logs and unpacked data for UpdateWhiteList events raised by the StapleController contract.
type StapleControllerUpdateWhiteListIterator struct {
	Event *StapleControllerUpdateWhiteList // Event containing the contract specifics and raw log

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
func (it *StapleControllerUpdateWhiteListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StapleControllerUpdateWhiteList)
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
		it.Event = new(StapleControllerUpdateWhiteList)
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
func (it *StapleControllerUpdateWhiteListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StapleControllerUpdateWhiteListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StapleControllerUpdateWhiteList represents a UpdateWhiteList event raised by the StapleController contract.
type StapleControllerUpdateWhiteList struct {
	Addresses []common.Address
	Discounts []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateWhiteList is a free log retrieval operation binding the contract event 0xa13b10df67b8060b0016082524a0b6dba104b23bcb9e34397dda989972efc7fd.
//
// Solidity: event UpdateWhiteList(address[] addresses, uint256[] discounts)
func (_StapleController *StapleControllerFilterer) FilterUpdateWhiteList(opts *bind.FilterOpts) (*StapleControllerUpdateWhiteListIterator, error) {

	logs, sub, err := _StapleController.contract.FilterLogs(opts, "UpdateWhiteList")
	if err != nil {
		return nil, err
	}
	return &StapleControllerUpdateWhiteListIterator{contract: _StapleController.contract, event: "UpdateWhiteList", logs: logs, sub: sub}, nil
}

// WatchUpdateWhiteList is a free log subscription operation binding the contract event 0xa13b10df67b8060b0016082524a0b6dba104b23bcb9e34397dda989972efc7fd.
//
// Solidity: event UpdateWhiteList(address[] addresses, uint256[] discounts)
func (_StapleController *StapleControllerFilterer) WatchUpdateWhiteList(opts *bind.WatchOpts, sink chan<- *StapleControllerUpdateWhiteList) (event.Subscription, error) {

	logs, sub, err := _StapleController.contract.WatchLogs(opts, "UpdateWhiteList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StapleControllerUpdateWhiteList)
				if err := _StapleController.contract.UnpackLog(event, "UpdateWhiteList", log); err != nil {
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

// ParseUpdateWhiteList is a log parse operation binding the contract event 0xa13b10df67b8060b0016082524a0b6dba104b23bcb9e34397dda989972efc7fd.
//
// Solidity: event UpdateWhiteList(address[] addresses, uint256[] discounts)
func (_StapleController *StapleControllerFilterer) ParseUpdateWhiteList(log types.Log) (*StapleControllerUpdateWhiteList, error) {
	event := new(StapleControllerUpdateWhiteList)
	if err := _StapleController.contract.UnpackLog(event, "UpdateWhiteList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
