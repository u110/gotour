package sort_test

import (
	"github.com/u110/gotour/sort"
	"github.com/u110/gotour/sort/bubble"
	"github.com/u110/gotour/sort/sleep"
	"testing"
)

func BenchmarkBubble(b *testing.B) {
	b.ResetTimer()
	arr := sort.Init(b.N, b.N*100)
	bubble.Bubblesort(&arr)
}

func BenchmarkSleep(b *testing.B) {
	b.ResetTimer()
	arr := sort.Init(b.N, b.N*100)
	sleep.Sort(&arr)
}
