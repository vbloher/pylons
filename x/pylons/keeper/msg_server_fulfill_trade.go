package keeper

import (
	"context"
	"strconv"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Pylons-tech/pylons/x/pylons/types"
)

func (k msgServer) MatchItemInputsForTrade(ctx sdk.Context, creatorAddr string, itemRefs []types.ItemRef, trade types.Trade) ([]types.Item, error) {
	if len(itemRefs) != len(trade.ItemInputs) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "size mismatch between provided input items and items required by trade")
	}
	matchedInputItems := make([]types.Item, len(trade.ItemInputs))

	// build Item list from inputItemIDs
	inputItemMap := make(map[types.ItemRef]types.Item)
	checkedInputItems := make([]bool, len(itemRefs))

	for i, recipeItemInput := range trade.ItemInputs {
		var err error
		for j, itemRef := range itemRefs {
			if checkedInputItems[j] {
				continue
			}
			inputItem, found := inputItemMap[itemRef]
			if !found {
				inputItem, found = k.GetItem(ctx, itemRef.CookbookID, itemRef.ItemID)
				if !found {
					return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "item with id %v not found", itemRef.ItemID)
				}
				if inputItem.Owner != creatorAddr {
					return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "item with id %s not owned by sender", inputItem.ID)
				}
			}
			inputItemMap[itemRef] = inputItem
			// match
			var ec types.CelEnvCollection
			ec, err = k.NewCelEnvCollectionFromItem(ctx, "", strconv.FormatUint(trade.ID, 10), inputItem)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
			}
			err = recipeItemInput.MatchItem(inputItem, ec)
			if err != nil {
				matchedInputItems[i] = inputItem
				checkedInputItems[j] = true
				break
			}
		}
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot find match for recipe input item ")
		}
	}
	return matchedInputItems, nil
}

