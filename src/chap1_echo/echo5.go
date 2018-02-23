package main

import (
    "fmt"
    "os"
)

func main() {
    sep := " "
    for i, arg := range os.Args[1:] {
        fmt.Println(i + 1, sep + arg)
    }
}
