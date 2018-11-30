package ibc

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"testing"
)

func TestIBC(t *testing.T) {

	cdc := codec.New()
	cdc.RegisterConcrete(sdkTypes.IBCPacket{}, "test/ibc", nil)
	cdc.RegisterInterface((*sdkTypes.Interface)(nil), nil)

	multiStore, ibcKey := DefaultSetup()
	ctx := csdkTypes.NewContext(multiStore, abci.Header{}, false, log.NewNopLogger())
	ibcKeeper := NewKeeper(ibcKey, cdc)

	ibcPacket1 := TestNewIBCPacket()
	ibcPacketRes1 := ibcKeeper.PostIBCPacket(ctx, ibcPacket1)
	require.Nil(t, ibcPacketRes1)

	getIBCPacket1, err := ibcKeeper.GetIBCPacket(ctx, sdkTypes.EgressKey(ibcPacket1.DestChainID, uint64(0)))
	require.Nil(t, err)
	require.Equal(t, getIBCPacket1, &ibcPacket1)
	require.Equal(t, getIBCPacket1.DestChainID, ibcPacket1.DestChainID)

	err = ibcKeeper.SetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcPacket1.SrcChainID), uint64(1))
	require.Nil(t, err)

	getIBCPacket2, err := ibcKeeper.GetIBCPacket(ctx, sdkTypes.EgressKey(ibcPacket1.DestChainID, uint64(1)))
	require.Nil(t, getIBCPacket2)

	err = ibcKeeper.SetEgressLength(ctx, sdkTypes.EgressLengthKey(ibcPacket1.DestChainID), uint64(1))
	require.Nil(t, err)

	ibcPacket2 := TestNewIBCPacket()
	ibcPacketRes2 := ibcKeeper.PostIBCPacket(ctx, ibcPacket2)
	require.Nil(t, ibcPacketRes2)

	egressLength1, err := ibcKeeper.GetEgressLength(ctx, sdkTypes.EgressLengthKey(ibcPacket1.DestChainID))
	require.Nil(t, err)

	require.Equal(t, uint64(2), egressLength1)

	ingressLength1, err := ibcKeeper.GetIngressLength(ctx, sdkTypes.IngressLengthKey(ibcPacket1.SrcChainID))
	require.Nil(t, err)
	require.Equal(t, uint64(1), ingressLength1)

}
