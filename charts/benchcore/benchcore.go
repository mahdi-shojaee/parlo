package benchcore

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type BenchmarkName string

type FunctionName string

type BenchmarkResult struct {
	Title     string                   `json:"title"`
	Sizes     []int                    `json:"sizes"`
	Durations map[FunctionName][]int64 `json:"durations"`
}

type BenchmarkOutput struct {
	Action  string `json:"Action"`
	Package string `json:"Package"`
	Test    string `json:"Test"`
	Output  string `json:"Output"`
}

func BenchmarkList() []BenchmarkName {
	cmd := exec.Command("go", "test", "-bench", ".", "-list", "Benchmark", "./benchcases")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing go test command:", err)
		return nil
	}

	lines := strings.Split(string(output), "\n")
	var benchmarks []BenchmarkName

	for _, line := range lines {
		if strings.HasPrefix(line, "Benchmark") {
			name := strings.TrimPrefix(line, "Benchmark")
			benchmarks = append(benchmarks, BenchmarkName(name))
		}
	}

	return benchmarks
}

func RunBenchmark(benchmarkName BenchmarkName, benchtime time.Duration) (BenchmarkResult, error) {
	// Validate the existence of the benchmarkName
	benchmarks := BenchmarkList()
	found := false

	for _, b := range benchmarks {
		if b == benchmarkName {
			found = true
			break
		}
	}

	if !found {
		return BenchmarkResult{}, fmt.Errorf("benchmark %q not found", benchmarkName)
	}

	cmd := exec.Command(
		"go",
		"test",
		"-bench",
		"^Benchmark"+string(benchmarkName)+"$",
		"-benchtime",
		benchtime.String(),
		"./benchcases",
	)

	result := BenchmarkResult{
		Title:     string(benchmarkName),
		Sizes:     []int{},
		Durations: make(map[FunctionName][]int64),
	}

	// Create a pipe to capture the command's output
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating stdout pipe:", err)
		return BenchmarkResult{}, err
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return BenchmarkResult{}, err
	}

	scanner := bufio.NewScanner(stdoutPipe)

	reStr := `Benchmark` + string(benchmarkName) + `/(\w+\.\w+)-Len=(\d+)-\d+\s+\d+\s+(\d+(?:\.\d+)?)\s+ns/op`
	re := regexp.MustCompile(reStr)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)

		matches := re.FindStringSubmatch(line)

		if len(matches) == 4 {
			functionName := FunctionName(matches[1])
			size, _ := strconv.Atoi(matches[2])
			duration, _ := strconv.ParseFloat(matches[3], 64)

			if !contains(result.Sizes, size) {
				result.Sizes = append(result.Sizes, size)
			}
			result.Durations[functionName] = append(result.Durations[functionName], int64(duration))
		}
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error executing go test command:", err)
		return BenchmarkResult{}, err
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading command output:", err)
		return BenchmarkResult{}, err
	}

	// Sort sizes
	sort.Ints(result.Sizes)

	return result, nil
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func UpdateBenchmarkResult(benchmarkResult BenchmarkResult, benchmarkName string) {
	jsonFilePath := fmt.Sprintf("data/%s.json", benchmarkName)

	jsonData, err := json.MarshalIndent(benchmarkResult, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON for %q: %v\n", benchmarkName, err)
		return
	}

	err = os.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing file for %q: %v\n", benchmarkName, err)
		return
	}

	fmt.Printf("Benchmark result for %q saved to %q\n", benchmarkName, jsonFilePath)
}

func UpdateChart(benchmarkName string) {
	// Run npm update command
	cmd := exec.Command("npm", "run", "update", "--", benchmarkName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running npm update for %q: %v\n", benchmarkName, err)
		fmt.Println("Command output:", string(output))
		return
	}

	fmt.Printf("Successfully updated the benchmark charts for %q functions\n", benchmarkName)
}

func InitializePar[S ~[]E, E any](slice S, newElem func(index int) E) {
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

		go func(s []E, chunkIndex int) {
			for i := 0; i < len(s); i++ {
				s[i] = newElem(chunkIndex*chunkSize + i)
			}
			wg.Done()
		}(s[:endIndex], i)

		s = s[endIndex:]
	}

	wg.Wait()
}

func MakeCollection[S ~[]E, E any](size int, randomness float32, newElem func(index int) E) S {
	slice := make(S, size)
	InitializePar(slice, newElem)

	numSwaps := int(randomness * float32(size))
	for i := 0; i < numSwaps; i++ {
		j := rand.Intn(size)
		k := rand.Intn(size)
		slice[j], slice[k] = slice[k], slice[j]
	}

	return slice
}
