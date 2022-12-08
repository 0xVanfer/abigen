// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancerComposableStablePool

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

// ComposableStablePoolSwapRequest is an auto generated low-level Go binding around an user-defined struct.
type ComposableStablePoolSwapRequest struct {
	Kind            uint8
	TokenIn         common.Address
	TokenOut        common.Address
	Amount          *big.Int
	PoolId          [32]byte
	LastChangeBlock *big.Int
	From            common.Address
	To              common.Address
	UserData        []byte
}

// BalancerComposableStablePoolMetaData contains all meta data concerning the BalancerComposableStablePool contract.
var BalancerComposableStablePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActualSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAmplificationParameter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isUpdating\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"precision\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBptIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastJoinExitData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastJoinExitAmplification\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPostJoinExitInvariant\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumBpt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"}],\"name\":\"getProtocolFeePercentageCache\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeesCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolSwapFeeDelegation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateProviders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getScalingFactors\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenRateCache\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inRecoveryMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenExemptFromYieldProtocolFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onExitPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onJoinPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumComposableStablePool.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structComposableStablePool.SwapRequest\",\"name\":\"swapRequest\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"indexIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexOut\",\"type\":\"uint256\"}],\"name\":\"onSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryJoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"poolConfig\",\"type\":\"bytes\"}],\"name\":\"setAssetManagerPoolConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"setSwapFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setTokenRateCacheDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rawEndValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"startAmplificationParameterUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopAmplificationParameterUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateProtocolFeePercentageCache\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"updateTokenRateCache\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BalancerComposableStablePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerComposableStablePoolMetaData.ABI instead.
var BalancerComposableStablePoolABI = BalancerComposableStablePoolMetaData.ABI

// BalancerComposableStablePool is an auto generated Go binding around an Ethereum contract.
type BalancerComposableStablePool struct {
	BalancerComposableStablePoolCaller     // Read-only binding to the contract
	BalancerComposableStablePoolTransactor // Write-only binding to the contract
	BalancerComposableStablePoolFilterer   // Log filterer for contract events
}

// BalancerComposableStablePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerComposableStablePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerComposableStablePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerComposableStablePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerComposableStablePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerComposableStablePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerComposableStablePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerComposableStablePoolSession struct {
	Contract     *BalancerComposableStablePool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BalancerComposableStablePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerComposableStablePoolCallerSession struct {
	Contract *BalancerComposableStablePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// BalancerComposableStablePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerComposableStablePoolTransactorSession struct {
	Contract     *BalancerComposableStablePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// BalancerComposableStablePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerComposableStablePoolRaw struct {
	Contract *BalancerComposableStablePool // Generic contract binding to access the raw methods on
}

// BalancerComposableStablePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerComposableStablePoolCallerRaw struct {
	Contract *BalancerComposableStablePoolCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerComposableStablePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerComposableStablePoolTransactorRaw struct {
	Contract *BalancerComposableStablePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerComposableStablePool creates a new instance of BalancerComposableStablePool, bound to a specific deployed contract.
func NewBalancerComposableStablePool(address common.Address, backend bind.ContractBackend) (*BalancerComposableStablePool, error) {
	contract, err := bindBalancerComposableStablePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerComposableStablePool{BalancerComposableStablePoolCaller: BalancerComposableStablePoolCaller{contract: contract}, BalancerComposableStablePoolTransactor: BalancerComposableStablePoolTransactor{contract: contract}, BalancerComposableStablePoolFilterer: BalancerComposableStablePoolFilterer{contract: contract}}, nil
}

// NewBalancerComposableStablePoolCaller creates a new read-only instance of BalancerComposableStablePool, bound to a specific deployed contract.
func NewBalancerComposableStablePoolCaller(address common.Address, caller bind.ContractCaller) (*BalancerComposableStablePoolCaller, error) {
	contract, err := bindBalancerComposableStablePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerComposableStablePoolCaller{contract: contract}, nil
}

// NewBalancerComposableStablePoolTransactor creates a new write-only instance of BalancerComposableStablePool, bound to a specific deployed contract.
func NewBalancerComposableStablePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerComposableStablePoolTransactor, error) {
	contract, err := bindBalancerComposableStablePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerComposableStablePoolTransactor{contract: contract}, nil
}

