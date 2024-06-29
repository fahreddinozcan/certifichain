package types

const (
	// ModuleName defines the module name
	ModuleName = "issuer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_issuer"
)

var (
	ParamsKey = []byte("p_issuer")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
