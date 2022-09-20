package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ResourceMapKeyPrefix is the prefix to retrieve all ResourceMap
	ResourceMapKeyPrefix = "ResourceMap/value/"
)

// ResourceMapKey returns the store key to retrieve a ResourceMap from the index fields
func ResourceMapKey(
	originator string,
	origResId string,
) []byte {
	var key []byte

	originatorBytes := []byte(originator)
	key = append(key, originatorBytes...)
	key = append(key, []byte("/")...)

	origResIdBytes := []byte(origResId)
	key = append(key, origResIdBytes...)
	key = append(key, []byte("/")...)

	return key
}

func ResourceMapKeyOf(resource *Resource) []byte {
	return ResourceMapKey(resource.Originator, resource.OrigResId)
}
