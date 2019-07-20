package main

import (
	"fmt"
)

func main() {
	var N uint8
	var D uint8
	fmt.Scan(&N, &D)

	ranges := 1 + D*2
	d := N / ranges
	m := N % ranges

	if d < 1 {
		fmt.Println(1)
	} else {
		if m == 0 {
			fmt.Println(d)
		} else {
			fmt.Println(d + 1)
		}
	}

}
