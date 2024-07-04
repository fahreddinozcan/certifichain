package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"certifichain/x/learner/types"
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

func (k Keeper) CreateLearner(ctx sdk.Context, learner types.Learner) error {
	kvStore := k.storeService.OpenKVStore(ctx)

	key := []byte(learner.Id)
	value, err := k.cdc.Marshal(&learner)
	if err != nil {
		return err
	}

	return kvStore.Set(key, value)
}

func (k Keeper) GetLearner(ctx sdk.Context, id string) (*types.Learner, bool) {
	kvStore := k.storeService.OpenKVStore(ctx)

	key := []byte(id)
	value, err := kvStore.Get(key)
	if err != nil || value == nil {
		return nil, false
	}

	var learner types.Learner
	err = k.cdc.Unmarshal(value, &learner)
	if err != nil {
		return nil, false
	}
	return &learner, true
}

func (k Keeper) UpdateLearner(ctx sdk.Context, learner types.Learner) error {
	_, found := k.GetLearner(ctx, learner.Id)
	if !found {
		return fmt.Errorf("learner not found")
	}

	return k.CreateLearner(ctx, learner) // Reuse CreateLearner to update
}

func (k Keeper) DeleteLearner(ctx sdk.Context, id string) error {
	kvStore := k.storeService.OpenKVStore(ctx)

	key := []byte(id)
	exists, err := kvStore.Has(key)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("learner not found")
	}

	return kvStore.Delete(key)
}

func (k Keeper) ListLearners(ctx sdk.Context) ([]*types.Learner, error) {
	store := k.storeService.OpenKVStore(ctx)

	iterator, err := store.Iterator(nil, nil)
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	var learners []*types.Learner
	for ; iterator.Valid(); iterator.Next() {
		var learner types.Learner
		err = k.cdc.Unmarshal(iterator.Value(), &learner)
		if err != nil {
			return nil, err
		}
		learners = append(learners, &learner)
	}

	return learners, nil
}
