// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveAtri

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

// CurveAtriMetaData contains all meta data concerning the CurveAtri contract.
var CurveAtriMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"D\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adjustment_step\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_actions_deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee_receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowed_extra_profit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apply_new_parameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"bool\",\"name\":\"deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"xp\",\"type\":\"uint256[3]\"}],\"name\":\"calc_token_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim_admin_fees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_new_mid_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_new_out_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_new_admin_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_new_fee_gamma\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_new_allowed_extra_profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_new_adjustment_step\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_new_ma_half_time\",\"type\":\"uint256\"}],\"name\":\"commit_new_parameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"xp\",\"type\":\"uint256[3]\"}],\"name\":\"fee_calc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee_gamma\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_gamma\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_gamma_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_adjustment_step\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_allowed_extra_profit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_fee_gamma\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_ma_half_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_mid_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_out_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gamma\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_gamma\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_gamma_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"is_killed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill_deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill_me\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"last_prices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_prices_timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ma_half_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mid_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"out_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"price_oracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"price_scale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"future_A\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"future_gamma\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"future_time\",\"type\":\"uint256\"}],\"name\":\"ramp_A_gamma\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"min_amounts\",\"type\":\"uint256[3]\"}],\"name\":\"remove_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revert_new_parameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward_receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin_fee_receiver\",\"type\":\"address\"}],\"name\":\"set_admin_fee_receiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reward_receiver\",\"type\":\"address\"}],\"name\":\"set_reward_receiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stop_ramp_A_gamma\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unkill_me\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xcp_profit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xcp_profit_a\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CurveAtriABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveAtriMetaData.ABI instead.
var CurveAtriABI = CurveAtriMetaData.ABI

// CurveAtri is an auto generated Go binding around an Ethereum contract.
type CurveAtri struct {
	CurveAtriCaller     // Read-only binding to the contract
	CurveAtriTransactor // Write-only binding to the contract
	CurveAtriFilterer   // Log filterer for contract events
}

// CurveAtriCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveAtriCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveAtriTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveAtriTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveAtriFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveAtriFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveAtriSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveAtriSession struct {
	Contract     *CurveAtri        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveAtriCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveAtriCallerSession struct {
	Contract *CurveAtriCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CurveAtriTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveAtriTransactorSession struct {
	Contract     *CurveAtriTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CurveAtriRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveAtriRaw struct {
	Contract *CurveAtri // Generic contract binding to access the raw methods on
}

// CurveAtriCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveAtriCallerRaw struct {
	Contract *CurveAtriCaller // Generic read-only contract binding to access the raw methods on
}

// CurveAtriTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveAtriTransactorRaw struct {
	Contract *CurveAtriTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveAtri creates a new instance of CurveAtri, bound to a specific deployed contract.
func NewCurveAtri(address common.Address, backend bind.ContractBackend) (*CurveAtri, error) {
	contract, err := bindCurveAtri(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveAtri{CurveAtriCaller: CurveAtriCaller{contract: contract}, CurveAtriTransactor: CurveAtriTransactor{contract: contract}, CurveAtriFilterer: CurveAtriFilterer{contract: contract}}, nil
}

// NewCurveAtriCaller creates a new read-only instance of CurveAtri, bound to a specific deployed contract.
func NewCurveAtriCaller(address common.Address, caller bind.ContractCaller) (*CurveAtriCaller, error) {
	contract, err := bindCurveAtri(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveAtriCaller{contract: contract}, nil
}

// NewCurveAtriTransactor creates a new write-only instance of CurveAtri, bound to a specific deployed contract.
func NewCurveAtriTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveAtriTransactor, error) {
	contract, err := bindCurveAtri(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveAtriTransactor{contract: contract}, nil
}

// NewCurveAtriFilterer creates a new log filterer instance of CurveAtri, bound to a specific deployed contract.
func NewCurveAtriFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveAtriFilterer, error) {
	contract, err := bindCurveAtri(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveAtriFilterer{contract: contract}, nil
}

// bindCurveAtri binds a generic wrapper to an already deployed contract.
func bindCurveAtri(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveAtriABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveAtri *CurveAtriRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveAtri.Contract.CurveAtriCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveAtri *CurveAtriRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAtri.Contract.CurveAtriTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveAtri *CurveAtriRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveAtri.Contract.CurveAtriTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveAtri *CurveAtriCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveAtri.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveAtri *CurveAtriTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAtri.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveAtri *CurveAtriTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveAtri.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAtri *CurveAtriSession) A() (*big.Int, error) {
	return _CurveAtri.Contract.A(&_CurveAtri.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) A() (*big.Int, error) {
	return _CurveAtri.Contract.A(&_CurveAtri.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) D(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "D")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_CurveAtri *CurveAtriSession) D() (*big.Int, error) {
	return _CurveAtri.Contract.D(&_CurveAtri.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) D() (*big.Int, error) {
	return _CurveAtri.Contract.D(&_CurveAtri.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) AdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_CurveAtri *CurveAtriSession) AdjustmentStep() (*big.Int, error) {
	return _CurveAtri.Contract.AdjustmentStep(&_CurveAtri.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) AdjustmentStep() (*big.Int, error) {
	return _CurveAtri.Contract.AdjustmentStep(&_CurveAtri.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveAtri *CurveAtriSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurveAtri.Contract.AdminActionsDeadline(&_CurveAtri.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurveAtri.Contract.AdminActionsDeadline(&_CurveAtri.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAtri *CurveAtriSession) AdminFee() (*big.Int, error) {
	return _CurveAtri.Contract.AdminFee(&_CurveAtri.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) AdminFee() (*big.Int, error) {
	return _CurveAtri.Contract.AdminFee(&_CurveAtri.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveAtri *CurveAtriCaller) AdminFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "admin_fee_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveAtri *CurveAtriSession) AdminFeeReceiver() (common.Address, error) {
	return _CurveAtri.Contract.AdminFeeReceiver(&_CurveAtri.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveAtri *CurveAtriCallerSession) AdminFeeReceiver() (common.Address, error) {
	return _CurveAtri.Contract.AdminFeeReceiver(&_CurveAtri.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) AllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_CurveAtri *CurveAtriSession) AllowedExtraProfit() (*big.Int, error) {
	return _CurveAtri.Contract.AllowedExtraProfit(&_CurveAtri.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) AllowedExtraProfit() (*big.Int, error) {
	return _CurveAtri.Contract.AllowedExtraProfit(&_CurveAtri.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurveAtri *CurveAtriCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurveAtri *CurveAtriSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.Balances(&_CurveAtri.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.Balances(&_CurveAtri.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurveAtri *CurveAtriCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "calc_token_amount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurveAtri *CurveAtriSession) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _CurveAtri.Contract.CalcTokenAmount(&_CurveAtri.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _CurveAtri.Contract.CalcTokenAmount(&_CurveAtri.CallOpts, amounts, deposit)
}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_CurveAtri *CurveAtriCaller) CalcTokenFee(opts *bind.CallOpts, amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "calc_token_fee", amounts, xp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_CurveAtri *CurveAtriSession) CalcTokenFee(amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.CalcTokenFee(&_CurveAtri.CallOpts, amounts, xp)
}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) CalcTokenFee(amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.CalcTokenFee(&_CurveAtri.CallOpts, amounts, xp)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_CurveAtri *CurveAtriCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "calc_withdraw_one_coin", token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_CurveAtri *CurveAtriSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.CalcWithdrawOneCoin(&_CurveAtri.CallOpts, token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.CalcWithdrawOneCoin(&_CurveAtri.CallOpts, token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAtri *CurveAtriCaller) Coins(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "coins", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAtri *CurveAtriSession) Coins(i *big.Int) (common.Address, error) {
	return _CurveAtri.Contract.Coins(&_CurveAtri.CallOpts, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAtri *CurveAtriCallerSession) Coins(i *big.Int) (common.Address, error) {
	return _CurveAtri.Contract.Coins(&_CurveAtri.CallOpts, i)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAtri *CurveAtriSession) Fee() (*big.Int, error) {
	return _CurveAtri.Contract.Fee(&_CurveAtri.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) Fee() (*big.Int, error) {
	return _CurveAtri.Contract.Fee(&_CurveAtri.CallOpts)
}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_CurveAtri *CurveAtriCaller) FeeCalc(opts *bind.CallOpts, xp [3]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "fee_calc", xp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_CurveAtri *CurveAtriSession) FeeCalc(xp [3]*big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.FeeCalc(&_CurveAtri.CallOpts, xp)
}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) FeeCalc(xp [3]*big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.FeeCalc(&_CurveAtri.CallOpts, xp)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) FeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriSession) FeeGamma() (*big.Int, error) {
	return _CurveAtri.Contract.FeeGamma(&_CurveAtri.CallOpts)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) FeeGamma() (*big.Int, error) {
	return _CurveAtri.Contract.FeeGamma(&_CurveAtri.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) FutureAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "future_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriSession) FutureAGamma() (*big.Int, error) {
	return _CurveAtri.Contract.FutureAGamma(&_CurveAtri.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) FutureAGamma() (*big.Int, error) {
	return _CurveAtri.Contract.FutureAGamma(&_CurveAtri.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) FutureAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "future_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_CurveAtri *CurveAtriSession) FutureAGammaTime() (*big.Int, error) {
	return _CurveAtri.Contract.FutureAGammaTime(&_CurveAtri.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) FutureAGammaTime() (*big.Int, error) {
	return _CurveAtri.Contract.FutureAGammaTime(&_CurveAtri.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) FutureAdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "future_adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_CurveAtri *CurveAtriSession) FutureAdjustmentStep() (*big.Int, error) {
	return _CurveAtri.Contract.FutureAdjustmentStep(&_CurveAtri.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) FutureAdjustmentStep() (*big.Int, error) {
	return _CurveAtri.Contract.FutureAdjustmentStep(&_CurveAtri.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveAtri *CurveAtriSession) FutureAdminFee() (*big.Int, error) {
	return _CurveAtri.Contract.FutureAdminFee(&_CurveAtri.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) FutureAdminFee() (*big.Int, error) {
	return _CurveAtri.Contract.FutureAdminFee(&_CurveAtri.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) FutureAllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "future_allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_CurveAtri *CurveAtriSession) FutureAllowedExtraProfit() (*big.Int, error) {
	return _CurveAtri.Contract.FutureAllowedExtraProfit(&_CurveAtri.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) FutureAllowedExtraProfit() (*big.Int, error) {
	return _CurveAtri.Contract.FutureAllowedExtraProfit(&_CurveAtri.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) FutureFeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "future_fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriSession) FutureFeeGamma() (*big.Int, error) {
	return _CurveAtri.Contract.FutureFeeGamma(&_CurveAtri.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) FutureFeeGamma() (*big.Int, error) {
	return _CurveAtri.Contract.FutureFeeGamma(&_CurveAtri.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) FutureMaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "future_ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_CurveAtri *CurveAtriSession) FutureMaHalfTime() (*big.Int, error) {
	return _CurveAtri.Contract.FutureMaHalfTime(&_CurveAtri.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) FutureMaHalfTime() (*big.Int, error) {
	return _CurveAtri.Contract.FutureMaHalfTime(&_CurveAtri.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) FutureMidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "future_mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_CurveAtri *CurveAtriSession) FutureMidFee() (*big.Int, error) {
	return _CurveAtri.Contract.FutureMidFee(&_CurveAtri.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) FutureMidFee() (*big.Int, error) {
	return _CurveAtri.Contract.FutureMidFee(&_CurveAtri.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) FutureOutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "future_out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_CurveAtri *CurveAtriSession) FutureOutFee() (*big.Int, error) {
	return _CurveAtri.Contract.FutureOutFee(&_CurveAtri.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) FutureOutFee() (*big.Int, error) {
	return _CurveAtri.Contract.FutureOutFee(&_CurveAtri.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveAtri *CurveAtriCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveAtri *CurveAtriSession) FutureOwner() (common.Address, error) {
	return _CurveAtri.Contract.FutureOwner(&_CurveAtri.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveAtri *CurveAtriCallerSession) FutureOwner() (common.Address, error) {
	return _CurveAtri.Contract.FutureOwner(&_CurveAtri.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) Gamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_CurveAtri *CurveAtriSession) Gamma() (*big.Int, error) {
	return _CurveAtri.Contract.Gamma(&_CurveAtri.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) Gamma() (*big.Int, error) {
	return _CurveAtri.Contract.Gamma(&_CurveAtri.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_CurveAtri *CurveAtriCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_CurveAtri *CurveAtriSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.GetDy(&_CurveAtri.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.GetDy(&_CurveAtri.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAtri *CurveAtriSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveAtri.Contract.GetVirtualPrice(&_CurveAtri.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveAtri.Contract.GetVirtualPrice(&_CurveAtri.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) InitialAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "initial_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriSession) InitialAGamma() (*big.Int, error) {
	return _CurveAtri.Contract.InitialAGamma(&_CurveAtri.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) InitialAGamma() (*big.Int, error) {
	return _CurveAtri.Contract.InitialAGamma(&_CurveAtri.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) InitialAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "initial_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_CurveAtri *CurveAtriSession) InitialAGammaTime() (*big.Int, error) {
	return _CurveAtri.Contract.InitialAGammaTime(&_CurveAtri.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) InitialAGammaTime() (*big.Int, error) {
	return _CurveAtri.Contract.InitialAGammaTime(&_CurveAtri.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveAtri *CurveAtriCaller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "is_killed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveAtri *CurveAtriSession) IsKilled() (bool, error) {
	return _CurveAtri.Contract.IsKilled(&_CurveAtri.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveAtri *CurveAtriCallerSession) IsKilled() (bool, error) {
	return _CurveAtri.Contract.IsKilled(&_CurveAtri.CallOpts)
}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) KillDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "kill_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_CurveAtri *CurveAtriSession) KillDeadline() (*big.Int, error) {
	return _CurveAtri.Contract.KillDeadline(&_CurveAtri.CallOpts)
}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) KillDeadline() (*big.Int, error) {
	return _CurveAtri.Contract.KillDeadline(&_CurveAtri.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_CurveAtri *CurveAtriCaller) LastPrices(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "last_prices", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_CurveAtri *CurveAtriSession) LastPrices(k *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.LastPrices(&_CurveAtri.CallOpts, k)
}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) LastPrices(k *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.LastPrices(&_CurveAtri.CallOpts, k)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) LastPricesTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "last_prices_timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_CurveAtri *CurveAtriSession) LastPricesTimestamp() (*big.Int, error) {
	return _CurveAtri.Contract.LastPricesTimestamp(&_CurveAtri.CallOpts)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) LastPricesTimestamp() (*big.Int, error) {
	return _CurveAtri.Contract.LastPricesTimestamp(&_CurveAtri.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) MaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_CurveAtri *CurveAtriSession) MaHalfTime() (*big.Int, error) {
	return _CurveAtri.Contract.MaHalfTime(&_CurveAtri.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) MaHalfTime() (*big.Int, error) {
	return _CurveAtri.Contract.MaHalfTime(&_CurveAtri.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) MidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_CurveAtri *CurveAtriSession) MidFee() (*big.Int, error) {
	return _CurveAtri.Contract.MidFee(&_CurveAtri.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) MidFee() (*big.Int, error) {
	return _CurveAtri.Contract.MidFee(&_CurveAtri.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) OutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_CurveAtri *CurveAtriSession) OutFee() (*big.Int, error) {
	return _CurveAtri.Contract.OutFee(&_CurveAtri.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) OutFee() (*big.Int, error) {
	return _CurveAtri.Contract.OutFee(&_CurveAtri.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveAtri *CurveAtriCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveAtri *CurveAtriSession) Owner() (common.Address, error) {
	return _CurveAtri.Contract.Owner(&_CurveAtri.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveAtri *CurveAtriCallerSession) Owner() (common.Address, error) {
	return _CurveAtri.Contract.Owner(&_CurveAtri.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_CurveAtri *CurveAtriCaller) PriceOracle(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "price_oracle", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_CurveAtri *CurveAtriSession) PriceOracle(k *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.PriceOracle(&_CurveAtri.CallOpts, k)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) PriceOracle(k *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.PriceOracle(&_CurveAtri.CallOpts, k)
}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_CurveAtri *CurveAtriCaller) PriceScale(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "price_scale", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_CurveAtri *CurveAtriSession) PriceScale(k *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.PriceScale(&_CurveAtri.CallOpts, k)
}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) PriceScale(k *big.Int) (*big.Int, error) {
	return _CurveAtri.Contract.PriceScale(&_CurveAtri.CallOpts, k)
}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveAtri *CurveAtriCaller) RewardReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "reward_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveAtri *CurveAtriSession) RewardReceiver() (common.Address, error) {
	return _CurveAtri.Contract.RewardReceiver(&_CurveAtri.CallOpts)
}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveAtri *CurveAtriCallerSession) RewardReceiver() (common.Address, error) {
	return _CurveAtri.Contract.RewardReceiver(&_CurveAtri.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAtri *CurveAtriCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAtri *CurveAtriSession) Token() (common.Address, error) {
	return _CurveAtri.Contract.Token(&_CurveAtri.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAtri *CurveAtriCallerSession) Token() (common.Address, error) {
	return _CurveAtri.Contract.Token(&_CurveAtri.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveAtri *CurveAtriSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurveAtri.Contract.TransferOwnershipDeadline(&_CurveAtri.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurveAtri.Contract.TransferOwnershipDeadline(&_CurveAtri.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) VirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_CurveAtri *CurveAtriSession) VirtualPrice() (*big.Int, error) {
	return _CurveAtri.Contract.VirtualPrice(&_CurveAtri.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) VirtualPrice() (*big.Int, error) {
	return _CurveAtri.Contract.VirtualPrice(&_CurveAtri.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) XcpProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "xcp_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_CurveAtri *CurveAtriSession) XcpProfit() (*big.Int, error) {
	return _CurveAtri.Contract.XcpProfit(&_CurveAtri.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) XcpProfit() (*big.Int, error) {
	return _CurveAtri.Contract.XcpProfit(&_CurveAtri.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_CurveAtri *CurveAtriCaller) XcpProfitA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAtri.contract.Call(opts, &out, "xcp_profit_a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_CurveAtri *CurveAtriSession) XcpProfitA() (*big.Int, error) {
	return _CurveAtri.Contract.XcpProfitA(&_CurveAtri.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_CurveAtri *CurveAtriCallerSession) XcpProfitA() (*big.Int, error) {
	return _CurveAtri.Contract.XcpProfitA(&_CurveAtri.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_CurveAtri *CurveAtriTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_CurveAtri *CurveAtriSession) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.AddLiquidity(&_CurveAtri.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_CurveAtri *CurveAtriTransactorSession) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.AddLiquidity(&_CurveAtri.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_CurveAtri *CurveAtriTransactor) ApplyNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "apply_new_parameters")
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_CurveAtri *CurveAtriSession) ApplyNewParameters() (*types.Transaction, error) {
	return _CurveAtri.Contract.ApplyNewParameters(&_CurveAtri.TransactOpts)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_CurveAtri *CurveAtriTransactorSession) ApplyNewParameters() (*types.Transaction, error) {
	return _CurveAtri.Contract.ApplyNewParameters(&_CurveAtri.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveAtri *CurveAtriTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveAtri *CurveAtriSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveAtri.Contract.ApplyTransferOwnership(&_CurveAtri.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveAtri *CurveAtriTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveAtri.Contract.ApplyTransferOwnership(&_CurveAtri.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_CurveAtri *CurveAtriTransactor) ClaimAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "claim_admin_fees")
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_CurveAtri *CurveAtriSession) ClaimAdminFees() (*types.Transaction, error) {
	return _CurveAtri.Contract.ClaimAdminFees(&_CurveAtri.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_CurveAtri *CurveAtriTransactorSession) ClaimAdminFees() (*types.Transaction, error) {
	return _CurveAtri.Contract.ClaimAdminFees(&_CurveAtri.TransactOpts)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_CurveAtri *CurveAtriTransactor) CommitNewParameters(opts *bind.TransactOpts, _new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "commit_new_parameters", _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_CurveAtri *CurveAtriSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.CommitNewParameters(&_CurveAtri.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_CurveAtri *CurveAtriTransactorSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.CommitNewParameters(&_CurveAtri.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveAtri *CurveAtriTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveAtri *CurveAtriSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveAtri.Contract.CommitTransferOwnership(&_CurveAtri.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveAtri *CurveAtriTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveAtri.Contract.CommitTransferOwnership(&_CurveAtri.TransactOpts, _owner)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveAtri *CurveAtriTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveAtri *CurveAtriSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.Exchange(&_CurveAtri.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) returns(uint256)
func (_CurveAtri *CurveAtriTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.Exchange(&_CurveAtri.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveAtri *CurveAtriTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveAtri *CurveAtriSession) KillMe() (*types.Transaction, error) {
	return _CurveAtri.Contract.KillMe(&_CurveAtri.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveAtri *CurveAtriTransactorSession) KillMe() (*types.Transaction, error) {
	return _CurveAtri.Contract.KillMe(&_CurveAtri.TransactOpts)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_CurveAtri *CurveAtriTransactor) RampAGamma(opts *bind.TransactOpts, future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "ramp_A_gamma", future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_CurveAtri *CurveAtriSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.RampAGamma(&_CurveAtri.TransactOpts, future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_CurveAtri *CurveAtriTransactorSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.RampAGamma(&_CurveAtri.TransactOpts, future_A, future_gamma, future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_CurveAtri *CurveAtriTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_CurveAtri *CurveAtriSession) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.RemoveLiquidity(&_CurveAtri.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_CurveAtri *CurveAtriTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.RemoveLiquidity(&_CurveAtri.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns()
func (_CurveAtri *CurveAtriTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "remove_liquidity_one_coin", token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns()
func (_CurveAtri *CurveAtriSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.RemoveLiquidityOneCoin(&_CurveAtri.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns()
func (_CurveAtri *CurveAtriTransactorSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurveAtri.Contract.RemoveLiquidityOneCoin(&_CurveAtri.TransactOpts, token_amount, i, min_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveAtri *CurveAtriTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveAtri *CurveAtriSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurveAtri.Contract.RevertNewParameters(&_CurveAtri.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveAtri *CurveAtriTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurveAtri.Contract.RevertNewParameters(&_CurveAtri.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveAtri *CurveAtriTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveAtri *CurveAtriSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurveAtri.Contract.RevertTransferOwnership(&_CurveAtri.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveAtri *CurveAtriTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurveAtri.Contract.RevertTransferOwnership(&_CurveAtri.TransactOpts)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveAtri *CurveAtriTransactor) SetAdminFeeReceiver(opts *bind.TransactOpts, _admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "set_admin_fee_receiver", _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveAtri *CurveAtriSession) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveAtri.Contract.SetAdminFeeReceiver(&_CurveAtri.TransactOpts, _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveAtri *CurveAtriTransactorSession) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveAtri.Contract.SetAdminFeeReceiver(&_CurveAtri.TransactOpts, _admin_fee_receiver)
}

// SetRewardReceiver is a paid mutator transaction binding the contract method 0xc51b8861.
//
// Solidity: function set_reward_receiver(address _reward_receiver) returns()
func (_CurveAtri *CurveAtriTransactor) SetRewardReceiver(opts *bind.TransactOpts, _reward_receiver common.Address) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "set_reward_receiver", _reward_receiver)
}

// SetRewardReceiver is a paid mutator transaction binding the contract method 0xc51b8861.
//
// Solidity: function set_reward_receiver(address _reward_receiver) returns()
func (_CurveAtri *CurveAtriSession) SetRewardReceiver(_reward_receiver common.Address) (*types.Transaction, error) {
	return _CurveAtri.Contract.SetRewardReceiver(&_CurveAtri.TransactOpts, _reward_receiver)
}

// SetRewardReceiver is a paid mutator transaction binding the contract method 0xc51b8861.
//
// Solidity: function set_reward_receiver(address _reward_receiver) returns()
func (_CurveAtri *CurveAtriTransactorSession) SetRewardReceiver(_reward_receiver common.Address) (*types.Transaction, error) {
	return _CurveAtri.Contract.SetRewardReceiver(&_CurveAtri.TransactOpts, _reward_receiver)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_CurveAtri *CurveAtriTransactor) StopRampAGamma(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "stop_ramp_A_gamma")
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_CurveAtri *CurveAtriSession) StopRampAGamma() (*types.Transaction, error) {
	return _CurveAtri.Contract.StopRampAGamma(&_CurveAtri.TransactOpts)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_CurveAtri *CurveAtriTransactorSession) StopRampAGamma() (*types.Transaction, error) {
	return _CurveAtri.Contract.StopRampAGamma(&_CurveAtri.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveAtri *CurveAtriTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAtri.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveAtri *CurveAtriSession) UnkillMe() (*types.Transaction, error) {
	return _CurveAtri.Contract.UnkillMe(&_CurveAtri.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveAtri *CurveAtriTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _CurveAtri.Contract.UnkillMe(&_CurveAtri.TransactOpts)
}
