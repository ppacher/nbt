package nbt

import (
	"encoding/binary"
	"io"
)

// StringTag is a TagString named binary tag
type StringTag struct {
	NamedTag
	Value string
}

// TagID return TagString
func (tag StringTag) TagID() byte {
	return TagString
}

// Read reads the string tag from an io.Reader
func (tag *StringTag) Read(reader io.Reader) error {
	val, err := readUTF8String(reader)
	tag.Value = val
	return err
}

func readUTF8String(reader io.Reader) (string, error) {
	var length int16
	var value string

	if err := binary.Read(reader, binary.BigEndian, &length); err != nil {
		return "", err
	}

	var i int16
	for i = 0; i < length; i++ {
		var c byte

		if err := binary.Read(reader, binary.BigEndian, &c); err != nil {
			return "", err
		}

		value += string(c)
	}

	return value, nil
}

var _ Tag = new(StringTag)
