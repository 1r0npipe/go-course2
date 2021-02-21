// Package prime implements the solution
//
//of getting all primes up to requested N
//
//This is just an example of using doc
package prime

import (
	"fmt"
	"math"
)

//IsPrime the fucntion which returns True if the number is prime
func IsPrime(a int) bool {
	if a <= 2 {
		return true
	}
	max := int(math.Sqrt(float64(a)))
	for j := 2; j <= max; j++ {
		if a%j == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int
	fmt.Print("Type any number, to find all primes to that:")
	fmt.Scanln(&num)
	fmt.Printf("The list of primes between 1 and %d\n", num)
	for i := 1; i < num; i++ {
		if IsPrime(i) {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println()
}
