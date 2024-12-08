package types

import (
	"errors"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func DefaultStudents() (Students []Student) {
	return
}

func (g Student) GetStudentAddress() (sdk.AccAddress, error) {
	return getStudentAddress(g.Id)
}

func (g Student) GetStudentName() string {
	return g.Name
}

func (g Student) GetStudentAge() uint64 {
	return g.Age
}

func (g Student) Validate() error {
	// Get Student id
	address, err := g.GetStudentAddress()
	if err != nil {
		return err
	}

	if address.Empty() {
		return errors.New("Address empty")
	}

	if g.Age < 0 {
		return errors.New("Invalid Age")
	}

	return nil
}

func getStudentAddress(address string) (sdk.AccAddress, error) {
	// Validate the address has our prefix (it means the wallet is from out blockchain)
	addr, err := sdk.AccAddressFromBech32(address)
	return addr, sdkerrors.Wrap(err, "Invalid Address")
}
