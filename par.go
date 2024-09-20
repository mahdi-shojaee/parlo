package parlo

import (
	"sync"

	"github.com/mahdi-shojaee/parlo/internal/utils"
)

func Do[S ~[]E, E, R any](
	collection S,
	numThreads int,
	cb func(chunk S, index, chunkStartIndex int) R,
) []R {
	threads := utils.NumThreads(numThreads)

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
		go func(chunk S, index, chunkStartIndex int) {
			result[index] = cb(chunk, index, chunkStartIndex)
			wg.Done()
		}(chunk, i, len(collection)-len(s))

		s = s[endIndex:]
	}

	wg.Wait()

	return result
}
