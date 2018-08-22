package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go func(i int) {
			for {
				fmt.Println(i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
