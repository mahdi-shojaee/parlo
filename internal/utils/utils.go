package utils

import (
	"runtime"
)

func NumThreads(numThreads int) int {
	numCPU := runtime.NumCPU()

	if numThreads <= 0 || numThreads > numCPU {
		return numCPU
	}

	return numThreads
}
