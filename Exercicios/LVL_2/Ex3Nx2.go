package main

import "fmt"

const (
	x int = 10
	y     = 2.4
)

func main() {

	fmt.Printf("%v - %T\n", x, x)
	fmt.Printf("%v - %T", y, y)

}
