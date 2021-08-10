package main

import "fmt"

func zeroVal(i int) {
	i = 0
}

func zeroPtr(i *int) {
	*i = 0
}

func main() {
	val := 25
	zeroVal(val)
	fmt.Println(val)
	zeroPtr(&val)
	fmt.Println(val)
}
