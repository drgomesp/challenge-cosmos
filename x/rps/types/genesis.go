package types

import "errors"

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Students: DefaultStudents(), // Return the genesis value for a game
	}
}

// Validate performs basic genesis state validation returning an error upon any
func (gs *GenesisState) Validate() error {
	// Validate the genesis students
	unique := make(map[string]bool)
	for _, game := range gs.Students {
		// Validate if the game number already exists
		_, ok := unique[game.Id]
		if ok {
			return errors.New("Id duplicated")
		}

		// Validate each game
		err := game.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
