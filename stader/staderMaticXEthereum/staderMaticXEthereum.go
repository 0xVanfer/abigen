// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staderMaticXEthereum

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

// IStaderMaticXEthereumWithdrawalRequest is an auto generated low-level Go binding around an user-defined struct.
type IStaderMaticXEthereumWithdrawalRequest struct {
	ValidatorNonce   *big.Int
	RequestEpoch     *big.Int
	ValidatorAddress common.Address
}

// StaderMaticXEthereumMetaData contains all meta data concerning the StaderMaticXEthereum contract.
var StaderMaticXEthereumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INSTANT_POOL_OWNER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PREDICATE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"claimWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"convertMaticToMaticX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"convertMaticXToMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercent\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fxStateRootTunnel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_stakeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonERC20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorRegistry\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"getSharesAmountOfUserWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPooledMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStakeAcrossAllValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserWithdrawalRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"internalType\":\"structIStaderMaticXEthereum.WithdrawalRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonERC20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_instantPoolOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"instantPoolMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"instantPoolMaticX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"instantPoolOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fromValidatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toValidatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"migrateDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintMaticXToInstantPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"provideInstantPoolMatic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"provideInstantPoolMaticX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_feePercent\",\"type\":\"uint8\"}],\"name\":\"setFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setFxStateRootTunnel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setInstantPoolOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setValidatorRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setupBotAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"stakeRewardsAndDistributeFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"swapMaticForMaticXViaInstantPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawInstantPoolMatic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawInstantPoolMaticX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"withdrawRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"}],\"name\":\"withdrawValidatorsReward\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StaderMaticXEthereumABI is the input ABI used to generate the binding from.
// Deprecated: Use StaderMaticXEthereumMetaData.ABI instead.
var StaderMaticXEthereumABI = StaderMaticXEthereumMetaData.ABI

// StaderMaticXEthereum is an auto generated Go binding around an Ethereum contract.
type StaderMaticXEthereum struct {
	StaderMaticXEthereumCaller     // Read-only binding to the contract
	StaderMaticXEthereumTransactor // Write-only binding to the contract
	StaderMaticXEthereumFilterer   // Log filterer for contract events
}

// StaderMaticXEthereumCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaderMaticXEthereumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderMaticXEthereumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaderMaticXEthereumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderMaticXEthereumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaderMaticXEthereumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderMaticXEthereumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaderMaticXEthereumSession struct {
	Contract     *StaderMaticXEthereum // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StaderMaticXEthereumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaderMaticXEthereumCallerSession struct {
	Contract *StaderMaticXEthereumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// StaderMaticXEthereumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaderMaticXEthereumTransactorSession struct {
	Contract     *StaderMaticXEthereumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// StaderMaticXEthereumRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaderMaticXEthereumRaw struct {
	Contract *StaderMaticXEthereum // Generic contract binding to access the raw methods on
}

// StaderMaticXEthereumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaderMaticXEthereumCallerRaw struct {
	Contract *StaderMaticXEthereumCaller // Generic read-only contract binding to access the raw methods on
}

// StaderMaticXEthereumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaderMaticXEthereumTransactorRaw struct {
	Contract *StaderMaticXEthereumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaderMaticXEthereum creates a new instance of StaderMaticXEthereum, bound to a specific deployed contract.
func NewStaderMaticXEthereum(address common.Address, backend bind.ContractBackend) (*StaderMaticXEthereum, error) {
	contract, err := bindStaderMaticXEthereum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaderMaticXEthereum{StaderMaticXEthereumCaller: StaderMaticXEthereumCaller{contract: contract}, StaderMaticXEthereumTransactor: StaderMaticXEthereumTransactor{contract: contract}, StaderMaticXEthereumFilterer: StaderMaticXEthereumFilterer{contract: contract}}, nil
}

// NewStaderMaticXEthereumCaller creates a new read-only instance of StaderMaticXEthereum, bound to a specific deployed contract.
func NewStaderMaticXEthereumCaller(address common.Address, caller bind.ContractCaller) (*StaderMaticXEthereumCaller, error) {
	contract, err := bindStaderMaticXEthereum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaderMaticXEthereumCaller{contract: contract}, nil
}

// NewStaderMaticXEthereumTransactor creates a new write-only instance of StaderMaticXEthereum, bound to a specific deployed contract.
func NewStaderMaticXEthereumTransactor(address common.Address, transactor bind.ContractTransactor) (*StaderMaticXEthereumTransactor, error) {
	contract, err := bindStaderMaticXEthereum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaderMaticXEthereumTransactor{contract: contract}, nil
}

