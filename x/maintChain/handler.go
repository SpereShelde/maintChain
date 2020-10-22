package maintChain

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "mcTypes" type messages.
func NewHandler(keeper Keeper) func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgAddRecord:
			return handleMsgAddRecord(ctx, keeper, msg)
		default:
			return handleUnknownMsg()
		}
	}
}

func handleUnknownMsg() sdk.Result {
	return sdk.Result{ Data: nil }
}


func handleMsgAddRecord(ctx sdk.Context, keeper Keeper, msg MsgAddRecord) sdk.Result {
	record := keeper.AddRecord(ctx, msg)
	recordData, err := json.Marshal(record)
	if err != nil {
		panic(err)
	}
	return sdk.Result{ Data: recordData }
}
