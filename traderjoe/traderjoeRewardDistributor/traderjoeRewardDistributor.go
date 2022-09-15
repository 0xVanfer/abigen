// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderjoeRewardDistributor

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

// TraderjoeRewardDistributorMetaData contains all meta data concerning the TraderjoeRewardDistributor contract.
var TraderjoeRewardDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rewardType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_grantReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rewardType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardSupplySpeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBorrowSpeed\",\"type\":\"uint256\"}],\"name\":\"_setRewardSpeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rewardType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rewardType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"jTokens\",\"type\":\"address[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rewardType\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"jTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joetroller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardBorrowSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardSupplySpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newJoeAddress\",\"type\":\"address\"}],\"name\":\"setJoeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_joetroller\",\"type\":\"address\"}],\"name\":\"setJoetroller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"}],\"name\":\"updateAndDistributeSupplierRewardsForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TraderjoeRewardDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderjoeRewardDistributorMetaData.ABI instead.
var TraderjoeRewardDistributorABI = TraderjoeRewardDistributorMetaData.ABI

// TraderjoeRewardDistributor is an auto generated Go binding around an Ethereum contract.
type TraderjoeRewardDistributor struct {
	TraderjoeRewardDistributorCaller     // Read-only binding to the contract
	TraderjoeRewardDistributorTransactor // Write-only binding to the contract
	TraderjoeRewardDistributorFilterer   // Log filterer for contract events
}

// TraderjoeRewardDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderjoeRewardDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeRewardDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderjoeRewardDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeRewardDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderjoeRewardDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeRewardDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderjoeRewardDistributorSession struct {
	Contract     *TraderjoeRewardDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TraderjoeRewardDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderjoeRewardDistributorCallerSession struct {
	Contract *TraderjoeRewardDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// TraderjoeRewardDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderjoeRewardDistributorTransactorSession struct {
	Contract     *TraderjoeRewardDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// TraderjoeRewardDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderjoeRewardDistributorRaw struct {
	Contract *TraderjoeRewardDistributor // Generic contract binding to access the raw methods on
}

// TraderjoeRewardDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderjoeRewardDistributorCallerRaw struct {
	Contract *TraderjoeRewardDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// TraderjoeRewardDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderjoeRewardDistributorTransactorRaw struct {
	Contract *TraderjoeRewardDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderjoeRewardDistributor creates a new instance of TraderjoeRewardDistributor, bound to a specific deployed contract.
func NewTraderjoeRewardDistributor(address common.Address, backend bind.ContractBackend) (*TraderjoeRewardDistributor, error) {
	contract, err := bindTraderjoeRewardDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraderjoeRewardDistributor{TraderjoeRewardDistributorCaller: TraderjoeRewardDistributorCaller{contract: contract}, TraderjoeRewardDistributorTransactor: TraderjoeRewardDistributorTransactor{contract: contract}, TraderjoeRewardDistributorFilterer: TraderjoeRewardDistributorFilterer{contract: contract}}, nil
}

// NewTraderjoeRewardDistributorCaller creates a new read-only instance of TraderjoeRewardDistributor, bound to a specific deployed contract.
func NewTraderjoeRewardDistributorCaller(address common.Address, caller bind.ContractCaller) (*TraderjoeRewardDistributorCaller, error) {
	contract, err := bindTraderjoeRewardDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeRewardDistributorCaller{contract: contract}, nil
}

// NewTraderjoeRewardDistributorTransactor creates a new write-only instance of TraderjoeRewardDistributor, bound to a specific deployed contract.
func NewTraderjoeRewardDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderjoeRewardDistributorTransactor, error) {
	contract, err := bindTraderjoeRewardDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeRewardDistributorTransactor{contract: contract}, nil
}

// NewTraderjoeRewardDistributorFilterer creates a new log filterer instance of TraderjoeRewardDistributor, bound to a specific deployed contract.
func NewTraderjoeRewardDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderjoeRewardDistributorFilterer, error) {
	contract, err := bindTraderjoeRewardDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderjoeRewardDistributorFilterer{contract: contract}, nil
}

