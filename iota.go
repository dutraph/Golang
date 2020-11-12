package main

import "fmt"

const (
	// _ = iota
	// b = iota
	// _ = iota
	// x = iota
	// _ = iota
	// z = iota

	// Or

	f = iota + 1
	g
	_
	i
	_
	k
)

func main() {

	fmt.Println("Numa declaração de constantes, o identificador iota representa numeros sequenciais.")
	fmt.Println(f, g, i, k)
}
