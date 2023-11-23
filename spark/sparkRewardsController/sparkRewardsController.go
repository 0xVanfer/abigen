// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sparkRewardsController

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

// RewardsDataTypesRewardsConfigInput is an auto generated low-level Go binding around an user-defined struct.
type RewardsDataTypesRewardsConfigInput struct {
	EmissionPerSecond *big.Int
	TotalSupply       *big.Int
	DistributionEnd   uint32
	Asset             common.Address
	Reward            common.Address
	TransferStrategy  common.Address
	RewardOracle      common.Address
}

// SparkRewardsControllerMetaData contains all meta data concerning the SparkRewardsController contract.
var SparkRewardsControllerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsAccrued\",\"type\":\"uint256\"}],\"name\":\"Accrued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldEmission\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEmission\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDistributionEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDistributionEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"}],\"name\":\"AssetConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"ClaimerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardOracle\",\"type\":\"address\"}],\"name\":\"RewardOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferStrategy\",\"type\":\"address\"}],\"name\":\"TransferStrategyInstalled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EMISSION_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claimAllRewards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"rewardsList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claimedAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claimAllRewardsOnBehalf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"rewardsList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claimedAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"claimAllRewardsToSelf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"rewardsList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claimedAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"claimRewardsOnBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"claimRewardsToSelf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint88\",\"name\":\"emissionPerSecond\",\"type\":\"uint88\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"distributionEnd\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"contractITransferStrategyBase\",\"name\":\"transferStrategy\",\"type\":\"address\"},{\"internalType\":\"contractIEACAggregatorProxy\",\"name\":\"rewardOracle\",\"type\":\"address\"}],\"internalType\":\"structRewardsDataTypes.RewardsConfigInput[]\",\"name\":\"config\",\"type\":\"tuple[]\"}],\"name\":\"configureAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAllUserRewards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getAssetIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getClaimer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getDistributionEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmissionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getRewardOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getRewardsByAsset\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getRewardsData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardsList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getTransferStrategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getUserAccruedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getUserAssetIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getUserRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBalance\",\"type\":\"uint256\"}],\"name\":\"handleAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"setClaimer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newDistributionEnd\",\"type\":\"uint32\"}],\"name\":\"setDistributionEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"rewards\",\"type\":\"address[]\"},{\"internalType\":\"uint88[]\",\"name\":\"newEmissionsPerSecond\",\"type\":\"uint88[]\"}],\"name\":\"setEmissionPerSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"contractIEACAggregatorProxy\",\"name\":\"rewardOracle\",\"type\":\"address\"}],\"name\":\"setRewardOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"contractITransferStrategyBase\",\"name\":\"transferStrategy\",\"type\":\"address\"}],\"name\":\"setTransferStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SparkRewardsControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use SparkRewardsControllerMetaData.ABI instead.
var SparkRewardsControllerABI = SparkRewardsControllerMetaData.ABI

// SparkRewardsController is an auto generated Go binding around an Ethereum contract.
type SparkRewardsController struct {
	SparkRewardsControllerCaller     // Read-only binding to the contract
	SparkRewardsControllerTransactor // Write-only binding to the contract
	SparkRewardsControllerFilterer   // Log filterer for contract events
}

// SparkRewardsControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SparkRewardsControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SparkRewardsControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SparkRewardsControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SparkRewardsControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SparkRewardsControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SparkRewardsControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SparkRewardsControllerSession struct {
	Contract     *SparkRewardsController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SparkRewardsControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SparkRewardsControllerCallerSession struct {
	Contract *SparkRewardsControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// SparkRewardsControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SparkRewardsControllerTransactorSession struct {
	Contract     *SparkRewardsControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// SparkRewardsControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SparkRewardsControllerRaw struct {
	Contract *SparkRewardsController // Generic contract binding to access the raw methods on
}

// SparkRewardsControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SparkRewardsControllerCallerRaw struct {
	Contract *SparkRewardsControllerCaller // Generic read-only contract binding to access the raw methods on
}

// SparkRewardsControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SparkRewardsControllerTransactorRaw struct {
	Contract *SparkRewardsControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSparkRewardsController creates a new instance of SparkRewardsController, bound to a specific deployed contract.
func NewSparkRewardsController(address common.Address, backend bind.ContractBackend) (*SparkRewardsController, error) {
	contract, err := bindSparkRewardsController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SparkRewardsController{SparkRewardsControllerCaller: SparkRewardsControllerCaller{contract: contract}, SparkRewardsControllerTransactor: SparkRewardsControllerTransactor{contract: contract}, SparkRewardsControllerFilterer: SparkRewardsControllerFilterer{contract: contract}}, nil
}

// NewSparkRewardsControllerCaller creates a new read-only instance of SparkRewardsController, bound to a specific deployed contract.
func NewSparkRewardsControllerCaller(address common.Address, caller bind.ContractCaller) (*SparkRewardsControllerCaller, error) {
	contract, err := bindSparkRewardsController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SparkRewardsControllerCaller{contract: contract}, nil
}

// NewSparkRewardsControllerTransactor creates a new write-only instance of SparkRewardsController, bound to a specific deployed contract.
func NewSparkRewardsControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*SparkRewardsControllerTransactor, error) {
	contract, err := bindSparkRewardsController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SparkRewardsControllerTransactor{contract: contract}, nil
}

// NewSparkRewardsControllerFilterer creates a new log filterer instance of SparkRewardsController, bound to a specific deployed contract.
func NewSparkRewardsControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*SparkRewardsControllerFilterer, error) {
	contract, err := bindSparkRewardsController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SparkRewardsControllerFilterer{contract: contract}, nil
}

// bindSparkRewardsController binds a generic wrapper to an already deployed contract.
func bindSparkRewardsController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SparkRewardsControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SparkRewardsController *SparkRewardsControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SparkRewardsController.Contract.SparkRewardsControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SparkRewardsController *SparkRewardsControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SparkRewardsControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SparkRewardsController *SparkRewardsControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SparkRewardsControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SparkRewardsController *SparkRewardsControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SparkRewardsController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SparkRewardsController *SparkRewardsControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SparkRewardsController *SparkRewardsControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.contract.Transact(opts, method, params...)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_SparkRewardsController *SparkRewardsControllerCaller) EMISSIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "EMISSION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_SparkRewardsController *SparkRewardsControllerSession) EMISSIONMANAGER() (common.Address, error) {
	return _SparkRewardsController.Contract.EMISSIONMANAGER(&_SparkRewardsController.CallOpts)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) EMISSIONMANAGER() (common.Address, error) {
	return _SparkRewardsController.Contract.EMISSIONMANAGER(&_SparkRewardsController.CallOpts)
}