// NewStaderMaticXEthereumFilterer creates a new log filterer instance of StaderMaticXEthereum, bound to a specific deployed contract.
func NewStaderMaticXEthereumFilterer(address common.Address, filterer bind.ContractFilterer) (*StaderMaticXEthereumFilterer, error) {
	contract, err := bindStaderMaticXEthereum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaderMaticXEthereumFilterer{contract: contract}, nil
}

// bindStaderMaticXEthereum binds a generic wrapper to an already deployed contract.
func bindStaderMaticXEthereum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StaderMaticXEthereumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderMaticXEthereum *StaderMaticXEthereumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderMaticXEthereum.Contract.StaderMaticXEthereumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderMaticXEthereum *StaderMaticXEthereumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.StaderMaticXEthereumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderMaticXEthereum *StaderMaticXEthereumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.StaderMaticXEthereumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderMaticXEthereum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.contract.Transact(opts, method, params...)
}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) BOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "BOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) BOT() ([32]byte, error) {
	return _StaderMaticXEthereum.Contract.BOT(&_StaderMaticXEthereum.CallOpts)
}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) BOT() ([32]byte, error) {
	return _StaderMaticXEthereum.Contract.BOT(&_StaderMaticXEthereum.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StaderMaticXEthereum.Contract.DEFAULTADMINROLE(&_StaderMaticXEthereum.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StaderMaticXEthereum.Contract.DEFAULTADMINROLE(&_StaderMaticXEthereum.CallOpts)
}

// INSTANTPOOLOWNER is a free data retrieval call binding the contract method 0xcba45a7c.
//
// Solidity: function INSTANT_POOL_OWNER() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) INSTANTPOOLOWNER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "INSTANT_POOL_OWNER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INSTANTPOOLOWNER is a free data retrieval call binding the contract method 0xcba45a7c.
//
// Solidity: function INSTANT_POOL_OWNER() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) INSTANTPOOLOWNER() ([32]byte, error) {
	return _StaderMaticXEthereum.Contract.INSTANTPOOLOWNER(&_StaderMaticXEthereum.CallOpts)
}

// INSTANTPOOLOWNER is a free data retrieval call binding the contract method 0xcba45a7c.
//
// Solidity: function INSTANT_POOL_OWNER() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) INSTANTPOOLOWNER() ([32]byte, error) {
	return _StaderMaticXEthereum.Contract.INSTANTPOOLOWNER(&_StaderMaticXEthereum.CallOpts)
}

// PREDICATEROLE is a free data retrieval call binding the contract method 0xe72db5fd.
//
// Solidity: function PREDICATE_ROLE() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) PREDICATEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "PREDICATE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PREDICATEROLE is a free data retrieval call binding the contract method 0xe72db5fd.
//
// Solidity: function PREDICATE_ROLE() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) PREDICATEROLE() ([32]byte, error) {
	return _StaderMaticXEthereum.Contract.PREDICATEROLE(&_StaderMaticXEthereum.CallOpts)
}

// PREDICATEROLE is a free data retrieval call binding the contract method 0xe72db5fd.
//
// Solidity: function PREDICATE_ROLE() view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) PREDICATEROLE() ([32]byte, error) {
	return _StaderMaticXEthereum.Contract.PREDICATEROLE(&_StaderMaticXEthereum.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.Allowance(&_StaderMaticXEthereum.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.Allowance(&_StaderMaticXEthereum.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.BalanceOf(&_StaderMaticXEthereum.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.BalanceOf(&_StaderMaticXEthereum.CallOpts, account)
}

// ConvertMaticToMaticX is a free data retrieval call binding the contract method 0x9683e28e.
//
// Solidity: function convertMaticToMaticX(uint256 _balance) view returns(uint256, uint256, uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) ConvertMaticToMaticX(opts *bind.CallOpts, _balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "convertMaticToMaticX", _balance)

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
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) ConvertMaticToMaticX(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _StaderMaticXEthereum.Contract.ConvertMaticToMaticX(&_StaderMaticXEthereum.CallOpts, _balance)
}

// ConvertMaticToMaticX is a free data retrieval call binding the contract method 0x9683e28e.
//
// Solidity: function convertMaticToMaticX(uint256 _balance) view returns(uint256, uint256, uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) ConvertMaticToMaticX(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _StaderMaticXEthereum.Contract.ConvertMaticToMaticX(&_StaderMaticXEthereum.CallOpts, _balance)
}

// ConvertMaticXToMatic is a free data retrieval call binding the contract method 0x75a85ef5.
//
// Solidity: function convertMaticXToMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) ConvertMaticXToMatic(opts *bind.CallOpts, _balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "convertMaticXToMatic", _balance)

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
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) ConvertMaticXToMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _StaderMaticXEthereum.Contract.ConvertMaticXToMatic(&_StaderMaticXEthereum.CallOpts, _balance)
}

