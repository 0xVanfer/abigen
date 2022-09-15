// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KyberRouter

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

// KyberRouterMetaData contains all meta data concerning the KyberRouter contract.
var KyberRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"vReserveRatioBounds\",\"type\":\"uint256[2]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"vReserveRatioBounds\",\"type\":\"uint256[2]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"ampBps\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityNewPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"ampBps\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityNewPoolETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"poolsPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"poolsPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"poolsPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"poolsPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"poolsPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"poolsPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"poolsPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"poolsPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"poolsPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"poolsPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"poolsPath\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// KyberRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use KyberRouterMetaData.ABI instead.
var KyberRouterABI = KyberRouterMetaData.ABI

// KyberRouter is an auto generated Go binding around an Ethereum contract.
type KyberRouter struct {
	KyberRouterCaller     // Read-only binding to the contract
	KyberRouterTransactor // Write-only binding to the contract
	KyberRouterFilterer   // Log filterer for contract events
}

// KyberRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberRouterSession struct {
	Contract     *KyberRouter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KyberRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberRouterCallerSession struct {
	Contract *KyberRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// KyberRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberRouterTransactorSession struct {
	Contract     *KyberRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// KyberRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberRouterRaw struct {
	Contract *KyberRouter // Generic contract binding to access the raw methods on
}

// KyberRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberRouterCallerRaw struct {
	Contract *KyberRouterCaller // Generic read-only contract binding to access the raw methods on
}

// KyberRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberRouterTransactorRaw struct {
	Contract *KyberRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberRouter creates a new instance of KyberRouter, bound to a specific deployed contract.
func NewKyberRouter(address common.Address, backend bind.ContractBackend) (*KyberRouter, error) {
	contract, err := bindKyberRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KyberRouter{KyberRouterCaller: KyberRouterCaller{contract: contract}, KyberRouterTransactor: KyberRouterTransactor{contract: contract}, KyberRouterFilterer: KyberRouterFilterer{contract: contract}}, nil
}

// NewKyberRouterCaller creates a new read-only instance of KyberRouter, bound to a specific deployed contract.
func NewKyberRouterCaller(address common.Address, caller bind.ContractCaller) (*KyberRouterCaller, error) {
	contract, err := bindKyberRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberRouterCaller{contract: contract}, nil
}

// NewKyberRouterTransactor creates a new write-only instance of KyberRouter, bound to a specific deployed contract.
func NewKyberRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberRouterTransactor, error) {
	contract, err := bindKyberRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberRouterTransactor{contract: contract}, nil
}

// NewKyberRouterFilterer creates a new log filterer instance of KyberRouter, bound to a specific deployed contract.
func NewKyberRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberRouterFilterer, error) {
	contract, err := bindKyberRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberRouterFilterer{contract: contract}, nil
}

// bindKyberRouter binds a generic wrapper to an already deployed contract.
func bindKyberRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberRouter *KyberRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberRouter.Contract.KyberRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberRouter *KyberRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberRouter.Contract.KyberRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberRouter *KyberRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberRouter.Contract.KyberRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberRouter *KyberRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KyberRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberRouter *KyberRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberRouter *KyberRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberRouter.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KyberRouter *KyberRouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KyberRouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KyberRouter *KyberRouterSession) Factory() (common.Address, error) {
	return _KyberRouter.Contract.Factory(&_KyberRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KyberRouter *KyberRouterCallerSession) Factory() (common.Address, error) {
	return _KyberRouter.Contract.Factory(&_KyberRouter.CallOpts)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x9e269b68.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] poolsPath, address[] path) view returns(uint256[] amounts)
func (_KyberRouter *KyberRouterCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, poolsPath []common.Address, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _KyberRouter.contract.Call(opts, &out, "getAmountsIn", amountOut, poolsPath, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x9e269b68.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] poolsPath, address[] path) view returns(uint256[] amounts)
func (_KyberRouter *KyberRouterSession) GetAmountsIn(amountOut *big.Int, poolsPath []common.Address, path []common.Address) ([]*big.Int, error) {
	return _KyberRouter.Contract.GetAmountsIn(&_KyberRouter.CallOpts, amountOut, poolsPath, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x9e269b68.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] poolsPath, address[] path) view returns(uint256[] amounts)
