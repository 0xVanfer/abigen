// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platypusMasterPlatypusV4

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

// IMasterPlatypusV4UserInfo is an auto generated low-level Go binding around an user-defined struct.
type IMasterPlatypusV4UserInfo struct {
	Amount       *big.Int
	Factor       *big.Int
	RewardDebt   *big.Int
	ClaimablePtp *big.Int
}

// PlatypusMasterPlatypusV4MetaData contains all meta data concerning the PlatypusMasterPlatypusV4 contract.
var PlatypusMasterPlatypusV4MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"additionalRewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dilutingRepartition\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPtpWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"}],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getSumOfFactors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"factor\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rewardDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimablePtp\",\"type\":\"uint128\"}],\"internalType\":\"structIMasterPlatypusV4.UserInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ptp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vePtp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dilutingRepartition\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPoolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"name\":\"multiClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"additionalRewards\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newMasterPlatypus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonDilutingRepartition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCandidate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingPtp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"bonusTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"bonusTokenSymbols\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pendingBonusTokens\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platypusTreasure\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewarder\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"sumOfFactors\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accPtpPerShare\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accPtpPerFactorShare\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"proposeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ptp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"rewarderBonusTokenInfo\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"bonusTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"bonusTokenSymbols\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPoolLength\",\"type\":\"uint256\"}],\"name\":\"setMaxPoolLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMasterPlatypus\",\"type\":\"address\"}],\"name\":\"setNewMasterPlatypus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_platypusTreasure\",\"type\":\"address\"}],\"name\":\"setPlatypusTreasure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"setRewarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVePtp\",\"type\":\"address\"}],\"name\":\"setVePtp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dilutingRepartition\",\"type\":\"uint16\"}],\"name\":\"updateEmissionRepartition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newVePtpBalance\",\"type\":\"uint256\"}],\"name\":\"updateFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"factor\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rewardDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimablePtp\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vePtp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"additionalRewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PlatypusMasterPlatypusV4ABI is the input ABI used to generate the binding from.
// Deprecated: Use PlatypusMasterPlatypusV4MetaData.ABI instead.
var PlatypusMasterPlatypusV4ABI = PlatypusMasterPlatypusV4MetaData.ABI

// PlatypusMasterPlatypusV4 is an auto generated Go binding around an Ethereum contract.
type PlatypusMasterPlatypusV4 struct {
	PlatypusMasterPlatypusV4Caller     // Read-only binding to the contract
	PlatypusMasterPlatypusV4Transactor // Write-only binding to the contract
	PlatypusMasterPlatypusV4Filterer   // Log filterer for contract events
}

