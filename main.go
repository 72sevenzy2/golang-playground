package main

import (
	"context"
	"fmt"
	"time"
)
 
func worker(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("stopped working")
			return
		
		case <- time.After(1 * time.Second):
			fmt.Println("working..")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(3 * time.Second)
	cancel()

	time.Sleep(500 * time.Millisecond)
}
