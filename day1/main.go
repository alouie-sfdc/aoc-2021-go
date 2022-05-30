// https://adventofcode.com/2021/day/1
//
// Reads measurements from an input file and calculates how many measurements are larger than the
// previous measurements.

package main

import (
	"container/list"
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

	part1(data)
	part2(data)
}

func part1(data []byte) {
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
	fmt.Printf("Part 1: Number of increases: %d\n", numIncreases)
}

func part2(data []byte) {
	// The window size we're using is 4 because we're comparing two 3-element windows. That requires
	// 4 elements in total.
	const windowSize = 4
	var window = list.New()
	numIncreases := 0

	for lineNumber, line := range strings.Split(string(data), "\n") {
		currentMeasurement, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not convert value to integer on line %d: %s (%v)\n", lineNumber, line, err)
			os.Exit(1)
		}
		window.PushFront(currentMeasurement)
		if window.Len() < windowSize {
			// We don't have enough data to compare two windows yet.
			continue
		}
		if window.Len() > windowSize {
			window.Remove(window.Back())
		}
		if currentMeasurement > window.Back().Value.(int) {
			// We only have to compare the front of the list with the back of the list, because the two middle
			// values are common to both of the 3-element windows that we're comparing.
			numIncreases++
		}

	}
	fmt.Printf("Part 2: Number of increases: %d\n", numIncreases)
}
