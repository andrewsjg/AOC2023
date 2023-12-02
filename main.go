package main

import (
	"fmt"

	puzzle1 "github.com/andrewsjg/AOC2023/Puzzle1"
)

func main() {
	// Puzzle 1
	puzzle1.P1Solve("./Puzzle1/input.txt")
	fmt.Println()
	//puzzle1.P2Solve("./Puzzle1/sampleinput2.txt")
	puzzle1.P2Solve("./Puzzle1/input.txt")
}