// PlatypusMasterPlatypusV4Caller is an auto generated read-only Go binding around an Ethereum contract.
type PlatypusMasterPlatypusV4Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusMasterPlatypusV4Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PlatypusMasterPlatypusV4Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusMasterPlatypusV4Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlatypusMasterPlatypusV4Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatypusMasterPlatypusV4Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlatypusMasterPlatypusV4Session struct {
	Contract     *PlatypusMasterPlatypusV4 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PlatypusMasterPlatypusV4CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlatypusMasterPlatypusV4CallerSession struct {
	Contract *PlatypusMasterPlatypusV4Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// PlatypusMasterPlatypusV4TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlatypusMasterPlatypusV4TransactorSession struct {
	Contract     *PlatypusMasterPlatypusV4Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PlatypusMasterPlatypusV4Raw is an auto generated low-level Go binding around an Ethereum contract.
type PlatypusMasterPlatypusV4Raw struct {
	Contract *PlatypusMasterPlatypusV4 // Generic contract binding to access the raw methods on
}

// PlatypusMasterPlatypusV4CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlatypusMasterPlatypusV4CallerRaw struct {
	Contract *PlatypusMasterPlatypusV4Caller // Generic read-only contract binding to access the raw methods on
}

// PlatypusMasterPlatypusV4TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlatypusMasterPlatypusV4TransactorRaw struct {
	Contract *PlatypusMasterPlatypusV4Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPlatypusMasterPlatypusV4 creates a new instance of PlatypusMasterPlatypusV4, bound to a specific deployed contract.
func NewPlatypusMasterPlatypusV4(address common.Address, backend bind.ContractBackend) (*PlatypusMasterPlatypusV4, error) {
	contract, err := bindPlatypusMasterPlatypusV4(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlatypusMasterPlatypusV4{PlatypusMasterPlatypusV4Caller: PlatypusMasterPlatypusV4Caller{contract: contract}, PlatypusMasterPlatypusV4Transactor: PlatypusMasterPlatypusV4Transactor{contract: contract}, PlatypusMasterPlatypusV4Filterer: PlatypusMasterPlatypusV4Filterer{contract: contract}}, nil
}

// NewPlatypusMasterPlatypusV4Caller creates a new read-only instance of PlatypusMasterPlatypusV4, bound to a specific deployed contract.
func NewPlatypusMasterPlatypusV4Caller(address common.Address, caller bind.ContractCaller) (*PlatypusMasterPlatypusV4Caller, error) {
	contract, err := bindPlatypusMasterPlatypusV4(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlatypusMasterPlatypusV4Caller{contract: contract}, nil
}

// NewPlatypusMasterPlatypusV4Transactor creates a new write-only instance of PlatypusMasterPlatypusV4, bound to a specific deployed contract.
func NewPlatypusMasterPlatypusV4Transactor(address common.Address, transactor bind.ContractTransactor) (*PlatypusMasterPlatypusV4Transactor, error) {
	contract, err := bindPlatypusMasterPlatypusV4(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlatypusMasterPlatypusV4Transactor{contract: contract}, nil
}

// NewPlatypusMasterPlatypusV4Filterer creates a new log filterer instance of PlatypusMasterPlatypusV4, bound to a specific deployed contract.
func NewPlatypusMasterPlatypusV4Filterer(address common.Address, filterer bind.ContractFilterer) (*PlatypusMasterPlatypusV4Filterer, error) {
	contract, err := bindPlatypusMasterPlatypusV4(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlatypusMasterPlatypusV4Filterer{contract: contract}, nil
}

// bindPlatypusMasterPlatypusV4 binds a generic wrapper to an already deployed contract.
func bindPlatypusMasterPlatypusV4(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlatypusMasterPlatypusV4ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatypusMasterPlatypusV4.Contract.PlatypusMasterPlatypusV4Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.PlatypusMasterPlatypusV4Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.PlatypusMasterPlatypusV4Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatypusMasterPlatypusV4.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.contract.Transact(opts, method, params...)
}

// DilutingRepartition is a free data retrieval call binding the contract method 0xf57347a8.
//
// Solidity: function dilutingRepartition() view returns(uint16)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) DilutingRepartition(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "dilutingRepartition")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DilutingRepartition is a free data retrieval call binding the contract method 0xf57347a8.
//
// Solidity: function dilutingRepartition() view returns(uint16)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) DilutingRepartition() (uint16, error) {
	return _PlatypusMasterPlatypusV4.Contract.DilutingRepartition(&_PlatypusMasterPlatypusV4.CallOpts)
}

// DilutingRepartition is a free data retrieval call binding the contract method 0xf57347a8.
//
// Solidity: function dilutingRepartition() view returns(uint16)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) DilutingRepartition() (uint16, error) {
	return _PlatypusMasterPlatypusV4.Contract.DilutingRepartition(&_PlatypusMasterPlatypusV4.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address _lp) view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) GetPoolId(opts *bind.CallOpts, _lp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "getPoolId", _lp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address _lp) view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) GetPoolId(_lp common.Address) (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.GetPoolId(&_PlatypusMasterPlatypusV4.CallOpts, _lp)
}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address _lp) view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) GetPoolId(_lp common.Address) (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.GetPoolId(&_PlatypusMasterPlatypusV4.CallOpts, _lp)
}

// GetSumOfFactors is a free data retrieval call binding the contract method 0x5da36238.
//
// Solidity: function getSumOfFactors(uint256 _pid) view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) GetSumOfFactors(opts *bind.CallOpts, _pid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "getSumOfFactors", _pid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSumOfFactors is a free data retrieval call binding the contract method 0x5da36238.
//
// Solidity: function getSumOfFactors(uint256 _pid) view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) GetSumOfFactors(_pid *big.Int) (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.GetSumOfFactors(&_PlatypusMasterPlatypusV4.CallOpts, _pid)
}

