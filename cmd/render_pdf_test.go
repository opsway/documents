package cmd

import (
	"testing"
)

func BenchmarkRender(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = RenderPdf("./public/index.html")
	}
}
