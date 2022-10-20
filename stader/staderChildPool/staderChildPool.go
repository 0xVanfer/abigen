// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staderChildPool

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

// IChildPoolMaticXSwapRequest is an auto generated low-level Go binding around an user-defined struct.
type IChildPoolMaticXSwapRequest struct {
	Amount         *big.Int
	RequestTime    *big.Int
	WithdrawalTime *big.Int
}

// StaderChildPoolMetaData contains all meta data concerning the StaderChildPool contract.
var StaderChildPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"claimMaticXSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"convertMaticToMaticX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"convertMaticXToMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getAmountAfterInstantWithdrawalFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_fxStateChildTunnel\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_maticX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaticXSwapLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserMaticXSwapRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalTime\",\"type\":\"uint256\"}],\"internalType\":\"structIChildPool.MaticXSwapRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"instantPoolMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"instantPoolMaticX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"instantPoolOwner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"instantWithdrawalFeeBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"instantWithdrawalFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maticXSwapLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provideInstantPoolMatic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"provideInstantPoolMaticX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"requestMaticXSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setFxStateChildTunnel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setInstantPoolOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeBps\",\"type\":\"uint256\"}],\"name\":\"setInstantWithdrawalFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"}],\"name\":\"setMaticXSwapLockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapMaticForMaticXViaInstantPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"swapMaticXForMaticViaInstantPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawInstantPoolMatic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawInstantPoolMaticX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawInstantWithdrawalFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StaderChildPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use StaderChildPoolMetaData.ABI instead.
var StaderChildPoolABI = StaderChildPoolMetaData.ABI

// StaderChildPool is an auto generated Go binding around an Ethereum contract.
type StaderChildPool struct {
	StaderChildPoolCaller     // Read-only binding to the contract
	StaderChildPoolTransactor // Write-only binding to the contract
	StaderChildPoolFilterer   // Log filterer for contract events
}

// StaderChildPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaderChildPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderChildPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaderChildPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderChildPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaderChildPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderChildPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaderChildPoolSession struct {
	Contract     *StaderChildPool  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaderChildPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaderChildPoolCallerSession struct {
	Contract *StaderChildPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StaderChildPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaderChildPoolTransactorSession struct {
	Contract     *StaderChildPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StaderChildPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaderChildPoolRaw struct {
	Contract *StaderChildPool // Generic contract binding to access the raw methods on
}

// StaderChildPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaderChildPoolCallerRaw struct {
	Contract *StaderChildPoolCaller // Generic read-only contract binding to access the raw methods on
}

// StaderChildPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaderChildPoolTransactorRaw struct {
	Contract *StaderChildPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaderChildPool creates a new instance of StaderChildPool, bound to a specific deployed contract.
func NewStaderChildPool(address common.Address, backend bind.ContractBackend) (*StaderChildPool, error) {
	contract, err := bindStaderChildPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaderChildPool{StaderChildPoolCaller: StaderChildPoolCaller{contract: contract}, StaderChildPoolTransactor: StaderChildPoolTransactor{contract: contract}, StaderChildPoolFilterer: StaderChildPoolFilterer{contract: contract}}, nil
}

// NewStaderChildPoolCaller creates a new read-only instance of StaderChildPool, bound to a specific deployed contract.
func NewStaderChildPoolCaller(address common.Address, caller bind.ContractCaller) (*StaderChildPoolCaller, error) {
	contract, err := bindStaderChildPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaderChildPoolCaller{contract: contract}, nil
}

// NewStaderChildPoolTransactor creates a new write-only instance of StaderChildPool, bound to a specific deployed contract.
func NewStaderChildPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*StaderChildPoolTransactor, error) {
	contract, err := bindStaderChildPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaderChildPoolTransactor{contract: contract}, nil
}

// NewStaderChildPoolFilterer creates a new log filterer instance of StaderChildPool, bound to a specific deployed contract.
func NewStaderChildPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*StaderChildPoolFilterer, error) {
	contract, err := bindStaderChildPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaderChildPoolFilterer{contract: contract}, nil
}

// bindStaderChildPool binds a generic wrapper to an already deployed contract.
func bindStaderChildPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StaderChildPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderChildPool *StaderChildPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderChildPool.Contract.StaderChildPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderChildPool *StaderChildPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderChildPool.Contract.StaderChildPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderChildPool *StaderChildPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderChildPool.Contract.StaderChildPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderChildPool *StaderChildPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderChildPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderChildPool *StaderChildPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderChildPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderChildPool *StaderChildPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderChildPool.Contract.contract.Transact(opts, method, params...)
}

