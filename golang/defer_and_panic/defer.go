package main

import "fmt"

func main() {
	c()
	b()
}

func b()  {
	for i := 0 ; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func c() (i int)  {
	defer func() {i++}()
	return 1
}
