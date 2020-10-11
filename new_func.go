package main

import "fmt"

var y int = 42
var j float64 = 3.2

func main() {
	// Para atribuir o valor de j (float64) que é 3.2 para um int fazemos o seguinte
	h = int(j) // j na forma int vale 3, logo:. h = 3
	qualquercoisa(h)

}

func qualquercoisa(x int) {
	fmt.Println(y) // Isso funciona pois y foi declarado Package-level Scope, fora de main()
	fmt.Println("O valor de qualquer coisa é = ", x)
	if x == 0 {
		fmt.Println("é zero")
	} else {
		fmt.Printf("é %v\n", x)
	}
}

var a int
var b float64
var c string
var ç bool

func valorZero(x int, y float64, z string, ç bool) {
	fmt.Printf("Int zero = %v\n", x)
	fmt.Printf("float zero = %v\n", y)
	fmt.Printf("String zero = %v\n", z)
	fmt.Printf("Boolean zero = %v\n", ç)
}
