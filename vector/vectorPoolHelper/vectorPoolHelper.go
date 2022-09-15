// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vectorPoolHelper

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

// VectorPoolHelperMetaData contains all meta data concerning the VectorPoolHelper contract.
var VectorPoolHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vtxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainStaking\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterVtx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingPTP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xptp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VectorPoolHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use VectorPoolHelperMetaData.ABI instead.
var VectorPoolHelperABI = VectorPoolHelperMetaData.ABI

// VectorPoolHelper is an auto generated Go binding around an Ethereum contract.
type VectorPoolHelper struct {
	VectorPoolHelperCaller     // Read-only binding to the contract
	VectorPoolHelperTransactor // Write-only binding to the contract
	VectorPoolHelperFilterer   // Log filterer for contract events
}

// VectorPoolHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type VectorPoolHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorPoolHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VectorPoolHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorPoolHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VectorPoolHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorPoolHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VectorPoolHelperSession struct {
	Contract     *VectorPoolHelper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VectorPoolHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VectorPoolHelperCallerSession struct {
	Contract *VectorPoolHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// VectorPoolHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VectorPoolHelperTransactorSession struct {
	Contract     *VectorPoolHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// VectorPoolHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type VectorPoolHelperRaw struct {
	Contract *VectorPoolHelper // Generic contract binding to access the raw methods on
}

// VectorPoolHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VectorPoolHelperCallerRaw struct {
	Contract *VectorPoolHelperCaller // Generic read-only contract binding to access the raw methods on
}

// VectorPoolHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VectorPoolHelperTransactorRaw struct {
	Contract *VectorPoolHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVectorPoolHelper creates a new instance of VectorPoolHelper, bound to a specific deployed contract.
func NewVectorPoolHelper(address common.Address, backend bind.ContractBackend) (*VectorPoolHelper, error) {
	contract, err := bindVectorPoolHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VectorPoolHelper{VectorPoolHelperCaller: VectorPoolHelperCaller{contract: contract}, VectorPoolHelperTransactor: VectorPoolHelperTransactor{contract: contract}, VectorPoolHelperFilterer: VectorPoolHelperFilterer{contract: contract}}, nil
}

// NewVectorPoolHelperCaller creates a new read-only instance of VectorPoolHelper, bound to a specific deployed contract.
func NewVectorPoolHelperCaller(address common.Address, caller bind.ContractCaller) (*VectorPoolHelperCaller, error) {
	contract, err := bindVectorPoolHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VectorPoolHelperCaller{contract: contract}, nil
}

// NewVectorPoolHelperTransactor creates a new write-only instance of VectorPoolHelper, bound to a specific deployed contract.
func NewVectorPoolHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*VectorPoolHelperTransactor, error) {
	contract, err := bindVectorPoolHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VectorPoolHelperTransactor{contract: contract}, nil
}

// NewVectorPoolHelperFilterer creates a new log filterer instance of VectorPoolHelper, bound to a specific deployed contract.
func NewVectorPoolHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*VectorPoolHelperFilterer, error) {
	contract, err := bindVectorPoolHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VectorPoolHelperFilterer{contract: contract}, nil
}

// bindVectorPoolHelper binds a generic wrapper to an already deployed contract.
func bindVectorPoolHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VectorPoolHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VectorPoolHelper *VectorPoolHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VectorPoolHelper.Contract.VectorPoolHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VectorPoolHelper *VectorPoolHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.VectorPoolHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VectorPoolHelper *VectorPoolHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.VectorPoolHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VectorPoolHelper *VectorPoolHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VectorPoolHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VectorPoolHelper *VectorPoolHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VectorPoolHelper *VectorPoolHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address _address) view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperCaller) Balance(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "balance", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address _address) view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperSession) Balance(_address common.Address) (*big.Int, error) {
	return _VectorPoolHelper.Contract.Balance(&_VectorPoolHelper.CallOpts, _address)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address _address) view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) Balance(_address common.Address) (*big.Int, error) {
	return _VectorPoolHelper.Contract.Balance(&_VectorPoolHelper.CallOpts, _address)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCaller) DepositToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "depositToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperSession) DepositToken() (common.Address, error) {
	return _VectorPoolHelper.Contract.DepositToken(&_VectorPoolHelper.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) DepositToken() (common.Address, error) {
	return _VectorPoolHelper.Contract.DepositToken(&_VectorPoolHelper.CallOpts)
}

