// Реализуйте функцию для разблокировки мьютекса с помощью defer
package main

import (
	"fmt"
	"sync"
)

const (
	maxGoRoutine = 1000
)

func main() {
	counter := 0
	var mtx sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < maxGoRoutine; i++ {
		wg.Add(1)
		go func() {
			defer mtx.Unlock()
			mtx.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter is", counter)
}
