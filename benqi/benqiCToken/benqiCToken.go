// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package benqiCToken

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

// BenqiCTokenMetaData contains all meta data concerning the BenqiCToken contract.
var BenqiCTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowResign\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"name\":\"_setImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProtocolSeizeShareMantissa\",\"type\":\"uint256\"}],\"name\":\"_setProtocolSeizeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrualBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowRatePerTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToViewImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isQiToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"qiTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolSeizeShareMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BenqiCTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BenqiCTokenMetaData.ABI instead.
var BenqiCTokenABI = BenqiCTokenMetaData.ABI

// BenqiCToken is an auto generated Go binding around an Ethereum contract.
type BenqiCToken struct {
	BenqiCTokenCaller     // Read-only binding to the contract
	BenqiCTokenTransactor // Write-only binding to the contract
	BenqiCTokenFilterer   // Log filterer for contract events
}

// BenqiCTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BenqiCTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqiCTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BenqiCTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqiCTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BenqiCTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqiCTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BenqiCTokenSession struct {
	Contract     *BenqiCToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BenqiCTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BenqiCTokenCallerSession struct {
	Contract *BenqiCTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BenqiCTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BenqiCTokenTransactorSession struct {
	Contract     *BenqiCTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BenqiCTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BenqiCTokenRaw struct {
	Contract *BenqiCToken // Generic contract binding to access the raw methods on
}

// BenqiCTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BenqiCTokenCallerRaw struct {
	Contract *BenqiCTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BenqiCTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BenqiCTokenTransactorRaw struct {
	Contract *BenqiCTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBenqiCToken creates a new instance of BenqiCToken, bound to a specific deployed contract.
func NewBenqiCToken(address common.Address, backend bind.ContractBackend) (*BenqiCToken, error) {
	contract, err := bindBenqiCToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BenqiCToken{BenqiCTokenCaller: BenqiCTokenCaller{contract: contract}, BenqiCTokenTransactor: BenqiCTokenTransactor{contract: contract}, BenqiCTokenFilterer: BenqiCTokenFilterer{contract: contract}}, nil
}

// NewBenqiCTokenCaller creates a new read-only instance of BenqiCToken, bound to a specific deployed contract.
func NewBenqiCTokenCaller(address common.Address, caller bind.ContractCaller) (*BenqiCTokenCaller, error) {
	contract, err := bindBenqiCToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BenqiCTokenCaller{contract: contract}, nil
}

// NewBenqiCTokenTransactor creates a new write-only instance of BenqiCToken, bound to a specific deployed contract.
func NewBenqiCTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BenqiCTokenTransactor, error) {
	contract, err := bindBenqiCToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BenqiCTokenTransactor{contract: contract}, nil
}

// NewBenqiCTokenFilterer creates a new log filterer instance of BenqiCToken, bound to a specific deployed contract.
func NewBenqiCTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BenqiCTokenFilterer, error) {
	contract, err := bindBenqiCToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BenqiCTokenFilterer{contract: contract}, nil
}

// bindBenqiCToken binds a generic wrapper to an already deployed contract.
func bindBenqiCToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BenqiCTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BenqiCToken *BenqiCTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BenqiCToken.Contract.BenqiCTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BenqiCToken *BenqiCTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenqiCToken.Contract.BenqiCTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BenqiCToken *BenqiCTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BenqiCToken.Contract.BenqiCTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BenqiCToken *BenqiCTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BenqiCToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BenqiCToken *BenqiCTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenqiCToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BenqiCToken *BenqiCTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BenqiCToken.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockTimestamp is a free data retrieval call binding the contract method 0xcfa99201.
//
// Solidity: function accrualBlockTimestamp() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) AccrualBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "accrualBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockTimestamp is a free data retrieval call binding the contract method 0xcfa99201.
//
// Solidity: function accrualBlockTimestamp() view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) AccrualBlockTimestamp() (*big.Int, error) {
	return _BenqiCToken.Contract.AccrualBlockTimestamp(&_BenqiCToken.CallOpts)
}

