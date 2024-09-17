package parlo_test

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/mahdi-shojaee/parlo"
	"github.com/stretchr/testify/assert"
)

func test(t *testing.T, fnName string, expected, actual func(Elems) Elem) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		tc := MakeSemiSortedCollection(500_000 + rand.IntN(100))
		testCases = append(testCases, tc)
	}

	name := fmt.Sprintf("should return %s value", fnName)

	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			expected := expected(tc)
			actual := actual(tc)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestMin(t *testing.T) {
	test(t, "min", slices.Min[Elems, Elem], func(slice Elems) Elem {
		return parlo.Min(slice)
	})
}

func TestMinBy(t *testing.T) {
	test(t, "min", slices.Min[Elems, Elem], func(slice Elems) Elem {
		return parlo.MinBy(slice, func(a, b Elem) bool {
			return a < b
		})
	})
}

func TestMinPar(t *testing.T) {
	for numThreads := 0; numThreads < MAX_THREADS; numThreads++ {
		test(t, "min", slices.Min[Elems, Elem], func(slice Elems) Elem {
			return parlo.ParMin(slice, numThreads)
		})
	}
}

func TestMinByPar(t *testing.T) {
	for numThreads := 0; numThreads < MAX_THREADS; numThreads++ {
		test(t, "min", slices.Min[Elems, Elem], func(slice Elems) Elem {
			return parlo.ParMinBy(slice, numThreads, func(a, b Elem) bool {
				return a < b
			})
		})
	}
}

func TestMax(t *testing.T) {
	test(t, "max", slices.Max[Elems, Elem], func(slice Elems) Elem {
		return parlo.Max(slice)
	})
}

func TestMaxBy(t *testing.T) {
	test(t, "max", slices.Max[Elems, Elem], func(slice Elems) Elem {
		return parlo.MaxBy(slice, func(a, b Elem) bool {
			return a > b
		})
	})
}

func TestMaxPar(t *testing.T) {
	for numThreads := 0; numThreads < MAX_THREADS; numThreads++ {
		test(t, "max", slices.Max[Elems, Elem], func(slice Elems) Elem {
			return parlo.ParMax(slice, numThreads)
		})
	}
}

func TestMaxByPar(t *testing.T) {
	for numThreads := 0; numThreads < MAX_THREADS; numThreads++ {
		test(t, "max", slices.Max[Elems, Elem], func(slice Elems) Elem {
			return parlo.ParMaxBy(slice, numThreads, func(a, b Elem) bool {
				return a > b
			})
		})
	}
}

func TestFind(t *testing.T) {
	slice := []int{2, 1, 8, 3}

	t.Run("should find correct value", func(t *testing.T) {
		actual, ok := parlo.Find(slice, func(n int) bool { return n == 8 })
		assert.Equal(t, 8, actual)
		assert.Equal(t, true, ok)
	})

	t.Run("should find correct value", func(t *testing.T) {
		actual, ok := parlo.Find(slice, func(n int) bool { return n == 5 })
		assert.Equal(t, 0, actual)
		assert.Equal(t, false, ok)
	})
}

func TestParFind(t *testing.T) {
	slice := []int{2, 1, 8, 3}

	for numThreads := 0; numThreads < MAX_THREADS; numThreads++ {
		t.Run("should find correct value", func(t *testing.T) {
			actual, ok := parlo.ParFind(slice, numThreads, func(n int) bool { return n == 8 })
			assert.Equal(t, 8, actual)
			assert.Equal(t, true, ok)
		})

		t.Run("should find correct value", func(t *testing.T) {
			actual, ok := parlo.ParFind(slice, numThreads, func(n int) bool { return n == 5 })
			assert.Equal(t, 0, actual)
			assert.Equal(t, false, ok)
		})
	}
}