// GetAllUserRewards is a free data retrieval call binding the contract method 0x4c0369c3.
//
// Solidity: function getAllUserRewards(address[] assets, address user) view returns(address[], uint256[])
func (_SparkRewardsController *SparkRewardsControllerCaller) GetAllUserRewards(opts *bind.CallOpts, assets []common.Address, user common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getAllUserRewards", assets, user)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetAllUserRewards is a free data retrieval call binding the contract method 0x4c0369c3.
//
// Solidity: function getAllUserRewards(address[] assets, address user) view returns(address[], uint256[])
func (_SparkRewardsController *SparkRewardsControllerSession) GetAllUserRewards(assets []common.Address, user common.Address) ([]common.Address, []*big.Int, error) {
	return _SparkRewardsController.Contract.GetAllUserRewards(&_SparkRewardsController.CallOpts, assets, user)
}

// GetAllUserRewards is a free data retrieval call binding the contract method 0x4c0369c3.
//
// Solidity: function getAllUserRewards(address[] assets, address user) view returns(address[], uint256[])
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetAllUserRewards(assets []common.Address, user common.Address) ([]common.Address, []*big.Int, error) {
	return _SparkRewardsController.Contract.GetAllUserRewards(&_SparkRewardsController.CallOpts, assets, user)
}

// GetAssetDecimals is a free data retrieval call binding the contract method 0x9efd6f72.
//
// Solidity: function getAssetDecimals(address asset) view returns(uint8)
func (_SparkRewardsController *SparkRewardsControllerCaller) GetAssetDecimals(opts *bind.CallOpts, asset common.Address) (uint8, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getAssetDecimals", asset)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAssetDecimals is a free data retrieval call binding the contract method 0x9efd6f72.
//
// Solidity: function getAssetDecimals(address asset) view returns(uint8)
func (_SparkRewardsController *SparkRewardsControllerSession) GetAssetDecimals(asset common.Address) (uint8, error) {
	return _SparkRewardsController.Contract.GetAssetDecimals(&_SparkRewardsController.CallOpts, asset)
}

// GetAssetDecimals is a free data retrieval call binding the contract method 0x9efd6f72.
//
// Solidity: function getAssetDecimals(address asset) view returns(uint8)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetAssetDecimals(asset common.Address) (uint8, error) {
	return _SparkRewardsController.Contract.GetAssetDecimals(&_SparkRewardsController.CallOpts, asset)
}

// GetAssetIndex is a free data retrieval call binding the contract method 0x886fe70b.
//
// Solidity: function getAssetIndex(address asset, address reward) view returns(uint256, uint256)
func (_SparkRewardsController *SparkRewardsControllerCaller) GetAssetIndex(opts *bind.CallOpts, asset common.Address, reward common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getAssetIndex", asset, reward)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAssetIndex is a free data retrieval call binding the contract method 0x886fe70b.
//
// Solidity: function getAssetIndex(address asset, address reward) view returns(uint256, uint256)
func (_SparkRewardsController *SparkRewardsControllerSession) GetAssetIndex(asset common.Address, reward common.Address) (*big.Int, *big.Int, error) {
	return _SparkRewardsController.Contract.GetAssetIndex(&_SparkRewardsController.CallOpts, asset, reward)
}

// GetAssetIndex is a free data retrieval call binding the contract method 0x886fe70b.
//
// Solidity: function getAssetIndex(address asset, address reward) view returns(uint256, uint256)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetAssetIndex(asset common.Address, reward common.Address) (*big.Int, *big.Int, error) {
	return _SparkRewardsController.Contract.GetAssetIndex(&_SparkRewardsController.CallOpts, asset, reward)
}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_SparkRewardsController *SparkRewardsControllerCaller) GetClaimer(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getClaimer", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_SparkRewardsController *SparkRewardsControllerSession) GetClaimer(user common.Address) (common.Address, error) {
	return _SparkRewardsController.Contract.GetClaimer(&_SparkRewardsController.CallOpts, user)
}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetClaimer(user common.Address) (common.Address, error) {
	return _SparkRewardsController.Contract.GetClaimer(&_SparkRewardsController.CallOpts, user)
}

// GetDistributionEnd is a free data retrieval call binding the contract method 0x1b839c77.
//
// Solidity: function getDistributionEnd(address asset, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerCaller) GetDistributionEnd(opts *bind.CallOpts, asset common.Address, reward common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getDistributionEnd", asset, reward)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionEnd is a free data retrieval call binding the contract method 0x1b839c77.
//
// Solidity: function getDistributionEnd(address asset, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerSession) GetDistributionEnd(asset common.Address, reward common.Address) (*big.Int, error) {
	return _SparkRewardsController.Contract.GetDistributionEnd(&_SparkRewardsController.CallOpts, asset, reward)
}

// GetDistributionEnd is a free data retrieval call binding the contract method 0x1b839c77.
//
// Solidity: function getDistributionEnd(address asset, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetDistributionEnd(asset common.Address, reward common.Address) (*big.Int, error) {
	return _SparkRewardsController.Contract.GetDistributionEnd(&_SparkRewardsController.CallOpts, asset, reward)
}

// GetEmissionManager is a free data retrieval call binding the contract method 0x92074b08.
//
// Solidity: function getEmissionManager() view returns(address)
func (_SparkRewardsController *SparkRewardsControllerCaller) GetEmissionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getEmissionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEmissionManager is a free data retrieval call binding the contract method 0x92074b08.
//
// Solidity: function getEmissionManager() view returns(address)
func (_SparkRewardsController *SparkRewardsControllerSession) GetEmissionManager() (common.Address, error) {
	return _SparkRewardsController.Contract.GetEmissionManager(&_SparkRewardsController.CallOpts)
}

