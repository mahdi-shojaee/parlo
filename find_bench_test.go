package parlo_test

import (
	"fmt"
	"testing"

	"github.com/mahdi-shojaee/parlo"
)

func BenchmarkMinVsParMin(b *testing.B) {
	sizes := []int{100_000, 150_000, 180_000, 200_000, 210_000, 220_000, 250_000, 300_000, 500_000, 1_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		// b.Run(fmt.Sprintf("slices.Min-Size%d", size), func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		slices.Min(slice)
		// 	}
		// })

		b.Run(fmt.Sprintf("parlo.Min-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.Min(slice)
			}
		})

		b.Run(fmt.Sprintf("parlo.ParMin-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParMin(slice)
			}
		})

		fmt.Println()
	}
}

func BenchmarkMinByVsParMinBy(b *testing.B) {
	sizes := []int{5_000, 9_000, 10_000, 100_000, 200_000, 500_000, 1_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		// b.Run(fmt.Sprintf("slices.MinFunc-Size%d", size), func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		slices.MinFunc(slice, func(a, b Elem) int {
		// 			return int(a) - int(b)
		// 		})
		// 	}
		// })

		b.Run(fmt.Sprintf("parlo.MinBy-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.MinBy(slice, func(a, b Elem) bool {
					return a < b
				})
			}
		})

		b.Run(fmt.Sprintf("parlo.ParMinBy-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParMinBy(slice, func(a, b Elem) bool {
					return a < b
				})
			}
		})

		fmt.Println()
	}
}

func BenchmarkMaxVsParMax(b *testing.B) {
	sizes := []int{100_000, 130_000, 150_000, 180_000, 200_000, 210_000, 220_000, 250_000, 300_000, 500_000, 1_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		// b.Run(fmt.Sprintf("slices.Max-Size%d", size), func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		slices.Max(slice)
		// 	}
		// })

		b.Run(fmt.Sprintf("parlo.Max-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.Max(slice)
			}
		})

		b.Run(fmt.Sprintf("parlo.ParMax-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParMax(slice)
			}
		})

		fmt.Println()
	}
}

func BenchmarkMaxByVsParMaxBy(b *testing.B) {
	sizes := []int{5_000, 9_000, 10_000, 100_000, 200_000, 500_000, 1_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		// b.Run(fmt.Sprintf("slices.MaxFunc-Size%d", size), func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		slices.MaxFunc(slice, func(a, b Elem) int {
		// 			return int(a) - int(b)
		// 		})
		// 	}
		// })

		b.Run(fmt.Sprintf("parlo.MaxBy-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.MaxBy(slice, func(a, b Elem) bool {
					return a < b
				})
			}
		})

		b.Run(fmt.Sprintf("parlo.ParMaxBy-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParMaxBy(slice, func(a, b Elem) bool {
					return a < b
				})
			}
		})

		fmt.Println()
	}
}

func BenchmarkFindVsParFind(b *testing.B) {
	sizes := []int{100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		b.Run(fmt.Sprintf("parlo.Find-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.Find(slice, func(a Elem) bool {
					return a == Elem(size)
				})
			}
		})

		b.Run(fmt.Sprintf("parlo.ParFind-Size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParFind(slice, func(a Elem) bool {
					return a == Elem(size)
				})
			}
		})

		fmt.Println()
	}
}