// GetSumOfFactors is a free data retrieval call binding the contract method 0x5da36238.
//
// Solidity: function getSumOfFactors(uint256 _pid) view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) GetSumOfFactors(_pid *big.Int) (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.GetSumOfFactors(&_PlatypusMasterPlatypusV4.CallOpts, _pid)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x1069f3b5.
//
// Solidity: function getUserInfo(uint256 _pid, address _user) view returns((uint128,uint128,uint128,uint128))
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) GetUserInfo(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (IMasterPlatypusV4UserInfo, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "getUserInfo", _pid, _user)

	if err != nil {
		return *new(IMasterPlatypusV4UserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IMasterPlatypusV4UserInfo)).(*IMasterPlatypusV4UserInfo)

	return out0, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x1069f3b5.
//
// Solidity: function getUserInfo(uint256 _pid, address _user) view returns((uint128,uint128,uint128,uint128))
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) GetUserInfo(_pid *big.Int, _user common.Address) (IMasterPlatypusV4UserInfo, error) {
	return _PlatypusMasterPlatypusV4.Contract.GetUserInfo(&_PlatypusMasterPlatypusV4.CallOpts, _pid, _user)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x1069f3b5.
//
// Solidity: function getUserInfo(uint256 _pid, address _user) view returns((uint128,uint128,uint128,uint128))
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) GetUserInfo(_pid *big.Int, _user common.Address) (IMasterPlatypusV4UserInfo, error) {
	return _PlatypusMasterPlatypusV4.Contract.GetUserInfo(&_PlatypusMasterPlatypusV4.CallOpts, _pid, _user)
}

// MaxPoolLength is a free data retrieval call binding the contract method 0x7b62a738.
//
// Solidity: function maxPoolLength() view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) MaxPoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "maxPoolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPoolLength is a free data retrieval call binding the contract method 0x7b62a738.
//
// Solidity: function maxPoolLength() view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) MaxPoolLength() (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.MaxPoolLength(&_PlatypusMasterPlatypusV4.CallOpts)
}

// MaxPoolLength is a free data retrieval call binding the contract method 0x7b62a738.
//
// Solidity: function maxPoolLength() view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) MaxPoolLength() (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.MaxPoolLength(&_PlatypusMasterPlatypusV4.CallOpts)
}

// NewMasterPlatypus is a free data retrieval call binding the contract method 0x01126816.
//
// Solidity: function newMasterPlatypus() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) NewMasterPlatypus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "newMasterPlatypus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewMasterPlatypus is a free data retrieval call binding the contract method 0x01126816.
//
// Solidity: function newMasterPlatypus() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) NewMasterPlatypus() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.NewMasterPlatypus(&_PlatypusMasterPlatypusV4.CallOpts)
}

// NewMasterPlatypus is a free data retrieval call binding the contract method 0x01126816.
//
// Solidity: function newMasterPlatypus() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) NewMasterPlatypus() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.NewMasterPlatypus(&_PlatypusMasterPlatypusV4.CallOpts)
}

// NonDilutingRepartition is a free data retrieval call binding the contract method 0x04e3682a.
//
// Solidity: function nonDilutingRepartition() view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) NonDilutingRepartition(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "nonDilutingRepartition")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonDilutingRepartition is a free data retrieval call binding the contract method 0x04e3682a.
//
// Solidity: function nonDilutingRepartition() view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) NonDilutingRepartition() (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.NonDilutingRepartition(&_PlatypusMasterPlatypusV4.CallOpts)
}

// NonDilutingRepartition is a free data retrieval call binding the contract method 0x04e3682a.
//
// Solidity: function nonDilutingRepartition() view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) NonDilutingRepartition() (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.NonDilutingRepartition(&_PlatypusMasterPlatypusV4.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Owner() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.Owner(&_PlatypusMasterPlatypusV4.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) Owner() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.Owner(&_PlatypusMasterPlatypusV4.CallOpts)
}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) OwnerCandidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "ownerCandidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) OwnerCandidate() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.OwnerCandidate(&_PlatypusMasterPlatypusV4.CallOpts)
}

