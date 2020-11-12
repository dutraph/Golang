package main

import "fmt"

func main() {

	for x := 33; x <= 122; x++ {
		fmt.Printf("Dec: %d\t Char Unicode: %v\n", x, string(x))
	}

}
