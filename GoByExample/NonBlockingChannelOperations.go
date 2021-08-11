package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println(msg)
	default:
		fmt.Println("No message received")
	}

	select {
	case messages <- "Hello":
		fmt.Println("Message sent")
	default:
		fmt.Println("No message sent")
	}

	select {
	case messages <- "Hello":
		fmt.Println("do something")
	case signals <- true:
		fmt.Println("Do something")
	default:
		fmt.Println("No activity")
	}
}
