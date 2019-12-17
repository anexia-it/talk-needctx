package main

import (
	"context"
	"time"
)

func startChildGoroutine(stopContext context.Context) context.Context {
	doneCtx, done := context.WithCancel(context.Background())

	go func() {
		defer done()
		println("child started...")
		<-stopContext.Done()
		println("child ended: ", stopContext.Err().Error())
	}()

	return doneCtx
}

func startGoroutine(stopContext context.Context) context.Context {
	doneCtx, done := context.WithCancel(context.Background())

	childStopCtx, childDone := context.WithTimeout(stopContext, time.Second*2)

	childCtx := startChildGoroutine(childStopCtx)

	go func() {
		defer childDone()
		defer done()

		println("goroutine started...")
		for {
			time.Sleep(time.Millisecond * 500)
			println("goroutine tick")
			select {
			case <-stopContext.Done():
				println("waiting for child...")
				<-childCtx.Done()
				println("goroutine exit!")
				return
			default:

			}
		}
	}()

	return doneCtx
}

func main() {
	topContext, stop := context.WithCancel(context.Background())
	defer stop()

	doneCtx := startGoroutine(topContext)

	println("waiting for 3 seconds...")
	time.Sleep(time.Second * 3)
	println("stopping goroutine")
	stop()

	select {
	case <-doneCtx.Done():
		println("goroutine stopped: ", doneCtx.Err().Error())
	case <-time.After(time.Second * 2):
		println("stopping timed out")
	}

}
