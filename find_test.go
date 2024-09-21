package parlo_test

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/mahdi-shojaee/parlo"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.IntN(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return min value", func(t *testing.T) {
			expected := slices.Min(tc)
			actual := parlo.Min(tc)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestMinBy(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.IntN(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return min value", func(t *testing.T) {
			expected := slices.Min(tc)
			actual := parlo.MinBy(tc, func(a, b Elem) bool {
				return a < b
			})
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParMin(t *testing.T) {
	for numThreads := 0; numThreads < MAX_THREADS; numThreads++ {
		testCases := []Elems{
			{2, 1, 8, 3},
			{1, 2, 3, 4, 5},
			{4, 3, 2, 1, 8, 9},
		}

		for i := 0; i < 3; i++ {
			testCases = append(testCases, MakeCollection(
				200_000+rand.IntN(100),
				0.0,
				func(index int) Elem { return Elem(index) }))
		}

		for _, tc := range testCases {
			t.Run("should return min value", func(t *testing.T) {
				expected := slices.Min(tc)
				actual := parlo.ParMin(tc, numThreads)
				assert.Equal(t, expected, actual)
			})
		}
	}
}

func TestParMinBy(t *testing.T) {
	for numThreads := 0; numThreads < MAX_THREADS; numThreads++ {
		testCases := []Elems{
			{2, 1, 8, 3},
			{1, 2, 3, 4, 5},
			{4, 3, 2, 1, 8, 9},
		}

		for i := 0; i < 3; i++ {
			testCases = append(testCases, MakeCollection(
				200_000+rand.IntN(100),
				0.0,
				func(index int) Elem { return Elem(index) }))
		}

		for _, tc := range testCases {
			t.Run("should return min value", func(t *testing.T) {
				expected := slices.Min(tc)
				actual := parlo.ParMinBy(tc, numThreads, func(a, b Elem) bool {
					return a < b
				})
				assert.Equal(t, expected, actual)
			})
		}
	}
}

func TestMax(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.IntN(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return max value", func(t *testing.T) {
			expected := slices.Max(tc)
			actual := parlo.Max(tc)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestMaxBy(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.IntN(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return max value", func(t *testing.T) {
			expected := slices.Max(tc)
			actual := parlo.MaxBy(tc, func(a, b Elem) bool {
				return a > b
			})
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParMax(t *testing.T) {
	for numThreads := 0; numThreads < MAX_THREADS; numThreads++ {
		testCases := []Elems{
			{2, 1, 8, 3},
			{1, 2, 3, 4, 5},
			{4, 3, 2, 1, 8, 9},
		}

		for i := 0; i < 3; i++ {
			testCases = append(testCases, MakeCollection(
				200_000+rand.IntN(100),
				0.0,
				func(index int) Elem { return Elem(index) }))
		}

		for _, tc := range testCases {
			t.Run("should return max value", func(t *testing.T) {
				expected := slices.Max(tc)
				actual := parlo.ParMax(tc, numThreads)
				assert.Equal(t, expected, actual)
			})
		}
	}
}

func TestParMaxBy(t *testing.T) {
	for numThreads := 0; numThreads < MAX_THREADS; numThreads++ {
		testCases := []Elems{
			{2, 1, 8, 3},
			{1, 2, 3, 4, 5},
			{4, 3, 2, 1, 8, 9},
		}

		for i := 0; i < 3; i++ {
			testCases = append(testCases, MakeCollection(
				200_000+rand.IntN(100),
				0.0,
				func(index int) Elem { return Elem(index) }))
		}

		for _, tc := range testCases {
			t.Run("should return max value", func(t *testing.T) {
				expected := slices.Max(tc)
				actual := parlo.ParMaxBy(tc, numThreads, func(a, b Elem) bool {
					return a > b
				})
				assert.Equal(t, expected, actual)
			})
		}
	}
}

func TestFind(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.IntN(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		exists := tc[rand.IntN(len(tc))]
		notExists := slices.Max(tc) + 1

		t.Run("should find value", func(t *testing.T) {
			actual, ok := parlo.Find(tc, func(item Elem) bool { return item == exists })
			assert.Equal(t, exists, actual)
			assert.Equal(t, true, ok)
		})

		t.Run("should not find value", func(t *testing.T) {
			actual, ok := parlo.Find(tc, func(item Elem) bool { return item == notExists })
			assert.Equal(t, Elem(0), actual)
			assert.Equal(t, false, ok)
		})
	}
}

func TestParFind(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.IntN(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for numThreads := 0; numThreads < MAX_THREADS; numThreads++ {
		for _, tc := range testCases {
			exists := tc[rand.IntN(len(tc))]
			notExists := slices.Max(tc) + 1

			t.Run("should find value", func(t *testing.T) {
				actual, ok := parlo.ParFind(tc, 0, func(item Elem) bool { return item == exists })
				assert.Equal(t, exists, actual)
				assert.Equal(t, true, ok)
			})

			t.Run("should not find value", func(t *testing.T) {
				actual, ok := parlo.ParFind(tc, 0, func(item Elem) bool { return item == notExists })
				assert.Equal(t, Elem(0), actual)
				assert.Equal(t, false, ok)
			})
		}
	}
}
