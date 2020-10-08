package main

import "fmt"

func main() {
	var nome, cargo, nome_chefe string
	var idade int
	var salario float64
	fmt.Print("Informe seu nome: ")
	fmt.Scan(&nome)
	fmt.Print("Informe sua idade: ")
	fmt.Scan(&idade)
	fmt.Print("Certo, agora informe seu cargo: \n")
	fmt.Scan(&cargo)
	fmt.Print("Certo, agora informe o nome do seu chefe: \n")
	fmt.Scan(&nome_chefe)
	fmt.Print("E por último, digite seu salário: ")
	fmt.Scan(&salario)
	fmt.Printf("Os dado informados foram:\nNome: %s\nIdade: %d\nCargo: %s\nNome do chefe: %s\nSalário: %v\n", nome, idade, cargo, nome_chefe, salario)
}
