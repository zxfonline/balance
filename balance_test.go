package balance

import (
	"testing"
)

func Benchmark_ToBase(b *testing.B) {
	for d := 0; d < b.N; d++ {
		GetUCServerById(int64(d))
	}
}

func TestGetUCServerById(t *testing.T) {
	for u := 0; u < 10000; u++ {
		GetUCServerById(int64(u))
	}
}
