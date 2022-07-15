package main

import (
	"fmt"
	"time"
)

var strChan4 = make(chan string, 3)

func main() {
	synChan1 := make(chan struct{}, 1)
	synChan2 := make(chan struct{}, 2)
	go receive(strChan4, synChan1, synChan2)
	go send(strChan4, synChan1, synChan2)
	<- synChan2
	<- synChan2

	var ok bool
	ch := make(chan int, 1)
	_, ok = interface{}(ch).(<- chan int)
	fmt.Println("chan int => <- chan int :", ok) // false
	_, ok = interface{}(ch).(chan <- int)
	fmt.Println("chan int => <- chan int :", ok) //false

	sch := make(<- chan int, 1)
	_, ok = interface{}(sch).(chan int)
	fmt.Println("<- chan int => chan int :", ok)  // false

	rch := make(chan <- int, 1)
	_, ok = interface{}(rch).(chan int)
	fmt.Println("chan <- int => chan int :", ok)  // false


}

func receive(strChan4 <- chan string, synChan1 <- chan struct{}, synChan2 chan<- struct{})  {
	<-synChan1
	fmt.Println("Received a sync signal and wait a second ... [received]")
	time.Sleep(time.Second)
	for {
		if elm,ok := <- strChan4; ok {
			fmt.Println("Received: ",elm,"[receiver]")
		} else {
			break
		}
	}
	fmt.Println("Stopped. [receiver]")
	synChan2 <- struct{}{}
}

func send(strChan4 chan <- string, synChan1 chan <- struct{}, synChan2 chan <- struct{})  {
	for _, elm := range []string{"a", "b", "c", "d"} {
		strChan4 <- elm
		fmt.Println("Sent: ", elm,"[sender]")
		if elm == "c" {
			synChan1 <- struct{}{}
			fmt.Println("Sent sync signal. [sender]")
		}
	}
	fmt.Println("Wait 2 seconds ... [sender]")
	time.Sleep(2*time.Second)
	close(strChan4)
	synChan2 <- struct{}{}
}
