package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/utils/slices"
	"github.com/lavanet/lava/x/dualstaking/types"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	"gopkg.in/yaml.v2"
)

// getNextEpoch returns the block of next epoch.
// (all delegate transaction take effect in the subsequent epoch)
func (k Keeper) getNextEpoch(ctx sdk.Context) (uint64, error) {
	block := uint64(ctx.BlockHeight())
	nextEpoch, err := k.epochstorageKeeper.GetNextEpoch(ctx, block)
	if err != nil {
		return 0, utils.LavaFormatError("critical: failed to get next epoch", err,
			utils.Attribute{Key: "block", Value: block},
		)
	}
	return nextEpoch, nil
}

// validateCoins validates that the input amount is valid and non-negative
func validateCoins(amount sdk.Coin) error {
	if !amount.IsValid() {
		return utils.LavaFormatWarning("invalid coins to delegate",
			sdkerrors.ErrInvalidCoins,
			utils.Attribute{Key: "amount", Value: amount},
		)
	} else if amount.IsNegative() {
		return utils.LavaFormatError("failed to transfer coins to module",
			types.ErrBadDelegationAmount,
			utils.Attribute{Key: "amount", Value: amount},
		)
	}
	return nil
}

// increaseDelegation increases the delegation of a delegator to a provider for a
// given chain. It updates the fixation stores for both delegations and delegators,
// and updates the (epochstorage) stake-entry.
func (k Keeper) increaseDelegation(ctx sdk.Context, delegator, provider, chainID string, amount sdk.Coin, nextEpoch uint64) error {
	// get, update and append the delegation entry
	var delegationEntry types.Delegation
	prefix := types.DelegationKey(delegator, provider, chainID)
	found := k.delegationFS.FindEntry(ctx, prefix, nextEpoch, &delegationEntry)
	if !found {
		// new delegation (i.e. not increase of existing one)
		delegationEntry = types.NewDelegation(delegator, provider, chainID)
	}

	delegationEntry.AddAmount(amount)

	err := k.delegationFS.AppendEntry(ctx, prefix, nextEpoch, &delegationEntry)
	if err != nil {
		// append should never fail here
		return utils.LavaFormatError("critical: append delegation entry", err,
			utils.Attribute{Key: "delegator", Value: delegationEntry.Delegator},
			utils.Attribute{Key: "provider", Value: delegationEntry.Provider},
		)
	}

	// get, update and append the delegator entry
	var delegatorEntry types.Delegator
	prefix = types.DelegatorKey(delegator)
	_ = k.delegatorFS.FindEntry(ctx, prefix, nextEpoch, &delegatorEntry)

	delegatorEntry.AddProvider(provider)

	err = k.delegatorFS.AppendEntry(ctx, prefix, nextEpoch, &delegatorEntry)
	if err != nil {
		// append should never fail here
		return utils.LavaFormatError("critical: append delegator entry", err,
			utils.Attribute{Key: "delegator", Value: delegator},
			utils.Attribute{Key: "provider", Value: provider},
		)
	}

	// update the stake entry
	err = k.updateStakeEntry(ctx, provider, chainID, amount)
	if err != nil {
		return err
	}

	return nil
}

