package main

import "fmt"

func main() {

	a, b, c := 42, "James Bond", true
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%v - %T\n%v - %T\n%v - %T\n", a, a, b, b, c, c)
}
