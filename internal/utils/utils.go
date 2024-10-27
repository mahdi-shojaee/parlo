package utils

import (
	"runtime"
)

func NumThreads(numThreads int) int {
	numCPU := runtime.GOMAXPROCS(0)

	if numThreads <= 0 || numThreads > numCPU {
		return numCPU
	}

	return numThreads
}
