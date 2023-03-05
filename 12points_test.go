package gobyexample

import (
	"fmt"
	"testing"
)

func zeroVal(i int) {
	i = 1
}

func zeroPtr(i *int) {
	*i = 1
}

func TestPoints(t *testing.T) {
	i := 0
	fmt.Println("init:", i)

	zeroVal(i)
	fmt.Println("val:", i)

	zeroPtr(&i)
	fmt.Println("ptr:", i)
}
