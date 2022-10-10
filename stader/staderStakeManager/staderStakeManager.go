// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staderStakeManager

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

// StaderStakeManagerMetaData contains all meta data concerning the StaderStakeManager contract.
var StaderStakeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CHECKPOINT_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NFTContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NFTCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRootChain\",\"type\":\"address\"}],\"name\":\"changeRootChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkPointBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockInterval\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"voteHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256[3][]\",\"name\":\"sigs\",\"type\":\"uint256[3][]\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkpointRewardDelta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accumFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"heimdallFee\",\"type\":\"uint256\"}],\"name\":\"confirmAuctionBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentValidatorSetSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentValidatorSetTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseValidatorDelegatedAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"delegatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegationDeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"delegatorsReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"auctionUser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"heimdallFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auctionAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"dethroneAndStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"drain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"drainValidatorShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dynasty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eventsHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extensionCode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"forceUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"getValidatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rootchain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_NFTContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingLogger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorShareFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_extensionCode\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"}],\"name\":\"insertSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"latestSignerUpdateEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxRewardedCheckpoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromValidatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"migrateDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorIdFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorIdTo\",\"type\":\"uint256\"}],\"name\":\"migrateValidatorsData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minHeimdallFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposerBonus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_NFTContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingLogger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorShareFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_extensionCode\",\"type\":\"address\"}],\"name\":\"reinitialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"replacementCoolDown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stakeRewards\",\"type\":\"bool\"}],\"name\":\"restake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardDecreasePerCheckpoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootChain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentEpoch\",\"type\":\"uint256\"}],\"name\":\"setCurrentEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelegationEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"setStakingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerToValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerUpdateLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_slashingInfoList\",\"type\":\"bytes\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"heimdallFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"stakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_acceptDelegation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_signerPubkey\",\"type\":\"bytes\"}],\"name\":\"startAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"forNCheckpoints\",\"type\":\"uint256\"}],\"name\":\"stopAuctions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"heimdallFee\",\"type\":\"uint256\"}],\"name\":\"topUpForFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalHeimdallFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRewardsLiquidated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"totalStakedFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"transferFunds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"unjail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"unstakeClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blocks\",\"type\":\"uint256\"}],\"name\":\"updateCheckPointBlockInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"updateCheckpointReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardDecreasePerCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxRewardedCheckpoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpointRewardDelta\",\"type\":\"uint256\"}],\"name\":\"updateCheckpointRewardParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newCommissionRate\",\"type\":\"uint256\"}],\"name\":\"updateCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDynasty\",\"type\":\"uint256\"}],\"name\":\"updateDynastyValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minHeimdallFee\",\"type\":\"uint256\"}],\"name\":\"updateMinAmounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposerBonus\",\"type\":\"uint256\"}],\"name\":\"updateProposerBonus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"updateSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"updateSignerUpdateLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newContractAddress\",\"type\":\"address\"}],\"name\":\"updateValidatorContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"delegation\",\"type\":\"bool\"}],\"name\":\"updateValidatorDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"updateValidatorState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"updateValidatorThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userFeeExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"validatorReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorShareFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"validatorStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorStateChanges\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"stakerCount\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activationEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivationEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastCommissionUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialRewardPerStake\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"withdrawDelegatorsReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StaderStakeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StaderStakeManagerMetaData.ABI instead.
var StaderStakeManagerABI = StaderStakeManagerMetaData.ABI

// StaderStakeManager is an auto generated Go binding around an Ethereum contract.
type StaderStakeManager struct {
	StaderStakeManagerCaller     // Read-only binding to the contract
	StaderStakeManagerTransactor // Write-only binding to the contract
	StaderStakeManagerFilterer   // Log filterer for contract events
}

// StaderStakeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaderStakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderStakeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaderStakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderStakeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaderStakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderStakeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaderStakeManagerSession struct {
	Contract     *StaderStakeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StaderStakeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaderStakeManagerCallerSession struct {
	Contract *StaderStakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// StaderStakeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaderStakeManagerTransactorSession struct {
	Contract     *StaderStakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// StaderStakeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaderStakeManagerRaw struct {
	Contract *StaderStakeManager // Generic contract binding to access the raw methods on
}

// StaderStakeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaderStakeManagerCallerRaw struct {
	Contract *StaderStakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StaderStakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaderStakeManagerTransactorRaw struct {
	Contract *StaderStakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaderStakeManager creates a new instance of StaderStakeManager, bound to a specific deployed contract.
func NewStaderStakeManager(address common.Address, backend bind.ContractBackend) (*StaderStakeManager, error) {
	contract, err := bindStaderStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaderStakeManager{StaderStakeManagerCaller: StaderStakeManagerCaller{contract: contract}, StaderStakeManagerTransactor: StaderStakeManagerTransactor{contract: contract}, StaderStakeManagerFilterer: StaderStakeManagerFilterer{contract: contract}}, nil
}

// NewStaderStakeManagerCaller creates a new read-only instance of StaderStakeManager, bound to a specific deployed contract.
func NewStaderStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*StaderStakeManagerCaller, error) {
	contract, err := bindStaderStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaderStakeManagerCaller{contract: contract}, nil
}

// NewStaderStakeManagerTransactor creates a new write-only instance of StaderStakeManager, bound to a specific deployed contract.
func NewStaderStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StaderStakeManagerTransactor, error) {
	contract, err := bindStaderStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaderStakeManagerTransactor{contract: contract}, nil
}

// NewStaderStakeManagerFilterer creates a new log filterer instance of StaderStakeManager, bound to a specific deployed contract.
func NewStaderStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StaderStakeManagerFilterer, error) {
	contract, err := bindStaderStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaderStakeManagerFilterer{contract: contract}, nil
}

// bindStaderStakeManager binds a generic wrapper to an already deployed contract.
func bindStaderStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StaderStakeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderStakeManager *StaderStakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderStakeManager.Contract.StaderStakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderStakeManager *StaderStakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.StaderStakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderStakeManager *StaderStakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.StaderStakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderStakeManager *StaderStakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderStakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderStakeManager *StaderStakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderStakeManager *StaderStakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.contract.Transact(opts, method, params...)
}

// CHECKPOINTREWARD is a free data retrieval call binding the contract method 0x7d669752.
//
// Solidity: function CHECKPOINT_REWARD() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) CHECKPOINTREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "CHECKPOINT_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHECKPOINTREWARD is a free data retrieval call binding the contract method 0x7d669752.
//
// Solidity: function CHECKPOINT_REWARD() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) CHECKPOINTREWARD() (*big.Int, error) {
	return _StaderStakeManager.Contract.CHECKPOINTREWARD(&_StaderStakeManager.CallOpts)
}

// CHECKPOINTREWARD is a free data retrieval call binding the contract method 0x7d669752.
//
// Solidity: function CHECKPOINT_REWARD() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) CHECKPOINTREWARD() (*big.Int, error) {
	return _StaderStakeManager.Contract.CHECKPOINTREWARD(&_StaderStakeManager.CallOpts)
}

// NFTContract is a free data retrieval call binding the contract method 0x31c2273b.
//
// Solidity: function NFTContract() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) NFTContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "NFTContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NFTContract is a free data retrieval call binding the contract method 0x31c2273b.
//
// Solidity: function NFTContract() view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) NFTContract() (common.Address, error) {
	return _StaderStakeManager.Contract.NFTContract(&_StaderStakeManager.CallOpts)
}

// NFTContract is a free data retrieval call binding the contract method 0x31c2273b.
//
// Solidity: function NFTContract() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) NFTContract() (common.Address, error) {
	return _StaderStakeManager.Contract.NFTContract(&_StaderStakeManager.CallOpts)
}

// NFTCounter is a free data retrieval call binding the contract method 0x5508d8e1.
//
// Solidity: function NFTCounter() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) NFTCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "NFTCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NFTCounter is a free data retrieval call binding the contract method 0x5508d8e1.
//
// Solidity: function NFTCounter() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) NFTCounter() (*big.Int, error) {
	return _StaderStakeManager.Contract.NFTCounter(&_StaderStakeManager.CallOpts)
}

// NFTCounter is a free data retrieval call binding the contract method 0x5508d8e1.
//
// Solidity: function NFTCounter() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) NFTCounter() (*big.Int, error) {
	return _StaderStakeManager.Contract.NFTCounter(&_StaderStakeManager.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) WITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "WITHDRAWAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _StaderStakeManager.Contract.WITHDRAWALDELAY(&_StaderStakeManager.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _StaderStakeManager.Contract.WITHDRAWALDELAY(&_StaderStakeManager.CallOpts)
}

// AccountStateRoot is a free data retrieval call binding the contract method 0x17c2b910.
//
// Solidity: function accountStateRoot() view returns(bytes32)
func (_StaderStakeManager *StaderStakeManagerCaller) AccountStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "accountStateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AccountStateRoot is a free data retrieval call binding the contract method 0x17c2b910.
//
// Solidity: function accountStateRoot() view returns(bytes32)
func (_StaderStakeManager *StaderStakeManagerSession) AccountStateRoot() ([32]byte, error) {
	return _StaderStakeManager.Contract.AccountStateRoot(&_StaderStakeManager.CallOpts)
}

// AccountStateRoot is a free data retrieval call binding the contract method 0x17c2b910.
//
// Solidity: function accountStateRoot() view returns(bytes32)
func (_StaderStakeManager *StaderStakeManagerCallerSession) AccountStateRoot() ([32]byte, error) {
	return _StaderStakeManager.Contract.AccountStateRoot(&_StaderStakeManager.CallOpts)
}

// AuctionPeriod is a free data retrieval call binding the contract method 0x0cccfc58.
//
// Solidity: function auctionPeriod() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) AuctionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "auctionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionPeriod is a free data retrieval call binding the contract method 0x0cccfc58.
//
// Solidity: function auctionPeriod() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) AuctionPeriod() (*big.Int, error) {
	return _StaderStakeManager.Contract.AuctionPeriod(&_StaderStakeManager.CallOpts)
}

// AuctionPeriod is a free data retrieval call binding the contract method 0x0cccfc58.
//
// Solidity: function auctionPeriod() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) AuctionPeriod() (*big.Int, error) {
	return _StaderStakeManager.Contract.AuctionPeriod(&_StaderStakeManager.CallOpts)
}

// CheckPointBlockInterval is a free data retrieval call binding the contract method 0x25316411.
//
// Solidity: function checkPointBlockInterval() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) CheckPointBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "checkPointBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckPointBlockInterval is a free data retrieval call binding the contract method 0x25316411.
//
// Solidity: function checkPointBlockInterval() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) CheckPointBlockInterval() (*big.Int, error) {
	return _StaderStakeManager.Contract.CheckPointBlockInterval(&_StaderStakeManager.CallOpts)
}

// CheckPointBlockInterval is a free data retrieval call binding the contract method 0x25316411.
//
// Solidity: function checkPointBlockInterval() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) CheckPointBlockInterval() (*big.Int, error) {
	return _StaderStakeManager.Contract.CheckPointBlockInterval(&_StaderStakeManager.CallOpts)
}

