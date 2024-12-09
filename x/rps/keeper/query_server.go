package rpsKeeper

import (
	"context"
	"cosmossdk.io/collections"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"challenge/x/rps/types"
)

var _ types.QueryServer = &queryServer{}

type queryServer struct {
	k Keeper
}

func NewQueryServer(keeper Keeper) types.QueryServer {
	return &queryServer{
		k: keeper,
	}
}

func (s *queryServer) GetStudent(ctx context.Context, req *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	student, err := s.k.Students.Get(ctx, req.Id)
	if err == nil {
		return &types.GetStudentResponse{Student: &student}, nil
	}

	if errors.Is(err, collections.ErrNotFound) {
		return &types.GetStudentResponse{Student: nil}, nil
	}

	return nil, status.Error(codes.Internal, err.Error())
}
