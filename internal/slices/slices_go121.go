//go:build go1.21

package slices

import (
	sls "slices"

	"github.com/mahdi-shojaee/parlo/internal/constraints"
)

// Sort sorts the given slice in ascending order.
// The slice must contain elements that satisfy the constraints.Ordered interface.
func Sort[S ~[]E, E constraints.Ordered](slice S) {
	sls.Sort(slice)
}

// SortFunc sorts the given slice using the provided comparison function.
// The comparison function should return a negative integer, zero, or a positive integer
// if a is considered to be respectively less than, equal to, or greater than b.
func SortFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) {
	sls.SortFunc(slice, cmp)
}

// SortStableFunc sorts the given slice using the provided comparison function
// while preserving the original order of equal elements.
// The comparison function should return a negative integer, zero, or a positive integer
// if a is considered to be respectively less than, equal to, or greater than b.
func SortStableFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) {
	sls.SortStableFunc(slice, cmp)
}

// Reverse reverses the order of elements in the given slice in place.
func Reverse[S ~[]E, E any](slice S) {
	sls.Reverse(slice)
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
// The result may have additional unused capacity.
func Clone[S ~[]E, E any](s S) S {
	// The s[:0:0] preserves nil in case it matters.
	return append(s[:0:0], s...)
}

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func Grow[S ~[]E, E any](s S, n int) S {
	if n < 0 {
		panic("cannot be negative")
	}
	if n -= cap(s) - len(s); n > 0 {
		s = append(s[:cap(s)], make([]E, n)...)[:len(s)]
	}
	return s
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
func BinarySearch[S ~[]E, E constraints.Ordered](x S, target E) (int, bool) {
	return sls.BinarySearch(x, target)
}

func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool) {
	return sls.BinarySearchFunc(x, target, cmp)
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[S ~[]E, E comparable](s S, v E) int {
	return sls.Index(s, v)
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	return sls.IndexFunc(s, f)
}

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Empty and nil slices are considered equal.
// Floating point NaNs are not considered equal.
func Equal[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
