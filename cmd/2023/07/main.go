package main

import (
	"adventOfCode/cmd/2023/07/internal"
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	fpPrefix := "./cmd/2023/07"
	inputFilename := "input.txt"

	input, sum1 := internal.Part1(fmt.Sprintf("%s/%s", fpPrefix, inputFilename))
	fmt.Print("2023: Day 07 - Part 1: ", *sum1, "\n")
	sum2 := internal.Part2(input)
	fmt.Print("2023: Day 07 - Part 2: ", *sum2, "\n")

	fmt.Printf("Total Time Taken: %v\n", time.Since(start))
	fmt.Printf("Difference: %f", math.Abs(float64(249781879-*sum2)))
}

// 249434909 not correct -IDK low or high
// 249646521 not correct -IDK low or high
// 250051537 CORRECT ANSWER