// OwnerCandidate is a free data retrieval call binding the contract method 0x5f504a82.
//
// Solidity: function ownerCandidate() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) OwnerCandidate() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.OwnerCandidate(&_PlatypusMasterPlatypusV4.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Paused() (bool, error) {
	return _PlatypusMasterPlatypusV4.Contract.Paused(&_PlatypusMasterPlatypusV4.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) Paused() (bool, error) {
	return _PlatypusMasterPlatypusV4.Contract.Paused(&_PlatypusMasterPlatypusV4.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingPtp, address[] bonusTokenAddresses, string[] bonusTokenSymbols, uint256[] pendingBonusTokens)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) PendingTokens(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (struct {
	PendingPtp          *big.Int
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
	PendingBonusTokens  []*big.Int
}, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "pendingTokens", _pid, _user)

	outstruct := new(struct {
		PendingPtp          *big.Int
		BonusTokenAddresses []common.Address
		BonusTokenSymbols   []string
		PendingBonusTokens  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingPtp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BonusTokenAddresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.BonusTokenSymbols = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.PendingBonusTokens = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingPtp, address[] bonusTokenAddresses, string[] bonusTokenSymbols, uint256[] pendingBonusTokens)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingPtp          *big.Int
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
	PendingBonusTokens  []*big.Int
}, error) {
	return _PlatypusMasterPlatypusV4.Contract.PendingTokens(&_PlatypusMasterPlatypusV4.CallOpts, _pid, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingPtp, address[] bonusTokenAddresses, string[] bonusTokenSymbols, uint256[] pendingBonusTokens)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingPtp          *big.Int
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
	PendingBonusTokens  []*big.Int
}, error) {
	return _PlatypusMasterPlatypusV4.Contract.PendingTokens(&_PlatypusMasterPlatypusV4.CallOpts, _pid, _user)
}

// PlatypusTreasure is a free data retrieval call binding the contract method 0x8c1ab965.
//
// Solidity: function platypusTreasure() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) PlatypusTreasure(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "platypusTreasure")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatypusTreasure is a free data retrieval call binding the contract method 0x8c1ab965.
//
// Solidity: function platypusTreasure() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) PlatypusTreasure() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.PlatypusTreasure(&_PlatypusMasterPlatypusV4.CallOpts)
}

// PlatypusTreasure is a free data retrieval call binding the contract method 0x8c1ab965.
//
// Solidity: function platypusTreasure() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) PlatypusTreasure() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.PlatypusTreasure(&_PlatypusMasterPlatypusV4.CallOpts)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, address rewarder, uint128 sumOfFactors, uint128 accPtpPerShare, uint128 accPtpPerFactorShare)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken              common.Address
	Rewarder             common.Address
	SumOfFactors         *big.Int
	AccPtpPerShare       *big.Int
	AccPtpPerFactorShare *big.Int
}, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken              common.Address
		Rewarder             common.Address
		SumOfFactors         *big.Int
		AccPtpPerShare       *big.Int
		AccPtpPerFactorShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Rewarder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.SumOfFactors = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccPtpPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AccPtpPerFactorShare = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, address rewarder, uint128 sumOfFactors, uint128 accPtpPerShare, uint128 accPtpPerFactorShare)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) PoolInfo(arg0 *big.Int) (struct {
	LpToken              common.Address
	Rewarder             common.Address
	SumOfFactors         *big.Int
	AccPtpPerShare       *big.Int
	AccPtpPerFactorShare *big.Int
}, error) {
	return _PlatypusMasterPlatypusV4.Contract.PoolInfo(&_PlatypusMasterPlatypusV4.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, address rewarder, uint128 sumOfFactors, uint128 accPtpPerShare, uint128 accPtpPerFactorShare)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken              common.Address
	Rewarder             common.Address
	SumOfFactors         *big.Int
	AccPtpPerShare       *big.Int
	AccPtpPerFactorShare *big.Int
}, error) {
	return _PlatypusMasterPlatypusV4.Contract.PoolInfo(&_PlatypusMasterPlatypusV4.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) PoolLength() (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.PoolLength(&_PlatypusMasterPlatypusV4.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) PoolLength() (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.PoolLength(&_PlatypusMasterPlatypusV4.CallOpts)
}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) Ptp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "ptp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Ptp() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.Ptp(&_PlatypusMasterPlatypusV4.CallOpts)
}

