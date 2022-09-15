// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package instadappVault

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

// InstadappVaultMetaData contains all meta data concerning the InstadappVault contract.
var InstadappVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"auth_\",\"type\":\"address\"}],\"name\":\"addDSAAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"status_\",\"type\":\"uint256\"}],\"name\":\"changeStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"collectRevenue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt_\",\"type\":\"uint256\"}],\"name\":\"deleverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deleverageAmt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"deleverageAndWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deleverageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentExchangePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"exchangePrice_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newRevenue_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"flashTkn_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"flashAmt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"route_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stEthAmt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wethAmt_\",\"type\":\"uint256\"}],\"name\":\"importPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRebalancer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRevenueExchangePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt_\",\"type\":\"uint256\"}],\"name\":\"leverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ratios\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"maxLimit\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minLimit\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minLimitGap\",\"type\":\"uint16\"},{\"internalType\":\"uint128\",\"name\":\"maxBorrowRate\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"flashTkn_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"flashAmt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"route_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"vaults_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amts_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"excessDebt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paybackDebt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountToSwap_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extraWithdraw_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unitAmt_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"oneInchData_\",\"type\":\"bytes\"}],\"name\":\"rebalanceOne\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawAmt_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"flashTkn_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"flashAmt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"route_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountToSwap_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unitAmt_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"oneInchData_\",\"type\":\"bytes\"}],\"name\":\"rebalanceTwo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revenueFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldata_\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operation_\",\"type\":\"uint256\"}],\"name\":\"spell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"supply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vtokenAmount_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"supplyEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vtokenAmount_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"auth_\",\"type\":\"address\"}],\"name\":\"updateAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"revenueFee_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalFee_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deleverageFee_\",\"type\":\"uint256\"}],\"name\":\"updateFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16[]\",\"name\":\"ratios_\",\"type\":\"uint16[]\"}],\"name\":\"updateRatios\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rebalancer_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRebalancer_\",\"type\":\"bool\"}],\"name\":\"updateRebalancer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultDsa\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vtokenAmount_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"afterDeleverage_\",\"type\":\"bool\"}],\"name\":\"withdrawFinal\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"transferAmts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// InstadappVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use InstadappVaultMetaData.ABI instead.
var InstadappVaultABI = InstadappVaultMetaData.ABI

// InstadappVault is an auto generated Go binding around an Ethereum contract.
type InstadappVault struct {
	InstadappVaultCaller     // Read-only binding to the contract
	InstadappVaultTransactor // Write-only binding to the contract
	InstadappVaultFilterer   // Log filterer for contract events
}

// InstadappVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type InstadappVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstadappVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InstadappVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstadappVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InstadappVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstadappVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InstadappVaultSession struct {
	Contract     *InstadappVault   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InstadappVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InstadappVaultCallerSession struct {
	Contract *InstadappVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// InstadappVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InstadappVaultTransactorSession struct {
	Contract     *InstadappVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// InstadappVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type InstadappVaultRaw struct {
	Contract *InstadappVault // Generic contract binding to access the raw methods on
}

// InstadappVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InstadappVaultCallerRaw struct {
	Contract *InstadappVaultCaller // Generic read-only contract binding to access the raw methods on
}

// InstadappVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InstadappVaultTransactorRaw struct {
	Contract *InstadappVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInstadappVault creates a new instance of InstadappVault, bound to a specific deployed contract.
func NewInstadappVault(address common.Address, backend bind.ContractBackend) (*InstadappVault, error) {
	contract, err := bindInstadappVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InstadappVault{InstadappVaultCaller: InstadappVaultCaller{contract: contract}, InstadappVaultTransactor: InstadappVaultTransactor{contract: contract}, InstadappVaultFilterer: InstadappVaultFilterer{contract: contract}}, nil
}

// NewInstadappVaultCaller creates a new read-only instance of InstadappVault, bound to a specific deployed contract.
func NewInstadappVaultCaller(address common.Address, caller bind.ContractCaller) (*InstadappVaultCaller, error) {
	contract, err := bindInstadappVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InstadappVaultCaller{contract: contract}, nil
}

// NewInstadappVaultTransactor creates a new write-only instance of InstadappVault, bound to a specific deployed contract.
func NewInstadappVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*InstadappVaultTransactor, error) {
	contract, err := bindInstadappVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InstadappVaultTransactor{contract: contract}, nil
}

// NewInstadappVaultFilterer creates a new log filterer instance of InstadappVault, bound to a specific deployed contract.
func NewInstadappVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*InstadappVaultFilterer, error) {
	contract, err := bindInstadappVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InstadappVaultFilterer{contract: contract}, nil
}

