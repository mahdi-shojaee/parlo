package parlo_test

// import (
// 	"cmp"
// 	"fmt"
// 	"math/rand/v2"
// 	"slices"
// 	"testing"
// 	"time"

// 	"github.com/mahdi-shojaee/parlo"
// 	"github.com/stretchr/testify/assert"
// )

// type Person struct {
// 	name string
// 	id   int
// }

// func TestIsReversed(t *testing.T) {
// 	t.Run("should return true", func(t *testing.T) {
// 		elems := []Elem{10, 9, 6, 5, 4, 2, 1}
// 		assert.True(t, parlo.IsReversed(elems))
// 	})

// 	t.Run("should return false", func(t *testing.T) {
// 		elems := []Elem{1, 2, 4, 1, 6, 9, 10}
// 		assert.False(t, parlo.IsReversed(elems))
// 	})
// }

// func TestParIsReversed(t *testing.T) {
// 	for threads := 2; threads <= 2; threads++ {
// 		name := fmt.Sprintf("threads=%d should return true", threads)

// 		t.Run(name, func(t *testing.T) {
// 			elems := []Elem{10, 9, 6, 5, 4, 2, 1}
// 			assert.True(t, parlo.ParIsReversed(elems, threads))
// 		})

// 		name = fmt.Sprintf("threads=%d should return false", threads)

// 		t.Run(name, func(t *testing.T) {
// 			elems := []Elem{6, 5, 4, 3, 2, 1, 12, 11, 10, 9, 8, 7}
// 			assert.False(t, parlo.ParIsReversed(elems, threads))
// 		})
// 	}
// }

// func testMerge[S ~[]T, T any](
// 	t *testing.T,
// 	getName func(sortedChunks []Elems, length int) string,
// 	merge func(sortedChunks []Elems, dest Elems),
// ) {
// 	testCases := [][]Elems{
// 		{},
// 		{
// 			{},
// 		},
// 		{
// 			{2},
// 		},
// 		{
// 			{}, {}, {},
// 		},
// 		{
// 			{2}, {1}, {},
// 		},
// 		{
// 			{1}, {2, 4}, {3, 5},
// 		},
// 		{
// 			{1, 3, 5, 7}, {2, 4, 6, 8},
// 		},
// 		{
// 			{10, 40, 70}, {20, 50, 80}, {30, 60, 90},
// 		},
// 		{
// 			{1, 13, 27, 45}, {2}, {3, 5, 11, 14, 17, 19, 22},
// 		},
// 		{
// 			{2, 4, 9, 14},
// 			{1, 10, 11, 15},
// 			{5, 6, 8, 16},
// 			{3, 7, 12, 13},
// 		},
// 		{
// 			{2, 3, 4, 18, 23, 30},
// 			{15, 17, 19, 21, 22},
// 			{1, 2, 4, 5, 6, 9, 10},
// 			{7, 11, 16, 18, 20},
// 		},
// 		{
// 			{6, 7, 8, 9, 10},
// 			{11, 12, 18, 19, 20},
// 			{1, 2, 3, 4, 5},
// 			{13, 14, 15, 16, 17},
// 		},
// 		{
// 			{1, 2, 3, 4, 5},
// 			{6, 7, 8, 9, 10},
// 			{11, 14, 15, 16, 17},
// 			{18, 19, 20, 21, 22},
// 		},
// 		{
// 			{1, 2, 3, 4, 5},
// 			{6, 7, 8, 9, 10},
// 			{11, 12, 18, 19, 20},
// 			{13, 14, 15, 16, 17},
// 			{18, 28, 30, 31, 32},
// 			{33, 34, 35, 36, 37},
// 		},
// 	}

// 	makeDest := func(sortedChunks []Elems) Elems {
// 		length := 0

// 		for _, chunk := range sortedChunks {
// 			length += len(chunk)
// 		}

// 		if length == 0 {
// 			return nil
// 		}

// 		return make([]Elem, length)
// 	}

// 	for _, testCase := range testCases {
// 		for i := 0; i < len(testCase); i++ {
// 			sortedChunks := testCase[0 : i+1]
// 			dest := makeDest(sortedChunks)
// 			name := getName(sortedChunks, len(dest))

// 			expected := slices.Concat(sortedChunks...)
// 			slices.Sort(expected)

