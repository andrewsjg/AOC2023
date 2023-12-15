package main

import (
	"fmt"

	"github.com/andrewsjg/AOC2023/puzzle1"
	"github.com/andrewsjg/AOC2023/puzzle2"
	"github.com/andrewsjg/AOC2023/puzzle3"
	"github.com/andrewsjg/AOC2023/puzzle4"
	"github.com/andrewsjg/AOC2023/puzzle6"
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

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("                     Puzzle 3")
	fmt.Println("======================================================")
	puzzle3.Part1Solve("./Puzzle3/input.txt")
	puzzle3.Part2Solve("./Puzzle3/Input.txt")

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("                     Puzzle 4")
	fmt.Println("======================================================")
	puzzle4.Part1Solve("./Puzzle4/input.txt")
	puzzle4.Part2Solve("./Puzzle4/input.txt")

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("                     Puzzle 5")
	fmt.Println("======================================================")
	fmt.Println("Day 5 solution takes to long to run!")
	//puzzle5.Part1Solve("./Puzzle5/input.txt")
	//puzzle5.Part2Solve("./Puzzle5/input.txt")

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("                     Puzzle 6")
	fmt.Println("======================================================")
	puzzle6.Part1Solve("./puzzle6/input.txt")
	puzzle6.Part2Solve("./puzzle6/input.txt")
}
