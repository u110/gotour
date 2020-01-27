package main

import (
	"fmt"
	"math/rand"
	"time"
)

func swap(arr *[]int, i int, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	maxLen := 10
	maxNum := 100
	count := 0
	arr := make([]int, maxLen)
	// initialize data
	for count < maxLen {
		arr[count] = rand.Intn(maxNum)
		count++
	}
	fmt.Println("start")
	fmt.Println(arr)

	doneIdx := maxLen
	for doneIdx > 0 {
		// fmt.Println("doneIdx", doneIdx)
		i := 0
		for i < doneIdx-1 {
			// fmt.Println(i, arr[i], i+1, arr[i+1])
			if arr[i] > arr[i+1] {
				swap(&arr, i, i+1)
			}
			i++
		}
		doneIdx--
	}

	fmt.Println("end")
	fmt.Println(arr)
}
