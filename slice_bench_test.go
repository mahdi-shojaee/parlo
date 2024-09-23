package parlo_test

import (
	"fmt"
	"testing"

	"github.com/mahdi-shojaee/parlo"
)

func BenchmarkIsSortedVsParIsSorted(b *testing.B) {
	sizes := []int{100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		b.Run(fmt.Sprintf("parlo.IsSorted-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.IsSorted(slice)
			}
		})

		b.Run(fmt.Sprintf("parlo.ParIsSorted-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParIsSorted(slice)
			}
		})

		fmt.Println()
	}
}

func BenchmarkIsSortedVsParIsSortedTwoFirstElemsSwapped(b *testing.B) {
	sizes := []int{100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	bigSlice[0], bigSlice[1] = bigSlice[1], bigSlice[0]

	for _, size := range sizes {
		slice := bigSlice[:size]

		b.Run(fmt.Sprintf("parlo.IsSorted-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.IsSorted(slice)
			}
		})

		b.Run(fmt.Sprintf("parlo.ParIsSorted-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParIsSorted(slice)
			}
		})

		fmt.Println()
	}
}

func BenchmarkIsSortedByVsParIsSortedBy(b *testing.B) {
	sizes := []int{100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		b.Run(fmt.Sprintf("parlo.IsSortedBy-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.IsSortedBy(slice, func(a, b Elem) bool {
					return a > b
				})
			}
		})

		b.Run(fmt.Sprintf("parlo.ParIsSortedBy-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParIsSortedBy(slice, func(a, b Elem) bool {
					return a > b
				})
			}
		})

		fmt.Println()
	}
}

func BenchmarkIsSortedByVsParIsSortedByTwoFirstElemsSwapped(b *testing.B) {
	sizes := []int{100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	bigSlice[0], bigSlice[1] = bigSlice[1], bigSlice[0]

	for _, size := range sizes {
		slice := bigSlice[:size]

		b.Run(fmt.Sprintf("parlo.IsSortedBy-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.IsSortedBy(slice, func(a, b Elem) bool {
					return a > b
				})
			}
		})

		b.Run(fmt.Sprintf("parlo.ParIsSortedBy-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParIsSortedBy(slice, func(a, b Elem) bool {
					return a > b
				})
			}
		})

		fmt.Println()
	}
}