// ClaimedMatic is a free data retrieval call binding the contract method 0x1dd5d34c.
//
// Solidity: function claimedMatic() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCaller) ClaimedMatic(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "claimedMatic")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedMatic is a free data retrieval call binding the contract method 0x1dd5d34c.
//
// Solidity: function claimedMatic() view returns(uint256)
func (_StaderChildPool *StaderChildPoolSession) ClaimedMatic() (*big.Int, error) {
	return _StaderChildPool.Contract.ClaimedMatic(&_StaderChildPool.CallOpts)
}

// ClaimedMatic is a free data retrieval call binding the contract method 0x1dd5d34c.
//
// Solidity: function claimedMatic() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCallerSession) ClaimedMatic() (*big.Int, error) {
	return _StaderChildPool.Contract.ClaimedMatic(&_StaderChildPool.CallOpts)
}

// ConvertMaticToMaticX is a free data retrieval call binding the contract method 0x9683e28e.
//
// Solidity: function convertMaticToMaticX(uint256 _balance) view returns(uint256, uint256, uint256)
func (_StaderChildPool *StaderChildPoolCaller) ConvertMaticToMaticX(opts *bind.CallOpts, _balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "convertMaticToMaticX", _balance)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// ConvertMaticToMaticX is a free data retrieval call binding the contract method 0x9683e28e.
//
// Solidity: function convertMaticToMaticX(uint256 _balance) view returns(uint256, uint256, uint256)
func (_StaderChildPool *StaderChildPoolSession) ConvertMaticToMaticX(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _StaderChildPool.Contract.ConvertMaticToMaticX(&_StaderChildPool.CallOpts, _balance)
}

// ConvertMaticToMaticX is a free data retrieval call binding the contract method 0x9683e28e.
//
// Solidity: function convertMaticToMaticX(uint256 _balance) view returns(uint256, uint256, uint256)
func (_StaderChildPool *StaderChildPoolCallerSession) ConvertMaticToMaticX(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _StaderChildPool.Contract.ConvertMaticToMaticX(&_StaderChildPool.CallOpts, _balance)
}

// ConvertMaticXToMatic is a free data retrieval call binding the contract method 0x75a85ef5.
//
// Solidity: function convertMaticXToMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_StaderChildPool *StaderChildPoolCaller) ConvertMaticXToMatic(opts *bind.CallOpts, _balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "convertMaticXToMatic", _balance)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// ConvertMaticXToMatic is a free data retrieval call binding the contract method 0x75a85ef5.
//
// Solidity: function convertMaticXToMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_StaderChildPool *StaderChildPoolSession) ConvertMaticXToMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _StaderChildPool.Contract.ConvertMaticXToMatic(&_StaderChildPool.CallOpts, _balance)
}

// ConvertMaticXToMatic is a free data retrieval call binding the contract method 0x75a85ef5.
//
// Solidity: function convertMaticXToMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_StaderChildPool *StaderChildPoolCallerSession) ConvertMaticXToMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _StaderChildPool.Contract.ConvertMaticXToMatic(&_StaderChildPool.CallOpts, _balance)
}

