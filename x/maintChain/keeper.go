package maintChain

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	mcTypes "maintChain/x/maintChain/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	key sdk.StoreKey
	cdc *codec.Codec
}

//func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
//	return Keeper{
//		cdc: cdc,
//		key: key,
//	}
//}

func (k Keeper) storeRecord(ctx sdk.Context, record *mcTypes.Record) {
	store := ctx.KVStore(k.key)
	key := record.Id
	value := k.cdc.MustMarshalJSON(record)
	store.Set([]byte(key), value)
}

func (k Keeper) getRecord(ctx sdk.Context, id string) *mcTypes.Record {
	store := ctx.KVStore(k.key)
	key := id
	value := store.Get([]byte(key))
	if value == nil {
		return nil
	}

	record := new(mcTypes.Record)
	err := k.cdc.UnmarshalJSON(value, record)
	if err != nil {
		panic(fmt.Sprintf("Invalid game stored: %s", err))
	}

	return record
}