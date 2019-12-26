package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("\ntick")
		case <-boom:
			fmt.Println("\nBoom")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
