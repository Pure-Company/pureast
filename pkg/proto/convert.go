package proto

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
	"unicode"

	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/purekernels/pkg/fold"
	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
	"github.com/vinodhalaharvi/purekernels/pkg/result"
)

// convertFieldToProtoFields converts a single Go field to proto fields (pure!)
// Returns empty slice for unsupported types
// filterAndRenumberFields filters successful results and assigns field numbers (pure!)
func filterAndRenumberFields(results []result.Result[ProtoField]) []ProtoField {
	// Filter only successful results
	fields := fold.Reduce(
		monoid.NewListMonoid[ProtoField](),
		fold.Map(
			func(res result.Result[ProtoField]) []ProtoField {
				if res.IsOk() {
					return []ProtoField{res.Unwrap()}
				}
				return []ProtoField{}
			},
			results,
		),
	)

	// Renumber fields sequentially using MapWithIndex (pure!)
	return fold.MapWithIndex(
		func(i int, field ProtoField) ProtoField {
			return ProtoField{
				Name:     field.Name,
				Type:     field.Type,
				Number:   i + 1,
				Repeated: field.Repeated,
				Optional: field.Optional,
				Comment:  field.Comment,
			}
		},
		fields,
	)
}

// TypeInfo represents the result of type conversion (pure data!)
type TypeInfo struct {
	ProtoType   string
	Repeated    bool
	NeedsImport string
}

// mapQualifiedType maps qualified types (pure!)
// Returns Result[TypeInfo] - Err for unsupported types
func mapQualifiedType(pkg, name string) result.Result[TypeInfo] {
	// Well-known types
	if pkg == "time" && name == "Time" {
		return result.Ok(TypeInfo{
			ProtoType:   "google.protobuf.Timestamp",
			Repeated:    false,
			NeedsImport: "google/protobuf/timestamp.proto",
		})
	}

	if pkg == "time" && name == "Duration" {
		return result.Ok(TypeInfo{
			ProtoType:   "google.protobuf.Duration",
			Repeated:    false,
			NeedsImport: "google/protobuf/duration.proto",
		})
	}

	if pkg == "uuid" && name == "UUID" {
		return result.Ok(TypeInfo{
			ProtoType:   "string",
			Repeated:    false,
			NeedsImport: "",
		})
	}

	// Skip known unserializable qualified types
	unserializableTypes := map[string]map[string]bool{
		"io": {
			"Reader":      true,
			"Writer":      true,
			"ReadWriter":  true,
			"ReadCloser":  true,
			"WriteCloser": true,
			"Closer":      true,
		},
		"context": {
			"Context": true,
		},
		"sync": {
			"Mutex":     true,
			"RWMutex":   true,
			"WaitGroup": true,
			"Cond":      true,
			"Once":      true,
			"Pool":      true,
		},
		"net": {
			"Conn":     true,
			"Listener": true,
		},
		"http": {
			"Handler":        true,
			"RoundTripper":   true,
			"ResponseWriter": true,
		},
		"tls": {
			"Config":          true,
			"ConnectionState": true,
			"Conn":            true,
		},
	}

	if pkgTypes, ok := unserializableTypes[pkg]; ok {
		if pkgTypes[name] {
			return result.Err[TypeInfo](fmt.Errorf("unserializable type: %s.%s", pkg, name))
		}
	}

	// Default: use the type name as a message reference
	return result.Ok(TypeInfo{
		ProtoType:   name,
		Repeated:    false,
		NeedsImport: "",
	})
}

// toScreamingSnakeCase converts CamelCase to SCREAMING_SNAKE_CASE (pure!)
func toScreamingSnakeCase(s string) string {
	var result strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if unicode.IsUpper(r) && i > 0 {
			prevIsLower := unicode.IsLower(runes[i-1])
			nextIsLower := i < len(runes)-1 && unicode.IsLower(runes[i+1])

			if prevIsLower || nextIsLower {
				result.WriteRune('_')
			}
		}

		result.WriteRune(unicode.ToUpper(r))
	}

	return result.String()
}