// GetAmountAfterInstantWithdrawalFees is a free data retrieval call binding the contract method 0xe9ce0532.
//
// Solidity: function getAmountAfterInstantWithdrawalFees(uint256 _amount) view returns(uint256, uint256)
func (_StaderChildPool *StaderChildPoolCaller) GetAmountAfterInstantWithdrawalFees(opts *bind.CallOpts, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "getAmountAfterInstantWithdrawalFees", _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAmountAfterInstantWithdrawalFees is a free data retrieval call binding the contract method 0xe9ce0532.
//
// Solidity: function getAmountAfterInstantWithdrawalFees(uint256 _amount) view returns(uint256, uint256)
func (_StaderChildPool *StaderChildPoolSession) GetAmountAfterInstantWithdrawalFees(_amount *big.Int) (*big.Int, *big.Int, error) {
	return _StaderChildPool.Contract.GetAmountAfterInstantWithdrawalFees(&_StaderChildPool.CallOpts, _amount)
}

// GetAmountAfterInstantWithdrawalFees is a free data retrieval call binding the contract method 0xe9ce0532.
//
// Solidity: function getAmountAfterInstantWithdrawalFees(uint256 _amount) view returns(uint256, uint256)
func (_StaderChildPool *StaderChildPoolCallerSession) GetAmountAfterInstantWithdrawalFees(_amount *big.Int) (*big.Int, *big.Int, error) {
	return _StaderChildPool.Contract.GetAmountAfterInstantWithdrawalFees(&_StaderChildPool.CallOpts, _amount)
}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns(address _fxStateChildTunnel, address _maticX, address _trustedForwarder)
func (_StaderChildPool *StaderChildPoolCaller) GetContracts(opts *bind.CallOpts) (struct {
	FxStateChildTunnel common.Address
	MaticX             common.Address
	TrustedForwarder   common.Address
}, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "getContracts")

	outstruct := new(struct {
		FxStateChildTunnel common.Address
		MaticX             common.Address
		TrustedForwarder   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FxStateChildTunnel = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MaticX = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TrustedForwarder = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns(address _fxStateChildTunnel, address _maticX, address _trustedForwarder)
func (_StaderChildPool *StaderChildPoolSession) GetContracts() (struct {
	FxStateChildTunnel common.Address
	MaticX             common.Address
	TrustedForwarder   common.Address
}, error) {
	return _StaderChildPool.Contract.GetContracts(&_StaderChildPool.CallOpts)
}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns(address _fxStateChildTunnel, address _maticX, address _trustedForwarder)
func (_StaderChildPool *StaderChildPoolCallerSession) GetContracts() (struct {
	FxStateChildTunnel common.Address
	MaticX             common.Address
	TrustedForwarder   common.Address
}, error) {
	return _StaderChildPool.Contract.GetContracts(&_StaderChildPool.CallOpts)
}

// GetMaticXSwapLockPeriod is a free data retrieval call binding the contract method 0xec7467ed.
//
// Solidity: function getMaticXSwapLockPeriod() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCaller) GetMaticXSwapLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "getMaticXSwapLockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaticXSwapLockPeriod is a free data retrieval call binding the contract method 0xec7467ed.
//
// Solidity: function getMaticXSwapLockPeriod() view returns(uint256)
func (_StaderChildPool *StaderChildPoolSession) GetMaticXSwapLockPeriod() (*big.Int, error) {
	return _StaderChildPool.Contract.GetMaticXSwapLockPeriod(&_StaderChildPool.CallOpts)
}

// GetMaticXSwapLockPeriod is a free data retrieval call binding the contract method 0xec7467ed.
//
// Solidity: function getMaticXSwapLockPeriod() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCallerSession) GetMaticXSwapLockPeriod() (*big.Int, error) {
	return _StaderChildPool.Contract.GetMaticXSwapLockPeriod(&_StaderChildPool.CallOpts)
}

// GetUserMaticXSwapRequests is a free data retrieval call binding the contract method 0xedb64ec0.
//
// Solidity: function getUserMaticXSwapRequests(address _address) view returns((uint256,uint256,uint256)[])
func (_StaderChildPool *StaderChildPoolCaller) GetUserMaticXSwapRequests(opts *bind.CallOpts, _address common.Address) ([]IChildPoolMaticXSwapRequest, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "getUserMaticXSwapRequests", _address)

	if err != nil {
		return *new([]IChildPoolMaticXSwapRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]IChildPoolMaticXSwapRequest)).(*[]IChildPoolMaticXSwapRequest)

	return out0, err

}

// GetUserMaticXSwapRequests is a free data retrieval call binding the contract method 0xedb64ec0.
//
// Solidity: function getUserMaticXSwapRequests(address _address) view returns((uint256,uint256,uint256)[])
func (_StaderChildPool *StaderChildPoolSession) GetUserMaticXSwapRequests(_address common.Address) ([]IChildPoolMaticXSwapRequest, error) {
	return _StaderChildPool.Contract.GetUserMaticXSwapRequests(&_StaderChildPool.CallOpts, _address)
}

// GetUserMaticXSwapRequests is a free data retrieval call binding the contract method 0xedb64ec0.
//
// Solidity: function getUserMaticXSwapRequests(address _address) view returns((uint256,uint256,uint256)[])
func (_StaderChildPool *StaderChildPoolCallerSession) GetUserMaticXSwapRequests(_address common.Address) ([]IChildPoolMaticXSwapRequest, error) {
	return _StaderChildPool.Contract.GetUserMaticXSwapRequests(&_StaderChildPool.CallOpts, _address)
}

