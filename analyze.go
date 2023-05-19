package akko

/*

import (
	"fmt"
	"go/ast"
	"go/token"
)

func analyzeFile(f *ast.File) {
	// Create an ast.FileSet object for the Golang file.
	fset := token.NewFileSet()

	// Create an ast.Package object for the Golang file.
	pkg := ast.NewPackage(fset, f.Name.String(), nil, nil)

	// Iterate over the ast.Package.Decls field to get all of the declarations in the Golang file.
	for _, decl := range pkg.Decls {
		// Check if the declaration is a function declaration.
		if fun, ok := decl.(*ast.FuncDecl); ok {
			// If the declaration is a function declaration, check if the function is a method of the service.Service type.
			if service, ok := fun.Recv.List[0].Type.(*ast.Ident); ok && service.Name == "Service" {
				// If the function is a method of the service.Service type, analyze the function.
				fmt.Println("Function name:", fun.Name.String())
				fmt.Println("Function receiver:", service.Name)
				fmt.Println("Function parameters:")
				for _, param := range fun.Type.Params.List {
					fmt.Println("  - Name:", param.Names[0].String())
					fmt.Println("    Type:", param.Type)
				}
				fmt.Println("Function results:")
				for _, result := range fun.Type.Results.List {
					fmt.Println("  - Name:", result.Names[0].String())
					fmt.Println("    Type:", result.Type)
				}

				// Get the function comment.
				comment := fun.Doc.Text()
				fmt.Println("Function comment:", comment)
			}
		}
	}
	// Create an ast.FileSet object for the Golang file.
	fset := token.NewFileSet()

	// Create an ast.Package object for the Golang file.
	pkg := ast.NewPackage(fset, f.Name.String(), nil, nil)

	// Iterate over the ast.Package.Decls field to get all of the declarations in the Golang file.
	for _, decl := range pkg.Decls {
		// Check if the declaration is a function declaration.
		if fun, ok := decl.(*ast.FuncDecl); ok {
			// If the declaration is a function declaration, check if the function is a method of the service.Service type.
			if service, ok := fun.Recv.List[0].Type.(*ast.Ident); ok && service.Name == "Service" {
				// If the function is a method of the service.Service type, analyze the function.
				fmt.Println("Function name:", fun.Name.String())
				fmt.Println("Function receiver:", service.Name)
				fmt.Println("Function parameters:")
				for _, param := range fun.Type.Params.List {
					fmt.Println("  - Name:", param.Names[0].String())
					fmt.Println("    Type:", param.Type)
				}
				fmt.Println("Function results:")
				for _, result := range fun.Type.Results.List {
					fmt.Println("  - Name:", result.Names[0].String())
					fmt.Println("    Type:", result.Type)
				}

				// Get the function comment.
				comment := fun.Doc.Text()
				fmt.Println("Function comment:", comment)
			}
		}
	}
}
*/
