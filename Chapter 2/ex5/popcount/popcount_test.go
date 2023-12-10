package popcount

import (
	"math/rand"
	"testing"
)

func BenchmarkPopCount1(b *testing.B) {
	b.ReportAllocs()
    b.ResetTimer()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
		_ = PopCount1(rand.Uint64())
	}
    b.StopTimer()
}

func BenchmarkPopCount2(b *testing.B) {
    b.ReportAllocs()
    b.ResetTimer()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
		_ = PopCount2(rand.Uint64())
	}
    b.StopTimer()
}