// InstantPoolMatic is a free data retrieval call binding the contract method 0x89dfa025.
//
// Solidity: function instantPoolMatic() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCaller) InstantPoolMatic(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "instantPoolMatic")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InstantPoolMatic is a free data retrieval call binding the contract method 0x89dfa025.
//
// Solidity: function instantPoolMatic() view returns(uint256)
func (_StaderChildPool *StaderChildPoolSession) InstantPoolMatic() (*big.Int, error) {
	return _StaderChildPool.Contract.InstantPoolMatic(&_StaderChildPool.CallOpts)
}

// InstantPoolMatic is a free data retrieval call binding the contract method 0x89dfa025.
//
// Solidity: function instantPoolMatic() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCallerSession) InstantPoolMatic() (*big.Int, error) {
	return _StaderChildPool.Contract.InstantPoolMatic(&_StaderChildPool.CallOpts)
}

// InstantPoolMaticX is a free data retrieval call binding the contract method 0xc759352d.
//
// Solidity: function instantPoolMaticX() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCaller) InstantPoolMaticX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "instantPoolMaticX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InstantPoolMaticX is a free data retrieval call binding the contract method 0xc759352d.
//
// Solidity: function instantPoolMaticX() view returns(uint256)
func (_StaderChildPool *StaderChildPoolSession) InstantPoolMaticX() (*big.Int, error) {
	return _StaderChildPool.Contract.InstantPoolMaticX(&_StaderChildPool.CallOpts)
}

// InstantPoolMaticX is a free data retrieval call binding the contract method 0xc759352d.
//
// Solidity: function instantPoolMaticX() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCallerSession) InstantPoolMaticX() (*big.Int, error) {
	return _StaderChildPool.Contract.InstantPoolMaticX(&_StaderChildPool.CallOpts)
}

// InstantPoolOwner is a free data retrieval call binding the contract method 0x1c083124.
//
// Solidity: function instantPoolOwner() view returns(address)
func (_StaderChildPool *StaderChildPoolCaller) InstantPoolOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "instantPoolOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InstantPoolOwner is a free data retrieval call binding the contract method 0x1c083124.
//
// Solidity: function instantPoolOwner() view returns(address)
func (_StaderChildPool *StaderChildPoolSession) InstantPoolOwner() (common.Address, error) {
	return _StaderChildPool.Contract.InstantPoolOwner(&_StaderChildPool.CallOpts)
}

// InstantPoolOwner is a free data retrieval call binding the contract method 0x1c083124.
//
// Solidity: function instantPoolOwner() view returns(address)
func (_StaderChildPool *StaderChildPoolCallerSession) InstantPoolOwner() (common.Address, error) {
	return _StaderChildPool.Contract.InstantPoolOwner(&_StaderChildPool.CallOpts)
}

// InstantWithdrawalFeeBps is a free data retrieval call binding the contract method 0x1e89a137.
//
// Solidity: function instantWithdrawalFeeBps() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCaller) InstantWithdrawalFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "instantWithdrawalFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InstantWithdrawalFeeBps is a free data retrieval call binding the contract method 0x1e89a137.
//
// Solidity: function instantWithdrawalFeeBps() view returns(uint256)
func (_StaderChildPool *StaderChildPoolSession) InstantWithdrawalFeeBps() (*big.Int, error) {
	return _StaderChildPool.Contract.InstantWithdrawalFeeBps(&_StaderChildPool.CallOpts)
}

// InstantWithdrawalFeeBps is a free data retrieval call binding the contract method 0x1e89a137.
//
// Solidity: function instantWithdrawalFeeBps() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCallerSession) InstantWithdrawalFeeBps() (*big.Int, error) {
	return _StaderChildPool.Contract.InstantWithdrawalFeeBps(&_StaderChildPool.CallOpts)
}

// InstantWithdrawalFees is a free data retrieval call binding the contract method 0x4aa6164d.
//
// Solidity: function instantWithdrawalFees() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCaller) InstantWithdrawalFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "instantWithdrawalFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InstantWithdrawalFees is a free data retrieval call binding the contract method 0x4aa6164d.
//
// Solidity: function instantWithdrawalFees() view returns(uint256)
func (_StaderChildPool *StaderChildPoolSession) InstantWithdrawalFees() (*big.Int, error) {
	return _StaderChildPool.Contract.InstantWithdrawalFees(&_StaderChildPool.CallOpts)
}

// InstantWithdrawalFees is a free data retrieval call binding the contract method 0x4aa6164d.
//
// Solidity: function instantWithdrawalFees() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCallerSession) InstantWithdrawalFees() (*big.Int, error) {
	return _StaderChildPool.Contract.InstantWithdrawalFees(&_StaderChildPool.CallOpts)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address _address) view returns(bool)
