package keeper

// DONTCOVER

// Although written in msg_server_test.go, it is approached at the keeper level rather than at the msgServer level
// so is not included in the coverage.

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/liquidity/x/liquidity/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the distribution MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// Message server, handler for CreateLiquidityPool msg
func (k msgServer) CreateLiquidityPool(goCtx context.Context, msg *types.MsgCreateLiquidityPool) (*types.MsgCreateLiquidityPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.Keeper.CreateLiquidityPool(ctx, msg)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			//types.EventTypeCreateLiquidityPool,
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.PoolCreatorAddress),
			sdk.NewAttribute(types.AttributeValueLiquidityPoolId, ""),
			sdk.NewAttribute(types.AttributeValueLiquidityPoolTypeIndex, fmt.Sprintf("%d", msg.PoolTypeIndex)),
			sdk.NewAttribute(types.AttributeValueReserveCoinDenoms, ""),
			sdk.NewAttribute(types.AttributeValueReserveAccount, ""),
			sdk.NewAttribute(types.AttributeValuePoolCoinDenom, ""),
			sdk.NewAttribute(types.AttributeValueSwapFeeRate, ""),
			sdk.NewAttribute(types.AttributeValueLiquidityPoolFeeRate, ""),
			sdk.NewAttribute(types.AttributeValueBatchSize, ""),
		),
	)
	return &types.MsgCreateLiquidityPoolResponse{}, nil
}

// Message server, handler for MsgDepositToLiquidityPool
func (k msgServer) DepositToLiquidityPool(goCtx context.Context, msg *types.MsgDepositToLiquidityPool) (*types.MsgDepositToLiquidityPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.Keeper.DepositLiquidityPoolToBatch(ctx, msg)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			//types.EventTypeDepositToLiquidityPoolToBatch,
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DepositorAddress),
			sdk.NewAttribute(types.AttributeValueBatchID, ""),
		),
	)
	return &types.MsgDepositToLiquidityPoolResponse{}, nil
}

// Message server, handler for MsgWithdrawFromLiquidityPool
func (k msgServer) WithdrawFromLiquidityPool(goCtx context.Context, msg *types.MsgWithdrawFromLiquidityPool) (*types.MsgWithdrawFromLiquidityPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.Keeper.WithdrawLiquidityPoolToBatch(ctx, msg)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			//types.EventTypeWithdrrawFromLiquidityPoolToBatch,
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.WithdrawerAddress),
			sdk.NewAttribute(types.AttributeValueBatchID, ""),
		),
	)
	return &types.MsgWithdrawFromLiquidityPoolResponse{}, nil
}

// Message server, handler for MsgSwap
func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if _, err := k.Keeper.SwapLiquidityPoolToBatch(ctx, msg, 0); err != nil {
		return &types.MsgSwapResponse{}, err
	}
	return &types.MsgSwapResponse{}, nil
}
