package parlo_test

import (
	"math/rand"
	"runtime"
	"sync"

	"github.com/mahdi-shojaee/parlo/internal/constraints"
)

var MAX_THREADS = 8

const LENGTH = 50_000_000

type Elem uint32

type Elems []Elem

func Initialize[T any](slice []T, create func(n int) T) {
	for i := 0; i < len(slice); i++ {
		slice[i] = create(i)
	}
}

func InitializePar[T any](slice []T, create func(n int) T) {
	threads := runtime.GOMAXPROCS(0)
	chunkSize := len(slice) / threads

	s := slice

	var wg sync.WaitGroup
	wg.Add(threads)

	for i := 0; i < threads; i++ {
		endIndex := chunkSize
		if i == threads-1 {
			endIndex = len(s)
		}

		go func(s []T, chunkIndex int) {
			for i := 0; i < len(s); i++ {
				s[i] = create(chunkIndex*chunkSize + i)
			}
			wg.Done()
		}(s[:endIndex], i)

		s = s[endIndex:]
	}

	wg.Wait()
}

func MakeCollection(size int, randomness float32, newElem func(index int) Elem) Elems {
	slice := make(Elems, size)
	InitializePar(slice, newElem)

	numSwaps := int(randomness * float32(size))
	for i := 0; i < numSwaps; i++ {
		j := rand.Intn(size)
		k := rand.Intn(size)
		slice[j], slice[k] = slice[k], slice[j]
	}

	return slice
}

func Split[S ~[]E, E any](slice []E, chunksNo int) []S {
	chunks := make([]S, chunksNo)

	chunkSize := len(slice) / chunksNo

	s := slice

	for i := 0; i < chunksNo; i++ {
		endIndex := chunkSize
		if i == chunksNo-1 {
			endIndex = len(s)
		}

		chunks[i] = s[:endIndex]

		s = s[endIndex:]
	}

	return chunks
}

func Min[S ~[]E, E constraints.Ordered](slice S) E {
	var min E

	if len(slice) == 0 {
		return min
	}

	min = slice[0]

	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

func MinFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) E {
	var min E

	if len(slice) == 0 {
		return min
	}

	min = slice[0]

	for _, v := range slice[1:] {
		if cmp(v, min) < 0 {
			min = v
		}
	}

	return min
}

func Max[S ~[]E, E constraints.Ordered](slice S) E {
	var max E

	if len(slice) == 0 {
		return max
	}

	max = slice[0]

	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

func MaxFunc[S ~[]E, E any](slice S, cmp func(a, b E) int) E {
	var max E

	if len(slice) == 0 {
		return max
	}

	max = slice[0]

	for _, v := range slice[1:] {
		if cmp(v, max) > 0 {
			max = v
		}
	}

	return max
}
