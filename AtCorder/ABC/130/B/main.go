package main

import (
	"fmt"
)

func main() {
	var N int
	var X int32
	fmt.Scan(&N, &X)

	l := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&l[i])
	}
	// fmt.Printf("%v\n", l)

	var high, result int32 = 0, 0

	for i := 0; i < N; i++ {
		if high <= X {
			// fmt.Printf("%d\n", high)
			result++
		} else {
			break
		}
		high = high + int32(l[i])
	}

	fmt.Printf("%d\n", result)
}