// NewBalancerComposableStablePoolFilterer creates a new log filterer instance of BalancerComposableStablePool, bound to a specific deployed contract.
func NewBalancerComposableStablePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerComposableStablePoolFilterer, error) {
	contract, err := bindBalancerComposableStablePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerComposableStablePoolFilterer{contract: contract}, nil
}

// bindBalancerComposableStablePool binds a generic wrapper to an already deployed contract.
func bindBalancerComposableStablePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerComposableStablePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerComposableStablePool *BalancerComposableStablePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerComposableStablePool.Contract.BalancerComposableStablePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerComposableStablePool *BalancerComposableStablePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.BalancerComposableStablePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerComposableStablePool *BalancerComposableStablePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.BalancerComposableStablePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerComposableStablePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.contract.Transact(opts, method, params...)
}

// DELEGATEPROTOCOLSWAPFEESSENTINEL is a free data retrieval call binding the contract method 0xddf4627b.
//
// Solidity: function DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) DELEGATEPROTOCOLSWAPFEESSENTINEL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DELEGATEPROTOCOLSWAPFEESSENTINEL is a free data retrieval call binding the contract method 0xddf4627b.
//
// Solidity: function DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) DELEGATEPROTOCOLSWAPFEESSENTINEL() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.DELEGATEPROTOCOLSWAPFEESSENTINEL(&_BalancerComposableStablePool.CallOpts)
}

// DELEGATEPROTOCOLSWAPFEESSENTINEL is a free data retrieval call binding the contract method 0xddf4627b.
//
// Solidity: function DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) DELEGATEPROTOCOLSWAPFEESSENTINEL() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.DELEGATEPROTOCOLSWAPFEESSENTINEL(&_BalancerComposableStablePool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BalancerComposableStablePool.Contract.DOMAINSEPARATOR(&_BalancerComposableStablePool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BalancerComposableStablePool.Contract.DOMAINSEPARATOR(&_BalancerComposableStablePool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.Allowance(&_BalancerComposableStablePool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.Allowance(&_BalancerComposableStablePool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.BalanceOf(&_BalancerComposableStablePool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.BalanceOf(&_BalancerComposableStablePool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) Decimals() (uint8, error) {
	return _BalancerComposableStablePool.Contract.Decimals(&_BalancerComposableStablePool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) Decimals() (uint8, error) {
	return _BalancerComposableStablePool.Contract.Decimals(&_BalancerComposableStablePool.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BalancerComposableStablePool.Contract.GetActionId(&_BalancerComposableStablePool.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BalancerComposableStablePool.Contract.GetActionId(&_BalancerComposableStablePool.CallOpts, selector)
}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetActualSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getActualSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetActualSupply() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetActualSupply(&_BalancerComposableStablePool.CallOpts)
}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetActualSupply() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetActualSupply(&_BalancerComposableStablePool.CallOpts)
}

// GetAmplificationParameter is a free data retrieval call binding the contract method 0x6daccffa.
//
// Solidity: function getAmplificationParameter() view returns(uint256 value, bool isUpdating, uint256 precision)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetAmplificationParameter(opts *bind.CallOpts) (struct {
	Value      *big.Int
	IsUpdating bool
	Precision  *big.Int
}, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getAmplificationParameter")

	outstruct := new(struct {
		Value      *big.Int
		IsUpdating bool
		Precision  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsUpdating = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Precision = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAmplificationParameter is a free data retrieval call binding the contract method 0x6daccffa.
//
// Solidity: function getAmplificationParameter() view returns(uint256 value, bool isUpdating, uint256 precision)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetAmplificationParameter() (struct {
	Value      *big.Int
	IsUpdating bool
	Precision  *big.Int
}, error) {
	return _BalancerComposableStablePool.Contract.GetAmplificationParameter(&_BalancerComposableStablePool.CallOpts)
}

