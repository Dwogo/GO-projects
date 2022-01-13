package main

import (
	"fmt"
	"time"
)

func main() {
	for ; ; {
		fmt.Println("🔥CPU🔥")
		time.Sleep(500 * time.Millisecond)
	}
}