// bindInstadappVault binds a generic wrapper to an already deployed contract.
func bindInstadappVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InstadappVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InstadappVault *InstadappVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InstadappVault.Contract.InstadappVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InstadappVault *InstadappVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InstadappVault.Contract.InstadappVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InstadappVault *InstadappVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InstadappVault.Contract.InstadappVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InstadappVault *InstadappVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InstadappVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InstadappVault *InstadappVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InstadappVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InstadappVault *InstadappVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InstadappVault.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_InstadappVault *InstadappVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_InstadappVault *InstadappVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _InstadappVault.Contract.Allowance(&_InstadappVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_InstadappVault *InstadappVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _InstadappVault.Contract.Allowance(&_InstadappVault.CallOpts, owner, spender)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_InstadappVault *InstadappVaultCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_InstadappVault *InstadappVaultSession) Auth() (common.Address, error) {
	return _InstadappVault.Contract.Auth(&_InstadappVault.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_InstadappVault *InstadappVaultCallerSession) Auth() (common.Address, error) {
	return _InstadappVault.Contract.Auth(&_InstadappVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InstadappVault *InstadappVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InstadappVault *InstadappVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InstadappVault.Contract.BalanceOf(&_InstadappVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InstadappVault *InstadappVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InstadappVault.Contract.BalanceOf(&_InstadappVault.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InstadappVault *InstadappVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InstadappVault *InstadappVaultSession) Decimals() (uint8, error) {
	return _InstadappVault.Contract.Decimals(&_InstadappVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InstadappVault *InstadappVaultCallerSession) Decimals() (uint8, error) {
	return _InstadappVault.Contract.Decimals(&_InstadappVault.CallOpts)
}

// DeleverageFee is a free data retrieval call binding the contract method 0xa4c6c1e0.
//
// Solidity: function deleverageFee() view returns(uint256)
func (_InstadappVault *InstadappVaultCaller) DeleverageFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "deleverageFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeleverageFee is a free data retrieval call binding the contract method 0xa4c6c1e0.
//
// Solidity: function deleverageFee() view returns(uint256)
func (_InstadappVault *InstadappVaultSession) DeleverageFee() (*big.Int, error) {
	return _InstadappVault.Contract.DeleverageFee(&_InstadappVault.CallOpts)
}

// DeleverageFee is a free data retrieval call binding the contract method 0xa4c6c1e0.
//
// Solidity: function deleverageFee() view returns(uint256)
func (_InstadappVault *InstadappVaultCallerSession) DeleverageFee() (*big.Int, error) {
	return _InstadappVault.Contract.DeleverageFee(&_InstadappVault.CallOpts)
}

// GetCurrentExchangePrice is a free data retrieval call binding the contract method 0xcc4a0158.
//
// Solidity: function getCurrentExchangePrice() view returns(uint256 exchangePrice_, uint256 newRevenue_)
func (_InstadappVault *InstadappVaultCaller) GetCurrentExchangePrice(opts *bind.CallOpts) (struct {
	ExchangePrice *big.Int
	NewRevenue    *big.Int
}, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "getCurrentExchangePrice")

	outstruct := new(struct {
		ExchangePrice *big.Int
		NewRevenue    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExchangePrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NewRevenue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentExchangePrice is a free data retrieval call binding the contract method 0xcc4a0158.
//
// Solidity: function getCurrentExchangePrice() view returns(uint256 exchangePrice_, uint256 newRevenue_)
func (_InstadappVault *InstadappVaultSession) GetCurrentExchangePrice() (struct {
	ExchangePrice *big.Int
	NewRevenue    *big.Int
}, error) {
	return _InstadappVault.Contract.GetCurrentExchangePrice(&_InstadappVault.CallOpts)
}

// GetCurrentExchangePrice is a free data retrieval call binding the contract method 0xcc4a0158.
//
// Solidity: function getCurrentExchangePrice() view returns(uint256 exchangePrice_, uint256 newRevenue_)
func (_InstadappVault *InstadappVaultCallerSession) GetCurrentExchangePrice() (struct {
	ExchangePrice *big.Int
	NewRevenue    *big.Int
}, error) {
	return _InstadappVault.Contract.GetCurrentExchangePrice(&_InstadappVault.CallOpts)
}

// IsRebalancer is a free data retrieval call binding the contract method 0x467c9eff.
//
// Solidity: function isRebalancer(address ) view returns(bool)
func (_InstadappVault *InstadappVaultCaller) IsRebalancer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "isRebalancer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRebalancer is a free data retrieval call binding the contract method 0x467c9eff.
//
// Solidity: function isRebalancer(address ) view returns(bool)
func (_InstadappVault *InstadappVaultSession) IsRebalancer(arg0 common.Address) (bool, error) {
	return _InstadappVault.Contract.IsRebalancer(&_InstadappVault.CallOpts, arg0)
}

// IsRebalancer is a free data retrieval call binding the contract method 0x467c9eff.
//
// Solidity: function isRebalancer(address ) view returns(bool)
func (_InstadappVault *InstadappVaultCallerSession) IsRebalancer(arg0 common.Address) (bool, error) {
	return _InstadappVault.Contract.IsRebalancer(&_InstadappVault.CallOpts, arg0)
}

// LastRevenueExchangePrice is a free data retrieval call binding the contract method 0x0f9775d5.
//
// Solidity: function lastRevenueExchangePrice() view returns(uint256)
func (_InstadappVault *InstadappVaultCaller) LastRevenueExchangePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "lastRevenueExchangePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRevenueExchangePrice is a free data retrieval call binding the contract method 0x0f9775d5.
//
// Solidity: function lastRevenueExchangePrice() view returns(uint256)
func (_InstadappVault *InstadappVaultSession) LastRevenueExchangePrice() (*big.Int, error) {
	return _InstadappVault.Contract.LastRevenueExchangePrice(&_InstadappVault.CallOpts)
}

// LastRevenueExchangePrice is a free data retrieval call binding the contract method 0x0f9775d5.
//
// Solidity: function lastRevenueExchangePrice() view returns(uint256)
func (_InstadappVault *InstadappVaultCallerSession) LastRevenueExchangePrice() (*big.Int, error) {
	return _InstadappVault.Contract.LastRevenueExchangePrice(&_InstadappVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_InstadappVault *InstadappVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_InstadappVault *InstadappVaultSession) Name() (string, error) {
	return _InstadappVault.Contract.Name(&_InstadappVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_InstadappVault *InstadappVaultCallerSession) Name() (string, error) {
	return _InstadappVault.Contract.Name(&_InstadappVault.CallOpts)
}

// Ratios is a free data retrieval call binding the contract method 0xcf6d625e.
//
// Solidity: function ratios() view returns(uint16 maxLimit, uint16 minLimit, uint16 minLimitGap, uint128 maxBorrowRate)
func (_InstadappVault *InstadappVaultCaller) Ratios(opts *bind.CallOpts) (struct {
	MaxLimit      uint16
	MinLimit      uint16
	MinLimitGap   uint16
	MaxBorrowRate *big.Int
}, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "ratios")

	outstruct := new(struct {
		MaxLimit      uint16
		MinLimit      uint16
		MinLimitGap   uint16
		MaxBorrowRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxLimit = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.MinLimit = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.MinLimitGap = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.MaxBorrowRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ratios is a free data retrieval call binding the contract method 0xcf6d625e.
//
// Solidity: function ratios() view returns(uint16 maxLimit, uint16 minLimit, uint16 minLimitGap, uint128 maxBorrowRate)
func (_InstadappVault *InstadappVaultSession) Ratios() (struct {
	MaxLimit      uint16
	MinLimit      uint16
	MinLimitGap   uint16
	MaxBorrowRate *big.Int
}, error) {
	return _InstadappVault.Contract.Ratios(&_InstadappVault.CallOpts)
}

// Ratios is a free data retrieval call binding the contract method 0xcf6d625e.
//
// Solidity: function ratios() view returns(uint16 maxLimit, uint16 minLimit, uint16 minLimitGap, uint128 maxBorrowRate)
func (_InstadappVault *InstadappVaultCallerSession) Ratios() (struct {
	MaxLimit      uint16
	MinLimit      uint16
	MinLimitGap   uint16
	MaxBorrowRate *big.Int
}, error) {
	return _InstadappVault.Contract.Ratios(&_InstadappVault.CallOpts)
}

// Revenue is a free data retrieval call binding the contract method 0x3e9491a2.
//
// Solidity: function revenue() view returns(uint256)
func (_InstadappVault *InstadappVaultCaller) Revenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "revenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Revenue is a free data retrieval call binding the contract method 0x3e9491a2.
//
// Solidity: function revenue() view returns(uint256)
func (_InstadappVault *InstadappVaultSession) Revenue() (*big.Int, error) {
	return _InstadappVault.Contract.Revenue(&_InstadappVault.CallOpts)
}

// Revenue is a free data retrieval call binding the contract method 0x3e9491a2.
//
// Solidity: function revenue() view returns(uint256)
func (_InstadappVault *InstadappVaultCallerSession) Revenue() (*big.Int, error) {
	return _InstadappVault.Contract.Revenue(&_InstadappVault.CallOpts)
}

// RevenueFee is a free data retrieval call binding the contract method 0x36e4ec64.
//
// Solidity: function revenueFee() view returns(uint256)
func (_InstadappVault *InstadappVaultCaller) RevenueFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "revenueFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevenueFee is a free data retrieval call binding the contract method 0x36e4ec64.
//
// Solidity: function revenueFee() view returns(uint256)
func (_InstadappVault *InstadappVaultSession) RevenueFee() (*big.Int, error) {
	return _InstadappVault.Contract.RevenueFee(&_InstadappVault.CallOpts)
}

// RevenueFee is a free data retrieval call binding the contract method 0x36e4ec64.
//
// Solidity: function revenueFee() view returns(uint256)
func (_InstadappVault *InstadappVaultCallerSession) RevenueFee() (*big.Int, error) {
	return _InstadappVault.Contract.RevenueFee(&_InstadappVault.CallOpts)
}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_InstadappVault *InstadappVaultCaller) SwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "swapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_InstadappVault *InstadappVaultSession) SwapFee() (*big.Int, error) {
	return _InstadappVault.Contract.SwapFee(&_InstadappVault.CallOpts)
}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_InstadappVault *InstadappVaultCallerSession) SwapFee() (*big.Int, error) {
	return _InstadappVault.Contract.SwapFee(&_InstadappVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_InstadappVault *InstadappVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_InstadappVault *InstadappVaultSession) Symbol() (string, error) {
	return _InstadappVault.Contract.Symbol(&_InstadappVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_InstadappVault *InstadappVaultCallerSession) Symbol() (string, error) {
	return _InstadappVault.Contract.Symbol(&_InstadappVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InstadappVault *InstadappVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InstadappVault *InstadappVaultSession) TotalSupply() (*big.Int, error) {
	return _InstadappVault.Contract.TotalSupply(&_InstadappVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InstadappVault *InstadappVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _InstadappVault.Contract.TotalSupply(&_InstadappVault.CallOpts)
}

// VaultDsa is a free data retrieval call binding the contract method 0xdc935698.
//
// Solidity: function vaultDsa() view returns(address)
func (_InstadappVault *InstadappVaultCaller) VaultDsa(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "vaultDsa")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultDsa is a free data retrieval call binding the contract method 0xdc935698.
//
// Solidity: function vaultDsa() view returns(address)
func (_InstadappVault *InstadappVaultSession) VaultDsa() (common.Address, error) {
	return _InstadappVault.Contract.VaultDsa(&_InstadappVault.CallOpts)
}

// VaultDsa is a free data retrieval call binding the contract method 0xdc935698.
//
// Solidity: function vaultDsa() view returns(address)
func (_InstadappVault *InstadappVaultCallerSession) VaultDsa() (common.Address, error) {
	return _InstadappVault.Contract.VaultDsa(&_InstadappVault.CallOpts)
}

// WithdrawFinal is a free data retrieval call binding the contract method 0x3477c81d.
//
// Solidity: function withdrawFinal(uint256 amount_, bool afterDeleverage_) view returns(uint256[] transferAmts_)
func (_InstadappVault *InstadappVaultCaller) WithdrawFinal(opts *bind.CallOpts, amount_ *big.Int, afterDeleverage_ bool) ([]*big.Int, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "withdrawFinal", amount_, afterDeleverage_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// WithdrawFinal is a free data retrieval call binding the contract method 0x3477c81d.
//
// Solidity: function withdrawFinal(uint256 amount_, bool afterDeleverage_) view returns(uint256[] transferAmts_)
func (_InstadappVault *InstadappVaultSession) WithdrawFinal(amount_ *big.Int, afterDeleverage_ bool) ([]*big.Int, error) {
	return _InstadappVault.Contract.WithdrawFinal(&_InstadappVault.CallOpts, amount_, afterDeleverage_)
}

// WithdrawFinal is a free data retrieval call binding the contract method 0x3477c81d.
//
// Solidity: function withdrawFinal(uint256 amount_, bool afterDeleverage_) view returns(uint256[] transferAmts_)
func (_InstadappVault *InstadappVaultCallerSession) WithdrawFinal(amount_ *big.Int, afterDeleverage_ bool) ([]*big.Int, error) {
	return _InstadappVault.Contract.WithdrawFinal(&_InstadappVault.CallOpts, amount_, afterDeleverage_)
}

// WithdrawalFee is a free data retrieval call binding the contract method 0x8bc7e8c4.
//
// Solidity: function withdrawalFee() view returns(uint256)
func (_InstadappVault *InstadappVaultCaller) WithdrawalFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstadappVault.contract.Call(opts, &out, "withdrawalFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalFee is a free data retrieval call binding the contract method 0x8bc7e8c4.
//
// Solidity: function withdrawalFee() view returns(uint256)
func (_InstadappVault *InstadappVaultSession) WithdrawalFee() (*big.Int, error) {
	return _InstadappVault.Contract.WithdrawalFee(&_InstadappVault.CallOpts)
}

// WithdrawalFee is a free data retrieval call binding the contract method 0x8bc7e8c4.
//
// Solidity: function withdrawalFee() view returns(uint256)
func (_InstadappVault *InstadappVaultCallerSession) WithdrawalFee() (*big.Int, error) {
	return _InstadappVault.Contract.WithdrawalFee(&_InstadappVault.CallOpts)
}

// AddDSAAuth is a paid mutator transaction binding the contract method 0x590ee346.
//
// Solidity: function addDSAAuth(address auth_) returns()
func (_InstadappVault *InstadappVaultTransactor) AddDSAAuth(opts *bind.TransactOpts, auth_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "addDSAAuth", auth_)
}

// AddDSAAuth is a paid mutator transaction binding the contract method 0x590ee346.
//
// Solidity: function addDSAAuth(address auth_) returns()
func (_InstadappVault *InstadappVaultSession) AddDSAAuth(auth_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.AddDSAAuth(&_InstadappVault.TransactOpts, auth_)
}

// AddDSAAuth is a paid mutator transaction binding the contract method 0x590ee346.
//
// Solidity: function addDSAAuth(address auth_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) AddDSAAuth(auth_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.AddDSAAuth(&_InstadappVault.TransactOpts, auth_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InstadappVault *InstadappVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InstadappVault *InstadappVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.Approve(&_InstadappVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InstadappVault *InstadappVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.Approve(&_InstadappVault.TransactOpts, spender, amount)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xe8025d77.
//
// Solidity: function changeStatus(uint256 status_) returns()
func (_InstadappVault *InstadappVaultTransactor) ChangeStatus(opts *bind.TransactOpts, status_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "changeStatus", status_)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xe8025d77.
//
// Solidity: function changeStatus(uint256 status_) returns()
func (_InstadappVault *InstadappVaultSession) ChangeStatus(status_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.ChangeStatus(&_InstadappVault.TransactOpts, status_)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xe8025d77.
//
// Solidity: function changeStatus(uint256 status_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) ChangeStatus(status_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.ChangeStatus(&_InstadappVault.TransactOpts, status_)
}

// CollectRevenue is a paid mutator transaction binding the contract method 0xb0037ada.
//
// Solidity: function collectRevenue(uint256 amount_, address to_) returns()
func (_InstadappVault *InstadappVaultTransactor) CollectRevenue(opts *bind.TransactOpts, amount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "collectRevenue", amount_, to_)
}

// CollectRevenue is a paid mutator transaction binding the contract method 0xb0037ada.
//
// Solidity: function collectRevenue(uint256 amount_, address to_) returns()
func (_InstadappVault *InstadappVaultSession) CollectRevenue(amount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.CollectRevenue(&_InstadappVault.TransactOpts, amount_, to_)
}

// CollectRevenue is a paid mutator transaction binding the contract method 0xb0037ada.
//
// Solidity: function collectRevenue(uint256 amount_, address to_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) CollectRevenue(amount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.CollectRevenue(&_InstadappVault.TransactOpts, amount_, to_)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_InstadappVault *InstadappVaultTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_InstadappVault *InstadappVaultSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.DecreaseAllowance(&_InstadappVault.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_InstadappVault *InstadappVaultTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.DecreaseAllowance(&_InstadappVault.TransactOpts, spender, subtractedValue)
}

// Deleverage is a paid mutator transaction binding the contract method 0x9fdabec2.
//
// Solidity: function deleverage(uint256 amt_) returns()
func (_InstadappVault *InstadappVaultTransactor) Deleverage(opts *bind.TransactOpts, amt_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "deleverage", amt_)
}

// Deleverage is a paid mutator transaction binding the contract method 0x9fdabec2.
//
// Solidity: function deleverage(uint256 amt_) returns()
func (_InstadappVault *InstadappVaultSession) Deleverage(amt_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.Deleverage(&_InstadappVault.TransactOpts, amt_)
}

// Deleverage is a paid mutator transaction binding the contract method 0x9fdabec2.
//
// Solidity: function deleverage(uint256 amt_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) Deleverage(amt_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.Deleverage(&_InstadappVault.TransactOpts, amt_)
}

// DeleverageAndWithdraw is a paid mutator transaction binding the contract method 0x3b3fdb5c.
//
// Solidity: function deleverageAndWithdraw(uint256 deleverageAmt_, uint256 withdrawAmount_, address to_) returns()
func (_InstadappVault *InstadappVaultTransactor) DeleverageAndWithdraw(opts *bind.TransactOpts, deleverageAmt_ *big.Int, withdrawAmount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "deleverageAndWithdraw", deleverageAmt_, withdrawAmount_, to_)
}

// DeleverageAndWithdraw is a paid mutator transaction binding the contract method 0x3b3fdb5c.
//
// Solidity: function deleverageAndWithdraw(uint256 deleverageAmt_, uint256 withdrawAmount_, address to_) returns()
func (_InstadappVault *InstadappVaultSession) DeleverageAndWithdraw(deleverageAmt_ *big.Int, withdrawAmount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.DeleverageAndWithdraw(&_InstadappVault.TransactOpts, deleverageAmt_, withdrawAmount_, to_)
}

// DeleverageAndWithdraw is a paid mutator transaction binding the contract method 0x3b3fdb5c.
//
// Solidity: function deleverageAndWithdraw(uint256 deleverageAmt_, uint256 withdrawAmount_, address to_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) DeleverageAndWithdraw(deleverageAmt_ *big.Int, withdrawAmount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.DeleverageAndWithdraw(&_InstadappVault.TransactOpts, deleverageAmt_, withdrawAmount_, to_)
}

// ImportPosition is a paid mutator transaction binding the contract method 0x98be6783.
//
// Solidity: function importPosition(address flashTkn_, uint256 flashAmt_, uint256 route_, address to_, uint256 stEthAmt_, uint256 wethAmt_) returns()
func (_InstadappVault *InstadappVaultTransactor) ImportPosition(opts *bind.TransactOpts, flashTkn_ common.Address, flashAmt_ *big.Int, route_ *big.Int, to_ common.Address, stEthAmt_ *big.Int, wethAmt_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "importPosition", flashTkn_, flashAmt_, route_, to_, stEthAmt_, wethAmt_)
}

// ImportPosition is a paid mutator transaction binding the contract method 0x98be6783.
//
// Solidity: function importPosition(address flashTkn_, uint256 flashAmt_, uint256 route_, address to_, uint256 stEthAmt_, uint256 wethAmt_) returns()
func (_InstadappVault *InstadappVaultSession) ImportPosition(flashTkn_ common.Address, flashAmt_ *big.Int, route_ *big.Int, to_ common.Address, stEthAmt_ *big.Int, wethAmt_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.ImportPosition(&_InstadappVault.TransactOpts, flashTkn_, flashAmt_, route_, to_, stEthAmt_, wethAmt_)
}

// ImportPosition is a paid mutator transaction binding the contract method 0x98be6783.
//
// Solidity: function importPosition(address flashTkn_, uint256 flashAmt_, uint256 route_, address to_, uint256 stEthAmt_, uint256 wethAmt_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) ImportPosition(flashTkn_ common.Address, flashAmt_ *big.Int, route_ *big.Int, to_ common.Address, stEthAmt_ *big.Int, wethAmt_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.ImportPosition(&_InstadappVault.TransactOpts, flashTkn_, flashAmt_, route_, to_, stEthAmt_, wethAmt_)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_InstadappVault *InstadappVaultTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_InstadappVault *InstadappVaultSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.IncreaseAllowance(&_InstadappVault.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_InstadappVault *InstadappVaultTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.IncreaseAllowance(&_InstadappVault.TransactOpts, spender, addedValue)
}

// Leverage is a paid mutator transaction binding the contract method 0xacd431a8.
//
// Solidity: function leverage(uint256 amt_) returns()
func (_InstadappVault *InstadappVaultTransactor) Leverage(opts *bind.TransactOpts, amt_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "leverage", amt_)
}

// Leverage is a paid mutator transaction binding the contract method 0xacd431a8.
//
// Solidity: function leverage(uint256 amt_) returns()
func (_InstadappVault *InstadappVaultSession) Leverage(amt_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.Leverage(&_InstadappVault.TransactOpts, amt_)
}

// Leverage is a paid mutator transaction binding the contract method 0xacd431a8.
//
// Solidity: function leverage(uint256 amt_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) Leverage(amt_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.Leverage(&_InstadappVault.TransactOpts, amt_)
}

// RebalanceOne is a paid mutator transaction binding the contract method 0x25258d0c.
//
// Solidity: function rebalanceOne(address flashTkn_, uint256 flashAmt_, uint256 route_, address[] vaults_, uint256[] amts_, uint256 excessDebt_, uint256 paybackDebt_, uint256 totalAmountToSwap_, uint256 extraWithdraw_, uint256 unitAmt_, bytes oneInchData_) returns()
func (_InstadappVault *InstadappVaultTransactor) RebalanceOne(opts *bind.TransactOpts, flashTkn_ common.Address, flashAmt_ *big.Int, route_ *big.Int, vaults_ []common.Address, amts_ []*big.Int, excessDebt_ *big.Int, paybackDebt_ *big.Int, totalAmountToSwap_ *big.Int, extraWithdraw_ *big.Int, unitAmt_ *big.Int, oneInchData_ []byte) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "rebalanceOne", flashTkn_, flashAmt_, route_, vaults_, amts_, excessDebt_, paybackDebt_, totalAmountToSwap_, extraWithdraw_, unitAmt_, oneInchData_)
}

// RebalanceOne is a paid mutator transaction binding the contract method 0x25258d0c.
//
// Solidity: function rebalanceOne(address flashTkn_, uint256 flashAmt_, uint256 route_, address[] vaults_, uint256[] amts_, uint256 excessDebt_, uint256 paybackDebt_, uint256 totalAmountToSwap_, uint256 extraWithdraw_, uint256 unitAmt_, bytes oneInchData_) returns()
func (_InstadappVault *InstadappVaultSession) RebalanceOne(flashTkn_ common.Address, flashAmt_ *big.Int, route_ *big.Int, vaults_ []common.Address, amts_ []*big.Int, excessDebt_ *big.Int, paybackDebt_ *big.Int, totalAmountToSwap_ *big.Int, extraWithdraw_ *big.Int, unitAmt_ *big.Int, oneInchData_ []byte) (*types.Transaction, error) {
	return _InstadappVault.Contract.RebalanceOne(&_InstadappVault.TransactOpts, flashTkn_, flashAmt_, route_, vaults_, amts_, excessDebt_, paybackDebt_, totalAmountToSwap_, extraWithdraw_, unitAmt_, oneInchData_)
}

// RebalanceOne is a paid mutator transaction binding the contract method 0x25258d0c.
//
// Solidity: function rebalanceOne(address flashTkn_, uint256 flashAmt_, uint256 route_, address[] vaults_, uint256[] amts_, uint256 excessDebt_, uint256 paybackDebt_, uint256 totalAmountToSwap_, uint256 extraWithdraw_, uint256 unitAmt_, bytes oneInchData_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) RebalanceOne(flashTkn_ common.Address, flashAmt_ *big.Int, route_ *big.Int, vaults_ []common.Address, amts_ []*big.Int, excessDebt_ *big.Int, paybackDebt_ *big.Int, totalAmountToSwap_ *big.Int, extraWithdraw_ *big.Int, unitAmt_ *big.Int, oneInchData_ []byte) (*types.Transaction, error) {
	return _InstadappVault.Contract.RebalanceOne(&_InstadappVault.TransactOpts, flashTkn_, flashAmt_, route_, vaults_, amts_, excessDebt_, paybackDebt_, totalAmountToSwap_, extraWithdraw_, unitAmt_, oneInchData_)
}

