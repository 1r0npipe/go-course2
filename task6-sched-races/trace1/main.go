// Реализуйте функцию для разблокировки мьютекса с помощью defer
package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

const (
	maxGoRoutine = 50000
)

func main() {
	counter := 0
	trace.Start(os.Stderr)
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
	trace.Stop()
}
