package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	var ms runtime.MemStats
	printMemStats(ms)

	s := make([]int, 900000)
	for i := 0; i <len(s); i++ {
		s[i] += 2
	}

	time.Sleep(5*time.Second)

	printMemStats(ms)

}

func printMemStats(ms runtime.MemStats)  {
	runtime.ReadMemStats(&ms)

	fmt.Println("--------------------------------------")
	fmt.Println("Memory Statistics Reporting time: ", time.Now())
	fmt.Println("--------------------------------------")
	fmt.Println("Bytes of allocated heap objects: ", ms.Alloc)
	fmt.Println("Total bytes of Heap object: ", ms.TotalAlloc)
	fmt.Println("Bytes of memory obtained from OS: ", ms.Sys)
	fmt.Println("Count of heap objects: ", ms.Mallocs)
	fmt.Println("Count of heap objects freed: ", ms.Frees)
	fmt.Println("Count of live heap objects", ms.Mallocs-ms.Frees)
	fmt.Println("Number of completed GC cycles: ", ms.NumGC)
	fmt.Println("--------------------------------------")
}
