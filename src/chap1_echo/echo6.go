package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Command: " + os.Args[0])
    for i, arg := range os.Args[1:] {
        fmt.Println(i + 1, " " + arg)
    }
}
