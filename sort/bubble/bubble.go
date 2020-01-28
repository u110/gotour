package bubble

import (
	"fmt"
)

func Swap(arr *[]int, i int, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
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

func Bubblesort(dat *[]int) {
	arr := *dat
	doneIdx := len(arr)
	for doneIdx > 0 {
		// fmt.Println("doneIdx", doneIdx)
		i := 0
		for i < doneIdx-1 {
			// fmt.Println(i, arr[i], i+1, arr[i+1])
			if arr[i] > arr[i+1] {
				Swap(&arr, i, i+1)
				// Show(arr)
				// fmt.Println("donIdx", doneIdx)
				// time.Sleep(time.Duration(100) * time.Millisecond)
			}
			i++
		}
		doneIdx--
	}
}
