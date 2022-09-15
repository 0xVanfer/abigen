// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveAave

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

// CurveAaveMetaData contains all meta data concerning the CurveAave contract.
var CurveAaveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A_precise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"_min_mint_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_use_underlying\",\"type\":\"bool\"}],\"name\":\"add_liquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_actions_deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"admin_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee_receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apply_new_fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"bool\",\"name\":\"is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"new_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_admin_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_offpeg_fee_multiplier\",\"type\":\"uint256\"}],\"name\":\"commit_new_fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate_admin_fees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"}],\"name\":\"dynamic_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_offpeg_fee_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill_me\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offpeg_fee_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_future_A\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_future_time\",\"type\":\"uint256\"}],\"name\":\"ramp_A\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"_min_amounts\",\"type\":\"uint256[3]\"}],\"name\":\"remove_liquidity\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"_min_amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"bool\",\"name\":\"_use_underlying\",\"type\":\"bool\"}],\"name\":\"remove_liquidity\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"_max_burn_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_use_underlying\",\"type\":\"bool\"}],\"name\":\"remove_liquidity_imbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"_max_burn_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_imbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"_min_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"_min_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_use_underlying\",\"type\":\"bool\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revert_new_parameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward_receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"referral_code\",\"type\":\"uint256\"}],\"name\":\"set_aave_referral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin_fee_receiver\",\"type\":\"address\"}],\"name\":\"set_admin_fee_receiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reward_receiver\",\"type\":\"address\"}],\"name\":\"set_reward_receiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stop_ramp_A\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unkill_me\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CurveAaveABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveAaveMetaData.ABI instead.
var CurveAaveABI = CurveAaveMetaData.ABI

// CurveAave is an auto generated Go binding around an Ethereum contract.
type CurveAave struct {
	CurveAaveCaller     // Read-only binding to the contract
	CurveAaveTransactor // Write-only binding to the contract
	CurveAaveFilterer   // Log filterer for contract events
}

// CurveAaveCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveAaveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveAaveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveAaveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveAaveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveAaveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveAaveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveAaveSession struct {
	Contract     *CurveAave        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveAaveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveAaveCallerSession struct {
	Contract *CurveAaveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CurveAaveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveAaveTransactorSession struct {
	Contract     *CurveAaveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CurveAaveRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveAaveRaw struct {
	Contract *CurveAave // Generic contract binding to access the raw methods on
}

// CurveAaveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveAaveCallerRaw struct {
	Contract *CurveAaveCaller // Generic read-only contract binding to access the raw methods on
}

// CurveAaveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveAaveTransactorRaw struct {
	Contract *CurveAaveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveAave creates a new instance of CurveAave, bound to a specific deployed contract.
func NewCurveAave(address common.Address, backend bind.ContractBackend) (*CurveAave, error) {
	contract, err := bindCurveAave(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveAave{CurveAaveCaller: CurveAaveCaller{contract: contract}, CurveAaveTransactor: CurveAaveTransactor{contract: contract}, CurveAaveFilterer: CurveAaveFilterer{contract: contract}}, nil
}

// NewCurveAaveCaller creates a new read-only instance of CurveAave, bound to a specific deployed contract.
func NewCurveAaveCaller(address common.Address, caller bind.ContractCaller) (*CurveAaveCaller, error) {
	contract, err := bindCurveAave(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveAaveCaller{contract: contract}, nil
}

// NewCurveAaveTransactor creates a new write-only instance of CurveAave, bound to a specific deployed contract.
func NewCurveAaveTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveAaveTransactor, error) {
	contract, err := bindCurveAave(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveAaveTransactor{contract: contract}, nil
}

// NewCurveAaveFilterer creates a new log filterer instance of CurveAave, bound to a specific deployed contract.
func NewCurveAaveFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveAaveFilterer, error) {
	contract, err := bindCurveAave(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveAaveFilterer{contract: contract}, nil
}

// bindCurveAave binds a generic wrapper to an already deployed contract.
func bindCurveAave(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveAaveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveAave *CurveAaveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveAave.Contract.CurveAaveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveAave *CurveAaveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAave.Contract.CurveAaveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveAave *CurveAaveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveAave.Contract.CurveAaveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveAave *CurveAaveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveAave.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveAave *CurveAaveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAave.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveAave *CurveAaveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveAave.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAave *CurveAaveCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAave *CurveAaveSession) A() (*big.Int, error) {
	return _CurveAave.Contract.A(&_CurveAave.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) A() (*big.Int, error) {
	return _CurveAave.Contract.A(&_CurveAave.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAave *CurveAaveCaller) APrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "A_precise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAave *CurveAaveSession) APrecise() (*big.Int, error) {
	return _CurveAave.Contract.APrecise(&_CurveAave.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) APrecise() (*big.Int, error) {
	return _CurveAave.Contract.APrecise(&_CurveAave.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveAave *CurveAaveCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveAave *CurveAaveSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurveAave.Contract.AdminActionsDeadline(&_CurveAave.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurveAave.Contract.AdminActionsDeadline(&_CurveAave.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveAave *CurveAaveCaller) AdminBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "admin_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveAave *CurveAaveSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.AdminBalances(&_CurveAave.CallOpts, arg0)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.AdminBalances(&_CurveAave.CallOpts, arg0)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAave *CurveAaveCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAave *CurveAaveSession) AdminFee() (*big.Int, error) {
	return _CurveAave.Contract.AdminFee(&_CurveAave.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) AdminFee() (*big.Int, error) {
	return _CurveAave.Contract.AdminFee(&_CurveAave.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveAave *CurveAaveCaller) AdminFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "admin_fee_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveAave *CurveAaveSession) AdminFeeReceiver() (common.Address, error) {
	return _CurveAave.Contract.AdminFeeReceiver(&_CurveAave.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveAave *CurveAaveCallerSession) AdminFeeReceiver() (common.Address, error) {
	return _CurveAave.Contract.AdminFeeReceiver(&_CurveAave.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAave *CurveAaveCaller) Balances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAave *CurveAaveSession) Balances(i *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.Balances(&_CurveAave.CallOpts, i)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) Balances(i *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.Balances(&_CurveAave.CallOpts, i)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool is_deposit) view returns(uint256)
func (_CurveAave *CurveAaveCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts [3]*big.Int, is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "calc_token_amount", _amounts, is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool is_deposit) view returns(uint256)
func (_CurveAave *CurveAaveSession) CalcTokenAmount(_amounts [3]*big.Int, is_deposit bool) (*big.Int, error) {
	return _CurveAave.Contract.CalcTokenAmount(&_CurveAave.CallOpts, _amounts, is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool is_deposit) view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) CalcTokenAmount(_amounts [3]*big.Int, is_deposit bool) (*big.Int, error) {
	return _CurveAave.Contract.CalcTokenAmount(&_CurveAave.CallOpts, _amounts, is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurveAave *CurveAaveCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "calc_withdraw_one_coin", _token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurveAave *CurveAaveSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.CalcWithdrawOneCoin(&_CurveAave.CallOpts, _token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.CalcWithdrawOneCoin(&_CurveAave.CallOpts, _token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveAave *CurveAaveCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveAave *CurveAaveSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurveAave.Contract.Coins(&_CurveAave.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurveAave *CurveAaveCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurveAave.Contract.Coins(&_CurveAave.CallOpts, arg0)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_CurveAave *CurveAaveCaller) DynamicFee(opts *bind.CallOpts, i *big.Int, j *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "dynamic_fee", i, j)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_CurveAave *CurveAaveSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.DynamicFee(&_CurveAave.CallOpts, i, j)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.DynamicFee(&_CurveAave.CallOpts, i, j)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAave *CurveAaveCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAave *CurveAaveSession) Fee() (*big.Int, error) {
	return _CurveAave.Contract.Fee(&_CurveAave.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) Fee() (*big.Int, error) {
	return _CurveAave.Contract.Fee(&_CurveAave.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAave *CurveAaveCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAave *CurveAaveSession) FutureA() (*big.Int, error) {
	return _CurveAave.Contract.FutureA(&_CurveAave.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) FutureA() (*big.Int, error) {
	return _CurveAave.Contract.FutureA(&_CurveAave.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAave *CurveAaveCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAave *CurveAaveSession) FutureATime() (*big.Int, error) {
	return _CurveAave.Contract.FutureATime(&_CurveAave.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) FutureATime() (*big.Int, error) {
	return _CurveAave.Contract.FutureATime(&_CurveAave.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveAave *CurveAaveCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveAave *CurveAaveSession) FutureAdminFee() (*big.Int, error) {
	return _CurveAave.Contract.FutureAdminFee(&_CurveAave.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) FutureAdminFee() (*big.Int, error) {
	return _CurveAave.Contract.FutureAdminFee(&_CurveAave.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurveAave *CurveAaveCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "future_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurveAave *CurveAaveSession) FutureFee() (*big.Int, error) {
	return _CurveAave.Contract.FutureFee(&_CurveAave.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) FutureFee() (*big.Int, error) {
	return _CurveAave.Contract.FutureFee(&_CurveAave.CallOpts)
}

// FutureOffpegFeeMultiplier is a free data retrieval call binding the contract method 0x1e4c4ef8.
//
// Solidity: function future_offpeg_fee_multiplier() view returns(uint256)
func (_CurveAave *CurveAaveCaller) FutureOffpegFeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "future_offpeg_fee_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureOffpegFeeMultiplier is a free data retrieval call binding the contract method 0x1e4c4ef8.
//
// Solidity: function future_offpeg_fee_multiplier() view returns(uint256)
func (_CurveAave *CurveAaveSession) FutureOffpegFeeMultiplier() (*big.Int, error) {
	return _CurveAave.Contract.FutureOffpegFeeMultiplier(&_CurveAave.CallOpts)
}

// FutureOffpegFeeMultiplier is a free data retrieval call binding the contract method 0x1e4c4ef8.
//
// Solidity: function future_offpeg_fee_multiplier() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) FutureOffpegFeeMultiplier() (*big.Int, error) {
	return _CurveAave.Contract.FutureOffpegFeeMultiplier(&_CurveAave.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveAave *CurveAaveCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveAave *CurveAaveSession) FutureOwner() (common.Address, error) {
	return _CurveAave.Contract.FutureOwner(&_CurveAave.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveAave *CurveAaveCallerSession) FutureOwner() (common.Address, error) {
	return _CurveAave.Contract.FutureOwner(&_CurveAave.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAave *CurveAaveCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAave *CurveAaveSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.GetDy(&_CurveAave.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.GetDy(&_CurveAave.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAave *CurveAaveCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAave *CurveAaveSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.GetDyUnderlying(&_CurveAave.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAave.Contract.GetDyUnderlying(&_CurveAave.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAave *CurveAaveCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAave *CurveAaveSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveAave.Contract.GetVirtualPrice(&_CurveAave.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveAave.Contract.GetVirtualPrice(&_CurveAave.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAave *CurveAaveCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAave *CurveAaveSession) InitialA() (*big.Int, error) {
	return _CurveAave.Contract.InitialA(&_CurveAave.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) InitialA() (*big.Int, error) {
	return _CurveAave.Contract.InitialA(&_CurveAave.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAave *CurveAaveCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAave *CurveAaveSession) InitialATime() (*big.Int, error) {
	return _CurveAave.Contract.InitialATime(&_CurveAave.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) InitialATime() (*big.Int, error) {
	return _CurveAave.Contract.InitialATime(&_CurveAave.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAave *CurveAaveCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "lp_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAave *CurveAaveSession) LpToken() (common.Address, error) {
	return _CurveAave.Contract.LpToken(&_CurveAave.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAave *CurveAaveCallerSession) LpToken() (common.Address, error) {
	return _CurveAave.Contract.LpToken(&_CurveAave.CallOpts)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_CurveAave *CurveAaveCaller) OffpegFeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "offpeg_fee_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_CurveAave *CurveAaveSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _CurveAave.Contract.OffpegFeeMultiplier(&_CurveAave.CallOpts)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _CurveAave.Contract.OffpegFeeMultiplier(&_CurveAave.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveAave *CurveAaveCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveAave *CurveAaveSession) Owner() (common.Address, error) {
	return _CurveAave.Contract.Owner(&_CurveAave.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveAave *CurveAaveCallerSession) Owner() (common.Address, error) {
	return _CurveAave.Contract.Owner(&_CurveAave.CallOpts)
}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveAave *CurveAaveCaller) RewardReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "reward_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveAave *CurveAaveSession) RewardReceiver() (common.Address, error) {
	return _CurveAave.Contract.RewardReceiver(&_CurveAave.CallOpts)
}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveAave *CurveAaveCallerSession) RewardReceiver() (common.Address, error) {
	return _CurveAave.Contract.RewardReceiver(&_CurveAave.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveAave *CurveAaveCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveAave *CurveAaveSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurveAave.Contract.TransferOwnershipDeadline(&_CurveAave.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveAave *CurveAaveCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurveAave.Contract.TransferOwnershipDeadline(&_CurveAave.CallOpts)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 arg0) view returns(address)
func (_CurveAave *CurveAaveCaller) UnderlyingCoins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAave.contract.Call(opts, &out, "underlying_coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 arg0) view returns(address)
func (_CurveAave *CurveAaveSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _CurveAave.Contract.UnderlyingCoins(&_CurveAave.CallOpts, arg0)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 arg0) view returns(address)
func (_CurveAave *CurveAaveCallerSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _CurveAave.Contract.UnderlyingCoins(&_CurveAave.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x2b6e993a.
//
// Solidity: function add_liquidity(uint256[3] _amounts, uint256 _min_mint_amount, bool _use_underlying) returns(uint256)
func (_CurveAave *CurveAaveTransactor) AddLiquidity(opts *bind.TransactOpts, _amounts [3]*big.Int, _min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "add_liquidity", _amounts, _min_mint_amount, _use_underlying)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x2b6e993a.
//
// Solidity: function add_liquidity(uint256[3] _amounts, uint256 _min_mint_amount, bool _use_underlying) returns(uint256)
func (_CurveAave *CurveAaveSession) AddLiquidity(_amounts [3]*big.Int, _min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.Contract.AddLiquidity(&_CurveAave.TransactOpts, _amounts, _min_mint_amount, _use_underlying)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x2b6e993a.
//
// Solidity: function add_liquidity(uint256[3] _amounts, uint256 _min_mint_amount, bool _use_underlying) returns(uint256)
func (_CurveAave *CurveAaveTransactorSession) AddLiquidity(_amounts [3]*big.Int, _min_mint_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.Contract.AddLiquidity(&_CurveAave.TransactOpts, _amounts, _min_mint_amount, _use_underlying)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurveAave *CurveAaveTransactor) AddLiquidity0(opts *bind.TransactOpts, _amounts [3]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "add_liquidity0", _amounts, _min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurveAave *CurveAaveSession) AddLiquidity0(_amounts [3]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.AddLiquidity0(&_CurveAave.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurveAave *CurveAaveTransactorSession) AddLiquidity0(_amounts [3]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.AddLiquidity0(&_CurveAave.TransactOpts, _amounts, _min_mint_amount)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurveAave *CurveAaveTransactor) ApplyNewFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "apply_new_fee")
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurveAave *CurveAaveSession) ApplyNewFee() (*types.Transaction, error) {
	return _CurveAave.Contract.ApplyNewFee(&_CurveAave.TransactOpts)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurveAave *CurveAaveTransactorSession) ApplyNewFee() (*types.Transaction, error) {
	return _CurveAave.Contract.ApplyNewFee(&_CurveAave.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveAave *CurveAaveTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveAave *CurveAaveSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveAave.Contract.ApplyTransferOwnership(&_CurveAave.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveAave *CurveAaveTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveAave.Contract.ApplyTransferOwnership(&_CurveAave.TransactOpts)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x0746dd5a.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee, uint256 new_offpeg_fee_multiplier) returns()
func (_CurveAave *CurveAaveTransactor) CommitNewFee(opts *bind.TransactOpts, new_fee *big.Int, new_admin_fee *big.Int, new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "commit_new_fee", new_fee, new_admin_fee, new_offpeg_fee_multiplier)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x0746dd5a.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee, uint256 new_offpeg_fee_multiplier) returns()
func (_CurveAave *CurveAaveSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int, new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.CommitNewFee(&_CurveAave.TransactOpts, new_fee, new_admin_fee, new_offpeg_fee_multiplier)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x0746dd5a.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee, uint256 new_offpeg_fee_multiplier) returns()
func (_CurveAave *CurveAaveTransactorSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int, new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.CommitNewFee(&_CurveAave.TransactOpts, new_fee, new_admin_fee, new_offpeg_fee_multiplier)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveAave *CurveAaveTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveAave *CurveAaveSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveAave.Contract.CommitTransferOwnership(&_CurveAave.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveAave *CurveAaveTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveAave.Contract.CommitTransferOwnership(&_CurveAave.TransactOpts, _owner)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurveAave *CurveAaveTransactor) DonateAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "donate_admin_fees")
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurveAave *CurveAaveSession) DonateAdminFees() (*types.Transaction, error) {
	return _CurveAave.Contract.DonateAdminFees(&_CurveAave.TransactOpts)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurveAave *CurveAaveTransactorSession) DonateAdminFees() (*types.Transaction, error) {
	return _CurveAave.Contract.DonateAdminFees(&_CurveAave.TransactOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveAave *CurveAaveTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveAave *CurveAaveSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.Exchange(&_CurveAave.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveAave *CurveAaveTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.Exchange(&_CurveAave.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveAave *CurveAaveTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveAave *CurveAaveSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.ExchangeUnderlying(&_CurveAave.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveAave *CurveAaveTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.ExchangeUnderlying(&_CurveAave.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveAave *CurveAaveTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveAave *CurveAaveSession) KillMe() (*types.Transaction, error) {
	return _CurveAave.Contract.KillMe(&_CurveAave.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveAave *CurveAaveTransactorSession) KillMe() (*types.Transaction, error) {
	return _CurveAave.Contract.KillMe(&_CurveAave.TransactOpts)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveAave *CurveAaveTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveAave *CurveAaveSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.RampA(&_CurveAave.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurveAave *CurveAaveTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.RampA(&_CurveAave.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] _min_amounts) returns(uint256[3])
func (_CurveAave *CurveAaveTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, _min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "remove_liquidity", _amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] _min_amounts) returns(uint256[3])
func (_CurveAave *CurveAaveSession) RemoveLiquidity(_amount *big.Int, _min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidity(&_CurveAave.TransactOpts, _amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] _min_amounts) returns(uint256[3])
func (_CurveAave *CurveAaveTransactorSession) RemoveLiquidity(_amount *big.Int, _min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidity(&_CurveAave.TransactOpts, _amount, _min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xfce64736.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] _min_amounts, bool _use_underlying) returns(uint256[3])
func (_CurveAave *CurveAaveTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _amount *big.Int, _min_amounts [3]*big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "remove_liquidity0", _amount, _min_amounts, _use_underlying)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xfce64736.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] _min_amounts, bool _use_underlying) returns(uint256[3])
func (_CurveAave *CurveAaveSession) RemoveLiquidity0(_amount *big.Int, _min_amounts [3]*big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidity0(&_CurveAave.TransactOpts, _amount, _min_amounts, _use_underlying)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xfce64736.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] _min_amounts, bool _use_underlying) returns(uint256[3])
func (_CurveAave *CurveAaveTransactorSession) RemoveLiquidity0(_amount *big.Int, _min_amounts [3]*big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidity0(&_CurveAave.TransactOpts, _amount, _min_amounts, _use_underlying)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x5b8369f5.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] _amounts, uint256 _max_burn_amount, bool _use_underlying) returns(uint256)
func (_CurveAave *CurveAaveTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, _amounts [3]*big.Int, _max_burn_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "remove_liquidity_imbalance", _amounts, _max_burn_amount, _use_underlying)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x5b8369f5.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] _amounts, uint256 _max_burn_amount, bool _use_underlying) returns(uint256)
func (_CurveAave *CurveAaveSession) RemoveLiquidityImbalance(_amounts [3]*big.Int, _max_burn_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidityImbalance(&_CurveAave.TransactOpts, _amounts, _max_burn_amount, _use_underlying)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x5b8369f5.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] _amounts, uint256 _max_burn_amount, bool _use_underlying) returns(uint256)
func (_CurveAave *CurveAaveTransactorSession) RemoveLiquidityImbalance(_amounts [3]*big.Int, _max_burn_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidityImbalance(&_CurveAave.TransactOpts, _amounts, _max_burn_amount, _use_underlying)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveAave *CurveAaveTransactor) RemoveLiquidityImbalance0(opts *bind.TransactOpts, _amounts [3]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "remove_liquidity_imbalance0", _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveAave *CurveAaveSession) RemoveLiquidityImbalance0(_amounts [3]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidityImbalance0(&_CurveAave.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurveAave *CurveAaveTransactorSession) RemoveLiquidityImbalance0(_amounts [3]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidityImbalance0(&_CurveAave.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount) returns(uint256)
func (_CurveAave *CurveAaveTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "remove_liquidity_one_coin", _token_amount, i, _min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount) returns(uint256)
func (_CurveAave *CurveAaveSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidityOneCoin(&_CurveAave.TransactOpts, _token_amount, i, _min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount) returns(uint256)
func (_CurveAave *CurveAaveTransactorSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidityOneCoin(&_CurveAave.TransactOpts, _token_amount, i, _min_amount)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x517a55a3.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount, bool _use_underlying) returns(uint256)
func (_CurveAave *CurveAaveTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, _min_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "remove_liquidity_one_coin0", _token_amount, i, _min_amount, _use_underlying)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x517a55a3.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount, bool _use_underlying) returns(uint256)
func (_CurveAave *CurveAaveSession) RemoveLiquidityOneCoin0(_token_amount *big.Int, i *big.Int, _min_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidityOneCoin0(&_CurveAave.TransactOpts, _token_amount, i, _min_amount, _use_underlying)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x517a55a3.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 _min_amount, bool _use_underlying) returns(uint256)
func (_CurveAave *CurveAaveTransactorSession) RemoveLiquidityOneCoin0(_token_amount *big.Int, i *big.Int, _min_amount *big.Int, _use_underlying bool) (*types.Transaction, error) {
	return _CurveAave.Contract.RemoveLiquidityOneCoin0(&_CurveAave.TransactOpts, _token_amount, i, _min_amount, _use_underlying)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveAave *CurveAaveTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveAave *CurveAaveSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurveAave.Contract.RevertNewParameters(&_CurveAave.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveAave *CurveAaveTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurveAave.Contract.RevertNewParameters(&_CurveAave.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveAave *CurveAaveTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveAave *CurveAaveSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurveAave.Contract.RevertTransferOwnership(&_CurveAave.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveAave *CurveAaveTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurveAave.Contract.RevertTransferOwnership(&_CurveAave.TransactOpts)
}

// SetAaveReferral is a paid mutator transaction binding the contract method 0xb6aa64c5.
//
// Solidity: function set_aave_referral(uint256 referral_code) returns()
func (_CurveAave *CurveAaveTransactor) SetAaveReferral(opts *bind.TransactOpts, referral_code *big.Int) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "set_aave_referral", referral_code)
}

// SetAaveReferral is a paid mutator transaction binding the contract method 0xb6aa64c5.
//
// Solidity: function set_aave_referral(uint256 referral_code) returns()
func (_CurveAave *CurveAaveSession) SetAaveReferral(referral_code *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.SetAaveReferral(&_CurveAave.TransactOpts, referral_code)
}

// SetAaveReferral is a paid mutator transaction binding the contract method 0xb6aa64c5.
//
// Solidity: function set_aave_referral(uint256 referral_code) returns()
func (_CurveAave *CurveAaveTransactorSession) SetAaveReferral(referral_code *big.Int) (*types.Transaction, error) {
	return _CurveAave.Contract.SetAaveReferral(&_CurveAave.TransactOpts, referral_code)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveAave *CurveAaveTransactor) SetAdminFeeReceiver(opts *bind.TransactOpts, _admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "set_admin_fee_receiver", _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveAave *CurveAaveSession) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveAave.Contract.SetAdminFeeReceiver(&_CurveAave.TransactOpts, _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveAave *CurveAaveTransactorSession) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveAave.Contract.SetAdminFeeReceiver(&_CurveAave.TransactOpts, _admin_fee_receiver)
}

// SetRewardReceiver is a paid mutator transaction binding the contract method 0xc51b8861.
//
// Solidity: function set_reward_receiver(address _reward_receiver) returns()
func (_CurveAave *CurveAaveTransactor) SetRewardReceiver(opts *bind.TransactOpts, _reward_receiver common.Address) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "set_reward_receiver", _reward_receiver)
}

// SetRewardReceiver is a paid mutator transaction binding the contract method 0xc51b8861.
//
// Solidity: function set_reward_receiver(address _reward_receiver) returns()
func (_CurveAave *CurveAaveSession) SetRewardReceiver(_reward_receiver common.Address) (*types.Transaction, error) {
	return _CurveAave.Contract.SetRewardReceiver(&_CurveAave.TransactOpts, _reward_receiver)
}

// SetRewardReceiver is a paid mutator transaction binding the contract method 0xc51b8861.
//
// Solidity: function set_reward_receiver(address _reward_receiver) returns()
func (_CurveAave *CurveAaveTransactorSession) SetRewardReceiver(_reward_receiver common.Address) (*types.Transaction, error) {
	return _CurveAave.Contract.SetRewardReceiver(&_CurveAave.TransactOpts, _reward_receiver)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveAave *CurveAaveTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveAave *CurveAaveSession) StopRampA() (*types.Transaction, error) {
	return _CurveAave.Contract.StopRampA(&_CurveAave.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurveAave *CurveAaveTransactorSession) StopRampA() (*types.Transaction, error) {
	return _CurveAave.Contract.StopRampA(&_CurveAave.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveAave *CurveAaveTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveAave *CurveAaveSession) UnkillMe() (*types.Transaction, error) {
	return _CurveAave.Contract.UnkillMe(&_CurveAave.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveAave *CurveAaveTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _CurveAave.Contract.UnkillMe(&_CurveAave.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveAave *CurveAaveTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAave.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveAave *CurveAaveSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurveAave.Contract.WithdrawAdminFees(&_CurveAave.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurveAave *CurveAaveTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurveAave.Contract.WithdrawAdminFees(&_CurveAave.TransactOpts)
}
