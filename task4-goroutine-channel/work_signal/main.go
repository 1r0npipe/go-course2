package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	tasks := make(chan struct{})
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	fmt.Println("PID-", os.Getpid())
	go func(ctx context.Context) {
		fmt.Println("Doing something...")
		<-signals
		tasks <- struct{}{}
	}(ctx)
	<-tasks
	<-time.After(1 * time.Second)
	fmt.Println("End of main...")
}
