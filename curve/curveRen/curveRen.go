// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveRen

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

// CurveRenMetaData contains all meta data concerning the CurveRen contract.
var CurveRenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A_precise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_min_mint_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_use_underlying\",\"type\":\"bool\"}],\"name\":\"add_liquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_actions_deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"admin_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee_receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apply_new_fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"bool\",\"name\":\"is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"new_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_admin_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_offpeg_fee_multiplier\",\"type\":\"uint256\"}],\"name\":\"commit_new_fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate_admin_fees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"}],\"name\":\"dynamic_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_offpeg_fee_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill_me\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offpeg_fee_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_future_A\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_future_time\",\"type\":\"uint256\"}],\"name\":\"ramp_A\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"_min_amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"bool\",\"name\":\"_use_underlying\",\"type\":\"bool\"}],\"name\":\"remove_liquidity\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"_min_amounts\",\"type\":\"uint256[2]\"}],\"name\":\"remove_liquidity\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_max_burn_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_use_underlying\",\"type\":\"bool\"}],\"name\":\"remove_liquidity_imbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_max_burn_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_imbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"_min_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"_min_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_use_underlying\",\"type\":\"bool\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revert_new_parameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward_receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"referral_code\",\"type\":\"uint256\"}],\"name\":\"set_aave_referral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin_fee_receiver\",\"type\":\"address\"}],\"name\":\"set_admin_fee_receiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reward_receiver\",\"type\":\"address\"}],\"name\":\"set_reward_receiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stop_ramp_A\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unkill_me\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CurveRenABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveRenMetaData.ABI instead.
var CurveRenABI = CurveRenMetaData.ABI

// CurveRen is an auto generated Go binding around an Ethereum contract.
type CurveRen struct {
	CurveRenCaller     // Read-only binding to the contract
	CurveRenTransactor // Write-only binding to the contract
	CurveRenFilterer   // Log filterer for contract events
}

// CurveRenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveRenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveRenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveRenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveRenSession struct {
	Contract     *CurveRen         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveRenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveRenCallerSession struct {
	Contract *CurveRenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CurveRenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveRenTransactorSession struct {
	Contract     *CurveRenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CurveRenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveRenRaw struct {
	Contract *CurveRen // Generic contract binding to access the raw methods on
}

// CurveRenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveRenCallerRaw struct {
	Contract *CurveRenCaller // Generic read-only contract binding to access the raw methods on
}

// CurveRenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveRenTransactorRaw struct {
	Contract *CurveRenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveRen creates a new instance of CurveRen, bound to a specific deployed contract.
func NewCurveRen(address common.Address, backend bind.ContractBackend) (*CurveRen, error) {
	contract, err := bindCurveRen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveRen{CurveRenCaller: CurveRenCaller{contract: contract}, CurveRenTransactor: CurveRenTransactor{contract: contract}, CurveRenFilterer: CurveRenFilterer{contract: contract}}, nil
}

// NewCurveRenCaller creates a new read-only instance of CurveRen, bound to a specific deployed contract.
func NewCurveRenCaller(address common.Address, caller bind.ContractCaller) (*CurveRenCaller, error) {
	contract, err := bindCurveRen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveRenCaller{contract: contract}, nil
}

// NewCurveRenTransactor creates a new write-only instance of CurveRen, bound to a specific deployed contract.
func NewCurveRenTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveRenTransactor, error) {
	contract, err := bindCurveRen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveRenTransactor{contract: contract}, nil
}

// NewCurveRenFilterer creates a new log filterer instance of CurveRen, bound to a specific deployed contract.
func NewCurveRenFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveRenFilterer, error) {
	contract, err := bindCurveRen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveRenFilterer{contract: contract}, nil
}

