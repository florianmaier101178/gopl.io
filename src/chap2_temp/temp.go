package main

import (
    "tempconv"
    "fmt"
)

func main() {
    var c tempconv.Celsius
    c = 80
    fmt.Printf("%g\n", c)
    fmt.Println(c.String())
    fmt.Printf("%s\n", c)
    fmt.Printf("%v\n", c)
}