// bindTraderjoeRewardDistributor binds a generic wrapper to an already deployed contract.
func bindTraderjoeRewardDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraderjoeRewardDistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeRewardDistributor.Contract.TraderjoeRewardDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.TraderjoeRewardDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.TraderjoeRewardDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderjoeRewardDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) Admin() (common.Address, error) {
	return _TraderjoeRewardDistributor.Contract.Admin(&_TraderjoeRewardDistributor.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) Admin() (common.Address, error) {
	return _TraderjoeRewardDistributor.Contract.Admin(&_TraderjoeRewardDistributor.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) GetBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "getBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) GetBlockTimestamp() (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.GetBlockTimestamp(&_TraderjoeRewardDistributor.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) GetBlockTimestamp() (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.GetBlockTimestamp(&_TraderjoeRewardDistributor.CallOpts)
}

// JoeAddress is a free data retrieval call binding the contract method 0xcb15d8e2.
//
// Solidity: function joeAddress() view returns(address)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) JoeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "joeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JoeAddress is a free data retrieval call binding the contract method 0xcb15d8e2.
//
// Solidity: function joeAddress() view returns(address)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) JoeAddress() (common.Address, error) {
	return _TraderjoeRewardDistributor.Contract.JoeAddress(&_TraderjoeRewardDistributor.CallOpts)
}

// JoeAddress is a free data retrieval call binding the contract method 0xcb15d8e2.
//
// Solidity: function joeAddress() view returns(address)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) JoeAddress() (common.Address, error) {
	return _TraderjoeRewardDistributor.Contract.JoeAddress(&_TraderjoeRewardDistributor.CallOpts)
}

// Joetroller is a free data retrieval call binding the contract method 0x6330533c.
//
// Solidity: function joetroller() view returns(address)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) Joetroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "joetroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Joetroller is a free data retrieval call binding the contract method 0x6330533c.
//
// Solidity: function joetroller() view returns(address)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) Joetroller() (common.Address, error) {
	return _TraderjoeRewardDistributor.Contract.Joetroller(&_TraderjoeRewardDistributor.CallOpts)
}

// Joetroller is a free data retrieval call binding the contract method 0x6330533c.
//
// Solidity: function joetroller() view returns(address)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) Joetroller() (common.Address, error) {
	return _TraderjoeRewardDistributor.Contract.Joetroller(&_TraderjoeRewardDistributor.CallOpts)
}

// RewardAccrued is a free data retrieval call binding the contract method 0x05b9783d.
//
// Solidity: function rewardAccrued(uint8 , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) RewardAccrued(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "rewardAccrued", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardAccrued is a free data retrieval call binding the contract method 0x05b9783d.
//
// Solidity: function rewardAccrued(uint8 , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) RewardAccrued(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardAccrued(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1)
}

// RewardAccrued is a free data retrieval call binding the contract method 0x05b9783d.
//
// Solidity: function rewardAccrued(uint8 , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) RewardAccrued(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardAccrued(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1)
}

// RewardBorrowSpeeds is a free data retrieval call binding the contract method 0xbf095955.
//
// Solidity: function rewardBorrowSpeeds(uint8 , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) RewardBorrowSpeeds(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "rewardBorrowSpeeds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardBorrowSpeeds is a free data retrieval call binding the contract method 0xbf095955.
//
// Solidity: function rewardBorrowSpeeds(uint8 , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) RewardBorrowSpeeds(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardBorrowSpeeds(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1)
}

// RewardBorrowSpeeds is a free data retrieval call binding the contract method 0xbf095955.
//
// Solidity: function rewardBorrowSpeeds(uint8 , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) RewardBorrowSpeeds(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardBorrowSpeeds(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1)
}

// RewardBorrowState is a free data retrieval call binding the contract method 0x4b3a0a74.
//
// Solidity: function rewardBorrowState(uint8 , address ) view returns(uint224 index, uint32 timestamp)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) RewardBorrowState(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "rewardBorrowState", arg0, arg1)

	outstruct := new(struct {
		Index     *big.Int
		Timestamp uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// RewardBorrowState is a free data retrieval call binding the contract method 0x4b3a0a74.
//
// Solidity: function rewardBorrowState(uint8 , address ) view returns(uint224 index, uint32 timestamp)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) RewardBorrowState(arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	return _TraderjoeRewardDistributor.Contract.RewardBorrowState(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1)
}

// RewardBorrowState is a free data retrieval call binding the contract method 0x4b3a0a74.
//
// Solidity: function rewardBorrowState(uint8 , address ) view returns(uint224 index, uint32 timestamp)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) RewardBorrowState(arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	return _TraderjoeRewardDistributor.Contract.RewardBorrowState(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1)
}

// RewardBorrowerIndex is a free data retrieval call binding the contract method 0x7937969d.
//
// Solidity: function rewardBorrowerIndex(uint8 , address , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) RewardBorrowerIndex(opts *bind.CallOpts, arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "rewardBorrowerIndex", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardBorrowerIndex is a free data retrieval call binding the contract method 0x7937969d.
//
// Solidity: function rewardBorrowerIndex(uint8 , address , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) RewardBorrowerIndex(arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardBorrowerIndex(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1, arg2)
}