// DepositTokenBalance is a free data retrieval call binding the contract method 0xcee7b7d2.
//
// Solidity: function depositTokenBalance() view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperCaller) DepositTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "depositTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositTokenBalance is a free data retrieval call binding the contract method 0xcee7b7d2.
//
// Solidity: function depositTokenBalance() view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperSession) DepositTokenBalance() (*big.Int, error) {
	return _VectorPoolHelper.Contract.DepositTokenBalance(&_VectorPoolHelper.CallOpts)
}

// DepositTokenBalance is a free data retrieval call binding the contract method 0xcee7b7d2.
//
// Solidity: function depositTokenBalance() view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) DepositTokenBalance() (*big.Int, error) {
	return _VectorPoolHelper.Contract.DepositTokenBalance(&_VectorPoolHelper.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address token) view returns(uint256 vtxAmount, uint256 tokenAmount)
func (_VectorPoolHelper *VectorPoolHelperCaller) Earned(opts *bind.CallOpts, token common.Address) (struct {
	VtxAmount   *big.Int
	TokenAmount *big.Int
}, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "earned", token)

	outstruct := new(struct {
		VtxAmount   *big.Int
		TokenAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VtxAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address token) view returns(uint256 vtxAmount, uint256 tokenAmount)
func (_VectorPoolHelper *VectorPoolHelperSession) Earned(token common.Address) (struct {
	VtxAmount   *big.Int
	TokenAmount *big.Int
}, error) {
	return _VectorPoolHelper.Contract.Earned(&_VectorPoolHelper.CallOpts, token)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address token) view returns(uint256 vtxAmount, uint256 tokenAmount)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) Earned(token common.Address) (struct {
	VtxAmount   *big.Int
	TokenAmount *big.Int
}, error) {
	return _VectorPoolHelper.Contract.Earned(&_VectorPoolHelper.CallOpts, token)
}

// MainStaking is a free data retrieval call binding the contract method 0x1208e15b.
//
// Solidity: function mainStaking() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCaller) MainStaking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "mainStaking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainStaking is a free data retrieval call binding the contract method 0x1208e15b.
//
// Solidity: function mainStaking() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperSession) MainStaking() (common.Address, error) {
	return _VectorPoolHelper.Contract.MainStaking(&_VectorPoolHelper.CallOpts)
}

// MainStaking is a free data retrieval call binding the contract method 0x1208e15b.
//
// Solidity: function mainStaking() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) MainStaking() (common.Address, error) {
	return _VectorPoolHelper.Contract.MainStaking(&_VectorPoolHelper.CallOpts)
}

// MasterVtx is a free data retrieval call binding the contract method 0xa296454f.
//
// Solidity: function masterVtx() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCaller) MasterVtx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "masterVtx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterVtx is a free data retrieval call binding the contract method 0xa296454f.
//
// Solidity: function masterVtx() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperSession) MasterVtx() (common.Address, error) {
	return _VectorPoolHelper.Contract.MasterVtx(&_VectorPoolHelper.CallOpts)
}

// MasterVtx is a free data retrieval call binding the contract method 0xa296454f.
//
// Solidity: function masterVtx() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) MasterVtx() (common.Address, error) {
	return _VectorPoolHelper.Contract.MasterVtx(&_VectorPoolHelper.CallOpts)
}

// PendingPTP is a free data retrieval call binding the contract method 0xe785312c.
//
// Solidity: function pendingPTP() view returns(uint256 pendingTokens)
func (_VectorPoolHelper *VectorPoolHelperCaller) PendingPTP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "pendingPTP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingPTP is a free data retrieval call binding the contract method 0xe785312c.
//
// Solidity: function pendingPTP() view returns(uint256 pendingTokens)
func (_VectorPoolHelper *VectorPoolHelperSession) PendingPTP() (*big.Int, error) {
	return _VectorPoolHelper.Contract.PendingPTP(&_VectorPoolHelper.CallOpts)
}

// PendingPTP is a free data retrieval call binding the contract method 0xe785312c.
//
// Solidity: function pendingPTP() view returns(uint256 pendingTokens)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) PendingPTP() (*big.Int, error) {
	return _VectorPoolHelper.Contract.PendingPTP(&_VectorPoolHelper.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperCaller) Pid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "pid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperSession) Pid() (*big.Int, error) {
	return _VectorPoolHelper.Contract.Pid(&_VectorPoolHelper.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) Pid() (*big.Int, error) {
	return _VectorPoolHelper.Contract.Pid(&_VectorPoolHelper.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address token) view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperCaller) RewardPerToken(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "rewardPerToken", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address token) view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperSession) RewardPerToken(token common.Address) (*big.Int, error) {
	return _VectorPoolHelper.Contract.RewardPerToken(&_VectorPoolHelper.CallOpts, token)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address token) view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) RewardPerToken(token common.Address) (*big.Int, error) {
	return _VectorPoolHelper.Contract.RewardPerToken(&_VectorPoolHelper.CallOpts, token)
}

