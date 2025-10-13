package analyze

import (
	"go/ast"

	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/purekernels/pkg/fold"
	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
)

// InterfaceChecker checks interface implementations
type InterfaceChecker struct {
	graph DependencyGraph
}

// NewInterfaceChecker creates a new checker
func NewInterfaceChecker(graph DependencyGraph) InterfaceChecker {
	return InterfaceChecker{graph: graph}
}

// CheckImplementation verifies if type implements interface
func (c InterfaceChecker) CheckImplementation(
	typeName string,
	interfaceName string,
	methods []astpkg.MethodNode,
) astpkg.InterfaceImplementation {
	// Get interface declaration
	ifaceDecl, ok := c.graph.Decls[interfaceName]
	if !ok {
		return astpkg.InterfaceImplementation{
			TypeName:       typeName,
			InterfaceName:  interfaceName,
			Methods:        []astpkg.MethodNode{},
			MissingMethods: []string{},
		}
	}

	// Extract required methods from interface
	requiredMethods := c.extractInterfaceMethodNames(ifaceDecl.Decl)

	// Find implemented methods
	implementedNames := monoid.NewSetMonoid[string]()
	typeMethods := []astpkg.MethodNode{}

	for _, method := range methods {
		if method.ReceiverType == typeName {
			implementedNames = implementedNames.Insert(method.MethodName)
			typeMethods = append(typeMethods, method)
		}
	}

	// Find missing methods
	missingMethods := fold.Filter(
		func(required string) bool {
			return !implementedNames.Contains(required)
		},
		requiredMethods,
	)

	return astpkg.InterfaceImplementation{
		TypeName:       typeName,
		InterfaceName:  interfaceName,
		Methods:        typeMethods,
		MissingMethods: missingMethods,
	}
}

// extractInterfaceMethodNames gets method names from interface
func (c InterfaceChecker) extractInterfaceMethodNames(decl ast.Decl) []string {
	genDecl, ok := decl.(*ast.GenDecl)
	if !ok {
		return []string{}
	}

	methods := []string{}

	for _, spec := range genDecl.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		ifaceType, ok := typeSpec.Type.(*ast.InterfaceType)
		if !ok {
			continue
		}

		if ifaceType.Methods != nil {
			for _, method := range ifaceType.Methods.List {
				if len(method.Names) > 0 {
					methods = append(methods, method.Names[0].Name)
				}
			}
		}
	}

	return methods
}

// FindImplementors finds all types implementing an interface
func (c InterfaceChecker) FindImplementors(
	interfaceName string,
	allMethods map[string][]astpkg.MethodNode,
) []string {
	implementors := []string{}

	for typeName, methods := range allMethods {
		impl := c.CheckImplementation(typeName, interfaceName, methods)
		if len(impl.MissingMethods) == 0 && len(impl.Methods) > 0 {
			implementors = append(implementors, typeName)
		}
	}

	return implementors
}
