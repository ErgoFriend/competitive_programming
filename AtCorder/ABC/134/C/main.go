package main

import (
	"fmt"
)

func main() {
	var N uint32
	fmt.Scan(&N)

	An := make([]uint32, N, 200000)

	var i uint32
	for i = 0; i < N; i++ {
		fmt.Scan(&An[i])
	}

	for i := range An {
		var j uint32
		var result uint32
		result = 0
		for j = 0; j < N; j++ {
			if result < An[j] && j != uint32(i) {
				result = An[j]
			}
		}

		fmt.Println(result)
	}

}
