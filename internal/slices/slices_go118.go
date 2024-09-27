//go:build go1.18 && !go1.21

package slices

import (
	"cmp"
	"sort"

	"github.com/mahdi-shojaee/parlo/internal/constraints"
)

// Sort sorts the given slice in ascending order.
// The slice must contain elements that satisfy the constraints.Ordered interface.
func Sort[S ~[]E, E constraints.Ordered](slice S) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

// SortFunc sorts the given slice using the provided comparison function.
// The comparison function should return a negative integer, zero, or a positive integer
// if a is considered to be respectively less than, equal to, or greater than b.
func SortFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) {
	sort.Slice(slice, func(i, j int) bool {
		return cmp(slice[i], slice[j]) < 0
	})
}

// SortStableFunc sorts the given slice using the provided comparison function
// while preserving the original order of equal elements.
// The comparison function should return a negative integer, zero, or a positive integer
// if a is considered to be respectively less than, equal to, or greater than b.
func SortStableFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) {
	sort.SliceStable(slice, func(i, j int) bool {
		return cmp(slice[i], slice[j]) < 0
	})
}

// Reverse reverses the order of elements in the given slice in place.
func Reverse[S ~[]E, E any](slice S) {
	l := len(slice)

	if l <= 1 {
		return
	}

	end := l / 2

	for i, j := 0, l-1; i < end; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
// The result may have additional unused capacity.
func Clone[S ~[]E, E any](s S) S {
	// The s[:0:0] preserves nil in case it matters.
	return append(s[:0:0], s...)
}

// Concat returns a new slice concatenating the passed in slices.
func Concat[S ~[]E, E any](slices ...S) S {
	size := 0
	for _, s := range slices {
		size += len(s)
		if size < 0 {
			panic("len out of range")
		}
	}
	newslice := Grow[S](nil, size)
	for _, s := range slices {
		newslice = append(newslice, s...)
	}
	return newslice
}

// BinarySearch searches for target in a sorted slice and returns the earliest
// position where target is found, or the position where target would appear
// in the sort order; it also returns a bool saying whether the target is
// really found in the slice. The slice must be sorted in increasing order.
func BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool) {
	// Inlining is faster than calling BinarySearchFunc with a lambda.
	n := len(x)
	// Define x[-1] < target and x[n] >= target.
	// Invariant: x[i-1] < target, x[j] >= target.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if cmp.Less(x[h], target) {
			i = h + 1 // preserves x[i-1] < target
		} else {
			j = h // preserves x[j] >= target
		}
	}
	// i == j, x[i-1] < target, and x[j] (= x[i]) >= target  =>  answer is i.
	return i, i < n && (x[i] == target || (isNaN(x[i]) && isNaN(target)))
}

// BinarySearchFunc works like [BinarySearch], but uses a custom comparison
// function. The slice must be sorted in increasing order, where "increasing"
// is defined by cmp. cmp should return 0 if the slice element matches
// the target, a negative number if the slice element precedes the target,
// or a positive number if the slice element follows the target.
// cmp must implement the same ordering as the slice, such that if
// cmp(a, t) < 0 and cmp(b, t) >= 0, then a must precede b in the slice.
func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool) {
	n := len(x)
	// Define cmp(x[-1], target) < 0 and cmp(x[n], target) >= 0 .
	// Invariant: cmp(x[i - 1], target) < 0, cmp(x[j], target) >= 0.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if cmp(x[h], target) < 0 {
			i = h + 1 // preserves cmp(x[i - 1], target) < 0
		} else {
			j = h // preserves cmp(x[j], target) >= 0
		}
	}
	// i == j, cmp(x[i-1], target) < 0, and cmp(x[j], target) (= cmp(x[i], target)) >= 0  =>  answer is i.
	return i, i < n && cmp(x[i], target) == 0
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}
