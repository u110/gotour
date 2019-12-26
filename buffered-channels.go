package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	s := []int{1, 2}
	for _, v := range s {
		ch <- v
	}

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
