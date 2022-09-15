// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vectorChef

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

// VectorChefMetaData contains all meta data concerning the VectorChef contract.
var VectorChefMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PoolManagers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vtx\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vtxPerSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"}],\"name\":\"__MasterChefVTX_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_helper\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToPoolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accVTXPerShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"helper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowEmergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"authorizeForLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_locker\",\"type\":\"address\"}],\"name\":\"authorizeLocker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"}],\"name\":\"claimLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mainRewardToken\",\"type\":\"address\"}],\"name\":\"createRewarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"depositInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"}],\"name\":\"emergencyWithdrawWithReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"emission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizeOfPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPoint\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAuthorizedForLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"}],\"name\":\"lockFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_amount\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_index\",\"type\":\"uint256[]\"}],\"name\":\"multiUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_lps\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"}],\"name\":\"multiclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingVTX\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pendingBonusToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"}],\"name\":\"realEmergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registeredToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_locker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_helper\",\"type\":\"address\"}],\"name\":\"setPoolHelper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"setPoolManagerStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vtxPerSec\",\"type\":\"uint256\"}],\"name\":\"updateEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vtx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vtxPerSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"withdrawFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VectorChefABI is the input ABI used to generate the binding from.
// Deprecated: Use VectorChefMetaData.ABI instead.
var VectorChefABI = VectorChefMetaData.ABI

// VectorChef is an auto generated Go binding around an Ethereum contract.
type VectorChef struct {
	VectorChefCaller     // Read-only binding to the contract
	VectorChefTransactor // Write-only binding to the contract
	VectorChefFilterer   // Log filterer for contract events
}

