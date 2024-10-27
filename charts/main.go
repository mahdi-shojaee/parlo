package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/mahdi-shojaee/parlo/charts/benchmark"
)

func main() {
	setupFlags()

	benchmarkNames, err := getBenchmarkNames()
	if err != nil {
		color.Red("Error getting benchmark names: %v", err)
		return
	}

	allErrors := runBenchmarks(benchmarkNames)

	if len(allErrors) > 0 {
		color.Red("Errors occurred during benchmarks:")
		for _, err := range allErrors {
			color.Red("- %v", err)
		}
		return
	}

	if err := benchmark.Build(); err != nil {
		color.Red("Error building benchmark-results.js: %v", err)
	}
}

func runBenchmarks(benchmarkNames []benchmark.BenchmarkName) []error {
	var allErrors []error

	benchmarkResults, errors := benchmark.RunBenchmarks(benchmarkNames)
	if len(errors) > 0 {
		allErrors = append(allErrors, errors...)
	}

	if err := benchmark.SaveResults(benchmarkResults); err != nil {
		allErrors = append(allErrors, fmt.Errorf("error saving results: %v", err))
	}

	return allErrors
}

func setupFlags() {
	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [options] [benchmark names...]\n\n", "go run main.go")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s -all\n", "go run main.go")
		fmt.Fprintf(os.Stderr, "  %s -benchmarktime=200ms Min MinFunc\n", "go run main.go")
		fmt.Fprintf(os.Stderr, "  %s -build\n", "go run main.go")
		fmt.Fprintf(os.Stderr, "\nAvailable benchmarks:\n")
		for _, name := range benchmark.BenchmarkList() {
			fmt.Fprintf(os.Stderr, "  %s\n", string(name))
		}
	}

	// Add a new flag for running all benchmarks
	flag.Bool("all", false, "Run all available benchmarks")
	// Add the benchtime flag
	flag.Duration("benchtime", 1*time.Second, "Benchmark time for each test")
	// Add the build flag
	build := flag.Bool("build", false, "Build the benchmarks for the viewer")
	flag.Parse()

	// Check if build flag is set
	if *build {
		benchmark.Build()
		os.Exit(0)
	}
}

func getBenchmarkNames() ([]benchmark.BenchmarkName, error) {
	runAll := flag.Lookup("all").Value.(flag.Getter).Get().(bool)

	benchmarkNames := []benchmark.BenchmarkName{}
	for _, arg := range flag.Args() {
		benchmarkNames = append(benchmarkNames, benchmark.BenchmarkName(arg))
	}

	if runAll {
		if len(benchmarkNames) > 0 {
			color.Yellow("Warning: Individual benchmark names are ignored when using the -all flag.")
		}

		benchmarkNames = benchmark.BenchmarkList()
	}

	if len(benchmarkNames) == 0 {
		return nil, fmt.Errorf("no benchmarks specified. Use -all to run all benchmarks or provide benchmark names")
	}

	return benchmarkNames, nil
}
