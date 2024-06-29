package types

const (
	// ModuleName defines the module name
	ModuleName = "certification"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_certification"
)

var (
	ParamsKey = []byte("p_certification")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
