package types

const (
	// ModuleName defines the module name
	ModuleName = "learner"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_learner"
)

var (
	ParamsKey = []byte("p_learner")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
