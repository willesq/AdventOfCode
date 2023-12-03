package main

import (
	"adventOfCode/cmd/2023/03/internal"
	"fmt"
)

func main() {
	fpPrefix := "./cmd/2023/03"
	inputFilename := "input.txt"

	//internal.Part1(fmt.Sprintf("%s/%s", fpPrefix, inputFilename))
	input, sum1 := internal.Part1(fmt.Sprintf("%s/%s", fpPrefix, inputFilename))
	fmt.Print("2023: Day 02 - Part 1: ", *sum1, "\n")
	sum2 := internal.Part2(input)
	fmt.Print("2023: Day 02 - Part 2: ", *sum2, "\n")
}
