// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderjoeComptroller

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

// TraderjoeComptrollerMetaData contains all meta data concerning the TraderjoeComptroller contract.
var TraderjoeComptrollerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_acceptImplementation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingImplementation\",\"type\":\"address\"}],\"name\":\"_setPendingImplementation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rewardType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rewardType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"jTokens\",\"type\":\"address[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rewardType\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"jTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"jTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TraderjoeComptrollerABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderjoeComptrollerMetaData.ABI instead.
var TraderjoeComptrollerABI = TraderjoeComptrollerMetaData.ABI

// TraderjoeComptroller is an auto generated Go binding around an Ethereum contract.
type TraderjoeComptroller struct {
	TraderjoeComptrollerCaller     // Read-only binding to the contract
	TraderjoeComptrollerTransactor // Write-only binding to the contract
	TraderjoeComptrollerFilterer   // Log filterer for contract events
}

// TraderjoeComptrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderjoeComptrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeComptrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderjoeComptrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeComptrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderjoeComptrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeComptrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderjoeComptrollerSession struct {
	Contract     *TraderjoeComptroller // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TraderjoeComptrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderjoeComptrollerCallerSession struct {
	Contract *TraderjoeComptrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TraderjoeComptrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderjoeComptrollerTransactorSession struct {
	Contract     *TraderjoeComptrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TraderjoeComptrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderjoeComptrollerRaw struct {
	Contract *TraderjoeComptroller // Generic contract binding to access the raw methods on
}

// TraderjoeComptrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderjoeComptrollerCallerRaw struct {
	Contract *TraderjoeComptrollerCaller // Generic read-only contract binding to access the raw methods on
}

// TraderjoeComptrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderjoeComptrollerTransactorRaw struct {
	Contract *TraderjoeComptrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderjoeComptroller creates a new instance of TraderjoeComptroller, bound to a specific deployed contract.
func NewTraderjoeComptroller(address common.Address, backend bind.ContractBackend) (*TraderjoeComptroller, error) {
	contract, err := bindTraderjoeComptroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraderjoeComptroller{TraderjoeComptrollerCaller: TraderjoeComptrollerCaller{contract: contract}, TraderjoeComptrollerTransactor: TraderjoeComptrollerTransactor{contract: contract}, TraderjoeComptrollerFilterer: TraderjoeComptrollerFilterer{contract: contract}}, nil
}

// NewTraderjoeComptrollerCaller creates a new read-only instance of TraderjoeComptroller, bound to a specific deployed contract.
func NewTraderjoeComptrollerCaller(address common.Address, caller bind.ContractCaller) (*TraderjoeComptrollerCaller, error) {
	contract, err := bindTraderjoeComptroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeComptrollerCaller{contract: contract}, nil
}

// NewTraderjoeComptrollerTransactor creates a new write-only instance of TraderjoeComptroller, bound to a specific deployed contract.
func NewTraderjoeComptrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderjoeComptrollerTransactor, error) {
	contract, err := bindTraderjoeComptroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeComptrollerTransactor{contract: contract}, nil
}

// NewTraderjoeComptrollerFilterer creates a new log filterer instance of TraderjoeComptroller, bound to a specific deployed contract.
func NewTraderjoeComptrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderjoeComptrollerFilterer, error) {
	contract, err := bindTraderjoeComptroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderjoeComptrollerFilterer{contract: contract}, nil
}

