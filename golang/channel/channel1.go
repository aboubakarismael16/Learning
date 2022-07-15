package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second ... [receiver]")
		time.Sleep(1 * time.Second)
		for elm := range strChan {
			fmt.Println("Received", elm, "[receiver]")
		}
		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for _, elm := range []string{"a", "b", "c", "d"} {
			strChan <- elm
			fmt.Println("Sent: ", elm, "[sender]")
			if elm == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Sent a sync signal. [sender]")
			}
		}
		fmt.Println("Wait 2 second... [sender]")
		time.Sleep(2 * time.Second)
		close(strChan)
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}