// CheckpointRewardDelta is a free data retrieval call binding the contract method 0x7c7eaf1a.
//
// Solidity: function checkpointRewardDelta() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) CheckpointRewardDelta(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "checkpointRewardDelta")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckpointRewardDelta is a free data retrieval call binding the contract method 0x7c7eaf1a.
//
// Solidity: function checkpointRewardDelta() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) CheckpointRewardDelta() (*big.Int, error) {
	return _StaderStakeManager.Contract.CheckpointRewardDelta(&_StaderStakeManager.CallOpts)
}

// CheckpointRewardDelta is a free data retrieval call binding the contract method 0x7c7eaf1a.
//
// Solidity: function checkpointRewardDelta() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) CheckpointRewardDelta() (*big.Int, error) {
	return _StaderStakeManager.Contract.CheckpointRewardDelta(&_StaderStakeManager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) CurrentEpoch() (*big.Int, error) {
	return _StaderStakeManager.Contract.CurrentEpoch(&_StaderStakeManager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) CurrentEpoch() (*big.Int, error) {
	return _StaderStakeManager.Contract.CurrentEpoch(&_StaderStakeManager.CallOpts)
}

// CurrentValidatorSetSize is a free data retrieval call binding the contract method 0x7f952d95.
//
// Solidity: function currentValidatorSetSize() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) CurrentValidatorSetSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "currentValidatorSetSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetSize is a free data retrieval call binding the contract method 0x7f952d95.
//
// Solidity: function currentValidatorSetSize() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) CurrentValidatorSetSize() (*big.Int, error) {
	return _StaderStakeManager.Contract.CurrentValidatorSetSize(&_StaderStakeManager.CallOpts)
}

// CurrentValidatorSetSize is a free data retrieval call binding the contract method 0x7f952d95.
//
// Solidity: function currentValidatorSetSize() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) CurrentValidatorSetSize() (*big.Int, error) {
	return _StaderStakeManager.Contract.CurrentValidatorSetSize(&_StaderStakeManager.CallOpts)
}

// CurrentValidatorSetTotalStake is a free data retrieval call binding the contract method 0xa4769071.
//
// Solidity: function currentValidatorSetTotalStake() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) CurrentValidatorSetTotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "currentValidatorSetTotalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetTotalStake is a free data retrieval call binding the contract method 0xa4769071.
//
// Solidity: function currentValidatorSetTotalStake() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) CurrentValidatorSetTotalStake() (*big.Int, error) {
	return _StaderStakeManager.Contract.CurrentValidatorSetTotalStake(&_StaderStakeManager.CallOpts)
}

// CurrentValidatorSetTotalStake is a free data retrieval call binding the contract method 0xa4769071.
//
// Solidity: function currentValidatorSetTotalStake() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) CurrentValidatorSetTotalStake() (*big.Int, error) {
	return _StaderStakeManager.Contract.CurrentValidatorSetTotalStake(&_StaderStakeManager.CallOpts)
}

// DelegatedAmount is a free data retrieval call binding the contract method 0x7f4b4323.
//
// Solidity: function delegatedAmount(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) DelegatedAmount(opts *bind.CallOpts, validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "delegatedAmount", validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedAmount is a free data retrieval call binding the contract method 0x7f4b4323.
//
// Solidity: function delegatedAmount(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) DelegatedAmount(validatorId *big.Int) (*big.Int, error) {
	return _StaderStakeManager.Contract.DelegatedAmount(&_StaderStakeManager.CallOpts, validatorId)
}

// DelegatedAmount is a free data retrieval call binding the contract method 0x7f4b4323.
//
// Solidity: function delegatedAmount(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) DelegatedAmount(validatorId *big.Int) (*big.Int, error) {
	return _StaderStakeManager.Contract.DelegatedAmount(&_StaderStakeManager.CallOpts, validatorId)
}

// DelegationEnabled is a free data retrieval call binding the contract method 0x54b8c601.
//
// Solidity: function delegationEnabled() view returns(bool)
func (_StaderStakeManager *StaderStakeManagerCaller) DelegationEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "delegationEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationEnabled is a free data retrieval call binding the contract method 0x54b8c601.
//
// Solidity: function delegationEnabled() view returns(bool)
func (_StaderStakeManager *StaderStakeManagerSession) DelegationEnabled() (bool, error) {
	return _StaderStakeManager.Contract.DelegationEnabled(&_StaderStakeManager.CallOpts)
}

// DelegationEnabled is a free data retrieval call binding the contract method 0x54b8c601.
//
// Solidity: function delegationEnabled() view returns(bool)
func (_StaderStakeManager *StaderStakeManagerCallerSession) DelegationEnabled() (bool, error) {
	return _StaderStakeManager.Contract.DelegationEnabled(&_StaderStakeManager.CallOpts)
}

// DelegatorsReward is a free data retrieval call binding the contract method 0x39610f78.
//
// Solidity: function delegatorsReward(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) DelegatorsReward(opts *bind.CallOpts, validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "delegatorsReward", validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatorsReward is a free data retrieval call binding the contract method 0x39610f78.
//
// Solidity: function delegatorsReward(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) DelegatorsReward(validatorId *big.Int) (*big.Int, error) {
	return _StaderStakeManager.Contract.DelegatorsReward(&_StaderStakeManager.CallOpts, validatorId)
}

// DelegatorsReward is a free data retrieval call binding the contract method 0x39610f78.
//
// Solidity: function delegatorsReward(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) DelegatorsReward(validatorId *big.Int) (*big.Int, error) {
	return _StaderStakeManager.Contract.DelegatorsReward(&_StaderStakeManager.CallOpts, validatorId)
}

// Dynasty is a free data retrieval call binding the contract method 0x7060054d.
//
// Solidity: function dynasty() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) Dynasty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "dynasty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dynasty is a free data retrieval call binding the contract method 0x7060054d.
//
// Solidity: function dynasty() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) Dynasty() (*big.Int, error) {
	return _StaderStakeManager.Contract.Dynasty(&_StaderStakeManager.CallOpts)
}

// Dynasty is a free data retrieval call binding the contract method 0x7060054d.
//
// Solidity: function dynasty() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) Dynasty() (*big.Int, error) {
	return _StaderStakeManager.Contract.Dynasty(&_StaderStakeManager.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) Epoch() (*big.Int, error) {
	return _StaderStakeManager.Contract.Epoch(&_StaderStakeManager.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) Epoch() (*big.Int, error) {
	return _StaderStakeManager.Contract.Epoch(&_StaderStakeManager.CallOpts)
}

// EventsHub is a free data retrieval call binding the contract method 0x883b455f.
//
// Solidity: function eventsHub() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) EventsHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "eventsHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EventsHub is a free data retrieval call binding the contract method 0x883b455f.
//
// Solidity: function eventsHub() view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) EventsHub() (common.Address, error) {
	return _StaderStakeManager.Contract.EventsHub(&_StaderStakeManager.CallOpts)
}

// EventsHub is a free data retrieval call binding the contract method 0x883b455f.
//
// Solidity: function eventsHub() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) EventsHub() (common.Address, error) {
	return _StaderStakeManager.Contract.EventsHub(&_StaderStakeManager.CallOpts)
}

// ExtensionCode is a free data retrieval call binding the contract method 0xf8a3176c.
//
// Solidity: function extensionCode() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) ExtensionCode(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "extensionCode")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtensionCode is a free data retrieval call binding the contract method 0xf8a3176c.
//
// Solidity: function extensionCode() view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) ExtensionCode() (common.Address, error) {
	return _StaderStakeManager.Contract.ExtensionCode(&_StaderStakeManager.CallOpts)
}

// ExtensionCode is a free data retrieval call binding the contract method 0xf8a3176c.
//
// Solidity: function extensionCode() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) ExtensionCode() (common.Address, error) {
	return _StaderStakeManager.Contract.ExtensionCode(&_StaderStakeManager.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) GetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "getRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) GetRegistry() (common.Address, error) {
	return _StaderStakeManager.Contract.GetRegistry(&_StaderStakeManager.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) GetRegistry() (common.Address, error) {
	return _StaderStakeManager.Contract.GetRegistry(&_StaderStakeManager.CallOpts)
}

// GetValidatorContract is a free data retrieval call binding the contract method 0x56342d8c.
//
// Solidity: function getValidatorContract(uint256 validatorId) view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) GetValidatorContract(opts *bind.CallOpts, validatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "getValidatorContract", validatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetValidatorContract is a free data retrieval call binding the contract method 0x56342d8c.
//
// Solidity: function getValidatorContract(uint256 validatorId) view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) GetValidatorContract(validatorId *big.Int) (common.Address, error) {
	return _StaderStakeManager.Contract.GetValidatorContract(&_StaderStakeManager.CallOpts, validatorId)
}

// GetValidatorContract is a free data retrieval call binding the contract method 0x56342d8c.
//
// Solidity: function getValidatorContract(uint256 validatorId) view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) GetValidatorContract(validatorId *big.Int) (common.Address, error) {
	return _StaderStakeManager.Contract.GetValidatorContract(&_StaderStakeManager.CallOpts, validatorId)
}

// GetValidatorId is a free data retrieval call binding the contract method 0x174e6832.
//
// Solidity: function getValidatorId(address user) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) GetValidatorId(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "getValidatorId", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorId is a free data retrieval call binding the contract method 0x174e6832.
//
// Solidity: function getValidatorId(address user) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) GetValidatorId(user common.Address) (*big.Int, error) {
	return _StaderStakeManager.Contract.GetValidatorId(&_StaderStakeManager.CallOpts, user)
}

