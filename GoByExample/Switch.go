package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2

	fmt.Println("write", i, "as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend")
	default:
		fmt.Println("Its a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its before noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am a int")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Printf("Dont know data type %T", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
