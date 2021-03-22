package main

import (
	"fmt"
)

// ManyGoCalls just example of function where
// we will count amount of go calls
func ManyGoCalls() {
	go func() {
		fmt.Println("Test printout 1")
	}()
	go func() {
		fmt.Println("Test printout 2")
	}()
	go func() {
		fmt.Println("Test printout 3")
	}()
	go func() {
		fmt.Println("Print something")
	}()
	go func() {
		fmt.Println("Final print :)")
	}()
}
