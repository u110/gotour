package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("start")
	defer fmt.Println("end")
	go say("1")
	go say("2")
	say("3")
}
