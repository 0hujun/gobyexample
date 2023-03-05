package gobyexample

import (
	"fmt"
	"testing"
)

// intSeq The returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0 // as global var
	return func() int {
		i++
		return i
	}
}

func TestClosure(t *testing.T) {

	nextInt := intSeq()

	// add by itself
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInt := intSeq()
	fmt.Println(newInt())  // 1
	fmt.Println(newInt())  // 2
	fmt.Println(nextInt()) // 4
}
