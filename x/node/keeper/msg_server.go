package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
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

func (k *msgServer) MsgRegister(c context.Context, msg *types.MsgRegisterRequest) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if msg.GigabytePrices != nil {
		if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
			return nil, types.NewErrorInvalidGigabytePrices(msg.GigabytePrices)
		}
	}
	if msg.HourlyPrices != nil {
		if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
			return nil, types.NewErrorInvalidGigabytePrices(msg.HourlyPrices)
		}
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	nodeAddr := hubtypes.NodeAddress(accAddr.Bytes())
	if k.HasNode(ctx, nodeAddr) {
		return nil, types.NewErrorDuplicateNode(nodeAddr)
	}

	deposit := k.Deposit(ctx)
	if err = k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	node := types.Node{
		Address:        nodeAddr.String(),
		GigabytePrices: msg.GigabytePrices,
		HourlyPrices:   msg.HourlyPrices,
		RemoteURL:      msg.RemoteURL,
		Status:         hubtypes.StatusInactive,
		StatusAt:       ctx.BlockTime(),
	}

	k.SetInactiveNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventRegister{
			Address: node.Address,
		},
	)

	return &types.MsgRegisterResponse{}, nil
}

func (k *msgServer) MsgUpdateDetails(c context.Context, msg *types.MsgUpdateDetailsRequest) (*types.MsgUpdateDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	if msg.GigabytePrices != nil {
		if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
			return nil, types.NewErrorInvalidGigabytePrices(msg.GigabytePrices)
		}

		node.GigabytePrices = msg.GigabytePrices
	}
	if msg.HourlyPrices != nil {
		if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
			return nil, types.NewErrorInvalidHourlyPrices(msg.HourlyPrices)
		}

		node.HourlyPrices = msg.HourlyPrices
	}
	if msg.RemoteURL != "" {
		node.RemoteURL = msg.RemoteURL
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateDetails{
			Address: node.Address,
		},
	)

	return &types.MsgUpdateDetailsResponse{}, nil
}

func (k *msgServer) MsgUpdateStatus(c context.Context, msg *types.MsgUpdateStatusRequest) (*types.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	inactiveDuration := k.InactiveDuration(ctx)
	if node.Status.Equal(hubtypes.StatusActive) {
		if msg.Status.Equal(hubtypes.StatusInactive) {
			k.DeleteActiveNode(ctx, nodeAddr)
		}

		inactiveAt := node.StatusAt.Add(inactiveDuration)
		k.DeleteInactiveNodeAt(ctx, inactiveAt, nodeAddr)
	} else {
		if msg.Status.Equal(hubtypes.StatusActive) {
			k.DeleteInactiveNode(ctx, nodeAddr)
		}
	}

	node.Status = msg.Status
	node.StatusAt = ctx.BlockTime()

	if node.Status.Equal(hubtypes.StatusActive) {
		inactiveAt := node.StatusAt.Add(inactiveDuration)
		k.SetInactiveNodeAt(ctx, inactiveAt, nodeAddr)
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateStatus{
			Address: node.Address,
			Status:  node.Status,
		},
	)

	return &types.MsgUpdateStatusResponse{}, nil
}

func (k *msgServer) MsgSubscribe(c context.Context, msg *types.MsgSubscribeRequest) (*types.MsgSubscribeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	nodeAddr, err := hubtypes.NodeAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	lease := types.Lease{
		Bytes: hubtypes.Gigabyte.MulRaw(msg.Gigabytes),
		Hours: msg.Hours,
	}

	if msg.Hours != 0 {
		if !k.IsValidLeaseHours(ctx, msg.Hours) {
			return nil, types.NewErrorInvalidLeaseHours(msg.Hours)
		}

		lease.Price, found = node.HourlyPrice(msg.Denom)
		if !found {
			return nil, types.NewErrorHourlyPriceNotFound(msg.Denom)
		}
	} else if msg.Gigabytes != 0 {
		if !k.IsValidLeaseGigabytes(ctx, msg.Hours) {
			return nil, types.NewErrorInvalidLeaseGigabytes(msg.Gigabytes)
		}

		lease.Price, found = node.GigabytePrice(msg.Denom)
		if !found {
			return nil, types.NewErrorHourlyPriceNotFound(msg.Denom)
		}
	} else {
		return nil, nil
	}

	k.SetLease(ctx, lease)
	k.SetLeaseForAccount(ctx, accAddr, lease.ID)
	k.SetLeaseForNode(ctx, nodeAddr, lease.ID)
	ctx.EventManager().EmitTypedEvent(
		&types.EventLease{
			ID:     lease.ID,
			Lessor: nodeAddr.String(),
			Lessee: accAddr.String(),
		},
	)

	return &types.MsgSubscribeResponse{}, nil
}
