package nbt

import "io"

// CompoundTag is a TagCompound named binary tag
type CompoundTag struct {
	NamedTag
	Tags map[string]Tag
}

// TagID returns TagCompound
func (tag CompoundTag) TagID() byte {
	return TagCompound
}

// Read reads the compound tag from an io.Reader
func (tag *CompoundTag) Read(reader io.Reader) error {
	if tag.Tags == nil {
		tag.Tags = make(map[string]Tag)
	}

	for {
		t, err := ReadNamedTag(reader)
		if err != nil {
			return err
		}

		if t.TagID() == TagEnd {
			break
		}

		tag.Tags[tag.Name()] = t
	}

	return nil
}

var _ Tag = new(CompoundTag)