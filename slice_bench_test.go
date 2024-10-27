package parlo_test

import (
	"fmt"
	"testing"

	"github.com/mahdi-shojaee/parlo"
	"github.com/mahdi-shojaee/parlo/internal/slices"
)

func BenchmarkFilterVsParFilter(b *testing.B) {
	sizes := []int{10_000, 12_000, 15_000, 20_000, 50_000, 100_000, 500_000, 1_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		fns := []struct {
			name string
			fn   func(Elems, func(Elem, int) bool) Elems
		}{
			{"parlo.Filter", parlo.Filter[Elems]},
			{"parlo.ParFilter", parlo.ParFilter[Elems]},
		}

		for _, f := range fns {
			b.Run(fmt.Sprintf("%s-Len=%d", f.name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.fn(slice, func(item Elem, index int) bool {
						return item%2 == 0
					})
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkIsSortedVsParIsSorted(b *testing.B) {
	sizes := []int{10_000, 50_000, 55_000, 60_000, 100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]
		fns := []struct {
			name string
			fn   func(Elems) bool
		}{
			{"parlo.IsSorted", parlo.IsSorted[Elems]},
			{"parlo.ParIsSorted", parlo.ParIsSorted[Elems]},
		}

		for _, f := range fns {
			b.Run(fmt.Sprintf("%s-Len=%d", f.name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.fn(slice)
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkIsSortedVsParIsSortedTwoFirstElemsSwapped(b *testing.B) {
	sizes := []int{100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	bigSlice[0], bigSlice[1] = bigSlice[1], bigSlice[0]

	for _, size := range sizes {
		slice := bigSlice[:size]

		fns := []struct {
			name string
			fn   func(Elems) bool
		}{
			{"parlo.IsSorted", parlo.IsSorted[Elems]},
			{"parlo.ParIsSorted", parlo.ParIsSorted[Elems]},
		}

		for _, f := range fns {
			b.Run(fmt.Sprintf("%s-Len=%d", f.name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.fn(slice)
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkIsSortedFuncVsParIsSortedFunc(b *testing.B) {
	sizes := []int{10_000, 20_000, 50_000, 100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		fns := []struct {
			name string
			fn   func(Elems, func(Elem, Elem) int) bool
		}{
			{"parlo.IsSortedFunc", parlo.IsSortedFunc[Elems]},
			{"parlo.ParIsSortedFunc", parlo.ParIsSortedFunc[Elems]},
		}

		for _, f := range fns {
			b.Run(fmt.Sprintf("%s-Len=%d", f.name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.fn(slice, func(a, b Elem) int {
						return int(a - b)
					})
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkIsSortedFuncVsParIsSortedFuncTwoFirstElemsSwapped(b *testing.B) {
	sizes := []int{100_000, 500_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 2_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	bigSlice[0], bigSlice[1] = bigSlice[1], bigSlice[0]

	for _, size := range sizes {
		slice := bigSlice[:size]
		fns := []struct {
			name string
			fn   func(Elems, func(Elem, Elem) int) bool
		}{
			{"parlo.IsSortedFunc", parlo.IsSortedFunc[Elems]},
			{"parlo.ParIsSortedFunc", parlo.ParIsSortedFunc[Elems]},
		}

		for _, f := range fns {
			b.Run(fmt.Sprintf("%s-Len=%d", f.name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.fn(slice, func(a, b Elem) int {
						return int(a - b)
					})
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkReverseVsParReverse(b *testing.B) {
	sizes := []int{10_000, 100_000, 500_000, 600_000, 700_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]

		fns := []struct {
			name string
			fn   func(Elems)
		}{
			{"parlo.Reverse", parlo.Reverse[Elems]},
			{"parlo.ParReverse", parlo.ParReverse[Elems]},
		}

		for _, f := range fns {
			b.Run(fmt.Sprintf("%s-Len=%d", f.name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.fn(slice)
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkEqualVsParEqual(b *testing.B) {
	sizes := []int{10_000, 100_000, 200_000, 200_000, 220_000, 240_000, 250_000, 300_000, 500_000, 600_000, 700_000, 1_000_000, 10_000_000}
	bigSlice := MakeCollection(Max(sizes), 0.0, func(index int) Elem { return Elem(index) })

	for _, size := range sizes {
		slice := bigSlice[:size]
		fns := []struct {
			name string
			fn   func(Elems, Elems) bool
		}{
			{"parlo.Equal", parlo.Equal[Elems]},
			{"parlo.ParEqual", parlo.ParEqual[Elems]},
		}

		for _, f := range fns {
			b.Run(fmt.Sprintf("%s-Len=%d", f.name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.fn(slice, slice)
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkSortPseudoRandomInput(b *testing.B) {
	lengths := []int{10_000, 1_000_000, 100_000_000}

	fns := []struct {
		name string
		fn   func(collection Elems)
	}{
		{"Sort", parlo.Sort[Elems]},
		{"ParSort", parlo.ParSort[Elems, Elem]},
	}

	for _, length := range lengths {
		slice := MakeCollection(length, 0.1, func(index int) Elem { return Elem(index) })

		for _, f := range fns {
			name := fmt.Sprintf("%s len=%d", f.name, length)

			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					sliceCopy := slices.Clone(slice)
					b.StartTimer()
					f.fn(sliceCopy)
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkSortFuncPseudoRandomInput(b *testing.B) {
	lengths := []int{10_000, 1_000_000, 100_000_000}

	fns := []struct {
		name string
		fn   func(collection Elems, cmp func(a Elem, b Elem) int)
	}{
		{"SortFunc", parlo.SortFunc[Elems, Elem]},
		{"ParSortFunc", parlo.ParSortFunc[Elems, Elem]},
	}

	for _, length := range lengths {
		slice := MakeCollection(length, 0.1, func(index int) Elem { return Elem(index) })

		for _, f := range fns {
			name := fmt.Sprintf("%s len=%d", f.name, length)

			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					sliceCopy := slices.Clone(slice)
					b.StartTimer()
					f.fn(sliceCopy, func(a Elem, b Elem) int {
						return int(a) - int(b)
					})
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkSortSortedAscending(b *testing.B) {
	lengths := []int{10_000, 1_000_000, 100_000_000}

	fns := []struct {
		name string
		fn   func(collection Elems)
	}{
		{"Sort", parlo.Sort[Elems]},
		{"ParSort", parlo.ParSort[Elems, Elem]},
	}

	for _, length := range lengths {
		slice := MakeCollection(length, 0.0, func(index int) Elem { return Elem(index) })

		for _, f := range fns {
			name := fmt.Sprintf("%s len=%d", f.name, length)

			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					sliceCopy := slices.Clone(slice)
					b.StartTimer()
					f.fn(sliceCopy)
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkSortFuncSortedAscending(b *testing.B) {
	lengths := []int{10_000, 1_000_000, 100_000_000}

	fns := []struct {
		name string
		fn   func(collection Elems, cmp func(a Elem, b Elem) int)
	}{
		{"SortFunc", parlo.SortFunc[Elems, Elem]},
		{"ParSortFunc", parlo.ParSortFunc[Elems, Elem]},
	}

	for _, length := range lengths {
		slice := MakeCollection(length, 0.0, func(index int) Elem { return Elem(index) })

		for _, f := range fns {
			name := fmt.Sprintf("%s len=%d", f.name, length)

			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					sliceCopy := slices.Clone(slice)
					b.StartTimer()
					f.fn(sliceCopy, func(a Elem, b Elem) int {
						return int(a) - int(b)
					})
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkSortSortedDescending(b *testing.B) {
	lengths := []int{10_000, 1_000_000, 100_000_000}

	fns := []struct {
		name string
		fn   func(collection Elems)
	}{
		{"Sort", parlo.Sort[Elems]},
		{"ParSort", parlo.ParSort[Elems, Elem]},
	}

	for _, length := range lengths {
		slice := MakeCollection(length, 0.0, func(index int) Elem { return Elem(length - index - 1) })

		for _, f := range fns {
			name := fmt.Sprintf("%s len=%d", f.name, length)

			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					sliceCopy := slices.Clone(slice)
					b.StartTimer()
					f.fn(sliceCopy)
				}
			})
		}

		fmt.Println()
	}
}

func BenchmarkSortFuncSortedDescending(b *testing.B) {
	lengths := []int{10_000, 1_000_000, 100_000_000}

	fns := []struct {
		name string
		fn   func(collection Elems, cmp func(a Elem, b Elem) int)
	}{
		{"SortFunc", parlo.SortFunc[Elems, Elem]},
		{"ParSortFunc", parlo.ParSortFunc[Elems, Elem]},
	}

	for _, length := range lengths {
		slice := MakeCollection(length, 0.0, func(index int) Elem { return Elem(length - index - 1) })

		for _, f := range fns {
			name := fmt.Sprintf("%s len=%d", f.name, length)

			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					sliceCopy := slices.Clone(slice)
					b.StartTimer()
					f.fn(sliceCopy, func(a Elem, b Elem) int {
						return int(a) - int(b)
					})
				}
			})
		}

		fmt.Println()
	}
}
