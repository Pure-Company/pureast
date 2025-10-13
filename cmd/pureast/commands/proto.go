// cmd/pureast-cobra/commands/proto.go
package commands

import (
    "context"
    "fmt"
    "go/token"
    "os"
    "strings"

    "github.com/spf13/cobra"
    "github.com/vinodhalaharvi/pureast/pkg/cli"
    "github.com/vinodhalaharvi/pureast/pkg/extract"
    "github.com/vinodhalaharvi/pureast/pkg/proto"
    "github.com/vinodhalaharvi/purekernels/pkg/result"
)

type ProtoArgs struct {
    FilePath   string
    OutputFile string
    Types      []string
    Workers    int
}

func NewProtoCommand() *cobra.Command {
    cmd := cli.NewCommand[ProtoArgs]("proto").
        Short("Generate protobuf schema from Go types").
        Long(`Generate Protocol Buffer schema from Go struct definitions.

Examples:
  pureast proto --file ./pkg
  pureast proto --file ./pkg --types User,Profile
  pureast proto --file ./pkg --output schema.proto`).
        ParseArgs(parseProtoArgs).
        Action(protoAction).
        Build()

    cmd.Flags().StringP("file", "f", "", "Go file or directory (required)")
    cmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")
    cmd.Flags().StringSlice("types", []string{}, "Comma-separated type names (empty = all)")
    cmd.Flags().IntP("workers", "w", 0, "Number of workers (0 = auto)")
    
    cmd.MarkFlagRequired("file")

    return cmd
}

func parseProtoArgs(cmd *cobra.Command, args []string) result.Result[ProtoArgs] {
    file, _ := cmd.Flags().GetString("file")
    output, _ := cmd.Flags().GetString("output")
    types, _ := cmd.Flags().GetStringSlice("types")
    workers, _ := cmd.Flags().GetInt("workers")

    // Trim whitespace from types
    cleanTypes := []string{}
    for _, t := range types {
        trimmed := strings.TrimSpace(t)
        if trimmed != "" {
            cleanTypes = append(cleanTypes, trimmed)
        }
    }

    return result.Ok(ProtoArgs{
        FilePath:   file,
        OutputFile: output,
        Types:      cleanTypes,
        Workers:    workers,
    })
}

func protoAction(ctx context.Context, args ProtoArgs) result.Result[cli.Output] {
    fset := token.NewFileSet()
    pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, args.Workers)
    if err != nil {
        return result.Ok(cli.Output{
            Text:     fmt.Sprintf("Error: %v\n", err),
            ExitCode: 1,
        })
    }

    protoFile := proto.GenerateProtoFromPackageConcurrent(
        pkgNode,
        args.Types,
        args.Workers,
    ).Value()

    protoCode := proto.FormatProtoFile(protoFile)

    if args.OutputFile != "" {
        if err := os.WriteFile(args.OutputFile, []byte(protoCode), 0644); err != nil {
            return result.Ok(cli.Output{
                Text:     fmt.Sprintf("Error writing file: %v\n", err),
                ExitCode: 1,
            })
        }
        return result.Ok(cli.Output{
            Text:     fmt.Sprintf("Written to %s\n", args.OutputFile),
            ExitCode: 0,
        })
    }

    return result.Ok(cli.Output{Text: protoCode, ExitCode: 0})
}