// GetAmplificationParameter is a free data retrieval call binding the contract method 0x6daccffa.
//
// Solidity: function getAmplificationParameter() view returns(uint256 value, bool isUpdating, uint256 precision)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetAmplificationParameter() (struct {
	Value      *big.Int
	IsUpdating bool
	Precision  *big.Int
}, error) {
	return _BalancerComposableStablePool.Contract.GetAmplificationParameter(&_BalancerComposableStablePool.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetAuthorizer() (common.Address, error) {
	return _BalancerComposableStablePool.Contract.GetAuthorizer(&_BalancerComposableStablePool.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetAuthorizer() (common.Address, error) {
	return _BalancerComposableStablePool.Contract.GetAuthorizer(&_BalancerComposableStablePool.CallOpts)
}

// GetBptIndex is a free data retrieval call binding the contract method 0x82687a56.
//
// Solidity: function getBptIndex() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetBptIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getBptIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBptIndex is a free data retrieval call binding the contract method 0x82687a56.
//
// Solidity: function getBptIndex() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetBptIndex() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetBptIndex(&_BalancerComposableStablePool.CallOpts)
}

// GetBptIndex is a free data retrieval call binding the contract method 0x82687a56.
//
// Solidity: function getBptIndex() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetBptIndex() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetBptIndex(&_BalancerComposableStablePool.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetDomainSeparator() ([32]byte, error) {
	return _BalancerComposableStablePool.Contract.GetDomainSeparator(&_BalancerComposableStablePool.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _BalancerComposableStablePool.Contract.GetDomainSeparator(&_BalancerComposableStablePool.CallOpts)
}

// GetLastJoinExitData is a free data retrieval call binding the contract method 0x3c975d51.
//
// Solidity: function getLastJoinExitData() view returns(uint256 lastJoinExitAmplification, uint256 lastPostJoinExitInvariant)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetLastJoinExitData(opts *bind.CallOpts) (struct {
	LastJoinExitAmplification *big.Int
	LastPostJoinExitInvariant *big.Int
}, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getLastJoinExitData")

	outstruct := new(struct {
		LastJoinExitAmplification *big.Int
		LastPostJoinExitInvariant *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastJoinExitAmplification = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastPostJoinExitInvariant = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLastJoinExitData is a free data retrieval call binding the contract method 0x3c975d51.
//
// Solidity: function getLastJoinExitData() view returns(uint256 lastJoinExitAmplification, uint256 lastPostJoinExitInvariant)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetLastJoinExitData() (struct {
	LastJoinExitAmplification *big.Int
	LastPostJoinExitInvariant *big.Int
}, error) {
	return _BalancerComposableStablePool.Contract.GetLastJoinExitData(&_BalancerComposableStablePool.CallOpts)
}

// GetLastJoinExitData is a free data retrieval call binding the contract method 0x3c975d51.
//
// Solidity: function getLastJoinExitData() view returns(uint256 lastJoinExitAmplification, uint256 lastPostJoinExitInvariant)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetLastJoinExitData() (struct {
	LastJoinExitAmplification *big.Int
	LastPostJoinExitInvariant *big.Int
}, error) {
	return _BalancerComposableStablePool.Contract.GetLastJoinExitData(&_BalancerComposableStablePool.CallOpts)
}

// GetMinimumBpt is a free data retrieval call binding the contract method 0x04842d4c.
//
// Solidity: function getMinimumBpt() pure returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetMinimumBpt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getMinimumBpt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumBpt is a free data retrieval call binding the contract method 0x04842d4c.
//
// Solidity: function getMinimumBpt() pure returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetMinimumBpt() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetMinimumBpt(&_BalancerComposableStablePool.CallOpts)
}

