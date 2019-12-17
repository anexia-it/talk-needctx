package main

import (
	"context"
	"time"
)

func goroutine(ctx context.Context) {

	println("goroutine started...")
	<-ctx.Done()
	println("goroutine ended...")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go goroutine(ctx)

	println("waiting for 3 seconds...")
	time.Sleep(time.Second * 3)
	println("done")
}