// 			t.Run(name, func(t *testing.T) {
// 				oldSortedChunks := slices.Clone(sortedChunks)
// 				merge(sortedChunks, dest)
// 				assert.Equal(
// 					t,
// 					oldSortedChunks,
// 					sortedChunks,
// 					"merge should not mutate the original sorted chunks",
// 				)
// 				assert.Equal(t, expected, dest)
// 			})
// 		}
// 	}
// }

// func TestMerge(t *testing.T) {
// 	mergeFuncs := []struct {
// 		fnName string
// 		fn     func(sortedChunks []Elems, dest Elems)
// 	}{
// 		{"SimpleMerge", parlo.SimpleMerge[Elems, Elem]},
// 		{"MinHeapMerge", parlo.MinHeapMerge[Elems, Elem]},
// 	}

// 	for _, m := range mergeFuncs {
// 		getName := func(sortedChunks []Elems, length int) string {
// 			return fmt.Sprintf("%s chunks=%d len=%d", m.fnName, len(sortedChunks), length)
// 		}

// 		testMerge[Elems, Elem](t, getName, m.fn)
// 	}

// }

// func TestMergeBy(t *testing.T) {
// 	mergeFuncs := []struct {
// 		fnName string
// 		fn     func(sortedChunks []Elems, dest Elems)
// 	}{
// 		{"SimpleMergeBy", func(sortedChunks []Elems, dest Elems) {
// 			parlo.SimpleMergeBy(sortedChunks, dest, cmp.Compare[Elem])
// 		}},
// 		{"MinHeapMergeBy", func(sortedChunks []Elems, dest Elems) {
// 			parlo.MinHeapMergeBy(sortedChunks, dest, cmp.Compare[Elem])
// 		}},
// 	}

// 	for _, m := range mergeFuncs {
// 		getName := func(sortedChunks []Elems, length int) string {
// 			return fmt.Sprintf("%s chunks=%d len=%d", m.fnName, len(sortedChunks), length)
// 		}

// 		testMerge[Elems, Elem](t, getName, m.fn)
// 	}

// }

// func TestMergeStableBy(t *testing.T) {
// 	mergeFuncs := []struct {
// 		fnName string
// 		fn     func(sortedChunks []Elems, dest Elems)
// 	}{
// 		{"SimpleMergeStableBy", func(sortedChunks []Elems, dest Elems) {
// 			parlo.SimpleMergeStableBy(sortedChunks, dest, cmp.Compare[Elem])
// 		}},
// 		{"MinHeapMergeStableBy", func(sortedChunks []Elems, dest Elems) {
// 			parlo.MinHeapMergeStableBy(sortedChunks, dest, cmp.Compare[Elem])
// 		}},
// 	}

// 	for _, m := range mergeFuncs {
// 		getName := func(sortedChunks []Elems, length int) string {
// 			return fmt.Sprintf("%s chunks=%d len=%d", m.fnName, len(sortedChunks), length)
// 		}

// 		testMerge[Elems, Elem](t, getName, m.fn)
// 	}

// }

// func TestMergePar1(t *testing.T) {
// 	for threads := 1; threads <= MAX_THREADS; threads++ {
// 		getName := func(sortedChunks []Elems, length int) string {
// 			return fmt.Sprintf("threads=%d chunks=%d len=%d", threads, len(sortedChunks), length)
// 		}

// 		merge := func(sortedChunks []Elems, dest Elems) {
// 			parlo.ParMerge1(sortedChunks, dest, threads)
// 		}

// 		testMerge[Elems, Elem](t, getName, merge)
// 	}
// }

// // func TestMergePar2(t *testing.T) {
// // 	testCases := []struct {
// // 		slice      []Elem
// // 		startIndex int
// // 		midIndex   int
// // 		endIndex   int
// // 	}{
// // 		{[]Elem{}, 0, 0, 0},
// // 		{[]Elem{2, 1}, 0, 1, 2},
// // 		{[]Elem{1, 2, 4}, 0, 1, 3},
// // 		{[]Elem{16, 7, 8, 9, 0, 1, 2, 3, 0, 5}, 1, 4, 8},
// // 		{[]Elem{12, 3, 5, 7, 9, 10, 12, 4, 6, 8}, 2, 4, 7},
// // 		{[]Elem{12, 3, 5, 7, 9, 10, 12, 4, 6, 8, 14}, 1, 7, 10},
// // 		{[]Elem{10, 2, 13, 27, 35, 43, 5, 11, 14, 17, 19, 22}, 3, 6, 10},
// // 		{[]Elem{2, 4, 9, 14, 1, 10, 11, 15}, 2, 4, 6},
// // 		{[]Elem{18, 17, 19, 21, 22, 27, 11, 16, 21, 20}, 1, 6, 9},
// // 		{[]Elem{18, 17, 19, 2, 4, 5, 11, 16, 21, 20}, 1, 3, 9},
// // 	}

