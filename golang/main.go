package main

import "fmt"
var v = 10

func foo() *int {
	v = 10
	return &v
}

func main() {
	m := foo()
	fmt.Println(&m,*m)
}