func (_StaderChildPool *StaderChildPoolCaller) IsTrustedForwarder(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "isTrustedForwarder", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address _address) view returns(bool)
func (_StaderChildPool *StaderChildPoolSession) IsTrustedForwarder(_address common.Address) (bool, error) {
	return _StaderChildPool.Contract.IsTrustedForwarder(&_StaderChildPool.CallOpts, _address)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address _address) view returns(bool)
func (_StaderChildPool *StaderChildPoolCallerSession) IsTrustedForwarder(_address common.Address) (bool, error) {
	return _StaderChildPool.Contract.IsTrustedForwarder(&_StaderChildPool.CallOpts, _address)
}

// MaticXSwapLockPeriod is a free data retrieval call binding the contract method 0x9ea87b2d.
//
// Solidity: function maticXSwapLockPeriod() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCaller) MaticXSwapLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "maticXSwapLockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaticXSwapLockPeriod is a free data retrieval call binding the contract method 0x9ea87b2d.
//
// Solidity: function maticXSwapLockPeriod() view returns(uint256)
func (_StaderChildPool *StaderChildPoolSession) MaticXSwapLockPeriod() (*big.Int, error) {
	return _StaderChildPool.Contract.MaticXSwapLockPeriod(&_StaderChildPool.CallOpts)
}

// MaticXSwapLockPeriod is a free data retrieval call binding the contract method 0x9ea87b2d.
//
// Solidity: function maticXSwapLockPeriod() view returns(uint256)
func (_StaderChildPool *StaderChildPoolCallerSession) MaticXSwapLockPeriod() (*big.Int, error) {
	return _StaderChildPool.Contract.MaticXSwapLockPeriod(&_StaderChildPool.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_StaderChildPool *StaderChildPoolCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_StaderChildPool *StaderChildPoolSession) Treasury() (common.Address, error) {
	return _StaderChildPool.Contract.Treasury(&_StaderChildPool.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_StaderChildPool *StaderChildPoolCallerSession) Treasury() (common.Address, error) {
	return _StaderChildPool.Contract.Treasury(&_StaderChildPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StaderChildPool *StaderChildPoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StaderChildPool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StaderChildPool *StaderChildPoolSession) Version() (string, error) {
	return _StaderChildPool.Contract.Version(&_StaderChildPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StaderChildPool *StaderChildPoolCallerSession) Version() (string, error) {
	return _StaderChildPool.Contract.Version(&_StaderChildPool.CallOpts)
}

// ClaimMaticXSwap is a paid mutator transaction binding the contract method 0x77baf209.
//
// Solidity: function claimMaticXSwap(uint256 _idx) returns()
func (_StaderChildPool *StaderChildPoolTransactor) ClaimMaticXSwap(opts *bind.TransactOpts, _idx *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "claimMaticXSwap", _idx)
}

// ClaimMaticXSwap is a paid mutator transaction binding the contract method 0x77baf209.
//
// Solidity: function claimMaticXSwap(uint256 _idx) returns()
func (_StaderChildPool *StaderChildPoolSession) ClaimMaticXSwap(_idx *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.ClaimMaticXSwap(&_StaderChildPool.TransactOpts, _idx)
}

// ClaimMaticXSwap is a paid mutator transaction binding the contract method 0x77baf209.
//
// Solidity: function claimMaticXSwap(uint256 _idx) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) ClaimMaticXSwap(_idx *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.ClaimMaticXSwap(&_StaderChildPool.TransactOpts, _idx)
}

// ProvideInstantPoolMatic is a paid mutator transaction binding the contract method 0xe82d73e6.
//
// Solidity: function provideInstantPoolMatic() payable returns()
func (_StaderChildPool *StaderChildPoolTransactor) ProvideInstantPoolMatic(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "provideInstantPoolMatic")
}

// ProvideInstantPoolMatic is a paid mutator transaction binding the contract method 0xe82d73e6.
//
// Solidity: function provideInstantPoolMatic() payable returns()
func (_StaderChildPool *StaderChildPoolSession) ProvideInstantPoolMatic() (*types.Transaction, error) {
	return _StaderChildPool.Contract.ProvideInstantPoolMatic(&_StaderChildPool.TransactOpts)
}

// ProvideInstantPoolMatic is a paid mutator transaction binding the contract method 0xe82d73e6.
//
// Solidity: function provideInstantPoolMatic() payable returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) ProvideInstantPoolMatic() (*types.Transaction, error) {
	return _StaderChildPool.Contract.ProvideInstantPoolMatic(&_StaderChildPool.TransactOpts)
}

