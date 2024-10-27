package benchmark

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"encoding/json"

	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

type BenchmarkConfig struct {
	Name         BenchmarkName `yaml:"name"`
	Sizes        []int         `yaml:"sizes"`
	BaseFuncName string        `yaml:"baseFuncName"`
	FuncName     string        `yaml:"funcName"`
	Scenarios    []Scenario    `yaml:"scenarios"`
}

type Scenario struct {
	Description string   `yaml:"description"`
	Imports     []string `yaml:"imports,omitempty"`
	SwapRatio   float64  `yaml:"swapRatio,omitempty"`
	ElemType    string   `yaml:"elemType"`
	ElemAtIndex string   `yaml:"elemAtIndex"`
	ExtraArgs   string   `yaml:"extraArgs,omitempty"`
}

type TemplateConfig struct {
	Name         BenchmarkName
	Sizes        []int
	BaseFuncName string
	FuncName     string
	Scenario     Scenario
}

type FunctionName string
type BenchmarkName string

type BenchmarkResult struct {
	BenchmarkName   BenchmarkName    `json:"benchmarkName"`
	Sizes           []int            `json:"sizes"`
	ScenarioResults []ScenarioResult `json:"scenarioResults"`
}

type ScenarioResult struct {
	Description string                           `json:"description"`
	Durations   map[FunctionName]map[int][]int64 `json:"durations"`
}

type BenchmarkOutput struct {
	Action  string `json:"Action"`
	Package string `json:"Package"`
	Test    string `json:"Test"`
	Output  string `json:"Output"`
}

type Progress struct {
	Total   int
	Current int
}

var numCores = []int{2, 4, 8}

var benchmarkConfigs map[string]BenchmarkConfig

func GetBenchmarkConfigs() map[string]BenchmarkConfig {
	return benchmarkConfigs
}

func init() {
	// Get the directory of the current file
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	// Construct the path to the config file
	configPath := filepath.Join(currentDir, "benchconfigs.yaml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &benchmarkConfigs)
	if err != nil {
		panic(err)
	}
}

func BenchmarkList() []BenchmarkName {
	var benchmarkNames []BenchmarkName
	for name := range GetBenchmarkConfigs() {
		benchmarkNames = append(benchmarkNames, BenchmarkName(name))
	}
	sort.Slice(benchmarkNames, func(i, j int) bool {
		return benchmarkNames[i] < benchmarkNames[j]
	})
	return benchmarkNames
}

func CalculateNumberOfBenchmarks(benchmarkNames []BenchmarkName) int {
	configs := GetBenchmarkConfigs()
	totalBenchmarks := 0

	for _, benchmarkName := range benchmarkNames {
		totalBenchmarks += len(configs[string(benchmarkName)].Scenarios) * len(numCores)
	}

	return totalBenchmarks
}

func RunBenchmarks(benchmarkNames []BenchmarkName) ([]BenchmarkResult, []error) {
	benchtime := flag.Lookup("benchtime").Value.(flag.Getter).Get().(time.Duration)
	benchmarkResults := make([]BenchmarkResult, 0, len(benchmarkNames))
	var errors []error

	configs := GetBenchmarkConfigs()

	progress := Progress{
		Total:   CalculateNumberOfBenchmarks(benchmarkNames),
		Current: 0,
	}

	for _, benchmarkName := range benchmarkNames {
		fmt.Printf("Running benchmarks for %q...\n", benchmarkName)
		fmt.Println("========================================")
		fmt.Println()

		config, ok := configs[string(benchmarkName)]
		if !ok {
			errors = append(errors, fmt.Errorf("benchmark %q not found", benchmarkName))
			fmt.Println()
			continue
		}

		if config.Name == "" {
			config.Name = BenchmarkName(benchmarkName)
		}

		benchmarkResult, err := RunBenchmark(config, benchtime, &progress)
		if err != nil {
			errors = append(errors, fmt.Errorf("error running benchmark %q: %w", benchmarkName, err))
			fmt.Println()
			continue
		}

		benchmarkResults = append(benchmarkResults, benchmarkResult)

		fmt.Println()
	}

	return benchmarkResults, errors
}

func SaveResults(benchmarkResults []BenchmarkResult) error {
	fmt.Println("Saving results...")
	fmt.Println("========================================")

	for _, br := range benchmarkResults {
		jsonData, err := json.MarshalIndent(br, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON for %q: %w", br.BenchmarkName, err)
		}

		fileName := fmt.Sprintf("benchmark-results-viewer/benchmark-results/%s.json", br.BenchmarkName)
		err = os.WriteFile(fileName, jsonData, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %w", fileName, err)
		}

		fmt.Printf("Results for %q saved successfully.\n", br.BenchmarkName)
	}

	fmt.Println("All results saved successfully.")
	fmt.Println()

	return nil
}

func Build() error {
	fmt.Println("Building benchmark-results.js...")

	// Get all JSON files in the benchmark-results directory
	files, err := filepath.Glob("benchmark-results-viewer/benchmark-results/*.json")
	if err != nil {
		return fmt.Errorf("failed to read benchmark results directory: %w", err)
	}

	// Create a map to store all benchmark results
	allResults := make(map[string]BenchmarkResult)

	// Read and parse each JSON file
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", file, err)
		}

		var result BenchmarkResult
		if err := json.Unmarshal(data, &result); err != nil {
			return fmt.Errorf("failed to parse JSON from file %s: %w", file, err)
		}

		// Use the benchmark name as the key in the allResults map
		allResults[string(result.BenchmarkName)] = result
	}

	// Convert the map to JSON
	jsonData, err := json.Marshal(allResults)
	if err != nil {
		return fmt.Errorf("failed to marshal combined results to JSON: %w", err)
	}

	// Create the JavaScript file content
	jsContent := fmt.Sprintf("const benchmarkResults = %s;", string(jsonData))

	// Write the JavaScript file
	err = os.WriteFile("benchmark-results-viewer/benchmark-results.js", []byte(jsContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write benchmark-results.js: %w", err)
	}

	fmt.Println("benchmark-results.js built successfully.")

	return nil
}

