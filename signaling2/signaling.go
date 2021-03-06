package main

import (
	"context"
	"time"
)

func startGoroutine(stopContext context.Context) context.Context {
	doneCtx, done := context.WithCancel(context.Background())
	go func() {
		defer done()
		println("goroutine started...")
		<-stopContext.Done()

		time.Sleep(time.Second * 5)
		println("goroutine ended...")
	}()

	return doneCtx
}

func main() {
	stopContext, stop := context.WithCancel(context.Background())
	defer stop()

	doneCtx := startGoroutine(stopContext)

	println("waiting for 3 seconds...")
	time.Sleep(time.Second * 3)
	println("stopping goroutine")
	stop()

	select {
	case <-doneCtx.Done():
		println("goroutine stopped")
	case <-time.After(time.Second):
		println("stopping timed out")
	}

}