// ProvideInstantPoolMaticX is a paid mutator transaction binding the contract method 0x68c05c97.
//
// Solidity: function provideInstantPoolMaticX(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolTransactor) ProvideInstantPoolMaticX(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "provideInstantPoolMaticX", _amount)
}

// ProvideInstantPoolMaticX is a paid mutator transaction binding the contract method 0x68c05c97.
//
// Solidity: function provideInstantPoolMaticX(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolSession) ProvideInstantPoolMaticX(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.ProvideInstantPoolMaticX(&_StaderChildPool.TransactOpts, _amount)
}

// ProvideInstantPoolMaticX is a paid mutator transaction binding the contract method 0x68c05c97.
//
// Solidity: function provideInstantPoolMaticX(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) ProvideInstantPoolMaticX(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.ProvideInstantPoolMaticX(&_StaderChildPool.TransactOpts, _amount)
}

// RequestMaticXSwap is a paid mutator transaction binding the contract method 0x48eaf6d6.
//
// Solidity: function requestMaticXSwap(uint256 _amount) returns(uint256)
func (_StaderChildPool *StaderChildPoolTransactor) RequestMaticXSwap(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "requestMaticXSwap", _amount)
}

// RequestMaticXSwap is a paid mutator transaction binding the contract method 0x48eaf6d6.
//
// Solidity: function requestMaticXSwap(uint256 _amount) returns(uint256)
func (_StaderChildPool *StaderChildPoolSession) RequestMaticXSwap(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.RequestMaticXSwap(&_StaderChildPool.TransactOpts, _amount)
}

// RequestMaticXSwap is a paid mutator transaction binding the contract method 0x48eaf6d6.
//
// Solidity: function requestMaticXSwap(uint256 _amount) returns(uint256)
func (_StaderChildPool *StaderChildPoolTransactorSession) RequestMaticXSwap(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.RequestMaticXSwap(&_StaderChildPool.TransactOpts, _amount)
}

// SetFxStateChildTunnel is a paid mutator transaction binding the contract method 0x174b151c.
//
// Solidity: function setFxStateChildTunnel(address _address) returns()
func (_StaderChildPool *StaderChildPoolTransactor) SetFxStateChildTunnel(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "setFxStateChildTunnel", _address)
}

// SetFxStateChildTunnel is a paid mutator transaction binding the contract method 0x174b151c.
//
// Solidity: function setFxStateChildTunnel(address _address) returns()
func (_StaderChildPool *StaderChildPoolSession) SetFxStateChildTunnel(_address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetFxStateChildTunnel(&_StaderChildPool.TransactOpts, _address)
}

// SetFxStateChildTunnel is a paid mutator transaction binding the contract method 0x174b151c.
//
// Solidity: function setFxStateChildTunnel(address _address) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) SetFxStateChildTunnel(_address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetFxStateChildTunnel(&_StaderChildPool.TransactOpts, _address)
}

// SetInstantPoolOwner is a paid mutator transaction binding the contract method 0x701845b8.
//
// Solidity: function setInstantPoolOwner(address _address) returns()
func (_StaderChildPool *StaderChildPoolTransactor) SetInstantPoolOwner(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "setInstantPoolOwner", _address)
}

// SetInstantPoolOwner is a paid mutator transaction binding the contract method 0x701845b8.
//
// Solidity: function setInstantPoolOwner(address _address) returns()
func (_StaderChildPool *StaderChildPoolSession) SetInstantPoolOwner(_address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetInstantPoolOwner(&_StaderChildPool.TransactOpts, _address)
}

// SetInstantPoolOwner is a paid mutator transaction binding the contract method 0x701845b8.
//
// Solidity: function setInstantPoolOwner(address _address) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) SetInstantPoolOwner(_address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetInstantPoolOwner(&_StaderChildPool.TransactOpts, _address)
}

// SetInstantWithdrawalFeeBps is a paid mutator transaction binding the contract method 0x13acce6a.
//
// Solidity: function setInstantWithdrawalFeeBps(uint256 _feeBps) returns()
func (_StaderChildPool *StaderChildPoolTransactor) SetInstantWithdrawalFeeBps(opts *bind.TransactOpts, _feeBps *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "setInstantWithdrawalFeeBps", _feeBps)
}

