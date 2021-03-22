package main

import (
	"fmt"
	"github.com/1r0npipe/go-course2/task7-reflect/ast/count_calls"
	"log"
)

func main() {
	filename := "./test_file.go"
	funcName := "ManyGoCalls"
	num, err := countcalls.CountAsyncCalls(filename, funcName)
	if err != nil {
		log.Fatal("Cant count the number of go func due to %+v\n", err)
	}
	fmt.Printf("Number of go calls in file: \"%s\" at function: \"%s\" is %d\n",
		filename,
		funcName,
		num)
}
