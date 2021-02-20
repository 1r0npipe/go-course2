package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type myErrorType struct {
	errorMessage string
	time         time.Time
}

func (e *myErrorType) Error() string {
	return "The error occurs: " + e.errorMessage + " with timestamp: " + e.time.Format("15:04:05.00000")
}

func recovering() {
	if err := recover(); err != nil {
		log.Println("Panic occurs here: ", err)
	}
	fmt.Println("Time stamp of panic: " + time.Now().Format("15:04:05.00000"))
}

func panicOutOfRange() {
	defer recovering()
	var array = [3]int{1, 2, 3}
	for i := 0; i <= 3; i++ {
		fmt.Printf("%v", array[i])
	}
}

func recoveringWithTime() {
	if smth := recover(); smth != nil {
		if err, ok := smth.(myErrorType); ok {
			fmt.Println(err.Error())
		}
	}
}

func panicManually(flag bool) {
	defer recoveringWithTime()
	var myTime myErrorType
	if flag {
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
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("File cannot be created, because of %v", err)
	}

	defer file.Close()
	_, _ = fmt.Fprintln(file, "Test data")
}
