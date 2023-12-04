package puzzle2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Feels like I brute forced this a bit! There is probably a more elegant way to do it.
func Part1Solve(input string) {

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	possibleTotal := 0

	for scanner.Scan() {
		possible := true
		line := scanner.Text()

		gameNum, err := strconv.Atoi(line[5:strings.Index(line, ":")])

		if err != nil {
			fmt.Println("couldnt convert to int: " + err.Error())
			os.Exit(1)
		}

		games := line[strings.Index(line, ":")+2:]
		gameList := strings.Split(games, ";")

		for _, turn := range gameList {

			if !possibleTurn(turn, maxRed, maxGreen, maxBlue) {
				possible = false
				break
			}
		}

		if possible {
			possibleTotal = possibleTotal + gameNum
		}

	}
	fmt.Printf("The total of all possible games is: %d\n", possibleTotal)
}

func Part2Solve(input string) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	powerTotal := 0

	for scanner.Scan() {

		line := scanner.Text()

		if err != nil {
			fmt.Println("couldnt convert to int: " + err.Error())
			os.Exit(1)
		}

		gameStr := line[strings.Index(line, ":")+2:]
		gameList := strings.Split(gameStr, ";")
		fewestRed, fewestGreen, fewestBlue := fewestRequred(gameList)

		gamePower := fewestRed * fewestGreen * fewestBlue
		powerTotal = powerTotal + gamePower

	}
	fmt.Printf("The total power of all games is: %d\n", powerTotal)
}

func possibleTurn(turnData string, maxRed int, maxGreen int, maxBlue int) bool {

	possible := true

	regex := regexp.MustCompile(`(?P<number>\d+) (?P<color>\w+)`)
	matches := regex.FindAllStringSubmatch(turnData, -1)

	for _, match := range matches {
		number, err := strconv.Atoi(match[1])

		if err != nil {
			fmt.Println("couldnt convert to int: " + err.Error())
			os.Exit(1)
		}

		colour := match[2]

		if colour == "red" && number > maxRed || colour == "green" && number > maxGreen || colour == "blue" && number > maxBlue {
			possible = false
		}

	}

	return possible
}

func fewestRequred(gameData []string) (red int, green int, blue int) {

	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, turn := range gameData {
		regex := regexp.MustCompile(`(?P<number>\d+) (?P<color>\w+)`)
		matches := regex.FindAllStringSubmatch(turn, -1)

		for _, match := range matches {
			number, err := strconv.Atoi(match[1])

			if err != nil {
				fmt.Println("couldnt convert to int: " + err.Error())
				os.Exit(1)
			}

			colour := match[2]

			switch {
			case colour == "red" && number > maxRed:
				maxRed = number
			case colour == "green" && number > maxGreen:
				maxGreen = number
			case colour == "blue" && number > maxBlue:
				maxBlue = number

			}
		}

	}

	return maxRed, maxGreen, maxBlue
}
