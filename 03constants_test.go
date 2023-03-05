package gobyexample

import (
	"fmt"
	"math"
	"testing"
)

const s string = "constant"

func TestConstant(t *testing.T) {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(n))

	fmt.Println(math.Sin(n))
	t.Log("done")
}