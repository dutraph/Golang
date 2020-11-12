package main

import "fmt"

func main() {
	s := "Hello, playground é legal"
	sb := []byte(s)
	fmt.Println("Range Slice of byte ([]byte)")
	fmt.Println("BYTE - TYPE - UNICODE - HEXADECIMAL")
	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")
	fmt.Println("Range of String (s)")
	fmt.Println("BYTE - TYPE - UNICODE - HEXADECIMAL")
	for _, b := range s {
		fmt.Printf("%v - %T - %#U - %#x\n", b, b, b, b)
	}

	fmt.Println("")
	fmt.Println("for loop")
	fmt.Println("BYTE - TYPE - UNICODE - HEXADECIMAL")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])

	}

}
