package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i % 20 == 0 {
			continue
		}

		if i == 95 {
			break
		}
		fmt.Print(i," ")
	}
	fmt.Println()
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()

	j := 0
	anExp := true
	for ok := true; ok; ok =anExp {
		if j > 10 {
			break
		}
		fmt.Print(j," ")
		j++
	}
	fmt.Println()

}
