package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums)
	sum := 0
	for _, i := range nums {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum(a...)
}
