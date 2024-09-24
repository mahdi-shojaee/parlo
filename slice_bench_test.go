package parlo_test

import (
	"fmt"
	"testing"

	"github.com/mahdi-shojaee/parlo"
)

func BenchmarkFilterVsParFilter(b *testing.B) {
	sizes := []int{10_000, 12_000, 15_000, 20_000, 50_000, 100_000, 500_000, 1_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		b.Run(fmt.Sprintf("parlo.IsSorted-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.Filter(slice, func(item Elem, index int) bool {
					return item%2 == 0
				})
			}
		})

		b.Run(fmt.Sprintf("parlo.ParIsSorted-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParFilter(slice, func(item Elem, index int) bool {
					return item%2 == 0
				})
			}
		})

		fmt.Println()
	}
}

func BenchmarkIsSortedVsParIsSorted(b *testing.B) {
	sizes := []int{10_000, 50_000, 55_000, 60_000, 100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
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

func BenchmarkIsSortedFuncVsParIsSortedFunc(b *testing.B) {
	sizes := []int{10_000, 20_000, 50_000, 100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		b.Run(fmt.Sprintf("parlo.IsSortedFunc-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.IsSortedFunc(slice, func(a, b Elem) int {
					return int(a - b)
				})
			}
		})

		b.Run(fmt.Sprintf("parlo.ParIsSortedFunc-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParIsSortedFunc(slice, func(a, b Elem) int {
					return int(a - b)
				})
			}
		})

		fmt.Println()
	}
}

func BenchmarkIsSortedFuncVsParIsSortedFuncTwoFirstElemsSwapped(b *testing.B) {
	sizes := []int{100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	bigSlice[0], bigSlice[1] = bigSlice[1], bigSlice[0]

	for _, size := range sizes {
		slice := bigSlice[:size]

		b.Run(fmt.Sprintf("parlo.IsSortedFunc-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.IsSortedFunc(slice, func(a, b Elem) int {
					return int(a - b)
				})
			}
		})

		b.Run(fmt.Sprintf("parlo.ParIsSortedFunc-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParIsSortedFunc(slice, func(a, b Elem) int {
					return int(a - b)
				})
			}
		})

		fmt.Println()
	}
}
