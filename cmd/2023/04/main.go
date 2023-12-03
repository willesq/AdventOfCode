package main

import (
	"adventOfCode/cmd/2023/04/internal"
	"fmt"
)

func main() {
	fpPrefix := "./cmd/2023/04"
	inputFilename := "input.txt"

	input, sum1 := internal.Part1(fmt.Sprintf("%s/%s", fpPrefix, inputFilename))
	fmt.Print("2023: Day 04 - Part 1: ", *sum1, "\n")
	sum2 := internal.Part2(input)
	fmt.Print("2023: Day 04 - Part 2: ", *sum2, "\n")
}
