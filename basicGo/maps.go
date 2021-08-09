package main

import "fmt"

func main() {
	mp := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(mp)
	fmt.Println(mp["a"])
	fmt.Println(mp["100"])
	fmt.Println(mp)

	mp["a"] = 100
	fmt.Println(mp)

	val, ok := mp["a"]
	fmt.Println(val, ok)

	delete(mp, "b")
	fmt.Println(mp)

	mp2 := make(map[string]int)
	fmt.Println(mp2)

}
