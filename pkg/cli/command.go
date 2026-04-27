// pkg/cli/command.go
package cli

import (
    "context"
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/vinodhalaharvi/purekernels/pkg/result"
)

// Action is the core abstraction
type Action[A any] func(context.Context, A) result.Result[Output]

// CommandBuilder builds type-safe commands
type CommandBuilder[A any] struct {
    name      string
    short     string
    long      string
    parseArgs func(*cobra.Command, []string) result.Result[A]
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

func (b *CommandBuilder[A]) ParseArgs(fn func(*cobra.Command, []string) result.Result[A]) *CommandBuilder[A] {
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
            argsResult := b.parseArgs(cmd, args)
            if !argsResult.IsOk() {
                return argsResult.Error()
            }

            ctx := cmd.Context()
            outputResult := b.action(ctx, argsResult.Unwrap())

            if !outputResult.IsOk() {
                return outputResult.Error()
            }

            output := outputResult.Unwrap()

            // Tool output goes to stdout so it composes with pipes:
            //   pureast list ./pkg | head -10
            //   pureast dump ./pkg --format md | wc -l
            //   pureast deps X ./pkg --format json | jq .
            //
            // Earlier this used cmd.Print, which routes through cobra's
            // OutOrStderr — fine for help text and error messages, wrong
            // for primary output. Errors continue to go to stderr via
            // RunE's returned error and main.go's stderr writer.
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

type ExitError struct {
    Code int
}

func (e *ExitError) Error() string {
    return ""
}

func (e *ExitError) ExitCode() int {
    return e.Code
}


