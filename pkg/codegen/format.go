package codegen

import (
	"go/format"
	"strings"

	"github.com/Pure-Company/purekernels/pkg/fold"
	"github.com/Pure-Company/purekernels/pkg/monoid"
)

// FormatCode formats Go source code (pure function)
func FormatCode(code string) (string, error) {
	formatted, err := format.Source([]byte(code))
	if err != nil {
		return code, err
	}
	return string(formatted), nil
}

// AddHeader adds a header comment to code (pure function)
func AddHeader(code string, header string) string {
	// Map each line to a commented line
	headerLines := fold.Map(
		func(line string) string {
			return "// " + line
		},
		strings.Split(header, "\n"),
	)

	// Combine with code
	allLines := append(headerLines, "", code)
	return monoid.Reduce(monoid.NewStringJoinMonoid("\n"), allLines)
}

// StripComments removes all comments from code (pure function)
func StripComments(code string) string {
	lines := strings.Split(code, "\n")

	// Filter out comment lines
	filtered := fold.Filter(
		func(line string) bool {
			trimmed := strings.TrimSpace(line)
			return !strings.HasPrefix(trimmed, "//")
		},
		lines,
	)

	return monoid.Reduce(monoid.NewStringJoinMonoid("\n"), filtered)
}

// CompactCode removes extra blank lines (pure function with state)
func CompactCode(code string) string {
	lines := strings.Split(code, "\n")

	// Fold with state to track previous blank
	type State struct {
		result    []string
		prevBlank bool
	}

	finalState := fold.FoldLeft(
		func(state State, line string) State {
			isBlank := strings.TrimSpace(line) == ""

			if isBlank {
				if !state.prevBlank {
					return State{
						result:    append(state.result, line),
						prevBlank: true,
					}
				}
				return State{
					result:    state.result,
					prevBlank: true,
				}
			}

			return State{
				result:    append(state.result, line),
				prevBlank: false,
			}
		},
		State{result: []string{}, prevBlank: false},
		lines,
	)

	return monoid.Reduce(monoid.NewStringJoinMonoid("\n"), finalState.result)
}
