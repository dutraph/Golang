package main

import "fmt"

type tipoInt int

var x tipoInt

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x)

}