// ConvertMaticXToMatic is a free data retrieval call binding the contract method 0x75a85ef5.
//
// Solidity: function convertMaticXToMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) ConvertMaticXToMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _StaderMaticXEthereum.Contract.ConvertMaticXToMatic(&_StaderMaticXEthereum.CallOpts, _balance)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Decimals() (uint8, error) {
	return _StaderMaticXEthereum.Contract.Decimals(&_StaderMaticXEthereum.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) Decimals() (uint8, error) {
	return _StaderMaticXEthereum.Contract.Decimals(&_StaderMaticXEthereum.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint8)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) FeePercent(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "feePercent")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint8)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) FeePercent() (uint8, error) {
	return _StaderMaticXEthereum.Contract.FeePercent(&_StaderMaticXEthereum.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint8)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) FeePercent() (uint8, error) {
	return _StaderMaticXEthereum.Contract.FeePercent(&_StaderMaticXEthereum.CallOpts)
}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) FxStateRootTunnel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "fxStateRootTunnel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) FxStateRootTunnel() (common.Address, error) {
	return _StaderMaticXEthereum.Contract.FxStateRootTunnel(&_StaderMaticXEthereum.CallOpts)
}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) FxStateRootTunnel() (common.Address, error) {
	return _StaderMaticXEthereum.Contract.FxStateRootTunnel(&_StaderMaticXEthereum.CallOpts)
}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns(address _stakeManager, address _polygonERC20, address _validatorRegistry)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) GetContracts(opts *bind.CallOpts) (struct {
	StakeManager      common.Address
	PolygonERC20      common.Address
	ValidatorRegistry common.Address
}, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "getContracts")

	outstruct := new(struct {
		StakeManager      common.Address
		PolygonERC20      common.Address
		ValidatorRegistry common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeManager = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PolygonERC20 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ValidatorRegistry = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns(address _stakeManager, address _polygonERC20, address _validatorRegistry)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) GetContracts() (struct {
	StakeManager      common.Address
	PolygonERC20      common.Address
	ValidatorRegistry common.Address
}, error) {
	return _StaderMaticXEthereum.Contract.GetContracts(&_StaderMaticXEthereum.CallOpts)
}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns(address _stakeManager, address _polygonERC20, address _validatorRegistry)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) GetContracts() (struct {
	StakeManager      common.Address
	PolygonERC20      common.Address
	ValidatorRegistry common.Address
}, error) {
	return _StaderMaticXEthereum.Contract.GetContracts(&_StaderMaticXEthereum.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StaderMaticXEthereum.Contract.GetRoleAdmin(&_StaderMaticXEthereum.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StaderMaticXEthereum.Contract.GetRoleAdmin(&_StaderMaticXEthereum.CallOpts, role)
}

// GetSharesAmountOfUserWithdrawalRequest is a free data retrieval call binding the contract method 0x73f0ecdf.
//
// Solidity: function getSharesAmountOfUserWithdrawalRequest(address _address, uint256 _idx) view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) GetSharesAmountOfUserWithdrawalRequest(opts *bind.CallOpts, _address common.Address, _idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "getSharesAmountOfUserWithdrawalRequest", _address, _idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesAmountOfUserWithdrawalRequest is a free data retrieval call binding the contract method 0x73f0ecdf.
//
// Solidity: function getSharesAmountOfUserWithdrawalRequest(address _address, uint256 _idx) view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) GetSharesAmountOfUserWithdrawalRequest(_address common.Address, _idx *big.Int) (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.GetSharesAmountOfUserWithdrawalRequest(&_StaderMaticXEthereum.CallOpts, _address, _idx)
}

// GetSharesAmountOfUserWithdrawalRequest is a free data retrieval call binding the contract method 0x73f0ecdf.
//
// Solidity: function getSharesAmountOfUserWithdrawalRequest(address _address, uint256 _idx) view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) GetSharesAmountOfUserWithdrawalRequest(_address common.Address, _idx *big.Int) (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.GetSharesAmountOfUserWithdrawalRequest(&_StaderMaticXEthereum.CallOpts, _address, _idx)
}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) GetTotalPooledMatic(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "getTotalPooledMatic")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) GetTotalPooledMatic() (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.GetTotalPooledMatic(&_StaderMaticXEthereum.CallOpts)
}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) GetTotalPooledMatic() (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.GetTotalPooledMatic(&_StaderMaticXEthereum.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) GetTotalStake(opts *bind.CallOpts, _validatorShare common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "getTotalStake", _validatorShare)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) GetTotalStake(_validatorShare common.Address) (*big.Int, *big.Int, error) {
	return _StaderMaticXEthereum.Contract.GetTotalStake(&_StaderMaticXEthereum.CallOpts, _validatorShare)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) GetTotalStake(_validatorShare common.Address) (*big.Int, *big.Int, error) {
	return _StaderMaticXEthereum.Contract.GetTotalStake(&_StaderMaticXEthereum.CallOpts, _validatorShare)
}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) GetTotalStakeAcrossAllValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "getTotalStakeAcrossAllValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) GetTotalStakeAcrossAllValidators() (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.GetTotalStakeAcrossAllValidators(&_StaderMaticXEthereum.CallOpts)
}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) GetTotalStakeAcrossAllValidators() (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.GetTotalStakeAcrossAllValidators(&_StaderMaticXEthereum.CallOpts)
}

