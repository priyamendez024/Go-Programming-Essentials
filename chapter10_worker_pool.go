// Chapter 10: Advanced Concurrency - Worker Pools
package main

import (
    "context"
    "fmt"
    "sync"
)

type Job struct {
    ID int
}

type Result struct {
    Job    Job
    Output string
}

func process(job Job) string {
    return fmt.Sprintf("processed job %d", job.ID)
}

func RunWorkerPool(ctx context.Context, numWorkers int, jobs <-chan Job) <-chan Result {
    results := make(chan Result)
    var wg sync.WaitGroup
    wg.Add(numWorkers)

    for i := 0; i < numWorkers; i++ {
        go func(workerID int) {
            defer wg.Done()
            for {
                select {
                case <-ctx.Done():
                    return
                case job, ok := <-jobs:
                    if !ok {
                        return
                    }
                    output := process(job)
                    results <- Result{Job: job, Output: output}
                }
            }
        }(i)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    return results
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    jobs := make(chan Job)
    results := RunWorkerPool(ctx, 3, jobs)

    go func() {
        for i := 1; i <= 5; i++ {
            jobs <- Job{ID: i}
        }
        close(jobs)
    }()

    for r := range results {
        fmt.Println(r.Output)
    }
}