// // 	for threads := 1; threads <= MAX_THREADS; threads++ {
// // 		for _, testCase := range testCases {
// // 			expected := slices.Clone(testCase.slice)
// // 			slices.Sort(expected[testCase.startIndex:testCase.endIndex])

// // 			t.Run("PingPongMerge", func(t *testing.T) {
// // 				dst := make([]Elem, len(testCase.slice))
// // 				copy(dst[:testCase.startIndex], testCase.slice[:testCase.startIndex])
// // 				copy(dst[testCase.endIndex:], testCase.slice[testCase.endIndex:])
// // 				parlo.ParMerge2(testCase.slice, dst, testCase.startIndex, testCase.midIndex, testCase.endIndex, threads)
// // 				assert.Equal(t, expected, dst)
// // 			})
// // 		}
// // 	}
// // }

// // func TestMergePar2(t *testing.T) {
// // 	testCases := []struct {
// // 		slice      []Elem
// // 		startIndex int
// // 		midIndex   int
// // 		endIndex   int
// // 	}{
// // 		{[]Elem{}, 0, 0, 0},
// // 		{[]Elem{2, 1}, 0, 1, 2},
// // 		{[]Elem{1, 2, 4}, 0, 1, 3},
// // 		{[]Elem{16, 7, 8, 9, 0, 1, 2, 3, 0, 5}, 1, 4, 8},
// // 		{[]Elem{12, 3, 5, 7, 9, 10, 12, 4, 6, 8}, 2, 4, 7},
// // 		{[]Elem{12, 3, 5, 7, 9, 10, 12, 4, 6, 8, 14}, 1, 7, 10},
// // 		{[]Elem{10, 2, 13, 27, 35, 43, 5, 11, 14, 17, 19, 22}, 3, 6, 10},
// // 		{[]Elem{2, 4, 9, 14, 1, 10, 11, 15}, 2, 4, 6},
// // 		{[]Elem{18, 17, 19, 21, 22, 27, 11, 16, 21, 20}, 1, 6, 9},
// // 		{[]Elem{18, 17, 19, 2, 4, 5, 11, 16, 21, 20}, 1, 3, 9},
// // 	}

// // 	for threads := 1; threads <= MAX_THREADS; threads++ {
// // 		for _, testCase := range testCases {
// // 			expected := slices.Clone(testCase.slice)
// // 			slices.Sort(expected[testCase.startIndex:testCase.endIndex])

// // 			t.Run("PingPongMerge", func(t *testing.T) {
// // 				dst := make([]Elem, len(testCase.slice))
// // 				copy(dst[:testCase.startIndex], testCase.slice[:testCase.startIndex])
// // 				copy(dst[testCase.endIndex:], testCase.slice[testCase.endIndex:])
// // 				parlo.ParMerge2(testCase.slice, dst, testCase.startIndex, testCase.midIndex, testCase.endIndex, threads)
// // 				assert.Equal(t, expected, dst)
// // 			})
// // 		}
// // 	}
// // }

// func TestParReverse(t *testing.T) {
// 	testCases := []Elems{
// 		{1, 4, 2, 5, 8, 3, 6, 9, 10},
// 		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
// 		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
// 		{1, 2, 3, 7, 4, 0, 6, 5, 8, 9},
// 		{3, 1, 2, 5, 0, 4, 9, 8, 7, 6},
// 		{0, 3, 8, 1, 9, 7, 6, 5, 4, 2},
// 		{9, 1, 2, 3, 4, 5, 6, 11, 8, 0, 10, 7},
// 		{3, 4, 1, 2, 5, 6, 7, 8, 0, 9, 10, 11, 12, 14},
// 	}

// 	for i := 0; i < 10; i++ {
// 		slice := MakeSemiSortedCollection(32 + rand.IntN(1_000))
// 		testCases = append(testCases, slice)
// 	}

// 	fns := []struct {
// 		name string
// 		fn   func(collection Elems, threads int)
// 	}{
// 		{"ParReverse", parlo.ParReverse[Elems, Elem]},
// 	}