// Ptp is a free data retrieval call binding the contract method 0x6af66772.
//
// Solidity: function ptp() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) Ptp() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.Ptp(&_PlatypusMasterPlatypusV4.CallOpts)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address[] bonusTokenAddresses, string[] bonusTokenSymbols)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) RewarderBonusTokenInfo(opts *bind.CallOpts, _pid *big.Int) (struct {
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
}, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "rewarderBonusTokenInfo", _pid)

	outstruct := new(struct {
		BonusTokenAddresses []common.Address
		BonusTokenSymbols   []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BonusTokenAddresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.BonusTokenSymbols = *abi.ConvertType(out[1], new([]string)).(*[]string)

	return *outstruct, err

}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address[] bonusTokenAddresses, string[] bonusTokenSymbols)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
}, error) {
	return _PlatypusMasterPlatypusV4.Contract.RewarderBonusTokenInfo(&_PlatypusMasterPlatypusV4.CallOpts, _pid)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address[] bonusTokenAddresses, string[] bonusTokenSymbols)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
}, error) {
	return _PlatypusMasterPlatypusV4.Contract.RewarderBonusTokenInfo(&_PlatypusMasterPlatypusV4.CallOpts, _pid)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint128 amount, uint128 factor, uint128 rewardDebt, uint128 claimablePtp)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount       *big.Int
	Factor       *big.Int
	RewardDebt   *big.Int
	ClaimablePtp *big.Int
}, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount       *big.Int
		Factor       *big.Int
		RewardDebt   *big.Int
		ClaimablePtp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Factor = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ClaimablePtp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint128 amount, uint128 factor, uint128 rewardDebt, uint128 claimablePtp)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount       *big.Int
	Factor       *big.Int
	RewardDebt   *big.Int
	ClaimablePtp *big.Int
}, error) {
	return _PlatypusMasterPlatypusV4.Contract.UserInfo(&_PlatypusMasterPlatypusV4.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint128 amount, uint128 factor, uint128 rewardDebt, uint128 claimablePtp)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount       *big.Int
	Factor       *big.Int
	RewardDebt   *big.Int
	ClaimablePtp *big.Int
}, error) {
	return _PlatypusMasterPlatypusV4.Contract.UserInfo(&_PlatypusMasterPlatypusV4.CallOpts, arg0, arg1)
}

// VePtp is a free data retrieval call binding the contract method 0x82c780a1.
//
// Solidity: function vePtp() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) VePtp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "vePtp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VePtp is a free data retrieval call binding the contract method 0x82c780a1.
//
// Solidity: function vePtp() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) VePtp() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.VePtp(&_PlatypusMasterPlatypusV4.CallOpts)
}