// GetValidatorId is a free data retrieval call binding the contract method 0x174e6832.
//
// Solidity: function getValidatorId(address user) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) GetValidatorId(user common.Address) (*big.Int, error) {
	return _StaderStakeManager.Contract.GetValidatorId(&_StaderStakeManager.CallOpts, user)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) Governance() (common.Address, error) {
	return _StaderStakeManager.Contract.Governance(&_StaderStakeManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) Governance() (common.Address, error) {
	return _StaderStakeManager.Contract.Governance(&_StaderStakeManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StaderStakeManager *StaderStakeManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StaderStakeManager *StaderStakeManagerSession) IsOwner() (bool, error) {
	return _StaderStakeManager.Contract.IsOwner(&_StaderStakeManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StaderStakeManager *StaderStakeManagerCallerSession) IsOwner() (bool, error) {
	return _StaderStakeManager.Contract.IsOwner(&_StaderStakeManager.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0x2649263a.
//
// Solidity: function isValidator(uint256 validatorId) view returns(bool)
func (_StaderStakeManager *StaderStakeManagerCaller) IsValidator(opts *bind.CallOpts, validatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "isValidator", validatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0x2649263a.
//
// Solidity: function isValidator(uint256 validatorId) view returns(bool)
func (_StaderStakeManager *StaderStakeManagerSession) IsValidator(validatorId *big.Int) (bool, error) {
	return _StaderStakeManager.Contract.IsValidator(&_StaderStakeManager.CallOpts, validatorId)
}

// IsValidator is a free data retrieval call binding the contract method 0x2649263a.
//
// Solidity: function isValidator(uint256 validatorId) view returns(bool)
func (_StaderStakeManager *StaderStakeManagerCallerSession) IsValidator(validatorId *big.Int) (bool, error) {
	return _StaderStakeManager.Contract.IsValidator(&_StaderStakeManager.CallOpts, validatorId)
}

// LatestSignerUpdateEpoch is a free data retrieval call binding the contract method 0xd7f5549d.
//
// Solidity: function latestSignerUpdateEpoch(uint256 ) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) LatestSignerUpdateEpoch(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "latestSignerUpdateEpoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestSignerUpdateEpoch is a free data retrieval call binding the contract method 0xd7f5549d.
//
// Solidity: function latestSignerUpdateEpoch(uint256 ) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) LatestSignerUpdateEpoch(arg0 *big.Int) (*big.Int, error) {
	return _StaderStakeManager.Contract.LatestSignerUpdateEpoch(&_StaderStakeManager.CallOpts, arg0)
}

// LatestSignerUpdateEpoch is a free data retrieval call binding the contract method 0xd7f5549d.
//
// Solidity: function latestSignerUpdateEpoch(uint256 ) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) LatestSignerUpdateEpoch(arg0 *big.Int) (*big.Int, error) {
	return _StaderStakeManager.Contract.LatestSignerUpdateEpoch(&_StaderStakeManager.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_StaderStakeManager *StaderStakeManagerCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_StaderStakeManager *StaderStakeManagerSession) Locked() (bool, error) {
	return _StaderStakeManager.Contract.Locked(&_StaderStakeManager.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_StaderStakeManager *StaderStakeManagerCallerSession) Locked() (bool, error) {
	return _StaderStakeManager.Contract.Locked(&_StaderStakeManager.CallOpts)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) Logger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "logger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) Logger() (common.Address, error) {
	return _StaderStakeManager.Contract.Logger(&_StaderStakeManager.CallOpts)
}

// Logger is a free data retrieval call binding the contract method 0xf24ccbfe.
//
// Solidity: function logger() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) Logger() (common.Address, error) {
	return _StaderStakeManager.Contract.Logger(&_StaderStakeManager.CallOpts)
}

// MaxRewardedCheckpoints is a free data retrieval call binding the contract method 0x451b5985.
//
// Solidity: function maxRewardedCheckpoints() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) MaxRewardedCheckpoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "maxRewardedCheckpoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRewardedCheckpoints is a free data retrieval call binding the contract method 0x451b5985.
//
// Solidity: function maxRewardedCheckpoints() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) MaxRewardedCheckpoints() (*big.Int, error) {
	return _StaderStakeManager.Contract.MaxRewardedCheckpoints(&_StaderStakeManager.CallOpts)
}

// MaxRewardedCheckpoints is a free data retrieval call binding the contract method 0x451b5985.
//
// Solidity: function maxRewardedCheckpoints() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) MaxRewardedCheckpoints() (*big.Int, error) {
	return _StaderStakeManager.Contract.MaxRewardedCheckpoints(&_StaderStakeManager.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "minDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) MinDeposit() (*big.Int, error) {
	return _StaderStakeManager.Contract.MinDeposit(&_StaderStakeManager.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) MinDeposit() (*big.Int, error) {
	return _StaderStakeManager.Contract.MinDeposit(&_StaderStakeManager.CallOpts)
}

// MinHeimdallFee is a free data retrieval call binding the contract method 0xfba58f34.
//
// Solidity: function minHeimdallFee() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) MinHeimdallFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "minHeimdallFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinHeimdallFee is a free data retrieval call binding the contract method 0xfba58f34.
//
// Solidity: function minHeimdallFee() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) MinHeimdallFee() (*big.Int, error) {
	return _StaderStakeManager.Contract.MinHeimdallFee(&_StaderStakeManager.CallOpts)
}

// MinHeimdallFee is a free data retrieval call binding the contract method 0xfba58f34.
//
// Solidity: function minHeimdallFee() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) MinHeimdallFee() (*big.Int, error) {
	return _StaderStakeManager.Contract.MinHeimdallFee(&_StaderStakeManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) Owner() (common.Address, error) {
	return _StaderStakeManager.Contract.Owner(&_StaderStakeManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) Owner() (common.Address, error) {
	return _StaderStakeManager.Contract.Owner(&_StaderStakeManager.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StaderStakeManager.Contract.OwnerOf(&_StaderStakeManager.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StaderStakeManager.Contract.OwnerOf(&_StaderStakeManager.CallOpts, tokenId)
}

// PrevBlockInterval is a free data retrieval call binding the contract method 0x91f1a3a5.
//
// Solidity: function prevBlockInterval() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) PrevBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "prevBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrevBlockInterval is a free data retrieval call binding the contract method 0x91f1a3a5.
//
// Solidity: function prevBlockInterval() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) PrevBlockInterval() (*big.Int, error) {
	return _StaderStakeManager.Contract.PrevBlockInterval(&_StaderStakeManager.CallOpts)
}

// PrevBlockInterval is a free data retrieval call binding the contract method 0x91f1a3a5.
//
// Solidity: function prevBlockInterval() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) PrevBlockInterval() (*big.Int, error) {
	return _StaderStakeManager.Contract.PrevBlockInterval(&_StaderStakeManager.CallOpts)
}

// ProposerBonus is a free data retrieval call binding the contract method 0x34274586.
//
// Solidity: function proposerBonus() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) ProposerBonus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "proposerBonus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposerBonus is a free data retrieval call binding the contract method 0x34274586.
//
// Solidity: function proposerBonus() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) ProposerBonus() (*big.Int, error) {
	return _StaderStakeManager.Contract.ProposerBonus(&_StaderStakeManager.CallOpts)
}

// ProposerBonus is a free data retrieval call binding the contract method 0x34274586.
//
// Solidity: function proposerBonus() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) ProposerBonus() (*big.Int, error) {
	return _StaderStakeManager.Contract.ProposerBonus(&_StaderStakeManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) Registry() (common.Address, error) {
	return _StaderStakeManager.Contract.Registry(&_StaderStakeManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) Registry() (common.Address, error) {
	return _StaderStakeManager.Contract.Registry(&_StaderStakeManager.CallOpts)
}

// ReplacementCoolDown is a free data retrieval call binding the contract method 0x77939d10.
//
// Solidity: function replacementCoolDown() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) ReplacementCoolDown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "replacementCoolDown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReplacementCoolDown is a free data retrieval call binding the contract method 0x77939d10.
//
// Solidity: function replacementCoolDown() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) ReplacementCoolDown() (*big.Int, error) {
	return _StaderStakeManager.Contract.ReplacementCoolDown(&_StaderStakeManager.CallOpts)
}

// ReplacementCoolDown is a free data retrieval call binding the contract method 0x77939d10.
//
// Solidity: function replacementCoolDown() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) ReplacementCoolDown() (*big.Int, error) {
	return _StaderStakeManager.Contract.ReplacementCoolDown(&_StaderStakeManager.CallOpts)
}

// RewardDecreasePerCheckpoint is a free data retrieval call binding the contract method 0xe568959a.
//
// Solidity: function rewardDecreasePerCheckpoint() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) RewardDecreasePerCheckpoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "rewardDecreasePerCheckpoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardDecreasePerCheckpoint is a free data retrieval call binding the contract method 0xe568959a.
//
// Solidity: function rewardDecreasePerCheckpoint() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) RewardDecreasePerCheckpoint() (*big.Int, error) {
	return _StaderStakeManager.Contract.RewardDecreasePerCheckpoint(&_StaderStakeManager.CallOpts)
}

// RewardDecreasePerCheckpoint is a free data retrieval call binding the contract method 0xe568959a.
//
// Solidity: function rewardDecreasePerCheckpoint() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) RewardDecreasePerCheckpoint() (*big.Int, error) {
	return _StaderStakeManager.Contract.RewardDecreasePerCheckpoint(&_StaderStakeManager.CallOpts)
}

// RewardPerStake is a free data retrieval call binding the contract method 0xa8dc889b.
//
// Solidity: function rewardPerStake() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) RewardPerStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "rewardPerStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerStake is a free data retrieval call binding the contract method 0xa8dc889b.
//
// Solidity: function rewardPerStake() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) RewardPerStake() (*big.Int, error) {
	return _StaderStakeManager.Contract.RewardPerStake(&_StaderStakeManager.CallOpts)
}

// RewardPerStake is a free data retrieval call binding the contract method 0xa8dc889b.
//
// Solidity: function rewardPerStake() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) RewardPerStake() (*big.Int, error) {
	return _StaderStakeManager.Contract.RewardPerStake(&_StaderStakeManager.CallOpts)
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) RootChain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "rootChain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) RootChain() (common.Address, error) {
	return _StaderStakeManager.Contract.RootChain(&_StaderStakeManager.CallOpts)
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) RootChain() (common.Address, error) {
	return _StaderStakeManager.Contract.RootChain(&_StaderStakeManager.CallOpts)
}

// SignerToValidator is a free data retrieval call binding the contract method 0x3862da0b.
//
// Solidity: function signerToValidator(address ) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) SignerToValidator(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "signerToValidator", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignerToValidator is a free data retrieval call binding the contract method 0x3862da0b.
//
// Solidity: function signerToValidator(address ) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) SignerToValidator(arg0 common.Address) (*big.Int, error) {
	return _StaderStakeManager.Contract.SignerToValidator(&_StaderStakeManager.CallOpts, arg0)
}

// SignerToValidator is a free data retrieval call binding the contract method 0x3862da0b.
//
// Solidity: function signerToValidator(address ) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) SignerToValidator(arg0 common.Address) (*big.Int, error) {
	return _StaderStakeManager.Contract.SignerToValidator(&_StaderStakeManager.CallOpts, arg0)
}

// SignerUpdateLimit is a free data retrieval call binding the contract method 0x4e3c83f1.
//
// Solidity: function signerUpdateLimit() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) SignerUpdateLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "signerUpdateLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignerUpdateLimit is a free data retrieval call binding the contract method 0x4e3c83f1.
//
// Solidity: function signerUpdateLimit() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) SignerUpdateLimit() (*big.Int, error) {
	return _StaderStakeManager.Contract.SignerUpdateLimit(&_StaderStakeManager.CallOpts)
}

// SignerUpdateLimit is a free data retrieval call binding the contract method 0x4e3c83f1.
//
// Solidity: function signerUpdateLimit() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) SignerUpdateLimit() (*big.Int, error) {
	return _StaderStakeManager.Contract.SignerUpdateLimit(&_StaderStakeManager.CallOpts)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) Signers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "signers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _StaderStakeManager.Contract.Signers(&_StaderStakeManager.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _StaderStakeManager.Contract.Signers(&_StaderStakeManager.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) Token() (common.Address, error) {
	return _StaderStakeManager.Contract.Token(&_StaderStakeManager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) Token() (common.Address, error) {
	return _StaderStakeManager.Contract.Token(&_StaderStakeManager.CallOpts)
}

// TotalHeimdallFee is a free data retrieval call binding the contract method 0x9a8a6243.
//
// Solidity: function totalHeimdallFee() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) TotalHeimdallFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "totalHeimdallFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalHeimdallFee is a free data retrieval call binding the contract method 0x9a8a6243.
//
// Solidity: function totalHeimdallFee() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) TotalHeimdallFee() (*big.Int, error) {
	return _StaderStakeManager.Contract.TotalHeimdallFee(&_StaderStakeManager.CallOpts)
}

