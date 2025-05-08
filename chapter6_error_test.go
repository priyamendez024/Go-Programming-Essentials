// Chapter 6: Error Handling, Testing & Benchmarking
package main

import (
    "errors"
    "fmt"
    "testing"
)

// Custom error
var ErrNotFound = errors.New("item not found")

type ValidationError struct {
    Field, Msg string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func TestAdd(t *testing.T) {
    got := add(2, 3)
    want := 5
    if got != want {
        t.Errorf("add(2,3) = %d; want %d", got, want)
    }
}

func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = fibonacci(20) // assume fibonacci defined elsewhere
    }
}
