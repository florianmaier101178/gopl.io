package main

import "fmt"

var globalX int

func main() {
	globalX = 1
	fmt.Printf("main.globalX=%d\n", globalX)
	changeGlobalX()
	fmt.Printf("main.globalX=%d\n", globalX)
	passGlobalX(globalX)
	fmt.Printf("main.globalX=%d\n", globalX)
}

func changeGlobalX() {
    globalX = 2
    fmt.Printf("changeGlobalX.globalX=%d\n", globalX)
}

func passGlobalX(x int) {
    x = 3
    fmt.Printf("passGlobalX.globalX=%d\n", x)
}
