package main

import "fmt"

type tipoInt int

var x tipoInt
var y int

func main() {
	fmt.Println("Valor Inicial de X")
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n\n", x)
	fmt.Println("Valor e X")
	x = 42
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n\n", x)

	fmt.Println("Valor e Y")
	y = int(x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%v\n", y)

}
