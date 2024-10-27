package parlo

// import (
// 	"cmp"
// 	"math"
// 	"slices"
// 	"sync"
// 	"sync/atomic"

// 	"github.com/mahdi-shojaee/parlo/internal/constraints"
// )

// func CopyChunks[S ~[]E, E any](dest []E, chunks []S) {
// 	destIndex := 0

// 	for _, chunk := range chunks {
// 		copy(dest[destIndex:], chunk)
// 		destIndex += len(chunk)
// 	}
// }

// func IsEqual[T comparable, S ~[]T](a S, b S) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	for i := 0; i < len(a); i++ {
// 		if a[i] != b[i] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func ParIsEqual[T comparable, S ~[]T](a S, b S, numThreads int) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	if len(a) <= 200_000 {
// 		return IsEqual(a, b)
// 	}

// 	result := Do(a, numThreads, func(chunk S, _, chunkStartIndex int) bool {
// 		other := b[chunkStartIndex : chunkStartIndex+len(chunk)]
// 		return IsEqual(chunk, other)
// 	})

// 	for _, r := range result {
// 		if !r {
// 			return false
// 		}
// 	}

// 	return true
// }

// func IsSortedChunksSorted[T constraints.Ordered, S ~[]T](sortedChunks []S) bool {
// 	chunks := Filter(sortedChunks, func(chunk S, index int) bool {
// 		return len(chunk) > 0
// 	})

// 	if len(chunks) < 2 {
// 		return true
// 	}

// 	for i := 1; i < len(chunks); i++ {
// 		if chunks[i-1][len(chunks[i-1])-1] > chunks[i][0] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func IsSortedChunksSortedBy[S ~[]E, E any](sortedChunks []S, cmp func(a, b E) int) bool {
// 	chunks := Filter(sortedChunks, func(chunk S, index int) bool {
// 		return len(chunk) > 0
// 	})

// 	if len(chunks) < 2 {
// 		return true
// 	}

// 	for i := 1; i < len(chunks); i++ {
// 		if cmp(chunks[i-1][len(chunks[i-1])-1], chunks[i][0]) > 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

// func Sort[T constraints.Ordered, S ~[]T](collection S) {
// 	slices.Sort(collection)
// 	// InsertionSort(collection)
// 	// TimSort(collection)
// }

// func SortBy[S ~[]E, E any](collection S, cmp func(a, b E) int) {
// 	slices.SortFunc(collection, cmp)
// }

// func SortStable[S ~[]E, E any](collection S, cmp func(a, b E) int) {
// 	slices.SortStableFunc(collection, cmp)
// }

// func SimpleMerge[S ~[]E, E constraints.Ordered](sortedChunks []S, dest S) {
// 	if len(dest) == 0 {
// 		return
// 	}

// 	sortedChunks = Filter(sortedChunks, func(chunk S, index int) bool {
// 		return len(chunk) > 0
// 	})

// 	if len(sortedChunks) == 0 {
// 		return
// 	}

// 	indexes := make([]int, len(sortedChunks))

// 	destIndex := 0

// 	////////////////////////////////////////////////////////////////////
// 	// slices.SortFunc(sortedChunks, func(a, b S) int {
// 	// 	return cmp.Compare(a[0], b[0])
// 	// })

// 	// var chunkIndex int
// 	// var midIndex int

// 	// for chunkIndex = 0; chunkIndex < len(sortedChunks); chunkIndex++ {
// 	// 	chunk := sortedChunks[chunkIndex]
// 	// 	if chunkIndex == len(sortedChunks)-1 || chunk[len(chunk)-1] <= sortedChunks[chunkIndex+1][0] {
// 	// 		copy(dest[destIndex:], chunk)
// 	// 		destIndex += len(chunk)
// 	// 	} else {
// 	// 		midIndex, _ = slices.BinarySearch(chunk, sortedChunks[chunkIndex+1][0])
// 	// 		copy(dest[destIndex:], chunk[:midIndex])
// 	// 		destIndex += midIndex
// 	// 		break
// 	// 	}
// 	// }

// 	// if destIndex == len(dest) {
// 	// 	return
// 	// }

// 	// sortedChunks = slices.Concat(
// 	// 	[]S{sortedChunks[chunkIndex][midIndex:]},
// 	// 	sortedChunks[chunkIndex+1:],
// 	// )
// 	// indexes = indexes[chunkIndex:]
// 	////////////////////////////////////////////////////////////////////

