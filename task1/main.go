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
	return fmt.Sprintf("The error occurs: %v with timestamp: %v",e.errorMessage ,e.time.Format("15:04:05.00000"))
}

func main() {
	panicOutOfRange()
	panicManually(true)
	readCreateFile()
}

func panicOutOfRange() {
	defer recovering()
	var array = [3]int{1, 2, 3}
	for i := 0; i <= 3; i++ {
		fmt.Printf("%v", array[i])
	}
}

func recovering() {
	if err := recover(); err != nil {
		log.Println("Panic occurs here: ", err)
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

func recoveringWithTime() {
	if smth := recover(); smth != nil {
		if err, ok := smth.(myErrorType); ok {
			fmt.Printf("The error \"%s\" occurs at: %v\n", err.errorMessage, err.time.Format("15:04:05.00000"))
		}
	}
}

func readCreateFile() error {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("File cannot be created, because of %v", err)
		return err
	}

	defer file.Close()
	_, _ = fmt.Fprintln(file, "Test data")
	return nil
}
