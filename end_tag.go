package nbt

import (
	"encoding/json"
	"io"
)

// EndTag is a TagEnd named binary tag
type EndTag struct {
	NamedTag
}

// TagID returns TagEnd
func (tag EndTag) TagID() byte {
	return TagEnd
}

// MarshalJSON returns the JSON representation of the tag
func (tag EndTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

// Read does nothing as end tag does not have a payload
func (tag *EndTag) Read(reader io.Reader) error {
	return nil
}

var _ Tag = new(EndTag)
