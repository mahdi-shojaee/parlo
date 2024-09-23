package parlo

import (
	"github.com/mahdi-shojaee/parlo/internal/constraints"
)

type ChunkResult[E any] struct {
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
// Note: ParMin is generally faster than Min for slices with length greater than approximately 200,000 elements.
func ParMin[S ~[]E, E constraints.Ordered](slice S) E {
	result := Do(slice, 0, func(s S, _, _ int) E {
		return Min(s)
	})

	return Min(result)
}

// ParMinBy returns the smallest element in the slice using parallel processing and a custom comparison function.
// The lt function should return true if a is considered less than b.
// If several values of the slice are equal to the smallest value, it returns the first such value.
// Note: ParMinBy is generally faster than MinBy for slices with length greater than approximately 10,000 elements.
func ParMinBy[S ~[]E, E any](slice S, lt func(a, b E) bool) E {
	result := Do(slice, 0, func(s S, _, _ int) E {
		return MinBy(s, lt)
	})

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
// Note: ParMax is generally faster than Max for slices with length greater than approximately 130,000 elements.
func ParMax[S ~[]E, E constraints.Ordered](slice S) E {
	result := Do(slice, 0, func(s S, _, _ int) E {
		return Max(s)
	})

	return Max(result)
}

// ParMaxBy returns the largest element in the slice using parallel processing and a custom comparison function.
// The gt function should return true if a is considered greater than b.
// If several values of the slice are equal to the largest value, it returns the first such value.
// Note: ParMaxBy is generally faster than MaxBy for slices with length greater than approximately 10,000 elements.
func ParMaxBy[S ~[]E, E any](slice S, gt func(a, b E) bool) E {
	result := Do(slice, 0, func(s S, _, _ int) E {
		return MaxBy(s, gt)
	})

	return MaxBy(result, gt)
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
	results := Do(slice, 0, func(chunk []E, index int, chunkStartIndex int) ChunkResult[E] {
		for _, v := range chunk {
			if predicate(v) {
				return ChunkResult[E]{v, true}
			}
		}

		var result E
		return ChunkResult[E]{result, false}
	})

	for _, v := range results {
		if v.ok {
			return v.value, true
		}
	}

	var zero E
	return zero, false
}
