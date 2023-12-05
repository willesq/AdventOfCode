package main

import (
	"adventOfCode/cmd/2023/05/internal"
	"fmt"
)

func main() {
	fpPrefix := "./cmd/2023/05"
	inputFilename := "input.txt"

	input, sum1 := internal.Part1(fmt.Sprintf("%s/%s", fpPrefix, inputFilename))
	fmt.Print("2023: Day 05 - Part 1: ", *sum1, "\n")
	sum2 := internal.Part2(input)
	fmt.Print("2023: Day 05 - Part 2: ", *sum2, "\n")
}
