package main

import (
	"crypto/rand"
	"testing"
)

func Benchmark_Read_4(b *testing.B) {
	const size = 4
	b.SetBytes(int64(size))
	buf := make([]byte, size)
	for b.Loop() {
		if _, err := rand.Read(buf); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Read_32(b *testing.B) {
	const size = 32
	b.SetBytes(int64(size))
	buf := make([]byte, size)
	for b.Loop() {
		if _, err := rand.Read(buf); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Read_4K(b *testing.B) {
	const size = 4 << 10
	b.SetBytes(int64(size))
	buf := make([]byte, size)
	for b.Loop() {
		if _, err := rand.Read(buf); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Text(b *testing.B) {
	b.SetBytes(26)
	for b.Loop() {
		sinkStr = rand.Text()
	}
}
