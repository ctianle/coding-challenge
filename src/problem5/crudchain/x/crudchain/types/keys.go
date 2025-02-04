package types

const (
	// ModuleName defines the module name
	ModuleName = "crudchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_crudchain"
)

var (
	ParamsKey = []byte("p_crudchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