// 	for len(sortedChunks) > 1 {
// 		min := sortedChunks[0][indexes[0]]
// 		chunkIndexOfMin := 0

// 		for chunkIndex := 1; chunkIndex < len(sortedChunks); chunkIndex++ {
// 			elem := sortedChunks[chunkIndex][indexes[chunkIndex]]
// 			if elem < min {
// 				min = elem
// 				chunkIndexOfMin = chunkIndex
// 			}
// 		}

// 		dest[destIndex] = min
// 		destIndex++

// 		if destIndex == len(dest) {
// 			break
// 		}

// 		indexes[chunkIndexOfMin]++

// 		if indexes[chunkIndexOfMin] == len(sortedChunks[chunkIndexOfMin]) {
// 			endIndex := len(sortedChunks) - 1

// 			sortedChunks[chunkIndexOfMin] = sortedChunks[endIndex]
// 			sortedChunks = sortedChunks[:endIndex]

// 			indexes[chunkIndexOfMin] = indexes[endIndex]
// 			indexes = indexes[:endIndex]
// 		}
// 	}

// 	copy(dest[destIndex:], sortedChunks[0][indexes[0]:])
// }

// func SimpleMergeStable[T constraints.Ordered, S ~[]T](sortedChunks []S, dest S) {
// 	if len(dest) == 0 {
// 		return
// 	}

// 	sortedChunks = Filter(sortedChunks, func(chunk S, index int) bool {
// 		return len(chunk) > 0
// 	})

// 	indexes := make([]int, len(sortedChunks))

// 	destIndex := 0

// 	for len(sortedChunks) > 1 {
// 		min := sortedChunks[0][indexes[0]]
// 		chunkIndexOfMin := 0

// 		for chunkIndex := 1; chunkIndex < len(sortedChunks); chunkIndex++ {
// 			elem := sortedChunks[chunkIndex][indexes[chunkIndex]]
// 			if elem < min {
// 				min = elem
// 				chunkIndexOfMin = chunkIndex
// 			}
// 		}

// 		dest[destIndex] = min
// 		destIndex++

// 		if destIndex == len(dest) {
// 			break
// 		}

// 		indexes[chunkIndexOfMin]++

// 		if indexes[chunkIndexOfMin] == len(sortedChunks[chunkIndexOfMin]) {
// 			/////////////////////////////////////////////////////////////////////////////////////
// 			// sortedChunks = slices.Delete(sortedChunks, chunkIndexOfMin, chunkIndexOfMin+1)
// 			// indexes = slices.Delete(indexes, chunkIndexOfMin, chunkIndexOfMin+1)
// 			/////////////////////////////////////////////////////////////////////////////////////
// 			l := len(sortedChunks)
// 			for i := chunkIndexOfMin; i < l-1; i++ {
// 				sortedChunks[i], sortedChunks[i+1] = sortedChunks[i+1], sortedChunks[i]
// 				indexes[i], indexes[i+1] = indexes[i+1], indexes[i]
// 			}
// 			sortedChunks = sortedChunks[:l-1]
// 			indexes = indexes[:l-1]
// 			/////////////////////////////////////////////////////////////////////////////////////
// 		}
// 	}

// 	copy(dest[destIndex:], sortedChunks[0][indexes[0]:])
// }

// func SimpleMergeBy[S ~[]E, E any](sortedChunks []S, dest S, cmp func(a, b E) int) {
// 	if len(dest) == 0 {
// 		return
// 	}

// 	sortedChunks = Filter(sortedChunks, func(chunk S, index int) bool {
// 		return len(chunk) > 0
// 	})

// 	indexes := make([]int, len(sortedChunks))

// 	destIndex := 0

// 	////////////////////////////////////////////////////////////////////
// 	// slices.SortFunc(sortedChunks, func(a, b S) int {
// 	// 	return cmp.Compare(a[0], b[0])
// 	// })

// 	// var chunkIndex int
// 	// var midIndex int

// 	// for chunkIndex = 0; chunkIndex < len(sortedChunks); chunkIndex++ {
// 	// 	chunk := sortedChunks[chunkIndex]
// 	// 	if chunkIndex == len(sortedChunks)-1 || chunk[len(chunk)-1] <= sortedChunks[chunkIndex+1][0] {
// 	// 		copy(dest[destIndex:], chunk)
// 	// 		destIndex += len(chunk)
// 	// 	} else {
// 	// 		midIndex, _ = slices.BinarySearch(chunk, sortedChunks[chunkIndex+1][0])
// 	// 		copy(dest[destIndex:], chunk[:midIndex])
// 	// 		destIndex += midIndex
// 	// 		break
// 	// 	}
// 	// }

