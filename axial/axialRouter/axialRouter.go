// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package axialRouter

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

// AxialRouterMetaData contains all meta data concerning the AxialRouter contract.
var AxialRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minToMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"}],\"name\":\"calculateRemoveLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"deposit\",\"type\":\"bool\"}],\"name\":\"calculateTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashLoanFeeBPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAPrecise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenIndex\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"maxBurnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AxialRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use AxialRouterMetaData.ABI instead.
var AxialRouterABI = AxialRouterMetaData.ABI

// AxialRouter is an auto generated Go binding around an Ethereum contract.
type AxialRouter struct {
	AxialRouterCaller     // Read-only binding to the contract
	AxialRouterTransactor // Write-only binding to the contract
	AxialRouterFilterer   // Log filterer for contract events
}

// AxialRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AxialRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxialRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AxialRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxialRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AxialRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxialRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AxialRouterSession struct {
	Contract     *AxialRouter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AxialRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AxialRouterCallerSession struct {
	Contract *AxialRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AxialRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AxialRouterTransactorSession struct {
	Contract     *AxialRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AxialRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AxialRouterRaw struct {
	Contract *AxialRouter // Generic contract binding to access the raw methods on
}

// AxialRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AxialRouterCallerRaw struct {
	Contract *AxialRouterCaller // Generic read-only contract binding to access the raw methods on
}

// AxialRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AxialRouterTransactorRaw struct {
	Contract *AxialRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAxialRouter creates a new instance of AxialRouter, bound to a specific deployed contract.
func NewAxialRouter(address common.Address, backend bind.ContractBackend) (*AxialRouter, error) {
	contract, err := bindAxialRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AxialRouter{AxialRouterCaller: AxialRouterCaller{contract: contract}, AxialRouterTransactor: AxialRouterTransactor{contract: contract}, AxialRouterFilterer: AxialRouterFilterer{contract: contract}}, nil
}

// NewAxialRouterCaller creates a new read-only instance of AxialRouter, bound to a specific deployed contract.
func NewAxialRouterCaller(address common.Address, caller bind.ContractCaller) (*AxialRouterCaller, error) {
	contract, err := bindAxialRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AxialRouterCaller{contract: contract}, nil
}

// NewAxialRouterTransactor creates a new write-only instance of AxialRouter, bound to a specific deployed contract.
func NewAxialRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*AxialRouterTransactor, error) {
	contract, err := bindAxialRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AxialRouterTransactor{contract: contract}, nil
}

// NewAxialRouterFilterer creates a new log filterer instance of AxialRouter, bound to a specific deployed contract.
func NewAxialRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*AxialRouterFilterer, error) {
	contract, err := bindAxialRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AxialRouterFilterer{contract: contract}, nil
}

