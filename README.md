## Parlo: A Performant, Parallel Utility Library for Go

Parlo is a Go library that provides utility functions for efficiently working with slices, maps, and channels.

### Key Advantages:

* **Parallel Processing:** Parlo leverages Go's concurrency features to provide parallel versions of several functions. This allows you to utilize multiple CPU cores and significantly improve performance for large datasets.
* **Generics:** Parlo utilizes Go's generics system (available in Go 1.18 and later) to eliminate the need for reflection in many functions. This translates to better **type safety** and improved performance compared to reflection-based approaches.

### Current Features:

* **Slices:** Sequential and parallel versions of `Find`, `Min`, `Max`, etc.

**(Note: The list of features is subject to change based on development progress.)**

### Installation:

```
go get -u github.com/mahdi-shojaee/parlo
```

### Usage:

```Go
package main

import (
  "fmt"

  "github.com/mahdi-shojaee/parlo"
)

func main() {
  data := []int{1, 2, 3, 4, 5}

  // Sequential filtering
  filtered := parlo.Filter(data, func(i int) bool { return i % 2 == 0 })

  // Parallel filtering with 4 CPUs
  parallelFiltered := parlo.ParFilter(data, 4, func(i int) bool { return i % 2 == 0 })

  fmt.Println("Sequential:", filtered)
  fmt.Println("Parallel:", parallelFiltered)
}
```
### Parallelism Control

All parallel versions of functions are prefixed with `Par`, indicating they utilize multi-core processing for better performance. For example, `ParMap`, `ParFilter`, and `ParSort` are the parallel counterparts of their sequential versions.

#### `numThreads` Argument

Each parallel function accepts a `numThreads` argument, which controls the degree of parallelism:

- **`numThreads == 0` or a negative number**: Automatically uses all available CPU cores, determined by calling `runtime.NumCPU()`, for **maximum performance**.
- **`numThreads == 1`**: The function runs in a **separate goroutine**, allowing asynchronous execution without parallelism.
- **`numThreads > 1`**: Manually specify the exact number of threads, offering more granular control over CPU usage.

This provides flexibility depending on the workload and environment:
- By default (with `numThreads == 0`), functions will use all available CPU cores.
- Users can specify `numThreads` based on their performance goals or system constraints.

#### Example Usage:

```go
// Automatically use all available CPU cores
result := parlo.ParMap(data, 0, func(n int) int {
    return n * n
})

// Specify the number of threads (e.g., 4 threads)
result := parlo.ParMap(data, 4, func(n int) int {
    return n * n
})

// Use a single thread in a new goroutine (asynchronous, but not parallel)
result := parlo.ParMap(data, 1, func(n int) int {
    return n * n
})
```

### Contributing:

We welcome contributions to Parlo! Feel free to open issues, submit pull requests, or reach out for discussions.

### License:

Parlo is licensed under the [MIT License](https://opensource.org/licenses/MIT).

### Stay Updated:

Follow this repository for updates and new feature announcements.
