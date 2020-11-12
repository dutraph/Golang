package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	if runtime.GOOS == "darwin" {
		fmt.Println("this computer brand is Apple")
	} else if runtime.GOOS == "linux" {
		fmt.Println("This computer is a Linux")
	}
}
