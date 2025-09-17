package main

import (
	"fmt"
	"go_goroutines_and_channels/sum"
)

func main() {
	largeData := make([]int, 1000000)
	for i := range largeData {
		largeData[i] = 1
	}

	numWorkers := 4
	sum := sum.ConcurrentSum(largeData, numWorkers)
	fmt.Printf("The total sum of the large data set is: %d\n", sum)
}
