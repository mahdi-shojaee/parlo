package parlo_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/mahdi-shojaee/parlo"
	"github.com/mahdi-shojaee/parlo/internal/slices"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	type TestCase struct {
		elems    Elems
		expected Elems
	}

	testCases := []TestCase{
		{Elems{2, 1, 8, 3}, Elems{2, 8}},
		{Elems{1, 2, 3, 4, 5}, Elems{2, 4}},
		{Elems{4, 3, 2, 1, 8, 9}, Elems{4, 2, 8}},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, TestCase{
			elems:    MakeCollection(200_000+100, 0.0, func(index int) Elem { return Elem(index) }),
			expected: MakeCollection((200_000+100)/2, 0.0, func(index int) Elem { return Elem(index * 2) }),
		})
	}

	for _, tc := range testCases {
		t.Run("should return filtered slice", func(t *testing.T) {
			expected := tc.expected
			actual := parlo.Filter(tc.elems, func(item Elem, index int) bool {
				return item%2 == 0
			})
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParFilter(t *testing.T) {
	type TestCase struct {
		elems    Elems
		expected Elems
	}

	testCases := []TestCase{
		{Elems{2, 1, 8, 3}, Elems{2, 8}},
		{Elems{1, 2, 3, 4, 5}, Elems{2, 4}},
		{Elems{4, 3, 2, 1, 8, 9}, Elems{4, 2, 8}},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, TestCase{
			elems:    MakeCollection(200_000+100, 0.0, func(index int) Elem { return Elem(index) }),
			expected: MakeCollection((200_000+100)/2, 0.0, func(index int) Elem { return Elem(index * 2) }),
		})
	}

	for _, tc := range testCases {
		t.Run("should return filtered slice", func(t *testing.T) {
			expected := tc.expected
			actual := parlo.ParFilter(tc.elems, func(item Elem, index int) bool {
				return item%2 == 0
			})
			assert.Equal(t, expected, actual)
		})
	}
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		a        Elems
		b        Elems
		expected bool
	}{
		{Elems{2, 1, 8, 3}, Elems{2, 1, 8, 3}, true},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5}, true},
		{Elems{4, 3, 2, 1, 8, 9}, Elems{4, 3, 2, 1, 8, 9}, true},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, true},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 6}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5, 6}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %t", tc.expected), func(t *testing.T) {
			expected := tc.expected
			actual := parlo.Equal(tc.a, tc.b)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParEqual(t *testing.T) {
	testCases := []struct {
		a        Elems
		b        Elems
		expected bool
	}{
		{Elems{2, 1, 8, 3}, Elems{2, 1, 8, 3}, true},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5}, true},
		{Elems{4, 3, 2, 1, 8, 9}, Elems{4, 3, 2, 1, 8, 9}, true},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, true},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 6}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5, 6}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %t", tc.expected), func(t *testing.T) {
			expected := tc.expected
			actual := parlo.ParEqual(tc.a, tc.b)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestEqualFunc(t *testing.T) {
	testCases := []struct {
		a        Elems
		b        Elems
		expected bool
	}{
		{Elems{2, 1, 8, 3}, Elems{2, 1, 8, 3}, true},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5}, true},
		{Elems{4, 3, 2, 1, 8, 9}, Elems{4, 3, 2, 1, 8, 9}, true},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, true},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 6}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5, 6}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %t", tc.expected), func(t *testing.T) {
			expected := tc.expected
			actual := parlo.EqualFunc(tc.a, tc.b, func(a, b Elem) bool {
				return a == b
			})
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParEqualFunc(t *testing.T) {
	testCases := []struct {
		a        Elems
		b        Elems
		expected bool
	}{
		{Elems{2, 1, 8, 3}, Elems{2, 1, 8, 3}, true},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5}, true},
		{Elems{4, 3, 2, 1, 8, 9}, Elems{4, 3, 2, 1, 8, 9}, true},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, true},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 6}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5, 6}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4}, false},
		{Elems{1, 2, 3, 4, 5}, Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %t", tc.expected), func(t *testing.T) {
			expected := tc.expected
			actual := parlo.ParEqualFunc(tc.a, tc.b, func(a, b Elem) bool {
				return a == b
			})
			assert.Equal(t, expected, actual)
		})
	}
}

