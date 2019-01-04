// wait for goroutine to finish.
// Another wait to do this is by using "select" inside of for loop

package main

import (
	"fmt"
	"sync"
)

var completion sync.WaitGroup

func main() {
	fmt.Println("start")

	completion.Add(1) // wait for one thing
	go doSomething()

	completion.Wait() // wait for all things to be done
	fmt.Println("end")
}

func doSomething() {
	fmt.Println("do something")
	completion.Done() // this is done
}
