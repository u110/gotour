package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Swap(arr *[]int, i int, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

func Init() []int {
	rand.Seed(time.Now().UnixNano())
	maxLen := 30
	maxNum := 200
	count := 0

	arr := make([]int, maxLen)
	// initialize data
	for count < maxLen {
		arr[count] = rand.Intn(maxNum)
		count++
	}

	return arr
}

func Show(arr []int) {
	fmt.Print("\033[H\033[2J") // clear
	for i, val := range arr {
		fmt.Printf("%2d(%4d): ", i, val)
		idx := 0
		for idx < val {
			fmt.Print("*")
			idx++
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("start")
	arr := Init()
	Show(arr)

	doneIdx := len(arr)
	for doneIdx > 0 {
		// fmt.Println("doneIdx", doneIdx)
		i := 0
		for i < doneIdx-1 {
			// fmt.Println(i, arr[i], i+1, arr[i+1])
			if arr[i] > arr[i+1] {
				Swap(&arr, i, i+1)
				Show(arr)
				fmt.Println("donIdx", doneIdx)
				time.Sleep(time.Duration(100) * time.Millisecond)
			}
			i++
		}
		doneIdx--
	}

	Show(arr)
	fmt.Println("end")
}