// 	// if destIndex == len(dest) {
// 	// 	return
// 	// }

// 	// sortedChunks = slices.Concat(
// 	// 	[]S{sortedChunks[chunkIndex][midIndex:]},
// 	// 	sortedChunks[chunkIndex+1:],
// 	// )
// 	// indexes = indexes[chunkIndex:]
// 	////////////////////////////////////////////////////////////////////

// 	for len(sortedChunks) > 1 {
// 		min := sortedChunks[0][indexes[0]]
// 		chunkIndexOfMin := 0

// 		for chunkIndex := 1; chunkIndex < len(sortedChunks); chunkIndex++ {
// 			elem := sortedChunks[chunkIndex][indexes[chunkIndex]]
// 			if cmp(elem, min) < 0 {
// 				min = elem
// 				chunkIndexOfMin = chunkIndex
// 			}
// 		}

// 		dest[destIndex] = min
// 		destIndex++

// 		if destIndex == len(dest) {
// 			break
// 		}

// 		indexes[chunkIndexOfMin]++

// 		if indexes[chunkIndexOfMin] == len(sortedChunks[chunkIndexOfMin]) {
// 			/////////////////////////////////////////////////////////////////////////////////////
// 			endIndex := len(sortedChunks) - 1

// 			sortedChunks[chunkIndexOfMin] = sortedChunks[endIndex]
// 			sortedChunks = sortedChunks[:endIndex]

// 			indexes[chunkIndexOfMin] = indexes[endIndex]
// 			indexes = indexes[:endIndex]
// 			/////////////////////////////////////////////////////////////////////////////////////
// 			// l := len(sortedChunks)
// 			// for i := chunkIndexOfMin; i < l-1; i++ {
// 			// 	sortedChunks[i], sortedChunks[i+1] = sortedChunks[i+1], sortedChunks[i]
// 			// 	indexes[i], indexes[i+1] = indexes[i+1], indexes[i]
// 			// }
// 			// sortedChunks = sortedChunks[:l-1]
// 			// indexes = indexes[:l-1]
// 			/////////////////////////////////////////////////////////////////////////////////////
// 		}
// 	}

// 	copy(dest[destIndex:], sortedChunks[0][indexes[0]:])
// }

// func SimpleMergeStableBy[S ~[]E, E any](sortedChunks []S, dest S, cmp func(a, b E) int) {
// 	if len(dest) == 0 {
// 		return
// 	}

// 	sortedChunks = Filter(sortedChunks, func(chunk S, index int) bool {
// 		return len(chunk) > 0
// 	})

// 	indexes := make([]int, len(sortedChunks))

// 	destIndex := 0

// 	for len(sortedChunks) > 1 {
// 		min := sortedChunks[0][indexes[0]]
// 		chunkIndexOfMin := 0

// 		for chunkIndex := 1; chunkIndex < len(sortedChunks); chunkIndex++ {
// 			elem := sortedChunks[chunkIndex][indexes[chunkIndex]]
// 			if cmp(elem, min) < 0 {
// 				min = elem
// 				chunkIndexOfMin = chunkIndex
// 			}
// 		}

// 		dest[destIndex] = min
// 		destIndex++

// 		if destIndex == len(dest) {
// 			break
// 		}

// 		indexes[chunkIndexOfMin]++

// 		if indexes[chunkIndexOfMin] == len(sortedChunks[chunkIndexOfMin]) {
// 			/////////////////////////////////////////////////////////////////////////////////////
// 			// sortedChunks = slices.Delete(sortedChunks, chunkIndexOfMin, chunkIndexOfMin+1)
// 			// indexes = slices.Delete(indexes, chunkIndexOfMin, chunkIndexOfMin+1)
// 			/////////////////////////////////////////////////////////////////////////////////////
// 			l := len(sortedChunks)
// 			for i := chunkIndexOfMin; i < l-1; i++ {
// 				sortedChunks[i], sortedChunks[i+1] = sortedChunks[i+1], sortedChunks[i]
// 				indexes[i], indexes[i+1] = indexes[i+1], indexes[i]
// 			}
// 			sortedChunks = sortedChunks[:l-1]
// 			indexes = indexes[:l-1]
// 			/////////////////////////////////////////////////////////////////////////////////////
// 		}
// 	}

