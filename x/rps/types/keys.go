package types

import "cosmossdk.io/collections"

const ModuleName = "rps"

var (
	StudentsKey = collections.NewPrefix("student/")
)
