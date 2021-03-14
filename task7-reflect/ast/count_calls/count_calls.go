package countcalls

import (
	"fmt"
	"go/parser"
	"go/token"
)

// CountAsyncCalls count all async call (goroutines) in fileName for funcName
// function inside of fileName, returns numbers of calls or error
func CountAsyncCalls(fileName, funcName string) (int, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileName, nil, 0)
	if err != nil {
		return 0, fmt.Errorf("Can't read file")
	}

	return countAsycCalls, nil
}
