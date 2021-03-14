package main

import (
	"fmt"
	"runtime"
)

var (
	maxGoRoutine = 1000
	result       = 0
)

func main() {
	runtime.GOMAXPROCS(0)
	ch := make(chan struct{})
	for i := 0; i < maxGoRoutine; i++ {
		go func() {
			result++
			ch <- struct{}{}
		}()
		<-ch
	}

	fmt.Println(result)
	fmt.Println("Main is finished")
}
