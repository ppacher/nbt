package nbt

import (
	"encoding/binary"
	"io"
)

// LongTag is TagLong named binary tag
type LongTag struct {
	NamedTag
	Value int64
}

// TagID returns TagLong
func (tag LongTag) TagID() byte {
	return TagLong
}

// Read reads the long tag from an io.Reader
func (tag *LongTag) Read(reader io.Reader) error {
	return binary.Read(reader, binary.BigEndian, &tag.Value)
}

var _ Tag = new(LongTag)
