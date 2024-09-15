package parlo

import (
	"runtime"
	"sync"
)

func do[S ~[]E, E, R any](
	collection S,
	cb func(chunk S, index int, chunkStartIndex int) R,
	threads int,
) []R {
	numCPU := runtime.NumCPU()

	if threads == 0 || threads > numCPU {
		threads = numCPU
	}

	if len(collection) < threads {
		return []R{cb(collection, 0, 0)}
	}

	result := make([]R, threads)

	chunkSize := len(collection) / threads

	s := collection

	var wg sync.WaitGroup
	wg.Add(threads)

	for i := 0; i < threads; i++ {
		endIndex := chunkSize
		if i == threads-1 {
			endIndex = len(s)
		}

		chunk := s[:endIndex]
		go func(chunk S, index int, chunkStartIndex int) {
			result[index] = cb(chunk, index, chunkStartIndex)
			wg.Done()
		}(chunk, i, len(collection)-len(s))

		s = s[endIndex:]
	}

	wg.Wait()

	return result
}
