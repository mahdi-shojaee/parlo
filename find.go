package parlo

import (
	"github.com/mahdi-shojaee/parlo/internal/constraints"
)

// Min searches for the minimum value in a slice.
//
// If multiple values share the minimum value, the first encountered one is returned.
// Returns the zero value if the slice is empty.
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

// MinBy finds the minimum value of a slice using the given comparison function.
//
// If multiple values in the slice are equal to the minimum, the first one is returned.
// Returns the zero value of the element type if the slice is empty.
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

// ParMin finds the minimum value in a slice using the specified number
// of threads. The minimum value found across all chunks is returned.
//
// If multiple values in the slice are equal to the minimum, the first one is returned.
// Returns the zero value of the element type if the slice is empty.
func ParMin[S ~[]E, E constraints.Ordered](slice S, numThreads int) E {
	// Less than MIN_LEN, single thread is faster.
	const minLen = 200_000

	if len(slice) <= minLen {
		return Min(slice)
	}

	cb := func(s S, _, _ int) E {
		return Min(s)
	}

	result := do(slice, cb, numThreads)

	return Min(result)
}

// ParMinBy finds the minimum value in a slice using the specified number
// of threads, using a custom less-than function. The minimum value found across
// all chunks is returned.
//
// If multiple values in the slice are equal to the minimum, the first one is returned.
// Returns the zero value of the element type if the slice is empty.
func ParMinBy[S ~[]E, E any](slice S, numThreads int, lt func(a, b E) bool) E {
	// Less than MIN_LEN, single thread is faster.
	const minLen = 200_000

	if len(slice) <= minLen {
		return MinBy(slice, lt)
	}

	cb := func(s S, _, _ int) E {
		return MinBy(s, lt)
	}

	result := do(slice, cb, numThreads)

	return MinBy(result, lt)
}

// Max searches for the maximum value in a slice.
//
// If multiple values share the maximum value, the first one is returned.
// Returns the zero value if the slice is empty.
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

// MaxBy finds the maximum value in a slice using the provided comparison function.
//
// If multiple values share the maximum value, the first one is returned.
// Returns the zero value if the slice is empty.
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

// ParMax finds the maximum value in a slice using the specified number of threads.
// The maximum value found across all chunks is returned.
//
// If multiple values in the slice are equal to the maximum, the first one is returned.
// Returns the zero value of the element type if the slice is empty.
func ParMax[S ~[]E, E constraints.Ordered](slice S, numThreads int) E {
	// Less than MIN_LEN, single thread is faster.
	const minLen = 200_000

	if len(slice) <= minLen {
		return Max(slice)
	}

	cb := func(s S, _, _ int) E {
		return Max(s)
	}

	result := do(slice, cb, numThreads)

	return Max(result)
}

// ParMaxBy finds the maximum value in a slice using the specified number
// of threads, based on a custom comparison function.
//
// If multiple values in the slice are equal to the maximum, the first one is returned.
// Returns the zero value of the element type if the slice is empty.
//
// The `gt` function should return true if `a` is greater than `b`.
func ParMaxBy[S ~[]E, E any](slice S, numThreads int, gt func(a, b E) bool) E {
	// Less than MIN_LEN, single thread is faster.
	const minLen = 200_000

	if len(slice) <= minLen {
		return MaxBy(slice, gt)
	}

	cb := func(s S, _, _ int) E {
		return MaxBy(s, gt)
	}

	result := do(slice, cb, numThreads)

	return MaxBy(result, gt)
}