func (k msgServer) FulfillTrade(goCtx context.Context, msg *types.MsgFulfillTrade) (*types.MsgFulfillTradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the trade from keeper
	if !k.HasTrade(ctx, msg.ID) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "trade does not exist")
	}
	trade := k.GetTrade(ctx, msg.ID)

	// match msg items to trade itemInputs
	matchedInputItems, err := k.MatchItemInputsForTrade(ctx, msg.Creator, msg.Items, trade)
	if err != nil {
		return nil, err
	}

	// check coinOutput is GTE amount to pay (from flat fees of itemInputs)
	for _, item := range matchedInputItems {
		if !item.Tradeable {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "item with id %v and cookbook id %v cannot be traded", item.ID, item.CookbookID)
		}
	}

	minItemInputsTransferFees := sdk.NewCoins()
	itemInputsTransferFeePermutation, err := types.FindValidPaymentsPermutation(matchedInputItems, trade.CoinOutputs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot use coinOutputs to pay for the items provided")
	}
	for i := range matchedInputItems {
		minItemInputsTransferFees = minItemInputsTransferFees.Add(matchedInputItems[i].TransferFee[itemInputsTransferFeePermutation[i]])
	}

	// check that sender has enough balance to pay coinInputs
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	balance := k.bankKeeper.SpendableCoins(ctx, addr)
	if !balance.IsAllGTE(trade.CoinInputs) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "not enough balance to pay for trade coinInputs")
	}

	outputItems := make([]types.Item, len(trade.ItemOutputs))
	for i, itemRef := range trade.ItemOutputs {
		item, _ := k.GetItem(ctx, itemRef.CookbookID, itemRef.ItemID)
		outputItems[i] = item
	}

	minItemOutputsTransferFees := sdk.NewCoins()
	itemOutputsTransferFeePermutation, err := types.FindValidPaymentsPermutation(outputItems, trade.CoinInputs)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "balance not sufficient to pay coinInputs")
	}
	for i := range matchedInputItems {
		minItemOutputsTransferFees = minItemOutputsTransferFees.Add(outputItems[i].TransferFee[itemOutputsTransferFeePermutation[i]])
	}

	// calculate item "weights" as a relative percentage of the total sum of items transferFees
	inputItemWeights := make([]sdk.Dec, len(matchedInputItems))
	for i, item := range matchedInputItems {
		transferFee := item.TransferFee[itemInputsTransferFeePermutation[i]]
		weight := transferFee.Amount.ToDec().Quo(minItemInputsTransferFees.AmountOf(transferFee.Denom).ToDec())
		inputItemWeights[i] = weight
	}
	outputItemWeights := make([]sdk.Dec, len(outputItems))
	for i, item := range outputItems {
		transferFee := item.TransferFee[itemOutputsTransferFeePermutation[i]]
		weight := transferFee.Amount.ToDec().Quo(minItemOutputsTransferFees.AmountOf(transferFee.Denom).ToDec())
		outputItemWeights[i] = weight
	}

	// calculate the actual payment for that item as coin / weight
	// get the residual to transfer to cookbook owner of that item from this actual payment
	inputChainTotAmt := sdk.NewCoins()
	inputTransferTotAmt := sdk.NewCoins()
	inputCookbookOwnersTotAmtMap := make(map[string]sdk.Coins)
	for i, item := range matchedInputItems {
		baseItemTransferFee := item.TransferFee[itemInputsTransferFeePermutation[i]]
		itemTransferFeeAmt := trade.CoinOutputs.AmountOf(baseItemTransferFee.Denom).ToDec().Mul(inputItemWeights[i]).RoundInt()
		tmpCookbookAmt := sdk.NewCoin(baseItemTransferFee.Denom, itemTransferFeeAmt.ToDec().Mul(item.TradePercentage).RoundInt())
		chainAmt := sdk.NewCoin(baseItemTransferFee.Denom, tmpCookbookAmt.Amount.ToDec().Mul(k.ItemTransferFeePercentage(ctx)).RoundInt())
		cookbookAmt := sdk.NewCoin(baseItemTransferFee.Denom, itemTransferFeeAmt.Sub(chainAmt.Amount))
		transferAmt := sdk.NewCoin(baseItemTransferFee.Denom, itemTransferFeeAmt.Sub(cookbookAmt.Amount).Sub(chainAmt.Amount))
		inputChainTotAmt = inputChainTotAmt.Add(chainAmt)
		inputTransferTotAmt = inputTransferTotAmt.Add(transferAmt)
		inputCookbookOwnersTotAmtMap[item.CookbookID] = inputCookbookOwnersTotAmtMap[item.CookbookID].Add(cookbookAmt)
	}
	outputChainTotAmt := sdk.NewCoins()
	outputTransferTotAmt := sdk.NewCoins()
	outputCookbookOwnersTotAmtMap := make(map[string]sdk.Coins)
	for i, item := range outputItems {
		baseItemTransferFee := item.TransferFee[itemOutputsTransferFeePermutation[i]]
		itemTransferFeeAmt := trade.CoinInputs.AmountOf(baseItemTransferFee.Denom).ToDec().Mul(outputItemWeights[i]).RoundInt()
		tmpCookbookAmt := sdk.NewCoin(baseItemTransferFee.Denom, itemTransferFeeAmt.ToDec().Mul(item.TradePercentage).RoundInt())
		chainAmt := sdk.NewCoin(baseItemTransferFee.Denom, tmpCookbookAmt.Amount.ToDec().Mul(k.ItemTransferFeePercentage(ctx)).RoundInt())
		cookbookAmt := sdk.NewCoin(baseItemTransferFee.Denom, itemTransferFeeAmt.Sub(chainAmt.Amount))
		transferAmt := sdk.NewCoin(baseItemTransferFee.Denom, itemTransferFeeAmt.Sub(cookbookAmt.Amount).Sub(chainAmt.Amount))
		outputChainTotAmt = outputChainTotAmt.Add(chainAmt)
		outputTransferTotAmt = outputTransferTotAmt.Add(transferAmt)
		outputCookbookOwnersTotAmtMap[item.CookbookID] = outputCookbookOwnersTotAmtMap[item.CookbookID].Add(cookbookAmt)
	}

	tradeCreatorAddr, _ := sdk.AccAddressFromBech32(trade.Creator)
	tradeFulfillerAddr, _ := sdk.AccAddressFromBech32(msg.Creator)
	// transfer ownership of items
	for _, item := range matchedInputItems {
		item.Owner = trade.Creator
		k.UpdateItem(ctx, item, tradeFulfillerAddr)
	}
	for _, item := range outputItems {
		item.Owner = msg.Creator
		k.UpdateItem(ctx, item, tradeCreatorAddr)
	}

	// send payments
	err = k.PayFees(ctx, tradeFulfillerAddr, inputChainTotAmt)
	err = k.bankKeeper.SendCoins(ctx, tradeFulfillerAddr, tradeCreatorAddr, inputTransferTotAmt)
	// TODO cb pay

	err = k.PayFees(ctx, tradeCreatorAddr, outputChainTotAmt)
	err = k.bankKeeper.SendCoins(ctx, tradeCreatorAddr, tradeFulfillerAddr, outputTransferTotAmt)
	// TODO cb pay

	// TODO - handle accepting multiple choices of CoinInputs, not just a static one - which one to select is decided as input to the msg

	// TODO - add this clamping
	/*
		if coin.Amount.LT(minTransferFee) {
			coin.Amount = minTransferFee
		} else if coin.Amount.GT(maxTransferFee) {
			coin.Amount = maxTransferFee
		}
	*/

	// TODO EMIT EVENT

	return &types.MsgFulfillTradeResponse{}, nil
}