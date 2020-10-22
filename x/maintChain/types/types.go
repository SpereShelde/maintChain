package maintChain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// Record is a struct that contains all the metadata of a maintenance record
type Record struct {
	Id      		string				`json:"id"`
	Time			time.Time			`json:"time"`
	Vin				string				`json:"vin"`
	Organization	sdk.AccAddress		`json:"organization"`
	Content			string				`json:"content"`
}
