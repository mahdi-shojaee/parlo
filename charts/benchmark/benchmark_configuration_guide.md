# Benchmark Configuration Guide

This guide explains how to define benchmarks and their scenarios in the `benchconfigs.yml` file for the Parlo benchmarking system.

## Overview

The `benchconfigs.yml` file contains configurations for various benchmarks. Each benchmark is defined as a top-level key in the YAML file, with its configuration as a nested object.

## Benchmark Configuration Structure

Each benchmark configuration has the following structure:

### Fields Explanation

1. `BenchmarkName`: The name of the benchmark (e.g., "Min", "Max", "Sort").

2. `sizes`: An array of integers representing the slice sizes to benchmark.
   - Example: `[0, 10000, 100000, 200000, 500000, 1000000, 2000000, 5000000, 10000000]`

3. `baseFuncName`: The name of the base (non-parallel) function to benchmark.
   - Example: `parlo.Min`

4. `funcName`: The name of the parallel function to benchmark.
   - Example: `parlo.ParMin`

5. `immutable`: A boolean indicating whether the function modifies its input slice.
   - `true`: The function does not modify its input (e.g., Min, Max, Filter)
   - `false`: The function modifies its input (e.g., Sort, SortFunc)
   - Example: `immutable: true`

6. `typeArgs` (optional): A string specifying the generic type arguments for functions that require them.
   - Used for functions with multiple generic type parameters
   - Format: Comma-separated list of Go types
   - Example: `typeArgs: "[]int, []int"` for Map functions that convert between slice types

7. `scenarios`: An array of scenario configurations for the benchmark.

### Scenario Configuration

Each scenario in the `scenarios` array has the following fields:

1. `description`: A string describing the scenario.
   - Example: "Input is sorted in ascending order"

2. `elemType`: The type of elements in the slice.
   - Possible values: "int", "float64", or any other Go type

3. `elemAtIndex`: An expression to generate the element at a given index.
   - Examples:
     - `index` (for ascending order)
     - `size - index - 1` (for descending order)
     - `rand.Intn(size)` (for random values)

4. `extraArgs` (optional): Additional arguments passed to the benchmark function.
   - Note: The first argument is always the slice. Use `extraArgs` for any additional arguments.
   - Example: `"func(a, b int) int { return a - b }"`

5. `swapRatio` (optional): A float value between 0 and 1 indicating the ratio of elements to swap randomly.
   - Example: `0.1` (swap 10% of elements)

6. `imports` (optional): An array of additional import statements needed for the scenario.
   - Example: `[math/rand]`

## Example Configuration

Here's an example configuration for a "Min" benchmark:

```yaml
Min:
  sizes: [0, 10000, 100000, 200000, 500000, 1000000, 2000000, 5000000, 10000000]
  baseFuncName: parlo.Min
  funcName: parlo.ParMin
  immutable: true
  typeArgs: "[]int, []int"
  scenarios:
    - description: Input is sorted in ascending order
      elemType: int
      elemAtIndex: index
    - description: Input is sorted in descending order
      elemType: int
      elemAtIndex: size - index - 1
```

## Best Practices

1. Use consistent naming conventions for benchmark and function names.
2. Provide a variety of slice sizes to test performance across different data volumes.
3. Include scenarios that cover different input distributions (e.g., sorted, reverse sorted, random).
4. Use the `swapRatio` field to create partially sorted inputs for sorting benchmarks.
5. Leverage the `imports` field when additional packages are needed for element generation.
6. Remember that the first argument to the benchmark function is always the slice. Use `extraArgs` for any additional arguments required by the function.

## Validation

The `benchconfigs.yml` file is validated against a JSON schema defined in `benchconfigs-schema.json`. Ensure your configurations comply with this schema to avoid errors during benchmark execution.
