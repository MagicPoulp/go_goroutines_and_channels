package sum

import (
	"sync"
)

// worker constitutes a goroutine that sums a sub-slice and sends the result to a channel.
func worker(data []int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, num := range data {
		sum += num
	}
	resultChan <- sum
}

// ConcurrentSum divides the work of summing a slice across multiple goroutines.
func ConcurrentSum(data []int, numWorkers int) int {
	// Handle the empty slice case
	if len(data) == 0 {
		return 0
	}

	// Cap the number of workers if it's more than the number of elements
	if numWorkers > len(data) {
		numWorkers = len(data)
	}

	var wg sync.WaitGroup
	resultChan := make(chan int, numWorkers)

	chunkSize := len(data) / numWorkers
	if chunkSize == 0 {
		chunkSize = 1
	}

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(data)
		}
		wg.Add(1)
		go worker(data[start:end], resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalSum := 0
	for result := range resultChan {
		totalSum += result
	}

	return totalSum
}
