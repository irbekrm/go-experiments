package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"testing"
)

var result *rsa.PrivateKey

func main() {
	res1024 := testing.Benchmark(BenchmarkCreateKey1024)
	fmt.Printf("Seconds to create a key with 1024 bits: %v\n", float64(res1024.NsPerOp())/1000000000)
	res4096 := testing.Benchmark(BenchmarkCreateKey4096)
	fmt.Printf("Seconds to create a key with 4096 bits: %v\n", float64(res4096.NsPerOp())/1000000000)
}

func CreateKey(bits int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

func BenchmarkCreateKey1024(b *testing.B) {
	benchmarkCreateKey(b, 1024)
}

func BenchmarkCreateKey4096(b *testing.B) {
	benchmarkCreateKey(b, 4096)
}

func benchmarkCreateKey(b *testing.B, bits int) {
	for n := 0; n < b.N; n++ {
		result, _ = CreateKey(bits)
	}
}
