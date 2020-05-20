# Worker Pool

Worker pool is one of the most important implementations of buffered channels, goroutines and WeightGroups.
In general, worker pool is a collection of threads which waiting for the task to be assigned and once they finish their work they make themselves available for the next task.

## Worker Pool Implementation

In this example, our workers will be calculating the square root of given random numbers.

**Creating and assigning jobs/tasks**

We've `Job struct` which contains `jobID of type int` to store the id of job and 
`number of type float64` to store number whose square root, workers will calculate.

```go
type Job struct {
    jobID  int
    number float64
}
```

Creating `jobs buffered channel` of type `Job` and buffer size of `10` through which we'll be transporting jobs to the worker goroutine.
Buffer size is of 10 it means it will block channel once the buffer has 10 jobs and will be unblocked once the buffer has space for a new job.

```go
jobs := make(chan Job, 10)
```

We are running goroutine to send jobs to the channel.

```go
go func() {
    for jobID := 1; jobID <= noOfJobs; jobID++ {
        jobs <- Job{jobID, float64(rand.Intn(1000))}
    }
    close(jobs)
}()
```

**Collecting Results**

We've `Result struct` whose fields are `job of type Job` to store job, `sqrt of type float64` to store the square root of the number 
and `workerID of type int` to store which worker worked on this task.

```go
//Result ...
type Result struct {
    job      Job
    sqrt     float64
    workerID int
}
```

Creating `results buffered channel` of type `Result` and buffer size of `10` through which we'll be transporting results to the worker goroutine.
Buffer size is of 10 it means it will block channel once the buffer has 10 results and will be unblocked once the buffer has space for new results.

```go
results := make(chan Result, 10)
```

We are running goroutine to range over the results channel and print it.

`done bool channel` will be set `true` after all results are received from `results` channel.
```go
go func(done chan bool) {
    for r := range results {
        fmt.Printf("Worker ID: %d\tJob Id: %d\tNumber: %f\tSqrt: %f\n", r.workerID, r.job.jobID, r.job.number, r.sqrt)
    }
    done <- true
}(done)
```

**Worker goroutine**

Worker goroutine has parameters, `workerID` of type `int`, `receive-only channel jobs`, `send-only channel results` and `wg pointer of WaitGroup` for synchronizing goroutines.

A worker receives a job from the jobs channel, and then calculates the square root of a number and sends job, square root and worker ID to result channel.

```go
func worker(workerID int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
    for job := range jobs {
        sqrt := math.Sqrt(job.number)
        time.Sleep(2000 * time.Millisecond) //Just to increase time of work
        results <- Result{job, sqrt, workerID}
    }
    wg.Done() //Decrease the counter added by wg.Add(1)
}
```

`wg.Done()` decrements the counter.

**Calling worker in main**

We create zero value variable of WaitGroup. We are creating 5 worker threads to perform the job. `wg.Add(1)` increments the counter by one each time we create worker thread. `wg.Wait()` blocks main thread execution till counter becomes `zero`.

```go
var wg sync.WaitGroup
noOfWorkers := 5
for w := 1; w <= noOfWorkers; w++ {
    wg.Add(1) //Increase the counter
    go worker(w, jobs, results, &wg)
}
wg.Wait()
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```