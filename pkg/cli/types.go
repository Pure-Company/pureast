// pkg/cli/types.go
package cli

// Output represents command output (this is a monoid!)
type Output struct {
    Text     string
    ExitCode int
}

// OutputMonoid combines outputs
type OutputMonoid struct{}

func NewOutputMonoid() OutputMonoid {
    return OutputMonoid{}
}

func (OutputMonoid) Empty() Output {
    return Output{Text: "", ExitCode: 0}
}

func (OutputMonoid) Combine(a, b Output) Output {
    exitCode := a.ExitCode
    if exitCode == 0 {
        exitCode = b.ExitCode
    }
    return Output{
        Text:     a.Text + b.Text,
        ExitCode: exitCode,
    }
}


