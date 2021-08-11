package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	// secs := now.Unix()
	// nanos := now.UnixNano()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Println(now)

	millis := now.UnixNano() / 1000000
	fmt.Println(millis)
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))

}
