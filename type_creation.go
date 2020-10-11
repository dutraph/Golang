package main

import "fmt"

//cria um novo type hotDog com o type subjacente int
type hotDog int

var h hotDog = 42
var i int

func main() {

	fmt.Printf("O valor de h = %v\n", h)
	fmt.Printf("O type de h = %T\n", h)

	i = int(h) // converte o valor de h que é do type hotDog para int para atribuir esse valor para i
	fmt.Printf("O valor de i = %v\n", i)
	fmt.Printf("O type de i = %T\n", i)
}
