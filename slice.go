package parlo

import (
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/mahdi-shojaee/parlo/internal/cmp"
	"github.com/mahdi-shojaee/parlo/internal/constraints"
	"github.com/mahdi-shojaee/parlo/internal/slices"
	"github.com/mahdi-shojaee/parlo/internal/utils"
)

type isSortedChunkResult[S ~[]E, E any] struct {
	isSorted        bool
	chunkStartIndex int
	chunk           S
}

type node[S ~[]E, E any] struct {
	sortedChunk S
	chunkIndex  int
}

// Filter applies a predicate function to each element of the input slice
// and returns a new slice containing only the elements for which the predicate returns true.
func Filter[S ~[]E, E any](slice S, predicate func(item E, index int) bool) S {
	result := make(S, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		if predicate(slice[i], i) {
			result = append(result, slice[i])
		}
	}

	return result
}

// ParFilter applies a predicate function to each element of the input slice in parallel
// and returns a new slice containing only the elements for which the predicate returns true.
func ParFilter[S ~[]E, E any](slice S, predicate func(item E, index int) bool) S {
	chunkResults := Do(slice, 0, func(chunk S, _, _ int) S {
		return Filter(chunk, predicate)
	})

	size := 0

	for _, chunkResult := range chunkResults {
		size += len(chunkResult)
	}

	result := make(S, 0, size)

	for _, chunkResult := range chunkResults {
		result = append(result, chunkResult...)
	}

	return result
}

func parCopyChunks[S ~[]E, E any](dest S, chunks []S) {
	destIndex := 0

	var wg sync.WaitGroup
	wg.Add(len(chunks))

	for _, chunk := range chunks {
		go func(destIndex int, chunk S) {
			defer wg.Done()
			copy(dest[destIndex:], chunk)
		}(destIndex, chunk)

		destIndex += len(chunk)
	}

	wg.Wait()
}

// Equal checks if two slices are equal.
func Equal[S ~[]E, E comparable](a S, b S) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// ParEqual checks if two slices are equal in parallel.
func ParEqual[S ~[]E, E comparable](a S, b S) bool {
	if len(a) != len(b) {
		return false
	}

	result := Do(a, 0, func(chunk S, _, chunkStartIndex int) bool {
		other := b[chunkStartIndex : chunkStartIndex+len(chunk)]
		return Equal(chunk, other)
	})

	for _, r := range result {
		if !r {
			return false
		}
	}

	return true
}

// EqualFunc checks if two slices are equal according to a comparison function.
func EqualFunc[S ~[]E, E any](a S, b S, eq func(a, b E) bool) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !eq(a[i], b[i]) {
			return false
		}
	}

	return true
}

// ParEqualFunc checks if two slices are equal in parallel according to a comparison function.
func ParEqualFunc[S ~[]E, E any](a S, b S, eq func(a, b E) bool) bool {
	if len(a) != len(b) {
		return false
	}

	result := Do(a, 0, func(chunk S, _, chunkStartIndex int) bool {
		other := b[chunkStartIndex : chunkStartIndex+len(chunk)]
		return EqualFunc(chunk, other, eq)
	})

	for _, r := range result {
		if !r {
			return false
		}
	}

	return true
}

func isSortedChunksSorted[S ~[]E, E constraints.Ordered](sortedChunks []S) bool {
	chunks := Filter(sortedChunks, func(chunk S, index int) bool {
		return len(chunk) > 0
	})

	if len(chunks) < 2 {
		return true
	}

	for i := 1; i < len(chunks); i++ {
		if chunks[i-1][len(chunks[i-1])-1] > chunks[i][0] {
			return false
		}
	}

	return true
}

