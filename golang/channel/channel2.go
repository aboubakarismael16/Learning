package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main() {
	synChan := make(chan struct{}, 2)

	go func() {
		for {
			if elm, ok := <- mapChan; ok {
				elm["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		synChan <- struct{}{}
	}()

	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		synChan <- struct{}{}
	}()
	<- synChan
	<- synChan
}
