package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	tasks := make(chan struct{})
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		signal := <-signals
		log.Printf("system call: %+v", signal)
		cancelFunc()
	}()
	fmt.Println("PID-", os.Getpid())
	ctxShutDown, cancelServ := context.WithTimeout(context.Background(), 1*time.Second)
	go func(ctx context.Context) {
		fmt.Println("Doing something...")
		<-signals
		tasks <- struct{}{}
		defer func() { cancelServ() }()
	}(ctxShutDown)
	<-tasks
	<-ctx.Done()
	fmt.Println("End of main...")
}
