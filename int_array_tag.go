package nbt

import (
	"encoding/binary"
	"io"
)

// IntArrayTag is a TagIntArray named binary tag
type IntArrayTag struct {
	NamedTag
	Value []int32
}

// TagID returns TagIntArray
func (tag IntArrayTag) TagID() byte {
	return TagIntArray
}

// Read reads the int array tag from an io.Reader
func (tag *IntArrayTag) Read(reader io.Reader) error {
	var length int32

	if err := binary.Read(reader, binary.BigEndian, &length); err != nil {
		return err
	}

	var i int32
	for i = 0; i < length; i++ {
		var b int32

		if err := binary.Read(reader, binary.BigEndian, &b); err != nil {
			return err
		}

		tag.Value = append(tag.Value, b)
	}

	return nil
}

var _ Tag = new(IntArrayTag)
