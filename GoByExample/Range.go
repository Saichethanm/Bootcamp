package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}

	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	mp := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range mp {
		fmt.Println(key, ":", value)
	}

	for i, c := range "abcdefghijklmnopqrstuvwxyz" {
		fmt.Println(i+1, ":", c)
	}

}