// 	for _, slice := range testCases {
// 		sliceCopy := make(Elems, len(slice))

// 		expected := slices.Clone(slice)
// 		slices.Reverse(expected)

// 		for threads := 1; threads <= MAX_THREADS; threads++ {
// 			for _, m := range fns {
// 				name := fmt.Sprintf("%s threads=%d len=%d", m.name, threads, len(slice))

// 				t.Run(name, func(t *testing.T) {
// 					copy(sliceCopy, slice)
// 					m.fn(sliceCopy, threads)
// 					assert.Equal(t, expected, sliceCopy)
// 				})
// 			}
// 		}
// 	}
// }

// func TestParSort(t *testing.T) {
// 	testCases := []Elems{
// 		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
// 		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
// 		{1, 2, 3, 7, 4, 0, 6, 5, 8, 9},
// 		{3, 1, 2, 5, 0, 4, 9, 8, 7, 6},
// 		{0, 3, 8, 1, 9, 7, 6, 5, 4, 2},
// 		{9, 1, 2, 3, 4, 5, 6, 11, 8, 0, 10, 7},
// 		{1, 4, 2, 5, 8, 3, 6, 9, 10},
// 		{3, 4, 1, 2, 5, 6, 7, 8, 0, 9, 10, 11, 12, 14},
// 	}

// 	for i := 0; i < 100; i++ {
// 		slice := MakeSemiSortedCollection(32 + rand.IntN(1_000))
// 		testCases = append(testCases, slice)
// 	}

// 	sortFns := []struct {
// 		name string
// 		fn   func(collection Elems, threads int)
// 	}{
// 		{"SortPar1", parlo.ParSort1[Elems, Elem]},
// 		{"SortPar2", parlo.ParSort2[Elems, Elem]},
// 		{"SortPar3", parlo.ParSort3[Elems, Elem]},
// 	}

// 	for _, slice := range testCases {
// 		sliceCopy := make(Elems, len(slice))

// 		expected := slices.Clone(slice)
// 		slices.Sort(expected)

// 		for threads := 1; threads <= MAX_THREADS; threads++ {
// 			for _, m := range sortFns {
// 				name := fmt.Sprintf("%s threads=%d len=%d", m.name, threads, len(slice))
// 				t.Run(name, func(t *testing.T) {
// 					copy(sliceCopy, slice)
// 					m.fn(sliceCopy, threads)
// 					assert.Equal(t, expected, sliceCopy)
// 				})
// 			}
// 		}
// 	}
// }

// func testSortStable(t *testing.T, getName func() string, sort func(collection []Person)) {
// 	randomTestCasesNo := 10
// 	letters := []byte("abcdefg")

// 	// Slices with length less than 7 never fail
// 	minSliceLength := 7

// 	testCases := [][]Person{
// 		{{"g", 13}, {"f", 11}, {"e", 10}, {"a", 2}, {"d", 8}, {"e", 9}, {"c", 6}, {"d", 7}, {"b", 3}, {"a", 1}, {"g", 14}, {"c", 5}, {"b", 4}, {"f", 12}},
// 		{{"e", 10}, {"g", 14}, {"e", 9}, {"d", 7}, {"b", 3}, {"f", 11}, {"a", 2}, {"f", 12}, {"c", 5}, {"b", 4}, {"g", 13}, {"d", 8}, {"c", 6}, {"a", 1}},
// 		{{"d", 8}, {"f", 12}, {"a", 2}, {"e", 10}, {"d", 7}, {"c", 5}, {"f", 11}, {"b", 3}, {"g", 14}, {"g", 13}, {"b", 4}, {"c", 6}, {"e", 9}, {"a", 1}},
// 	}

// 	for i := 0; i < randomTestCasesNo; i++ {
// 		slice := []Person{}

// 		for i := 0; i < minSliceLength; i++ {
// 			person1 := Person{
// 				name: string(letters[i%len(letters)]),
// 				id:   2 * i,
// 			}
// 			person2 := Person{
// 				name: person1.name,
// 				id:   2*i + 1,
// 			}
// 			slice = append(slice, person1, person2)
// 		}

// 		rand.Shuffle(len(slice), func(i, j int) {
// 			slice[i], slice[j] = slice[j], slice[i]
// 		})

// 		testCases = append(testCases, slice)
// 	}

// 	for _, testCase := range testCases {
// 		t.Run(getName(), func(t *testing.T) {
// 			expected := slices.Clone(testCase)

