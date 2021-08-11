package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plus3(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println(plus(3, 5))
	fmt.Println(plus3(4, 5, 6))
}
