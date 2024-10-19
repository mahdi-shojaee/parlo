package parlo_test

import (
	"fmt"
	"testing"

	"github.com/mahdi-shojaee/parlo"
)

func BenchmarkMinVsParMin(b *testing.B) {
	sizes := []int{0, 1_000, 10_000, 30_000, 50_000, 100_000, 200_000, 500_000, 1_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		b.Run(fmt.Sprintf("parlo.Min-Len=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.Min(slice)
			}
		})

		b.Run(fmt.Sprintf("parlo.ParMin-Len=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParMin(slice)
			}
		})

		fmt.Println()
	}
}

func BenchmarkMinFuncVsParMinFunc(b *testing.B) {
	sizes := []int{5_000, 9_000, 10_000, 100_000, 200_000, 500_000, 1_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		b.Run(fmt.Sprintf("parlo.MinFunc-Len=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.MinFunc(slice, func(a, b Elem) int {
					return int(a) - int(b)
				})
			}
		})

		b.Run(fmt.Sprintf("parlo.ParMinFunc-Len=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParMinFunc(slice, func(a, b Elem) int {
					return int(a) - int(b)
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

		b.Run(fmt.Sprintf("parlo.Max-Len=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.Max(slice)
			}
		})

		b.Run(fmt.Sprintf("parlo.ParMax-Len=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParMax(slice)
			}
		})

		fmt.Println()
	}
}

func BenchmarkMaxFuncVsParMaxFunc(b *testing.B) {
	sizes := []int{5_000, 9_000, 10_000, 100_000, 200_000, 500_000, 1_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		b.Run(fmt.Sprintf("parlo.MaxFunc-Len=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.MaxFunc(slice, func(a, b Elem) int {
					return int(a) - int(b)
				})
			}
		})

		b.Run(fmt.Sprintf("parlo.ParMaxFunc-Len=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parlo.ParMaxFunc(slice, func(a, b Elem) int {
					return int(a) - int(b)
				})
			}
		})

		fmt.Println()
	}
}
