package rpsKeeper

import (
	"context"

	"challenge/x/rps/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	k Keeper
}

func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return queryServer{
		k: keeper,
	}
}

func (qs queryServer) GetStudent(ctx context.Context, req *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	return &types.GetStudentResponse{Student: nil}, nil
}
