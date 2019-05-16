package main

import "fmt"

func Count(start int, end int) <-chan int {
	ch := make(chan int)
	go func(ch chan int) {
		for i := start; i <= end; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	return ch
}

func CountLargeBuffer(start int, end int) <-chan int {
	ch := make(chan int, 5)
	go func(ch chan int) {
		for i := start; i <= end; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	return ch
}

func main() {
	// basic for
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("----------------------------------")
	// Count returns ch, this make continious number
	for i := range Count(1, 9) {
		fmt.Println(i)
	}
	fmt.Println("----------------------------------")
	// If you want to change func of Count, you fix only Count
	for i := range CountLargeBuffer(1, 9) {
		fmt.Println(i)
	}
	fmt.Println("----------------------------------")
}
