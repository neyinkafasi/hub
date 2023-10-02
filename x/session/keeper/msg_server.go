package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

func (k *msgServer) MsgStart(c context.Context, msg *types.MsgStartRequest) (*types.MsgStartResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	subscription, found := k.GetSubscription(ctx, msg.Id)
	if !found {
		return nil, types.ErrorSubscriptionDoesNotExit
	}
	if !subscription.Status.Equal(hubtypes.StatusActive) {
		return nil, types.ErrorInvalidSubscriptionStatus
	}

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Node)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.ErrorNodeDoesNotExist
	}
	if !node.Status.Equal(hubtypes.StatusActive) {
		return nil, types.ErrorInvalidNodeStatus
	}

	if subscription.Plan == 0 {
		if node.Address != subscription.Node {
			return nil, types.ErrorNodeAddressMismatch
		}
	} else {
		if !k.HasNodeForPlan(ctx, subscription.Plan, nodeAddr) {
			return nil, types.ErrorNodeDoesNotExistForPlan
		}
	}

	quota, found := k.GetQuota(ctx, subscription.Id, fromAddr)
	if !found {
		return nil, types.ErrorQuotaDoesNotExist
	}
	if quota.Consumed.GTE(quota.Allocated) {
		return nil, types.ErrorNotEnoughQuota
	}

	items := k.GetActiveSessionsForAddress(ctx, fromAddr, 0, 1)
	if len(items) > 0 {
		return nil, types.ErrorDuplicateSession
	}

	var (
		count            = k.GetCount(ctx)
		inactiveDuration = k.InactiveDuration(ctx)
		session          = types.Session{
			Id:           count + 1,
			Subscription: subscription.Id,
			Node:         node.Address,
			Address:      msg.From,
			Duration:     0,
			Bandwidth:    hubtypes.NewBandwidthFromInt64(0, 0),
			Status:       hubtypes.StatusActive,
			StatusAt:     ctx.BlockTime(),
		}
	)

	k.SetCount(ctx, count+1)
	k.SetSession(ctx, session)
	k.SetActiveSessionForAddress(ctx, fromAddr, session.Id)
	k.SetInactiveSessionAt(ctx, session.StatusAt.Add(inactiveDuration), session.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventStartSession{
			From:         msg.From,
			Id:           session.Id,
			Subscription: session.Subscription,
			Node:         session.Node,
		},
	)

	return &types.MsgStartResponse{}, nil
}

func (k *msgServer) MsgUpdate(c context.Context, msg *types.MsgUpdateRequest) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	session, found := k.GetSession(ctx, msg.Proof.Id)
	if !found {
		return nil, types.ErrorSessionDoesNotExist
	}
	if session.Status.Equal(hubtypes.StatusInactive) {
		return nil, types.ErrorInvalidSessionStatus
	}
	if msg.From != session.Node {
		return nil, types.ErrorUnauthorized
	}

	if k.ProofVerificationEnabled(ctx) {
		accAddr := session.GetAddress()
		if err := k.VerifyProof(ctx, accAddr, msg.Proof, msg.Signature); err != nil {
			return nil, types.ErrorFailedToVerifyProof
		}
	}

	inactiveDuration := k.InactiveDuration(ctx)
	k.DeleteInactiveSessionAt(ctx, session.StatusAt.Add(inactiveDuration), session.Id)

	session.Duration = msg.Proof.Duration
	session.Bandwidth = msg.Proof.Bandwidth
	session.StatusAt = ctx.BlockTime()

	k.SetSession(ctx, session)
	k.SetInactiveSessionAt(ctx, session.StatusAt.Add(inactiveDuration), session.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateSession{
			From:         msg.From,
			Id:           session.Id,
			Subscription: session.Subscription,
			Node:         session.Node,
			Address:      session.Address,
			Duration:     session.Duration,
			Bandwidth:    session.Bandwidth,
		},
	)

	return &types.MsgUpdateResponse{}, nil
}

func (k *msgServer) MsgEnd(c context.Context, msg *types.MsgEndRequest) (*types.MsgEndResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	session, found := k.GetSession(ctx, msg.Id)
	if !found {
		return nil, types.ErrorSessionDoesNotExist
	}
	if !session.Status.Equal(hubtypes.StatusActive) {
		return nil, types.ErrorInvalidSessionStatus
	}
	if msg.From != session.Address {
		return nil, types.ErrorUnauthorized
	}

	inactiveDuration := k.InactiveDuration(ctx)
	k.DeleteActiveSessionForAddress(ctx, fromAddr, session.Id)
	k.DeleteInactiveSessionAt(ctx, session.StatusAt.Add(inactiveDuration), session.Id)

	session.Status = hubtypes.StatusInactivePending
	session.StatusAt = ctx.BlockTime()

	k.SetSession(ctx, session)
	k.SetInactiveSessionForAddress(ctx, fromAddr, session.Id)
	k.SetInactiveSessionAt(ctx, session.StatusAt.Add(inactiveDuration), session.Id)
	ctx.EventManager().EmitTypedEvent(
		&types.EventEndSession{
			From:         msg.From,
			Id:           session.Id,
			Subscription: session.Subscription,
			Node:         session.Node,
			Status:       session.Status,
		},
	)

	return &types.MsgEndResponse{}, nil
}
