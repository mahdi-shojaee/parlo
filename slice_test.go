package parlo_test

import (
	"fmt"
	"testing"

	"github.com/mahdi-shojaee/parlo"
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
