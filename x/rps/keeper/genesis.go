package rpsKeeper

import (
	"context"

	"challenge/x/rps/types"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) error {

	// Set the genesis games into the state
	for _, student := range data.Students {
		err := k.Students.Set(ctx, student.Id, student)
		if err != nil {
			return nil
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var students []types.Student
	// The fuction walk just iterate the Games map
	err := k.Students.Walk(ctx, nil, func(id string, student types.Student) (stop bool, err error) {
		students = append(students, student)
		return false, nil
	})

	if err != nil {
		return nil, err
	}

	return &types.GenesisState{
		Students: students,
	}, nil
}
