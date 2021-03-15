package countcalls

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// CountAsyncCalls count all async call (goroutines) in fileName for funcName
// function inside of fileName, returns numbers of calls or error
func CountAsyncCalls(fileName, funcName string) (int, error) {
	countGoCalls := 0
	fSet := token.NewFileSet()
	fileAst, err := parser.ParseFile(fSet, fileName, nil, 0)
	if err != nil {
		return 0, fmt.Errorf("Can't read file, due to error %+v\n", err)
	}
	for _, decl := range fileAst.Decls {
		genDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if genDecl.Name.Name != funcName {
			continue
		}
		for _, node := range genDecl.Body.List {
			switch node.(type) {
			case *ast.GoStmt:
				countGoCalls += 1
			}
		}
	}
	return countGoCalls, nil
}
