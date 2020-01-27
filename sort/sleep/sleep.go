package sleep

import (
	"sync"
	"time"
)

func Sort(dat *[]int) {
	result := make([]int, len(*dat))
	ch := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for _, v := range *dat {
			wg.Add(1)
			go func(v int) {
				time.Sleep(time.Duration(v) * time.Millisecond)
				ch <- v
				wg.Done()
			}(v)
		}
		wg.Wait()
		close(ch)
	}()
	i := 0
	for {
		receive, ok := <-ch
		if !ok {
			*dat = result
			return
		}
		result[i] = receive
		i++
	}
}