// Rewarder is a free data retrieval call binding the contract method 0xdcc3e06e.
//
// Solidity: function rewarder() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCaller) Rewarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "rewarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewarder is a free data retrieval call binding the contract method 0xdcc3e06e.
//
// Solidity: function rewarder() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperSession) Rewarder() (common.Address, error) {
	return _VectorPoolHelper.Contract.Rewarder(&_VectorPoolHelper.CallOpts)
}

// Rewarder is a free data retrieval call binding the contract method 0xdcc3e06e.
//
// Solidity: function rewarder() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) Rewarder() (common.Address, error) {
	return _VectorPoolHelper.Contract.Rewarder(&_VectorPoolHelper.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperSession) StakingToken() (common.Address, error) {
	return _VectorPoolHelper.Contract.StakingToken(&_VectorPoolHelper.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) StakingToken() (common.Address, error) {
	return _VectorPoolHelper.Contract.StakingToken(&_VectorPoolHelper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperSession) TotalSupply() (*big.Int, error) {
	return _VectorPoolHelper.Contract.TotalSupply(&_VectorPoolHelper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) TotalSupply() (*big.Int, error) {
	return _VectorPoolHelper.Contract.TotalSupply(&_VectorPoolHelper.CallOpts)
}

// Xptp is a free data retrieval call binding the contract method 0x383d78e0.
//
// Solidity: function xptp() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCaller) Xptp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorPoolHelper.contract.Call(opts, &out, "xptp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Xptp is a free data retrieval call binding the contract method 0x383d78e0.
//
// Solidity: function xptp() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperSession) Xptp() (common.Address, error) {
	return _VectorPoolHelper.Contract.Xptp(&_VectorPoolHelper.CallOpts)
}

// Xptp is a free data retrieval call binding the contract method 0x383d78e0.
//
// Solidity: function xptp() view returns(address)
func (_VectorPoolHelper *VectorPoolHelperCallerSession) Xptp() (common.Address, error) {
	return _VectorPoolHelper.Contract.Xptp(&_VectorPoolHelper.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_VectorPoolHelper *VectorPoolHelperTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VectorPoolHelper.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_VectorPoolHelper *VectorPoolHelperSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.Deposit(&_VectorPoolHelper.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_VectorPoolHelper *VectorPoolHelperTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.Deposit(&_VectorPoolHelper.TransactOpts, amount)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_VectorPoolHelper *VectorPoolHelperTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorPoolHelper.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_VectorPoolHelper *VectorPoolHelperSession) GetReward() (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.GetReward(&_VectorPoolHelper.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_VectorPoolHelper *VectorPoolHelperTransactorSession) GetReward() (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.GetReward(&_VectorPoolHelper.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_VectorPoolHelper *VectorPoolHelperTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorPoolHelper.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_VectorPoolHelper *VectorPoolHelperSession) Harvest() (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.Harvest(&_VectorPoolHelper.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_VectorPoolHelper *VectorPoolHelperTransactorSession) Harvest() (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.Harvest(&_VectorPoolHelper.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_VectorPoolHelper *VectorPoolHelperTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _VectorPoolHelper.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_VectorPoolHelper *VectorPoolHelperSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.Stake(&_VectorPoolHelper.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_VectorPoolHelper *VectorPoolHelperTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.Stake(&_VectorPoolHelper.TransactOpts, _amount)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_VectorPoolHelper *VectorPoolHelperTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorPoolHelper.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_VectorPoolHelper *VectorPoolHelperSession) Update() (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.Update(&_VectorPoolHelper.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_VectorPoolHelper *VectorPoolHelperTransactorSession) Update() (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.Update(&_VectorPoolHelper.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 amount, uint256 minAmount) returns()
func (_VectorPoolHelper *VectorPoolHelperTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _VectorPoolHelper.contract.Transact(opts, "withdraw", amount, minAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 amount, uint256 minAmount) returns()
func (_VectorPoolHelper *VectorPoolHelperSession) Withdraw(amount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.Withdraw(&_VectorPoolHelper.TransactOpts, amount, minAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 amount, uint256 minAmount) returns()
func (_VectorPoolHelper *VectorPoolHelperTransactorSession) Withdraw(amount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _VectorPoolHelper.Contract.Withdraw(&_VectorPoolHelper.TransactOpts, amount, minAmount)
}
