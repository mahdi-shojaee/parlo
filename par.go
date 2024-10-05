package parlo

import (
	"runtime"
	"sync"
)

// Do is a generic function that applies a callback function to each chunk of the input slice in parallel.
// It splits the input slice into multiple chunks and processes each chunk in a separate goroutine.
// The callback function is executed for each chunk, and its result is collected in a new slice.
// The function returns the final result slice after all goroutines have completed.
// If numThreads is 0 or a negative number, it automatically uses all available CPU cores.
// If numThreads is 1, the function runs in a separate goroutine, allowing asynchronous execution without parallelism.
// If numThreads is greater than 1, it manually specifies the exact number of threads.
func Do[S ~[]E, E, R any](
	slice S,
	numThreads int,
	cb func(chunk S, index, chunkStartIndex int) R,
) []R {
	if len(slice) == 0 {
		return []R{cb(S{}, 0, 0)}
	}

	numCPU := GOMAXPROCS()

	if numThreads <= 0 || numThreads > numCPU {
		numThreads = numCPU
	}

	if len(slice) <= numThreads {
		numThreads = len(slice)
	}

	result := make([]R, numThreads)

	chunkSize := len(slice) / numThreads

	s := slice

	var wg sync.WaitGroup
	wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		endIndex := chunkSize
		if i == numThreads-1 {
			endIndex = len(s)
		}

		chunk := s[:endIndex]
		go func(chunk S, index, chunkStartIndex int) {
			defer wg.Done()
			result[index] = cb(chunk, index, chunkStartIndex)
		}(chunk, i, len(slice)-len(s))

		s = s[endIndex:]
	}

	wg.Wait()

	return result
}

func GOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}