// TotalHeimdallFee is a free data retrieval call binding the contract method 0x9a8a6243.
//
// Solidity: function totalHeimdallFee() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) TotalHeimdallFee() (*big.Int, error) {
	return _StaderStakeManager.Contract.TotalHeimdallFee(&_StaderStakeManager.CallOpts)
}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) TotalRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "totalRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) TotalRewards() (*big.Int, error) {
	return _StaderStakeManager.Contract.TotalRewards(&_StaderStakeManager.CallOpts)
}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) TotalRewards() (*big.Int, error) {
	return _StaderStakeManager.Contract.TotalRewards(&_StaderStakeManager.CallOpts)
}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) TotalRewardsLiquidated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "totalRewardsLiquidated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) TotalRewardsLiquidated() (*big.Int, error) {
	return _StaderStakeManager.Contract.TotalRewardsLiquidated(&_StaderStakeManager.CallOpts)
}

// TotalRewardsLiquidated is a free data retrieval call binding the contract method 0xcd6b8388.
//
// Solidity: function totalRewardsLiquidated() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) TotalRewardsLiquidated() (*big.Int, error) {
	return _StaderStakeManager.Contract.TotalRewardsLiquidated(&_StaderStakeManager.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) TotalStaked() (*big.Int, error) {
	return _StaderStakeManager.Contract.TotalStaked(&_StaderStakeManager.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) TotalStaked() (*big.Int, error) {
	return _StaderStakeManager.Contract.TotalStaked(&_StaderStakeManager.CallOpts)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address user) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) TotalStakedFor(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "totalStakedFor", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address user) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) TotalStakedFor(user common.Address) (*big.Int, error) {
	return _StaderStakeManager.Contract.TotalStakedFor(&_StaderStakeManager.CallOpts, user)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address user) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) TotalStakedFor(user common.Address) (*big.Int, error) {
	return _StaderStakeManager.Contract.TotalStakedFor(&_StaderStakeManager.CallOpts, user)
}

// UserFeeExit is a free data retrieval call binding the contract method 0x78f84a44.
//
// Solidity: function userFeeExit(address ) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) UserFeeExit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "userFeeExit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserFeeExit is a free data retrieval call binding the contract method 0x78f84a44.
//
// Solidity: function userFeeExit(address ) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) UserFeeExit(arg0 common.Address) (*big.Int, error) {
	return _StaderStakeManager.Contract.UserFeeExit(&_StaderStakeManager.CallOpts, arg0)
}

// UserFeeExit is a free data retrieval call binding the contract method 0x78f84a44.
//
// Solidity: function userFeeExit(address ) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) UserFeeExit(arg0 common.Address) (*big.Int, error) {
	return _StaderStakeManager.Contract.UserFeeExit(&_StaderStakeManager.CallOpts, arg0)
}

// ValidatorAuction is a free data retrieval call binding the contract method 0x5325e144.
//
// Solidity: function validatorAuction(uint256 ) view returns(uint256 amount, uint256 startEpoch, address user, bool acceptDelegation, bytes signerPubkey)
func (_StaderStakeManager *StaderStakeManagerCaller) ValidatorAuction(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount           *big.Int
	StartEpoch       *big.Int
	User             common.Address
	AcceptDelegation bool
	SignerPubkey     []byte
}, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "validatorAuction", arg0)

	outstruct := new(struct {
		Amount           *big.Int
		StartEpoch       *big.Int
		User             common.Address
		AcceptDelegation bool
		SignerPubkey     []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.User = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.AcceptDelegation = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.SignerPubkey = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ValidatorAuction is a free data retrieval call binding the contract method 0x5325e144.
//
// Solidity: function validatorAuction(uint256 ) view returns(uint256 amount, uint256 startEpoch, address user, bool acceptDelegation, bytes signerPubkey)
func (_StaderStakeManager *StaderStakeManagerSession) ValidatorAuction(arg0 *big.Int) (struct {
	Amount           *big.Int
	StartEpoch       *big.Int
	User             common.Address
	AcceptDelegation bool
	SignerPubkey     []byte
}, error) {
	return _StaderStakeManager.Contract.ValidatorAuction(&_StaderStakeManager.CallOpts, arg0)
}

// ValidatorAuction is a free data retrieval call binding the contract method 0x5325e144.
//
// Solidity: function validatorAuction(uint256 ) view returns(uint256 amount, uint256 startEpoch, address user, bool acceptDelegation, bytes signerPubkey)
func (_StaderStakeManager *StaderStakeManagerCallerSession) ValidatorAuction(arg0 *big.Int) (struct {
	Amount           *big.Int
	StartEpoch       *big.Int
	User             common.Address
	AcceptDelegation bool
	SignerPubkey     []byte
}, error) {
	return _StaderStakeManager.Contract.ValidatorAuction(&_StaderStakeManager.CallOpts, arg0)
}

// ValidatorReward is a free data retrieval call binding the contract method 0xb65de35e.
//
// Solidity: function validatorReward(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) ValidatorReward(opts *bind.CallOpts, validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "validatorReward", validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorReward is a free data retrieval call binding the contract method 0xb65de35e.
//
// Solidity: function validatorReward(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) ValidatorReward(validatorId *big.Int) (*big.Int, error) {
	return _StaderStakeManager.Contract.ValidatorReward(&_StaderStakeManager.CallOpts, validatorId)
}

// ValidatorReward is a free data retrieval call binding the contract method 0xb65de35e.
//
// Solidity: function validatorReward(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) ValidatorReward(validatorId *big.Int) (*big.Int, error) {
	return _StaderStakeManager.Contract.ValidatorReward(&_StaderStakeManager.CallOpts, validatorId)
}

// ValidatorShareFactory is a free data retrieval call binding the contract method 0x1ae4818f.
//
// Solidity: function validatorShareFactory() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCaller) ValidatorShareFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "validatorShareFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorShareFactory is a free data retrieval call binding the contract method 0x1ae4818f.
//
// Solidity: function validatorShareFactory() view returns(address)
func (_StaderStakeManager *StaderStakeManagerSession) ValidatorShareFactory() (common.Address, error) {
	return _StaderStakeManager.Contract.ValidatorShareFactory(&_StaderStakeManager.CallOpts)
}

// ValidatorShareFactory is a free data retrieval call binding the contract method 0x1ae4818f.
//
// Solidity: function validatorShareFactory() view returns(address)
func (_StaderStakeManager *StaderStakeManagerCallerSession) ValidatorShareFactory() (common.Address, error) {
	return _StaderStakeManager.Contract.ValidatorShareFactory(&_StaderStakeManager.CallOpts)
}

// ValidatorStake is a free data retrieval call binding the contract method 0xeceec1d3.
//
// Solidity: function validatorStake(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) ValidatorStake(opts *bind.CallOpts, validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "validatorStake", validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorStake is a free data retrieval call binding the contract method 0xeceec1d3.
//
// Solidity: function validatorStake(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) ValidatorStake(validatorId *big.Int) (*big.Int, error) {
	return _StaderStakeManager.Contract.ValidatorStake(&_StaderStakeManager.CallOpts, validatorId)
}

// ValidatorStake is a free data retrieval call binding the contract method 0xeceec1d3.
//
// Solidity: function validatorStake(uint256 validatorId) view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) ValidatorStake(validatorId *big.Int) (*big.Int, error) {
	return _StaderStakeManager.Contract.ValidatorStake(&_StaderStakeManager.CallOpts, validatorId)
}

