package main

import (
	"fmt"
	"time"
)

func main() {
	sendInterval := time.Second
	receptionInterval := 2 * time.Second

	intChan := make(chan int, 0)

	go func() {
		var ts0, ts1 int64
		for i := 1; i <= 5; i++ {
			intChan <- i
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Sent: ", i)
			} else {
				fmt.Printf("Send: %d [interval: %d s]\n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sendInterval)
		}
		close(intChan)
	}()

	var ts0, ts1 int64
Loop:
	for {
		select {
		case e, ok := <-intChan:
			if !ok {
				break Loop
			}
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Received: ", e)
			} else {
				fmt.Printf("Received: %d [interval: %d s]\n", e, ts1-ts0)
			}
		}
		ts0 = time.Now().Unix()
		time.Sleep(receptionInterval)
	}

	fmt.Println("End. ")
}
