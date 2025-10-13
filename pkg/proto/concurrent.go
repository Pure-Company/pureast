package proto

import (
	astpkg "github.com/vinodhalaharvi/pureast/pkg/ast"
	"github.com/vinodhalaharvi/purekernels/pkg/functor"
)

// GenerateProtoFromMultiplePackages generates proto from multiple packages (applicative!)
func GenerateProtoFromMultiplePackages(
	packages []astpkg.PackageNode,
	typeFilter []string,
	workers int,
) functor.Concurrent[ProtoFile] {
	protoMonoid := NewProtoFileMonoid()

	// Use TraverseConcurrent to process packages in parallel!
	return functor.TraverseConcurrent(
		protoMonoid,
		func(pkg astpkg.PackageNode) ProtoFile {
			return GenerateProtoFromPackageConcurrent(pkg, typeFilter, workers).Value()
		},
		packages,
		workers,
	)
}

// GenerateProtoFromPackageConcurrent generates proto concurrently (applicative!)
func GenerateProtoFromPackageConcurrent(
	pkgNode astpkg.PackageNode,
	typeFilter []string,
	workers int,
) functor.Concurrent[ProtoFile] {
	// Just use the sequential version for now since conflict detection needs all files
	return functor.NewConcurrent(
		NewProtoFileMonoid(),
		func() ProtoFile {
			return GenerateProtoFromPackage(pkgNode, typeFilter)
		},
	)
}
