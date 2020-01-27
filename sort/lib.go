package sort

import (
	"math/rand"
	"time"
)

func Init(maxLen int, maxNum int) []int {
	rand.Seed(time.Now().UnixNano())
	count := 0

	arr := make([]int, maxLen)
	// initialize data
	for count < maxLen {
		arr[count] = rand.Intn(maxNum)
		count++
	}

	return arr
}
