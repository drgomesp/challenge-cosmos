package rpsKeeper

import (
	"context"

	"challenge/x/rps/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		k: keeper,
	}
}

// Implement MsgServer interface
func (ms msgServer) CreateStudent(ctx context.Context, msg *types.MsgCreateStudent) (*types.MsgCreateStudentResponse, error) {
	// TODO: Create the logic to create student in the blockchain
	return &types.MsgCreateStudentResponse{}, nil // Return the response as the interface needs (and proto file specify)
}

func (ms msgServer) DeleteStudent(ctx context.Context, msg *types.MsgDeleteStudent) (*types.MsgDeleteStudentResponse, error) {
	// TODO: Create the logic to delete student in the blockchain
	return &types.MsgDeleteStudentResponse{}, nil
}
