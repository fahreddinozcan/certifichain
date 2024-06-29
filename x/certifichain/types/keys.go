package types

const (
	// ModuleName defines the module name
	ModuleName = "certifichain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_certifichain"
)

var (
	ParamsKey = []byte("p_certifichain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
