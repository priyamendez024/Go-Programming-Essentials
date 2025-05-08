// Chapter 4: Variables, Types & Control Flow
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    // Variables
    var a int = 10
    b, c := 3.14, "Go"
    d := true
    fmt.Println(a, b, c, d)

    // If statement
    if n := len(os.Args); n > 1 {
        fmt.Println("Args:", os.Args[1:])
    } else {
        fmt.Println("No args provided")
    }

    // For loops
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    for _, v := range []string{"a", "b", "c"} {
        fmt.Println(v)
    }

    // Switch
    switch day := time.Now().Weekday(); day {
    case time.Saturday, time.Sunday:
        fmt.Println("Weekend")
    default:
        fmt.Println("Weekday")
    }
}
