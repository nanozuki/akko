package akko

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// ParseService parse a golang package, find struct with name structName, parse it into ServiceInfo, and return it.
func ParseService(dir string, structName string) (*ServiceInfo, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse dir: %w", err)
	}
	pkgCount := 0
	for pkg := range pkgs {
		if !strings.HasPrefix(pkg, "_test") {
			pkgCount++
		}
	}
	if len(pkgs) > 1 {
		return nil, fmt.Errorf("more than 1 package found in dir %s", dir)
	}
	return nil, nil
}

func parsePackage(pkg *ast.Package, structName string) (*ServiceInfo, error) {
	var serviceDecl *ast.TypeSpec
	ast.Inspect(pkg, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.TypeSpec:
			st, ok := node.Type.(*ast.StructType)
			if ok && node.Name.Name == structName {
				serviceDecl = node
				return false
			}
		case *ast.FuncDecl:
			if node.Recv != nil && node.Recv.List[0].
		}
		return false
	})
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			switch decl := decl.(type) {
			case *ast.GenDecl:
				for _, spec := range decl.Specs {
					switch spec := spec.(type) {
					case *ast.TypeSpec:
						if spec.Name.Name == structName {
							return parseStruct(spec.Type.(*ast.StructType))
						}
					}
				}
			}
		}
	}
	return nil, fmt.Errorf("struct %s not found", structName)
}
