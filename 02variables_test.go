package gobyexample

import (
	"fmt"
	"testing"
)

func TestVariables(t *testing.T) {
	var a = "inital"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(c, b)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
	
	t.Log("done")
}