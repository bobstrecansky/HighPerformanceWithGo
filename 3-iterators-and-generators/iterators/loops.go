package iterators

var sumLoops int

func simpleLoop(n int) int {
	for i := 0; i < n; i++ {
		sumLoops += i
	}
	return sumLoops
}
