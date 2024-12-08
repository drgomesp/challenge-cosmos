package rpsKeeper

import (
	"context"

	"challenge/x/rps/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = &msgServer{}

func NewMsgServer(keeper Keeper) types.MsgServer {
	return &msgServer{
		k: keeper,
	}
}

func (s *msgServer) CreateStudent(ctx context.Context, msg *types.MsgCreateStudent) (*types.MsgCreateStudentResponse, error) {
	return &types.MsgCreateStudentResponse{}, nil
}

func (s *msgServer) DeleteStudent(ctx context.Context, msg *types.MsgDeleteStudent) (*types.MsgDeleteStudentResponse, error) {
	return &types.MsgDeleteStudentResponse{}, nil
}
