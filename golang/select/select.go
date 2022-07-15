package main

import "fmt"

var intChan1 chan int
var intChan2 chan int
var channels = []chan int{intChan1, intChan2}

var nums = []int{1,2,3,4,5}

func main() {
	select {
	case getChan(0) <- getNum(0):
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNum(1):
		fmt.Println("2nd case is selected.")
	default:
		fmt.Println("Default")
	}

	chanCap := 5
	intChan3 := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		select {
		case intChan3 <- 1:
		case intChan3 <- 2:
		case intChan3 <- 3:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%d\n", <-intChan3)
	}
}

func getNum(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return nums[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}