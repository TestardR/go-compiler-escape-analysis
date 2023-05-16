package main

import (
	"testing"
)

// go test -bench . -benchmem

// BenchmarkEscapeWithAppend-16            1000000000               0.1479 ns/op          0 B/op         0 allocs/op
func BenchmarkEscapeWithAppend(b *testing.B) {
	data := make([]*bigStruct, 0, numberOfItems)

	element := bigStruct{
		A: 123, B: 456, C: 789,
		D: "ABC", E: "DEF", F: "HIJ",
		G: true, H: true, I: true,
	}

	for i := 0; i < numberOfItems; i++ {
		data = append(data, &element)
	}
}

// BenchmarkEscapeWithIndex-16             1000000000               0.07195 ns/op         0 B/op         0 allocs/op
func BenchmarkEscapeWithIndex(b *testing.B) {
	data := make([]*bigStruct, numberOfItems)

	element := bigStruct{
		A: 123, B: 456, C: 789,
		D: "ABC", E: "DEF", F: "HIJ",
		G: true, H: true, I: true,
	}

	for i := 0; i < numberOfItems; i++ {
		data[i] = &element
	}
}

// BenchmarkEscapeWithAppendAndReturn-16           1000000000               0.1521 ns/op          0 B/op          0 allocs/op
func BenchmarkEscapeWithAppendAndReturn(b *testing.B) {
	_ = fillSliceWithAppend()
}

// BenchmarkEscapeWithIndexAndReturn-16            1000000000               0.05926 ns/op         0 B/op          0 allocs/op
func BenchmarkEscapeWithIndexAndReturn(b *testing.B) {
	_ = fillSliceWithIndex()
}
