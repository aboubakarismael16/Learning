package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int

	wg sync.WaitGroup
)


func main() {

	wg.Add(2)

	go inCounter(1)
	go inCounter(2)

	wg.Wait()
	fmt.Println("Final counter :", counter)

}


func inCounter(id int)  {
	defer  wg.Done()

	for count := 0; count < 2; count++ {
		//get the value of counter
		value := counter
		//Yield the thread and be placed back in queue
		runtime.Gosched()

		//increment the local value of counter
		value++

		//store the value back into counter
		counter = value
	}
}