func setGOMAXPROCS(numCore int) []string {
	env := os.Environ()
	gomaxprocsPrefix := "GOMAXPROCS="
	gomaxprocsSet := false
	for i, envVar := range env {
		if strings.HasPrefix(envVar, gomaxprocsPrefix) {
			env[i] = fmt.Sprintf("%s%d", gomaxprocsPrefix, numCore)
			gomaxprocsSet = true
			break
		}
	}
	if !gomaxprocsSet {
		env = append(env, fmt.Sprintf("%s%d", gomaxprocsPrefix, numCore))
	}
	return env
}

func RunBenchmark(config BenchmarkConfig, benchtime time.Duration, progress *Progress) (BenchmarkResult, error) {
	benchmarks := BenchmarkList()
	found := false

	sort.Ints(config.Sizes)

	for _, b := range benchmarks {
		if b == config.Name {
			found = true
			break
		}
	}

	if !found {
		return BenchmarkResult{}, fmt.Errorf("benchmark %q not found", config.Name)
	}

	result := BenchmarkResult{
		BenchmarkName:   config.Name,
		Sizes:           config.Sizes,
		ScenarioResults: []ScenarioResult{},
	}

	for _, scenario := range config.Scenarios {
		templateConfig := TemplateConfig{
			Name:         config.Name,
			Sizes:        config.Sizes,
			BaseFuncName: config.BaseFuncName,
			FuncName:     config.FuncName,
			Scenario:     scenario,
		}
		scenarioResult, err := runScenarioBenchmark(templateConfig, benchtime, progress)
		if err != nil {
			return BenchmarkResult{}, err
		}
		result.ScenarioResults = append(result.ScenarioResults, scenarioResult)
	}

	return result, nil
}