// SetInstantWithdrawalFeeBps is a paid mutator transaction binding the contract method 0x13acce6a.
//
// Solidity: function setInstantWithdrawalFeeBps(uint256 _feeBps) returns()
func (_StaderChildPool *StaderChildPoolSession) SetInstantWithdrawalFeeBps(_feeBps *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetInstantWithdrawalFeeBps(&_StaderChildPool.TransactOpts, _feeBps)
}

// SetInstantWithdrawalFeeBps is a paid mutator transaction binding the contract method 0x13acce6a.
//
// Solidity: function setInstantWithdrawalFeeBps(uint256 _feeBps) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) SetInstantWithdrawalFeeBps(_feeBps *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetInstantWithdrawalFeeBps(&_StaderChildPool.TransactOpts, _feeBps)
}

// SetMaticXSwapLockPeriod is a paid mutator transaction binding the contract method 0x72be8891.
//
// Solidity: function setMaticXSwapLockPeriod(uint256 _hours) returns()
func (_StaderChildPool *StaderChildPoolTransactor) SetMaticXSwapLockPeriod(opts *bind.TransactOpts, _hours *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "setMaticXSwapLockPeriod", _hours)
}

// SetMaticXSwapLockPeriod is a paid mutator transaction binding the contract method 0x72be8891.
//
// Solidity: function setMaticXSwapLockPeriod(uint256 _hours) returns()
func (_StaderChildPool *StaderChildPoolSession) SetMaticXSwapLockPeriod(_hours *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetMaticXSwapLockPeriod(&_StaderChildPool.TransactOpts, _hours)
}

// SetMaticXSwapLockPeriod is a paid mutator transaction binding the contract method 0x72be8891.
//
// Solidity: function setMaticXSwapLockPeriod(uint256 _hours) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) SetMaticXSwapLockPeriod(_hours *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetMaticXSwapLockPeriod(&_StaderChildPool.TransactOpts, _hours)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _address) returns()
func (_StaderChildPool *StaderChildPoolTransactor) SetTreasury(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "setTreasury", _address)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _address) returns()
func (_StaderChildPool *StaderChildPoolSession) SetTreasury(_address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetTreasury(&_StaderChildPool.TransactOpts, _address)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _address) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) SetTreasury(_address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetTreasury(&_StaderChildPool.TransactOpts, _address)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _address) returns()
func (_StaderChildPool *StaderChildPoolTransactor) SetTrustedForwarder(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "setTrustedForwarder", _address)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _address) returns()
func (_StaderChildPool *StaderChildPoolSession) SetTrustedForwarder(_address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetTrustedForwarder(&_StaderChildPool.TransactOpts, _address)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _address) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) SetTrustedForwarder(_address common.Address) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetTrustedForwarder(&_StaderChildPool.TransactOpts, _address)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_StaderChildPool *StaderChildPoolTransactor) SetVersion(opts *bind.TransactOpts, _version string) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "setVersion", _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_StaderChildPool *StaderChildPoolSession) SetVersion(_version string) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetVersion(&_StaderChildPool.TransactOpts, _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) SetVersion(_version string) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SetVersion(&_StaderChildPool.TransactOpts, _version)
}

// SwapMaticForMaticXViaInstantPool is a paid mutator transaction binding the contract method 0xc78cf1a0.
//
// Solidity: function swapMaticForMaticXViaInstantPool() payable returns()
func (_StaderChildPool *StaderChildPoolTransactor) SwapMaticForMaticXViaInstantPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "swapMaticForMaticXViaInstantPool")
}

// SwapMaticForMaticXViaInstantPool is a paid mutator transaction binding the contract method 0xc78cf1a0.
//
// Solidity: function swapMaticForMaticXViaInstantPool() payable returns()
func (_StaderChildPool *StaderChildPoolSession) SwapMaticForMaticXViaInstantPool() (*types.Transaction, error) {
	return _StaderChildPool.Contract.SwapMaticForMaticXViaInstantPool(&_StaderChildPool.TransactOpts)
}

// SwapMaticForMaticXViaInstantPool is a paid mutator transaction binding the contract method 0xc78cf1a0.
//
// Solidity: function swapMaticForMaticXViaInstantPool() payable returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) SwapMaticForMaticXViaInstantPool() (*types.Transaction, error) {
	return _StaderChildPool.Contract.SwapMaticForMaticXViaInstantPool(&_StaderChildPool.TransactOpts)
}