// ValidatorState is a free data retrieval call binding the contract method 0xe59ee0c6.
//
// Solidity: function validatorState() view returns(uint256 amount, uint256 stakerCount)
func (_StaderStakeManager *StaderStakeManagerCaller) ValidatorState(opts *bind.CallOpts) (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "validatorState")

	outstruct := new(struct {
		Amount      *big.Int
		StakerCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StakerCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatorState is a free data retrieval call binding the contract method 0xe59ee0c6.
//
// Solidity: function validatorState() view returns(uint256 amount, uint256 stakerCount)
func (_StaderStakeManager *StaderStakeManagerSession) ValidatorState() (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	return _StaderStakeManager.Contract.ValidatorState(&_StaderStakeManager.CallOpts)
}

// ValidatorState is a free data retrieval call binding the contract method 0xe59ee0c6.
//
// Solidity: function validatorState() view returns(uint256 amount, uint256 stakerCount)
func (_StaderStakeManager *StaderStakeManagerCallerSession) ValidatorState() (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	return _StaderStakeManager.Contract.ValidatorState(&_StaderStakeManager.CallOpts)
}

// ValidatorStateChanges is a free data retrieval call binding the contract method 0x25726df2.
//
// Solidity: function validatorStateChanges(uint256 ) view returns(int256 amount, int256 stakerCount)
func (_StaderStakeManager *StaderStakeManagerCaller) ValidatorStateChanges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "validatorStateChanges", arg0)

	outstruct := new(struct {
		Amount      *big.Int
		StakerCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StakerCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatorStateChanges is a free data retrieval call binding the contract method 0x25726df2.
//
// Solidity: function validatorStateChanges(uint256 ) view returns(int256 amount, int256 stakerCount)
func (_StaderStakeManager *StaderStakeManagerSession) ValidatorStateChanges(arg0 *big.Int) (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	return _StaderStakeManager.Contract.ValidatorStateChanges(&_StaderStakeManager.CallOpts, arg0)
}

// ValidatorStateChanges is a free data retrieval call binding the contract method 0x25726df2.
//
// Solidity: function validatorStateChanges(uint256 ) view returns(int256 amount, int256 stakerCount)
func (_StaderStakeManager *StaderStakeManagerCallerSession) ValidatorStateChanges(arg0 *big.Int) (struct {
	Amount      *big.Int
	StakerCount *big.Int
}, error) {
	return _StaderStakeManager.Contract.ValidatorStateChanges(&_StaderStakeManager.CallOpts, arg0)
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x4fd101d7.
//
// Solidity: function validatorThreshold() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) ValidatorThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "validatorThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x4fd101d7.
//
// Solidity: function validatorThreshold() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) ValidatorThreshold() (*big.Int, error) {
	return _StaderStakeManager.Contract.ValidatorThreshold(&_StaderStakeManager.CallOpts)
}

// ValidatorThreshold is a free data retrieval call binding the contract method 0x4fd101d7.
//
// Solidity: function validatorThreshold() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) ValidatorThreshold() (*big.Int, error) {
	return _StaderStakeManager.Contract.ValidatorThreshold(&_StaderStakeManager.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(uint256 amount, uint256 reward, uint256 activationEpoch, uint256 deactivationEpoch, uint256 jailTime, address signer, address contractAddress, uint8 status, uint256 commissionRate, uint256 lastCommissionUpdate, uint256 delegatorsReward, uint256 delegatedAmount, uint256 initialRewardPerStake)
func (_StaderStakeManager *StaderStakeManagerCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount                *big.Int
	Reward                *big.Int
	ActivationEpoch       *big.Int
	DeactivationEpoch     *big.Int
	JailTime              *big.Int
	Signer                common.Address
	ContractAddress       common.Address
	Status                uint8
	CommissionRate        *big.Int
	LastCommissionUpdate  *big.Int
	DelegatorsReward      *big.Int
	DelegatedAmount       *big.Int
	InitialRewardPerStake *big.Int
}, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Amount                *big.Int
		Reward                *big.Int
		ActivationEpoch       *big.Int
		DeactivationEpoch     *big.Int
		JailTime              *big.Int
		Signer                common.Address
		ContractAddress       common.Address
		Status                uint8
		CommissionRate        *big.Int
		LastCommissionUpdate  *big.Int
		DelegatorsReward      *big.Int
		DelegatedAmount       *big.Int
		InitialRewardPerStake *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ActivationEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeactivationEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.JailTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Signer = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.ContractAddress = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.CommissionRate = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LastCommissionUpdate = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.DelegatorsReward = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.DelegatedAmount = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.InitialRewardPerStake = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(uint256 amount, uint256 reward, uint256 activationEpoch, uint256 deactivationEpoch, uint256 jailTime, address signer, address contractAddress, uint8 status, uint256 commissionRate, uint256 lastCommissionUpdate, uint256 delegatorsReward, uint256 delegatedAmount, uint256 initialRewardPerStake)
func (_StaderStakeManager *StaderStakeManagerSession) Validators(arg0 *big.Int) (struct {
	Amount                *big.Int
	Reward                *big.Int
	ActivationEpoch       *big.Int
	DeactivationEpoch     *big.Int
	JailTime              *big.Int
	Signer                common.Address
	ContractAddress       common.Address
	Status                uint8
	CommissionRate        *big.Int
	LastCommissionUpdate  *big.Int
	DelegatorsReward      *big.Int
	DelegatedAmount       *big.Int
	InitialRewardPerStake *big.Int
}, error) {
	return _StaderStakeManager.Contract.Validators(&_StaderStakeManager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(uint256 amount, uint256 reward, uint256 activationEpoch, uint256 deactivationEpoch, uint256 jailTime, address signer, address contractAddress, uint8 status, uint256 commissionRate, uint256 lastCommissionUpdate, uint256 delegatorsReward, uint256 delegatedAmount, uint256 initialRewardPerStake)
func (_StaderStakeManager *StaderStakeManagerCallerSession) Validators(arg0 *big.Int) (struct {
	Amount                *big.Int
	Reward                *big.Int
	ActivationEpoch       *big.Int
	DeactivationEpoch     *big.Int
	JailTime              *big.Int
	Signer                common.Address
	ContractAddress       common.Address
	Status                uint8
	CommissionRate        *big.Int
	LastCommissionUpdate  *big.Int
	DelegatorsReward      *big.Int
	DelegatedAmount       *big.Int
	InitialRewardPerStake *big.Int
}, error) {
	return _StaderStakeManager.Contract.Validators(&_StaderStakeManager.CallOpts, arg0)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCaller) WithdrawalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderStakeManager.contract.Call(opts, &out, "withdrawalDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) WithdrawalDelay() (*big.Int, error) {
	return _StaderStakeManager.Contract.WithdrawalDelay(&_StaderStakeManager.CallOpts)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_StaderStakeManager *StaderStakeManagerCallerSession) WithdrawalDelay() (*big.Int, error) {
	return _StaderStakeManager.Contract.WithdrawalDelay(&_StaderStakeManager.CallOpts)
}

// ChangeRootChain is a paid mutator transaction binding the contract method 0xe8afa8e8.
//
// Solidity: function changeRootChain(address newRootChain) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) ChangeRootChain(opts *bind.TransactOpts, newRootChain common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "changeRootChain", newRootChain)
}

// ChangeRootChain is a paid mutator transaction binding the contract method 0xe8afa8e8.
//
// Solidity: function changeRootChain(address newRootChain) returns()
func (_StaderStakeManager *StaderStakeManagerSession) ChangeRootChain(newRootChain common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.ChangeRootChain(&_StaderStakeManager.TransactOpts, newRootChain)
}

// ChangeRootChain is a paid mutator transaction binding the contract method 0xe8afa8e8.
//
// Solidity: function changeRootChain(address newRootChain) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) ChangeRootChain(newRootChain common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.ChangeRootChain(&_StaderStakeManager.TransactOpts, newRootChain)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0x2fa9d18b.
//
// Solidity: function checkSignatures(uint256 blockInterval, bytes32 voteHash, bytes32 stateRoot, address proposer, uint256[3][] sigs) returns(uint256)
func (_StaderStakeManager *StaderStakeManagerTransactor) CheckSignatures(opts *bind.TransactOpts, blockInterval *big.Int, voteHash [32]byte, stateRoot [32]byte, proposer common.Address, sigs [][3]*big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "checkSignatures", blockInterval, voteHash, stateRoot, proposer, sigs)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0x2fa9d18b.
//
// Solidity: function checkSignatures(uint256 blockInterval, bytes32 voteHash, bytes32 stateRoot, address proposer, uint256[3][] sigs) returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) CheckSignatures(blockInterval *big.Int, voteHash [32]byte, stateRoot [32]byte, proposer common.Address, sigs [][3]*big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.CheckSignatures(&_StaderStakeManager.TransactOpts, blockInterval, voteHash, stateRoot, proposer, sigs)
}

// CheckSignatures is a paid mutator transaction binding the contract method 0x2fa9d18b.
//
// Solidity: function checkSignatures(uint256 blockInterval, bytes32 voteHash, bytes32 stateRoot, address proposer, uint256[3][] sigs) returns(uint256)
func (_StaderStakeManager *StaderStakeManagerTransactorSession) CheckSignatures(blockInterval *big.Int, voteHash [32]byte, stateRoot [32]byte, proposer common.Address, sigs [][3]*big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.CheckSignatures(&_StaderStakeManager.TransactOpts, blockInterval, voteHash, stateRoot, proposer, sigs)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x68cb812a.
//
// Solidity: function claimFee(uint256 accumFeeAmount, uint256 index, bytes proof) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) ClaimFee(opts *bind.TransactOpts, accumFeeAmount *big.Int, index *big.Int, proof []byte) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "claimFee", accumFeeAmount, index, proof)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x68cb812a.
//
// Solidity: function claimFee(uint256 accumFeeAmount, uint256 index, bytes proof) returns()
func (_StaderStakeManager *StaderStakeManagerSession) ClaimFee(accumFeeAmount *big.Int, index *big.Int, proof []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.ClaimFee(&_StaderStakeManager.TransactOpts, accumFeeAmount, index, proof)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x68cb812a.
//
// Solidity: function claimFee(uint256 accumFeeAmount, uint256 index, bytes proof) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) ClaimFee(accumFeeAmount *big.Int, index *big.Int, proof []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.ClaimFee(&_StaderStakeManager.TransactOpts, accumFeeAmount, index, proof)
}

// ConfirmAuctionBid is a paid mutator transaction binding the contract method 0x99d18f6f.
//
// Solidity: function confirmAuctionBid(uint256 validatorId, uint256 heimdallFee) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) ConfirmAuctionBid(opts *bind.TransactOpts, validatorId *big.Int, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "confirmAuctionBid", validatorId, heimdallFee)
}

// ConfirmAuctionBid is a paid mutator transaction binding the contract method 0x99d18f6f.
//
// Solidity: function confirmAuctionBid(uint256 validatorId, uint256 heimdallFee) returns()
func (_StaderStakeManager *StaderStakeManagerSession) ConfirmAuctionBid(validatorId *big.Int, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.ConfirmAuctionBid(&_StaderStakeManager.TransactOpts, validatorId, heimdallFee)
}

// ConfirmAuctionBid is a paid mutator transaction binding the contract method 0x99d18f6f.
//
// Solidity: function confirmAuctionBid(uint256 validatorId, uint256 heimdallFee) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) ConfirmAuctionBid(validatorId *big.Int, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.ConfirmAuctionBid(&_StaderStakeManager.TransactOpts, validatorId, heimdallFee)
}

// DecreaseValidatorDelegatedAmount is a paid mutator transaction binding the contract method 0x858a7c03.
//
// Solidity: function decreaseValidatorDelegatedAmount(uint256 validatorId, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) DecreaseValidatorDelegatedAmount(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "decreaseValidatorDelegatedAmount", validatorId, amount)
}

// DecreaseValidatorDelegatedAmount is a paid mutator transaction binding the contract method 0x858a7c03.
//
// Solidity: function decreaseValidatorDelegatedAmount(uint256 validatorId, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerSession) DecreaseValidatorDelegatedAmount(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.DecreaseValidatorDelegatedAmount(&_StaderStakeManager.TransactOpts, validatorId, amount)
}

// DecreaseValidatorDelegatedAmount is a paid mutator transaction binding the contract method 0x858a7c03.
//
// Solidity: function decreaseValidatorDelegatedAmount(uint256 validatorId, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) DecreaseValidatorDelegatedAmount(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.DecreaseValidatorDelegatedAmount(&_StaderStakeManager.TransactOpts, validatorId, amount)
}

// DelegationDeposit is a paid mutator transaction binding the contract method 0x6901b253.
//
// Solidity: function delegationDeposit(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StaderStakeManager *StaderStakeManagerTransactor) DelegationDeposit(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "delegationDeposit", validatorId, amount, delegator)
}

// DelegationDeposit is a paid mutator transaction binding the contract method 0x6901b253.
//
// Solidity: function delegationDeposit(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StaderStakeManager *StaderStakeManagerSession) DelegationDeposit(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.DelegationDeposit(&_StaderStakeManager.TransactOpts, validatorId, amount, delegator)
}

// DelegationDeposit is a paid mutator transaction binding the contract method 0x6901b253.
//
// Solidity: function delegationDeposit(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StaderStakeManager *StaderStakeManagerTransactorSession) DelegationDeposit(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.DelegationDeposit(&_StaderStakeManager.TransactOpts, validatorId, amount, delegator)
}

// DethroneAndStake is a paid mutator transaction binding the contract method 0x52b8115d.
//
// Solidity: function dethroneAndStake(address auctionUser, uint256 heimdallFee, uint256 validatorId, uint256 auctionAmount, bool acceptDelegation, bytes signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) DethroneAndStake(opts *bind.TransactOpts, auctionUser common.Address, heimdallFee *big.Int, validatorId *big.Int, auctionAmount *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "dethroneAndStake", auctionUser, heimdallFee, validatorId, auctionAmount, acceptDelegation, signerPubkey)
}

// DethroneAndStake is a paid mutator transaction binding the contract method 0x52b8115d.
//
// Solidity: function dethroneAndStake(address auctionUser, uint256 heimdallFee, uint256 validatorId, uint256 auctionAmount, bool acceptDelegation, bytes signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerSession) DethroneAndStake(auctionUser common.Address, heimdallFee *big.Int, validatorId *big.Int, auctionAmount *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.DethroneAndStake(&_StaderStakeManager.TransactOpts, auctionUser, heimdallFee, validatorId, auctionAmount, acceptDelegation, signerPubkey)
}

// DethroneAndStake is a paid mutator transaction binding the contract method 0x52b8115d.
//
// Solidity: function dethroneAndStake(address auctionUser, uint256 heimdallFee, uint256 validatorId, uint256 auctionAmount, bool acceptDelegation, bytes signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) DethroneAndStake(auctionUser common.Address, heimdallFee *big.Int, validatorId *big.Int, auctionAmount *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.DethroneAndStake(&_StaderStakeManager.TransactOpts, auctionUser, heimdallFee, validatorId, auctionAmount, acceptDelegation, signerPubkey)
}

// Drain is a paid mutator transaction binding the contract method 0xb184be81.
//
// Solidity: function drain(address destination, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) Drain(opts *bind.TransactOpts, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "drain", destination, amount)
}

// Drain is a paid mutator transaction binding the contract method 0xb184be81.
//
// Solidity: function drain(address destination, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerSession) Drain(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Drain(&_StaderStakeManager.TransactOpts, destination, amount)
}

// Drain is a paid mutator transaction binding the contract method 0xb184be81.
//
// Solidity: function drain(address destination, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) Drain(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Drain(&_StaderStakeManager.TransactOpts, destination, amount)
}

