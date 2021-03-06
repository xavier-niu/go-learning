package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	// slice don't have a length compared with array
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	// t.Log(s2[3) Error!
	s2 = append(s2, 1)
	t.Log(s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr"}
	Q2 := year[2:3]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[2:4]
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	t.Log(a, b)
	// Error! Slice can't be compared.
	// if a == b {
	// 	t.Log("equal")
	// }
}

func sliceAddress(s []int) {
	fmt.Printf("%p %T\n", s, s)
}

func TestSliceAddress(t *testing.T) {
	s := []int{1}
	fmt.Printf("%p %T\n", s, s)
	sliceAddress(s)
}