// 			slices.SortStableFunc(expected, func(a Person, b Person) int {
// 				return cmp.Compare(a.name, b.name)
// 			})

// 			actual := slices.Clone(testCase)

// 			sort(actual)

// 			assert.Equal(t, expected, actual)
// 		})
// 	}
// }

// func TestSortStable(t *testing.T) {
// 	getName := func() string {
// 		return "SortStable"
// 	}

// 	testSortStable(t, getName, func(collection []Person) {
// 		parlo.SortStable(collection, func(a Person, b Person) int {
// 			return cmp.Compare(a.name, b.name)
// 		})
// 	})
// }

// func TestParSortStable(t *testing.T) {
// 	for threads := 1; threads <= MAX_THREADS; threads++ {
// 		getName := func() string {
// 			return fmt.Sprintf("ParSortStable threads=%d", threads)
// 		}

// 		testSortStable(t, getName, func(collection []Person) {
// 			parlo.ParSortStable(collection, func(a Person, b Person) int {
// 				return cmp.Compare(a.name, b.name)
// 			}, threads)
// 		})
// 	}
// }

// func BenchmarkMerge(b *testing.B) {
// 	chunksNo := []int{2, 4, 8, 16, 24, 32, 64}
// 	slices.Sort(chunksNo)

// 	size := 10_000_000

// 	slice := MakeSemiSortedCollection(size)

// 	sliceCopy := make([]Elem, size)

// 	dest := make([]Elem, size)

// 	getSortedChunks := func(chunkNo int) []Elems {
// 		copy(sliceCopy, slice)
// 		sortedChunks := Split[Elems, Elem](sliceCopy, chunkNo)
// 		for _, chunk := range sortedChunks {
// 			slices.Sort(chunk)
// 		}
// 		return sortedChunks
// 	}

// 	for _, chunkNo := range chunksNo {
// 		sortedChunks := getSortedChunks(chunkNo)

// 		mergeFuncs := []struct {
// 			fnName string
// 			fn     func(sortedChunks []Elems, dest Elems)
// 		}{
// 			{"SimpleMerge", parlo.SimpleMerge[Elems, Elem]},
// 			{"MinHeapMerge", parlo.MinHeapMerge[Elems, Elem]},
// 		}

// 		for _, m := range mergeFuncs {
// 			time.Sleep(1 * time.Second)

// 			name := fmt.Sprintf("%s chunks=%d size=%d", m.fnName, chunkNo, size)

// 			b.Run(name, func(b *testing.B) {
// 				for i := 0; i < b.N; i++ {
// 					m.fn(sortedChunks, dest)
// 				}
// 			})
// 		}

// 		fmt.Println()
// 	}
// }

// func BenchmarkMergeBy(b *testing.B) {
// 	chunksNo := []int{2, 4, 8, 16, 24, 32, 64}
// 	slices.Sort(chunksNo)

// 	size := 10_000_000

// 	slice := MakeSemiSortedCollection(size)

// 	sliceCopy := make([]Elem, size)

// 	dest := make([]Elem, size)

// 	getSortedChunks := func(chunkNo int) []Elems {
// 		copy(sliceCopy, slice)
// 		sortedChunks := Split[Elems, Elem](sliceCopy, chunkNo)
// 		for _, chunk := range sortedChunks {
// 			slices.Sort(chunk)
// 		}
// 		return sortedChunks
// 	}

// 	for _, chunkNo := range chunksNo {
// 		sortedChunks := getSortedChunks(chunkNo)

// 		mergeByFuncs := []struct {
// 			fnName string
// 			fn     func(sortedChunks []Elems, dest Elems, cmp func(a, b Elem) int)
// 		}{
// 			// {"SimpleMergeBy", parlo.SimpleMergeBy[Elems, Elem]},
// 			// {"SimpleMergeStableBy", parlo.SimpleMergeStableBy[Elems, Elem]},
// 			{"MinHeapMergeBy", parlo.MinHeapMergeBy[Elems, Elem]},
// 			{"MinHeapMergeStableBy", parlo.MinHeapMergeStableBy[Elems, Elem]},
// 		}

// 		for _, m := range mergeByFuncs {
// 			time.Sleep(3 * time.Second)

// 			name := fmt.Sprintf("%s chunks=%d size=%d", m.fnName, chunkNo, size)