// DrainValidatorShares is a paid mutator transaction binding the contract method 0x48ab8b2a.
//
// Solidity: function drainValidatorShares(uint256 validatorId, address tokenAddr, address destination, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) DrainValidatorShares(opts *bind.TransactOpts, validatorId *big.Int, tokenAddr common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "drainValidatorShares", validatorId, tokenAddr, destination, amount)
}

// DrainValidatorShares is a paid mutator transaction binding the contract method 0x48ab8b2a.
//
// Solidity: function drainValidatorShares(uint256 validatorId, address tokenAddr, address destination, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerSession) DrainValidatorShares(validatorId *big.Int, tokenAddr common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.DrainValidatorShares(&_StaderStakeManager.TransactOpts, validatorId, tokenAddr, destination, amount)
}

// DrainValidatorShares is a paid mutator transaction binding the contract method 0x48ab8b2a.
//
// Solidity: function drainValidatorShares(uint256 validatorId, address tokenAddr, address destination, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) DrainValidatorShares(validatorId *big.Int, tokenAddr common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.DrainValidatorShares(&_StaderStakeManager.TransactOpts, validatorId, tokenAddr, destination, amount)
}

// ForceUnstake is a paid mutator transaction binding the contract method 0x91460149.
//
// Solidity: function forceUnstake(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) ForceUnstake(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "forceUnstake", validatorId)
}

// ForceUnstake is a paid mutator transaction binding the contract method 0x91460149.
//
// Solidity: function forceUnstake(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerSession) ForceUnstake(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.ForceUnstake(&_StaderStakeManager.TransactOpts, validatorId)
}

// ForceUnstake is a paid mutator transaction binding the contract method 0x91460149.
//
// Solidity: function forceUnstake(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) ForceUnstake(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.ForceUnstake(&_StaderStakeManager.TransactOpts, validatorId)
}

// Initialize is a paid mutator transaction binding the contract method 0xf5e95acb.
//
// Solidity: function initialize(address _registry, address _rootchain, address _token, address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _governance, address _owner, address _extensionCode) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _rootchain common.Address, _token common.Address, _NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _governance common.Address, _owner common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "initialize", _registry, _rootchain, _token, _NFTContract, _stakingLogger, _validatorShareFactory, _governance, _owner, _extensionCode)
}

// Initialize is a paid mutator transaction binding the contract method 0xf5e95acb.
//
// Solidity: function initialize(address _registry, address _rootchain, address _token, address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _governance, address _owner, address _extensionCode) returns()
func (_StaderStakeManager *StaderStakeManagerSession) Initialize(_registry common.Address, _rootchain common.Address, _token common.Address, _NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _governance common.Address, _owner common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Initialize(&_StaderStakeManager.TransactOpts, _registry, _rootchain, _token, _NFTContract, _stakingLogger, _validatorShareFactory, _governance, _owner, _extensionCode)
}

// Initialize is a paid mutator transaction binding the contract method 0xf5e95acb.
//
// Solidity: function initialize(address _registry, address _rootchain, address _token, address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _governance, address _owner, address _extensionCode) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) Initialize(_registry common.Address, _rootchain common.Address, _token common.Address, _NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _governance common.Address, _owner common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Initialize(&_StaderStakeManager.TransactOpts, _registry, _rootchain, _token, _NFTContract, _stakingLogger, _validatorShareFactory, _governance, _owner, _extensionCode)
}

// InsertSigners is a paid mutator transaction binding the contract method 0x2cf44a43.
//
// Solidity: function insertSigners(address[] _signers) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) InsertSigners(opts *bind.TransactOpts, _signers []common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "insertSigners", _signers)
}

// InsertSigners is a paid mutator transaction binding the contract method 0x2cf44a43.
//
// Solidity: function insertSigners(address[] _signers) returns()
func (_StaderStakeManager *StaderStakeManagerSession) InsertSigners(_signers []common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.InsertSigners(&_StaderStakeManager.TransactOpts, _signers)
}

// InsertSigners is a paid mutator transaction binding the contract method 0x2cf44a43.
//
// Solidity: function insertSigners(address[] _signers) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) InsertSigners(_signers []common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.InsertSigners(&_StaderStakeManager.TransactOpts, _signers)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_StaderStakeManager *StaderStakeManagerSession) Lock() (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Lock(&_StaderStakeManager.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) Lock() (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Lock(&_StaderStakeManager.TransactOpts)
}

// MigrateDelegation is a paid mutator transaction binding the contract method 0xfb1ef52c.
//
// Solidity: function migrateDelegation(uint256 fromValidatorId, uint256 toValidatorId, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) MigrateDelegation(opts *bind.TransactOpts, fromValidatorId *big.Int, toValidatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "migrateDelegation", fromValidatorId, toValidatorId, amount)
}

// MigrateDelegation is a paid mutator transaction binding the contract method 0xfb1ef52c.
//
// Solidity: function migrateDelegation(uint256 fromValidatorId, uint256 toValidatorId, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerSession) MigrateDelegation(fromValidatorId *big.Int, toValidatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.MigrateDelegation(&_StaderStakeManager.TransactOpts, fromValidatorId, toValidatorId, amount)
}

// MigrateDelegation is a paid mutator transaction binding the contract method 0xfb1ef52c.
//
// Solidity: function migrateDelegation(uint256 fromValidatorId, uint256 toValidatorId, uint256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) MigrateDelegation(fromValidatorId *big.Int, toValidatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.MigrateDelegation(&_StaderStakeManager.TransactOpts, fromValidatorId, toValidatorId, amount)
}

// MigrateValidatorsData is a paid mutator transaction binding the contract method 0x9ddbbf85.
//
// Solidity: function migrateValidatorsData(uint256 validatorIdFrom, uint256 validatorIdTo) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) MigrateValidatorsData(opts *bind.TransactOpts, validatorIdFrom *big.Int, validatorIdTo *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "migrateValidatorsData", validatorIdFrom, validatorIdTo)
}

// MigrateValidatorsData is a paid mutator transaction binding the contract method 0x9ddbbf85.
//
// Solidity: function migrateValidatorsData(uint256 validatorIdFrom, uint256 validatorIdTo) returns()
func (_StaderStakeManager *StaderStakeManagerSession) MigrateValidatorsData(validatorIdFrom *big.Int, validatorIdTo *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.MigrateValidatorsData(&_StaderStakeManager.TransactOpts, validatorIdFrom, validatorIdTo)
}

// MigrateValidatorsData is a paid mutator transaction binding the contract method 0x9ddbbf85.
//
// Solidity: function migrateValidatorsData(uint256 validatorIdFrom, uint256 validatorIdTo) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) MigrateValidatorsData(validatorIdFrom *big.Int, validatorIdTo *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.MigrateValidatorsData(&_StaderStakeManager.TransactOpts, validatorIdFrom, validatorIdTo)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x078a13b1.
//
// Solidity: function reinitialize(address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _extensionCode) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) Reinitialize(opts *bind.TransactOpts, _NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "reinitialize", _NFTContract, _stakingLogger, _validatorShareFactory, _extensionCode)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x078a13b1.
//
// Solidity: function reinitialize(address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _extensionCode) returns()
func (_StaderStakeManager *StaderStakeManagerSession) Reinitialize(_NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Reinitialize(&_StaderStakeManager.TransactOpts, _NFTContract, _stakingLogger, _validatorShareFactory, _extensionCode)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x078a13b1.
//
// Solidity: function reinitialize(address _NFTContract, address _stakingLogger, address _validatorShareFactory, address _extensionCode) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) Reinitialize(_NFTContract common.Address, _stakingLogger common.Address, _validatorShareFactory common.Address, _extensionCode common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Reinitialize(&_StaderStakeManager.TransactOpts, _NFTContract, _stakingLogger, _validatorShareFactory, _extensionCode)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StaderStakeManager *StaderStakeManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StaderStakeManager.Contract.RenounceOwnership(&_StaderStakeManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StaderStakeManager.Contract.RenounceOwnership(&_StaderStakeManager.TransactOpts)
}

// Restake is a paid mutator transaction binding the contract method 0x28cc4e41.
//
// Solidity: function restake(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) Restake(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "restake", validatorId, amount, stakeRewards)
}

// Restake is a paid mutator transaction binding the contract method 0x28cc4e41.
//
// Solidity: function restake(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_StaderStakeManager *StaderStakeManagerSession) Restake(validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Restake(&_StaderStakeManager.TransactOpts, validatorId, amount, stakeRewards)
}

// Restake is a paid mutator transaction binding the contract method 0x28cc4e41.
//
// Solidity: function restake(uint256 validatorId, uint256 amount, bool stakeRewards) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) Restake(validatorId *big.Int, amount *big.Int, stakeRewards bool) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Restake(&_StaderStakeManager.TransactOpts, validatorId, amount, stakeRewards)
}

// SetCurrentEpoch is a paid mutator transaction binding the contract method 0x1dd6b9b1.
//
// Solidity: function setCurrentEpoch(uint256 _currentEpoch) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) SetCurrentEpoch(opts *bind.TransactOpts, _currentEpoch *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "setCurrentEpoch", _currentEpoch)
}

// SetCurrentEpoch is a paid mutator transaction binding the contract method 0x1dd6b9b1.
//
// Solidity: function setCurrentEpoch(uint256 _currentEpoch) returns()
func (_StaderStakeManager *StaderStakeManagerSession) SetCurrentEpoch(_currentEpoch *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.SetCurrentEpoch(&_StaderStakeManager.TransactOpts, _currentEpoch)
}

// SetCurrentEpoch is a paid mutator transaction binding the contract method 0x1dd6b9b1.
//
// Solidity: function setCurrentEpoch(uint256 _currentEpoch) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) SetCurrentEpoch(_currentEpoch *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.SetCurrentEpoch(&_StaderStakeManager.TransactOpts, _currentEpoch)
}

// SetDelegationEnabled is a paid mutator transaction binding the contract method 0xf28699fa.
//
// Solidity: function setDelegationEnabled(bool enabled) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) SetDelegationEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "setDelegationEnabled", enabled)
}

// SetDelegationEnabled is a paid mutator transaction binding the contract method 0xf28699fa.
//
// Solidity: function setDelegationEnabled(bool enabled) returns()
func (_StaderStakeManager *StaderStakeManagerSession) SetDelegationEnabled(enabled bool) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.SetDelegationEnabled(&_StaderStakeManager.TransactOpts, enabled)
}

// SetDelegationEnabled is a paid mutator transaction binding the contract method 0xf28699fa.
//
// Solidity: function setDelegationEnabled(bool enabled) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) SetDelegationEnabled(enabled bool) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.SetDelegationEnabled(&_StaderStakeManager.TransactOpts, enabled)
}

// SetStakingToken is a paid mutator transaction binding the contract method 0x1e9b12ef.
//
// Solidity: function setStakingToken(address _token) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) SetStakingToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "setStakingToken", _token)
}

// SetStakingToken is a paid mutator transaction binding the contract method 0x1e9b12ef.
//
// Solidity: function setStakingToken(address _token) returns()
func (_StaderStakeManager *StaderStakeManagerSession) SetStakingToken(_token common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.SetStakingToken(&_StaderStakeManager.TransactOpts, _token)
}