// bindCurveRen binds a generic wrapper to an already deployed contract.
func bindCurveRen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveRenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveRen *CurveRenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveRen.Contract.CurveRenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveRen *CurveRenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRen.Contract.CurveRenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveRen *CurveRenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveRen.Contract.CurveRenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveRen *CurveRenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveRen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveRen *CurveRenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveRen *CurveRenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveRen.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveRen *CurveRenCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveRen *CurveRenSession) A() (*big.Int, error) {
	return _CurveRen.Contract.A(&_CurveRen.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) A() (*big.Int, error) {
	return _CurveRen.Contract.A(&_CurveRen.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveRen *CurveRenCaller) APrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "A_precise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveRen *CurveRenSession) APrecise() (*big.Int, error) {
	return _CurveRen.Contract.APrecise(&_CurveRen.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) APrecise() (*big.Int, error) {
	return _CurveRen.Contract.APrecise(&_CurveRen.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveRen *CurveRenCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveRen *CurveRenSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurveRen.Contract.AdminActionsDeadline(&_CurveRen.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurveRen.Contract.AdminActionsDeadline(&_CurveRen.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveRen *CurveRenCaller) AdminBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "admin_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveRen *CurveRenSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.AdminBalances(&_CurveRen.CallOpts, arg0)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveRen *CurveRenCallerSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.AdminBalances(&_CurveRen.CallOpts, arg0)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveRen *CurveRenCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveRen *CurveRenSession) AdminFee() (*big.Int, error) {
	return _CurveRen.Contract.AdminFee(&_CurveRen.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) AdminFee() (*big.Int, error) {
	return _CurveRen.Contract.AdminFee(&_CurveRen.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveRen *CurveRenCaller) AdminFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "admin_fee_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveRen *CurveRenSession) AdminFeeReceiver() (common.Address, error) {
	return _CurveRen.Contract.AdminFeeReceiver(&_CurveRen.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveRen *CurveRenCallerSession) AdminFeeReceiver() (common.Address, error) {
	return _CurveRen.Contract.AdminFeeReceiver(&_CurveRen.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveRen *CurveRenCaller) Balances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveRen *CurveRenSession) Balances(i *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.Balances(&_CurveRen.CallOpts, i)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveRen *CurveRenCallerSession) Balances(i *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.Balances(&_CurveRen.CallOpts, i)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool is_deposit) view returns(uint256)
func (_CurveRen *CurveRenCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts [2]*big.Int, is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "calc_token_amount", _amounts, is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool is_deposit) view returns(uint256)
func (_CurveRen *CurveRenSession) CalcTokenAmount(_amounts [2]*big.Int, is_deposit bool) (*big.Int, error) {
	return _CurveRen.Contract.CalcTokenAmount(&_CurveRen.CallOpts, _amounts, is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool is_deposit) view returns(uint256)
func (_CurveRen *CurveRenCallerSession) CalcTokenAmount(_amounts [2]*big.Int, is_deposit bool) (*big.Int, error) {
	return _CurveRen.Contract.CalcTokenAmount(&_CurveRen.CallOpts, _amounts, is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurveRen *CurveRenCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "calc_withdraw_one_coin", _token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurveRen *CurveRenSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.CalcWithdrawOneCoin(&_CurveRen.CallOpts, _token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurveRen *CurveRenCallerSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.CalcWithdrawOneCoin(&_CurveRen.CallOpts, _token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveRen *CurveRenCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveRen *CurveRenSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurveRen.Contract.Coins(&_CurveRen.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveRen *CurveRenCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurveRen.Contract.Coins(&_CurveRen.CallOpts, arg0)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_CurveRen *CurveRenCaller) DynamicFee(opts *bind.CallOpts, i *big.Int, j *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "dynamic_fee", i, j)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_CurveRen *CurveRenSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.DynamicFee(&_CurveRen.CallOpts, i, j)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_CurveRen *CurveRenCallerSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.DynamicFee(&_CurveRen.CallOpts, i, j)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveRen *CurveRenCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveRen *CurveRenSession) Fee() (*big.Int, error) {
	return _CurveRen.Contract.Fee(&_CurveRen.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) Fee() (*big.Int, error) {
	return _CurveRen.Contract.Fee(&_CurveRen.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveRen *CurveRenCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveRen *CurveRenSession) FutureA() (*big.Int, error) {
	return _CurveRen.Contract.FutureA(&_CurveRen.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) FutureA() (*big.Int, error) {
	return _CurveRen.Contract.FutureA(&_CurveRen.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveRen *CurveRenCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveRen *CurveRenSession) FutureATime() (*big.Int, error) {
	return _CurveRen.Contract.FutureATime(&_CurveRen.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) FutureATime() (*big.Int, error) {
	return _CurveRen.Contract.FutureATime(&_CurveRen.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveRen *CurveRenCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveRen *CurveRenSession) FutureAdminFee() (*big.Int, error) {
	return _CurveRen.Contract.FutureAdminFee(&_CurveRen.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) FutureAdminFee() (*big.Int, error) {
	return _CurveRen.Contract.FutureAdminFee(&_CurveRen.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurveRen *CurveRenCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "future_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurveRen *CurveRenSession) FutureFee() (*big.Int, error) {
	return _CurveRen.Contract.FutureFee(&_CurveRen.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) FutureFee() (*big.Int, error) {
	return _CurveRen.Contract.FutureFee(&_CurveRen.CallOpts)
}

// FutureOffpegFeeMultiplier is a free data retrieval call binding the contract method 0x1e4c4ef8.
//
// Solidity: function future_offpeg_fee_multiplier() view returns(uint256)
func (_CurveRen *CurveRenCaller) FutureOffpegFeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "future_offpeg_fee_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureOffpegFeeMultiplier is a free data retrieval call binding the contract method 0x1e4c4ef8.
//
// Solidity: function future_offpeg_fee_multiplier() view returns(uint256)
func (_CurveRen *CurveRenSession) FutureOffpegFeeMultiplier() (*big.Int, error) {
	return _CurveRen.Contract.FutureOffpegFeeMultiplier(&_CurveRen.CallOpts)
}

// FutureOffpegFeeMultiplier is a free data retrieval call binding the contract method 0x1e4c4ef8.
//
// Solidity: function future_offpeg_fee_multiplier() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) FutureOffpegFeeMultiplier() (*big.Int, error) {
	return _CurveRen.Contract.FutureOffpegFeeMultiplier(&_CurveRen.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveRen *CurveRenCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveRen *CurveRenSession) FutureOwner() (common.Address, error) {
	return _CurveRen.Contract.FutureOwner(&_CurveRen.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveRen *CurveRenCallerSession) FutureOwner() (common.Address, error) {
	return _CurveRen.Contract.FutureOwner(&_CurveRen.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveRen *CurveRenCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveRen *CurveRenSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.GetDy(&_CurveRen.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveRen *CurveRenCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.GetDy(&_CurveRen.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveRen *CurveRenCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveRen *CurveRenSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.GetDyUnderlying(&_CurveRen.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveRen *CurveRenCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveRen.Contract.GetDyUnderlying(&_CurveRen.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveRen *CurveRenCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveRen *CurveRenSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveRen.Contract.GetVirtualPrice(&_CurveRen.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveRen.Contract.GetVirtualPrice(&_CurveRen.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveRen *CurveRenCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveRen *CurveRenSession) InitialA() (*big.Int, error) {
	return _CurveRen.Contract.InitialA(&_CurveRen.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) InitialA() (*big.Int, error) {
	return _CurveRen.Contract.InitialA(&_CurveRen.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveRen *CurveRenCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveRen *CurveRenSession) InitialATime() (*big.Int, error) {
	return _CurveRen.Contract.InitialATime(&_CurveRen.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) InitialATime() (*big.Int, error) {
	return _CurveRen.Contract.InitialATime(&_CurveRen.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveRen *CurveRenCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "lp_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveRen *CurveRenSession) LpToken() (common.Address, error) {
	return _CurveRen.Contract.LpToken(&_CurveRen.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveRen *CurveRenCallerSession) LpToken() (common.Address, error) {
	return _CurveRen.Contract.LpToken(&_CurveRen.CallOpts)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_CurveRen *CurveRenCaller) OffpegFeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "offpeg_fee_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_CurveRen *CurveRenSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _CurveRen.Contract.OffpegFeeMultiplier(&_CurveRen.CallOpts)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _CurveRen.Contract.OffpegFeeMultiplier(&_CurveRen.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveRen *CurveRenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveRen *CurveRenSession) Owner() (common.Address, error) {
	return _CurveRen.Contract.Owner(&_CurveRen.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveRen *CurveRenCallerSession) Owner() (common.Address, error) {
	return _CurveRen.Contract.Owner(&_CurveRen.CallOpts)
}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveRen *CurveRenCaller) RewardReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "reward_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveRen *CurveRenSession) RewardReceiver() (common.Address, error) {
	return _CurveRen.Contract.RewardReceiver(&_CurveRen.CallOpts)
}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveRen *CurveRenCallerSession) RewardReceiver() (common.Address, error) {
	return _CurveRen.Contract.RewardReceiver(&_CurveRen.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveRen *CurveRenCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveRen *CurveRenSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurveRen.Contract.TransferOwnershipDeadline(&_CurveRen.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveRen *CurveRenCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurveRen.Contract.TransferOwnershipDeadline(&_CurveRen.CallOpts)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 arg0) view returns(address)
func (_CurveRen *CurveRenCaller) UnderlyingCoins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveRen.contract.Call(opts, &out, "underlying_coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 arg0) view returns(address)
func (_CurveRen *CurveRenSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _CurveRen.Contract.UnderlyingCoins(&_CurveRen.CallOpts, arg0)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 arg0) view returns(address)
func (_CurveRen *CurveRenCallerSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _CurveRen.Contract.UnderlyingCoins(&_CurveRen.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurveRen *CurveRenTransactor) AddLiquidity(opts *bind.TransactOpts, _amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "add_liquidity", _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurveRen *CurveRenSession) AddLiquidity(_amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.AddLiquidity(&_CurveRen.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurveRen *CurveRenTransactorSession) AddLiquidity(_amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.AddLiquidity(&_CurveRen.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, bool _use_underlying) returns(uint256)
func (_CurveRen *CurveRenTransactor) AddLiquidity0(opts *bind.TransactOpts, _amounts [2]*big.Int, _min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "add_liquidity0", _amounts, _min_mint_amount, _use_underlying)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, bool _use_underlying) returns(uint256)
func (_CurveRen *CurveRenSession) AddLiquidity0(_amounts [2]*big.Int, _min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.Contract.AddLiquidity0(&_CurveRen.TransactOpts, _amounts, _min_mint_amount, _use_underlying)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, bool _use_underlying) returns(uint256)
func (_CurveRen *CurveRenTransactorSession) AddLiquidity0(_amounts [2]*big.Int, _min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.Contract.AddLiquidity0(&_CurveRen.TransactOpts, _amounts, _min_mint_amount, _use_underlying)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurveRen *CurveRenTransactor) ApplyNewFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "apply_new_fee")
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurveRen *CurveRenSession) ApplyNewFee() (*types.Transaction, error) {
	return _CurveRen.Contract.ApplyNewFee(&_CurveRen.TransactOpts)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurveRen *CurveRenTransactorSession) ApplyNewFee() (*types.Transaction, error) {
	return _CurveRen.Contract.ApplyNewFee(&_CurveRen.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveRen *CurveRenTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveRen *CurveRenSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveRen.Contract.ApplyTransferOwnership(&_CurveRen.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveRen *CurveRenTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveRen.Contract.ApplyTransferOwnership(&_CurveRen.TransactOpts)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x0746dd5a.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee, uint256 new_offpeg_fee_multiplier) returns()
func (_CurveRen *CurveRenTransactor) CommitNewFee(opts *bind.TransactOpts, new_fee *big.Int, new_admin_fee *big.Int, new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "commit_new_fee", new_fee, new_admin_fee, new_offpeg_fee_multiplier)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x0746dd5a.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee, uint256 new_offpeg_fee_multiplier) returns()
func (_CurveRen *CurveRenSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int, new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.CommitNewFee(&_CurveRen.TransactOpts, new_fee, new_admin_fee, new_offpeg_fee_multiplier)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x0746dd5a.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee, uint256 new_offpeg_fee_multiplier) returns()
func (_CurveRen *CurveRenTransactorSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int, new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.CommitNewFee(&_CurveRen.TransactOpts, new_fee, new_admin_fee, new_offpeg_fee_multiplier)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveRen *CurveRenTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveRen *CurveRenSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveRen.Contract.CommitTransferOwnership(&_CurveRen.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveRen *CurveRenTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveRen.Contract.CommitTransferOwnership(&_CurveRen.TransactOpts, _owner)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurveRen *CurveRenTransactor) DonateAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "donate_admin_fees")
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurveRen *CurveRenSession) DonateAdminFees() (*types.Transaction, error) {
	return _CurveRen.Contract.DonateAdminFees(&_CurveRen.TransactOpts)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurveRen *CurveRenTransactorSession) DonateAdminFees() (*types.Transaction, error) {
	return _CurveRen.Contract.DonateAdminFees(&_CurveRen.TransactOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveRen *CurveRenTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveRen *CurveRenSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.Exchange(&_CurveRen.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveRen *CurveRenTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.Exchange(&_CurveRen.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveRen *CurveRenTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveRen *CurveRenSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.ExchangeUnderlying(&_CurveRen.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveRen *CurveRenTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.ExchangeUnderlying(&_CurveRen.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveRen *CurveRenTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveRen *CurveRenSession) KillMe() (*types.Transaction, error) {
	return _CurveRen.Contract.KillMe(&_CurveRen.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveRen *CurveRenTransactorSession) KillMe() (*types.Transaction, error) {
	return _CurveRen.Contract.KillMe(&_CurveRen.TransactOpts)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveRen *CurveRenTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveRen *CurveRenSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.RampA(&_CurveRen.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveRen *CurveRenTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.RampA(&_CurveRen.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] _min_amounts, bool _use_underlying) returns(uint256[2])
func (_CurveRen *CurveRenTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, _min_amounts [2]*big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "remove_liquidity", _amount, _min_amounts, _use_underlying)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] _min_amounts, bool _use_underlying) returns(uint256[2])
func (_CurveRen *CurveRenSession) RemoveLiquidity(_amount *big.Int, _min_amounts [2]*big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidity(&_CurveRen.TransactOpts, _amount, _min_amounts, _use_underlying)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] _min_amounts, bool _use_underlying) returns(uint256[2])
func (_CurveRen *CurveRenTransactorSession) RemoveLiquidity(_amount *big.Int, _min_amounts [2]*big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidity(&_CurveRen.TransactOpts, _amount, _min_amounts, _use_underlying)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurveRen *CurveRenTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "remove_liquidity0", _amount, _min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurveRen *CurveRenSession) RemoveLiquidity0(_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidity0(&_CurveRen.TransactOpts, _amount, _min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurveRen *CurveRenTransactorSession) RemoveLiquidity0(_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidity0(&_CurveRen.TransactOpts, _amount, _min_amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe1df4f63.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, bool _use_underlying) returns(uint256)
func (_CurveRen *CurveRenTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, _amounts [2]*big.Int, _max_burn_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "remove_liquidity_imbalance", _amounts, _max_burn_amount, _use_underlying)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe1df4f63.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, bool _use_underlying) returns(uint256)
func (_CurveRen *CurveRenSession) RemoveLiquidityImbalance(_amounts [2]*big.Int, _max_burn_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidityImbalance(&_CurveRen.TransactOpts, _amounts, _max_burn_amount, _use_underlying)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe1df4f63.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, bool _use_underlying) returns(uint256)
func (_CurveRen *CurveRenTransactorSession) RemoveLiquidityImbalance(_amounts [2]*big.Int, _max_burn_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidityImbalance(&_CurveRen.TransactOpts, _amounts, _max_burn_amount, _use_underlying)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveRen *CurveRenTransactor) RemoveLiquidityImbalance0(opts *bind.TransactOpts, _amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "remove_liquidity_imbalance0", _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveRen *CurveRenSession) RemoveLiquidityImbalance0(_amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidityImbalance0(&_CurveRen.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveRen *CurveRenTransactorSession) RemoveLiquidityImbalance0(_amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidityImbalance0(&_CurveRen.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount) returns(uint256)
func (_CurveRen *CurveRenTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "remove_liquidity_one_coin", _token_amount, i, _min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount) returns(uint256)
func (_CurveRen *CurveRenSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidityOneCoin(&_CurveRen.TransactOpts, _token_amount, i, _min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount) returns(uint256)
func (_CurveRen *CurveRenTransactorSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidityOneCoin(&_CurveRen.TransactOpts, _token_amount, i, _min_amount)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x517a55a3.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount, bool _use_underlying) returns(uint256)
func (_CurveRen *CurveRenTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, _min_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "remove_liquidity_one_coin0", _token_amount, i, _min_amount, _use_underlying)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x517a55a3.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount, bool _use_underlying) returns(uint256)
func (_CurveRen *CurveRenSession) RemoveLiquidityOneCoin0(_token_amount *big.Int, i *big.Int, _min_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidityOneCoin0(&_CurveRen.TransactOpts, _token_amount, i, _min_amount, _use_underlying)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x517a55a3.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount, bool _use_underlying) returns(uint256)
func (_CurveRen *CurveRenTransactorSession) RemoveLiquidityOneCoin0(_token_amount *big.Int, i *big.Int, _min_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveRen.Contract.RemoveLiquidityOneCoin0(&_CurveRen.TransactOpts, _token_amount, i, _min_amount, _use_underlying)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveRen *CurveRenTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveRen *CurveRenSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurveRen.Contract.RevertNewParameters(&_CurveRen.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveRen *CurveRenTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurveRen.Contract.RevertNewParameters(&_CurveRen.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveRen *CurveRenTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveRen *CurveRenSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurveRen.Contract.RevertTransferOwnership(&_CurveRen.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveRen *CurveRenTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurveRen.Contract.RevertTransferOwnership(&_CurveRen.TransactOpts)
}

// SetAaveReferral is a paid mutator transaction binding the contract method 0xb6aa64c5.
//
// Solidity: function set_aave_referral(uint256 referral_code) returns()
func (_CurveRen *CurveRenTransactor) SetAaveReferral(opts *bind.TransactOpts, referral_code *big.Int) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "set_aave_referral", referral_code)
}

// SetAaveReferral is a paid mutator transaction binding the contract method 0xb6aa64c5.
//
// Solidity: function set_aave_referral(uint256 referral_code) returns()
func (_CurveRen *CurveRenSession) SetAaveReferral(referral_code *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.SetAaveReferral(&_CurveRen.TransactOpts, referral_code)
}

// SetAaveReferral is a paid mutator transaction binding the contract method 0xb6aa64c5.
//
// Solidity: function set_aave_referral(uint256 referral_code) returns()
func (_CurveRen *CurveRenTransactorSession) SetAaveReferral(referral_code *big.Int) (*types.Transaction, error) {
	return _CurveRen.Contract.SetAaveReferral(&_CurveRen.TransactOpts, referral_code)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveRen *CurveRenTransactor) SetAdminFeeReceiver(opts *bind.TransactOpts, _admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "set_admin_fee_receiver", _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveRen *CurveRenSession) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveRen.Contract.SetAdminFeeReceiver(&_CurveRen.TransactOpts, _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveRen *CurveRenTransactorSession) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveRen.Contract.SetAdminFeeReceiver(&_CurveRen.TransactOpts, _admin_fee_receiver)
}

// SetRewardReceiver is a paid mutator transaction binding the contract method 0xc51b8861.
//
// Solidity: function set_reward_receiver(address _reward_receiver) returns()
func (_CurveRen *CurveRenTransactor) SetRewardReceiver(opts *bind.TransactOpts, _reward_receiver common.Address) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "set_reward_receiver", _reward_receiver)
}

// SetRewardReceiver is a paid mutator transaction binding the contract method 0xc51b8861.
//
// Solidity: function set_reward_receiver(address _reward_receiver) returns()
func (_CurveRen *CurveRenSession) SetRewardReceiver(_reward_receiver common.Address) (*types.Transaction, error) {
	return _CurveRen.Contract.SetRewardReceiver(&_CurveRen.TransactOpts, _reward_receiver)
}

// SetRewardReceiver is a paid mutator transaction binding the contract method 0xc51b8861.
//
// Solidity: function set_reward_receiver(address _reward_receiver) returns()
func (_CurveRen *CurveRenTransactorSession) SetRewardReceiver(_reward_receiver common.Address) (*types.Transaction, error) {
	return _CurveRen.Contract.SetRewardReceiver(&_CurveRen.TransactOpts, _reward_receiver)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveRen *CurveRenTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveRen *CurveRenSession) StopRampA() (*types.Transaction, error) {
	return _CurveRen.Contract.StopRampA(&_CurveRen.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveRen *CurveRenTransactorSession) StopRampA() (*types.Transaction, error) {
	return _CurveRen.Contract.StopRampA(&_CurveRen.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveRen *CurveRenTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveRen *CurveRenSession) UnkillMe() (*types.Transaction, error) {
	return _CurveRen.Contract.UnkillMe(&_CurveRen.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveRen *CurveRenTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _CurveRen.Contract.UnkillMe(&_CurveRen.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveRen *CurveRenTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRen.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveRen *CurveRenSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurveRen.Contract.WithdrawAdminFees(&_CurveRen.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveRen *CurveRenTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurveRen.Contract.WithdrawAdminFees(&_CurveRen.TransactOpts)
}
