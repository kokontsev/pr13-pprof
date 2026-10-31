package work

import "testing"

var result int

func TestFib(t *testing.T) {
	for n := 0; n <= 30; n++ {
		if Fib(n) != FibFast(n) {
			t.Fatalf("n=%d", n)
		}
	}
}
func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = Fib(30)
	}
}
func BenchmarkFibFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = FibFast(30)
	}
}
