package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

//Job ...
type Job struct {
	jobID  int
	number float64
}

//Result ...
type Result struct {
	job      Job
	sqrt     float64
	workerID int
}

func main() {
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)
	done := make(chan bool)
	noOfJobs := 50

	//Adding jobs to jobs channel
	go func() {
		for jobID := 1; jobID <= noOfJobs; jobID++ {
			jobs <- Job{jobID, float64(rand.Intn(1000))}
		}
		close(jobs)
	}()

	//Getting result from the results channel
	go func(done chan bool) {
		for r := range results {
			fmt.Printf("Worker ID: %d\tJob Id: %d\tNumber: %f\tSqrt: %f\n", r.workerID, r.job.jobID, r.job.number, r.sqrt)
		}
		done <- true
	}(done)

	//Using waitgroup to pause execution of main go routine
	var wg sync.WaitGroup
	noOfWorkers := 5
	for w := 1; w <= noOfWorkers; w++ {
		wg.Add(1) //Increase the counter
		go worker(w, jobs, results, &wg)
	}
	wg.Wait() //Wait till all goroutine/workers returns
	close(results)
	<-done // wait till results goroutine finish its work
}

//This defines what work to be done by worker
//In this case our workers has to square root the give number
func worker(workerID int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	for job := range jobs {
		sqrt := math.Sqrt(job.number)
		time.Sleep(2000 * time.Millisecond) //Just to increase time of work
		results <- Result{job, sqrt, workerID}
	}
	wg.Done() //Decrease the counter added by wg.Add(1)
}