// RewardBorrowerIndex is a free data retrieval call binding the contract method 0x7937969d.
//
// Solidity: function rewardBorrowerIndex(uint8 , address , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) RewardBorrowerIndex(arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardBorrowerIndex(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1, arg2)
}

// RewardInitialIndex is a free data retrieval call binding the contract method 0x66f91a24.
//
// Solidity: function rewardInitialIndex() view returns(uint224)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) RewardInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "rewardInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardInitialIndex is a free data retrieval call binding the contract method 0x66f91a24.
//
// Solidity: function rewardInitialIndex() view returns(uint224)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) RewardInitialIndex() (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardInitialIndex(&_TraderjoeRewardDistributor.CallOpts)
}

// RewardInitialIndex is a free data retrieval call binding the contract method 0x66f91a24.
//
// Solidity: function rewardInitialIndex() view returns(uint224)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) RewardInitialIndex() (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardInitialIndex(&_TraderjoeRewardDistributor.CallOpts)
}

// RewardSupplierIndex is a free data retrieval call binding the contract method 0x88e972b8.
//
// Solidity: function rewardSupplierIndex(uint8 , address , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) RewardSupplierIndex(opts *bind.CallOpts, arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "rewardSupplierIndex", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardSupplierIndex is a free data retrieval call binding the contract method 0x88e972b8.
//
// Solidity: function rewardSupplierIndex(uint8 , address , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) RewardSupplierIndex(arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardSupplierIndex(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1, arg2)
}

// RewardSupplierIndex is a free data retrieval call binding the contract method 0x88e972b8.
//
// Solidity: function rewardSupplierIndex(uint8 , address , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) RewardSupplierIndex(arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardSupplierIndex(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1, arg2)
}

// RewardSupplySpeeds is a free data retrieval call binding the contract method 0x030ce638.
//
// Solidity: function rewardSupplySpeeds(uint8 , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) RewardSupplySpeeds(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "rewardSupplySpeeds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardSupplySpeeds is a free data retrieval call binding the contract method 0x030ce638.
//
// Solidity: function rewardSupplySpeeds(uint8 , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) RewardSupplySpeeds(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardSupplySpeeds(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1)
}

// RewardSupplySpeeds is a free data retrieval call binding the contract method 0x030ce638.
//
// Solidity: function rewardSupplySpeeds(uint8 , address ) view returns(uint256)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) RewardSupplySpeeds(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _TraderjoeRewardDistributor.Contract.RewardSupplySpeeds(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1)
}

// RewardSupplyState is a free data retrieval call binding the contract method 0xd81c5e45.
//
// Solidity: function rewardSupplyState(uint8 , address ) view returns(uint224 index, uint32 timestamp)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCaller) RewardSupplyState(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	var out []interface{}
	err := _TraderjoeRewardDistributor.contract.Call(opts, &out, "rewardSupplyState", arg0, arg1)

	outstruct := new(struct {
		Index     *big.Int
		Timestamp uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// RewardSupplyState is a free data retrieval call binding the contract method 0xd81c5e45.
//
// Solidity: function rewardSupplyState(uint8 , address ) view returns(uint224 index, uint32 timestamp)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) RewardSupplyState(arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	return _TraderjoeRewardDistributor.Contract.RewardSupplyState(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1)
}

// RewardSupplyState is a free data retrieval call binding the contract method 0xd81c5e45.
//
// Solidity: function rewardSupplyState(uint8 , address ) view returns(uint224 index, uint32 timestamp)
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorCallerSession) RewardSupplyState(arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	return _TraderjoeRewardDistributor.Contract.RewardSupplyState(&_TraderjoeRewardDistributor.CallOpts, arg0, arg1)
}

// GrantReward is a paid mutator transaction binding the contract method 0x2f32b11f.
//
// Solidity: function _grantReward(uint8 rewardType, address recipient, uint256 amount) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactor) GrantReward(opts *bind.TransactOpts, rewardType uint8, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.contract.Transact(opts, "_grantReward", rewardType, recipient, amount)
}

// GrantReward is a paid mutator transaction binding the contract method 0x2f32b11f.
//
// Solidity: function _grantReward(uint8 rewardType, address recipient, uint256 amount) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) GrantReward(rewardType uint8, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.GrantReward(&_TraderjoeRewardDistributor.TransactOpts, rewardType, recipient, amount)
}

