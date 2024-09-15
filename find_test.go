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
	for threads := 1; threads < MAX_THREADS; threads++ {
		test(t, "min", slices.Min[Elems, Elem], func(slice Elems) Elem {
			return parlo.MinPar(slice, threads)
		})
	}
}

func TestMinByPar(t *testing.T) {
	for threads := 1; threads < MAX_THREADS; threads++ {
		test(t, "min", slices.Min[Elems, Elem], func(slice Elems) Elem {
			return parlo.MinByPar(slice, func(a, b Elem) bool {
				return a < b
			}, threads)
		})
	}
}