// GetEmissionManager is a free data retrieval call binding the contract method 0x92074b08.
//
// Solidity: function getEmissionManager() view returns(address)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetEmissionManager() (common.Address, error) {
	return _SparkRewardsController.Contract.GetEmissionManager(&_SparkRewardsController.CallOpts)
}

// GetRewardOracle is a free data retrieval call binding the contract method 0x2a17bf60.
//
// Solidity: function getRewardOracle(address reward) view returns(address)
func (_SparkRewardsController *SparkRewardsControllerCaller) GetRewardOracle(opts *bind.CallOpts, reward common.Address) (common.Address, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getRewardOracle", reward)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardOracle is a free data retrieval call binding the contract method 0x2a17bf60.
//
// Solidity: function getRewardOracle(address reward) view returns(address)
func (_SparkRewardsController *SparkRewardsControllerSession) GetRewardOracle(reward common.Address) (common.Address, error) {
	return _SparkRewardsController.Contract.GetRewardOracle(&_SparkRewardsController.CallOpts, reward)
}

// GetRewardOracle is a free data retrieval call binding the contract method 0x2a17bf60.
//
// Solidity: function getRewardOracle(address reward) view returns(address)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetRewardOracle(reward common.Address) (common.Address, error) {
	return _SparkRewardsController.Contract.GetRewardOracle(&_SparkRewardsController.CallOpts, reward)
}

// GetRewardsByAsset is a free data retrieval call binding the contract method 0x6657732f.
//
// Solidity: function getRewardsByAsset(address asset) view returns(address[])
func (_SparkRewardsController *SparkRewardsControllerCaller) GetRewardsByAsset(opts *bind.CallOpts, asset common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getRewardsByAsset", asset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardsByAsset is a free data retrieval call binding the contract method 0x6657732f.
//
// Solidity: function getRewardsByAsset(address asset) view returns(address[])
func (_SparkRewardsController *SparkRewardsControllerSession) GetRewardsByAsset(asset common.Address) ([]common.Address, error) {
	return _SparkRewardsController.Contract.GetRewardsByAsset(&_SparkRewardsController.CallOpts, asset)
}

// GetRewardsByAsset is a free data retrieval call binding the contract method 0x6657732f.
//
// Solidity: function getRewardsByAsset(address asset) view returns(address[])
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetRewardsByAsset(asset common.Address) ([]common.Address, error) {
	return _SparkRewardsController.Contract.GetRewardsByAsset(&_SparkRewardsController.CallOpts, asset)
}

// GetRewardsData is a free data retrieval call binding the contract method 0x7eff4ba8.
//
// Solidity: function getRewardsData(address asset, address reward) view returns(uint256, uint256, uint256, uint256)
func (_SparkRewardsController *SparkRewardsControllerCaller) GetRewardsData(opts *bind.CallOpts, asset common.Address, reward common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getRewardsData", asset, reward)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetRewardsData is a free data retrieval call binding the contract method 0x7eff4ba8.
//
// Solidity: function getRewardsData(address asset, address reward) view returns(uint256, uint256, uint256, uint256)
func (_SparkRewardsController *SparkRewardsControllerSession) GetRewardsData(asset common.Address, reward common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SparkRewardsController.Contract.GetRewardsData(&_SparkRewardsController.CallOpts, asset, reward)
}

// GetRewardsData is a free data retrieval call binding the contract method 0x7eff4ba8.
//
// Solidity: function getRewardsData(address asset, address reward) view returns(uint256, uint256, uint256, uint256)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetRewardsData(asset common.Address, reward common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SparkRewardsController.Contract.GetRewardsData(&_SparkRewardsController.CallOpts, asset, reward)
}

// GetRewardsList is a free data retrieval call binding the contract method 0xb45ac1a9.
//
// Solidity: function getRewardsList() view returns(address[])
func (_SparkRewardsController *SparkRewardsControllerCaller) GetRewardsList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getRewardsList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardsList is a free data retrieval call binding the contract method 0xb45ac1a9.
//
// Solidity: function getRewardsList() view returns(address[])
func (_SparkRewardsController *SparkRewardsControllerSession) GetRewardsList() ([]common.Address, error) {
	return _SparkRewardsController.Contract.GetRewardsList(&_SparkRewardsController.CallOpts)
}

// GetRewardsList is a free data retrieval call binding the contract method 0xb45ac1a9.
//
// Solidity: function getRewardsList() view returns(address[])
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetRewardsList() ([]common.Address, error) {
	return _SparkRewardsController.Contract.GetRewardsList(&_SparkRewardsController.CallOpts)
}

// GetTransferStrategy is a free data retrieval call binding the contract method 0x5f130b24.
//
// Solidity: function getTransferStrategy(address reward) view returns(address)
func (_SparkRewardsController *SparkRewardsControllerCaller) GetTransferStrategy(opts *bind.CallOpts, reward common.Address) (common.Address, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getTransferStrategy", reward)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTransferStrategy is a free data retrieval call binding the contract method 0x5f130b24.
//
// Solidity: function getTransferStrategy(address reward) view returns(address)
func (_SparkRewardsController *SparkRewardsControllerSession) GetTransferStrategy(reward common.Address) (common.Address, error) {
	return _SparkRewardsController.Contract.GetTransferStrategy(&_SparkRewardsController.CallOpts, reward)
}

// GetTransferStrategy is a free data retrieval call binding the contract method 0x5f130b24.
//
// Solidity: function getTransferStrategy(address reward) view returns(address)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetTransferStrategy(reward common.Address) (common.Address, error) {
	return _SparkRewardsController.Contract.GetTransferStrategy(&_SparkRewardsController.CallOpts, reward)
}

// GetUserAccruedRewards is a free data retrieval call binding the contract method 0xb022418c.
//
// Solidity: function getUserAccruedRewards(address user, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerCaller) GetUserAccruedRewards(opts *bind.CallOpts, user common.Address, reward common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getUserAccruedRewards", user, reward)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAccruedRewards is a free data retrieval call binding the contract method 0xb022418c.
//
// Solidity: function getUserAccruedRewards(address user, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerSession) GetUserAccruedRewards(user common.Address, reward common.Address) (*big.Int, error) {
	return _SparkRewardsController.Contract.GetUserAccruedRewards(&_SparkRewardsController.CallOpts, user, reward)
}

