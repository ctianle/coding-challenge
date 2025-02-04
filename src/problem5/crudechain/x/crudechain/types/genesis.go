package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CrudeblockList: []Crudeblock{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in crudeblock
	crudeblockIdMap := make(map[uint64]bool)
	crudeblockCount := gs.GetCrudeblockCount()
	for _, elem := range gs.CrudeblockList {
		if _, ok := crudeblockIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for crudeblock")
		}
		if elem.Id >= crudeblockCount {
			return fmt.Errorf("crudeblock id should be lower or equal than the last id")
		}
		crudeblockIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
