package bbnclient

import (
	"encoding/hex"

	checkpointtypes "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	stakingtypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
)

// StakingParams represents the staking parameters of the BBN chain
// Reference: https://github.com/babylonlabs-io/babylon/blob/main/proto/babylon/btcstaking/v1/params.proto
type StakingParams struct {
	CovenantPks                  []string `bson:"covenant_pks"`
	CovenantQuorum               uint32   `bson:"covenant_quorum"`
	MinStakingValueSat           int64    `bson:"min_staking_value_sat"`
	MaxStakingValueSat           int64    `bson:"max_staking_value_sat"`
	MinStakingTimeBlocks         uint32   `bson:"min_staking_time_blocks"`
	MaxStakingTimeBlocks         uint32   `bson:"max_staking_time_blocks"`
	SlashingPkScript             string   `bson:"slashing_pk_script"`
	MinSlashingTxFeeSat          int64    `bson:"min_slashing_tx_fee_sat"`
	SlashingRate                 string   `bson:"slashing_rate"`
	MinUnbondingTimeBlocks       uint32   `bson:"min_unbonding_time_blocks"`
	UnbondingFeeSat              int64    `bson:"unbonding_fee_sat"`
	MinCommissionRate            string   `bson:"min_commission_rate"`
	MaxActiveFinalityProviders   uint32   `bson:"max_active_finality_providers"`
	DelegationCreationBaseGasFee uint64   `bson:"delegation_creation_base_gas_fee"`
}

type CheckpointParams struct {
	BtcConfirmationDepth          uint64 `bson:"btc_confirmation_depth"`
	CheckpointFinalizationTimeout uint64 `bson:"checkpoint_finalization_timeout"`
	CheckpointTag                 string `bson:"checkpoint_tag"`
}

func FromBbnStakingParams(params stakingtypes.Params) *StakingParams {
	return &StakingParams{
		CovenantPks:                  params.CovenantPksHex(),
		CovenantQuorum:               params.CovenantQuorum,
		MinStakingValueSat:           params.MinStakingValueSat,
		MaxStakingValueSat:           params.MaxStakingValueSat,
		MinStakingTimeBlocks:         params.MinStakingTimeBlocks,
		MaxStakingTimeBlocks:         params.MaxStakingTimeBlocks,
		SlashingPkScript:             hex.EncodeToString(params.SlashingPkScript),
		MinSlashingTxFeeSat:          params.MinSlashingTxFeeSat,
		SlashingRate:                 params.SlashingRate.String(),
		MinUnbondingTimeBlocks:       params.MinUnbondingTimeBlocks,
		UnbondingFeeSat:              params.UnbondingFeeSat,
		MinCommissionRate:            params.MinCommissionRate.String(),
		MaxActiveFinalityProviders:   params.MaxActiveFinalityProviders,
		DelegationCreationBaseGasFee: params.DelegationCreationBaseGasFee,
	}
}

func FromBbnCheckpointParams(params checkpointtypes.Params) *CheckpointParams {
	return &CheckpointParams{
		BtcConfirmationDepth:          params.BtcConfirmationDepth,
		CheckpointFinalizationTimeout: params.CheckpointFinalizationTimeout,
		CheckpointTag:                 params.CheckpointTag,
	}
}