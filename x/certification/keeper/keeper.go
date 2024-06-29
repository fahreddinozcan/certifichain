package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"certifichain/x/certification/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) IssueCertification(ctx sdk.Context, cert types.Certification) error {
	kvStore := k.storeService.OpenKVStore(ctx)

	key := []byte(cert.Id)
	value := k.cdc.MustMarshal(&cert)

	return kvStore.Set(key, value)
}

func (k Keeper) GetCertification(ctx sdk.Context, id string) (types.Certification, bool) {
	kvStore := k.storeService.OpenKVStore(ctx)

	key := []byte(id)
	value, err := kvStore.Get(key)
	if err != nil || value == nil {
		return types.Certification{}, false
	}

	var cert types.Certification
	k.cdc.MustUnmarshal(value, &cert)
	return cert, true
}

func (k Keeper) ListCertifications(ctx sdk.Context) ([]types.Certification, error) {
	kvStore := k.storeService.OpenKVStore(ctx)

	iterator, err := kvStore.Iterator(nil, nil)
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	var certifications []types.Certification
	for ; iterator.Valid(); iterator.Next() {
		var cert types.Certification
		k.cdc.MustUnmarshal(iterator.Value(), &cert)
		certifications = append(certifications, cert)
	}

	return certifications, nil
}