// 			b.Run(name, func(b *testing.B) {
// 				for i := 0; i < b.N; i++ {
// 					m.fn(sortedChunks, dest, func(a, b Elem) int {
// 						return int(a) - int(b)
// 					})
// 				}
// 			})
// 		}

// 		fmt.Println()
// 	}
// }

// func BenchmarkMergePar(b *testing.B) {
// 	threadsNo := []int{2, 4, 8, 16, 24, 32, 64, 128}
// 	slices.Sort(threadsNo)

// 	chunks := 8

// 	size := 10_000_000

// 	slice := MakeSemiSortedCollection(size)

// 	sliceCopy := make(Elems, size)

// 	dest := make(Elems, size)

// 	getSortedChunks := func(chunkNo int) []Elems {
// 		copy(sliceCopy, slice)
// 		sortedChunks := Split[Elems, Elem](sliceCopy, chunkNo)
// 		for _, chunk := range sortedChunks {
// 			slices.Sort(chunk)
// 		}
// 		return sortedChunks
// 	}

// 	sortedChunks := getSortedChunks(chunks)

// 	for _, threads := range threadsNo {
// 		time.Sleep(1 * time.Second)

// 		name := fmt.Sprintf("MergePar threads=%d chunks=%d", threads, chunks)

// 		b.Run(name, func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				parlo.ParMerge1(sortedChunks, dest, threads)
// 			}
// 		})
// 	}
// }

// func BenchmarkMergeParByMerge(b *testing.B) {
// 	threadsNo := []int{2, 4, 8, 16, 24}
// 	slices.Sort(threadsNo)

// 	chunks := 8

// 	size := 10_000_000

// 	slice := MakeSemiSortedCollection(size)

// 	sliceCopy := make(Elems, size)

// 	dest := make(Elems, size)

// 	getSortedChunks := func(chunkNo int) []Elems {
// 		copy(sliceCopy, slice)
// 		sortedChunks := Split[Elems, Elem](sliceCopy, chunkNo)
// 		for _, chunk := range sortedChunks {
// 			slices.Sort(chunk)
// 		}
// 		return sortedChunks
// 	}

// 	sortedChunks := getSortedChunks(chunks)

// 	mergeFuncs := []struct {
// 		fnName string
// 		fn     func(sortedChunks []Elems, dest Elems)
// 	}{
// 		{"SimpleMerge", parlo.SimpleMerge[Elems, Elem]},
// 		{"MinHeapMerge", parlo.MinHeapMerge[Elems, Elem]},
// 	}

// 	for _, threads := range threadsNo {
// 		for _, m := range mergeFuncs {
// 			time.Sleep(1 * time.Second)

// 			name := fmt.Sprintf("MergeParBy%s threads=%d chunks=%d", m.fnName, threads, chunks)

// 			b.Run(name, func(b *testing.B) {
// 				for i := 0; i < b.N; i++ {
// 					parlo.ParMergeByMerge(sortedChunks, dest, m.fn, threads)
// 				}
// 			})
// 		}

// 		fmt.Println()
// 	}
// }

// func BenchmarkReverse(b *testing.B) {
// 	lengths := []int{1_000, 10_000, 100_000, 1_000_000, 100_000_000, 1_000_000_000}
// 	// lengths := []int{1_000_000_000, 100_000_000, 10_000_000, 1_000_000, 100_000, 10_000, 1_000}

// 	for _, length := range lengths {
// 		slice := MakeSemiSortedCollection(length)

// 		expected := slices.Clone(slice)
// 		slices.Sort(expected)

// 		time.Sleep(10 * time.Second)

// 		name := fmt.Sprintf("slices.Reverse len=%d", length)

// 		b.Run(name, func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				b.StopTimer()
// 				sliceCopy := slices.Clone(slice)
// 				b.StartTimer()
// 				slices.Reverse(sliceCopy)
// 				// b.StopTimer()
// 				// assert.Equal(b, expected, sliceCopy)
// 			}
// 		})

// 		time.Sleep(10 * time.Second)

// 		name = fmt.Sprintf("Reverse len=%d", length)

// 		b.Run(name, func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				b.StopTimer()
// 				sliceCopy := slices.Clone(slice)
// 				b.StartTimer()
// 				parlo.Reverse(sliceCopy)
// 				// b.StopTimer()
// 				// assert.Equal(b, expected, sliceCopy)
// 			}
// 		})

// 		fmt.Println()

// 		for threads := 1; threads <= MAX_THREADS; threads++ {
// 			time.Sleep(10 * time.Second)

