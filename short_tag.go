package nbt

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// ShortTag is a TagShort named binary tag
type ShortTag struct {
	NamedTag
	Value int16
}

// MarshalJSON returns the JSON representation of the tag
func (tag ShortTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(tag.Value)
}

// TagID returns TagShort
func (tag ShortTag) TagID() byte {
	return TagShort
}

// Read reads the short tag from an io.Reader
func (tag *ShortTag) Read(reader io.Reader) error {
	return binary.Read(reader, binary.BigEndian, &tag.Value)
}

var _ Tag = new(ShortTag)
