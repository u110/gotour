package sort_test

import (
	"github.com/u110/gotour/sort/bubble"
	"testing"
)

func BenchmarkBubble(b *testing.B) {
	b.ResetTimer()
	arr := bubble.Init(b.N, b.N*100)
	bubble.Bubblesort(&arr)
}
