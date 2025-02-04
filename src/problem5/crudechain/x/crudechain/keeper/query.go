package keeper

import (
	"crudechain/x/crudechain/types"
)

var _ types.QueryServer = Keeper{}
