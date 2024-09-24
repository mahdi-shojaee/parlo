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

func TestParFind(t *testing.T) {
	type Item struct {
		name  string
		grade int
	}

	type TestCase struct {
		input          []Item
		expected       Item
		notExistsGrade int
	}

	testCases := []TestCase{
		{
			input: []Item{
				{name: "one1", grade: 1}, {name: "two", grade: 2}, {name: "eight", grade: 8},
				{name: "one2", grade: 1},
			},
			expected:       Item{name: "one1", grade: 1},
			notExistsGrade: 100,
		},
		{
			input: []Item{
				{name: "one1", grade: 1}, {name: "two", grade: 2}, {name: "three", grade: 3},
				{name: "four", grade: 4}, {name: "one2", grade: 1},
			},
			expected:       Item{name: "one1", grade: 1},
			notExistsGrade: 100,
		},
		{
			input: []Item{
				{name: "one1", grade: 1}, {name: "three", grade: 3}, {name: "two", grade: 2},
				{name: "seven", grade: 7}, {name: "eight", grade: 8}, {name: "one2", grade: 1},
			},
			expected:       Item{name: "one1", grade: 1},
			notExistsGrade: 100,
		},
		{
			input: []Item{
				{name: "one1", grade: 1}, {name: "nine", grade: 9},
				{name: "three", grade: 3}, {name: "seven1", grade: 7},
				{name: "five", grade: 5}, {name: "two", grade: 2},
				{name: "eight", grade: 8}, {name: "six", grade: 6},
				{name: "four", grade: 4}, {name: "zero", grade: 0},
				{name: "twelve", grade: 12}, {name: "fifteen", grade: 15},
				{name: "seven2", grade: 7}, {name: "fourteen", grade: 14},
				{name: "ten", grade: 10}, {name: "one2", grade: 1},
			},
			expected:       Item{name: "one1", grade: 1},
			notExistsGrade: 100,
		},
	}

	for _, tc := range testCases {
		for i := 0; i < 100; i++ {
			t.Run("should find value", func(t *testing.T) {
				actual, ok := parlo.ParFind(tc.input, func(item Item) bool { return item.grade == tc.expected.grade })
				assert.Equal(t, tc.expected, actual)
				assert.Equal(t, true, ok)
			})
		}

		t.Run("should not find value", func(t *testing.T) {
			actual, ok := parlo.ParFind(tc.input, func(item Item) bool { return item.grade == tc.notExistsGrade })
			assert.Equal(t, Item{}, actual)
			assert.Equal(t, false, ok)
		})
	}
}
