package work

import (
	"log"
	"time"
)

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
func FibFast(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
func TimeIt(name string) func() {
	start := time.Now()
	return func() { log.Printf("%s took %s", name, time.Since(start)) }
}
