package HighPerformanceGo

import "testing"

func TestPassArray(t *testing.T) {
	m := [6]int{1, 2, 3, 4, 5, 6}
	fn := func(s [6]int) {
		s[2] = 4
		t.Logf("passing array. %v\n", s)
	}
	fn(m)
	t.Logf("passing array. %v\n", m)
}

func TestPassSlice(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	fn := func(s []int) {
		s[2] = 4
		t.Logf("passing array. %v\n", s)
	}
	fn(m)
	t.Logf("passing array. %v\n", m)
}
