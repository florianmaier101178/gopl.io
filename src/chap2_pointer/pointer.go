package main

import "fmt"

func main() {
	x := 2
	y := &x
	var z *int
	z = &x
	fmt.Printf("x=%d\n", x)
	fmt.Printf("y=%d\n", *y)
	fmt.Printf("z=%d\n", *z)
	x = 4
	fmt.Printf("x=%d\n", x)
	fmt.Printf("y=%d\n", *y)
	fmt.Printf("z=%d\n", *z)
	*y = 73
	fmt.Printf("x=%d\n", x)
	fmt.Printf("y=%d\n", *y)
	fmt.Printf("z=%d\n", *z)
	*z = 1012
	fmt.Printf("x=%d\n", x)
	fmt.Printf("y=%d\n", *y)
	fmt.Printf("z=%d\n", *z)
}
