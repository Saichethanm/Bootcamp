package main

import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// func add(x, y int) (int, int) {
// 	return x + y, x - y
// }

func add(x, y int) (a int, b int) {
	defer fmt.Println("Hello")
	a = x + y
	b = x - y
	fmt.Println("Before return")
	return
}

func test() (int, int) {
	return 1, 2
}

func main() {
	fmt.Println(add(6, 7))
	fmt.Println(test())
}
