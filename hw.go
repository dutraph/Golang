package main

import "fmt"

//Para declaração PLS (Package Level Scope) usamos o = e nao :=
var y string = "Im a String"

//ou
var yy = "Im a String too"
var g int = 3
var j float64 = 3.2

func main() {
	/*
	* a funcao Println alem de imprimir na tela, retorna um int e erros, mas se nao quisermos que um desses valores sejao usados
	* devemos apenas usar um _ no seu lugar conhecido como Blank Identifier
	 */
	/*numdebytes*/
	_, erros := fmt.Println("Hello, playground", "Alo", 100)
	fmt.Println(erros)

	// Os tipos de dados primitivos em Go são (int, string e boolean )
	// Os tipos de dados compostos em Go são (slice, arraym struct e map)

	// := (operador de declaração, so pode ser usado dentro da function body)
	x := 2
	//y := "Strings"
	z := true
	fmt.Printf("x: Value = %v, Type = %T\n", x, x)
	fmt.Printf("y: Value = %v, Type = %T\n", y, y)
	fmt.Printf("z: Value = %v, Type = %T\n\n", z, z)

	// = (operador de atribuição)
	x = 30
	fmt.Printf("x: Value = %v, Type = %T\n", x, x)

	// Nao posso declarar 2 vezes a mesma variavel
	//x := 34
	//fmt.Printf("x: Value = %v, Type = %T\n", x, x)

	//Porem existe a essa exception, que so aceita a "re-declaração do x pois estamos tambem declarando k"
	x, k := 34, 43.3
	fmt.Printf("re-declared x: Value = %v, Type = %T\n", x, x)
	fmt.Printf("k: Value = %v, Type = %T\n", k, k)

	h := 2
	qualquercoisa(h)

	// O valor zero é o valor inicial da variavel antes de ser inicializada.
	//  - int: 0
	//  - floats: 0.0
	//  - booleans: false
	//  - strings: ""
	//  - pointers, functions, interfaces, slices, channels, maps: nil

	fmt.Println("Valores zero das variaveis...")
	valorZero(a, b, c, ç)

	// para utilizar o interpreted literal usamos " " e para raw literal ' '

	// Para atribuir o valor de j (float64) que é 3.2 para um int fazemos o seguinte
	h = int(j) // j na forma int vale 3, logo:. h = 3
	qualquercoisa(h)
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

func qualquercoisa(x int) {
	fmt.Println(y) // Isso funciona pois y foi declarado PLS (Escopo do nivel Pacote)
	fmt.Printf("Valor de qualquer coisa  = %v\n", x)
	if x == 0 {
		fmt.Println("é zero")
	} else {
		fmt.Printf("é %v\n", x)
	}
}
