package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"Eric", "John","Tom"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello %s\n", who)
		}(name)
		time.Sleep(1*time.Millisecond)
	}

}