func isSortedChunksSortedBy[S ~[]E, E any](sortedChunks []S, cmp func(a, b E) int) bool {
	chunks := Filter(sortedChunks, func(chunk S, index int) bool {
		return len(chunk) > 0
	})

	if len(chunks) < 2 {
		return true
	}

	for i := 1; i < len(chunks); i++ {
		if cmp(chunks[i-1][len(chunks[i-1])-1], chunks[i][0]) > 0 {
			return false
		}
	}

	return true
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
func ParIsSorted[S ~[]E, E constraints.Ordered](slice S) bool {
	if len(slice) <= 1 {
		return true
	}

	var end uint32 = 0

	results := Do(slice, 0, func(chunk S, _, chunkStartIndex int) isSortedChunkResult[S, E] {
		if len(chunk) <= 1 {
			return isSortedChunkResult[S, E]{
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

		return isSortedChunkResult[S, E]{
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
func ParIsSortedFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) bool {
	if len(slice) <= 1 {
		return true
	}

	var end uint32 = 0

	results := Do(slice, 0, func(chunk S, _, chunkStartIndex int) isSortedChunkResult[S, E] {
		if len(chunk) <= 1 {
			return isSortedChunkResult[S, E]{
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

		return isSortedChunkResult[S, E]{
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
func ParIsSortedDesc[S ~[]E, E constraints.Ordered](slice S) bool {
	if len(slice) <= 1 {
		return true
	}

	var end uint32 = 0

	results := Do(slice, 0, func(chunk S, _, chunkStartIndex int) isSortedChunkResult[S, E] {
		if len(chunk) <= 1 {
			return isSortedChunkResult[S, E]{
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

		return isSortedChunkResult[S, E]{
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
	l := len(slice)

	if l <= 1 {
		return
	}

	end := l / 2

	for i, j := 0, l-1; i < end; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// ParReverse reverses the elements of the input slice in parallel.
func ParReverse[S ~[]E, E any](slice S) {
	l := len(slice)

	if l <= 1 {
		return
	}

	// In practice, 2 is the best number of threads for this function.
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

// Sort sorts the given slice in ascending order. The slice must contain elements that satisfy the constraints.Ordered interface.
func Sort[S ~[]E, E constraints.Ordered](slice S) {
	slices.Sort(slice)
}

// SortFunc sorts the given slice using the provided comparison function.
// The comparison function should return a negative number when a < b,
// a positive number when a > b, and zero when a == b.
func SortFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) {
	slices.SortFunc(slice, cmp)
}

// SortStableFunc sorts the given slice using the provided comparison function.
// It maintains the relative order of equal elements, making it a stable sort.
// The comparison function should return a negative number when a < b,
// a positive number when a > b, and zero when a == b.
func SortStableFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) {
	slices.SortStableFunc(slice, cmp)
}

// ParSort performs a parallel sort on the given slice in ascending order.
// The slice must contain elements that satisfy the constraints.Ordered interface.
func ParSort[S ~[]E, E constraints.Ordered](slice S) {
	isSortedDesc := ParIsSortedDesc(slice)

	if isSortedDesc {
		ParReverse(slice)
		return
	}

	numThreads := runtime.GOMAXPROCS(0)

	chunks := Do(slice, numThreads, func(chunk S, _, _ int) S {
		if IsSorted(chunk) {
			return chunk
		}
		if IsSortedDesc(chunk) {
			Reverse(chunk)
			return chunk
		}
		Sort(chunk)
		return chunk
	})

	chunks = Filter(chunks, func(chunk S, index int) bool {
		return len(chunk) > 0
	})

	if isSortedChunksSorted(chunks) {
		return
	}

	chunksCopy := Do(chunks, len(chunks), func(chunk []S, _, _ int) S {
		return slices.Clone(chunk[0])
	})

	slices.SortFunc(chunksCopy, func(a, b S) int {
		return cmp.Compare(a[0], b[0])
	})

	if isSortedChunksSorted(chunksCopy) {
		parCopyChunks(slice, chunksCopy)
		return
	}

	parMergeByMerge(chunksCopy, slice, minHeapMerge[S, E])
}

// ParSortFunc performs a parallel sort on the given slice using a custom comparison function.
func ParSortFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) {
	isSortedDesc := ParIsSortedFunc(slice, func(a E, b E) int { return cmp(b, a) })

	if isSortedDesc {
		ParReverse(slice)
		return
	}

	numThreads := runtime.GOMAXPROCS(0)

	chunks := Do(slice, numThreads, func(chunk S, _, _ int) S {
		if IsSortedFunc(chunk, cmp) {
			return chunk
		}
		if IsSortedFunc(chunk, func(a E, b E) int { return cmp(b, a) }) {
			Reverse(chunk)
			return chunk
		}
		SortFunc(chunk, cmp)
		return chunk
	})

	chunks = Filter(chunks, func(chunk S, index int) bool {
		return len(chunk) > 0
	})

	if isSortedChunksSortedBy(chunks, cmp) {
		return
	}

	chunksCopy := Do(chunks, len(chunks), func(chunk []S, _, _ int) S {
		return slices.Clone(chunk[0])
	})

	slices.SortFunc(chunksCopy, func(a, b S) int {
		return cmp(a[0], b[0])
	})

	if isSortedChunksSortedBy(chunksCopy, cmp) {
		parCopyChunks(slice, chunksCopy)
		return
	}

	parMergeByMergeFunc(chunksCopy, slice, minHeapMergeFunc[S, E], cmp)
}

// ParSortStableFunc performs a parallel stable sort on the given slice using a custom comparison function.
// It maintains the relative order of equal elements while sorting in parallel for improved performance.
func ParSortStableFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) {
	parSortStableByMerge(slice, minHeapMergeStableFunc[S, E], cmp)
}

func parSortStableByMerge[S ~[]E, E any](
	slice S,
	merge func(sortedChunks []S, dest S, cmp func(a, b E) int),
	cmp func(a, b E) int,
) {
	numThreads := utils.NumThreads(runtime.GOMAXPROCS(0))

	chunks := Do(slice, numThreads, func(chunk S, _, _ int) S {
		SortStableFunc(chunk, cmp)
		return chunk
	})

	chunks = Filter(chunks, func(chunk S, index int) bool {
		return len(chunk) > 0
	})

	if isSortedChunksSortedBy(chunks, cmp) {
		return
	}

	chunksCopy := Do(chunks, len(chunks), func(chunk []S, _, _ int) S {
		return slices.Clone(chunk[0])
	})

	// Cannot use parallel merge because it's not stable.
	merge(chunksCopy, slice, cmp)

}

func minHeapMerge[S ~[]E, E constraints.Ordered](sortedChunks []S, dest S) {
	if len(dest) == 0 {
		return
	}

	h := slices.Clone(sortedChunks)

	size := len(h)

	destIndex := 0

	// For initializing the min-heap, a sorting approach is used instead of heapify.
	// While heapify is generally more efficient for building a heap, the small input size
	// in this context (number of CPU cores) makes sorting a suitable and simpler option.
	slices.SortFunc(h, func(a, b S) int {
		if len(a) == 0 {
			return 1
		}

		if len(b) == 0 {
			return -1
		}

		return cmp.Compare(a[0], b[0])
	})

	less := func(i, j int) bool {
		if len(h[i]) == 0 {
			return false
		}

		if len(h[j]) == 0 {
			return true
		}

		return h[i][0] < h[j][0]
	}

	sortedChunksLeft := size

	var min *S

	for {
		min = &h[0]

		dest[destIndex] = (*min)[0]

		if destIndex++; destIndex == len(dest) {
			break
		}

		*min = (*min)[1:]

		if len(*min) == 0 {
			sortedChunksLeft--
		}

		// down procedure
		i := 0

		for {
			left := 2*i + 1

			if left >= size {
				break
			}

			min := left

			if right := left + 1; right < size && less(right, left) {
				min = right
			}

			if !less(min, i) {
				break
			}

			h[min], h[i] = h[i], h[min]

			i = min
		}

		if sortedChunksLeft == 1 {
			break
		}
	}

	copy(dest[destIndex:], *min)
}

func minHeapMergeFunc[S ~[]E, E any](sortedChunks []S, dest S, cmp func(a, b E) int) {
	if len(dest) == 0 {
		return
	}

	h := slices.Clone(sortedChunks)

	size := len(h)

	destIndex := 0

	// For initializing the min-heap, a sorting approach is used instead of heapify.
	// While heapify is generally more efficient for building a heap, the small input size
	// in this context (number of CPU cores) makes sorting a suitable and simpler option.
	slices.SortStableFunc(h, func(a, b S) int {
		if len(a) == 0 {
			return 1
		}

		if len(b) == 0 {
			return -1
		}

		return cmp(a[0], b[0])
	})

	less := func(i, j int) bool {
		if len(h[i]) == 0 {
			return false
		}

		if len(h[j]) == 0 {
			return true
		}

		return cmp(h[i][0], h[j][0]) < 0
	}

	sortedChunksLeft := size

	var min *S

	for {
		min = &h[0]

		dest[destIndex] = (*min)[0]

		if destIndex++; destIndex == len(dest) {
			break
		}

		*min = (*min)[1:]

		if len(*min) == 0 {
			sortedChunksLeft--
		}

		// down procedure
		i := 0

		for {
			left := 2*i + 1

			if left >= size {
				break
			}

			min := left

			if right := left + 1; right < size && less(right, left) {
				min = right
			}

			if !less(min, i) {
				break
			}

			h[min], h[i] = h[i], h[min]

			i = min
		}

		if sortedChunksLeft == 1 {
			break
		}
	}

	copy(dest[destIndex:], *min)
}

func minHeapMergeStableFunc[S ~[]E, E any](sortedChunks []S, dest S, cmp func(a, b E) int) {
	if len(dest) == 0 {
		return
	}

	h := make([]node[S, E], len(sortedChunks))

	for chunkIndex, sortedChunk := range sortedChunks {
		h[chunkIndex] = node[S, E]{
			sortedChunk: sortedChunk,
			chunkIndex:  chunkIndex,
		}
	}

	size := len(h)

	destIndex := 0

	// For initializing the min-heap, a sorting approach is used instead of heapify.
	// While heapify is generally more efficient for building a heap, the small input size
	// in this context (number of CPU cores) makes sorting a suitable and simpler option.
	slices.SortStableFunc(h, func(a, b node[S, E]) int {
		if len(a.sortedChunk) == 0 {
			return 1
		}

		if len(b.sortedChunk) == 0 {
			return -1
		}

		return cmp(a.sortedChunk[0], b.sortedChunk[0])
	})

	less := func(i, j int) bool {
		if len(h[i].sortedChunk) == 0 {
			return false
		}

		if len(h[j].sortedChunk) == 0 {
			return true
		}

		r := cmp(h[i].sortedChunk[0], h[j].sortedChunk[0])

		return (r < 0) || (r == 0 && h[i].chunkIndex < h[j].chunkIndex)
	}

	sortedChunksLeft := size

	var min *S

	for {
		min = &h[0].sortedChunk

		dest[destIndex] = (*min)[0]

		if destIndex++; destIndex == len(dest) {
			break
		}

		*min = (*min)[1:]

		if len(*min) == 0 {
			sortedChunksLeft--
		}

		// down procedure
		i := 0

		for {
			left := 2*i + 1

			if left >= size {
				break
			}

			min := left

			if right := left + 1; right < size && less(right, left) {
				min = right
			}

			if !less(min, i) {
				break
			}

			h[min], h[i] = h[i], h[min]

			i = min
		}

		if sortedChunksLeft == 1 {
			break
		}
	}

	copy(dest[destIndex:], *min)
}

// Is not stable
func parMergeByMerge[S ~[]E, E constraints.Ordered](
	sortedChunks []S,
	dest S,
	merge func(sortedChunks []S, dest S),
) {
	if len(sortedChunks) == 0 {
		return
	}

	numThreads := utils.NumThreads(runtime.GOMAXPROCS(0))

	if numThreads == 1 {
		merge(sortedChunks, dest)
		return
	}

	l := len(sortedChunks)

	bigChunkIndex := 0

	for i := 1; i < l; i++ {
		if len(sortedChunks[i]) > len(sortedChunks[bigChunkIndex]) {
			bigChunkIndex = i
		}
	}

	bigChunk := sortedChunks[bigChunkIndex]

	if len(bigChunk) == 0 {
		return
	}

	otherChunks := slices.Concat(sortedChunks[0:bigChunkIndex], sortedChunks[bigChunkIndex+1:])

	sortedChunks = slices.Concat([]S{sortedChunks[bigChunkIndex]}, otherChunks)

	bigChunkIndex = 0

	splitSize := len(bigChunk) / numThreads

	lastIndexes := make([]int, l)
	lastDestIndex := 0

	var wg sync.WaitGroup

	for i := 0; i < numThreads-1; i++ {
		splitIndex := (i + 1) * splitSize

		a := bigChunk[splitIndex]
		indexes := make([]int, l)
		indexes[0] = splitIndex

		for chunkIndex, chunk := range otherChunks {
			bSearchIndex, _ := slices.BinarySearch(chunk, a)
			indexes[chunkIndex+1] = bSearchIndex
		}

		t := 0
		chunks := make([]S, l)

		for k, index := range indexes {
			t += index

			lastIndex := lastIndexes[k]
			if i > 0 && k == 0 {
				lastIndex++
			}

			if lastIndex < index {
				chunks[k] = sortedChunks[k][lastIndex:index]
			}

			lastIndexes[k] = index
		}

		dest[t] = a

		if i > 0 {
			lastDestIndex++
		}

		if lastDestIndex < t {
			wg.Add(1)
			go func(chs []S, dst S) {
				merge(chs, dst)
				wg.Done()
			}(chunks, dest[lastDestIndex:t])
		}

		lastDestIndex = t
	}

	chunks := make([]S, l)

	for k, lastIndex := range lastIndexes {
		if k == 0 {
			lastIndex++
		}
		chunks[k] = sortedChunks[k][lastIndex:]
	}

	lastDestIndex++

	wg.Add(1)
	go func(chs []S, dst S) {
		merge(chs, dst)
		wg.Done()
	}(chunks, dest[lastDestIndex:])

	wg.Wait()
}

// Is not stable
func parMergeByMergeFunc[S ~[]E, E any](
	sortedChunks []S,
	dest S,
	merge func(sortedChunks []S, dest S, cmp func(a, b E) int),
	cmp func(a, b E) int,
) {
	if len(sortedChunks) == 0 {
		return
	}

	numThreads := utils.NumThreads(runtime.GOMAXPROCS(0))

	if numThreads == 1 {
		merge(sortedChunks, dest, cmp)
		return
	}

	l := len(sortedChunks)

	bigChunkIndex := 0

	for i := 1; i < l; i++ {
		if len(sortedChunks[i]) > len(sortedChunks[bigChunkIndex]) {
			bigChunkIndex = i
		}
	}

	bigChunk := sortedChunks[bigChunkIndex]

	if len(bigChunk) == 0 {
		return
	}

	otherChunks := slices.Concat(sortedChunks[0:bigChunkIndex], sortedChunks[bigChunkIndex+1:])

	sortedChunks = slices.Concat([]S{sortedChunks[bigChunkIndex]}, otherChunks)

	bigChunkIndex = 0

	splitSize := len(bigChunk) / numThreads

	lastIndexes := make([]int, l)
	lastDestIndex := 0

	var wg sync.WaitGroup

	for i := 0; i < numThreads-1; i++ {
		splitIndex := (i + 1) * splitSize

		a := bigChunk[splitIndex]
		indexes := make([]int, l)
		indexes[0] = splitIndex

		for chunkIndex, chunk := range otherChunks {
			bSearchIndex, _ := slices.BinarySearchFunc(chunk, a, cmp)
			indexes[chunkIndex+1] = bSearchIndex
		}

		t := 0
		chunks := make([]S, l)

		for k, index := range indexes {
			t += index

			lastIndex := lastIndexes[k]
			if i > 0 && k == 0 {
				lastIndex++
			}

			if lastIndex < index {
				chunks[k] = sortedChunks[k][lastIndex:index]
			}

			lastIndexes[k] = index
		}

		dest[t] = a

		if i > 0 {
			lastDestIndex++
		}

		if lastDestIndex < t {
			wg.Add(1)
			go func(chs []S, dst S) {
				merge(chs, dst, cmp)
				wg.Done()
			}(chunks, dest[lastDestIndex:t])
		}

		lastDestIndex = t
	}

	chunks := make([]S, l)

	for k, lastIndex := range lastIndexes {
		if k == 0 {
			lastIndex++
		}
		chunks[k] = sortedChunks[k][lastIndex:]
	}

	lastDestIndex++

	wg.Add(1)
	go func(chs []S, dst S) {
		merge(chs, dst, cmp)
		wg.Done()
	}(chunks, dest[lastDestIndex:])

	wg.Wait()
}

// Map returns a new slice with the results of applying the given function to each element of the slice.
func Map[S ~[]E, R ~[]T, E, T any](slice S, transform func(item E, index int) T) R {
	result := make(R, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		result = append(result, transform(slice[i], i))
	}

	return result
}

// ParMap returns a new slice with the results of applying the given function to each element of the slice in parallel.
func ParMap[S ~[]E, R ~[]T, E, T any](slice S, transform func(item E, index int) T) R {
	result := make(R, len(slice))

	Do(slice, 0, func(chunk S, _, chunkStartIndex int) int {
		for i, item := range chunk {
			result[chunkStartIndex+i] = transform(item, chunkStartIndex+i)
		}
		return 0
	})

	return result
}
