package nbt

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

const (
	TagEnd       = 0
	TagByte      = 1
	TagShort     = 2
	TagInt       = 3
	TagLong      = 4
	TagFloat     = 5
	TagDouble    = 6
	TagByteArray = 7
	TagString    = 8
	TagList      = 9
	TagCompound  = 10
	TagIntArray  = 11
	TagLongArray = 12
)

// Tag reprensents a named tag in the NBT format (Named Binary Tag)
// as specified by Mojang AB
type Tag interface {
	// Tags must support being marshaled into JSON
	json.Marshaler

	// Read reads the tag from a given io.Reader
	Read(io.Reader) error

	// SetName sets the name of the tag
	SetName(name string)

	// Name returns the name of the tag
	Name() string

	// TagID returns the ID of the tag as a byte
	TagID() byte
}

// CreateTag creates a new Tag of the given type
func CreateTag(tagID byte) Tag {
	var t Tag

	switch tagID {
	case TagByteArray:
		t = new(ByteArray)
		break
	case TagByte:
		t = new(ByteTag)
		break
	case TagDouble:
		t = new(DoubleTag)
		break
	case TagEnd:
		t = new(EndTag)
		break
	case TagFloat:
		t = new(FloatTag)
		break
	case TagIntArray:
		t = new(IntArrayTag)
		break
	case TagInt:
		t = new(IntTag)
		break
	case TagList:
		t = new(ListTag)
		break
	case TagLongArray:
		t = new(LongArrayTag)
		break
	case TagLong:
		t = new(LongTag)
		break
	case TagShort:
		t = new(ShortTag)
		break
	case TagString:
		t = new(StringTag)
		break
	case TagCompound:
		t = new(CompoundTag)
		break
	default:
		panic(fmt.Sprintf("Invalid tag id: %d", tagID))
	}

	return t
}

// ReadNamedTag reads a named binary tag from the given reader
func ReadNamedTag(reader io.Reader) (Tag, error) {
	var tagID byte

	if err := binary.Read(reader, binary.BigEndian, &tagID); err != nil {
		return nil, err
	}

	name, err := readUTF8String(reader)
	if err != nil {
		return nil, err
	}

	tag := CreateTag(tagID)

	//fmt.Printf("Reading tagID %d with name %s and type %+v\n", tagID, name, tag)

	tag.SetName(name)
	if err := tag.Read(reader); err != nil {
		return nil, err
	}

	return tag, nil
}
