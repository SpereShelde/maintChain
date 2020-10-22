package maintChain

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the maintChain Querier
const (
	QueryRecord = "record"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryRecord:
			return queryRecord(ctx, path[1:], req, keeper)
		default:
			return nil, sdkErrors.Wrap(sdkErrors.ErrUnknownRequest, "unknown maintChain query endpoint")
		}
	}
}

func queryRecord(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	record := keeper.FetchRecord(ctx, path[0])

	if record == nil {
		return nil, sdkErrors.Wrap(sdkErrors.ErrUnknownRequest, "No such record")
	}

	recordJson, err := json.Marshal(record)
	if err != nil {
		panic(fmt.Sprintf("Failed to encode record"))
	}
	return recordJson, nil
}