// decreaseDelegation decreases the delegation of a delegator to a provider for a
// given chain. It updates the fixation stores for both delegations and delegators,
// and updates the (epochstorage) stake-entry.
func (k Keeper) decreaseDelegation(ctx sdk.Context, delegator, provider, chainID string, amount sdk.Coin, nextEpoch uint64) error {
	// get, update and append the delegation entry
	var delegationEntry types.Delegation
	prefix := types.DelegationKey(delegator, provider, chainID)
	found := k.delegationFS.FindEntry(ctx, prefix, nextEpoch, &delegationEntry)
	if !found {
		return types.ErrDelegationNotFound
	}

	if delegationEntry.Amount.IsLT(amount) {
		return types.ErrInsufficientDelegation
	}

	delegationEntry.SubAmount(amount)

	// if delegation now becomes zero, then remove this entry altogether;
	// otherwise just append the new version (for next epoch).
	if delegationEntry.Amount.IsZero() {
		err := k.delegationFS.DelEntry(ctx, prefix, nextEpoch)
		if err != nil {
			// delete should never fail here
			return utils.LavaFormatError("critical: delete delegation entry", err,
				utils.Attribute{Key: "delegator", Value: delegator},
				utils.Attribute{Key: "provider", Value: provider},
			)
		}
	} else {
		err := k.delegationFS.AppendEntry(ctx, prefix, nextEpoch, &delegationEntry)
		if err != nil {
			// append should never fail here
			return utils.LavaFormatError("failed to update delegation entry", err,
				utils.Attribute{Key: "delegator", Value: delegator},
				utils.Attribute{Key: "provider", Value: provider},
			)
		}
	}

	// get, update and append the delegator entry
	var delegatorEntry types.Delegator
	prefix = types.DelegatorKey(delegator)
	found = k.delegatorFS.FindEntry(ctx, prefix, nextEpoch, &delegatorEntry)
	if !found {
		// we found the delegation above, so the delegator must exist as well
		return utils.LavaFormatError("critical: delegator entry for delegation not found",
			types.ErrDelegationNotFound,
			utils.Attribute{Key: "delegator", Value: delegator},
			utils.Attribute{Key: "provider", Value: provider},
		)
	}

	// if delegation now becomes zero, then remove this provider from the delegator
	// entry; and if the delegator entry becomes entry then remove it altogether.
	// otherwise just append the new version (for next epoch).
	if delegationEntry.Amount.IsZero() {
		delegatorEntry.DelProvider(provider)
		if delegatorEntry.IsEmpty() {
			err := k.delegatorFS.DelEntry(ctx, prefix, nextEpoch)
			if err != nil {
				// delete should never fail here
				return utils.LavaFormatError("critical: delete delegator entry", err,
					utils.Attribute{Key: "delegator", Value: delegator},
					utils.Attribute{Key: "provider", Value: provider},
				)
			}
		}
	} else {
		delegatorEntry.AddProvider(provider)
		err := k.delegatorFS.AppendEntry(ctx, prefix, nextEpoch, &delegatorEntry)
		if err != nil {
			// append should never fail here
			return utils.LavaFormatError("failed to update delegator entry", err,
				utils.Attribute{Key: "delegator", Value: delegator},
				utils.Attribute{Key: "provider", Value: provider},
			)
		}
	}

	// call updateStakeEntry() with the negetive amount
	amount = sdk.NewCoin(amount.Denom, sdk.ZeroInt()).Sub(amount)
	if err := k.updateStakeEntry(ctx, provider, chainID, amount); err != nil {
		return err
	}

	return nil
}

// updateStakeEntry updates the (epochstorage) stake-entry of the provider for the
// given chain. amount can be positive or negative (to increase or decrease).
func (k Keeper) updateStakeEntry(ctx sdk.Context, provider, chainID string, amount sdk.Coin) error {
	providerAddr, err := sdk.AccAddressFromBech32(provider)
	if err != nil {
		// panic:ok: this call was alreadys successful by the caller
		utils.LavaFormatPanic("updateStakeEntry: invalid provider address", err,
			utils.Attribute{Key: "provider", Value: provider},
		)
	}

	stakeEntry, exists, index := k.epochstorageKeeper.GetStakeEntryByAddressCurrent(ctx, chainID, providerAddr)
	if !exists {
		return types.ErrProviderNotStaked
	}

	// sanity check
	if stakeEntry.Address != provider {
		return utils.LavaFormatError("critical: delagate to provider with address mismatch", sdkerrors.ErrInvalidAddress,
			utils.Attribute{Key: "provider", Value: provider},
			utils.Attribute{Key: "address", Value: stakeEntry.Address},
		)
	}

	stakeEntry.DelegateTotal = stakeEntry.DelegateTotal.Add(amount)

	// fail if amount was negative and the result underflowed the existing total
	if stakeEntry.DelegateTotal.IsNegative() {
		return types.ErrInsufficientDelegation
	}

	// if the entry is does not have any stake, it means that it was staked before,
	// and then received delegations, and unstaked - so the remaining delegations
	// are now "orphan".
	// in this state delegators can withdraw delegations but not add new ones. if
	// as a result the total delegatin for this stake-entry reaches zero then the
	// entire stake-entry can be released.

	if stakeEntry.Stake.Amount.IsZero() {
		if amount.IsNegative() {
			if stakeEntry.DelegateTotal.IsZero() {
				// TODO:OL: REMOVE StakeEntry, no longer needed around
				// TODO:OL: epochstorage should not remove unstaked with delegation
			}
		} else {
			return types.ErrProviderNotStaked
		}
	}

	k.epochstorageKeeper.ModifyStakeEntryCurrent(ctx, chainID, stakeEntry, index)

	return nil
}

