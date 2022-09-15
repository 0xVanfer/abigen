// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveGauge

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

// CurveGaugeMetaData contains all meta data concerning the CurveGauge contract.
var CurveGaugeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"accept_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"arg1\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"claim_rewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claim_rewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim_rewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim_sig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimable_reward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimable_reward_write\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimed_reward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtracted_value\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_claim_rewards\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_added_value\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"reward_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward_contract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"reward_integral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"arg1\",\"type\":\"address\"}],\"name\":\"reward_integral_for\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"reward_tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"rewards_receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reward_contract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_claim_sig\",\"type\":\"bytes32\"},{\"internalType\":\"address[8]\",\"name\":\"_reward_tokens\",\"type\":\"address[8]\"}],\"name\":\"set_rewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"set_rewards_receiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_claim_rewards\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CurveGaugeABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveGaugeMetaData.ABI instead.
var CurveGaugeABI = CurveGaugeMetaData.ABI

// CurveGauge is an auto generated Go binding around an Ethereum contract.
type CurveGauge struct {
	CurveGaugeCaller     // Read-only binding to the contract
	CurveGaugeTransactor // Write-only binding to the contract
	CurveGaugeFilterer   // Log filterer for contract events
}

// CurveGaugeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveGaugeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveGaugeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveGaugeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveGaugeSession struct {
	Contract     *CurveGauge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveGaugeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveGaugeCallerSession struct {
	Contract *CurveGaugeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CurveGaugeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveGaugeTransactorSession struct {
	Contract     *CurveGaugeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CurveGaugeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveGaugeRaw struct {
	Contract *CurveGauge // Generic contract binding to access the raw methods on
}

// CurveGaugeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveGaugeCallerRaw struct {
	Contract *CurveGaugeCaller // Generic read-only contract binding to access the raw methods on
}

// CurveGaugeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveGaugeTransactorRaw struct {
	Contract *CurveGaugeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveGauge creates a new instance of CurveGauge, bound to a specific deployed contract.
func NewCurveGauge(address common.Address, backend bind.ContractBackend) (*CurveGauge, error) {
	contract, err := bindCurveGauge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveGauge{CurveGaugeCaller: CurveGaugeCaller{contract: contract}, CurveGaugeTransactor: CurveGaugeTransactor{contract: contract}, CurveGaugeFilterer: CurveGaugeFilterer{contract: contract}}, nil
}

// NewCurveGaugeCaller creates a new read-only instance of CurveGauge, bound to a specific deployed contract.
func NewCurveGaugeCaller(address common.Address, caller bind.ContractCaller) (*CurveGaugeCaller, error) {
	contract, err := bindCurveGauge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeCaller{contract: contract}, nil
}

// NewCurveGaugeTransactor creates a new write-only instance of CurveGauge, bound to a specific deployed contract.
func NewCurveGaugeTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveGaugeTransactor, error) {
	contract, err := bindCurveGauge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeTransactor{contract: contract}, nil
}

// NewCurveGaugeFilterer creates a new log filterer instance of CurveGauge, bound to a specific deployed contract.
func NewCurveGaugeFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveGaugeFilterer, error) {
	contract, err := bindCurveGauge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeFilterer{contract: contract}, nil
}

