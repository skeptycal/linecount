package linecount

import (
	"testing"
)

func BenchmarkStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sampleRunStringsBuilder(defaultSampleSize)
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sampleRunBytesBuffer(defaultSampleSize)
	}
}

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sample()
	}
}
