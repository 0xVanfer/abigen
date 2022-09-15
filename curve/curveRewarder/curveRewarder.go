// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveRewarder

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

// CurveRewarderMetaData contains all meta data concerning the CurveRewarder contract.
var CurveRewarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"accept_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"add_reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_update_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"notify_reward_amount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"remove_reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward_count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"reward_data\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"period_finish\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward_receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"reward_tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"set_receiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"set_reward_distributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"set_reward_duration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CurveRewarderABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveRewarderMetaData.ABI instead.
var CurveRewarderABI = CurveRewarderMetaData.ABI

// CurveRewarder is an auto generated Go binding around an Ethereum contract.
type CurveRewarder struct {
	CurveRewarderCaller     // Read-only binding to the contract
	CurveRewarderTransactor // Write-only binding to the contract
	CurveRewarderFilterer   // Log filterer for contract events
}

// CurveRewarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveRewarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRewarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveRewarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRewarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveRewarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveRewarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveRewarderSession struct {
	Contract     *CurveRewarder    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveRewarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveRewarderCallerSession struct {
	Contract *CurveRewarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CurveRewarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveRewarderTransactorSession struct {
	Contract     *CurveRewarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CurveRewarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveRewarderRaw struct {
	Contract *CurveRewarder // Generic contract binding to access the raw methods on
}

// CurveRewarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveRewarderCallerRaw struct {
	Contract *CurveRewarderCaller // Generic read-only contract binding to access the raw methods on
}

// CurveRewarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveRewarderTransactorRaw struct {
	Contract *CurveRewarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveRewarder creates a new instance of CurveRewarder, bound to a specific deployed contract.
func NewCurveRewarder(address common.Address, backend bind.ContractBackend) (*CurveRewarder, error) {
	contract, err := bindCurveRewarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveRewarder{CurveRewarderCaller: CurveRewarderCaller{contract: contract}, CurveRewarderTransactor: CurveRewarderTransactor{contract: contract}, CurveRewarderFilterer: CurveRewarderFilterer{contract: contract}}, nil
}

// NewCurveRewarderCaller creates a new read-only instance of CurveRewarder, bound to a specific deployed contract.
func NewCurveRewarderCaller(address common.Address, caller bind.ContractCaller) (*CurveRewarderCaller, error) {
	contract, err := bindCurveRewarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveRewarderCaller{contract: contract}, nil
}

// NewCurveRewarderTransactor creates a new write-only instance of CurveRewarder, bound to a specific deployed contract.
func NewCurveRewarderTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveRewarderTransactor, error) {
	contract, err := bindCurveRewarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveRewarderTransactor{contract: contract}, nil
}

// NewCurveRewarderFilterer creates a new log filterer instance of CurveRewarder, bound to a specific deployed contract.
func NewCurveRewarderFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveRewarderFilterer, error) {
	contract, err := bindCurveRewarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveRewarderFilterer{contract: contract}, nil
}

