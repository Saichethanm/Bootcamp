package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) // int of base 10 and size 64
	// fmt.Printf("You typed: %q", input)
	fmt.Printf("You will be %d years old at the end of 2021 ", 2021-input)

}
