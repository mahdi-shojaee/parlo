package parlo_test

import (
	"math/rand"
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
			200_000+rand.Intn(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return min value", func(t *testing.T) {
			expected := Min(tc)
			actual := parlo.Min(tc)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestMinFunc(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.Intn(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return min value", func(t *testing.T) {
			expected := Min(tc)
			actual := parlo.MinFunc(tc, func(a, b Elem) int {
				return int(a) - int(b)
			})
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParMin(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.Intn(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return min value", func(t *testing.T) {
			expected := Min(tc)
			actual := parlo.ParMin(tc)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParMinFunc(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.Intn(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return min value", func(t *testing.T) {
			expected := Min(tc)
			actual := parlo.ParMinFunc(tc, func(a, b Elem) int {
				return int(a) - int(b)
			})
			assert.Equal(t, expected, actual)
		})
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
			200_000+rand.Intn(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return max value", func(t *testing.T) {
			expected := Max(tc)
			actual := parlo.Max(tc)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestMaxFunc(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.Intn(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return max value", func(t *testing.T) {
			expected := Max(tc)
			actual := parlo.MaxFunc(tc, func(a, b Elem) int {
				return int(a) - int(b)
			})
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParMax(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.Intn(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return max value", func(t *testing.T) {
			expected := Max(tc)
			actual := parlo.ParMax(tc)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestParMaxFunc(t *testing.T) {
	testCases := []Elems{
		{2, 1, 8, 3},
		{1, 2, 3, 4, 5},
		{4, 3, 2, 1, 8, 9},
	}

	for i := 0; i < 3; i++ {
		testCases = append(testCases, MakeCollection(
			200_000+rand.Intn(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		t.Run("should return max value", func(t *testing.T) {
			expected := Max(tc)
			actual := parlo.ParMaxFunc(tc, func(a, b Elem) int {
				return int(a) - int(b)
			})
			assert.Equal(t, expected, actual)
		})
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
			200_000+rand.Intn(100),
			0.0,
			func(index int) Elem { return Elem(index) }))
	}

	for _, tc := range testCases {
		exists := tc[rand.Intn(len(tc))]
		notExists := Max(tc) + 1

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
