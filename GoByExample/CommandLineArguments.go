package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[5]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println("Hello: ", arg)
}
