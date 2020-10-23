package main

import (
	"bytes"
	"testing"
)

func BenchmarkPrintWithPool(b *testing.B) {
	bb := &bytes.Buffer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrintWithPool([]byte(`{"name":"Aina"}`), bb)
	}
}

func BenchmarkPrintWithoutPool(b *testing.B) {
	bb := &bytes.Buffer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrintWithoutPool([]byte(`{"name":"Elga"}`), bb)
	}
}
