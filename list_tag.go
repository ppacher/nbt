package nbt

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// ListTag is a TagList named binary tag
type ListTag struct {
	NamedTag
	typeID byte
	Tags   []Tag
}

// MarshalJSON returns the JSON representation of the tag
func (tag ListTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(tag.Tags)
}

// TagID returns TagList
func (tag ListTag) TagID() byte {
	return TagList
}

// Read reads the list tag from an io.Reader
func (tag *ListTag) Read(reader io.Reader) error {
	var tagID byte

	if err := binary.Read(reader, binary.BigEndian, &tagID); err != nil {
		return err
	}

	var length int32
	if err := binary.Read(reader, binary.BigEndian, &length); err != nil {
		return err
	}

	var i int32

	for i = 0; i < length; i++ {
		t := CreateTag(tagID)

		if err := t.Read(reader); err != nil {
			return err
		}

		tag.Tags = append(tag.Tags, t)
	}

	return nil
}

var _ = new(ListTag)
