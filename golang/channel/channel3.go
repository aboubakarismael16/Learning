package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

var mapChan3 = make(chan map[string]*Counter, 1)

func main() {
	synChan3 := make(chan struct{}, 2)

	go func() {
		for {
			if elm, ok := <-mapChan3; ok {
				counter := elm["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		synChan3 <- struct{}{}
	}()

	go func() {
		countMap3 := map[string]*Counter{
			"count": &Counter{},
		}

		for i := 0; i < 5; i++ {
			mapChan3 <- countMap3
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map : %v. [sender]\n", countMap3)
		}
		close(mapChan3)
		synChan3 <- struct{}{}
	}()
	<-synChan3
	<-synChan3
}