// SetStakingToken is a paid mutator transaction binding the contract method 0x1e9b12ef.
//
// Solidity: function setStakingToken(address _token) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) SetStakingToken(_token common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.SetStakingToken(&_StaderStakeManager.TransactOpts, _token)
}

// Slash is a paid mutator transaction binding the contract method 0x5e47655f.
//
// Solidity: function slash(bytes _slashingInfoList) returns(uint256)
func (_StaderStakeManager *StaderStakeManagerTransactor) Slash(opts *bind.TransactOpts, _slashingInfoList []byte) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "slash", _slashingInfoList)
}

// Slash is a paid mutator transaction binding the contract method 0x5e47655f.
//
// Solidity: function slash(bytes _slashingInfoList) returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) Slash(_slashingInfoList []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Slash(&_StaderStakeManager.TransactOpts, _slashingInfoList)
}

// Slash is a paid mutator transaction binding the contract method 0x5e47655f.
//
// Solidity: function slash(bytes _slashingInfoList) returns(uint256)
func (_StaderStakeManager *StaderStakeManagerTransactorSession) Slash(_slashingInfoList []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Slash(&_StaderStakeManager.TransactOpts, _slashingInfoList)
}

// StakeFor is a paid mutator transaction binding the contract method 0x4fdd20f1.
//
// Solidity: function stakeFor(address user, uint256 amount, uint256 heimdallFee, bool acceptDelegation, bytes signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) StakeFor(opts *bind.TransactOpts, user common.Address, amount *big.Int, heimdallFee *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "stakeFor", user, amount, heimdallFee, acceptDelegation, signerPubkey)
}

// StakeFor is a paid mutator transaction binding the contract method 0x4fdd20f1.
//
// Solidity: function stakeFor(address user, uint256 amount, uint256 heimdallFee, bool acceptDelegation, bytes signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerSession) StakeFor(user common.Address, amount *big.Int, heimdallFee *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.StakeFor(&_StaderStakeManager.TransactOpts, user, amount, heimdallFee, acceptDelegation, signerPubkey)
}

// StakeFor is a paid mutator transaction binding the contract method 0x4fdd20f1.
//
// Solidity: function stakeFor(address user, uint256 amount, uint256 heimdallFee, bool acceptDelegation, bytes signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) StakeFor(user common.Address, amount *big.Int, heimdallFee *big.Int, acceptDelegation bool, signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.StakeFor(&_StaderStakeManager.TransactOpts, user, amount, heimdallFee, acceptDelegation, signerPubkey)
}

// StartAuction is a paid mutator transaction binding the contract method 0xa6854877.
//
// Solidity: function startAuction(uint256 validatorId, uint256 amount, bool _acceptDelegation, bytes _signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) StartAuction(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, _acceptDelegation bool, _signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "startAuction", validatorId, amount, _acceptDelegation, _signerPubkey)
}

// StartAuction is a paid mutator transaction binding the contract method 0xa6854877.
//
// Solidity: function startAuction(uint256 validatorId, uint256 amount, bool _acceptDelegation, bytes _signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerSession) StartAuction(validatorId *big.Int, amount *big.Int, _acceptDelegation bool, _signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.StartAuction(&_StaderStakeManager.TransactOpts, validatorId, amount, _acceptDelegation, _signerPubkey)
}

// StartAuction is a paid mutator transaction binding the contract method 0xa6854877.
//
// Solidity: function startAuction(uint256 validatorId, uint256 amount, bool _acceptDelegation, bytes _signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) StartAuction(validatorId *big.Int, amount *big.Int, _acceptDelegation bool, _signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.StartAuction(&_StaderStakeManager.TransactOpts, validatorId, amount, _acceptDelegation, _signerPubkey)
}

// StopAuctions is a paid mutator transaction binding the contract method 0xf771fc87.
//
// Solidity: function stopAuctions(uint256 forNCheckpoints) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) StopAuctions(opts *bind.TransactOpts, forNCheckpoints *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "stopAuctions", forNCheckpoints)
}

// StopAuctions is a paid mutator transaction binding the contract method 0xf771fc87.
//
// Solidity: function stopAuctions(uint256 forNCheckpoints) returns()
func (_StaderStakeManager *StaderStakeManagerSession) StopAuctions(forNCheckpoints *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.StopAuctions(&_StaderStakeManager.TransactOpts, forNCheckpoints)
}

// StopAuctions is a paid mutator transaction binding the contract method 0xf771fc87.
//
// Solidity: function stopAuctions(uint256 forNCheckpoints) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) StopAuctions(forNCheckpoints *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.StopAuctions(&_StaderStakeManager.TransactOpts, forNCheckpoints)
}

// TopUpForFee is a paid mutator transaction binding the contract method 0x63656798.
//
// Solidity: function topUpForFee(address user, uint256 heimdallFee) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) TopUpForFee(opts *bind.TransactOpts, user common.Address, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "topUpForFee", user, heimdallFee)
}

// TopUpForFee is a paid mutator transaction binding the contract method 0x63656798.
//
// Solidity: function topUpForFee(address user, uint256 heimdallFee) returns()
func (_StaderStakeManager *StaderStakeManagerSession) TopUpForFee(user common.Address, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.TopUpForFee(&_StaderStakeManager.TransactOpts, user, heimdallFee)
}

// TopUpForFee is a paid mutator transaction binding the contract method 0x63656798.
//
// Solidity: function topUpForFee(address user, uint256 heimdallFee) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) TopUpForFee(user common.Address, heimdallFee *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.TopUpForFee(&_StaderStakeManager.TransactOpts, user, heimdallFee)
}

// TransferFunds is a paid mutator transaction binding the contract method 0xbc8756a9.
//
// Solidity: function transferFunds(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StaderStakeManager *StaderStakeManagerTransactor) TransferFunds(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "transferFunds", validatorId, amount, delegator)
}

// TransferFunds is a paid mutator transaction binding the contract method 0xbc8756a9.
//
// Solidity: function transferFunds(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StaderStakeManager *StaderStakeManagerSession) TransferFunds(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.TransferFunds(&_StaderStakeManager.TransactOpts, validatorId, amount, delegator)
}

// TransferFunds is a paid mutator transaction binding the contract method 0xbc8756a9.
//
// Solidity: function transferFunds(uint256 validatorId, uint256 amount, address delegator) returns(bool)
func (_StaderStakeManager *StaderStakeManagerTransactorSession) TransferFunds(validatorId *big.Int, amount *big.Int, delegator common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.TransferFunds(&_StaderStakeManager.TransactOpts, validatorId, amount, delegator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StaderStakeManager *StaderStakeManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.TransferOwnership(&_StaderStakeManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.TransferOwnership(&_StaderStakeManager.TransactOpts, newOwner)
}

// Unjail is a paid mutator transaction binding the contract method 0x178c2c83.
//
// Solidity: function unjail(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) Unjail(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "unjail", validatorId)
}

// Unjail is a paid mutator transaction binding the contract method 0x178c2c83.
//
// Solidity: function unjail(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerSession) Unjail(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Unjail(&_StaderStakeManager.TransactOpts, validatorId)
}

// Unjail is a paid mutator transaction binding the contract method 0x178c2c83.
//
// Solidity: function unjail(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) Unjail(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Unjail(&_StaderStakeManager.TransactOpts, validatorId)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_StaderStakeManager *StaderStakeManagerSession) Unlock() (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Unlock(&_StaderStakeManager.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) Unlock() (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Unlock(&_StaderStakeManager.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) Unstake(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "unstake", validatorId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerSession) Unstake(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Unstake(&_StaderStakeManager.TransactOpts, validatorId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) Unstake(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.Unstake(&_StaderStakeManager.TransactOpts, validatorId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0xd86d53e7.
//
// Solidity: function unstakeClaim(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UnstakeClaim(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "unstakeClaim", validatorId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0xd86d53e7.
//
// Solidity: function unstakeClaim(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UnstakeClaim(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UnstakeClaim(&_StaderStakeManager.TransactOpts, validatorId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0xd86d53e7.
//
// Solidity: function unstakeClaim(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UnstakeClaim(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UnstakeClaim(&_StaderStakeManager.TransactOpts, validatorId)
}

// UpdateCheckPointBlockInterval is a paid mutator transaction binding the contract method 0xa440ab1e.
//
// Solidity: function updateCheckPointBlockInterval(uint256 _blocks) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateCheckPointBlockInterval(opts *bind.TransactOpts, _blocks *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateCheckPointBlockInterval", _blocks)
}

// UpdateCheckPointBlockInterval is a paid mutator transaction binding the contract method 0xa440ab1e.
//
// Solidity: function updateCheckPointBlockInterval(uint256 _blocks) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateCheckPointBlockInterval(_blocks *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateCheckPointBlockInterval(&_StaderStakeManager.TransactOpts, _blocks)
}

// UpdateCheckPointBlockInterval is a paid mutator transaction binding the contract method 0xa440ab1e.
//
// Solidity: function updateCheckPointBlockInterval(uint256 _blocks) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateCheckPointBlockInterval(_blocks *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateCheckPointBlockInterval(&_StaderStakeManager.TransactOpts, _blocks)
}

// UpdateCheckpointReward is a paid mutator transaction binding the contract method 0xcbf383d5.
//
// Solidity: function updateCheckpointReward(uint256 newReward) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateCheckpointReward(opts *bind.TransactOpts, newReward *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateCheckpointReward", newReward)
}

// UpdateCheckpointReward is a paid mutator transaction binding the contract method 0xcbf383d5.
//
// Solidity: function updateCheckpointReward(uint256 newReward) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateCheckpointReward(newReward *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateCheckpointReward(&_StaderStakeManager.TransactOpts, newReward)
}

// UpdateCheckpointReward is a paid mutator transaction binding the contract method 0xcbf383d5.
//
// Solidity: function updateCheckpointReward(uint256 newReward) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateCheckpointReward(newReward *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateCheckpointReward(&_StaderStakeManager.TransactOpts, newReward)
}

// UpdateCheckpointRewardParams is a paid mutator transaction binding the contract method 0x60c8d122.
//
// Solidity: function updateCheckpointRewardParams(uint256 _rewardDecreasePerCheckpoint, uint256 _maxRewardedCheckpoints, uint256 _checkpointRewardDelta) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateCheckpointRewardParams(opts *bind.TransactOpts, _rewardDecreasePerCheckpoint *big.Int, _maxRewardedCheckpoints *big.Int, _checkpointRewardDelta *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateCheckpointRewardParams", _rewardDecreasePerCheckpoint, _maxRewardedCheckpoints, _checkpointRewardDelta)
}

// UpdateCheckpointRewardParams is a paid mutator transaction binding the contract method 0x60c8d122.
//
// Solidity: function updateCheckpointRewardParams(uint256 _rewardDecreasePerCheckpoint, uint256 _maxRewardedCheckpoints, uint256 _checkpointRewardDelta) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateCheckpointRewardParams(_rewardDecreasePerCheckpoint *big.Int, _maxRewardedCheckpoints *big.Int, _checkpointRewardDelta *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateCheckpointRewardParams(&_StaderStakeManager.TransactOpts, _rewardDecreasePerCheckpoint, _maxRewardedCheckpoints, _checkpointRewardDelta)
}

