package main

import (
	"fmt"

	"github.com/andrewsjg/AOC2023/puzzle1"
	"github.com/andrewsjg/AOC2023/puzzle2"
)

func main() {

	fmt.Println("======================================================")
	fmt.Println("                     Puzzle 1")
	fmt.Println("======================================================")
	puzzle1.Part1Solve("./Puzzle1/input.txt")
	puzzle1.Part2Solve("./Puzzle1/input.txt")

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("                     Puzzle 2")
	fmt.Println("======================================================")
	puzzle2.Part1Solve("./Puzzle2/input.txt")
	puzzle2.Part2Solve("./Puzzle2/input.txt")
}