// VePtp is a free data retrieval call binding the contract method 0x82c780a1.
//
// Solidity: function vePtp() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) VePtp() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.VePtp(&_PlatypusMasterPlatypusV4.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Version() (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.Version(&_PlatypusMasterPlatypusV4.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) Version() (*big.Int, error) {
	return _PlatypusMasterPlatypusV4.Contract.Version(&_PlatypusMasterPlatypusV4.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Caller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatypusMasterPlatypusV4.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Voter() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.Voter(&_PlatypusMasterPlatypusV4.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4CallerSession) Voter() (common.Address, error) {
	return _PlatypusMasterPlatypusV4.Contract.Voter(&_PlatypusMasterPlatypusV4.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) AcceptOwnership() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.AcceptOwnership(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.AcceptOwnership(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0x52c28fab.
//
// Solidity: function add(address _lpToken, address _rewarder) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) Add(opts *bind.TransactOpts, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "add", _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0x52c28fab.
//
// Solidity: function add(address _lpToken, address _rewarder) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Add(_lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Add(&_PlatypusMasterPlatypusV4.TransactOpts, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0x52c28fab.
//
// Solidity: function add(address _lpToken, address _rewarder) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) Add(_lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Add(&_PlatypusMasterPlatypusV4.TransactOpts, _lpToken, _rewarder)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256 reward, uint256[] additionalRewards)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256 reward, uint256[] additionalRewards)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Deposit(&_PlatypusMasterPlatypusV4.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256 reward, uint256[] additionalRewards)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Deposit(&_PlatypusMasterPlatypusV4.TransactOpts, _pid, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) DepositFor(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "depositFor", _pid, _amount, _user)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) DepositFor(_pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.DepositFor(&_PlatypusMasterPlatypusV4.TransactOpts, _pid, _amount, _user)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) DepositFor(_pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.DepositFor(&_PlatypusMasterPlatypusV4.TransactOpts, _pid, _amount, _user)
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) EmergencyPtpWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "emergencyPtpWithdraw")
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) EmergencyPtpWithdraw() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.EmergencyPtpWithdraw(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// EmergencyPtpWithdraw is a paid mutator transaction binding the contract method 0x7dd38dcc.
//
// Solidity: function emergencyPtpWithdraw() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) EmergencyPtpWithdraw() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.EmergencyPtpWithdraw(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.EmergencyWithdraw(&_PlatypusMasterPlatypusV4.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.EmergencyWithdraw(&_PlatypusMasterPlatypusV4.TransactOpts, _pid)
}

// Initialize is a paid mutator transaction binding the contract method 0x2e112757.
//
// Solidity: function initialize(address _ptp, address _vePtp, address _voter, uint16 _dilutingRepartition) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) Initialize(opts *bind.TransactOpts, _ptp common.Address, _vePtp common.Address, _voter common.Address, _dilutingRepartition uint16) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "initialize", _ptp, _vePtp, _voter, _dilutingRepartition)
}

// Initialize is a paid mutator transaction binding the contract method 0x2e112757.
//
// Solidity: function initialize(address _ptp, address _vePtp, address _voter, uint16 _dilutingRepartition) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Initialize(_ptp common.Address, _vePtp common.Address, _voter common.Address, _dilutingRepartition uint16) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Initialize(&_PlatypusMasterPlatypusV4.TransactOpts, _ptp, _vePtp, _voter, _dilutingRepartition)
}

// Initialize is a paid mutator transaction binding the contract method 0x2e112757.
//
// Solidity: function initialize(address _ptp, address _vePtp, address _voter, uint16 _dilutingRepartition) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) Initialize(_ptp common.Address, _vePtp common.Address, _voter common.Address, _dilutingRepartition uint16) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Initialize(&_PlatypusMasterPlatypusV4.TransactOpts, _ptp, _vePtp, _voter, _dilutingRepartition)
}

// Liquidate is a paid mutator transaction binding the contract method 0x79bd1eac.
//
// Solidity: function liquidate(uint256 _pid, address _user, uint256 _amount) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) Liquidate(opts *bind.TransactOpts, _pid *big.Int, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "liquidate", _pid, _user, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x79bd1eac.
//
// Solidity: function liquidate(uint256 _pid, address _user, uint256 _amount) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Liquidate(_pid *big.Int, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Liquidate(&_PlatypusMasterPlatypusV4.TransactOpts, _pid, _user, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x79bd1eac.
//
// Solidity: function liquidate(uint256 _pid, address _user, uint256 _amount) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) Liquidate(_pid *big.Int, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Liquidate(&_PlatypusMasterPlatypusV4.TransactOpts, _pid, _user, _amount)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) MassUpdatePools() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.MassUpdatePools(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.MassUpdatePools(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) Migrate(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "migrate", _pids)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Migrate(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Migrate(&_PlatypusMasterPlatypusV4.TransactOpts, _pids)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) Migrate(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Migrate(&_PlatypusMasterPlatypusV4.TransactOpts, _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256 reward, uint256[] amounts, uint256[][] additionalRewards)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) MultiClaim(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "multiClaim", _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256 reward, uint256[] amounts, uint256[][] additionalRewards)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) MultiClaim(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.MultiClaim(&_PlatypusMasterPlatypusV4.TransactOpts, _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256 reward, uint256[] amounts, uint256[][] additionalRewards)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) MultiClaim(_pids []*big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.MultiClaim(&_PlatypusMasterPlatypusV4.TransactOpts, _pids)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _lpToken, uint256 _amount) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) NotifyRewardAmount(opts *bind.TransactOpts, _lpToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "notifyRewardAmount", _lpToken, _amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _lpToken, uint256 _amount) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) NotifyRewardAmount(_lpToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.NotifyRewardAmount(&_PlatypusMasterPlatypusV4.TransactOpts, _lpToken, _amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _lpToken, uint256 _amount) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) NotifyRewardAmount(_lpToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.NotifyRewardAmount(&_PlatypusMasterPlatypusV4.TransactOpts, _lpToken, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Pause() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Pause(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) Pause() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Pause(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) ProposeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "proposeOwner", newOwner)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) ProposeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.ProposeOwner(&_PlatypusMasterPlatypusV4.TransactOpts, newOwner)
}

