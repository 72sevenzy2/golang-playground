package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // wait for cancel() func to stop the worker
			fmt.Println("stopped working")
			return  // exits before the printing happens so we need to schedule some time for print to execute.

		case <-time.After(1 * time.Second): // wait 1 second after each print msg:
			fmt.Println("working..")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // initialising the context cancellation

	go worker(ctx) // start the worker

	time.Sleep(3 * time.Second) // wait 3 seconds after calling the cancel() func
	cancel()

	time.Sleep(500 * time.Millisecond) // give some time for the "stopped working" msg to be printed
}
