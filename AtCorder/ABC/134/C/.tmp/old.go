package main

import (
	"fmt"
	"sync"
)

func main() {
	var N uint32
	fmt.Scan(&N)

	An := make([]uint32, N)
	result := make([]uint32, N)

	var i uint32
	for i = 0; i < N; i++ {
		fmt.Scan(&An[i])
	}

	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup
	for i := range An {
		wg.Add(1)
		go func(array []uint32, except int) {
			mutex.Lock()

			var i int
			for i = 0; i < len(array); i++ {
				if result[except] < array[i] && except != i {
					result[except] = array[i]
				}
			}

			mutex.Unlock()
			wg.Done()
		}(An, i)
	}

	wg.Wait()

	for i = 0; i < N; i++ {
		fmt.Println(result[i])
	}
}
