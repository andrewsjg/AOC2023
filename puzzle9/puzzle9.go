package puzzle9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1Solve(input string) {
	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	part2 := false

	stringInput := string(file)

	scanner := bufio.NewScanner(strings.NewReader(stringInput))

	result := 0
	for scanner.Scan() {
		line := scanner.Text()

		intArray, err := stringToInt(line)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		result = oasisPrediction(intArray, part2) + result

	}

	fmt.Printf("Sum of all histories: %d\n", result)
}

func Part2Solve(input string) {

	part2 := true

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	scanner := bufio.NewScanner(strings.NewReader(stringInput))

	result := 0
	for scanner.Scan() {
		line := scanner.Text()

		intArray, err := stringToInt(line)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		result = oasisPrediction(intArray, part2) + result
	}

	fmt.Printf("Sum of all histories: %d\n", result)
}

func oasisPrediction(input []int, part2 bool) (result int) {

	nextSequence := []int{}
	allZeros := true

	for i := 0; i < len(input); i++ {

		if input[i] != 0 {
			allZeros = false
		}

		if i > 0 {

			diff := input[i] - input[i-1]
			nextSequence = append(nextSequence, diff)
		}
	}

	if !allZeros {
		result = (input[len(input)-1] + oasisPrediction(nextSequence, part2))

		if part2 {
			result = (input[0]) - oasisPrediction(nextSequence, part2)

		}

	}

	return result

}

func stringToInt(input string) (result []int, err error) {

	inputArray := strings.Split(input, " ")

	for i := 0; i < len(inputArray); i++ {
		num, err := strconv.Atoi(string(inputArray[i]))

		if err != nil {
			return []int{}, err
		}

		result = append(result, num)
	}

	return result, nil
}
