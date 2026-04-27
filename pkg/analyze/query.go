// pkg/analyze/query.go - High-level query API

package analyze

import (
	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
)

// DependencyQuery represents different ways to query dependencies
type DependencyQuery int

const (
	// TypesOnly gets only type dependencies (transitive closure)
	TypesOnly DependencyQuery = iota

	// WithAssociatedCode gets types + constructors + methods
	WithAssociatedCode

	// Minimal gets minimal set needed for compilation
	Minimal
)

// Query executes a dependency query with the specified strategy
func (g DependencyGraph) Query(targetName string, query DependencyQuery) astpkg.Dependencies {
	switch query {
	case TypesOnly:
		return g.ResolveTransitive(targetName)
	case WithAssociatedCode:
		return g.ResolveWithAssociatedCode(targetName)
	case Minimal:
		return g.MinimalDependencies(targetName)
	default:
		return g.ResolveTransitive(targetName)
	}
}

// DependencyReport contains structured dependency information
type DependencyReport struct {
	Target       string
	Dependencies astpkg.Dependencies
	Stats        DependencyStats
	Order        []string // Topological order
	Cycles       [][]string
}

// AnalyzeSymbol performs comprehensive analysis of a symbol
func (g DependencyGraph) AnalyzeSymbol(targetName string, query DependencyQuery) DependencyReport {
	deps := g.Query(targetName, query)
	stats := g.ComputeStats(targetName)
	order := g.DependencyOrder(targetName)
	cycles := g.CircularDependencies(targetName)

	return DependencyReport{
		Target:       targetName,
		Dependencies: deps,
		Stats:        stats,
		Order:        order,
		Cycles:       cycles,
	}
}

// FormatDependencies formats dependencies as a string (pure function)
func FormatDependencies(target string, deps astpkg.Dependencies) string {
	output := ""

	if deps.Types.Size() > 0 {
		output += formatSection("Types", deps.Types.ToSlice())
	}

	if deps.Functions.Size() > 0 {
		output += formatSection("Functions", deps.Functions.ToSlice())
	}

	if deps.Structs.Size() > 0 {
		output += formatSection("Structs", deps.Structs.ToSlice())
	}

	if deps.Interfaces.Size() > 0 {
		output += formatSection("Interfaces", deps.Interfaces.ToSlice())
	}

	if deps.Imports.Size() > 0 {
		output += formatSection("Imports", deps.Imports.ToSlice())
	}

	if output == "" {
		return "No dependencies found.\n"
	}

	return output
}

func formatSection(title string, items []string) string {
	if len(items) == 0 {
		return ""
	}

	output := title + " (" + intToString(len(items)) + "):\n"
	for _, item := range items {
		output += "  - " + item + "\n"
	}
	output += "\n"
	return output
}

func intToString(n int) string {
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

// FormatReport formats a full dependency report
func FormatReport(report DependencyReport) string {
	output := "Dependencies for " + report.Target + ":\n\n"
	output += FormatDependencies(report.Target, report.Dependencies)

	if len(report.Cycles) > 0 {
		output += "\n⚠️  Circular Dependencies Detected:\n"
		for i, cycle := range report.Cycles {
			output += "  " + intToString(i+1) + ". "
			for j, node := range cycle {
				if j > 0 {
					output += " -> "
				}
				output += node
			}
			output += "\n"
		}
	}

	return output
}
