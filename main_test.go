package main

import (
	"math"
	"os"
	"testing"
)

func TestReadData(t *testing.T) {
	// Create a temporary file with test data
	file, err := os.CreateTemp("", "testdata")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	// Write test data to the file
	data := "1\n2\n3\n4\n5\n"
	if _, err := file.WriteString(data); err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}

	// Close the file to flush data
	if err := file.Close(); err != nil {
		t.Fatalf("Failed to close file: %v", err)
	}

	// Read the data using the readData function
	readData, err := readData(file.Name())
	if err != nil {
		t.Fatalf("Failed to read data: %v", err)
	}

	// Expected data
	expectedData := []float64{1, 2, 3, 4, 5}

	// Compare the read data with the expected data
	for i, v := range readData {
		if v != expectedData[i] {
			t.Errorf("Expected %v, got %v", expectedData[i], v)
		}
	}
}

func TestCalculateStatistics(t *testing.T) {
	// Test data
	data := []float64{1, 2, 3, 4, 5}

	// Expected results
	expectedAverage := 3.0
	expectedMedian := 3.0
	expectedVariance := 2.0
	expectedStdDev := math.Sqrt(expectedVariance)

	// Calculate statistics using the function
	average, median, variance, stdDev := calculateStatistics(data)

	// Compare the results with the expected values
	if average != expectedAverage {
		t.Errorf("Expected average %.0f, got %.0f", expectedAverage, average)
	}
	if median != expectedMedian {
		t.Errorf("Expected median %.0f, got %.0f", expectedMedian, median)
	}
	if variance != expectedVariance {
		t.Errorf("Expected variance %.0f, got %.0f", expectedVariance, variance)
	}
	if stdDev != expectedStdDev {
		t.Errorf("Expected standard deviation %.0f, got %.0f", expectedStdDev, stdDev)
	}
}
