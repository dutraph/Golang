package main

import "fmt"

func main() {
	x := 0
	for {
		if x < 10 {
			x++
			fmt.Println("X menor que 10")
		} else {
			fmt.Println("X maior que 10, saindo do loop")
			break
		}
	}
}
