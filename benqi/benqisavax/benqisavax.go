// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package benqisavax

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

// BenqisavaxMetaData contains all meta data concerning the Benqisavax contract.
var BenqisavaxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_ACCRUE_REWARDS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_DEPOSIT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_PAUSE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_PAUSE_MINTING\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_RESUME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_RESUME_MINTING\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_SET_TOTAL_POOLED_AVAX_CAP\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_WITHDRAW\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"accrueRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelPendingUnlockRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelRedeemableUnlockRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unlockIndex\",\"type\":\"uint256\"}],\"name\":\"cancelUnlockRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shareAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledAvaxByShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"avaxAmount\",\"type\":\"uint256\"}],\"name\":\"getSharesByPooledAvax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUnlockRequestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"historicalExchangeRateTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"historicalExchangeRatesByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_redeemPeriod\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintingPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseMinting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unlockIndex\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemOverdueShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unlockIndex\",\"type\":\"uint256\"}],\"name\":\"redeemOverdueShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shareAmount\",\"type\":\"uint256\"}],\"name\":\"requestUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeMinting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCooldownPeriod\",\"type\":\"uint256\"}],\"name\":\"setCooldownPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRedeemPeriod\",\"type\":\"uint256\"}],\"name\":\"setRedeemPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTotalPooledAvaxCap\",\"type\":\"uint256\"}],\"name\":\"setTotalPooledAvaxCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPooledAvax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPooledAvaxCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userSharesInCustody\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userUnlockRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shareAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BenqisavaxABI is the input ABI used to generate the binding from.
// Deprecated: Use BenqisavaxMetaData.ABI instead.
var BenqisavaxABI = BenqisavaxMetaData.ABI

// Benqisavax is an auto generated Go binding around an Ethereum contract.
type Benqisavax struct {
	BenqisavaxCaller     // Read-only binding to the contract
	BenqisavaxTransactor // Write-only binding to the contract
	BenqisavaxFilterer   // Log filterer for contract events
}

