package sum

import (
	"testing"
)

func TestConcurrentSum(t *testing.T) {
	// Test case 1: A simple, predictable sum.
	t.Run("simple sum", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expectedSum := 55
		numWorkers := 3
		sum := ConcurrentSum(data, numWorkers)
		if sum != expectedSum {
			t.Errorf("ConcurrentSum(%v) = %d; want %d", data, sum, expectedSum)
		}
	})

	// Test case 2: Empty slice.
	t.Run("empty slice", func(t *testing.T) {
		data := []int{}
		expectedSum := 0
		numWorkers := 4
		sum := ConcurrentSum(data, numWorkers)
		if sum != expectedSum {
			t.Errorf("ConcurrentSum(%v) = %d; want %d", data, sum, expectedSum)
		}
	})

	// Test case 3: Single worker, which should behave like a serial sum.
	t.Run("single worker", func(t *testing.T) {
		data := []int{10, 20, 30, 40}
		expectedSum := 100
		numWorkers := 1
		sum := ConcurrentSum(data, numWorkers)
		if sum != expectedSum {
			t.Errorf("ConcurrentSum(%v) = %d; want %d", data, sum, expectedSum)
		}
	})

	// Test case 4: More workers than data elements.
	t.Run("more workers than elements", func(t *testing.T) {
		data := []int{5, 5, 5}
		expectedSum := 15
		numWorkers := 10
		sum := ConcurrentSum(data, numWorkers)
		if sum != expectedSum {
			t.Errorf("ConcurrentSum(%v) = %d; want %d", data, sum, expectedSum)
		}
	})

	// Test case 5: A large data set.
	t.Run("large data set", func(t *testing.T) {
		data := make([]int, 1000000)
		expectedSum := 0
		for i := range data {
			data[i] = i
			expectedSum += i
		}
		numWorkers := 8
		sum := ConcurrentSum(data, numWorkers)
		if sum != expectedSum {
			t.Errorf("ConcurrentSum(large data) = %d; want %d", sum, expectedSum)
		}
	})
}
