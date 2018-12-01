package nbt

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// ByteTag is a TagByte named binary tag
type ByteTag struct {
	NamedTag
	Value byte
}

// MarshalJSON returns the JSON representation of the tag
func (tag ByteTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(tag.Value)
}

// TagID returns TagByte
func (tag ByteTag) TagID() byte {
	return TagByte
}

// Read reads the byte tag from an io.Reader
func (tag *ByteTag) Read(reader io.Reader) error {
	return binary.Read(reader, binary.BigEndian, &tag.Value)
}

var _ Tag = new(ByteTag)
