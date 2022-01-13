package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("Hello ", i)
		time.Sleep(50 * time.Millisecond)
	}
}
