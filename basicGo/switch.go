package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Enter a number: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	val, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	switch val {
	case 1, -1, 0, 5:
		fmt.Println("You can separate values with comma to check condition")
	case 2:
		fmt.Println("In go, break is automatically given after every case")
	default:
		fmt.Println("That's it from the go switch statements")
	}
}
