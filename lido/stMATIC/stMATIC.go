// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stMATIC

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

// stMATICRequestWithdraw is an auto generated low-level Go binding around an user-defined struct.
type stMATICRequestWithdraw struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}

// StMATICMetaData contains all meta data concerning the StMATIC contract.
var StMATICMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DAO\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNPAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePendingBufferedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingBufferedTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"claimTokensFromValidatorToContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInMatic\",\"type\":\"uint256\"}],\"name\":\"convertMaticToStMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInStMatic\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStMaticSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPooledMatic\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInStMatic\",\"type\":\"uint256\"}],\"name\":\"convertStMaticToMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInMatic\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStMaticAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPooledMatic\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationLowerBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entityFees\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"dao\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operators\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"insurance\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fxStateRootTunnel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"getLiquidRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getMaticFromTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getToken2WithdrawRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"internalType\":\"structstMATIC.RequestWithdraw[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPooledMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStakeAcrossAllValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalWithdrawRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"internalType\":\"structstMATIC.RequestWithdraw[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_insurance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poLidoNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fxStateRootTunnel\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insurance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastWithdrawnValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeOperatorRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poLidoNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalanceDelegatedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"requestWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardDistributionLowerBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDAO\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delegationLowerBound\",\"type\":\"uint256\"}],\"name\":\"setDelegationLowerBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_daoFee\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_operatorsFee\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_insuranceFee\",\"type\":\"uint8\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFxStateRootTunnel\",\"type\":\"address\"}],\"name\":\"setFxStateRootTunnel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setInsuranceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setNodeOperatorRegistryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newLidoNFT\",\"type\":\"address\"}],\"name\":\"setPoLidoNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_newProtocolFee\",\"type\":\"uint8\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRewardDistributionLowerBound\",\"type\":\"uint256\"}],\"name\":\"setRewardDistributionLowerBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newVersion\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stMaticWithdrawRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"token2WithdrawRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"token2WithdrawRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBuffered\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"withdrawTotalDelegated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StMATICABI is the input ABI used to generate the binding from.
// Deprecated: Use StMATICMetaData.ABI instead.
var StMATICABI = StMATICMetaData.ABI

// StMATIC is an auto generated Go binding around an Ethereum contract.
type StMATIC struct {
	StMATICCaller     // Read-only binding to the contract
	StMATICTransactor // Write-only binding to the contract
	StMATICFilterer   // Log filterer for contract events
}

