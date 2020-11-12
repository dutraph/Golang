package main

import "fmt"

func main() {

	x := 200
	fmt.Printf("Dec %d  -  Bin %b  -  Hex %#x", x, x, x)

	fmt.Println("")
	y := x << 1 // Add 1 Byte a esquerda
	fmt.Printf("Dec %d  -  Bin %b  -  Hex %#x", y, y, y)

}
