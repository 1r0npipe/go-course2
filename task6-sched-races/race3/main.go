package main

import "sync"

func main() {
	m := make(map[int]string, 1)
	var wg sync.WaitGroup
	var attempt int = 1000
	for i := 0; i < int(attempt); i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m[0] = "testM1"
		}()
	}
	for i := 0; i < attempt; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m[0] = "testM2"
		}()
	}
	wg.Wait()
}
