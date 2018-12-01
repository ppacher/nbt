package nbt

import "io"

// EndTag is a TagEnd named binary tag
type EndTag struct {
	NamedTag
}

// TagID returns TagEnd
func (tag EndTag) TagID() byte {
	return TagEnd
}

// Read does nothing as end tag does not have a payload
func (tag *EndTag) Read(reader io.Reader) error {
	return nil
}

var _ Tag = new(EndTag)
