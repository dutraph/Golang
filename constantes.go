package main

import "fmt"

const (
	x = 10
)

// OR
const xx = 10

// var y = 10

var c int

// var d float64

func main() {
	// d = y //cannot use y (type int) as type float64 in assignment
	c = x // Consts only choose a type at the usage moment.
	fmt.Println(c)
}
