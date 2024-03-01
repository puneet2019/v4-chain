package events

import (
	v1 "github.com/dydxprotocol/v4-chain/protocol/indexer/protocol/v1"
	satypes "github.com/dydxprotocol/v4-chain/protocol/x/subaccounts/types"
)

// NewDeleveragingEvent creates a DeleveragingEvent representing a deleveraging
// where a liquidated subaccount's position is offset by another subaccount.
// Due to the support of final settlement deleveraging matches, sometimes the
// liquidatedSubaccountId is not actually an account that is liquidatable. More
// specifically, it may be a well-collateralized subaccount with an open position
// in a market with the final settlement status.
func NewDeleveragingEvent(
	liquidatedSubaccountId satypes.SubaccountId,
	offsettingSubaccountId satypes.SubaccountId,
	perpetualId uint32,
	fillAmount satypes.BaseQuantums,
	price satypes.BaseQuantums,
	isBuy bool,
	isFinalSettlement bool,
) *DeleveragingEventV2 {
	indexerLiquidatedSubaccountId := v1.SubaccountIdToIndexerSubaccountId(liquidatedSubaccountId)
	indexerOffsettingSubaccountId := v1.SubaccountIdToIndexerSubaccountId(offsettingSubaccountId)
	return &DeleveragingEventV2{
		Liquidated:        indexerLiquidatedSubaccountId,
		Offsetting:        indexerOffsettingSubaccountId,
		PerpetualId:       perpetualId,
		FillAmount:        fillAmount.ToUint64(),
		Price:             price.ToUint64(),
		IsBuy:             isBuy,
		IsFinalSettlement: isFinalSettlement,
	}
}