// RebalanceTwo is a paid mutator transaction binding the contract method 0xf02105a5.
//
// Solidity: function rebalanceTwo(uint256 withdrawAmt_, address flashTkn_, uint256 flashAmt_, uint256 route_, uint256 totalAmountToSwap_, uint256 unitAmt_, bytes oneInchData_) returns()
func (_InstadappVault *InstadappVaultTransactor) RebalanceTwo(opts *bind.TransactOpts, withdrawAmt_ *big.Int, flashTkn_ common.Address, flashAmt_ *big.Int, route_ *big.Int, totalAmountToSwap_ *big.Int, unitAmt_ *big.Int, oneInchData_ []byte) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "rebalanceTwo", withdrawAmt_, flashTkn_, flashAmt_, route_, totalAmountToSwap_, unitAmt_, oneInchData_)
}

// RebalanceTwo is a paid mutator transaction binding the contract method 0xf02105a5.
//
// Solidity: function rebalanceTwo(uint256 withdrawAmt_, address flashTkn_, uint256 flashAmt_, uint256 route_, uint256 totalAmountToSwap_, uint256 unitAmt_, bytes oneInchData_) returns()
func (_InstadappVault *InstadappVaultSession) RebalanceTwo(withdrawAmt_ *big.Int, flashTkn_ common.Address, flashAmt_ *big.Int, route_ *big.Int, totalAmountToSwap_ *big.Int, unitAmt_ *big.Int, oneInchData_ []byte) (*types.Transaction, error) {
	return _InstadappVault.Contract.RebalanceTwo(&_InstadappVault.TransactOpts, withdrawAmt_, flashTkn_, flashAmt_, route_, totalAmountToSwap_, unitAmt_, oneInchData_)
}

