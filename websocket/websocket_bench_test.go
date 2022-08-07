package websocket

import (
	"fmt"
	"strings"
	"testing"
)

// BenchmarkSend benchmarks the Send function
func BenchmarkSend(b *testing.B) {
	tableTest := []struct {
		inputSize int
	}{
		{1},
		{1000},
		{1000000},
		{1000000000},
		{10000000000},
	}

	for _, test := range tableTest {
		b.Run(fmt.Sprintf("input_size_%d", test.inputSize), func(b *testing.B) {
			b.ReportAllocs()
			payload := strings.Repeat("#", test.inputSize)
			Send("3000", payload)
		})
	}
}
