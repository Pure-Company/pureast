package proto

import (
	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
)

// ProtoMessage represents a protobuf message (pure data!)
type ProtoMessage struct {
	Name    string
	Fields  []ProtoField
	Comment string
}

// ProtoField represents a message field (pure data!)
type ProtoField struct {
	Name     string
	Type     string
	Number   int
	Repeated bool
	Optional bool
	Comment  string
}

// ProtoEnum represents an enum (pure data!)
type ProtoEnum struct {
	Name    string
	Values  []ProtoEnumValue
	Comment string
}

// ProtoEnumValue represents an enum value (pure data!)
type ProtoEnumValue struct {
	Name    string
	Number  int
	Comment string
}

// ProtoFile represents a complete .proto file (monoid!)
type ProtoFile struct {
	Syntax   string
	Package  string
	Options  map[string]string
	Imports  []string
	Messages []ProtoMessage
	Enums    []ProtoEnum
}

// ProtoFileMonoid combines proto files
type ProtoFileMonoid struct{}

func NewProtoFileMonoid() ProtoFileMonoid {
	return ProtoFileMonoid{}
}

func (ProtoFileMonoid) Empty() ProtoFile {
	return ProtoFile{
		Syntax:   "proto3",
		Package:  "",
		Options:  make(map[string]string),
		Imports:  []string{},
		Messages: []ProtoMessage{},
		Enums:    []ProtoEnum{},
	}
}

func (m ProtoFileMonoid) Combine(a, b ProtoFile) ProtoFile {
	// Use b's package if a's is empty
	pkg := a.Package
	if pkg == "" {
		pkg = b.Package
	}

	// Merge options
	options := make(map[string]string)
	for k, v := range a.Options {
		options[k] = v
	}
	for k, v := range b.Options {
		options[k] = v
	}

	// Merge imports (deduplicate using set monoid)
	importSet := monoid.NewSetMonoid[string]()
	for _, imp := range a.Imports {
		importSet = importSet.Insert(imp)
	}
	for _, imp := range b.Imports {
		importSet = importSet.Insert(imp)
	}

	return ProtoFile{
		Syntax:   "proto3",
		Package:  pkg,
		Options:  options,
		Imports:  importSet.ToSlice(),
		Messages: append(a.Messages, b.Messages...),
		Enums:    append(a.Enums, b.Enums...),
	}
}

// MessageMonoid for combining messages
type MessageMonoid struct{}

func NewMessageMonoid() MessageMonoid {
	return MessageMonoid{}
}

func (MessageMonoid) Empty() []ProtoMessage {
	return []ProtoMessage{}
}

func (MessageMonoid) Combine(a, b []ProtoMessage) []ProtoMessage { // This is correct
	return append(a, b...)
}

// EnumMonoid for combining enums
type EnumMonoid struct{}

func NewEnumMonoid() EnumMonoid {
	return EnumMonoid{}
}

func (EnumMonoid) Empty() []ProtoEnum {
	return []ProtoEnum{}
}

func (EnumMonoid) Combine(a, b []ProtoEnum) []ProtoEnum {
	return append(a, b...)
}