// RebalanceTwo is a paid mutator transaction binding the contract method 0xf02105a5.
//
// Solidity: function rebalanceTwo(uint256 withdrawAmt_, address flashTkn_, uint256 flashAmt_, uint256 route_, uint256 totalAmountToSwap_, uint256 unitAmt_, bytes oneInchData_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) RebalanceTwo(withdrawAmt_ *big.Int, flashTkn_ common.Address, flashAmt_ *big.Int, route_ *big.Int, totalAmountToSwap_ *big.Int, unitAmt_ *big.Int, oneInchData_ []byte) (*types.Transaction, error) {
	return _InstadappVault.Contract.RebalanceTwo(&_InstadappVault.TransactOpts, withdrawAmt_, flashTkn_, flashAmt_, route_, totalAmountToSwap_, unitAmt_, oneInchData_)
}

// Spell is a paid mutator transaction binding the contract method 0x15ff627d.
//
// Solidity: function spell(address to_, bytes calldata_, uint256 value_, uint256 operation_) returns()
func (_InstadappVault *InstadappVaultTransactor) Spell(opts *bind.TransactOpts, to_ common.Address, calldata_ []byte, value_ *big.Int, operation_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "spell", to_, calldata_, value_, operation_)
}

// Spell is a paid mutator transaction binding the contract method 0x15ff627d.
//
// Solidity: function spell(address to_, bytes calldata_, uint256 value_, uint256 operation_) returns()
func (_InstadappVault *InstadappVaultSession) Spell(to_ common.Address, calldata_ []byte, value_ *big.Int, operation_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.Spell(&_InstadappVault.TransactOpts, to_, calldata_, value_, operation_)
}

