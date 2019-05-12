package string_concat

import "testing"

const TEST_STRING = "test"

func benchmarkConcat(size int, SelfConcat func(string, int) string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SelfConcat(TEST_STRING, size)
	}
}

func BenchmarkConcatOperator2(b *testing.B)      { benchmarkConcat(2, SelfConcatOperator, b) }
func BenchmarkConcatOperator10(b *testing.B)     { benchmarkConcat(10, SelfConcatOperator, b) }
func BenchmarkConcatOperator100(b *testing.B)    { benchmarkConcat(100, SelfConcatOperator, b) }
func BenchmarkConcatOperator1000(b *testing.B)   { benchmarkConcat(1000, SelfConcatOperator, b) }
func BenchmarkConcatOperator10000(b *testing.B)  { benchmarkConcat(10000, SelfConcatOperator, b) }
func BenchmarkConcatOperator100000(b *testing.B) { benchmarkConcat(100000, SelfConcatOperator, b) }

func BenchmarkConcatBuffer2(b *testing.B)      { benchmarkConcat(2, SelfConcatBuffer, b) }
func BenchmarkConcatBuffer10(b *testing.B)     { benchmarkConcat(10, SelfConcatBuffer, b) }
func BenchmarkConcatBuffer100(b *testing.B)    { benchmarkConcat(100, SelfConcatBuffer, b) }
func BenchmarkConcatBuffer1000(b *testing.B)   { benchmarkConcat(1000, SelfConcatBuffer, b) }
func BenchmarkConcatBuffer10000(b *testing.B)  { benchmarkConcat(10000, SelfConcatBuffer, b) }
func BenchmarkConcatBuffer100000(b *testing.B) { benchmarkConcat(100000, SelfConcatBuffer, b) }