// bindAxialRouter binds a generic wrapper to an already deployed contract.
func bindAxialRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AxialRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AxialRouter *AxialRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AxialRouter.Contract.AxialRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AxialRouter *AxialRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AxialRouter.Contract.AxialRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AxialRouter *AxialRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AxialRouter.Contract.AxialRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AxialRouter *AxialRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AxialRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AxialRouter *AxialRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AxialRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AxialRouter *AxialRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AxialRouter.Contract.contract.Transact(opts, method, params...)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_AxialRouter *AxialRouterCaller) CalculateRemoveLiquidity(opts *bind.CallOpts, amount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _AxialRouter.contract.Call(opts, &out, "calculateRemoveLiquidity", amount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_AxialRouter *AxialRouterSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _AxialRouter.Contract.CalculateRemoveLiquidity(&_AxialRouter.CallOpts, amount)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_AxialRouter *AxialRouterCallerSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _AxialRouter.Contract.CalculateRemoveLiquidity(&_AxialRouter.CallOpts, amount)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_AxialRouter *AxialRouterCaller) CalculateRemoveLiquidityOneToken(opts *bind.CallOpts, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _AxialRouter.contract.Call(opts, &out, "calculateRemoveLiquidityOneToken", tokenAmount, tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_AxialRouter *AxialRouterSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _AxialRouter.Contract.CalculateRemoveLiquidityOneToken(&_AxialRouter.CallOpts, tokenAmount, tokenIndex)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_AxialRouter *AxialRouterCallerSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _AxialRouter.Contract.CalculateRemoveLiquidityOneToken(&_AxialRouter.CallOpts, tokenAmount, tokenIndex)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_AxialRouter *AxialRouterCaller) CalculateSwap(opts *bind.CallOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AxialRouter.contract.Call(opts, &out, "calculateSwap", tokenIndexFrom, tokenIndexTo, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_AxialRouter *AxialRouterSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _AxialRouter.Contract.CalculateSwap(&_AxialRouter.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_AxialRouter *AxialRouterCallerSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _AxialRouter.Contract.CalculateSwap(&_AxialRouter.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_AxialRouter *AxialRouterCaller) CalculateTokenAmount(opts *bind.CallOpts, amounts []*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _AxialRouter.contract.Call(opts, &out, "calculateTokenAmount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_AxialRouter *AxialRouterSession) CalculateTokenAmount(amounts []*big.Int, deposit bool) (*big.Int, error) {
	return _AxialRouter.Contract.CalculateTokenAmount(&_AxialRouter.CallOpts, amounts, deposit)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_AxialRouter *AxialRouterCallerSession) CalculateTokenAmount(amounts []*big.Int, deposit bool) (*big.Int, error) {
	return _AxialRouter.Contract.CalculateTokenAmount(&_AxialRouter.CallOpts, amounts, deposit)
}

// FlashLoanFeeBPS is a free data retrieval call binding the contract method 0x7f1c825a.
//
// Solidity: function flashLoanFeeBPS() view returns(uint256)
func (_AxialRouter *AxialRouterCaller) FlashLoanFeeBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AxialRouter.contract.Call(opts, &out, "flashLoanFeeBPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashLoanFeeBPS is a free data retrieval call binding the contract method 0x7f1c825a.
//
// Solidity: function flashLoanFeeBPS() view returns(uint256)
func (_AxialRouter *AxialRouterSession) FlashLoanFeeBPS() (*big.Int, error) {
	return _AxialRouter.Contract.FlashLoanFeeBPS(&_AxialRouter.CallOpts)
}

// FlashLoanFeeBPS is a free data retrieval call binding the contract method 0x7f1c825a.
//
// Solidity: function flashLoanFeeBPS() view returns(uint256)
func (_AxialRouter *AxialRouterCallerSession) FlashLoanFeeBPS() (*big.Int, error) {
	return _AxialRouter.Contract.FlashLoanFeeBPS(&_AxialRouter.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_AxialRouter *AxialRouterCaller) GetA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AxialRouter.contract.Call(opts, &out, "getA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_AxialRouter *AxialRouterSession) GetA() (*big.Int, error) {
	return _AxialRouter.Contract.GetA(&_AxialRouter.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_AxialRouter *AxialRouterCallerSession) GetA() (*big.Int, error) {
	return _AxialRouter.Contract.GetA(&_AxialRouter.CallOpts)
}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_AxialRouter *AxialRouterCaller) GetAPrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AxialRouter.contract.Call(opts, &out, "getAPrecise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_AxialRouter *AxialRouterSession) GetAPrecise() (*big.Int, error) {
	return _AxialRouter.Contract.GetAPrecise(&_AxialRouter.CallOpts)
}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_AxialRouter *AxialRouterCallerSession) GetAPrecise() (*big.Int, error) {
	return _AxialRouter.Contract.GetAPrecise(&_AxialRouter.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_AxialRouter *AxialRouterCaller) GetToken(opts *bind.CallOpts, index uint8) (common.Address, error) {
	var out []interface{}
	err := _AxialRouter.contract.Call(opts, &out, "getToken", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_AxialRouter *AxialRouterSession) GetToken(index uint8) (common.Address, error) {
	return _AxialRouter.Contract.GetToken(&_AxialRouter.CallOpts, index)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_AxialRouter *AxialRouterCallerSession) GetToken(index uint8) (common.Address, error) {
	return _AxialRouter.Contract.GetToken(&_AxialRouter.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_AxialRouter *AxialRouterCaller) GetTokenBalance(opts *bind.CallOpts, index uint8) (*big.Int, error) {
	var out []interface{}
	err := _AxialRouter.contract.Call(opts, &out, "getTokenBalance", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_AxialRouter *AxialRouterSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _AxialRouter.Contract.GetTokenBalance(&_AxialRouter.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_AxialRouter *AxialRouterCallerSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _AxialRouter.Contract.GetTokenBalance(&_AxialRouter.CallOpts, index)
}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_AxialRouter *AxialRouterCaller) GetTokenIndex(opts *bind.CallOpts, tokenAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _AxialRouter.contract.Call(opts, &out, "getTokenIndex", tokenAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_AxialRouter *AxialRouterSession) GetTokenIndex(tokenAddress common.Address) (uint8, error) {
	return _AxialRouter.Contract.GetTokenIndex(&_AxialRouter.CallOpts, tokenAddress)
}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_AxialRouter *AxialRouterCallerSession) GetTokenIndex(tokenAddress common.Address) (uint8, error) {
	return _AxialRouter.Contract.GetTokenIndex(&_AxialRouter.CallOpts, tokenAddress)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterTransactor) AddLiquidity(opts *bind.TransactOpts, amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.contract.Transact(opts, "addLiquidity", amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.Contract.AddLiquidity(&_AxialRouter.TransactOpts, amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterTransactorSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.Contract.AddLiquidity(&_AxialRouter.TransactOpts, amounts, minToMint, deadline)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes params) returns()
func (_AxialRouter *AxialRouterTransactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _AxialRouter.contract.Transact(opts, "flashLoan", receiver, token, amount, params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes params) returns()
func (_AxialRouter *AxialRouterSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _AxialRouter.Contract.FlashLoan(&_AxialRouter.TransactOpts, receiver, token, amount, params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes params) returns()
func (_AxialRouter *AxialRouterTransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _AxialRouter.Contract.FlashLoan(&_AxialRouter.TransactOpts, receiver, token, amount, params)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_AxialRouter *AxialRouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.contract.Transact(opts, "removeLiquidity", amount, minAmounts, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_AxialRouter *AxialRouterSession) RemoveLiquidity(amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.Contract.RemoveLiquidity(&_AxialRouter.TransactOpts, amount, minAmounts, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_AxialRouter *AxialRouterTransactorSession) RemoveLiquidity(amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.Contract.RemoveLiquidity(&_AxialRouter.TransactOpts, amount, minAmounts, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.contract.Transact(opts, "removeLiquidityImbalance", amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterSession) RemoveLiquidityImbalance(amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.Contract.RemoveLiquidityImbalance(&_AxialRouter.TransactOpts, amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterTransactorSession) RemoveLiquidityImbalance(amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.Contract.RemoveLiquidityImbalance(&_AxialRouter.TransactOpts, amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterTransactor) RemoveLiquidityOneToken(opts *bind.TransactOpts, tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.contract.Transact(opts, "removeLiquidityOneToken", tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.Contract.RemoveLiquidityOneToken(&_AxialRouter.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterTransactorSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.Contract.RemoveLiquidityOneToken(&_AxialRouter.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterTransactor) Swap(opts *bind.TransactOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.contract.Transact(opts, "swap", tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.Contract.Swap(&_AxialRouter.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_AxialRouter *AxialRouterTransactorSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AxialRouter.Contract.Swap(&_AxialRouter.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}