// Spell is a paid mutator transaction binding the contract method 0x15ff627d.
//
// Solidity: function spell(address to_, bytes calldata_, uint256 value_, uint256 operation_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) Spell(to_ common.Address, calldata_ []byte, value_ *big.Int, operation_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.Spell(&_InstadappVault.TransactOpts, to_, calldata_, value_, operation_)
}

// Supply is a paid mutator transaction binding the contract method 0x8b2a4df5.
//
// Solidity: function supply(address token_, uint256 amount_, address to_) returns(uint256 vtokenAmount_)
func (_InstadappVault *InstadappVaultTransactor) Supply(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "supply", token_, amount_, to_)
}

// Supply is a paid mutator transaction binding the contract method 0x8b2a4df5.
//
// Solidity: function supply(address token_, uint256 amount_, address to_) returns(uint256 vtokenAmount_)
func (_InstadappVault *InstadappVaultSession) Supply(token_ common.Address, amount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.Supply(&_InstadappVault.TransactOpts, token_, amount_, to_)
}

// Supply is a paid mutator transaction binding the contract method 0x8b2a4df5.
//
// Solidity: function supply(address token_, uint256 amount_, address to_) returns(uint256 vtokenAmount_)
func (_InstadappVault *InstadappVaultTransactorSession) Supply(token_ common.Address, amount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.Supply(&_InstadappVault.TransactOpts, token_, amount_, to_)
}

