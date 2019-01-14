package HighPerformanceGo

import (
	"strconv"
	"testing"
)

var stringElems []string

const (
	startNum = 1000000000000000000
	times    = 50000
)

func init() {
	for i := startNum; i < startNum+times; i++ {
		stringElems = append(stringElems, strconv.Itoa(i))
	}
}

func BenchmarkStringKey(b *testing.B) {
	m := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, el := range stringElems {
			m[el] = 1
		}
		for _, el := range stringElems {
			_ = m[el]
		}
	}
}

func BenchmarkIntKey(b *testing.B) {
	m := make(map[int]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := startNum; j < startNum+times; j++ {
			m[j] = 1
		}
		for j := startNum; j < startNum+times; j++ {
			_ = m[j]
		}
	}
}

//Ouput:
//BenchmarkStringKey-4       	     500	   3574504 ns/op
//BenchmarkIntKey-4          	     500	   2908102 ns/op