// bindCurveRewarder binds a generic wrapper to an already deployed contract.
func bindCurveRewarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveRewarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveRewarder *CurveRewarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveRewarder.Contract.CurveRewarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveRewarder *CurveRewarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRewarder.Contract.CurveRewarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveRewarder *CurveRewarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveRewarder.Contract.CurveRewarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveRewarder *CurveRewarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveRewarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveRewarder *CurveRewarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRewarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveRewarder *CurveRewarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveRewarder.Contract.contract.Transact(opts, method, params...)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveRewarder *CurveRewarderCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveRewarder.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveRewarder *CurveRewarderSession) FutureOwner() (common.Address, error) {
	return _CurveRewarder.Contract.FutureOwner(&_CurveRewarder.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveRewarder *CurveRewarderCallerSession) FutureOwner() (common.Address, error) {
	return _CurveRewarder.Contract.FutureOwner(&_CurveRewarder.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0x629d46c2.
//
// Solidity: function last_update_time() view returns(uint256)
func (_CurveRewarder *CurveRewarderCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRewarder.contract.Call(opts, &out, "last_update_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0x629d46c2.
//
// Solidity: function last_update_time() view returns(uint256)
func (_CurveRewarder *CurveRewarderSession) LastUpdateTime() (*big.Int, error) {
	return _CurveRewarder.Contract.LastUpdateTime(&_CurveRewarder.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0x629d46c2.
//
// Solidity: function last_update_time() view returns(uint256)
func (_CurveRewarder *CurveRewarderCallerSession) LastUpdateTime() (*big.Int, error) {
	return _CurveRewarder.Contract.LastUpdateTime(&_CurveRewarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveRewarder *CurveRewarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveRewarder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveRewarder *CurveRewarderSession) Owner() (common.Address, error) {
	return _CurveRewarder.Contract.Owner(&_CurveRewarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveRewarder *CurveRewarderCallerSession) Owner() (common.Address, error) {
	return _CurveRewarder.Contract.Owner(&_CurveRewarder.CallOpts)
}

// RewardCount is a free data retrieval call binding the contract method 0x963c94b9.
//
// Solidity: function reward_count() view returns(uint256)
func (_CurveRewarder *CurveRewarderCaller) RewardCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveRewarder.contract.Call(opts, &out, "reward_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardCount is a free data retrieval call binding the contract method 0x963c94b9.
//
// Solidity: function reward_count() view returns(uint256)
func (_CurveRewarder *CurveRewarderSession) RewardCount() (*big.Int, error) {
	return _CurveRewarder.Contract.RewardCount(&_CurveRewarder.CallOpts)
}

// RewardCount is a free data retrieval call binding the contract method 0x963c94b9.
//
// Solidity: function reward_count() view returns(uint256)
func (_CurveRewarder *CurveRewarderCallerSession) RewardCount() (*big.Int, error) {
	return _CurveRewarder.Contract.RewardCount(&_CurveRewarder.CallOpts)
}

// RewardData is a free data retrieval call binding the contract method 0x48e9c65e.
//
// Solidity: function reward_data(address arg0) view returns(address distributor, uint256 period_finish, uint256 rate, uint256 duration, uint256 received, uint256 paid)
func (_CurveRewarder *CurveRewarderCaller) RewardData(opts *bind.CallOpts, arg0 common.Address) (struct {
	Distributor  common.Address
	PeriodFinish *big.Int
	Rate         *big.Int
	Duration     *big.Int
	Received     *big.Int
	Paid         *big.Int
}, error) {
	var out []interface{}
	err := _CurveRewarder.contract.Call(opts, &out, "reward_data", arg0)

	outstruct := new(struct {
		Distributor  common.Address
		PeriodFinish *big.Int
		Rate         *big.Int
		Duration     *big.Int
		Received     *big.Int
		Paid         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Distributor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PeriodFinish = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Received = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Paid = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardData is a free data retrieval call binding the contract method 0x48e9c65e.
//
// Solidity: function reward_data(address arg0) view returns(address distributor, uint256 period_finish, uint256 rate, uint256 duration, uint256 received, uint256 paid)
func (_CurveRewarder *CurveRewarderSession) RewardData(arg0 common.Address) (struct {
	Distributor  common.Address
	PeriodFinish *big.Int
	Rate         *big.Int
	Duration     *big.Int
	Received     *big.Int
	Paid         *big.Int
}, error) {
	return _CurveRewarder.Contract.RewardData(&_CurveRewarder.CallOpts, arg0)
}

// RewardData is a free data retrieval call binding the contract method 0x48e9c65e.
//
// Solidity: function reward_data(address arg0) view returns(address distributor, uint256 period_finish, uint256 rate, uint256 duration, uint256 received, uint256 paid)
func (_CurveRewarder *CurveRewarderCallerSession) RewardData(arg0 common.Address) (struct {
	Distributor  common.Address
	PeriodFinish *big.Int
	Rate         *big.Int
	Duration     *big.Int
	Received     *big.Int
	Paid         *big.Int
}, error) {
	return _CurveRewarder.Contract.RewardData(&_CurveRewarder.CallOpts, arg0)
}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveRewarder *CurveRewarderCaller) RewardReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveRewarder.contract.Call(opts, &out, "reward_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveRewarder *CurveRewarderSession) RewardReceiver() (common.Address, error) {
	return _CurveRewarder.Contract.RewardReceiver(&_CurveRewarder.CallOpts)
}

// RewardReceiver is a free data retrieval call binding the contract method 0xb618ba62.
//
// Solidity: function reward_receiver() view returns(address)
func (_CurveRewarder *CurveRewarderCallerSession) RewardReceiver() (common.Address, error) {
	return _CurveRewarder.Contract.RewardReceiver(&_CurveRewarder.CallOpts)
}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveRewarder *CurveRewarderCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveRewarder.contract.Call(opts, &out, "reward_tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveRewarder *CurveRewarderSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _CurveRewarder.Contract.RewardTokens(&_CurveRewarder.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveRewarder *CurveRewarderCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _CurveRewarder.Contract.RewardTokens(&_CurveRewarder.CallOpts, arg0)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveRewarder *CurveRewarderTransactor) AcceptTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRewarder.contract.Transact(opts, "accept_transfer_ownership")
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveRewarder *CurveRewarderSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveRewarder.Contract.AcceptTransferOwnership(&_CurveRewarder.TransactOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveRewarder *CurveRewarderTransactorSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveRewarder.Contract.AcceptTransferOwnership(&_CurveRewarder.TransactOpts)
}

// AddReward is a paid mutator transaction binding the contract method 0x661ab0b2.
//
// Solidity: function add_reward(address _token, address _distributor, uint256 _duration) returns()
func (_CurveRewarder *CurveRewarderTransactor) AddReward(opts *bind.TransactOpts, _token common.Address, _distributor common.Address, _duration *big.Int) (*types.Transaction, error) {
	return _CurveRewarder.contract.Transact(opts, "add_reward", _token, _distributor, _duration)
}

// AddReward is a paid mutator transaction binding the contract method 0x661ab0b2.
//
// Solidity: function add_reward(address _token, address _distributor, uint256 _duration) returns()
func (_CurveRewarder *CurveRewarderSession) AddReward(_token common.Address, _distributor common.Address, _duration *big.Int) (*types.Transaction, error) {
	return _CurveRewarder.Contract.AddReward(&_CurveRewarder.TransactOpts, _token, _distributor, _duration)
}

// AddReward is a paid mutator transaction binding the contract method 0x661ab0b2.
//
// Solidity: function add_reward(address _token, address _distributor, uint256 _duration) returns()
func (_CurveRewarder *CurveRewarderTransactorSession) AddReward(_token common.Address, _distributor common.Address, _duration *big.Int) (*types.Transaction, error) {
	return _CurveRewarder.Contract.AddReward(&_CurveRewarder.TransactOpts, _token, _distributor, _duration)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveRewarder *CurveRewarderTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _CurveRewarder.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveRewarder *CurveRewarderSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveRewarder.Contract.CommitTransferOwnership(&_CurveRewarder.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveRewarder *CurveRewarderTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveRewarder.Contract.CommitTransferOwnership(&_CurveRewarder.TransactOpts, _owner)
}

// GetReward is a paid mutator transaction binding the contract method 0x1afe22a6.
//
// Solidity: function get_reward() returns()
func (_CurveRewarder *CurveRewarderTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveRewarder.contract.Transact(opts, "get_reward")
}

// GetReward is a paid mutator transaction binding the contract method 0x1afe22a6.
//
// Solidity: function get_reward() returns()
func (_CurveRewarder *CurveRewarderSession) GetReward() (*types.Transaction, error) {
	return _CurveRewarder.Contract.GetReward(&_CurveRewarder.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x1afe22a6.
//
// Solidity: function get_reward() returns()
func (_CurveRewarder *CurveRewarderTransactorSession) GetReward() (*types.Transaction, error) {
	return _CurveRewarder.Contract.GetReward(&_CurveRewarder.TransactOpts)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xa4f6af70.
//
// Solidity: function notify_reward_amount(address _token) returns()
func (_CurveRewarder *CurveRewarderTransactor) NotifyRewardAmount(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _CurveRewarder.contract.Transact(opts, "notify_reward_amount", _token)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xa4f6af70.
//
// Solidity: function notify_reward_amount(address _token) returns()
func (_CurveRewarder *CurveRewarderSession) NotifyRewardAmount(_token common.Address) (*types.Transaction, error) {
	return _CurveRewarder.Contract.NotifyRewardAmount(&_CurveRewarder.TransactOpts, _token)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xa4f6af70.
//
// Solidity: function notify_reward_amount(address _token) returns()
func (_CurveRewarder *CurveRewarderTransactorSession) NotifyRewardAmount(_token common.Address) (*types.Transaction, error) {
	return _CurveRewarder.Contract.NotifyRewardAmount(&_CurveRewarder.TransactOpts, _token)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x5de59dbc.
//
// Solidity: function remove_reward(address _token) returns()
func (_CurveRewarder *CurveRewarderTransactor) RemoveReward(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _CurveRewarder.contract.Transact(opts, "remove_reward", _token)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x5de59dbc.
//
// Solidity: function remove_reward(address _token) returns()
func (_CurveRewarder *CurveRewarderSession) RemoveReward(_token common.Address) (*types.Transaction, error) {
	return _CurveRewarder.Contract.RemoveReward(&_CurveRewarder.TransactOpts, _token)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x5de59dbc.
//
// Solidity: function remove_reward(address _token) returns()
func (_CurveRewarder *CurveRewarderTransactorSession) RemoveReward(_token common.Address) (*types.Transaction, error) {
	return _CurveRewarder.Contract.RemoveReward(&_CurveRewarder.TransactOpts, _token)
}

// SetReceiver is a paid mutator transaction binding the contract method 0xd1dd6f56.
//
// Solidity: function set_receiver(address _receiver) returns()
func (_CurveRewarder *CurveRewarderTransactor) SetReceiver(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _CurveRewarder.contract.Transact(opts, "set_receiver", _receiver)
}

// SetReceiver is a paid mutator transaction binding the contract method 0xd1dd6f56.
//
// Solidity: function set_receiver(address _receiver) returns()
func (_CurveRewarder *CurveRewarderSession) SetReceiver(_receiver common.Address) (*types.Transaction, error) {
	return _CurveRewarder.Contract.SetReceiver(&_CurveRewarder.TransactOpts, _receiver)
}

// SetReceiver is a paid mutator transaction binding the contract method 0xd1dd6f56.
//
// Solidity: function set_receiver(address _receiver) returns()
func (_CurveRewarder *CurveRewarderTransactorSession) SetReceiver(_receiver common.Address) (*types.Transaction, error) {
	return _CurveRewarder.Contract.SetReceiver(&_CurveRewarder.TransactOpts, _receiver)
}

// SetRewardDistributor is a paid mutator transaction binding the contract method 0x058a3a24.
//
// Solidity: function set_reward_distributor(address _token, address _distributor) returns()
func (_CurveRewarder *CurveRewarderTransactor) SetRewardDistributor(opts *bind.TransactOpts, _token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _CurveRewarder.contract.Transact(opts, "set_reward_distributor", _token, _distributor)
}

// SetRewardDistributor is a paid mutator transaction binding the contract method 0x058a3a24.
//
// Solidity: function set_reward_distributor(address _token, address _distributor) returns()
func (_CurveRewarder *CurveRewarderSession) SetRewardDistributor(_token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _CurveRewarder.Contract.SetRewardDistributor(&_CurveRewarder.TransactOpts, _token, _distributor)
}

// SetRewardDistributor is a paid mutator transaction binding the contract method 0x058a3a24.
//
// Solidity: function set_reward_distributor(address _token, address _distributor) returns()
func (_CurveRewarder *CurveRewarderTransactorSession) SetRewardDistributor(_token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _CurveRewarder.Contract.SetRewardDistributor(&_CurveRewarder.TransactOpts, _token, _distributor)
}

// SetRewardDuration is a paid mutator transaction binding the contract method 0xd88449ee.
//
// Solidity: function set_reward_duration(address _token, uint256 _duration) returns()
func (_CurveRewarder *CurveRewarderTransactor) SetRewardDuration(opts *bind.TransactOpts, _token common.Address, _duration *big.Int) (*types.Transaction, error) {
	return _CurveRewarder.contract.Transact(opts, "set_reward_duration", _token, _duration)
}

// SetRewardDuration is a paid mutator transaction binding the contract method 0xd88449ee.
//
// Solidity: function set_reward_duration(address _token, uint256 _duration) returns()
func (_CurveRewarder *CurveRewarderSession) SetRewardDuration(_token common.Address, _duration *big.Int) (*types.Transaction, error) {
	return _CurveRewarder.Contract.SetRewardDuration(&_CurveRewarder.TransactOpts, _token, _duration)
}

// SetRewardDuration is a paid mutator transaction binding the contract method 0xd88449ee.
//
// Solidity: function set_reward_duration(address _token, uint256 _duration) returns()
func (_CurveRewarder *CurveRewarderTransactorSession) SetRewardDuration(_token common.Address, _duration *big.Int) (*types.Transaction, error) {
	return _CurveRewarder.Contract.SetRewardDuration(&_CurveRewarder.TransactOpts, _token, _duration)
}
