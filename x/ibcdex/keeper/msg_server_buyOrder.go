package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

func (k msgServer) SendBuyOrder(goCtx context.Context, msg *types.MsgSendBuyOrder) (*types.MsgSendBuyOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Cannot send a order if the pair doesn't exist
	pairIndex := types.OrderBookIndex(msg.Port, msg.ChannelID, msg.AmountDenom, msg.PriceDenom)
	_, found := k.GetBuyOrderBook(ctx, pairIndex)
	if !found {
		return &types.MsgSendBuyOrderResponse{}, errors.New("the pair doesn't exist")
	}

	// Lock the token to send
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return &types.MsgSendBuyOrderResponse{}, err
	}

	// Use SafeBurn to ensure no new native tokens are minted
	if err := k.SafeBurn(
		ctx,
		msg.Port,
		msg.ChannelID,
		sender,
		msg.PriceDenom,
		msg.Amount*msg.Price,
	); err != nil {
		return &types.MsgSendBuyOrderResponse{}, err
	}

	// Save the voucher received on the other chain, to have the ability to resolve it into the original denom
	k.SaveVoucherDenom(ctx, msg.Port, msg.ChannelID, msg.PriceDenom)

	// Construct the packet
	var packet types.BuyOrderPacketData

	packet.Buyer = msg.Sender // <- Manually specify the buyer here
	packet.AmountDenom = msg.AmountDenom
	packet.Amount = msg.Amount
	packet.PriceDenom = msg.PriceDenom
	packet.Price = msg.Price

	// Transmit the packet
	err = k.TransmitBuyOrderPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendBuyOrderResponse{}, nil
}
