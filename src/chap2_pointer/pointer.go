package main

import "fmt"

func main() {
    x := 2
    y := &x
    fmt.Printf("x=%d\n", x)
    fmt.Printf("y=%d\n", *y)
    x = 4
    fmt.Printf("x=%d\n", x)
    fmt.Printf("y=%d\n", *y)
    *y = 73
    fmt.Printf("x=%d\n", x)
    fmt.Printf("y=%d\n", *y)
}
