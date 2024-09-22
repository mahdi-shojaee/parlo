package parlo

// Filter applies a predicate function to each element of the input slice
// and returns a new slice containing only the elements for which the predicate returns true.
func Filter[S ~[]E, E any](slice S, predicate func(item E, index int) bool) S {
	result := make([]E, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		if predicate(slice[i], i) {
			result = append(result, slice[i])
		}
	}

	return result
}

// ParFilter applies a predicate function to each element of the input slice in parallel
// and returns a new slice containing only the elements for which the predicate returns true.
//
// If the input slice length is less than or equal to (200,000), it falls back to the sequential Filter function.
func ParFilter[S ~[]E, E any](slice S, predicate func(item E, index int) bool) S {
	const minLen = 200_000

	if len(slice) <= minLen {
		return Filter(slice, predicate)
	}

	chunkResults := Do(slice, 0, func(chunk S, _, _ int) []E {
		return Filter(chunk, predicate)
	})

	size := 0

	for _, chunkResult := range chunkResults {
		size += len(chunkResult)
	}

	result := make([]E, 0, size)

	for _, chunkResult := range chunkResults {
		result = append(result, chunkResult...)
	}

	return result
}
