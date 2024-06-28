package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

// readData reads data from the given file and returns a slice of float64 values.
func readData(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

// calculateStatistics calculates average, median, variance, and standard deviation.
func calculateStatistics(data []float64) (float64, float64, float64, float64) {
	n := float64(len(data))

	// Calculate average (mean)
	sum := 0.0
	for _, num := range data {
		sum += num
	}
	average := sum / n

	// Calculate median
	sort.Float64s(data)
	var median float64
	if int(n)%2 == 0 {
		median = (data[int(n)/2-1] + data[int(n)/2]) / 2
	} else {
		median = data[int(n)/2]
	}

	// Calculate variance
	var sumSquaredDifferences float64
	for _, num := range data {
		sumSquaredDifferences += (num - average) * (num - average)
	}
	variance := sumSquaredDifferences / n

	// Calculate standard deviation
	stdDev := math.Sqrt(variance)

	return average, median, variance, stdDev
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go data.txt")
		return
	}

	filePath := os.Args[1]
	data, err := readData(filePath)
	if err != nil {
		log.Fatalf("Failed to read data: %v", err)
	}

	average, median, variance, stdDev := calculateStatistics(data)

	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", stdDev)
}
