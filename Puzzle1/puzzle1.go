package puzzle1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func P1Solve(input string) {

	calibrationTotal := 0

	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := 0
		lastDigit := 0

		for _, r := range line {

			firstDigit, err = getDigit(r)
			if err == nil {
				break
			}

		}

		for i := len(line) - 1; i >= 0; i-- {
			lastDigit, err = getDigit(rune(line[i]))

			if err == nil {
				break
			}

		}

		// This is a dirty way to make a 2 digit number
		calVal, err := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))

		if err != nil {
			fmt.Println("couldnt create a value: " + err.Error())
			os.Exit(1)
		}

		calibrationTotal = calibrationTotal + calVal
	}

	fmt.Printf("The calibration total for puzzle 1, part 1 is: %d\n", calibrationTotal)
}

func P2Solve(input string) {
	calibrationTotal := 0

	validDigitStrings := []string{"one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9"}

	digitMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		occurances := make(map[int]string)

		for _, validDigit := range validDigitStrings {
			firstoccurance := strings.Index(line, validDigit)
			lastoccurance := strings.LastIndex(line, validDigit)

			if firstoccurance != -1 {
				occurances[firstoccurance] = validDigit
			}

			if lastoccurance != -1 {
				occurances[lastoccurance] = validDigit
			}
		}

		highest := 0
		lowest := 9999999999

		firstDigit := 0
		lastDigit := 0

		for idx, digit := range occurances {

			if idx <= lowest {
				lowest = idx
				firstDigit = digitMap[digit]
			}

			if idx >= highest {
				highest = idx
				lastDigit = digitMap[digit]
			}

		}

		calVal, err := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))
		if err != nil {
			fmt.Println("couldnt create a value: " + err.Error())
			os.Exit(1)
		}

		calibrationTotal = calibrationTotal + calVal
	}

	fmt.Printf("The calibration total for puzzle 1, part 2 is: %d\n", calibrationTotal)

}

// Check if the passed rune is a digit then return the digit as an int
func getDigit(r rune) (int, error) {

	var err error
	digit := -9

	if unicode.IsDigit(r) {
		digit, err = strconv.Atoi(string(r))
	} else {
		err = errors.New("not a digit")
	}

	return digit, err
}