// 	copy(dest[destIndex:], sortedChunks[0][indexes[0]:])
// }

// func MinHeapMerge[S ~[]E, E constraints.Ordered](sortedChunks []S, dest S) {
// 	if len(dest) == 0 {
// 		return
// 	}

// 	h := slices.Clone(sortedChunks)

// 	size := len(h)

// 	destIndex := 0

// 	// For initializing the min-heap, a sorting approach is used instead of heapify.
// 	// While heapify is generally more efficient for building a heap, the small input size
// 	// in this context (number of CPU cores) makes sorting a suitable and simpler option.
// 	slices.SortFunc(h, func(a, b S) int {
// 		if len(a) == 0 {
// 			return 1
// 		}

// 		if len(b) == 0 {
// 			return -1
// 		}

// 		return cmp.Compare(a[0], b[0])
// 	})

// 	less := func(i, j int) bool {
// 		if len(h[i]) == 0 {
// 			return false
// 		}

// 		if len(h[j]) == 0 {
// 			return true
// 		}

// 		return h[i][0] < h[j][0]
// 	}

// 	sortedChunksLeft := size

// 	var min *S

// 	for {
// 		min = &h[0]

// 		dest[destIndex] = (*min)[0]

// 		if destIndex++; destIndex == len(dest) {
// 			break
// 		}

// 		*min = (*min)[1:]

// 		if len(*min) == 0 {
// 			sortedChunksLeft--
// 		}

// 		// down procedure
// 		i := 0

// 		for {
// 			left := 2*i + 1

// 			if left >= size {
// 				break
// 			}

// 			min := left

// 			if right := left + 1; right < size && less(right, left) {
// 				min = right
// 			}

// 			if !less(min, i) {
// 				break
// 			}

// 			h[min], h[i] = h[i], h[min]

// 			i = min
// 		}

// 		if sortedChunksLeft == 1 {
// 			break
// 		}
// 	}

// 	copy(dest[destIndex:], *min)
// }

// func MinHeapMergeBy[S ~[]E, E any](sortedChunks []S, dest S, cmp func(a, b E) int) {
// 	if len(dest) == 0 {
// 		return
// 	}

// 	h := slices.Clone(sortedChunks)

// 	size := len(h)

// 	destIndex := 0

// 	// For initializing the min-heap, a sorting approach is used instead of heapify.
// 	// While heapify is generally more efficient for building a heap, the small input size
// 	// in this context (number of CPU cores) makes sorting a suitable and simpler option.
// 	slices.SortStableFunc(h, func(a, b S) int {
// 		if len(a) == 0 {
// 			return 1
// 		}

// 		if len(b) == 0 {
// 			return -1
// 		}

// 		return cmp(a[0], b[0])
// 	})

// 	less := func(i, j int) bool {
// 		if len(h[i]) == 0 {
// 			return false
// 		}

// 		if len(h[j]) == 0 {
// 			return true
// 		}

// 		return cmp(h[i][0], h[j][0]) < 0
// 	}

// 	sortedChunksLeft := size

// 	var min *S

// 	for {
// 		min = &h[0]

// 		dest[destIndex] = (*min)[0]

// 		if destIndex++; destIndex == len(dest) {
// 			break
// 		}

// 		*min = (*min)[1:]

// 		if len(*min) == 0 {
// 			sortedChunksLeft--
// 		}

// 		// down procedure
// 		i := 0

// 		for {
// 			left := 2*i + 1

// 			if left >= size {
// 				break
// 			}

// 			min := left

// 			if right := left + 1; right < size && less(right, left) {
// 				min = right
// 			}

// 			if !less(min, i) {
// 				break
// 			}

// 			h[min], h[i] = h[i], h[min]

// 			i = min
// 		}

// 		if sortedChunksLeft == 1 {
// 			break
// 		}
// 	}

// 	copy(dest[destIndex:], *min)
// }

// func MinHeapMergeStableBy[S ~[]E, E any](sortedChunks []S, dest S, cmp func(a, b E) int) {
// 	if len(dest) == 0 {
// 		return
// 	}

// 	type Node struct {
// 		sortedChunk S
// 		chunkIndex  int
// 	}

// 	h := make([]Node, len(sortedChunks))

