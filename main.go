package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	x, y := 1, 2
	fmt.Println(Max(x, y))
}

func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}