// GetUserWithdrawalRequests is a free data retrieval call binding the contract method 0x6c930228.
//
// Solidity: function getUserWithdrawalRequests(address _address) view returns((uint256,uint256,address)[])
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) GetUserWithdrawalRequests(opts *bind.CallOpts, _address common.Address) ([]IStaderMaticXEthereumWithdrawalRequest, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "getUserWithdrawalRequests", _address)

	if err != nil {
		return *new([]IStaderMaticXEthereumWithdrawalRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStaderMaticXEthereumWithdrawalRequest)).(*[]IStaderMaticXEthereumWithdrawalRequest)

	return out0, err

}

// GetUserWithdrawalRequests is a free data retrieval call binding the contract method 0x6c930228.
//
// Solidity: function getUserWithdrawalRequests(address _address) view returns((uint256,uint256,address)[])
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) GetUserWithdrawalRequests(_address common.Address) ([]IStaderMaticXEthereumWithdrawalRequest, error) {
	return _StaderMaticXEthereum.Contract.GetUserWithdrawalRequests(&_StaderMaticXEthereum.CallOpts, _address)
}

// GetUserWithdrawalRequests is a free data retrieval call binding the contract method 0x6c930228.
//
// Solidity: function getUserWithdrawalRequests(address _address) view returns((uint256,uint256,address)[])
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) GetUserWithdrawalRequests(_address common.Address) ([]IStaderMaticXEthereumWithdrawalRequest, error) {
	return _StaderMaticXEthereum.Contract.GetUserWithdrawalRequests(&_StaderMaticXEthereum.CallOpts, _address)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StaderMaticXEthereum.Contract.HasRole(&_StaderMaticXEthereum.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StaderMaticXEthereum.Contract.HasRole(&_StaderMaticXEthereum.CallOpts, role, account)
}

// InstantPoolMatic is a free data retrieval call binding the contract method 0x89dfa025.
//
// Solidity: function instantPoolMatic() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) InstantPoolMatic(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "instantPoolMatic")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InstantPoolMatic is a free data retrieval call binding the contract method 0x89dfa025.
//
// Solidity: function instantPoolMatic() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) InstantPoolMatic() (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.InstantPoolMatic(&_StaderMaticXEthereum.CallOpts)
}

// InstantPoolMatic is a free data retrieval call binding the contract method 0x89dfa025.
//
// Solidity: function instantPoolMatic() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) InstantPoolMatic() (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.InstantPoolMatic(&_StaderMaticXEthereum.CallOpts)
}

// InstantPoolMaticX is a free data retrieval call binding the contract method 0xc759352d.
//
// Solidity: function instantPoolMaticX() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) InstantPoolMaticX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "instantPoolMaticX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InstantPoolMaticX is a free data retrieval call binding the contract method 0xc759352d.
//
// Solidity: function instantPoolMaticX() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) InstantPoolMaticX() (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.InstantPoolMaticX(&_StaderMaticXEthereum.CallOpts)
}

// InstantPoolMaticX is a free data retrieval call binding the contract method 0xc759352d.
//
// Solidity: function instantPoolMaticX() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) InstantPoolMaticX() (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.InstantPoolMaticX(&_StaderMaticXEthereum.CallOpts)
}

// InstantPoolOwner is a free data retrieval call binding the contract method 0x1c083124.
//
// Solidity: function instantPoolOwner() view returns(address)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) InstantPoolOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "instantPoolOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InstantPoolOwner is a free data retrieval call binding the contract method 0x1c083124.
//
// Solidity: function instantPoolOwner() view returns(address)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) InstantPoolOwner() (common.Address, error) {
	return _StaderMaticXEthereum.Contract.InstantPoolOwner(&_StaderMaticXEthereum.CallOpts)
}

