package nbt

import (
	"encoding/binary"
	"io"
)

// ByteTag is a TagByte named binary tag
type ByteTag struct {
	NamedTag
	Value byte
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
