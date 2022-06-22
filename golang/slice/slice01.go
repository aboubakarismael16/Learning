package main

import "fmt"

func main() {
	s := []int{2,3,4,5,6,7}


	y := s[1:3]

	x := make([]int,2,3)
	x = s[:]



	fmt.Printf("len = %d, cap = %d ,add = %p\n", len(s),cap(s), s)
	fmt.Printf("len = %d, cap = %d ,add = %p\n", len(y),cap(y), y)

	fmt.Printf("x = %T, len = %d, cap = %d\n", x,len(x), cap(x))
	fmt.Printf("x value of x = %v\n", x)
}
