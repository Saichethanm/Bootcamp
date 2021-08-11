package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	fmt.Println(<-timer1.C)
	fmt.Println("Timer 1 is fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 is fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 is stopped")
	}

	time.Sleep(time.Second)
}
