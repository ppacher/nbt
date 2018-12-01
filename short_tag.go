package nbt

import (
	"encoding/binary"
	"io"
)

// ShortTag is a TagShort named binary tag
type ShortTag struct {
	NamedTag
	Value int16
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