// ProposeOwner is a paid mutator transaction binding the contract method 0xb5ed298a.
//
// Solidity: function proposeOwner(address newOwner) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) ProposeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.ProposeOwner(&_PlatypusMasterPlatypusV4.TransactOpts, newOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) RenounceOwnership() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.RenounceOwnership(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.RenounceOwnership(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// SetMaxPoolLength is a paid mutator transaction binding the contract method 0x53c5eb44.
//
// Solidity: function setMaxPoolLength(uint256 _maxPoolLength) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) SetMaxPoolLength(opts *bind.TransactOpts, _maxPoolLength *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "setMaxPoolLength", _maxPoolLength)
}

// SetMaxPoolLength is a paid mutator transaction binding the contract method 0x53c5eb44.
//
// Solidity: function setMaxPoolLength(uint256 _maxPoolLength) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) SetMaxPoolLength(_maxPoolLength *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.SetMaxPoolLength(&_PlatypusMasterPlatypusV4.TransactOpts, _maxPoolLength)
}

// SetMaxPoolLength is a paid mutator transaction binding the contract method 0x53c5eb44.
//
// Solidity: function setMaxPoolLength(uint256 _maxPoolLength) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) SetMaxPoolLength(_maxPoolLength *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.SetMaxPoolLength(&_PlatypusMasterPlatypusV4.TransactOpts, _maxPoolLength)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) SetNewMasterPlatypus(opts *bind.TransactOpts, _newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "setNewMasterPlatypus", _newMasterPlatypus)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) SetNewMasterPlatypus(_newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.SetNewMasterPlatypus(&_PlatypusMasterPlatypusV4.TransactOpts, _newMasterPlatypus)
}

// SetNewMasterPlatypus is a paid mutator transaction binding the contract method 0x7b261591.
//
// Solidity: function setNewMasterPlatypus(address _newMasterPlatypus) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) SetNewMasterPlatypus(_newMasterPlatypus common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.SetNewMasterPlatypus(&_PlatypusMasterPlatypusV4.TransactOpts, _newMasterPlatypus)
}

// SetPlatypusTreasure is a paid mutator transaction binding the contract method 0x05ca5edc.
//
// Solidity: function setPlatypusTreasure(address _platypusTreasure) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) SetPlatypusTreasure(opts *bind.TransactOpts, _platypusTreasure common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "setPlatypusTreasure", _platypusTreasure)
}

// SetPlatypusTreasure is a paid mutator transaction binding the contract method 0x05ca5edc.
//
// Solidity: function setPlatypusTreasure(address _platypusTreasure) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) SetPlatypusTreasure(_platypusTreasure common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.SetPlatypusTreasure(&_PlatypusMasterPlatypusV4.TransactOpts, _platypusTreasure)
}

// SetPlatypusTreasure is a paid mutator transaction binding the contract method 0x05ca5edc.
//
// Solidity: function setPlatypusTreasure(address _platypusTreasure) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) SetPlatypusTreasure(_platypusTreasure common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.SetPlatypusTreasure(&_PlatypusMasterPlatypusV4.TransactOpts, _platypusTreasure)
}

