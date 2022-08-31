package main

import "testing"

func init() {
	print = false
}

func BenchmarkPascalTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pascalTriangle(50)
	}
}

func BenchmarkPascalTriangle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pascalTriangle2(50)
	}
}
