package on

// Create a slice of n sorted ints
// generateIntSlice(3) == [0 1 2]
func generateIntSlice(n int) []int {
	x := make([]int, 0, n)
	for i := 0; i < n; i++ {
		x = append(x, n)
	}
	return x
}

// Sum the elements in a slice
func simpleLoopSum(count int) int {
	slice := generateIntSlice(count)
	sum := 0
	for i := 0; i < count; i++ {
		sum += slice[i]
	}
	return sum
}
