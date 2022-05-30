// https://adventofcode.com/2021/day/1
//
// Reads measurements from an input file and calculates how many measurements are larger than the
// previous measurements.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const EmptyMeasurementSentinelValue = -1

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: main.go FILENAME\n")
		os.Exit(1)
	}
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	previousMeasurement := EmptyMeasurementSentinelValue
	numIncreases := 0

	for lineNumber, line := range strings.Split(string(data), "\n") {
		currentMeasurement, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not convert value to integer on line %d: %s (%v)\n", lineNumber, line, err)
			os.Exit(1)
		}
		if currentMeasurement > previousMeasurement && previousMeasurement != EmptyMeasurementSentinelValue {
			numIncreases += 1
		}
		previousMeasurement = currentMeasurement
	}
	fmt.Printf("Number of increases: %d\n", numIncreases)
}
