package proto

import (
	"fmt"
	"strings"

	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/purekernels/pkg/fold"
	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
)

// matchesFilter checks if name matches any filter (pure!)
func matchesFilter(name string, filters []string) bool {
	if len(filters) == 0 {
		return true
	}

	for _, filter := range filters {
		if name == filter {
			return true
		}
	}

	return false
}

// collectImports collects needed imports from messages (pure!)
func collectImports(messages []ProtoMessage) []string {
	importSet := monoid.NewSetMonoid[string]()

	for _, msg := range messages {
		for _, field := range msg.Fields {
			if strings.Contains(field.Type, "google.protobuf.") {
				if strings.Contains(field.Type, "Timestamp") {
					importSet = importSet.Insert("google/protobuf/timestamp.proto")
				}
				if strings.Contains(field.Type, "Duration") {
					importSet = importSet.Insert("google/protobuf/duration.proto")
				}
			}
		}
	}

	return importSet.ToSlice()
}

// FormatProtoFile formats proto file as string (pure!)
func FormatProtoFile(pf ProtoFile) string {
	var parts []string

	// Syntax
	parts = append(parts, fmt.Sprintf("syntax = \"%s\";\n", pf.Syntax))

	// Package
	if pf.Package != "" {
		parts = append(parts, fmt.Sprintf("package %s;\n", pf.Package))
	}

	// Options
	if len(pf.Options) > 0 {
		for key, value := range pf.Options {
			parts = append(parts, fmt.Sprintf("option %s = \"%s\";\n", key, value))
		}
	}

	// Imports
	if len(pf.Imports) > 0 {
		parts = append(parts, "")
		for _, imp := range pf.Imports {
			parts = append(parts, fmt.Sprintf("import \"%s\";", imp))
		}
	}

	// Messages
	if len(pf.Messages) > 0 {
		parts = append(parts, "")
		for _, msg := range pf.Messages {
			parts = append(parts, formatMessage(msg))
			parts = append(parts, "")
		}
	}

	// Enums
	if len(pf.Enums) > 0 {
		for _, enum := range pf.Enums {
			parts = append(parts, formatEnum(enum))
			parts = append(parts, "")
		}
	}

	// Combine using monoid!
	return monoid.Reduce(
		monoid.NewStringJoinMonoid("\n"),
		parts,
	)
}

// formatMessage formats proto message (pure!)
func formatMessage(msg ProtoMessage) string {
	var lines []string

	// Comment
	if msg.Comment != "" {
		lines = append(lines, "// "+msg.Comment)
	}

	// Message declaration
	lines = append(lines, fmt.Sprintf("message %s {", msg.Name))

	// Fields
	for _, field := range msg.Fields {
		fieldLine := formatField(field)
		lines = append(lines, "  "+fieldLine)
	}

	lines = append(lines, "}")

	return strings.Join(lines, "\n")
}

// formatField formats proto field (pure!)
func formatField(field ProtoField) string {
	var parts []string

	// Repeated modifier
	if field.Repeated {
		parts = append(parts, "repeated")
	}

	// Type and name
	parts = append(parts, field.Type)
	parts = append(parts, field.Name)

	// Field number
	parts = append(parts, fmt.Sprintf("= %d;", field.Number))

	fieldStr := strings.Join(parts, " ")

	// Add comment if present
	if field.Comment != "" {
		fieldStr += " // " + field.Comment
	}

	return fieldStr
}

// formatEnum formats proto enum (pure!)
func formatEnum(enum ProtoEnum) string {
	var lines []string

	// Comment
	if enum.Comment != "" {
		lines = append(lines, "// "+enum.Comment)
	}

	// Enum declaration
	lines = append(lines, fmt.Sprintf("enum %s {", enum.Name))

	// Values
	for _, value := range enum.Values {
		valueLine := fmt.Sprintf("  %s = %d;", value.Name, value.Number)
		if value.Comment != "" {
			valueLine += " // " + value.Comment
		}
		lines = append(lines, valueLine)
	}

	lines = append(lines, "}")

	return strings.Join(lines, "\n")
}

// generateProtoFromFile generates proto from single file (pure!)
// This is used when conflict detection across files isn't needed
func generateProtoFromFile(
	fileNode astpkg.FileNode,
	typeFilter []string,
) ProtoFile {
	// For single file, no cross-file conflicts possible
	// Just use empty conflict map
	return generateProtoFromFileWithConflictDetection(
		fileNode,
		typeFilter,
		make(map[string][]string), // No conflicts
	)
}

// generateProtoFromFileWithConflictDetection generates proto with duplicate detection
func generateProtoFromFileWithConflictDetection(
	fileNode astpkg.FileNode,
	typeFilter []string,
	typeSources map[string][]string,
) ProtoFile {
	// Filter declarations if needed
	decls := fileNode.Decls
	if len(typeFilter) > 0 {
		decls = fold.Filter(
			func(decl astpkg.DeclNode) bool {
				return matchesFilter(decl.Name, typeFilter)
			},
			decls,
		)
	}

	// Convert with package prefix if there are conflicts
	messages := fold.FoldLeft(
		func(acc []ProtoMessage, decl astpkg.DeclNode) []ProtoMessage {
			// Check if this type name appears in multiple files
			pkgPrefix := ""
			if sources := typeSources[decl.Name]; len(sources) > 1 {
				// Conflict detected - use file name as prefix
				pkgPrefix = fileNode.Name
			}

			msgResult := ConvertStructToMessageWithPackage(decl, pkgPrefix)
			if msgResult.IsOk() {
				return append(acc, msgResult.Unwrap())
			}
			return acc
		},
		[]ProtoMessage{},
		decls,
	)

	// Collect needed imports
	imports := collectImports(messages)

	return ProtoFile{
		Syntax:   "proto3",
		Package:  fileNode.Name + ".v1",
		Imports:  imports,
		Messages: messages,
		Enums:    []ProtoEnum{},
	}
}

// GenerateProtoFromPackage generates .proto from package (pure using fold!)
func GenerateProtoFromPackage(
	pkgNode astpkg.PackageNode,
	typeFilter []string,
) ProtoFile {
	protoMonoid := NewProtoFileMonoid()

	// Initialize with package name
	baseProto := ProtoFile{
		Syntax:  "proto3",
		Package: pkgNode.Name + ".v1",
		Options: map[string]string{
			"go_package": "./" + pkgNode.Name + ";pb",
		},
		Imports:  []string{},
		Messages: []ProtoMessage{},
		Enums:    []ProtoEnum{},
	}

	// First pass: detect duplicate type names across files
	typeSources := make(map[string][]string) // typeName -> []fileName

	for _, fileNode := range pkgNode.Files {
		for _, decl := range fileNode.Decls {
			if len(typeFilter) == 0 || matchesFilter(decl.Name, typeFilter) {
				if sources, ok := typeSources[decl.Name]; ok {
					typeSources[decl.Name] = append(sources, fileNode.Name)
				} else {
					typeSources[decl.Name] = []string{fileNode.Name}
				}
			}
		}
	}

	// Second pass: generate proto with package prefix if needed
	result := fold.FoldLeft(
		func(acc ProtoFile, fileNode astpkg.FileNode) ProtoFile {
			fileProto := generateProtoFromFileWithConflictDetection(
				fileNode,
				typeFilter,
				typeSources,
			)
			return protoMonoid.Combine(acc, fileProto)
		},
		baseProto,
		pkgNode.Files,
	)

	return result
}
