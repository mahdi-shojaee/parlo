package parlo_test

import (
	"testing"

	"github.com/mahdi-shojaee/parlo"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Run("should return filtered slice", func(t *testing.T) {
		elems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected := []int{2, 4, 6, 8, 10}
		actual := parlo.Filter(elems, func(item int, index int) bool {
			return item%2 == 0
		})
		assert.Equal(t, expected, actual)
	})
}

func TestParFilter(t *testing.T) {
	for numThreads := 1; numThreads <= MAX_THREADS; numThreads++ {
		t.Run("should return filtered slice", func(t *testing.T) {
			elems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			expected := []int{2, 4, 6, 8, 10}
			actual := parlo.ParFilter(elems, numThreads, func(item int, index int) bool {
				return item%2 == 0
			})
			assert.Equal(t, expected, actual)
		})
	}
}
