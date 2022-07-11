package main

import "fmt"

func main() {

	defer fmt.Println("defer in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic oncer")
}
