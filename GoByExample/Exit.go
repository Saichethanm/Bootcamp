package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hello world")
	// fmt.Println("Hello world")

	os.Exit(1)
}
