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
