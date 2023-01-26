package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type resourceKey struct {
	originator string
	origResId  string
}

type ResourceKey interface {
	GetOriginator() (originator string)
	GetOriginatorAddress() (originator sdk.AccAddress)
	GetOrigResKey() (origResKey string)
}

func NewResourceKeyForDelete(deleteMsg *MsgDeleteResource) (ResourceKey, error) {
	return NewResourceKey(deleteMsg.Originator, deleteMsg.OrigResId)
}
func NewResourceKey(originator string, origResId string) (ResourceKey, error) {
	err := checkOriginator(originator)
	if err != nil {
		return nil, err
	}
	return &resourceKey{
		originator: originator,
		origResId:  origResId,
	}, nil
}

func checkOriginator(originator string) error {
	_, err := sdk.AccAddressFromBech32(originator)
	return err
}

func (m resourceKey) GetOriginator() (originator string) {
	return m.originator
}

func (m resourceKey) GetOriginatorAddress() (originator sdk.AccAddress) {
	originator, _ = sdk.AccAddressFromBech32(m.originator)
	return originator
}

func (m resourceKey) GetOrigResKey() (origResKey string) {
	return m.origResId
}
