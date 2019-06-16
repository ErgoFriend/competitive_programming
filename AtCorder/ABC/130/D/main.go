package main

import (
	"fmt"
)

func listSum(len int32, list []int32) int64 {
	var sum int64 = 0

	for i := int32(0); i < len; i++ {
		sum += int64(list[i])
	}
	return sum
}

func main() {
	var N, tmpContSubseqLen int32
	var K, result int64
	fmt.Scan(&N, &K)

	A := make([]int32, N)
	for i := int32(0); i < N; i++ {
		fmt.Scan(&A[i])
	}

	tmpContSubseqLen = N

	for i := int32(0); i < N; i++ {
		// iは部分配列を右に移動させる回数
		tmpContSubseqLen = N - i
		for move := int32(0); move < i; move++ {
			start := move
			end := move + tmpContSubseqLen + 1
			if K <= listSum(tmpContSubseqLen+1, A[start:end]) {
				result++
			}
		}
	}

	fmt.Printf("%d\n", result)

}
