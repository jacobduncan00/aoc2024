package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jacobduncan00/aoc2024/solutions"
)

func main() {
	day := flag.Int("day", 0, "day of advent (1-25)")
	flag.Parse()

	if *day < 1 || *day > 25 {
		log.Fatal("Please specify a day between 1 and 25")
	}

	inputFile := filepath.Join("inputs", fmt.Sprintf("day%02d.txt", *day))
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	start := time.Now()
	part1, part2 := solutions.Solve(*day, string(input))
	duration := time.Since(start)

	fmt.Printf("Day %d Solutions:\n", *day)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
	fmt.Printf("Time: %v\n", duration)
}
