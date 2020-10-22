package maintChain

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgAddRecord struct {
	Vin     	sdk.AccAddress	`json:"vin"`
	Org     	sdk.AccAddress	`json:"org"`
	Content 	string			`json:"content"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgAddRecord(vin sdk.AccAddress, organization sdk.AccAddress, content string) MsgAddRecord {
	return MsgAddRecord{
		Vin:     vin,
		Org:     organization,
		Content: content,
	}
}

// Route should return the name of the module
func (msg MsgAddRecord) Route() string { return "maintChain" }

// Type should return the action
func (msg MsgAddRecord) Type() string { return "addRecord" }

// ValidateBasic runs stateless checks on the message
func (msg MsgAddRecord) ValidateBasic() error {
	if msg.Org.Empty() {
		return sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, msg.Org.String())
	}
	if msg.Vin.Empty() {
		return sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, msg.Vin.String())
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAddRecord) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgAddRecord) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Org}
}