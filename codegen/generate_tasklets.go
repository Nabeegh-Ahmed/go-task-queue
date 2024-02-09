package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

func main() {
	dir := "tasklets"

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer

	// Begin the registry map
	buf.WriteString("package tasklets\n\n")
	buf.WriteString("type AnyFunc func(...interface{}) (interface{}, error)\n\n")
	buf.WriteString("var TaskRegistry = map[string]AnyFunc{\n")

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if fn, isFn := decl.(*ast.FuncDecl); isFn {
					// Check for the special comment
					if hasDynamicComment(fn.Doc) {
						funcName := fn.Name.Name
						fmt.Fprintf(&buf, "\t\"%s\": %s,\n", funcName, funcName)
					}
				}
			}
		}
	}

	// End the registry map
	buf.WriteString("}\n")

	// Write the generated code to a file
	if err := ioutil.WriteFile("tasklets/task_registry_gen.go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}

	fmt.Println("Generated task_registry_gen.go")
}

// hasDynamicComment checks if the comment group contains the special tag.
func hasDynamicComment(cg *ast.CommentGroup) bool {
	if cg == nil {
		return false
	}
	for _, c := range cg.List {
		if strings.Contains(c.Text, "go:tasklet") {
			return true
		}
	}
	return false
}
