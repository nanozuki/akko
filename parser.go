package akko

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
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
	pkgName := ""
	for pkg := range pkgs {
		if !strings.HasPrefix(pkg, "_test") {
			pkgCount++
			pkgName = pkg
		}
	}
	if len(pkgs) > 1 {
		return nil, fmt.Errorf("more than 1 package found in dir %s", dir)
	}
	return parsePackage(pkgs[pkgName], structName)
}

func parsePackage(pkg *ast.Package, structName string) (*ServiceInfo, error) {
	var serviceDecl *ast.TypeSpec
	var methods []*ast.FuncDecl
	ast.Inspect(pkg, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.TypeSpec:
			if _, ok := node.Type.(*ast.StructType); ok && node.Name.Name == structName {
				serviceDecl = node
			}
		case *ast.FuncDecl:
			if node.Recv != nil {
				switch recv := node.Recv.List[0].Type.(type) {
				case *ast.StarExpr:
					if recv.X.(*ast.Ident).Name == structName {
						methods = append(methods, node)
					}
				case *ast.Ident:
					if recv.Obj.Name == structName {
						methods = append(methods, node)
					}
				}
			}
		}
		return false
	})
	if serviceDecl == nil || len(methods) == 0 {
		return nil, fmt.Errorf("struct %s not found", structName)
	}
	si := ServiceInfo{
		Typ: TypeInfo{
			Name:      structName,
			PkgPath:   "", // TODO
			Indentity: fmt.Sprintf("%s.%s", pkg.Name, structName),
			Kind:      reflect.Struct, // TODO
		},
	}
	for _, method := range methods {
		makeMethodInfo(&si, serviceDecl, method)
	}
	return &si, nil
}

func makeMethodInfo(si *ServiceInfo, serviceDecl *ast.TypeSpec, method *ast.FuncDecl) {
	mi := MethodInfo{
		Name:       method.Name.Name,
		Inputs:     []ParamInfo{},
		Returns:    []ParamInfo{},
		HTTPMethod: "",
		URL:        "",
		Process:    []Process{},
	}
	si.Methods = append(si.Methods, mi)
}
