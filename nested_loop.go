package main

import "fmt"

func main() {

	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês: ", mes)
		for dia := 1; dia <= 30; dia++ {
			fmt.Print("Dia: ", dia, " ")
		}
		fmt.Println("\n")
	}
}
