// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderjoeChefV3

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

// TraderjoeChefV3MetaData contains all meta data concerning the TraderjoeChefV3 contract.
var TraderjoeChefV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"JOE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_CHEF_V2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_PID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestFromMasterChef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dummyToken\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joePerSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pids\",\"type\":\"uint256[]\"}],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingJoe\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pendingBonusToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"accJoePerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewarder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pools\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TraderjoeChefV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderjoeChefV3MetaData.ABI instead.
var TraderjoeChefV3ABI = TraderjoeChefV3MetaData.ABI

// TraderjoeChefV3 is an auto generated Go binding around an Ethereum contract.
type TraderjoeChefV3 struct {
	TraderjoeChefV3Caller     // Read-only binding to the contract
	TraderjoeChefV3Transactor // Write-only binding to the contract
	TraderjoeChefV3Filterer   // Log filterer for contract events
}

// TraderjoeChefV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type TraderjoeChefV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeChefV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderjoeChefV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeChefV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderjoeChefV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeChefV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderjoeChefV3Session struct {
	Contract     *TraderjoeChefV3  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraderjoeChefV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderjoeChefV3CallerSession struct {
	Contract *TraderjoeChefV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TraderjoeChefV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderjoeChefV3TransactorSession struct {
	Contract     *TraderjoeChefV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TraderjoeChefV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type TraderjoeChefV3Raw struct {
	Contract *TraderjoeChefV3 // Generic contract binding to access the raw methods on
}

// TraderjoeChefV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderjoeChefV3CallerRaw struct {
	Contract *TraderjoeChefV3Caller // Generic read-only contract binding to access the raw methods on
}

// TraderjoeChefV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderjoeChefV3TransactorRaw struct {
	Contract *TraderjoeChefV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderjoeChefV3 creates a new instance of TraderjoeChefV3, bound to a specific deployed contract.
func NewTraderjoeChefV3(address common.Address, backend bind.ContractBackend) (*TraderjoeChefV3, error) {
	contract, err := bindTraderjoeChefV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraderjoeChefV3{TraderjoeChefV3Caller: TraderjoeChefV3Caller{contract: contract}, TraderjoeChefV3Transactor: TraderjoeChefV3Transactor{contract: contract}, TraderjoeChefV3Filterer: TraderjoeChefV3Filterer{contract: contract}}, nil
}

// NewTraderjoeChefV3Caller creates a new read-only instance of TraderjoeChefV3, bound to a specific deployed contract.
func NewTraderjoeChefV3Caller(address common.Address, caller bind.ContractCaller) (*TraderjoeChefV3Caller, error) {
	contract, err := bindTraderjoeChefV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeChefV3Caller{contract: contract}, nil
}

// NewTraderjoeChefV3Transactor creates a new write-only instance of TraderjoeChefV3, bound to a specific deployed contract.
func NewTraderjoeChefV3Transactor(address common.Address, transactor bind.ContractTransactor) (*TraderjoeChefV3Transactor, error) {
	contract, err := bindTraderjoeChefV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeChefV3Transactor{contract: contract}, nil
}

// NewTraderjoeChefV3Filterer creates a new log filterer instance of TraderjoeChefV3, bound to a specific deployed contract.
func NewTraderjoeChefV3Filterer(address common.Address, filterer bind.ContractFilterer) (*TraderjoeChefV3Filterer, error) {
	contract, err := bindTraderjoeChefV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderjoeChefV3Filterer{contract: contract}, nil
}

// bindTraderjoeChefV3 binds a generic wrapper to an already deployed contract.
func bindTraderjoeChefV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraderjoeChefV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeChefV3 *TraderjoeChefV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeChefV3.Contract.TraderjoeChefV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeChefV3 *TraderjoeChefV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.TraderjoeChefV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeChefV3 *TraderjoeChefV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.TraderjoeChefV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeChefV3 *TraderjoeChefV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeChefV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.contract.Transact(opts, method, params...)
}

