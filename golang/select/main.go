package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for range time.Tick(1* time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <- ch:
			fmt.Println("case 1")
		case <- ch:
			fmt.Println("case 2")
		}
	}
}

