package main

import (
	"testing"
)

func BenchmarkEcho2(b *testing.B) {
	args := ReadArgs()
	b.ReportAllocs()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		result := Echo2(args)
		b.StopTimer()
		_ = result
	}
}