// 	for chunkIndex, sortedChunk := range sortedChunks {
// 		h[chunkIndex] = Node{
// 			sortedChunk: sortedChunk,
// 			chunkIndex:  chunkIndex,
// 		}
// 	}

// 	size := len(h)

// 	destIndex := 0

// 	// For initializing the min-heap, a sorting approach is used instead of heapify.
// 	// While heapify is generally more efficient for building a heap, the small input size
// 	// in this context (number of CPU cores) makes sorting a suitable and simpler option.
// 	slices.SortStableFunc(h, func(a, b Node) int {
// 		if len(a.sortedChunk) == 0 {
// 			return 1
// 		}

// 		if len(b.sortedChunk) == 0 {
// 			return -1
// 		}

// 		return cmp(a.sortedChunk[0], b.sortedChunk[0])
// 	})

// 	less := func(i, j int) bool {
// 		if len(h[i].sortedChunk) == 0 {
// 			return false
// 		}

// 		if len(h[j].sortedChunk) == 0 {
// 			return true
// 		}

// 		r := cmp(h[i].sortedChunk[0], h[j].sortedChunk[0])

// 		return (r < 0) || (r == 0 && h[i].chunkIndex < h[j].chunkIndex)
// 	}

// 	sortedChunksLeft := size

// 	var min *S

// 	for {
// 		min = &h[0].sortedChunk

// 		dest[destIndex] = (*min)[0]

// 		if destIndex++; destIndex == len(dest) {
// 			break
// 		}

// 		*min = (*min)[1:]

// 		if len(*min) == 0 {
// 			sortedChunksLeft--
// 		}

// 		// down procedure
// 		i := 0

// 		for {
// 			left := 2*i + 1

// 			if left >= size {
// 				break
// 			}

// 			min := left

// 			if right := left + 1; right < size && less(right, left) {
// 				min = right
// 			}

// 			if !less(min, i) {
// 				break
// 			}

// 			h[min], h[i] = h[i], h[min]

// 			i = min
// 		}

// 		if sortedChunksLeft == 1 {
// 			break
// 		}
// 	}

// 	copy(dest[destIndex:], *min)
// }

// // Is not stable
// func ParMergeByMerge[T constraints.Ordered, S ~[]T](
// 	sortedChunks []S,
// 	dest S,
// 	merge func(sortedChunks []S, dest S),
// 	numThreads int,
// ) {
// 	if len(sortedChunks) == 0 {
// 		return
// 	}

// 	if numThreads == 1 {
// 		merge(sortedChunks, dest)
// 		return
// 	}

// 	l := len(sortedChunks)

// 	bigChunkIndex := 0

// 	for i := 1; i < l; i++ {
// 		if len(sortedChunks[i]) > len(sortedChunks[bigChunkIndex]) {
// 			bigChunkIndex = i
// 		}
// 	}

// 	bigChunk := sortedChunks[bigChunkIndex]

// 	if len(bigChunk) == 0 {
// 		return
// 	}

// 	otherChunks := slices.Concat(sortedChunks[0:bigChunkIndex], sortedChunks[bigChunkIndex+1:])

// 	sortedChunks = slices.Concat([]S{sortedChunks[bigChunkIndex]}, otherChunks)

// 	bigChunkIndex = 0

// 	splitSize := len(bigChunk) / numThreads

// 	lastIndexes := make([]int, l)
// 	lastDestIndex := 0

// 	var wg sync.WaitGroup

// 	for i := 0; i < numThreads-1; i++ {
// 		splitIndex := (i + 1) * splitSize

// 		a := bigChunk[splitIndex]
// 		indexes := make([]int, l)
// 		indexes[0] = splitIndex

// 		for chunkIndex, chunk := range otherChunks {
// 			bSearchIndex, _ := slices.BinarySearch(chunk, a)
// 			indexes[chunkIndex+1] = bSearchIndex
// 		}

// 		t := 0
// 		chunks := make([]S, l)

// 		for k, index := range indexes {
// 			t += index

// 			lastIndex := lastIndexes[k]
// 			if i > 0 && k == 0 {
// 				lastIndex++
// 			}

// 			if lastIndex < index {
// 				chunks[k] = sortedChunks[k][lastIndex:index]
// 			}

// 			lastIndexes[k] = index
// 		}

// 		dest[t] = a

// 		if i > 0 {
// 			lastDestIndex++
// 		}