// SupplyEth is a paid mutator transaction binding the contract method 0x87ee9312.
//
// Solidity: function supplyEth(address to_) returns(uint256 vtokenAmount_)
func (_InstadappVault *InstadappVaultTransactor) SupplyEth(opts *bind.TransactOpts, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "supplyEth", to_)
}

// SupplyEth is a paid mutator transaction binding the contract method 0x87ee9312.
//
// Solidity: function supplyEth(address to_) returns(uint256 vtokenAmount_)
func (_InstadappVault *InstadappVaultSession) SupplyEth(to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.SupplyEth(&_InstadappVault.TransactOpts, to_)
}

// SupplyEth is a paid mutator transaction binding the contract method 0x87ee9312.
//
// Solidity: function supplyEth(address to_) returns(uint256 vtokenAmount_)
func (_InstadappVault *InstadappVaultTransactorSession) SupplyEth(to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.SupplyEth(&_InstadappVault.TransactOpts, to_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_InstadappVault *InstadappVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_InstadappVault *InstadappVaultSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.Transfer(&_InstadappVault.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_InstadappVault *InstadappVaultTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.Transfer(&_InstadappVault.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_InstadappVault *InstadappVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_InstadappVault *InstadappVaultSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.TransferFrom(&_InstadappVault.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_InstadappVault *InstadappVaultTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.TransferFrom(&_InstadappVault.TransactOpts, from, to, amount)
}

// UpdateAuth is a paid mutator transaction binding the contract method 0xb8d07f4f.
//
// Solidity: function updateAuth(address auth_) returns()
func (_InstadappVault *InstadappVaultTransactor) UpdateAuth(opts *bind.TransactOpts, auth_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "updateAuth", auth_)
}

// UpdateAuth is a paid mutator transaction binding the contract method 0xb8d07f4f.
//
// Solidity: function updateAuth(address auth_) returns()
func (_InstadappVault *InstadappVaultSession) UpdateAuth(auth_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.UpdateAuth(&_InstadappVault.TransactOpts, auth_)
}

// UpdateAuth is a paid mutator transaction binding the contract method 0xb8d07f4f.
//
// Solidity: function updateAuth(address auth_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) UpdateAuth(auth_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.UpdateAuth(&_InstadappVault.TransactOpts, auth_)
}

// UpdateFees is a paid mutator transaction binding the contract method 0xc6616ba1.
//
// Solidity: function updateFees(uint256 revenueFee_, uint256 withdrawalFee_, uint256 swapFee_, uint256 deleverageFee_) returns()
func (_InstadappVault *InstadappVaultTransactor) UpdateFees(opts *bind.TransactOpts, revenueFee_ *big.Int, withdrawalFee_ *big.Int, swapFee_ *big.Int, deleverageFee_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "updateFees", revenueFee_, withdrawalFee_, swapFee_, deleverageFee_)
}

// UpdateFees is a paid mutator transaction binding the contract method 0xc6616ba1.
//
// Solidity: function updateFees(uint256 revenueFee_, uint256 withdrawalFee_, uint256 swapFee_, uint256 deleverageFee_) returns()
func (_InstadappVault *InstadappVaultSession) UpdateFees(revenueFee_ *big.Int, withdrawalFee_ *big.Int, swapFee_ *big.Int, deleverageFee_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.UpdateFees(&_InstadappVault.TransactOpts, revenueFee_, withdrawalFee_, swapFee_, deleverageFee_)
}

// UpdateFees is a paid mutator transaction binding the contract method 0xc6616ba1.
//
// Solidity: function updateFees(uint256 revenueFee_, uint256 withdrawalFee_, uint256 swapFee_, uint256 deleverageFee_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) UpdateFees(revenueFee_ *big.Int, withdrawalFee_ *big.Int, swapFee_ *big.Int, deleverageFee_ *big.Int) (*types.Transaction, error) {
	return _InstadappVault.Contract.UpdateFees(&_InstadappVault.TransactOpts, revenueFee_, withdrawalFee_, swapFee_, deleverageFee_)
}

