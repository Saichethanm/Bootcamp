package main

func main() {
	queue := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		queue <- i
	}
	close(queue)

	for val := range queue {
		println(val)
	}
}
