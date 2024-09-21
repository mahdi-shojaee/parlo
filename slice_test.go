package parlo_test

import (
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
	for numThreads := 1; numThreads <= MAX_THREADS; numThreads++ {
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
				actual := parlo.ParFilter(tc.elems, numThreads, func(item Elem, index int) bool {
					return item%2 == 0
				})
				assert.Equal(t, expected, actual)
			})
		}
	}
}