// GrantReward is a paid mutator transaction binding the contract method 0x2f32b11f.
//
// Solidity: function _grantReward(uint8 rewardType, address recipient, uint256 amount) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorSession) GrantReward(rewardType uint8, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.GrantReward(&_TraderjoeRewardDistributor.TransactOpts, rewardType, recipient, amount)
}

// SetRewardSpeed is a paid mutator transaction binding the contract method 0x1d94cb94.
//
// Solidity: function _setRewardSpeed(uint8 rewardType, address jToken, uint256 rewardSupplySpeed, uint256 rewardBorrowSpeed) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactor) SetRewardSpeed(opts *bind.TransactOpts, rewardType uint8, jToken common.Address, rewardSupplySpeed *big.Int, rewardBorrowSpeed *big.Int) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.contract.Transact(opts, "_setRewardSpeed", rewardType, jToken, rewardSupplySpeed, rewardBorrowSpeed)
}

// SetRewardSpeed is a paid mutator transaction binding the contract method 0x1d94cb94.
//
// Solidity: function _setRewardSpeed(uint8 rewardType, address jToken, uint256 rewardSupplySpeed, uint256 rewardBorrowSpeed) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) SetRewardSpeed(rewardType uint8, jToken common.Address, rewardSupplySpeed *big.Int, rewardBorrowSpeed *big.Int) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.SetRewardSpeed(&_TraderjoeRewardDistributor.TransactOpts, rewardType, jToken, rewardSupplySpeed, rewardBorrowSpeed)
}

// SetRewardSpeed is a paid mutator transaction binding the contract method 0x1d94cb94.
//
// Solidity: function _setRewardSpeed(uint8 rewardType, address jToken, uint256 rewardSupplySpeed, uint256 rewardBorrowSpeed) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorSession) SetRewardSpeed(rewardType uint8, jToken common.Address, rewardSupplySpeed *big.Int, rewardBorrowSpeed *big.Int) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.SetRewardSpeed(&_TraderjoeRewardDistributor.TransactOpts, rewardType, jToken, rewardSupplySpeed, rewardBorrowSpeed)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0952c563.
//
// Solidity: function claimReward(uint8 rewardType, address holder) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactor) ClaimReward(opts *bind.TransactOpts, rewardType uint8, holder common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.contract.Transact(opts, "claimReward", rewardType, holder)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0952c563.
//
// Solidity: function claimReward(uint8 rewardType, address holder) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) ClaimReward(rewardType uint8, holder common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.ClaimReward(&_TraderjoeRewardDistributor.TransactOpts, rewardType, holder)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0952c563.
//
// Solidity: function claimReward(uint8 rewardType, address holder) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorSession) ClaimReward(rewardType uint8, holder common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.ClaimReward(&_TraderjoeRewardDistributor.TransactOpts, rewardType, holder)
}

// ClaimReward0 is a paid mutator transaction binding the contract method 0x744532ae.
//
// Solidity: function claimReward(uint8 rewardType, address holder, address[] jTokens) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactor) ClaimReward0(opts *bind.TransactOpts, rewardType uint8, holder common.Address, jTokens []common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.contract.Transact(opts, "claimReward0", rewardType, holder, jTokens)
}

// ClaimReward0 is a paid mutator transaction binding the contract method 0x744532ae.
//
// Solidity: function claimReward(uint8 rewardType, address holder, address[] jTokens) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) ClaimReward0(rewardType uint8, holder common.Address, jTokens []common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.ClaimReward0(&_TraderjoeRewardDistributor.TransactOpts, rewardType, holder, jTokens)
}

// ClaimReward0 is a paid mutator transaction binding the contract method 0x744532ae.
//
// Solidity: function claimReward(uint8 rewardType, address holder, address[] jTokens) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorSession) ClaimReward0(rewardType uint8, holder common.Address, jTokens []common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.ClaimReward0(&_TraderjoeRewardDistributor.TransactOpts, rewardType, holder, jTokens)
}

// ClaimReward1 is a paid mutator transaction binding the contract method 0x8805714b.
//
// Solidity: function claimReward(uint8 rewardType, address[] holders, address[] jTokens, bool borrowers, bool suppliers) payable returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactor) ClaimReward1(opts *bind.TransactOpts, rewardType uint8, holders []common.Address, jTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.contract.Transact(opts, "claimReward1", rewardType, holders, jTokens, borrowers, suppliers)
}

