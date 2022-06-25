package main

import (
	"os"
	"runtime/trace"
)

func allocate()  {
	//_ = make([]byte, 1<<200)
}

func main() {


	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	//for n := 1; n < 100000; n++ {
	//	allocate()
	//}

	keepalloc()
	//keepalloc2()
	keepalloc3()

}


func keepalloc2() {
	for i := 0; i < 100000; i++ {
		go func() {
			select {}
		}()
	}
}

var cache = map[interface{}]interface{}{}

func keepalloc() {
	for i := 0; i < 10000; i++ {
		m := make([]byte, 1<<10)
		cache[i] = m
	}
}

var ch = make(chan struct{})

func keepalloc3() {
	for i := 0; i < 100000; i++ {
		// 没有接收方，goroutine 会一直阻塞
		go func() { ch <- struct{}{} }()
	}
}