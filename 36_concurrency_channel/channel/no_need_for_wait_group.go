package main

import "fmt"

var ch = make(chan int) // create channel

// Pings prints few lines, then sends an int over channel back to main
func Pings() {
	for i := 0; i < 3; i++ {
		fmt.Println("Step", i)
	}
	ch <- 1 // send channel value
}

func main() {
	go Pings() // concurrent execution
	fmt.Println("Start")
	<-ch // wait for channel value
	fmt.Println("The End")
}
