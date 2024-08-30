package gogo

import (
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "os"
)

// FuncPrint prints all functions in the given Go source file.
func FuncPrint(filename string) error {
    fset := token.NewFileSet()
    node, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
    if err != nil {
        return err
    }

    for _, decl := range node.Decls {
        if fn, isFn := decl.(*ast.FuncDecl); isFn {
            fmt.Println("Function:", fn.Name.Name)
        }
    }
    return nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run funcprint.go <filename>")
        return
    }
    filename := os.Args[1]
    if err := FuncPrint(filename); err != nil {
        fmt.Println("Error:", err)
    }
}
