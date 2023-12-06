package main

import (
	"adventOfCode/cmd/2023/06/internal"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fpPrefix := "./cmd/2023/06"
	inputFilename := "input.txt"

	input, sum1 := internal.Part1(fmt.Sprintf("%s/%s", fpPrefix, inputFilename))
	fmt.Print("2023: Day 06 - Part 1: ", *sum1, "\n")
	sum2 := internal.Part2(input)
	fmt.Print("2023: Day 06 - Part 2: ", *sum2, "\n")

	fmt.Printf("Total Time Taken: %v ", time.Since(start))
}
