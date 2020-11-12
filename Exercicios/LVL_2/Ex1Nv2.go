package main

import "fmt"

var num int

func main() {
	fmt.Print("Entre um numero para converter: ")
	fmt.Scan(&num)
	fmt.Printf(" Decimal %d\t - Binario %b\t - hex %#x\n", num, num, num)
}
