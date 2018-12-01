package nbt

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// ByteArray is a TagByteArray named tag
type ByteArray struct {
	NamedTag
	Value []byte
}

// MarshalJSON returns the JSON representation of the tag
func (tag ByteArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(tag.Value)
}

// TagID returns TagByteArray
func (tag ByteArray) TagID() byte {
	return TagByteArray
}

// Read reads the byte array tag from a io.Reader
func (tag *ByteArray) Read(reader io.Reader) error {
	var length int32

	if err := binary.Read(reader, binary.BigEndian, &length); err != nil {
		return err
	}

	var i int32
	for i = 0; i < length; i++ {
		var b byte

		if err := binary.Read(reader, binary.BigEndian, &b); err != nil {
			return err
		}

		tag.Value = append(tag.Value, b)
	}

	return nil
}

var _ Tag = new(ByteArray)