// bindTraderjoeComptroller binds a generic wrapper to an already deployed contract.
func bindTraderjoeComptroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraderjoeComptrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeComptroller *TraderjoeComptrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeComptroller.Contract.TraderjoeComptrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeComptroller *TraderjoeComptrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.TraderjoeComptrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeComptroller *TraderjoeComptrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.TraderjoeComptrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeComptroller *TraderjoeComptrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeComptroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeComptroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) Admin() (common.Address, error) {
	return _TraderjoeComptroller.Contract.Admin(&_TraderjoeComptroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerCallerSession) Admin() (common.Address, error) {
	return _TraderjoeComptroller.Contract.Admin(&_TraderjoeComptroller.CallOpts)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address jToken) view returns(bool)
func (_TraderjoeComptroller *TraderjoeComptrollerCaller) CheckMembership(opts *bind.CallOpts, account common.Address, jToken common.Address) (bool, error) {
	var out []interface{}
	err := _TraderjoeComptroller.contract.Call(opts, &out, "checkMembership", account, jToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address jToken) view returns(bool)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) CheckMembership(account common.Address, jToken common.Address) (bool, error) {
	return _TraderjoeComptroller.Contract.CheckMembership(&_TraderjoeComptroller.CallOpts, account, jToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address jToken) view returns(bool)
func (_TraderjoeComptroller *TraderjoeComptrollerCallerSession) CheckMembership(account common.Address, jToken common.Address) (bool, error) {
	return _TraderjoeComptroller.Contract.CheckMembership(&_TraderjoeComptroller.CallOpts, account, jToken)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TraderjoeComptroller.contract.Call(opts, &out, "getAccountLiquidity", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _TraderjoeComptroller.Contract.GetAccountLiquidity(&_TraderjoeComptroller.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _TraderjoeComptroller.Contract.GetAccountLiquidity(&_TraderjoeComptroller.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_TraderjoeComptroller *TraderjoeComptrollerCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TraderjoeComptroller.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_TraderjoeComptroller *TraderjoeComptrollerSession) GetAllMarkets() ([]common.Address, error) {
	return _TraderjoeComptroller.Contract.GetAllMarkets(&_TraderjoeComptroller.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_TraderjoeComptroller *TraderjoeComptrollerCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _TraderjoeComptroller.Contract.GetAllMarkets(&_TraderjoeComptroller.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address user) view returns(address[])
func (_TraderjoeComptroller *TraderjoeComptrollerCaller) GetAssetsIn(opts *bind.CallOpts, user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _TraderjoeComptroller.contract.Call(opts, &out, "getAssetsIn", user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address user) view returns(address[])
func (_TraderjoeComptroller *TraderjoeComptrollerSession) GetAssetsIn(user common.Address) ([]common.Address, error) {
	return _TraderjoeComptroller.Contract.GetAssetsIn(&_TraderjoeComptroller.CallOpts, user)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address user) view returns(address[])
func (_TraderjoeComptroller *TraderjoeComptrollerCallerSession) GetAssetsIn(user common.Address) ([]common.Address, error) {
	return _TraderjoeComptroller.Contract.GetAssetsIn(&_TraderjoeComptroller.CallOpts, user)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeComptroller.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) Implementation() (common.Address, error) {
	return _TraderjoeComptroller.Contract.Implementation(&_TraderjoeComptroller.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerCallerSession) Implementation() (common.Address, error) {
	return _TraderjoeComptroller.Contract.Implementation(&_TraderjoeComptroller.CallOpts)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address jTokenBorrowed, address jTokenCollateral, uint256 repayAmount) view returns(uint256, uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, jTokenBorrowed common.Address, jTokenCollateral common.Address, repayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TraderjoeComptroller.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", jTokenBorrowed, jTokenCollateral, repayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address jTokenBorrowed, address jTokenCollateral, uint256 repayAmount) view returns(uint256, uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) LiquidateCalculateSeizeTokens(jTokenBorrowed common.Address, jTokenCollateral common.Address, repayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _TraderjoeComptroller.Contract.LiquidateCalculateSeizeTokens(&_TraderjoeComptroller.CallOpts, jTokenBorrowed, jTokenCollateral, repayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address jTokenBorrowed, address jTokenCollateral, uint256 repayAmount) view returns(uint256, uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerCallerSession) LiquidateCalculateSeizeTokens(jTokenBorrowed common.Address, jTokenCollateral common.Address, repayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _TraderjoeComptroller.Contract.LiquidateCalculateSeizeTokens(&_TraderjoeComptroller.CallOpts, jTokenBorrowed, jTokenCollateral, repayAmount)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address jToken) view returns(bool, uint256, uint8)
func (_TraderjoeComptroller *TraderjoeComptrollerCaller) Markets(opts *bind.CallOpts, jToken common.Address) (bool, *big.Int, uint8, error) {
	var out []interface{}
	err := _TraderjoeComptroller.contract.Call(opts, &out, "markets", jToken)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return out0, out1, out2, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address jToken) view returns(bool, uint256, uint8)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) Markets(jToken common.Address) (bool, *big.Int, uint8, error) {
	return _TraderjoeComptroller.Contract.Markets(&_TraderjoeComptroller.CallOpts, jToken)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address jToken) view returns(bool, uint256, uint8)
func (_TraderjoeComptroller *TraderjoeComptrollerCallerSession) Markets(jToken common.Address) (bool, *big.Int, uint8, error) {
	return _TraderjoeComptroller.Contract.Markets(&_TraderjoeComptroller.CallOpts, jToken)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeComptroller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) Oracle() (common.Address, error) {
	return _TraderjoeComptroller.Contract.Oracle(&_TraderjoeComptroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerCallerSession) Oracle() (common.Address, error) {
	return _TraderjoeComptroller.Contract.Oracle(&_TraderjoeComptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeComptroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) PendingAdmin() (common.Address, error) {
	return _TraderjoeComptroller.Contract.PendingAdmin(&_TraderjoeComptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _TraderjoeComptroller.Contract.PendingAdmin(&_TraderjoeComptroller.CallOpts)
}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerCaller) PendingImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeComptroller.contract.Call(opts, &out, "pendingImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) PendingImplementation() (common.Address, error) {
	return _TraderjoeComptroller.Contract.PendingImplementation(&_TraderjoeComptroller.CallOpts)
}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_TraderjoeComptroller *TraderjoeComptrollerCallerSession) PendingImplementation() (common.Address, error) {
	return _TraderjoeComptroller.Contract.PendingImplementation(&_TraderjoeComptroller.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) AcceptAdmin() (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.AcceptAdmin(&_TraderjoeComptroller.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.AcceptAdmin(&_TraderjoeComptroller.TransactOpts)
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) AcceptImplementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "_acceptImplementation")
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) AcceptImplementation() (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.AcceptImplementation(&_TraderjoeComptroller.TransactOpts)
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) AcceptImplementation() (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.AcceptImplementation(&_TraderjoeComptroller.TransactOpts)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.SetPendingAdmin(&_TraderjoeComptroller.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.SetPendingAdmin(&_TraderjoeComptroller.TransactOpts, newPendingAdmin)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) SetPendingImplementation(opts *bind.TransactOpts, newPendingImplementation common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "_setPendingImplementation", newPendingImplementation)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) SetPendingImplementation(newPendingImplementation common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.SetPendingImplementation(&_TraderjoeComptroller.TransactOpts, newPendingImplementation)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) SetPendingImplementation(newPendingImplementation common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.SetPendingImplementation(&_TraderjoeComptroller.TransactOpts, newPendingImplementation)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address jToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) BorrowAllowed(opts *bind.TransactOpts, jToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "borrowAllowed", jToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address jToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) BorrowAllowed(jToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.BorrowAllowed(&_TraderjoeComptroller.TransactOpts, jToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address jToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) BorrowAllowed(jToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.BorrowAllowed(&_TraderjoeComptroller.TransactOpts, jToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address jToken, address borrower, uint256 borrowAmount) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) BorrowVerify(opts *bind.TransactOpts, jToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "borrowVerify", jToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address jToken, address borrower, uint256 borrowAmount) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerSession) BorrowVerify(jToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.BorrowVerify(&_TraderjoeComptroller.TransactOpts, jToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address jToken, address borrower, uint256 borrowAmount) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) BorrowVerify(jToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.BorrowVerify(&_TraderjoeComptroller.TransactOpts, jToken, borrower, borrowAmount)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0952c563.
//
// Solidity: function claimReward(uint8 rewardType, address holder) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) ClaimReward(opts *bind.TransactOpts, rewardType uint8, holder common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "claimReward", rewardType, holder)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0952c563.
//
// Solidity: function claimReward(uint8 rewardType, address holder) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerSession) ClaimReward(rewardType uint8, holder common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.ClaimReward(&_TraderjoeComptroller.TransactOpts, rewardType, holder)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0952c563.
//
// Solidity: function claimReward(uint8 rewardType, address holder) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) ClaimReward(rewardType uint8, holder common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.ClaimReward(&_TraderjoeComptroller.TransactOpts, rewardType, holder)
}

// ClaimReward0 is a paid mutator transaction binding the contract method 0x744532ae.
//
// Solidity: function claimReward(uint8 rewardType, address holder, address[] jTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) ClaimReward0(opts *bind.TransactOpts, rewardType uint8, holder common.Address, jTokens []common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "claimReward0", rewardType, holder, jTokens)
}

// ClaimReward0 is a paid mutator transaction binding the contract method 0x744532ae.
//
// Solidity: function claimReward(uint8 rewardType, address holder, address[] jTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerSession) ClaimReward0(rewardType uint8, holder common.Address, jTokens []common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.ClaimReward0(&_TraderjoeComptroller.TransactOpts, rewardType, holder, jTokens)
}

// ClaimReward0 is a paid mutator transaction binding the contract method 0x744532ae.
//
// Solidity: function claimReward(uint8 rewardType, address holder, address[] jTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) ClaimReward0(rewardType uint8, holder common.Address, jTokens []common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.ClaimReward0(&_TraderjoeComptroller.TransactOpts, rewardType, holder, jTokens)
}

// ClaimReward1 is a paid mutator transaction binding the contract method 0x8805714b.
//
// Solidity: function claimReward(uint8 rewardType, address[] holders, address[] jTokens, bool borrowers, bool suppliers) payable returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) ClaimReward1(opts *bind.TransactOpts, rewardType uint8, holders []common.Address, jTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "claimReward1", rewardType, holders, jTokens, borrowers, suppliers)
}

// ClaimReward1 is a paid mutator transaction binding the contract method 0x8805714b.
//
// Solidity: function claimReward(uint8 rewardType, address[] holders, address[] jTokens, bool borrowers, bool suppliers) payable returns()
func (_TraderjoeComptroller *TraderjoeComptrollerSession) ClaimReward1(rewardType uint8, holders []common.Address, jTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.ClaimReward1(&_TraderjoeComptroller.TransactOpts, rewardType, holders, jTokens, borrowers, suppliers)
}

// ClaimReward1 is a paid mutator transaction binding the contract method 0x8805714b.
//
// Solidity: function claimReward(uint8 rewardType, address[] holders, address[] jTokens, bool borrowers, bool suppliers) payable returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) ClaimReward1(rewardType uint8, holders []common.Address, jTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.ClaimReward1(&_TraderjoeComptroller.TransactOpts, rewardType, holders, jTokens, borrowers, suppliers)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] jTokens) returns(uint256[])
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) EnterMarkets(opts *bind.TransactOpts, jTokens []common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "enterMarkets", jTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] jTokens) returns(uint256[])
func (_TraderjoeComptroller *TraderjoeComptrollerSession) EnterMarkets(jTokens []common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.EnterMarkets(&_TraderjoeComptroller.TransactOpts, jTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] jTokens) returns(uint256[])
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) EnterMarkets(jTokens []common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.EnterMarkets(&_TraderjoeComptroller.TransactOpts, jTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address jToken) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) ExitMarket(opts *bind.TransactOpts, jToken common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "exitMarket", jToken)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address jToken) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) ExitMarket(jToken common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.ExitMarket(&_TraderjoeComptroller.TransactOpts, jToken)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address jToken) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) ExitMarket(jToken common.Address) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.ExitMarket(&_TraderjoeComptroller.TransactOpts, jToken)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address jTokenBorrowed, address jTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, jTokenBorrowed common.Address, jTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "liquidateBorrowAllowed", jTokenBorrowed, jTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address jTokenBorrowed, address jTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) LiquidateBorrowAllowed(jTokenBorrowed common.Address, jTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.LiquidateBorrowAllowed(&_TraderjoeComptroller.TransactOpts, jTokenBorrowed, jTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address jTokenBorrowed, address jTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) LiquidateBorrowAllowed(jTokenBorrowed common.Address, jTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.LiquidateBorrowAllowed(&_TraderjoeComptroller.TransactOpts, jTokenBorrowed, jTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address jTokenBorrowed, address jTokenCollateral, address liquidator, address borrower, uint256 repayAmount, uint256 seizeTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, jTokenBorrowed common.Address, jTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "liquidateBorrowVerify", jTokenBorrowed, jTokenCollateral, liquidator, borrower, repayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address jTokenBorrowed, address jTokenCollateral, address liquidator, address borrower, uint256 repayAmount, uint256 seizeTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerSession) LiquidateBorrowVerify(jTokenBorrowed common.Address, jTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.LiquidateBorrowVerify(&_TraderjoeComptroller.TransactOpts, jTokenBorrowed, jTokenCollateral, liquidator, borrower, repayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address jTokenBorrowed, address jTokenCollateral, address liquidator, address borrower, uint256 repayAmount, uint256 seizeTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) LiquidateBorrowVerify(jTokenBorrowed common.Address, jTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.LiquidateBorrowVerify(&_TraderjoeComptroller.TransactOpts, jTokenBorrowed, jTokenCollateral, liquidator, borrower, repayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address jToken, address minter, uint256 mintAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) MintAllowed(opts *bind.TransactOpts, jToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "mintAllowed", jToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address jToken, address minter, uint256 mintAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) MintAllowed(jToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.MintAllowed(&_TraderjoeComptroller.TransactOpts, jToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address jToken, address minter, uint256 mintAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) MintAllowed(jToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.MintAllowed(&_TraderjoeComptroller.TransactOpts, jToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address jToken, address minter, uint256 mintAmount, uint256 mintTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) MintVerify(opts *bind.TransactOpts, jToken common.Address, minter common.Address, mintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "mintVerify", jToken, minter, mintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address jToken, address minter, uint256 mintAmount, uint256 mintTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerSession) MintVerify(jToken common.Address, minter common.Address, mintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.MintVerify(&_TraderjoeComptroller.TransactOpts, jToken, minter, mintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address jToken, address minter, uint256 mintAmount, uint256 mintTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) MintVerify(jToken common.Address, minter common.Address, mintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.MintVerify(&_TraderjoeComptroller.TransactOpts, jToken, minter, mintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address jToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) RedeemAllowed(opts *bind.TransactOpts, jToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "redeemAllowed", jToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address jToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) RedeemAllowed(jToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.RedeemAllowed(&_TraderjoeComptroller.TransactOpts, jToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address jToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) RedeemAllowed(jToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.RedeemAllowed(&_TraderjoeComptroller.TransactOpts, jToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address jToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) RedeemVerify(opts *bind.TransactOpts, jToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "redeemVerify", jToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address jToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerSession) RedeemVerify(jToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.RedeemVerify(&_TraderjoeComptroller.TransactOpts, jToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address jToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) RedeemVerify(jToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.RedeemVerify(&_TraderjoeComptroller.TransactOpts, jToken, redeemer, redeemAmount, redeemTokens)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address jToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, jToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "repayBorrowAllowed", jToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address jToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) RepayBorrowAllowed(jToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.RepayBorrowAllowed(&_TraderjoeComptroller.TransactOpts, jToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address jToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) RepayBorrowAllowed(jToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.RepayBorrowAllowed(&_TraderjoeComptroller.TransactOpts, jToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address jToken, address payer, address borrower, uint256 repayAmount, uint256 borrowerIndex) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) RepayBorrowVerify(opts *bind.TransactOpts, jToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "repayBorrowVerify", jToken, payer, borrower, repayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address jToken, address payer, address borrower, uint256 repayAmount, uint256 borrowerIndex) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerSession) RepayBorrowVerify(jToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.RepayBorrowVerify(&_TraderjoeComptroller.TransactOpts, jToken, payer, borrower, repayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address jToken, address payer, address borrower, uint256 repayAmount, uint256 borrowerIndex) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) RepayBorrowVerify(jToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.RepayBorrowVerify(&_TraderjoeComptroller.TransactOpts, jToken, payer, borrower, repayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address jTokenCollateral, address jTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) SeizeAllowed(opts *bind.TransactOpts, jTokenCollateral common.Address, jTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "seizeAllowed", jTokenCollateral, jTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address jTokenCollateral, address jTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) SeizeAllowed(jTokenCollateral common.Address, jTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.SeizeAllowed(&_TraderjoeComptroller.TransactOpts, jTokenCollateral, jTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address jTokenCollateral, address jTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) SeizeAllowed(jTokenCollateral common.Address, jTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.SeizeAllowed(&_TraderjoeComptroller.TransactOpts, jTokenCollateral, jTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address jTokenCollateral, address jTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) SeizeVerify(opts *bind.TransactOpts, jTokenCollateral common.Address, jTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "seizeVerify", jTokenCollateral, jTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address jTokenCollateral, address jTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerSession) SeizeVerify(jTokenCollateral common.Address, jTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.SeizeVerify(&_TraderjoeComptroller.TransactOpts, jTokenCollateral, jTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address jTokenCollateral, address jTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) SeizeVerify(jTokenCollateral common.Address, jTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.SeizeVerify(&_TraderjoeComptroller.TransactOpts, jTokenCollateral, jTokenBorrowed, liquidator, borrower, seizeTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address jToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) TransferAllowed(opts *bind.TransactOpts, jToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "transferAllowed", jToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address jToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerSession) TransferAllowed(jToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.TransferAllowed(&_TraderjoeComptroller.TransactOpts, jToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address jToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) TransferAllowed(jToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.TransferAllowed(&_TraderjoeComptroller.TransactOpts, jToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address jToken, address src, address dst, uint256 transferTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactor) TransferVerify(opts *bind.TransactOpts, jToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.contract.Transact(opts, "transferVerify", jToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address jToken, address src, address dst, uint256 transferTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerSession) TransferVerify(jToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.TransferVerify(&_TraderjoeComptroller.TransactOpts, jToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address jToken, address src, address dst, uint256 transferTokens) returns()
func (_TraderjoeComptroller *TraderjoeComptrollerTransactorSession) TransferVerify(jToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _TraderjoeComptroller.Contract.TransferVerify(&_TraderjoeComptroller.TransactOpts, jToken, src, dst, transferTokens)
}
