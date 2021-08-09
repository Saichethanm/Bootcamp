package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)

	a := [3]int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(len(a))

	sum := 0

	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	fmt.Println(sum)

	a2D := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(a2D)
	fmt.Println(len(a2D))
}
