package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	envs := os.Environ()
	for _, val := range envs {
		pair := strings.SplitN(val, "=", -1)
		fmt.Println(pair)
	}
	fmt.Println()
}
