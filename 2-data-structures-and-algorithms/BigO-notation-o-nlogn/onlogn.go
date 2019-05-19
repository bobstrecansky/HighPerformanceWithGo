package onlogn

func OnlognLoop(n int) {
	for i := 0; i < n; i++ { // this loop O(n) because it is executed n times
		for j := n; j > 0; j /= 2 { // this loop is O(log n)
			// The result here is O(n) * O(log n) = O(n log n)
		}
	}
}
