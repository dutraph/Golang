package main

import "fmt"

func main() {
	//mostrando apenas os pares
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue // quando o resto da divisao for diferente de zero, interrompe a iteração e pula pra proxima
		}
		fmt.Println(i)
	}

	// Outro exemplo

	for j := 0; j < 20; j++ {
		if j == 3 {
			continue // quando j for igual a 3 ele interrompe a iteração e pula pra proxima.
		}
		fmt.Println(j)
	}

}
