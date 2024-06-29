package keeper

import (
	"certifichain/x/certification/types"
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// IssueCertification handles the MsgIssueCertification message
func (ms msgServer) IssueCertification(goCtx context.Context, msg *types.MsgIssueCertification) (*types.MsgIssueCertificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create a new Certification object
	cert := &types.Certification{
		Id:          msg.Id,
		IssuerId:    msg.IssuerId,
		LearnerId:   msg.LearnerId,
		Title:       msg.Title,
		Description: msg.Description,
		IssueDate:   msg.IssueDate,
		ExpiryDate:  msg.ExpiryDate,
		Skills:      msg.Skills,
		Hash:        msg.Hash,
		Creator:     msg.Creator,
	}

	// Call the keeper method to issue the certification
	err := ms.Keeper.IssueCertification(ctx, *cert)
	if err != nil {
		return nil, fmt.Errorf("failed to issue certification: %w", err)
	}

	// Emit an event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeIssueCertification,
			sdk.NewAttribute(types.AttributeKeyCertificationId, cert.Id),
			sdk.NewAttribute(types.AttributeKeyIssuerId, cert.IssuerId),
			sdk.NewAttribute(types.AttributeKeyLearnerId, cert.LearnerId),
		),
	)

	return &types.MsgIssueCertificationResponse{}, nil
}