// InstantPoolOwner is a free data retrieval call binding the contract method 0x1c083124.
//
// Solidity: function instantPoolOwner() view returns(address)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) InstantPoolOwner() (common.Address, error) {
	return _StaderMaticXEthereum.Contract.InstantPoolOwner(&_StaderMaticXEthereum.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Name() (string, error) {
	return _StaderMaticXEthereum.Contract.Name(&_StaderMaticXEthereum.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) Name() (string, error) {
	return _StaderMaticXEthereum.Contract.Name(&_StaderMaticXEthereum.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Paused() (bool, error) {
	return _StaderMaticXEthereum.Contract.Paused(&_StaderMaticXEthereum.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) Paused() (bool, error) {
	return _StaderMaticXEthereum.Contract.Paused(&_StaderMaticXEthereum.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StaderMaticXEthereum.Contract.SupportsInterface(&_StaderMaticXEthereum.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StaderMaticXEthereum.Contract.SupportsInterface(&_StaderMaticXEthereum.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Symbol() (string, error) {
	return _StaderMaticXEthereum.Contract.Symbol(&_StaderMaticXEthereum.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) Symbol() (string, error) {
	return _StaderMaticXEthereum.Contract.Symbol(&_StaderMaticXEthereum.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) TotalSupply() (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.TotalSupply(&_StaderMaticXEthereum.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) TotalSupply() (*big.Int, error) {
	return _StaderMaticXEthereum.Contract.TotalSupply(&_StaderMaticXEthereum.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Treasury() (common.Address, error) {
	return _StaderMaticXEthereum.Contract.Treasury(&_StaderMaticXEthereum.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) Treasury() (common.Address, error) {
	return _StaderMaticXEthereum.Contract.Treasury(&_StaderMaticXEthereum.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StaderMaticXEthereum *StaderMaticXEthereumCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StaderMaticXEthereum.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Version() (string, error) {
	return _StaderMaticXEthereum.Contract.Version(&_StaderMaticXEthereum.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StaderMaticXEthereum *StaderMaticXEthereumCallerSession) Version() (string, error) {
	return _StaderMaticXEthereum.Contract.Version(&_StaderMaticXEthereum.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.Approve(&_StaderMaticXEthereum.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.Approve(&_StaderMaticXEthereum.TransactOpts, spender, amount)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xf8444436.
//
// Solidity: function claimWithdrawal(uint256 _idx) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) ClaimWithdrawal(opts *bind.TransactOpts, _idx *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "claimWithdrawal", _idx)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xf8444436.
//
// Solidity: function claimWithdrawal(uint256 _idx) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) ClaimWithdrawal(_idx *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.ClaimWithdrawal(&_StaderMaticXEthereum.TransactOpts, _idx)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xf8444436.
//
// Solidity: function claimWithdrawal(uint256 _idx) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) ClaimWithdrawal(_idx *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.ClaimWithdrawal(&_StaderMaticXEthereum.TransactOpts, _idx)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.DecreaseAllowance(&_StaderMaticXEthereum.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.DecreaseAllowance(&_StaderMaticXEthereum.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.GrantRole(&_StaderMaticXEthereum.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.GrantRole(&_StaderMaticXEthereum.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.IncreaseAllowance(&_StaderMaticXEthereum.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.IncreaseAllowance(&_StaderMaticXEthereum.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _validatorRegistry, address _stakeManager, address _polygonERC20, address _manager, address _instantPoolOwner, address _treasury) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) Initialize(opts *bind.TransactOpts, _validatorRegistry common.Address, _stakeManager common.Address, _polygonERC20 common.Address, _manager common.Address, _instantPoolOwner common.Address, _treasury common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "initialize", _validatorRegistry, _stakeManager, _polygonERC20, _manager, _instantPoolOwner, _treasury)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _validatorRegistry, address _stakeManager, address _polygonERC20, address _manager, address _instantPoolOwner, address _treasury) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Initialize(_validatorRegistry common.Address, _stakeManager common.Address, _polygonERC20 common.Address, _manager common.Address, _instantPoolOwner common.Address, _treasury common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.Initialize(&_StaderMaticXEthereum.TransactOpts, _validatorRegistry, _stakeManager, _polygonERC20, _manager, _instantPoolOwner, _treasury)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _validatorRegistry, address _stakeManager, address _polygonERC20, address _manager, address _instantPoolOwner, address _treasury) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) Initialize(_validatorRegistry common.Address, _stakeManager common.Address, _polygonERC20 common.Address, _manager common.Address, _instantPoolOwner common.Address, _treasury common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.Initialize(&_StaderMaticXEthereum.TransactOpts, _validatorRegistry, _stakeManager, _polygonERC20, _manager, _instantPoolOwner, _treasury)
}

// MigrateDelegation is a paid mutator transaction binding the contract method 0xfb1ef52c.
//
// Solidity: function migrateDelegation(uint256 _fromValidatorId, uint256 _toValidatorId, uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) MigrateDelegation(opts *bind.TransactOpts, _fromValidatorId *big.Int, _toValidatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "migrateDelegation", _fromValidatorId, _toValidatorId, _amount)
}

// MigrateDelegation is a paid mutator transaction binding the contract method 0xfb1ef52c.
//
// Solidity: function migrateDelegation(uint256 _fromValidatorId, uint256 _toValidatorId, uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) MigrateDelegation(_fromValidatorId *big.Int, _toValidatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.MigrateDelegation(&_StaderMaticXEthereum.TransactOpts, _fromValidatorId, _toValidatorId, _amount)
}

// MigrateDelegation is a paid mutator transaction binding the contract method 0xfb1ef52c.
//
// Solidity: function migrateDelegation(uint256 _fromValidatorId, uint256 _toValidatorId, uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) MigrateDelegation(_fromValidatorId *big.Int, _toValidatorId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.MigrateDelegation(&_StaderMaticXEthereum.TransactOpts, _fromValidatorId, _toValidatorId, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _user, uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) Mint(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "mint", _user, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _user, uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Mint(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.Mint(&_StaderMaticXEthereum.TransactOpts, _user, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _user, uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) Mint(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.Mint(&_StaderMaticXEthereum.TransactOpts, _user, _amount)
}

// MintMaticXToInstantPool is a paid mutator transaction binding the contract method 0x02b09f2e.
//
// Solidity: function mintMaticXToInstantPool() returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) MintMaticXToInstantPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "mintMaticXToInstantPool")
}

// MintMaticXToInstantPool is a paid mutator transaction binding the contract method 0x02b09f2e.
//
// Solidity: function mintMaticXToInstantPool() returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) MintMaticXToInstantPool() (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.MintMaticXToInstantPool(&_StaderMaticXEthereum.TransactOpts)
}

// MintMaticXToInstantPool is a paid mutator transaction binding the contract method 0x02b09f2e.
//
// Solidity: function mintMaticXToInstantPool() returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) MintMaticXToInstantPool() (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.MintMaticXToInstantPool(&_StaderMaticXEthereum.TransactOpts)
}

// ProvideInstantPoolMatic is a paid mutator transaction binding the contract method 0x74b7b2d2.
//
// Solidity: function provideInstantPoolMatic(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) ProvideInstantPoolMatic(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "provideInstantPoolMatic", _amount)
}

// ProvideInstantPoolMatic is a paid mutator transaction binding the contract method 0x74b7b2d2.
//
// Solidity: function provideInstantPoolMatic(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) ProvideInstantPoolMatic(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.ProvideInstantPoolMatic(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// ProvideInstantPoolMatic is a paid mutator transaction binding the contract method 0x74b7b2d2.
//
// Solidity: function provideInstantPoolMatic(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) ProvideInstantPoolMatic(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.ProvideInstantPoolMatic(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// ProvideInstantPoolMaticX is a paid mutator transaction binding the contract method 0x68c05c97.
//
// Solidity: function provideInstantPoolMaticX(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) ProvideInstantPoolMaticX(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "provideInstantPoolMaticX", _amount)
}

// ProvideInstantPoolMaticX is a paid mutator transaction binding the contract method 0x68c05c97.
//
// Solidity: function provideInstantPoolMaticX(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) ProvideInstantPoolMaticX(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.ProvideInstantPoolMaticX(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// ProvideInstantPoolMaticX is a paid mutator transaction binding the contract method 0x68c05c97.
//
// Solidity: function provideInstantPoolMaticX(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) ProvideInstantPoolMaticX(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.ProvideInstantPoolMaticX(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.RenounceRole(&_StaderMaticXEthereum.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.RenounceRole(&_StaderMaticXEthereum.TransactOpts, role, account)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) RequestWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "requestWithdraw", _amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) RequestWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.RequestWithdraw(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) RequestWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.RequestWithdraw(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.RevokeRole(&_StaderMaticXEthereum.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.RevokeRole(&_StaderMaticXEthereum.TransactOpts, role, account)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0xf4838176.
//
// Solidity: function setFeePercent(uint8 _feePercent) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) SetFeePercent(opts *bind.TransactOpts, _feePercent uint8) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "setFeePercent", _feePercent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0xf4838176.
//
// Solidity: function setFeePercent(uint8 _feePercent) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) SetFeePercent(_feePercent uint8) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetFeePercent(&_StaderMaticXEthereum.TransactOpts, _feePercent)
}

// SetFeePercent is a paid mutator transaction binding the contract method 0xf4838176.
//
// Solidity: function setFeePercent(uint8 _feePercent) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) SetFeePercent(_feePercent uint8) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetFeePercent(&_StaderMaticXEthereum.TransactOpts, _feePercent)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) SetFxStateRootTunnel(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "setFxStateRootTunnel", _address)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) SetFxStateRootTunnel(_address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetFxStateRootTunnel(&_StaderMaticXEthereum.TransactOpts, _address)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) SetFxStateRootTunnel(_address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetFxStateRootTunnel(&_StaderMaticXEthereum.TransactOpts, _address)
}

// SetInstantPoolOwner is a paid mutator transaction binding the contract method 0x701845b8.
//
// Solidity: function setInstantPoolOwner(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) SetInstantPoolOwner(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "setInstantPoolOwner", _address)
}

// SetInstantPoolOwner is a paid mutator transaction binding the contract method 0x701845b8.
//
// Solidity: function setInstantPoolOwner(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) SetInstantPoolOwner(_address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetInstantPoolOwner(&_StaderMaticXEthereum.TransactOpts, _address)
}

// SetInstantPoolOwner is a paid mutator transaction binding the contract method 0x701845b8.
//
// Solidity: function setInstantPoolOwner(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) SetInstantPoolOwner(_address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetInstantPoolOwner(&_StaderMaticXEthereum.TransactOpts, _address)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) SetTreasury(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "setTreasury", _address)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) SetTreasury(_address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetTreasury(&_StaderMaticXEthereum.TransactOpts, _address)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) SetTreasury(_address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetTreasury(&_StaderMaticXEthereum.TransactOpts, _address)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) SetValidatorRegistry(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "setValidatorRegistry", _address)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) SetValidatorRegistry(_address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetValidatorRegistry(&_StaderMaticXEthereum.TransactOpts, _address)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _address) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) SetValidatorRegistry(_address common.Address) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetValidatorRegistry(&_StaderMaticXEthereum.TransactOpts, _address)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) SetVersion(opts *bind.TransactOpts, _version string) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "setVersion", _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) SetVersion(_version string) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetVersion(&_StaderMaticXEthereum.TransactOpts, _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) SetVersion(_version string) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetVersion(&_StaderMaticXEthereum.TransactOpts, _version)
}

// SetupBotAdmin is a paid mutator transaction binding the contract method 0xc21e4634.
//
// Solidity: function setupBotAdmin() returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) SetupBotAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "setupBotAdmin")
}

// SetupBotAdmin is a paid mutator transaction binding the contract method 0xc21e4634.
//
// Solidity: function setupBotAdmin() returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) SetupBotAdmin() (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetupBotAdmin(&_StaderMaticXEthereum.TransactOpts)
}

// SetupBotAdmin is a paid mutator transaction binding the contract method 0xc21e4634.
//
// Solidity: function setupBotAdmin() returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) SetupBotAdmin() (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SetupBotAdmin(&_StaderMaticXEthereum.TransactOpts)
}

// StakeRewardsAndDistributeFees is a paid mutator transaction binding the contract method 0xdcea4be9.
//
// Solidity: function stakeRewardsAndDistributeFees(uint256 _validatorId) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) StakeRewardsAndDistributeFees(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "stakeRewardsAndDistributeFees", _validatorId)
}

// StakeRewardsAndDistributeFees is a paid mutator transaction binding the contract method 0xdcea4be9.
//
// Solidity: function stakeRewardsAndDistributeFees(uint256 _validatorId) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) StakeRewardsAndDistributeFees(_validatorId *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.StakeRewardsAndDistributeFees(&_StaderMaticXEthereum.TransactOpts, _validatorId)
}

// StakeRewardsAndDistributeFees is a paid mutator transaction binding the contract method 0xdcea4be9.
//
// Solidity: function stakeRewardsAndDistributeFees(uint256 _validatorId) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) StakeRewardsAndDistributeFees(_validatorId *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.StakeRewardsAndDistributeFees(&_StaderMaticXEthereum.TransactOpts, _validatorId)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 _amount) returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) Submit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "submit", _amount)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 _amount) returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Submit(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.Submit(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 _amount) returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) Submit(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.Submit(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// SwapMaticForMaticXViaInstantPool is a paid mutator transaction binding the contract method 0xbaec30ca.
//
// Solidity: function swapMaticForMaticXViaInstantPool(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) SwapMaticForMaticXViaInstantPool(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "swapMaticForMaticXViaInstantPool", _amount)
}

// SwapMaticForMaticXViaInstantPool is a paid mutator transaction binding the contract method 0xbaec30ca.
//
// Solidity: function swapMaticForMaticXViaInstantPool(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) SwapMaticForMaticXViaInstantPool(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SwapMaticForMaticXViaInstantPool(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// SwapMaticForMaticXViaInstantPool is a paid mutator transaction binding the contract method 0xbaec30ca.
//
// Solidity: function swapMaticForMaticXViaInstantPool(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) SwapMaticForMaticXViaInstantPool(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.SwapMaticForMaticXViaInstantPool(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) TogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "togglePause")
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) TogglePause() (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.TogglePause(&_StaderMaticXEthereum.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) TogglePause() (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.TogglePause(&_StaderMaticXEthereum.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.Transfer(&_StaderMaticXEthereum.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.Transfer(&_StaderMaticXEthereum.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.TransferFrom(&_StaderMaticXEthereum.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.TransferFrom(&_StaderMaticXEthereum.TransactOpts, from, to, amount)
}

// WithdrawInstantPoolMatic is a paid mutator transaction binding the contract method 0x00fd822c.
//
// Solidity: function withdrawInstantPoolMatic(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) WithdrawInstantPoolMatic(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "withdrawInstantPoolMatic", _amount)
}

// WithdrawInstantPoolMatic is a paid mutator transaction binding the contract method 0x00fd822c.
//
// Solidity: function withdrawInstantPoolMatic(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) WithdrawInstantPoolMatic(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.WithdrawInstantPoolMatic(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// WithdrawInstantPoolMatic is a paid mutator transaction binding the contract method 0x00fd822c.
//
// Solidity: function withdrawInstantPoolMatic(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) WithdrawInstantPoolMatic(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.WithdrawInstantPoolMatic(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// WithdrawInstantPoolMaticX is a paid mutator transaction binding the contract method 0xc1e324a5.
//
// Solidity: function withdrawInstantPoolMaticX(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) WithdrawInstantPoolMaticX(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "withdrawInstantPoolMaticX", _amount)
}

// WithdrawInstantPoolMaticX is a paid mutator transaction binding the contract method 0xc1e324a5.
//
// Solidity: function withdrawInstantPoolMaticX(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) WithdrawInstantPoolMaticX(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.WithdrawInstantPoolMaticX(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// WithdrawInstantPoolMaticX is a paid mutator transaction binding the contract method 0xc1e324a5.
//
// Solidity: function withdrawInstantPoolMaticX(uint256 _amount) returns()
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) WithdrawInstantPoolMaticX(_amount *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.WithdrawInstantPoolMaticX(&_StaderMaticXEthereum.TransactOpts, _amount)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 _validatorId) returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) WithdrawRewards(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "withdrawRewards", _validatorId)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 _validatorId) returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) WithdrawRewards(_validatorId *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.WithdrawRewards(&_StaderMaticXEthereum.TransactOpts, _validatorId)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 _validatorId) returns(uint256)
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) WithdrawRewards(_validatorId *big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.WithdrawRewards(&_StaderMaticXEthereum.TransactOpts, _validatorId)
}

// WithdrawValidatorsReward is a paid mutator transaction binding the contract method 0x74d72acc.
//
// Solidity: function withdrawValidatorsReward(uint256[] _validatorIds) returns(uint256[])
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactor) WithdrawValidatorsReward(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.contract.Transact(opts, "withdrawValidatorsReward", _validatorIds)
}

// WithdrawValidatorsReward is a paid mutator transaction binding the contract method 0x74d72acc.
//
// Solidity: function withdrawValidatorsReward(uint256[] _validatorIds) returns(uint256[])
func (_StaderMaticXEthereum *StaderMaticXEthereumSession) WithdrawValidatorsReward(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.WithdrawValidatorsReward(&_StaderMaticXEthereum.TransactOpts, _validatorIds)
}

// WithdrawValidatorsReward is a paid mutator transaction binding the contract method 0x74d72acc.
//
// Solidity: function withdrawValidatorsReward(uint256[] _validatorIds) returns(uint256[])
func (_StaderMaticXEthereum *StaderMaticXEthereumTransactorSession) WithdrawValidatorsReward(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _StaderMaticXEthereum.Contract.WithdrawValidatorsReward(&_StaderMaticXEthereum.TransactOpts, _validatorIds)
}
