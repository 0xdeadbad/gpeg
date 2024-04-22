package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src is the input for which we want to print the AST.
	src := `
package main
func main() {
        println("Hello, World!")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	for _, f := range f.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}

		fmt.Println(fn.Name.Name)
		for _, stmt := range fn.Body.List {
			fmt.Printf("\t%s\n", stmt)
		}
	}
}