// AccrualBlockTimestamp is a free data retrieval call binding the contract method 0xcfa99201.
//
// Solidity: function accrualBlockTimestamp() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) AccrualBlockTimestamp() (*big.Int, error) {
	return _BenqiCToken.Contract.AccrualBlockTimestamp(&_BenqiCToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BenqiCToken *BenqiCTokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BenqiCToken *BenqiCTokenSession) Admin() (common.Address, error) {
	return _BenqiCToken.Contract.Admin(&_BenqiCToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BenqiCToken *BenqiCTokenCallerSession) Admin() (common.Address, error) {
	return _BenqiCToken.Contract.Admin(&_BenqiCToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BenqiCToken.Contract.Allowance(&_BenqiCToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BenqiCToken.Contract.Allowance(&_BenqiCToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BenqiCToken.Contract.BalanceOf(&_BenqiCToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BenqiCToken.Contract.BalanceOf(&_BenqiCToken.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _BenqiCToken.Contract.BorrowBalanceStored(&_BenqiCToken.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _BenqiCToken.Contract.BorrowBalanceStored(&_BenqiCToken.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) BorrowIndex() (*big.Int, error) {
	return _BenqiCToken.Contract.BorrowIndex(&_BenqiCToken.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) BorrowIndex() (*big.Int, error) {
	return _BenqiCToken.Contract.BorrowIndex(&_BenqiCToken.CallOpts)
}

// BorrowRatePerTimestamp is a free data retrieval call binding the contract method 0xcd91801c.
//
// Solidity: function borrowRatePerTimestamp() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) BorrowRatePerTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "borrowRatePerTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerTimestamp is a free data retrieval call binding the contract method 0xcd91801c.
//
// Solidity: function borrowRatePerTimestamp() view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) BorrowRatePerTimestamp() (*big.Int, error) {
	return _BenqiCToken.Contract.BorrowRatePerTimestamp(&_BenqiCToken.CallOpts)
}

// BorrowRatePerTimestamp is a free data retrieval call binding the contract method 0xcd91801c.
//
// Solidity: function borrowRatePerTimestamp() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) BorrowRatePerTimestamp() (*big.Int, error) {
	return _BenqiCToken.Contract.BorrowRatePerTimestamp(&_BenqiCToken.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_BenqiCToken *BenqiCTokenCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_BenqiCToken *BenqiCTokenSession) Comptroller() (common.Address, error) {
	return _BenqiCToken.Contract.Comptroller(&_BenqiCToken.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_BenqiCToken *BenqiCTokenCallerSession) Comptroller() (common.Address, error) {
	return _BenqiCToken.Contract.Comptroller(&_BenqiCToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BenqiCToken *BenqiCTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BenqiCToken *BenqiCTokenSession) Decimals() (uint8, error) {
	return _BenqiCToken.Contract.Decimals(&_BenqiCToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BenqiCToken *BenqiCTokenCallerSession) Decimals() (uint8, error) {
	return _BenqiCToken.Contract.Decimals(&_BenqiCToken.CallOpts)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_BenqiCToken *BenqiCTokenCaller) DelegateToViewImplementation(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "delegateToViewImplementation", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_BenqiCToken *BenqiCTokenSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _BenqiCToken.Contract.DelegateToViewImplementation(&_BenqiCToken.CallOpts, data)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_BenqiCToken *BenqiCTokenCallerSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _BenqiCToken.Contract.DelegateToViewImplementation(&_BenqiCToken.CallOpts, data)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) ExchangeRateStored() (*big.Int, error) {
	return _BenqiCToken.Contract.ExchangeRateStored(&_BenqiCToken.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _BenqiCToken.Contract.ExchangeRateStored(&_BenqiCToken.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_BenqiCToken *BenqiCTokenCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "getAccountSnapshot", account)

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
func (_BenqiCToken *BenqiCTokenSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _BenqiCToken.Contract.GetAccountSnapshot(&_BenqiCToken.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _BenqiCToken.Contract.GetAccountSnapshot(&_BenqiCToken.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) GetCash() (*big.Int, error) {
	return _BenqiCToken.Contract.GetCash(&_BenqiCToken.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) GetCash() (*big.Int, error) {
	return _BenqiCToken.Contract.GetCash(&_BenqiCToken.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BenqiCToken *BenqiCTokenCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BenqiCToken *BenqiCTokenSession) Implementation() (common.Address, error) {
	return _BenqiCToken.Contract.Implementation(&_BenqiCToken.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BenqiCToken *BenqiCTokenCallerSession) Implementation() (common.Address, error) {
	return _BenqiCToken.Contract.Implementation(&_BenqiCToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_BenqiCToken *BenqiCTokenCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_BenqiCToken *BenqiCTokenSession) InterestRateModel() (common.Address, error) {
	return _BenqiCToken.Contract.InterestRateModel(&_BenqiCToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_BenqiCToken *BenqiCTokenCallerSession) InterestRateModel() (common.Address, error) {
	return _BenqiCToken.Contract.InterestRateModel(&_BenqiCToken.CallOpts)
}

// IsQiToken is a free data retrieval call binding the contract method 0x840bbeac.
//
// Solidity: function isQiToken() view returns(bool)
func (_BenqiCToken *BenqiCTokenCaller) IsQiToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "isQiToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQiToken is a free data retrieval call binding the contract method 0x840bbeac.
//
// Solidity: function isQiToken() view returns(bool)
func (_BenqiCToken *BenqiCTokenSession) IsQiToken() (bool, error) {
	return _BenqiCToken.Contract.IsQiToken(&_BenqiCToken.CallOpts)
}

// IsQiToken is a free data retrieval call binding the contract method 0x840bbeac.
//
// Solidity: function isQiToken() view returns(bool)
func (_BenqiCToken *BenqiCTokenCallerSession) IsQiToken() (bool, error) {
	return _BenqiCToken.Contract.IsQiToken(&_BenqiCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BenqiCToken *BenqiCTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BenqiCToken *BenqiCTokenSession) Name() (string, error) {
	return _BenqiCToken.Contract.Name(&_BenqiCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BenqiCToken *BenqiCTokenCallerSession) Name() (string, error) {
	return _BenqiCToken.Contract.Name(&_BenqiCToken.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_BenqiCToken *BenqiCTokenCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_BenqiCToken *BenqiCTokenSession) PendingAdmin() (common.Address, error) {
	return _BenqiCToken.Contract.PendingAdmin(&_BenqiCToken.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_BenqiCToken *BenqiCTokenCallerSession) PendingAdmin() (common.Address, error) {
	return _BenqiCToken.Contract.PendingAdmin(&_BenqiCToken.CallOpts)
}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) ProtocolSeizeShareMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "protocolSeizeShareMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) ProtocolSeizeShareMantissa() (*big.Int, error) {
	return _BenqiCToken.Contract.ProtocolSeizeShareMantissa(&_BenqiCToken.CallOpts)
}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) ProtocolSeizeShareMantissa() (*big.Int, error) {
	return _BenqiCToken.Contract.ProtocolSeizeShareMantissa(&_BenqiCToken.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) ReserveFactorMantissa() (*big.Int, error) {
	return _BenqiCToken.Contract.ReserveFactorMantissa(&_BenqiCToken.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _BenqiCToken.Contract.ReserveFactorMantissa(&_BenqiCToken.CallOpts)
}

// SupplyRatePerTimestamp is a free data retrieval call binding the contract method 0xd3bd2c72.
//
// Solidity: function supplyRatePerTimestamp() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) SupplyRatePerTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "supplyRatePerTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerTimestamp is a free data retrieval call binding the contract method 0xd3bd2c72.
//
// Solidity: function supplyRatePerTimestamp() view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) SupplyRatePerTimestamp() (*big.Int, error) {
	return _BenqiCToken.Contract.SupplyRatePerTimestamp(&_BenqiCToken.CallOpts)
}

// SupplyRatePerTimestamp is a free data retrieval call binding the contract method 0xd3bd2c72.
//
// Solidity: function supplyRatePerTimestamp() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) SupplyRatePerTimestamp() (*big.Int, error) {
	return _BenqiCToken.Contract.SupplyRatePerTimestamp(&_BenqiCToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BenqiCToken *BenqiCTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BenqiCToken *BenqiCTokenSession) Symbol() (string, error) {
	return _BenqiCToken.Contract.Symbol(&_BenqiCToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BenqiCToken *BenqiCTokenCallerSession) Symbol() (string, error) {
	return _BenqiCToken.Contract.Symbol(&_BenqiCToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) TotalBorrows() (*big.Int, error) {
	return _BenqiCToken.Contract.TotalBorrows(&_BenqiCToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) TotalBorrows() (*big.Int, error) {
	return _BenqiCToken.Contract.TotalBorrows(&_BenqiCToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) TotalReserves() (*big.Int, error) {
	return _BenqiCToken.Contract.TotalReserves(&_BenqiCToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) TotalReserves() (*big.Int, error) {
	return _BenqiCToken.Contract.TotalReserves(&_BenqiCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) TotalSupply() (*big.Int, error) {
	return _BenqiCToken.Contract.TotalSupply(&_BenqiCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BenqiCToken *BenqiCTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BenqiCToken.Contract.TotalSupply(&_BenqiCToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_BenqiCToken *BenqiCTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiCToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_BenqiCToken *BenqiCTokenSession) Underlying() (common.Address, error) {
	return _BenqiCToken.Contract.Underlying(&_BenqiCToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_BenqiCToken *BenqiCTokenCallerSession) Underlying() (common.Address, error) {
	return _BenqiCToken.Contract.Underlying(&_BenqiCToken.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) AcceptAdmin() (*types.Transaction, error) {
	return _BenqiCToken.Contract.AcceptAdmin(&_BenqiCToken.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _BenqiCToken.Contract.AcceptAdmin(&_BenqiCToken.TransactOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) AddReserves(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "_addReserves", addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.AddReserves(&_BenqiCToken.TransactOpts, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.AddReserves(&_BenqiCToken.TransactOpts, addAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.ReduceReserves(&_BenqiCToken.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.ReduceReserves(&_BenqiCToken.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetComptroller(&_BenqiCToken.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetComptroller(&_BenqiCToken.TransactOpts, newComptroller)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_BenqiCToken *BenqiCTokenTransactor) SetImplementation(opts *bind.TransactOpts, implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "_setImplementation", implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_BenqiCToken *BenqiCTokenSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetImplementation(&_BenqiCToken.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_BenqiCToken *BenqiCTokenTransactorSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetImplementation(&_BenqiCToken.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetInterestRateModel(&_BenqiCToken.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetInterestRateModel(&_BenqiCToken.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetPendingAdmin(&_BenqiCToken.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetPendingAdmin(&_BenqiCToken.TransactOpts, newPendingAdmin)
}

// SetProtocolSeizeShare is a paid mutator transaction binding the contract method 0x83030846.
//
// Solidity: function _setProtocolSeizeShare(uint256 newProtocolSeizeShareMantissa) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) SetProtocolSeizeShare(opts *bind.TransactOpts, newProtocolSeizeShareMantissa *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "_setProtocolSeizeShare", newProtocolSeizeShareMantissa)
}

// SetProtocolSeizeShare is a paid mutator transaction binding the contract method 0x83030846.
//
// Solidity: function _setProtocolSeizeShare(uint256 newProtocolSeizeShareMantissa) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) SetProtocolSeizeShare(newProtocolSeizeShareMantissa *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetProtocolSeizeShare(&_BenqiCToken.TransactOpts, newProtocolSeizeShareMantissa)
}

// SetProtocolSeizeShare is a paid mutator transaction binding the contract method 0x83030846.
//
// Solidity: function _setProtocolSeizeShare(uint256 newProtocolSeizeShareMantissa) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) SetProtocolSeizeShare(newProtocolSeizeShareMantissa *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetProtocolSeizeShare(&_BenqiCToken.TransactOpts, newProtocolSeizeShareMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetReserveFactor(&_BenqiCToken.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SetReserveFactor(&_BenqiCToken.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) AccrueInterest() (*types.Transaction, error) {
	return _BenqiCToken.Contract.AccrueInterest(&_BenqiCToken.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _BenqiCToken.Contract.AccrueInterest(&_BenqiCToken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BenqiCToken *BenqiCTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BenqiCToken *BenqiCTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Approve(&_BenqiCToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BenqiCToken *BenqiCTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Approve(&_BenqiCToken.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.BalanceOfUnderlying(&_BenqiCToken.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.BalanceOfUnderlying(&_BenqiCToken.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Borrow(&_BenqiCToken.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Borrow(&_BenqiCToken.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.BorrowBalanceCurrent(&_BenqiCToken.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.BorrowBalanceCurrent(&_BenqiCToken.TransactOpts, account)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_BenqiCToken *BenqiCTokenTransactor) DelegateToImplementation(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "delegateToImplementation", data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_BenqiCToken *BenqiCTokenSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _BenqiCToken.Contract.DelegateToImplementation(&_BenqiCToken.TransactOpts, data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_BenqiCToken *BenqiCTokenTransactorSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _BenqiCToken.Contract.DelegateToImplementation(&_BenqiCToken.TransactOpts, data)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _BenqiCToken.Contract.ExchangeRateCurrent(&_BenqiCToken.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _BenqiCToken.Contract.ExchangeRateCurrent(&_BenqiCToken.TransactOpts)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address qiTokenCollateral) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, qiTokenCollateral common.Address) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, qiTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address qiTokenCollateral) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, qiTokenCollateral common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.LiquidateBorrow(&_BenqiCToken.TransactOpts, borrower, repayAmount, qiTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address qiTokenCollateral) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, qiTokenCollateral common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.LiquidateBorrow(&_BenqiCToken.TransactOpts, borrower, repayAmount, qiTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Mint(&_BenqiCToken.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Mint(&_BenqiCToken.TransactOpts, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Redeem(&_BenqiCToken.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Redeem(&_BenqiCToken.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.RedeemUnderlying(&_BenqiCToken.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.RedeemUnderlying(&_BenqiCToken.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.RepayBorrow(&_BenqiCToken.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.RepayBorrow(&_BenqiCToken.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.RepayBorrowBehalf(&_BenqiCToken.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.RepayBorrowBehalf(&_BenqiCToken.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Seize(&_BenqiCToken.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Seize(&_BenqiCToken.TransactOpts, liquidator, borrower, seizeTokens)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address token) returns()
func (_BenqiCToken *BenqiCTokenTransactor) SweepToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "sweepToken", token)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address token) returns()
func (_BenqiCToken *BenqiCTokenSession) SweepToken(token common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SweepToken(&_BenqiCToken.TransactOpts, token)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address token) returns()
func (_BenqiCToken *BenqiCTokenTransactorSession) SweepToken(token common.Address) (*types.Transaction, error) {
	return _BenqiCToken.Contract.SweepToken(&_BenqiCToken.TransactOpts, token)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_BenqiCToken *BenqiCTokenSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _BenqiCToken.Contract.TotalBorrowsCurrent(&_BenqiCToken.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_BenqiCToken *BenqiCTokenTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _BenqiCToken.Contract.TotalBorrowsCurrent(&_BenqiCToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_BenqiCToken *BenqiCTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_BenqiCToken *BenqiCTokenSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Transfer(&_BenqiCToken.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_BenqiCToken *BenqiCTokenTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.Transfer(&_BenqiCToken.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_BenqiCToken *BenqiCTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_BenqiCToken *BenqiCTokenSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.TransferFrom(&_BenqiCToken.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_BenqiCToken *BenqiCTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenqiCToken.Contract.TransferFrom(&_BenqiCToken.TransactOpts, src, dst, amount)
}