// 		if lastDestIndex < t {
// 			wg.Add(1)
// 			go func(chs []S, dst []T) {
// 				merge(chs, dst)
// 				wg.Done()
// 			}(chunks, dest[lastDestIndex:t])
// 		}

// 		lastDestIndex = t
// 	}

// 	chunks := make([]S, l)

// 	for k, lastIndex := range lastIndexes {
// 		if k == 0 {
// 			lastIndex++
// 		}
// 		chunks[k] = sortedChunks[k][lastIndex:]
// 	}

// 	lastDestIndex++

// 	wg.Add(1)
// 	go func(chs []S, dst []T) {
// 		merge(chs, dst)
// 		wg.Done()
// 	}(chunks, dest[lastDestIndex:])

// 	wg.Wait()
// }

// func PingPongMerge2[T constraints.Ordered, S ~[]T](sortedChunks []S, dst S) {
// 	leftSlice := sortedChunks[0]
// 	rightSlice := sortedChunks[1]

// 	leftSliceLength := len(leftSlice)
// 	rightSliceLength := len(rightSlice)

// 	if leftSliceLength == 0 || rightSliceLength == 0 {
// 		copy(dst, leftSlice)
// 		copy(dst, rightSlice)
// 		return
// 	}

// 	if leftSlice[leftSliceLength-1] <= rightSlice[0] {
// 		copy(dst, leftSlice)
// 		copy(dst, rightSlice)
// 		return
// 	}

// 	if rightSlice[rightSliceLength-1] <= leftSlice[0] {
// 		copy(dst, rightSlice)
// 		copy(dst, leftSlice)
// 		return
// 	}

// 	leftIndex := 0
// 	rightIndex := 0

// 	for leftIndex < leftSliceLength && rightIndex < rightSliceLength {
// 		if leftSlice[leftIndex] < rightSlice[rightIndex] {
// 			dst[leftIndex+rightIndex] = leftSlice[leftIndex]
// 			leftIndex++
// 		} else {
// 			dst[leftIndex+rightIndex] = rightSlice[rightIndex]
// 			rightIndex++
// 		}
// 	}

// 	for ; leftIndex < leftSliceLength; leftIndex++ {
// 		dst[leftIndex+rightIndex] = leftSlice[leftIndex]
// 	}

// 	for ; rightIndex < rightSliceLength; rightIndex++ {
// 		dst[leftIndex+rightIndex] = rightSlice[rightIndex]
// 	}
// }

// func ParMerge1[S ~[]E, E constraints.Ordered](sortedChunks []S, dest S, numThreads int) {
// 	ParMergeByMerge(sortedChunks, dest, MinHeapMerge[S, E], numThreads)
// }

// func ParMerge3[S ~[]E, E constraints.Ordered](sortedChunks []S, dest S, numThreads int) {
// 	ParMergeByMerge(sortedChunks, dest, PingPongMerge2[E, S], numThreads)
// }

// func ParSortByMerge1[S ~[]E, E constraints.Ordered](
// 	collection S,
// 	merge func(sortedChunks []S, dest S),
// 	numThreads int,
// ) {
// 	if len(collection) <= 200_000 {
// 		Sort(collection)
// 		return
// 	}

// 	isReversed := ParIsReversed(collection, numThreads)

// 	if isReversed {
// 		ParReverse(collection, 2)
// 		return
// 	}

// 	chunks := Do(collection, numThreads, func(chunk S, _, _ int) []E {
// 		if IsSorted(chunk) {
// 			return chunk
// 		}
// 		Sort(chunk)
// 		return chunk
// 	})

// 	chunks = Filter(chunks, func(chunk []E, index int) bool {
// 		return len(chunk) > 0
// 	})

// 	if IsSortedChunksSorted(chunks) {
// 		return
// 	}

// 	chunksCopy := Do(chunks, len(chunks), func(chunk [][]E, _, _ int) []E {
// 		return slices.Clone(chunk[0])
// 	})

// 	slices.SortFunc(chunksCopy, func(a, b []E) int {
// 		return cmp.Compare(a[0], b[0])
// 	})

// 	if IsSortedChunksSorted(chunksCopy) {
// 		CopyChunks(collection, chunksCopy)
// 		return
// 	}

// 	ParMergeByMerge(chunksCopy, collection, MinHeapMerge[[]E, E], numThreads)
// }

// func PingPongMerge[T constraints.Ordered, S ~[]T](src S, dst S, startIndex int, midIndex int, endIndex int) {
// 	leftSlice := src[startIndex:midIndex]
// 	rightSlice := src[midIndex:endIndex]

