package main

import (
	"fmt"

	"github.com/ergofriend/ml2array"
)

func main() {
	input := `4 5 2
0 1 2 1
0 2 1 2
1 2 1 1
1 3 1 3
2 3 2 1`

	result := ml2array.Make(input)

	for i := 0; i < len(result); i++ {
		fmt.Printf("%d行目: %s+++\n", i, result[i])
	}

}
