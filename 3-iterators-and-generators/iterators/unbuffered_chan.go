package iterators

var sumUnbufferedChan int

func UnbufferedChanLoop(n int) int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()

	for j := range ch {
		sumUnbufferedChan += j
	}
	return sumUnbufferedChan
}
