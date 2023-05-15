// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package benqiJumpRateModel

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

// BenqiJumpRateModelMetaData contains all meta data concerning the BenqiJumpRateModel contract.
var BenqiJumpRateModelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"baseRatePerTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"name\":\"getBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"getSupplyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInterestRateModel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jumpMultiplierPerTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kink\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierPerTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestampsPerYear\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"name\":\"utilizationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// BenqiJumpRateModelABI is the input ABI used to generate the binding from.
// Deprecated: Use BenqiJumpRateModelMetaData.ABI instead.
var BenqiJumpRateModelABI = BenqiJumpRateModelMetaData.ABI

// BenqiJumpRateModel is an auto generated Go binding around an Ethereum contract.
type BenqiJumpRateModel struct {
	BenqiJumpRateModelCaller     // Read-only binding to the contract
	BenqiJumpRateModelTransactor // Write-only binding to the contract
	BenqiJumpRateModelFilterer   // Log filterer for contract events
}

// BenqiJumpRateModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type BenqiJumpRateModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqiJumpRateModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BenqiJumpRateModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqiJumpRateModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BenqiJumpRateModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqiJumpRateModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BenqiJumpRateModelSession struct {
	Contract     *BenqiJumpRateModel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BenqiJumpRateModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BenqiJumpRateModelCallerSession struct {
	Contract *BenqiJumpRateModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BenqiJumpRateModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BenqiJumpRateModelTransactorSession struct {
	Contract     *BenqiJumpRateModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BenqiJumpRateModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type BenqiJumpRateModelRaw struct {
	Contract *BenqiJumpRateModel // Generic contract binding to access the raw methods on
}

// BenqiJumpRateModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BenqiJumpRateModelCallerRaw struct {
	Contract *BenqiJumpRateModelCaller // Generic read-only contract binding to access the raw methods on
}

// BenqiJumpRateModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BenqiJumpRateModelTransactorRaw struct {
	Contract *BenqiJumpRateModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBenqiJumpRateModel creates a new instance of BenqiJumpRateModel, bound to a specific deployed contract.
func NewBenqiJumpRateModel(address common.Address, backend bind.ContractBackend) (*BenqiJumpRateModel, error) {
	contract, err := bindBenqiJumpRateModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BenqiJumpRateModel{BenqiJumpRateModelCaller: BenqiJumpRateModelCaller{contract: contract}, BenqiJumpRateModelTransactor: BenqiJumpRateModelTransactor{contract: contract}, BenqiJumpRateModelFilterer: BenqiJumpRateModelFilterer{contract: contract}}, nil
}

// NewBenqiJumpRateModelCaller creates a new read-only instance of BenqiJumpRateModel, bound to a specific deployed contract.
func NewBenqiJumpRateModelCaller(address common.Address, caller bind.ContractCaller) (*BenqiJumpRateModelCaller, error) {
	contract, err := bindBenqiJumpRateModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BenqiJumpRateModelCaller{contract: contract}, nil
}

// NewBenqiJumpRateModelTransactor creates a new write-only instance of BenqiJumpRateModel, bound to a specific deployed contract.
func NewBenqiJumpRateModelTransactor(address common.Address, transactor bind.ContractTransactor) (*BenqiJumpRateModelTransactor, error) {
	contract, err := bindBenqiJumpRateModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BenqiJumpRateModelTransactor{contract: contract}, nil
}

// NewBenqiJumpRateModelFilterer creates a new log filterer instance of BenqiJumpRateModel, bound to a specific deployed contract.
func NewBenqiJumpRateModelFilterer(address common.Address, filterer bind.ContractFilterer) (*BenqiJumpRateModelFilterer, error) {
	contract, err := bindBenqiJumpRateModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BenqiJumpRateModelFilterer{contract: contract}, nil
}

// bindBenqiJumpRateModel binds a generic wrapper to an already deployed contract.
func bindBenqiJumpRateModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BenqiJumpRateModelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BenqiJumpRateModel *BenqiJumpRateModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BenqiJumpRateModel.Contract.BenqiJumpRateModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BenqiJumpRateModel *BenqiJumpRateModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenqiJumpRateModel.Contract.BenqiJumpRateModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BenqiJumpRateModel *BenqiJumpRateModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BenqiJumpRateModel.Contract.BenqiJumpRateModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BenqiJumpRateModel *BenqiJumpRateModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BenqiJumpRateModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BenqiJumpRateModel *BenqiJumpRateModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenqiJumpRateModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BenqiJumpRateModel *BenqiJumpRateModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BenqiJumpRateModel.Contract.contract.Transact(opts, method, params...)
}

// BaseRatePerTimestamp is a free data retrieval call binding the contract method 0x40bc0af4.
//
// Solidity: function baseRatePerTimestamp() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCaller) BaseRatePerTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiJumpRateModel.contract.Call(opts, &out, "baseRatePerTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRatePerTimestamp is a free data retrieval call binding the contract method 0x40bc0af4.
//
// Solidity: function baseRatePerTimestamp() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelSession) BaseRatePerTimestamp() (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.BaseRatePerTimestamp(&_BenqiJumpRateModel.CallOpts)
}