// UpdateCheckpointRewardParams is a paid mutator transaction binding the contract method 0x60c8d122.
//
// Solidity: function updateCheckpointRewardParams(uint256 _rewardDecreasePerCheckpoint, uint256 _maxRewardedCheckpoints, uint256 _checkpointRewardDelta) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateCheckpointRewardParams(_rewardDecreasePerCheckpoint *big.Int, _maxRewardedCheckpoints *big.Int, _checkpointRewardDelta *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateCheckpointRewardParams(&_StaderStakeManager.TransactOpts, _rewardDecreasePerCheckpoint, _maxRewardedCheckpoints, _checkpointRewardDelta)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0xdcd962b2.
//
// Solidity: function updateCommissionRate(uint256 validatorId, uint256 newCommissionRate) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateCommissionRate(opts *bind.TransactOpts, validatorId *big.Int, newCommissionRate *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateCommissionRate", validatorId, newCommissionRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0xdcd962b2.
//
// Solidity: function updateCommissionRate(uint256 validatorId, uint256 newCommissionRate) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateCommissionRate(validatorId *big.Int, newCommissionRate *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateCommissionRate(&_StaderStakeManager.TransactOpts, validatorId, newCommissionRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0xdcd962b2.
//
// Solidity: function updateCommissionRate(uint256 validatorId, uint256 newCommissionRate) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateCommissionRate(validatorId *big.Int, newCommissionRate *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateCommissionRate(&_StaderStakeManager.TransactOpts, validatorId, newCommissionRate)
}

// UpdateDynastyValue is a paid mutator transaction binding the contract method 0xe6692f49.
//
// Solidity: function updateDynastyValue(uint256 newDynasty) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateDynastyValue(opts *bind.TransactOpts, newDynasty *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateDynastyValue", newDynasty)
}

// UpdateDynastyValue is a paid mutator transaction binding the contract method 0xe6692f49.
//
// Solidity: function updateDynastyValue(uint256 newDynasty) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateDynastyValue(newDynasty *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateDynastyValue(&_StaderStakeManager.TransactOpts, newDynasty)
}

// UpdateDynastyValue is a paid mutator transaction binding the contract method 0xe6692f49.
//
// Solidity: function updateDynastyValue(uint256 newDynasty) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateDynastyValue(newDynasty *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateDynastyValue(&_StaderStakeManager.TransactOpts, newDynasty)
}

// UpdateMinAmounts is a paid mutator transaction binding the contract method 0xb1d23f02.
//
// Solidity: function updateMinAmounts(uint256 _minDeposit, uint256 _minHeimdallFee) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateMinAmounts(opts *bind.TransactOpts, _minDeposit *big.Int, _minHeimdallFee *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateMinAmounts", _minDeposit, _minHeimdallFee)
}

// UpdateMinAmounts is a paid mutator transaction binding the contract method 0xb1d23f02.
//
// Solidity: function updateMinAmounts(uint256 _minDeposit, uint256 _minHeimdallFee) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateMinAmounts(_minDeposit *big.Int, _minHeimdallFee *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateMinAmounts(&_StaderStakeManager.TransactOpts, _minDeposit, _minHeimdallFee)
}

// UpdateMinAmounts is a paid mutator transaction binding the contract method 0xb1d23f02.
//
// Solidity: function updateMinAmounts(uint256 _minDeposit, uint256 _minHeimdallFee) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateMinAmounts(_minDeposit *big.Int, _minHeimdallFee *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateMinAmounts(&_StaderStakeManager.TransactOpts, _minDeposit, _minHeimdallFee)
}

// UpdateProposerBonus is a paid mutator transaction binding the contract method 0x9b33f434.
//
// Solidity: function updateProposerBonus(uint256 newProposerBonus) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateProposerBonus(opts *bind.TransactOpts, newProposerBonus *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateProposerBonus", newProposerBonus)
}

// UpdateProposerBonus is a paid mutator transaction binding the contract method 0x9b33f434.
//
// Solidity: function updateProposerBonus(uint256 newProposerBonus) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateProposerBonus(newProposerBonus *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateProposerBonus(&_StaderStakeManager.TransactOpts, newProposerBonus)
}

// UpdateProposerBonus is a paid mutator transaction binding the contract method 0x9b33f434.
//
// Solidity: function updateProposerBonus(uint256 newProposerBonus) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateProposerBonus(newProposerBonus *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateProposerBonus(&_StaderStakeManager.TransactOpts, newProposerBonus)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf41a9642.
//
// Solidity: function updateSigner(uint256 validatorId, bytes signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateSigner(opts *bind.TransactOpts, validatorId *big.Int, signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateSigner", validatorId, signerPubkey)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf41a9642.
//
// Solidity: function updateSigner(uint256 validatorId, bytes signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateSigner(validatorId *big.Int, signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateSigner(&_StaderStakeManager.TransactOpts, validatorId, signerPubkey)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf41a9642.
//
// Solidity: function updateSigner(uint256 validatorId, bytes signerPubkey) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateSigner(validatorId *big.Int, signerPubkey []byte) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateSigner(&_StaderStakeManager.TransactOpts, validatorId, signerPubkey)
}

// UpdateSignerUpdateLimit is a paid mutator transaction binding the contract method 0x06cfb104.
//
// Solidity: function updateSignerUpdateLimit(uint256 _limit) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateSignerUpdateLimit(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateSignerUpdateLimit", _limit)
}

// UpdateSignerUpdateLimit is a paid mutator transaction binding the contract method 0x06cfb104.
//
// Solidity: function updateSignerUpdateLimit(uint256 _limit) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateSignerUpdateLimit(_limit *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateSignerUpdateLimit(&_StaderStakeManager.TransactOpts, _limit)
}

// UpdateSignerUpdateLimit is a paid mutator transaction binding the contract method 0x06cfb104.
//
// Solidity: function updateSignerUpdateLimit(uint256 _limit) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateSignerUpdateLimit(_limit *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateSignerUpdateLimit(&_StaderStakeManager.TransactOpts, _limit)
}

// UpdateValidatorContractAddress is a paid mutator transaction binding the contract method 0xc710e922.
//
// Solidity: function updateValidatorContractAddress(uint256 validatorId, address newContractAddress) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateValidatorContractAddress(opts *bind.TransactOpts, validatorId *big.Int, newContractAddress common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateValidatorContractAddress", validatorId, newContractAddress)
}

// UpdateValidatorContractAddress is a paid mutator transaction binding the contract method 0xc710e922.
//
// Solidity: function updateValidatorContractAddress(uint256 validatorId, address newContractAddress) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateValidatorContractAddress(validatorId *big.Int, newContractAddress common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateValidatorContractAddress(&_StaderStakeManager.TransactOpts, validatorId, newContractAddress)
}

// UpdateValidatorContractAddress is a paid mutator transaction binding the contract method 0xc710e922.
//
// Solidity: function updateValidatorContractAddress(uint256 validatorId, address newContractAddress) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateValidatorContractAddress(validatorId *big.Int, newContractAddress common.Address) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateValidatorContractAddress(&_StaderStakeManager.TransactOpts, validatorId, newContractAddress)
}

// UpdateValidatorDelegation is a paid mutator transaction binding the contract method 0xd6de07d0.
//
// Solidity: function updateValidatorDelegation(bool delegation) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateValidatorDelegation(opts *bind.TransactOpts, delegation bool) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateValidatorDelegation", delegation)
}

// UpdateValidatorDelegation is a paid mutator transaction binding the contract method 0xd6de07d0.
//
// Solidity: function updateValidatorDelegation(bool delegation) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateValidatorDelegation(delegation bool) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateValidatorDelegation(&_StaderStakeManager.TransactOpts, delegation)
}

// UpdateValidatorDelegation is a paid mutator transaction binding the contract method 0xd6de07d0.
//
// Solidity: function updateValidatorDelegation(bool delegation) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateValidatorDelegation(delegation bool) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateValidatorDelegation(&_StaderStakeManager.TransactOpts, delegation)
}

// UpdateValidatorState is a paid mutator transaction binding the contract method 0x9ff11500.
//
// Solidity: function updateValidatorState(uint256 validatorId, int256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateValidatorState(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateValidatorState", validatorId, amount)
}

// UpdateValidatorState is a paid mutator transaction binding the contract method 0x9ff11500.
//
// Solidity: function updateValidatorState(uint256 validatorId, int256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateValidatorState(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateValidatorState(&_StaderStakeManager.TransactOpts, validatorId, amount)
}

// UpdateValidatorState is a paid mutator transaction binding the contract method 0x9ff11500.
//
// Solidity: function updateValidatorState(uint256 validatorId, int256 amount) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateValidatorState(validatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateValidatorState(&_StaderStakeManager.TransactOpts, validatorId, amount)
}

// UpdateValidatorThreshold is a paid mutator transaction binding the contract method 0x16827b1b.
//
// Solidity: function updateValidatorThreshold(uint256 newThreshold) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) UpdateValidatorThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "updateValidatorThreshold", newThreshold)
}

// UpdateValidatorThreshold is a paid mutator transaction binding the contract method 0x16827b1b.
//
// Solidity: function updateValidatorThreshold(uint256 newThreshold) returns()
func (_StaderStakeManager *StaderStakeManagerSession) UpdateValidatorThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateValidatorThreshold(&_StaderStakeManager.TransactOpts, newThreshold)
}

// UpdateValidatorThreshold is a paid mutator transaction binding the contract method 0x16827b1b.
//
// Solidity: function updateValidatorThreshold(uint256 newThreshold) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) UpdateValidatorThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.UpdateValidatorThreshold(&_StaderStakeManager.TransactOpts, newThreshold)
}

// WithdrawDelegatorsReward is a paid mutator transaction binding the contract method 0x7ed4b27c.
//
// Solidity: function withdrawDelegatorsReward(uint256 validatorId) returns(uint256)
func (_StaderStakeManager *StaderStakeManagerTransactor) WithdrawDelegatorsReward(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "withdrawDelegatorsReward", validatorId)
}

// WithdrawDelegatorsReward is a paid mutator transaction binding the contract method 0x7ed4b27c.
//
// Solidity: function withdrawDelegatorsReward(uint256 validatorId) returns(uint256)
func (_StaderStakeManager *StaderStakeManagerSession) WithdrawDelegatorsReward(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.WithdrawDelegatorsReward(&_StaderStakeManager.TransactOpts, validatorId)
}

// WithdrawDelegatorsReward is a paid mutator transaction binding the contract method 0x7ed4b27c.
//
// Solidity: function withdrawDelegatorsReward(uint256 validatorId) returns(uint256)
func (_StaderStakeManager *StaderStakeManagerTransactorSession) WithdrawDelegatorsReward(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.WithdrawDelegatorsReward(&_StaderStakeManager.TransactOpts, validatorId)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerTransactor) WithdrawRewards(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.contract.Transact(opts, "withdrawRewards", validatorId)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerSession) WithdrawRewards(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.WithdrawRewards(&_StaderStakeManager.TransactOpts, validatorId)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x9342c8f4.
//
// Solidity: function withdrawRewards(uint256 validatorId) returns()
func (_StaderStakeManager *StaderStakeManagerTransactorSession) WithdrawRewards(validatorId *big.Int) (*types.Transaction, error) {
	return _StaderStakeManager.Contract.WithdrawRewards(&_StaderStakeManager.TransactOpts, validatorId)
}
