package main

import (
	"fmt"
	"sync"
)

// Job ...
type Job struct {
	ID int
}

func main() {
	var jobs []Job
	for i := 0; i < 2; i++ {
		jobs = append(jobs, Job{i})
	}

	wait := sync.WaitGroup{}

	for i := range jobs {
		wait.Add(1)
		go func(job *Job) {
			defer wait.Done()
			fmt.Println(job.ID)
		}(&jobs[i])
	}

	wait.Wait()
}
