package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"
)

// ReadAll redirect channels to one out channel
func ReadAll(ctx context.Context, wait *sync.WaitGroup, channels ...chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	var cases []reflect.SelectCase
	cases = append(cases, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(ctx.Done()),
	})

	for _, c := range channels {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c),
		})
	}

	go func() {
		defer func() {
			close(out)
			wait.Done()
			log.Println("close out and done")
		}()

		for len(cases) > 1 {
			i, v, ok := reflect.Select(cases)
			log.Println(i, v, ok)
			if i == 0 { // timeout and exit
				log.Println("cancel !!!")
				return
			}
			if !ok { // some channel is closed and remove from cases
				cases = append(cases[:i], cases[i+1:]...)
			}

			out <- v.Interface()
		}
	}()

	return out
}

func main() {
	channels := []chan interface{}{
		make(chan interface{}),
		make(chan interface{}),
		make(chan interface{}),
	}

	defer func() {
		for _, c := range channels {
			close(c)
		}
		log.Println("close channels completed")
	}()

	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		log.Println("cancel context")
	}()
	wait := sync.WaitGroup{}
	wait.Add(1)
	out := ReadAll(timeout, &wait, channels...)

	// generate goroutines for each channel to write data
	for i, c := range channels {
		go func(a int, c chan<- interface{}) {
			for {
				<-time.After(time.Second)
				c <- fmt.Sprintf("%d:%s", a, time.Now().Format("2006-01-02 15:04:05"))
			}
		}(i, c)
	}

	// generate goroutine to read data from out channel
	wait.Add(1)
	go func() {
		defer wait.Done()
		for x := range out {
			log.Println("out got:", x)
		}
	}()

	// wait for goroutine in ReadAll
	wait.Wait()
	log.Println("end")

}