// GetUserAccruedRewards is a free data retrieval call binding the contract method 0xb022418c.
//
// Solidity: function getUserAccruedRewards(address user, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetUserAccruedRewards(user common.Address, reward common.Address) (*big.Int, error) {
	return _SparkRewardsController.Contract.GetUserAccruedRewards(&_SparkRewardsController.CallOpts, user, reward)
}

// GetUserAssetIndex is a free data retrieval call binding the contract method 0x533f542a.
//
// Solidity: function getUserAssetIndex(address user, address asset, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerCaller) GetUserAssetIndex(opts *bind.CallOpts, user common.Address, asset common.Address, reward common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getUserAssetIndex", user, asset, reward)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAssetIndex is a free data retrieval call binding the contract method 0x533f542a.
//
// Solidity: function getUserAssetIndex(address user, address asset, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerSession) GetUserAssetIndex(user common.Address, asset common.Address, reward common.Address) (*big.Int, error) {
	return _SparkRewardsController.Contract.GetUserAssetIndex(&_SparkRewardsController.CallOpts, user, asset, reward)
}

// GetUserAssetIndex is a free data retrieval call binding the contract method 0x533f542a.
//
// Solidity: function getUserAssetIndex(address user, address asset, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetUserAssetIndex(user common.Address, asset common.Address, reward common.Address) (*big.Int, error) {
	return _SparkRewardsController.Contract.GetUserAssetIndex(&_SparkRewardsController.CallOpts, user, asset, reward)
}

// GetUserRewards is a free data retrieval call binding the contract method 0x70674ab9.
//
// Solidity: function getUserRewards(address[] assets, address user, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerCaller) GetUserRewards(opts *bind.CallOpts, assets []common.Address, user common.Address, reward common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SparkRewardsController.contract.Call(opts, &out, "getUserRewards", assets, user, reward)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserRewards is a free data retrieval call binding the contract method 0x70674ab9.
//
// Solidity: function getUserRewards(address[] assets, address user, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerSession) GetUserRewards(assets []common.Address, user common.Address, reward common.Address) (*big.Int, error) {
	return _SparkRewardsController.Contract.GetUserRewards(&_SparkRewardsController.CallOpts, assets, user, reward)
}

// GetUserRewards is a free data retrieval call binding the contract method 0x70674ab9.
//
// Solidity: function getUserRewards(address[] assets, address user, address reward) view returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerCallerSession) GetUserRewards(assets []common.Address, user common.Address, reward common.Address) (*big.Int, error) {
	return _SparkRewardsController.Contract.GetUserRewards(&_SparkRewardsController.CallOpts, assets, user, reward)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0xbb492bf5.
//
// Solidity: function claimAllRewards(address[] assets, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_SparkRewardsController *SparkRewardsControllerTransactor) ClaimAllRewards(opts *bind.TransactOpts, assets []common.Address, to common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "claimAllRewards", assets, to)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0xbb492bf5.
//
// Solidity: function claimAllRewards(address[] assets, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_SparkRewardsController *SparkRewardsControllerSession) ClaimAllRewards(assets []common.Address, to common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimAllRewards(&_SparkRewardsController.TransactOpts, assets, to)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0xbb492bf5.
//
// Solidity: function claimAllRewards(address[] assets, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) ClaimAllRewards(assets []common.Address, to common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimAllRewards(&_SparkRewardsController.TransactOpts, assets, to)
}

// ClaimAllRewardsOnBehalf is a paid mutator transaction binding the contract method 0x9ff55db9.
//
// Solidity: function claimAllRewardsOnBehalf(address[] assets, address user, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_SparkRewardsController *SparkRewardsControllerTransactor) ClaimAllRewardsOnBehalf(opts *bind.TransactOpts, assets []common.Address, user common.Address, to common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "claimAllRewardsOnBehalf", assets, user, to)
}

// ClaimAllRewardsOnBehalf is a paid mutator transaction binding the contract method 0x9ff55db9.
//
// Solidity: function claimAllRewardsOnBehalf(address[] assets, address user, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_SparkRewardsController *SparkRewardsControllerSession) ClaimAllRewardsOnBehalf(assets []common.Address, user common.Address, to common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimAllRewardsOnBehalf(&_SparkRewardsController.TransactOpts, assets, user, to)
}

// ClaimAllRewardsOnBehalf is a paid mutator transaction binding the contract method 0x9ff55db9.
//
// Solidity: function claimAllRewardsOnBehalf(address[] assets, address user, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) ClaimAllRewardsOnBehalf(assets []common.Address, user common.Address, to common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimAllRewardsOnBehalf(&_SparkRewardsController.TransactOpts, assets, user, to)
}

// ClaimAllRewardsToSelf is a paid mutator transaction binding the contract method 0xbf90f63a.
//
// Solidity: function claimAllRewardsToSelf(address[] assets) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_SparkRewardsController *SparkRewardsControllerTransactor) ClaimAllRewardsToSelf(opts *bind.TransactOpts, assets []common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "claimAllRewardsToSelf", assets)
}

// ClaimAllRewardsToSelf is a paid mutator transaction binding the contract method 0xbf90f63a.
//
// Solidity: function claimAllRewardsToSelf(address[] assets) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_SparkRewardsController *SparkRewardsControllerSession) ClaimAllRewardsToSelf(assets []common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimAllRewardsToSelf(&_SparkRewardsController.TransactOpts, assets)
}

