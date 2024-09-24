package parlo

import (
	"sync/atomic"

	"github.com/mahdi-shojaee/parlo/internal/constraints"
)

type ParFindChunkResult[E any] struct {
	value E
	ok    bool
}

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

// MinFunc returns the smallest element in the slice based on the provided comparison function.
// If the slice is empty, it returns the zero value of type E.
// The cmp function should return a negative integer if a is considered less than b,
// a positive integer if a is considered greater than b, and zero if a is considered equal to b.
// If several values of the slice are equal to the smallest value, it returns the first such value.
func MinFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) E {
	var min E

	if len(slice) == 0 {
		return min
	}

	min = slice[0]

	for _, v := range slice[1:] {
		if cmp(v, min) < 0 {
			min = v
		}
	}

	return min
}

// ParMin returns the smallest element in the slice using parallel processing.
// Note: ParMin is generally faster than Min for slices with length greater than approximately 200,000 elements.
func ParMin[S ~[]E, E constraints.Ordered](slice S) E {
	result := Do(slice, 0, func(s S, _, _ int) E {
		return Min(s)
	})

	return Min(result)
}

// ParMinFunc returns the smallest element in the slice using parallel processing and a custom comparison function.
// The cmp function should return a negative integer if a is considered less than b,
// a positive integer if a is considered greater than b, and zero if a is considered equal to b.
// If several values of the slice are equal to the smallest value, it returns the first such value.
// Note: ParMinFunc is generally faster than MinFunc for slices with length greater than approximately 10,000 elements.
func ParMinFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) E {
	result := Do(slice, 0, func(s S, _, _ int) E {
		return MinFunc(s, cmp)
	})

	return MinFunc(result, cmp)
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

// MaxFunc returns the largest element in the slice based on the provided comparison function.
// If the slice is empty, it returns the zero value of type E.
// The cmp function should return a positive integer if a is considered greater than b,
// a negative integer if a is considered less than b, and zero if a is considered equal to b.
// If several values of the slice are equal to the largest value, it returns the first such value.
func MaxFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) E {
	var max E

	if len(slice) == 0 {
		return max
	}

	max = slice[0]

	for _, v := range slice[1:] {
		if cmp(v, max) > 0 {
			max = v
		}
	}

	return max
}

// ParMax returns the largest element in the slice using parallel processing.
// Note: ParMax is generally faster than Max for slices with length greater than approximately 130,000 elements.
func ParMax[S ~[]E, E constraints.Ordered](slice S) E {
	result := Do(slice, 0, func(s S, _, _ int) E {
		return Max(s)
	})

	return Max(result)
}

// ParMaxFunc returns the largest element in the slice using parallel processing and a custom comparison function.
// The cmp function should return a positive integer if a is considered greater than b,
// a negative integer if a is considered less than b, and zero if a is considered equal to b.
// If several values of the slice are equal to the largest value, it returns the first such value.
// Note: ParMaxFunc is generally faster than MaxFunc for slices with length greater than approximately 10,000 elements.
func ParMaxFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) E {
	result := Do(slice, 0, func(s S, _, _ int) E {
		return MaxFunc(s, cmp)
	})

	return MaxFunc(result, cmp)
}

// Find returns the first element in the slice that satisfies the predicate function.
// It returns the found element and true if an element is found, otherwise it returns the zero value of E and false.
func Find[E any](slice []E, predicate func(item E) bool) (E, bool) {
	for _, x := range slice {
		if predicate(x) {
			return x, true
		}
	}

	var result E
	return result, false
}

// ParFind returns the first element in the slice that satisfies the predicate function using parallel processing.
// It returns the found element and true if an element is found, otherwise it returns the zero value of E and false.
// Note: ParFind is generally faster than Find for slices with length greater than approximately 1,000,000,000 elements.
func ParFind[E any](slice []E, predicate func(item E) bool) (E, bool) {
	var mask uint64 = 0

	results := Do(slice, 0, func(chunk []E, index int, chunkStartIndex int) ParFindChunkResult[E] {
		for _, v := range chunk {
			if atomic.LoadUint64(&mask)>>(64-index) != 0 {
				// Found by prev chunks
				var zero E
				return ParFindChunkResult[E]{zero, false}
			}
			if predicate(v) {
				// Found, So set the related bit in flag
				atomic.AddUint64(&mask, 1<<(63-index))
				return ParFindChunkResult[E]{v, true}
			}
		}

		var result E
		return ParFindChunkResult[E]{result, false}
	})

	for _, v := range results {
		if v.ok {
			return v.value, true
		}
	}

	var zero E
	return zero, false
}
