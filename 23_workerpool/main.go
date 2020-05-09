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
	id     int
	number float64
}

//Result ...
type Result struct {
	job  Job
	sqrt float64
}

//This defines what work to be done by worker
//In this case our workers has to square root the give number
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	for job := range jobs {
		sqrt := math.Sqrt(job.number)
		time.Sleep(500 * time.Millisecond) //Just to increase time of work
		results <- Result{job, sqrt}
	}
	wg.Done() //Decrease the counter added by wg.Add(1)
}

func main() {
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)
	noOfJobs := 50

	//Adding jobs to jobs channel
	go func() {
		for j := 1; j <= noOfJobs; j++ {
			jobs <- Job{j, float64(rand.Intn(1000))}
		}
		close(jobs)
	}()

	//Getting result from the results channel
	go func() {
		for r := range results {
			fmt.Printf("Job Id: %d\tNumber: %f\tSqrt: %f\n", r.job.id, r.job.number, r.sqrt)
		}
	}()

	//Using waitgroup to pause execution of main go routine
	var wg sync.WaitGroup
	noOfWorkers := 5
	for w := 1; w <= noOfWorkers; w++ {
		wg.Add(1) //Increase the counter
		go worker(w, jobs, results, &wg)
	}
	wg.Wait() //Wait till all goroutine/workers returns
	close(results)
}
