package keeper

import (
	"crudchain/x/crudchain/types"
)

var _ types.QueryServer = Keeper{}
