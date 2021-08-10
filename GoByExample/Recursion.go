package main

import "fmt"

func reverseArray(a []int) {
	if len(a) == 1 {
		fmt.Println(a[0])
		return
	}
	reverseArray(a[1:])
	fmt.Println(a[0])
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	reverseArray(a)
}
