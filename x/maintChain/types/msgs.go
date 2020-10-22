package maintChain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"time"
)

// MsgSetName defines a SetName message
type MsgStore struct {
	Id      		string
	Time			time.Time
	Vin				string
	Organization	sdk.AccAddress
	Content			string
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgStore(time time.Time, vin string, organization sdk.AccAddress, content string) MsgStore {
	id := uuid.New()
	return MsgStore {
		Id: id.String(),
		Time: time,
		Vin: vin,
		Organization:  organization,
		Content:  content,
	}
}

// Route should return the name of the module
func (msg MsgStore) Route() string { return "maintChain" }

// Type should return the action
func (msg MsgStore) Type() string { return "store"}