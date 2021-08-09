package main

import (
	"fmt"
)

func test() {
	fmt.Println("Hello")
}

func main() {
	fmt.Println("HELLO")

	x := test
	x()
}
