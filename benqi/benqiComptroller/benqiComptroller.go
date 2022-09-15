// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package benqiComptroller

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

// BenqiComptrollerMetaData contains all meta data concerning the BenqiComptroller contract.
var BenqiComptrollerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowRewardSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"qiToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rewardType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rewardType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"qiTokens\",\"type\":\"address[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rewardType\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"qiTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"qiTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"qiTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialIndexConstant\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"qiTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"qiTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"qiTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isQied\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qiAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardAvax\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardQi\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"qiTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"qiTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supplyRewardSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"qiToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BenqiComptrollerABI is the input ABI used to generate the binding from.
// Deprecated: Use BenqiComptrollerMetaData.ABI instead.
var BenqiComptrollerABI = BenqiComptrollerMetaData.ABI

// BenqiComptroller is an auto generated Go binding around an Ethereum contract.
type BenqiComptroller struct {
	BenqiComptrollerCaller     // Read-only binding to the contract
	BenqiComptrollerTransactor // Write-only binding to the contract
	BenqiComptrollerFilterer   // Log filterer for contract events
}

// BenqiComptrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BenqiComptrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqiComptrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BenqiComptrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqiComptrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BenqiComptrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenqiComptrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BenqiComptrollerSession struct {
	Contract     *BenqiComptroller // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BenqiComptrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BenqiComptrollerCallerSession struct {
	Contract *BenqiComptrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BenqiComptrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BenqiComptrollerTransactorSession struct {
	Contract     *BenqiComptrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BenqiComptrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BenqiComptrollerRaw struct {
	Contract *BenqiComptroller // Generic contract binding to access the raw methods on
}

// BenqiComptrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BenqiComptrollerCallerRaw struct {
	Contract *BenqiComptrollerCaller // Generic read-only contract binding to access the raw methods on
}

// BenqiComptrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BenqiComptrollerTransactorRaw struct {
	Contract *BenqiComptrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBenqiComptroller creates a new instance of BenqiComptroller, bound to a specific deployed contract.
func NewBenqiComptroller(address common.Address, backend bind.ContractBackend) (*BenqiComptroller, error) {
	contract, err := bindBenqiComptroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BenqiComptroller{BenqiComptrollerCaller: BenqiComptrollerCaller{contract: contract}, BenqiComptrollerTransactor: BenqiComptrollerTransactor{contract: contract}, BenqiComptrollerFilterer: BenqiComptrollerFilterer{contract: contract}}, nil
}

// NewBenqiComptrollerCaller creates a new read-only instance of BenqiComptroller, bound to a specific deployed contract.
func NewBenqiComptrollerCaller(address common.Address, caller bind.ContractCaller) (*BenqiComptrollerCaller, error) {
	contract, err := bindBenqiComptroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BenqiComptrollerCaller{contract: contract}, nil
}

// NewBenqiComptrollerTransactor creates a new write-only instance of BenqiComptroller, bound to a specific deployed contract.
func NewBenqiComptrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*BenqiComptrollerTransactor, error) {
	contract, err := bindBenqiComptroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BenqiComptrollerTransactor{contract: contract}, nil
}

// NewBenqiComptrollerFilterer creates a new log filterer instance of BenqiComptroller, bound to a specific deployed contract.
func NewBenqiComptrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*BenqiComptrollerFilterer, error) {
	contract, err := bindBenqiComptroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BenqiComptrollerFilterer{contract: contract}, nil
}

