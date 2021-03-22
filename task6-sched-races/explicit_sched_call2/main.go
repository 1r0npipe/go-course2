package main

import (
	"math"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	var counter uint = 0
	trace.Start(os.Stderr)
	go func() {
		for i := 0; ; i += 1 {
			_ = math.Cos(float64(i))
			if i%100 == 0 {
				runtime.Gosched()
			}
		}
	}()
	for i := 0; i < 1e10; i += 1 {
		counter += 1
	}
	trace.Stop()
}
