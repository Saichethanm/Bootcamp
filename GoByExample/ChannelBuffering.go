package main

import "fmt"

func main() {

	messages := make(chan []int, 2)

	messages <- []int{1, 2, 3}
	messages <- []int{4, 5, 6}

	fmt.Println("second ", <-messages)
	fmt.Println("first ", <-messages)
}
