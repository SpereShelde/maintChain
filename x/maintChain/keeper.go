package maintChain

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	mcTypes "maintChain/x/maintChain/types"
	"time"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	key sdk.StoreKey
	cdc *codec.Codec
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		cdc: cdc,
		key: key,
	}
}

func (keeper Keeper) storeRecord(ctx sdk.Context, record *mcTypes.Record) {
	store := ctx.KVStore(keeper.key)
	key := record.Id
	value := keeper.cdc.MustMarshalJSON(record)
	store.Set([]byte(key), value)
}

func (keeper Keeper) getRecord(ctx sdk.Context, id string) *mcTypes.Record {
	store := ctx.KVStore(keeper.key)
	key := id
	value := store.Get([]byte(key))
	if value == nil {
		return nil
	}

	record := new(mcTypes.Record)
	err := keeper.cdc.UnmarshalJSON(value, record)
	if err != nil {
		panic(fmt.Sprintf("Invalid game stored: %s", err))
	}

	return record
}

func (keeper Keeper) AddRecord(ctx sdk.Context, msg MsgAddRecord) *mcTypes.Record {
	id := uuid.New()
	record := &mcTypes.Record{
		Id: id.String(),
		Time: time.Now(),
		Vin: msg.Vin,
		Org: msg.Org,
		Content: msg.Content,
	}
	keeper.storeRecord(ctx, record)
	return record
}

func (keeper Keeper) FetchRecord(ctx sdk.Context, id string) *mcTypes.Record {
	record := keeper.getRecord(ctx, id)
	//TODO: nil

	return record
}