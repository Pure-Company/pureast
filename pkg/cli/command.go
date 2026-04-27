// pkg/cli/command.go
//
// Type-safe command builder around Cobra. Every verb runs through
// the same shape:
//
//   1. ParseArgs: turn (cobra.Command, []string) into a typed args
//      struct, or return an error explaining what's wrong.
//   2. Action: take the typed args and return Output (which goes to
//      stdout) or an error (which goes to stderr).
//
// The signatures use plain (T, error) rather than a Result wrapper
// because Go's idioms expect that shape, and because errors here
// genuinely terminate the verb — they're not values being threaded
// through pure computations. Earlier this file used result.Result[T]
// with a Text-and-ExitCode pattern for errors, which had two real
// problems: error messages went to stdout (so they polluted pipes
// like `pureast deps Foo ./pkg | jq .`), and every action carried
// its own `"Error: %v\n"` formatting boilerplate. Both gone.
package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Action is the core abstraction. An Action consumes typed args and
// produces either an Output (success) or an error (failure). Errors
// surface via Cobra's RunE → main.go's stderr writer, with an exit
// code of 1.
//
// For non-fatal exit codes (e.g. "found differences, exit 2 like
// `diff(1)`"), set Output.ExitCode on a successful Output. Don't
// conflate that with the error path.
type Action[A any] func(context.Context, A) (Output, error)

// CommandBuilder builds type-safe commands.
type CommandBuilder[A any] struct {
	name      string
	short     string
	long      string
	parseArgs func(*cobra.Command, []string) (A, error)
	action    Action[A]
}

func NewCommand[A any](name string) *CommandBuilder[A] {
	return &CommandBuilder[A]{name: name}
}

func (b *CommandBuilder[A]) Short(desc string) *CommandBuilder[A] {
	b.short = desc
	return b
}

func (b *CommandBuilder[A]) Long(desc string) *CommandBuilder[A] {
	b.long = desc
	return b
}

func (b *CommandBuilder[A]) ParseArgs(fn func(*cobra.Command, []string) (A, error)) *CommandBuilder[A] {
	b.parseArgs = fn
	return b
}

func (b *CommandBuilder[A]) Action(action Action[A]) *CommandBuilder[A] {
	b.action = action
	return b
}

func (b *CommandBuilder[A]) Build() *cobra.Command {
	cmd := &cobra.Command{
		Use:   b.name,
		Short: b.short,
		Long:  b.long,
		RunE: func(cmd *cobra.Command, args []string) error {
			parsed, err := b.parseArgs(cmd, args)
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			output, err := b.action(ctx, parsed)
			if err != nil {
				return err
			}

			// Tool output goes to stdout so it composes with pipes:
			//   pureast list ./pkg | head -10
			//   pureast dump ./pkg --format md | wc -l
			//   pureast deps X ./pkg --format json | jq .
			if output.Text != "" {
				fmt.Fprint(os.Stdout, output.Text)
			}

			if output.ExitCode != 0 {
				return &ExitError{Code: output.ExitCode}
			}

			return nil
		},
	}
	return cmd
}

// ExitError carries a non-zero exit code without an error message.
// Used when a successful Output requested a specific exit code (rare:
// `diff` style "I succeeded, but found something").
type ExitError struct {
	Code int
}

func (e *ExitError) Error() string {
	return ""
}

func (e *ExitError) ExitCode() int {
	return e.Code
}
