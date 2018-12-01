package nbt

import (
	"encoding/binary"
	"io"
)

// LongArrayTag is a TagLongArray named binary tag
type LongArrayTag struct {
	NamedTag
	Value []int64
}

// TagID returns TagLongArray
func (tag LongArrayTag) TagID() byte {
	return TagLongArray
}

// Read reads the long array tag from an io.Reader
func (tag *LongArrayTag) Read(reader io.Reader) error {
	var length int32

	if err := binary.Read(reader, binary.BigEndian, &length); err != nil {
		return err
	}

	var i int32
	for i = 0; i < length; i++ {
		var b int64

		if err := binary.Read(reader, binary.BigEndian, &b); err != nil {
			return err
		}

		tag.Value = append(tag.Value, b)
	}

	return nil
}

var _ Tag = new(LongArrayTag)
