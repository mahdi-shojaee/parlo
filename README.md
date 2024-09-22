## Parlo: A Performant, Parallel Utility Library for Go

Parlo is a Go library that provides utility functions for efficiently working with slices, maps, and channels.

### Key Advantages:

* **Parallel Processing:** Parlo leverages Go's concurrency features to provide parallel versions of several functions. This allows you to utilize multiple CPU cores and significantly improve performance for large datasets.
* **Generics:** Parlo utilizes Go's generics system (available in Go 1.18 and later) to eliminate the need for reflection in many functions. This translates to better **type safety** and improved performance compared to reflection-based approaches.

### Current Features:

* **Slices:** Sequential and parallel versions of `Find`, `Min`, `Max`, `Filter`, etc.

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

  // Sequential Max
  max := parlo.Max(data)
  fmt.Println("Sequential:", max)

  // Parallel Max
  max = parlo.ParMax(data)
  fmt.Println("Parallel:", max)
}
```
### Parallelism Control

All parallel versions of functions are prefixed with `Par`, indicating they utilize multi-core processing for better performance. For example, `ParMap`, `ParFilter`, and `ParSort` are the parallel counterparts of their sequential versions.

#### Automatic Parallelism

Parlo's parallel functions now automatically manage the degree of parallelism internally. This simplifies the API and ensures optimal performance without requiring manual thread configuration.

- For small datasets (typically less than 200,000 elements), the functions fall back to their sequential counterparts to avoid the overhead of parallelization.
- For larger datasets, the functions utilize parallel processing to improve performance.

### Contributing:

We welcome contributions to Parlo! Feel free to open issues, submit pull requests, or reach out for discussions.

### License:

Parlo is licensed under the [MIT License](https://opensource.org/licenses/MIT).

### Stay Updated:

Follow this repository for updates and new feature announcements.
