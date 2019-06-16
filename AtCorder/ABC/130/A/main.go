package main

import (
	"fmt"
)

func main() {
	var X, A int
	fmt.Scan(&X, &A)

	var ans int
	if X < A {
		ans = 0
	}
	if A <= X {
		ans = 10
	}

	fmt.Printf("%d\n", ans)
}