// Delegate lets a delegator delegate an amount of coins to a provider.
// (effective on next epoch)
func (k Keeper) Delegate(ctx sdk.Context, delegator, provider, chainID string, amount sdk.Coin) error {
	nextEpoch, err := k.getNextEpoch(ctx)
	if err != nil {
		return err
	}

	delegatorAddr, err := sdk.AccAddressFromBech32(delegator)
	if err != nil {
		return utils.LavaFormatWarning("invalid delegator address", err,
			utils.Attribute{Key: "delegator", Value: delegator},
		)
	}

	if _, err = sdk.AccAddressFromBech32(provider); err != nil {
		return utils.LavaFormatWarning("invalid provider address", err,
			utils.Attribute{Key: "provider", Value: provider},
		)
	}

	if err := validateCoins(amount); err != nil {
		return err
	} else if amount.IsZero() {
		return nil
	}

	balance := k.bankKeeper.GetBalance(ctx, delegatorAddr, epochstoragetypes.TokenDenom)
	if balance.IsLT(amount) {
		return utils.LavaFormatWarning("insufficient funds to delegate",
			sdkerrors.ErrInsufficientFunds,
			utils.Attribute{Key: "balance", Value: balance},
			utils.Attribute{Key: "amount", Value: amount},
		)
	}

	err = k.increaseDelegation(ctx, delegator, provider, chainID, amount, nextEpoch)
	if err != nil {
		return utils.LavaFormatWarning("failed to increase delegation", err,
			utils.Attribute{Key: "delegator", Value: delegator},
			utils.Attribute{Key: "provider", Value: provider},
			utils.Attribute{Key: "amount", Value: amount.String()},
		)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, delegatorAddr, types.NotBondedPoolName, sdk.NewCoins(amount))
	if err != nil {
		return utils.LavaFormatError("failed to transfer coins to module", err,
			utils.Attribute{Key: "balance", Value: balance},
			utils.Attribute{Key: "amount", Value: amount},
		)
	}

	details := map[string]string{
		"delegator": delegator,
		"provider":  provider,
		"chainID":   chainID,
		"amount":    amount.String(),
	}

	utils.LogLavaEvent(ctx, k.Logger(ctx), types.DelegateEventName, details, "Delegate")

	return nil
}

// Redelegate lets a delegator transfer its delegation between providers, but
// without the funds being subject to unstakeHoldBlocks witholding period.
// (effective on next epoch)
func (k Keeper) Redelegate(ctx sdk.Context, delegator, from, to, fromChainID, toChainID string, amount sdk.Coin) error {
	nextEpoch, err := k.getNextEpoch(ctx)
	if err != nil {
		return err
	}

	if _, err = sdk.AccAddressFromBech32(delegator); err != nil {
		return utils.LavaFormatWarning("invalid delegator address", err,
			utils.Attribute{Key: "delegator", Value: delegator},
		)
	}

	if _, err = sdk.AccAddressFromBech32(from); err != nil {
		return utils.LavaFormatWarning("invalid from-provider address", err,
			utils.Attribute{Key: "from_provider", Value: from},
		)
	}

	if _, err = sdk.AccAddressFromBech32(to); err != nil {
		return utils.LavaFormatWarning("invalid to-provider address", err,
			utils.Attribute{Key: "to_provider", Value: to},
		)
	}

	if err := validateCoins(amount); err != nil {
		return err
	} else if amount.IsZero() {
		return nil
	}

	err = k.decreaseDelegation(ctx, delegator, from, fromChainID, amount, nextEpoch)
	if err != nil {
		return utils.LavaFormatWarning("failed to decrease delegation", err,
			utils.Attribute{Key: "delegator", Value: delegator},
			utils.Attribute{Key: "provider", Value: from},
			utils.Attribute{Key: "amount", Value: amount.String()},
		)
	}

	err = k.increaseDelegation(ctx, delegator, to, fromChainID, amount, nextEpoch)
	if err != nil {
		return utils.LavaFormatWarning("failed to increase delegation", err,
			utils.Attribute{Key: "delegator", Value: delegator},
			utils.Attribute{Key: "provider", Value: to},
			utils.Attribute{Key: "amount", Value: amount.String()},
		)
	}

	// no need to transfer funds, because they remain in the dualstaking module
	// (specifically in types.NotBondedPoolName).

	details := map[string]string{
		"delegator":     delegator,
		"from_provider": from,
		"to_provider":   to,
		"from_chainID":  fromChainID,
		"to_chainID":    toChainID,
		"amount":        amount.String(),
	}

	utils.LogLavaEvent(ctx, k.Logger(ctx), types.RedelegateEventName, details, "Redelegate")

	return nil
}

