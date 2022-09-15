// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveRoutersteCRV

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

// CurveRoutersteCRVMetaData contains all meta data concerning the CurveRoutersteCRV contract.
var CurveRoutersteCRVMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A_precise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_actions_deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"admin_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apply_new_fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"bool\",\"name\":\"is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"new_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_admin_fee\",\"type\":\"uint256\"}],\"name\":\"commit_new_fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate_admin_fees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill_me\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_future_A\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_future_time\",\"type\":\"uint256\"}],\"name\":\"ramp_A\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"_min_amounts\",\"type\":\"uint256[2]\"}],\"name\":\"remove_liquidity\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_max_burn_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_imbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"_min_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revert_new_parameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stop_ramp_A\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unkill_me\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CurveRoutersteCRVABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveRoutersteCRVMetaData.ABI instead.
var CurveRoutersteCRVABI = CurveRoutersteCRVMetaData.ABI

// CurveRoutersteCRV is an auto generated Go binding around an Ethereum contract.
type CurveRoutersteCRV struct {
	CurveRoutersteCRVCaller     // Read-only binding to the contract
	CurveRoutersteCRVTransactor // Write-only binding to the contract
	CurveRoutersteCRVFilterer   // Log filterer for contract events
}

// CurveRoutersteCRVCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveRoutersteCRVCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRoutersteCRVTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveRoutersteCRVTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRoutersteCRVFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveRoutersteCRVFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRoutersteCRVSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveRoutersteCRVSession struct {
	Contract     *CurveRoutersteCRV // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CurveRoutersteCRVCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveRoutersteCRVCallerSession struct {
	Contract *CurveRoutersteCRVCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CurveRoutersteCRVTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveRoutersteCRVTransactorSession struct {
	Contract     *CurveRoutersteCRVTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CurveRoutersteCRVRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveRoutersteCRVRaw struct {
	Contract *CurveRoutersteCRV // Generic contract binding to access the raw methods on
}

// CurveRoutersteCRVCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveRoutersteCRVCallerRaw struct {
	Contract *CurveRoutersteCRVCaller // Generic read-only contract binding to access the raw methods on
}

// CurveRoutersteCRVTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveRoutersteCRVTransactorRaw struct {
	Contract *CurveRoutersteCRVTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveRoutersteCRV creates a new instance of CurveRoutersteCRV, bound to a specific deployed contract.
