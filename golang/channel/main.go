package main

import (
	"fmt"
	"time"
)

func goroutineA(a <-chan int)  {
	val := <- a
	fmt.Println("Received from A goroutine", val)
	return
}

func goroutineB(b <-chan int)  {
	val := <- b
	fmt.Println("Received from A goroutine", val)
	return
}

func main() {
	ch := make(chan int, 2)
	goroutineA(ch)
	goroutineB(ch)
	ch <- 3

	time.Sleep(time.Second)

	//ch1 := make(chan struct{})
}