// extractComment extracts comment text (pure!)
func extractComment(cg *ast.CommentGroup) string {
	if cg == nil {
		return ""
	}

	comments := fold.Map(
		func(c *ast.Comment) string {
			text := c.Text
			text = strings.TrimPrefix(text, "// ")
			text = strings.TrimPrefix(text, "/* ")
			text = strings.TrimSuffix(text, " */")
			return text
		},
		cg.List,
	)

	return strings.Join(comments, " ")
}

// Errors
var (
	ErrNotAStruct = errors.New("declaration is not a struct")
	ErrNotAnEnum  = errors.New("declaration is not an enum")
)

// mapBasicType maps Go primitives to proto types (pure!)
// Returns empty string for unsupported types
func mapBasicType(goType string) string {
	mapping := map[string]string{
		"string":  "string",
		"int":     "int64",
		"int8":    "int32",
		"int16":   "int32",
		"int32":   "int32",
		"int64":   "int64",
		"uint":    "uint64",
		"uint8":   "uint32",
		"uint16":  "uint32",
		"uint32":  "uint32",
		"uint64":  "uint64",
		"float32": "float",
		"float64": "double",
		"bool":    "bool",
		"byte":    "bytes",
		"rune":    "int32",
	}

	if protoType, ok := mapping[goType]; ok {
		return protoType
	}

	// If not in mapping and not unserializable, assume it's a custom message type
	return goType
}

// toSnakeCase converts CamelCase to snake_case (pure!)
// Handles acronyms properly (URL → url, HTTP → http, TLS → tls)
func toSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if unicode.IsUpper(r) {
			// Add underscore before uppercase if:
			// 1. Not first character AND
			// 2. (Previous is lowercase OR next is lowercase)
			if i > 0 {
				prev := runes[i-1]
				prevIsLower := unicode.IsLower(prev)
				nextIsLower := i < len(runes)-1 && unicode.IsLower(runes[i+1])

				// Insert underscore if transitioning from lower to upper
				// OR if in the middle of an acronym followed by lowercase (HTTPServer → http_server)
				if prevIsLower || (unicode.IsUpper(prev) && nextIsLower) {
					result.WriteRune('_')
				}
			}
		}

		result.WriteRune(unicode.ToLower(r))
	}

	return result.String()
}

// ConvertStructToMessage converts Go struct to proto message (pure!)
// ConvertStructToMessage converts Go struct to proto message (pure!)
// Note: Without declMap, interface types will be treated as custom messages
func ConvertStructToMessage(decl astpkg.DeclNode) result.Result[ProtoMessage] {
	return ConvertStructToMessageWithPackage(decl, "", nil)
}

// ConvertStructToMessageWithDeclMap converts with interface detection support
func ConvertStructToMessageWithDeclMap(
	decl astpkg.DeclNode,
	declMap map[string]astpkg.DeclNode,
) result.Result[ProtoMessage] {
	return ConvertStructToMessageWithPackage(decl, "", declMap)
}

// convertFieldToProtoFields converts a single Go field to proto fields (pure!)
// Returns empty slice for unsupported types
func convertFieldToProtoFields(field *ast.Field) []result.Result[ProtoField] {
	// Skip embedded fields
	if len(field.Names) == 0 {
		return []result.Result[ProtoField]{}
	}

	name := field.Names[0].Name

	// Skip unexported fields
	if !ast.IsExported(name) {
		return []result.Result[ProtoField]{}
	}

	// Convert type - returns Result
	typeResult := convertType(field.Type)

	if !typeResult.IsOk() {
		// Unsupported type - return empty (pure way to skip)
		return []result.Result[ProtoField]{}
	}

	typeInfo := typeResult.Unwrap()

	// Create proto field (number will be assigned later)
	protoField := ProtoField{
		Name:     toSnakeCase(name),
		Type:     typeInfo.ProtoType,
		Number:   0, // Will be renumbered
		Repeated: typeInfo.Repeated,
		Comment:  extractComment(field.Doc),
	}

	return []result.Result[ProtoField]{result.Ok(protoField)}
}

// convertType converts Go type to proto type info (pure!)
// Returns Result[TypeInfo] - Err for unsupported types
func convertType(expr ast.Expr) result.Result[TypeInfo] {
	return convertTypeWithContext(expr, nil)
}

