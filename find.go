package parlo

import (
	"sync/atomic"

	"github.com/mahdi-shojaee/parlo/internal/constraints"
	"github.com/mahdi-shojaee/parlo/internal/utils"
)

// Min returns the smallest element in the slice.
// If the slice is empty, it returns the zero value of type E.
func Min[S ~[]E, E constraints.Ordered](slice S) E {
	var min E

	if len(slice) == 0 {
		return min
	}

	min = slice[0]

	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

// MinBy returns the smallest element in the slice based on the provided comparison function.
// If the slice is empty, it returns the zero value of type E.
// The lt function should return true if a is considered less than b.
// If several values of the slice are equal to the smallest value, it returns the first such value.
func MinBy[S ~[]E, E any](slice S, lt func(a, b E) bool) E {
	var min E

	if len(slice) == 0 {
		return min
	}

	min = slice[0]

	for _, v := range slice[1:] {
		if lt(v, min) {
			min = v
		}
	}

	return min
}

// ParMin returns the smallest element in the slice using parallel processing.
// It uses the specified number of threads for processing.
// For slices with length less than 200,000, it falls back to the non-parallel Min function.
func ParMin[S ~[]E, E constraints.Ordered](slice S, numThreads int) E {
	const minLen = 200_000

	if len(slice) <= minLen {
		return Min(slice)
	}

	cb := func(s S, _, _ int) E {
		return Min(s)
	}

	result := do(slice, cb, utils.NumThreads(numThreads))

	return Min(result)
}

// ParMinBy returns the smallest element in the slice using parallel processing and a custom comparison function.
// It uses the specified number of threads for processing.
// For slices with length less than 200,000, it falls back to the non-parallel MinBy function.
// The lt function should return true if a is considered less than b.
// If several values of the slice are equal to the smallest value, it returns the first such value.
func ParMinBy[S ~[]E, E any](slice S, numThreads int, lt func(a, b E) bool) E {
	const minLen = 200_000

	if len(slice) <= minLen {
		return MinBy(slice, lt)
	}

	cb := func(s S, _, _ int) E {
		return MinBy(s, lt)
	}

	result := do(slice, cb, utils.NumThreads(numThreads))

	return MinBy(result, lt)
}

// Max returns the largest element in the slice.
// If the slice is empty, it returns the zero value of type E.
func Max[S ~[]E, E constraints.Ordered](slice S) E {
	var max E

	if len(slice) == 0 {
		return max
	}

	max = slice[0]

	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

// MaxBy returns the largest element in the slice based on the provided comparison function.
// If the slice is empty, it returns the zero value of type E.
// The gt function should return true if a is considered greater than b.
// If several values of the slice are equal to the largest value, it returns the first such value.
func MaxBy[S ~[]E, E any](slice S, gt func(a, b E) bool) E {
	var max E

	if len(slice) == 0 {
		return max
	}

	max = slice[0]

	for _, v := range slice[1:] {
		if gt(v, max) {
			max = v
		}
	}

	return max
}

// ParMax returns the largest element in the slice using parallel processing.
// It uses the specified number of threads for processing.
// For slices with length less than 200,000, it falls back to the non-parallel Max function.
func ParMax[S ~[]E, E constraints.Ordered](slice S, numThreads int) E {
	const minLen = 200_000

	if len(slice) <= minLen {
		return Max(slice)
	}

	cb := func(s S, _, _ int) E {
		return Max(s)
	}

	result := do(slice, cb, utils.NumThreads(numThreads))

	return Max(result)
}

// ParMaxBy returns the largest element in the slice using parallel processing and a custom comparison function.
// It uses the specified number of threads for processing.
// For slices with length less than 200,000, it falls back to the non-parallel MaxBy function.
// The gt function should return true if a is considered greater than b.
// If several values of the slice are equal to the largest value, it returns the first such value.
func ParMaxBy[S ~[]E, E any](slice S, numThreads int, gt func(a, b E) bool) E {
	const minLen = 200_000

	if len(slice) <= minLen {
		return MaxBy(slice, gt)
	}

	cb := func(s S, _, _ int) E {
		return MaxBy(s, gt)
	}

	result := do(slice, cb, utils.NumThreads(numThreads))

	return MaxBy(result, gt)
}

// Find returns the first element in the slice that satisfies the predicate function.
// It returns the found element and true if an element is found, otherwise it returns the zero value of E and false.
func Find[E any](slice []E, predicate func(item E) bool) (E, bool) {
	for i := range slice {
		if predicate(slice[i]) {
			return slice[i], true
		}
	}

	var result E
	return result, false
}

// ParFind returns the first element in the slice that satisfies the predicate function using parallel processing.
// It uses the specified number of threads for processing.
// For slices with length less than 200,000, it falls back to the non-parallel Find function.
// It returns the found element and true if an element is found, otherwise it returns the zero value of E and false.
func ParFind[E any](slice []E, numThreads int, predicate func(item E) bool) (E, bool) {
	const minLen = 200_000

	if len(slice) <= minLen {
		return Find(slice, predicate)
	}

	type ChunkResult struct {
		value E
		ok    bool
	}

	var end atomic.Uint64
	end.Store(0)

	cb := func(chunk []E, index int, chunkStartIndex int) ChunkResult {
		value, ok := func() (E, bool) {
			for _, v := range chunk {
				if end.Load()>>(64-index) != 0 {
					// Found by prev chunks
					var result E
					return result, false
				}
				if predicate(v) {
					// Found, So set the related bit in flag
					end.Add(1 << (63 - index))
					return v, true
				}
			}

			var result E
			return result, false
		}()

		return ChunkResult{value, ok}
	}

	results := do(slice, cb, utils.NumThreads(numThreads))

	r, ok := Find(results, func(r ChunkResult) bool {
		return r.ok
	})

	return r.value, ok
}
