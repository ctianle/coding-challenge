package types

const (
	// ModuleName defines the module name
	ModuleName = "crudechain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_crudechain"
)

var (
	ParamsKey = []byte("p_crudechain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
