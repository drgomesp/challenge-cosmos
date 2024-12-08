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

func (qs queryServer) GetStudent(ctx context.Context, req *types.QueryGetStudentRequest) (*types.QueryGetStudentResponse, error) {
	// TODO: Create function to get student by Id
	// Return response
	return &types.QueryGetStudentResponse{Student: nil}, nil
}
