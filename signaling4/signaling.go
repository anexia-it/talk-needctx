package main

import (
	"context"
	"runtime/debug"
	"sync"
	"time"
)

func goroutine(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if r := recover(); r != nil {
			println("recovered:", r)
			println(string(debug.Stack()))
		}
	}()
	panic("test")

	println("goroutine started...")
	<-ctx.Done()
	println("goroutine ended...")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go goroutine(ctx, wg.Done)
	
	println("waiting for 3 seconds...")
	time.Sleep(time.Second * 3)
	println("done")
	cancel()
	wg.Wait()
}