// 	leftSliceLength := len(leftSlice)
// 	rightSliceLength := len(rightSlice)

// 	if leftSliceLength == 0 || rightSliceLength == 0 {
// 		copy(dst[startIndex:endIndex], src[startIndex:endIndex])
// 		return
// 	}

// 	k := 0

// 	///////////////////////////////////////////////////////////////////////
// 	// k = -1
// 	// x := rightSlice[0]

// 	// for i := 0; i < leftSliceLength && i < 10; i++ {
// 	// 	if leftSlice[i] == x {
// 	// 		k = i
// 	// 		break
// 	// 	}
// 	// }

// 	// if k == -1 {
// 	// 	k, _ = slices.BinarySearch(leftSlice, x)
// 	// }

// 	// copy(dst[startIndex:], src[startIndex:startIndex+k])
// 	// leftSlice = leftSlice[k:]
// 	// leftSliceLength = len(leftSlice)

// 	// if leftSliceLength == 0 {
// 	// 	copy(dst[startIndex+k:], src[startIndex+k:endIndex])
// 	// 	return
// 	// }
// 	///////////////////////////////////////////////////////////////////////

// 	if leftSlice[leftSliceLength-1] <= rightSlice[0] {
// 		copy(dst[startIndex+k:], src[startIndex+k:endIndex])
// 		return
// 	}

// 	if rightSlice[rightSliceLength-1] <= leftSlice[0] {
// 		copy(dst[startIndex+k:], rightSlice)
// 		copy(dst[startIndex+k+rightSliceLength:], leftSlice)
// 		return
// 	}

// 	dst = dst[startIndex+k : endIndex]

// 	leftIndex := 0
// 	rightIndex := 0

// 	for leftIndex < leftSliceLength && rightIndex < rightSliceLength {
// 		if leftSlice[leftIndex] < rightSlice[rightIndex] {
// 			dst[leftIndex+rightIndex] = leftSlice[leftIndex]
// 			leftIndex++
// 		} else {
// 			dst[leftIndex+rightIndex] = rightSlice[rightIndex]
// 			rightIndex++
// 		}
// 	}

// 	for ; leftIndex < leftSliceLength; leftIndex++ {
// 		dst[leftIndex+rightIndex] = leftSlice[leftIndex]
// 	}

// 	for ; rightIndex < rightSliceLength; rightIndex++ {
// 		dst[leftIndex+rightIndex] = rightSlice[rightIndex]
// 	}
// }

// func ParSortByMerge2[T constraints.Ordered, S ~[]T](collection S, numThreads int) {
// 	if len(collection) <= 200_000 {
// 		Sort(collection)
// 		return
// 	}

// 	chunks := Do(collection, numThreads, func(chunk S, _, _ int) []T {
// 		Sort(chunk)
// 		return chunk
// 	})

// 	chunkLengths := make([]int, len(chunks))

// 	for i, chunk := range chunks {
// 		chunkLengths[i] = len(chunk)
// 	}

// 	chunks = Filter(chunks, func(chunk []T, index int) bool {
// 		return len(chunk) > 0
// 	})

// 	if IsSortedChunksSorted(chunks) {
// 		return
// 	}

// 	// Ping Pong Merge
// 	///////////////////////////////////////////////////////
// 	tmp := make(S, len(collection))

// 	src := collection
// 	dst := tmp

// 	n := int(math.Log2(float64(numThreads-1))) + 1

// 	for i := 0; i < n; i++ {
// 		left := 0

// 		for k := 0; k < len(chunkLengths); k += 2 {
// 			mid := left + chunkLengths[k]

// 			if k+1 == len(chunkLengths) {
// 				copy(dst[left:mid], src[left:mid])
// 				break
// 			}

// 			right := mid + chunkLengths[k+1]

// 			PingPongMerge(src, dst, left, mid, right)

// 			left += chunkLengths[k] + chunkLengths[k+1]
// 		}

// 		for k := 0; k < len(chunkLengths); k += 2 {
// 			chunkLengths[k/2] = chunkLengths[k]
// 			if k+1 < len(chunkLengths) {
// 				chunkLengths[k/2] += chunkLengths[k+1]
// 			}
// 		}

// 		chunkLengths = chunkLengths[:len(chunkLengths)/2+len(chunkLengths)%2]