// VectorChefCaller is an auto generated read-only Go binding around an Ethereum contract.
type VectorChefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorChefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VectorChefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorChefFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VectorChefFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VectorChefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VectorChefSession struct {
	Contract     *VectorChef       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VectorChefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VectorChefCallerSession struct {
	Contract *VectorChefCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// VectorChefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VectorChefTransactorSession struct {
	Contract     *VectorChefTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VectorChefRaw is an auto generated low-level Go binding around an Ethereum contract.
type VectorChefRaw struct {
	Contract *VectorChef // Generic contract binding to access the raw methods on
}

// VectorChefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VectorChefCallerRaw struct {
	Contract *VectorChefCaller // Generic read-only contract binding to access the raw methods on
}

// VectorChefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VectorChefTransactorRaw struct {
	Contract *VectorChefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVectorChef creates a new instance of VectorChef, bound to a specific deployed contract.
func NewVectorChef(address common.Address, backend bind.ContractBackend) (*VectorChef, error) {
	contract, err := bindVectorChef(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VectorChef{VectorChefCaller: VectorChefCaller{contract: contract}, VectorChefTransactor: VectorChefTransactor{contract: contract}, VectorChefFilterer: VectorChefFilterer{contract: contract}}, nil
}

// NewVectorChefCaller creates a new read-only instance of VectorChef, bound to a specific deployed contract.
func NewVectorChefCaller(address common.Address, caller bind.ContractCaller) (*VectorChefCaller, error) {
	contract, err := bindVectorChef(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VectorChefCaller{contract: contract}, nil
}

// NewVectorChefTransactor creates a new write-only instance of VectorChef, bound to a specific deployed contract.
func NewVectorChefTransactor(address common.Address, transactor bind.ContractTransactor) (*VectorChefTransactor, error) {
	contract, err := bindVectorChef(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VectorChefTransactor{contract: contract}, nil
}

// NewVectorChefFilterer creates a new log filterer instance of VectorChef, bound to a specific deployed contract.
func NewVectorChefFilterer(address common.Address, filterer bind.ContractFilterer) (*VectorChefFilterer, error) {
	contract, err := bindVectorChef(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VectorChefFilterer{contract: contract}, nil
}

// bindVectorChef binds a generic wrapper to an already deployed contract.
func bindVectorChef(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VectorChefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VectorChef *VectorChefRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VectorChef.Contract.VectorChefCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VectorChef *VectorChefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorChef.Contract.VectorChefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VectorChef *VectorChefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VectorChef.Contract.VectorChefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VectorChef *VectorChefCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VectorChef.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VectorChef *VectorChefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorChef.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VectorChef *VectorChefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VectorChef.Contract.contract.Transact(opts, method, params...)
}

// PoolManagers is a free data retrieval call binding the contract method 0x07337f2b.
//
// Solidity: function PoolManagers(address ) view returns(bool)
func (_VectorChef *VectorChefCaller) PoolManagers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "PoolManagers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolManagers is a free data retrieval call binding the contract method 0x07337f2b.
//
// Solidity: function PoolManagers(address ) view returns(bool)
func (_VectorChef *VectorChefSession) PoolManagers(arg0 common.Address) (bool, error) {
	return _VectorChef.Contract.PoolManagers(&_VectorChef.CallOpts, arg0)
}

// PoolManagers is a free data retrieval call binding the contract method 0x07337f2b.
//
// Solidity: function PoolManagers(address ) view returns(bool)
func (_VectorChef *VectorChefCallerSession) PoolManagers(arg0 common.Address) (bool, error) {
	return _VectorChef.Contract.PoolManagers(&_VectorChef.CallOpts, arg0)
}

// AddressToPoolInfo is a free data retrieval call binding the contract method 0x75b1aadc.
//
// Solidity: function addressToPoolInfo(address ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardTimestamp, uint256 accVTXPerShare, address rewarder, address helper, address locker)
func (_VectorChef *VectorChefCaller) AddressToPoolInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardTimestamp *big.Int
	AccVTXPerShare      *big.Int
	Rewarder            common.Address
	Helper              common.Address
	Locker              common.Address
}, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "addressToPoolInfo", arg0)

	outstruct := new(struct {
		LpToken             common.Address
		AllocPoint          *big.Int
		LastRewardTimestamp *big.Int
		AccVTXPerShare      *big.Int
		Rewarder            common.Address
		Helper              common.Address
		Locker              common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccVTXPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Rewarder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Helper = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Locker = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AddressToPoolInfo is a free data retrieval call binding the contract method 0x75b1aadc.
//
// Solidity: function addressToPoolInfo(address ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardTimestamp, uint256 accVTXPerShare, address rewarder, address helper, address locker)
func (_VectorChef *VectorChefSession) AddressToPoolInfo(arg0 common.Address) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardTimestamp *big.Int
	AccVTXPerShare      *big.Int
	Rewarder            common.Address
	Helper              common.Address
	Locker              common.Address
}, error) {
	return _VectorChef.Contract.AddressToPoolInfo(&_VectorChef.CallOpts, arg0)
}

// AddressToPoolInfo is a free data retrieval call binding the contract method 0x75b1aadc.
//
// Solidity: function addressToPoolInfo(address ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardTimestamp, uint256 accVTXPerShare, address rewarder, address helper, address locker)
func (_VectorChef *VectorChefCallerSession) AddressToPoolInfo(arg0 common.Address) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardTimestamp *big.Int
	AccVTXPerShare      *big.Int
	Rewarder            common.Address
	Helper              common.Address
	Locker              common.Address
}, error) {
	return _VectorChef.Contract.AddressToPoolInfo(&_VectorChef.CallOpts, arg0)
}

// DepositInfo is a free data retrieval call binding the contract method 0x493b0170.
//
// Solidity: function depositInfo(address _lp, address _user) view returns(uint256 availableAmount)
func (_VectorChef *VectorChefCaller) DepositInfo(opts *bind.CallOpts, _lp common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "depositInfo", _lp, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositInfo is a free data retrieval call binding the contract method 0x493b0170.
//
// Solidity: function depositInfo(address _lp, address _user) view returns(uint256 availableAmount)
func (_VectorChef *VectorChefSession) DepositInfo(_lp common.Address, _user common.Address) (*big.Int, error) {
	return _VectorChef.Contract.DepositInfo(&_VectorChef.CallOpts, _lp, _user)
}

// DepositInfo is a free data retrieval call binding the contract method 0x493b0170.
//
// Solidity: function depositInfo(address _lp, address _user) view returns(uint256 availableAmount)
func (_VectorChef *VectorChefCallerSession) DepositInfo(_lp common.Address, _user common.Address) (*big.Int, error) {
	return _VectorChef.Contract.DepositInfo(&_VectorChef.CallOpts, _lp, _user)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x06bfa938.
//
// Solidity: function getPoolInfo(address token) view returns(uint256 emission, uint256 allocpoint, uint256 sizeOfPool, uint256 totalPoint)
func (_VectorChef *VectorChefCaller) GetPoolInfo(opts *bind.CallOpts, token common.Address) (struct {
	Emission   *big.Int
	Allocpoint *big.Int
	SizeOfPool *big.Int
	TotalPoint *big.Int
}, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "getPoolInfo", token)

	outstruct := new(struct {
		Emission   *big.Int
		Allocpoint *big.Int
		SizeOfPool *big.Int
		TotalPoint *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Emission = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Allocpoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SizeOfPool = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalPoint = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x06bfa938.
//
// Solidity: function getPoolInfo(address token) view returns(uint256 emission, uint256 allocpoint, uint256 sizeOfPool, uint256 totalPoint)
func (_VectorChef *VectorChefSession) GetPoolInfo(token common.Address) (struct {
	Emission   *big.Int
	Allocpoint *big.Int
	SizeOfPool *big.Int
	TotalPoint *big.Int
}, error) {
	return _VectorChef.Contract.GetPoolInfo(&_VectorChef.CallOpts, token)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x06bfa938.
//
// Solidity: function getPoolInfo(address token) view returns(uint256 emission, uint256 allocpoint, uint256 sizeOfPool, uint256 totalPoint)
func (_VectorChef *VectorChefCallerSession) GetPoolInfo(token common.Address) (struct {
	Emission   *big.Int
	Allocpoint *big.Int
	SizeOfPool *big.Int
	TotalPoint *big.Int
}, error) {
	return _VectorChef.Contract.GetPoolInfo(&_VectorChef.CallOpts, token)
}

// IsAuthorizedForLock is a free data retrieval call binding the contract method 0x7e29b1a6.
//
// Solidity: function isAuthorizedForLock(address ) view returns(bool)
func (_VectorChef *VectorChefCaller) IsAuthorizedForLock(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "isAuthorizedForLock", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedForLock is a free data retrieval call binding the contract method 0x7e29b1a6.
//
// Solidity: function isAuthorizedForLock(address ) view returns(bool)
func (_VectorChef *VectorChefSession) IsAuthorizedForLock(arg0 common.Address) (bool, error) {
	return _VectorChef.Contract.IsAuthorizedForLock(&_VectorChef.CallOpts, arg0)
}

// IsAuthorizedForLock is a free data retrieval call binding the contract method 0x7e29b1a6.
//
// Solidity: function isAuthorizedForLock(address ) view returns(bool)
func (_VectorChef *VectorChefCallerSession) IsAuthorizedForLock(arg0 common.Address) (bool, error) {
	return _VectorChef.Contract.IsAuthorizedForLock(&_VectorChef.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VectorChef *VectorChefCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VectorChef *VectorChefSession) Owner() (common.Address, error) {
	return _VectorChef.Contract.Owner(&_VectorChef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VectorChef *VectorChefCallerSession) Owner() (common.Address, error) {
	return _VectorChef.Contract.Owner(&_VectorChef.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xad05e627.
//
// Solidity: function pendingTokens(address _lp, address _user, address token) view returns(uint256 pendingVTX, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_VectorChef *VectorChefCaller) PendingTokens(opts *bind.CallOpts, _lp common.Address, _user common.Address, token common.Address) (struct {
	PendingVTX        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "pendingTokens", _lp, _user, token)

	outstruct := new(struct {
		PendingVTX        *big.Int
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
		PendingBonusToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingVTX = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BonusTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PendingBonusToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xad05e627.
//
// Solidity: function pendingTokens(address _lp, address _user, address token) view returns(uint256 pendingVTX, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_VectorChef *VectorChefSession) PendingTokens(_lp common.Address, _user common.Address, token common.Address) (struct {
	PendingVTX        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _VectorChef.Contract.PendingTokens(&_VectorChef.CallOpts, _lp, _user, token)
}

// PendingTokens is a free data retrieval call binding the contract method 0xad05e627.
//
// Solidity: function pendingTokens(address _lp, address _user, address token) view returns(uint256 pendingVTX, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_VectorChef *VectorChefCallerSession) PendingTokens(_lp common.Address, _user common.Address, token common.Address) (struct {
	PendingVTX        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _VectorChef.Contract.PendingTokens(&_VectorChef.CallOpts, _lp, _user, token)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_VectorChef *VectorChefCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_VectorChef *VectorChefSession) PoolLength() (*big.Int, error) {
	return _VectorChef.Contract.PoolLength(&_VectorChef.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_VectorChef *VectorChefCallerSession) PoolLength() (*big.Int, error) {
	return _VectorChef.Contract.PoolLength(&_VectorChef.CallOpts)
}

// RegisteredToken is a free data retrieval call binding the contract method 0x9c7e2655.
//
// Solidity: function registeredToken(uint256 ) view returns(address)
func (_VectorChef *VectorChefCaller) RegisteredToken(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "registeredToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegisteredToken is a free data retrieval call binding the contract method 0x9c7e2655.
//
// Solidity: function registeredToken(uint256 ) view returns(address)
func (_VectorChef *VectorChefSession) RegisteredToken(arg0 *big.Int) (common.Address, error) {
	return _VectorChef.Contract.RegisteredToken(&_VectorChef.CallOpts, arg0)
}

// RegisteredToken is a free data retrieval call binding the contract method 0x9c7e2655.
//
// Solidity: function registeredToken(uint256 ) view returns(address)
func (_VectorChef *VectorChefCallerSession) RegisteredToken(arg0 *big.Int) (common.Address, error) {
	return _VectorChef.Contract.RegisteredToken(&_VectorChef.CallOpts, arg0)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_VectorChef *VectorChefCaller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "startTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_VectorChef *VectorChefSession) StartTimestamp() (*big.Int, error) {
	return _VectorChef.Contract.StartTimestamp(&_VectorChef.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_VectorChef *VectorChefCallerSession) StartTimestamp() (*big.Int, error) {
	return _VectorChef.Contract.StartTimestamp(&_VectorChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_VectorChef *VectorChefCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_VectorChef *VectorChefSession) TotalAllocPoint() (*big.Int, error) {
	return _VectorChef.Contract.TotalAllocPoint(&_VectorChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_VectorChef *VectorChefCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _VectorChef.Contract.TotalAllocPoint(&_VectorChef.CallOpts)
}

// Vtx is a free data retrieval call binding the contract method 0xb76ae05e.
//
// Solidity: function vtx() view returns(address)
func (_VectorChef *VectorChefCaller) Vtx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "vtx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vtx is a free data retrieval call binding the contract method 0xb76ae05e.
//
// Solidity: function vtx() view returns(address)
func (_VectorChef *VectorChefSession) Vtx() (common.Address, error) {
	return _VectorChef.Contract.Vtx(&_VectorChef.CallOpts)
}

// Vtx is a free data retrieval call binding the contract method 0xb76ae05e.
//
// Solidity: function vtx() view returns(address)
func (_VectorChef *VectorChefCallerSession) Vtx() (common.Address, error) {
	return _VectorChef.Contract.Vtx(&_VectorChef.CallOpts)
}

// VtxPerSec is a free data retrieval call binding the contract method 0xd6c494c8.
//
// Solidity: function vtxPerSec() view returns(uint256)
func (_VectorChef *VectorChefCaller) VtxPerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VectorChef.contract.Call(opts, &out, "vtxPerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VtxPerSec is a free data retrieval call binding the contract method 0xd6c494c8.
//
// Solidity: function vtxPerSec() view returns(uint256)
func (_VectorChef *VectorChefSession) VtxPerSec() (*big.Int, error) {
	return _VectorChef.Contract.VtxPerSec(&_VectorChef.CallOpts)
}

// VtxPerSec is a free data retrieval call binding the contract method 0xd6c494c8.
//
// Solidity: function vtxPerSec() view returns(uint256)
func (_VectorChef *VectorChefCallerSession) VtxPerSec() (*big.Int, error) {
	return _VectorChef.Contract.VtxPerSec(&_VectorChef.CallOpts)
}

// MasterChefVTXInit is a paid mutator transaction binding the contract method 0x5aec2722.
//
// Solidity: function __MasterChefVTX_init(address _vtx, uint256 _vtxPerSec, uint256 _startTimestamp) returns()
func (_VectorChef *VectorChefTransactor) MasterChefVTXInit(opts *bind.TransactOpts, _vtx common.Address, _vtxPerSec *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "__MasterChefVTX_init", _vtx, _vtxPerSec, _startTimestamp)
}

// MasterChefVTXInit is a paid mutator transaction binding the contract method 0x5aec2722.
//
// Solidity: function __MasterChefVTX_init(address _vtx, uint256 _vtxPerSec, uint256 _startTimestamp) returns()
func (_VectorChef *VectorChefSession) MasterChefVTXInit(_vtx common.Address, _vtxPerSec *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.MasterChefVTXInit(&_VectorChef.TransactOpts, _vtx, _vtxPerSec, _startTimestamp)
}

// MasterChefVTXInit is a paid mutator transaction binding the contract method 0x5aec2722.
//
// Solidity: function __MasterChefVTX_init(address _vtx, uint256 _vtxPerSec, uint256 _startTimestamp) returns()
func (_VectorChef *VectorChefTransactorSession) MasterChefVTXInit(_vtx common.Address, _vtxPerSec *big.Int, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.MasterChefVTXInit(&_VectorChef.TransactOpts, _vtx, _vtxPerSec, _startTimestamp)
}

// Add is a paid mutator transaction binding the contract method 0x266f24b7.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, address _rewarder, address _helper) returns()
func (_VectorChef *VectorChefTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address, _helper common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "add", _allocPoint, _lpToken, _rewarder, _helper)
}

// Add is a paid mutator transaction binding the contract method 0x266f24b7.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, address _rewarder, address _helper) returns()
func (_VectorChef *VectorChefSession) Add(_allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address, _helper common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.Add(&_VectorChef.TransactOpts, _allocPoint, _lpToken, _rewarder, _helper)
}

// Add is a paid mutator transaction binding the contract method 0x266f24b7.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, address _rewarder, address _helper) returns()
func (_VectorChef *VectorChefTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address, _helper common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.Add(&_VectorChef.TransactOpts, _allocPoint, _lpToken, _rewarder, _helper)
}

// AllowEmergency is a paid mutator transaction binding the contract method 0xcdfaeab6.
//
// Solidity: function allowEmergency() returns()
func (_VectorChef *VectorChefTransactor) AllowEmergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "allowEmergency")
}

// AllowEmergency is a paid mutator transaction binding the contract method 0xcdfaeab6.
//
// Solidity: function allowEmergency() returns()
func (_VectorChef *VectorChefSession) AllowEmergency() (*types.Transaction, error) {
	return _VectorChef.Contract.AllowEmergency(&_VectorChef.TransactOpts)
}

// AllowEmergency is a paid mutator transaction binding the contract method 0xcdfaeab6.
//
// Solidity: function allowEmergency() returns()
func (_VectorChef *VectorChefTransactorSession) AllowEmergency() (*types.Transaction, error) {
	return _VectorChef.Contract.AllowEmergency(&_VectorChef.TransactOpts)
}

// AuthorizeForLock is a paid mutator transaction binding the contract method 0x6f80d877.
//
// Solidity: function authorizeForLock(address _address) returns()
func (_VectorChef *VectorChefTransactor) AuthorizeForLock(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "authorizeForLock", _address)
}

// AuthorizeForLock is a paid mutator transaction binding the contract method 0x6f80d877.
//
// Solidity: function authorizeForLock(address _address) returns()
func (_VectorChef *VectorChefSession) AuthorizeForLock(_address common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.AuthorizeForLock(&_VectorChef.TransactOpts, _address)
}

// AuthorizeForLock is a paid mutator transaction binding the contract method 0x6f80d877.
//
// Solidity: function authorizeForLock(address _address) returns()
func (_VectorChef *VectorChefTransactorSession) AuthorizeForLock(_address common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.AuthorizeForLock(&_VectorChef.TransactOpts, _address)
}

// AuthorizeLocker is a paid mutator transaction binding the contract method 0x81c95579.
//
// Solidity: function authorizeLocker(address _locker) returns()
func (_VectorChef *VectorChefTransactor) AuthorizeLocker(opts *bind.TransactOpts, _locker common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "authorizeLocker", _locker)
}

// AuthorizeLocker is a paid mutator transaction binding the contract method 0x81c95579.
//
// Solidity: function authorizeLocker(address _locker) returns()
func (_VectorChef *VectorChefSession) AuthorizeLocker(_locker common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.AuthorizeLocker(&_VectorChef.TransactOpts, _locker)
}

// AuthorizeLocker is a paid mutator transaction binding the contract method 0x81c95579.
//
// Solidity: function authorizeLocker(address _locker) returns()
func (_VectorChef *VectorChefTransactorSession) AuthorizeLocker(_locker common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.AuthorizeLocker(&_VectorChef.TransactOpts, _locker)
}

// ClaimLock is a paid mutator transaction binding the contract method 0x986e8539.
//
// Solidity: function claimLock(address _lp) returns()
func (_VectorChef *VectorChefTransactor) ClaimLock(opts *bind.TransactOpts, _lp common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "claimLock", _lp)
}

// ClaimLock is a paid mutator transaction binding the contract method 0x986e8539.
//
// Solidity: function claimLock(address _lp) returns()
func (_VectorChef *VectorChefSession) ClaimLock(_lp common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.ClaimLock(&_VectorChef.TransactOpts, _lp)
}

// ClaimLock is a paid mutator transaction binding the contract method 0x986e8539.
//
// Solidity: function claimLock(address _lp) returns()
func (_VectorChef *VectorChefTransactorSession) ClaimLock(_lp common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.ClaimLock(&_VectorChef.TransactOpts, _lp)
}

// CreateRewarder is a paid mutator transaction binding the contract method 0x3b3f0ee6.
//
// Solidity: function createRewarder(address _lpToken, address mainRewardToken) returns(address)
func (_VectorChef *VectorChefTransactor) CreateRewarder(opts *bind.TransactOpts, _lpToken common.Address, mainRewardToken common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "createRewarder", _lpToken, mainRewardToken)
}

// CreateRewarder is a paid mutator transaction binding the contract method 0x3b3f0ee6.
//
// Solidity: function createRewarder(address _lpToken, address mainRewardToken) returns(address)
func (_VectorChef *VectorChefSession) CreateRewarder(_lpToken common.Address, mainRewardToken common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.CreateRewarder(&_VectorChef.TransactOpts, _lpToken, mainRewardToken)
}

// CreateRewarder is a paid mutator transaction binding the contract method 0x3b3f0ee6.
//
// Solidity: function createRewarder(address _lpToken, address mainRewardToken) returns(address)
func (_VectorChef *VectorChefTransactorSession) CreateRewarder(_lpToken common.Address, mainRewardToken common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.CreateRewarder(&_VectorChef.TransactOpts, _lpToken, mainRewardToken)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _lp, uint256 _amount) returns()
func (_VectorChef *VectorChefTransactor) Deposit(opts *bind.TransactOpts, _lp common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "deposit", _lp, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _lp, uint256 _amount) returns()
func (_VectorChef *VectorChefSession) Deposit(_lp common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.Deposit(&_VectorChef.TransactOpts, _lp, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _lp, uint256 _amount) returns()
func (_VectorChef *VectorChefTransactorSession) Deposit(_lp common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.Deposit(&_VectorChef.TransactOpts, _lp, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0xc8820f6c.
//
// Solidity: function depositFor(address _lp, uint256 _amount, address sender) returns()
func (_VectorChef *VectorChefTransactor) DepositFor(opts *bind.TransactOpts, _lp common.Address, _amount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "depositFor", _lp, _amount, sender)
}

// DepositFor is a paid mutator transaction binding the contract method 0xc8820f6c.
//
// Solidity: function depositFor(address _lp, uint256 _amount, address sender) returns()
func (_VectorChef *VectorChefSession) DepositFor(_lp common.Address, _amount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.DepositFor(&_VectorChef.TransactOpts, _lp, _amount, sender)
}

// DepositFor is a paid mutator transaction binding the contract method 0xc8820f6c.
//
// Solidity: function depositFor(address _lp, uint256 _amount, address sender) returns()
func (_VectorChef *VectorChefTransactorSession) DepositFor(_lp common.Address, _amount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.DepositFor(&_VectorChef.TransactOpts, _lp, _amount, sender)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address _lp) returns()
func (_VectorChef *VectorChefTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _lp common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "emergencyWithdraw", _lp)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address _lp) returns()
func (_VectorChef *VectorChefSession) EmergencyWithdraw(_lp common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.EmergencyWithdraw(&_VectorChef.TransactOpts, _lp)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address _lp) returns()
func (_VectorChef *VectorChefTransactorSession) EmergencyWithdraw(_lp common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.EmergencyWithdraw(&_VectorChef.TransactOpts, _lp)
}

// EmergencyWithdrawWithReward is a paid mutator transaction binding the contract method 0x97f3b304.
//
// Solidity: function emergencyWithdrawWithReward(address _lp) returns()
func (_VectorChef *VectorChefTransactor) EmergencyWithdrawWithReward(opts *bind.TransactOpts, _lp common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "emergencyWithdrawWithReward", _lp)
}

// EmergencyWithdrawWithReward is a paid mutator transaction binding the contract method 0x97f3b304.
//
// Solidity: function emergencyWithdrawWithReward(address _lp) returns()
func (_VectorChef *VectorChefSession) EmergencyWithdrawWithReward(_lp common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.EmergencyWithdrawWithReward(&_VectorChef.TransactOpts, _lp)
}

// EmergencyWithdrawWithReward is a paid mutator transaction binding the contract method 0x97f3b304.
//
// Solidity: function emergencyWithdrawWithReward(address _lp) returns()
func (_VectorChef *VectorChefTransactorSession) EmergencyWithdrawWithReward(_lp common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.EmergencyWithdrawWithReward(&_VectorChef.TransactOpts, _lp)
}

// Lock is a paid mutator transaction binding the contract method 0x9bf92698.
//
// Solidity: function lock(address _lp, uint256 _amount, uint256 _index, bool force) returns()
func (_VectorChef *VectorChefTransactor) Lock(opts *bind.TransactOpts, _lp common.Address, _amount *big.Int, _index *big.Int, force bool) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "lock", _lp, _amount, _index, force)
}

// Lock is a paid mutator transaction binding the contract method 0x9bf92698.
//
// Solidity: function lock(address _lp, uint256 _amount, uint256 _index, bool force) returns()
func (_VectorChef *VectorChefSession) Lock(_lp common.Address, _amount *big.Int, _index *big.Int, force bool) (*types.Transaction, error) {
	return _VectorChef.Contract.Lock(&_VectorChef.TransactOpts, _lp, _amount, _index, force)
}

// Lock is a paid mutator transaction binding the contract method 0x9bf92698.
//
// Solidity: function lock(address _lp, uint256 _amount, uint256 _index, bool force) returns()
func (_VectorChef *VectorChefTransactorSession) Lock(_lp common.Address, _amount *big.Int, _index *big.Int, force bool) (*types.Transaction, error) {
	return _VectorChef.Contract.Lock(&_VectorChef.TransactOpts, _lp, _amount, _index, force)
}

// LockFor is a paid mutator transaction binding the contract method 0x7ce92844.
//
// Solidity: function lockFor(address _lp, uint256 _amount, uint256 _index, address _for, bool force) returns()
func (_VectorChef *VectorChefTransactor) LockFor(opts *bind.TransactOpts, _lp common.Address, _amount *big.Int, _index *big.Int, _for common.Address, force bool) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "lockFor", _lp, _amount, _index, _for, force)
}

// LockFor is a paid mutator transaction binding the contract method 0x7ce92844.
//
// Solidity: function lockFor(address _lp, uint256 _amount, uint256 _index, address _for, bool force) returns()
func (_VectorChef *VectorChefSession) LockFor(_lp common.Address, _amount *big.Int, _index *big.Int, _for common.Address, force bool) (*types.Transaction, error) {
	return _VectorChef.Contract.LockFor(&_VectorChef.TransactOpts, _lp, _amount, _index, _for, force)
}

// LockFor is a paid mutator transaction binding the contract method 0x7ce92844.
//
// Solidity: function lockFor(address _lp, uint256 _amount, uint256 _index, address _for, bool force) returns()
func (_VectorChef *VectorChefTransactorSession) LockFor(_lp common.Address, _amount *big.Int, _index *big.Int, _for common.Address, force bool) (*types.Transaction, error) {
	return _VectorChef.Contract.LockFor(&_VectorChef.TransactOpts, _lp, _amount, _index, _for, force)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_VectorChef *VectorChefTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_VectorChef *VectorChefSession) MassUpdatePools() (*types.Transaction, error) {
	return _VectorChef.Contract.MassUpdatePools(&_VectorChef.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_VectorChef *VectorChefTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _VectorChef.Contract.MassUpdatePools(&_VectorChef.TransactOpts)
}

// MultiUnlock is a paid mutator transaction binding the contract method 0x3e8d54e1.
//
// Solidity: function multiUnlock(address _lp, uint256[] _amount, uint256[] _index) returns()
func (_VectorChef *VectorChefTransactor) MultiUnlock(opts *bind.TransactOpts, _lp common.Address, _amount []*big.Int, _index []*big.Int) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "multiUnlock", _lp, _amount, _index)
}

// MultiUnlock is a paid mutator transaction binding the contract method 0x3e8d54e1.
//
// Solidity: function multiUnlock(address _lp, uint256[] _amount, uint256[] _index) returns()
func (_VectorChef *VectorChefSession) MultiUnlock(_lp common.Address, _amount []*big.Int, _index []*big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.MultiUnlock(&_VectorChef.TransactOpts, _lp, _amount, _index)
}

// MultiUnlock is a paid mutator transaction binding the contract method 0x3e8d54e1.
//
// Solidity: function multiUnlock(address _lp, uint256[] _amount, uint256[] _index) returns()
func (_VectorChef *VectorChefTransactorSession) MultiUnlock(_lp common.Address, _amount []*big.Int, _index []*big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.MultiUnlock(&_VectorChef.TransactOpts, _lp, _amount, _index)
}

// Multiclaim is a paid mutator transaction binding the contract method 0x17e54830.
//
// Solidity: function multiclaim(address[] _lps, address user_address) returns()
func (_VectorChef *VectorChefTransactor) Multiclaim(opts *bind.TransactOpts, _lps []common.Address, user_address common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "multiclaim", _lps, user_address)
}

// Multiclaim is a paid mutator transaction binding the contract method 0x17e54830.
//
// Solidity: function multiclaim(address[] _lps, address user_address) returns()
func (_VectorChef *VectorChefSession) Multiclaim(_lps []common.Address, user_address common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.Multiclaim(&_VectorChef.TransactOpts, _lps, user_address)
}

// Multiclaim is a paid mutator transaction binding the contract method 0x17e54830.
//
// Solidity: function multiclaim(address[] _lps, address user_address) returns()
func (_VectorChef *VectorChefTransactorSession) Multiclaim(_lps []common.Address, user_address common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.Multiclaim(&_VectorChef.TransactOpts, _lps, user_address)
}

// RealEmergencyWithdraw is a paid mutator transaction binding the contract method 0xe952fddc.
//
// Solidity: function realEmergencyWithdraw(address _lp) returns()
func (_VectorChef *VectorChefTransactor) RealEmergencyWithdraw(opts *bind.TransactOpts, _lp common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "realEmergencyWithdraw", _lp)
}

// RealEmergencyWithdraw is a paid mutator transaction binding the contract method 0xe952fddc.
//
// Solidity: function realEmergencyWithdraw(address _lp) returns()
func (_VectorChef *VectorChefSession) RealEmergencyWithdraw(_lp common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.RealEmergencyWithdraw(&_VectorChef.TransactOpts, _lp)
}

// RealEmergencyWithdraw is a paid mutator transaction binding the contract method 0xe952fddc.
//
// Solidity: function realEmergencyWithdraw(address _lp) returns()
func (_VectorChef *VectorChefTransactorSession) RealEmergencyWithdraw(_lp common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.RealEmergencyWithdraw(&_VectorChef.TransactOpts, _lp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VectorChef *VectorChefTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VectorChef *VectorChefSession) RenounceOwnership() (*types.Transaction, error) {
	return _VectorChef.Contract.RenounceOwnership(&_VectorChef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VectorChef *VectorChefTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VectorChef.Contract.RenounceOwnership(&_VectorChef.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x56a2ff68.
//
// Solidity: function set(address _lp, uint256 _allocPoint, address _rewarder, address _locker, bool overwrite) returns()
func (_VectorChef *VectorChefTransactor) Set(opts *bind.TransactOpts, _lp common.Address, _allocPoint *big.Int, _rewarder common.Address, _locker common.Address, overwrite bool) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "set", _lp, _allocPoint, _rewarder, _locker, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x56a2ff68.
//
// Solidity: function set(address _lp, uint256 _allocPoint, address _rewarder, address _locker, bool overwrite) returns()
func (_VectorChef *VectorChefSession) Set(_lp common.Address, _allocPoint *big.Int, _rewarder common.Address, _locker common.Address, overwrite bool) (*types.Transaction, error) {
	return _VectorChef.Contract.Set(&_VectorChef.TransactOpts, _lp, _allocPoint, _rewarder, _locker, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x56a2ff68.
//
// Solidity: function set(address _lp, uint256 _allocPoint, address _rewarder, address _locker, bool overwrite) returns()
func (_VectorChef *VectorChefTransactorSession) Set(_lp common.Address, _allocPoint *big.Int, _rewarder common.Address, _locker common.Address, overwrite bool) (*types.Transaction, error) {
	return _VectorChef.Contract.Set(&_VectorChef.TransactOpts, _lp, _allocPoint, _rewarder, _locker, overwrite)
}

// SetPoolHelper is a paid mutator transaction binding the contract method 0xa2212459.
//
// Solidity: function setPoolHelper(address _lp, address _helper) returns()
func (_VectorChef *VectorChefTransactor) SetPoolHelper(opts *bind.TransactOpts, _lp common.Address, _helper common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "setPoolHelper", _lp, _helper)
}

// SetPoolHelper is a paid mutator transaction binding the contract method 0xa2212459.
//
// Solidity: function setPoolHelper(address _lp, address _helper) returns()
func (_VectorChef *VectorChefSession) SetPoolHelper(_lp common.Address, _helper common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.SetPoolHelper(&_VectorChef.TransactOpts, _lp, _helper)
}

// SetPoolHelper is a paid mutator transaction binding the contract method 0xa2212459.
//
// Solidity: function setPoolHelper(address _lp, address _helper) returns()
func (_VectorChef *VectorChefTransactorSession) SetPoolHelper(_lp common.Address, _helper common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.SetPoolHelper(&_VectorChef.TransactOpts, _lp, _helper)
}

// SetPoolManagerStatus is a paid mutator transaction binding the contract method 0xefe33cfa.
//
// Solidity: function setPoolManagerStatus(address _address, bool _bool) returns()
func (_VectorChef *VectorChefTransactor) SetPoolManagerStatus(opts *bind.TransactOpts, _address common.Address, _bool bool) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "setPoolManagerStatus", _address, _bool)
}

// SetPoolManagerStatus is a paid mutator transaction binding the contract method 0xefe33cfa.
//
// Solidity: function setPoolManagerStatus(address _address, bool _bool) returns()
func (_VectorChef *VectorChefSession) SetPoolManagerStatus(_address common.Address, _bool bool) (*types.Transaction, error) {
	return _VectorChef.Contract.SetPoolManagerStatus(&_VectorChef.TransactOpts, _address, _bool)
}

// SetPoolManagerStatus is a paid mutator transaction binding the contract method 0xefe33cfa.
//
// Solidity: function setPoolManagerStatus(address _address, bool _bool) returns()
func (_VectorChef *VectorChefTransactorSession) SetPoolManagerStatus(_address common.Address, _bool bool) (*types.Transaction, error) {
	return _VectorChef.Contract.SetPoolManagerStatus(&_VectorChef.TransactOpts, _address, _bool)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VectorChef *VectorChefTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VectorChef *VectorChefSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.TransferOwnership(&_VectorChef.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VectorChef *VectorChefTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.TransferOwnership(&_VectorChef.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xb5674c68.
//
// Solidity: function unlock(address _lp, uint256 _amount, uint256 _index) returns()
func (_VectorChef *VectorChefTransactor) Unlock(opts *bind.TransactOpts, _lp common.Address, _amount *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "unlock", _lp, _amount, _index)
}

// Unlock is a paid mutator transaction binding the contract method 0xb5674c68.
//
// Solidity: function unlock(address _lp, uint256 _amount, uint256 _index) returns()
func (_VectorChef *VectorChefSession) Unlock(_lp common.Address, _amount *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.Unlock(&_VectorChef.TransactOpts, _lp, _amount, _index)
}

// Unlock is a paid mutator transaction binding the contract method 0xb5674c68.
//
// Solidity: function unlock(address _lp, uint256 _amount, uint256 _index) returns()
func (_VectorChef *VectorChefTransactorSession) Unlock(_lp common.Address, _amount *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.Unlock(&_VectorChef.TransactOpts, _lp, _amount, _index)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _vtxPerSec) returns()
func (_VectorChef *VectorChefTransactor) UpdateEmissionRate(opts *bind.TransactOpts, _vtxPerSec *big.Int) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "updateEmissionRate", _vtxPerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _vtxPerSec) returns()
func (_VectorChef *VectorChefSession) UpdateEmissionRate(_vtxPerSec *big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.UpdateEmissionRate(&_VectorChef.TransactOpts, _vtxPerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _vtxPerSec) returns()
func (_VectorChef *VectorChefTransactorSession) UpdateEmissionRate(_vtxPerSec *big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.UpdateEmissionRate(&_VectorChef.TransactOpts, _vtxPerSec)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x7b46c54f.
//
// Solidity: function updatePool(address _lp) returns()
func (_VectorChef *VectorChefTransactor) UpdatePool(opts *bind.TransactOpts, _lp common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "updatePool", _lp)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x7b46c54f.
//
// Solidity: function updatePool(address _lp) returns()
func (_VectorChef *VectorChefSession) UpdatePool(_lp common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.UpdatePool(&_VectorChef.TransactOpts, _lp)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x7b46c54f.
//
// Solidity: function updatePool(address _lp) returns()
func (_VectorChef *VectorChefTransactorSession) UpdatePool(_lp common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.UpdatePool(&_VectorChef.TransactOpts, _lp)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _lp, uint256 _amount) returns()
func (_VectorChef *VectorChefTransactor) Withdraw(opts *bind.TransactOpts, _lp common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "withdraw", _lp, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _lp, uint256 _amount) returns()
func (_VectorChef *VectorChefSession) Withdraw(_lp common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.Withdraw(&_VectorChef.TransactOpts, _lp, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _lp, uint256 _amount) returns()
func (_VectorChef *VectorChefTransactorSession) Withdraw(_lp common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VectorChef.Contract.Withdraw(&_VectorChef.TransactOpts, _lp, _amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x6766efbe.
//
// Solidity: function withdrawFor(address _lp, uint256 _amount, address _sender) returns()
func (_VectorChef *VectorChefTransactor) WithdrawFor(opts *bind.TransactOpts, _lp common.Address, _amount *big.Int, _sender common.Address) (*types.Transaction, error) {
	return _VectorChef.contract.Transact(opts, "withdrawFor", _lp, _amount, _sender)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x6766efbe.
//
// Solidity: function withdrawFor(address _lp, uint256 _amount, address _sender) returns()
func (_VectorChef *VectorChefSession) WithdrawFor(_lp common.Address, _amount *big.Int, _sender common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.WithdrawFor(&_VectorChef.TransactOpts, _lp, _amount, _sender)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0x6766efbe.
//
// Solidity: function withdrawFor(address _lp, uint256 _amount, address _sender) returns()
func (_VectorChef *VectorChefTransactorSession) WithdrawFor(_lp common.Address, _amount *big.Int, _sender common.Address) (*types.Transaction, error) {
	return _VectorChef.Contract.WithdrawFor(&_VectorChef.TransactOpts, _lp, _amount, _sender)
}
