package iterators

var sumBufferedChan int

func BufferedChanLoop(n int) int {

	ch := make(chan int, n)

	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()

	for j := range ch {
		sumBufferedChan += j
	}
	return sumBufferedChan
}
