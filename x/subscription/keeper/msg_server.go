package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(k Keeper) types.MsgServiceServer {
	return &msgServer{k}
}

func (k *msgServer) MsgCancel(c context.Context, msg *types.MsgCancelRequest) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if !subscription.GetStatus().Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.GetID(), subscription.GetStatus())
	}
	if msg.From != subscription.GetAddress() {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.DeleteSubscriptionForExpiryAt(ctx, subscription.GetExpiryAt(), subscription.GetID())

	inactiveDuration := k.InactiveDuration(ctx)
	subscription.SetExpiryAt(ctx.BlockTime().Add(inactiveDuration))
	k.SetSubscriptionForExpiryAt(ctx, subscription.GetExpiryAt(), subscription.GetID())

	subscription.SetStatus(hubtypes.StatusInactivePending)
	subscription.SetStatusAt(ctx.BlockTime())

	k.SetSubscription(ctx, subscription)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateStatus{
			ID:     subscription.GetID(),
			Status: subscription.GetStatus(),
		},
	)

	return &types.MsgCancelResponse{}, nil
}

func (k *msgServer) MsgShare(c context.Context, msg *types.MsgShareRequest) (*types.MsgShareResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if subscription.Type() != types.TypePlan {
		return nil, types.NewErrorInvalidSubscriptionType(subscription.GetID(), subscription.Type())
	}
	if msg.From != subscription.GetAddress() {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	fromQuota, found := k.GetQuota(ctx, subscription.GetID(), fromAddr)
	if !found {
		return nil, types.NewErrorQuotaNotFound(subscription.GetID(), fromAddr)
	}

	toAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	toQuota, found := k.GetQuota(ctx, subscription.GetID(), toAddr)
	if !found {
		toQuota = types.Quota{
			Address:        toAddr.String(),
			AllocatedBytes: sdk.ZeroInt(),
			ConsumedBytes:  sdk.ZeroInt(),
		}

		k.SetSubscriptionForAccount(ctx, toAddr, subscription.GetID())
		ctx.EventManager().EmitTypedEvent(
			&types.EventShare{
				ID:      subscription.GetID(),
				Address: toAddr.String(),
			},
		)
	}

	var (
		allocated = fromQuota.AllocatedBytes.Add(toQuota.AllocatedBytes)
		consumed  = fromQuota.ConsumedBytes.Add(toQuota.ConsumedBytes)
		available = allocated.Sub(consumed)
	)

	if msg.Bytes.GT(available) {
		return nil, types.NewErrorInsufficientBytes(subscription.GetID(), msg.Bytes)
	}

	fromQuota.AllocatedBytes = available.Sub(msg.Bytes)
	if fromQuota.AllocatedBytes.LT(fromQuota.ConsumedBytes) {
		return nil, types.NewErrorInvalidQuota(subscription.GetID(), fromAddr)
	}

	k.SetQuota(ctx, subscription.GetID(), fromQuota)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAllocate{
			ID:      subscription.GetID(),
			Address: fromQuota.Address,
			Bytes:   fromQuota.AllocatedBytes,
		},
	)

	toQuota.AllocatedBytes = msg.Bytes
	if toQuota.AllocatedBytes.LT(toQuota.ConsumedBytes) {
		return nil, types.NewErrorInvalidQuota(subscription.GetID(), toAddr)
	}

	k.SetQuota(ctx, subscription.GetID(), toQuota)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAllocate{
			ID:      subscription.GetID(),
			Address: toQuota.Address,
			Bytes:   toQuota.AllocatedBytes,
		},
	)

	return &types.MsgShareResponse{}, nil
}
