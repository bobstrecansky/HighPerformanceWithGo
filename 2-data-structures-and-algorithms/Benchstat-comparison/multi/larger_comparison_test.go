package comparison_larger

import "testing"

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CarJSON()
	}
}
