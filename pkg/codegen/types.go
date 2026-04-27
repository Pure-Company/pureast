package codegen

import (
	"go/format"
	"strings"

	"github.com/Pure-Company/pureast/pkg/extract"
	"github.com/vinodhalaharvi/purekernels/pkg/fold"
	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
)

// GenerateTypesSummaryReport generates a summary report
func GenerateTypesSummaryReport(summary extract.TypeSummary) string {
	lines := []string{
		"# Type Summary",
		"",
		"## Overview",
		"- Total Types: " + intToString(summary.TotalTypes),
		"- Structs: " + intToString(summary.StructCount),
		"- Interfaces: " + intToString(summary.InterfaceCount),
		"- Other: " + intToString(summary.OtherCount),
		"",
		"## Type Names",
	}

	// Add names
	nameLines := fold.Map(
		func(name string) string {
			return "- " + name
		},
		summary.Names,
	)

	lines = append(lines, nameLines...)

	return strings.Join(lines, "\n")
}

func intToString(n int) string {
	// Simple int to string conversion
	if n == 0 {
		return "0"
	}

	digits := []rune{}
	for n > 0 {
		digits = append([]rune{rune('0' + n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}

// GenerateTypesOnly generates only type declarations (no methods/functions)
func (g Generator) GenerateTypesOnly(
	packageName string,
	types []extract.TypeDeclaration,
	imports []string,
) (string, error) {
	codeMonoid := NewCodeMonoid()

	// Header comment
	headerCode := ToCode("// Type definitions extracted by pureast")
	headerCode = codeMonoid.Combine(
		headerCode,
		ToCode("// Suitable for LLM context - contains only type structures"),
	)

	// Package
	packageCode := ToCode("package " + packageName)

	// Filter to only standard library and third-party imports (no internal project imports)
	filteredImports := filterRelevantImports(imports)
	importsCode := g.generateImports(filteredImports)

	// Type declarations only (map to Code)
	typeCodes := fold.Map(
		func(t extract.TypeDeclaration) Code {
			return g.generateDecl(t.Decl)
		},
		types,
	)

	// Combine all type codes
	allTypesCode := monoid.Reduce(codeMonoid, typeCodes)

	// Combine everything
	allCode := monoid.Reduce(
		codeMonoid,
		[]Code{
			headerCode,
			packageCode,
			importsCode,
			allTypesCode,
		},
	)

	result := allCode.Join("\n\n")

	// Format
	formatted, err := format.Source([]byte(result))
	if err != nil {
		return result, err
	}

	return string(formatted), nil
}

// filterRelevantImports keeps only external/stdlib imports
func filterRelevantImports(imports []string) []string {
	return fold.Filter(
		func(imp string) bool {
			// Keep standard library imports
			if !strings.Contains(imp, ".") {
				return true
			}
			// Keep third-party, but exclude project-internal
			return !strings.Contains(imp, "github.com/Pure-Company/pureast")
		},
		imports,
	)
}
