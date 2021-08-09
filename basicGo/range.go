package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for i, value := range a {
		fmt.Println(i, ":", value)
	}
}
