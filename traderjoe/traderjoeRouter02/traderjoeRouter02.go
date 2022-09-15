// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderjoeRouter02

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

// TraderjoeRouter02MetaData contains all meta data concerning the TraderjoeRouter02 contract.
var TraderjoeRouter02MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WAVAX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityAVAXSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityAVAXWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAVAXForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAXSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TraderjoeRouter02ABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderjoeRouter02MetaData.ABI instead.
var TraderjoeRouter02ABI = TraderjoeRouter02MetaData.ABI

// TraderjoeRouter02 is an auto generated Go binding around an Ethereum contract.
type TraderjoeRouter02 struct {
	TraderjoeRouter02Caller     // Read-only binding to the contract
	TraderjoeRouter02Transactor // Write-only binding to the contract
	TraderjoeRouter02Filterer   // Log filterer for contract events
}

// TraderjoeRouter02Caller is an auto generated read-only Go binding around an Ethereum contract.
type TraderjoeRouter02Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeRouter02Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderjoeRouter02Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeRouter02Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderjoeRouter02Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeRouter02Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderjoeRouter02Session struct {
	Contract     *TraderjoeRouter02 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TraderjoeRouter02CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderjoeRouter02CallerSession struct {
	Contract *TraderjoeRouter02Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TraderjoeRouter02TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderjoeRouter02TransactorSession struct {
	Contract     *TraderjoeRouter02Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TraderjoeRouter02Raw is an auto generated low-level Go binding around an Ethereum contract.
type TraderjoeRouter02Raw struct {
	Contract *TraderjoeRouter02 // Generic contract binding to access the raw methods on
}

// TraderjoeRouter02CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderjoeRouter02CallerRaw struct {
	Contract *TraderjoeRouter02Caller // Generic read-only contract binding to access the raw methods on
}

// TraderjoeRouter02TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderjoeRouter02TransactorRaw struct {
	Contract *TraderjoeRouter02Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderjoeRouter02 creates a new instance of TraderjoeRouter02, bound to a specific deployed contract.
func NewTraderjoeRouter02(address common.Address, backend bind.ContractBackend) (*TraderjoeRouter02, error) {
	contract, err := bindTraderjoeRouter02(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraderjoeRouter02{TraderjoeRouter02Caller: TraderjoeRouter02Caller{contract: contract}, TraderjoeRouter02Transactor: TraderjoeRouter02Transactor{contract: contract}, TraderjoeRouter02Filterer: TraderjoeRouter02Filterer{contract: contract}}, nil
}

// NewTraderjoeRouter02Caller creates a new read-only instance of TraderjoeRouter02, bound to a specific deployed contract.
func NewTraderjoeRouter02Caller(address common.Address, caller bind.ContractCaller) (*TraderjoeRouter02Caller, error) {
	contract, err := bindTraderjoeRouter02(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeRouter02Caller{contract: contract}, nil
}

// NewTraderjoeRouter02Transactor creates a new write-only instance of TraderjoeRouter02, bound to a specific deployed contract.
func NewTraderjoeRouter02Transactor(address common.Address, transactor bind.ContractTransactor) (*TraderjoeRouter02Transactor, error) {
	contract, err := bindTraderjoeRouter02(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeRouter02Transactor{contract: contract}, nil
}

// NewTraderjoeRouter02Filterer creates a new log filterer instance of TraderjoeRouter02, bound to a specific deployed contract.
func NewTraderjoeRouter02Filterer(address common.Address, filterer bind.ContractFilterer) (*TraderjoeRouter02Filterer, error) {
	contract, err := bindTraderjoeRouter02(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderjoeRouter02Filterer{contract: contract}, nil
}

// bindTraderjoeRouter02 binds a generic wrapper to an already deployed contract.
func bindTraderjoeRouter02(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraderjoeRouter02ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeRouter02 *TraderjoeRouter02Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeRouter02.Contract.TraderjoeRouter02Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeRouter02 *TraderjoeRouter02Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.TraderjoeRouter02Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeRouter02 *TraderjoeRouter02Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.TraderjoeRouter02Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeRouter02 *TraderjoeRouter02CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeRouter02.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.contract.Transact(opts, method, params...)
}

// WAVAX is a free data retrieval call binding the contract method 0x73b295c2.
//
// Solidity: function WAVAX() pure returns(address)
func (_TraderjoeRouter02 *TraderjoeRouter02Caller) WAVAX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeRouter02.contract.Call(opts, &out, "WAVAX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WAVAX is a free data retrieval call binding the contract method 0x73b295c2.
//
// Solidity: function WAVAX() pure returns(address)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) WAVAX() (common.Address, error) {
	return _TraderjoeRouter02.Contract.WAVAX(&_TraderjoeRouter02.CallOpts)
}

// WAVAX is a free data retrieval call binding the contract method 0x73b295c2.
//
// Solidity: function WAVAX() pure returns(address)
func (_TraderjoeRouter02 *TraderjoeRouter02CallerSession) WAVAX() (common.Address, error) {
	return _TraderjoeRouter02.Contract.WAVAX(&_TraderjoeRouter02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_TraderjoeRouter02 *TraderjoeRouter02Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeRouter02.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) Factory() (common.Address, error) {
	return _TraderjoeRouter02.Contract.Factory(&_TraderjoeRouter02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_TraderjoeRouter02 *TraderjoeRouter02CallerSession) Factory() (common.Address, error) {
	return _TraderjoeRouter02.Contract.Factory(&_TraderjoeRouter02.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_TraderjoeRouter02 *TraderjoeRouter02Caller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRouter02.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _TraderjoeRouter02.Contract.GetAmountIn(&_TraderjoeRouter02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_TraderjoeRouter02 *TraderjoeRouter02CallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _TraderjoeRouter02.Contract.GetAmountIn(&_TraderjoeRouter02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_TraderjoeRouter02 *TraderjoeRouter02Caller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRouter02.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _TraderjoeRouter02.Contract.GetAmountOut(&_TraderjoeRouter02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_TraderjoeRouter02 *TraderjoeRouter02CallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _TraderjoeRouter02.Contract.GetAmountOut(&_TraderjoeRouter02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Caller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRouter02.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _TraderjoeRouter02.Contract.GetAmountsIn(&_TraderjoeRouter02.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02CallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _TraderjoeRouter02.Contract.GetAmountsIn(&_TraderjoeRouter02.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRouter02.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _TraderjoeRouter02.Contract.GetAmountsOut(&_TraderjoeRouter02.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _TraderjoeRouter02.Contract.GetAmountsOut(&_TraderjoeRouter02.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_TraderjoeRouter02 *TraderjoeRouter02Caller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRouter02.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _TraderjoeRouter02.Contract.Quote(&_TraderjoeRouter02.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_TraderjoeRouter02 *TraderjoeRouter02CallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _TraderjoeRouter02.Contract.Quote(&_TraderjoeRouter02.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.AddLiquidity(&_TraderjoeRouter02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.AddLiquidity(&_TraderjoeRouter02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) AddLiquidityAVAX(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "addLiquidityAVAX", token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) AddLiquidityAVAX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.AddLiquidityAVAX(&_TraderjoeRouter02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) AddLiquidityAVAX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.AddLiquidityAVAX(&_TraderjoeRouter02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidity(&_TraderjoeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidity(&_TraderjoeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0x33c6b725.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) RemoveLiquidityAVAX(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "removeLiquidityAVAX", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0x33c6b725.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) RemoveLiquidityAVAX(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidityAVAX(&_TraderjoeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0x33c6b725.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) RemoveLiquidityAVAX(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidityAVAX(&_TraderjoeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x73bc79cf.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) RemoveLiquidityAVAXSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "removeLiquidityAVAXSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x73bc79cf.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) RemoveLiquidityAVAXSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidityAVAXSupportingFeeOnTransferTokens(&_TraderjoeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x73bc79cf.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) RemoveLiquidityAVAXSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidityAVAXSupportingFeeOnTransferTokens(&_TraderjoeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXWithPermit is a paid mutator transaction binding the contract method 0x2c407024.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) RemoveLiquidityAVAXWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "removeLiquidityAVAXWithPermit", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermit is a paid mutator transaction binding the contract method 0x2c407024.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) RemoveLiquidityAVAXWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidityAVAXWithPermit(&_TraderjoeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermit is a paid mutator transaction binding the contract method 0x2c407024.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) RemoveLiquidityAVAXWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidityAVAXWithPermit(&_TraderjoeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9fc27226.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9fc27226.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(&_TraderjoeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9fc27226.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(&_TraderjoeRouter02.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidityWithPermit(&_TraderjoeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.RemoveLiquidityWithPermit(&_TraderjoeRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) SwapAVAXForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "swapAVAXForExactTokens", amountOut, path, to, deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) SwapAVAXForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapAVAXForExactTokens(&_TraderjoeRouter02.TransactOpts, amountOut, path, to, deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) SwapAVAXForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapAVAXForExactTokens(&_TraderjoeRouter02.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) SwapExactAVAXForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "swapExactAVAXForTokens", amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) SwapExactAVAXForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactAVAXForTokens(&_TraderjoeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) SwapExactAVAXForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactAVAXForTokens(&_TraderjoeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc57559dd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "swapExactAVAXForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc57559dd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_TraderjoeRouter02 *TraderjoeRouter02Session) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactAVAXForTokensSupportingFeeOnTransferTokens(&_TraderjoeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc57559dd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactAVAXForTokensSupportingFeeOnTransferTokens(&_TraderjoeRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) SwapExactTokensForAVAX(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "swapExactTokensForAVAX", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) SwapExactTokensForAVAX(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactTokensForAVAX(&_TraderjoeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) SwapExactTokensForAVAX(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactTokensForAVAX(&_TraderjoeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x762b1562.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "swapExactTokensForAVAXSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x762b1562.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_TraderjoeRouter02 *TraderjoeRouter02Session) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactTokensForAVAXSupportingFeeOnTransferTokens(&_TraderjoeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x762b1562.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactTokensForAVAXSupportingFeeOnTransferTokens(&_TraderjoeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactTokensForTokens(&_TraderjoeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactTokensForTokens(&_TraderjoeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_TraderjoeRouter02 *TraderjoeRouter02Session) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_TraderjoeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_TraderjoeRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) SwapTokensForExactAVAX(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "swapTokensForExactAVAX", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) SwapTokensForExactAVAX(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapTokensForExactAVAX(&_TraderjoeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) SwapTokensForExactAVAX(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapTokensForExactAVAX(&_TraderjoeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapTokensForExactTokens(&_TraderjoeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_TraderjoeRouter02 *TraderjoeRouter02TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _TraderjoeRouter02.Contract.SwapTokensForExactTokens(&_TraderjoeRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}
