package main

import "fmt"

func main() {
	dataChan := make(chan int, 5)
	synChan1 := make(chan struct{}, 1)
	synChan2 := make(chan struct{}, 2)

	go func() { // receive goroutine
		<- synChan1
		for {
			if elm, ok := <- dataChan; ok {
				fmt.Printf("Received: %d [receiver]\n", elm)
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		synChan2 <- struct{}{}
	}()

	go func() { // send goroutine
		for i := 0; i < 5 ; i++ {
			dataChan <- i
			fmt.Printf("Sent: %d [sender]\n", i)
		}
		close(dataChan)
		synChan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		synChan2 <- struct{}{}
	}()
	<- synChan2
	<- synChan2
}
