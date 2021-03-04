package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	maxGoRoutine         = 1000
	maxCountInter        = 1000
)

func main() {
	runtime.GOMAXPROCS(8)
	var wg sync.WaitGroup
	ch := make(chan int)
	go func(channel chan int) {
		for {
			if _, ok := <- channel; !ok {
				fmt.Println("Issue with reading from channel")
				return
			}
		}
	}(ch)
	for i := 0; i < maxGoRoutine; i++ {
		wg.Add(1)
		go incThread(i, maxCountInter, ch, &wg)
	}
	wg.Wait()
	fmt.Println("Main is finished")
}

func incThread(id int, n int, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var counter int = 0
	for ; counter < n; counter++ {
	}
	channel <- counter
	fmt.Printf("Goroutine finished with ID: %d\n", id)
	return
}
