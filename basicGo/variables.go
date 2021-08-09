package main

import "fmt"

func main() {
	var number = 50.0
	number2 := 40
	fmt.Printf("%T\n", number)
	fmt.Println(number2)
	fmt.Printf("Number in binary: %b", number2)
	fmt.Printf("Number in octal: %o", number2)
	fmt.Printf("Number in dec: %d", number2)
	fmt.Printf("Number in hex: %x", number2)

	var out string = fmt.Sprintf("Number: %d", 2500)
	fmt.Println(out)
}
