// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package axialMasterChef

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

// AxialMasterChefMetaData contains all meta data concerning the AxialMasterChef contract.
var AxialMasterChefMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"axialPerSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingAxial\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pendingBonusToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"accAxialPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewarder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pools\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AxialMasterChefABI is the input ABI used to generate the binding from.
// Deprecated: Use AxialMasterChefMetaData.ABI instead.
var AxialMasterChefABI = AxialMasterChefMetaData.ABI

// AxialMasterChef is an auto generated Go binding around an Ethereum contract.
type AxialMasterChef struct {
	AxialMasterChefCaller     // Read-only binding to the contract
	AxialMasterChefTransactor // Write-only binding to the contract
	AxialMasterChefFilterer   // Log filterer for contract events
}

// AxialMasterChefCaller is an auto generated read-only Go binding around an Ethereum contract.
type AxialMasterChefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxialMasterChefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AxialMasterChefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxialMasterChefFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AxialMasterChefFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxialMasterChefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AxialMasterChefSession struct {
	Contract     *AxialMasterChef  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AxialMasterChefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AxialMasterChefCallerSession struct {
	Contract *AxialMasterChefCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AxialMasterChefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AxialMasterChefTransactorSession struct {
	Contract     *AxialMasterChefTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AxialMasterChefRaw is an auto generated low-level Go binding around an Ethereum contract.
type AxialMasterChefRaw struct {
	Contract *AxialMasterChef // Generic contract binding to access the raw methods on
}

// AxialMasterChefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AxialMasterChefCallerRaw struct {
	Contract *AxialMasterChefCaller // Generic read-only contract binding to access the raw methods on
}

// AxialMasterChefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AxialMasterChefTransactorRaw struct {
	Contract *AxialMasterChefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAxialMasterChef creates a new instance of AxialMasterChef, bound to a specific deployed contract.
func NewAxialMasterChef(address common.Address, backend bind.ContractBackend) (*AxialMasterChef, error) {
	contract, err := bindAxialMasterChef(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AxialMasterChef{AxialMasterChefCaller: AxialMasterChefCaller{contract: contract}, AxialMasterChefTransactor: AxialMasterChefTransactor{contract: contract}, AxialMasterChefFilterer: AxialMasterChefFilterer{contract: contract}}, nil
}

// NewAxialMasterChefCaller creates a new read-only instance of AxialMasterChef, bound to a specific deployed contract.
func NewAxialMasterChefCaller(address common.Address, caller bind.ContractCaller) (*AxialMasterChefCaller, error) {
	contract, err := bindAxialMasterChef(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AxialMasterChefCaller{contract: contract}, nil
}

// NewAxialMasterChefTransactor creates a new write-only instance of AxialMasterChef, bound to a specific deployed contract.
func NewAxialMasterChefTransactor(address common.Address, transactor bind.ContractTransactor) (*AxialMasterChefTransactor, error) {
	contract, err := bindAxialMasterChef(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AxialMasterChefTransactor{contract: contract}, nil
}

// NewAxialMasterChefFilterer creates a new log filterer instance of AxialMasterChef, bound to a specific deployed contract.
func NewAxialMasterChefFilterer(address common.Address, filterer bind.ContractFilterer) (*AxialMasterChefFilterer, error) {
	contract, err := bindAxialMasterChef(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AxialMasterChefFilterer{contract: contract}, nil
}

// bindAxialMasterChef binds a generic wrapper to an already deployed contract.
func bindAxialMasterChef(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AxialMasterChefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AxialMasterChef *AxialMasterChefRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AxialMasterChef.Contract.AxialMasterChefCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AxialMasterChef *AxialMasterChefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AxialMasterChef.Contract.AxialMasterChefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AxialMasterChef *AxialMasterChefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AxialMasterChef.Contract.AxialMasterChefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AxialMasterChef *AxialMasterChefCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AxialMasterChef.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AxialMasterChef *AxialMasterChefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AxialMasterChef.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AxialMasterChef *AxialMasterChefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AxialMasterChef.Contract.contract.Transact(opts, method, params...)
}

// AxialPerSec is a free data retrieval call binding the contract method 0x194ed6af.
//
// Solidity: function axialPerSec() view returns(uint256 amount)
func (_AxialMasterChef *AxialMasterChefCaller) AxialPerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AxialMasterChef.contract.Call(opts, &out, "axialPerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AxialPerSec is a free data retrieval call binding the contract method 0x194ed6af.
//
// Solidity: function axialPerSec() view returns(uint256 amount)
func (_AxialMasterChef *AxialMasterChefSession) AxialPerSec() (*big.Int, error) {
	return _AxialMasterChef.Contract.AxialPerSec(&_AxialMasterChef.CallOpts)
}

// AxialPerSec is a free data retrieval call binding the contract method 0x194ed6af.
//
// Solidity: function axialPerSec() view returns(uint256 amount)
func (_AxialMasterChef *AxialMasterChefCallerSession) AxialPerSec() (*big.Int, error) {
	return _AxialMasterChef.Contract.AxialPerSec(&_AxialMasterChef.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 pid, address user) view returns(uint256 pendingAxial, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_AxialMasterChef *AxialMasterChefCaller) PendingTokens(opts *bind.CallOpts, pid *big.Int, user common.Address) (struct {
	PendingAxial      *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	var out []interface{}
	err := _AxialMasterChef.contract.Call(opts, &out, "pendingTokens", pid, user)

	outstruct := new(struct {
		PendingAxial      *big.Int
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
		PendingBonusToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingAxial = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BonusTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PendingBonusToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 pid, address user) view returns(uint256 pendingAxial, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_AxialMasterChef *AxialMasterChefSession) PendingTokens(pid *big.Int, user common.Address) (struct {
	PendingAxial      *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _AxialMasterChef.Contract.PendingTokens(&_AxialMasterChef.CallOpts, pid, user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 pid, address user) view returns(uint256 pendingAxial, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_AxialMasterChef *AxialMasterChefCallerSession) PendingTokens(pid *big.Int, user common.Address) (struct {
	PendingAxial      *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _AxialMasterChef.Contract.PendingTokens(&_AxialMasterChef.CallOpts, pid, user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 accAxialPerShare, uint256 lastRewardTimestamp, uint256 allocPoint, address rewarder)
func (_AxialMasterChef *AxialMasterChefCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken             common.Address
	AccAxialPerShare    *big.Int
	LastRewardTimestamp *big.Int
	AllocPoint          *big.Int
	Rewarder            common.Address
}, error) {
	var out []interface{}
	err := _AxialMasterChef.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken             common.Address
		AccAxialPerShare    *big.Int
		LastRewardTimestamp *big.Int
		AllocPoint          *big.Int
		Rewarder            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AccAxialPerShare = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AllocPoint = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Rewarder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 accAxialPerShare, uint256 lastRewardTimestamp, uint256 allocPoint, address rewarder)
func (_AxialMasterChef *AxialMasterChefSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AccAxialPerShare    *big.Int
	LastRewardTimestamp *big.Int
	AllocPoint          *big.Int
	Rewarder            common.Address
}, error) {
	return _AxialMasterChef.Contract.PoolInfo(&_AxialMasterChef.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 accAxialPerShare, uint256 lastRewardTimestamp, uint256 allocPoint, address rewarder)
func (_AxialMasterChef *AxialMasterChefCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AccAxialPerShare    *big.Int
	LastRewardTimestamp *big.Int
	AllocPoint          *big.Int
	Rewarder            common.Address
}, error) {
	return _AxialMasterChef.Contract.PoolInfo(&_AxialMasterChef.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_AxialMasterChef *AxialMasterChefCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AxialMasterChef.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_AxialMasterChef *AxialMasterChefSession) PoolLength() (*big.Int, error) {
	return _AxialMasterChef.Contract.PoolLength(&_AxialMasterChef.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_AxialMasterChef *AxialMasterChefCallerSession) PoolLength() (*big.Int, error) {
	return _AxialMasterChef.Contract.PoolLength(&_AxialMasterChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_AxialMasterChef *AxialMasterChefCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AxialMasterChef.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_AxialMasterChef *AxialMasterChefSession) TotalAllocPoint() (*big.Int, error) {
	return _AxialMasterChef.Contract.TotalAllocPoint(&_AxialMasterChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_AxialMasterChef *AxialMasterChefCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _AxialMasterChef.Contract.TotalAllocPoint(&_AxialMasterChef.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_AxialMasterChef *AxialMasterChefCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _AxialMasterChef.contract.Call(opts, &out, "userInfo", arg0, arg1)

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
func (_AxialMasterChef *AxialMasterChefSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _AxialMasterChef.Contract.UserInfo(&_AxialMasterChef.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_AxialMasterChef *AxialMasterChefCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _AxialMasterChef.Contract.UserInfo(&_AxialMasterChef.CallOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) returns()
func (_AxialMasterChef *AxialMasterChefTransactor) Deposit(opts *bind.TransactOpts, pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AxialMasterChef.contract.Transact(opts, "deposit", pid, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) returns()
func (_AxialMasterChef *AxialMasterChefSession) Deposit(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AxialMasterChef.Contract.Deposit(&_AxialMasterChef.TransactOpts, pid, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) returns()
func (_AxialMasterChef *AxialMasterChefTransactorSession) Deposit(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AxialMasterChef.Contract.Deposit(&_AxialMasterChef.TransactOpts, pid, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 pid, uint256 amount) returns()
func (_AxialMasterChef *AxialMasterChefTransactor) Withdraw(opts *bind.TransactOpts, pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AxialMasterChef.contract.Transact(opts, "withdraw", pid, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 pid, uint256 amount) returns()
func (_AxialMasterChef *AxialMasterChefSession) Withdraw(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AxialMasterChef.Contract.Withdraw(&_AxialMasterChef.TransactOpts, pid, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 pid, uint256 amount) returns()
func (_AxialMasterChef *AxialMasterChefTransactorSession) Withdraw(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _AxialMasterChef.Contract.Withdraw(&_AxialMasterChef.TransactOpts, pid, amount)
}