// StMATICCaller is an auto generated read-only Go binding around an Ethereum contract.
type StMATICCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StMATICTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StMATICTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StMATICFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StMATICFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StMATICSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StMATICSession struct {
	Contract     *StMATIC          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StMATICCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StMATICCallerSession struct {
	Contract *StMATICCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StMATICTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StMATICTransactorSession struct {
	Contract     *StMATICTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StMATICRaw is an auto generated low-level Go binding around an Ethereum contract.
type StMATICRaw struct {
	Contract *StMATIC // Generic contract binding to access the raw methods on
}

// StMATICCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StMATICCallerRaw struct {
	Contract *StMATICCaller // Generic read-only contract binding to access the raw methods on
}

// StMATICTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StMATICTransactorRaw struct {
	Contract *StMATICTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStMATIC creates a new instance of StMATIC, bound to a specific deployed contract.
func NewStMATIC(address common.Address, backend bind.ContractBackend) (*StMATIC, error) {
	contract, err := bindStMATIC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StMATIC{StMATICCaller: StMATICCaller{contract: contract}, StMATICTransactor: StMATICTransactor{contract: contract}, StMATICFilterer: StMATICFilterer{contract: contract}}, nil
}

// NewStMATICCaller creates a new read-only instance of StMATIC, bound to a specific deployed contract.
func NewStMATICCaller(address common.Address, caller bind.ContractCaller) (*StMATICCaller, error) {
	contract, err := bindStMATIC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StMATICCaller{contract: contract}, nil
}

// NewStMATICTransactor creates a new write-only instance of StMATIC, bound to a specific deployed contract.
func NewStMATICTransactor(address common.Address, transactor bind.ContractTransactor) (*StMATICTransactor, error) {
	contract, err := bindStMATIC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StMATICTransactor{contract: contract}, nil
}

// NewStMATICFilterer creates a new log filterer instance of StMATIC, bound to a specific deployed contract.
func NewStMATICFilterer(address common.Address, filterer bind.ContractFilterer) (*StMATICFilterer, error) {
	contract, err := bindStMATIC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StMATICFilterer{contract: contract}, nil
}

// bindStMATIC binds a generic wrapper to an already deployed contract.
func bindStMATIC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StMATICABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StMATIC *StMATICRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StMATIC.Contract.StMATICCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StMATIC *StMATICRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StMATIC.Contract.StMATICTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StMATIC *StMATICRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StMATIC.Contract.StMATICTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StMATIC *StMATICCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StMATIC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StMATIC *StMATICTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StMATIC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StMATIC *StMATICTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StMATIC.Contract.contract.Transact(opts, method, params...)
}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_StMATIC *StMATICCaller) DAO(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "DAO")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_StMATIC *StMATICSession) DAO() ([32]byte, error) {
	return _StMATIC.Contract.DAO(&_StMATIC.CallOpts)
}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_StMATIC *StMATICCallerSession) DAO() ([32]byte, error) {
	return _StMATIC.Contract.DAO(&_StMATIC.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StMATIC *StMATICCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StMATIC *StMATICSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StMATIC.Contract.DEFAULTADMINROLE(&_StMATIC.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StMATIC *StMATICCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StMATIC.Contract.DEFAULTADMINROLE(&_StMATIC.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StMATIC *StMATICCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StMATIC *StMATICSession) PAUSEROLE() ([32]byte, error) {
	return _StMATIC.Contract.PAUSEROLE(&_StMATIC.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StMATIC *StMATICCallerSession) PAUSEROLE() ([32]byte, error) {
	return _StMATIC.Contract.PAUSEROLE(&_StMATIC.CallOpts)
}

// UNPAUSEROLE is a free data retrieval call binding the contract method 0x309756fb.
//
// Solidity: function UNPAUSE_ROLE() view returns(bytes32)
func (_StMATIC *StMATICCaller) UNPAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "UNPAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNPAUSEROLE is a free data retrieval call binding the contract method 0x309756fb.
//
// Solidity: function UNPAUSE_ROLE() view returns(bytes32)
func (_StMATIC *StMATICSession) UNPAUSEROLE() ([32]byte, error) {
	return _StMATIC.Contract.UNPAUSEROLE(&_StMATIC.CallOpts)
}

// UNPAUSEROLE is a free data retrieval call binding the contract method 0x309756fb.
//
// Solidity: function UNPAUSE_ROLE() view returns(bytes32)
func (_StMATIC *StMATICCallerSession) UNPAUSEROLE() ([32]byte, error) {
	return _StMATIC.Contract.UNPAUSEROLE(&_StMATIC.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StMATIC *StMATICCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StMATIC *StMATICSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StMATIC.Contract.Allowance(&_StMATIC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StMATIC *StMATICCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StMATIC.Contract.Allowance(&_StMATIC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StMATIC *StMATICCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StMATIC *StMATICSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StMATIC.Contract.BalanceOf(&_StMATIC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StMATIC *StMATICCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StMATIC.Contract.BalanceOf(&_StMATIC.CallOpts, account)
}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_StMATIC *StMATICCaller) CalculatePendingBufferedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "calculatePendingBufferedTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_StMATIC *StMATICSession) CalculatePendingBufferedTokens() (*big.Int, error) {
	return _StMATIC.Contract.CalculatePendingBufferedTokens(&_StMATIC.CallOpts)
}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_StMATIC *StMATICCallerSession) CalculatePendingBufferedTokens() (*big.Int, error) {
	return _StMATIC.Contract.CalculatePendingBufferedTokens(&_StMATIC.CallOpts)
}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _amountInMatic) view returns(uint256 amountInStMatic, uint256 totalStMaticSupply, uint256 totalPooledMatic)
func (_StMATIC *StMATICCaller) ConvertMaticToStMatic(opts *bind.CallOpts, _amountInMatic *big.Int) (struct {
	AmountInStMatic    *big.Int
	TotalStMaticSupply *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "convertMaticToStMatic", _amountInMatic)

	outstruct := new(struct {
		AmountInStMatic    *big.Int
		TotalStMaticSupply *big.Int
		TotalPooledMatic   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountInStMatic = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalStMaticSupply = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalPooledMatic = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _amountInMatic) view returns(uint256 amountInStMatic, uint256 totalStMaticSupply, uint256 totalPooledMatic)
func (_StMATIC *StMATICSession) ConvertMaticToStMatic(_amountInMatic *big.Int) (struct {
	AmountInStMatic    *big.Int
	TotalStMaticSupply *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	return _StMATIC.Contract.ConvertMaticToStMatic(&_StMATIC.CallOpts, _amountInMatic)
}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _amountInMatic) view returns(uint256 amountInStMatic, uint256 totalStMaticSupply, uint256 totalPooledMatic)
func (_StMATIC *StMATICCallerSession) ConvertMaticToStMatic(_amountInMatic *big.Int) (struct {
	AmountInStMatic    *big.Int
	TotalStMaticSupply *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	return _StMATIC.Contract.ConvertMaticToStMatic(&_StMATIC.CallOpts, _amountInMatic)
}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _amountInStMatic) view returns(uint256 amountInMatic, uint256 totalStMaticAmount, uint256 totalPooledMatic)
func (_StMATIC *StMATICCaller) ConvertStMaticToMatic(opts *bind.CallOpts, _amountInStMatic *big.Int) (struct {
	AmountInMatic      *big.Int
	TotalStMaticAmount *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "convertStMaticToMatic", _amountInStMatic)

	outstruct := new(struct {
		AmountInMatic      *big.Int
		TotalStMaticAmount *big.Int
		TotalPooledMatic   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountInMatic = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalStMaticAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalPooledMatic = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _amountInStMatic) view returns(uint256 amountInMatic, uint256 totalStMaticAmount, uint256 totalPooledMatic)
func (_StMATIC *StMATICSession) ConvertStMaticToMatic(_amountInStMatic *big.Int) (struct {
	AmountInMatic      *big.Int
	TotalStMaticAmount *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	return _StMATIC.Contract.ConvertStMaticToMatic(&_StMATIC.CallOpts, _amountInStMatic)
}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _amountInStMatic) view returns(uint256 amountInMatic, uint256 totalStMaticAmount, uint256 totalPooledMatic)
func (_StMATIC *StMATICCallerSession) ConvertStMaticToMatic(_amountInStMatic *big.Int) (struct {
	AmountInMatic      *big.Int
	TotalStMaticAmount *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	return _StMATIC.Contract.ConvertStMaticToMatic(&_StMATIC.CallOpts, _amountInStMatic)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_StMATIC *StMATICCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_StMATIC *StMATICSession) Dao() (common.Address, error) {
	return _StMATIC.Contract.Dao(&_StMATIC.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_StMATIC *StMATICCallerSession) Dao() (common.Address, error) {
	return _StMATIC.Contract.Dao(&_StMATIC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StMATIC *StMATICCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StMATIC *StMATICSession) Decimals() (uint8, error) {
	return _StMATIC.Contract.Decimals(&_StMATIC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StMATIC *StMATICCallerSession) Decimals() (uint8, error) {
	return _StMATIC.Contract.Decimals(&_StMATIC.CallOpts)
}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_StMATIC *StMATICCaller) DelegationLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "delegationLowerBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_StMATIC *StMATICSession) DelegationLowerBound() (*big.Int, error) {
	return _StMATIC.Contract.DelegationLowerBound(&_StMATIC.CallOpts)
}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_StMATIC *StMATICCallerSession) DelegationLowerBound() (*big.Int, error) {
	return _StMATIC.Contract.DelegationLowerBound(&_StMATIC.CallOpts)
}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_StMATIC *StMATICCaller) EntityFees(opts *bind.CallOpts) (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "entityFees")

	outstruct := new(struct {
		Dao       uint8
		Operators uint8
		Insurance uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dao = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Operators = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Insurance = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_StMATIC *StMATICSession) EntityFees() (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	return _StMATIC.Contract.EntityFees(&_StMATIC.CallOpts)
}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_StMATIC *StMATICCallerSession) EntityFees() (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	return _StMATIC.Contract.EntityFees(&_StMATIC.CallOpts)
}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_StMATIC *StMATICCaller) FxStateRootTunnel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "fxStateRootTunnel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_StMATIC *StMATICSession) FxStateRootTunnel() (common.Address, error) {
	return _StMATIC.Contract.FxStateRootTunnel(&_StMATIC.CallOpts)
}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_StMATIC *StMATICCallerSession) FxStateRootTunnel() (common.Address, error) {
	return _StMATIC.Contract.FxStateRootTunnel(&_StMATIC.CallOpts)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_StMATIC *StMATICCaller) GetLiquidRewards(opts *bind.CallOpts, _validatorShare common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "getLiquidRewards", _validatorShare)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_StMATIC *StMATICSession) GetLiquidRewards(_validatorShare common.Address) (*big.Int, error) {
	return _StMATIC.Contract.GetLiquidRewards(&_StMATIC.CallOpts, _validatorShare)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_StMATIC *StMATICCallerSession) GetLiquidRewards(_validatorShare common.Address) (*big.Int, error) {
	return _StMATIC.Contract.GetLiquidRewards(&_StMATIC.CallOpts, _validatorShare)
}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_StMATIC *StMATICCaller) GetMaticFromTokenId(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "getMaticFromTokenId", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_StMATIC *StMATICSession) GetMaticFromTokenId(_tokenId *big.Int) (*big.Int, error) {
	return _StMATIC.Contract.GetMaticFromTokenId(&_StMATIC.CallOpts, _tokenId)
}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_StMATIC *StMATICCallerSession) GetMaticFromTokenId(_tokenId *big.Int) (*big.Int, error) {
	return _StMATIC.Contract.GetMaticFromTokenId(&_StMATIC.CallOpts, _tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StMATIC *StMATICCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StMATIC *StMATICSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StMATIC.Contract.GetRoleAdmin(&_StMATIC.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StMATIC *StMATICCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StMATIC.Contract.GetRoleAdmin(&_StMATIC.CallOpts, role)
}

// GetToken2WithdrawRequests is a free data retrieval call binding the contract method 0x253d1735.
//
// Solidity: function getToken2WithdrawRequests(uint256 _tokenId) view returns((uint256,uint256,uint256,address)[])
func (_StMATIC *StMATICCaller) GetToken2WithdrawRequests(opts *bind.CallOpts, _tokenId *big.Int) ([]stMATICRequestWithdraw, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "getToken2WithdrawRequests", _tokenId)

	if err != nil {
		return *new([]stMATICRequestWithdraw), err
	}

	out0 := *abi.ConvertType(out[0], new([]stMATICRequestWithdraw)).(*[]stMATICRequestWithdraw)

	return out0, err

}

// GetToken2WithdrawRequests is a free data retrieval call binding the contract method 0x253d1735.
//
// Solidity: function getToken2WithdrawRequests(uint256 _tokenId) view returns((uint256,uint256,uint256,address)[])
func (_StMATIC *StMATICSession) GetToken2WithdrawRequests(_tokenId *big.Int) ([]stMATICRequestWithdraw, error) {
	return _StMATIC.Contract.GetToken2WithdrawRequests(&_StMATIC.CallOpts, _tokenId)
}

// GetToken2WithdrawRequests is a free data retrieval call binding the contract method 0x253d1735.
//
// Solidity: function getToken2WithdrawRequests(uint256 _tokenId) view returns((uint256,uint256,uint256,address)[])
func (_StMATIC *StMATICCallerSession) GetToken2WithdrawRequests(_tokenId *big.Int) ([]stMATICRequestWithdraw, error) {
	return _StMATIC.Contract.GetToken2WithdrawRequests(&_StMATIC.CallOpts, _tokenId)
}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_StMATIC *StMATICCaller) GetTotalPooledMatic(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "getTotalPooledMatic")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_StMATIC *StMATICSession) GetTotalPooledMatic() (*big.Int, error) {
	return _StMATIC.Contract.GetTotalPooledMatic(&_StMATIC.CallOpts)
}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_StMATIC *StMATICCallerSession) GetTotalPooledMatic() (*big.Int, error) {
	return _StMATIC.Contract.GetTotalPooledMatic(&_StMATIC.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_StMATIC *StMATICCaller) GetTotalStake(opts *bind.CallOpts, _validatorShare common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "getTotalStake", _validatorShare)

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
func (_StMATIC *StMATICSession) GetTotalStake(_validatorShare common.Address) (*big.Int, *big.Int, error) {
	return _StMATIC.Contract.GetTotalStake(&_StMATIC.CallOpts, _validatorShare)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_StMATIC *StMATICCallerSession) GetTotalStake(_validatorShare common.Address) (*big.Int, *big.Int, error) {
	return _StMATIC.Contract.GetTotalStake(&_StMATIC.CallOpts, _validatorShare)
}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_StMATIC *StMATICCaller) GetTotalStakeAcrossAllValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "getTotalStakeAcrossAllValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_StMATIC *StMATICSession) GetTotalStakeAcrossAllValidators() (*big.Int, error) {
	return _StMATIC.Contract.GetTotalStakeAcrossAllValidators(&_StMATIC.CallOpts)
}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_StMATIC *StMATICCallerSession) GetTotalStakeAcrossAllValidators() (*big.Int, error) {
	return _StMATIC.Contract.GetTotalStakeAcrossAllValidators(&_StMATIC.CallOpts)
}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_StMATIC *StMATICCaller) GetTotalWithdrawRequest(opts *bind.CallOpts) ([]stMATICRequestWithdraw, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "getTotalWithdrawRequest")

	if err != nil {
		return *new([]stMATICRequestWithdraw), err
	}

	out0 := *abi.ConvertType(out[0], new([]stMATICRequestWithdraw)).(*[]stMATICRequestWithdraw)

	return out0, err

}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_StMATIC *StMATICSession) GetTotalWithdrawRequest() ([]stMATICRequestWithdraw, error) {
	return _StMATIC.Contract.GetTotalWithdrawRequest(&_StMATIC.CallOpts)
}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_StMATIC *StMATICCallerSession) GetTotalWithdrawRequest() ([]stMATICRequestWithdraw, error) {
	return _StMATIC.Contract.GetTotalWithdrawRequest(&_StMATIC.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StMATIC *StMATICCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StMATIC *StMATICSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StMATIC.Contract.HasRole(&_StMATIC.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StMATIC *StMATICCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StMATIC.Contract.HasRole(&_StMATIC.CallOpts, role, account)
}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_StMATIC *StMATICCaller) Insurance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "insurance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_StMATIC *StMATICSession) Insurance() (common.Address, error) {
	return _StMATIC.Contract.Insurance(&_StMATIC.CallOpts)
}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_StMATIC *StMATICCallerSession) Insurance() (common.Address, error) {
	return _StMATIC.Contract.Insurance(&_StMATIC.CallOpts)
}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_StMATIC *StMATICCaller) LastWithdrawnValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "lastWithdrawnValidatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_StMATIC *StMATICSession) LastWithdrawnValidatorId() (*big.Int, error) {
	return _StMATIC.Contract.LastWithdrawnValidatorId(&_StMATIC.CallOpts)
}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_StMATIC *StMATICCallerSession) LastWithdrawnValidatorId() (*big.Int, error) {
	return _StMATIC.Contract.LastWithdrawnValidatorId(&_StMATIC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StMATIC *StMATICCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StMATIC *StMATICSession) Name() (string, error) {
	return _StMATIC.Contract.Name(&_StMATIC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StMATIC *StMATICCallerSession) Name() (string, error) {
	return _StMATIC.Contract.Name(&_StMATIC.CallOpts)
}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_StMATIC *StMATICCaller) NodeOperatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "nodeOperatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_StMATIC *StMATICSession) NodeOperatorRegistry() (common.Address, error) {
	return _StMATIC.Contract.NodeOperatorRegistry(&_StMATIC.CallOpts)
}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_StMATIC *StMATICCallerSession) NodeOperatorRegistry() (common.Address, error) {
	return _StMATIC.Contract.NodeOperatorRegistry(&_StMATIC.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StMATIC *StMATICCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StMATIC *StMATICSession) Paused() (bool, error) {
	return _StMATIC.Contract.Paused(&_StMATIC.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StMATIC *StMATICCallerSession) Paused() (bool, error) {
	return _StMATIC.Contract.Paused(&_StMATIC.CallOpts)
}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_StMATIC *StMATICCaller) PoLidoNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "poLidoNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_StMATIC *StMATICSession) PoLidoNFT() (common.Address, error) {
	return _StMATIC.Contract.PoLidoNFT(&_StMATIC.CallOpts)
}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_StMATIC *StMATICCallerSession) PoLidoNFT() (common.Address, error) {
	return _StMATIC.Contract.PoLidoNFT(&_StMATIC.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint8)
func (_StMATIC *StMATICCaller) ProtocolFee(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint8)
func (_StMATIC *StMATICSession) ProtocolFee() (uint8, error) {
	return _StMATIC.Contract.ProtocolFee(&_StMATIC.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint8)
func (_StMATIC *StMATICCallerSession) ProtocolFee() (uint8, error) {
	return _StMATIC.Contract.ProtocolFee(&_StMATIC.CallOpts)
}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_StMATIC *StMATICCaller) ReservedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "reservedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_StMATIC *StMATICSession) ReservedFunds() (*big.Int, error) {
	return _StMATIC.Contract.ReservedFunds(&_StMATIC.CallOpts)
}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_StMATIC *StMATICCallerSession) ReservedFunds() (*big.Int, error) {
	return _StMATIC.Contract.ReservedFunds(&_StMATIC.CallOpts)
}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_StMATIC *StMATICCaller) RewardDistributionLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "rewardDistributionLowerBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_StMATIC *StMATICSession) RewardDistributionLowerBound() (*big.Int, error) {
	return _StMATIC.Contract.RewardDistributionLowerBound(&_StMATIC.CallOpts)
}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_StMATIC *StMATICCallerSession) RewardDistributionLowerBound() (*big.Int, error) {
	return _StMATIC.Contract.RewardDistributionLowerBound(&_StMATIC.CallOpts)
}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StMATIC *StMATICCaller) StMaticWithdrawRequest(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "stMaticWithdrawRequest", arg0)

	outstruct := new(struct {
		Amount2WithdrawFromStMATIC *big.Int
		ValidatorNonce             *big.Int
		RequestEpoch               *big.Int
		ValidatorAddress           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount2WithdrawFromStMATIC = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidatorNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StMATIC *StMATICSession) StMaticWithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StMATIC.Contract.StMaticWithdrawRequest(&_StMATIC.CallOpts, arg0)
}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StMATIC *StMATICCallerSession) StMaticWithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StMATIC.Contract.StMaticWithdrawRequest(&_StMATIC.CallOpts, arg0)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_StMATIC *StMATICCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_StMATIC *StMATICSession) StakeManager() (common.Address, error) {
	return _StMATIC.Contract.StakeManager(&_StMATIC.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_StMATIC *StMATICCallerSession) StakeManager() (common.Address, error) {
	return _StMATIC.Contract.StakeManager(&_StMATIC.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_StMATIC *StMATICCaller) SubmitHandler(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "submitHandler")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_StMATIC *StMATICSession) SubmitHandler() (bool, error) {
	return _StMATIC.Contract.SubmitHandler(&_StMATIC.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_StMATIC *StMATICCallerSession) SubmitHandler() (bool, error) {
	return _StMATIC.Contract.SubmitHandler(&_StMATIC.CallOpts)
}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_StMATIC *StMATICCaller) SubmitThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "submitThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_StMATIC *StMATICSession) SubmitThreshold() (*big.Int, error) {
	return _StMATIC.Contract.SubmitThreshold(&_StMATIC.CallOpts)
}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_StMATIC *StMATICCallerSession) SubmitThreshold() (*big.Int, error) {
	return _StMATIC.Contract.SubmitThreshold(&_StMATIC.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StMATIC *StMATICCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StMATIC *StMATICSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StMATIC.Contract.SupportsInterface(&_StMATIC.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StMATIC *StMATICCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StMATIC.Contract.SupportsInterface(&_StMATIC.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StMATIC *StMATICCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StMATIC *StMATICSession) Symbol() (string, error) {
	return _StMATIC.Contract.Symbol(&_StMATIC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StMATIC *StMATICCallerSession) Symbol() (string, error) {
	return _StMATIC.Contract.Symbol(&_StMATIC.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StMATIC *StMATICCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StMATIC *StMATICSession) Token() (common.Address, error) {
	return _StMATIC.Contract.Token(&_StMATIC.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StMATIC *StMATICCallerSession) Token() (common.Address, error) {
	return _StMATIC.Contract.Token(&_StMATIC.CallOpts)
}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StMATIC *StMATICCaller) Token2WithdrawRequest(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "token2WithdrawRequest", arg0)

	outstruct := new(struct {
		Amount2WithdrawFromStMATIC *big.Int
		ValidatorNonce             *big.Int
		RequestEpoch               *big.Int
		ValidatorAddress           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount2WithdrawFromStMATIC = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidatorNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StMATIC *StMATICSession) Token2WithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StMATIC.Contract.Token2WithdrawRequest(&_StMATIC.CallOpts, arg0)
}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StMATIC *StMATICCallerSession) Token2WithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StMATIC.Contract.Token2WithdrawRequest(&_StMATIC.CallOpts, arg0)
}

// Token2WithdrawRequests is a free data retrieval call binding the contract method 0xc697d2c7.
//
// Solidity: function token2WithdrawRequests(uint256 , uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StMATIC *StMATICCaller) Token2WithdrawRequests(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "token2WithdrawRequests", arg0, arg1)

	outstruct := new(struct {
		Amount2WithdrawFromStMATIC *big.Int
		ValidatorNonce             *big.Int
		RequestEpoch               *big.Int
		ValidatorAddress           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount2WithdrawFromStMATIC = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidatorNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Token2WithdrawRequests is a free data retrieval call binding the contract method 0xc697d2c7.
//
// Solidity: function token2WithdrawRequests(uint256 , uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StMATIC *StMATICSession) Token2WithdrawRequests(arg0 *big.Int, arg1 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StMATIC.Contract.Token2WithdrawRequests(&_StMATIC.CallOpts, arg0, arg1)
}

// Token2WithdrawRequests is a free data retrieval call binding the contract method 0xc697d2c7.
//
// Solidity: function token2WithdrawRequests(uint256 , uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StMATIC *StMATICCallerSession) Token2WithdrawRequests(arg0 *big.Int, arg1 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StMATIC.Contract.Token2WithdrawRequests(&_StMATIC.CallOpts, arg0, arg1)
}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_StMATIC *StMATICCaller) TotalBuffered(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "totalBuffered")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_StMATIC *StMATICSession) TotalBuffered() (*big.Int, error) {
	return _StMATIC.Contract.TotalBuffered(&_StMATIC.CallOpts)
}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_StMATIC *StMATICCallerSession) TotalBuffered() (*big.Int, error) {
	return _StMATIC.Contract.TotalBuffered(&_StMATIC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StMATIC *StMATICCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StMATIC *StMATICSession) TotalSupply() (*big.Int, error) {
	return _StMATIC.Contract.TotalSupply(&_StMATIC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StMATIC *StMATICCallerSession) TotalSupply() (*big.Int, error) {
	return _StMATIC.Contract.TotalSupply(&_StMATIC.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StMATIC *StMATICCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StMATIC.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StMATIC *StMATICSession) Version() (string, error) {
	return _StMATIC.Contract.Version(&_StMATIC.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StMATIC *StMATICCallerSession) Version() (string, error) {
	return _StMATIC.Contract.Version(&_StMATIC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StMATIC *StMATICTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StMATIC *StMATICSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.Approve(&_StMATIC.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StMATIC *StMATICTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.Approve(&_StMATIC.TransactOpts, spender, amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_StMATIC *StMATICTransactor) ClaimTokens(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "claimTokens", _tokenId)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_StMATIC *StMATICSession) ClaimTokens(_tokenId *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.ClaimTokens(&_StMATIC.TransactOpts, _tokenId)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_StMATIC *StMATICTransactorSession) ClaimTokens(_tokenId *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.ClaimTokens(&_StMATIC.TransactOpts, _tokenId)
}

// ClaimTokensFromValidatorToContract is a paid mutator transaction binding the contract method 0x4cfeb862.
//
// Solidity: function claimTokensFromValidatorToContract(uint256 _index) returns()
func (_StMATIC *StMATICTransactor) ClaimTokensFromValidatorToContract(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "claimTokensFromValidatorToContract", _index)
}

// ClaimTokensFromValidatorToContract is a paid mutator transaction binding the contract method 0x4cfeb862.
//
// Solidity: function claimTokensFromValidatorToContract(uint256 _index) returns()
func (_StMATIC *StMATICSession) ClaimTokensFromValidatorToContract(_index *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.ClaimTokensFromValidatorToContract(&_StMATIC.TransactOpts, _index)
}

// ClaimTokensFromValidatorToContract is a paid mutator transaction binding the contract method 0x4cfeb862.
//
// Solidity: function claimTokensFromValidatorToContract(uint256 _index) returns()
func (_StMATIC *StMATICTransactorSession) ClaimTokensFromValidatorToContract(_index *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.ClaimTokensFromValidatorToContract(&_StMATIC.TransactOpts, _index)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StMATIC *StMATICTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StMATIC *StMATICSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.DecreaseAllowance(&_StMATIC.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StMATIC *StMATICTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.DecreaseAllowance(&_StMATIC.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_StMATIC *StMATICTransactor) Delegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "delegate")
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_StMATIC *StMATICSession) Delegate() (*types.Transaction, error) {
	return _StMATIC.Contract.Delegate(&_StMATIC.TransactOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_StMATIC *StMATICTransactorSession) Delegate() (*types.Transaction, error) {
	return _StMATIC.Contract.Delegate(&_StMATIC.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_StMATIC *StMATICTransactor) DistributeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "distributeRewards")
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_StMATIC *StMATICSession) DistributeRewards() (*types.Transaction, error) {
	return _StMATIC.Contract.DistributeRewards(&_StMATIC.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_StMATIC *StMATICTransactorSession) DistributeRewards() (*types.Transaction, error) {
	return _StMATIC.Contract.DistributeRewards(&_StMATIC.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StMATIC *StMATICTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StMATIC *StMATICSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.GrantRole(&_StMATIC.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StMATIC *StMATICTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.GrantRole(&_StMATIC.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StMATIC *StMATICTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StMATIC *StMATICSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.IncreaseAllowance(&_StMATIC.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StMATIC *StMATICTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.IncreaseAllowance(&_StMATIC.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel) returns()
func (_StMATIC *StMATICTransactor) Initialize(opts *bind.TransactOpts, _nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "initialize", _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel) returns()
func (_StMATIC *StMATICSession) Initialize(_nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.Initialize(&_StMATIC.TransactOpts, _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel) returns()
func (_StMATIC *StMATICTransactorSession) Initialize(_nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.Initialize(&_StMATIC.TransactOpts, _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StMATIC *StMATICTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StMATIC *StMATICSession) Pause() (*types.Transaction, error) {
	return _StMATIC.Contract.Pause(&_StMATIC.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StMATIC *StMATICTransactorSession) Pause() (*types.Transaction, error) {
	return _StMATIC.Contract.Pause(&_StMATIC.TransactOpts)
}

// RebalanceDelegatedTokens is a paid mutator transaction binding the contract method 0xd280f14f.
//
// Solidity: function rebalanceDelegatedTokens() returns()
func (_StMATIC *StMATICTransactor) RebalanceDelegatedTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "rebalanceDelegatedTokens")
}

// RebalanceDelegatedTokens is a paid mutator transaction binding the contract method 0xd280f14f.
//
// Solidity: function rebalanceDelegatedTokens() returns()
func (_StMATIC *StMATICSession) RebalanceDelegatedTokens() (*types.Transaction, error) {
	return _StMATIC.Contract.RebalanceDelegatedTokens(&_StMATIC.TransactOpts)
}

// RebalanceDelegatedTokens is a paid mutator transaction binding the contract method 0xd280f14f.
//
// Solidity: function rebalanceDelegatedTokens() returns()
func (_StMATIC *StMATICTransactorSession) RebalanceDelegatedTokens() (*types.Transaction, error) {
	return _StMATIC.Contract.RebalanceDelegatedTokens(&_StMATIC.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StMATIC *StMATICTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StMATIC *StMATICSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.RenounceRole(&_StMATIC.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StMATIC *StMATICTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.RenounceRole(&_StMATIC.TransactOpts, role, account)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 _amount, address _referral) returns(uint256)
func (_StMATIC *StMATICTransactor) RequestWithdraw(opts *bind.TransactOpts, _amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "requestWithdraw", _amount, _referral)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 _amount, address _referral) returns(uint256)
func (_StMATIC *StMATICSession) RequestWithdraw(_amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.RequestWithdraw(&_StMATIC.TransactOpts, _amount, _referral)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 _amount, address _referral) returns(uint256)
func (_StMATIC *StMATICTransactorSession) RequestWithdraw(_amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.RequestWithdraw(&_StMATIC.TransactOpts, _amount, _referral)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StMATIC *StMATICTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StMATIC *StMATICSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.RevokeRole(&_StMATIC.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StMATIC *StMATICTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.RevokeRole(&_StMATIC.TransactOpts, role, account)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _newDAO) returns()
func (_StMATIC *StMATICTransactor) SetDaoAddress(opts *bind.TransactOpts, _newDAO common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "setDaoAddress", _newDAO)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _newDAO) returns()
func (_StMATIC *StMATICSession) SetDaoAddress(_newDAO common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.SetDaoAddress(&_StMATIC.TransactOpts, _newDAO)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _newDAO) returns()
func (_StMATIC *StMATICTransactorSession) SetDaoAddress(_newDAO common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.SetDaoAddress(&_StMATIC.TransactOpts, _newDAO)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_StMATIC *StMATICTransactor) SetDelegationLowerBound(opts *bind.TransactOpts, _delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "setDelegationLowerBound", _delegationLowerBound)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_StMATIC *StMATICSession) SetDelegationLowerBound(_delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.SetDelegationLowerBound(&_StMATIC.TransactOpts, _delegationLowerBound)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_StMATIC *StMATICTransactorSession) SetDelegationLowerBound(_delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.SetDelegationLowerBound(&_StMATIC.TransactOpts, _delegationLowerBound)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_StMATIC *StMATICTransactor) SetFees(opts *bind.TransactOpts, _daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "setFees", _daoFee, _operatorsFee, _insuranceFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_StMATIC *StMATICSession) SetFees(_daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _StMATIC.Contract.SetFees(&_StMATIC.TransactOpts, _daoFee, _operatorsFee, _insuranceFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_StMATIC *StMATICTransactorSession) SetFees(_daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _StMATIC.Contract.SetFees(&_StMATIC.TransactOpts, _daoFee, _operatorsFee, _insuranceFee)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _newFxStateRootTunnel) returns()
func (_StMATIC *StMATICTransactor) SetFxStateRootTunnel(opts *bind.TransactOpts, _newFxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "setFxStateRootTunnel", _newFxStateRootTunnel)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _newFxStateRootTunnel) returns()
func (_StMATIC *StMATICSession) SetFxStateRootTunnel(_newFxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.SetFxStateRootTunnel(&_StMATIC.TransactOpts, _newFxStateRootTunnel)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _newFxStateRootTunnel) returns()
func (_StMATIC *StMATICTransactorSession) SetFxStateRootTunnel(_newFxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.SetFxStateRootTunnel(&_StMATIC.TransactOpts, _newFxStateRootTunnel)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_StMATIC *StMATICTransactor) SetInsuranceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "setInsuranceAddress", _address)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_StMATIC *StMATICSession) SetInsuranceAddress(_address common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.SetInsuranceAddress(&_StMATIC.TransactOpts, _address)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_StMATIC *StMATICTransactorSession) SetInsuranceAddress(_address common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.SetInsuranceAddress(&_StMATIC.TransactOpts, _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_StMATIC *StMATICTransactor) SetNodeOperatorRegistryAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "setNodeOperatorRegistryAddress", _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_StMATIC *StMATICSession) SetNodeOperatorRegistryAddress(_address common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.SetNodeOperatorRegistryAddress(&_StMATIC.TransactOpts, _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_StMATIC *StMATICTransactorSession) SetNodeOperatorRegistryAddress(_address common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.SetNodeOperatorRegistryAddress(&_StMATIC.TransactOpts, _address)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _newLidoNFT) returns()
func (_StMATIC *StMATICTransactor) SetPoLidoNFT(opts *bind.TransactOpts, _newLidoNFT common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "setPoLidoNFT", _newLidoNFT)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _newLidoNFT) returns()
func (_StMATIC *StMATICSession) SetPoLidoNFT(_newLidoNFT common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.SetPoLidoNFT(&_StMATIC.TransactOpts, _newLidoNFT)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _newLidoNFT) returns()
func (_StMATIC *StMATICTransactorSession) SetPoLidoNFT(_newLidoNFT common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.SetPoLidoNFT(&_StMATIC.TransactOpts, _newLidoNFT)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x4e91f811.
//
// Solidity: function setProtocolFee(uint8 _newProtocolFee) returns()
func (_StMATIC *StMATICTransactor) SetProtocolFee(opts *bind.TransactOpts, _newProtocolFee uint8) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "setProtocolFee", _newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x4e91f811.
//
// Solidity: function setProtocolFee(uint8 _newProtocolFee) returns()
func (_StMATIC *StMATICSession) SetProtocolFee(_newProtocolFee uint8) (*types.Transaction, error) {
	return _StMATIC.Contract.SetProtocolFee(&_StMATIC.TransactOpts, _newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x4e91f811.
//
// Solidity: function setProtocolFee(uint8 _newProtocolFee) returns()
func (_StMATIC *StMATICTransactorSession) SetProtocolFee(_newProtocolFee uint8) (*types.Transaction, error) {
	return _StMATIC.Contract.SetProtocolFee(&_StMATIC.TransactOpts, _newProtocolFee)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _newRewardDistributionLowerBound) returns()
func (_StMATIC *StMATICTransactor) SetRewardDistributionLowerBound(opts *bind.TransactOpts, _newRewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "setRewardDistributionLowerBound", _newRewardDistributionLowerBound)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _newRewardDistributionLowerBound) returns()
func (_StMATIC *StMATICSession) SetRewardDistributionLowerBound(_newRewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.SetRewardDistributionLowerBound(&_StMATIC.TransactOpts, _newRewardDistributionLowerBound)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _newRewardDistributionLowerBound) returns()
func (_StMATIC *StMATICTransactorSession) SetRewardDistributionLowerBound(_newRewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.SetRewardDistributionLowerBound(&_StMATIC.TransactOpts, _newRewardDistributionLowerBound)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _newVersion) returns()
func (_StMATIC *StMATICTransactor) SetVersion(opts *bind.TransactOpts, _newVersion string) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "setVersion", _newVersion)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _newVersion) returns()
func (_StMATIC *StMATICSession) SetVersion(_newVersion string) (*types.Transaction, error) {
	return _StMATIC.Contract.SetVersion(&_StMATIC.TransactOpts, _newVersion)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _newVersion) returns()
func (_StMATIC *StMATICTransactorSession) SetVersion(_newVersion string) (*types.Transaction, error) {
	return _StMATIC.Contract.SetVersion(&_StMATIC.TransactOpts, _newVersion)
}

// Submit is a paid mutator transaction binding the contract method 0xf532e86a.
//
// Solidity: function submit(uint256 _amount, address _referral) returns(uint256)
func (_StMATIC *StMATICTransactor) Submit(opts *bind.TransactOpts, _amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "submit", _amount, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xf532e86a.
//
// Solidity: function submit(uint256 _amount, address _referral) returns(uint256)
func (_StMATIC *StMATICSession) Submit(_amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.Submit(&_StMATIC.TransactOpts, _amount, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xf532e86a.
//
// Solidity: function submit(uint256 _amount, address _referral) returns(uint256)
func (_StMATIC *StMATICTransactorSession) Submit(_amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.Submit(&_StMATIC.TransactOpts, _amount, _referral)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StMATIC *StMATICTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StMATIC *StMATICSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.Transfer(&_StMATIC.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StMATIC *StMATICTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.Transfer(&_StMATIC.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StMATIC *StMATICTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StMATIC *StMATICSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.TransferFrom(&_StMATIC.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StMATIC *StMATICTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StMATIC.Contract.TransferFrom(&_StMATIC.TransactOpts, from, to, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StMATIC *StMATICTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StMATIC *StMATICSession) Unpause() (*types.Transaction, error) {
	return _StMATIC.Contract.Unpause(&_StMATIC.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StMATIC *StMATICTransactorSession) Unpause() (*types.Transaction, error) {
	return _StMATIC.Contract.Unpause(&_StMATIC.TransactOpts)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_StMATIC *StMATICTransactor) WithdrawTotalDelegated(opts *bind.TransactOpts, _validatorShare common.Address) (*types.Transaction, error) {
	return _StMATIC.contract.Transact(opts, "withdrawTotalDelegated", _validatorShare)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_StMATIC *StMATICSession) WithdrawTotalDelegated(_validatorShare common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.WithdrawTotalDelegated(&_StMATIC.TransactOpts, _validatorShare)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_StMATIC *StMATICTransactorSession) WithdrawTotalDelegated(_validatorShare common.Address) (*types.Transaction, error) {
	return _StMATIC.Contract.WithdrawTotalDelegated(&_StMATIC.TransactOpts, _validatorShare)
}