// Unbond lets a delegator get its delegated coins back from a provider. The
// delegation ends immediately, but coins are held for unstakeHoldBlocks period
// before released and transferred back to the delegator. The rewards from the
// provider will be updated accordingly (or terminate) from the next epoch.
// (effective on next epoch)
func (k Keeper) Unbond(ctx sdk.Context, delegator, provider, chainID string, amount sdk.Coin) error {
	nextEpoch, err := k.getNextEpoch(ctx)
	if err != nil {
		return err
	}

	if _, err = sdk.AccAddressFromBech32(delegator); err != nil {
		return utils.LavaFormatWarning("invalid delegator address", err,
			utils.Attribute{Key: "delegator", Value: delegator},
		)
	}

	if _, err = sdk.AccAddressFromBech32(provider); err != nil {
		return utils.LavaFormatWarning("invalid provider address", err,
			utils.Attribute{Key: "provider", Value: provider},
		)
	}

	if err := validateCoins(amount); err != nil {
		return err
	} else if amount.IsZero() {
		return nil
	}

	err = k.decreaseDelegation(ctx, delegator, provider, chainID, amount, nextEpoch)
	if err != nil {
		return utils.LavaFormatWarning("failed to decrease delegation", err,
			utils.Attribute{Key: "delegator", Value: delegator},
			utils.Attribute{Key: "provider", Value: provider},
			utils.Attribute{Key: "amount", Value: amount.String()},
		)
	}

	// in unbonding the funds to not return immediately to the delegator; instead
	// they transfer to the BonderPoolName module account, where they are held for
	// the hold period before they are finally released.

	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.NotBondedPoolName, types.BondedPoolName, sdk.NewCoins(amount))
	if err != nil {
		return utils.LavaFormatError("failed to unbond delegated coins", err,
			utils.Attribute{Key: "amount", Value: amount},
		)
	}

	err = k.setUnbondingTimer(ctx, delegator, provider, chainID, amount)
	if err != nil {
		return utils.LavaFormatError("failed to unbond delegated coins", err,
			utils.Attribute{Key: "amount", Value: amount},
		)
	}

	details := map[string]string{
		"delegator": delegator,
		"provider":  provider,
		"chainID":   chainID,
		"amount":    amount.String(),
	}

	utils.LogLavaEvent(ctx, k.Logger(ctx), types.UnbondingEventName, details, "Unbond")

	return nil
}

var space = " "

// encodeForTimer generates timer key unique to a specific delegation; thus it
// must include the delegator, provider, and chainID.
func encodeForTimer(delegator, provider, chainID string) []byte {
	dlen, plen, clen := len(delegator), len(provider), len(chainID)

	encodedKey := make([]byte, dlen+plen+clen+2)

	index := 0
	index += copy(encodedKey[index:], delegator)
	encodedKey[index] = []byte(space)[0]
	index += copy(encodedKey[index+1:], provider) + 1
	encodedKey[index] = []byte(space)[0]
	copy(encodedKey[index+1:], chainID)

	return encodedKey
}

// decodeForTimer extracts the delegation details from a timer key.
func decodeForTimer(encodedKey []byte) (string, string, string) {
	split := strings.Split(string(encodedKey), space)
	if len(split) != 3 {
		return "", "", ""
	}
	delegator, provider, chainID := split[0], split[1], split[2]
	return delegator, provider, chainID
}