// UpdateRatios is a paid mutator transaction binding the contract method 0x7c37411c.
//
// Solidity: function updateRatios(uint16[] ratios_) returns()
func (_InstadappVault *InstadappVaultTransactor) UpdateRatios(opts *bind.TransactOpts, ratios_ []uint16) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "updateRatios", ratios_)
}

// UpdateRatios is a paid mutator transaction binding the contract method 0x7c37411c.
//
// Solidity: function updateRatios(uint16[] ratios_) returns()
func (_InstadappVault *InstadappVaultSession) UpdateRatios(ratios_ []uint16) (*types.Transaction, error) {
	return _InstadappVault.Contract.UpdateRatios(&_InstadappVault.TransactOpts, ratios_)
}

// UpdateRatios is a paid mutator transaction binding the contract method 0x7c37411c.
//
// Solidity: function updateRatios(uint16[] ratios_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) UpdateRatios(ratios_ []uint16) (*types.Transaction, error) {
	return _InstadappVault.Contract.UpdateRatios(&_InstadappVault.TransactOpts, ratios_)
}

// UpdateRebalancer is a paid mutator transaction binding the contract method 0x0de30836.
//
// Solidity: function updateRebalancer(address rebalancer_, bool isRebalancer_) returns()
func (_InstadappVault *InstadappVaultTransactor) UpdateRebalancer(opts *bind.TransactOpts, rebalancer_ common.Address, isRebalancer_ bool) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "updateRebalancer", rebalancer_, isRebalancer_)
}

