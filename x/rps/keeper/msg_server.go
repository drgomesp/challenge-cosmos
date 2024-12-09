package rpsKeeper

import (
	"challenge/x/rps/types"
	"context"
	"fmt"
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
	if _, err := s.k.Students.Get(ctx, msg.Id); err == nil {
		return nil, fmt.Errorf("student already exists with id: %s", msg.Id)
	}

	student := types.Student{
		Id:   msg.Id,
		Name: msg.Name,
		Age:  uint64(msg.Age),
	}

	if err := student.Validate(); err != nil {
		return nil, err
	}

	if err := s.k.Students.Set(ctx, msg.Id, student); err != nil {
		return nil, err
	}

	return &types.MsgCreateStudentResponse{Student: &student}, nil
}

func (s *msgServer) DeleteStudent(ctx context.Context, msg *types.MsgDeleteStudent) (*types.MsgDeleteStudentResponse, error) {
	return &types.MsgDeleteStudentResponse{}, nil
}
