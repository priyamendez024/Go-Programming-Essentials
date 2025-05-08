// Chapter 8: Goroutines & Concurrency Fundamentals
package main

import (
    "fmt"
    "time"
)

func doWork(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    for i := 1; i <= 5; i++ {
        go doWork(i)
    }
    time.Sleep(2 * time.Second)
}
