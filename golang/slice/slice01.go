package main

import "fmt"

func main() {
	s := []int{2,3,4,5,6,7}


	y := s[1:3]

	x := make([]int,2,3)
	x = s[:]

	s1 := []string{99:"a"}

	var s2 []int

	s3 := make([]int, 0)

	s4 := []int{10,20,30,40,50}

	newSlice := s4[1:3]

	newSlice[1] = 35


	fmt.Printf("len = %d, cap = %d ,add = %p\n", len(s),cap(s), s)
	fmt.Printf("len = %d, cap = %d ,add = %p\n", len(y),cap(y), y)

	fmt.Printf("x = %T, len = %d, cap = %d\n", x,len(x), cap(x))
	fmt.Printf("x value of x = %v\n", x)

	fmt.Printf("s1 = %v, len = %d, cap = %d ,add = %p\n",s1, len(s1),cap(s1), s1)

	fmt.Printf("s1 = %v, len = %d, cap = %d ,add = %p\n",s2, len(s2),cap(s2), s2)
	fmt.Printf("s1 = %v, len = %d, cap = %d ,add = %p\n",s3, len(s3),cap(s3), s3)

	fmt.Printf("s4 add = %p and newSlice add = %p\n", s4, newSlice)
	fmt.Println(s4,"len(s4)=", len(s4),newSlice, "len(newSlice)=", len(newSlice))
}
