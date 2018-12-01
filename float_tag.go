package nbt

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// FloatTag is a TagFloat named binary tag
type FloatTag struct {
	NamedTag
	Value float32
}

// MarshalJSON returns the JSON representation of the tag
func (tag FloatTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(tag.Value)
}

// TagID returns TagFloat
func (tag FloatTag) TagID() byte {
	return TagFloat
}

// Read reads the float tag from an io.Reader
func (tag *FloatTag) Read(reader io.Reader) error {
	return binary.Read(reader, binary.BigEndian, &tag.Value)
}

var _ Tag = new(FloatTag)