// convertTypeWithContext converts Go type with access to declaration context (pure!)
func convertTypeWithContext(expr ast.Expr, declMap map[string]astpkg.DeclNode) result.Result[TypeInfo] {
	switch t := expr.(type) {
	case *ast.Ident:
		// Check if it's a basic type first
		protoType := mapBasicType(t.Name)
		if protoType != "" {
			return result.Ok(TypeInfo{
				ProtoType:   protoType,
				Repeated:    false,
				NeedsImport: "",
			})
		}

		// Check if it's a known unserializable type
		if isUnserializableType(t.Name) {
			return result.Err[TypeInfo](fmt.Errorf("unserializable type: %s", t.Name))
		}

		// 🚫 Heuristic safeguard for common interface-like identifiers
		lower := strings.ToLower(t.Name)
		if lower == "logger" || lower == "writer" || lower == "reader" || lower == "closer" {
			return result.Err[TypeInfo](fmt.Errorf("interface-like type: %s", t.Name))
		}

		// If we have declaration context, check if it's an interface type declared in this package
		if declMap != nil {
			if decl, ok := declMap[t.Name]; ok {
				if isInterfaceDecl(decl) {
					// 🚫 Skip any fields that reference a locally-declared interface
					return result.Err[TypeInfo](fmt.Errorf("interface type: %s", t.Name))
				}
			}
		}

		// Otherwise, assume it's a custom message type
		return result.Ok(TypeInfo{
			ProtoType:   t.Name,
			Repeated:    false,
			NeedsImport: "",
		})

	case *ast.StarExpr:
		// Pointer types - proto3 doesn't have explicit optional
		return convertTypeWithContext(t.X, declMap)

	case *ast.ArrayType:
		// Arrays/slices become repeated
		innerResult := convertTypeWithContext(t.Elt, declMap)
		if !innerResult.IsOk() {
			return result.Err[TypeInfo](errors.New("unsupported array element type"))
		}

		inner := innerResult.Unwrap()
		return result.Ok(TypeInfo{
			ProtoType:   inner.ProtoType,
			Repeated:    true,
			NeedsImport: inner.NeedsImport,
		})

	case *ast.MapType:
		// Map types
		keyResult := convertTypeWithContext(t.Key, declMap)
		valueResult := convertTypeWithContext(t.Value, declMap)

		if !keyResult.IsOk() || !valueResult.IsOk() {
			return result.Err[TypeInfo](errors.New("unsupported map key or value type"))
		}

		key := keyResult.Unwrap()
		value := valueResult.Unwrap()

		return result.Ok(TypeInfo{
			ProtoType:   fmt.Sprintf("map<%s, %s>", key.ProtoType, value.ProtoType),
			Repeated:    false,
			NeedsImport: "",
		})

	case *ast.SelectorExpr:
		// Qualified types (e.g., time.Time, log.Logger)
		if ident, ok := t.X.(*ast.Ident); ok {
			// 🚫 Skip known qualified interface-like or unserializable types
			if (ident.Name == "log" && t.Sel.Name == "Logger") ||
				(ident.Name == "zap" && t.Sel.Name == "Logger") ||
				(ident.Name == "io" && (t.Sel.Name == "Writer" || t.Sel.Name == "Reader" || t.Sel.Name == "Closer")) {
				return result.Err[TypeInfo](fmt.Errorf("unserializable qualified type: %s.%s", ident.Name, t.Sel.Name))
			}
			return mapQualifiedType(ident.Name, t.Sel.Name)
		}
		return result.Err[TypeInfo](errors.New("unknown qualified type"))

	case *ast.InterfaceType:
		// Inline interface definition - can't be serialized
		return result.Err[TypeInfo](errors.New("inline interface cannot be serialized"))

	case *ast.FuncType:
		// Functions can't be serialized
		return result.Err[TypeInfo](errors.New("functions cannot be serialized"))

	case *ast.ChanType:
		// Channels can't be serialized
		return result.Err[TypeInfo](errors.New("channels cannot be serialized"))

	default:
		return result.Err[TypeInfo](errors.New("unknown type"))
	}
}