// GetMinimumBpt is a free data retrieval call binding the contract method 0x04842d4c.
//
// Solidity: function getMinimumBpt() pure returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetMinimumBpt() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetMinimumBpt(&_BalancerComposableStablePool.CallOpts)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetNextNonce(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getNextNonce", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetNextNonce(&_BalancerComposableStablePool.CallOpts, account)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetNextNonce(&_BalancerComposableStablePool.CallOpts, account)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetOwner() (common.Address, error) {
	return _BalancerComposableStablePool.Contract.GetOwner(&_BalancerComposableStablePool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetOwner() (common.Address, error) {
	return _BalancerComposableStablePool.Contract.GetOwner(&_BalancerComposableStablePool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getPausedState")

	outstruct := new(struct {
		Paused              bool
		PauseWindowEndTime  *big.Int
		BufferPeriodEndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PauseWindowEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BufferPeriodEndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BalancerComposableStablePool.Contract.GetPausedState(&_BalancerComposableStablePool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BalancerComposableStablePool.Contract.GetPausedState(&_BalancerComposableStablePool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetPoolId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getPoolId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetPoolId() ([32]byte, error) {
	return _BalancerComposableStablePool.Contract.GetPoolId(&_BalancerComposableStablePool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetPoolId() ([32]byte, error) {
	return _BalancerComposableStablePool.Contract.GetPoolId(&_BalancerComposableStablePool.CallOpts)
}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetProtocolFeePercentageCache(opts *bind.CallOpts, feeType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getProtocolFeePercentageCache", feeType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetProtocolFeePercentageCache(feeType *big.Int) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetProtocolFeePercentageCache(&_BalancerComposableStablePool.CallOpts, feeType)
}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetProtocolFeePercentageCache(feeType *big.Int) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetProtocolFeePercentageCache(&_BalancerComposableStablePool.CallOpts, feeType)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetProtocolFeesCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getProtocolFeesCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetProtocolFeesCollector() (common.Address, error) {
	return _BalancerComposableStablePool.Contract.GetProtocolFeesCollector(&_BalancerComposableStablePool.CallOpts)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetProtocolFeesCollector() (common.Address, error) {
	return _BalancerComposableStablePool.Contract.GetProtocolFeesCollector(&_BalancerComposableStablePool.CallOpts)
}

// GetProtocolSwapFeeDelegation is a free data retrieval call binding the contract method 0x15b0015b.
//
// Solidity: function getProtocolSwapFeeDelegation() view returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetProtocolSwapFeeDelegation(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getProtocolSwapFeeDelegation")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetProtocolSwapFeeDelegation is a free data retrieval call binding the contract method 0x15b0015b.
//
// Solidity: function getProtocolSwapFeeDelegation() view returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetProtocolSwapFeeDelegation() (bool, error) {
	return _BalancerComposableStablePool.Contract.GetProtocolSwapFeeDelegation(&_BalancerComposableStablePool.CallOpts)
}

// GetProtocolSwapFeeDelegation is a free data retrieval call binding the contract method 0x15b0015b.
//
// Solidity: function getProtocolSwapFeeDelegation() view returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetProtocolSwapFeeDelegation() (bool, error) {
	return _BalancerComposableStablePool.Contract.GetProtocolSwapFeeDelegation(&_BalancerComposableStablePool.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetRate() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetRate(&_BalancerComposableStablePool.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetRate() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetRate(&_BalancerComposableStablePool.CallOpts)
}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetRateProviders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getRateProviders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetRateProviders() ([]common.Address, error) {
	return _BalancerComposableStablePool.Contract.GetRateProviders(&_BalancerComposableStablePool.CallOpts)
}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetRateProviders() ([]common.Address, error) {
	return _BalancerComposableStablePool.Contract.GetRateProviders(&_BalancerComposableStablePool.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetScalingFactors(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getScalingFactors")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetScalingFactors() ([]*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetScalingFactors(&_BalancerComposableStablePool.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetScalingFactors() ([]*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetScalingFactors(&_BalancerComposableStablePool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetSwapFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getSwapFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetSwapFeePercentage(&_BalancerComposableStablePool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetSwapFeePercentage(&_BalancerComposableStablePool.CallOpts)
}

// GetTokenRate is a free data retrieval call binding the contract method 0x54dea00a.
//
// Solidity: function getTokenRate(address token) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetTokenRate(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getTokenRate", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenRate is a free data retrieval call binding the contract method 0x54dea00a.
//
// Solidity: function getTokenRate(address token) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetTokenRate(token common.Address) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetTokenRate(&_BalancerComposableStablePool.CallOpts, token)
}

// GetTokenRate is a free data retrieval call binding the contract method 0x54dea00a.
//
// Solidity: function getTokenRate(address token) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetTokenRate(token common.Address) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.GetTokenRate(&_BalancerComposableStablePool.CallOpts, token)
}

// GetTokenRateCache is a free data retrieval call binding the contract method 0x7f1260d1.
//
// Solidity: function getTokenRateCache(address token) view returns(uint256 rate, uint256 oldRate, uint256 duration, uint256 expires)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetTokenRateCache(opts *bind.CallOpts, token common.Address) (struct {
	Rate     *big.Int
	OldRate  *big.Int
	Duration *big.Int
	Expires  *big.Int
}, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getTokenRateCache", token)

	outstruct := new(struct {
		Rate     *big.Int
		OldRate  *big.Int
		Duration *big.Int
		Expires  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OldRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Expires = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenRateCache is a free data retrieval call binding the contract method 0x7f1260d1.
//
// Solidity: function getTokenRateCache(address token) view returns(uint256 rate, uint256 oldRate, uint256 duration, uint256 expires)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetTokenRateCache(token common.Address) (struct {
	Rate     *big.Int
	OldRate  *big.Int
	Duration *big.Int
	Expires  *big.Int
}, error) {
	return _BalancerComposableStablePool.Contract.GetTokenRateCache(&_BalancerComposableStablePool.CallOpts, token)
}

// GetTokenRateCache is a free data retrieval call binding the contract method 0x7f1260d1.
//
// Solidity: function getTokenRateCache(address token) view returns(uint256 rate, uint256 oldRate, uint256 duration, uint256 expires)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetTokenRateCache(token common.Address) (struct {
	Rate     *big.Int
	OldRate  *big.Int
	Duration *big.Int
	Expires  *big.Int
}, error) {
	return _BalancerComposableStablePool.Contract.GetTokenRateCache(&_BalancerComposableStablePool.CallOpts, token)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) GetVault() (common.Address, error) {
	return _BalancerComposableStablePool.Contract.GetVault(&_BalancerComposableStablePool.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) GetVault() (common.Address, error) {
	return _BalancerComposableStablePool.Contract.GetVault(&_BalancerComposableStablePool.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) InRecoveryMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "inRecoveryMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) InRecoveryMode() (bool, error) {
	return _BalancerComposableStablePool.Contract.InRecoveryMode(&_BalancerComposableStablePool.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) InRecoveryMode() (bool, error) {
	return _BalancerComposableStablePool.Contract.InRecoveryMode(&_BalancerComposableStablePool.CallOpts)
}

// IsTokenExemptFromYieldProtocolFee is a free data retrieval call binding the contract method 0xab7759f1.
//
// Solidity: function isTokenExemptFromYieldProtocolFee(address token) view returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) IsTokenExemptFromYieldProtocolFee(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "isTokenExemptFromYieldProtocolFee", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenExemptFromYieldProtocolFee is a free data retrieval call binding the contract method 0xab7759f1.
//
// Solidity: function isTokenExemptFromYieldProtocolFee(address token) view returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) IsTokenExemptFromYieldProtocolFee(token common.Address) (bool, error) {
	return _BalancerComposableStablePool.Contract.IsTokenExemptFromYieldProtocolFee(&_BalancerComposableStablePool.CallOpts, token)
}

// IsTokenExemptFromYieldProtocolFee is a free data retrieval call binding the contract method 0xab7759f1.
//
// Solidity: function isTokenExemptFromYieldProtocolFee(address token) view returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) IsTokenExemptFromYieldProtocolFee(token common.Address) (bool, error) {
	return _BalancerComposableStablePool.Contract.IsTokenExemptFromYieldProtocolFee(&_BalancerComposableStablePool.CallOpts, token)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) Name() (string, error) {
	return _BalancerComposableStablePool.Contract.Name(&_BalancerComposableStablePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) Name() (string, error) {
	return _BalancerComposableStablePool.Contract.Name(&_BalancerComposableStablePool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.Nonces(&_BalancerComposableStablePool.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.Nonces(&_BalancerComposableStablePool.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) Symbol() (string, error) {
	return _BalancerComposableStablePool.Contract.Symbol(&_BalancerComposableStablePool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) Symbol() (string, error) {
	return _BalancerComposableStablePool.Contract.Symbol(&_BalancerComposableStablePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerComposableStablePool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) TotalSupply() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.TotalSupply(&_BalancerComposableStablePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolCallerSession) TotalSupply() (*big.Int, error) {
	return _BalancerComposableStablePool.Contract.TotalSupply(&_BalancerComposableStablePool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.Approve(&_BalancerComposableStablePool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.Approve(&_BalancerComposableStablePool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.DecreaseAllowance(&_BalancerComposableStablePool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.DecreaseAllowance(&_BalancerComposableStablePool.TransactOpts, spender, amount)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) DisableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "disableRecoveryMode")
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.DisableRecoveryMode(&_BalancerComposableStablePool.TransactOpts)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.DisableRecoveryMode(&_BalancerComposableStablePool.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) EnableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "enableRecoveryMode")
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.EnableRecoveryMode(&_BalancerComposableStablePool.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.EnableRecoveryMode(&_BalancerComposableStablePool.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.IncreaseAllowance(&_BalancerComposableStablePool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.IncreaseAllowance(&_BalancerComposableStablePool.TransactOpts, spender, addedValue)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) OnExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "onExitPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.OnExitPool(&_BalancerComposableStablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.OnExitPool(&_BalancerComposableStablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) OnJoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "onJoinPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.OnJoinPool(&_BalancerComposableStablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.OnJoinPool(&_BalancerComposableStablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnSwap is a paid mutator transaction binding the contract method 0x01ec954a.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) swapRequest, uint256[] balances, uint256 indexIn, uint256 indexOut) returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) OnSwap(opts *bind.TransactOpts, swapRequest ComposableStablePoolSwapRequest, balances []*big.Int, indexIn *big.Int, indexOut *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "onSwap", swapRequest, balances, indexIn, indexOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x01ec954a.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) swapRequest, uint256[] balances, uint256 indexIn, uint256 indexOut) returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) OnSwap(swapRequest ComposableStablePoolSwapRequest, balances []*big.Int, indexIn *big.Int, indexOut *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.OnSwap(&_BalancerComposableStablePool.TransactOpts, swapRequest, balances, indexIn, indexOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x01ec954a.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) swapRequest, uint256[] balances, uint256 indexIn, uint256 indexOut) returns(uint256)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) OnSwap(swapRequest ComposableStablePoolSwapRequest, balances []*big.Int, indexIn *big.Int, indexOut *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.OnSwap(&_BalancerComposableStablePool.TransactOpts, swapRequest, balances, indexIn, indexOut)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) Pause() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.Pause(&_BalancerComposableStablePool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) Pause() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.Pause(&_BalancerComposableStablePool.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.Permit(&_BalancerComposableStablePool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.Permit(&_BalancerComposableStablePool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) QueryExit(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "queryExit", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.QueryExit(&_BalancerComposableStablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.QueryExit(&_BalancerComposableStablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) QueryJoin(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "queryJoin", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.QueryJoin(&_BalancerComposableStablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.QueryJoin(&_BalancerComposableStablePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) SetAssetManagerPoolConfig(opts *bind.TransactOpts, token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "setAssetManagerPoolConfig", token, poolConfig)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) SetAssetManagerPoolConfig(token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.SetAssetManagerPoolConfig(&_BalancerComposableStablePool.TransactOpts, token, poolConfig)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) SetAssetManagerPoolConfig(token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.SetAssetManagerPoolConfig(&_BalancerComposableStablePool.TransactOpts, token, poolConfig)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) SetSwapFeePercentage(opts *bind.TransactOpts, swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "setSwapFeePercentage", swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.SetSwapFeePercentage(&_BalancerComposableStablePool.TransactOpts, swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.SetSwapFeePercentage(&_BalancerComposableStablePool.TransactOpts, swapFeePercentage)
}

// SetTokenRateCacheDuration is a paid mutator transaction binding the contract method 0xf4b7964d.
//
// Solidity: function setTokenRateCacheDuration(address token, uint256 duration) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) SetTokenRateCacheDuration(opts *bind.TransactOpts, token common.Address, duration *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "setTokenRateCacheDuration", token, duration)
}

// SetTokenRateCacheDuration is a paid mutator transaction binding the contract method 0xf4b7964d.
//
// Solidity: function setTokenRateCacheDuration(address token, uint256 duration) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) SetTokenRateCacheDuration(token common.Address, duration *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.SetTokenRateCacheDuration(&_BalancerComposableStablePool.TransactOpts, token, duration)
}

// SetTokenRateCacheDuration is a paid mutator transaction binding the contract method 0xf4b7964d.
//
// Solidity: function setTokenRateCacheDuration(address token, uint256 duration) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) SetTokenRateCacheDuration(token common.Address, duration *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.SetTokenRateCacheDuration(&_BalancerComposableStablePool.TransactOpts, token, duration)
}

// StartAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0x2f1a0bc9.
//
// Solidity: function startAmplificationParameterUpdate(uint256 rawEndValue, uint256 endTime) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) StartAmplificationParameterUpdate(opts *bind.TransactOpts, rawEndValue *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "startAmplificationParameterUpdate", rawEndValue, endTime)
}

// StartAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0x2f1a0bc9.
//
// Solidity: function startAmplificationParameterUpdate(uint256 rawEndValue, uint256 endTime) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) StartAmplificationParameterUpdate(rawEndValue *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.StartAmplificationParameterUpdate(&_BalancerComposableStablePool.TransactOpts, rawEndValue, endTime)
}

// StartAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0x2f1a0bc9.
//
// Solidity: function startAmplificationParameterUpdate(uint256 rawEndValue, uint256 endTime) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) StartAmplificationParameterUpdate(rawEndValue *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.StartAmplificationParameterUpdate(&_BalancerComposableStablePool.TransactOpts, rawEndValue, endTime)
}

// StopAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0xeb0f24d6.
//
// Solidity: function stopAmplificationParameterUpdate() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) StopAmplificationParameterUpdate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "stopAmplificationParameterUpdate")
}

// StopAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0xeb0f24d6.
//
// Solidity: function stopAmplificationParameterUpdate() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) StopAmplificationParameterUpdate() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.StopAmplificationParameterUpdate(&_BalancerComposableStablePool.TransactOpts)
}

// StopAmplificationParameterUpdate is a paid mutator transaction binding the contract method 0xeb0f24d6.
//
// Solidity: function stopAmplificationParameterUpdate() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) StopAmplificationParameterUpdate() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.StopAmplificationParameterUpdate(&_BalancerComposableStablePool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.Transfer(&_BalancerComposableStablePool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.Transfer(&_BalancerComposableStablePool.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.TransferFrom(&_BalancerComposableStablePool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.TransferFrom(&_BalancerComposableStablePool.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) Unpause() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.Unpause(&_BalancerComposableStablePool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.Unpause(&_BalancerComposableStablePool.TransactOpts)
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) UpdateProtocolFeePercentageCache(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "updateProtocolFeePercentageCache")
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) UpdateProtocolFeePercentageCache() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.UpdateProtocolFeePercentageCache(&_BalancerComposableStablePool.TransactOpts)
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) UpdateProtocolFeePercentageCache() (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.UpdateProtocolFeePercentageCache(&_BalancerComposableStablePool.TransactOpts)
}

// UpdateTokenRateCache is a paid mutator transaction binding the contract method 0x2df2c7c0.
//
// Solidity: function updateTokenRateCache(address token) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactor) UpdateTokenRateCache(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BalancerComposableStablePool.contract.Transact(opts, "updateTokenRateCache", token)
}

// UpdateTokenRateCache is a paid mutator transaction binding the contract method 0x2df2c7c0.
//
// Solidity: function updateTokenRateCache(address token) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolSession) UpdateTokenRateCache(token common.Address) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.UpdateTokenRateCache(&_BalancerComposableStablePool.TransactOpts, token)
}

// UpdateTokenRateCache is a paid mutator transaction binding the contract method 0x2df2c7c0.
//
// Solidity: function updateTokenRateCache(address token) returns()
func (_BalancerComposableStablePool *BalancerComposableStablePoolTransactorSession) UpdateTokenRateCache(token common.Address) (*types.Transaction, error) {
	return _BalancerComposableStablePool.Contract.UpdateTokenRateCache(&_BalancerComposableStablePool.TransactOpts, token)
}
