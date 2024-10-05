package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/mahdi-shojaee/parlo"
	"github.com/mahdi-shojaee/parlo/charts/benchcore"
)

type BenchmarkResultEntry struct {
	benchmarkName   benchcore.BenchmarkName
	benchmarkResult benchcore.BenchmarkResult
}

func main() {
	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [options] [benchmark names...]\n\n", "go run main.go")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s -all\n", "go run main.go")
		fmt.Fprintf(os.Stderr, "  %s -benchmarktime=200ms Min MinFunc\n", "go run main.go")
		fmt.Fprintf(os.Stderr, "\nAvailable benchmarks:\n")
		for _, name := range benchcore.BenchmarkList() {
			fmt.Fprintf(os.Stderr, "  %s\n", string(name))
		}
	}

	// Add a new flag for running all benchmarks
	runAll := flag.Bool("all", false, "Run all available benchmarks")
	// Add the benchtime flag
	benchtime := flag.Duration("benchtime", 1*time.Second, "Benchmark time for each test")
	flag.Parse()

	// Get benchmark names from command-line arguments
	benchmarkNames := flag.Args()

	// If -all flag is set, use all available benchmark names
	if *runAll {
		if len(benchmarkNames) > 0 {
			fmt.Println("Warning: Individual benchmark names are ignored when using the -all flag.")
		}

		for _, name := range benchcore.BenchmarkList() {
			benchmarkNames = append(benchmarkNames, string(name))
		}
	}

	if len(benchmarkNames) == 0 {
		fmt.Println("Error: No benchmarks specified. Use -all to run all benchmarks or provide benchmark names.")
		flag.Usage()
		os.Exit(1)
	}

	// Run npm install command
	cmd := exec.Command("npm", "install")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running npm install: %v\n", err)
		fmt.Println("Command output:", string(output))
		return
	}

	fmt.Printf("Successfully ran npm install\n")

	// Report the number of CPU cores that the runtime uses
	numThreads := parlo.GOMAXPROCS()
	fmt.Printf("Number of CPU cores: %d\n\n", numThreads)

	benchmarkResults := make([]BenchmarkResultEntry, 0, len(benchmarkNames))

	for _, bn := range benchmarkNames {
		benchmarkName := benchcore.BenchmarkName(bn)

		fmt.Printf("Running benchmarks for %q...\n", benchmarkName)
		fmt.Println("----------------------------------------")

		// Pass the benchtime parameter to RunBenchmark
		benchmarkResult, err := benchcore.RunBenchmark(benchmarkName, *benchtime)
		if err != nil {
			c := color.New(color.FgRed)
			c.Printf("Benchmark %q not found\n", benchmarkName)
			fmt.Println()
			continue
		}

		fmt.Println()

		benchmarkResults = append(benchmarkResults, BenchmarkResultEntry{
			benchmarkName:   benchmarkName,
			benchmarkResult: benchmarkResult,
		})
	}

	fmt.Println("Saving the results...")
	fmt.Println("----------------------------------------")

	for _, result := range benchmarkResults {
		benchcore.UpdateBenchmarkResult(result.benchmarkResult, string(result.benchmarkName))
		// benchcore.UpdateChart(string(result.benchmarkName))
		fmt.Println()
	}
}
