package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%T\n%T\n%T\n", x, y, z)

	fmt.Println("Valores zero")
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}
