package nbt

// NamedTag is a named binary tag
type NamedTag struct {
	name string
}

// SetName sets the name of the tag
func (tag *NamedTag) SetName(name string) {
	tag.name = name
}

// Name returns the name of the tag
func (tag NamedTag) Name() string {
	return tag.name
}