// BenqisavaxCaller is an auto generated read-only Go binding around an Ethereum contract.
type BenqisavaxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqisavaxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BenqisavaxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqisavaxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BenqisavaxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqisavaxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BenqisavaxSession struct {
	Contract     *Benqisavax       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BenqisavaxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BenqisavaxCallerSession struct {
	Contract *BenqisavaxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BenqisavaxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BenqisavaxTransactorSession struct {
	Contract     *BenqisavaxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BenqisavaxRaw is an auto generated low-level Go binding around an Ethereum contract.
type BenqisavaxRaw struct {
	Contract *Benqisavax // Generic contract binding to access the raw methods on
}

// BenqisavaxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BenqisavaxCallerRaw struct {
	Contract *BenqisavaxCaller // Generic read-only contract binding to access the raw methods on
}

// BenqisavaxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BenqisavaxTransactorRaw struct {
	Contract *BenqisavaxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBenqisavax creates a new instance of Benqisavax, bound to a specific deployed contract.
func NewBenqisavax(address common.Address, backend bind.ContractBackend) (*Benqisavax, error) {
	contract, err := bindBenqisavax(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Benqisavax{BenqisavaxCaller: BenqisavaxCaller{contract: contract}, BenqisavaxTransactor: BenqisavaxTransactor{contract: contract}, BenqisavaxFilterer: BenqisavaxFilterer{contract: contract}}, nil
}

// NewBenqisavaxCaller creates a new read-only instance of Benqisavax, bound to a specific deployed contract.
func NewBenqisavaxCaller(address common.Address, caller bind.ContractCaller) (*BenqisavaxCaller, error) {
	contract, err := bindBenqisavax(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BenqisavaxCaller{contract: contract}, nil
}

// NewBenqisavaxTransactor creates a new write-only instance of Benqisavax, bound to a specific deployed contract.
func NewBenqisavaxTransactor(address common.Address, transactor bind.ContractTransactor) (*BenqisavaxTransactor, error) {
	contract, err := bindBenqisavax(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BenqisavaxTransactor{contract: contract}, nil
}

// NewBenqisavaxFilterer creates a new log filterer instance of Benqisavax, bound to a specific deployed contract.
func NewBenqisavaxFilterer(address common.Address, filterer bind.ContractFilterer) (*BenqisavaxFilterer, error) {
	contract, err := bindBenqisavax(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BenqisavaxFilterer{contract: contract}, nil
}

// bindBenqisavax binds a generic wrapper to an already deployed contract.
func bindBenqisavax(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BenqisavaxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Benqisavax *BenqisavaxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Benqisavax.Contract.BenqisavaxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Benqisavax *BenqisavaxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.Contract.BenqisavaxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Benqisavax *BenqisavaxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Benqisavax.Contract.BenqisavaxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Benqisavax *BenqisavaxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Benqisavax.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Benqisavax *BenqisavaxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Benqisavax *BenqisavaxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Benqisavax.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Benqisavax *BenqisavaxCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Benqisavax *BenqisavaxSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Benqisavax.Contract.DEFAULTADMINROLE(&_Benqisavax.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Benqisavax *BenqisavaxCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Benqisavax.Contract.DEFAULTADMINROLE(&_Benqisavax.CallOpts)
}

// ROLEACCRUEREWARDS is a free data retrieval call binding the contract method 0xa905ff93.
//
// Solidity: function ROLE_ACCRUE_REWARDS() view returns(bytes32)
func (_Benqisavax *BenqisavaxCaller) ROLEACCRUEREWARDS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "ROLE_ACCRUE_REWARDS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLEACCRUEREWARDS is a free data retrieval call binding the contract method 0xa905ff93.
//
// Solidity: function ROLE_ACCRUE_REWARDS() view returns(bytes32)
func (_Benqisavax *BenqisavaxSession) ROLEACCRUEREWARDS() ([32]byte, error) {
	return _Benqisavax.Contract.ROLEACCRUEREWARDS(&_Benqisavax.CallOpts)
}

// ROLEACCRUEREWARDS is a free data retrieval call binding the contract method 0xa905ff93.
//
// Solidity: function ROLE_ACCRUE_REWARDS() view returns(bytes32)
func (_Benqisavax *BenqisavaxCallerSession) ROLEACCRUEREWARDS() ([32]byte, error) {
	return _Benqisavax.Contract.ROLEACCRUEREWARDS(&_Benqisavax.CallOpts)
}

// ROLEDEPOSIT is a free data retrieval call binding the contract method 0xc2d78654.
//
// Solidity: function ROLE_DEPOSIT() view returns(bytes32)
func (_Benqisavax *BenqisavaxCaller) ROLEDEPOSIT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "ROLE_DEPOSIT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLEDEPOSIT is a free data retrieval call binding the contract method 0xc2d78654.
//
// Solidity: function ROLE_DEPOSIT() view returns(bytes32)
func (_Benqisavax *BenqisavaxSession) ROLEDEPOSIT() ([32]byte, error) {
	return _Benqisavax.Contract.ROLEDEPOSIT(&_Benqisavax.CallOpts)
}

// ROLEDEPOSIT is a free data retrieval call binding the contract method 0xc2d78654.
//
// Solidity: function ROLE_DEPOSIT() view returns(bytes32)
func (_Benqisavax *BenqisavaxCallerSession) ROLEDEPOSIT() ([32]byte, error) {
	return _Benqisavax.Contract.ROLEDEPOSIT(&_Benqisavax.CallOpts)
}

// ROLEPAUSE is a free data retrieval call binding the contract method 0x4ff0241a.
//
// Solidity: function ROLE_PAUSE() view returns(bytes32)
func (_Benqisavax *BenqisavaxCaller) ROLEPAUSE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "ROLE_PAUSE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLEPAUSE is a free data retrieval call binding the contract method 0x4ff0241a.
//
// Solidity: function ROLE_PAUSE() view returns(bytes32)
func (_Benqisavax *BenqisavaxSession) ROLEPAUSE() ([32]byte, error) {
	return _Benqisavax.Contract.ROLEPAUSE(&_Benqisavax.CallOpts)
}

// ROLEPAUSE is a free data retrieval call binding the contract method 0x4ff0241a.
//
// Solidity: function ROLE_PAUSE() view returns(bytes32)
func (_Benqisavax *BenqisavaxCallerSession) ROLEPAUSE() ([32]byte, error) {
	return _Benqisavax.Contract.ROLEPAUSE(&_Benqisavax.CallOpts)
}

// ROLEPAUSEMINTING is a free data retrieval call binding the contract method 0x3fc777b3.
//
// Solidity: function ROLE_PAUSE_MINTING() view returns(bytes32)
func (_Benqisavax *BenqisavaxCaller) ROLEPAUSEMINTING(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "ROLE_PAUSE_MINTING")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLEPAUSEMINTING is a free data retrieval call binding the contract method 0x3fc777b3.
//
// Solidity: function ROLE_PAUSE_MINTING() view returns(bytes32)
func (_Benqisavax *BenqisavaxSession) ROLEPAUSEMINTING() ([32]byte, error) {
	return _Benqisavax.Contract.ROLEPAUSEMINTING(&_Benqisavax.CallOpts)
}

// ROLEPAUSEMINTING is a free data retrieval call binding the contract method 0x3fc777b3.
//
// Solidity: function ROLE_PAUSE_MINTING() view returns(bytes32)
func (_Benqisavax *BenqisavaxCallerSession) ROLEPAUSEMINTING() ([32]byte, error) {
	return _Benqisavax.Contract.ROLEPAUSEMINTING(&_Benqisavax.CallOpts)
}

// ROLERESUME is a free data retrieval call binding the contract method 0xc1db6588.
//
// Solidity: function ROLE_RESUME() view returns(bytes32)
func (_Benqisavax *BenqisavaxCaller) ROLERESUME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "ROLE_RESUME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLERESUME is a free data retrieval call binding the contract method 0xc1db6588.
//
// Solidity: function ROLE_RESUME() view returns(bytes32)
func (_Benqisavax *BenqisavaxSession) ROLERESUME() ([32]byte, error) {
	return _Benqisavax.Contract.ROLERESUME(&_Benqisavax.CallOpts)
}

// ROLERESUME is a free data retrieval call binding the contract method 0xc1db6588.
//
// Solidity: function ROLE_RESUME() view returns(bytes32)
func (_Benqisavax *BenqisavaxCallerSession) ROLERESUME() ([32]byte, error) {
	return _Benqisavax.Contract.ROLERESUME(&_Benqisavax.CallOpts)
}

// ROLERESUMEMINTING is a free data retrieval call binding the contract method 0x1ab8ab25.
//
// Solidity: function ROLE_RESUME_MINTING() view returns(bytes32)
func (_Benqisavax *BenqisavaxCaller) ROLERESUMEMINTING(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "ROLE_RESUME_MINTING")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLERESUMEMINTING is a free data retrieval call binding the contract method 0x1ab8ab25.
//
// Solidity: function ROLE_RESUME_MINTING() view returns(bytes32)
func (_Benqisavax *BenqisavaxSession) ROLERESUMEMINTING() ([32]byte, error) {
	return _Benqisavax.Contract.ROLERESUMEMINTING(&_Benqisavax.CallOpts)
}

// ROLERESUMEMINTING is a free data retrieval call binding the contract method 0x1ab8ab25.
//
// Solidity: function ROLE_RESUME_MINTING() view returns(bytes32)
func (_Benqisavax *BenqisavaxCallerSession) ROLERESUMEMINTING() ([32]byte, error) {
	return _Benqisavax.Contract.ROLERESUMEMINTING(&_Benqisavax.CallOpts)
}

// ROLESETTOTALPOOLEDAVAXCAP is a free data retrieval call binding the contract method 0x4757c0d2.
//
// Solidity: function ROLE_SET_TOTAL_POOLED_AVAX_CAP() view returns(bytes32)
func (_Benqisavax *BenqisavaxCaller) ROLESETTOTALPOOLEDAVAXCAP(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "ROLE_SET_TOTAL_POOLED_AVAX_CAP")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLESETTOTALPOOLEDAVAXCAP is a free data retrieval call binding the contract method 0x4757c0d2.
//
// Solidity: function ROLE_SET_TOTAL_POOLED_AVAX_CAP() view returns(bytes32)
func (_Benqisavax *BenqisavaxSession) ROLESETTOTALPOOLEDAVAXCAP() ([32]byte, error) {
	return _Benqisavax.Contract.ROLESETTOTALPOOLEDAVAXCAP(&_Benqisavax.CallOpts)
}

// ROLESETTOTALPOOLEDAVAXCAP is a free data retrieval call binding the contract method 0x4757c0d2.
//
// Solidity: function ROLE_SET_TOTAL_POOLED_AVAX_CAP() view returns(bytes32)
func (_Benqisavax *BenqisavaxCallerSession) ROLESETTOTALPOOLEDAVAXCAP() ([32]byte, error) {
	return _Benqisavax.Contract.ROLESETTOTALPOOLEDAVAXCAP(&_Benqisavax.CallOpts)
}

// ROLEWITHDRAW is a free data retrieval call binding the contract method 0x4b7e23dc.
//
// Solidity: function ROLE_WITHDRAW() view returns(bytes32)
func (_Benqisavax *BenqisavaxCaller) ROLEWITHDRAW(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "ROLE_WITHDRAW")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLEWITHDRAW is a free data retrieval call binding the contract method 0x4b7e23dc.
//
// Solidity: function ROLE_WITHDRAW() view returns(bytes32)
func (_Benqisavax *BenqisavaxSession) ROLEWITHDRAW() ([32]byte, error) {
	return _Benqisavax.Contract.ROLEWITHDRAW(&_Benqisavax.CallOpts)
}

// ROLEWITHDRAW is a free data retrieval call binding the contract method 0x4b7e23dc.
//
// Solidity: function ROLE_WITHDRAW() view returns(bytes32)
func (_Benqisavax *BenqisavaxCallerSession) ROLEWITHDRAW() ([32]byte, error) {
	return _Benqisavax.Contract.ROLEWITHDRAW(&_Benqisavax.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Benqisavax *BenqisavaxSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Benqisavax.Contract.Allowance(&_Benqisavax.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Benqisavax.Contract.Allowance(&_Benqisavax.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Benqisavax *BenqisavaxSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Benqisavax.Contract.BalanceOf(&_Benqisavax.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Benqisavax.Contract.BalanceOf(&_Benqisavax.CallOpts, account)
}

// CooldownPeriod is a free data retrieval call binding the contract method 0x04646a49.
//
// Solidity: function cooldownPeriod() view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) CooldownPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "cooldownPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownPeriod is a free data retrieval call binding the contract method 0x04646a49.
//
// Solidity: function cooldownPeriod() view returns(uint256)
func (_Benqisavax *BenqisavaxSession) CooldownPeriod() (*big.Int, error) {
	return _Benqisavax.Contract.CooldownPeriod(&_Benqisavax.CallOpts)
}

// CooldownPeriod is a free data retrieval call binding the contract method 0x04646a49.
//
// Solidity: function cooldownPeriod() view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) CooldownPeriod() (*big.Int, error) {
	return _Benqisavax.Contract.CooldownPeriod(&_Benqisavax.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Benqisavax *BenqisavaxCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Benqisavax *BenqisavaxSession) Decimals() (uint8, error) {
	return _Benqisavax.Contract.Decimals(&_Benqisavax.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Benqisavax *BenqisavaxCallerSession) Decimals() (uint8, error) {
	return _Benqisavax.Contract.Decimals(&_Benqisavax.CallOpts)
}

// GetPooledAvaxByShares is a free data retrieval call binding the contract method 0x4a36d6c1.
//
// Solidity: function getPooledAvaxByShares(uint256 shareAmount) view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) GetPooledAvaxByShares(opts *bind.CallOpts, shareAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "getPooledAvaxByShares", shareAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledAvaxByShares is a free data retrieval call binding the contract method 0x4a36d6c1.
//
// Solidity: function getPooledAvaxByShares(uint256 shareAmount) view returns(uint256)
func (_Benqisavax *BenqisavaxSession) GetPooledAvaxByShares(shareAmount *big.Int) (*big.Int, error) {
	return _Benqisavax.Contract.GetPooledAvaxByShares(&_Benqisavax.CallOpts, shareAmount)
}

// GetPooledAvaxByShares is a free data retrieval call binding the contract method 0x4a36d6c1.
//
// Solidity: function getPooledAvaxByShares(uint256 shareAmount) view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) GetPooledAvaxByShares(shareAmount *big.Int) (*big.Int, error) {
	return _Benqisavax.Contract.GetPooledAvaxByShares(&_Benqisavax.CallOpts, shareAmount)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Benqisavax *BenqisavaxCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Benqisavax *BenqisavaxSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Benqisavax.Contract.GetRoleAdmin(&_Benqisavax.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Benqisavax *BenqisavaxCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Benqisavax.Contract.GetRoleAdmin(&_Benqisavax.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Benqisavax *BenqisavaxCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Benqisavax *BenqisavaxSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Benqisavax.Contract.GetRoleMember(&_Benqisavax.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Benqisavax *BenqisavaxCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Benqisavax.Contract.GetRoleMember(&_Benqisavax.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Benqisavax *BenqisavaxSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Benqisavax.Contract.GetRoleMemberCount(&_Benqisavax.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Benqisavax.Contract.GetRoleMemberCount(&_Benqisavax.CallOpts, role)
}

// GetSharesByPooledAvax is a free data retrieval call binding the contract method 0xf1ee8d92.
//
// Solidity: function getSharesByPooledAvax(uint256 avaxAmount) view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) GetSharesByPooledAvax(opts *bind.CallOpts, avaxAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "getSharesByPooledAvax", avaxAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledAvax is a free data retrieval call binding the contract method 0xf1ee8d92.
//
// Solidity: function getSharesByPooledAvax(uint256 avaxAmount) view returns(uint256)
func (_Benqisavax *BenqisavaxSession) GetSharesByPooledAvax(avaxAmount *big.Int) (*big.Int, error) {
	return _Benqisavax.Contract.GetSharesByPooledAvax(&_Benqisavax.CallOpts, avaxAmount)
}

// GetSharesByPooledAvax is a free data retrieval call binding the contract method 0xf1ee8d92.
//
// Solidity: function getSharesByPooledAvax(uint256 avaxAmount) view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) GetSharesByPooledAvax(avaxAmount *big.Int) (*big.Int, error) {
	return _Benqisavax.Contract.GetSharesByPooledAvax(&_Benqisavax.CallOpts, avaxAmount)
}

// GetUnlockRequestCount is a free data retrieval call binding the contract method 0xc423f9a8.
//
// Solidity: function getUnlockRequestCount(address user) view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) GetUnlockRequestCount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "getUnlockRequestCount", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnlockRequestCount is a free data retrieval call binding the contract method 0xc423f9a8.
//
// Solidity: function getUnlockRequestCount(address user) view returns(uint256)
func (_Benqisavax *BenqisavaxSession) GetUnlockRequestCount(user common.Address) (*big.Int, error) {
	return _Benqisavax.Contract.GetUnlockRequestCount(&_Benqisavax.CallOpts, user)
}

// GetUnlockRequestCount is a free data retrieval call binding the contract method 0xc423f9a8.
//
// Solidity: function getUnlockRequestCount(address user) view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) GetUnlockRequestCount(user common.Address) (*big.Int, error) {
	return _Benqisavax.Contract.GetUnlockRequestCount(&_Benqisavax.CallOpts, user)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Benqisavax *BenqisavaxCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Benqisavax *BenqisavaxSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Benqisavax.Contract.HasRole(&_Benqisavax.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Benqisavax *BenqisavaxCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Benqisavax.Contract.HasRole(&_Benqisavax.CallOpts, role, account)
}

// HistoricalExchangeRateTimestamps is a free data retrieval call binding the contract method 0x1b2b3a2f.
//
// Solidity: function historicalExchangeRateTimestamps(uint256 ) view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) HistoricalExchangeRateTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "historicalExchangeRateTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HistoricalExchangeRateTimestamps is a free data retrieval call binding the contract method 0x1b2b3a2f.
//
// Solidity: function historicalExchangeRateTimestamps(uint256 ) view returns(uint256)
func (_Benqisavax *BenqisavaxSession) HistoricalExchangeRateTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _Benqisavax.Contract.HistoricalExchangeRateTimestamps(&_Benqisavax.CallOpts, arg0)
}

// HistoricalExchangeRateTimestamps is a free data retrieval call binding the contract method 0x1b2b3a2f.
//
// Solidity: function historicalExchangeRateTimestamps(uint256 ) view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) HistoricalExchangeRateTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _Benqisavax.Contract.HistoricalExchangeRateTimestamps(&_Benqisavax.CallOpts, arg0)
}

// HistoricalExchangeRatesByTimestamp is a free data retrieval call binding the contract method 0x0a732ce6.
//
// Solidity: function historicalExchangeRatesByTimestamp(uint256 ) view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) HistoricalExchangeRatesByTimestamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "historicalExchangeRatesByTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HistoricalExchangeRatesByTimestamp is a free data retrieval call binding the contract method 0x0a732ce6.
//
// Solidity: function historicalExchangeRatesByTimestamp(uint256 ) view returns(uint256)
func (_Benqisavax *BenqisavaxSession) HistoricalExchangeRatesByTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _Benqisavax.Contract.HistoricalExchangeRatesByTimestamp(&_Benqisavax.CallOpts, arg0)
}

// HistoricalExchangeRatesByTimestamp is a free data retrieval call binding the contract method 0x0a732ce6.
//
// Solidity: function historicalExchangeRatesByTimestamp(uint256 ) view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) HistoricalExchangeRatesByTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _Benqisavax.Contract.HistoricalExchangeRatesByTimestamp(&_Benqisavax.CallOpts, arg0)
}

// MintingPaused is a free data retrieval call binding the contract method 0xe1a283d6.
//
// Solidity: function mintingPaused() view returns(bool)
func (_Benqisavax *BenqisavaxCaller) MintingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "mintingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintingPaused is a free data retrieval call binding the contract method 0xe1a283d6.
//
// Solidity: function mintingPaused() view returns(bool)
func (_Benqisavax *BenqisavaxSession) MintingPaused() (bool, error) {
	return _Benqisavax.Contract.MintingPaused(&_Benqisavax.CallOpts)
}

// MintingPaused is a free data retrieval call binding the contract method 0xe1a283d6.
//
// Solidity: function mintingPaused() view returns(bool)
func (_Benqisavax *BenqisavaxCallerSession) MintingPaused() (bool, error) {
	return _Benqisavax.Contract.MintingPaused(&_Benqisavax.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Benqisavax *BenqisavaxCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Benqisavax *BenqisavaxSession) Name() (string, error) {
	return _Benqisavax.Contract.Name(&_Benqisavax.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Benqisavax *BenqisavaxCallerSession) Name() (string, error) {
	return _Benqisavax.Contract.Name(&_Benqisavax.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Benqisavax *BenqisavaxCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Benqisavax *BenqisavaxSession) Paused() (bool, error) {
	return _Benqisavax.Contract.Paused(&_Benqisavax.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Benqisavax *BenqisavaxCallerSession) Paused() (bool, error) {
	return _Benqisavax.Contract.Paused(&_Benqisavax.CallOpts)
}

// RedeemPeriod is a free data retrieval call binding the contract method 0x40a233a6.
//
// Solidity: function redeemPeriod() view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) RedeemPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "redeemPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemPeriod is a free data retrieval call binding the contract method 0x40a233a6.
//
// Solidity: function redeemPeriod() view returns(uint256)
func (_Benqisavax *BenqisavaxSession) RedeemPeriod() (*big.Int, error) {
	return _Benqisavax.Contract.RedeemPeriod(&_Benqisavax.CallOpts)
}

// RedeemPeriod is a free data retrieval call binding the contract method 0x40a233a6.
//
// Solidity: function redeemPeriod() view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) RedeemPeriod() (*big.Int, error) {
	return _Benqisavax.Contract.RedeemPeriod(&_Benqisavax.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Benqisavax *BenqisavaxSession) StakerCount() (*big.Int, error) {
	return _Benqisavax.Contract.StakerCount(&_Benqisavax.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) StakerCount() (*big.Int, error) {
	return _Benqisavax.Contract.StakerCount(&_Benqisavax.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Benqisavax *BenqisavaxCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Benqisavax *BenqisavaxSession) Symbol() (string, error) {
	return _Benqisavax.Contract.Symbol(&_Benqisavax.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Benqisavax *BenqisavaxCallerSession) Symbol() (string, error) {
	return _Benqisavax.Contract.Symbol(&_Benqisavax.CallOpts)
}

// TotalPooledAvax is a free data retrieval call binding the contract method 0x629e8056.
//
// Solidity: function totalPooledAvax() view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) TotalPooledAvax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "totalPooledAvax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPooledAvax is a free data retrieval call binding the contract method 0x629e8056.
//
// Solidity: function totalPooledAvax() view returns(uint256)
func (_Benqisavax *BenqisavaxSession) TotalPooledAvax() (*big.Int, error) {
	return _Benqisavax.Contract.TotalPooledAvax(&_Benqisavax.CallOpts)
}

// TotalPooledAvax is a free data retrieval call binding the contract method 0x629e8056.
//
// Solidity: function totalPooledAvax() view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) TotalPooledAvax() (*big.Int, error) {
	return _Benqisavax.Contract.TotalPooledAvax(&_Benqisavax.CallOpts)
}

// TotalPooledAvaxCap is a free data retrieval call binding the contract method 0x5cd47487.
//
// Solidity: function totalPooledAvaxCap() view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) TotalPooledAvaxCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "totalPooledAvaxCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPooledAvaxCap is a free data retrieval call binding the contract method 0x5cd47487.
//
// Solidity: function totalPooledAvaxCap() view returns(uint256)
func (_Benqisavax *BenqisavaxSession) TotalPooledAvaxCap() (*big.Int, error) {
	return _Benqisavax.Contract.TotalPooledAvaxCap(&_Benqisavax.CallOpts)
}

// TotalPooledAvaxCap is a free data retrieval call binding the contract method 0x5cd47487.
//
// Solidity: function totalPooledAvaxCap() view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) TotalPooledAvaxCap() (*big.Int, error) {
	return _Benqisavax.Contract.TotalPooledAvaxCap(&_Benqisavax.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Benqisavax *BenqisavaxSession) TotalShares() (*big.Int, error) {
	return _Benqisavax.Contract.TotalShares(&_Benqisavax.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) TotalShares() (*big.Int, error) {
	return _Benqisavax.Contract.TotalShares(&_Benqisavax.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Benqisavax *BenqisavaxSession) TotalSupply() (*big.Int, error) {
	return _Benqisavax.Contract.TotalSupply(&_Benqisavax.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) TotalSupply() (*big.Int, error) {
	return _Benqisavax.Contract.TotalSupply(&_Benqisavax.CallOpts)
}

// UserSharesInCustody is a free data retrieval call binding the contract method 0x5d039525.
//
// Solidity: function userSharesInCustody(address ) view returns(uint256)
func (_Benqisavax *BenqisavaxCaller) UserSharesInCustody(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "userSharesInCustody", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserSharesInCustody is a free data retrieval call binding the contract method 0x5d039525.
//
// Solidity: function userSharesInCustody(address ) view returns(uint256)
func (_Benqisavax *BenqisavaxSession) UserSharesInCustody(arg0 common.Address) (*big.Int, error) {
	return _Benqisavax.Contract.UserSharesInCustody(&_Benqisavax.CallOpts, arg0)
}

// UserSharesInCustody is a free data retrieval call binding the contract method 0x5d039525.
//
// Solidity: function userSharesInCustody(address ) view returns(uint256)
func (_Benqisavax *BenqisavaxCallerSession) UserSharesInCustody(arg0 common.Address) (*big.Int, error) {
	return _Benqisavax.Contract.UserSharesInCustody(&_Benqisavax.CallOpts, arg0)
}

// UserUnlockRequests is a free data retrieval call binding the contract method 0xfd012e34.
//
// Solidity: function userUnlockRequests(address , uint256 ) view returns(uint256 startedAt, uint256 shareAmount)
func (_Benqisavax *BenqisavaxCaller) UserUnlockRequests(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StartedAt   *big.Int
	ShareAmount *big.Int
}, error) {
	var out []interface{}
	err := _Benqisavax.contract.Call(opts, &out, "userUnlockRequests", arg0, arg1)

	outstruct := new(struct {
		StartedAt   *big.Int
		ShareAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartedAt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ShareAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserUnlockRequests is a free data retrieval call binding the contract method 0xfd012e34.
//
// Solidity: function userUnlockRequests(address , uint256 ) view returns(uint256 startedAt, uint256 shareAmount)
func (_Benqisavax *BenqisavaxSession) UserUnlockRequests(arg0 common.Address, arg1 *big.Int) (struct {
	StartedAt   *big.Int
	ShareAmount *big.Int
}, error) {
	return _Benqisavax.Contract.UserUnlockRequests(&_Benqisavax.CallOpts, arg0, arg1)
}

// UserUnlockRequests is a free data retrieval call binding the contract method 0xfd012e34.
//
// Solidity: function userUnlockRequests(address , uint256 ) view returns(uint256 startedAt, uint256 shareAmount)
func (_Benqisavax *BenqisavaxCallerSession) UserUnlockRequests(arg0 common.Address, arg1 *big.Int) (struct {
	StartedAt   *big.Int
	ShareAmount *big.Int
}, error) {
	return _Benqisavax.Contract.UserUnlockRequests(&_Benqisavax.CallOpts, arg0, arg1)
}

// AccrueRewards is a paid mutator transaction binding the contract method 0xe1a472b9.
//
// Solidity: function accrueRewards(uint256 amount) returns()
func (_Benqisavax *BenqisavaxTransactor) AccrueRewards(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "accrueRewards", amount)
}

// AccrueRewards is a paid mutator transaction binding the contract method 0xe1a472b9.
//
// Solidity: function accrueRewards(uint256 amount) returns()
func (_Benqisavax *BenqisavaxSession) AccrueRewards(amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.AccrueRewards(&_Benqisavax.TransactOpts, amount)
}

// AccrueRewards is a paid mutator transaction binding the contract method 0xe1a472b9.
//
// Solidity: function accrueRewards(uint256 amount) returns()
func (_Benqisavax *BenqisavaxTransactorSession) AccrueRewards(amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.AccrueRewards(&_Benqisavax.TransactOpts, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Benqisavax *BenqisavaxTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Benqisavax *BenqisavaxSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.Approve(&_Benqisavax.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Benqisavax *BenqisavaxTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.Approve(&_Benqisavax.TransactOpts, spender, amount)
}

// CancelPendingUnlockRequests is a paid mutator transaction binding the contract method 0x01550f64.
//
// Solidity: function cancelPendingUnlockRequests() returns()
func (_Benqisavax *BenqisavaxTransactor) CancelPendingUnlockRequests(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "cancelPendingUnlockRequests")
}

// CancelPendingUnlockRequests is a paid mutator transaction binding the contract method 0x01550f64.
//
// Solidity: function cancelPendingUnlockRequests() returns()
func (_Benqisavax *BenqisavaxSession) CancelPendingUnlockRequests() (*types.Transaction, error) {
	return _Benqisavax.Contract.CancelPendingUnlockRequests(&_Benqisavax.TransactOpts)
}

// CancelPendingUnlockRequests is a paid mutator transaction binding the contract method 0x01550f64.
//
// Solidity: function cancelPendingUnlockRequests() returns()
func (_Benqisavax *BenqisavaxTransactorSession) CancelPendingUnlockRequests() (*types.Transaction, error) {
	return _Benqisavax.Contract.CancelPendingUnlockRequests(&_Benqisavax.TransactOpts)
}

// CancelRedeemableUnlockRequests is a paid mutator transaction binding the contract method 0xd1f1ca04.
//
// Solidity: function cancelRedeemableUnlockRequests() returns()
func (_Benqisavax *BenqisavaxTransactor) CancelRedeemableUnlockRequests(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "cancelRedeemableUnlockRequests")
}

// CancelRedeemableUnlockRequests is a paid mutator transaction binding the contract method 0xd1f1ca04.
//
// Solidity: function cancelRedeemableUnlockRequests() returns()
func (_Benqisavax *BenqisavaxSession) CancelRedeemableUnlockRequests() (*types.Transaction, error) {
	return _Benqisavax.Contract.CancelRedeemableUnlockRequests(&_Benqisavax.TransactOpts)
}

// CancelRedeemableUnlockRequests is a paid mutator transaction binding the contract method 0xd1f1ca04.
//
// Solidity: function cancelRedeemableUnlockRequests() returns()
func (_Benqisavax *BenqisavaxTransactorSession) CancelRedeemableUnlockRequests() (*types.Transaction, error) {
	return _Benqisavax.Contract.CancelRedeemableUnlockRequests(&_Benqisavax.TransactOpts)
}

// CancelUnlockRequest is a paid mutator transaction binding the contract method 0x1610247b.
//
// Solidity: function cancelUnlockRequest(uint256 unlockIndex) returns()
func (_Benqisavax *BenqisavaxTransactor) CancelUnlockRequest(opts *bind.TransactOpts, unlockIndex *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "cancelUnlockRequest", unlockIndex)
}

// CancelUnlockRequest is a paid mutator transaction binding the contract method 0x1610247b.
//
// Solidity: function cancelUnlockRequest(uint256 unlockIndex) returns()
func (_Benqisavax *BenqisavaxSession) CancelUnlockRequest(unlockIndex *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.CancelUnlockRequest(&_Benqisavax.TransactOpts, unlockIndex)
}

// CancelUnlockRequest is a paid mutator transaction binding the contract method 0x1610247b.
//
// Solidity: function cancelUnlockRequest(uint256 unlockIndex) returns()
func (_Benqisavax *BenqisavaxTransactorSession) CancelUnlockRequest(unlockIndex *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.CancelUnlockRequest(&_Benqisavax.TransactOpts, unlockIndex)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Benqisavax *BenqisavaxTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Benqisavax *BenqisavaxSession) Deposit() (*types.Transaction, error) {
	return _Benqisavax.Contract.Deposit(&_Benqisavax.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Benqisavax *BenqisavaxTransactorSession) Deposit() (*types.Transaction, error) {
	return _Benqisavax.Contract.Deposit(&_Benqisavax.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Benqisavax *BenqisavaxTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Benqisavax *BenqisavaxSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Benqisavax.Contract.GrantRole(&_Benqisavax.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Benqisavax *BenqisavaxTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Benqisavax.Contract.GrantRole(&_Benqisavax.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _cooldownPeriod, uint256 _redeemPeriod) returns()
func (_Benqisavax *BenqisavaxTransactor) Initialize(opts *bind.TransactOpts, _cooldownPeriod *big.Int, _redeemPeriod *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "initialize", _cooldownPeriod, _redeemPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _cooldownPeriod, uint256 _redeemPeriod) returns()
func (_Benqisavax *BenqisavaxSession) Initialize(_cooldownPeriod *big.Int, _redeemPeriod *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.Initialize(&_Benqisavax.TransactOpts, _cooldownPeriod, _redeemPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _cooldownPeriod, uint256 _redeemPeriod) returns()
func (_Benqisavax *BenqisavaxTransactorSession) Initialize(_cooldownPeriod *big.Int, _redeemPeriod *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.Initialize(&_Benqisavax.TransactOpts, _cooldownPeriod, _redeemPeriod)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Benqisavax *BenqisavaxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Benqisavax *BenqisavaxSession) Pause() (*types.Transaction, error) {
	return _Benqisavax.Contract.Pause(&_Benqisavax.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Benqisavax *BenqisavaxTransactorSession) Pause() (*types.Transaction, error) {
	return _Benqisavax.Contract.Pause(&_Benqisavax.TransactOpts)
}

// PauseMinting is a paid mutator transaction binding the contract method 0xda8fbf2a.
//
// Solidity: function pauseMinting() returns()
func (_Benqisavax *BenqisavaxTransactor) PauseMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "pauseMinting")
}

// PauseMinting is a paid mutator transaction binding the contract method 0xda8fbf2a.
//
// Solidity: function pauseMinting() returns()
func (_Benqisavax *BenqisavaxSession) PauseMinting() (*types.Transaction, error) {
	return _Benqisavax.Contract.PauseMinting(&_Benqisavax.TransactOpts)
}

// PauseMinting is a paid mutator transaction binding the contract method 0xda8fbf2a.
//
// Solidity: function pauseMinting() returns()
func (_Benqisavax *BenqisavaxTransactorSession) PauseMinting() (*types.Transaction, error) {
	return _Benqisavax.Contract.PauseMinting(&_Benqisavax.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns()
func (_Benqisavax *BenqisavaxTransactor) Redeem(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "redeem")
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns()
func (_Benqisavax *BenqisavaxSession) Redeem() (*types.Transaction, error) {
	return _Benqisavax.Contract.Redeem(&_Benqisavax.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns()
func (_Benqisavax *BenqisavaxTransactorSession) Redeem() (*types.Transaction, error) {
	return _Benqisavax.Contract.Redeem(&_Benqisavax.TransactOpts)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 unlockIndex) returns()
func (_Benqisavax *BenqisavaxTransactor) Redeem0(opts *bind.TransactOpts, unlockIndex *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "redeem0", unlockIndex)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 unlockIndex) returns()
func (_Benqisavax *BenqisavaxSession) Redeem0(unlockIndex *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.Redeem0(&_Benqisavax.TransactOpts, unlockIndex)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 unlockIndex) returns()
func (_Benqisavax *BenqisavaxTransactorSession) Redeem0(unlockIndex *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.Redeem0(&_Benqisavax.TransactOpts, unlockIndex)
}

// RedeemOverdueShares is a paid mutator transaction binding the contract method 0x0d10d32c.
//
// Solidity: function redeemOverdueShares() returns()
func (_Benqisavax *BenqisavaxTransactor) RedeemOverdueShares(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "redeemOverdueShares")
}

// RedeemOverdueShares is a paid mutator transaction binding the contract method 0x0d10d32c.
//
// Solidity: function redeemOverdueShares() returns()
func (_Benqisavax *BenqisavaxSession) RedeemOverdueShares() (*types.Transaction, error) {
	return _Benqisavax.Contract.RedeemOverdueShares(&_Benqisavax.TransactOpts)
}

// RedeemOverdueShares is a paid mutator transaction binding the contract method 0x0d10d32c.
//
// Solidity: function redeemOverdueShares() returns()
func (_Benqisavax *BenqisavaxTransactorSession) RedeemOverdueShares() (*types.Transaction, error) {
	return _Benqisavax.Contract.RedeemOverdueShares(&_Benqisavax.TransactOpts)
}

// RedeemOverdueShares0 is a paid mutator transaction binding the contract method 0x0f7e2048.
//
// Solidity: function redeemOverdueShares(uint256 unlockIndex) returns()
func (_Benqisavax *BenqisavaxTransactor) RedeemOverdueShares0(opts *bind.TransactOpts, unlockIndex *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "redeemOverdueShares0", unlockIndex)
}

// RedeemOverdueShares0 is a paid mutator transaction binding the contract method 0x0f7e2048.
//
// Solidity: function redeemOverdueShares(uint256 unlockIndex) returns()
func (_Benqisavax *BenqisavaxSession) RedeemOverdueShares0(unlockIndex *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.RedeemOverdueShares0(&_Benqisavax.TransactOpts, unlockIndex)
}

// RedeemOverdueShares0 is a paid mutator transaction binding the contract method 0x0f7e2048.
//
// Solidity: function redeemOverdueShares(uint256 unlockIndex) returns()
func (_Benqisavax *BenqisavaxTransactorSession) RedeemOverdueShares0(unlockIndex *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.RedeemOverdueShares0(&_Benqisavax.TransactOpts, unlockIndex)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Benqisavax *BenqisavaxTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Benqisavax *BenqisavaxSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Benqisavax.Contract.RenounceRole(&_Benqisavax.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Benqisavax *BenqisavaxTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Benqisavax.Contract.RenounceRole(&_Benqisavax.TransactOpts, role, account)
}

// RequestUnlock is a paid mutator transaction binding the contract method 0xc9d2ff9d.
//
// Solidity: function requestUnlock(uint256 shareAmount) returns()
func (_Benqisavax *BenqisavaxTransactor) RequestUnlock(opts *bind.TransactOpts, shareAmount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "requestUnlock", shareAmount)
}

// RequestUnlock is a paid mutator transaction binding the contract method 0xc9d2ff9d.
//
// Solidity: function requestUnlock(uint256 shareAmount) returns()
func (_Benqisavax *BenqisavaxSession) RequestUnlock(shareAmount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.RequestUnlock(&_Benqisavax.TransactOpts, shareAmount)
}

// RequestUnlock is a paid mutator transaction binding the contract method 0xc9d2ff9d.
//
// Solidity: function requestUnlock(uint256 shareAmount) returns()
func (_Benqisavax *BenqisavaxTransactorSession) RequestUnlock(shareAmount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.RequestUnlock(&_Benqisavax.TransactOpts, shareAmount)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Benqisavax *BenqisavaxTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Benqisavax *BenqisavaxSession) Resume() (*types.Transaction, error) {
	return _Benqisavax.Contract.Resume(&_Benqisavax.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Benqisavax *BenqisavaxTransactorSession) Resume() (*types.Transaction, error) {
	return _Benqisavax.Contract.Resume(&_Benqisavax.TransactOpts)
}

// ResumeMinting is a paid mutator transaction binding the contract method 0x59ae340e.
//
// Solidity: function resumeMinting() returns()
func (_Benqisavax *BenqisavaxTransactor) ResumeMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "resumeMinting")
}

// ResumeMinting is a paid mutator transaction binding the contract method 0x59ae340e.
//
// Solidity: function resumeMinting() returns()
func (_Benqisavax *BenqisavaxSession) ResumeMinting() (*types.Transaction, error) {
	return _Benqisavax.Contract.ResumeMinting(&_Benqisavax.TransactOpts)
}

// ResumeMinting is a paid mutator transaction binding the contract method 0x59ae340e.
//
// Solidity: function resumeMinting() returns()
func (_Benqisavax *BenqisavaxTransactorSession) ResumeMinting() (*types.Transaction, error) {
	return _Benqisavax.Contract.ResumeMinting(&_Benqisavax.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Benqisavax *BenqisavaxTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Benqisavax *BenqisavaxSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Benqisavax.Contract.RevokeRole(&_Benqisavax.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Benqisavax *BenqisavaxTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Benqisavax.Contract.RevokeRole(&_Benqisavax.TransactOpts, role, account)
}

// SetCooldownPeriod is a paid mutator transaction binding the contract method 0x80ea3de1.
//
// Solidity: function setCooldownPeriod(uint256 newCooldownPeriod) returns()
func (_Benqisavax *BenqisavaxTransactor) SetCooldownPeriod(opts *bind.TransactOpts, newCooldownPeriod *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "setCooldownPeriod", newCooldownPeriod)
}

// SetCooldownPeriod is a paid mutator transaction binding the contract method 0x80ea3de1.
//
// Solidity: function setCooldownPeriod(uint256 newCooldownPeriod) returns()
func (_Benqisavax *BenqisavaxSession) SetCooldownPeriod(newCooldownPeriod *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.SetCooldownPeriod(&_Benqisavax.TransactOpts, newCooldownPeriod)
}

// SetCooldownPeriod is a paid mutator transaction binding the contract method 0x80ea3de1.
//
// Solidity: function setCooldownPeriod(uint256 newCooldownPeriod) returns()
func (_Benqisavax *BenqisavaxTransactorSession) SetCooldownPeriod(newCooldownPeriod *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.SetCooldownPeriod(&_Benqisavax.TransactOpts, newCooldownPeriod)
}

// SetRedeemPeriod is a paid mutator transaction binding the contract method 0xf0d82e84.
//
// Solidity: function setRedeemPeriod(uint256 newRedeemPeriod) returns()
func (_Benqisavax *BenqisavaxTransactor) SetRedeemPeriod(opts *bind.TransactOpts, newRedeemPeriod *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "setRedeemPeriod", newRedeemPeriod)
}

// SetRedeemPeriod is a paid mutator transaction binding the contract method 0xf0d82e84.
//
// Solidity: function setRedeemPeriod(uint256 newRedeemPeriod) returns()
func (_Benqisavax *BenqisavaxSession) SetRedeemPeriod(newRedeemPeriod *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.SetRedeemPeriod(&_Benqisavax.TransactOpts, newRedeemPeriod)
}

// SetRedeemPeriod is a paid mutator transaction binding the contract method 0xf0d82e84.
//
// Solidity: function setRedeemPeriod(uint256 newRedeemPeriod) returns()
func (_Benqisavax *BenqisavaxTransactorSession) SetRedeemPeriod(newRedeemPeriod *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.SetRedeemPeriod(&_Benqisavax.TransactOpts, newRedeemPeriod)
}

// SetTotalPooledAvaxCap is a paid mutator transaction binding the contract method 0xada03b38.
//
// Solidity: function setTotalPooledAvaxCap(uint256 newTotalPooledAvaxCap) returns()
func (_Benqisavax *BenqisavaxTransactor) SetTotalPooledAvaxCap(opts *bind.TransactOpts, newTotalPooledAvaxCap *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "setTotalPooledAvaxCap", newTotalPooledAvaxCap)
}

// SetTotalPooledAvaxCap is a paid mutator transaction binding the contract method 0xada03b38.
//
// Solidity: function setTotalPooledAvaxCap(uint256 newTotalPooledAvaxCap) returns()
func (_Benqisavax *BenqisavaxSession) SetTotalPooledAvaxCap(newTotalPooledAvaxCap *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.SetTotalPooledAvaxCap(&_Benqisavax.TransactOpts, newTotalPooledAvaxCap)
}

// SetTotalPooledAvaxCap is a paid mutator transaction binding the contract method 0xada03b38.
//
// Solidity: function setTotalPooledAvaxCap(uint256 newTotalPooledAvaxCap) returns()
func (_Benqisavax *BenqisavaxTransactorSession) SetTotalPooledAvaxCap(newTotalPooledAvaxCap *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.SetTotalPooledAvaxCap(&_Benqisavax.TransactOpts, newTotalPooledAvaxCap)
}

// Submit is a paid mutator transaction binding the contract method 0x5bcb2fc6.
//
// Solidity: function submit() payable returns(uint256)
func (_Benqisavax *BenqisavaxTransactor) Submit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "submit")
}

// Submit is a paid mutator transaction binding the contract method 0x5bcb2fc6.
//
// Solidity: function submit() payable returns(uint256)
func (_Benqisavax *BenqisavaxSession) Submit() (*types.Transaction, error) {
	return _Benqisavax.Contract.Submit(&_Benqisavax.TransactOpts)
}

// Submit is a paid mutator transaction binding the contract method 0x5bcb2fc6.
//
// Solidity: function submit() payable returns(uint256)
func (_Benqisavax *BenqisavaxTransactorSession) Submit() (*types.Transaction, error) {
	return _Benqisavax.Contract.Submit(&_Benqisavax.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Benqisavax *BenqisavaxTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Benqisavax *BenqisavaxSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.Transfer(&_Benqisavax.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Benqisavax *BenqisavaxTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.Transfer(&_Benqisavax.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Benqisavax *BenqisavaxTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Benqisavax *BenqisavaxSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.TransferFrom(&_Benqisavax.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Benqisavax *BenqisavaxTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.TransferFrom(&_Benqisavax.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Benqisavax *BenqisavaxTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Benqisavax *BenqisavaxSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.Withdraw(&_Benqisavax.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Benqisavax *BenqisavaxTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Benqisavax.Contract.Withdraw(&_Benqisavax.TransactOpts, amount)
}