// ClaimAllRewardsToSelf is a paid mutator transaction binding the contract method 0xbf90f63a.
//
// Solidity: function claimAllRewardsToSelf(address[] assets) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) ClaimAllRewardsToSelf(assets []common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimAllRewardsToSelf(&_SparkRewardsController.TransactOpts, assets)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x236300dc.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to, address reward) returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerTransactor) ClaimRewards(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "claimRewards", assets, amount, to, reward)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x236300dc.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to, address reward) returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerSession) ClaimRewards(assets []common.Address, amount *big.Int, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimRewards(&_SparkRewardsController.TransactOpts, assets, amount, to, reward)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x236300dc.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to, address reward) returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) ClaimRewards(assets []common.Address, amount *big.Int, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimRewards(&_SparkRewardsController.TransactOpts, assets, amount, to, reward)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x33028b99.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to, address reward) returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerTransactor) ClaimRewardsOnBehalf(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, user common.Address, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "claimRewardsOnBehalf", assets, amount, user, to, reward)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x33028b99.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to, address reward) returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerSession) ClaimRewardsOnBehalf(assets []common.Address, amount *big.Int, user common.Address, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimRewardsOnBehalf(&_SparkRewardsController.TransactOpts, assets, amount, user, to, reward)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x33028b99.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to, address reward) returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) ClaimRewardsOnBehalf(assets []common.Address, amount *big.Int, user common.Address, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimRewardsOnBehalf(&_SparkRewardsController.TransactOpts, assets, amount, user, to, reward)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x57b89883.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount, address reward) returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerTransactor) ClaimRewardsToSelf(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, reward common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "claimRewardsToSelf", assets, amount, reward)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x57b89883.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount, address reward) returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerSession) ClaimRewardsToSelf(assets []common.Address, amount *big.Int, reward common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimRewardsToSelf(&_SparkRewardsController.TransactOpts, assets, amount, reward)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x57b89883.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount, address reward) returns(uint256)
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) ClaimRewardsToSelf(assets []common.Address, amount *big.Int, reward common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ClaimRewardsToSelf(&_SparkRewardsController.TransactOpts, assets, amount, reward)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x955c2ad7.
//
// Solidity: function configureAssets((uint88,uint256,uint32,address,address,address,address)[] config) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactor) ConfigureAssets(opts *bind.TransactOpts, config []RewardsDataTypesRewardsConfigInput) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "configureAssets", config)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x955c2ad7.
//
// Solidity: function configureAssets((uint88,uint256,uint32,address,address,address,address)[] config) returns()
func (_SparkRewardsController *SparkRewardsControllerSession) ConfigureAssets(config []RewardsDataTypesRewardsConfigInput) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ConfigureAssets(&_SparkRewardsController.TransactOpts, config)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x955c2ad7.
//
// Solidity: function configureAssets((uint88,uint256,uint32,address,address,address,address)[] config) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) ConfigureAssets(config []RewardsDataTypesRewardsConfigInput) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.ConfigureAssets(&_SparkRewardsController.TransactOpts, config)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactor) HandleAction(opts *bind.TransactOpts, user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "handleAction", user, totalSupply, userBalance)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_SparkRewardsController *SparkRewardsControllerSession) HandleAction(user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.HandleAction(&_SparkRewardsController.TransactOpts, user, totalSupply, userBalance)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) HandleAction(user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.HandleAction(&_SparkRewardsController.TransactOpts, user, totalSupply, userBalance)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address claimer) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactor) SetClaimer(opts *bind.TransactOpts, user common.Address, claimer common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "setClaimer", user, claimer)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address claimer) returns()
func (_SparkRewardsController *SparkRewardsControllerSession) SetClaimer(user common.Address, claimer common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SetClaimer(&_SparkRewardsController.TransactOpts, user, claimer)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address claimer) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) SetClaimer(user common.Address, claimer common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SetClaimer(&_SparkRewardsController.TransactOpts, user, claimer)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0xc5a7b538.
//
// Solidity: function setDistributionEnd(address asset, address reward, uint32 newDistributionEnd) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactor) SetDistributionEnd(opts *bind.TransactOpts, asset common.Address, reward common.Address, newDistributionEnd uint32) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "setDistributionEnd", asset, reward, newDistributionEnd)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0xc5a7b538.
//
// Solidity: function setDistributionEnd(address asset, address reward, uint32 newDistributionEnd) returns()
func (_SparkRewardsController *SparkRewardsControllerSession) SetDistributionEnd(asset common.Address, reward common.Address, newDistributionEnd uint32) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SetDistributionEnd(&_SparkRewardsController.TransactOpts, asset, reward, newDistributionEnd)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0xc5a7b538.
//
// Solidity: function setDistributionEnd(address asset, address reward, uint32 newDistributionEnd) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) SetDistributionEnd(asset common.Address, reward common.Address, newDistributionEnd uint32) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SetDistributionEnd(&_SparkRewardsController.TransactOpts, asset, reward, newDistributionEnd)
}

// SetEmissionPerSecond is a paid mutator transaction binding the contract method 0xf996868b.
//
// Solidity: function setEmissionPerSecond(address asset, address[] rewards, uint88[] newEmissionsPerSecond) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactor) SetEmissionPerSecond(opts *bind.TransactOpts, asset common.Address, rewards []common.Address, newEmissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "setEmissionPerSecond", asset, rewards, newEmissionsPerSecond)
}

// SetEmissionPerSecond is a paid mutator transaction binding the contract method 0xf996868b.
//
// Solidity: function setEmissionPerSecond(address asset, address[] rewards, uint88[] newEmissionsPerSecond) returns()
func (_SparkRewardsController *SparkRewardsControllerSession) SetEmissionPerSecond(asset common.Address, rewards []common.Address, newEmissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SetEmissionPerSecond(&_SparkRewardsController.TransactOpts, asset, rewards, newEmissionsPerSecond)
}

// SetEmissionPerSecond is a paid mutator transaction binding the contract method 0xf996868b.
//
// Solidity: function setEmissionPerSecond(address asset, address[] rewards, uint88[] newEmissionsPerSecond) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) SetEmissionPerSecond(asset common.Address, rewards []common.Address, newEmissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SetEmissionPerSecond(&_SparkRewardsController.TransactOpts, asset, rewards, newEmissionsPerSecond)
}

// SetRewardOracle is a paid mutator transaction binding the contract method 0x5453ba10.
//
// Solidity: function setRewardOracle(address reward, address rewardOracle) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactor) SetRewardOracle(opts *bind.TransactOpts, reward common.Address, rewardOracle common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "setRewardOracle", reward, rewardOracle)
}

// SetRewardOracle is a paid mutator transaction binding the contract method 0x5453ba10.
//
// Solidity: function setRewardOracle(address reward, address rewardOracle) returns()
func (_SparkRewardsController *SparkRewardsControllerSession) SetRewardOracle(reward common.Address, rewardOracle common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SetRewardOracle(&_SparkRewardsController.TransactOpts, reward, rewardOracle)
}

// SetRewardOracle is a paid mutator transaction binding the contract method 0x5453ba10.
//
// Solidity: function setRewardOracle(address reward, address rewardOracle) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) SetRewardOracle(reward common.Address, rewardOracle common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SetRewardOracle(&_SparkRewardsController.TransactOpts, reward, rewardOracle)
}