// SetRewarder is a paid mutator transaction binding the contract method 0x43de3207.
//
// Solidity: function setRewarder(uint256 _pid, address _rewarder) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) SetRewarder(opts *bind.TransactOpts, _pid *big.Int, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "setRewarder", _pid, _rewarder)
}

// SetRewarder is a paid mutator transaction binding the contract method 0x43de3207.
//
// Solidity: function setRewarder(uint256 _pid, address _rewarder) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) SetRewarder(_pid *big.Int, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.SetRewarder(&_PlatypusMasterPlatypusV4.TransactOpts, _pid, _rewarder)
}

// SetRewarder is a paid mutator transaction binding the contract method 0x43de3207.
//
// Solidity: function setRewarder(uint256 _pid, address _rewarder) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) SetRewarder(_pid *big.Int, _rewarder common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.SetRewarder(&_PlatypusMasterPlatypusV4.TransactOpts, _pid, _rewarder)
}

// SetVePtp is a paid mutator transaction binding the contract method 0x90d9c1c3.
//
// Solidity: function setVePtp(address _newVePtp) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) SetVePtp(opts *bind.TransactOpts, _newVePtp common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "setVePtp", _newVePtp)
}

// SetVePtp is a paid mutator transaction binding the contract method 0x90d9c1c3.
//
// Solidity: function setVePtp(address _newVePtp) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) SetVePtp(_newVePtp common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.SetVePtp(&_PlatypusMasterPlatypusV4.TransactOpts, _newVePtp)
}

// SetVePtp is a paid mutator transaction binding the contract method 0x90d9c1c3.
//
// Solidity: function setVePtp(address _newVePtp) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) SetVePtp(_newVePtp common.Address) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.SetVePtp(&_PlatypusMasterPlatypusV4.TransactOpts, _newVePtp)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Unpause() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Unpause(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) Unpause() (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Unpause(&_PlatypusMasterPlatypusV4.TransactOpts)
}

// UpdateEmissionRepartition is a paid mutator transaction binding the contract method 0xa1ef9608.
//
// Solidity: function updateEmissionRepartition(uint16 _dilutingRepartition) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) UpdateEmissionRepartition(opts *bind.TransactOpts, _dilutingRepartition uint16) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "updateEmissionRepartition", _dilutingRepartition)
}

// UpdateEmissionRepartition is a paid mutator transaction binding the contract method 0xa1ef9608.
//
// Solidity: function updateEmissionRepartition(uint16 _dilutingRepartition) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) UpdateEmissionRepartition(_dilutingRepartition uint16) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.UpdateEmissionRepartition(&_PlatypusMasterPlatypusV4.TransactOpts, _dilutingRepartition)
}

// UpdateEmissionRepartition is a paid mutator transaction binding the contract method 0xa1ef9608.
//
// Solidity: function updateEmissionRepartition(uint16 _dilutingRepartition) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) UpdateEmissionRepartition(_dilutingRepartition uint16) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.UpdateEmissionRepartition(&_PlatypusMasterPlatypusV4.TransactOpts, _dilutingRepartition)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVePtpBalance) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) UpdateFactor(opts *bind.TransactOpts, _user common.Address, _newVePtpBalance *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "updateFactor", _user, _newVePtpBalance)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVePtpBalance) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) UpdateFactor(_user common.Address, _newVePtpBalance *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.UpdateFactor(&_PlatypusMasterPlatypusV4.TransactOpts, _user, _newVePtpBalance)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVePtpBalance) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) UpdateFactor(_user common.Address, _newVePtpBalance *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.UpdateFactor(&_PlatypusMasterPlatypusV4.TransactOpts, _user, _newVePtpBalance)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.UpdatePool(&_PlatypusMasterPlatypusV4.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.UpdatePool(&_PlatypusMasterPlatypusV4.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256 reward, uint256[] additionalRewards)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Transactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256 reward, uint256[] additionalRewards)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4Session) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Withdraw(&_PlatypusMasterPlatypusV4.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256 reward, uint256[] additionalRewards)
func (_PlatypusMasterPlatypusV4 *PlatypusMasterPlatypusV4TransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PlatypusMasterPlatypusV4.Contract.Withdraw(&_PlatypusMasterPlatypusV4.TransactOpts, _pid, _amount)
}
