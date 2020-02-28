package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {

	// create a new context with cancellation possibility
	// based on previous base context
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := operation1(ctx)
		// if there is an error -
		// cancel all operations within this context
		if err != nil {
			cancel()
		}
	}()

	operation2(ctx)
	operation3(ctx)
}

func operation1(ctx context.Context) error {
	time.Sleep(1 * time.Second)
	fmt.Println("operation1 failed")
	return errors.New("operation1 failed")
}

func operation2(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("operation2 done")
	case <-ctx.Done():
		fmt.Println("operation2 halted")
	}
}

func operation3(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("operation3 done")
	case <-ctx.Done():
		fmt.Println("operation3 halted")
	}
}