func TestIsSorted(t *testing.T) {
	testCases := []struct {
		slice    Elems
		isSorted bool
	}{
		{Elems{2, 1, 8, 3}, false},
		{Elems{1, 2, 3, 4, 5}, true},
		{Elems{4, 3, 2, 1, 8, 9}, false},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %t", tc.isSorted), func(t *testing.T) {
			expected := tc.isSorted
			actual := parlo.IsSorted(tc.slice)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestIsSortedFunc(t *testing.T) {
	testCases := []struct {
		slice    Elems
		isSorted bool
	}{
		{Elems{2, 1, 8, 3}, false},
		{Elems{1, 2, 3, 4, 5}, true},
		{Elems{4, 3, 2, 1, 8, 9}, false},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %t", tc.isSorted), func(t *testing.T) {
			expected := tc.isSorted
			actual := parlo.IsSortedFunc(tc.slice, func(a, b Elem) int {
				return int(a) - int(b)
			})
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParIsSorted(t *testing.T) {
	testCases := []struct {
		slice    Elems
		isSorted bool
	}{
		{Elems{2, 1, 8, 3}, false},
		{Elems{1, 2, 3, 4, 5}, true},
		{Elems{4, 3, 2, 1, 8, 9}, false},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %t", tc.isSorted), func(t *testing.T) {
			expected := tc.isSorted
			actual := parlo.ParIsSorted(tc.slice)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParIsSortedFunc(t *testing.T) {
	testCases := []struct {
		slice    Elems
		isSorted bool
	}{
		{Elems{2, 1, 8, 3}, false},
		{Elems{1, 2, 3, 4, 5}, true},
		{Elems{4, 3, 2, 1, 8, 9}, false},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %t", tc.isSorted), func(t *testing.T) {
			expected := tc.isSorted
			actual := parlo.ParIsSortedFunc(tc.slice, func(a, b Elem) int {
				return int(a) - int(b)
			})
			assert.Equal(t, expected, actual)
		})
	}
}

func TestIsSortedDesc(t *testing.T) {
	testCases := []struct {
		slice    Elems
		isSorted bool
	}{
		{Elems{2, 1, 8, 3}, false},
		{Elems{1, 2, 3, 4, 5}, false},
		{Elems{4, 3, 2, 1, 8, 9}, false},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %t", tc.isSorted), func(t *testing.T) {
			expected := tc.isSorted
			actual := parlo.IsSortedDesc(tc.slice)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParIsSortedDesc(t *testing.T) {
	testCases := []struct {
		slice    Elems
		isSorted bool
	}{
		{Elems{2, 1, 8, 3}, false},
		{Elems{1, 2, 3, 4, 5}, false},
		{Elems{4, 3, 2, 1, 8, 9}, false},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %t", tc.isSorted), func(t *testing.T) {
			expected := tc.isSorted
			actual := parlo.ParIsSortedDesc(tc.slice)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		slice    Elems
		expected Elems
	}{
		{Elems{2, 1, 8, 3}, Elems{3, 8, 1, 2}},
		{Elems{1, 2, 3, 4, 5}, Elems{5, 4, 3, 2, 1}},
		{Elems{4, 3, 2, 1, 8, 9}, Elems{9, 8, 1, 2, 3, 4}},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %v", tc.expected), func(t *testing.T) {
			expected := tc.expected
			actual := tc.slice
			parlo.Reverse(actual)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParReverse(t *testing.T) {
	testCases := []struct {
		slice    Elems
		expected Elems
	}{
		{Elems{2, 1, 8, 3}, Elems{3, 8, 1, 2}},
		{Elems{1, 2, 3, 4, 5}, Elems{5, 4, 3, 2, 1}},
		{Elems{4, 3, 2, 1, 8, 9}, Elems{9, 8, 1, 2, 3, 4}},
		{Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{Elems{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, Elems{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %v", tc.expected), func(t *testing.T) {
			expected := tc.expected
			actual := tc.slice
			parlo.ParReverse(actual)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestSort(t *testing.T) {
	testCases := []Elems{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{1, 2, 3, 7, 4, 0, 6, 5, 8, 9},
		{3, 1, 2, 5, 0, 4, 9, 8, 7, 6},
		{0, 3, 8, 1, 9, 7, 6, 5, 4, 2},
		{9, 1, 2, 3, 4, 5, 6, 11, 8, 0, 10, 7},
		{1, 4, 2, 5, 8, 3, 6, 9, 10},
		{3, 4, 1, 2, 5, 6, 7, 8, 0, 9, 10, 11, 12, 14},
	}

	for i := 0; i < 10; i++ {
		size := rand.Intn(10_000) + 100
		slice := MakeCollection(size, 0.5, func(index int) Elem {
			return Elem(index)
		})
		testCases = append(testCases, slice)
	}

	for _, slice := range testCases {
		sliceCopy := make(Elems, len(slice))

		expected := slices.Clone(slice)
		slices.Sort(expected)

		name := fmt.Sprintf("Sort len=%d", len(slice))
		t.Run(name, func(t *testing.T) {
			copy(sliceCopy, slice)
			parlo.Sort(sliceCopy)
			assert.Equal(t, expected, sliceCopy)
		})
	}
}

func TestSortFunc(t *testing.T) {
	testCases := []Elems{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{1, 2, 3, 7, 4, 0, 6, 5, 8, 9},
		{3, 1, 2, 5, 0, 4, 9, 8, 7, 6},
		{0, 3, 8, 1, 9, 7, 6, 5, 4, 2},
		{9, 1, 2, 3, 4, 5, 6, 11, 8, 0, 10, 7},
		{1, 4, 2, 5, 8, 3, 6, 9, 10},
		{3, 4, 1, 2, 5, 6, 7, 8, 0, 9, 10, 11, 12, 14},
	}

	for i := 0; i < 10; i++ {
		size := rand.Intn(10_000) + 100
		slice := MakeCollection(size, 0.5, func(index int) Elem {
			return Elem(index)
		})
		testCases = append(testCases, slice)
	}

	for _, slice := range testCases {
		sliceCopy := make(Elems, len(slice))

		expected := slices.Clone(slice)
		slices.SortFunc(expected, func(a, b Elem) int {
			return int(a) - int(b)
		})

		name := fmt.Sprintf("SortFunc len=%d", len(slice))
		t.Run(name, func(t *testing.T) {
			copy(sliceCopy, slice)
			parlo.SortFunc(sliceCopy, func(a, b Elem) int {
				return int(a) - int(b)
			})
			assert.Equal(t, expected, sliceCopy)
		})
	}
}

func TestSortStableFunc(t *testing.T) {
	type Person struct {
		Id   int
		Name string
		Age  int
	}

	testCases := [][]Person{
		{
			{Id: 1, Name: "Alice", Age: 30},
			{Id: 2, Name: "Bob", Age: 25},
			{Id: 3, Name: "Charlie", Age: 35},
			{Id: 4, Name: "David", Age: 25},
			{Id: 5, Name: "Eve", Age: 30},
		},
		{
			{Id: 1, Name: "Frank", Age: 40},
			{Id: 2, Name: "Grace", Age: 35},
			{Id: 3, Name: "Henry", Age: 40},
			{Id: 4, Name: "Ivy", Age: 35},
			{Id: 5, Name: "Jack", Age: 40},
		},
	}

	for i := 0; i < 3; i++ {
		size := rand.Intn(1000) + 100
		slice := make([]Person, size)
		for j := 0; j < size; j++ {
			slice[j] = Person{
				Id:   j + 1, // Ensure unique Id within this slice
				Name: fmt.Sprintf("Person%d", j),
				Age:  rand.Intn(50) + 20,
			}
		}
		testCases = append(testCases, slice)
	}

	for _, slice := range testCases {
		t.Run(fmt.Sprintf("SortStableFunc len=%d", len(slice)), func(t *testing.T) {
			sliceCopy := slices.Clone(slice)

			expected := slices.Clone(slice)
			slices.SortStableFunc(expected, func(a, b Person) int {
				return a.Age - b.Age
			})

			parlo.SortStableFunc(sliceCopy, func(a, b Person) int {
				return a.Age - b.Age
			})

			assert.Equal(t, expected, sliceCopy)

			// Check stability: people with the same age should maintain their relative order
			for i := 1; i < len(sliceCopy); i++ {
				if sliceCopy[i].Age == sliceCopy[i-1].Age {
					indexInOriginal := slices.IndexFunc(slice, func(p Person) bool {
						return p.Id == sliceCopy[i].Id
					})
					prevIndexInOriginal := slices.IndexFunc(slice, func(p Person) bool {
						return p.Id == sliceCopy[i-1].Id
					})
					assert.Less(t, prevIndexInOriginal, indexInOriginal,
						"Stability violated for %v and %v", sliceCopy[i-1], sliceCopy[i])
				}
			}
		})
	}
}

func TestParSort(t *testing.T) {
	testCases := []Elems{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{1, 2, 3, 7, 4, 0, 6, 5, 8, 9},
		{3, 1, 2, 5, 0, 4, 9, 8, 7, 6},
		{0, 3, 8, 1, 9, 7, 6, 5, 4, 2},
		{9, 1, 2, 3, 4, 5, 6, 11, 8, 0, 10, 7},
		{1, 4, 2, 5, 8, 3, 6, 9, 10},
		{3, 4, 1, 2, 5, 6, 7, 8, 0, 9, 10, 11, 12, 14},
	}

	for i := 0; i < 10; i++ {
		size := rand.Intn(10_000) + 100
		slice := MakeCollection(size, 0.5, func(index int) Elem {
			return Elem(index)
		})
		testCases = append(testCases, slice)
	}

	for _, slice := range testCases {
		sliceCopy := make(Elems, len(slice))

		expected := slices.Clone(slice)
		slices.Sort(expected)

		name := fmt.Sprintf("ParSort len=%d", len(slice))
		t.Run(name, func(t *testing.T) {
			copy(sliceCopy, slice)
			parlo.ParSort(sliceCopy)
			assert.Equal(t, expected, sliceCopy)
		})
	}
}

func TestParSortFunc(t *testing.T) {
	testCases := []Elems{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{1, 2, 3, 7, 4, 0, 6, 5, 8, 9},
		{3, 1, 2, 5, 0, 4, 9, 8, 7, 6},
		{0, 3, 8, 1, 9, 7, 6, 5, 4, 2},
		{9, 1, 2, 3, 4, 5, 6, 11, 8, 0, 10, 7},
		{1, 4, 2, 5, 8, 3, 6, 9, 10},
		{3, 4, 1, 2, 5, 6, 7, 8, 0, 9, 10, 11, 12, 14},
	}

	for i := 0; i < 10; i++ {
		size := rand.Intn(10_000) + 100
		slice := MakeCollection(size, 0.5, func(index int) Elem {
			return Elem(index)
		})
		testCases = append(testCases, slice)
	}

	for _, slice := range testCases {
		sliceCopy := make(Elems, len(slice))

		expected := slices.Clone(slice)
		slices.Sort(expected)

		name := fmt.Sprintf("ParSortFunc len=%d", len(slice))
		t.Run(name, func(t *testing.T) {
			copy(sliceCopy, slice)
			parlo.ParSortFunc(sliceCopy, func(a, b Elem) int {
				return int(a) - int(b)
			})
			assert.Equal(t, expected, sliceCopy)
		})
	}
}

func TestParSortStableFunc(t *testing.T) {
	type Person struct {
		Id   int
		Name string
		Age  int
	}

	testCases := [][]Person{
		{
			{Id: 1, Name: "Alice", Age: 30},
			{Id: 2, Name: "Bob", Age: 25},
			{Id: 3, Name: "Charlie", Age: 35},
			{Id: 4, Name: "David", Age: 25},
			{Id: 5, Name: "Eve", Age: 30},
		},
		{
			{Id: 1, Name: "Frank", Age: 40},
			{Id: 2, Name: "Grace", Age: 35},
			{Id: 3, Name: "Henry", Age: 40},
			{Id: 4, Name: "Ivy", Age: 35},
			{Id: 5, Name: "Jack", Age: 40},
		},
	}

	for i := 0; i < 3; i++ {
		size := rand.Intn(1000) + 100
		slice := make([]Person, size)
		for j := 0; j < size; j++ {
			slice[j] = Person{
				Id:   j + 1, // Ensure unique Id within this slice
				Name: fmt.Sprintf("Person%d", j),
				Age:  rand.Intn(50) + 20,
			}
		}
		testCases = append(testCases, slice)
	}

	for _, slice := range testCases {
		t.Run(fmt.Sprintf("ParSortStableFunc len=%d", len(slice)), func(t *testing.T) {
			sliceCopy := slices.Clone(slice)

			expected := slices.Clone(slice)
			slices.SortStableFunc(expected, func(a, b Person) int {
				return a.Age - b.Age
			})

			parlo.ParSortStableFunc(sliceCopy, func(a, b Person) int {
				return a.Age - b.Age
			})

			assert.Equal(t, expected, sliceCopy)

			// Check stability: people with the same age should maintain their relative order
			for i := 1; i < len(sliceCopy); i++ {
				if sliceCopy[i].Age == sliceCopy[i-1].Age {
					indexInOriginal := slices.IndexFunc(slice, func(p Person) bool {
						return p.Id == sliceCopy[i].Id
					})
					prevIndexInOriginal := slices.IndexFunc(slice, func(p Person) bool {
						return p.Id == sliceCopy[i-1].Id
					})
					assert.Less(t, prevIndexInOriginal, indexInOriginal,
						"Stability violated for %v and %v", sliceCopy[i-1], sliceCopy[i])
				}
			}
		})
	}
}
