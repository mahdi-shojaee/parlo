package parlo_test

import (
	"math/rand/v2"
	"runtime"
	"sync"
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
	threads := runtime.NumCPU()
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

func MakeCollection(size int, newElem func(n int) Elem) Elems {
	list := make(Elems, size)
	InitializePar(list, newElem)
	return list
}

func MakeSemiSortedCollection(size int) Elems {
	slice := MakeCollection(size, func(n int) Elem {
		return Elem(n)
	})

	numSwaps := int(0.1 * float32(size))
	for i := 0; i < numSwaps; i++ {
		j := rand.IntN(size)
		k := rand.IntN(size)
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