// isInterfaceDecl checks if a declaration is an interface (pure!)
func isInterfaceDecl(decl astpkg.DeclNode) bool {
	genDecl, ok := decl.Decl.(*ast.GenDecl)
	if !ok {
		return false
	}

	for _, spec := range genDecl.Specs {
		if typeSpec, ok := spec.(*ast.TypeSpec); ok {
			if _, isIface := typeSpec.Type.(*ast.InterfaceType); isIface {
				return true
			}
		}
	}

	return false
}

func isUnserializableType(typeName string) bool {
	unserializable := map[string]bool{
		// Core interfaces / sync primitives
		"Context":   true,
		"error":     true,
		"Mutex":     true,
		"RWMutex":   true,
		"WaitGroup": true,
		"Cond":      true,
		"Once":      true,
		"Pool":      true,

		// Common logging interfaces
		"Logger":     true,
		"log.Logger": true,

		// Other unmarshalable I/O types
		"Reader":      true,
		"Writer":      true,
		"Closer":      true,
		"ReadCloser":  true,
		"WriteCloser": true,
	}

	return unserializable[typeName]
}

// ConvertStructToMessageWithPackage converts Go struct to proto message with package context (pure!)
func ConvertStructToMessageWithPackage(
	decl astpkg.DeclNode,
	pkgName string,
	declMap map[string]astpkg.DeclNode,
) result.Result[ProtoMessage] {
	genDecl, ok := decl.Decl.(*ast.GenDecl)
	if !ok {
		return result.Err[ProtoMessage](ErrNotAStruct)
	}

	// 🚫 Skip if this declaration defines an interface
	for _, spec := range genDecl.Specs {
		if typeSpec, ok := spec.(*ast.TypeSpec); ok {
			if _, isIface := typeSpec.Type.(*ast.InterfaceType); isIface {
				// Simply skip interface declarations — not serializable in proto
				return result.Err[ProtoMessage](ErrNotAStruct)
			}
		}
	}

	for _, spec := range genDecl.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}

		// Convert fields using fold (categorical!)
		fieldResults := fold.Map(
			func(field *ast.Field) []result.Result[ProtoField] {
				return convertFieldToProtoFieldsWithContext(field, declMap)
			},
			structType.Fields.List,
		)

		// Flatten results using Reduce with monoid
		flatResults := fold.Reduce(
			monoid.NewListMonoid[result.Result[ProtoField]](),
			fieldResults,
		)

		// Filter successful results and renumber
		fields := filterAndRenumberFields(flatResults)

		// Extract comment
		comment := extractComment(genDecl.Doc)

		// Create message name with optional package prefix
		messageName := typeSpec.Name.Name
		if pkgName != "" && pkgName != "main" {
			messageName = pkgName + "_" + messageName
		}

		return result.Ok(ProtoMessage{
			Name:    messageName,
			Fields:  fields,
			Comment: comment,
		})
	}

	return result.Err[ProtoMessage](ErrNotAStruct)
}

// convertFieldToProtoFieldsWithContext converts field with declaration context
func convertFieldToProtoFieldsWithContext(
	field *ast.Field,
	declMap map[string]astpkg.DeclNode,
) []result.Result[ProtoField] {
	// Skip embedded fields
	if len(field.Names) == 0 {
		return []result.Result[ProtoField]{}
	}

	name := field.Names[0].Name

	// Skip unexported fields
	if !ast.IsExported(name) {
		return []result.Result[ProtoField]{}
	}

	// Convert type with context - returns Result
	typeResult := convertTypeWithContext(field.Type, declMap)

	if !typeResult.IsOk() {
		// Unsupported type - return empty (pure way to skip)
		return []result.Result[ProtoField]{}
	}

	typeInfo := typeResult.Unwrap()

	// Create proto field (number will be assigned later)
	protoField := ProtoField{
		Name:     toSnakeCase(name),
		Type:     typeInfo.ProtoType,
		Number:   0, // Will be renumbered
		Repeated: typeInfo.Repeated,
		Comment:  extractComment(field.Doc),
	}

	return []result.Result[ProtoField]{result.Ok(protoField)}
}
