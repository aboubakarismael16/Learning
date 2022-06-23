package main

import "fmt"

func main() {
	s := []int{10,20,30,40,50}

	newSlice := s[1:3]

	//newSlice = append(newSlice, 45,66,67)
	fmt.Printf("len = %d, cap = %d, vallue = %v,  add = %p\n", len(s), cap(s),s, &s)
	fmt.Printf("before append len = %d, cap = %d,vallue = %v, add = %p\n", len(newSlice), cap(newSlice),newSlice, &newSlice)

	newSlice = append(newSlice, 60,70,80)
	fmt.Printf("after append len = %d, cap = %d,vallue = %v, add = %p\n", len(newSlice), cap(newSlice), newSlice, &newSlice)

	//passing slice in function
	funcSlice := make([]int, 1e6)

	fmt.Println(foo(funcSlice))
}



func foo(slice []int) []int {
	return slice
}
