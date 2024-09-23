package parlo

import (
	"sync/atomic"

	"github.com/mahdi-shojaee/parlo/internal/constraints"
)

type IsSortedChunkResult[E any] struct {
	isSorted        bool
	chunkStartIndex int
	chunk           []E
}

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

// IsSorted checks if the input slice is sorted in ascending order.
// It returns true if the slice is sorted, false otherwise.
func IsSorted[S ~[]E, E constraints.Ordered](slice S) bool {
	if len(slice) <= 1 {
		return true
	}

	item := slice[0]

	for _, v := range slice[1:] {
		if item > v {
			return false
		}
		item = v
	}

	return true
}

// IsSortedBy checks if the input slice is sorted according to the provided comparison function.
// The gt function should return true if a is considered greater than b.
// It returns true if the slice is sorted, false otherwise.
func IsSortedBy[S ~[]E, E any](slice S, gt func(a, b E) bool) bool {
	if len(slice) <= 1 {
		return true
	}

	item := slice[0]

	for _, v := range slice[1:] {
		if gt(item, v) {
			return false
		}
		item = v
	}

	return true
}

// ParIsSorted checks if the input slice is sorted in ascending order in parallel.
// It returns true if the slice is sorted, false otherwise.
// Note: ParIsSorted is generally faster than IsSorted for slices with length greater than approximately 55,000 elements.
func ParIsSorted[S ~[]E, E constraints.Ordered](slice S) bool {
	if len(slice) <= 1 {
		return true
	}

	var end uint32 = 0

	results := Do(slice, 0, func(chunk S, _, chunkStartIndex int) IsSortedChunkResult[E] {
		if len(chunk) <= 1 {
			return IsSortedChunkResult[E]{
				isSorted:        true,
				chunkStartIndex: chunkStartIndex,
				chunk:           chunk,
			}
		}

		prev := chunk[0]
		isSorted := true

		for _, v := range chunk[1:] {
			if atomic.LoadUint32(&end) != 0 {
				isSorted = false
				break
			}

			if prev > v {
				atomic.StoreUint32(&end, 1)
				isSorted = false
				break
			}

			prev = v
		}

		return IsSortedChunkResult[E]{
			isSorted:        isSorted,
			chunkStartIndex: chunkStartIndex,
			chunk:           chunk,
		}
	})

	var prevChunkLastItem E
	prevChunkLastItemIsSet := false

	for _, r := range results {
		if len(r.chunk) == 0 {
			continue
		}

		if !r.isSorted {
			return false
		}

		if prevChunkLastItemIsSet && prevChunkLastItem > r.chunk[0] {
			return false
		}

		prevChunkLastItem = r.chunk[len(r.chunk)-1]
		prevChunkLastItemIsSet = true
	}

	return true
}

// ParIsSortedBy checks if the input slice is sorted according to the provided comparison function in parallel.
// The cmp function should return a negative integer if a is considered less than b,
// a positive integer if a is considered greater than b, and zero if a is considered equal to b.
// It returns true if the slice is sorted, false otherwise.
// Note: ParIsSortedBy is generally faster than IsSortedBy for slices with length greater than approximately 20,000 elements.
func ParIsSortedBy[S ~[]E, E any](slice S, gt func(a, b E) bool) bool {
	if len(slice) <= 1 {
		return true
	}

	var end uint32 = 0

	results := Do(slice, 0, func(chunk S, _, chunkStartIndex int) IsSortedChunkResult[E] {
		if len(chunk) <= 1 {
			return IsSortedChunkResult[E]{
				isSorted:        true,
				chunkStartIndex: chunkStartIndex,
				chunk:           chunk,
			}
		}

		prev := chunk[0]
		isSorted := true

		for _, v := range chunk[1:] {
			if atomic.LoadUint32(&end) != 0 {
				isSorted = false
				break
			}

			if gt(prev, v) {
				atomic.StoreUint32(&end, 1)
				isSorted = false
				break
			}

			prev = v
		}

		return IsSortedChunkResult[E]{
			isSorted:        isSorted,
			chunkStartIndex: chunkStartIndex,
			chunk:           chunk,
		}
	})

	var prevChunkLastItem E
	prevChunkLastItemIsSet := false

	for _, r := range results {
		if len(r.chunk) == 0 {
			continue
		}

		if !r.isSorted {
			return false
		}

		if prevChunkLastItemIsSet && gt(prevChunkLastItem, r.chunk[0]) {
			return false
		}

		prevChunkLastItem = r.chunk[len(r.chunk)-1]
		prevChunkLastItemIsSet = true
	}

	return true
}
