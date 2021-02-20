package main

import (
	"fmt"
	"log"
	"time"
)

type myErrorType struct {
	errorMessage string
	time time.Time
}

func (e *myErrorType) myErrorTypePrint () error {
	return fmt.Errorf("The error occurs: %s, with timestamp: %s", e.errorMessage, e.time.String())
} 

func recovering() {
	if err := recover(); err != nil {
		log.Println("Panic occurs here: ", err)
	}
	fmt.Println("Time stamp of panic: " + time.Now().String())
}

func panicOutOfRange() {
	defer recovering()
	var array = [3]int{1, 2, 3}
	for i := 0; i <= 3; i++ {
		fmt.Printf("%v", array[i])
	}
}

func recoveringWithTime () {
	if smth := recover(); smth != nil {
		fmt.Println(smth)
	}
}

func panicManually(flag bool) {
	defer recoveringWithTime()
	var myTime *myErrorType
	if flag {
		fmt.Println("The panic manually created: ")
		myTime.errorMessage = "Manually called PANIC"
		myTime.time = time.Now()
		panic(myTime)
	}
}

func main() {
	fmt.Println("=== Start of program ===")
	panicOutOfRange()
	fmt.Println("===  ===")
	panicManually(true)
	fmt.Println("=== End of program ===")
}
