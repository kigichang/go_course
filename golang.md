---
presentation:
  enableSpeakerNotes: true
  theme: serif.css
  slideNumber: true
  controls: true
  margin: 0
  width: 1920
  height: 1080
  center: true
  showNotes: true
---

<!-- slide -->

# Golang (Go) 簡介

<!-- slide -->

## Agenda

- Go History
- Projects and Companies Using Go
- Go Pros and Cons
- Go in Cyberon

<!-- slide -->

## Go History

- Created in Nov. 2009
- 1.0 in March 2012
- Last Version: 1.11 released on 2018/08/24
  - WebAssembly
  - Modules (package management)

<!-- slide -->

## Go Authors

- Robert Griesemer <!-- .element: class="fragment" data-fragment-index="1" -->
- Rob Pike <!-- .element: class="fragment" data-fragment-index="2" -->
  - UTF8 <!-- .element: class="fragment" data-fragment-index="3" -->
- Ken Thompson <!-- .element: class="fragment" data-fragment-index="4" -->
  - UTF8 <!-- .element: class="fragment" data-fragment-index="5" -->
  - **B Language** <!-- .element: class="fragment" data-fragment-index="6" -->

<!-- slide -->

## Companies Using Go

- BBC
- CoreOS
- Dropbox
- Google
- Hyperledger Fabric
- MongoDB
- Netflix
- Twitch.tv
- Uber

<!-- slide -->

## Projects Using Go

- Dockers
- Kubernetes
- OpenShift

<!-- slide -->

## Go Characteristics

- Statically typed, Compiled language (like C)
- Memory safety
- Garbage Collection
- CSP concurrency
- UTF8 encoding

<!-- slide -->

## Go Packages

- I/O
- Crypto
- SQL
- JSON/XML
- Network
- Text/HTML Template
- Process/Signal
- UTF8/UTF16

<!-- slide -->

## Go Prons

- Easy to Learn
- High Performance
- Fast to Deploy
- Concurrency

<!-- slide -->

## Go Cons

- Bad Error Handling
- No Inheritance
- No Generic
- Lack of Extensibility
  - Operator Overloading

<!-- slide -->

## Go Summary

### Go is C with Garbage Collection

<!-- slide -->

## Cross Platform

### Hello World

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, 世界")
}
```

<!-- slide -->

## Go Concurrency

- Channel
- Go Routine
- demo

<!-- slide -->

## Go in Cyberon

- Web Tool
- Corpus Crawler
- Go call STT libary

<!-- slide -->

```go
package main

import (
  "log"
  "sync"
)

func main() {

  intChannel := make(chan int)
  wait := &sync.WaitGroup{}

  wait.Add(1)
  go func() {
    for i := 0; i < 10; i++ {
      intChannel <- i
      log.Println("write ", i)
    }

    close(intChannel)
    wait.Done()
  }()

  wait.Add(1)
  go func() {
    for x := range intChannel {
      log.Println("got ", x)
    }
    wait.Done()
  }()

  wait.Wait()
  log.Println("end")
}
```