func runScenarioBenchmark(templateConfig TemplateConfig, benchtime time.Duration, progress *Progress) (ScenarioResult, error) {
	// Generate temporary benchmark file
	tmpFile, err := generateBenchmarkFile(templateConfig)
	if err != nil {
		return ScenarioResult{}, fmt.Errorf("failed to generate benchmark file: %v", err)
	}

	// Set up a channel to handle Ctrl+C (SIGINT)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Use a goroutine to handle the signal
	go func() {
		<-c
		fmt.Print("\nInterrupted, cleaning up temporary file...")
		os.Remove(tmpFile.Name())
		fmt.Println("done")
		os.Exit(1)
	}()

	// Ensure cleanup happens even if the function returns normally
	defer func() {
		signal.Stop(c)
		os.Remove(tmpFile.Name())
	}()

	scenarioResult := ScenarioResult{
		Description: templateConfig.Scenario.Description,
		Durations:   make(map[FunctionName]map[int][]int64),
	}

	for _, numCore := range numCores {
		progress.Current++
		fmt.Printf("Running benchmarks for %q (GOMAXPROCS=%d) %d/%d...\n", templateConfig.Name, numCore, progress.Current, progress.Total)
		fmt.Println("----------------------------------------")
		fmt.Println(templateConfig.Scenario.Description)
		fmt.Println("----------------------------------------")

		cmd := exec.Command(
			"go",
			"test",
			"-bench",
			"^Benchmark"+string(templateConfig.Name)+"$",
			"-benchtime",
			benchtime.String(),
			tmpFile.Name(),
		)

		cmd.Env = setGOMAXPROCS(numCore)

		// Create a pipe to capture the command's output
		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			os.Remove(tmpFile.Name())
			return ScenarioResult{}, err
		}

		// Start the command
		if err := cmd.Start(); err != nil {
			os.Remove(tmpFile.Name())
			return ScenarioResult{}, err
		}

		scanner := bufio.NewScanner(stdoutPipe)

		reStr := `Benchmark` + string(templateConfig.Name) + `/(\w+\.\w+)-Len=(\d+)-\d+\s+\d+\s+(\d+(?:\.\d+)?)\s+ns/op`
		re := regexp.MustCompile(reStr)

		for scanner.Scan() {
			line := scanner.Text()

			if strings.HasPrefix(line, "FAIL") {
				color.Red(line)
			} else {
				fmt.Println(line)
			}

			matches := re.FindStringSubmatch(line)

			if len(matches) == 4 {
				functionName := FunctionName(matches[1])
				duration, _ := strconv.ParseFloat(matches[3], 64)

				if _, ok := scenarioResult.Durations[functionName]; !ok {
					scenarioResult.Durations[functionName] = make(map[int][]int64)
				}

				scenarioResult.Durations[functionName][numCore] = append(scenarioResult.Durations[functionName][numCore], int64(duration))
			}
		}

		// Wait for the command to finish
		if err := cmd.Wait(); err != nil {
			os.Remove(tmpFile.Name())
			return ScenarioResult{}, err
		}

		if err := scanner.Err(); err != nil {
			os.Remove(tmpFile.Name())
			return ScenarioResult{}, err
		}

		fmt.Println()
	}

	return scenarioResult, nil
}

func generateBenchmarkFile(config TemplateConfig) (*os.File, error) {
	tmpFile, err := os.CreateTemp("./benchmark", "benchmark_*_test.go")
	if err != nil {
		return nil, err
	}

	// Load the template from a file
	tmpl, err := template.ParseFiles("benchmark/benchmark_template.go.tmpl")
	if err != nil {
		os.Remove(tmpFile.Name())
		return nil, err
	}

	err = tmpl.Execute(tmpFile, config)
	if err != nil {
		os.Remove(tmpFile.Name())
		return nil, err
	}

	return tmpFile, nil
}

func InitializePar[S ~[]E, E any](slice S, newElem func(index, size int) E) {
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

		go func(s []E, chunkIndex int) {
			for i := 0; i < len(s); i++ {
				s[i] = newElem(chunkIndex*chunkSize+i, len(slice))
			}
			wg.Done()
		}(s[:endIndex], i)

		s = s[endIndex:]
	}

	wg.Wait()
}

func MakeCollection[S ~[]E, E any](size int, swapRatio float32, newElem func(index, size int) E) S {
	slice := make(S, size)
	InitializePar(slice, newElem)

	numSwaps := int(swapRatio * float32(size))
	for i := 0; i < numSwaps; i++ {
		j := rand.Intn(size)
		k := rand.Intn(size)
		slice[j], slice[k] = slice[k], slice[j]
	}

	return slice
}
