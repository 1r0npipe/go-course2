package main

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	// ErrorReadFromChannel error if can't read from channel
	ErrorReadFromChannel = errors.New("Can't read from channel")
	maxGoRoutine         = 1000
	maxCountInter        = 1000
)

func main() {
	runtime.GOMAXPROCS(8)
	ch := make(chan int)
	go func() {
		for c := range ch {
			fmt.Println(c)
		}
	}()
	fmt.Println("Main is finished")
	for i := 0; i < maxGoRoutine; i++ {
		go incThread(maxCountInter, ch)
	}

}

func incThread(n int, channel chan int) {
	var counter int = 0
	for ; counter < n; counter++ {
	}
	channel <- counter
	fmt.Println("Routine is finished")
	return
}