// ClaimReward1 is a paid mutator transaction binding the contract method 0x8805714b.
//
// Solidity: function claimReward(uint8 rewardType, address[] holders, address[] jTokens, bool borrowers, bool suppliers) payable returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) ClaimReward1(rewardType uint8, holders []common.Address, jTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.ClaimReward1(&_TraderjoeRewardDistributor.TransactOpts, rewardType, holders, jTokens, borrowers, suppliers)
}

// ClaimReward1 is a paid mutator transaction binding the contract method 0x8805714b.
//
// Solidity: function claimReward(uint8 rewardType, address[] holders, address[] jTokens, bool borrowers, bool suppliers) payable returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorSession) ClaimReward1(rewardType uint8, holders []common.Address, jTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.ClaimReward1(&_TraderjoeRewardDistributor.TransactOpts, rewardType, holders, jTokens, borrowers, suppliers)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) Initialize() (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.Initialize(&_TraderjoeRewardDistributor.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorSession) Initialize() (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.Initialize(&_TraderjoeRewardDistributor.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.SetAdmin(&_TraderjoeRewardDistributor.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.SetAdmin(&_TraderjoeRewardDistributor.TransactOpts, _newAdmin)
}

// SetJoeAddress is a paid mutator transaction binding the contract method 0xfcc4cfd1.
//
// Solidity: function setJoeAddress(address newJoeAddress) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactor) SetJoeAddress(opts *bind.TransactOpts, newJoeAddress common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.contract.Transact(opts, "setJoeAddress", newJoeAddress)
}

// SetJoeAddress is a paid mutator transaction binding the contract method 0xfcc4cfd1.
//
// Solidity: function setJoeAddress(address newJoeAddress) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) SetJoeAddress(newJoeAddress common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.SetJoeAddress(&_TraderjoeRewardDistributor.TransactOpts, newJoeAddress)
}

// SetJoeAddress is a paid mutator transaction binding the contract method 0xfcc4cfd1.
//
// Solidity: function setJoeAddress(address newJoeAddress) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorSession) SetJoeAddress(newJoeAddress common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.SetJoeAddress(&_TraderjoeRewardDistributor.TransactOpts, newJoeAddress)
}

// SetJoetroller is a paid mutator transaction binding the contract method 0x1f9f3511.
//
// Solidity: function setJoetroller(address _joetroller) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactor) SetJoetroller(opts *bind.TransactOpts, _joetroller common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.contract.Transact(opts, "setJoetroller", _joetroller)
}

// SetJoetroller is a paid mutator transaction binding the contract method 0x1f9f3511.
//
// Solidity: function setJoetroller(address _joetroller) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) SetJoetroller(_joetroller common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.SetJoetroller(&_TraderjoeRewardDistributor.TransactOpts, _joetroller)
}

// SetJoetroller is a paid mutator transaction binding the contract method 0x1f9f3511.
//
// Solidity: function setJoetroller(address _joetroller) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorSession) SetJoetroller(_joetroller common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.SetJoetroller(&_TraderjoeRewardDistributor.TransactOpts, _joetroller)
}

// UpdateAndDistributeSupplierRewardsForToken is a paid mutator transaction binding the contract method 0x670e8afd.
//
// Solidity: function updateAndDistributeSupplierRewardsForToken(address jToken, address supplier) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactor) UpdateAndDistributeSupplierRewardsForToken(opts *bind.TransactOpts, jToken common.Address, supplier common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.contract.Transact(opts, "updateAndDistributeSupplierRewardsForToken", jToken, supplier)
}

// UpdateAndDistributeSupplierRewardsForToken is a paid mutator transaction binding the contract method 0x670e8afd.
//
// Solidity: function updateAndDistributeSupplierRewardsForToken(address jToken, address supplier) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorSession) UpdateAndDistributeSupplierRewardsForToken(jToken common.Address, supplier common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.UpdateAndDistributeSupplierRewardsForToken(&_TraderjoeRewardDistributor.TransactOpts, jToken, supplier)
}

// UpdateAndDistributeSupplierRewardsForToken is a paid mutator transaction binding the contract method 0x670e8afd.
//
// Solidity: function updateAndDistributeSupplierRewardsForToken(address jToken, address supplier) returns()
func (_TraderjoeRewardDistributor *TraderjoeRewardDistributorTransactorSession) UpdateAndDistributeSupplierRewardsForToken(jToken common.Address, supplier common.Address) (*types.Transaction, error) {
	return _TraderjoeRewardDistributor.Contract.UpdateAndDistributeSupplierRewardsForToken(&_TraderjoeRewardDistributor.TransactOpts, jToken, supplier)
}
