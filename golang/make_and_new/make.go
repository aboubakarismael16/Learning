package main

import (
	"fmt"
	"time"
)

func main() {
	startAT := time.Now()
	defer func() { fmt.Println(time.Since(startAT)) }()

	time.Sleep(time.Second)
}