// JOE is a free data retrieval call binding the contract method 0xffebad30.
//
// Solidity: function JOE() view returns(address)
func (_TraderjoeChefV3 *TraderjoeChefV3Caller) JOE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeChefV3.contract.Call(opts, &out, "JOE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JOE is a free data retrieval call binding the contract method 0xffebad30.
//
// Solidity: function JOE() view returns(address)
func (_TraderjoeChefV3 *TraderjoeChefV3Session) JOE() (common.Address, error) {
	return _TraderjoeChefV3.Contract.JOE(&_TraderjoeChefV3.CallOpts)
}

// JOE is a free data retrieval call binding the contract method 0xffebad30.
//
// Solidity: function JOE() view returns(address)
func (_TraderjoeChefV3 *TraderjoeChefV3CallerSession) JOE() (common.Address, error) {
	return _TraderjoeChefV3.Contract.JOE(&_TraderjoeChefV3.CallOpts)
}

// MASTERCHEFV2 is a free data retrieval call binding the contract method 0x27bf88ad.
//
// Solidity: function MASTER_CHEF_V2() view returns(address)
func (_TraderjoeChefV3 *TraderjoeChefV3Caller) MASTERCHEFV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeChefV3.contract.Call(opts, &out, "MASTER_CHEF_V2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MASTERCHEFV2 is a free data retrieval call binding the contract method 0x27bf88ad.
//
// Solidity: function MASTER_CHEF_V2() view returns(address)
func (_TraderjoeChefV3 *TraderjoeChefV3Session) MASTERCHEFV2() (common.Address, error) {
	return _TraderjoeChefV3.Contract.MASTERCHEFV2(&_TraderjoeChefV3.CallOpts)
}

// MASTERCHEFV2 is a free data retrieval call binding the contract method 0x27bf88ad.
//
// Solidity: function MASTER_CHEF_V2() view returns(address)
func (_TraderjoeChefV3 *TraderjoeChefV3CallerSession) MASTERCHEFV2() (common.Address, error) {
	return _TraderjoeChefV3.Contract.MASTERCHEFV2(&_TraderjoeChefV3.CallOpts)
}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_TraderjoeChefV3 *TraderjoeChefV3Caller) MASTERPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeChefV3.contract.Call(opts, &out, "MASTER_PID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_TraderjoeChefV3 *TraderjoeChefV3Session) MASTERPID() (*big.Int, error) {
	return _TraderjoeChefV3.Contract.MASTERPID(&_TraderjoeChefV3.CallOpts)
}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_TraderjoeChefV3 *TraderjoeChefV3CallerSession) MASTERPID() (*big.Int, error) {
	return _TraderjoeChefV3.Contract.MASTERPID(&_TraderjoeChefV3.CallOpts)
}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256 amount)
func (_TraderjoeChefV3 *TraderjoeChefV3Caller) JoePerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeChefV3.contract.Call(opts, &out, "joePerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256 amount)
func (_TraderjoeChefV3 *TraderjoeChefV3Session) JoePerSec() (*big.Int, error) {
	return _TraderjoeChefV3.Contract.JoePerSec(&_TraderjoeChefV3.CallOpts)
}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256 amount)
func (_TraderjoeChefV3 *TraderjoeChefV3CallerSession) JoePerSec() (*big.Int, error) {
	return _TraderjoeChefV3.Contract.JoePerSec(&_TraderjoeChefV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderjoeChefV3 *TraderjoeChefV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeChefV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderjoeChefV3 *TraderjoeChefV3Session) Owner() (common.Address, error) {
	return _TraderjoeChefV3.Contract.Owner(&_TraderjoeChefV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderjoeChefV3 *TraderjoeChefV3CallerSession) Owner() (common.Address, error) {
	return _TraderjoeChefV3.Contract.Owner(&_TraderjoeChefV3.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_TraderjoeChefV3 *TraderjoeChefV3Caller) PendingTokens(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	var out []interface{}
	err := _TraderjoeChefV3.contract.Call(opts, &out, "pendingTokens", _pid, _user)

	outstruct := new(struct {
		PendingJoe        *big.Int
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
		PendingBonusToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingJoe = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BonusTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PendingBonusToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_TraderjoeChefV3 *TraderjoeChefV3Session) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _TraderjoeChefV3.Contract.PendingTokens(&_TraderjoeChefV3.CallOpts, _pid, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_TraderjoeChefV3 *TraderjoeChefV3CallerSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _TraderjoeChefV3.Contract.PendingTokens(&_TraderjoeChefV3.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 accJoePerShare, uint256 lastRewardTimestamp, uint256 allocPoint, address rewarder)
func (_TraderjoeChefV3 *TraderjoeChefV3Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken             common.Address
	AccJoePerShare      *big.Int
	LastRewardTimestamp *big.Int
	AllocPoint          *big.Int
	Rewarder            common.Address
}, error) {
	var out []interface{}
	err := _TraderjoeChefV3.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken             common.Address
		AccJoePerShare      *big.Int
		LastRewardTimestamp *big.Int
		AllocPoint          *big.Int
		Rewarder            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AccJoePerShare = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AllocPoint = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Rewarder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 accJoePerShare, uint256 lastRewardTimestamp, uint256 allocPoint, address rewarder)
func (_TraderjoeChefV3 *TraderjoeChefV3Session) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AccJoePerShare      *big.Int
	LastRewardTimestamp *big.Int
	AllocPoint          *big.Int
	Rewarder            common.Address
}, error) {
	return _TraderjoeChefV3.Contract.PoolInfo(&_TraderjoeChefV3.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 accJoePerShare, uint256 lastRewardTimestamp, uint256 allocPoint, address rewarder)
func (_TraderjoeChefV3 *TraderjoeChefV3CallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AccJoePerShare      *big.Int
	LastRewardTimestamp *big.Int
	AllocPoint          *big.Int
	Rewarder            common.Address
}, error) {
	return _TraderjoeChefV3.Contract.PoolInfo(&_TraderjoeChefV3.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_TraderjoeChefV3 *TraderjoeChefV3Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeChefV3.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_TraderjoeChefV3 *TraderjoeChefV3Session) PoolLength() (*big.Int, error) {
	return _TraderjoeChefV3.Contract.PoolLength(&_TraderjoeChefV3.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_TraderjoeChefV3 *TraderjoeChefV3CallerSession) PoolLength() (*big.Int, error) {
	return _TraderjoeChefV3.Contract.PoolLength(&_TraderjoeChefV3.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_TraderjoeChefV3 *TraderjoeChefV3Caller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeChefV3.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_TraderjoeChefV3 *TraderjoeChefV3Session) TotalAllocPoint() (*big.Int, error) {
	return _TraderjoeChefV3.Contract.TotalAllocPoint(&_TraderjoeChefV3.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_TraderjoeChefV3 *TraderjoeChefV3CallerSession) TotalAllocPoint() (*big.Int, error) {
	return _TraderjoeChefV3.Contract.TotalAllocPoint(&_TraderjoeChefV3.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_TraderjoeChefV3 *TraderjoeChefV3Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _TraderjoeChefV3.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_TraderjoeChefV3 *TraderjoeChefV3Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _TraderjoeChefV3.Contract.UserInfo(&_TraderjoeChefV3.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_TraderjoeChefV3 *TraderjoeChefV3CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _TraderjoeChefV3.Contract.UserInfo(&_TraderjoeChefV3.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 allocPoint, address _lpToken, address _rewarder) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Transactor) Add(opts *bind.TransactOpts, allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _TraderjoeChefV3.contract.Transact(opts, "add", allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 allocPoint, address _lpToken, address _rewarder) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Session) Add(allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.Add(&_TraderjoeChefV3.TransactOpts, allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 allocPoint, address _lpToken, address _rewarder) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorSession) Add(allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.Add(&_TraderjoeChefV3.TransactOpts, allocPoint, _lpToken, _rewarder)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Transactor) Deposit(opts *bind.TransactOpts, pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.contract.Transact(opts, "deposit", pid, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Session) Deposit(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.Deposit(&_TraderjoeChefV3.TransactOpts, pid, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorSession) Deposit(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.Deposit(&_TraderjoeChefV3.TransactOpts, pid, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 pid) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Transactor) EmergencyWithdraw(opts *bind.TransactOpts, pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.contract.Transact(opts, "emergencyWithdraw", pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 pid) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Session) EmergencyWithdraw(pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.EmergencyWithdraw(&_TraderjoeChefV3.TransactOpts, pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 pid) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorSession) EmergencyWithdraw(pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.EmergencyWithdraw(&_TraderjoeChefV3.TransactOpts, pid)
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Transactor) HarvestFromMasterChef(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeChefV3.contract.Transact(opts, "harvestFromMasterChef")
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Session) HarvestFromMasterChef() (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.HarvestFromMasterChef(&_TraderjoeChefV3.TransactOpts)
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorSession) HarvestFromMasterChef() (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.HarvestFromMasterChef(&_TraderjoeChefV3.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Transactor) Init(opts *bind.TransactOpts, dummyToken common.Address) (*types.Transaction, error) {
	return _TraderjoeChefV3.contract.Transact(opts, "init", dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Session) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.Init(&_TraderjoeChefV3.TransactOpts, dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorSession) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.Init(&_TraderjoeChefV3.TransactOpts, dummyToken)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Transactor) MassUpdatePools(opts *bind.TransactOpts, pids []*big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.contract.Transact(opts, "massUpdatePools", pids)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Session) MassUpdatePools(pids []*big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.MassUpdatePools(&_TraderjoeChefV3.TransactOpts, pids)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorSession) MassUpdatePools(pids []*big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.MassUpdatePools(&_TraderjoeChefV3.TransactOpts, pids)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeChefV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.RenounceOwnership(&_TraderjoeChefV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.RenounceOwnership(&_TraderjoeChefV3.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Transactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _TraderjoeChefV3.contract.Transact(opts, "set", _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Session) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.Set(&_TraderjoeChefV3.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.Set(&_TraderjoeChefV3.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TraderjoeChefV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.TransferOwnership(&_TraderjoeChefV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.TransferOwnership(&_TraderjoeChefV3.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 pid) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Transactor) UpdatePool(opts *bind.TransactOpts, pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.contract.Transact(opts, "updatePool", pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 pid) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Session) UpdatePool(pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.UpdatePool(&_TraderjoeChefV3.TransactOpts, pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 pid) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorSession) UpdatePool(pid *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.UpdatePool(&_TraderjoeChefV3.TransactOpts, pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 pid, uint256 amount) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Transactor) Withdraw(opts *bind.TransactOpts, pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.contract.Transact(opts, "withdraw", pid, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 pid, uint256 amount) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3Session) Withdraw(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.Withdraw(&_TraderjoeChefV3.TransactOpts, pid, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 pid, uint256 amount) returns()
func (_TraderjoeChefV3 *TraderjoeChefV3TransactorSession) Withdraw(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeChefV3.Contract.Withdraw(&_TraderjoeChefV3.TransactOpts, pid, amount)
}
