// Chapter 5: Functions, Methods & Interfaces
package main

import (
    "fmt"
    "math"
)

// Function
func add(x, y int) int {
    return x + y
}

// Multiple return values
func divide(dividend, divisor float64) (quotient, remainder float64) {
    quotient = dividend / divisor
    remainder = math.Mod(dividend, divisor)
    return
}

// Methods
type Rectangle struct{ Width, Height float64 }

func (r *Rectangle) Scale(f float64) {
    r.Width *= f
    r.Height *= f
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Interfaces
type Shape interface {
    Area() float64
}

func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    fmt.Println("Add:", add(2, 3))
    rect := Rectangle{3, 4}
    fmt.Println("Area before:", rect.Area())
    rect.Scale(2)
    fmt.Println("Area after:", rect.Area())
    printArea(rect)
}
