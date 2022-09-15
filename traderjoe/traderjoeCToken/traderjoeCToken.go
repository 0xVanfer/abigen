// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderjoeCToken

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

// TraderjoeCTokenMetaData contains all meta data concerning the TraderjoeCToken contract.
var TraderjoeCTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCollateralCap\",\"type\":\"uint256\"}],\"name\":\"_setCollateralCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowResign\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"name\":\"_setImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newJoetroller\",\"type\":\"address\"}],\"name\":\"_setJoetroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountCollateralTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrualBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowRatePerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToViewImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashFeeBips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gulp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"internalCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isCollateralTokenInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isJToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joetroller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"jTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"registerCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCollateralTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"unregisterCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TraderjoeCTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderjoeCTokenMetaData.ABI instead.
var TraderjoeCTokenABI = TraderjoeCTokenMetaData.ABI

// TraderjoeCToken is an auto generated Go binding around an Ethereum contract.
type TraderjoeCToken struct {
	TraderjoeCTokenCaller     // Read-only binding to the contract
	TraderjoeCTokenTransactor // Write-only binding to the contract
	TraderjoeCTokenFilterer   // Log filterer for contract events
}

// TraderjoeCTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderjoeCTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeCTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderjoeCTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeCTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderjoeCTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeCTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderjoeCTokenSession struct {
	Contract     *TraderjoeCToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraderjoeCTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderjoeCTokenCallerSession struct {
	Contract *TraderjoeCTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TraderjoeCTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderjoeCTokenTransactorSession struct {
	Contract     *TraderjoeCTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TraderjoeCTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderjoeCTokenRaw struct {
	Contract *TraderjoeCToken // Generic contract binding to access the raw methods on
}

// TraderjoeCTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderjoeCTokenCallerRaw struct {
	Contract *TraderjoeCTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TraderjoeCTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderjoeCTokenTransactorRaw struct {
	Contract *TraderjoeCTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderjoeCToken creates a new instance of TraderjoeCToken, bound to a specific deployed contract.
func NewTraderjoeCToken(address common.Address, backend bind.ContractBackend) (*TraderjoeCToken, error) {
	contract, err := bindTraderjoeCToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraderjoeCToken{TraderjoeCTokenCaller: TraderjoeCTokenCaller{contract: contract}, TraderjoeCTokenTransactor: TraderjoeCTokenTransactor{contract: contract}, TraderjoeCTokenFilterer: TraderjoeCTokenFilterer{contract: contract}}, nil
}

// NewTraderjoeCTokenCaller creates a new read-only instance of TraderjoeCToken, bound to a specific deployed contract.
func NewTraderjoeCTokenCaller(address common.Address, caller bind.ContractCaller) (*TraderjoeCTokenCaller, error) {
	contract, err := bindTraderjoeCToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeCTokenCaller{contract: contract}, nil
}

// NewTraderjoeCTokenTransactor creates a new write-only instance of TraderjoeCToken, bound to a specific deployed contract.
func NewTraderjoeCTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderjoeCTokenTransactor, error) {
	contract, err := bindTraderjoeCToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeCTokenTransactor{contract: contract}, nil
}

// NewTraderjoeCTokenFilterer creates a new log filterer instance of TraderjoeCToken, bound to a specific deployed contract.
func NewTraderjoeCTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderjoeCTokenFilterer, error) {
	contract, err := bindTraderjoeCToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderjoeCTokenFilterer{contract: contract}, nil
}

// bindTraderjoeCToken binds a generic wrapper to an already deployed contract.
func bindTraderjoeCToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraderjoeCTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeCToken *TraderjoeCTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeCToken.Contract.TraderjoeCTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeCToken *TraderjoeCTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.TraderjoeCTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeCToken *TraderjoeCTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.TraderjoeCTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeCToken *TraderjoeCTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeCToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeCToken *TraderjoeCTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeCToken *TraderjoeCTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.contract.Transact(opts, method, params...)
}

// AccountCollateralTokens is a free data retrieval call binding the contract method 0xd240d64a.
//
// Solidity: function accountCollateralTokens(address ) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) AccountCollateralTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "accountCollateralTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountCollateralTokens is a free data retrieval call binding the contract method 0xd240d64a.
//
// Solidity: function accountCollateralTokens(address ) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) AccountCollateralTokens(arg0 common.Address) (*big.Int, error) {
	return _TraderjoeCToken.Contract.AccountCollateralTokens(&_TraderjoeCToken.CallOpts, arg0)
}

