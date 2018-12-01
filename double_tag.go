package nbt

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// DoubleTag is a TagDouble named binary tag
type DoubleTag struct {
	NamedTag
	Value float64
}

// MarshalJSON returns the JSON representation of the tag
func (tag DoubleTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(tag.Value)
}

// TagID returns TagDouble
func (tag DoubleTag) TagID() byte {
	return TagDouble
}

// Read reads the double tag from an io.Reader
func (tag *DoubleTag) Read(reader io.Reader) error {
	return binary.Read(reader, binary.BigEndian, &tag.Value)
}

var _ Tag = new(DoubleTag)
