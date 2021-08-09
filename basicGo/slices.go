package main

import "fmt"

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var slic []int = x[1:3]
	fmt.Println("array = ", x)
	fmt.Println("slice = ", slic)
	fmt.Println("length = ", len(slic))
	fmt.Println("capacity = ", cap(slic))

	// capacity means how many more elements can be possible to extend
	// after 2 3 we have 4 5 thats why capacity is 4
	// extending the slice
	fmt.Println("Extended slice = ", slic[:cap(slic)])

	// making slices directly
	var a []int = []int{1, 2, 3, 4, 5}
	fmt.Println(cap(a[:3]))

	b := append(a, 10)
	fmt.Println(b)

	a = append(b, 12)
	fmt.Println(a)

	//slices using make

	makeSlice := make([]int, 5)
	fmt.Println(makeSlice)

}
