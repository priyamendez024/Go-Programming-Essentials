// Chapter 9: Channels & Synchronization Patterns
package main

import "fmt"

func main() {
    // Unbuffered
    ch := make(chan int)
    go func() {
        ch <- 42
    }()
    fmt.Println(<-ch)

    // Buffered
    buf := make(chan int, 3)
    buf <- 1
    buf <- 2
    buf <- 3
    for i := 0; i < 3; i++ {
        fmt.Println(<-buf)
    }

    // Select
    select {
    case v := <-ch:
        fmt.Println("Received", v)
    default:
        fmt.Println("No communication")
    }
}
