package main

import "fmt"

var done = make(chan bool)
var msg string

func aGoroutine()  {
	msg = "Hello world"
	<- done
}

func main() {
	go aGoroutine()
	done <- true
	fmt.Println(msg)
}
