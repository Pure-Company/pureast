// cmd/pureast-cobra/commands/types.go
package commands

import (
    "context"
    "fmt"
    "go/token"
    "os"

    "github.com/spf13/cobra"
    "github.com/vinodhalaharvi/pureast/pkg/cli"
    "github.com/vinodhalaharvi/pureast/pkg/codegen"
    "github.com/vinodhalaharvi/pureast/pkg/extract"
    "github.com/vinodhalaharvi/purekernels/pkg/result"
)

type TypesArgs struct {
    FilePath       string
    OutputFile     string
    StructsOnly    bool
    InterfacesOnly bool
}

func NewTypesCommand() *cobra.Command {
    cmd := cli.NewCommand[TypesArgs]("types").
        Short("Extract type definitions (for LLM context)").
        Long(`Extract all type definitions without implementations.
Perfect for providing clean context to LLMs.

Examples:
  pureast types --file ./pkg
  pureast types --file ./pkg --structs-only
  pureast types --file ./pkg --interfaces-only`).
        ParseArgs(parseTypesArgs).
        Action(typesAction).
        Build()

    cmd.Flags().StringP("file", "f", "", "Go file or directory (required)")
    cmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")
    cmd.Flags().Bool("structs-only", false, "Extract only structs")
    cmd.Flags().Bool("interfaces-only", false, "Extract only interfaces")
    
    cmd.MarkFlagRequired("file")
    cmd.MarkFlagsMutuallyExclusive("structs-only", "interfaces-only")

    return cmd
}

func parseTypesArgs(cmd *cobra.Command, args []string) result.Result[TypesArgs] {
    file, _ := cmd.Flags().GetString("file")
    output, _ := cmd.Flags().GetString("output")
    structsOnly, _ := cmd.Flags().GetBool("structs-only")
    interfacesOnly, _ := cmd.Flags().GetBool("interfaces-only")

    return result.Ok(TypesArgs{
        FilePath:       file,
        OutputFile:     output,
        StructsOnly:    structsOnly,
        InterfacesOnly: interfacesOnly,
    })
}

func typesAction(ctx context.Context, args TypesArgs) result.Result[cli.Output] {
    fset := token.NewFileSet()
    pkgNode, err := extract.ExtractDirectoryConcurrent(fset, args.FilePath, true, 0)
    if err != nil {
        return result.Ok(cli.Output{
            Text:     fmt.Sprintf("Error: %v\n", err),
            ExitCode: 1,
        })
    }

    var types []extract.TypeDeclaration
    if args.StructsOnly {
        types = extract.ExtractAllStructs(pkgNode)
    } else if args.InterfacesOnly {
        types = extract.ExtractAllInterfaces(pkgNode)
    } else {
        types = extract.ExtractAllStructsAndInterfaces(pkgNode)
    }

    gen := codegen.NewGenerator(fset)
    code, err := gen.GenerateTypesOnly(
        pkgNode.Name,
        types,
        pkgNode.Deps.Imports.ToSlice(),
    )
    if err != nil {
        return result.Ok(cli.Output{
            Text:     fmt.Sprintf("Error: %v\n", err),
            ExitCode: 1,
        })
    }

    if args.OutputFile != "" {
        if err := os.WriteFile(args.OutputFile, []byte(code), 0644); err != nil {
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

    return result.Ok(cli.Output{Text: code, ExitCode: 0})
}


