package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2

	fmt.Println(m)

	v1 := m["a"]
	fmt.Println(v1)
	fmt.Println("length = ", len(m))

	delete(m, "b")

	fmt.Println(m)

	_, prs := m["b"]
	fmt.Println(prs) // value is present or not

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	o := map[string][]int{"a": {1, 2, 3}, "b": {4, 5, 6}}
	fmt.Println(o)

}