// AccountCollateralTokens is a free data retrieval call binding the contract method 0xd240d64a.
//
// Solidity: function accountCollateralTokens(address ) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) AccountCollateralTokens(arg0 common.Address) (*big.Int, error) {
	return _TraderjoeCToken.Contract.AccountCollateralTokens(&_TraderjoeCToken.CallOpts, arg0)
}

// AccrualBlockTimestamp is a free data retrieval call binding the contract method 0xcfa99201.
//
// Solidity: function accrualBlockTimestamp() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) AccrualBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "accrualBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockTimestamp is a free data retrieval call binding the contract method 0xcfa99201.
//
// Solidity: function accrualBlockTimestamp() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) AccrualBlockTimestamp() (*big.Int, error) {
	return _TraderjoeCToken.Contract.AccrualBlockTimestamp(&_TraderjoeCToken.CallOpts)
}

// AccrualBlockTimestamp is a free data retrieval call binding the contract method 0xcfa99201.
//
// Solidity: function accrualBlockTimestamp() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) AccrualBlockTimestamp() (*big.Int, error) {
	return _TraderjoeCToken.Contract.AccrualBlockTimestamp(&_TraderjoeCToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenSession) Admin() (common.Address, error) {
	return _TraderjoeCToken.Contract.Admin(&_TraderjoeCToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) Admin() (common.Address, error) {
	return _TraderjoeCToken.Contract.Admin(&_TraderjoeCToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TraderjoeCToken.Contract.Allowance(&_TraderjoeCToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TraderjoeCToken.Contract.Allowance(&_TraderjoeCToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TraderjoeCToken.Contract.BalanceOf(&_TraderjoeCToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TraderjoeCToken.Contract.BalanceOf(&_TraderjoeCToken.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _TraderjoeCToken.Contract.BorrowBalanceStored(&_TraderjoeCToken.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _TraderjoeCToken.Contract.BorrowBalanceStored(&_TraderjoeCToken.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) BorrowIndex() (*big.Int, error) {
	return _TraderjoeCToken.Contract.BorrowIndex(&_TraderjoeCToken.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) BorrowIndex() (*big.Int, error) {
	return _TraderjoeCToken.Contract.BorrowIndex(&_TraderjoeCToken.CallOpts)
}

// BorrowRatePerSecond is a free data retrieval call binding the contract method 0x52609750.
//
// Solidity: function borrowRatePerSecond() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) BorrowRatePerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "borrowRatePerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerSecond is a free data retrieval call binding the contract method 0x52609750.
//
// Solidity: function borrowRatePerSecond() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) BorrowRatePerSecond() (*big.Int, error) {
	return _TraderjoeCToken.Contract.BorrowRatePerSecond(&_TraderjoeCToken.CallOpts)
}

// BorrowRatePerSecond is a free data retrieval call binding the contract method 0x52609750.
//
// Solidity: function borrowRatePerSecond() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) BorrowRatePerSecond() (*big.Int, error) {
	return _TraderjoeCToken.Contract.BorrowRatePerSecond(&_TraderjoeCToken.CallOpts)
}

// CollateralCap is a free data retrieval call binding the contract method 0xd2bb18e9.
//
// Solidity: function collateralCap() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) CollateralCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "collateralCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralCap is a free data retrieval call binding the contract method 0xd2bb18e9.
//
// Solidity: function collateralCap() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) CollateralCap() (*big.Int, error) {
	return _TraderjoeCToken.Contract.CollateralCap(&_TraderjoeCToken.CallOpts)
}

// CollateralCap is a free data retrieval call binding the contract method 0xd2bb18e9.
//
// Solidity: function collateralCap() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) CollateralCap() (*big.Int, error) {
	return _TraderjoeCToken.Contract.CollateralCap(&_TraderjoeCToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TraderjoeCToken *TraderjoeCTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TraderjoeCToken *TraderjoeCTokenSession) Decimals() (uint8, error) {
	return _TraderjoeCToken.Contract.Decimals(&_TraderjoeCToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) Decimals() (uint8, error) {
	return _TraderjoeCToken.Contract.Decimals(&_TraderjoeCToken.CallOpts)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_TraderjoeCToken *TraderjoeCTokenCaller) DelegateToViewImplementation(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "delegateToViewImplementation", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_TraderjoeCToken *TraderjoeCTokenSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _TraderjoeCToken.Contract.DelegateToViewImplementation(&_TraderjoeCToken.CallOpts, data)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _TraderjoeCToken.Contract.DelegateToViewImplementation(&_TraderjoeCToken.CallOpts, data)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) ExchangeRateStored() (*big.Int, error) {
	return _TraderjoeCToken.Contract.ExchangeRateStored(&_TraderjoeCToken.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _TraderjoeCToken.Contract.ExchangeRateStored(&_TraderjoeCToken.CallOpts)
}

// FlashFee is a free data retrieval call binding the contract method 0xa7af467a.
//
// Solidity: function flashFee(uint256 amount) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) FlashFee(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "flashFee", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xa7af467a.
//
// Solidity: function flashFee(uint256 amount) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) FlashFee(amount *big.Int) (*big.Int, error) {
	return _TraderjoeCToken.Contract.FlashFee(&_TraderjoeCToken.CallOpts, amount)
}

// FlashFee is a free data retrieval call binding the contract method 0xa7af467a.
//
// Solidity: function flashFee(uint256 amount) view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) FlashFee(amount *big.Int) (*big.Int, error) {
	return _TraderjoeCToken.Contract.FlashFee(&_TraderjoeCToken.CallOpts, amount)
}

// FlashFeeBips is a free data retrieval call binding the contract method 0xea11eea4.
//
// Solidity: function flashFeeBips() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) FlashFeeBips(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "flashFeeBips")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFeeBips is a free data retrieval call binding the contract method 0xea11eea4.
//
// Solidity: function flashFeeBips() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) FlashFeeBips() (*big.Int, error) {
	return _TraderjoeCToken.Contract.FlashFeeBips(&_TraderjoeCToken.CallOpts)
}

// FlashFeeBips is a free data retrieval call binding the contract method 0xea11eea4.
//
// Solidity: function flashFeeBips() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) FlashFeeBips() (*big.Int, error) {
	return _TraderjoeCToken.Contract.FlashFeeBips(&_TraderjoeCToken.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _TraderjoeCToken.Contract.GetAccountSnapshot(&_TraderjoeCToken.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _TraderjoeCToken.Contract.GetAccountSnapshot(&_TraderjoeCToken.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) GetCash() (*big.Int, error) {
	return _TraderjoeCToken.Contract.GetCash(&_TraderjoeCToken.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) GetCash() (*big.Int, error) {
	return _TraderjoeCToken.Contract.GetCash(&_TraderjoeCToken.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenSession) Implementation() (common.Address, error) {
	return _TraderjoeCToken.Contract.Implementation(&_TraderjoeCToken.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) Implementation() (common.Address, error) {
	return _TraderjoeCToken.Contract.Implementation(&_TraderjoeCToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenSession) InterestRateModel() (common.Address, error) {
	return _TraderjoeCToken.Contract.InterestRateModel(&_TraderjoeCToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) InterestRateModel() (common.Address, error) {
	return _TraderjoeCToken.Contract.InterestRateModel(&_TraderjoeCToken.CallOpts)
}

// InternalCash is a free data retrieval call binding the contract method 0x22abdbf5.
//
// Solidity: function internalCash() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) InternalCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "internalCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InternalCash is a free data retrieval call binding the contract method 0x22abdbf5.
//
// Solidity: function internalCash() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) InternalCash() (*big.Int, error) {
	return _TraderjoeCToken.Contract.InternalCash(&_TraderjoeCToken.CallOpts)
}

// InternalCash is a free data retrieval call binding the contract method 0x22abdbf5.
//
// Solidity: function internalCash() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) InternalCash() (*big.Int, error) {
	return _TraderjoeCToken.Contract.InternalCash(&_TraderjoeCToken.CallOpts)
}

// IsCollateralTokenInit is a free data retrieval call binding the contract method 0x85d8a2e6.
//
// Solidity: function isCollateralTokenInit(address ) view returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenCaller) IsCollateralTokenInit(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "isCollateralTokenInit", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralTokenInit is a free data retrieval call binding the contract method 0x85d8a2e6.
//
// Solidity: function isCollateralTokenInit(address ) view returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenSession) IsCollateralTokenInit(arg0 common.Address) (bool, error) {
	return _TraderjoeCToken.Contract.IsCollateralTokenInit(&_TraderjoeCToken.CallOpts, arg0)
}

// IsCollateralTokenInit is a free data retrieval call binding the contract method 0x85d8a2e6.
//
// Solidity: function isCollateralTokenInit(address ) view returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) IsCollateralTokenInit(arg0 common.Address) (bool, error) {
	return _TraderjoeCToken.Contract.IsCollateralTokenInit(&_TraderjoeCToken.CallOpts, arg0)
}

// IsJToken is a free data retrieval call binding the contract method 0x406de0b6.
//
// Solidity: function isJToken() view returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenCaller) IsJToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "isJToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsJToken is a free data retrieval call binding the contract method 0x406de0b6.
//
// Solidity: function isJToken() view returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenSession) IsJToken() (bool, error) {
	return _TraderjoeCToken.Contract.IsJToken(&_TraderjoeCToken.CallOpts)
}

// IsJToken is a free data retrieval call binding the contract method 0x406de0b6.
//
// Solidity: function isJToken() view returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) IsJToken() (bool, error) {
	return _TraderjoeCToken.Contract.IsJToken(&_TraderjoeCToken.CallOpts)
}

// Joetroller is a free data retrieval call binding the contract method 0x6330533c.
//
// Solidity: function joetroller() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCaller) Joetroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "joetroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Joetroller is a free data retrieval call binding the contract method 0x6330533c.
//
// Solidity: function joetroller() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenSession) Joetroller() (common.Address, error) {
	return _TraderjoeCToken.Contract.Joetroller(&_TraderjoeCToken.CallOpts)
}

// Joetroller is a free data retrieval call binding the contract method 0x6330533c.
//
// Solidity: function joetroller() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) Joetroller() (common.Address, error) {
	return _TraderjoeCToken.Contract.Joetroller(&_TraderjoeCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TraderjoeCToken *TraderjoeCTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TraderjoeCToken *TraderjoeCTokenSession) Name() (string, error) {
	return _TraderjoeCToken.Contract.Name(&_TraderjoeCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) Name() (string, error) {
	return _TraderjoeCToken.Contract.Name(&_TraderjoeCToken.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenSession) PendingAdmin() (common.Address, error) {
	return _TraderjoeCToken.Contract.PendingAdmin(&_TraderjoeCToken.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) PendingAdmin() (common.Address, error) {
	return _TraderjoeCToken.Contract.PendingAdmin(&_TraderjoeCToken.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) ReserveFactorMantissa() (*big.Int, error) {
	return _TraderjoeCToken.Contract.ReserveFactorMantissa(&_TraderjoeCToken.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _TraderjoeCToken.Contract.ReserveFactorMantissa(&_TraderjoeCToken.CallOpts)
}

// SupplyRatePerSecond is a free data retrieval call binding the contract method 0xb1d38974.
//
// Solidity: function supplyRatePerSecond() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) SupplyRatePerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "supplyRatePerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerSecond is a free data retrieval call binding the contract method 0xb1d38974.
//
// Solidity: function supplyRatePerSecond() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) SupplyRatePerSecond() (*big.Int, error) {
	return _TraderjoeCToken.Contract.SupplyRatePerSecond(&_TraderjoeCToken.CallOpts)
}

// SupplyRatePerSecond is a free data retrieval call binding the contract method 0xb1d38974.
//
// Solidity: function supplyRatePerSecond() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) SupplyRatePerSecond() (*big.Int, error) {
	return _TraderjoeCToken.Contract.SupplyRatePerSecond(&_TraderjoeCToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TraderjoeCToken *TraderjoeCTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TraderjoeCToken *TraderjoeCTokenSession) Symbol() (string, error) {
	return _TraderjoeCToken.Contract.Symbol(&_TraderjoeCToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) Symbol() (string, error) {
	return _TraderjoeCToken.Contract.Symbol(&_TraderjoeCToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) TotalBorrows() (*big.Int, error) {
	return _TraderjoeCToken.Contract.TotalBorrows(&_TraderjoeCToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) TotalBorrows() (*big.Int, error) {
	return _TraderjoeCToken.Contract.TotalBorrows(&_TraderjoeCToken.CallOpts)
}

// TotalCollateralTokens is a free data retrieval call binding the contract method 0x19a4dd3c.
//
// Solidity: function totalCollateralTokens() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) TotalCollateralTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "totalCollateralTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollateralTokens is a free data retrieval call binding the contract method 0x19a4dd3c.
//
// Solidity: function totalCollateralTokens() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) TotalCollateralTokens() (*big.Int, error) {
	return _TraderjoeCToken.Contract.TotalCollateralTokens(&_TraderjoeCToken.CallOpts)
}

// TotalCollateralTokens is a free data retrieval call binding the contract method 0x19a4dd3c.
//
// Solidity: function totalCollateralTokens() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) TotalCollateralTokens() (*big.Int, error) {
	return _TraderjoeCToken.Contract.TotalCollateralTokens(&_TraderjoeCToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) TotalReserves() (*big.Int, error) {
	return _TraderjoeCToken.Contract.TotalReserves(&_TraderjoeCToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) TotalReserves() (*big.Int, error) {
	return _TraderjoeCToken.Contract.TotalReserves(&_TraderjoeCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) TotalSupply() (*big.Int, error) {
	return _TraderjoeCToken.Contract.TotalSupply(&_TraderjoeCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TraderjoeCToken.Contract.TotalSupply(&_TraderjoeCToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeCToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenSession) Underlying() (common.Address, error) {
	return _TraderjoeCToken.Contract.Underlying(&_TraderjoeCToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_TraderjoeCToken *TraderjoeCTokenCallerSession) Underlying() (common.Address, error) {
	return _TraderjoeCToken.Contract.Underlying(&_TraderjoeCToken.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) AcceptAdmin() (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.AcceptAdmin(&_TraderjoeCToken.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.AcceptAdmin(&_TraderjoeCToken.TransactOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) AddReserves(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "_addReserves", addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.AddReserves(&_TraderjoeCToken.TransactOpts, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.AddReserves(&_TraderjoeCToken.TransactOpts, addAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.ReduceReserves(&_TraderjoeCToken.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.ReduceReserves(&_TraderjoeCToken.TransactOpts, reduceAmount)
}

// SetCollateralCap is a paid mutator transaction binding the contract method 0x81cf00eb.
//
// Solidity: function _setCollateralCap(uint256 newCollateralCap) returns()
func (_TraderjoeCToken *TraderjoeCTokenTransactor) SetCollateralCap(opts *bind.TransactOpts, newCollateralCap *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "_setCollateralCap", newCollateralCap)
}

// SetCollateralCap is a paid mutator transaction binding the contract method 0x81cf00eb.
//
// Solidity: function _setCollateralCap(uint256 newCollateralCap) returns()
func (_TraderjoeCToken *TraderjoeCTokenSession) SetCollateralCap(newCollateralCap *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetCollateralCap(&_TraderjoeCToken.TransactOpts, newCollateralCap)
}

// SetCollateralCap is a paid mutator transaction binding the contract method 0x81cf00eb.
//
// Solidity: function _setCollateralCap(uint256 newCollateralCap) returns()
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) SetCollateralCap(newCollateralCap *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetCollateralCap(&_TraderjoeCToken.TransactOpts, newCollateralCap)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_TraderjoeCToken *TraderjoeCTokenTransactor) SetImplementation(opts *bind.TransactOpts, implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "_setImplementation", implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_TraderjoeCToken *TraderjoeCTokenSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetImplementation(&_TraderjoeCToken.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetImplementation(&_TraderjoeCToken.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetInterestRateModel(&_TraderjoeCToken.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetInterestRateModel(&_TraderjoeCToken.TransactOpts, newInterestRateModel)
}

// SetJoetroller is a paid mutator transaction binding the contract method 0x432f4b2d.
//
// Solidity: function _setJoetroller(address newJoetroller) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) SetJoetroller(opts *bind.TransactOpts, newJoetroller common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "_setJoetroller", newJoetroller)
}

// SetJoetroller is a paid mutator transaction binding the contract method 0x432f4b2d.
//
// Solidity: function _setJoetroller(address newJoetroller) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) SetJoetroller(newJoetroller common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetJoetroller(&_TraderjoeCToken.TransactOpts, newJoetroller)
}

// SetJoetroller is a paid mutator transaction binding the contract method 0x432f4b2d.
//
// Solidity: function _setJoetroller(address newJoetroller) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) SetJoetroller(newJoetroller common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetJoetroller(&_TraderjoeCToken.TransactOpts, newJoetroller)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetPendingAdmin(&_TraderjoeCToken.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetPendingAdmin(&_TraderjoeCToken.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetReserveFactor(&_TraderjoeCToken.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.SetReserveFactor(&_TraderjoeCToken.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) AccrueInterest() (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.AccrueInterest(&_TraderjoeCToken.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.AccrueInterest(&_TraderjoeCToken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Approve(&_TraderjoeCToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Approve(&_TraderjoeCToken.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.BalanceOfUnderlying(&_TraderjoeCToken.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.BalanceOfUnderlying(&_TraderjoeCToken.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Borrow(&_TraderjoeCToken.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Borrow(&_TraderjoeCToken.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.BorrowBalanceCurrent(&_TraderjoeCToken.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.BorrowBalanceCurrent(&_TraderjoeCToken.TransactOpts, account)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) DelegateToImplementation(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "delegateToImplementation", data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_TraderjoeCToken *TraderjoeCTokenSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.DelegateToImplementation(&_TraderjoeCToken.TransactOpts, data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.DelegateToImplementation(&_TraderjoeCToken.TransactOpts, data)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.ExchangeRateCurrent(&_TraderjoeCToken.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.ExchangeRateCurrent(&_TraderjoeCToken.TransactOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address initiator, uint256 amount, bytes data) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, initiator common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "flashLoan", receiver, initiator, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address initiator, uint256 amount, bytes data) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenSession) FlashLoan(receiver common.Address, initiator common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.FlashLoan(&_TraderjoeCToken.TransactOpts, receiver, initiator, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address initiator, uint256 amount, bytes data) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) FlashLoan(receiver common.Address, initiator common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.FlashLoan(&_TraderjoeCToken.TransactOpts, receiver, initiator, amount, data)
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_TraderjoeCToken *TraderjoeCTokenTransactor) Gulp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "gulp")
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_TraderjoeCToken *TraderjoeCTokenSession) Gulp() (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Gulp(&_TraderjoeCToken.TransactOpts)
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) Gulp() (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Gulp(&_TraderjoeCToken.TransactOpts)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address jTokenCollateral) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, jTokenCollateral common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, jTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address jTokenCollateral) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, jTokenCollateral common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.LiquidateBorrow(&_TraderjoeCToken.TransactOpts, borrower, repayAmount, jTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address jTokenCollateral) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, jTokenCollateral common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.LiquidateBorrow(&_TraderjoeCToken.TransactOpts, borrower, repayAmount, jTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Mint(&_TraderjoeCToken.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Mint(&_TraderjoeCToken.TransactOpts, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Redeem(&_TraderjoeCToken.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Redeem(&_TraderjoeCToken.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.RedeemUnderlying(&_TraderjoeCToken.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.RedeemUnderlying(&_TraderjoeCToken.TransactOpts, redeemAmount)
}

// RegisterCollateral is a paid mutator transaction binding the contract method 0x8897bd85.
//
// Solidity: function registerCollateral(address account) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) RegisterCollateral(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "registerCollateral", account)
}

// RegisterCollateral is a paid mutator transaction binding the contract method 0x8897bd85.
//
// Solidity: function registerCollateral(address account) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) RegisterCollateral(account common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.RegisterCollateral(&_TraderjoeCToken.TransactOpts, account)
}

// RegisterCollateral is a paid mutator transaction binding the contract method 0x8897bd85.
//
// Solidity: function registerCollateral(address account) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) RegisterCollateral(account common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.RegisterCollateral(&_TraderjoeCToken.TransactOpts, account)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.RepayBorrow(&_TraderjoeCToken.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.RepayBorrow(&_TraderjoeCToken.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.RepayBorrowBehalf(&_TraderjoeCToken.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.RepayBorrowBehalf(&_TraderjoeCToken.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Seize(&_TraderjoeCToken.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Seize(&_TraderjoeCToken.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.TotalBorrowsCurrent(&_TraderjoeCToken.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.TotalBorrowsCurrent(&_TraderjoeCToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Transfer(&_TraderjoeCToken.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.Transfer(&_TraderjoeCToken.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.TransferFrom(&_TraderjoeCToken.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.TransferFrom(&_TraderjoeCToken.TransactOpts, src, dst, amount)
}

// UnregisterCollateral is a paid mutator transaction binding the contract method 0x8b35776b.
//
// Solidity: function unregisterCollateral(address account) returns()
func (_TraderjoeCToken *TraderjoeCTokenTransactor) UnregisterCollateral(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.contract.Transact(opts, "unregisterCollateral", account)
}

// UnregisterCollateral is a paid mutator transaction binding the contract method 0x8b35776b.
//
// Solidity: function unregisterCollateral(address account) returns()
func (_TraderjoeCToken *TraderjoeCTokenSession) UnregisterCollateral(account common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.UnregisterCollateral(&_TraderjoeCToken.TransactOpts, account)
}

// UnregisterCollateral is a paid mutator transaction binding the contract method 0x8b35776b.
//
// Solidity: function unregisterCollateral(address account) returns()
func (_TraderjoeCToken *TraderjoeCTokenTransactorSession) UnregisterCollateral(account common.Address) (*types.Transaction, error) {
	return _TraderjoeCToken.Contract.UnregisterCollateral(&_TraderjoeCToken.TransactOpts, account)
}
