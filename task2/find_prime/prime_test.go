package prime

import "fmt"

func Example() {
	var N = 23
	if IsPrime(N) {
		fmt.Printf("This is prime: %d", N)
	}
}