func (_KyberRouter *KyberRouterCallerSession) GetAmountsIn(amountOut *big.Int, poolsPath []common.Address, path []common.Address) ([]*big.Int, error) {
	return _KyberRouter.Contract.GetAmountsIn(&_KyberRouter.CallOpts, amountOut, poolsPath, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xa8312b1d.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] poolsPath, address[] path) view returns(uint256[] amounts)
func (_KyberRouter *KyberRouterCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, poolsPath []common.Address, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _KyberRouter.contract.Call(opts, &out, "getAmountsOut", amountIn, poolsPath, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xa8312b1d.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] poolsPath, address[] path) view returns(uint256[] amounts)
func (_KyberRouter *KyberRouterSession) GetAmountsOut(amountIn *big.Int, poolsPath []common.Address, path []common.Address) ([]*big.Int, error) {
	return _KyberRouter.Contract.GetAmountsOut(&_KyberRouter.CallOpts, amountIn, poolsPath, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xa8312b1d.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] poolsPath, address[] path) view returns(uint256[] amounts)
func (_KyberRouter *KyberRouterCallerSession) GetAmountsOut(amountIn *big.Int, poolsPath []common.Address, path []common.Address) ([]*big.Int, error) {
	return _KyberRouter.Contract.GetAmountsOut(&_KyberRouter.CallOpts, amountIn, poolsPath, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_KyberRouter *KyberRouterCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KyberRouter.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_KyberRouter *KyberRouterSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _KyberRouter.Contract.Quote(&_KyberRouter.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_KyberRouter *KyberRouterCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _KyberRouter.Contract.Quote(&_KyberRouter.CallOpts, amountA, reserveA, reserveB)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_KyberRouter *KyberRouterCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KyberRouter.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_KyberRouter *KyberRouterSession) Weth() (common.Address, error) {
	return _KyberRouter.Contract.Weth(&_KyberRouter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_KyberRouter *KyberRouterCallerSession) Weth() (common.Address, error) {
	return _KyberRouter.Contract.Weth(&_KyberRouter.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x24341934.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, address pool, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, uint256[2] vReserveRatioBounds, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_KyberRouter *KyberRouterTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, pool common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, vReserveRatioBounds [2]*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "addLiquidity", tokenA, tokenB, pool, amountADesired, amountBDesired, amountAMin, amountBMin, vReserveRatioBounds, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x24341934.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, address pool, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, uint256[2] vReserveRatioBounds, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_KyberRouter *KyberRouterSession) AddLiquidity(tokenA common.Address, tokenB common.Address, pool common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, vReserveRatioBounds [2]*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.AddLiquidity(&_KyberRouter.TransactOpts, tokenA, tokenB, pool, amountADesired, amountBDesired, amountAMin, amountBMin, vReserveRatioBounds, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x24341934.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, address pool, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, uint256[2] vReserveRatioBounds, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_KyberRouter *KyberRouterTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, pool common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, vReserveRatioBounds [2]*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.AddLiquidity(&_KyberRouter.TransactOpts, tokenA, tokenB, pool, amountADesired, amountBDesired, amountAMin, amountBMin, vReserveRatioBounds, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xd72c1447.
//
// Solidity: function addLiquidityETH(address token, address pool, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, uint256[2] vReserveRatioBounds, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_KyberRouter *KyberRouterTransactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, pool common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, vReserveRatioBounds [2]*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "addLiquidityETH", token, pool, amountTokenDesired, amountTokenMin, amountETHMin, vReserveRatioBounds, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xd72c1447.
//
// Solidity: function addLiquidityETH(address token, address pool, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, uint256[2] vReserveRatioBounds, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_KyberRouter *KyberRouterSession) AddLiquidityETH(token common.Address, pool common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, vReserveRatioBounds [2]*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.AddLiquidityETH(&_KyberRouter.TransactOpts, token, pool, amountTokenDesired, amountTokenMin, amountETHMin, vReserveRatioBounds, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xd72c1447.
//
// Solidity: function addLiquidityETH(address token, address pool, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, uint256[2] vReserveRatioBounds, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_KyberRouter *KyberRouterTransactorSession) AddLiquidityETH(token common.Address, pool common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, vReserveRatioBounds [2]*big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.AddLiquidityETH(&_KyberRouter.TransactOpts, token, pool, amountTokenDesired, amountTokenMin, amountETHMin, vReserveRatioBounds, to, deadline)
}

// AddLiquidityNewPool is a paid mutator transaction binding the contract method 0x1c5d0a6b.
//
// Solidity: function addLiquidityNewPool(address tokenA, address tokenB, uint32 ampBps, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_KyberRouter *KyberRouterTransactor) AddLiquidityNewPool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, ampBps uint32, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "addLiquidityNewPool", tokenA, tokenB, ampBps, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityNewPool is a paid mutator transaction binding the contract method 0x1c5d0a6b.
//
// Solidity: function addLiquidityNewPool(address tokenA, address tokenB, uint32 ampBps, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_KyberRouter *KyberRouterSession) AddLiquidityNewPool(tokenA common.Address, tokenB common.Address, ampBps uint32, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.AddLiquidityNewPool(&_KyberRouter.TransactOpts, tokenA, tokenB, ampBps, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityNewPool is a paid mutator transaction binding the contract method 0x1c5d0a6b.
//
// Solidity: function addLiquidityNewPool(address tokenA, address tokenB, uint32 ampBps, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_KyberRouter *KyberRouterTransactorSession) AddLiquidityNewPool(tokenA common.Address, tokenB common.Address, ampBps uint32, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.AddLiquidityNewPool(&_KyberRouter.TransactOpts, tokenA, tokenB, ampBps, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityNewPoolETH is a paid mutator transaction binding the contract method 0x7d41a422.
//
// Solidity: function addLiquidityNewPoolETH(address token, uint32 ampBps, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_KyberRouter *KyberRouterTransactor) AddLiquidityNewPoolETH(opts *bind.TransactOpts, token common.Address, ampBps uint32, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "addLiquidityNewPoolETH", token, ampBps, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityNewPoolETH is a paid mutator transaction binding the contract method 0x7d41a422.
//
// Solidity: function addLiquidityNewPoolETH(address token, uint32 ampBps, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_KyberRouter *KyberRouterSession) AddLiquidityNewPoolETH(token common.Address, ampBps uint32, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.AddLiquidityNewPoolETH(&_KyberRouter.TransactOpts, token, ampBps, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityNewPoolETH is a paid mutator transaction binding the contract method 0x7d41a422.
//
// Solidity: function addLiquidityNewPoolETH(address token, uint32 ampBps, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_KyberRouter *KyberRouterTransactorSession) AddLiquidityNewPoolETH(token common.Address, ampBps uint32, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.AddLiquidityNewPoolETH(&_KyberRouter.TransactOpts, token, ampBps, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe1f4a784.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, address pool, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_KyberRouter *KyberRouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, pool common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, pool, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe1f4a784.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, address pool, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_KyberRouter *KyberRouterSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, pool common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidity(&_KyberRouter.TransactOpts, tokenA, tokenB, pool, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe1f4a784.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, address pool, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_KyberRouter *KyberRouterTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, pool common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidity(&_KyberRouter.TransactOpts, tokenA, tokenB, pool, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xb56b681d.
//
// Solidity: function removeLiquidityETH(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_KyberRouter *KyberRouterTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "removeLiquidityETH", token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xb56b681d.
//
// Solidity: function removeLiquidityETH(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_KyberRouter *KyberRouterSession) RemoveLiquidityETH(token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidityETH(&_KyberRouter.TransactOpts, token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xb56b681d.
//
// Solidity: function removeLiquidityETH(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_KyberRouter *KyberRouterTransactorSession) RemoveLiquidityETH(token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidityETH(&_KyberRouter.TransactOpts, token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xebb5d2e9.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_KyberRouter *KyberRouterTransactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xebb5d2e9.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_KyberRouter *KyberRouterSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_KyberRouter.TransactOpts, token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xebb5d2e9.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (_KyberRouter *KyberRouterTransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_KyberRouter.TransactOpts, token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xa4aabb08.
//
// Solidity: function removeLiquidityETHWithPermit(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_KyberRouter *KyberRouterTransactor) RemoveLiquidityETHWithPermit(opts *bind.TransactOpts, token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "removeLiquidityETHWithPermit", token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xa4aabb08.
//
// Solidity: function removeLiquidityETHWithPermit(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_KyberRouter *KyberRouterSession) RemoveLiquidityETHWithPermit(token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidityETHWithPermit(&_KyberRouter.TransactOpts, token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0xa4aabb08.
//
// Solidity: function removeLiquidityETHWithPermit(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_KyberRouter *KyberRouterTransactorSession) RemoveLiquidityETHWithPermit(token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidityETHWithPermit(&_KyberRouter.TransactOpts, token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x0e2f024c.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_KyberRouter *KyberRouterTransactor) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x0e2f024c.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_KyberRouter *KyberRouterSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_KyberRouter.TransactOpts, token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x0e2f024c.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, address pool, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (_KyberRouter *KyberRouterTransactorSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, pool common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_KyberRouter.TransactOpts, token, pool, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x4c17fd7c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, address pool, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_KyberRouter *KyberRouterTransactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, pool common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, pool, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x4c17fd7c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, address pool, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_KyberRouter *KyberRouterSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, pool common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidityWithPermit(&_KyberRouter.TransactOpts, tokenA, tokenB, pool, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x4c17fd7c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, address pool, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_KyberRouter *KyberRouterTransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, pool common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KyberRouter.Contract.RemoveLiquidityWithPermit(&_KyberRouter.TransactOpts, tokenA, tokenB, pool, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfc5b8bce.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "swapETHForExactTokens", amountOut, poolsPath, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfc5b8bce.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterSession) SwapETHForExactTokens(amountOut *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapETHForExactTokens(&_KyberRouter.TransactOpts, amountOut, poolsPath, path, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfc5b8bce.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactorSession) SwapETHForExactTokens(amountOut *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapETHForExactTokens(&_KyberRouter.TransactOpts, amountOut, poolsPath, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0xaf55e31f.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0xaf55e31f.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterSession) SwapExactETHForTokens(amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactETHForTokens(&_KyberRouter.TransactOpts, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0xaf55e31f.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactETHForTokens(&_KyberRouter.TransactOpts, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x2daaa818.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns()
func (_KyberRouter *KyberRouterTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x2daaa818.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns()
func (_KyberRouter *KyberRouterSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_KyberRouter.TransactOpts, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x2daaa818.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns()
func (_KyberRouter *KyberRouterTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_KyberRouter.TransactOpts, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x6dce49ae.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x6dce49ae.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactTokensForETH(&_KyberRouter.TransactOpts, amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x6dce49ae.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactTokensForETH(&_KyberRouter.TransactOpts, amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3e741fca.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns()
func (_KyberRouter *KyberRouterTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3e741fca.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns()
func (_KyberRouter *KyberRouterSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_KyberRouter.TransactOpts, amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3e741fca.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns()
func (_KyberRouter *KyberRouterTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_KyberRouter.TransactOpts, amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xceb757d5.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xceb757d5.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactTokensForTokens(&_KyberRouter.TransactOpts, amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xceb757d5.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactTokensForTokens(&_KyberRouter.TransactOpts, amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xe8898b5f.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns()
func (_KyberRouter *KyberRouterTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xe8898b5f.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns()
func (_KyberRouter *KyberRouterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_KyberRouter.TransactOpts, amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xe8898b5f.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] poolsPath, address[] path, address to, uint256 deadline) returns()
func (_KyberRouter *KyberRouterTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_KyberRouter.TransactOpts, amountIn, amountOutMin, poolsPath, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x89c27594.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, poolsPath, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x89c27594.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapTokensForExactETH(&_KyberRouter.TransactOpts, amountOut, amountInMax, poolsPath, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x89c27594.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapTokensForExactETH(&_KyberRouter.TransactOpts, amountOut, amountInMax, poolsPath, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xae8290b7.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, poolsPath, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xae8290b7.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapTokensForExactTokens(&_KyberRouter.TransactOpts, amountOut, amountInMax, poolsPath, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xae8290b7.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] poolsPath, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_KyberRouter *KyberRouterTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, poolsPath []common.Address, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _KyberRouter.Contract.SwapTokensForExactTokens(&_KyberRouter.TransactOpts, amountOut, amountInMax, poolsPath, path, to, deadline)
}