// SwapMaticXForMaticViaInstantPool is a paid mutator transaction binding the contract method 0xd3bf9d59.
//
// Solidity: function swapMaticXForMaticViaInstantPool(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolTransactor) SwapMaticXForMaticViaInstantPool(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "swapMaticXForMaticViaInstantPool", _amount)
}

// SwapMaticXForMaticViaInstantPool is a paid mutator transaction binding the contract method 0xd3bf9d59.
//
// Solidity: function swapMaticXForMaticViaInstantPool(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolSession) SwapMaticXForMaticViaInstantPool(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SwapMaticXForMaticViaInstantPool(&_StaderChildPool.TransactOpts, _amount)
}

// SwapMaticXForMaticViaInstantPool is a paid mutator transaction binding the contract method 0xd3bf9d59.
//
// Solidity: function swapMaticXForMaticViaInstantPool(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) SwapMaticXForMaticViaInstantPool(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.SwapMaticXForMaticViaInstantPool(&_StaderChildPool.TransactOpts, _amount)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_StaderChildPool *StaderChildPoolTransactor) TogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "togglePause")
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_StaderChildPool *StaderChildPoolSession) TogglePause() (*types.Transaction, error) {
	return _StaderChildPool.Contract.TogglePause(&_StaderChildPool.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) TogglePause() (*types.Transaction, error) {
	return _StaderChildPool.Contract.TogglePause(&_StaderChildPool.TransactOpts)
}

// WithdrawInstantPoolMatic is a paid mutator transaction binding the contract method 0x00fd822c.
//
// Solidity: function withdrawInstantPoolMatic(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolTransactor) WithdrawInstantPoolMatic(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "withdrawInstantPoolMatic", _amount)
}

// WithdrawInstantPoolMatic is a paid mutator transaction binding the contract method 0x00fd822c.
//
// Solidity: function withdrawInstantPoolMatic(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolSession) WithdrawInstantPoolMatic(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.WithdrawInstantPoolMatic(&_StaderChildPool.TransactOpts, _amount)
}

// WithdrawInstantPoolMatic is a paid mutator transaction binding the contract method 0x00fd822c.
//
// Solidity: function withdrawInstantPoolMatic(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) WithdrawInstantPoolMatic(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.WithdrawInstantPoolMatic(&_StaderChildPool.TransactOpts, _amount)
}

// WithdrawInstantPoolMaticX is a paid mutator transaction binding the contract method 0xc1e324a5.
//
// Solidity: function withdrawInstantPoolMaticX(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolTransactor) WithdrawInstantPoolMaticX(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "withdrawInstantPoolMaticX", _amount)
}

// WithdrawInstantPoolMaticX is a paid mutator transaction binding the contract method 0xc1e324a5.
//
// Solidity: function withdrawInstantPoolMaticX(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolSession) WithdrawInstantPoolMaticX(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.WithdrawInstantPoolMaticX(&_StaderChildPool.TransactOpts, _amount)
}

// WithdrawInstantPoolMaticX is a paid mutator transaction binding the contract method 0xc1e324a5.
//
// Solidity: function withdrawInstantPoolMaticX(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) WithdrawInstantPoolMaticX(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.WithdrawInstantPoolMaticX(&_StaderChildPool.TransactOpts, _amount)
}

// WithdrawInstantWithdrawalFees is a paid mutator transaction binding the contract method 0x13d0255e.
//
// Solidity: function withdrawInstantWithdrawalFees(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolTransactor) WithdrawInstantWithdrawalFees(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.contract.Transact(opts, "withdrawInstantWithdrawalFees", _amount)
}

// WithdrawInstantWithdrawalFees is a paid mutator transaction binding the contract method 0x13d0255e.
//
// Solidity: function withdrawInstantWithdrawalFees(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolSession) WithdrawInstantWithdrawalFees(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.WithdrawInstantWithdrawalFees(&_StaderChildPool.TransactOpts, _amount)
}

// WithdrawInstantWithdrawalFees is a paid mutator transaction binding the contract method 0x13d0255e.
//
// Solidity: function withdrawInstantWithdrawalFees(uint256 _amount) returns()
func (_StaderChildPool *StaderChildPoolTransactorSession) WithdrawInstantWithdrawalFees(_amount *big.Int) (*types.Transaction, error) {
	return _StaderChildPool.Contract.WithdrawInstantWithdrawalFees(&_StaderChildPool.TransactOpts, _amount)
}
