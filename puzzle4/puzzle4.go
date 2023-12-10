package puzzle4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part1Solve(input string) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	CardNo := 1
	totalCardVal := 0

	for scanner.Scan() {
		line := scanner.Text()

		winningNumbers, cardNumbers := getCardArrays(line)

		cardPointVal := 0
		for _, winningNum := range winningNumbers {
			for _, cardNum := range cardNumbers {
				if cardNum == winningNum {

					if cardPointVal == 0 {
						cardPointVal = 1
					} else {
						cardPointVal = cardPointVal * 2
					}

				}
			}
		}
		totalCardVal = totalCardVal + cardPointVal
		CardNo++
	}

	fmt.Printf("The total point value of all cards is: %d\n", totalCardVal)
}

func Part2Solve(input string) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	gameNo := 0
	deck := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()

		deck[gameNo] += 1

		winningNumbers, cardNumbers := getCardArrays(line)
		winners := []string{}

		for _, winningNum := range winningNumbers {
			if arrayContains(cardNumbers, winningNum) {
				winners = append(winners, winningNum)
			}
		}

		for win := range winners {
			deck[gameNo+win+1] += deck[gameNo]
		}

		gameNo++

	}

	totalCards := 0
	for i := range deck {
		totalCards += deck[i]
	}

	fmt.Printf("Total number of cards: %d\n", totalCards)
}

func arrayContains(input []string, find string) bool {

	for _, item := range input {

		if item == find {
			return true
		}
	}

	return false
}

func getCardArrays(input string) ([]string, []string) {

	winningNums := input[strings.Index(input, ":")+2 : strings.Index(input, "|")]
	cardNums := input[strings.Index(input, "|")+1:]

	tmpArray := strings.Split(winningNums, " ")
	winningNumbers := removeEmptyItems(tmpArray)

	tmpArray = strings.Split(cardNums, " ")
	cardNumbers := removeEmptyItems(tmpArray)

	return winningNumbers, cardNumbers
}

func removeEmptyItems(input []string) []string {
	output := []string{}

	for _, item := range input {
		if item != "" {
			output = append(output, item)
		}
	}

	return output
}