// SetTransferStrategy is a paid mutator transaction binding the contract method 0xe15ac623.
//
// Solidity: function setTransferStrategy(address reward, address transferStrategy) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactor) SetTransferStrategy(opts *bind.TransactOpts, reward common.Address, transferStrategy common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.contract.Transact(opts, "setTransferStrategy", reward, transferStrategy)
}

// SetTransferStrategy is a paid mutator transaction binding the contract method 0xe15ac623.
//
// Solidity: function setTransferStrategy(address reward, address transferStrategy) returns()
func (_SparkRewardsController *SparkRewardsControllerSession) SetTransferStrategy(reward common.Address, transferStrategy common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SetTransferStrategy(&_SparkRewardsController.TransactOpts, reward, transferStrategy)
}

// SetTransferStrategy is a paid mutator transaction binding the contract method 0xe15ac623.
//
// Solidity: function setTransferStrategy(address reward, address transferStrategy) returns()
func (_SparkRewardsController *SparkRewardsControllerTransactorSession) SetTransferStrategy(reward common.Address, transferStrategy common.Address) (*types.Transaction, error) {
	return _SparkRewardsController.Contract.SetTransferStrategy(&_SparkRewardsController.TransactOpts, reward, transferStrategy)
}

// SparkRewardsControllerAccruedIterator is returned from FilterAccrued and is used to iterate over the raw logs and unpacked data for Accrued events raised by the SparkRewardsController contract.
type SparkRewardsControllerAccruedIterator struct {
	Event *SparkRewardsControllerAccrued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SparkRewardsControllerAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SparkRewardsControllerAccrued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SparkRewardsControllerAccrued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SparkRewardsControllerAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SparkRewardsControllerAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SparkRewardsControllerAccrued represents a Accrued event raised by the SparkRewardsController contract.
type SparkRewardsControllerAccrued struct {
	Asset          common.Address
	Reward         common.Address
	User           common.Address
	AssetIndex     *big.Int
	UserIndex      *big.Int
	RewardsAccrued *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccrued is a free log retrieval operation binding the contract event 0x3303facd24627943a92e9dc87cfbb34b15c49b726eec3ad3487c16be9ab8efe8.
//
// Solidity: event Accrued(address indexed asset, address indexed reward, address indexed user, uint256 assetIndex, uint256 userIndex, uint256 rewardsAccrued)
func (_SparkRewardsController *SparkRewardsControllerFilterer) FilterAccrued(opts *bind.FilterOpts, asset []common.Address, reward []common.Address, user []common.Address) (*SparkRewardsControllerAccruedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SparkRewardsController.contract.FilterLogs(opts, "Accrued", assetRule, rewardRule, userRule)
	if err != nil {
		return nil, err
	}
	return &SparkRewardsControllerAccruedIterator{contract: _SparkRewardsController.contract, event: "Accrued", logs: logs, sub: sub}, nil
}

// WatchAccrued is a free log subscription operation binding the contract event 0x3303facd24627943a92e9dc87cfbb34b15c49b726eec3ad3487c16be9ab8efe8.
//
// Solidity: event Accrued(address indexed asset, address indexed reward, address indexed user, uint256 assetIndex, uint256 userIndex, uint256 rewardsAccrued)
func (_SparkRewardsController *SparkRewardsControllerFilterer) WatchAccrued(opts *bind.WatchOpts, sink chan<- *SparkRewardsControllerAccrued, asset []common.Address, reward []common.Address, user []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SparkRewardsController.contract.WatchLogs(opts, "Accrued", assetRule, rewardRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SparkRewardsControllerAccrued)
				if err := _SparkRewardsController.contract.UnpackLog(event, "Accrued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccrued is a log parse operation binding the contract event 0x3303facd24627943a92e9dc87cfbb34b15c49b726eec3ad3487c16be9ab8efe8.
//
// Solidity: event Accrued(address indexed asset, address indexed reward, address indexed user, uint256 assetIndex, uint256 userIndex, uint256 rewardsAccrued)
func (_SparkRewardsController *SparkRewardsControllerFilterer) ParseAccrued(log types.Log) (*SparkRewardsControllerAccrued, error) {
	event := new(SparkRewardsControllerAccrued)
	if err := _SparkRewardsController.contract.UnpackLog(event, "Accrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SparkRewardsControllerAssetConfigUpdatedIterator is returned from FilterAssetConfigUpdated and is used to iterate over the raw logs and unpacked data for AssetConfigUpdated events raised by the SparkRewardsController contract.
type SparkRewardsControllerAssetConfigUpdatedIterator struct {
	Event *SparkRewardsControllerAssetConfigUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SparkRewardsControllerAssetConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SparkRewardsControllerAssetConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SparkRewardsControllerAssetConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SparkRewardsControllerAssetConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SparkRewardsControllerAssetConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SparkRewardsControllerAssetConfigUpdated represents a AssetConfigUpdated event raised by the SparkRewardsController contract.
type SparkRewardsControllerAssetConfigUpdated struct {
	Asset              common.Address
	Reward             common.Address
	OldEmission        *big.Int
	NewEmission        *big.Int
	OldDistributionEnd *big.Int
	NewDistributionEnd *big.Int
	AssetIndex         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAssetConfigUpdated is a free log retrieval operation binding the contract event 0xac1777479f07f3e7c34da8402139d54027a6a260caaae168bdee825ca5580dc5.
//
// Solidity: event AssetConfigUpdated(address indexed asset, address indexed reward, uint256 oldEmission, uint256 newEmission, uint256 oldDistributionEnd, uint256 newDistributionEnd, uint256 assetIndex)
func (_SparkRewardsController *SparkRewardsControllerFilterer) FilterAssetConfigUpdated(opts *bind.FilterOpts, asset []common.Address, reward []common.Address) (*SparkRewardsControllerAssetConfigUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _SparkRewardsController.contract.FilterLogs(opts, "AssetConfigUpdated", assetRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &SparkRewardsControllerAssetConfigUpdatedIterator{contract: _SparkRewardsController.contract, event: "AssetConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetConfigUpdated is a free log subscription operation binding the contract event 0xac1777479f07f3e7c34da8402139d54027a6a260caaae168bdee825ca5580dc5.
//
// Solidity: event AssetConfigUpdated(address indexed asset, address indexed reward, uint256 oldEmission, uint256 newEmission, uint256 oldDistributionEnd, uint256 newDistributionEnd, uint256 assetIndex)
func (_SparkRewardsController *SparkRewardsControllerFilterer) WatchAssetConfigUpdated(opts *bind.WatchOpts, sink chan<- *SparkRewardsControllerAssetConfigUpdated, asset []common.Address, reward []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _SparkRewardsController.contract.WatchLogs(opts, "AssetConfigUpdated", assetRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SparkRewardsControllerAssetConfigUpdated)
				if err := _SparkRewardsController.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetConfigUpdated is a log parse operation binding the contract event 0xac1777479f07f3e7c34da8402139d54027a6a260caaae168bdee825ca5580dc5.
//
// Solidity: event AssetConfigUpdated(address indexed asset, address indexed reward, uint256 oldEmission, uint256 newEmission, uint256 oldDistributionEnd, uint256 newDistributionEnd, uint256 assetIndex)
func (_SparkRewardsController *SparkRewardsControllerFilterer) ParseAssetConfigUpdated(log types.Log) (*SparkRewardsControllerAssetConfigUpdated, error) {
	event := new(SparkRewardsControllerAssetConfigUpdated)
	if err := _SparkRewardsController.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SparkRewardsControllerClaimerSetIterator is returned from FilterClaimerSet and is used to iterate over the raw logs and unpacked data for ClaimerSet events raised by the SparkRewardsController contract.
type SparkRewardsControllerClaimerSetIterator struct {
	Event *SparkRewardsControllerClaimerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SparkRewardsControllerClaimerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SparkRewardsControllerClaimerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SparkRewardsControllerClaimerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SparkRewardsControllerClaimerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SparkRewardsControllerClaimerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SparkRewardsControllerClaimerSet represents a ClaimerSet event raised by the SparkRewardsController contract.
type SparkRewardsControllerClaimerSet struct {
	User    common.Address
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimerSet is a free log retrieval operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_SparkRewardsController *SparkRewardsControllerFilterer) FilterClaimerSet(opts *bind.FilterOpts, user []common.Address, claimer []common.Address) (*SparkRewardsControllerClaimerSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _SparkRewardsController.contract.FilterLogs(opts, "ClaimerSet", userRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &SparkRewardsControllerClaimerSetIterator{contract: _SparkRewardsController.contract, event: "ClaimerSet", logs: logs, sub: sub}, nil
}

// WatchClaimerSet is a free log subscription operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_SparkRewardsController *SparkRewardsControllerFilterer) WatchClaimerSet(opts *bind.WatchOpts, sink chan<- *SparkRewardsControllerClaimerSet, user []common.Address, claimer []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _SparkRewardsController.contract.WatchLogs(opts, "ClaimerSet", userRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SparkRewardsControllerClaimerSet)
				if err := _SparkRewardsController.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimerSet is a log parse operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_SparkRewardsController *SparkRewardsControllerFilterer) ParseClaimerSet(log types.Log) (*SparkRewardsControllerClaimerSet, error) {
	event := new(SparkRewardsControllerClaimerSet)
	if err := _SparkRewardsController.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SparkRewardsControllerRewardOracleUpdatedIterator is returned from FilterRewardOracleUpdated and is used to iterate over the raw logs and unpacked data for RewardOracleUpdated events raised by the SparkRewardsController contract.
type SparkRewardsControllerRewardOracleUpdatedIterator struct {
	Event *SparkRewardsControllerRewardOracleUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SparkRewardsControllerRewardOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SparkRewardsControllerRewardOracleUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SparkRewardsControllerRewardOracleUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SparkRewardsControllerRewardOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SparkRewardsControllerRewardOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SparkRewardsControllerRewardOracleUpdated represents a RewardOracleUpdated event raised by the SparkRewardsController contract.
type SparkRewardsControllerRewardOracleUpdated struct {
	Reward       common.Address
	RewardOracle common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardOracleUpdated is a free log retrieval operation binding the contract event 0x1a1cd5483e52e60b9ff7f3b9d1db3bbd9e9d21c6324ad3a8c79dba9b75e62f4d.
//
// Solidity: event RewardOracleUpdated(address indexed reward, address indexed rewardOracle)
func (_SparkRewardsController *SparkRewardsControllerFilterer) FilterRewardOracleUpdated(opts *bind.FilterOpts, reward []common.Address, rewardOracle []common.Address) (*SparkRewardsControllerRewardOracleUpdatedIterator, error) {

	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var rewardOracleRule []interface{}
	for _, rewardOracleItem := range rewardOracle {
		rewardOracleRule = append(rewardOracleRule, rewardOracleItem)
	}

	logs, sub, err := _SparkRewardsController.contract.FilterLogs(opts, "RewardOracleUpdated", rewardRule, rewardOracleRule)
	if err != nil {
		return nil, err
	}
	return &SparkRewardsControllerRewardOracleUpdatedIterator{contract: _SparkRewardsController.contract, event: "RewardOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardOracleUpdated is a free log subscription operation binding the contract event 0x1a1cd5483e52e60b9ff7f3b9d1db3bbd9e9d21c6324ad3a8c79dba9b75e62f4d.
//
// Solidity: event RewardOracleUpdated(address indexed reward, address indexed rewardOracle)
func (_SparkRewardsController *SparkRewardsControllerFilterer) WatchRewardOracleUpdated(opts *bind.WatchOpts, sink chan<- *SparkRewardsControllerRewardOracleUpdated, reward []common.Address, rewardOracle []common.Address) (event.Subscription, error) {

	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var rewardOracleRule []interface{}
	for _, rewardOracleItem := range rewardOracle {
		rewardOracleRule = append(rewardOracleRule, rewardOracleItem)
	}

	logs, sub, err := _SparkRewardsController.contract.WatchLogs(opts, "RewardOracleUpdated", rewardRule, rewardOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SparkRewardsControllerRewardOracleUpdated)
				if err := _SparkRewardsController.contract.UnpackLog(event, "RewardOracleUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardOracleUpdated is a log parse operation binding the contract event 0x1a1cd5483e52e60b9ff7f3b9d1db3bbd9e9d21c6324ad3a8c79dba9b75e62f4d.
//
// Solidity: event RewardOracleUpdated(address indexed reward, address indexed rewardOracle)
func (_SparkRewardsController *SparkRewardsControllerFilterer) ParseRewardOracleUpdated(log types.Log) (*SparkRewardsControllerRewardOracleUpdated, error) {
	event := new(SparkRewardsControllerRewardOracleUpdated)
	if err := _SparkRewardsController.contract.UnpackLog(event, "RewardOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SparkRewardsControllerRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the SparkRewardsController contract.
type SparkRewardsControllerRewardsClaimedIterator struct {
	Event *SparkRewardsControllerRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SparkRewardsControllerRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SparkRewardsControllerRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SparkRewardsControllerRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SparkRewardsControllerRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SparkRewardsControllerRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SparkRewardsControllerRewardsClaimed represents a RewardsClaimed event raised by the SparkRewardsController contract.
type SparkRewardsControllerRewardsClaimed struct {
	User    common.Address
	Reward  common.Address
	To      common.Address
	Claimer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xc052130bc4ef84580db505783484b067ea8b71b3bca78a7e12db7aea8658f004.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed reward, address indexed to, address claimer, uint256 amount)
func (_SparkRewardsController *SparkRewardsControllerFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address, reward []common.Address, to []common.Address) (*SparkRewardsControllerRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SparkRewardsController.contract.FilterLogs(opts, "RewardsClaimed", userRule, rewardRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SparkRewardsControllerRewardsClaimedIterator{contract: _SparkRewardsController.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xc052130bc4ef84580db505783484b067ea8b71b3bca78a7e12db7aea8658f004.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed reward, address indexed to, address claimer, uint256 amount)
func (_SparkRewardsController *SparkRewardsControllerFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *SparkRewardsControllerRewardsClaimed, user []common.Address, reward []common.Address, to []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SparkRewardsController.contract.WatchLogs(opts, "RewardsClaimed", userRule, rewardRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SparkRewardsControllerRewardsClaimed)
				if err := _SparkRewardsController.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsClaimed is a log parse operation binding the contract event 0xc052130bc4ef84580db505783484b067ea8b71b3bca78a7e12db7aea8658f004.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed reward, address indexed to, address claimer, uint256 amount)
func (_SparkRewardsController *SparkRewardsControllerFilterer) ParseRewardsClaimed(log types.Log) (*SparkRewardsControllerRewardsClaimed, error) {
	event := new(SparkRewardsControllerRewardsClaimed)
	if err := _SparkRewardsController.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SparkRewardsControllerTransferStrategyInstalledIterator is returned from FilterTransferStrategyInstalled and is used to iterate over the raw logs and unpacked data for TransferStrategyInstalled events raised by the SparkRewardsController contract.
type SparkRewardsControllerTransferStrategyInstalledIterator struct {
	Event *SparkRewardsControllerTransferStrategyInstalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SparkRewardsControllerTransferStrategyInstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SparkRewardsControllerTransferStrategyInstalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SparkRewardsControllerTransferStrategyInstalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SparkRewardsControllerTransferStrategyInstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SparkRewardsControllerTransferStrategyInstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SparkRewardsControllerTransferStrategyInstalled represents a TransferStrategyInstalled event raised by the SparkRewardsController contract.
type SparkRewardsControllerTransferStrategyInstalled struct {
	Reward           common.Address
	TransferStrategy common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTransferStrategyInstalled is a free log retrieval operation binding the contract event 0x8ca1d928f1d72493a6b78c4f74aabde976bc37ffe2570f2a1ce5a8abd3dde0aa.
//
// Solidity: event TransferStrategyInstalled(address indexed reward, address indexed transferStrategy)
func (_SparkRewardsController *SparkRewardsControllerFilterer) FilterTransferStrategyInstalled(opts *bind.FilterOpts, reward []common.Address, transferStrategy []common.Address) (*SparkRewardsControllerTransferStrategyInstalledIterator, error) {

	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var transferStrategyRule []interface{}
	for _, transferStrategyItem := range transferStrategy {
		transferStrategyRule = append(transferStrategyRule, transferStrategyItem)
	}

	logs, sub, err := _SparkRewardsController.contract.FilterLogs(opts, "TransferStrategyInstalled", rewardRule, transferStrategyRule)
	if err != nil {
		return nil, err
	}
	return &SparkRewardsControllerTransferStrategyInstalledIterator{contract: _SparkRewardsController.contract, event: "TransferStrategyInstalled", logs: logs, sub: sub}, nil
}

// WatchTransferStrategyInstalled is a free log subscription operation binding the contract event 0x8ca1d928f1d72493a6b78c4f74aabde976bc37ffe2570f2a1ce5a8abd3dde0aa.
//
// Solidity: event TransferStrategyInstalled(address indexed reward, address indexed transferStrategy)
func (_SparkRewardsController *SparkRewardsControllerFilterer) WatchTransferStrategyInstalled(opts *bind.WatchOpts, sink chan<- *SparkRewardsControllerTransferStrategyInstalled, reward []common.Address, transferStrategy []common.Address) (event.Subscription, error) {

	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var transferStrategyRule []interface{}
	for _, transferStrategyItem := range transferStrategy {
		transferStrategyRule = append(transferStrategyRule, transferStrategyItem)
	}

	logs, sub, err := _SparkRewardsController.contract.WatchLogs(opts, "TransferStrategyInstalled", rewardRule, transferStrategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SparkRewardsControllerTransferStrategyInstalled)
				if err := _SparkRewardsController.contract.UnpackLog(event, "TransferStrategyInstalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferStrategyInstalled is a log parse operation binding the contract event 0x8ca1d928f1d72493a6b78c4f74aabde976bc37ffe2570f2a1ce5a8abd3dde0aa.
//
// Solidity: event TransferStrategyInstalled(address indexed reward, address indexed transferStrategy)
func (_SparkRewardsController *SparkRewardsControllerFilterer) ParseTransferStrategyInstalled(log types.Log) (*SparkRewardsControllerTransferStrategyInstalled, error) {
	event := new(SparkRewardsControllerTransferStrategyInstalled)
	if err := _SparkRewardsController.contract.UnpackLog(event, "TransferStrategyInstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