func NewCurveRoutersteCRV(address common.Address, backend bind.ContractBackend) (*CurveRoutersteCRV, error) {
	contract, err := bindCurveRoutersteCRV(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveRoutersteCRV{CurveRoutersteCRVCaller: CurveRoutersteCRVCaller{contract: contract}, CurveRoutersteCRVTransactor: CurveRoutersteCRVTransactor{contract: contract}, CurveRoutersteCRVFilterer: CurveRoutersteCRVFilterer{contract: contract}}, nil
}

// NewCurveRoutersteCRVCaller creates a new read-only instance of CurveRoutersteCRV, bound to a specific deployed contract.
func NewCurveRoutersteCRVCaller(address common.Address, caller bind.ContractCaller) (*CurveRoutersteCRVCaller, error) {
	contract, err := bindCurveRoutersteCRV(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveRoutersteCRVCaller{contract: contract}, nil
}

// NewCurveRoutersteCRVTransactor creates a new write-only instance of CurveRoutersteCRV, bound to a specific deployed contract.
func NewCurveRoutersteCRVTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveRoutersteCRVTransactor, error) {
	contract, err := bindCurveRoutersteCRV(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveRoutersteCRVTransactor{contract: contract}, nil
}

// NewCurveRoutersteCRVFilterer creates a new log filterer instance of CurveRoutersteCRV, bound to a specific deployed contract.
func NewCurveRoutersteCRVFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveRoutersteCRVFilterer, error) {
	contract, err := bindCurveRoutersteCRV(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveRoutersteCRVFilterer{contract: contract}, nil
}

// bindCurveRoutersteCRV binds a generic wrapper to an already deployed contract.
func bindCurveRoutersteCRV(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveRoutersteCRVABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveRoutersteCRV *CurveRoutersteCRVRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveRoutersteCRV.Contract.CurveRoutersteCRVCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveRoutersteCRV *CurveRoutersteCRVRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.CurveRoutersteCRVTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveRoutersteCRV *CurveRoutersteCRVRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.CurveRoutersteCRVTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveRoutersteCRV.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) A() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.A(&_CurveRoutersteCRV.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) A() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.A(&_CurveRoutersteCRV.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) APrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "A_precise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) APrecise() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.APrecise(&_CurveRoutersteCRV.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) APrecise() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.APrecise(&_CurveRoutersteCRV.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.AdminActionsDeadline(&_CurveRoutersteCRV.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.AdminActionsDeadline(&_CurveRoutersteCRV.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) AdminBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "admin_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.AdminBalances(&_CurveRoutersteCRV.CallOpts, arg0)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.AdminBalances(&_CurveRoutersteCRV.CallOpts, arg0)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) AdminFee() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.AdminFee(&_CurveRoutersteCRV.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) AdminFee() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.AdminFee(&_CurveRoutersteCRV.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) Balances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) Balances(i *big.Int) (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.Balances(&_CurveRoutersteCRV.CallOpts, i)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) Balances(i *big.Int) (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.Balances(&_CurveRoutersteCRV.CallOpts, i)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool is_deposit) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [2]*big.Int, is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "calc_token_amount", amounts, is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool is_deposit) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) CalcTokenAmount(amounts [2]*big.Int, is_deposit bool) (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.CalcTokenAmount(&_CurveRoutersteCRV.CallOpts, amounts, is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] amounts, bool is_deposit) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) CalcTokenAmount(amounts [2]*big.Int, is_deposit bool) (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.CalcTokenAmount(&_CurveRoutersteCRV.CallOpts, amounts, is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "calc_withdraw_one_coin", _token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.CalcWithdrawOneCoin(&_CurveRoutersteCRV.CallOpts, _token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.CalcWithdrawOneCoin(&_CurveRoutersteCRV.CallOpts, _token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurveRoutersteCRV.Contract.Coins(&_CurveRoutersteCRV.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurveRoutersteCRV.Contract.Coins(&_CurveRoutersteCRV.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) Fee() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.Fee(&_CurveRoutersteCRV.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) Fee() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.Fee(&_CurveRoutersteCRV.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) FutureA() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.FutureA(&_CurveRoutersteCRV.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) FutureA() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.FutureA(&_CurveRoutersteCRV.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) FutureATime() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.FutureATime(&_CurveRoutersteCRV.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) FutureATime() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.FutureATime(&_CurveRoutersteCRV.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) FutureAdminFee() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.FutureAdminFee(&_CurveRoutersteCRV.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) FutureAdminFee() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.FutureAdminFee(&_CurveRoutersteCRV.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "future_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) FutureFee() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.FutureFee(&_CurveRoutersteCRV.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) FutureFee() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.FutureFee(&_CurveRoutersteCRV.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) FutureOwner() (common.Address, error) {
	return _CurveRoutersteCRV.Contract.FutureOwner(&_CurveRoutersteCRV.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) FutureOwner() (common.Address, error) {
	return _CurveRoutersteCRV.Contract.FutureOwner(&_CurveRoutersteCRV.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.GetDy(&_CurveRoutersteCRV.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.GetDy(&_CurveRoutersteCRV.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.GetVirtualPrice(&_CurveRoutersteCRV.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.GetVirtualPrice(&_CurveRoutersteCRV.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) InitialA() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.InitialA(&_CurveRoutersteCRV.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) InitialA() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.InitialA(&_CurveRoutersteCRV.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) InitialATime() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.InitialATime(&_CurveRoutersteCRV.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) InitialATime() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.InitialATime(&_CurveRoutersteCRV.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "lp_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) LpToken() (common.Address, error) {
	return _CurveRoutersteCRV.Contract.LpToken(&_CurveRoutersteCRV.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) LpToken() (common.Address, error) {
	return _CurveRoutersteCRV.Contract.LpToken(&_CurveRoutersteCRV.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) Owner() (common.Address, error) {
	return _CurveRoutersteCRV.Contract.Owner(&_CurveRoutersteCRV.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) Owner() (common.Address, error) {
	return _CurveRoutersteCRV.Contract.Owner(&_CurveRoutersteCRV.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRoutersteCRV.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.TransferOwnershipDeadline(&_CurveRoutersteCRV.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurveRoutersteCRV.Contract.TransferOwnershipDeadline(&_CurveRoutersteCRV.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.AddLiquidity(&_CurveRoutersteCRV.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.AddLiquidity(&_CurveRoutersteCRV.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) ApplyNewFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "apply_new_fee")
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) ApplyNewFee() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.ApplyNewFee(&_CurveRoutersteCRV.TransactOpts)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) ApplyNewFee() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.ApplyNewFee(&_CurveRoutersteCRV.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.ApplyTransferOwnership(&_CurveRoutersteCRV.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.ApplyTransferOwnership(&_CurveRoutersteCRV.TransactOpts)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) CommitNewFee(opts *bind.TransactOpts, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "commit_new_fee", new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.CommitNewFee(&_CurveRoutersteCRV.TransactOpts, new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.CommitNewFee(&_CurveRoutersteCRV.TransactOpts, new_fee, new_admin_fee)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.CommitTransferOwnership(&_CurveRoutersteCRV.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.CommitTransferOwnership(&_CurveRoutersteCRV.TransactOpts, _owner)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) DonateAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "donate_admin_fees")
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) DonateAdminFees() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.DonateAdminFees(&_CurveRoutersteCRV.TransactOpts)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) DonateAdminFees() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.DonateAdminFees(&_CurveRoutersteCRV.TransactOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.Exchange(&_CurveRoutersteCRV.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.Exchange(&_CurveRoutersteCRV.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) KillMe() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.KillMe(&_CurveRoutersteCRV.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) KillMe() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.KillMe(&_CurveRoutersteCRV.TransactOpts)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RampA(&_CurveRoutersteCRV.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RampA(&_CurveRoutersteCRV.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "remove_liquidity", _amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) RemoveLiquidity(_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RemoveLiquidity(&_CurveRoutersteCRV.TransactOpts, _amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) RemoveLiquidity(_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RemoveLiquidity(&_CurveRoutersteCRV.TransactOpts, _amount, _min_amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, _amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "remove_liquidity_imbalance", _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) RemoveLiquidityImbalance(_amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RemoveLiquidityImbalance(&_CurveRoutersteCRV.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) RemoveLiquidityImbalance(_amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RemoveLiquidityImbalance(&_CurveRoutersteCRV.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount) returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "remove_liquidity_one_coin", _token_amount, i, _min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount) returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RemoveLiquidityOneCoin(&_CurveRoutersteCRV.TransactOpts, _token_amount, i, _min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount) returns(uint256)
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RemoveLiquidityOneCoin(&_CurveRoutersteCRV.TransactOpts, _token_amount, i, _min_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RevertNewParameters(&_CurveRoutersteCRV.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RevertNewParameters(&_CurveRoutersteCRV.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RevertTransferOwnership(&_CurveRoutersteCRV.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.RevertTransferOwnership(&_CurveRoutersteCRV.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) StopRampA() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.StopRampA(&_CurveRoutersteCRV.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) StopRampA() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.StopRampA(&_CurveRoutersteCRV.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) UnkillMe() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.UnkillMe(&_CurveRoutersteCRV.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.UnkillMe(&_CurveRoutersteCRV.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRoutersteCRV.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.WithdrawAdminFees(&_CurveRoutersteCRV.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveRoutersteCRV *CurveRoutersteCRVTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurveRoutersteCRV.Contract.WithdrawAdminFees(&_CurveRoutersteCRV.TransactOpts)
}