// bindCurveGauge binds a generic wrapper to an already deployed contract.
func bindCurveGauge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveGaugeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGauge *CurveGaugeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGauge.Contract.CurveGaugeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGauge *CurveGaugeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGauge.Contract.CurveGaugeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGauge *CurveGaugeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGauge.Contract.CurveGaugeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGauge *CurveGaugeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGauge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGauge *CurveGaugeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGauge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGauge *CurveGaugeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGauge.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveGauge *CurveGaugeCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveGauge *CurveGaugeSession) Admin() (common.Address, error) {
	return _CurveGauge.Contract.Admin(&_CurveGauge.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveGauge *CurveGaugeCallerSession) Admin() (common.Address, error) {
	return _CurveGauge.Contract.Admin(&_CurveGauge.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.Allowance(&_CurveGauge.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.Allowance(&_CurveGauge.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.BalanceOf(&_CurveGauge.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.BalanceOf(&_CurveGauge.CallOpts, arg0)
}

// ClaimSig is a free data retrieval call binding the contract method 0xb7aca568.
//
// Solidity: function claim_sig() view returns(bytes)
func (_CurveGauge *CurveGaugeCaller) ClaimSig(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "claim_sig")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ClaimSig is a free data retrieval call binding the contract method 0xb7aca568.
//
// Solidity: function claim_sig() view returns(bytes)
func (_CurveGauge *CurveGaugeSession) ClaimSig() ([]byte, error) {
	return _CurveGauge.Contract.ClaimSig(&_CurveGauge.CallOpts)
}

// ClaimSig is a free data retrieval call binding the contract method 0xb7aca568.
//
// Solidity: function claim_sig() view returns(bytes)
func (_CurveGauge *CurveGaugeCallerSession) ClaimSig() ([]byte, error) {
	return _CurveGauge.Contract.ClaimSig(&_CurveGauge.CallOpts)
}

// ClaimableReward is a free data retrieval call binding the contract method 0x33fd6f74.
//
// Solidity: function claimable_reward(address _addr, address _token) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) ClaimableReward(opts *bind.CallOpts, _addr common.Address, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "claimable_reward", _addr, _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableReward is a free data retrieval call binding the contract method 0x33fd6f74.
//
// Solidity: function claimable_reward(address _addr, address _token) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) ClaimableReward(_addr common.Address, _token common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.ClaimableReward(&_CurveGauge.CallOpts, _addr, _token)
}

// ClaimableReward is a free data retrieval call binding the contract method 0x33fd6f74.
//
// Solidity: function claimable_reward(address _addr, address _token) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) ClaimableReward(_addr common.Address, _token common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.ClaimableReward(&_CurveGauge.CallOpts, _addr, _token)
}

// ClaimedReward is a free data retrieval call binding the contract method 0xe77e7437.
//
// Solidity: function claimed_reward(address _addr, address _token) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) ClaimedReward(opts *bind.CallOpts, _addr common.Address, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "claimed_reward", _addr, _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedReward is a free data retrieval call binding the contract method 0xe77e7437.
//
// Solidity: function claimed_reward(address _addr, address _token) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) ClaimedReward(_addr common.Address, _token common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.ClaimedReward(&_CurveGauge.CallOpts, _addr, _token)
}

// ClaimedReward is a free data retrieval call binding the contract method 0xe77e7437.
//
// Solidity: function claimed_reward(address _addr, address _token) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) ClaimedReward(_addr common.Address, _token common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.ClaimedReward(&_CurveGauge.CallOpts, _addr, _token)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveGauge *CurveGaugeSession) Decimals() (*big.Int, error) {
	return _CurveGauge.Contract.Decimals(&_CurveGauge.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) Decimals() (*big.Int, error) {
	return _CurveGauge.Contract.Decimals(&_CurveGauge.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveGauge *CurveGaugeCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveGauge *CurveGaugeSession) FutureAdmin() (common.Address, error) {
	return _CurveGauge.Contract.FutureAdmin(&_CurveGauge.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveGauge *CurveGaugeCallerSession) FutureAdmin() (common.Address, error) {
	return _CurveGauge.Contract.FutureAdmin(&_CurveGauge.CallOpts)
}

// LastClaim is a free data retrieval call binding the contract method 0x3488bd19.
//
// Solidity: function last_claim() view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) LastClaim(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "last_claim")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastClaim is a free data retrieval call binding the contract method 0x3488bd19.
//
// Solidity: function last_claim() view returns(uint256)
func (_CurveGauge *CurveGaugeSession) LastClaim() (*big.Int, error) {
	return _CurveGauge.Contract.LastClaim(&_CurveGauge.CallOpts)
}

// LastClaim is a free data retrieval call binding the contract method 0x3488bd19.
//
// Solidity: function last_claim() view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) LastClaim() (*big.Int, error) {
	return _CurveGauge.Contract.LastClaim(&_CurveGauge.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveGauge *CurveGaugeCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "lp_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveGauge *CurveGaugeSession) LpToken() (common.Address, error) {
	return _CurveGauge.Contract.LpToken(&_CurveGauge.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveGauge *CurveGaugeCallerSession) LpToken() (common.Address, error) {
	return _CurveGauge.Contract.LpToken(&_CurveGauge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveGauge *CurveGaugeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveGauge *CurveGaugeSession) Name() (string, error) {
	return _CurveGauge.Contract.Name(&_CurveGauge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveGauge *CurveGaugeCallerSession) Name() (string, error) {
	return _CurveGauge.Contract.Name(&_CurveGauge.CallOpts)
}

// RewardBalances is a free data retrieval call binding the contract method 0xfe9e2178.
//
// Solidity: function reward_balances(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) RewardBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "reward_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardBalances is a free data retrieval call binding the contract method 0xfe9e2178.
//
// Solidity: function reward_balances(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) RewardBalances(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.RewardBalances(&_CurveGauge.CallOpts, arg0)
}

// RewardBalances is a free data retrieval call binding the contract method 0xfe9e2178.
//
// Solidity: function reward_balances(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) RewardBalances(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.RewardBalances(&_CurveGauge.CallOpts, arg0)
}

// RewardContract is a free data retrieval call binding the contract method 0xbf88a6ff.
//
// Solidity: function reward_contract() view returns(address)
func (_CurveGauge *CurveGaugeCaller) RewardContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "reward_contract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardContract is a free data retrieval call binding the contract method 0xbf88a6ff.
//
// Solidity: function reward_contract() view returns(address)
func (_CurveGauge *CurveGaugeSession) RewardContract() (common.Address, error) {
	return _CurveGauge.Contract.RewardContract(&_CurveGauge.CallOpts)
}

// RewardContract is a free data retrieval call binding the contract method 0xbf88a6ff.
//
// Solidity: function reward_contract() view returns(address)
func (_CurveGauge *CurveGaugeCallerSession) RewardContract() (common.Address, error) {
	return _CurveGauge.Contract.RewardContract(&_CurveGauge.CallOpts)
}

// RewardIntegral is a free data retrieval call binding the contract method 0x73861fb3.
//
// Solidity: function reward_integral(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) RewardIntegral(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "reward_integral", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardIntegral is a free data retrieval call binding the contract method 0x73861fb3.
//
// Solidity: function reward_integral(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) RewardIntegral(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.RewardIntegral(&_CurveGauge.CallOpts, arg0)
}

// RewardIntegral is a free data retrieval call binding the contract method 0x73861fb3.
//
// Solidity: function reward_integral(address arg0) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) RewardIntegral(arg0 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.RewardIntegral(&_CurveGauge.CallOpts, arg0)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xf05cc058.
//
// Solidity: function reward_integral_for(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) RewardIntegralFor(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "reward_integral_for", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xf05cc058.
//
// Solidity: function reward_integral_for(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeSession) RewardIntegralFor(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.RewardIntegralFor(&_CurveGauge.CallOpts, arg0, arg1)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xf05cc058.
//
// Solidity: function reward_integral_for(address arg0, address arg1) view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) RewardIntegralFor(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGauge.Contract.RewardIntegralFor(&_CurveGauge.CallOpts, arg0, arg1)
}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveGauge *CurveGaugeCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "reward_tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveGauge *CurveGaugeSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _CurveGauge.Contract.RewardTokens(&_CurveGauge.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveGauge *CurveGaugeCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _CurveGauge.Contract.RewardTokens(&_CurveGauge.CallOpts, arg0)
}

// RewardsReceiver is a free data retrieval call binding the contract method 0x01ddabf1.
//
// Solidity: function rewards_receiver(address arg0) view returns(address)
func (_CurveGauge *CurveGaugeCaller) RewardsReceiver(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "rewards_receiver", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsReceiver is a free data retrieval call binding the contract method 0x01ddabf1.
//
// Solidity: function rewards_receiver(address arg0) view returns(address)
func (_CurveGauge *CurveGaugeSession) RewardsReceiver(arg0 common.Address) (common.Address, error) {
	return _CurveGauge.Contract.RewardsReceiver(&_CurveGauge.CallOpts, arg0)
}

// RewardsReceiver is a free data retrieval call binding the contract method 0x01ddabf1.
//
// Solidity: function rewards_receiver(address arg0) view returns(address)
func (_CurveGauge *CurveGaugeCallerSession) RewardsReceiver(arg0 common.Address) (common.Address, error) {
	return _CurveGauge.Contract.RewardsReceiver(&_CurveGauge.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveGauge *CurveGaugeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveGauge *CurveGaugeSession) Symbol() (string, error) {
	return _CurveGauge.Contract.Symbol(&_CurveGauge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveGauge *CurveGaugeCallerSession) Symbol() (string, error) {
	return _CurveGauge.Contract.Symbol(&_CurveGauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveGauge *CurveGaugeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGauge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveGauge *CurveGaugeSession) TotalSupply() (*big.Int, error) {
	return _CurveGauge.Contract.TotalSupply(&_CurveGauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveGauge *CurveGaugeCallerSession) TotalSupply() (*big.Int, error) {
	return _CurveGauge.Contract.TotalSupply(&_CurveGauge.CallOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveGauge *CurveGaugeTransactor) AcceptTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "accept_transfer_ownership")
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveGauge *CurveGaugeSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveGauge.Contract.AcceptTransferOwnership(&_CurveGauge.TransactOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveGauge *CurveGaugeTransactorSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveGauge.Contract.AcceptTransferOwnership(&_CurveGauge.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Approve(&_CurveGauge.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Approve(&_CurveGauge.TransactOpts, _spender, _value)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x84e9bd7e.
//
// Solidity: function claim_rewards(address _addr) returns()
func (_CurveGauge *CurveGaugeTransactor) ClaimRewards(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "claim_rewards", _addr)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x84e9bd7e.
//
// Solidity: function claim_rewards(address _addr) returns()
func (_CurveGauge *CurveGaugeSession) ClaimRewards(_addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards(&_CurveGauge.TransactOpts, _addr)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x84e9bd7e.
//
// Solidity: function claim_rewards(address _addr) returns()
func (_CurveGauge *CurveGaugeTransactorSession) ClaimRewards(_addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards(&_CurveGauge.TransactOpts, _addr)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x9faceb1b.
//
// Solidity: function claim_rewards(address _addr, address _receiver) returns()
func (_CurveGauge *CurveGaugeTransactor) ClaimRewards0(opts *bind.TransactOpts, _addr common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "claim_rewards0", _addr, _receiver)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x9faceb1b.
//
// Solidity: function claim_rewards(address _addr, address _receiver) returns()
func (_CurveGauge *CurveGaugeSession) ClaimRewards0(_addr common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards0(&_CurveGauge.TransactOpts, _addr, _receiver)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x9faceb1b.
//
// Solidity: function claim_rewards(address _addr, address _receiver) returns()
func (_CurveGauge *CurveGaugeTransactorSession) ClaimRewards0(_addr common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards0(&_CurveGauge.TransactOpts, _addr, _receiver)
}

// ClaimRewards1 is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_CurveGauge *CurveGaugeTransactor) ClaimRewards1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "claim_rewards1")
}

// ClaimRewards1 is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_CurveGauge *CurveGaugeSession) ClaimRewards1() (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards1(&_CurveGauge.TransactOpts)
}

// ClaimRewards1 is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_CurveGauge *CurveGaugeTransactorSession) ClaimRewards1() (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimRewards1(&_CurveGauge.TransactOpts)
}

// ClaimableRewardWrite is a paid mutator transaction binding the contract method 0x59b7e409.
//
// Solidity: function claimable_reward_write(address _addr, address _token) returns(uint256)
func (_CurveGauge *CurveGaugeTransactor) ClaimableRewardWrite(opts *bind.TransactOpts, _addr common.Address, _token common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "claimable_reward_write", _addr, _token)
}

// ClaimableRewardWrite is a paid mutator transaction binding the contract method 0x59b7e409.
//
// Solidity: function claimable_reward_write(address _addr, address _token) returns(uint256)
func (_CurveGauge *CurveGaugeSession) ClaimableRewardWrite(_addr common.Address, _token common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimableRewardWrite(&_CurveGauge.TransactOpts, _addr, _token)
}

// ClaimableRewardWrite is a paid mutator transaction binding the contract method 0x59b7e409.
//
// Solidity: function claimable_reward_write(address _addr, address _token) returns(uint256)
func (_CurveGauge *CurveGaugeTransactorSession) ClaimableRewardWrite(_addr common.Address, _token common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.ClaimableRewardWrite(&_CurveGauge.TransactOpts, _addr, _token)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveGauge *CurveGaugeTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveGauge *CurveGaugeSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.CommitTransferOwnership(&_CurveGauge.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveGauge *CurveGaugeTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.CommitTransferOwnership(&_CurveGauge.TransactOpts, addr)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveGauge *CurveGaugeTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "decreaseAllowance", _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveGauge *CurveGaugeSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.DecreaseAllowance(&_CurveGauge.TransactOpts, _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveGauge *CurveGaugeTransactorSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.DecreaseAllowance(&_CurveGauge.TransactOpts, _spender, _subtracted_value)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address _addr) returns()
func (_CurveGauge *CurveGaugeTransactor) Deposit(opts *bind.TransactOpts, _value *big.Int, _addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "deposit", _value, _addr)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address _addr) returns()
func (_CurveGauge *CurveGaugeSession) Deposit(_value *big.Int, _addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit(&_CurveGauge.TransactOpts, _value, _addr)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address _addr) returns()
func (_CurveGauge *CurveGaugeTransactorSession) Deposit(_value *big.Int, _addr common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit(&_CurveGauge.TransactOpts, _value, _addr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x83df6747.
//
// Solidity: function deposit(uint256 _value, address _addr, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeTransactor) Deposit0(opts *bind.TransactOpts, _value *big.Int, _addr common.Address, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "deposit0", _value, _addr, _claim_rewards)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x83df6747.
//
// Solidity: function deposit(uint256 _value, address _addr, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeSession) Deposit0(_value *big.Int, _addr common.Address, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit0(&_CurveGauge.TransactOpts, _value, _addr, _claim_rewards)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x83df6747.
//
// Solidity: function deposit(uint256 _value, address _addr, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeTransactorSession) Deposit0(_value *big.Int, _addr common.Address, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit0(&_CurveGauge.TransactOpts, _value, _addr, _claim_rewards)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_CurveGauge *CurveGaugeTransactor) Deposit1(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "deposit1", _value)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_CurveGauge *CurveGaugeSession) Deposit1(_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit1(&_CurveGauge.TransactOpts, _value)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_CurveGauge *CurveGaugeTransactorSession) Deposit1(_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Deposit1(&_CurveGauge.TransactOpts, _value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveGauge *CurveGaugeTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "increaseAllowance", _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveGauge *CurveGaugeSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.IncreaseAllowance(&_CurveGauge.TransactOpts, _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveGauge *CurveGaugeTransactorSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.IncreaseAllowance(&_CurveGauge.TransactOpts, _spender, _added_value)
}

// SetRewards is a paid mutator transaction binding the contract method 0x47d2d5d3.
//
// Solidity: function set_rewards(address _reward_contract, bytes32 _claim_sig, address[8] _reward_tokens) returns()
func (_CurveGauge *CurveGaugeTransactor) SetRewards(opts *bind.TransactOpts, _reward_contract common.Address, _claim_sig [32]byte, _reward_tokens [8]common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "set_rewards", _reward_contract, _claim_sig, _reward_tokens)
}

// SetRewards is a paid mutator transaction binding the contract method 0x47d2d5d3.
//
// Solidity: function set_rewards(address _reward_contract, bytes32 _claim_sig, address[8] _reward_tokens) returns()
func (_CurveGauge *CurveGaugeSession) SetRewards(_reward_contract common.Address, _claim_sig [32]byte, _reward_tokens [8]common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.SetRewards(&_CurveGauge.TransactOpts, _reward_contract, _claim_sig, _reward_tokens)
}

// SetRewards is a paid mutator transaction binding the contract method 0x47d2d5d3.
//
// Solidity: function set_rewards(address _reward_contract, bytes32 _claim_sig, address[8] _reward_tokens) returns()
func (_CurveGauge *CurveGaugeTransactorSession) SetRewards(_reward_contract common.Address, _claim_sig [32]byte, _reward_tokens [8]common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.SetRewards(&_CurveGauge.TransactOpts, _reward_contract, _claim_sig, _reward_tokens)
}

// SetRewardsReceiver is a paid mutator transaction binding the contract method 0xbdf98116.
//
// Solidity: function set_rewards_receiver(address _receiver) returns()
func (_CurveGauge *CurveGaugeTransactor) SetRewardsReceiver(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "set_rewards_receiver", _receiver)
}

// SetRewardsReceiver is a paid mutator transaction binding the contract method 0xbdf98116.
//
// Solidity: function set_rewards_receiver(address _receiver) returns()
func (_CurveGauge *CurveGaugeSession) SetRewardsReceiver(_receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.SetRewardsReceiver(&_CurveGauge.TransactOpts, _receiver)
}

// SetRewardsReceiver is a paid mutator transaction binding the contract method 0xbdf98116.
//
// Solidity: function set_rewards_receiver(address _receiver) returns()
func (_CurveGauge *CurveGaugeTransactorSession) SetRewardsReceiver(_receiver common.Address) (*types.Transaction, error) {
	return _CurveGauge.Contract.SetRewardsReceiver(&_CurveGauge.TransactOpts, _receiver)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Transfer(&_CurveGauge.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Transfer(&_CurveGauge.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.TransferFrom(&_CurveGauge.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveGauge *CurveGaugeTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.TransferFrom(&_CurveGauge.TransactOpts, _from, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_CurveGauge *CurveGaugeTransactor) Withdraw(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "withdraw", _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_CurveGauge *CurveGaugeSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Withdraw(&_CurveGauge.TransactOpts, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_CurveGauge *CurveGaugeTransactorSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _CurveGauge.Contract.Withdraw(&_CurveGauge.TransactOpts, _value)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 _value, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeTransactor) Withdraw0(opts *bind.TransactOpts, _value *big.Int, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.contract.Transact(opts, "withdraw0", _value, _claim_rewards)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 _value, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeSession) Withdraw0(_value *big.Int, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.Contract.Withdraw0(&_CurveGauge.TransactOpts, _value, _claim_rewards)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 _value, bool _claim_rewards) returns()
func (_CurveGauge *CurveGaugeTransactorSession) Withdraw0(_value *big.Int, _claim_rewards bool) (*types.Transaction, error) {
	return _CurveGauge.Contract.Withdraw0(&_CurveGauge.TransactOpts, _value, _claim_rewards)
}