// 		src, dst = dst, src
// 	}

// 	if n%2 == 1 {
// 		copy(collection, tmp)
// 	}
// }

// func ParSortByMerge3[T constraints.Ordered, S ~[]T](collection S, numThreads int) {
// 	if len(collection) <= 200_000 {
// 		Sort(collection)
// 		return
// 	}

// 	chunks := Do(collection, numThreads, func(chunk S, _, _ int) []T {
// 		Sort(chunk)
// 		return chunk
// 	})

// 	chunkLengths := make([]int, len(chunks))

// 	for i, chunk := range chunks {
// 		chunkLengths[i] = len(chunk)
// 	}

// 	chunks = Filter(chunks, func(chunk []T, index int) bool {
// 		return len(chunk) > 0
// 	})

// 	if IsSortedChunksSorted(chunks) {
// 		return
// 	}

// 	// Ping Pong Merge
// 	///////////////////////////////////////////////////////
// 	tmp := make(S, len(collection))

// 	src := collection
// 	dst := tmp

// 	var wg sync.WaitGroup

// 	n := int(math.Log2(float64(numThreads-1))) + 1

// 	for i := 0; i < n; i++ {
// 		left := 0

// 		for k := 0; k < len(chunkLengths); k += 2 {
// 			mid := left + chunkLengths[k]

// 			if k+1 == len(chunkLengths) {
// 				wg.Add(1)
// 				go func() {
// 					copy(dst[left:mid], src[left:mid])
// 					wg.Done()
// 				}()
// 				break
// 			}

// 			right := mid + chunkLengths[k+1]

// 			wg.Add(1)
// 			go func(left, mid, right int) {
// 				PingPongMerge(src, dst, left, mid, right)
// 				wg.Done()
// 			}(left, mid, right)

// 			left += chunkLengths[k] + chunkLengths[k+1]
// 		}

// 		wg.Wait()

// 		for k := 0; k < len(chunkLengths); k += 2 {
// 			chunkLengths[k/2] = chunkLengths[k]
// 			if k+1 < len(chunkLengths) {
// 				chunkLengths[k/2] += chunkLengths[k+1]
// 			}
// 		}

// 		chunkLengths = chunkLengths[:len(chunkLengths)/2+len(chunkLengths)%2]

// 		src, dst = dst, src
// 	}

// 	if n%2 == 1 {
// 		copy(collection, tmp)
// 	}
// }

// func ParSortStableByMerge[S ~[]E, E any](
// 	collection S,
// 	merge func(sortedChunks []S, dest S, cmp func(a, b E) int),
// 	cmp func(a, b E) int,
// 	numThreads int,
// ) {
// 	if len(collection) <= 200_000 {
// 		SortStable(collection, cmp)
// 		return
// 	}

// 	chunks := Do(collection, numThreads, func(chunk S, _, _ int) S {
// 		SortStable(chunk, cmp)
// 		return chunk
// 	})

// 	chunks = Filter(chunks, func(chunk S, index int) bool {
// 		return len(chunk) > 0
// 	})

// 	if IsSortedChunksSortedBy(chunks, cmp) {
// 		return
// 	}

// 	/////////////////////////////////////////
// 	// chunksCopy := make([]S, len(chunks))
// 	// for i, chunk := range chunks {
// 	// 	chunksCopy[i] = slices.Clone(chunk)
// 	// }
// 	/////////////////////////////////////////
// 	chunksCopy := Do(chunks, len(chunks), func(chunk []S, _, _ int) S {
// 		return slices.Clone(chunk[0])
// 	})
// 	/////////////////////////////////////////

// 	// Cannot use parallel merge because it's not stable.
// 	merge(chunksCopy, collection, cmp)

// }

// func ParSort1[S ~[]E, E constraints.Ordered](collection S, numThreads int) {
// 	ParSortByMerge1(collection, MinHeapMerge[S, E], numThreads)
// }

// func ParSort2[S ~[]E, E constraints.Ordered](collection S, numThreads int) {
// 	ParSortByMerge2(collection, numThreads)
// }

// func ParSort3[S ~[]E, E constraints.Ordered](collection S, numThreads int) {
// 	ParSortByMerge3(collection, numThreads)
// }

// func ParSortStable[S ~[]E, E any](collection S, cmp func(a, b E) int, numThreads int) {
// 	ParSortStableByMerge(collection, MinHeapMergeStableBy[S, E], cmp, numThreads)
// }
