package gobyexample

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s), cap(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("add: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sli1:", l)

	l = s[:5]
	fmt.Println("sli2:", l)

	l = s[2:]
	fmt.Println("sli3:", l)

	x := []string{"g", "h", "i"}
	fmt.Println("dcl: ", x)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func TestSliceChange(t *testing.T) {
	s := make([]int, 5)
	for i := 0; i < 5; i++ {
		s[i] = i
	}
	fmt.Println("s len:", len(s), "cap:", cap(s))

	ns := s[:3]
	fmt.Println("ns len:", len(ns), "cap:", cap(ns))

	ns[0] = 11
	// same
	fmt.Println("s[1]", s[0], "ns[1]", ns[0])

	ns2 := s[2:]
	fmt.Println("ns2 len:", len(ns2), "cap:", cap(ns2))

	// not affect s
	ns2 = append(ns2, 22)
	fmt.Println("s len:", len(s), "ns len:", len(ns), "ns2 len:", len(ns2))
	fmt.Println("ns2:", ns2[3], "s:", s[4])

}