// UpdateRebalancer is a paid mutator transaction binding the contract method 0x0de30836.
//
// Solidity: function updateRebalancer(address rebalancer_, bool isRebalancer_) returns()
func (_InstadappVault *InstadappVaultSession) UpdateRebalancer(rebalancer_ common.Address, isRebalancer_ bool) (*types.Transaction, error) {
	return _InstadappVault.Contract.UpdateRebalancer(&_InstadappVault.TransactOpts, rebalancer_, isRebalancer_)
}

// UpdateRebalancer is a paid mutator transaction binding the contract method 0x0de30836.
//
// Solidity: function updateRebalancer(address rebalancer_, bool isRebalancer_) returns()
func (_InstadappVault *InstadappVaultTransactorSession) UpdateRebalancer(rebalancer_ common.Address, isRebalancer_ bool) (*types.Transaction, error) {
	return _InstadappVault.Contract.UpdateRebalancer(&_InstadappVault.TransactOpts, rebalancer_, isRebalancer_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount_, address to_) returns(uint256 vtokenAmount_)
func (_InstadappVault *InstadappVaultTransactor) Withdraw(opts *bind.TransactOpts, amount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.contract.Transact(opts, "withdraw", amount_, to_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount_, address to_) returns(uint256 vtokenAmount_)
func (_InstadappVault *InstadappVaultSession) Withdraw(amount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.Withdraw(&_InstadappVault.TransactOpts, amount_, to_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount_, address to_) returns(uint256 vtokenAmount_)
func (_InstadappVault *InstadappVaultTransactorSession) Withdraw(amount_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _InstadappVault.Contract.Withdraw(&_InstadappVault.TransactOpts, amount_, to_)
}