// BaseRatePerTimestamp is a free data retrieval call binding the contract method 0x40bc0af4.
//
// Solidity: function baseRatePerTimestamp() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCallerSession) BaseRatePerTimestamp() (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.BaseRatePerTimestamp(&_BenqiJumpRateModel.CallOpts)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCaller) GetBorrowRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BenqiJumpRateModel.contract.Call(opts, &out, "getBorrowRate", cash, borrows, reserves)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelSession) GetBorrowRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.GetBorrowRate(&_BenqiJumpRateModel.CallOpts, cash, borrows, reserves)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCallerSession) GetBorrowRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.GetBorrowRate(&_BenqiJumpRateModel.CallOpts, cash, borrows, reserves)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCaller) GetSupplyRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BenqiJumpRateModel.contract.Call(opts, &out, "getSupplyRate", cash, borrows, reserves, reserveFactorMantissa)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelSession) GetSupplyRate(cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.GetSupplyRate(&_BenqiJumpRateModel.CallOpts, cash, borrows, reserves, reserveFactorMantissa)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCallerSession) GetSupplyRate(cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.GetSupplyRate(&_BenqiJumpRateModel.CallOpts, cash, borrows, reserves, reserveFactorMantissa)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_BenqiJumpRateModel *BenqiJumpRateModelCaller) IsInterestRateModel(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BenqiJumpRateModel.contract.Call(opts, &out, "isInterestRateModel")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_BenqiJumpRateModel *BenqiJumpRateModelSession) IsInterestRateModel() (bool, error) {
	return _BenqiJumpRateModel.Contract.IsInterestRateModel(&_BenqiJumpRateModel.CallOpts)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_BenqiJumpRateModel *BenqiJumpRateModelCallerSession) IsInterestRateModel() (bool, error) {
	return _BenqiJumpRateModel.Contract.IsInterestRateModel(&_BenqiJumpRateModel.CallOpts)
}

// JumpMultiplierPerTimestamp is a free data retrieval call binding the contract method 0x26c394f7.
//
// Solidity: function jumpMultiplierPerTimestamp() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCaller) JumpMultiplierPerTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiJumpRateModel.contract.Call(opts, &out, "jumpMultiplierPerTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JumpMultiplierPerTimestamp is a free data retrieval call binding the contract method 0x26c394f7.
//
// Solidity: function jumpMultiplierPerTimestamp() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelSession) JumpMultiplierPerTimestamp() (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.JumpMultiplierPerTimestamp(&_BenqiJumpRateModel.CallOpts)
}

// JumpMultiplierPerTimestamp is a free data retrieval call binding the contract method 0x26c394f7.
//
// Solidity: function jumpMultiplierPerTimestamp() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCallerSession) JumpMultiplierPerTimestamp() (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.JumpMultiplierPerTimestamp(&_BenqiJumpRateModel.CallOpts)
}

// Kink is a free data retrieval call binding the contract method 0xfd2da339.
//
// Solidity: function kink() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCaller) Kink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiJumpRateModel.contract.Call(opts, &out, "kink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Kink is a free data retrieval call binding the contract method 0xfd2da339.
//
// Solidity: function kink() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelSession) Kink() (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.Kink(&_BenqiJumpRateModel.CallOpts)
}

// Kink is a free data retrieval call binding the contract method 0xfd2da339.
//
// Solidity: function kink() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCallerSession) Kink() (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.Kink(&_BenqiJumpRateModel.CallOpts)
}

// MultiplierPerTimestamp is a free data retrieval call binding the contract method 0x6c2df6a7.
//
// Solidity: function multiplierPerTimestamp() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCaller) MultiplierPerTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiJumpRateModel.contract.Call(opts, &out, "multiplierPerTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MultiplierPerTimestamp is a free data retrieval call binding the contract method 0x6c2df6a7.
//
// Solidity: function multiplierPerTimestamp() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelSession) MultiplierPerTimestamp() (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.MultiplierPerTimestamp(&_BenqiJumpRateModel.CallOpts)
}

// MultiplierPerTimestamp is a free data retrieval call binding the contract method 0x6c2df6a7.
//
// Solidity: function multiplierPerTimestamp() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCallerSession) MultiplierPerTimestamp() (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.MultiplierPerTimestamp(&_BenqiJumpRateModel.CallOpts)
}

// TimestampsPerYear is a free data retrieval call binding the contract method 0x0c574861.
//
// Solidity: function timestampsPerYear() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCaller) TimestampsPerYear(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiJumpRateModel.contract.Call(opts, &out, "timestampsPerYear")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimestampsPerYear is a free data retrieval call binding the contract method 0x0c574861.
//
// Solidity: function timestampsPerYear() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelSession) TimestampsPerYear() (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.TimestampsPerYear(&_BenqiJumpRateModel.CallOpts)
}

// TimestampsPerYear is a free data retrieval call binding the contract method 0x0c574861.
//
// Solidity: function timestampsPerYear() view returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCallerSession) TimestampsPerYear() (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.TimestampsPerYear(&_BenqiJumpRateModel.CallOpts)
}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCaller) UtilizationRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BenqiJumpRateModel.contract.Call(opts, &out, "utilizationRate", cash, borrows, reserves)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelSession) UtilizationRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.UtilizationRate(&_BenqiJumpRateModel.CallOpts, cash, borrows, reserves)
}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_BenqiJumpRateModel *BenqiJumpRateModelCallerSession) UtilizationRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _BenqiJumpRateModel.Contract.UtilizationRate(&_BenqiJumpRateModel.CallOpts, cash, borrows, reserves)
}
