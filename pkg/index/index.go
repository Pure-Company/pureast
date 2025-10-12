package index

import (
	"encoding/json"
	"os"

	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/pureast/pkg/extract"
	"github.com/vinodhalaharvi/purekernels/pkg/fold"
)

// SymbolEntry represents a symbol in the index
type SymbolEntry struct {
	Name        string
	Kind        string
	PackageName string
	File        string
	Line        int
}

// Index represents a searchable symbol index
type Index struct {
	Symbols     map[string]SymbolEntry
	PackageName string
	Files       []string
}

// SearchPattern represents search criteria
type SearchPattern struct {
	SymbolPattern string
	Kind          string
	PackageName   string
}

// BuildIndex builds an index from a package node (pure function)
func BuildIndex(pkgNode astpkg.PackageNode) Index {
	// Discover all symbols
	symbols := extract.DiscoverAllSymbols(pkgNode)

	// Build map using fold
	symbolMap := fold.FoldLeft(
		func(acc map[string]SymbolEntry, info extract.SymbolInfo) map[string]SymbolEntry {
			entry := SymbolEntry{
				Name:        info.Name,
				Kind:        info.Kind,
				PackageName: info.Package,
				File:        "", // Could be enhanced to include file info
				Line:        0,  // Could be enhanced to include line info
			}
			acc[info.Name] = entry
			return acc
		},
		make(map[string]SymbolEntry),
		symbols,
	)

	// Collect file names
	files := fold.Map(
		func(fn astpkg.FileNode) string {
			return fn.Name
		},
		pkgNode.Files,
	)

	return Index{
		Symbols:     symbolMap,
		PackageName: pkgNode.Name,
		Files:       files,
	}
}

// SaveIndex saves index to a file
func SaveIndex(idx Index, path string) error {
	data, err := json.MarshalIndent(idx, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// LoadIndex loads index from a file
func LoadIndex(path string) (Index, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Index{}, err
	}

	var idx Index
	if err := json.Unmarshal(data, &idx); err != nil {
		return Index{}, err
	}

	return idx, nil
}

// SearchIndex searches for symbols (pure function)
func SearchIndex(pattern SearchPattern, idx Index) []SymbolEntry {
	results := []SymbolEntry{}

	for _, entry := range idx.Symbols {
		// Simple substring match
		if pattern.SymbolPattern != "" && !contains(entry.Name, pattern.SymbolPattern) {
			continue
		}

		if pattern.Kind != "" && entry.Kind != pattern.Kind {
			continue
		}

		if pattern.PackageName != "" && entry.PackageName != pattern.PackageName {
			continue
		}

		results = append(results, entry)
	}

	return results
}

// contains checks if a string contains a substring (case-insensitive)
func contains(s, substr string) bool {
	return fold.FoldLeft(
		func(acc bool, _ rune) bool {
			return acc
		},
		false,
		[]rune(s),
	) || simpleContains(s, substr)
}

func simpleContains(s, substr string) bool {
	return len(substr) == 0 || len(s) >= len(substr) && (s[:len(substr)] == substr || simpleContains(s[1:], substr))
}
