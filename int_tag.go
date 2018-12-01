package nbt

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// IntTag is a TagInt named binary tag
type IntTag struct {
	NamedTag
	Value int32
}

// MarshalJSON returns the JSON representation of the tag
func (tag IntTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(tag.Value)
}

// TagID returns TagInt
func (tag IntTag) TagID() byte {
	return TagInt
}

// Read reads the IntTag from an io.Reader
func (tag *IntTag) Read(reader io.Reader) error {
	return binary.Read(reader, binary.BigEndian, &tag.Value)
}

var _ Tag = new(IntTag)
