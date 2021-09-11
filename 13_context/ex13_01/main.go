package main

import (
	"context"
	"log"
	"time"
)

func contextDemo(ctx context.Context) {
	dealine, ok := ctx.Deadline()
	name := ctx.Value(contextKey("name"))

	if ok {
		log.Println(name, "has dealine:", dealine.Format("2006-01-02 15:04:05"))
	} else {
		log.Println(name, "does not have dealine")
	}
}

type contextKey string

func main() {
	timeout := 3 * time.Second
	deadline := time.Now().Add(10 * time.Second)

	timeoutContext, timeoutCancelFunc := context.WithTimeout(context.Background(), timeout)
	defer timeoutCancelFunc()

	cancelContext, cancelFunc := context.WithCancel(context.Background())

	deadlineContext, deadlineCancelFunc := context.WithDeadline(context.Background(), deadline)
	defer deadlineCancelFunc()

	contextDemo(context.WithValue(timeoutContext, contextKey("name"), "[Timeout Context]"))
	contextDemo(context.WithValue(cancelContext, contextKey("name"), "[Canncel Context]"))
	contextDemo(context.WithValue(deadlineContext, contextKey("name"), "[Deadline Context]"))

	<-timeoutContext.Done()
	log.Println("timeout ...")

	log.Println("cancel error:", cancelContext.Err())
	log.Println("canncel...")
	cancelFunc()
	log.Println("cancel error:", cancelContext.Err())

	<-cancelContext.Done()
	log.Println("The cancel context has been cancelled...")

	<-deadlineContext.Done()
	log.Println("The deadline context has been cancelled...")
}
