package main

import (
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci(10)
	}
}