// bindBenqiComptroller binds a generic wrapper to an already deployed contract.
func bindBenqiComptroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BenqiComptrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BenqiComptroller *BenqiComptrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BenqiComptroller.Contract.BenqiComptrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BenqiComptroller *BenqiComptrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.BenqiComptrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BenqiComptroller *BenqiComptrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.BenqiComptrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BenqiComptroller *BenqiComptrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BenqiComptroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BenqiComptroller *BenqiComptrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BenqiComptroller *BenqiComptrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.contract.Transact(opts, method, params...)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_BenqiComptroller *BenqiComptrollerCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "accountAssets", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_BenqiComptroller *BenqiComptrollerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _BenqiComptroller.Contract.AccountAssets(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_BenqiComptroller *BenqiComptrollerCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _BenqiComptroller.Contract.AccountAssets(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BenqiComptroller *BenqiComptrollerSession) Admin() (common.Address, error) {
	return _BenqiComptroller.Contract.Admin(&_BenqiComptroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCallerSession) Admin() (common.Address, error) {
	return _BenqiComptroller.Contract.Admin(&_BenqiComptroller.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_BenqiComptroller *BenqiComptrollerCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_BenqiComptroller *BenqiComptrollerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _BenqiComptroller.Contract.AllMarkets(&_BenqiComptroller.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_BenqiComptroller *BenqiComptrollerCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _BenqiComptroller.Contract.AllMarkets(&_BenqiComptroller.CallOpts, arg0)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCaller) BorrowCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "borrowCapGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_BenqiComptroller *BenqiComptrollerSession) BorrowCapGuardian() (common.Address, error) {
	return _BenqiComptroller.Contract.BorrowCapGuardian(&_BenqiComptroller.CallOpts)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCallerSession) BorrowCapGuardian() (common.Address, error) {
	return _BenqiComptroller.Contract.BorrowCapGuardian(&_BenqiComptroller.CallOpts)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) BorrowCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "borrowCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.BorrowCaps(&_BenqiComptroller.CallOpts, arg0)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.BorrowCaps(&_BenqiComptroller.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCaller) BorrowGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "borrowGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_BenqiComptroller *BenqiComptrollerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _BenqiComptroller.Contract.BorrowGuardianPaused(&_BenqiComptroller.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCallerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _BenqiComptroller.Contract.BorrowGuardianPaused(&_BenqiComptroller.CallOpts, arg0)
}

// BorrowRewardSpeeds is a free data retrieval call binding the contract method 0xc376fada.
//
// Solidity: function borrowRewardSpeeds(uint8 , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) BorrowRewardSpeeds(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "borrowRewardSpeeds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRewardSpeeds is a free data retrieval call binding the contract method 0xc376fada.
//
// Solidity: function borrowRewardSpeeds(uint8 , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) BorrowRewardSpeeds(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.BorrowRewardSpeeds(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// BorrowRewardSpeeds is a free data retrieval call binding the contract method 0xc376fada.
//
// Solidity: function borrowRewardSpeeds(uint8 , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) BorrowRewardSpeeds(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.BorrowRewardSpeeds(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address qiToken) view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCaller) CheckMembership(opts *bind.CallOpts, account common.Address, qiToken common.Address) (bool, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "checkMembership", account, qiToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address qiToken) view returns(bool)
func (_BenqiComptroller *BenqiComptrollerSession) CheckMembership(account common.Address, qiToken common.Address) (bool, error) {
	return _BenqiComptroller.Contract.CheckMembership(&_BenqiComptroller.CallOpts, account, qiToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address qiToken) view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCallerSession) CheckMembership(account common.Address, qiToken common.Address) (bool, error) {
	return _BenqiComptroller.Contract.CheckMembership(&_BenqiComptroller.CallOpts, account, qiToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "closeFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) CloseFactorMantissa() (*big.Int, error) {
	return _BenqiComptroller.Contract.CloseFactorMantissa(&_BenqiComptroller.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _BenqiComptroller.Contract.CloseFactorMantissa(&_BenqiComptroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_BenqiComptroller *BenqiComptrollerSession) ComptrollerImplementation() (common.Address, error) {
	return _BenqiComptroller.Contract.ComptrollerImplementation(&_BenqiComptroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _BenqiComptroller.Contract.ComptrollerImplementation(&_BenqiComptroller.CallOpts)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "getAccountLiquidity", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_BenqiComptroller *BenqiComptrollerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _BenqiComptroller.Contract.GetAccountLiquidity(&_BenqiComptroller.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _BenqiComptroller.Contract.GetAccountLiquidity(&_BenqiComptroller.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_BenqiComptroller *BenqiComptrollerCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_BenqiComptroller *BenqiComptrollerSession) GetAllMarkets() ([]common.Address, error) {
	return _BenqiComptroller.Contract.GetAllMarkets(&_BenqiComptroller.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_BenqiComptroller *BenqiComptrollerCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _BenqiComptroller.Contract.GetAllMarkets(&_BenqiComptroller.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_BenqiComptroller *BenqiComptrollerCaller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "getAssetsIn", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_BenqiComptroller *BenqiComptrollerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _BenqiComptroller.Contract.GetAssetsIn(&_BenqiComptroller.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_BenqiComptroller *BenqiComptrollerCallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _BenqiComptroller.Contract.GetAssetsIn(&_BenqiComptroller.CallOpts, account)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) GetBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "getBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) GetBlockTimestamp() (*big.Int, error) {
	return _BenqiComptroller.Contract.GetBlockTimestamp(&_BenqiComptroller.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) GetBlockTimestamp() (*big.Int, error) {
	return _BenqiComptroller.Contract.GetBlockTimestamp(&_BenqiComptroller.CallOpts)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address qiTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) GetHypotheticalAccountLiquidity(opts *bind.CallOpts, account common.Address, qiTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "getHypotheticalAccountLiquidity", account, qiTokenModify, redeemTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address qiTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_BenqiComptroller *BenqiComptrollerSession) GetHypotheticalAccountLiquidity(account common.Address, qiTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _BenqiComptroller.Contract.GetHypotheticalAccountLiquidity(&_BenqiComptroller.CallOpts, account, qiTokenModify, redeemTokens, borrowAmount)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address qiTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) GetHypotheticalAccountLiquidity(account common.Address, qiTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _BenqiComptroller.Contract.GetHypotheticalAccountLiquidity(&_BenqiComptroller.CallOpts, account, qiTokenModify, redeemTokens, borrowAmount)
}

// InitialIndexConstant is a free data retrieval call binding the contract method 0xed302dfd.
//
// Solidity: function initialIndexConstant() view returns(uint224)
func (_BenqiComptroller *BenqiComptrollerCaller) InitialIndexConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "initialIndexConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialIndexConstant is a free data retrieval call binding the contract method 0xed302dfd.
//
// Solidity: function initialIndexConstant() view returns(uint224)
func (_BenqiComptroller *BenqiComptrollerSession) InitialIndexConstant() (*big.Int, error) {
	return _BenqiComptroller.Contract.InitialIndexConstant(&_BenqiComptroller.CallOpts)
}

// InitialIndexConstant is a free data retrieval call binding the contract method 0xed302dfd.
//
// Solidity: function initialIndexConstant() view returns(uint224)
func (_BenqiComptroller *BenqiComptrollerCallerSession) InitialIndexConstant() (*big.Int, error) {
	return _BenqiComptroller.Contract.InitialIndexConstant(&_BenqiComptroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "isComptroller")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_BenqiComptroller *BenqiComptrollerSession) IsComptroller() (bool, error) {
	return _BenqiComptroller.Contract.IsComptroller(&_BenqiComptroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCallerSession) IsComptroller() (bool, error) {
	return _BenqiComptroller.Contract.IsComptroller(&_BenqiComptroller.CallOpts)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address qiTokenBorrowed, address qiTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, qiTokenBorrowed common.Address, qiTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", qiTokenBorrowed, qiTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address qiTokenBorrowed, address qiTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_BenqiComptroller *BenqiComptrollerSession) LiquidateCalculateSeizeTokens(qiTokenBorrowed common.Address, qiTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _BenqiComptroller.Contract.LiquidateCalculateSeizeTokens(&_BenqiComptroller.CallOpts, qiTokenBorrowed, qiTokenCollateral, actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address qiTokenBorrowed, address qiTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) LiquidateCalculateSeizeTokens(qiTokenBorrowed common.Address, qiTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _BenqiComptroller.Contract.LiquidateCalculateSeizeTokens(&_BenqiComptroller.CallOpts, qiTokenBorrowed, qiTokenCollateral, actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "liquidationIncentiveMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _BenqiComptroller.Contract.LiquidationIncentiveMantissa(&_BenqiComptroller.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _BenqiComptroller.Contract.LiquidationIncentiveMantissa(&_BenqiComptroller.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isQied)
func (_BenqiComptroller *BenqiComptrollerCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsQied                   bool
}, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
		IsQied                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsListed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CollateralFactorMantissa = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsQied = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isQied)
func (_BenqiComptroller *BenqiComptrollerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsQied                   bool
}, error) {
	return _BenqiComptroller.Contract.Markets(&_BenqiComptroller.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isQied)
func (_BenqiComptroller *BenqiComptrollerCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsQied                   bool
}, error) {
	return _BenqiComptroller.Contract.Markets(&_BenqiComptroller.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "maxAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) MaxAssets() (*big.Int, error) {
	return _BenqiComptroller.Contract.MaxAssets(&_BenqiComptroller.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) MaxAssets() (*big.Int, error) {
	return _BenqiComptroller.Contract.MaxAssets(&_BenqiComptroller.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCaller) MintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "mintGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_BenqiComptroller *BenqiComptrollerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _BenqiComptroller.Contract.MintGuardianPaused(&_BenqiComptroller.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCallerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _BenqiComptroller.Contract.MintGuardianPaused(&_BenqiComptroller.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BenqiComptroller *BenqiComptrollerSession) Oracle() (common.Address, error) {
	return _BenqiComptroller.Contract.Oracle(&_BenqiComptroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCallerSession) Oracle() (common.Address, error) {
	return _BenqiComptroller.Contract.Oracle(&_BenqiComptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_BenqiComptroller *BenqiComptrollerSession) PauseGuardian() (common.Address, error) {
	return _BenqiComptroller.Contract.PauseGuardian(&_BenqiComptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCallerSession) PauseGuardian() (common.Address, error) {
	return _BenqiComptroller.Contract.PauseGuardian(&_BenqiComptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_BenqiComptroller *BenqiComptrollerSession) PendingAdmin() (common.Address, error) {
	return _BenqiComptroller.Contract.PendingAdmin(&_BenqiComptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _BenqiComptroller.Contract.PendingAdmin(&_BenqiComptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_BenqiComptroller *BenqiComptrollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _BenqiComptroller.Contract.PendingComptrollerImplementation(&_BenqiComptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _BenqiComptroller.Contract.PendingComptrollerImplementation(&_BenqiComptroller.CallOpts)
}

// QiAddress is a free data retrieval call binding the contract method 0x26d71f1e.
//
// Solidity: function qiAddress() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCaller) QiAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "qiAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QiAddress is a free data retrieval call binding the contract method 0x26d71f1e.
//
// Solidity: function qiAddress() view returns(address)
func (_BenqiComptroller *BenqiComptrollerSession) QiAddress() (common.Address, error) {
	return _BenqiComptroller.Contract.QiAddress(&_BenqiComptroller.CallOpts)
}

// QiAddress is a free data retrieval call binding the contract method 0x26d71f1e.
//
// Solidity: function qiAddress() view returns(address)
func (_BenqiComptroller *BenqiComptrollerCallerSession) QiAddress() (common.Address, error) {
	return _BenqiComptroller.Contract.QiAddress(&_BenqiComptroller.CallOpts)
}

// RewardAccrued is a free data retrieval call binding the contract method 0x05b9783d.
//
// Solidity: function rewardAccrued(uint8 , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) RewardAccrued(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "rewardAccrued", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardAccrued is a free data retrieval call binding the contract method 0x05b9783d.
//
// Solidity: function rewardAccrued(uint8 , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) RewardAccrued(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.RewardAccrued(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// RewardAccrued is a free data retrieval call binding the contract method 0x05b9783d.
//
// Solidity: function rewardAccrued(uint8 , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) RewardAccrued(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.RewardAccrued(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// RewardAvax is a free data retrieval call binding the contract method 0x0b8d87df.
//
// Solidity: function rewardAvax() view returns(uint8)
func (_BenqiComptroller *BenqiComptrollerCaller) RewardAvax(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "rewardAvax")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RewardAvax is a free data retrieval call binding the contract method 0x0b8d87df.
//
// Solidity: function rewardAvax() view returns(uint8)
func (_BenqiComptroller *BenqiComptrollerSession) RewardAvax() (uint8, error) {
	return _BenqiComptroller.Contract.RewardAvax(&_BenqiComptroller.CallOpts)
}

// RewardAvax is a free data retrieval call binding the contract method 0x0b8d87df.
//
// Solidity: function rewardAvax() view returns(uint8)
func (_BenqiComptroller *BenqiComptrollerCallerSession) RewardAvax() (uint8, error) {
	return _BenqiComptroller.Contract.RewardAvax(&_BenqiComptroller.CallOpts)
}

// RewardBorrowState is a free data retrieval call binding the contract method 0x4b3a0a74.
//
// Solidity: function rewardBorrowState(uint8 , address ) view returns(uint224 index, uint32 timestamp)
func (_BenqiComptroller *BenqiComptrollerCaller) RewardBorrowState(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "rewardBorrowState", arg0, arg1)

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
func (_BenqiComptroller *BenqiComptrollerSession) RewardBorrowState(arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	return _BenqiComptroller.Contract.RewardBorrowState(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// RewardBorrowState is a free data retrieval call binding the contract method 0x4b3a0a74.
//
// Solidity: function rewardBorrowState(uint8 , address ) view returns(uint224 index, uint32 timestamp)
func (_BenqiComptroller *BenqiComptrollerCallerSession) RewardBorrowState(arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	return _BenqiComptroller.Contract.RewardBorrowState(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// RewardBorrowerIndex is a free data retrieval call binding the contract method 0x7937969d.
//
// Solidity: function rewardBorrowerIndex(uint8 , address , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) RewardBorrowerIndex(opts *bind.CallOpts, arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "rewardBorrowerIndex", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardBorrowerIndex is a free data retrieval call binding the contract method 0x7937969d.
//
// Solidity: function rewardBorrowerIndex(uint8 , address , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) RewardBorrowerIndex(arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.RewardBorrowerIndex(&_BenqiComptroller.CallOpts, arg0, arg1, arg2)
}

// RewardBorrowerIndex is a free data retrieval call binding the contract method 0x7937969d.
//
// Solidity: function rewardBorrowerIndex(uint8 , address , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) RewardBorrowerIndex(arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.RewardBorrowerIndex(&_BenqiComptroller.CallOpts, arg0, arg1, arg2)
}

// RewardQi is a free data retrieval call binding the contract method 0xceeca693.
//
// Solidity: function rewardQi() view returns(uint8)
func (_BenqiComptroller *BenqiComptrollerCaller) RewardQi(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "rewardQi")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RewardQi is a free data retrieval call binding the contract method 0xceeca693.
//
// Solidity: function rewardQi() view returns(uint8)
func (_BenqiComptroller *BenqiComptrollerSession) RewardQi() (uint8, error) {
	return _BenqiComptroller.Contract.RewardQi(&_BenqiComptroller.CallOpts)
}

// RewardQi is a free data retrieval call binding the contract method 0xceeca693.
//
// Solidity: function rewardQi() view returns(uint8)
func (_BenqiComptroller *BenqiComptrollerCallerSession) RewardQi() (uint8, error) {
	return _BenqiComptroller.Contract.RewardQi(&_BenqiComptroller.CallOpts)
}

// RewardSupplierIndex is a free data retrieval call binding the contract method 0x88e972b8.
//
// Solidity: function rewardSupplierIndex(uint8 , address , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) RewardSupplierIndex(opts *bind.CallOpts, arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "rewardSupplierIndex", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardSupplierIndex is a free data retrieval call binding the contract method 0x88e972b8.
//
// Solidity: function rewardSupplierIndex(uint8 , address , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) RewardSupplierIndex(arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.RewardSupplierIndex(&_BenqiComptroller.CallOpts, arg0, arg1, arg2)
}

// RewardSupplierIndex is a free data retrieval call binding the contract method 0x88e972b8.
//
// Solidity: function rewardSupplierIndex(uint8 , address , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) RewardSupplierIndex(arg0 uint8, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.RewardSupplierIndex(&_BenqiComptroller.CallOpts, arg0, arg1, arg2)
}

// RewardSupplyState is a free data retrieval call binding the contract method 0xd81c5e45.
//
// Solidity: function rewardSupplyState(uint8 , address ) view returns(uint224 index, uint32 timestamp)
func (_BenqiComptroller *BenqiComptrollerCaller) RewardSupplyState(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "rewardSupplyState", arg0, arg1)

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
func (_BenqiComptroller *BenqiComptrollerSession) RewardSupplyState(arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	return _BenqiComptroller.Contract.RewardSupplyState(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// RewardSupplyState is a free data retrieval call binding the contract method 0xd81c5e45.
//
// Solidity: function rewardSupplyState(uint8 , address ) view returns(uint224 index, uint32 timestamp)
func (_BenqiComptroller *BenqiComptrollerCallerSession) RewardSupplyState(arg0 uint8, arg1 common.Address) (struct {
	Index     *big.Int
	Timestamp uint32
}, error) {
	return _BenqiComptroller.Contract.RewardSupplyState(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "seizeGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_BenqiComptroller *BenqiComptrollerSession) SeizeGuardianPaused() (bool, error) {
	return _BenqiComptroller.Contract.SeizeGuardianPaused(&_BenqiComptroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCallerSession) SeizeGuardianPaused() (bool, error) {
	return _BenqiComptroller.Contract.SeizeGuardianPaused(&_BenqiComptroller.CallOpts)
}

// SupplyRewardSpeeds is a free data retrieval call binding the contract method 0xcf9cfb61.
//
// Solidity: function supplyRewardSpeeds(uint8 , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCaller) SupplyRewardSpeeds(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "supplyRewardSpeeds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRewardSpeeds is a free data retrieval call binding the contract method 0xcf9cfb61.
//
// Solidity: function supplyRewardSpeeds(uint8 , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) SupplyRewardSpeeds(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.SupplyRewardSpeeds(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// SupplyRewardSpeeds is a free data retrieval call binding the contract method 0xcf9cfb61.
//
// Solidity: function supplyRewardSpeeds(uint8 , address ) view returns(uint256)
func (_BenqiComptroller *BenqiComptrollerCallerSession) SupplyRewardSpeeds(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _BenqiComptroller.Contract.SupplyRewardSpeeds(&_BenqiComptroller.CallOpts, arg0, arg1)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BenqiComptroller.contract.Call(opts, &out, "transferGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_BenqiComptroller *BenqiComptrollerSession) TransferGuardianPaused() (bool, error) {
	return _BenqiComptroller.Contract.TransferGuardianPaused(&_BenqiComptroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_BenqiComptroller *BenqiComptrollerCallerSession) TransferGuardianPaused() (bool, error) {
	return _BenqiComptroller.Contract.TransferGuardianPaused(&_BenqiComptroller.CallOpts)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address qiToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactor) BorrowAllowed(opts *bind.TransactOpts, qiToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "borrowAllowed", qiToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address qiToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) BorrowAllowed(qiToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.BorrowAllowed(&_BenqiComptroller.TransactOpts, qiToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address qiToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactorSession) BorrowAllowed(qiToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.BorrowAllowed(&_BenqiComptroller.TransactOpts, qiToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address qiToken, address borrower, uint256 borrowAmount) returns()
func (_BenqiComptroller *BenqiComptrollerTransactor) BorrowVerify(opts *bind.TransactOpts, qiToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "borrowVerify", qiToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address qiToken, address borrower, uint256 borrowAmount) returns()
func (_BenqiComptroller *BenqiComptrollerSession) BorrowVerify(qiToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.BorrowVerify(&_BenqiComptroller.TransactOpts, qiToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address qiToken, address borrower, uint256 borrowAmount) returns()
func (_BenqiComptroller *BenqiComptrollerTransactorSession) BorrowVerify(qiToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.BorrowVerify(&_BenqiComptroller.TransactOpts, qiToken, borrower, borrowAmount)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0952c563.
//
// Solidity: function claimReward(uint8 rewardType, address holder) returns()
func (_BenqiComptroller *BenqiComptrollerTransactor) ClaimReward(opts *bind.TransactOpts, rewardType uint8, holder common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "claimReward", rewardType, holder)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0952c563.
//
// Solidity: function claimReward(uint8 rewardType, address holder) returns()
func (_BenqiComptroller *BenqiComptrollerSession) ClaimReward(rewardType uint8, holder common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.ClaimReward(&_BenqiComptroller.TransactOpts, rewardType, holder)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x0952c563.
//
// Solidity: function claimReward(uint8 rewardType, address holder) returns()
func (_BenqiComptroller *BenqiComptrollerTransactorSession) ClaimReward(rewardType uint8, holder common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.ClaimReward(&_BenqiComptroller.TransactOpts, rewardType, holder)
}

// ClaimReward0 is a paid mutator transaction binding the contract method 0x744532ae.
//
// Solidity: function claimReward(uint8 rewardType, address holder, address[] qiTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactor) ClaimReward0(opts *bind.TransactOpts, rewardType uint8, holder common.Address, qiTokens []common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "claimReward0", rewardType, holder, qiTokens)
}

// ClaimReward0 is a paid mutator transaction binding the contract method 0x744532ae.
//
// Solidity: function claimReward(uint8 rewardType, address holder, address[] qiTokens) returns()
func (_BenqiComptroller *BenqiComptrollerSession) ClaimReward0(rewardType uint8, holder common.Address, qiTokens []common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.ClaimReward0(&_BenqiComptroller.TransactOpts, rewardType, holder, qiTokens)
}

// ClaimReward0 is a paid mutator transaction binding the contract method 0x744532ae.
//
// Solidity: function claimReward(uint8 rewardType, address holder, address[] qiTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactorSession) ClaimReward0(rewardType uint8, holder common.Address, qiTokens []common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.ClaimReward0(&_BenqiComptroller.TransactOpts, rewardType, holder, qiTokens)
}

// ClaimReward1 is a paid mutator transaction binding the contract method 0x8805714b.
//
// Solidity: function claimReward(uint8 rewardType, address[] holders, address[] qiTokens, bool borrowers, bool suppliers) payable returns()
func (_BenqiComptroller *BenqiComptrollerTransactor) ClaimReward1(opts *bind.TransactOpts, rewardType uint8, holders []common.Address, qiTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "claimReward1", rewardType, holders, qiTokens, borrowers, suppliers)
}

// ClaimReward1 is a paid mutator transaction binding the contract method 0x8805714b.
//
// Solidity: function claimReward(uint8 rewardType, address[] holders, address[] qiTokens, bool borrowers, bool suppliers) payable returns()
func (_BenqiComptroller *BenqiComptrollerSession) ClaimReward1(rewardType uint8, holders []common.Address, qiTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.ClaimReward1(&_BenqiComptroller.TransactOpts, rewardType, holders, qiTokens, borrowers, suppliers)
}

// ClaimReward1 is a paid mutator transaction binding the contract method 0x8805714b.
//
// Solidity: function claimReward(uint8 rewardType, address[] holders, address[] qiTokens, bool borrowers, bool suppliers) payable returns()
func (_BenqiComptroller *BenqiComptrollerTransactorSession) ClaimReward1(rewardType uint8, holders []common.Address, qiTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.ClaimReward1(&_BenqiComptroller.TransactOpts, rewardType, holders, qiTokens, borrowers, suppliers)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] qiTokens) returns(uint256[])
func (_BenqiComptroller *BenqiComptrollerTransactor) EnterMarkets(opts *bind.TransactOpts, qiTokens []common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "enterMarkets", qiTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] qiTokens) returns(uint256[])
func (_BenqiComptroller *BenqiComptrollerSession) EnterMarkets(qiTokens []common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.EnterMarkets(&_BenqiComptroller.TransactOpts, qiTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] qiTokens) returns(uint256[])
func (_BenqiComptroller *BenqiComptrollerTransactorSession) EnterMarkets(qiTokens []common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.EnterMarkets(&_BenqiComptroller.TransactOpts, qiTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address qiTokenAddress) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactor) ExitMarket(opts *bind.TransactOpts, qiTokenAddress common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "exitMarket", qiTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address qiTokenAddress) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) ExitMarket(qiTokenAddress common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.ExitMarket(&_BenqiComptroller.TransactOpts, qiTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address qiTokenAddress) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactorSession) ExitMarket(qiTokenAddress common.Address) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.ExitMarket(&_BenqiComptroller.TransactOpts, qiTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address qiTokenBorrowed, address qiTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, qiTokenBorrowed common.Address, qiTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "liquidateBorrowAllowed", qiTokenBorrowed, qiTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address qiTokenBorrowed, address qiTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) LiquidateBorrowAllowed(qiTokenBorrowed common.Address, qiTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.LiquidateBorrowAllowed(&_BenqiComptroller.TransactOpts, qiTokenBorrowed, qiTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address qiTokenBorrowed, address qiTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactorSession) LiquidateBorrowAllowed(qiTokenBorrowed common.Address, qiTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.LiquidateBorrowAllowed(&_BenqiComptroller.TransactOpts, qiTokenBorrowed, qiTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address qiTokenBorrowed, address qiTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, qiTokenBorrowed common.Address, qiTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "liquidateBorrowVerify", qiTokenBorrowed, qiTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address qiTokenBorrowed, address qiTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_BenqiComptroller *BenqiComptrollerSession) LiquidateBorrowVerify(qiTokenBorrowed common.Address, qiTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.LiquidateBorrowVerify(&_BenqiComptroller.TransactOpts, qiTokenBorrowed, qiTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address qiTokenBorrowed, address qiTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactorSession) LiquidateBorrowVerify(qiTokenBorrowed common.Address, qiTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.LiquidateBorrowVerify(&_BenqiComptroller.TransactOpts, qiTokenBorrowed, qiTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address qiToken, address minter, uint256 mintAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactor) MintAllowed(opts *bind.TransactOpts, qiToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "mintAllowed", qiToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address qiToken, address minter, uint256 mintAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) MintAllowed(qiToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.MintAllowed(&_BenqiComptroller.TransactOpts, qiToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address qiToken, address minter, uint256 mintAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactorSession) MintAllowed(qiToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.MintAllowed(&_BenqiComptroller.TransactOpts, qiToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address qiToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactor) MintVerify(opts *bind.TransactOpts, qiToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "mintVerify", qiToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address qiToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_BenqiComptroller *BenqiComptrollerSession) MintVerify(qiToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.MintVerify(&_BenqiComptroller.TransactOpts, qiToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address qiToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactorSession) MintVerify(qiToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.MintVerify(&_BenqiComptroller.TransactOpts, qiToken, minter, actualMintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address qiToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactor) RedeemAllowed(opts *bind.TransactOpts, qiToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "redeemAllowed", qiToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address qiToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) RedeemAllowed(qiToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.RedeemAllowed(&_BenqiComptroller.TransactOpts, qiToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address qiToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactorSession) RedeemAllowed(qiToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.RedeemAllowed(&_BenqiComptroller.TransactOpts, qiToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address qiToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactor) RedeemVerify(opts *bind.TransactOpts, qiToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "redeemVerify", qiToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address qiToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_BenqiComptroller *BenqiComptrollerSession) RedeemVerify(qiToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.RedeemVerify(&_BenqiComptroller.TransactOpts, qiToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address qiToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactorSession) RedeemVerify(qiToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.RedeemVerify(&_BenqiComptroller.TransactOpts, qiToken, redeemer, redeemAmount, redeemTokens)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address qiToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, qiToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "repayBorrowAllowed", qiToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address qiToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) RepayBorrowAllowed(qiToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.RepayBorrowAllowed(&_BenqiComptroller.TransactOpts, qiToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address qiToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactorSession) RepayBorrowAllowed(qiToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.RepayBorrowAllowed(&_BenqiComptroller.TransactOpts, qiToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address qiToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_BenqiComptroller *BenqiComptrollerTransactor) RepayBorrowVerify(opts *bind.TransactOpts, qiToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "repayBorrowVerify", qiToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address qiToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_BenqiComptroller *BenqiComptrollerSession) RepayBorrowVerify(qiToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.RepayBorrowVerify(&_BenqiComptroller.TransactOpts, qiToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address qiToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_BenqiComptroller *BenqiComptrollerTransactorSession) RepayBorrowVerify(qiToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.RepayBorrowVerify(&_BenqiComptroller.TransactOpts, qiToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address qiTokenCollateral, address qiTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactor) SeizeAllowed(opts *bind.TransactOpts, qiTokenCollateral common.Address, qiTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "seizeAllowed", qiTokenCollateral, qiTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address qiTokenCollateral, address qiTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) SeizeAllowed(qiTokenCollateral common.Address, qiTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.SeizeAllowed(&_BenqiComptroller.TransactOpts, qiTokenCollateral, qiTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address qiTokenCollateral, address qiTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactorSession) SeizeAllowed(qiTokenCollateral common.Address, qiTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.SeizeAllowed(&_BenqiComptroller.TransactOpts, qiTokenCollateral, qiTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address qiTokenCollateral, address qiTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactor) SeizeVerify(opts *bind.TransactOpts, qiTokenCollateral common.Address, qiTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "seizeVerify", qiTokenCollateral, qiTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address qiTokenCollateral, address qiTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_BenqiComptroller *BenqiComptrollerSession) SeizeVerify(qiTokenCollateral common.Address, qiTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.SeizeVerify(&_BenqiComptroller.TransactOpts, qiTokenCollateral, qiTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address qiTokenCollateral, address qiTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactorSession) SeizeVerify(qiTokenCollateral common.Address, qiTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.SeizeVerify(&_BenqiComptroller.TransactOpts, qiTokenCollateral, qiTokenBorrowed, liquidator, borrower, seizeTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address qiToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactor) TransferAllowed(opts *bind.TransactOpts, qiToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "transferAllowed", qiToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address qiToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerSession) TransferAllowed(qiToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.TransferAllowed(&_BenqiComptroller.TransactOpts, qiToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address qiToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_BenqiComptroller *BenqiComptrollerTransactorSession) TransferAllowed(qiToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.TransferAllowed(&_BenqiComptroller.TransactOpts, qiToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address qiToken, address src, address dst, uint256 transferTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactor) TransferVerify(opts *bind.TransactOpts, qiToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.contract.Transact(opts, "transferVerify", qiToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address qiToken, address src, address dst, uint256 transferTokens) returns()
func (_BenqiComptroller *BenqiComptrollerSession) TransferVerify(qiToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.TransferVerify(&_BenqiComptroller.TransactOpts, qiToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address qiToken, address src, address dst, uint256 transferTokens) returns()
func (_BenqiComptroller *BenqiComptrollerTransactorSession) TransferVerify(qiToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _BenqiComptroller.Contract.TransferVerify(&_BenqiComptroller.TransactOpts, qiToken, src, dst, transferTokens)
}
