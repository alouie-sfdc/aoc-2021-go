// https://adventofcode.com/2021/day/2
//
// Reads commands from an input file and calculates horizontal position and depth.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

func getDirectionAndAmount(line string, lineNumber int) (string, int) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		fmt.Fprintf(os.Stderr, "Error on line %d: each line requires direction and an amount. Contents: '%s'\n", lineNumber, line)
		os.Exit(1)
	}

	direction := parts[0]
	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not convert value to integer on line %d: %s (%v)\n", lineNumber, parts[1], err)
		os.Exit(1)
	}
	return direction, amount
}

func part1(data []byte) {

	entries := make(map[string]int)

	for lineNumber, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		direction, amount := getDirectionAndAmount(line, lineNumber)
		entries[direction] += amount
	}

	horizontal := entries["forward"]
	depth := entries["down"] - entries["up"]

	fmt.Printf("Part 1: entries (debug): %v\n", entries)
	fmt.Printf("Part 1: Horizontal position: %d\n", horizontal)
	fmt.Printf("Part 1: Depth: %d\n", depth)
	fmt.Printf("Part 1: Product: %d\n", horizontal*depth)
}

func part2(data []byte) {

	aim := 0
	horizontal := 0
	depth := 0

	for lineNumber, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		direction, amount := getDirectionAndAmount(line, lineNumber)

		switch direction {
		case "down":
			aim += amount
		case "up":
			aim -= amount
		case "forward":
			horizontal += amount
			depth += aim * amount
		default:
			fmt.Fprintf(os.Stderr, "Unexpected direction on line %d: %s\n", lineNumber, direction)
			os.Exit(1)
		}
	}

	fmt.Printf("Part 2: Horizontal position: %d\n", horizontal)
	fmt.Printf("Part 2: Depth: %d\n", depth)
	fmt.Printf("Part 2: Product: %d\n", horizontal*depth)
}
