package main

import "fmt"

func main() {
	var a = "one"
	fmt.Println(a)

	var b, c int = 4, 5
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple" // equivalent to => var f string = "apple"
	// var(keyword) variable_name data_type = ----
	fmt.Println(f)
}
