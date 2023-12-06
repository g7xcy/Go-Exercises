package main

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	args := ReadArgs()
	b.ReportAllocs()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		result := Echo1(args)
		b.StopTimer()
		_ = result
	}
}
