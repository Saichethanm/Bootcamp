package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	fmt.Println("length = ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			fmt.Print(c[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Print(c)
	fmt.Println("\n---")
	fmt.Println(c)

}
