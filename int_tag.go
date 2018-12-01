package nbt

import (
	"encoding/binary"
	"io"
)

// IntTag is a TagInt named binary tag
type IntTag struct {
	NamedTag
	Value int32
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