// setUnbondingTimer creates a timer for an unbonding operation, that will expire
// in the spec's unbonding-hold-blocks blocks from now. only one unbonding request
// can be placed per block (for a given delegator/provider/chainID combination).
func (k Keeper) setUnbondingTimer(ctx sdk.Context, delegator, provider, chainID string, amount sdk.Coin) error {
	block := uint64(ctx.BlockHeight())
	key := encodeForTimer(delegator, provider, chainID)

	timeout := block + k.getUnbondHoldBlocks(ctx, chainID)

	// the timer key encodes the unique delegator/provider/chainID combination.
	// the timer data holds the amount to be released in the future (marshalled).

	if k.unbondingTS.HasTimerByBlockHeight(ctx, timeout, key) {
		return types.ErrUnbondingInProgress
	}

	data, _ := yaml.Marshal(amount)
	k.unbondingTS.AddTimerByBlockHeight(ctx, timeout, key, data)

	return nil
}

// finalizeUnbonding is called when the unbond hold period terminated; it extracts
// the delegation details from the timer key and data and releases the bonder funds
// back to the delegator.
func (k Keeper) finalizeUnbonding(ctx sdk.Context, key []byte, data []byte) {
	delegator, provider, chainID := decodeForTimer(key)

	attrs := slices.Slice(
		utils.Attribute{Key: "delegator", Value: delegator},
		utils.Attribute{Key: "provider", Value: provider},
		utils.Attribute{Key: "chainID", Value: chainID},
		utils.Attribute{Key: "data", Value: data},
	)

	var amount sdk.Coin
	err := yaml.Unmarshal(data, &amount)
	if err != nil {
		utils.LavaFormatError("critical: finalizeBonding failed to decode", err, attrs...)
	}

	// sanity: delegator address must be valid
	delegatorAddr, err := sdk.AccAddressFromBech32(delegator)
	if err != nil {
		utils.LavaFormatError("critical: finalizeBonding invalid delegator", err, attrs...)
		return
	}

	// sanity: provider address must be valid
	if _, err := sdk.AccAddressFromBech32(provider); err != nil {
		utils.LavaFormatError("critical: finalizeBonding invalid provider", err, attrs...)
		return
	}

	// sanity: saved amount is valid
	if err := validateCoins(amount); err != nil {
		utils.LavaFormatError("critical: finalizeBonding invalid amount", err, attrs...)
		return
	} else if amount.IsZero() || amount.IsNegative() {
		utils.LavaFormatError("critical: finalizeBonding zero amount", err, attrs...)
		return
	}

	// sanity: verify that BondedPool has enough funds
	if k.totalBondedTokens(ctx).LT(amount.Amount) {
		utils.LavaFormatError("critical: finalizeBonding insufficient bonds", err, attrs...)
		return
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.BondedPoolName, delegatorAddr, sdk.NewCoins(amount))
	if err != nil {
		utils.LavaFormatError("critical: finalizeBonding failed to transfer", err, attrs...)
	}

	details := map[string]string{
		"delegator": delegator,
		"provider":  provider,
		"chainID":   chainID,
		"amount":    amount.String(),
	}

	utils.LogLavaEvent(ctx, k.Logger(ctx), types.RefundedEventName, details, "Refunded")
}

// TODO: this is duplicated in x/pairing/keeper/unstaking.go; merge into one call
func (k Keeper) getUnbondHoldBlocks(ctx sdk.Context, chainID string) uint64 {
	_, found, providerType := k.specKeeper.IsSpecFoundAndActive(ctx, chainID)
	if !found {
		utils.LavaFormatError("critical: failed to get spec for chainID",
			fmt.Errorf("unknown chainID"),
			utils.Attribute{Key: "chainID", Value: chainID},
		)
	}

	// note: if spec was not found, the default choice is Spec_dynamic == 0

	block := uint64(ctx.BlockHeight())
	if providerType == spectypes.Spec_static {
		return k.epochstorageKeeper.UnstakeHoldBlocksStatic(ctx, block)
	} else {
		return k.epochstorageKeeper.UnstakeHoldBlocks(ctx, block)
	}

	// NOT REACHED
}
