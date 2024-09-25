package parlo

import (
	"sync"
	"sync/atomic"

	"github.com/mahdi-shojaee/parlo/internal/constraints"
)

type isSortedChunkResult[E any] struct {
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
// Note: ParFilter is generally faster than Filter for slices with length greater than approximately 12,000 elements.
func ParFilter[S ~[]E, E any](slice S, predicate func(item E, index int) bool) S {
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

// IsSortedFunc checks if the input slice is sorted according to the provided comparison function.
// The cmp function should return a negative integer if a is considered less than b,
// a positive integer if a is considered greater than b, and zero if a is considered equal to b.
// It returns true if the slice is sorted, false otherwise.
func IsSortedFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) bool {
	if len(slice) <= 1 {
		return true
	}

	item := slice[0]

	for _, v := range slice[1:] {
		if cmp(item, v) > 0 {
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

	results := Do(slice, 0, func(chunk S, _, chunkStartIndex int) isSortedChunkResult[E] {
		if len(chunk) <= 1 {
			return isSortedChunkResult[E]{
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

		return isSortedChunkResult[E]{
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

// ParIsSortedFunc checks if the input slice is sorted according to the provided comparison function in parallel.
// The cmp function should return a negative integer if a is considered less than b,
// a positive integer if a is considered greater than b, and zero if a is considered equal to b.
// It returns true if the slice is sorted, false otherwise.
// Note: ParIsSortedFunc is generally faster than IsSortedFunc for slices with length greater than approximately 20,000 elements.
func ParIsSortedFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) bool {
	if len(slice) <= 1 {
		return true
	}

	var end uint32 = 0

	results := Do(slice, 0, func(chunk S, _, chunkStartIndex int) isSortedChunkResult[E] {
		if len(chunk) <= 1 {
			return isSortedChunkResult[E]{
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

			if cmp(prev, v) > 0 {
				atomic.StoreUint32(&end, 1)
				isSorted = false
				break
			}

			prev = v
		}

		return isSortedChunkResult[E]{
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

		if prevChunkLastItemIsSet && cmp(prevChunkLastItem, r.chunk[0]) > 0 {
			return false
		}

		prevChunkLastItem = r.chunk[len(r.chunk)-1]
		prevChunkLastItemIsSet = true
	}

	return true
}

// IsSortedDesc checks if the input slice is sorted in descending order.
// It returns true if the slice is sorted, false otherwise.
func IsSortedDesc[S ~[]E, E constraints.Ordered](slice S) bool {
	if len(slice) <= 1 {
		return true
	}

	item := slice[0]

	for _, v := range slice[1:] {
		if item < v {
			return false
		}
		item = v
	}

	return true
}

// ParIsSortedDesc checks if the input slice is sorted in descending order in parallel.
// It returns true if the slice is sorted, false otherwise.
// Note: ParIsSorted is generally faster than IsSorted for slices with length greater than approximately 55,000 elements.
func ParIsSortedDesc[S ~[]E, E constraints.Ordered](slice S) bool {
	if len(slice) <= 1 {
		return true
	}

	var end uint32 = 0

	results := Do(slice, 0, func(chunk S, _, chunkStartIndex int) isSortedChunkResult[E] {
		if len(chunk) <= 1 {
			return isSortedChunkResult[E]{
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

			if prev < v {
				atomic.StoreUint32(&end, 1)
				isSorted = false
				break
			}

			prev = v
		}

		return isSortedChunkResult[E]{
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

		if prevChunkLastItemIsSet && prevChunkLastItem < r.chunk[0] {
			return false
		}

		prevChunkLastItem = r.chunk[len(r.chunk)-1]
		prevChunkLastItemIsSet = true
	}

	return true
}

// Reverse reverses the elements of the input slice in place.
func Reverse[S ~[]E, E any](slice S) {
	if len(slice) <= 1 {
		return
	}

	end := len(slice) / 2
	l := len(slice)

	for i, j := 0, l-1; i < end; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// ParReverse reverses the elements of the input slice in parallel.
// Note: ParReverse is generally faster than Reverse for slices with length greater than approximately 600,000 elements.
func ParReverse[S ~[]E, E any](slice S) {
	l := len(slice)

	if l <= 1 {
		return
	}

	numThreads := 2

	chunkSize := l / numThreads / 2

	var wg sync.WaitGroup

	for i, startIndex := 0, 0; i < numThreads-1; i, startIndex = i+1, startIndex+chunkSize {
		leftSlice := slice[startIndex : startIndex+chunkSize]
		rightSlice := slice[l-startIndex-chunkSize : l-startIndex]

		wg.Add(1)
		go func() {
			defer wg.Done()
			for j, k := 0, len(rightSlice)-1; j < chunkSize; j, k = j+1, k-1 {
				leftSlice[j], rightSlice[k] = rightSlice[k], leftSlice[j]
			}
		}()
	}

	Reverse(slice[(numThreads-1)*chunkSize : l-(numThreads-1)*chunkSize])

	wg.Wait()
}
