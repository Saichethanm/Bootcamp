package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	fmt.Println(time.Second)
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

/*

 **** When you dont send anything into the channel ***
go run ChannelSynchronization.go
Working...
1s
done
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /home/saichm/Saichethan/bootcamp/GoByExample/ChannelSynchronization.go:19 +0x6f
exit status 2
*/