// 			name := fmt.Sprintf("ParReverse1 threads=%d len=%d", threads, length)

// 			b.Run(name, func(b *testing.B) {
// 				for i := 0; i < b.N; i++ {
// 					b.StopTimer()
// 					sliceCopy := slices.Clone(slice)
// 					b.StartTimer()
// 					parlo.ParReverse(sliceCopy, threads)
// 					// b.StopTimer()
// 					// assert.Equal(b, expected, sliceCopy)
// 				}
// 			})
// 		}

// 		fmt.Println()
// 		fmt.Println()
// 	}
// }

// func BenchmarkSort(b *testing.B) {
// 	// lengths := []int{1_000, 10_000, 100_000, 1_000_000, 100_000_000}
// 	lengths := []int{10_000, 100_000, 1_000_000, 10_000_000, 100_000_000}

// 	sortFns := []struct {
// 		name string
// 		fn   func(collection Elems)
// 	}{
// 		{"slices.Sort", slices.Sort[Elems]},
// 	}

// 	for _, length := range lengths {
// 		slice := MakeSemiSortedCollection(length)

// 		expected := slices.Clone(slice)
// 		slices.Sort(expected)

// 		for _, m := range sortFns {
// 			time.Sleep(1 * time.Second)

// 			name := fmt.Sprintf("%s len=%d", m.name, length)

// 			b.Run(name, func(b *testing.B) {
// 				for i := 0; i < b.N; i++ {
// 					b.StopTimer()
// 					sliceCopy := slices.Clone(slice)
// 					b.StartTimer()
// 					m.fn(sliceCopy)
// 					// b.StopTimer()
// 					// assert.Equal(b, expected, sliceCopy)
// 				}
// 			})
// 		}

// 		fmt.Println()
// 	}
// }

// func BenchmarkStableSort(b *testing.B) {
// 	lengths := []int{100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000}

// 	sortFns := []struct {
// 		name string
// 		fn   func(collection Elems, cmp func(a Elem, b Elem) int)
// 	}{
// 		{"slices.SortStableFunc", slices.SortStableFunc[Elems]},
// 	}

// 	cmp := func(a Elem, b Elem) int {
// 		return int(a) - int(b)
// 	}

// 	for _, length := range lengths {
// 		slice := MakeSemiSortedCollection(length)

// 		sliceCopy := make(Elems, len(slice))

// 		for _, m := range sortFns {
// 			name := fmt.Sprintf("%s len=%d", m.name, length)

// 			b.Run(name, func(b *testing.B) {
// 				for i := 0; i < b.N; i++ {
// 					b.StopTimer()
// 					copy(sliceCopy, slice)
// 					b.StartTimer()
// 					m.fn(sliceCopy, cmp)
// 				}
// 			})
// 		}

// 		fmt.Println()
// 	}
// }

// func BenchmarkSortParByMerge(b *testing.B) {
// 	threadsNo := []int{2, 4, 8, 16, 24, 32, 64}
// 	slices.Sort(threadsNo)

// 	size := 10_000_000

// 	slice := MakeSemiSortedCollection(size)

// 	sliceCopy := make(Elems, len(slice))

// 	mergeFuncs := []struct {
// 		fnName string
// 		fn     func(sortedChunks []Elems, dest Elems)
// 	}{
// 		{"SimpleMerge", parlo.SimpleMerge[Elems, Elem]},
// 		{"MinHeapMerge", parlo.MinHeapMerge[Elems, Elem]},
// 	}

// 	name := fmt.Sprintf("slices.Sort len=%d", size)

// 	b.Run(name, func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			b.StopTimer()
// 			copy(sliceCopy, slice)
// 			b.StartTimer()
// 			slices.Sort(sliceCopy)
// 		}
// 	})

// 	fmt.Println()

// 	for _, threads := range threadsNo {
// 		for _, m := range mergeFuncs {
// 			time.Sleep(1 * time.Second)

// 			name := fmt.Sprintf("SortParBy%s threads=%d len=%d", m.fnName, threads, size)

// 			b.Run(name, func(b *testing.B) {
// 				for i := 0; i < b.N; i++ {
// 					b.StopTimer()
// 					copy(sliceCopy, slice)
// 					b.StartTimer()
// 					parlo.ParSortByMerge1(sliceCopy, m.fn, threads)
// 				}
// 			})
// 		}

// 		fmt.Println()
// 	}
// }
