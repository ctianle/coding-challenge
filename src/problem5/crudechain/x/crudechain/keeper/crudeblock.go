package keeper

import (
	"context"
	"encoding/binary"

	"crudechain/x/crudechain/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetCrudeblockCount get the total number of crudeblock
func (k Keeper) GetCrudeblockCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.CrudeblockCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCrudeblockCount set the total number of crudeblock
func (k Keeper) SetCrudeblockCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.CrudeblockCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCrudeblock appends a crudeblock in the store with a new id and update the count
func (k Keeper) AppendCrudeblock(
	ctx context.Context,
	crudeblock types.Crudeblock,
) uint64 {
	// Create the crudeblock
	count := k.GetCrudeblockCount(ctx)

	// Set the ID of the appended value
	crudeblock.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CrudeblockKey))
	appendedValue := k.cdc.MustMarshal(&crudeblock)
	store.Set(GetCrudeblockIDBytes(crudeblock.Id), appendedValue)

	// Update crudeblock count
	k.SetCrudeblockCount(ctx, count+1)

	return count
}

// SetCrudeblock set a specific crudeblock in the store
func (k Keeper) SetCrudeblock(ctx context.Context, crudeblock types.Crudeblock) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CrudeblockKey))
	b := k.cdc.MustMarshal(&crudeblock)
	store.Set(GetCrudeblockIDBytes(crudeblock.Id), b)
}

// GetCrudeblock returns a crudeblock from its id
func (k Keeper) GetCrudeblock(ctx context.Context, id uint64) (val types.Crudeblock, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CrudeblockKey))
	b := store.Get(GetCrudeblockIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCrudeblock removes a crudeblock from the store
func (k Keeper) RemoveCrudeblock(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CrudeblockKey))
	store.Delete(GetCrudeblockIDBytes(id))
}

// GetAllCrudeblock returns all crudeblock
func (k Keeper) GetAllCrudeblock(ctx context.Context) (list []types.Crudeblock) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CrudeblockKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Crudeblock
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCrudeblockIDBytes returns the byte representation of the ID
func GetCrudeblockIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.CrudeblockKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
